package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var base_url = "https://owners.hyundaiusa.com"

func get_csrf_token() (string, error) {
	// Get CSRF token and validate it
	req, err := http.NewRequest("GET", base_url+"/etc/designs/ownercommon/us/token.json", nil)
	if err != nil {
		log.Println("error generating CSRF token req: ", err)
		return "", err
	}
	// get jwt_token from json response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error calling CSRF token req: ", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading CSRF token: ", err)
		return "", err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	csrf := result["jwt_token"].(string)
	log.Println("csrf: ", csrf)
	// validate CSRF token
	req, err = http.NewRequest("GET", base_url+"/libs/granite/csrf/token.json", nil)
	if err != nil {
		log.Println("error generating csrf_token validation req: ", err)
		return "", err
	}
	req.Header.Add("csrf_token", csrf)
	// check response status
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending csrf validation request: ", err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("error could not validate csrf ", resp.Status)
		return "", err
	}
	return csrf, nil
}

func Login(username, password, pin string) (Auth, error) {

	// get csrf token
	csrf, err := get_csrf_token()
	if err != nil {
		log.Println("Error getting csrf_token: ", err)
		return Auth{}, err
	}

	// get jwt_token from username, password, and csrf token
	req, err := http.NewRequest("POST", base_url+"/bin/common/connectCar", nil)
	if err != nil {
		log.Println("Error getting csrf_token req: ", err)
		return Auth{}, err
	}
	// Add url query params to request
	q := req.URL.Query()
	q.Add(":cq_csrf_token", csrf)
	q.Add("url", base_url+"/us/en/index.html")
	q.Add("username", username)
	q.Add("password", password)
	req.URL.RawQuery = q.Encode()
	// check response status
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error getting csrf_token: ", err)
		return Auth{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Error logging in ", resp.Status)
		return Auth{}, err
	}
	// print response body as json
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading login response: ", err)
		return Auth{}, err
	}
	var login_result map[string]interface{}
	json.Unmarshal([]byte(body), &login_result)
	// print login_result["RESPONSE_STRING"]["jwt_id"]
	// print jwt_id from JSON
	jwt_id := login_result["RESPONSE_STRING"].(map[string]interface{})["jwt_id"].(string)
	log.Println("jwt_id: ", jwt_id)
	// remove first 4 characters from jwt_id if it contains "JWT-" at the beginning
	var jwt_id_truncated string
	jwt_id_truncated = jwt_id
	if strings.HasPrefix(jwt_id, "JWT-") {
		jwt_id_truncated = jwt_id[4:]
	}
	// decode JWT token and get expiration date from "exp" field
	token, err := jwt.Parse(jwt_id_truncated, func(token *jwt.Token) (interface{}, error) {
		// just dummy secret since we can't actually validate the JWT without Huyndai's secret
		// so just using this to parse the token exp date from the token claim
		hmacSampleSecret := []byte("secret")
		return hmacSampleSecret, nil
	})
	if err != nil {
		// we'lll get a signature invalid here but that is expected
		// due to not having a private key to validate the JWT
		log.Println("expected error:", err)
	}
	// print expiration date
	expires_at := int64(token.Claims.(jwt.MapClaims)["exp"].(float64) / 1000)
	log.Println("Raw expiration date: ", expires_at)
	log.Println("Expiration date: ", time.Unix(expires_at, 0))

	var auth Auth = Auth{ // create Auth struct
		Username:   username,
		PIN:        pin,
		JWT_Token:  jwt_id,
		JTW_Expiry: time.Unix(expires_at, 0),
	}

	return auth, nil
}

func setReqHeaders(req *http.Request, auth Auth) {
	// set request headers
	req.Header.Add("CSRF-Token", "undefined")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Referer", "https://owners.hyundaiusa.com/content/myhyundai/us/en/page/dashboard.html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	req.Header.Add("Host", "owners.hyundaiusa.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Origin", "https://owners.hyundaiusa.com")
	req.Header.Add("Cookie", "jwt_token="+auth.JWT_Token+"; s_name="+auth.Username)
}

func GetOwnerInfo(auth Auth) (OwnerInfoService, error) {
	// get owner info such as vehicles model, registration id etc
	req, err := http.NewRequest("POST", base_url+"/bin/common/MyAccountServlet", nil)
	if err != nil {
		log.Println("Error getting owner info req: ", err)
		return OwnerInfoService{}, err
	}
	setReqHeaders(req, auth)
	q := req.URL.Query()
	q.Add("username", auth.Username)
	q.Add("token", auth.JWT_Token)
	q.Add("service", "getOwnerInfoService")
	q.Add("url", "https://owners.hyundaiusa.com/us/en/page/dashboard.html")
	req.URL.RawQuery = q.Encode()
	// check response status
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error getting owner info: ", err)
		return OwnerInfoService{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Error getting owner info ", resp.Status)
		return OwnerInfoService{}, err
	}
	// print response body as json
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading owner info response: ", err)
		return OwnerInfoService{}, err
	}
	var owner_info_result OwnerInfoService
	json.Unmarshal([]byte(body), &owner_info_result)
	return owner_info_result, nil
}

func GetVehicleStatus(auth Auth, vehicle Vehicle, refresh bool) (VehicleStatus, error) {
	// get vehicle status
	req, err := http.NewRequest("POST", base_url+"/bin/common/enrollmentFeature", nil)
	if err != nil {
		log.Println("Error getting vehicle status req: ", err)
		return VehicleStatus{}, err
	}
	setReqHeaders(req, auth)
	q := req.URL.Query()
	q.Add("username", auth.Username)
	q.Add("token", auth.JWT_Token)
	q.Add("services", "getVehicleStatus")
	q.Add("url", "https://owners.hyundaiusa.com/us/en/page/dashboard.html")
	q.Add("vin", vehicle.VinNumber)
	q.Add("regId", vehicle.RegID)
	// by default Huyndai doesn't refresh the data from car unless you set refresh to true
	q.Add("refresh", strconv.FormatBool(refresh))
	req.URL.RawQuery = q.Encode()
	// check response status
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error getting vehicle status: ", err)
		return VehicleStatus{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Error getting vehicle status ", resp.Status)
		return VehicleStatus{}, err
	}
	// print response body as json
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading vehicle status response: ", err)
		return VehicleStatus{}, err
	}
	var vehicle_status_result VehicleStatus
	json.Unmarshal([]byte(body), &vehicle_status_result)
	return vehicle_status_result, nil
}

func StartClimate(auth Auth, vehicle Vehicle, climate ClimateInput) (RemoteActionResult, error) {
	// start climate
	req, err := http.NewRequest("POST", base_url+"/bin/common/remoteAction", nil)
	if err != nil {
		log.Println("Error starting climate req: ", err)
		return RemoteActionResult{}, err
	}
	setReqHeaders(req, auth)
	q := req.URL.Query()
	q.Add("username", auth.Username)
	q.Add("token", auth.JWT_Token)
	q.Add("pin", auth.PIN)
	q.Add("service", "postRemoteFatcStart")
	q.Add("url", "https://owners.hyundaiusa.com/us/en/page/blue-link.html")
	q.Add("regId", vehicle.RegID)
	q.Add("vin", vehicle.VinNumber)
	q.Add("gen", vehicle.Gen)
	q.Add("airCtrl", climate.AirCtrl)
	q.Add("igniOnDuration", climate.IgniOnDuration)
	q.Add("airTempvalue", climate.AirTempvalue)
	// convert climate.Defrost to string
	q.Add("defrost", climate.Defrost)
	q.Add("heating1", climate.Heating1)
	// convert climate.seatHeaterVentInfo to string
	seatHeaterVentInfo_string := fmt.Sprintf(`{"drvSeatHeatState":"%s","astSeatHeatState":"%s"}`, climate.SeatHeaterVentInfo.DrvSeatHeatState, climate.SeatHeaterVentInfo.AstSeatHeatState)
	q.Add("seatHeaterVentInfo", seatHeaterVentInfo_string)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error starting climate: ", err)
		return RemoteActionResult{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading start climate: ", err)
		return RemoteActionResult{}, err
	}
	log.Println("Start climate response: ", string(body))
	// decode body as json
	var start_climate_response RemoteActionResult
	err = json.Unmarshal(body, &start_climate_response)
	if err != nil {
		log.Println("Error decoding start climate response: ", err)
		return RemoteActionResult{}, err
	}
	return start_climate_response, nil

}

func StopClimate(auth Auth, vehicle Vehicle) (RemoteActionResult, error) {
	// stop climate
	req, err := http.NewRequest("POST", base_url+"/bin/common/remoteAction", nil)
	if err != nil {
		log.Println("Error starting climate req: ", err)
		return RemoteActionResult{}, err
	}
	setReqHeaders(req, auth)
	q := req.URL.Query()
	q.Add("username", auth.Username)
	q.Add("token", auth.JWT_Token)
	q.Add("pin", auth.PIN)
	q.Add("service", "postRemoteFatcStop")
	q.Add("url", "https://owners.hyundaiusa.com/us/en/page/blue-link.html")
	q.Add("regId", vehicle.RegID)
	q.Add("vin", vehicle.VinNumber)
	q.Add("gen", vehicle.Gen)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error stopping climate: ", err)
		return RemoteActionResult{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading stop climate: ", err)
		return RemoteActionResult{}, err
	}
	log.Println("stop climate response: ", string(body))
	// decode body as json
	var stop_climate_response RemoteActionResult
	err = json.Unmarshal(body, &stop_climate_response)
	if err != nil {
		log.Println("error decoding stop climate response: ", err)
		return RemoteActionResult{}, err
	}
	return stop_climate_response, nil
}

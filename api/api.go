package api

import (
	"encoding/json"
	"fmt"
	"io"

	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var base_url = "https://api.telematics.hyundaiusa.com"

/*
	CURL code

curl --location 'https://api.telematics.hyundaiusa.com/v2/ac/oauth/token' \
--header 'clientId: m66129Bb-em93-SPAHYN-bZ91-am4540zp19920' \
--header 'clientSecret: v558o935-6nne-423i-baa8' \
--header 'Content-Type: application/json' \

	--data-raw '{
	    "username": "xxx",
	    "password": "xxxx"
	}'
*/
func Login(username, password, pin string, vin string) (Auth, error) {
	// create an http GET request based on the curl code above
	req, err := http.NewRequest("POST", base_url+"/v2/ac/oauth/token", strings.NewReader(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)))
	if err != nil {
		log.Println("Error creating login req: ", err)
		return Auth{}, err
	}
	// set request headers
	req.Header.Add("clientId", "m66129Bb-em93-SPAHYN-bZ91-am4540zp19920")
	req.Header.Add("clientSecret", "v558o935-6nne-423i-baa8")
	req.Header.Add("Content-Type", "application/json")
	// check response status

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error logging in: ", err)
		return Auth{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Error logging in ", resp.Status)
		return Auth{}, err
	}
	// print response body as json
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading login response: ", err)
		return Auth{}, err
	}
	var login_result map[string]interface{}
	json.Unmarshal([]byte(body), &login_result)
	// print access_token from JSON
	access_token := login_result["access_token"].(string)

	// print refresh_token from JSON
	refresh_token := login_result["refresh_token"].(string)

	// print username from JSON
	username = login_result["username"].(string)

	expires_in := login_result["expires_in"].(string)

	// Convert expires_in from string to int
	expires_in_int, err := strconv.Atoi(expires_in)
	if err != nil {
		log.Println("Error converting expires_in to int: ", err)
		return Auth{}, err
	}
	// convert expires_in to time.Duration
	expires_in_duration := time.Duration(expires_in_int) * time.Second
	// get current time
	current_time := time.Now()
	// add expires_in_duration to current_time to get expires_at
	expires_at := current_time.Add(expires_in_duration).Unix()
	// convert expires_at to string

	// get enrollment details and wait until it completes
	var auth Auth = Auth{ // create Auth struct
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		ExpiresAt:    expires_at,
		Username:     username,
		PIN:          pin,
		VIN:          vin,
	}

	enrollment_details, err := GetEnrollmentDetails(auth)
	if err != nil {
		log.Println("Error getting enrollment details: ", err)
		return Auth{}, err
	}
	// find vehicle registration ID based on VIN
	var reg_id string
	for _, vehicle := range enrollment_details.EnrolledVehicleDetails {
		if vehicle.VehicleDetails.Vin == vin {
			reg_id = vehicle.VehicleDetails.Regid
		}
	}

	// fmt.Println("reg_id: ", reg_id)
	auth.RegID = reg_id

	return auth, nil
}

/*
	get enrollment details

curl --location 'https://api.telematics.hyundaiusa.com/ac/v2/enrollment/details/xxxx@email.com' \
--header 'access_token: xxx' \
--header 'User-Agent: okhttp/3.12.0' \
--header 'client_id: m66129Bb-em93-SPAHYN-bZ91-am4540zp19920' \
--header 'includeNonConnectedVehicles: Y' \
--header 'Host: api.telematics.hyundaiusa.com' \
*/
func GetEnrollmentDetails(auth Auth) (EnrollmentDetails, error) {
	// create requests
	req, err := http.NewRequest("GET", base_url+"/ac/v2/enrollment/details/"+auth.Username, nil)
	if err != nil {
		log.Println("Error GetEnrollmentDetails req: ", err)
		return EnrollmentDetails{}, err
	}
	// set request headers
	req.Header.Add("access_token", auth.AccessToken)
	req.Header.Add("User-Agent", "okhttp/3.12.0")
	req.Header.Add("client_id", "m66129Bb-em93-SPAHYN-bZ91-am4540zp19920")
	req.Header.Add("includeNonConnectedVehicles", "Y")
	req.Header.Add("Host", "api.telematics.hyundaiusa.com")

	// fmt.Println("auth.AccessToken: ", auth.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error getting GetEnrollmentDetails req: ", err)
		return EnrollmentDetails{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Error getting GetEnrollmentDetails status ", resp.Status)
		return EnrollmentDetails{}, err
	}
	// print response body as json
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading odometer response: ", err)
		return EnrollmentDetails{}, err
	}
	//marshal to EnrollmentDetails struct
	var enrollment_details_result EnrollmentDetails
	json.Unmarshal([]byte(body), &enrollment_details_result)

	return enrollment_details_result, nil
}

func RefreshToken(auth Auth) (Auth, error) {
	// create requests
	req, err := http.NewRequest("POST", base_url+"/v2/ac/oauth/token/refresh", strings.NewReader(fmt.Sprintf(`{"refresh_token": "%s"}`, auth.RefreshToken)))
	if err != nil {
		log.Println("Error creating refresh token req: ", err)
		return Auth{}, err
	}
	// set request headers
	req.Header.Add("User-Agent", "PostmanRuntime/7.26.10")
	req.Header.Add("client_secret", "v558o935-6nne-423i-baa8")
	req.Header.Add("client_id", "m66129Bb-em93-SPAHYN-bZ91-am4540zp19920")
	req.Header.Add("Content-Type", "application/json")
	// check response status
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error refreshing token: ", err)
		return Auth{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Error refreshing token ", resp.Status)
		return Auth{}, err
	}
	// print response body as json
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading refresh token response: ", err)
		return Auth{}, err
	}
	var refresh_token_result map[string]interface{}
	json.Unmarshal([]byte(body), &refresh_token_result)
	// print access_token from JSON
	access_token := refresh_token_result["access_token"].(string)

	// print refresh_token from JSON
	refresh_token := refresh_token_result["refresh_token"].(string)

	// print username from JSON
	username := refresh_token_result["username"].(string)

	expires_in := refresh_token_result["expires_in"].(string)

	// Convert expires_in from string to int
	expires_in_int, err := strconv.Atoi(expires_in)
	if err != nil {
		log.Println("Error converting expires_in to int: ", err)
		return Auth{}, err
	}
	// convert expires_in to time.Duration
	expires_in_duration := time.Duration(expires_in_int) * time.Second
	// get current time
	current_time := time.Now()
	// add expires_in_duration to current_time to get expires_at
	expires_at := current_time.Add(expires_in_duration).Unix()

	return Auth{ // create Auth struct
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		ExpiresAt:    expires_at,
		Username:     username,
		PIN:          auth.PIN,
		VIN:          auth.VIN,
		RegID:        auth.RegID,
	}, nil

}

func setReqHeaders(req *http.Request, auth Auth) {
	// set request headers
	/*--header 'client_id: m66129Bb-em93-SPAHYN-bZ91-am4540zp19920' \
	  --header 'Host: api.telematics.hyundaiusa.com' \
	  --header 'User-Agent: okhttp/3.12.0' \
	  --header 'registrationId: xxxxxx' \
	  --header 'gen: 2' \
	  --header 'username: xxxx@xxx.com' \
	  --header 'vin: xxxx' \
	  --header 'APPCLOUD-VIN: xxxx' \
	  --header 'Language: 0' \
	  --header 'to: ISS' \
	  --header 'encryptFlag: false' \
	  --header 'from: SPA' \
	  --header 'brandIndicator: H' \
	  --header 'bluelinkservicepin: xxx' \
	  --header 'offset: -4' \
	*/
	req.Header.Add("access_token", auth.AccessToken)
	req.Header.Add("client_id", "m66129Bb-em93-SPAHYN-bZ91-am4540zp19920")
	req.Header.Add("User-Agent", "okhttp/3.12.0")
	req.Header.Add("Host", "api.telematics.hyundaiusa.com")
	req.Header.Add("registrationId", auth.RegID)
	req.Header.Add("VIN", auth.VIN)
	req.Header.Add("APPCLOUD-VIN", auth.VIN)
	req.Header.Add("Language", "0")
	req.Header.Add("to", "ISS")
	req.Header.Add("encryptFlag", "false")
	req.Header.Add("from", "SPA")
	req.Header.Add("brandIndicator", "H")
	req.Header.Add("bluelinkservicepin", auth.PIN)
	req.Header.Add("offset", "-4")
}

/*
curl --location 'https://api.telematics.hyundaiusa.com/ac/v2/rcs/rdo/off' \
--header 'access_token: xxxxx' \
--header 'client_id: xxxxxx' \
--header 'Host: api.telematics.hyundaiusa.com' \
--header 'User-Agent: okhttp/3.12.0' \
--header 'registrationId: xxxxx' \
--header 'gen: 2' \
--header 'username: xxxx' \
--header 'vin: xxxx' \
--header 'APPCLOUD-VIN: xxx' \
--header 'Language: 0' \
--header 'to: ISS' \
--header 'encryptFlag: false' \
--header 'from: SPA' \
--header 'brandIndicator: H' \
--header 'bluelinkservicepin: xxx' \
--header 'offset: -4' \
--header 'Content-Type: application/json' \

	--data-raw '{
		"vin": "xxxx",
		"username": "xxxxx"
	}'

// check status api

curl --location 'https://api.telematics.hyundaiusa.com/ac/v2/rmt/getRunningStatus' \
--header 'access_token: xxx' \
--header 'client_id: xxxx' \
--header 'Host: api.telematics.hyundaiusa.com' \
--header 'User-Agent: okhttp/3.12.0' \
--header 'registrationId: xxxxx' \
--header 'gen: 2' \
--header 'login_id: your_email_address' \
--header 'vin: xxxx' \
--header 'APPCLOUD-VIN: xxx' \
--header 'Language: 0' \
--header 'to: ISS' \
--header 'encryptFlag: false' \
--header 'from: SPA' \
--header 'brandIndicator: H' \
--header 'offset: -4' \
--header 'service_type: REMOTE_LOCK' \
--header 'tid: TID_HEADER_RESPONSE_FROM_LOCK_UNLOCK_COMMAND' \
*/
func DoorLock(auth Auth) error {
	// lock doors
	req, err := http.NewRequest("POST", base_url+"/ac/v2/rcs/rdo/off", strings.NewReader(fmt.Sprintf(`{"vin": "%s", "username": "%s"}`, auth.VIN, auth.Username)))
	if err != nil {
		log.Println("Error locking doors req: ", err)
		return err
	}
	setReqHeaders(req, auth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error locking doors: ", err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading lock doors: ", err)
		return err
	}
	log.Println("Lock doors response: ", string(body))

	// read response header "tmsTid"
	tmsTid := resp.Header.Get("tmsTid")

	// wait 5 seconds before checking getRunningStatus
	time.Sleep(5 * time.Second)

	// check status api
	req, err = http.NewRequest("GET", base_url+"/ac/v2/rmt/getRunningStatus", nil)
	if err != nil {
		log.Println("Error getting getRunningStatus req: ", err)
		return err
	}
	setReqHeaders(req, auth)
	req.Header.Add("service_type", "REMOTE_LOCK")
	req.Header.Add("tid", tmsTid)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error getting getRunningStatus: ", err)
		return err
	}
	defer resp.Body.Close()
	// unmarshal response body to RunningStatusResponse struct

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading getRunningStatus: ", err)
		return err
	}
	var running_status_result RunningStatusResponse
	json.Unmarshal([]byte(body), &running_status_result)

	log.Println("getRunningStatus response: ", running_status_result)
	// check to make sure result.Status == "SUCCESS"
	if running_status_result.Status != "SUCCESS" {
		log.Println("Error getRunningStatus result.Status: ", running_status_result.Status)
		// return the above text as error message
		return fmt.Errorf("error getRunningStatus result.Status: %s", running_status_result.Status)
	}
	return nil

}

/*

curl --location 'https://api.telematics.hyundaiusa.com/ac/v2/evc/fatc/start' \
--header 'access_token: xxxxx' \
--header 'client_id: m66129Bb-em93-SPAHYN-bZ91-am4540zp19920' \
--header 'Host: api.telematics.hyundaiusa.com' \
--header 'User-Agent: okhttp/3.12.0' \
--header 'registrationId: xxx' \
--header 'gen: 2' \
--header 'username: xxxx' \
--header 'vin: xxxx' \
--header 'APPCLOUD-VIN: xxxx' \
--header 'Language: 0' \
--header 'to: ISS' \
--header 'encryptFlag: false' \
--header 'from: SPA' \
--header 'brandIndicator: H' \
--header 'bluelinkservicepin: xxxx' \
--header 'offset: -4' \
--header 'Content-Type: application/json' \
--data '{"airCtrl":1,"airTemp":{"value":"72","unit":1},"heating1":0,"defrost":false}'

*/

func StartClimate(auth Auth, temp int) error {
	// start climate
	req, err := http.NewRequest("POST", base_url+"/ac/v2/evc/fatc/start", strings.NewReader(fmt.Sprintf(`{"airCtrl":1,"airTemp":{"value":"%d","unit":1},"heating1":0,"defrost":false}`, temp)))
	if err != nil {
		log.Println("Error starting climate req: ", err)
		return err
	}
	setReqHeaders(req, auth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error starting climate: ", err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading start climate: ", err)
		return err
	}
	log.Println("Start climate response: ", string(body))

	return nil

}

/*
curl --location --request POST 'https://api.telematics.hyundaiusa.com/ac/v2/evc/fatc/stop' \
--header 'access_token: xxxx' \
--header 'client_id: m66129Bb-em93-SPAHYN-bZ91-am4540zp19920' \
--header 'Host: api.telematics.hyundaiusa.com' \
--header 'User-Agent: okhttp/3.12.0' \
--header 'registrationId: x' \
--header 'gen: 2' \
--header 'username: xxxx' \
--header 'vin: xxx' \
--header 'APPCLOUD-VIN: xxx' \
--header 'Language: 0' \
--header 'to: ISS' \
--header 'encryptFlag: false' \
--header 'from: SPA' \
--header 'brandIndicator: H' \
--header 'bluelinkservicepin: xxx' \
--header 'offset: -4' \
--data ''
*/

func StopClimate(auth Auth) error {
	// stop climate
	req, err := http.NewRequest("POST", base_url+"/ac/v2/evc/fatc/stop", nil)
	if err != nil {
		log.Println("Error starting climate req: ", err)
		return err
	}
	setReqHeaders(req, auth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error stopping climate: ", err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading stop climate: ", err)
		return err
	}
	log.Println("stop climate response: ", string(body))

	return nil

}

func GetVehicleStatus(auth Auth, refresh string) (VehicleStatusResponse, error) {
	// stop climate
	req, err := http.NewRequest("GET", base_url+"/ac/v2/rcs/rvs/vehicleStatus", nil)
	if err != nil {
		log.Println("Error getting vehicleStatus req: ", err)
		return VehicleStatusResponse{}, err
	}
	setReqHeaders(req, auth)
	req.Header.Add("refresh", refresh)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error getting vehicleStatus: ", err)
		return VehicleStatusResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading vehicleStatus: ", err)
		return VehicleStatusResponse{}, err
	}
	// marshal to VehicleStatusResponse struct
	var vehicle_status_result VehicleStatusResponse
	json.Unmarshal([]byte(body), &vehicle_status_result)

	return vehicle_status_result, nil
}

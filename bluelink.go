package bluelink_go

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/TaiPhamD/bluelink_go/api"
)

func Login(username, password, pin string, vin string) (api.Auth, error) {
	my_auth, err := api.Login(username, password, pin, vin)
	if err != nil {
		return api.Auth{}, err
	}
	return my_auth, nil
}

func RefreshToken(auth api.Auth) (api.Auth, error) {
	my_auth, err := api.RefreshToken(auth)
	if err != nil {
		return api.Auth{}, err
	}
	return my_auth, nil
}

func GetOdometer(auth api.Auth) (string, error) {
	details, err := api.GetEnrollmentDetails(auth)
	if err != nil {
		return "", err
	}
	// loop through details until we find matching vehicles vin
	for _, vehicle := range details.EnrolledVehicleDetails {
		if vehicle.VehicleDetails.Vin == auth.VIN {
			updateDate := vehicle.VehicleDetails.OdometerUpdateDate
			format := "20060102150405" // Go uses this unique layout format based on the reference time "Mon Jan 2 15:04:05 -0700 MST 2006"

			// 1. Parse the string into a time.Time object
			t, err := time.Parse(format, updateDate)
			if err == nil {
				//fmt.Println("Error:", err)
				duration := time.Since(t)
				// 3. Extract the number of hours
				hours := duration.Hours()
				s := fmt.Sprintf(" miles last updated %d hours ago", int(hours))
				return vehicle.VehicleDetails.Odometer + s, nil
			} else {
				return vehicle.VehicleDetails.Odometer + " miles", nil
			}

		}
	}
	// return error vehicle not found
	return "", errors.New("vehicle not found")
}

func GetVehicleStatus(auth api.Auth) (api.VehicleStatusResponse, error) {
	status, err := api.GetVehicleStatus(auth, "false")
	if err != nil {
		return api.VehicleStatusResponse{}, err
	}

	// get current time
	currentTime := time.Now()
	// get last updated time
	lastUpdated := status.VehicleStatus.DateTime

	// if lastUpdated is more than 10 minutes ago then call api again but with "true" to refresh it
	duration := currentTime.Sub(lastUpdated)
	if duration.Minutes() > 15 {
		log.Println("Vehicle status is more than 15 minutes old, refreshing...")
		status, err = api.GetVehicleStatus(auth, "true")
		if err != nil {
			return api.VehicleStatusResponse{}, err
		}
	}

	return status, nil
}

func DoorLock(auth api.Auth) error {
	err := api.DoorLock(auth)
	if err != nil {
		return err
	}
	return nil
}

func StartClimate(auth api.Auth, temp int) error {
	err := api.StartClimate(auth, temp)
	if err != nil {
		return err
	}
	return nil
}

func StopClimate(auth api.Auth) error {
	err := api.StopClimate(auth)
	if err != nil {
		return err
	}
	return nil
}

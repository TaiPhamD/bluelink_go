package bluelink_go

import (
	"errors"

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
			return vehicle.VehicleDetails.Odometer, nil
		}
	}
	// return error vehicle not found
	return "", errors.New("vehicle not found")
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

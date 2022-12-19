package bluelink_go

import (
	"errors"

	"github.com/TaiPhamD/bluelink_go/api"
)

func Login(username, password, pin string) (api.Auth, error) {
	my_auth, err := api.Login(username, password, pin)
	if err != nil {
		return api.Auth{}, err
	}
	return my_auth, nil
}

func GetOwnerInfo(auth api.Auth) (api.OwnerInfoService, error) {
	owner_info, err := api.GetOwnerInfo(auth)
	if err != nil {
		return api.OwnerInfoService{}, err
	}
	return owner_info, nil
}

func GetVehicleFromVin(owner_info api.OwnerInfoService, vin string) (api.Vehicle, error) {
	vehicles := owner_info.ResponseString.OwnersVehiclesInfo
	for _, vehicle := range vehicles {
		if vehicle.VinNumber == vin {
			var mycar api.Vehicle
			mycar.VinNumber = vehicle.VinNumber
			mycar.RegID = vehicle.RegistrationID
			// the property IsGen2 is actually a string number "2" for gen 2 from Huyndai
			// even though the name implies a boolean
			mycar.Gen = vehicle.IsGen2

			return mycar, nil
		}
	}
	// return error vehicle not found
	return api.Vehicle{}, errors.New("vehicle not found")

}

func GetVehicleStatus(auth api.Auth, vehicle api.Vehicle, refresh bool) (api.VehicleStatus, error) {
	// calling **GetVehicleStatus** API. Set 3rd param to true to force refresh
	vehicle_status, err := api.GetVehicleStatus(auth, vehicle, refresh)
	if err != nil {
		return api.VehicleStatus{}, err
	}
	return vehicle_status, nil
}

func StartClimate(auth api.Auth, vehicle api.Vehicle, climate api.ClimateInput) error {
	remote_action_result, err := api.StartClimate(auth, vehicle, climate)
	if err != nil {
		return err
	}
	if remote_action_result.EIfresult != "Z:Success" {
		return errors.New(remote_action_result.EIffailmsg)
	}

	return nil
}

func StopClimate(auth api.Auth, vehicle api.Vehicle) error {
	remote_action_result, err := api.StopClimate(auth, vehicle)
	if err != nil {
		return err
	}
	if remote_action_result.EIfresult != "Z:Success" {
		return errors.New(remote_action_result.EIffailmsg)
	}
	return nil
}

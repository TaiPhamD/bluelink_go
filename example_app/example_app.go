package main

import (
	"bluelink_go"
	"bluelink_go/api"
	"fmt"
	"os"
)

func main() {
	// get username, password, pin from environment variables
	// or from a config file
	username := os.Getenv("BL_USERNAME")
	password := os.Getenv("BL_PASSWORD")
	pin := os.Getenv("BL_PIN")
	auth, err := bluelink_go.Login(username, password, pin)
	if err != nil {
		fmt.Println("Error logging in: ", err)
	}
	fmt.Println("username: ", auth.Username)
	fmt.Println("token: ", auth.JWT_Token)
	fmt.Println("token expiration: ", auth.JTW_Expiry)
	// calling **GetOwnerInfo** API
	owner_info, err := bluelink_go.GetOwnerInfo(auth)
	if err != nil {
		fmt.Println("Error getting owner info: ", err)
	}

	// print owner info
	vehicles := owner_info.ResponseString.OwnersVehiclesInfo
	for _, vehicle := range vehicles {
		fmt.Println("Vehicle: ", vehicle)
		fmt.Println("VIN: ", vehicle.VinNumber)
		fmt.Println("Model: ", vehicle.VehicleNickName)
		fmt.Println("Model Year: ", vehicle.Year)
		fmt.Println("RegistrationID: ", vehicle.RegistrationID)
	}

	// get vehicle status for the first vehicle

	vin := vehicles[0].VinNumber
	fmt.Println("VIN: ", vin)
	mycar, err := bluelink_go.GetVehicleFromVin(owner_info, vin)
	if err != nil {
		fmt.Println("Error getting vehicle from VIN: ", err)
		os.Exit(1)
	}
	// print mycar gen
	fmt.Println("My car gen: ", mycar.Gen)

	// calling **GetVehicleStatus** API. Set 3rd param to true to force refresh
	vehicle_status, err := bluelink_go.GetVehicleStatus(auth, mycar, false)
	if err != nil {
		fmt.Println("Error getting vehicle status: ", err)
	}
	// print vehicle DateTime
	fmt.Println("Vehicle status last update time DateTime: ", vehicle_status.ResponseString.VehicleStatus.DateTime)
	// print vehicle status battery level
	fmt.Println("Battery Level: ", vehicle_status.ResponseString.VehicleStatus.EvStatus.BatteryStatus)
	// print long/lat
	fmt.Println("Longitude: ", vehicle_status.ResponseString.VehicleStatus.VehicleLocation.Coord.Lon)
	fmt.Println("Latitude: ", vehicle_status.ResponseString.VehicleStatus.VehicleLocation.Coord.Lat)
	// print ignition status
	fmt.Println("Ignition Status: ", vehicle_status.ResponseString.VehicleStatus.IgnitionStatus)

	// create climate input object
	seatingventinfo := api.SeatHeaterVentInfo{}
	seatingventinfo.DrvSeatHeatState = "3"
	seatingventinfo.AstSeatHeatState = "0"
	StartClimateInput := api.ClimateInput{
		AirCtrl:            "true",
		IgniOnDuration:     "NaN", //max is 10 minutes or NaN for default 10
		AirTempvalue:       "72",
		Defrost:            "false",
		Heating1:           "0",
		SeatHeaterVentInfo: seatingventinfo,
	}

	// call **StartClimate** API

	err = bluelink_go.StartClimate(auth, mycar, StartClimateInput)
	if err != nil {
		fmt.Println("Error starting climate: ", err)
	}
	fmt.Println("Started climate successfully")

	/*
		// call **StopClimate** API
		err = bluelink_go.StopClimate(auth, mycar)
		if err != nil {
			fmt.Println("Error stopping climate: ", err)
		}
	*/

}

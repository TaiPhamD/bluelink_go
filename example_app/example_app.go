package main

import (
	"fmt"
	"os"
	"time"

	"github.com/TaiPhamD/bluelink_go"
)

func main() {
	// get username, password, pin from environment variables
	// or from a config file
	username := os.Getenv("BL_USERNAME")
	password := os.Getenv("BL_PASSWORD")
	pin := os.Getenv("BL_PIN")
	vin := os.Getenv("BL_VIN")
	auth, err := bluelink_go.Login(username, password, pin, vin)
	if err != nil {
		fmt.Println("Error logging in: ", err)
	}
	fmt.Println("username: ", auth.Username)
	fmt.Println("access token: ", auth.AccessToken)
	fmt.Println("refresh token: ", auth.RefreshToken)
	fmt.Println("token expiration: ", auth.ExpiresAt)
	fmt.Println("registration id: ", auth.RegID)

	// sleep 5 seconds

	time.Sleep(5 * time.Second)

	// test refreshign token and print new expiration date
	auth, err = bluelink_go.RefreshToken(auth)
	if err != nil {
		fmt.Println("Error refreshing token: ", err)
	}
	fmt.Println("token expiration: ", auth.ExpiresAt)

	// odometer, _ := bluelink_go.GetOdometer(auth)
	// fmt.Println("odometer: ", odometer)

	// // StartClimate
	// err = bluelink_go.StartClimate(auth, 72)
	// if err != nil {
	// 	fmt.Println("Error starting climate: ", err)
	// }

	// // StopClimate
	// err = bluelink_go.StopClimate(auth)
	// if err != nil {
	// 	fmt.Println("Error stopping climate: ", err)
	// }

}

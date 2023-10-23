package api

// structs to parse json data from bluelink api
// auto generated via: https://mholt.github.io/json-to-go/

type Auth struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64
	Username     string
	RegID        string
	PIN          string
	VIN          string
}

// response from : https://api.telematics.hyundaiusa.com/ac/v2/enrollment/details/[EMAIL_ADDRESS]
type EnrollmentDetails struct {
	EnrolledVehicleDetails []struct {
		PackageDetails []struct {
			AssetNumber     string `json:"assetNumber"`
			EndDate         string `json:"endDate"`
			DisplayCategory string `json:"displayCategory"`
			PackageID       string `json:"packageId"`
			Term            string `json:"term"`
			RenewalDate     string `json:"renewalDate"`
			PackageType     string `json:"packageType"`
			StartDate       string `json:"startDate"`
		} `json:"packageDetails"`
		VehicleDetails struct {
			SvrStatus                string `json:"svrStatus"`
			DynamicBurgerMenu        string `json:"dynamicBurgerMenu"`
			RemoteStartWakeupDays    string `json:"remoteStartWakeupDays"`
			EnrollmentDate           string `json:"enrollmentDate"`
			SvdDay                   string `json:"svdDay"`
			Trim                     string `json:"trim"`
			ModelCode                string `json:"modelCode"`
			UbiCapabilityInd         string `json:"ubiCapabilityInd"`
			Vin                      string `json:"vin"`
			EnrollmentID             string `json:"enrollmentId"`
			SideMirrorHeatCapable    string `json:"sideMirrorHeatCapable"`
			Ownersuccession          string `json:"ownersuccession"`
			Odometer                 string `json:"odometer"`
			NickName                 string `json:"nickName"`
			DefaultBurgerMenu        string `json:"defaultBurgerMenu"`
			EvStatus                 string `json:"evStatus"`
			ModelYear                string `json:"modelYear"`
			SteeringWheelHeatCapable string `json:"steeringWheelHeatCapable"`
			DefaultDashboard         string `json:"defaultDashboard"`
			VehicleGeneration        string `json:"vehicleGeneration"`
			Starttype                string `json:"starttype"`
			EnrollmentType           string `json:"enrollmentType"`
			SapColorCode             string `json:"sapColorCode"`
			BluelinkEnabled          bool   `json:"bluelinkEnabled"`
			OdometerUpdateDate       string `json:"odometerUpdateDate"`
			FatcAvailable            string `json:"fatcAvailable"`
			Color                    string `json:"color"`
			MaintSyncCapable         string `json:"maintSyncCapable"`
			BrandIndicator           string `json:"brandIndicator"`
			DeviceStatus             string `json:"deviceStatus"`
			SetOffPeak               string `json:"setOffPeak"`
			MapProvider              string `json:"mapProvider"`
			GeneralBurgerMenu        string `json:"generalBurgerMenu"`
			InteriorColor            string `json:"interiorColor"`
			AccessoryCode            string `json:"accessoryCode"`
			Nadid                    string `json:"nadid"`
			Mit                      string `json:"mit"`
			Regid                    string `json:"regid"`
			BlueLink                 string `json:"blueLink"`
			WaypointInd              string `json:"waypointInd"`
			BillingInd               string `json:"billingInd"`
			DynamicDashboard         string `json:"dynamicDashboard"`
			Imat                     string `json:"imat"`
			AdditionalVehicleDetails struct {
				TemperatureRange                       string `json:"temperatureRange"`
				TmuSleepMode                           string `json:"tmuSleepMode"`
				EnableHCAModule                        string `json:"enableHCAModule"`
				MaxTemp                                int    `json:"maxTemp"`
				IcpParking                             int    `json:"icpParking"`
				RemoteLockConsentForRemoteStart        string `json:"remoteLockConsentForRemoteStart"`
				CalendarVehicleSyncEnable              string `json:"calendarVehicleSyncEnable"`
				VehicleModemType                       string `json:"vehicleModemType"`
				IcpAACapable                           string `json:"icpAACapable"`
				IcpDriveThru                           int    `json:"icpDriveThru"`
				DkType                                 string `json:"dkType"`
				DynamicSOCText                         string `json:"dynamicSOCText"`
				EnableRoadSideAssitanceAAAModule       string `json:"enableRoadSideAssitanceAAAModule"`
				IdleSpeedinValetAlert                  string `json:"idleSpeedinValetAlert"`
				EvAlarmOptionInfo                      string `json:"evAlarmOptionInfo"`
				MapOtaAccepted                         string `json:"mapOtaAccepted"`
				DkCapable                              string `json:"dkCapable"`
				CombinedHeatSettingsEnable             string `json:"combinedHeatSettingsEnable"`
				IcpChargingStation                     int    `json:"icpChargingStation"`
				HyundaiHome                            string `json:"hyundaiHome"`
				WifiHotspotCapable                     string `json:"wifiHotspotCapable"`
				DkEnrolled                             string `json:"dkEnrolled"`
				IcpAvntCapable                         string `json:"icpAvntCapable"`
				EnableEVTrip                           string `json:"enableEVTrip"`
				MinTemp                                int    `json:"minTemp"`
				IcpFuelStation                         int    `json:"icpFuelStation"`
				TargetSOCLevelMax                      int    `json:"targetSOCLevelMax"`
				RemoteLockConsentForRemoteStartCapable string `json:"remoteLockConsentForRemoteStartCapable"`
				EaPromotion                            struct {
					ExpireOn       string `json:"expireOn"`
					Description    string `json:"description"`
					EnrollmentCode string `json:"enrollmentCode"`
				} `json:"eaPromotion"`
				MsCapableOption      string `json:"msCapableOption"`
				IcpCPCapable         string `json:"icpCPCapable"`
				EnableValetActivate  string `json:"enableValetActivate"`
				EnergyConsoleCapable string `json:"energyConsoleCapable"`
				CpoVehicle           string `json:"cpoVehicle"`
			} `json:"additionalVehicleDetails"`
			Transmissiontype      string `json:"transmissiontype"`
			BluelinkEnrolled      bool   `json:"bluelinkEnrolled"`
			TargetSOCLevel        string `json:"targetSOCLevel"`
			RearWindowHeatCapable string `json:"rearWindowHeatCapable"`
			PreferredDealerCode   string `json:"preferredDealerCode"`
			HmaModel              string `json:"hmaModel"`
			Series                string `json:"series"`
			EnrollmentStatus      string `json:"enrollmentStatus"`
			GeneralDashboard      string `json:"generalDashboard"`
			Userprofilestatus     string `json:"userprofilestatus"`
		} `json:"vehicleDetails"`
		RoleDetails []struct {
			RoleCode string `json:"roleCode"`
			RoleName string `json:"roleName"`
		} `json:"roleDetails"`
		ResponseHeaderMap struct {
		} `json:"responseHeaderMap"`
	} `json:"enrolledVehicleDetails"`
	AddressDetails []struct {
		City       string `json:"city"`
		Street     string `json:"street"`
		PostalCode string `json:"postalCode"`
		Type       string `json:"type"`
		Region     string `json:"region"`
	} `json:"addressDetails"`
	EmergencyContacts []struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		ContactID string `json:"contactId"`
		Phones    []struct {
			Number string `json:"number"`
			Type   string `json:"type"`
			Order  int    `json:"order"`
		} `json:"phones"`
		Relationship string `json:"relationship"`
		Email        string `json:"email"`
	} `json:"emergencyContacts"`
	User struct {
		AccountID   string `json:"accountId"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		PhonesOptIn []struct {
			Number                string `json:"number"`
			PrimaryPhoneIndicator string `json:"primaryPhoneIndicator"`
			FccOptIn              string `json:"fccOptIn"`
			Type                  string `json:"type"`
		} `json:"phonesOptIn"`
		LoginID               string `json:"loginId"`
		AdditionalUserDetails struct {
			UserProfileUpdate     string `json:"userProfileUpdate"`
			TimezoneOffset        int    `json:"timezoneOffset"`
			BillingAccountNumber  string `json:"billingAccountNumber"`
			AppRating             string `json:"appRating"`
			GeoLocationConsent    string `json:"geoLocationConsent"`
			TimezoneAbbr          string `json:"timezoneAbbr"`
			OtaAcceptance         string `json:"otaAcceptance"`
			TelematicsPhoneNumber string `json:"telematicsPhoneNumber"`
		} `json:"additionalUserDetails"`
		TncFlag string `json:"tncFlag"`
		Phones  []struct {
			Number string `json:"number"`
			Type   string `json:"type"`
			Order  int    `json:"order"`
		} `json:"phones"`
		IdmID             string `json:"idmId"`
		UserID            string `json:"userId"`
		NotificationEmail string `json:"notificationEmail"`
		Email             string `json:"email"`
	} `json:"user"`
}

package api

import "time"

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

// response from https://api.telematics.hyundaiusa.com/ac/v2/rcs/rvs/vehicleStatus

type VehicleStatusResponse struct {
	HataTID       string `json:"hataTID"`
	VehicleStatus struct {
		DateTime       time.Time `json:"dateTime"`
		Acc            bool      `json:"acc"`
		DefrostStatus  string    `json:"defrostStatus"`
		TransCond      bool      `json:"transCond"`
		DoorLockStatus string    `json:"doorLockStatus"`
		DoorOpen       struct {
			FrontRight int `json:"frontRight"`
			FrontLeft  int `json:"frontLeft"`
			BackLeft   int `json:"backLeft"`
			BackRight  int `json:"backRight"`
		} `json:"doorOpen"`
		WasherFluidStatus bool `json:"washerFluidStatus"`
		Battery           struct {
			BatSoc          int `json:"batSoc"`
			SjbDeliveryMode int `json:"sjbDeliveryMode"`
		} `json:"battery"`
		HazardStatus    int `json:"hazardStatus"`
		VehicleLocation struct {
			Coord struct {
				Alt  float64 `json:"alt"`
				Lon  float64 `json:"lon"`
				Type int     `json:"type"`
				Lat  float64 `json:"lat"`
			} `json:"coord"`
		} `json:"vehicleLocation"`
		Ign3                   bool   `json:"ign3"`
		IgnitionStatus         string `json:"ignitionStatus"`
		LowFuelLight           bool   `json:"lowFuelLight"`
		SideBackWindowHeat     int    `json:"sideBackWindowHeat"`
		Engine                 bool   `json:"engine"`
		RemoteWaitingTimeAlert struct {
			RemoteControlWaitingTime int `json:"remoteControlWaitingTime"`
			RemoteControlAvailable   int `json:"remoteControlAvailable"`
		} `json:"remoteWaitingTimeAlert"`
		HoodOpen               bool   `json:"hoodOpen"`
		BreakOilStatus         bool   `json:"breakOilStatus"`
		AirConditionStatus     string `json:"airConditionStatus"`
		SmartKeyBatteryWarning bool   `json:"smartKeyBatteryWarning"`
		SteerWheelHeat         int    `json:"steerWheelHeat"`
		TailLampStatus         int    `json:"tailLampStatus"`
		TrunkOpen              bool   `json:"trunkOpen"`
		DoorLock               bool   `json:"doorLock"`
		AirCtrlOn              bool   `json:"airCtrlOn"`
		AirTemp                struct {
			Unit         int    `json:"unit"`
			HvacTempType int    `json:"hvacTempType"`
			Value        string `json:"value"`
		} `json:"airTemp"`
		EvStatus struct {
			ValueDiff   int `json:"valueDiff"`
			RemainTime2 struct {
				Etc3 struct {
					Unit  int `json:"unit"`
					Value int `json:"value"`
				} `json:"etc3"`
				Etc2 struct {
					Unit  int `json:"unit"`
					Value int `json:"value"`
				} `json:"etc2"`
				Atc struct {
					Unit  int `json:"unit"`
					Value int `json:"value"`
				} `json:"atc"`
				Etc1 struct {
					Unit  int `json:"unit"`
					Value int `json:"value"`
				} `json:"etc1"`
			} `json:"remainTime2"`
			EvIgnitionStatus     bool  `json:"evIgnitionStatus"`
			InletLockModeStatus  int   `json:"inletLockModeStatus"`
			BatteryPlugin        int   `json:"batteryPlugin"`
			TimeDiff             int   `json:"timeDiff"`
			BatteryCharge        bool  `json:"batteryCharge"`
			BatteryStatus        int   `json:"batteryStatus"`
			RemainTime           []any `json:"remainTime"`
			BatteryPreconditiong bool  `json:"batteryPreconditiong"`
			DrvDistance          []struct {
				RangeByFuel struct {
					GasModeRange struct {
						Unit  int     `json:"unit"`
						Value float64 `json:"value"`
					} `json:"gasModeRange"`
					TotalAvailableRange struct {
						Unit  int     `json:"unit"`
						Value float64 `json:"value"`
					} `json:"totalAvailableRange"`
					EvModeRange struct {
						Unit  int     `json:"unit"`
						Value float64 `json:"value"`
					} `json:"evModeRange"`
				} `json:"rangeByFuel"`
				Type int `json:"type"`
			} `json:"drvDistance"`
			ReservChargeInfos struct {
				TargetSOClist []struct {
					RangeByFuel struct {
						GasModeRange struct {
							Unit  int `json:"unit"`
							Value int `json:"value"`
						} `json:"gasModeRange"`
						TotalAvailableRange struct {
							Unit  int `json:"unit"`
							Value int `json:"value"`
						} `json:"totalAvailableRange"`
						EvModeRange struct {
							Unit  int `json:"unit"`
							Value int `json:"value"`
						} `json:"evModeRange"`
					} `json:"rangeByFuel"`
					PlugType       int `json:"plugType"`
					TargetSOClevel int `json:"targetSOClevel"`
					Type           int `json:"type"`
				} `json:"targetSOClist"`
				ReservChargeInfo struct {
					DateTime               string `json:"dateTime"`
					ReservChargeInfoDetail struct {
						ReservChargeSet bool   `json:"reservChargeSet"`
						ReservEndTime   string `json:"reservEndTime"`
						ReservFatcSet   struct {
							Defrost bool `json:"defrost"`
							AirTemp struct {
								Unit  int    `json:"unit"`
								Value string `json:"value"`
							} `json:"airTemp"`
							AirCtrl int `json:"airCtrl"`
						} `json:"reservFatcSet"`
						ReservInfo struct {
							Time struct {
								TimeSection int    `json:"timeSection"`
								Time        string `json:"time"`
							} `json:"time"`
							Day []int `json:"day"`
						} `json:"reservInfo"`
					} `json:"reservChargeInfoDetail"`
				} `json:"reservChargeInfo"`
				ReservFlag         int `json:"reservFlag"`
				ReserveChargeInfo2 struct {
					ReservChargeInfoDetail struct {
						ReservChargeSet bool   `json:"reservChargeSet"`
						ReservEndTime   string `json:"reservEndTime"`
						ReservFatcSet   struct {
							Defrost bool `json:"defrost"`
							AirTemp struct {
								Unit  int    `json:"unit"`
								Value string `json:"value"`
							} `json:"airTemp"`
							AirCtrl int `json:"airCtrl"`
						} `json:"reservFatcSet"`
						ReservInfo struct {
							Time struct {
								TimeSection int    `json:"timeSection"`
								Time        string `json:"time"`
							} `json:"time"`
							Day []int `json:"day"`
						} `json:"reservInfo"`
					} `json:"reservChargeInfoDetail"`
				} `json:"reserveChargeInfo2"`
				OffpeakPowerInfo struct {
					OffPeakPowerFlag  int `json:"offPeakPowerFlag"`
					OffPeakPowerTime1 struct {
						StartTime struct {
							TimeSection int    `json:"timeSection"`
							Time        string `json:"time"`
						} `json:"startTime"`
						EndTime struct {
							TimeSection int    `json:"timeSection"`
							Time        string `json:"time"`
						} `json:"endTime"`
					} `json:"offPeakPowerTime1"`
				} `json:"offpeakPowerInfo"`
			} `json:"reservChargeInfos"`
		} `json:"evStatus"`
		LampWireStatus struct {
			HeadLamp struct {
				RightBifuncLamp bool `json:"rightBifuncLamp"`
				HeadLampStatus  bool `json:"headLampStatus"`
				LeftLowLamp     bool `json:"leftLowLamp"`
				RightHighLamp   bool `json:"rightHighLamp"`
				LeftBifuncLamp  bool `json:"leftBifuncLamp"`
				LeftHighLamp    bool `json:"leftHighLamp"`
				RightLowLamp    bool `json:"rightLowLamp"`
			} `json:"headLamp"`
			StopLamp struct {
				RightLamp bool `json:"rightLamp"`
				LeftLamp  bool `json:"leftLamp"`
			} `json:"stopLamp"`
			TurnSignalLamp struct {
				RightRearLamp  bool `json:"rightRearLamp"`
				RightFrontLamp bool `json:"rightFrontLamp"`
				LeftRearLamp   bool `json:"leftRearLamp"`
				LeftFrontLamp  bool `json:"leftFrontLamp"`
			} `json:"turnSignalLamp"`
		} `json:"lampWireStatus"`
		SleepModeCheck   bool `json:"sleepModeCheck"`
		Defrost          bool `json:"defrost"`
		TirePressureLamp struct {
			TirePressureWarningLampRearLeft   int `json:"tirePressureWarningLampRearLeft"`
			TirePressureWarningLampFrontLeft  int `json:"tirePressureWarningLampFrontLeft"`
			TirePressureWarningLampFrontRight int `json:"tirePressureWarningLampFrontRight"`
			TirePressureWarningLampAll        int `json:"tirePressureWarningLampAll"`
			TirePressureWarningLampRearRight  int `json:"tirePressureWarningLampRearRight"`
		} `json:"tirePressureLamp"`
		TrunkOpenStatus string `json:"trunkOpenStatus"`
	} `json:"vehicleStatus"`
}

type RunningStatusResponse struct {
	NextPollingInterval string `json:"nextPollingInterval"`
	Tid                 string `json:"tid"`
	Status              string `json:"status"`
}

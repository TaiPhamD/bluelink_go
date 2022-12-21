# bluelink go wrapper
This is a bluelink go wrapper based on [bluelink](https://github.com/synchronizing/bluelink) python package but I've added additional APIs to get vehicle status information like battery level, lock status etc.

see this example_app for using this go module

# APIs

- Login - returns an auth object required to call all apis below
- GetOwnerInfo - return OwnerInfoService to get information such as vehicles and their reg id
- GetVehicleFromVin - helper routine to get vehicle vin/regid/gen based on OwnerInfoService struct
- GetVehicleStatus - provides battery status, door status, ignition status etc
- StartClimate - start climate based on climate input struct
- StopClimate - stop climate control
- RemoteLockAction - Remote lock/unlock

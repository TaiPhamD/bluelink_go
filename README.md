# bluelink go wrapper
This is a bluelink go wrapper based on [bluelink](https://github.com/synchronizing/bluelink) python package but I've added additional APIs to get vehicle status information like battery level, lock status etc.

see this example_app for using this go module

# APIs

- Login - returns an auth object required to call all apis below
- GetVehicleStatus - provides battery status, door status, ignition status etc
- StartClimate - start climate based on climate input struct
- StopClimate - stop climate control
- RemoteLockAction - Remote lock/unlock

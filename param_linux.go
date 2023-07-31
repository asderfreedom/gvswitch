package dtun

type DevicePermissions struct {
	Owner uint
	Group uint
}

type PlatformSpecificParams struct {
	// Name is the name to be set for the interface to be created. This overrides the default name assigned by the os such
	// as TUN0 TAP0. A zero-value of this field, i.e. an empty string, indicates that the default name should be used.
	Name string

	// Persist specifics whether persistence mode for the interface device should be enabled or disabled.
	Persist bool

	// Permissions if non-nil, specifics the owner and the group owner for the interface, A zero-value will specific
	// that no changes to the owner and the group owner will be made.
	Permissions *DevicePermissions

	// MultiQueue specific if open parallelize packets sending and receiving mode. this will not be used in this project.
	MultiQueue bool
}

func defaultPlatformSpecificParams() PlatformSpecificParams {
	return PlatformSpecificParams{}
}

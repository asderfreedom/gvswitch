package dtun

import (
	"errors"
	"fmt"
	"io"
)

type Interface struct {
	isTAP bool
	io.ReadWriteCloser
	name string
}

// DeviceType is the type for specifying device types
type DeviceType int

const (
	_ = iota
	TUN
	TAP
)

type Config struct {
	// DeviceType specific whether the device is a TUN or TAP interface, A zero-value is treated as TUN.
	DeviceType DeviceType

	PlatformSpecificParams
}

func defaultConfig() Config {
	return Config{
		DeviceType:             TUN,
		PlatformSpecificParams: defaultPlatformSpecificParams(),
	}
}

var zeroConfig Config

// New Creates a new TUN/TAP interface using config
func New(config Config) (face *Interface, err error) {
	if zeroConfig == config {
		config = defaultConfig()
	}
	if config.PlatformSpecificParams == zeroConfig.PlatformSpecificParams {
		config.PlatformSpecificParams = defaultPlatformSpecificParams()
	}
	switch config.DeviceType {
	case TUN, TAP:
		return openDev(config)
	default:
		return nil, errors.New("unknown device type")
	}
}

func (i *Interface) IsTAP() bool {
	return i.isTAP
}

func (i *Interface) IsTUN() bool {
	return !i.isTAP
}

// Name returns the name of the interface
func (i *Interface) Name() string {
	return i.name
}

// NewTAP Creates a new TAP interface with the specified name
func NewTAP(ifName string) (*Interface, error) {
	fmt.Println("Deprecated: use New(Config{DeviceType: TAP, PlatformSpecificParams: PlatformSpecificParams{Name: ifName}}) instead")
	config := Config{DeviceType: TAP}
	config.Name = ifName
	return openDev(config)
}

// NewTUN Creates a new TUN interface with the specified name
func NewTUN(ifName string) (*Interface, error) {
	fmt.Println("Deprecated: use New(Config{DeviceType: TUN, PlatformSpecificParams: PlatformSpecificParams{Name: ifName}}) instead")
	config := Config{DeviceType: TUN}
	config.Name = ifName
	return openDev(config)
}

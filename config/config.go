package config

var (
	Config = &ConfigEntries{}
)

type ConfigEntries struct {
	Username string
	Password string

	IpAddress        string
	NetworkInterface string

	Debug bool
}

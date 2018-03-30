package devicebundle

// Device struct
type Device struct {
	Hostname  string `json:"hostname"`
	PrivateIP string `json:"privateIP"`
	PublicIP  string `json:"publicIP"`
}

// Create a new device
func NewDevice(hostName string, privIP string, pubIP string) *Device {
	return &Device{
		Hostname:  hostName,
		PrivateIP: privIP,
		PublicIP:  pubIP,
	}
}

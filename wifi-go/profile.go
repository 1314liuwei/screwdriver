package wifi

type Profile struct {
	SSID      string   `json:"ssid"`
	RSSI      int      `json:"rssi"`
	Frequency []string `json:"frequency"`
	Akm       []string `json:"akm"`
	Password  string   `json:"password,omitempty"`
}

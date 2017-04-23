package model

type DeviceInfo struct {
	UUID     string `json:"uuid"`
	Platform string `json:"platform"`
	Name     string `json:"name"`
}

type UserService struct {
	Email  string `json:"email"`
	Addr   string `json:"addr"`
	UUID   string `json:"uuid"`
	Device *DeviceInfo `json:"device"`
}



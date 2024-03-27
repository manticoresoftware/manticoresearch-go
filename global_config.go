package openapi

var defaultIP = ""

func SetDefaultIP(IP string) {
	defaultIP = IP
}

func GetDefaultIP() string {
	if defaultIP != "" {
		return defaultIP
	}
	return "127.0.0.1"
}

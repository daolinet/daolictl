package client

type (
	Firewall struct {
		Name        string
		Container   string
		DatapathID  string
		GatewayIP   string
		GatewayPort int
		ServicePort int
	}

	ContainerNetwork struct {
		Id          string
		NetworkName string
		IPAddress   string
		MacAddress  string
		Gateway     string
		VIPAddress  string
	}
)

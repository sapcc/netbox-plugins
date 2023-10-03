package sism_v2

type SISM_Device struct {
	Id      string `json:"sism_id"`
	Serial  string `json:"sism_serial"`
	Equi_no string `json:"equi_no"`
	Device  Device `json:"device"`
}

type Device struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Name    string `json:"name"`
	Display string `json:"display"`
}

const PluginPath = "netbox-sism/netboxsism/"

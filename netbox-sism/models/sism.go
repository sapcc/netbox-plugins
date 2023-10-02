package models

type SISM_Device struct {
	Id     string `json:"sism_id"`
	Serial string `json:"sism_serial"`
}

const PluginPath = "/netbox-sism/netboxsism/"

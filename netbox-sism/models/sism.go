package models

import "github.com/go-openapi/strfmt"

type SISM_Device struct {
	Id              string          `json:"sism_id"`
	Serial          string          `json:"sism_serial"`
	Planned_removal strfmt.DateTime `json:"planned_removal"`
	Equi_no         string          `json:"equi_no"`
	Last_sync       strfmt.DateTime `json:"last_sync"`
	Device          NestedDevice    `json:"device"`
}

type NestedDevice struct {
	Id          int    `json:"id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

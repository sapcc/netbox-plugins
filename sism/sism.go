package sism

import "github.com/go-openapi/strfmt"

type SISM_Device struct {
	Id              string      `json:"sism_id"`
	Serial          string      `json:"sism_serial"`
	Planned_removal strfmt.Date `json:"planned_removal"`
	Equi_no         string      `json:"equi_no"`
	Last_sync       strfmt.Date `json:"last_sync"`
	Device          Device      `json:"device"`
}

type Device struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Name    string `json:"name"`
	Display string `json:"display"`
}

package sism_v2

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
)

func TestClient_GetSISMDevice(t *testing.T) {
	c, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = c.HttpClient
	vcr := govcr.NewVCR("GetSISMDevice", vcrConf)
	c.HttpClient = vcr.Client
	res, err := c.GetSISMDevice(12)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Device ID:", res.Id)
}

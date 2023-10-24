package clusters_v1

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
)

func TestClient_GetClusterByName(t *testing.T) {
	c, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = c.HttpClient
	vcr := govcr.NewVCR("GetSISMDevice", vcrConf)
	c.HttpClient = vcr.Client
	res, err := c.GetClusterTypeByName("test")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Cluster ID:", res.Results)
}
func TestClient_AddDeleteClusterType(t *testing.T) {
	c, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = c.HttpClient
	vcr := govcr.NewVCR("AddDeleteClusterType", vcrConf)
	c.HttpClient = vcr.Client
	ct := PhysicalClusterTypeRequest{}
	ct.Name = "testClusterType001"

	// add cluster type
	res, err := c.AddClusterType(ct)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Cluster ID:", res)

	// delete clusterType
	// err = c.DeleteClusterType(res.Id)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

func TestClient_AddDeleteCluster(t *testing.T) {
	c, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = c.HttpClient
	vcr := govcr.NewVCR("AddDeleteCluster", vcrConf)
	c.HttpClient = vcr.Client
	wt := WritablePhysicalClusterRequest{}
	wt.Name = "testCluster001"
	wt.ClusterType = 5
	// add cluster
	res, err := c.AddCluster(wt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Cluster ID:", res.Id)

	// delete cluster
	err = c.DeleteCluster(res.Id)
	if err != nil {
		t.Fatal(err)
	}
}

package clusters_v1

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
)

// Cluster Types Functions
func TestClient_GetClusterType(t *testing.T) {
	c, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = c.HttpClient
	vcr := govcr.NewVCR("GetClusterType", vcrConf)
	c.HttpClient = vcr.Client
	res, err := c.GetClusterType(5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Cluster ID:", res)
}
func TestClient_ListClusterTypes(t *testing.T) {
	c, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = c.HttpClient
	vcr := govcr.NewVCR("ListClusterTypes", vcrConf)
	c.HttpClient = vcr.Client
	res, err := c.ListClusterTypes()
	if err != nil {
		t.Fatal(err)
	}
	for _, ct := range res.Results {
		t.Log("Cluster Type:", ct.Id, ct.Name)
	}

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
	err = c.DeleteClusterType(res.Id)
	if err != nil {
		t.Fatal(err)
	}
}

// Cluster Functions
func TestClient_GetCluster(t *testing.T) {
	c, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = c.HttpClient
	vcr := govcr.NewVCR("GetCluster", vcrConf)
	c.HttpClient = vcr.Client
	res, err := c.GetCluster(3)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Cluster:", res)
}
func TestClient_ListClusters(t *testing.T) {
	c, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = c.HttpClient
	vcr := govcr.NewVCR("ListCluster", vcrConf)
	c.HttpClient = vcr.Client
	var opts ListClustersRequest
	opts.Site = "NA-US-1a"
	res, err := c.ListClusters(opts)
	if err != nil {
		t.Fatal(err)
	}
	//for _, ct := range res.Results {
	t.Log("Clusters:", res)
	//}

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
	t.Log("Cluster ID:", res)

	// delete cluster
	err = c.DeleteCluster(res.Id)
	if err != nil {
		t.Fatal(err)
	}
}

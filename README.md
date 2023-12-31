# netbox-plugins
code to interact with different API versions of netbox plugins

## sample
```
package main

import (
	"fmt"
	"os"

	sism "github.com/sapcc/netbox-plugins/sism/sism_v2"
	clusters "github.com/sapcc/netbox-plugins/clusters/clusters_v1"
)

func main() {

	// initialize client with netbox url and token
	c, _ := sism.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)

	// execute sism_v2 function GetSISMDevice
	res, _ := c.GetSISMDevice(12)

	// print result info
	fmt.Println("SISM ID:", res.Id)
	fmt.Println("SISM Serial:", res.Serial)
	fmt.Println("Device ID:", res.Device.Id)
	fmt.Println("Device Name:", res.Device.Name)

	cc, _ := clusters.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	wt := clusters.WritablePhysicalClusterRequest{}
	wt.Name = "coolCluster"
	wt.ClusterType = 5
	cluster, err := cc.AddCluster(wt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cluster created")
	fmt.Println(cluster.Id)
	fmt.Println(cluster.Name)
	fmt.Println(cluster.Url)

}

```

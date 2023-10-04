# netbox-plugins
code to interact with different API versions of netbox plugins

## sample
```
package main

import (
	"fmt"
	"os"

	sism "github.com/sapcc/netbox-plugins/sism/sism_v2"
)

func main() {
	fmt.Println("start")

	c, _ := sism.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	res, _ := c.GetSISMDevice(12)
	fmt.Println("SISM ID:", res.Id)
	fmt.Println("SISM Serial:", res.Serial)
	fmt.Println("Device ID:", res.Device.Id)
	fmt.Println("Device Name:", res.Device.Name)
}

```

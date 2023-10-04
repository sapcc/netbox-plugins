package sism_v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	common "github.com/sapcc/netbox-plugins/common"
)

func (c *Client) GetSISMDevice(deviceID int) (*SISM_Device, error) {

	fmt.Println("pluginVersion: ", PluginVersion)
	url := c.BaseUrl.String() + common.BasePath + PluginPath + strconv.Itoa(deviceID)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	var resObj = SISM_Device{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}

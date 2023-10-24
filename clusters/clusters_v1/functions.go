package clusters_v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	common "github.com/sapcc/netbox-plugins/common"
)

func (c *Client) GetClusterTypeByName(clusterName string) (*ListClustersResponse, error) {

	fmt.Println("pluginVersion: ", PluginVersion)
	url := c.BaseUrl.String() + common.BasePath + PluginPath_Cluster

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	q := request.URL.Query()
	q.Set("name", clusterName)
	request.URL.RawQuery = q.Encode()

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	var resObj = ListClustersResponse{}
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

func (c *Client) AddClusterType(ct PhysicalClusterTypeRequest) (*PhysicalCluster, error) {
	fmt.Println("pluginVersion: ", PluginVersion)

	body, err := json.Marshal(ct)
	if err != nil {
		return nil, err
	}
	url := c.BaseUrl.String() + common.BasePath + PluginPath_ClusterTypes

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 201 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}
	var resObj = PhysicalCluster{}
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

func (c *Client) DeleteClusterType(clusterTypeID int) error {
	fmt.Println("pluginVersion: ", PluginVersion)
	url := c.BaseUrl.String() + common.BasePath + PluginPath_ClusterTypes + strconv.Itoa(clusterTypeID) + "/"

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	c.SetAuthToken(&request.Header)

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 204 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}

	return nil
}

func (c *Client) AddCluster(wt WritablePhysicalClusterRequest) (*PhysicalCluster, error) {
	fmt.Println("pluginVersion: ", PluginVersion)

	body, err := json.Marshal(wt)
	if err != nil {
		return nil, err
	}
	url := c.BaseUrl.String() + common.BasePath + PluginPath_Cluster

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 201 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}
	var resObj = PhysicalCluster{}
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

func (c *Client) DeleteCluster(clusterID int) error {
	fmt.Println("pluginVersion: ", PluginVersion)
	url := c.BaseUrl.String() + common.BasePath + PluginPath_Cluster + strconv.Itoa(clusterID) + "/"

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	c.SetAuthToken(&request.Header)

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 204 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}

	return nil
}

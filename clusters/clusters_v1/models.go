package clusters_v1

type PhysicalCluster struct {
	Id      int               `json:"id"`
	Url     string            `json:"url"`
	Name    string            `json:"name"`
	Type    NestedClusterType `json:"cluster_type"`
	Devices []NestedDevice    `json:"devices"`
}
type WritablePhysicalClusterRequest struct {
	Name        string `json:"name"`
	ClusterType int    `json:"cluster_type"`
}

type ClusterType struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

type PhysicalClusterTypeRequest struct {
	Name string `json:"name"`
}

type NestedClusterType struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Name    string `json:"name"`
	Display string `json:"display"`
}

type NestedDevice struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Name    string `json:"name"`
	Display string `json:"display"`
}

type ListClustersRequest struct {
	Name string
}

type ListClustersResponse struct {
	Count   int
	Results []PhysicalCluster `json:"results"`
}

const PluginPath_Cluster = "clusters/clusters/"
const PluginPath_ClusterTypes = "clusters/cluster-types/"
const PluginVersion = "v1"

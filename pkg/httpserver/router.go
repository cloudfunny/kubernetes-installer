package httpserver

import (
	"net/http"

	"github.com/cloudfunny/kubernetes-installer/pkg/model"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterHandler() {
	s.addFunc("/clusters", "GET", listAllClusters)
	s.addFunc("/clusters/watch", "GET", watchClusters)
	s.addFunc("/clusters/:id", "GET", queryClusterByID)
	s.addFunc("/clusters", "POST", createCluster)

	s.addFunc("/clusters/:id/node/register", "POST", registerNode)
	s.addFunc("/nodes/watch", "GET", watchNodes)
}

// hello world for test
func listAllClusters(c *gin.Context) {
	result := make(map[string]interface{})
	clusters := model.ListAllCluster()
	// allDirs := kubeadmutils.ListDirectories()
	result["cluster_info"] = clusters
	// result["all_dirs"] = allDirs
	c.JSON(http.StatusOK, result)
}

func createCluster(c *gin.Context) {
	// name string, version string, podnetcidr string, registry string
	name := c.PostForm("name")
	version := c.PostForm("version")
	registry := c.PostForm("registry")
	podNetCIDR := c.PostForm("podNetCIDR")
	newCluster := model.CreateCluster(name, version, podNetCIDR, registry)
	c.JSON(http.StatusOK, newCluster)
}

func queryClusterByID(c *gin.Context) {
	id := c.Param("id")
	cluster := model.QueryClusterByID(id)
	c.JSON(http.StatusOK, cluster)
}

func registerNode(c *gin.Context) {
	cid := c.Param("id")
	hostName := c.PostForm("hostname")
	role := c.PostForm("role")
	model.RegisterNode(hostName, role, cid)
	cluster := model.QueryClusterByID(cid)
	c.JSON(http.StatusOK, cluster)
}

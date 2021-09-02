package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/cloudfunny/kubernetes-installer/pkg/model"
	"github.com/gin-gonic/gin"
)

func watchClusters(c *gin.Context) {
	c.Header("Transfer-Encoding", "chunked")
	c.Writer.WriteHeader(http.StatusOK)
	flusher, _ := c.Writer.(http.Flusher)
	flusher.Flush()

	for event := range model.WatchClusterChan {
		obj, _ := json.Marshal(event)
		c.Writer.Write(append(obj, '\n'))
		flusher.Flush()
	}
}

func watchNodes(c *gin.Context) {
	c.Header("Transfer-Encoding", "chunked")
	c.Writer.WriteHeader(http.StatusOK)
	flusher, _ := c.Writer.(http.Flusher)
	flusher.Flush()

	for event := range model.WatchNodeChan {
		if event.Obj.(model.Node).ClusterID != c.Param("id") {
			continue
		}
		obj, _ := json.Marshal(event)
		c.Writer.Write(append(obj, '\n'))
		flusher.Flush()
	}
}

package nodehandler

import (
	G_mq "Gungnir/internal/master/G-mq"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

type Nodehandler struct {
}

type TaskRequest struct {
	CPU    int64 `form:"cpu" binding:"required"`
	Memory int64 `form:"memory" binding:"required"`
}

func New() *Nodehandler {
	return &Nodehandler{}
}

func (n *Nodehandler) Hi(c *gin.Context) {
	//G_mq.GetMqOr().NodeConnectionChan <- &G_mq.NodeConnectionMsg{}
	c.String(http.StatusOK, "ok")
}

func (n *Nodehandler) Connect(c *gin.Context) {
	clusterName := c.Query("clusterName")
	addr := c.Query("address")
	G_mq.GetMqOr().NodeConnectionChan <- &G_mq.NodeConnectionMsg{
		Name: clusterName,
		Addr: addr,
	}
	c.String(http.StatusOK, "plz wait")
}

func (n *Nodehandler) UploadJob(c *gin.Context) {

	// get file
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("Error getting file:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to get file",
		})
		return
	}
	defer file.Close()

	// get params
	var taskReq TaskRequest
	if err := c.ShouldBind(&taskReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save file
	saveFilePath := "./job/" + header.Filename
	f, err := os.Create(saveFilePath)
	if err != nil {
		log.Println("Error creating file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to save file",
		})
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		log.Println("Error copying file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to save file",
		})
		return
	}
	defer f.Close()

	// start scheduler
	G_mq.GetMqOr().JobStartSchedulerChan <- &G_mq.JobStartSchedulerMsg{
		Name:   header.Filename,
		Path:   saveFilePath,
		Cpu:    int(taskReq.CPU),
		Memory: int(taskReq.Memory),
	}

	c.String(200, "ok")
}

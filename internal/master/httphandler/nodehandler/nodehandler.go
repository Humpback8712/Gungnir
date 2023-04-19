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

func New() *Nodehandler {
	return &Nodehandler{}
}

func (n *Nodehandler) Hi(c *gin.Context) {
	//G_mq.GetMqOr().NodeConnectionChan <- &G_mq.NodeConnectionMsg{}
	c.String(http.StatusOK, "ok")
}

func (n *Nodehandler) Connect(c *gin.Context) {
	G_mq.GetMqOr().NodeConnectionChan <- &G_mq.NodeConnectionMsg{
		Name: "12",
		Addr: "127.0.0.1:8889",
	}
	c.String(http.StatusOK, "plz wait")
}

func (n *Nodehandler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("Error getting file:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to get file",
		})
		return
	}
	defer file.Close()

	// 将上传的文件保存到本地文件系统
	saveFilePath := "./" + header.Filename
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
	c.String(200, "ok")
}

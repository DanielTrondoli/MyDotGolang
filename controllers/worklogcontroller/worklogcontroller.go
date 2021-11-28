package worklogcontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DanielTrondoli/MyDotGolang/service/worklogservice"
	"github.com/gin-gonic/gin"
)

func Start(c *gin.Context) {
	hashKey, ok := c.GetQuery("uuid")
	if !ok {
		log.Fatal("chave nao informada !")
	}

	worklogservice.Start(hashKey)
	c.Redirect(http.StatusFound, "/")
}

func Stop(c *gin.Context) {
	hashKey, ok := c.GetQuery("uuid")
	if !ok {
		log.Fatal("chave nao informada !")
	}

	fmt.Println("Stop: passou caraio")

	worklogservice.Stop(hashKey)
	c.Redirect(http.StatusFound, "/")
}

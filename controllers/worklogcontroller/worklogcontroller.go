package worklogcontroller

import (
	"log"
	"net/http"

	"github.com/DanielTrondoli/MyDotGolang/service/worklogservice"
	"github.com/gin-gonic/gin"
)

const KEY_UUID = "uuid"

func Start(c *gin.Context) {
	hashKey, ok := c.GetQuery(KEY_UUID)
	if !ok {
		log.Fatal("chave nao informada !")
	}

	worklogservice.Start(hashKey)
	c.Redirect(http.StatusFound, "/")
}

func Stop(c *gin.Context) {
	hashKey, ok := c.GetQuery(KEY_UUID)
	if !ok {
		log.Fatal("chave nao informada !")
	}

	worklogservice.Stop(hashKey)
	c.Redirect(http.StatusFound, "/")
}

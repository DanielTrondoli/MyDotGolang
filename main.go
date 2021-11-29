package main

import (
	"os"

	"github.com/DanielTrondoli/MyDotGolang/repository/issuerepository"
	"github.com/DanielTrondoli/MyDotGolang/routes"
	"github.com/gin-gonic/gin"
)

//var workLogs map[string]issue.Issue

func main() {

	os.Setenv("AWS_ACCESS_KEY_ID", "dummy1")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy2")
	os.Setenv("AWS_SESSION_TOKEN", "dummy3")

	issuerepository.DBInitIssue()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	routes.Routes(router)

	router.Run(":1666")

}

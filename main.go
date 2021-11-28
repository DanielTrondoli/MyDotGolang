package main

import (
	"github.com/DanielTrondoli/MyDotGolang/repository/issuerepository"
	"github.com/DanielTrondoli/MyDotGolang/routes"
	"github.com/gin-gonic/gin"
)

//var workLogs map[string]issue.Issue

func main() {

	issuerepository.DBInitIssue()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	routes.Routes(router)

	router.Run(":1666")

}

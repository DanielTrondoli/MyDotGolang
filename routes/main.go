package routes

import (
	"github.com/DanielTrondoli/MyDotGolang/controllers/issuecontroller"
	"github.com/DanielTrondoli/MyDotGolang/controllers/worklogcontroller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.GET("/", issuecontroller.GetAllIssues)

	// router.GET("/worklogs/:url", getAllWorkLogs)
	router.POST("/issue", issuecontroller.InsertIssues)
	router.GET("/deleteissue", issuecontroller.DeleteIssue)

	router.GET("/worklog_start", worklogcontroller.Start)
	router.GET("/worklog_stop", worklogcontroller.Stop)

	// router.GET("/ticwork", tickWork)

}

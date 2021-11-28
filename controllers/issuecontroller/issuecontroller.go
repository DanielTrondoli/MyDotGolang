package issuecontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DanielTrondoli/MyDotGolang/repository/issuerepository"
	"github.com/DanielTrondoli/MyDotGolang/service/issueservice"
	"github.com/DanielTrondoli/MyDotGolang/timeutils"
	"github.com/gin-gonic/gin"
)

func GetAllIssues(c *gin.Context) {

	date, ok := c.GetQuery("date")
	if !ok {
		date = ""
	}

	allIssues := issueservice.GetAllIssues()

	activates := issueservice.GetActiveIssues(allIssues)

	workLogsOfTheDay := issueservice.GetWorkLogsOfTheDay(allIssues, date)

	currentDate, err := timeutils.GetDayFormatedDMY(date, 0)
	if err != nil {
		log.Default().Print(err.Error())
	}
	nextDate, err := timeutils.GetDayFormatedDMY(date, 1)
	if err != nil {
		log.Default().Print(err.Error())
	}
	previousDate, err := timeutils.GetDayFormatedDMY(date, -1)
	if err != nil {
		log.Default().Print(err.Error())
	}
	nextWeek, err := timeutils.GetDayFormatedDMY(date, 7)
	if err != nil {
		log.Default().Print(err.Error())
	}
	previousWeek, err := timeutils.GetDayFormatedDMY(date, -7)
	if err != nil {
		log.Default().Print(err.Error())
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"data":             allIssues,
		"activates":        activates,
		"workLogsOfTheDay": workLogsOfTheDay,
		"CurrentDate":      currentDate,
		"NextDate":         nextDate,
		"PreviousDate":     previousDate,
		"NextWeek":         nextWeek,
		"PreviousWeek":     previousWeek,
	})
}

func InsertIssues(c *gin.Context) {

	title := c.Request.FormValue("title")
	url := c.Request.FormValue("url")

	err := issueservice.InsertIssues(title, url)
	if err != nil {
		log.Default().Print(err.Error())
	}

	c.Redirect(http.StatusFound, "/")
}

func DeleteIssue(c *gin.Context) {

	hashKey, ok := c.GetQuery("uuid")
	if ok {
		err := issuerepository.DeleteIssue(hashKey)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		log.Fatal("chave a ser deletada nao existe !")
	}

	fmt.Println(" ", hashKey)
	c.Redirect(http.StatusFound, "/")
}

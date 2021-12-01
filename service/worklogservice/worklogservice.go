package worklogservice

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielTrondoli/MyDotGolang/models/worklogs"
	"github.com/DanielTrondoli/MyDotGolang/repository/issuerepository"
)

func Start(hashKey string) {

	issue, err := issuerepository.GetIssue(hashKey)
	if err != nil {
		log.Fatal(err)
	}

	stopAllIssues()

	for _, w := range issue.WorkLogs {
		if w.EndTime == (time.Time{}) {
			fmt.Println("Issue ja possui um work log aberto")
			return
		}
	}

	start := worklogs.WorkLogs{
		StartTime: time.Now(),
	}

	issue.WorkLogs = append(issue.WorkLogs, start)

	issuerepository.UpdateWorkLogs(issue.UUID, issue.WorkLogs)
}

func Stop(hashKey string) {

	issue, err := issuerepository.GetIssue(hashKey)
	if err != nil {
		log.Fatal(err)
	}

	var achou bool
	for i, w := range issue.WorkLogs {
		if w.EndTime == (time.Time{}) {
			issue.WorkLogs[i].EndTime = time.Now()
			achou = true
			break
		}
	}

	if !achou {
		fmt.Println("Issue nao possui um work log aberto")
		return
	}

	issuerepository.UpdateWorkLogs(issue.UUID, issue.WorkLogs)
}

func stopAllIssues() {
	allIssues, err := issuerepository.GetAllIssues()
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range allIssues {
		for j, w := range e.WorkLogs {
			if w.EndTime == (time.Time{}) {
				e.WorkLogs[j].EndTime = time.Now()
				issuerepository.UpdateWorkLogs(e.UUID, e.WorkLogs)
			}
		}
	}
}

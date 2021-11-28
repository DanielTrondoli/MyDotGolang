package issueservice

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielTrondoli/MyDotGolang/models/issue"
	"github.com/DanielTrondoli/MyDotGolang/models/relworklogitem"
	"github.com/DanielTrondoli/MyDotGolang/models/worklogs"
	"github.com/DanielTrondoli/MyDotGolang/repository/issuerepository"
	"github.com/DanielTrondoli/MyDotGolang/timeutils"
	"github.com/gofrs/uuid"
)

func GetAllIssues() []issue.Issue {
	allIssues, err := issuerepository.GetAllIssues()
	if err != nil {
		log.Fatal(err)
	}

	return allIssues
}

func GetActiveIssues(allIssues []issue.Issue) []bool {

	activates := make([]bool, len(allIssues))

	for i, iss := range allIssues {
		for _, w := range iss.WorkLogs {
			if (w.EndTime == time.Time{}) {
				activates[i] = true
				break
			}
		}
	}

	return activates
}

func GetWorkLogsOfTheDay(allIssues []issue.Issue, dateTime string) []relworklogitem.Relworklogitem {

	workLogsOfTheDay := make([]relworklogitem.Relworklogitem, 0)

	now0, err := timeutils.GetNowTruncated()
	if err != nil {
		log.Default().Print(err.Error())
		return []relworklogitem.Relworklogitem{}
	}

	now23, err := timeutils.GetNowEndDay()
	if err != nil {
		log.Default().Print(err.Error())
		return []relworklogitem.Relworklogitem{}
	}
	fmt.Println("GetWorkLogsOfTheDay: ", dateTime)
	if dateTime != "" {
		days, err := timeutils.DifDateFromNow(dateTime)
		if err != nil {
			log.Default().Print(err.Error())
			return []relworklogitem.Relworklogitem{}
		}

		fmt.Println(days)
		if days != 0 {
			now0 = now0.Add(days)
			now23 = now23.Add(days)
		}
		fmt.Println("GetWorkLogsOfTheDay: ", now0, now23)
	}

	for _, iss := range allIssues {
		duration := time.Duration(0)
		for _, w := range iss.WorkLogs {
			if w.StartTime.After(now0) && w.StartTime.Before(now23) {
				duration += w.WorkLogDuration()

			}
		}
		if duration > 0 {
			strDuration := duration.Round(time.Second).String()
			workLogsOfTheDay = append(workLogsOfTheDay, relworklogitem.Relworklogitem{Title: iss.Title, URL: iss.URL, Duration: strDuration})
		}
	}

	return workLogsOfTheDay
}

func InsertIssues(title, url string) error {

	if title == "" || url == "" {
		return fmt.Errorf(" Os campos Titulo e URL nao podem esta Vazios")
	}

	urlExists, err := issuerepository.IsUrlAlreadyExists(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	if urlExists {
		return fmt.Errorf(" Url ja foi cadastrada ")
	}

	hashKey, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(title, " ", url)

	newIssue := issue.Issue{
		UUID:     hashKey.String(),
		URL:      url,
		Title:    title,
		WorkLogs: []worklogs.WorkLogs{},
		Hide:     false,
	}

	issuerepository.PutIssue(newIssue)
	return nil
}

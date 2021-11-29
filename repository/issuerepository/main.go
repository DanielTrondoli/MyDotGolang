package issuerepository

import (
	"fmt"

	"github.com/DanielTrondoli/MyDotGolang/db/dynoclient"
	"github.com/DanielTrondoli/MyDotGolang/models/issue"
	"github.com/DanielTrondoli/MyDotGolang/models/worklogs"
)

//var db dynoclient.DynamoConnection

const TABLE_NAME = "issue"
const KEY = "uuid"
const KEY_WORKLOGS = "workLogs"
const KEY_URL = "url"
const KEY_HIDE = "hide"

func PutIssue(newIssue issue.Issue) error {
	db := dynoclient.GetInstanse()
	return db.Put(TABLE_NAME, newIssue)
}

func DeleteIssue(hashKey string) error {
	db := dynoclient.GetInstanse()
	return db.Remove(TABLE_NAME, KEY, hashKey)
}

func GetIssue(hashKey string) (issue.Issue, error) {
	db := dynoclient.GetInstanse()
	result := issue.Issue{}

	err := db.GetOneBykey(TABLE_NAME, KEY, hashKey, &result)
	if err != nil {
		return issue.Issue{}, err
	}

	return result, nil
}

func IsUrlAlreadyExists(url string) (bool, error) {
	db := dynoclient.GetInstanse()

	count, err := db.Count(TABLE_NAME, KEY_URL, url)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetAllIssues() ([]issue.Issue, error) {
	db := dynoclient.GetInstanse()
	result := []issue.Issue{}
	err := db.GetAll(TABLE_NAME, &result)

	if err != nil {
		return []issue.Issue{}, err
	}

	return result, nil
}

func GetAllNoHideIssues() ([]issue.Issue, error) {
	db := dynoclient.GetInstanse()

	result2 := []issue.Issue{}

	err := db.Get(TABLE_NAME, KEY_HIDE, false, &result2)
	if err != nil {
		return []issue.Issue{}, err
	}

	return result2, nil
}

func DBInitIssue() {
	db := dynoclient.GetInstanse()

	tables, err := db.ListTables()
	if err != nil {
		panic(err.Error())
	}

	for _, s := range tables {
		if s == TABLE_NAME {
			fmt.Println("Table Issue Exists !")
			return
		}
	}

	err = db.CreateTable(TABLE_NAME, issue.Issue{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table Issue Created !")
}

func UpdateWorkLogs(hashKey string, workLogs []worklogs.WorkLogs) {

	db := dynoclient.GetInstanse()

	err := db.Update(TABLE_NAME, KEY, hashKey, KEY_WORKLOGS, workLogs)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Issue Updated !")
}

func HideIssue(hashKey string) {
	db := dynoclient.GetInstanse()

	err := db.Update(TABLE_NAME, KEY, hashKey, KEY_HIDE, true)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Issue Now is Hide !")

}

func ShowIssue(hashKey string) {
	db := dynoclient.GetInstanse()

	err := db.Update(TABLE_NAME, KEY, hashKey, KEY_HIDE, false)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Issue Now is Visible !")
}

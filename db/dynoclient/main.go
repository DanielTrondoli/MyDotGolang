package dynoclient

import (
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type DynamoConnection struct {
	conn     *dynamo.DB
	endpoint string
}

var db *DynamoConnection

var lock = &sync.Mutex{}

func GetInstanse() *DynamoConnection {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			fmt.Println("Creating single instance now.")
			db = &DynamoConnection{}
			db.GetConnection("")
		}
	} else {
		fmt.Println("Using single instance now.")
	}

	return db
}

func (conn *DynamoConnection) GetConnection(endpoint string) error {

	cfg := aws.Config{}

	cfg.Region = aws.String("eu-west-2")
	if endpoint == "" {
		endpoint = "http://0.0.0.0:8000"
	}

	fmt.Println("Conect dynamodb in: ", endpoint)
	cfg.Endpoint = aws.String(endpoint)
	cfg.CredentialsChainVerboseErrors = aws.Bool(true)

	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &cfg)

	conn.conn = db
	conn.endpoint = endpoint

	return nil
}

func (conn DynamoConnection) CreateTable(tableName string, tableStruct interface{}) error {

	db := conn.conn
	if db == nil {
		return fmt.Errorf("conexao nao realizada")
	}

	err := db.CreateTable(tableName, tableStruct).Run()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (conn DynamoConnection) ListTables() ([]string, error) {

	db := conn.conn
	if db == nil {
		return []string{}, fmt.Errorf("conexao nao realizada")
	}

	result, err := db.ListTables().All()
	if err != nil {
		log.Println(err)
		return []string{}, fmt.Errorf(err.Error())
	}

	return result, err
}

func (conn DynamoConnection) Remove(tableName, key string, hashKey interface{}) error {
	db := conn.conn
	if db == nil {
		return fmt.Errorf("conexao nao realizada")
	}

	return db.Table(tableName).Delete(key, hashKey).Run()
}

func (conn DynamoConnection) DeleteAllTables() error {
	db := conn.conn
	if db == nil {
		return fmt.Errorf("conexao nao realizada")
	}

	tn, _ := db.ListTables().All()

	for _, tn := range tn {
		t := db.Table(tn)
		t.DeleteTable().Run()
	}

	return nil
}

func (conn DynamoConnection) Put(tableName string, item interface{}) error {
	db := conn.conn
	if db == nil {
		return fmt.Errorf("conexao nao realizada")
	}

	table := db.Table(tableName)

	err := table.Put(item).Run()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (conn DynamoConnection) Update(tableName, key string, hashKey interface{}, path string, newValue interface{}) error {
	db := conn.conn
	if db == nil {
		return fmt.Errorf("conexao nao realizada")
	}

	table := db.Table(tableName)

	err := table.Update(key, hashKey).Set(path, newValue).Run()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (conn DynamoConnection) Get(tableName, key string, keyValue interface{}, result interface{}) error {
	fmt.Println("GET")
	db := conn.conn
	if db == nil {
		return fmt.Errorf("conexao nao realizada")
	}

	query := db.Table(tableName).Get(key, keyValue)
	err := query.One(result)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (conn DynamoConnection) Count(tableName, key string, keyValue interface{}) (int64, error) {
	db := conn.conn
	if db == nil {
		return -1, fmt.Errorf("conexao nao realizada")
	}

	scan := db.Table(tableName).Scan()
	count, err := scan.Filter(fmt.Sprintf("'%s' = ?", key), keyValue).Count()
	if err != nil {
		return -1, fmt.Errorf(err.Error())
	}

	return count, nil
}

func (conn DynamoConnection) GetAll(tableName string, result interface{}) error {
	db := conn.conn
	if db == nil {
		return fmt.Errorf("conexao nao realizada")
	}

	scan := db.Table(tableName).Scan()
	err := scan.All(result)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

// func InitDB(endpoint string) {
// 	conn.GetConnection(endpoint)
// 	conn.DeleteAllTables()
// 	conn.CreateTable("issue", issue.Issue{})
// 	conn.ListTables()

// 	newIssue := issue.Issue{
// 		URL:      "http://dnaile1231231",
// 		Title:    "Daniel123123",
// 		WorkLogs: []worklogs.WorkLogs{},
// 	}

// 	newIssue2 := issue.Issue{
// 		URL:      "http://dnaile1231232",
// 		Title:    "Daniel123123",
// 		WorkLogs: []worklogs.WorkLogs{},
// 	}

// 	newIssue3 := issue.Issue{
// 		URL:      "http://dnaile1231233",
// 		Title:    "Daniel123123",
// 		WorkLogs: []worklogs.WorkLogs{},
// 	}

// 	conn.Put("issue", newIssue)
// 	conn.Put("issue", newIssue2)
// 	conn.Put("issue", newIssue3)

// 	newIssue.Title = " $$$$$$$$$$$ "
// 	conn.Update("issue", "url", newIssue.URL, "title", newIssue.Title)
// 	fmt.Println(newIssue)

// 	result := []issue.Issue{}
// 	conn.GetAll("issue", "url", newIssue.URL, &result)
// 	fmt.Println(result)

// 	result = []issue.Issue{}
// 	conn.GetAll("issue", "url", newIssue.URL, &result)
// 	fmt.Println(result)
// }

// func main() {

// 	c := DynamoConnection{}
// 	c.InitDB("")

// }
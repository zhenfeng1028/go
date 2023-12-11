package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hootsuite/healthchecks"
	"xorm.io/xorm"
)

func main() {
	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:lzf123@(127.0.0.1:4000)/lzf?charset=utf8&parseTime=true&loc=Local") // lzf是数据库实例名称
	if err != nil {
		log.Fatal("init engine failed:", err)
	}

	db := healthchecks.StatusEndpoint{
		Name:          "The DB",
		Slug:          "db",
		Type:          "internal",
		IsTraversable: false,
		StatusCheck: MysqlStatusChecker{
			engine: engine,
		},
		TraverseCheck: nil,
	}

	statusEndpoints := []healthchecks.StatusEndpoint{db}

	// Set the path for the about and version files
	aboutFilePath := "/Users/lizhenfeng/GitHub/go/open/healthcheck/about.json"
	versionFilePath := "/Users/lizhenfeng/GitHub/go/open/healthcheck/version.txt"

	// Set up any service injected customData for /status/about response
	// Values can be any valid JSON conversion and will override values set in about.json
	customData := make(map[string]interface{})

	// String value
	customData["db"] = "mysql"

	// Register all the "/status/..." requests to use our health checking framework
	http.Handle("/status/", healthchecks.Handler(statusEndpoints, aboutFilePath, versionFilePath, customData))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type MysqlStatusChecker struct {
	engine *xorm.Engine
}

func (m MysqlStatusChecker) CheckStatus(name string) healthchecks.StatusList {
	err := m.engine.Ping()

	// Set a default response
	s := healthchecks.Status{
		Description: name,
		Result:      healthchecks.OK,
		Details:     "",
	}

	// Handle any errors that Ping() function returned
	if err != nil {
		s = healthchecks.Status{
			Description: name,
			Result:      healthchecks.CRITICAL,
			Details:     err.Error(),
		}
	}

	// Return our response
	return healthchecks.StatusList{StatusList: []healthchecks.Status{s}}
}

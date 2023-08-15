package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/consul/api"
)

const (
	serviceName = "WZ_Sp"
)

type Service struct {
	ServiceId    string
	ServiceName  string
	consulClient *api.Client
}

type ServiceInfo struct {
	ServiceID   string
	ServiceName string
	IP          string
	Port        int
	Status      string
	PackageName string
	Load        int
	Times       int   //被调用次数
	Timestamp   int64 //load updated ts
}

func main() {
	srv := &Service{}
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	srv.consulClient = client

	dcNames, err := srv.consulClient.Catalog().Datacenters()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dcNames:%v\n\n", dcNames)

	services, _, err := srv.consulClient.Catalog().Services(&api.QueryOptions{})
	if err != nil {
		panic(err)
	}
	bs, _ := json.Marshal(services)
	fmt.Printf("services:%s\n\n", string(bs))

	// serviceEntries, _, err := srv.consulClient.Health().Service(serviceName, "", false, &api.QueryOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// bs2, _ := json.Marshal(serviceEntries)
	// fmt.Printf("serviceEntries:%s\n\n", string(bs2))

	// items := make(map[string]*ServiceInfo)
	// for _, entry := range serviceEntries {
	// 	if serviceName != entry.Service.Service {
	// 		continue
	// 	}
	// 	for _, health := range entry.Checks {
	// 		if health.ServiceName != serviceName {
	// 			continue
	// 		}
	// 		if health.Status != "passing" {
	// 			continue
	// 		}
	// 		sInfo := new(ServiceInfo)
	// 		sInfo.ServiceName = health.ServiceName
	// 		sInfo.ServiceID = health.ServiceID
	// 		if len(health.ServiceTags) > 0 {
	// 			sInfo.PackageName = health.ServiceTags[0]
	// 		}
	// 		strIP, nPort := getBaseInfoWithServiceID(sInfo.ServiceName, sInfo.ServiceID)
	// 		if strIP == "" || nPort == 0 {
	// 			continue
	// 		}
	// 		sInfo.IP = strIP
	// 		sInfo.Port = nPort
	// 		sInfo.Status = health.Status
	// 		items[sInfo.ServiceID] = sInfo
	// 	}
	// }
	// bs3, _ := json.Marshal(items)
	// fmt.Printf("items:%s\n\n", string(bs3))

	healthChecks, _, err := srv.consulClient.Health().State("any", &api.QueryOptions{WaitTime: 3 * time.Second})
	if err != nil {
		panic(err)
	}
	bs4, _ := json.Marshal(healthChecks)
	fmt.Printf("healthChecks:%s\n\n", string(bs4))

	err = srv.consulClient.Agent().ServiceDeregister("WZ_Sp-172.24.193.13-29065")
	if err != nil {
		panic(err)
	}
}

func getBaseInfoWithServiceID(strServiceName, strServiceId string) (string, int) {
	listStr := strings.Split(strServiceId, "-")
	if len(listStr) == 3 && listStr[0] == strServiceName && listStr[1] != "" && listStr[2] != "" {
		nPort, _ := strconv.Atoi(listStr[2])
		if nPort > 0 {
			return listStr[1], nPort
		}
	}

	return "", 0
}

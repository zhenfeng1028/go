package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type FreePassResp struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []*FreePassInfo `json:"data"`
}

type FreePassInfo struct {
	CreatedBy   int        `json:"createdBy"`
	UpdatedBy   int        `json:"updatedBy"`
	UpdatedTime string     `json:"updatedTime"`
	CreatedTime string     `json:"createdTime"`
	ObjStatus   *ObjStatus `json:"objStatus"`
	Id          int        `json:"id"`
	PassType    int        `json:"passType"`
	PassTime    string     `json:"passTime"`
}

type ObjStatus struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Text string `json:"text"`
}

const (
	Minute_Format   = "200601021504"
	Calendar_Format = "2006-01-02+15:04:05"
)

func main() {
	minute := "202112272350"
	minuteTime, _ := time.Parse(Minute_Format, minute)
	fmt.Println(minuteTime)
	addTime := minuteTime.Add(time.Minute * 10)
	fmt.Println(addTime)

	freePass, _ := getFreePass(minuteTime)
	fmt.Println(freePass)
}

func getFreePass(minute time.Time) (bool, error) {
	passTime := minute.Format(Calendar_Format)
	url := fmt.Sprintf("http://100.100.64.235:32000/v1/facilities-manage-service/freepasscalendar/list?passTime=%s&passTimeType=3", passTime)
	fmt.Println(url)
	freePassResp, err := http.Get(url)
	if err != nil {
		fmt.Printf("get freePassCalendar error: %v, %v\n", url, err)
		return false, err
	}
	defer freePassResp.Body.Close()

	freePassData, err := ioutil.ReadAll(freePassResp.Body)
	if err != nil {
		fmt.Printf("read body err: %v", err)
		return false, err
	}
	if freePassResp.StatusCode/100 != 2 {
		fmt.Printf("get freePassCalendar error: %v\n", string(freePassData))
		return false, fmt.Errorf("get freePassCalendar error: %v", string(freePassData))
	}

	var freePassResponse FreePassResp
	err = json.Unmarshal(freePassData, &freePassResponse)
	if err != nil {
		fmt.Printf("bad road response : %v, err : %v\n", string(freePassData), err)
		return false, err
	}

	if len(freePassResponse.Data) != 1 {
		fmt.Println("length of freePassResponse Data not 1")
	}
	if freePassResponse.Data[0].PassType == 1 {
		return true, nil
	}
	return false, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func main() {
	u1 := &MyUser{1, "Ken", time.Now()}
	_ = json.NewEncoder(os.Stdout).Encode(u1)

	u2 := &MyUser{}
	_ = json.NewDecoder(os.Stdin).Decode(u2)
	fmt.Println(u2)
}

type Alias MyUser

func (u *MyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		LastSeen: u.LastSeen.Unix(),
		Alias:    (*Alias)(u),
	})
}

func (u *MyUser) UnmarshalJSON(data []byte) error {
	aux := &struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u.LastSeen = time.Unix(aux.LastSeen, 0)
	return nil
}

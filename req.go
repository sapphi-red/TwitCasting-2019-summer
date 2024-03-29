package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var TOKEN string

type getStruct struct {
	ID       string `json:"id"`
	Question string `json:"question"`
}

func Init() {
	f, err := os.Open("token.txt")
	if err != nil {
		panic("where is token?")
	}

	b, err := ioutil.ReadAll(f)
	TOKEN = strings.TrimSpace(string(b))

	f.Close()
}

func Get(level string) *getStruct {
	req, err := http.NewRequest("GET", "https://apiv2.twitcasting.tv/internships/2019/games?level="+level, nil)
	req.Header.Add("Authorization", "Bearer "+TOKEN)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	var struc getStruct
	if err := json.NewDecoder(res.Body).Decode(&struc); err != nil {
		panic(err)
	}

	fmt.Println("GET: " + struc.ID + " ||| " + struc.Question)

	res.Body.Close()
	return &struc
}

func Post(id string, answer string) {
	req, err := http.NewRequest("POST", "https://apiv2.twitcasting.tv/internships/2019/games/"+id, bytes.NewBuffer([]byte(`{"answer":"`+answer+`"}`)))

	req.Header.Add("Authorization", "Bearer "+TOKEN)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("POST: " + string(body))
	res.Body.Close()
}

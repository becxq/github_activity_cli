package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Activity struct {
	EventType string `json:"type"`
	Repo      Repo   `json:"repo"`
}

func (a Activity) convert(c int) string {
	var baseEvent string

	switch a.EventType {
	case "CreateEvent":
		baseEvent = fmt.Sprintf("Created repository named %s", a.Repo.Name)
	case "PushEvent":
		baseEvent = fmt.Sprintf("Pushed to %s", a.Repo.Name)
	case "PullRequestEvent":
		baseEvent = fmt.Sprintf("Created pull request in %s", a.Repo.Name)
	case "IssuesEvent":
		baseEvent = fmt.Sprintf("Opened issue in %s", a.Repo.Name)
	case "WatchEvent":
		baseEvent = fmt.Sprintf("Starred %s", a.Repo.Name)
	case "ForkEvent":
		baseEvent = fmt.Sprintf("Forked %s", a.Repo.Name)
	case "DeleteEvent":
		baseEvent = fmt.Sprintf("Deleted from %s", a.Repo.Name)
	case "ReleaseEvent":
		baseEvent = fmt.Sprintf("Released in %s", a.Repo.Name)
	default:
		baseEvent = fmt.Sprintf("%s in %s", a.EventType, a.Repo.Name)
	}

	if c > 1 {
		return fmt.Sprintf("%s (%d times)", baseEvent, c)
	}
	return baseEvent
}

func main() {
	args := os.Args[1:]
	var username string
	var count int

	if len(args) == 0 {
		fmt.Println("To use github activity checker, you need to put github username as a first argument\nOptional: Number of string")
		return
	} else if len(args) == 1 {
		if reflect.TypeOf(args[0]) != reflect.TypeOf(".") {
			fmt.Println("To use github activity checker, you need to put github username as a first argument\nOptional: Number of string")
			fmt.Println("Wrong data format")
			return
		}
		username = args[0]
		count = 5
	} else if len(args) == 2 {
		if _, err := strconv.Atoi(args[1]); reflect.TypeOf(args[0]) != reflect.TypeOf(".") && err != nil {
			fmt.Println("To use github activity checker, you need to put github username as a first argument\nOptional: Number of string")
			fmt.Println("Wrong data format")
			return
		}
		username = args[0]
		count, _ = strconv.Atoi(args[1])
	}

	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events/public", username))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("User not found. Try again")
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var activity []Activity

	err = json.Unmarshal(body, &activity)
	if err != nil {
		log.Fatal(err)
	}

	result := []string{}
	c := 1
	for i := 0; i < count+1; i++ {
		if activity[i].Repo != activity[i+1].Repo || activity[i].EventType != activity[i+1].EventType {
			result = append(result, activity[i].convert(c))
			c = 0
		}
		c++
	}

	for _, r := range result {
		fmt.Println(r)
	}
}

package fancy

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Job struct {
	By        string `json:"by"`
	Id        int    `json:"id"`
	Score     int    `json:"score"`
	Text      string `json:"text"`
	Time      int    `json:"time"`
	JsonTitle string `json:"title"`
	Type      string `json:"type"`
	Url       string `json:"url"`
}

func GetJobIds() []int {
	url := "https://hacker-news.firebaseio.com/v0/jobstories.json?print=pretty"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	resp.Body.Close()
	return ids
}

func FetchJob(id int) (*Job, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", id)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var job *Job
	if err := json.NewDecoder(res.Body).Decode(&job); err != nil {
		return nil, err
	}
	res.Body.Close()
	return job, nil
}

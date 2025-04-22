package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var repoarray = [2]string{"https://api.github.com/repos/prometheus/prometheus/releases/latest","https://api.github.com/repos/grafana/grafana/releases/latest" }

type Response struct {
	TagName     string    `json:"tag_name"`
	PublishedAt time.Time `json:"published_at"`
	URL         string 	  `json:"html_url"`
}

func main() {
	for _, repourl := range repoarray {

		resp, err := http.Get(repourl)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()
	
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
	
		var result Response
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}
	
		fmt.Println("Latest Prometheus release:")
		fmt.Printf("Version: %s\n", result.TagName)
		fmt.Printf("Published at: %s\n", result.PublishedAt.Format("2003-07-23"))
	
		fmt.Println("\nPretty printed JSON response:")
		fmt.Println(PrettyPrint(result))
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
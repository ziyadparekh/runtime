package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ziyadparekh/runtime/logger"
)

type Event struct {
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequest struct {
	Url       string `json:"url"`
	State     string `json:"state"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
	Head      Commit `json:"head"`
}

type User struct {
	Login  string `json:"login"`
	Avatar string `json:"avatar_url"`
	Url    string `json:"url"`
}

type Commit struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Sha   string `json:"sha"`
	URL   string `json:"url"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hi there, welcome to Runtime\n")
}

func Payload(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var event Event
	if err := json.Unmarshal(body, &event); err != nil {
		logger.Fatalf("%v", err)
	}
	out, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		logger.Fatalf("%v", err)
	}
	fmt.Println(string(out))
	// var out bytes.Buffer
	// json.Indent(&out, body, "", "\t")
	// out.WriteTo(os.Stdout)
}

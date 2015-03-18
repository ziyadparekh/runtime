package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ziyadparekh/runtime/logger"
	"github.com/ziyadparekh/runtime/repository"
	"github.com/ziyadparekh/runtime/structs"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hi there, welcome to Runtime\n")
}

func PullRequestHook(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var event structs.Event
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

func BranchHook(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	e := r.Header.Get("X-Github-Event")

	defer r.Body.Close()

	var ref structs.Branch
	if err := json.Unmarshal(body, &ref); err != nil {
		logger.Fatalf("%v", err)
	}
	ref.Event = e
	out, err := json.MarshalIndent(ref, "", "  ")
	if err != nil {
		logger.Fatalf("%v", err)
	}
	fmt.Println(string(out))
	repository.HandleBranchHook(&ref)
	// var out bytes.Buffer
	// json.Indent(&out, body, "", "\t")
	// out.WriteTo(os.Stdout)
}

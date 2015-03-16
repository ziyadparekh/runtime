package repository

import (
	"encoding/json"
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ziyadparekh/runtime/copy"
	"github.com/ziyadparekh/runtime/filecontrol"
	"github.com/ziyadparekh/runtime/logger"
)

type Repository struct {
	Name         string   `json:name`
	Url          string   `json:url`
	Email        string   `json:email`
	Directory    string   `json:directory`
	Language     string   `json:language`
	Dependencies []string `json:dependencies`
}

func NewRuntime() []byte {
	// get the name of the project
	n := getUserInput(copy.Name)
	// get the url of the git repo
	u := getUserInput(copy.Repository)
	// get the email address
	e := getUserInput(copy.Email)
	// get the directory
	d := getUserInput(copy.Directory)
	// get the language
	l := getUserInput(copy.Language)

	repo := &Repository{
		Name:      n,
		Url:       u,
		Email:     e,
		Directory: d,
		Language:  l,
	}

	fmt.Println("====> Saving your Runtime configuration")
	r, err := json.MarshalIndent(repo, "", "  ")
	if err != nil {
		logger.Fatalf("%v", err)
	}
	if err := filecontrol.WriteFile(copy.Runtime, r); err != nil {
		logger.Fatalf("%v", err)
	}
	fmt.Println("====> File runtime.json successfully written")
	return r
}

func SetupStructure() bool {
	var repo *Repository
	file, err := filecontrol.CheckAndReadFile(copy.Runtime)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	if err := json.Unmarshal(file, &repo); err != nil {
		logger.Fatalf("%v", err)
	}
	if repo.Name == "" || repo.Directory == "" {
		logger.Fatalf("%s", "Your runtime.json file is missing information")
	}
	fmt.Println("====> Attempting to create your project structure")
	if err := filecontrol.CreateDirp(repo.Directory, repo.Name); err != nil {
		logger.Fatalf("%v", err)
	}
	fmt.Println("====> Succesfully created directory structure")
	// TODO:: clone repo in directory
	fmt.Println("====> Attempting to clone repo into directory")

	return true

}

func SaveDependencies(args cli.Args) (bool, error) {
	var repo *Repository
	file, err := filecontrol.CheckAndReadFile(copy.Runtime)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	if err := json.Unmarshal(file, &repo); err != nil {
		logger.Fatalf("%v", err)
	}
	if repo.Dependencies == nil {
		repo.Dependencies = make([]string, len(args)-1)
	}
	for _, dep := range args {
		if !inDepArray(dep, repo.Dependencies) && dep != "" {
			fmt.Println("====> " + dep)
			repo.Dependencies = append(repo.Dependencies, dep)
		}
	}
	d, derr := json.MarshalIndent(repo, "", "  ")
	if derr != nil {
		logger.Fatalf("%v", derr)
	}
	if err := filecontrol.WriteFile(copy.Runtime, d); err != nil {
		logger.Fatalf("%v", err)
	}
	return true, nil
}

func getUserInput(msg string) string {
	m, err := filecontrol.GetUserInput(msg)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	return m
}

func inDepArray(dep string, depArray []string) bool {
	for _, d := range depArray {
		if d == dep {
			return true
		}
	}
	return false
}

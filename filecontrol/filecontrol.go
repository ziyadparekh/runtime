package filecontrol

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ziyadparekh/runtime/copy"
)

func GetUserInput(msg string) (string, error) {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("====> " + msg)
	usr, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(usr), nil
}

func WriteFile(name string, data []byte) error {
	err := ioutil.WriteFile(name, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

func CheckAndReadFile(name string) ([]byte, error) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return nil, err
	}
	f, ferr := ioutil.ReadFile(name)
	if ferr != nil {
		return nil, ferr
	}
	return f, nil
}

func CreateGenericDir(path string) error {
	fmt.Println("====> Creating directory --> " + path)
	if err := os.MkdirAll(path, 0777); err != nil {
		return err
	}
	return nil
}

func CreateDirp(path, name string) error {
	apppath := fmt.Sprintf("%s%s%s%s", copy.PATH_APPS, path, name, copy.DEFAULT_BRANCH)
	gitpath := fmt.Sprintf("%s%s%s%s", copy.PATH_GIT, path, name, copy.DEFAULT_BRANCH)

	fmt.Println("====> Creating folders for app --> " + apppath)
	if err := CreateGenericDir(apppath); err != nil {
		return err
	}
	fmt.Println("====> Creating folders for git repo --> " + gitpath)
	if err := CreateGenericDir(gitpath); err != nil {
		return err
	}
	return nil
}

func CreateBranchDir(branch, path string) error {
	branchAppPath := fmt.Sprintf("%s%s%s", copy.PATH_APPS, path, branch)
	branchGitPath := fmt.Sprintf("%s%s%s", copy.PATH_GIT, path, branch)
	fmt.Println("====> Creating folders for app branch --> " + branchAppPath)
	if err := CreateGenericDir(branchAppPath); err != nil {
		return err
	}
	fmt.Println("====> Creating folders for git branch --> " + branchGitPath)
	if err := CreateGenericDir(branchGitPath); err != nil {
		return err
	}
	return nil
}

func DeleteBranchDir(branch, path string) error {
	branchAppPath := fmt.Sprintf("%s%s%s", copy.PATH_APPS, path, branch)
	branchGitPath := fmt.Sprintf("%s%s%s", copy.PATH_GIT, path, branch)
	fmt.Println("====> Deleting folders for app branch --> " + branchAppPath)
	if err := os.RemoveAll(branchAppPath); err != nil {
		return err
	}
	fmt.Println("====> Deleting folders for git branch --> " + branchAppPath)
	if err := os.RemoveAll(branchGitPath); err != nil {
		return err
	}
	return nil
}

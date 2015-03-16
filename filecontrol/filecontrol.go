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

func CreateDirp(path, name string) error {
	apppath := fmt.Sprintf("%s%s%s", copy.PATH_APPS, path, name)
	gitpath := fmt.Sprintf("%s%s%s", copy.PATH_GIT, path, name)
	fmt.Println("====> Creating folders for app --> " + apppath)
	if err := os.MkdirAll(apppath, 0777); err != nil {
		return err
	}
	fmt.Println("====> Creating folders for git repo --> " + gitpath)
	if err := os.MkdirAll(gitpath, 0777); err != nil {
		return err
	}
	return nil
}

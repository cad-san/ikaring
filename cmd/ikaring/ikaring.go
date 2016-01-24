package main

import (
	"github.com/cad-san/ikaring"

	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/bgentry/speakeasy"
)

func getCacheFile() (string, error) {
	me, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(me.HomeDir, ".ikaring.session"), nil
}

func readSession(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(buff)), nil
}

func writeSession(path string, session string) error {
	return ioutil.WriteFile(path, []byte(session), 600)
}

func getAccount(r io.Reader) (string, string, error) {
	scanner := bufio.NewScanner(r)
	for {
		fmt.Print("User: ")
		if scanner.Scan() {
			break
		}
	}
	username := scanner.Text()
	password, err := speakeasy.Ask("Password: ")
	return username, password, err
}

func main() {
	client, err := ikaring.CreateClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	path, err := getCacheFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	session, err := readSession(path)
	if err == nil && len(session) > 0 {
		client.SetSession(session)
	} else {
		username, password, err := getAccount(os.Stdin)
		session, err = client.Login(username, password)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if len(session) <= 0 {
		fmt.Println("ログインできませんでした")
		return
	}

	writeSession(path, session)

	info, err := client.GetStageInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	if info.FesSchedules != nil {
		for _, s := range *info.FesSchedules {
			fmt.Printf("%v\n", s)
		}
	}

	if info.Schedules != nil {
		for _, s := range *info.Schedules {
			fmt.Printf("%v\n", s)
		}
	}
}

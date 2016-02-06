package main

import (
	"github.com/cad-san/ikaring"

	"bufio"
	"errors"
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

func login(client *ikaring.IkaClient) error {
	path, err := getCacheFile()
	if err != nil {
		return err
	}

	session, err := readSession(path)
	if err == nil && len(session) > 0 {
		client.SetSession(session)
		return nil // already authorized
	}

	username, password, err := getAccount(os.Stdin)
	session, err = client.Login(username, password)
	if err != nil {
		return err
	}

	if len(session) <= 0 {
		return errors.New("login failure")
	}
	writeSession(path, session)
	return nil
}

func stage(client *ikaring.IkaClient) error {
	info, err := client.GetStageInfo()
	if err != nil {
		return err
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
	return nil
}

func ranking(client *ikaring.IkaClient) error {
	info, err := client.GetRanking()
	if err != nil {
		return err
	}

	if len(info.Regular) > 0 {
		fmt.Println("レギュラーマッチ")
		for _, p := range info.Regular {
			fmt.Printf("\t[%d] %3d %s (%s)\n", p.Rank, p.Score, p.Name, p.Weapon)
		}
	}

	if len(info.Gachi) > 0 {
		fmt.Println("ガチマッチ")
		for _, p := range info.Gachi {
			fmt.Printf("\t[%d] %3d %s (%s)\n", p.Rank, p.Score, p.Name, p.Weapon)
		}
	}
	return nil
}

func main() {
	client, err := ikaring.CreateClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = login(client); err != nil {
		fmt.Println(err)
		return
	}

	if err = stage(client); err != nil {
		fmt.Println(err)
		return
	}
}

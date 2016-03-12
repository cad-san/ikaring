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
	"github.com/jessevdk/go-flags"
	"github.com/mattn/go-colorable"
)

type option struct {
	Stage  stageCmd  `command:"stage"  description:"display stage schedule"`
	Rank   rankCmd   `command:"rank"   description:"display ranking with friends"`
	Friend friendCmd `command:"friend" description:"display friend list"`
}
type stageCmd struct{}
type rankCmd struct{}
type friendCmd struct{}

const (
	Red   = "\x1b[31;1m"
	Green = "\x1b[32;1m"
	White = "\x1b[37;1m"
	End   = "\x1b[0m"
)

var (
	stdout = colorable.NewColorableStdout()
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

func (c *stageCmd) Execute(args []string) error {
	client, err := ikaring.CreateClient()
	if err != nil {
		return err
	}

	if err = login(client); err != nil {
		return err
	}

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

func (c *rankCmd) Execute(args []string) error {
	client, err := ikaring.CreateClient()
	if err != nil {
		return err
	}

	if err = login(client); err != nil {
		return err
	}

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

	if len(info.Festival) > 0 {
		fmt.Println("フェス")
		for _, p := range info.Festival {
			top100 := ""
			if p.Top100 {
				top100 = Red + "百ケツ!\t" + End
			} else {
				top100 = "\t"
			}
			fmt.Fprintf(stdout, "%s[%d] %3d %s(%s)\n", top100, p.Rank, p.Score, p.Name, p.Weapon)
		}
	}

	return nil
}

func (c *friendCmd) Execute(args []string) error {
	client, err := ikaring.CreateClient()
	if err != nil {
		return err
	}

	if err = login(client); err != nil {
		return err
	}

	list, err := client.GetFriendList()
	if err != nil {
		return err
	}

	if len(list) == 0 {
		fmt.Println("フレンドはオフラインです")
	}

	for _, f := range list {
		fmt.Printf("%s\n", f.Name)
		fmt.Printf("\t%s\n", f.Mode)
	}
	return nil
}

func main() {
	var opts option
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "ikaring"
	parser.SubcommandsOptional = true

	_, err := parser.Parse()

	if len(os.Args) == 1 {
		parser.WriteHelp(os.Stdout)
		return
	}
	if err != nil {
		os.Exit(1)
	}
}

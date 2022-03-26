package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitor = 3
const delay = 5

func main() {
	registerLog("false", false)
	showIntro()
	for {
		showMenu()
		// if command == 1 {
		// 	fmt.Println("monitoring")
		// } else if command == 2 {
		// 	fmt.Println("logs")
		// } else if command == 0 {
		// 	fmt.Println("Exiting...")
		// } else {
		// 	fmt.Println("Command Unavailable")
		// }

		command := getCommand()

		switch command {
		case 1:
			monitoring()
		case 2:
			printLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Command Unavailable")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Vítor"
	version := 1.1
	fmt.Println("Hello", name)
	fmt.Println("This program version is", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func getCommand() int {

	var command int
	fmt.Scan(&command)
	fmt.Println("Choosen command is", command)

	return command
}

func monitoring() {
	fmt.Println("monitoring")
	sites := getSitesFromFile()

	for i := 0; i < monitor; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i)
			testSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "successfully loaded")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "unavailable. Status Code", res.StatusCode)
		registerLog(site, false)
	}
}

func getSitesFromFile() []string {

	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)

		sites = append(sites, line)
		if err == io.EOF {
			break
		}

	}
	file.Close()
	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()

}

func printLogs() {
	fmt.Println("logs")
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(file))
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	introduction()
	for {
		habitat()
		toothPaste := tooth()

		switch toothPaste {
		case 1:
			initialize1()
		case 2:
			fmt.Println("showing logs..")
		case 0:
			fmt.Println("leaving app..")
			os.Exit(0)
		default:
			fmt.Println("unknown command..")
			os.Exit(-1)
		}
	}
}

func introduction() {
	name := "user"
	versao := 0.9
	fmt.Println("helcome", name)
	fmt.Println("app version", versao)
}

func tooth() int {
	var toothPaste int
	fmt.Scan(&toothPaste)
	fmt.Println("the command was", toothPaste)

	return toothPaste
}

func habitat() {
	fmt.Println("1- start monitoring")
	fmt.Println("2- show logs")
	fmt.Println("0- exit")
}

func initialize1() {
	fmt.Println("monitoring..")
	sites := []string{"https://roblox.com/",
		"https://youtube.com.br/",
		"https://google.com.br",
		"https://github.com.br"}

	//fmt.Println(sites)

	for i, site := range sites {
		fmt.Println("testing site..", i, ":", site)
		siteTesting(site)
	}

	site := "https://roblox.com/"
	resp, _ := http.Get(site)
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("website", site, "working")
	} else {
		fmt.Println("not working",
			resp.StatusCode)
	}
}

func siteTesting(site string) {
	resp, _ := http.Get(site)
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("website", site, "working")
	} else {
		fmt.Println("not working",
			resp.StatusCode)
	}
}

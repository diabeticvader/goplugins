package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cbelk/vars/pkg/varsapi"
)

func main() {
	config := flag.String("config", "/etc/vars/vars.conf", "The path to the vars configuration file")
	flag.Parse()

	// Read in config and connect to DB
	fmt.Println("[+] Reading vars config file ...")
	err := varsapi.ReadConfig(*config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[+] Connecting to postgresql ...")
	_, err = varsapi.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	vulns, err := varsapi.GetVulnerabilities()
	if !varsapi.IsNilErr(err) {
		fmt.Print(err)
	}

	count := make(map[int]int)

	for _, v := range vulns {
		year, _, _ := v.Dates.Initiated.Date()
		_, ok := count[year]
		if ok {
			count[year]++
		} else {
			count[year] = 1
		}
	}
	fmt.Println(count)

}

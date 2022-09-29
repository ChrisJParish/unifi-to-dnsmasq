package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/chrisparish/unifi"
)

func main() {
	c := &unifi.Config{
		User: "ChrisParish",
		Pass: "U4Fc2FzgxaN8Uz",
		URL:  "https://controller.cjparish.uk",
		// Log with log.Printf or make your own interface that accepts (msg, fmt)
		ErrorLog: log.Printf,
		DebugLog: log.Printf,
	}
	uni, err := unifi.NewUnifi(c)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	sites, err := uni.GetSites()
	if err != nil {
		log.Fatalln("Error:", err)
	}

	devices, err := uni.GetDevices(sites)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	users, err := uni.GetUsers(sites, 87600)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	log.Println(len(sites), "Unifi Sites Found: ", sites)

	log.Println(len(users), "Users Found:")

	log.Println(len(devices.USWs), "Unifi Switches Found")
	log.Println(len(devices.USGs), "Unifi Gateways Found")

	log.Println(len(devices.UAPs), "Unifi Wireless APs Found:")

	m1 := regexp.MustCompile(`\s|'|:|,|_|’`)

	os.Remove("dnsdata.conf")

	fi, err := os.Create("dnsdata.conf")
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	for _, users := range users {
		if users.UseFixedIp.Val == false {
			continue
		}
		name := strings.ToLower(users.Name)
		if len(users.Note) != 0 {
			name = users.Note
		} else {
			name = m1.ReplaceAllString(name, "-")
		}

		fi.WriteString(fmt.Sprintln(users.FixedIp, name))

	}
}

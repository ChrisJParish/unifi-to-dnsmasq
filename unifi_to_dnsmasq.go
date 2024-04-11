package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/unpoller/unifi"
)

func main() {

	user := flag.String("u", "", "User name")
	pass := flag.String("p", "", "Password")
	url := flag.String("h", "", "Host Url")

	flag.Parse()

	if *user == "" || *pass == "" || *url == "" {
		fmt.Println("You haven't passed the required variables")
		os.Exit(1)
	}

	c := &unifi.Config{
		User: *user,
		Pass: *pass,
		URL:  *url,
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

	os.Remove("dnsmasq-hosts")

	fi, err := os.Create("dnsmasq-hosts")
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

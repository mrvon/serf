package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func payload() string {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return ""
	}
	return strconv.Quote(string(buf))
}

func main() {
	f, err := os.OpenFile("/tmp/serf.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	serfEvent := os.Getenv("SERF_EVENT")
	selfName := os.Getenv("SERF_SELF_NAME")
	switch serfEvent {
	case "user":
		userEvent := os.Getenv("SERF_USER_EVENT")
		log.Printf("%s> %s | %s | %s", selfName, serfEvent, userEvent, payload())
	case "query":
		userQuery := os.Getenv("SERF_QUERY_NAME")
		log.Printf("%s> %s | %s | %s", selfName, serfEvent, userQuery, payload())
	default: // other event
		log.Printf("%s> %s | %s", selfName, serfEvent, payload())
	}
}

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
	userEvent := os.Getenv("USER_EVENT")
	userQuery := os.Getenv("SERF_QUERY_NAME")
	if len(serfEvent) > 0 {
		selfName := os.Getenv("SERF_SELF_NAME")
		log.Printf("SerfEvent | %s | %s | %s", selfName, serfEvent, payload())
	} else if len(userEvent) > 0 {
		log.Printf("UserEvent | %s | %s", userEvent, payload())
	} else if len(userQuery) > 0 {
		log.Printf("UserQuery | %s | %s", userQuery, payload())
	} else {
		log.Printf("UNKNOWN EVENT")
	}
}

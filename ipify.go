package main

import (
	"github.com/rdegges/go-ipify"

	log "github.com/sirupsen/logrus"
)

func getCurrentIP() string {
	ip, err := ipify.GetIp()
	if err != nil {
		log.Error("Couldn't get IP address: ", err)
	} else {
		log.Info("IP address is: ", ip)
	}

	return ip
}

package main

import (
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/kardianos/osext"
	"github.com/kardianos/service"
	config "github.com/spf13/viper"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	config.AddConfigPath(folderPath)
	config.SetConfigName("config")
	config.SetDefault("CheckInterval", 3600)
	config.ReadInConfig()

	go p.run()
	return nil
}

func (p *program) run() {
	apiKey := config.GetString("APIKey")
	dnsName := config.GetString("DNSName")
	currentIP := getCurrentIP()
	previousRecord := getDNSRecord(apiKey, dnsName)

	for {
		if currentIP == previousRecord {
			log.Info("DNS Unchanged")
		} else {
			log.Info("Updating DNS")
			//removeDNS(apiKey, previousRecord)
			//addDNS(apiKey, dnsName, currentIP)
		}

		time.Sleep(config.GetDuration("CheckInterval") * time.Second)
	}
}

func (p *program) Stop(s service.Service) error {
	log.Info("Shutting down")
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "DhDNS",
		DisplayName: "Dreamhost Dynamic DNS Service",
		Description: "Dreamhost Dynamic DNS Service",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

package main

import (
	"path/filepath"
	"time"

	"github.com/kardianos/osext"
	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	execPath, _ := osext.ExecutableFolder()
	folderPath := filepath.Join(execPath, "config")

	config.AddConfigPath(folderPath)
	config.SetEnvPrefix("dh")
	config.AllowEmptyEnv(true)
	config.AutomaticEnv()

	config.SetDefault("CheckInterval", 3600)

	err := config.ReadInConfig()
	if err != nil {
		switch err.(type) {
		case config.ConfigFileNotFoundError:
			log.Warnf("No Config file found, loaded config from Environment - Default path %s", folderPath)
		default:
			log.Fatalf("Error when Fetching Configuration - %s", err)
		}
	}

	go p.run()
	return nil
}

func (p *program) run() {
	apiKey := config.GetString("APIKey")
	hostFQDN := config.GetString("HostFQDN")

	for {
		updateDNS(apiKey, hostFQDN)

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

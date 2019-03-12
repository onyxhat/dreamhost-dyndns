package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func (dh dhRequest) getResponse() dhResponse {
	url := dh.URL + "?cmd=" + dh.CMD + "&format=" + dh.Format + "&key=" + dh.APIKey + dh.Args
	res, err := http.Get(url)
	jsonResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	records := dhResponse{}
	json.Unmarshal(jsonResponse, &records)

	return records
}

func getDNSRecord(apiKey string, recordValue string) dhDNSRecord {
	dh := dhRequest{URL: "https://api.dreamhost.com/", APIKey: apiKey, CMD: "dns-list_records", Format: "json"}
	response := dh.getResponse()

	output := dhDNSRecord{}

	for _, v := range response.Data {
		if v.Record == recordValue && v.Type == "A" && v.Editable == "1" {
			output = v
		}
	}

	return output
}

func delDNSRecord(apiKey string, dns dhDNSRecord) {
	args := fmt.Sprintf("&record=%s&type=%s&value=%s", dns.Record, dns.Type, dns.Value)
	dh := dhRequest{URL: "https://api.dreamhost.com/", APIKey: apiKey, CMD: "dns-remove_record", Format: "json", Args: args}
	response := dh.getResponse()

	if response.Result == "success" {
		log.Info("Successfully removed: ", dns)
	} else {
		log.Error("Failed to remove: ", dns)
	}
}

func addDNSRecord(apiKey string, hostFQDN string, currentIP string) {
	args := fmt.Sprintf("&record=%s&type=%s&value=%s", hostFQDN, "A", currentIP)
	dh := dhRequest{URL: "https://api.dreamhost.com/", APIKey: apiKey, CMD: "dns-add_record", Format: "json", Args: args}
	response := dh.getResponse()

	if response.Result == "success" {
		log.Info("Successfully added: ", hostFQDN)
	} else {
		log.Error("Failed to add: ", hostFQDN)
	}
}

func updateDNS(apiKey string, hostFQDN string) {
	currentRecord := getDNSRecord(apiKey, hostFQDN)
	currentIP := getCurrentIP()

	if currentRecord.Value != currentIP {
		log.Info("Updating DNS entry")

		if currentRecord != (dhDNSRecord{}) {
			log.Info("Deleting current record")
			delDNSRecord(apiKey, currentRecord)
		}

		addDNSRecord(apiKey, hostFQDN, currentIP)
	} else {
		log.Info("DNS already up to date")
	}
}

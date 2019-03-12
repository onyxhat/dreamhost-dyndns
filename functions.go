package main

import (
	"io/ioutil"
	"net/http"

	"github.com/rdegges/go-ipify"
	"github.com/thedevsaddam/gojsonq"

	log "github.com/Sirupsen/logrus"
)

func getHTTP(uri string) string {
	r, err := http.Get(uri)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()

	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	return string(contents)
}

func getCurrentIP() string {
	ip, err := ipify.GetIp()
	if err != nil {
		log.Error("Couldn't get IP address: ", err)
	} else {
		log.Info("IP address is: ", ip)
	}

	return ip
}

func (dh dhRequest) getResponse() []byte {
	url := dh.url + "?cmd=" + dh.cmd + "&format=" + dh.format + "&key=" + dh.key
	res, err := http.Get(url)
	jsonResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	return jsonResponse
}

func getDNSRecord(apiKey string, recordValue string) interface{} {
	dh := dhRequest{url: "https://api.dreamhost.com/", key: apiKey, cmd: "dns-list_records", format: "json"}
	response := dh.getResponse()

	jq := gojsonq.New().JSONString(string(response)).From("data").Where("type", "=", "A").Where("record", "=", recordValue)
	res := jq.Get()

	return res
}

func addDNS(apiKey string, dnsName string, currentIP string) []byte {
	dh := dhRequest{url: "https://api.dreamhost.com/", key: apiKey, cmd: "dns-add_record", format: "json"}
	response := dh.getResponse()

	return response
}

func removeDNS(apiKey string, recordValue string) []byte {
	dh := dhRequest{url: "https://api.dreamhost.com/", key: apiKey, cmd: "dns-remove_record", format: "json"}
	response := dh.getResponse()

	return response
}

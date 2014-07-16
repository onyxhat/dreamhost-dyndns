package main

//Imports
import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//Functions
func terminateIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getHttp(uri string) string {
	resp, err := http.Get(uri)
	defer resp.Body.Close()
	terminateIfError(err)

	contents, err := ioutil.ReadAll(resp.Body)
	terminateIfError(err)

	return string(contents)
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func getIndexInSlice(slice []string, text string) int {
	for p, v := range slice {
		if v == text {
			return p
		}
	}
	return -1
}

func getPreviousRecord(apiKey string, dnsEntry string) []string {
	host, _ := os.Hostname()
	uuid, _ := newUUID()

	uri := "https://api.dreamhost.com/?key=" + apiKey + "&unique_id=" + uuid + "&cmd=dns-list_records&ps=" + host
	records := strings.Fields(getHttp(uri))
	index := getIndexInSlice(records, dnsEntry)

	if index == -1 {
		return []string{"0", "0", "0", "0", "0", "0"}
	}

	//'account_id', 'zone', 'record', 'type', 'value', 'comment', 'editable'
	return records[index-2 : index+6]
}

func removeDns(apiKey string, previousRecord []string) string {
	host, _ := os.Hostname()
	uuid, _ := newUUID()

	uri := "https://api.dreamhost.com/?key=" + apiKey + "&unique_id=" + uuid + "&cmd=dns-remove_record&ps=" + host + "&record=" + previousRecord[2] + "&type=" + previousRecord[3] + "&value=" + previousRecord[4] + "&comment=" + previousRecord[5]
	response := strings.Fields(getHttp(uri))

	return response[0]
}

func addDns(apiKey string, dnsName string, currentIp string) string {
	host, _ := os.Hostname()
	uuid, _ := newUUID()
	currTime := time.Now().String()

	uri := "https://api.dreamhost.com/?key=" + apiKey + "&unique_id=" + uuid + "&cmd=dns-add_record&ps=" + host + "&record=" + dnsName + "&type=A&value=" + currentIp + "&comment=" + currTime
	response := strings.Fields(getHttp(uri))

	return response[0]
}

//Runtime
func main() {
	//Application Variables
	apiKey := "YOUR_KEY_GOES_HERE"
	dnsName := "sub.domain.com"
	currentIp := getHttp("http://domain.com/get_myip.php")
	previousRecord := getPreviousRecord(apiKey, dnsName)

	if currentIp != previousRecord[4] {
		fmt.Println("Updating DNS...")
		removeDns(apiKey, previousRecord)
		addDns(apiKey, dnsName, currentIp)
	} else {
		fmt.Println("DNS Unchanged...")
	}
}

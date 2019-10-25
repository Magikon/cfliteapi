package cffunc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	cfst "main/cloudflare-lite-api/cfstructs"
	"net/http"
	"strings"
	"time"
)

var zoneID string
var id string

//Email email for cloudflare
var Email string

//XAuthKey authority key for cludflare
var XAuthKey string

//Bearer key
var Bearer = false

//return structured body
func getResult(body []byte) *cfst.AllInfo {
	var s = new(cfst.AllInfo)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Error in parser:", err)
	}
	return s
}

//pretty print of returned json
func printMyJSON(data []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(prettyJSON.Bytes()))
}

//Akter func akt: 1: getZoneID, 2: findID, 3: listRecords, 4: deleteRecord, 5: updateRecord, 6: createRecord
func Akter(akt int, domainname string, name string, content string, resourcetype string, proxyMXPriority string) {
	var data *strings.Reader
	var req *http.Request
	var err error
	url := "https://api.cloudflare.com/client/v4/zones"
	if akt > 4 {
		if resourcetype == "MX" {
			data = strings.NewReader("{\"type\":\"" + resourcetype + "\",\"name\":\"" + name + "\",\"content\":\"" + content + "\",\"priority\":" + proxyMXPriority + "}")
		} else {
			if resourcetype == "CNAME" && content == "@" {
				data = strings.NewReader("{\"type\":\"" + resourcetype + "\",\"name\":\"" + name + "\",\"content\":\"" + domainname + "\",\"proxied\":" + proxyMXPriority + "}")
			} else {
				data = strings.NewReader("{\"type\":\"" + resourcetype + "\",\"name\":\"" + name + "\",\"content\":\"" + content + "\",\"proxied\":" + proxyMXPriority + "}")
			}
		}
	}
	if akt > 2 || akt == 0 {
		Akter(1, domainname, name, content, resourcetype, proxyMXPriority)
	}
	if akt == 4 || akt == 5 {
		Akter(2, domainname, name, content, resourcetype, proxyMXPriority)
	}
	switch akt {
	case 0:
		time.Sleep(200 * time.Millisecond)
		url += "/" + zoneID + "/purge_cache"
		data = strings.NewReader("{\"purge_everything\":true}")
		req, err = http.NewRequest("POST", url, data)
	case 1: //getZoneID
		url += "?name=" + domainname
		req, err = http.NewRequest("GET", url, nil)
	case 2: //findID
		url += "/" + zoneID + "/dns_records"
		req, err = http.NewRequest("GET", url, nil)
	case 3: //listRecords
		url += "/" + zoneID + "/dns_records"
		req, err = http.NewRequest("GET", url, nil)
	case 4: //deleteRecord
		url += "/" + zoneID + "/dns_records/" + id
		req, err = http.NewRequest("DELETE", url, nil)
	case 5: //updateRecord
		url += "/" + zoneID + "/dns_records/" + id
		req, err = http.NewRequest("PUT", url, data)
	case 6: //createRecord
		url += "/" + zoneID + "/dns_records"
		req, err = http.NewRequest("POST", url, data)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	if Bearer == true {
		req.Header.Add("Authorization", "Bearer " + XAuthKey)
	} else {
		req.Header.Add("X-Auth-Email", Email)
		req.Header.Add("X-Auth-Key", XAuthKey)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	switch akt {
	case 1: //getZoneID
		s := getResult([]byte(body))
		if err == nil && len(s.Result) != 0 {
			zoneID = s.Result[0].ID
		} else {
			log.Fatal(errors.New("Domain name: " + domainname + " not found"))
		}
	case 2: //findID
		s := getResult([]byte(body))
		if err == nil {
			for _, t := range (*s).Result {
				if t.Name == name && t.Type == resourcetype {
					id = t.ID
				}
			}
			if id == "" {
				log.Fatal(errors.New("Resource name: " + name + " not found"))
			}
		}
	case 3: //listRecords
		s := getResult([]byte(body))
		if err == nil {
			fmt.Println("")
			for _, t := range (*s).Result {
				if t.Type == "MX" {
					fmt.Println(t.Type, t.Name, t.Content, t.Priority)
				} else {
					fmt.Println(t.Type, t.Name, t.Content, t.Proxied)
				}
			}
		}
	case 0, 4, 5, 6: //delete update create record
		printMyJSON(body)
	}
}

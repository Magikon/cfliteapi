package main

import (
	cffc "cfmain/cffunc"
	cfst "cfmain/cfstructs"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func writheConf(email string, key string) error {
	var config = new(cfst.CFConf)
	config.Name = "config"
	config.Data.Email = email
	config.Data.XAuthKey = key
	confJSON, _ := json.Marshal(config)
	err := ioutil.WriteFile("./cf.conf", confJSON, 644)
	//fmt.Println(config, confJSON)
	return err
}

func readConfig() error {
	var conf = new(cfst.CFConf)
	// if _, err := os.Stat("./cf.conf"); os.IsNotExist(err) {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Println("Email: ")
	// 	cffc.Email, _ = reader.ReadString('\n')
	// 	cffc.Email = strings.TrimSpace(cffc.Email)
	// 	fmt.Println("Key: ")
	// 	cffc.XAuthKey, _ = reader.ReadString('\n')
	// 	cffc.XAuthKey = strings.TrimSpace(cffc.XAuthKey)
	// 	fmt.Println("Config file not found, create one? (Y/n)")

	// 	cmdString, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 	}
	// 	strings.TrimSpace(cmdString)
	// 	strings.ToUpper(cmdString)
	// 	if cmdString != "N" {
	// 		writheConf(cffc.Email, cffc.XAuthKey)
	// 	}
	// } else
	if file, err := os.Open("./cf.conf"); err == nil {
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&conf); err != nil {
			return err
		}
		cffc.Email = conf.Data.Email
		cffc.XAuthKey = conf.Data.XAuthKey
	}
	return nil
}

func init() {
	// if _, err := os.Stat("./cf.conf"); os.IsNotExist(err) {
	// 	fmt.Println("You need to enter your email and authorization key for authorization in cloudflare")
	// }
	readConfig()
}
func main() {
	emailPtr := flag.String("email", "", "cloudflare email")
	xkeyPtr := flag.String("xkey", "", "cloudflare xkey")
	domainPtr := flag.String("domain", "", "Domain name")
	namePtr := flag.String("name", "", "record name")
	restypePtr := flag.String("type", "", "resource type")
	contentPtr := flag.String("content", "", "content of resource")
	aktPtr := flag.String("akt", "", "create akt")
	proxyPtr := flag.String("proxy", "", "proxy")
	createConfigPtr := flag.Bool("config", true, "create config")
	flag.Parse()
	if *createConfigPtr == true && *emailPtr != "" && *xkeyPtr != "" {
		writheConf(*emailPtr, *xkeyPtr)
	}
	if *emailPtr != "" && *xkeyPtr != "" {
		cffc.Email = *emailPtr
		cffc.XAuthKey = *xkeyPtr
	}
	switch *aktPtr {
	case "create":
		//cffc.CreateRecord(*domainPtr, *namePtr, *contentPtr, *restypePtr, *proxyPtr)
		cffc.Akter(6, *domainPtr, *namePtr, *contentPtr, *restypePtr, *proxyPtr)
	case "update":
		//cffc.UpdateRecord(*domainPtr, *namePtr, *contentPtr, *restypePtr, *proxyPtr)
		cffc.Akter(5, *domainPtr, *namePtr, *contentPtr, *restypePtr, *proxyPtr)
	case "delete":
		//cffc.DeleteRecord(*domainPtr, *namePtr, *restypePtr)
		cffc.Akter(4, *domainPtr, *namePtr, *contentPtr, *restypePtr, *proxyPtr)
	case "list":
		//cffc.ListRecords(*domainPtr, *namePtr, *contentPtr, *restypePtr, *proxyPtr)
		cffc.Akter(3, *domainPtr, *namePtr, *contentPtr, *restypePtr, *proxyPtr)
	default:
		fmt.Println("Create:    cflite -domain=example.com -name=test.example.com -type=A  -content=12.13.14.15   -proxy=true  -akt=create")
		fmt.Println("Create MX: cflite -domain=example.com -name=mx.example.com   -type=MX -content=mx.server.com -proxy=10    -akt=create")
		fmt.Println("Update:    cflite -domain=example.com -name=test.example.com -type=A  -content=12.13.14.16   -proxy=false -akt=update")
		fmt.Println("Delete:    cflite -domain=example.com -name=test.example.com -type=A  -akt=delete")
		fmt.Println("List:      cflite -domain=example.com -akt=list")
	}
}

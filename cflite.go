package main

import (
	"flag"
	"fmt"
	cffc "main/cloudflare-lite-api/cffunc"
	"strings"
	"time"
)

func writheHelp() {
	fmt.Println("Create:       cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=test.example.com -type=A      -content=12.13.14.15   -proxy=true  -akt=create")
	fmt.Println("Create:       cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=test.example.com -type=CNAME  -content=example.com   -proxy=true  -akt=create")
	fmt.Println("Create MX:    cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=mx.example.com   -type=MX     -content=mx.server.com -proxy=10    -akt=create")
	fmt.Println("Update:       cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=test.example.com -type=A      -content=12.13.14.16   -proxy=false -akt=update")
	fmt.Println("Delete:       cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=test.example.com -type=A      -akt=delete")
	fmt.Println("List:         cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -akt=list")
	fmt.Println("Purge Cache:  cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -akt=purge")
}

func main() {
	startTime := time.Now()

	emailPtr := flag.String("email", "", "cloudflare email")
	xkeyPtr := flag.String("xkey", "", "cloudflare xkey")
	domainPtr := flag.String("domain", "", "Domain name")
	namePtr := flag.String("name", "", "record name")
	restypePtr := flag.String("type", "", "resource type")
	contentPtr := flag.String("content", "", "content of resource")
	aktPtr := flag.String("akt", "", "create akt")
	proxyPtr := flag.String("proxy", "", "proxy")
	helpPtr := flag.Bool("help", false, "help")

	flag.Parse()

	if *emailPtr == "" {
		cffc.Bearer = true
	}
	cffc.Email = *emailPtr
	cffc.XAuthKey = *xkeyPtr

	if *helpPtr {
		writheHelp()
	}
	var name string
	for _, domain := range strings.Split(*domainPtr, ",") {
		if !strings.Contains(*namePtr, domain) {
			name = *namePtr + "." + domain
		} else {
			name = *namePtr
		}
		switch *aktPtr {
		case "create":
			//cffc.CreateRecord(domain, name, *contentPtr, *restypePtr, *proxyPtr)
			cffc.Akter(6, domain, name, *contentPtr, *restypePtr, *proxyPtr)
			//purge cache
			fmt.Println("Purge cache for domain " + domain)
			cffc.Akter(0, domain, name, *contentPtr, *restypePtr, *proxyPtr)
		case "update":
			//cffc.UpdateRecord(domain, name, *contentPtr, *restypePtr, *proxyPtr)
			cffc.Akter(5, domain, name, *contentPtr, *restypePtr, *proxyPtr)
			//purge cache
			fmt.Println("Purge cache for domain " + domain)
			cffc.Akter(0, domain, name, *contentPtr, *restypePtr, *proxyPtr)
		case "delete":
			//cffc.DeleteRecord(domain, name, *restypePtr)
			cffc.Akter(4, domain, name, *contentPtr, *restypePtr, *proxyPtr)
			//purge cache
			fmt.Println("Purge cache for domain " + domain)
			cffc.Akter(0, domain, name, *contentPtr, *restypePtr, *proxyPtr)
		case "list":
			//cffc.ListRecords(domain, name, *contentPtr, *restypePtr, *proxyPtr)
			cffc.Akter(3, domain, name, *contentPtr, *restypePtr, *proxyPtr)
		case "purge":
			//purge cache
			fmt.Println("Purge cache for domain " + domain)
			cffc.Akter(0, domain, name, *contentPtr, *restypePtr, *proxyPtr)
		default:
		}
	}
	fmt.Println("Working time is: ", time.Since(startTime))
}

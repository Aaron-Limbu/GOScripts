package main

import (
	"fmt"
	"os"
	"net/http"
	"flag"
	"strings"
	"bufio"
)

var (
	url string 
	path string 
)

func fuzzPath(target, wpath string){
	file, err := os.Open(wpath)
	if err != nil{
		fmt.Println("[-] Failed to open wordlists ",err)
		return 
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		path := scanner.Text()
		fullUrl := strings.TrimRight(target,"/") + "/" + wpath
		resp, err := http.Get(fullUrl)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200{
			fmt.Printf("[+] Found: %s (200 OK)\n",fullUrl)
		}
	}
}

func main(){
	flag.StringVar(&url,"url","default","example: -url https://example.com")
	flag.StringVar(&path,"path","default")
	flag.Parse()
	if url == ""{
		fmt.Println("[i] Enter url ")
	}
	fuzzPath(url, path)
}

package main 

import (
	"fmt"
	"net/http"
	"flag"
)

var url string 

func checkHTTPSmethods(url string){
	fmt.Println("[+] Checking for HTTP methods")
	req, err := http.NewRequest("OPTIONS",url,nil)
	if err != nil {
		fmt.Println("[-] Request Error: ",err)
		return 
	}
	client := &http.Client{}
	resp, err := client.Do{req}
	if err != nil{
		fmt.Println("[-] Error: ",err)
		return 
	}
	defer resp.Body.Close()
	allowed := resp.Header.Get("Allow")
	if allowed != ""{
		fmt.Println("[i] Allowed methods: ",allowed)
		if contains(allowed,"PUT") || contains(allowed,"DELETE") || contains(allowed,"PATCH") || contains(allowed,"TRACE"){
			fmt.Println("[i] Insecure HTTP methods enabled")
		}
	}else{
		fmt.Println("[-] No allowed header found")
	}
}

func contains(s,substr string) bool{
	return len(s)>0 && strings.Contains(s, substr)
}

func main(){
	flag.StringVar(&url,"url","default","example:- -url https://example.com "
)
	flag.Parse()
	if url == ""{
		fmt.Println("[i] please enter url")
		return 
	}
	checkHTTPSmethods(url)
}

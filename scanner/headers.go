package main

import (
	"fmt"
	"net/http"
	"flag"
	"io"
	"crypto/tls"
	"time"
)

var url string

func analyzeHeaders(url string){
	fmt.Println("[+] Analyzing HTTP headers for ",url)
	security_headers := map[string]string{
		"Content-Security-Policy":"Protects against XSS attacks",
		"Strict-Transport-Security":"Enforces HTTPS",
		"X-Frame-Options":"Mitigates clickjacking",
		"X-XSS-Protection":"Enables XSS filter",
		"X-Content-Type-Options":"Prevents MIME type sniffing",
		"Referrer-Policy":"Controls referrer information",
		"Permission-Policy":"Restricts browser features"
	}
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	for key, value := range resp.Header {
		fmt.Printf("[+] %s: %s \n",key,value)
	}
	for header, description := range security_headers{
		if _, ok := resp.Header[header]; ok{
			fmt.Printf("[i] %s is present: %s\n",header,description)
		}else{
			fmt.Printf("[i] %s is missing: %s\n",header,description)
		}
	}
}

func analyzeTLS(url string){
	fmt.Println("[+] Analyzing SSL/TLS for ",url)
	conn, err :=  tls.Dial("tcp",url+"443",nil)
	if err != nil{
		fmt.Println("[-] TLS connection failed: ",err)
		return
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		fmt.Println("[-] No certificates found")
		return
	}
	
	cert := certs[0]
	fmt.Println("[i] Subject: ",cert.Subject)
	fmt.Println("[i] Issuer: ",cert.Issuer)
	fmt.Println("[i] Valid From: ",cert.NotBefore)
	fmt.Println("[i] Valid To: ",cert.NotAfter)

	if time.Now().After(cert.NotAfter){
		fmt.Println("[-] Certificate is Expired")
	}else{
		fmt.Println("[+] Certificate is valid")
	}
	state := conn.ConnectionState()
	fmt.Printf("[i] TLS Version: %s\n".state.Version)
	fmt.Printf("[i] Cipher Suite: 0x%x\n".state.CipherSuite)
}

func main(){
	flag.StringVar(&url,"url","default","example: -url https://example.com")
	flag.Parse()
	analyzeHeaders(url)
	analyzeTLS(url)
}

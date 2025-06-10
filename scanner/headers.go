package scanner

import (
	"fmt"
	"net/http"
	"flag"
	"io"
	"crypto/tls"
	"time"
	"strings"
	"toolkit/reports"
)

var url string

func AnalyzeHeaders(url string){
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
	headersFound := make(map[string]string)
	for header, description := range security_headers{
		if _, ok := resp.Header[header]; ok{
			fmt.Printf("[i] %s is present: %s\n",header,description)
			headersFound[header] = "Present"
		}else{
			fmt.Printf("[i] %s is missing: %s\n",header,description)
			headersFound[header] = "Missing"
		}
	}
	reports.AddResult("Headers Scan",url,headersFound,"Completed" )
}

func AnalyzeTLS(host string) {
	fmt.Println("[+] Analyzing SSL/TLS for", host)
	conn, err := tls.Dial("tcp", host+":443", nil)
	if err != nil {
		fmt.Println("[-] TLS connection failed:", err)
		return
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		fmt.Println("[-] No certificates found")
		return
	}
	cert := certs[0]
	state := conn.ConnectionState()
	certsFound := map[string]string{
		"Subject":     cert.Subject.CommonName,
		"Issuer":      cert.Issuer.CommonName,
		"Valid From":  cert.NotBefore.Format(time.RFC3339),
		"Valid Until": cert.NotAfter.Format(time.RFC3339),
		"TLS Version": tlsVersionToString(state.Version),
		"CipherSuite": fmt.Sprintf("0x%x", state.CipherSuite),
	}

	if time.Now().After(cert.NotAfter) {
		certsFound["Status"] = "Expired"
		fmt.Println("[-] Certificate is Expired")
	} else {
		certsFound["Status"] = "Valid"
		fmt.Println("[+] Certificate is Valid")
	}
	for k, v := range certsFound {
		fmt.Printf("[i] %s: %s\n", k, v)
	}
	reports.AddResult("TLS Scanner", host, certsFound, "Completed")
}

func tlsVersionToString(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown"
	}
}

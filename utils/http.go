package main

import (
	"fmt"
	"net/http"
	"strings"
)

func CleanURL(raw string) string{
	if !strings.HasPrefix(raw,"http://") && strings.HasPrefix(raw,"https://"){
		raw = "http://" + raw
	}
	return raw
}

func GetDomain(raw string) string{
	u, err := url.Parse(raw)
	if err != nil {
		fmt.Println("[-] Invalid URL: ",err)
		return ""
	}
	return u.Hostname()
}

func StripProtocol(raw string) string{
	raw = strings.TrimPrefix(raw,"https://")
	raw = strings.TrimPrefix(raw,"http://")
	return strings.TrimRight(raw,"/")
}

package main
import (
	"buffio"
	"fmt"
	"net/http"
	"flag"
	"os"
	"strings"
)

var (
	url string 
	path string
)

func fuzzWeb(target string, wpath string){
	
}


func main(){
	flag.StringVar(&url,"url","default","example:- -url https://example.com")
	flag.StringVar(&path,"path","../wordlists/wordlists.txt","example:- -path ../wordlists/wordlists.txt or /home/username/Wordlists/wordlists.txt")
	flag.Parse()

}

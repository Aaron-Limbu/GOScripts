package main
import ("fmt"
	"net/http"
	"flag"
	"toolkit/scanner"
	"toolkit/utils"
	"toolkit/reports"
)


func checkStatus(url string) bool{
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	return true
}

func main(){
	var url string 
	var wordlists string
	var output string
	flag.StringVar(&url,"url","default","website url")
	flag.StringVar(&wordlists,"wordlists","default","wordlist path")
	flag.StringVar(&output,"output","report.txt","output filename")
	flag.Parse()
	if url == ""{
		fmt.Println("[i] Provide a URL using -url")
		return
	}
	if wordlists == ""{
		fmt.Println("[i] Provide wordlists path using -wordlists")
		return
	}
	if checkStatus(url){
		url = utils.CleanURL(url)
		scanner.AnalyzeHeaders(url)
		scanner.AnalyzeTLS(url)

		err := reports.SaveReports(output)
		if err != nil{
			fmt.Println("[-] Failed to save report: ",err)
		}else{
			fmt.Println("[+] Report saved: ",err)
		}
	}
}

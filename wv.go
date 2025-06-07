package main
import ("fmt"
	"net/http"
	"flag"
	"io"
)

var (
	url string
	wordlists string
)

func checkStatus(url string){
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main(){
	flag.StringVar(&url,"url","default","website url")
	flag.StringVar(&wordlists,"wordlists","default","wordlist path")
	flag.Parse()
	
	checkStatus(url)
}

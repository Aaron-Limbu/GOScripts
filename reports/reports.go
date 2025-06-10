package reports

import (
	"fmt"
	"encoding/json"
	"os"
)

type ScanRepo struct {
	Module string 
	Target string
	Details map[string]string
	Status string						
}

var Results []ScanRepo

func AddResult(module, target string, details map[string]string, status string){
	result := ScanRepo{
		Module: module,
		Target: target,
		Details: details,
		Status: status,
	}
	Results = append(Results, result)
}

func PrintResults(){
	for _, res := range Results{
		fmt.Printf("\n[%s] %s - $s\n",res.Status,res.Module,res.Target)
		for k, v := range res.Details{
			fmt.Printf(" - %s: %s\n",k,v)
		}
	}
}

func SaveReports(filename string) error{
	file, err := os.Create(filename)
	if err != nil{
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent(""," ")
	return enc.Encode(Results)
}



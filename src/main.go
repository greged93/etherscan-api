package main

import (
	"os"
	"bufio"
	"strconv"
	"time"
	"net/http"
	"fmt"
	"encoding/json"
	m "urlRequest"
)
const API_TOKEN string = "122VMCKVN32VA5WY34TSRCFU7IFWC24MMH"

type TransactionRequest struct {
	Status string
	Message string
	Result string
}


func makeRequest(url string, h *TransactionRequest) error {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return err 
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(h)
}

func handleRequest(request m.UrlRequest, ch chan float64) {
	url, err := request.MakeUrl()
	if err != nil {
		fmt.Println("Error in MakeUrl: ", err)
		return
	}
	var h TransactionRequest = TransactionRequest{}
	errResp := makeRequest(url.String(), &h)
	if errResp != nil{
		fmt.Println(errResp)
		return
	} 
	f, err := strconv.ParseFloat(h.Result, 64)
	if err != nil {
		fmt.Println("Conversion error, ", err)
		return 
	}
	ch <- f/1e18
}

func main() {
	var baseUrl string = "https://api.etherscan.io/api"
	var queries = map[string]string{"module": "account", 
									"action": "balance",
									"address": "0x3C517c5d2040B995e697c7b916d120a4f7Fa095d",
									"tag": "latest",
									"apikey": API_TOKEN,}
	var request m.UrlRequest = m.UrlRequest{BaseUrl: baseUrl, Queries: queries}
	ch := make(chan float64)
	var text string
	
	for ;text != "End\n"; {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ = reader.ReadString('\n')
		fmt.Println(text)
		go handleRequest(request, ch)
		fmt.Println(<-ch)
	}

}

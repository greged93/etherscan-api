package main

import(
	"fmt"
	"net/http"
)

type HttpRequestError struct {
	url string
}

func (e HttpRequestError) Error() string {
	return fmt.Sprintf("Error at url %s", e.url)
}


func main() {
	fmt.Printf("Hello World!")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("Error !")
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
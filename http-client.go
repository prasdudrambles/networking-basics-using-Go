package main

import(
	"fmt"
    	"io"
//	"log"
	"net/http"
//	"github.com/go-chi/chi/v5"
//	"encoding/json"
//	"time"
)

func main(){
	client := http.Client{}

	requestURL := "http://localhost:7000/getList"
	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil{
		return
	}
	response, err := client.Do(request)
	if err != nil{
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	fmt.Printf("Body : %s", body)
}




package main

import(
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"encoding/json"
	"fmt"
//	"time"
)

func main(){
	router := chi.NewRouter()
	router.Get("/health", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("connected to Abdul")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Connected to Abdul's server"))
	})
	router.Get("/getList", func(w http.ResponseWriter, r *http.Request){
		list := make([]int, 0)
		for i:=0; i<100; i++{
			if i%2==0{
				list = append(list, i)
			}
		}
		res,_ := json.Marshal(list) 
		//fmt.Println(string(res))
		w.Write([]byte(string(res)))
	})


	router.Get("/youtube", func(w http.ResponseWriter, r *http.Request){
		//time.Sleep(5* time.Second)
		http.Redirect(w, r,  "https://www.youtube.com", http.StatusMovedPermanently)
	})
	err := http.ListenAndServe(":7000", router)
	if err == nil{
		log.Println("error in starting the server")
		return
	}

}




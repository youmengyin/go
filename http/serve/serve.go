package main

import (
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/test", handle)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(res http.ResponseWriter, req *http.Request){
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	res.Write([]byte ("server response!"))
}
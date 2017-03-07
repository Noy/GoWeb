package main

import (
	"fmt"
	"net/http"
	"./powerweb"
)

func main() {
	start()
}

func start()  {
	fmt.Println("Attempting to start...")
	http.HandleFunc("/", func (responseWriter http.ResponseWriter, request *http.Request) {
		ip, err := powerweb.Nest(nil, request)
		if err != nil { panic(err) }
		// HTML which includes a header and some content so the front page won't be empty. It's always better to make a static folder.
		fmt.Fprint(responseWriter, "<head><link href=\"https://fonts.googleapis.com/css?family=Source+Code+Pro\" rel=\"stylesheet\"></head>" +
			"<style> body { background-color: #E27B65; text-align: center; position: relative; font-family: 'Source Code Pro', monospace;" +
			"font-size: 30px;} </style>")
		fmt.Fprint(responseWriter, "<br><br><br><br><h1>Your IP: " + ip + "</h1>")
	})
	powerweb.AllowPages()
	http.ListenAndServe(":90", nil)
}

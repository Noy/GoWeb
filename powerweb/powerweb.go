package powerweb

import (
	"errors"
	"net/http"
	"strings"
)

func Nest(delegate interface{}, request *http.Request) (string, error) {
	var value string
	if delegate == "" {
		value = request.Header.Get(delegate.(string))
		if value == "" { value = request.RemoteAddr }
	}
	value = request.RemoteAddr
	ips := strings.Split(value, ",")
	ip := strings.TrimSpace(ips[len(ips)-1])
	index := strings.LastIndex(ip, ":")
	if index != -1 { ip = ip[:index] }
	if ip == "" {
		err := errors.New("No IP availible!")
		return "", err
	}
	return ip, nil
}

func AllowPages() {
	/* 
	wanted to clear this up, this is to do something like: http://yourserver.com:90/site IN FOLDER './html'
	it'll read the file from the specific directory you choose..
	*/
	http.Handle("/site/", http.StripPrefix("/site/", http.FileServer(http.Dir("./html"))))
}

// eventually make cool things

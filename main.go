package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "os"
)

var host string = ""
var port string = ""

func handler(w http.ResponseWriter, r *http.Request) {
    target := fmt.Sprintf("http://%s:%s", host, port)
    resp, err := http.Get(target)
    if err != nil {
        fmt.Fprintln(w, "Error fetching IP...")
	fmt.Fprintln(w, err)
	return
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Fprintln(w, "Error reading response body...")
	return
    }
    fmt.Fprintln(w, "Response => " + string(body))
}

func main() {
    host = os.Getenv("EXTERNAL_WHATISMYIP_SERVICE_HOST")
    port = os.Getenv("EXTERNAL_WHATISMYIP_SERVICE_PORT")
    fmt.Println(host)
    fmt.Println(port)
    if host == "" || port == "" {
        fmt.Println("EXTERNAL_WHATISMYIP_SERVICE_[HOST|PORT] environment variable not set. Exiting")
        return
    }
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

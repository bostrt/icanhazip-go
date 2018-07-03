package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://ipv4.icanhazip.com")
    if err != nil {
        fmt.Fprintln(w, "Error fetching IP...")
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
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

package main

import (
  "log"
  "net/http"
  "flag"
)

var urlPtr = flag.String("url", "http://google.com", "Url to redirect to - full path")

func redirect(w http.ResponseWriter, r *http.Request) {

  http.Redirect(w, r, *urlPtr, 301)
}

func main() {

  flag.Parse()

  http.HandleFunc("/", redirect, )

  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

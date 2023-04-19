package main

import (
   "net/http"
   "log"
   "fmt"
)

func log_request(next http.HandlerFunc) http.HandlerFunc{
   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      log.Println(fmt.Sprintf("%v %v", r.Method, r.RequestURI))
      next.ServeHTTP(w, r)
   })
}

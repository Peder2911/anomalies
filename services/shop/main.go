package main

import (
   "net/http"
   "log"
   "fmt"
)

func main(){
   err := initialize_database()
   if err != nil {
      panic(err)
   }

   log.Println("Database initialized")
   http.DefaultServeMux.Handle("/", log_request(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      fmt.Fprintf(w, "Hello world")
   })))

   http.DefaultServeMux.Handle("/me", log_request(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      fmt.Fprintf(w, "Hello user")
   })))

   log.Println("Serving on :8000")

   http.ListenAndServe(":8000",nil)
}

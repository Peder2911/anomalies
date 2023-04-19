package main

import (
   "net/http"
   "log"
   "fmt"
   "database/sql"
   "encoding/json"
)

var Db  *sql.DB

func main(){
   var err error

   Db, err = sql.Open("sqlite3", "./shop.db")
   if err != nil {
      panic(err)
   }

   err = initialize_database(Db)
   if err != nil {
      panic(err)
   }
   log.Println("Database initialized")


   http.DefaultServeMux.Handle("/", log_request(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      fmt.Fprintf(w, "Hello world")
   })))

   http.DefaultServeMux.Handle("/logout", log_request(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      cookie, err := r.Cookie("session")
      if err != nil {
         w.WriteHeader(400)
         return
      }
      _,err = Db.Exec("delete from sessions where id=?", cookie.Value)
      if err != nil {
         w.WriteHeader(500)
         log.Println(fmt.Sprintf("Failed to delete session: %v", err))
         return
      }

      w.Header().Add("Set-Cookie","session=deleted")
   })))

   http.DefaultServeMux.Handle("/login", log_request(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      var user_id int
      var login_request LoginRequest 
      var session Session
      var session_id string
      var err error
      var rows *sql.Rows

      err = json.NewDecoder(r.Body).Decode(&login_request)
      if err != nil {
         w.WriteHeader(400)
         fmt.Fprint(w, "Failed to read request body")
         return
      }
      log.Println(fmt.Sprintf("%v:%v",login_request.Username, login_request.Password))

      rows, err = Db.Query(
         "select id from users where name=? and password=?", 
         login_request.Username, 
         login_request.Password)
      if err != nil {
         w.WriteHeader(500)
         fmt.Fprintf(w, "Failed to get user: %v", err)
         return 
      }
      defer rows.Close()

      if rows.Next() {
         rows.Scan(&user_id)
      } else {
         w.WriteHeader(401)
         fmt.Fprint(w, "Unathorized")
         return
      }
      rows.Close()

      rows, err = Db.Query("select id from sessions where user=?", user_id)
      if err != nil {
         w.WriteHeader(500)
         fmt.Fprintf(w, "Internal server error: %v", err)
         return
      }
      defer rows.Close()

      if rows.Next() {
         rows.Scan(&session_id)
      } else {
         session = create_session(user_id)
         rows.Close()
         _, err = Db.Exec("insert into sessions values (?, ?, ?)",
            session.Id,
            session.CreatedAt,
            session.User)
         if err != nil {
            w.WriteHeader(500)
            fmt.Fprintf(w, "Internal server error: %v", err)
            return
         }
         session_id = session.Id
      }

      w.Header().Add("Set-Cookie", fmt.Sprintf("session=%v; HttpOnly; Secure; ; SameSite=Strict", session_id))
   })))

   http.DefaultServeMux.Handle("/products", log_request(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      var id int16
      var name string
      var description string
      var response ProductResponse = ProductResponse{[]Product{}}

      rows, err := Db.Query("select * from products")
      if err != nil {
         fmt.Fprintf(w, "Error when quering for data: %v", err)
         w.WriteHeader(500)
         return
      }

      for rows.Next(){
         rows.Scan(&id, &name, &description)
         response.Products = append(response.Products, Product{id, name, description})
      }

      data, err := json.Marshal(response)
      if err != nil {
         fmt.Fprintf(w,"Failed to serialize response")
         w.WriteHeader(500)
      }

      w.WriteHeader(200)
      w.Write(data)
   })))

   http.DefaultServeMux.Handle("/me", log_request(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      user_id,err := user_from_session(Db, r)
      if err != nil {
         w.WriteHeader(401)
         log.Println(fmt.Sprintf("%v", err))
         return
      }
      row := Db.QueryRow("select name from users where id=?", user_id)

      var username string
      err = row.Scan(&username)
      if err != nil {
         w.WriteHeader(500)
         fmt.Fprintf(w, "Internal server error: %v", err)
         return
      }

      fmt.Fprintf(w, "Hello %v", username)
   })))

   log.Println("Serving on :8000")

   http.ListenAndServe(":8000",nil)
}

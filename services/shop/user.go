package main

import (
	"database/sql"
	"net/http"
)

type User struct {
   Id        int
   LastLogin int64
   Name      string
   Password  string
}

type LoginRequest struct {
   Username string
   Password string
}

func user_from_session(db *sql.DB,r *http.Request) (int, error) {
   cookie, err := r.Cookie("session")
   if err != nil {
      return -1, err
   }
   row := db.QueryRow("select user from sessions where id=?", cookie.Value)
   var user_id int 
   err = row.Scan(&user_id)
   if err != nil {
      return -1,err
   }
   return user_id,nil
}

package main 

import (
   "crypto/rand"
   "encoding/hex"
   "time"
)

type Session struct {
   id         string
   created_at int64
   user       int
}

func (s Session) Age() int64 {
   return time.Now().Unix() - s.created_at
}

func random_id_string() string {
   id := make([]byte,32)
   _,err := rand.Read(id)
   if err != nil {
      panic(err)
   }
   return hex.EncodeToString(id) 
}

func create_session(user int) Session {
   return Session{
      random_id_string(),
      time.Now().Unix(),
      user,
   }
}

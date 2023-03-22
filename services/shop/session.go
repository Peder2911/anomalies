package main

import (
   "crypto/rand"
   "encoding/hex"
   "time"
)

type Session struct {
   Id        string
   CreatedAt int64
   User      int
}

func (s Session) Age() int64 {
   return time.Now().Unix() - s.CreatedAt
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

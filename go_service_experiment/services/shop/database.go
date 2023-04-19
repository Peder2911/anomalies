package main
import (
   "database/sql"
   _ "github.com/mattn/go-sqlite3"
)

var tables = []string {
      `
      create table if not exists users (
         id       integer  primary key,
         last_login int,
         name     text not null,
         password text not null
      )
      `,
      `
      insert into users (name, password) values ('admin', 'admin')
      `,
      `
      create table if not exists products (
         id          integer  primary key autoincrement,
         name        text not null,
         description text not null
      )
      `,
      `
      create table if not exists orders (
         id      integer  primary key autoincrement,
         created_at int not null,
         user    int  not null,
         product int  not null,
         status  text check(status in ('pending', 'completed')) not null default 'pending',
         foreign key (user) references user (id),
         foreign key (product) references product (id)
      )
      `,
      `
      create table if not exists sessions (
         id      text primary key,
         created_at int not null,
         user    int  not null,
         foreign key (user) references user (id)
      )
      `,
}

func initialize_database(con *sql.DB) error {
   var err error
   for _,s := range tables {
      _, err = con.Exec(s)
      if err != nil {
         return err
      }
   }
   return nil
}

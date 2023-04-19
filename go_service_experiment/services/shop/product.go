package main

type Product struct {
   id int16
   name string
   description string
}

type ProductResponse struct {
   Products []Product
}

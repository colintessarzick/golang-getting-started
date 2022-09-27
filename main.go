package main

type album struct {
	ID string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
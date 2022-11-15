package main

import (
	_ "github.com/gin-gonic/gin"
	_ "log"
	_ "math/rand"
	_ "net/http"
	_ "strconv"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

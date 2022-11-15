package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// postAlbums добавляет альбом из JSON,
// полученного в теле запроса.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Вызов BindJSON для привязки
	// полученного JSON к newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Добавляем в срез новый альбом.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID находит альбом,
// значение идентификатора которого совпадает с
// параметром id, отправленным клиентом,
// затем возвращает этот альбом в качестве ответа
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Перебираем список альбомов в поисках альбома,
	// значение идентификатора которого соответствует параметру.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// album представляет данные об альбоме.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Срез albums для заполнения данных об альбомах.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums возвращает список всех альбомов в формате JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

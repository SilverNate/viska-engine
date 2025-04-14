package library

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
	"viska/data-service/internal/redis"
)

type LibraryHandler struct {
	service ILibraryService
	redis   *redis.RedisClient
}

func NewLibraryHandler(service ILibraryService, redis *redis.RedisClient) *LibraryHandler {
	return &LibraryHandler{service: service, redis: redis}
}

func (h *LibraryHandler) CreateAuthor(c *gin.Context) {
	var author Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("implementasi validator go playground")
	validate := validator.New()
	if err := validate.Struct(author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("validation failed: %v", err),
		})
		return
	}

	if err := h.service.CreateAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, author)
}

func (h *LibraryHandler) GetAllAuthor(c *gin.Context) {
	cacheKey := "author:all"

	if cached, err := h.redis.Get(cacheKey); err == nil {
		var books []Book
		_ = json.Unmarshal([]byte(cached), &books)
		c.JSON(http.StatusOK, books)
		return
	}

	author, err := h.service.GetAllAuthor()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if jsonData, err := json.Marshal(author); err == nil {
		_ = h.redis.Set(cacheKey, string(jsonData), 60)
	}

	c.JSON(http.StatusOK, author)
}

func (h *LibraryHandler) GetAuthorByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	author, err := h.service.GetAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *LibraryHandler) UpdateAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var author Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author.ID = uint(id)
	if err := h.service.UpdateAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, author)
}

func (h *LibraryHandler) DeleteAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteAuthor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "author deleted"})
}

func (h *LibraryHandler) CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("validation failed: %v", err),
		})
		return
	}

	if err := h.service.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (h *LibraryHandler) GetAllBook(c *gin.Context) {
	cacheKey := "books:all"

	if cached, err := h.redis.Get(cacheKey); err == nil {
		var books []Book
		_ = json.Unmarshal([]byte(cached), &books)
		c.JSON(http.StatusOK, books)
		return
	}

	books, err := h.service.GetAllBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if jsonData, err := json.Marshal(books); err == nil {
		_ = h.redis.Set(cacheKey, string(jsonData), 60)
	}

	c.JSON(http.StatusOK, books)
}

func (h *LibraryHandler) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.service.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *LibraryHandler) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.ID = uint(id)
	if err := h.service.UpdateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *LibraryHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}

func (h *LibraryHandler) CreatePublisher(c *gin.Context) {
	var publisher Publisher
	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(publisher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("validation failed: %v", err),
		})
		return
	}

	if err := h.service.CreatePublisher(&publisher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, publisher)
}

func (h *LibraryHandler) GetAllPublisher(c *gin.Context) {
	cacheKey := "publisher:all"

	if cached, err := h.redis.Get(cacheKey); err == nil {
		var books []Book
		_ = json.Unmarshal([]byte(cached), &books)
		c.JSON(http.StatusOK, books)
		return
	}

	publisher, err := h.service.GetAllPublisher()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if jsonData, err := json.Marshal(publisher); err == nil {
		_ = h.redis.Set(cacheKey, string(jsonData), 60)
	}

	c.JSON(http.StatusOK, publisher)
}

func (h *LibraryHandler) GetPublisherByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	publisher, err := h.service.GetPublisherByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "publisher not found"})
		return
	}
	c.JSON(http.StatusOK, publisher)
}

func (h *LibraryHandler) UpdatePublisher(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var publisher Publisher
	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publisher.ID = uint(id)
	if err := h.service.UpdatePublisher(&publisher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, publisher)
}

func (h *LibraryHandler) DeletePublisher(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeletePublisher(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Publisher deleted"})
}

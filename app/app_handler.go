package app

import (
	"gin-example/app/data"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	UserService data.UserService
	BookService data.BookService
}

func (handler *Handler) authenticator(c *gin.Context) (interface{}, error) {
	var loginVals LoginReq
	if err := c.ShouldBind(&loginVals); err != nil {
		log.Println(err.Error())
		return nil, jwt.ErrMissingLoginValues
	}

	// search generalUser
	var user, err = handler.UserService.User(loginVals.UserName, loginVals.Password)
	if err != nil {
		if err == data.ErrDBConnection {
			c.Status(http.StatusInternalServerError)
		}
		return nil, err
	}

	return &UserClaims{
		UserId: user.ID,
	}, nil
}

func (handler *Handler) signUp(c *gin.Context) {
	var req SignUpReq
	if err := c.Bind(&req); err != nil {
		// request error handling
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if err := handler.UserService.CreateUser(req.UserName, req.Password); err != nil {
		switch err {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "sign up customer success",
	})
}

func (handler *Handler) authProfile(c *gin.Context) {
	userClaims := ConvertClaims(jwt.ExtractClaims(c))
	c.JSON(http.StatusOK, userClaims)
}

func (handler *Handler) book(c *gin.Context) {
	bookId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, handler.BookService.Book(uint(bookId)))
}

func (handler *Handler) postBook(c *gin.Context) {
	var bookReq BookReq
	if err := c.ShouldBind(&bookReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	id, err := handler.BookService.PostBook(bookReq.Title, bookReq.Author)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Post Book Success",
		"id":      id,
	})
}

func (handler *Handler) putBook(c *gin.Context) {
	bookId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	var bookReq BookReq
	if err := c.ShouldBind(&bookReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if err != handler.BookService.PutBook(uint(bookId), bookReq.Title, bookReq.Author) {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Put Book Success",
	})
}

func (handler *Handler) deleteBook(c *gin.Context) {
	bookId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if err != handler.BookService.DeleteBook(uint(bookId)) {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Delete Book Success",
	})
}

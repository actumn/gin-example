package app

import (
	"gin-example/app/data/mocks"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"testing"
)

var userServiceMock = mocks.UserService{}
var bookServiceMock = mocks.BookService{}
var handler = &Handler{
	UserService: &userServiceMock,
	BookService: &bookServiceMock,
}
var ginEngine = GinEngine(handler)

func init() {
	viper.Set("port", "8080")
	gin.SetMode(gin.TestMode)
}

func TestLogin(t *testing.T) {

}

func TestSignUp(t *testing.T) {

}

package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	rest_err "github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"github.com/matheuszzy/crud-basico-go/src/controller/adapters/in"
)

func CreateUser(c *gin.Context) {
	var userRequest in.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		resterr := rest_err.NewBadRequestError(fmt.Sprintf("invalid json body, error =%s\n", err.Error()))

		c.JSON(resterr.Code, resterr)
		return
	}

	fmt.Println(userRequest)
}

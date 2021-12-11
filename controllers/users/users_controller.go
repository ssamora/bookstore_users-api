package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"users/domain/users"

	"users/services"

	"users/utils/errors"
)

func GetUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var user users.User
	/*fmt.Println("user before unmarshal ", user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: Handle json error
	}*/

	//This function does the above code
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func FindUser(c *gin.Context) {

}

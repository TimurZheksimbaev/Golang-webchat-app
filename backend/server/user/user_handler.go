package user

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type Handler struct {
	service ServiceI
}

func NewHandler(service ServiceI) *Handler {
	return &Handler{
		service: service,
	}
}


func (h *Handler) SignUp(c *gin.Context) {
	var u CreateUserRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 
	}

	response, err := h.service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	 c.JSON(http.StatusOK, response)

}

func (h *Handler) Login(c *gin.Context) {
	var user LoginUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	u, err := h.service.Login(c.Request.Context(), &user) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.SetCookie("jwt", u.accessToken, 3600, "/", "localhost", false, true)

	res := &LoginUserResponse{
		Username: u.Username,
		ID: u.ID,
	}

	c.JSON(http.StatusOK, res)
	
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
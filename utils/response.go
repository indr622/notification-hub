package utils

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Respond(c *gin.Context, status int, message string, result interface{}, err error) {
	resp := APIResponse{
		Message: message,
		Result:  result,
	}

	if err != nil {
		resp.Error = err.Error()
	}

	c.JSON(status, resp)
}

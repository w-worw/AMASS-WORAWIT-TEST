package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Result  interface{} `json:"result"`
}

func NewResponse(success bool, code int, msg string, err string, result interface{}) *Response {
	return &Response{
		Success: success,
		Code:    code,
		Message: msg,
		Error:   err,
		Result:  result,
	}
}

func (res *Response) Abort(c *gin.Context, statusCode int) {
	c.AbortWithStatusJSON(statusCode, res)
}

func (res *Response) Send(c *gin.Context, statusCode int) {
	c.JSON(statusCode, res)
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	response := NewResponse(false, statusCode, err.Error(), err.Error(), nil)
	response.Abort(c, statusCode)
}

func SuccessResponse(c *gin.Context, msg string, result interface{}) {
	statusCode := http.StatusOK
	response := NewResponse(true, statusCode, msg, "", result)
	response.Send(c, statusCode)
}

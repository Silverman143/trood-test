package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type ResponseBuilder struct {
}

func NewResponseBuilder() *ResponseBuilder {
    return &ResponseBuilder{}
}

// Response базовая структура ответа
// @Description Стандартная структура ответа API
type Response struct {
    Data  interface{} `json:"data,omitempty" swaggertype:"object"`  // Данные ответа
    Error string      `json:"error,omitempty" example:"error message"` // Сообщение об ошибке
}

// ErrorResponse структура для ответов с ошибками
// @Description Структура ответа при ошибке
type ErrorResponse struct {
    Error string `json:"error" example:"Invalid request parameters"` // Сообщение об ошибке
}

// DataResponse структура для успешных ответов с данными
// @Description Структура успешного ответа
type DataResponse struct {
    Data interface{} `json:"data" swaggertype:"object"` // Данные ответа
}

func (f *ResponseBuilder) InternalServerError(ctx *gin.Context, message string) {
    ctx.JSON(http.StatusInternalServerError, Response{Error: message})
}

func (f *ResponseBuilder) UnprocessableEntity(ctx *gin.Context, errorMessage string) {
    ctx.JSON(http.StatusUnprocessableEntity, Response{Error: errorMessage})
}

func (f *ResponseBuilder) BadRequest(ctx *gin.Context, errorMessage string) {
    ctx.JSON(http.StatusBadRequest, Response{Error: errorMessage})
}

func (f *ResponseBuilder) Unauthorized(ctx *gin.Context) {
    ctx.JSON(http.StatusUnauthorized, Response{Error: "Unauthorized."})
}

func (f *ResponseBuilder) UnavailableForLegalReasons(ctx *gin.Context) {
    ctx.JSON(http.StatusUnavailableForLegalReasons, Response{Error: "Unavailable For Legal Reasons."})
}

func (f *ResponseBuilder) NotFound(ctx *gin.Context, errorMessage string) {
    ctx.JSON(http.StatusNotFound, Response{Error: errorMessage})
}

func (f *ResponseBuilder) TooManyRequests(ctx *gin.Context) {
    ctx.JSON(http.StatusTooManyRequests, Response{Error: "Too Many Requests."})
}

func (f *ResponseBuilder) Ok(ctx *gin.Context, data any) {
    ctx.JSON(http.StatusOK, Response{Data: data})
}

func (f *ResponseBuilder) NoContent(ctx *gin.Context, data any) {
    ctx.JSON(http.StatusNoContent, Response{Data: data})
}
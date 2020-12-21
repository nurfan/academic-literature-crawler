package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateArticle for transaction with midtrans
type CreateArticle struct{}

// Handle is a handler for creating transaction with midtrans
func (h *CreateArticle) Handle(c echo.Context) (err error) {
	//ctx := c.Request().Context()

	// r := new(registration.CreateArticleRequest)
	// err = h.validate(r, c)
	// if err != nil {
	// 	return
	// }

	// resp := qHttp.Response{
	// 	Data: map[string]interface{}{
	// 		"status": grpcResp.Status,
	// 	},
	// 	Code:    codes.Success,
	// 	Message: codes.StatusMessage[codes.Success],
	// 	TraceID: qHttp.GetTraceID(ctx),
	// }

	return c.JSON(http.StatusCreated, "Tested Create Article PASS")
}

// func (h *CreateArticle) validate(r *registration.CreateArticleRequest, c echo.Context) error {
// 	if err := c.Bind(r); err != nil {
// 		return status.Errorf(codes.InvalidArgument, err.Error())
// 	}
// 	return c.Validate(r)
// }

// NewCreateArticle for initiate create article handler
func NewCreateArticle() *CreateArticle {
	return &CreateArticle{}
}

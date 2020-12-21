package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ReadArticle for transaction with midtrans
type ReadArticle struct{}

// Handle is a handler for creating transaction with midtrans
func (h *ReadArticle) Handle(c echo.Context) (err error) {
	//ctx := c.Request().Context()

	// r := new(registration.ReadArticleRequest)
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

	return c.JSON(http.StatusCreated, "Tested Read Article PASS")
}

// func (h *ReadArticle) validate(r *registration.ReadArticleRequest, c echo.Context) error {
// 	if err := c.Bind(r); err != nil {
// 		return status.Errorf(codes.InvalidArgument, err.Error())
// 	}
// 	return c.Validate(r)
// }

// NewReadArticle for initiate read article handler
func NewReadArticle() *ReadArticle {
	return &ReadArticle{}
}

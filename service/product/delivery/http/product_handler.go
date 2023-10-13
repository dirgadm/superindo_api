package http

import (
	"project-version3/superindo-task/service/domain"

	"github.com/labstack/echo"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ProductHandler  represent the httphandler for Product
type ProductHandler struct {
	PUsecase domain.ProductUsecase
}

// NewProductHandler will initialize the Products/ resources endpoint
func NewProductHandler(e *echo.Echo, ps domain.ProductUsecase) {
	handler := &ProductHandler{
		PUsecase: ps,
	}
	e.GET("/product", handler.FetchProduct)
}

// FetchProduct will fetch the Product based on given params
func (a *ProductHandler) FetchProduct(c echo.Context) error {
	var err error
	// numS := c.QueryParam("num")
	// num, _ := strconv.Atoi(numS)
	// cursor := c.QueryParam("cursor")
	// ctx := c.Request().Context()

	// listAr, nextCursor, err := a.PUsecase.Fetch(ctx, cursor, int64(num))
	// if err != nil {
	// 	// return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	// }

	// // fmt.Println("ABCDEEEEEE", nextCursor)
	// c.Response().Header().Set(`X-Cursor`, nextCursor)
	// // fmt.Println("ABCDEEEEEEsad", nextCursor)
	// return c.JSON(http.StatusOK, listAr)
	return err
}

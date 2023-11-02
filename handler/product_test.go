package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mockErrorPkg "base-gin-golang/mock/pkg/errors"
	mockProduct "base-gin-golang/mock/usecase/product"
	customErrors "base-gin-golang/pkg/errors/custom"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductUseCase := mockProduct.NewMockUseCase(ctrl)
	mockErrorService := mockErrorPkg.NewMockService(ctrl)
	t.Run("Test", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		CreateProduct(c, mockProductUseCase, mockErrorService)
		if w.Code != http.StatusBadRequest {
			t.Errorf("test create fail")
		}
	})
	mockProductUseCase.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("Fail"))
	mockErrorService.EXPECT().
		ParseInternalServer(gomock.Any()).
		Return(&customErrors.InternalServerError{
			HTTPCode: http.StatusInternalServerError,
			Code:     "Internal server error",
			Message:  "",
		})
	t.Run("Test create fail", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(
			http.MethodPost,
			"/api/products",
			bytes.NewReader([]byte(`
				{
					"productCode":"123",
					"productName":"test",
					"price":123
				}
			`)),
		)
		CreateProduct(c, mockProductUseCase, mockErrorService)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("test create fail")
		}
	})
}

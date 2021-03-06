package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bino1490/case-study/pkg/service"
	gomock "github.com/golang/mock/gomock"
)

func TestInMemory_Success(t *testing.T) {

	// Initialize
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSer := service.NewMockMemoryService(mockCtrl)

	req := httptest.NewRequest(http.MethodGet, "/in-memory", nil)
	res := httptest.NewRecorder()
	mockSer.EXPECT().InMemGetPOST(res, req).AnyTimes().Do(req)

	handler := http.Handler(InMemReqHandler(service.NewInMemService()))
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

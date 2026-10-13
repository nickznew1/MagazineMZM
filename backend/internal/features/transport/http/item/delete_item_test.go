package item_transport_http

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	mock_item_transport_http "github.com/nickznew1/MagazineMZM/backend/internal/features/transport/http/item/mocks"
	"go.uber.org/mock/gomock"
)

func TestDeleteItem(t *testing.T) {
	body := `{"id":4, "name": "ТЕСТОВЫЙ ПРЕДМЕТ"}`
	var item model.Item
	request := httptest.NewRequest(http.MethodDelete, "/item/delete", strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	log := slog.Default()

	logger := &core_logger.Logger{Logger: log}

	ctx := context.WithValue(request.Context(), "log", logger)

	ctrl := gomock.NewController(t)

	mock := mock_item_transport_http.NewMockItemService(ctrl)

	mock.EXPECT().DeleteItem(gomock.Any(), gomock.Any()).Return(model.Item{Id: 1, Name: "ТЕСТОВЫЙ ПРЕДМЕТ"}, nil)

	handler := &ItemHTTPHandler{itemService: mock}

	handler.DeleteItem(recorder, request.WithContext(ctx))

	if recorder.Code != http.StatusOK {
		t.Errorf("failed: get wrong server response, excepted ='%d', got ='%d'", http.StatusOK, recorder.Code)
	}

	json.NewDecoder(recorder.Body).Decode(&item)
	if item.Id != 1 {
		t.Errorf("failed: excepted response id='%d' recieved = '%d'", 4, item.Id)
	}
}

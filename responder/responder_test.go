package responder_test

import (
	"errors"
	"github.com/sudtanj/Go-SwissKnife/responder"
	"github.com/sudtanj/Go-SwissKnife/responder/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestJson(t *testing.T) {
	t.Run("return success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		jsonResp := mocks.NewMockIJSONResponder(ctrl)

		jsonResp.EXPECT().JSON(200, "test")

		responder.Json(jsonResp, errors.New("test"))
	})
}

package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/kdnakt/go_todo_app/testutil"
)

func TestLogin_ServeHTTP(t *testing.T) {
	type moq struct {
		token string
		err   error
	}
	tests := map[string]struct {
		moq    moq
		status int
	}{
		"ok": {
			moq: moq{
				token: "from_moq",
			},
			status: http.StatusOK,
		},
		"bad": {
			status: http.StatusBadRequest,
		},
		"error": {
			moq: moq{
				err: errors.New("error from mock"),
			},
			status: http.StatusInternalServerError,
		},
	}
	for n, tt := range tests {
		n := n
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/login",
				bytes.NewReader(testutil.LoadFile(t, fmt.Sprintf("testdata/login/%s_req.json.golden", n))),
			)

			moq := &LoginServiceMock{}
			moq.LoginFunc = func(ctx context.Context, name, pw string) (string, error) {
				return tt.moq.token, tt.moq.err
			}
			sut := Login{
				Service:   moq,
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.status, testutil.LoadFile(t, fmt.Sprintf("testdata/login/%s_rsp.json.golden", n)),
			)
		})
	}
}

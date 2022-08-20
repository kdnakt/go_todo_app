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
	"github.com/kdnakt/go_todo_app/entity"
	"github.com/kdnakt/go_todo_app/testutil"
)

func TestAddTask(t *testing.T) {
	t.Parallel()
	tests := map[string]int{
		"ok":  http.StatusOK,
		"bad": http.StatusBadRequest,
	}
	for n, tt := range tests {
		// Avoid to be overwritten by parallel testing
		tt := tt
		n := n
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				bytes.NewReader(testutil.LoadFile(t, fmt.Sprintf("testdata/add_task/%s_req.json.golden", n))),
			)

			moq := &AddTaskServiceMock{}
			moq.AddTaskFunc = func(
				ctx context.Context, title string,
			) (*entity.Task, error) {
				if tt == http.StatusOK {
					return &entity.Task{ID: 1}, nil
				}
				return nil, errors.New("error from mock")
			}
			sut := AddTask{
				Service:   moq,
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt, testutil.LoadFile(t, fmt.Sprintf("testdata/add_task/%s_rsp.json.golden", n)),
			)
		})
	}
}

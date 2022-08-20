package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kdnakt/go_todo_app/entity"
	"github.com/kdnakt/go_todo_app/testutil"
)

func TestListTask(t *testing.T) {
	t.Parallel()
	tests := map[string]int{
		"ok": http.StatusOK,
	}
	for n, tt := range tests {
		// Avoid to be overwritten by parallel testing
		tt := tt
		n := n
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodGet,
				"/tasks",
				nil,
			)

			moq := &ListTasksServiceMock{}
			moq.ListTasksFunc = func(ctx context.Context) (entity.Tasks, error) {
				if tt == http.StatusOK {
					return entity.Tasks{
						{
							ID:     1,
							Title:  "test1",
							Status: "todo",
						},
						{
							ID:     2,
							Title:  "test2",
							Status: "done",
						},
					}, nil
				}
				return nil, errors.New("error from mock")
			}
			sut := ListTask{
				Service: moq,
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt, testutil.LoadFile(t, fmt.Sprintf("testdata/list_task/%s_rsp.json.golden", n)),
			)
		})
	}
}

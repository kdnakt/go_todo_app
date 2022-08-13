package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kdnakt/go_todo_app/entity"
	"github.com/kdnakt/go_todo_app/store"
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

			sut := ListTask{
				Store: &store.TaskStore{
					Tasks: map[entity.TaskID]*entity.Task{
						1: {
							ID:     1,
							Title:  "test1",
							Status: "todo",
						},
						2: {
							ID:     2,
							Title:  "test2",
							Status: "done",
						},
					},
				},
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt, testutil.LoadFile(t, fmt.Sprintf("testdata/list_task/%s_rsp.json.golden", n)),
			)
		})
	}
}

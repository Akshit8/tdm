package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Akshit8/tdm/internal"
	"github.com/Akshit8/tdm/internal/rest"
	"github.com/Akshit8/tdm/internal/service/servicetesting"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(*servicetesting.FakeTaskService)
		input  []byte
		output output
	}{
		{
			"OK: 201",
			func(s *servicetesting.FakeTaskService) {
				s.CreateReturns(
					internal.Task{
						ID:          "1-2-3",
						Description: "new task",
						Priority:    internal.PriorityHigh,
						Dates: internal.Dates{
							Start: time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC),
							Due:   time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC).Add(time.Hour),
						},
						IsDone: false,
					},
					nil)
			},
			func() []byte {
				b, _ := json.Marshal(&rest.CreateTasksRequest{
					Description: "new task",
					Priority:    rest.Priority("high"),
					Dates: rest.Dates{
						Start: rest.Time(time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC)),
						Due:   rest.Time(time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC).Add(time.Hour)),
					},
				})

				return b
			}(),
			output{
				http.StatusCreated,
				&rest.CreateTasksResponse{
					Task: rest.Task{
						ID:          "1-2-3",
						Description: "new task",
						Priority:    rest.Priority("high"),
						Dates: rest.Dates{
							Start: rest.Time(time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC)),
							Due:   rest.Time(time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC).Add(time.Hour)),
						},
						IsDone: false,
					},
				},
				&rest.CreateTasksResponse{},
			},
		},
		{
			"ERR: 400",
			func(s *servicetesting.FakeTaskService) {},
			[]byte(`{"invalid":"json`),
			output{
				http.StatusBadRequest,
				&rest.ErrorResponse{
					Status: http.StatusBadRequest,
					Error:  "invalid request",
				},
				&rest.ErrorResponse{},
			},
		},
		{
			"ERR: 500",
			func(s *servicetesting.FakeTaskService) {
				s.CreateReturns(internal.Task{}, errors.New("service error"))
			},
			[]byte(`{}`),
			output{
				http.StatusInternalServerError,
				&rest.ErrorResponse{
					Status: http.StatusInternalServerError,
					Error:  "create failed",
				},
				&rest.ErrorResponse{},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &servicetesting.FakeTaskService{}
			test.setup(svc)

			rest.NewTaskHandler(svc).Register(router)

			res := doRequest(
				router,
				httptest.NewRequest(
					http.MethodPost,
					"/tasks",
					bytes.NewReader(test.input),
				),
			)

			require.Equal(t, test.output.expectedStatus, res.StatusCode)

			assertResponse(t, res, testOuptut{test.output.expected, test.output.target})
		})
	}
}

func TestTask(t *testing.T) {
	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(*servicetesting.FakeTaskService)
		output output
	}{
		{
			"OK: 200",
			func(s *servicetesting.FakeTaskService) {
				s.TaskReturns(
					internal.Task{
						ID:          "a-b-c",
						Description: "existing task",
						Priority:    internal.PriorityLow,
						Dates: internal.Dates{
							Start: time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC),
							Due:   time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC).Add(time.Hour),
						},
						IsDone: true,
					},
					nil)
			},
			output{
				http.StatusOK,
				&rest.GetTasksResponse{
					Task: rest.Task{
						ID:          "a-b-c",
						Description: "existing task",
						Priority:    rest.Priority("low"),
						Dates: rest.Dates{
							Start: rest.Time(time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC)),
							Due:   rest.Time(time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC).Add(time.Hour)),
						},
						IsDone: true,
					},
				},
				&rest.GetTasksResponse{},
			},
		},
		{
			"ERR: 500",
			func(s *servicetesting.FakeTaskService) {
				s.TaskReturns(internal.Task{}, errors.New("service error"))
			},
			output{
				http.StatusInternalServerError,
				&rest.ErrorResponse{
					Status: http.StatusInternalServerError,
					Error:  "find failed",
				},
				&rest.ErrorResponse{},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &servicetesting.FakeTaskService{}
			test.setup(svc)

			rest.NewTaskHandler(svc).Register(router)

			res := doRequest(
				router,
				httptest.NewRequest(
					http.MethodGet,
					fmt.Sprintf("/tasks/%s", uuid.NewString()),
					nil,
				),
			)

			require.Equal(t, test.output.expectedStatus, res.StatusCode)

			assertResponse(t, res, testOuptut{test.output.expected, test.output.target})
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(*servicetesting.FakeTaskService)
		input  []byte
		output output
	}{
		{
			"OK: 200",
			func(s *servicetesting.FakeTaskService) {},
			func() []byte {
				b, _ := json.Marshal(&rest.UpdateTasksRequest{
					Description: "update task",
					Priority:    rest.Priority("low"),
					Dates: rest.Dates{
						Start: rest.Time(time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC)),
						Due:   rest.Time(time.Date(2009, 11, 19, 23, 0, 0, 0, time.UTC).Add(time.Hour)),
					},
					IsDone: true,
				})

				return b
			}(),
			output{
				http.StatusOK,
				func() *string {
					a := "task updated"
					return &a
				}(),
				new(string),
			},
		},
		{
			"ERR: 400",
			func(s *servicetesting.FakeTaskService) {},
			[]byte(`{"invalid":"json`),
			output{
				http.StatusBadRequest,
				&rest.ErrorResponse{
					Status: http.StatusBadRequest,
					Error:  "invalid request",
				},
				&rest.ErrorResponse{},
			},
		},
		{
			"ERR: 500",
			func(s *servicetesting.FakeTaskService) {
				s.UpdateReturns(errors.New("service error"))
			},
			[]byte(`{}`),
			output{
				http.StatusInternalServerError,
				&rest.ErrorResponse{
					Status: http.StatusInternalServerError,
					Error:  "update failed",
				},
				&rest.ErrorResponse{},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &servicetesting.FakeTaskService{}
			test.setup(svc)

			rest.NewTaskHandler(svc).Register(router)

			res := doRequest(
				router,
				httptest.NewRequest(
					http.MethodPut,
					fmt.Sprintf("/tasks/%s", uuid.NewString()),
					bytes.NewReader(test.input),
				),
			)

			require.Equal(t, test.output.expectedStatus, res.StatusCode)

			assertResponse(t, res, testOuptut{test.output.expected, test.output.target})
		})
	}
}

type testOuptut struct {
	expected interface{}
	target   interface{}
}

func doRequest(router *mux.Router, req *http.Request) *http.Response {
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	return rr.Result()
}

func assertResponse(t *testing.T, res *http.Response, test testOuptut) {
	// https://stackoverflow.com/questions/39194816/how-to-wrap-golang-test-functions
	t.Helper()

	err := json.NewDecoder(res.Body).Decode(test.target)
	require.NoError(t, err)
	defer res.Body.Close()

	if !cmp.Equal(test.expected, test.target, cmp.AllowUnexported(rest.Time{})) {
		t.Fatalf("expected results don't match: %s", cmp.Diff(test.expected, test.target, cmp.AllowUnexported(rest.Time{})))
	}

}

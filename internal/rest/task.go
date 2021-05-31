package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Akshit8/tdm/internal"
	"github.com/Akshit8/tdm/internal/service"
	"github.com/gorilla/mux"
)

const uuidRegEx string = `[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}`

// TaskHandler ...
type TaskHandler struct {
	svc service.TaskService
}

func NewTaskHandler(svc service.TaskService) *TaskHandler {
	return &TaskHandler{
		svc: svc,
	}
}

// Register connects the handlers to the router.
func (t *TaskHandler) Register(r *mux.Router) {
	r.HandleFunc("/tasks", t.create).Methods(http.MethodPost)
	r.HandleFunc(fmt.Sprintf("/tasks/{id:%s}", uuidRegEx), t.task).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("/tasks/{id:%s}", uuidRegEx), t.update).Methods(http.MethodPut)
}

// Task is an activity that needs to be completed within a period of time.
type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// CreateTasksRequest defines the request used for creating tasks.
type CreateTasksRequest struct {
	Description string `json:"description"`
}

// CreateTasksResponse defines the response returned back after creating tasks.
type CreateTasksResponse struct {
	Task Task `json:"task"`
}

func (t *TaskHandler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateTasksRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		renderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	task, err := t.svc.Create(r.Context(), req.Description, internal.PriorityMedium, internal.Dates{})
	if err != nil {
		renderErrorResponse(w, "create failed", http.StatusInternalServerError)
		return
	}

	renderResponse(w,
		&CreateTasksResponse{
			Task: Task{
				ID:          task.ID,
				Description: task.Description,
			},
		}, http.StatusCreated)
}

// GetTasksResponse defines the response returned back after searching one task.
type GetTasksResponse struct {
	Task Task `json:"task"`
}

func (t *TaskHandler) task(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := t.svc.Task(r.Context(), id)
	if err != nil {
		renderErrorResponse(w, "find failed", http.StatusInternalServerError)
		return
	}

	renderResponse(w,
		&GetTasksResponse{
			Task: Task{
				ID:          task.ID,
				Description: task.Description,
			},
		}, http.StatusOK)
}

// UpdateTasksRequest defines the request used for updating a task.
type UpdateTasksRequest struct {
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func (t *TaskHandler) update(w http.ResponseWriter, r *http.Request) {
	var req UpdateTasksRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		renderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	id := mux.Vars(r)["id"]

	err = t.svc.Update(r.Context(), id, req.Description, internal.PriorityMedium, internal.Dates{}, req.IsDone)
	if err != nil {
		renderErrorResponse(w, "update failed", http.StatusInternalServerError)
		return
	}

	renderResponse(w, "task updated", http.StatusOK)
}

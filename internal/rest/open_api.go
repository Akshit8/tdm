package rest

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

//go:generate oapi-codegen -package openapi3 -generate types  -o ../../pkg/openapi3/task_types.gen.go openapi3.yaml
//go:generate oapi-codegen -package openapi3 -generate client -o ../../pkg/openapi3/client.gen.go     openapi3.yaml

// NewOpenAPI3 instantiates the OpenAPI specification for this service.
func NewOpenAPI3() openapi3.T {
	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Todo API",
			Description: "REST API to create and manage TODO's",
			Version:     "0.0.1",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				URL: "https://github.com/Akshit8/tdm",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://127.0.0.1:8000",
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"Priority": openapi3.NewSchemaRef("",
			openapi3.NewStringSchema().
				WithEnum("none", "low", "medium", "high").
				WithDefault("none"),
		),
		"Dates": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("start", openapi3.NewStringSchema().WithFormat("date-time").WithNullable()).
				WithProperty("due", openapi3.NewStringSchema().WithFormat("date-time").WithNullable()),
		),
		"Task": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewUUIDSchema()).
				WithProperty("description", openapi3.NewStringSchema()).
				WithPropertyRef("priority", &openapi3.SchemaRef{Ref: "#/components/schemas/Priority"}).
				WithPropertyRef("dates", &openapi3.SchemaRef{Ref: "#/components/schemas/Dates"}).
				WithProperty("is_done", openapi3.NewBoolSchema()),
		),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CreateTasksRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("request to create a new task.").
				WithRequired(true).
				WithJSONSchema(
					openapi3.NewSchema().
						WithProperty("description", openapi3.NewStringSchema().WithMinLength(1)).
						WithPropertyRef("priority", &openapi3.SchemaRef{Ref: "#/components/schemas/Priority"}).
						WithPropertyRef("dates", &openapi3.SchemaRef{Ref: "#/components/schemas/Dates"}),
				),
		},
		"UpdateTasksRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("request to update an existing task.").
				WithRequired(true).
				WithJSONSchema(
					openapi3.NewSchema().
						WithProperty("description", openapi3.NewStringSchema().WithMinLength(1)).
						WithPropertyRef("priority", &openapi3.SchemaRef{Ref: "#/components/schemas/Priority"}).
						WithPropertyRef("dates", &openapi3.SchemaRef{Ref: "#/components/schemas/Dates"}).
						WithProperty("is_done", openapi3.NewBoolSchema()),
				),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when errors happen.").
				WithContent(openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithProperty("status", openapi3.NewIntegerSchema()).
						WithProperty("error", openapi3.NewStringSchema()),
				),
				),
		},
		"CreateTasksResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after creating tasks.").
				WithContent(openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithPropertyRef("task", &openapi3.SchemaRef{Ref: "#/components/schemas/Task"}),
				),
				),
		},
		"ReadTasksResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after searching one task.").
				WithContent(openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithPropertyRef("task", &openapi3.SchemaRef{Ref: "#/components/schemas/Task"}),
				),
				),
		},
	}

	swagger.Paths = openapi3.Paths{
		"/tasks": &openapi3.PathItem{
			Post: &openapi3.Operation{
				OperationID: "CreateTask",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CreateTasksRequest",
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/CreateTasksResponse",
					},
				},
			},
		},
		"/tasks/{taskId}": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "ReadTask",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("taskId").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/ReadTasksResponse",
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Task not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
			Put: &openapi3.Operation{
				OperationID: "UpdateTask",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("taskId").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/UpdateTasksRequest",
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Task updated"),
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Task not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
		},
	}

	return swagger
}

// RegisterOpenAPI registers open API config files to mux router.
func RegisterOpenAPI(r *mux.Router) {
	swagger := NewOpenAPI3()

	r.HandleFunc("/openapi3.json", func(w http.ResponseWriter, r *http.Request) {
		renderResponse(w, &swagger, http.StatusOK)
	}).Methods(http.MethodGet)

	r.HandleFunc("/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")

		data, _ := yaml.Marshal(&swagger) // NOTE: safe to ignore error

		_, _ = w.Write(data)

		w.WriteHeader(http.StatusOK)

	}).Methods(http.MethodGet)
}

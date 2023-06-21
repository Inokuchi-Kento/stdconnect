// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: task/v1/task.proto

package taskv1connect

import (
	v1 "blrpc/gen/task/v1"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TaskServiceName is the fully-qualified name of the TaskService service.
	TaskServiceName = "task.v1.TaskService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TaskServiceListTasksProcedure is the fully-qualified name of the TaskService's ListTasks RPC.
	TaskServiceListTasksProcedure = "/task.v1.TaskService/ListTasks"
	// TaskServiceCreateTaskProcedure is the fully-qualified name of the TaskService's CreateTask RPC.
	TaskServiceCreateTaskProcedure = "/task.v1.TaskService/CreateTask"
)

// TaskServiceClient is a client for the task.v1.TaskService service.
type TaskServiceClient interface {
	ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error)
	CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error)
}

// NewTaskServiceClient constructs a client for the task.v1.TaskService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTaskServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TaskServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &taskServiceClient{
		listTasks: connect_go.NewClient[v1.ListTasksRequest, v1.ListTasksResponse](
			httpClient,
			baseURL+TaskServiceListTasksProcedure,
			opts...,
		),
		createTask: connect_go.NewClient[v1.CreateTaskRequest, v1.CreateTaskResponse](
			httpClient,
			baseURL+TaskServiceCreateTaskProcedure,
			opts...,
		),
	}
}

// taskServiceClient implements TaskServiceClient.
type taskServiceClient struct {
	listTasks  *connect_go.Client[v1.ListTasksRequest, v1.ListTasksResponse]
	createTask *connect_go.Client[v1.CreateTaskRequest, v1.CreateTaskResponse]
}

// ListTasks calls task.v1.TaskService.ListTasks.
func (c *taskServiceClient) ListTasks(ctx context.Context, req *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error) {
	return c.listTasks.CallUnary(ctx, req)
}

// CreateTask calls task.v1.TaskService.CreateTask.
func (c *taskServiceClient) CreateTask(ctx context.Context, req *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error) {
	return c.createTask.CallUnary(ctx, req)
}

// TaskServiceHandler is an implementation of the task.v1.TaskService service.
type TaskServiceHandler interface {
	ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error)
	CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error)
}

// NewTaskServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTaskServiceHandler(svc TaskServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(TaskServiceListTasksProcedure, connect_go.NewUnaryHandler(
		TaskServiceListTasksProcedure,
		svc.ListTasks,
		opts...,
	))
	mux.Handle(TaskServiceCreateTaskProcedure, connect_go.NewUnaryHandler(
		TaskServiceCreateTaskProcedure,
		svc.CreateTask,
		opts...,
	))
	return "/task.v1.TaskService/", mux
}

// UnimplementedTaskServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTaskServiceHandler struct{}

func (UnimplementedTaskServiceHandler) ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("task.v1.TaskService.ListTasks is not implemented"))
}

func (UnimplementedTaskServiceHandler) CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("task.v1.TaskService.CreateTask is not implemented"))
}

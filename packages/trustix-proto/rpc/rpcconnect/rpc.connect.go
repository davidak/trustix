// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: rpc/rpc.proto

package rpcconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	api "github.com/nix-community/trustix/packages/trustix-proto/api"
	rpc "github.com/nix-community/trustix/packages/trustix-proto/rpc"
	schema "github.com/nix-community/trustix/packages/trustix-proto/schema"
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
	// RPCApiName is the fully-qualified name of the RPCApi service.
	RPCApiName = "trustix.RPCApi"
	// LogRPCName is the fully-qualified name of the LogRPC service.
	LogRPCName = "trustix.LogRPC"
)

// RPCApiClient is a client for the trustix.RPCApi service.
type RPCApiClient interface {
	// Get a list of all logs published/subscribed by this node
	Logs(context.Context, *connect_go.Request[api.LogsRequest]) (*connect_go.Response[api.LogsResponse], error)
	// Decide on an output for key based on the configured decision method
	Decide(context.Context, *connect_go.Request[rpc.DecideRequest]) (*connect_go.Response[rpc.DecisionResponse], error)
	// Get values by their content-address
	GetValue(context.Context, *connect_go.Request[api.ValueRequest]) (*connect_go.Response[api.ValueResponse], error)
}

// NewRPCApiClient constructs a client for the trustix.RPCApi service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRPCApiClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) RPCApiClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &rPCApiClient{
		logs: connect_go.NewClient[api.LogsRequest, api.LogsResponse](
			httpClient,
			baseURL+"/trustix.RPCApi/Logs",
			opts...,
		),
		decide: connect_go.NewClient[rpc.DecideRequest, rpc.DecisionResponse](
			httpClient,
			baseURL+"/trustix.RPCApi/Decide",
			opts...,
		),
		getValue: connect_go.NewClient[api.ValueRequest, api.ValueResponse](
			httpClient,
			baseURL+"/trustix.RPCApi/GetValue",
			opts...,
		),
	}
}

// rPCApiClient implements RPCApiClient.
type rPCApiClient struct {
	logs     *connect_go.Client[api.LogsRequest, api.LogsResponse]
	decide   *connect_go.Client[rpc.DecideRequest, rpc.DecisionResponse]
	getValue *connect_go.Client[api.ValueRequest, api.ValueResponse]
}

// Logs calls trustix.RPCApi.Logs.
func (c *rPCApiClient) Logs(ctx context.Context, req *connect_go.Request[api.LogsRequest]) (*connect_go.Response[api.LogsResponse], error) {
	return c.logs.CallUnary(ctx, req)
}

// Decide calls trustix.RPCApi.Decide.
func (c *rPCApiClient) Decide(ctx context.Context, req *connect_go.Request[rpc.DecideRequest]) (*connect_go.Response[rpc.DecisionResponse], error) {
	return c.decide.CallUnary(ctx, req)
}

// GetValue calls trustix.RPCApi.GetValue.
func (c *rPCApiClient) GetValue(ctx context.Context, req *connect_go.Request[api.ValueRequest]) (*connect_go.Response[api.ValueResponse], error) {
	return c.getValue.CallUnary(ctx, req)
}

// RPCApiHandler is an implementation of the trustix.RPCApi service.
type RPCApiHandler interface {
	// Get a list of all logs published/subscribed by this node
	Logs(context.Context, *connect_go.Request[api.LogsRequest]) (*connect_go.Response[api.LogsResponse], error)
	// Decide on an output for key based on the configured decision method
	Decide(context.Context, *connect_go.Request[rpc.DecideRequest]) (*connect_go.Response[rpc.DecisionResponse], error)
	// Get values by their content-address
	GetValue(context.Context, *connect_go.Request[api.ValueRequest]) (*connect_go.Response[api.ValueResponse], error)
}

// NewRPCApiHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRPCApiHandler(svc RPCApiHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/trustix.RPCApi/Logs", connect_go.NewUnaryHandler(
		"/trustix.RPCApi/Logs",
		svc.Logs,
		opts...,
	))
	mux.Handle("/trustix.RPCApi/Decide", connect_go.NewUnaryHandler(
		"/trustix.RPCApi/Decide",
		svc.Decide,
		opts...,
	))
	mux.Handle("/trustix.RPCApi/GetValue", connect_go.NewUnaryHandler(
		"/trustix.RPCApi/GetValue",
		svc.GetValue,
		opts...,
	))
	return "/trustix.RPCApi/", mux
}

// UnimplementedRPCApiHandler returns CodeUnimplemented from all methods.
type UnimplementedRPCApiHandler struct{}

func (UnimplementedRPCApiHandler) Logs(context.Context, *connect_go.Request[api.LogsRequest]) (*connect_go.Response[api.LogsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("trustix.RPCApi.Logs is not implemented"))
}

func (UnimplementedRPCApiHandler) Decide(context.Context, *connect_go.Request[rpc.DecideRequest]) (*connect_go.Response[rpc.DecisionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("trustix.RPCApi.Decide is not implemented"))
}

func (UnimplementedRPCApiHandler) GetValue(context.Context, *connect_go.Request[api.ValueRequest]) (*connect_go.Response[api.ValueResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("trustix.RPCApi.GetValue is not implemented"))
}

// LogRPCClient is a client for the trustix.LogRPC service.
type LogRPCClient interface {
	GetHead(context.Context, *connect_go.Request[api.LogHeadRequest]) (*connect_go.Response[schema.LogHead], error)
	GetLogEntries(context.Context, *connect_go.Request[api.GetLogEntriesRequest]) (*connect_go.Response[api.LogEntriesResponse], error)
	Submit(context.Context, *connect_go.Request[rpc.SubmitRequest]) (*connect_go.Response[rpc.SubmitResponse], error)
	Flush(context.Context, *connect_go.Request[rpc.FlushRequest]) (*connect_go.Response[rpc.FlushResponse], error)
}

// NewLogRPCClient constructs a client for the trustix.LogRPC service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLogRPCClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) LogRPCClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &logRPCClient{
		getHead: connect_go.NewClient[api.LogHeadRequest, schema.LogHead](
			httpClient,
			baseURL+"/trustix.LogRPC/GetHead",
			opts...,
		),
		getLogEntries: connect_go.NewClient[api.GetLogEntriesRequest, api.LogEntriesResponse](
			httpClient,
			baseURL+"/trustix.LogRPC/GetLogEntries",
			opts...,
		),
		submit: connect_go.NewClient[rpc.SubmitRequest, rpc.SubmitResponse](
			httpClient,
			baseURL+"/trustix.LogRPC/Submit",
			opts...,
		),
		flush: connect_go.NewClient[rpc.FlushRequest, rpc.FlushResponse](
			httpClient,
			baseURL+"/trustix.LogRPC/Flush",
			opts...,
		),
	}
}

// logRPCClient implements LogRPCClient.
type logRPCClient struct {
	getHead       *connect_go.Client[api.LogHeadRequest, schema.LogHead]
	getLogEntries *connect_go.Client[api.GetLogEntriesRequest, api.LogEntriesResponse]
	submit        *connect_go.Client[rpc.SubmitRequest, rpc.SubmitResponse]
	flush         *connect_go.Client[rpc.FlushRequest, rpc.FlushResponse]
}

// GetHead calls trustix.LogRPC.GetHead.
func (c *logRPCClient) GetHead(ctx context.Context, req *connect_go.Request[api.LogHeadRequest]) (*connect_go.Response[schema.LogHead], error) {
	return c.getHead.CallUnary(ctx, req)
}

// GetLogEntries calls trustix.LogRPC.GetLogEntries.
func (c *logRPCClient) GetLogEntries(ctx context.Context, req *connect_go.Request[api.GetLogEntriesRequest]) (*connect_go.Response[api.LogEntriesResponse], error) {
	return c.getLogEntries.CallUnary(ctx, req)
}

// Submit calls trustix.LogRPC.Submit.
func (c *logRPCClient) Submit(ctx context.Context, req *connect_go.Request[rpc.SubmitRequest]) (*connect_go.Response[rpc.SubmitResponse], error) {
	return c.submit.CallUnary(ctx, req)
}

// Flush calls trustix.LogRPC.Flush.
func (c *logRPCClient) Flush(ctx context.Context, req *connect_go.Request[rpc.FlushRequest]) (*connect_go.Response[rpc.FlushResponse], error) {
	return c.flush.CallUnary(ctx, req)
}

// LogRPCHandler is an implementation of the trustix.LogRPC service.
type LogRPCHandler interface {
	GetHead(context.Context, *connect_go.Request[api.LogHeadRequest]) (*connect_go.Response[schema.LogHead], error)
	GetLogEntries(context.Context, *connect_go.Request[api.GetLogEntriesRequest]) (*connect_go.Response[api.LogEntriesResponse], error)
	Submit(context.Context, *connect_go.Request[rpc.SubmitRequest]) (*connect_go.Response[rpc.SubmitResponse], error)
	Flush(context.Context, *connect_go.Request[rpc.FlushRequest]) (*connect_go.Response[rpc.FlushResponse], error)
}

// NewLogRPCHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLogRPCHandler(svc LogRPCHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/trustix.LogRPC/GetHead", connect_go.NewUnaryHandler(
		"/trustix.LogRPC/GetHead",
		svc.GetHead,
		opts...,
	))
	mux.Handle("/trustix.LogRPC/GetLogEntries", connect_go.NewUnaryHandler(
		"/trustix.LogRPC/GetLogEntries",
		svc.GetLogEntries,
		opts...,
	))
	mux.Handle("/trustix.LogRPC/Submit", connect_go.NewUnaryHandler(
		"/trustix.LogRPC/Submit",
		svc.Submit,
		opts...,
	))
	mux.Handle("/trustix.LogRPC/Flush", connect_go.NewUnaryHandler(
		"/trustix.LogRPC/Flush",
		svc.Flush,
		opts...,
	))
	return "/trustix.LogRPC/", mux
}

// UnimplementedLogRPCHandler returns CodeUnimplemented from all methods.
type UnimplementedLogRPCHandler struct{}

func (UnimplementedLogRPCHandler) GetHead(context.Context, *connect_go.Request[api.LogHeadRequest]) (*connect_go.Response[schema.LogHead], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("trustix.LogRPC.GetHead is not implemented"))
}

func (UnimplementedLogRPCHandler) GetLogEntries(context.Context, *connect_go.Request[api.GetLogEntriesRequest]) (*connect_go.Response[api.LogEntriesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("trustix.LogRPC.GetLogEntries is not implemented"))
}

func (UnimplementedLogRPCHandler) Submit(context.Context, *connect_go.Request[rpc.SubmitRequest]) (*connect_go.Response[rpc.SubmitResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("trustix.LogRPC.Submit is not implemented"))
}

func (UnimplementedLogRPCHandler) Flush(context.Context, *connect_go.Request[rpc.FlushRequest]) (*connect_go.Response[rpc.FlushResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("trustix.LogRPC.Flush is not implemented"))
}

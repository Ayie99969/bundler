// bundler v0.1.0 fa97054e223493177aebecd1739a101f077728de
// --
// Code generated by webrpc-gen@v0.14.0-dev with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=rpc.ridl -target=golang -pkg=bundler -client -out=./clients/proto.gen.go
package bundler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/0xsequence/go-sequence/lib/prototyp"
)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.1.0"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "fa97054e223493177aebecd1739a101f077728de"
}

//
// Types
//

type MessageType int

const (
	MessageType_DEBUG         MessageType = 0
	MessageType_NEW_OPERATION MessageType = 1
	MessageType_ARCHIVE       MessageType = 2
)

var MessageType_name = map[int]string{
	0: "DEBUG",
	1: "NEW_OPERATION",
	2: "ARCHIVE",
}

var MessageType_value = map[string]int{
	"DEBUG":         0,
	"NEW_OPERATION": 1,
	"ARCHIVE":       2,
}

func (x MessageType) String() string {
	return MessageType_name[int(x)]
}

func (x MessageType) MarshalText() ([]byte, error) {
	return []byte(MessageType_name[int(x)]), nil
}

func (x *MessageType) UnmarshalText(b []byte) error {
	*x = MessageType(MessageType_value[string(b)])
	return nil
}

func (x *MessageType) Is(values ...MessageType) bool {
	if x == nil {
		return false
	}
	for _, v := range values {
		if *x == v {
			return true
		}
	}
	return false
}

type OperationStrategy int

const (
	OperationStrategy_Fresh  OperationStrategy = 0
	OperationStrategy_Greedy OperationStrategy = 1
)

var OperationStrategy_name = map[int]string{
	0: "Fresh",
	1: "Greedy",
}

var OperationStrategy_value = map[string]int{
	"Fresh":  0,
	"Greedy": 1,
}

func (x OperationStrategy) String() string {
	return OperationStrategy_name[int(x)]
}

func (x OperationStrategy) MarshalText() ([]byte, error) {
	return []byte(OperationStrategy_name[int(x)]), nil
}

func (x *OperationStrategy) UnmarshalText(b []byte) error {
	*x = OperationStrategy(OperationStrategy_value[string(b)])
	return nil
}

func (x *OperationStrategy) Is(values ...OperationStrategy) bool {
	if x == nil {
		return false
	}
	for _, v := range values {
		if *x == v {
			return true
		}
	}
	return false
}

type ReadyAtChange int

const (
	ReadyAtChange_None ReadyAtChange = 0
	ReadyAtChange_Now  ReadyAtChange = 1
	ReadyAtChange_Zero ReadyAtChange = 2
)

var ReadyAtChange_name = map[int]string{
	0: "None",
	1: "Now",
	2: "Zero",
}

var ReadyAtChange_value = map[string]int{
	"None": 0,
	"Now":  1,
	"Zero": 2,
}

func (x ReadyAtChange) String() string {
	return ReadyAtChange_name[int(x)]
}

func (x ReadyAtChange) MarshalText() ([]byte, error) {
	return []byte(ReadyAtChange_name[int(x)]), nil
}

func (x *ReadyAtChange) UnmarshalText(b []byte) error {
	*x = ReadyAtChange(ReadyAtChange_value[string(b)])
	return nil
}

func (x *ReadyAtChange) Is(values ...ReadyAtChange) bool {
	if x == nil {
		return false
	}
	for _, v := range values {
		if *x == v {
			return true
		}
	}
	return false
}

type Version struct {
	WebrpcVersion string `json:"webrpcVersion"`
	SchemaVersion string `json:"schemaVersion"`
	SchemaHash    string `json:"schemaHash"`
	NodeVersion   string `json:"nodeVersion"`
}

type Status struct {
	HealthOK      bool      `json:"healthOK"`
	StartTime     time.Time `json:"startTime"`
	Uptime        uint64    `json:"uptime"`
	Ver           string    `json:"ver"`
	Branch        string    `json:"branch"`
	CommitHash    string    `json:"commitHash"`
	HostID        string    `json:"hostId"`
	HostAddrs     []string  `json:"hostAddrs"`
	PriorityPeers []string  `json:"priorityPeers"`
}

type Operation struct {
	Entrypoint                 prototyp.Hash   `json:"entrypoint"`
	CallData                   prototyp.Hash   `json:"callData"`
	GasLimit                   prototyp.BigInt `json:"gasLimit"`
	FeeToken                   prototyp.Hash   `json:"feeToken"`
	Endorser                   prototyp.Hash   `json:"endorser"`
	EndorserCallData           prototyp.Hash   `json:"endorserCallData"`
	EndorserGasLimit           prototyp.BigInt `json:"endorserGasLimit"`
	MaxFeePerGas               prototyp.BigInt `json:"maxFeePerGas"`
	PriorityFeePerGas          prototyp.BigInt `json:"priorityFeePerGas"`
	BaseFeeScalingFactor       prototyp.BigInt `json:"baseFeeScalingFactor"`
	BaseFeeNormalizationFactor prototyp.BigInt `json:"baseFeeNormalizationFactor"`
	HasUntrustedContext        bool            `json:"hasUntrustedContext"`
	ChainID                    prototyp.BigInt `json:"chainId"`
	Hash                       *string         `json:"hash,omitempty"`
}

type Message struct {
	Type    MessageType `json:"type"`
	Message interface{} `json:"message"`
}

type MempoolView struct {
	Size       int         `json:"size"`
	SeenSize   int         `json:"seenSize"`
	LockSize   int         `json:"lockSize"`
	Seen       []string    `json:"seen"`
	Operations interface{} `json:"operations"`
}

type Operations struct {
	Mempool []string `json:"mempool"`
	Archive string   `json:"archive,omitempty"`
}

type BaseFeeRate struct {
	ScalingFactor       prototyp.BigInt `json:"scalingFactor"`
	NormalizationFactor prototyp.BigInt `json:"normalizationFactor"`
}

type FeeAsks struct {
	MinBaseFee     prototyp.BigInt        `json:"minBaseFee"`
	MinPriorityFee prototyp.BigInt        `json:"minPriorityFee"`
	AcceptedTokens map[string]BaseFeeRate `json:"acceptedTokens"`
}

type Bundler interface {
	Ping(ctx context.Context) (bool, error)
	Status(ctx context.Context) (*Status, error)
	Peers(ctx context.Context) ([]string, []string, error)
	Mempool(ctx context.Context) (*MempoolView, error)
	SendOperation(ctx context.Context, operation *Operation) (string, error)
	Operations(ctx context.Context) (*Operations, error)
	FeeAsks(ctx context.Context) (*FeeAsks, error)
}

type Debug interface {
	Broadcast(ctx context.Context, message interface{}) (bool, error)
}

type Admin interface {
	SendOperation(ctx context.Context, operation *Operation, ignorePayment *bool) (string, error)
	ReserveOperations(ctx context.Context, num int, skip int, strategy *OperationStrategy) ([]*Operation, error)
	ReleaseOperations(ctx context.Context, operations []string, readyAtChange *ReadyAtChange) error
	DiscardOperations(ctx context.Context, operations []string) error
	BanEndorser(ctx context.Context, endorser string, duration int) error
	BannedEndorsers(ctx context.Context) ([]string, error)
}

var WebRPCServices = map[string][]string{
	"Bundler": {
		"Ping",
		"Status",
		"Peers",
		"Mempool",
		"SendOperation",
		"Operations",
		"FeeAsks",
	},
	"Debug": {
		"Broadcast",
	},
	"Admin": {
		"SendOperation",
		"ReserveOperations",
		"ReleaseOperations",
		"DiscardOperations",
		"BanEndorser",
		"BannedEndorsers",
	},
}

//
// Client
//

const BundlerPathPrefix = "/rpc/Bundler/"
const DebugPathPrefix = "/rpc/Debug/"
const AdminPathPrefix = "/rpc/Admin/"

type bundlerClient struct {
	client HTTPClient
	urls   [7]string
}

func NewBundlerClient(addr string, client HTTPClient) Bundler {
	prefix := urlBase(addr) + BundlerPathPrefix
	urls := [7]string{
		prefix + "Ping",
		prefix + "Status",
		prefix + "Peers",
		prefix + "Mempool",
		prefix + "SendOperation",
		prefix + "Operations",
		prefix + "FeeAsks",
	}
	return &bundlerClient{
		client: client,
		urls:   urls,
	}
}

func (c *bundlerClient) Ping(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[0], nil, &out)
	return out.Ret0, err
}

func (c *bundlerClient) Status(ctx context.Context) (*Status, error) {
	out := struct {
		Ret0 *Status `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[1], nil, &out)
	return out.Ret0, err
}

func (c *bundlerClient) Peers(ctx context.Context) ([]string, []string, error) {
	out := struct {
		Ret0 []string `json:"peers"`
		Ret1 []string `json:"priorityPeers"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[2], nil, &out)
	return out.Ret0, out.Ret1, err
}

func (c *bundlerClient) Mempool(ctx context.Context) (*MempoolView, error) {
	out := struct {
		Ret0 *MempoolView `json:"mempool"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[3], nil, &out)
	return out.Ret0, err
}

func (c *bundlerClient) SendOperation(ctx context.Context, operation *Operation) (string, error) {
	in := struct {
		Arg0 *Operation `json:"operation"`
	}{operation}
	out := struct {
		Ret0 string `json:"operation"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[4], in, &out)
	return out.Ret0, err
}

func (c *bundlerClient) Operations(ctx context.Context) (*Operations, error) {
	out := struct {
		Ret0 *Operations `json:"operations"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[5], nil, &out)
	return out.Ret0, err
}

func (c *bundlerClient) FeeAsks(ctx context.Context) (*FeeAsks, error) {
	out := struct {
		Ret0 *FeeAsks `json:"feeAsks"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[6], nil, &out)
	return out.Ret0, err
}

type debugClient struct {
	client HTTPClient
	urls   [1]string
}

func NewDebugClient(addr string, client HTTPClient) Debug {
	prefix := urlBase(addr) + DebugPathPrefix
	urls := [1]string{
		prefix + "Broadcast",
	}
	return &debugClient{
		client: client,
		urls:   urls,
	}
}

func (c *debugClient) Broadcast(ctx context.Context, message interface{}) (bool, error) {
	in := struct {
		Arg0 interface{} `json:"message"`
	}{message}
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[0], in, &out)
	return out.Ret0, err
}

type adminClient struct {
	client HTTPClient
	urls   [6]string
}

func NewAdminClient(addr string, client HTTPClient) Admin {
	prefix := urlBase(addr) + AdminPathPrefix
	urls := [6]string{
		prefix + "SendOperation",
		prefix + "ReserveOperations",
		prefix + "ReleaseOperations",
		prefix + "DiscardOperations",
		prefix + "BanEndorser",
		prefix + "BannedEndorsers",
	}
	return &adminClient{
		client: client,
		urls:   urls,
	}
}

func (c *adminClient) SendOperation(ctx context.Context, operation *Operation, ignorePayment *bool) (string, error) {
	in := struct {
		Arg0 *Operation `json:"operation"`
		Arg1 *bool      `json:"ignorePayment"`
	}{operation, ignorePayment}
	out := struct {
		Ret0 string `json:"operation"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[0], in, &out)
	return out.Ret0, err
}

func (c *adminClient) ReserveOperations(ctx context.Context, num int, skip int, strategy *OperationStrategy) ([]*Operation, error) {
	in := struct {
		Arg0 int                `json:"num"`
		Arg1 int                `json:"skip"`
		Arg2 *OperationStrategy `json:"strategy"`
	}{num, skip, strategy}
	out := struct {
		Ret0 []*Operation `json:"operations"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[1], in, &out)
	return out.Ret0, err
}

func (c *adminClient) ReleaseOperations(ctx context.Context, operations []string, readyAtChange *ReadyAtChange) error {
	in := struct {
		Arg0 []string       `json:"operations"`
		Arg1 *ReadyAtChange `json:"readyAtChange"`
	}{operations, readyAtChange}
	err := doJSONRequest(ctx, c.client, c.urls[2], in, nil)
	return err
}

func (c *adminClient) DiscardOperations(ctx context.Context, operations []string) error {
	in := struct {
		Arg0 []string `json:"operations"`
	}{operations}
	err := doJSONRequest(ctx, c.client, c.urls[3], in, nil)
	return err
}

func (c *adminClient) BanEndorser(ctx context.Context, endorser string, duration int) error {
	in := struct {
		Arg0 string `json:"endorser"`
		Arg1 int    `json:"duration"`
	}{endorser, duration}
	err := doJSONRequest(ctx, c.client, c.urls[4], in, nil)
	return err
}

func (c *adminClient) BannedEndorsers(ctx context.Context) ([]string, error) {
	out := struct {
		Ret0 []string `json:"endorser"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[5], nil, &out)
	return out.Ret0, err
}

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	if headers, ok := HTTPRequestHeaders(ctx); ok {
		for k := range headers {
			for _, v := range headers[k] {
				req.Header.Add(k, v)
			}
		}
	}
	return req, nil
}

// doJSONRequest is common code to make a request to the remote service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) error {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to marshal JSON body: %w", err))
	}
	if err = ctx.Err(); err != nil {
		return ErrWebrpcRequestFailed.WithCause(fmt.Errorf("aborted because context was done: %w", err))
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return ErrWebrpcRequestFailed.WithCause(fmt.Errorf("could not build request: %w", err))
	}
	resp, err := client.Do(req)
	if err != nil {
		return ErrWebrpcRequestFailed.WithCause(err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}()

	if err = ctx.Err(); err != nil {
		return ErrWebrpcRequestFailed.WithCause(fmt.Errorf("aborted because context was done: %w", err))
	}

	if resp.StatusCode != 200 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to read server error response body: %w", err))
		}

		var rpcErr WebRPCError
		if err := json.Unmarshal(respBody, &rpcErr); err != nil {
			return ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to unmarshal server error: %w", err))
		}
		if rpcErr.Cause != "" {
			rpcErr.cause = errors.New(rpcErr.Cause)
		}
		return rpcErr
	}

	if out != nil {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to read response body: %w", err))
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to unmarshal JSON response body: %w", err))
		}
	}

	return nil
}

func WithHTTPRequestHeaders(ctx context.Context, h http.Header) (context.Context, error) {
	if _, ok := h["Accept"]; ok {
		return nil, errors.New("provided header cannot set Accept")
	}
	if _, ok := h["Content-Type"]; ok {
		return nil, errors.New("provided header cannot set Content-Type")
	}

	copied := make(http.Header, len(h))
	for k, vv := range h {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}

	return context.WithValue(ctx, HTTPClientRequestHeadersCtxKey, copied), nil
}

func HTTPRequestHeaders(ctx context.Context) (http.Header, bool) {
	h, ok := ctx.Value(HTTPClientRequestHeadersCtxKey).(http.Header)
	return h, ok
}

//
// Helpers
//

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}
	HTTPRequestCtxKey              = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

func ServiceNameFromContext(ctx context.Context) string {
	service, _ := ctx.Value(ServiceNameCtxKey).(string)
	return service
}

func MethodNameFromContext(ctx context.Context) string {
	method, _ := ctx.Value(MethodNameCtxKey).(string)
	return method
}

func RequestFromContext(ctx context.Context) *http.Request {
	r, _ := ctx.Value(HTTPRequestCtxKey).(*http.Request)
	return r
}

//
// Errors
//

type WebRPCError struct {
	Name       string `json:"error"`
	Code       int    `json:"code"`
	Message    string `json:"msg"`
	Cause      string `json:"cause,omitempty"`
	HTTPStatus int    `json:"status"`
	cause      error
}

var _ error = WebRPCError{}

func (e WebRPCError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s %d: %s: %v", e.Name, e.Code, e.Message, e.cause)
	}
	return fmt.Sprintf("%s %d: %s", e.Name, e.Code, e.Message)
}

func (e WebRPCError) Is(target error) bool {
	if target == nil {
		return false
	}
	if rpcErr, ok := target.(WebRPCError); ok {
		return rpcErr.Code == e.Code
	}
	return errors.Is(e.cause, target)
}

func (e WebRPCError) Unwrap() error {
	return e.cause
}

func (e WebRPCError) WithCause(cause error) WebRPCError {
	err := e
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Deprecated: Use .WithCause() method on WebRPCError.
func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	return rpcErr.WithCause(cause)
}

// Webrpc errors
var (
	ErrWebrpcEndpoint           = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed      = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 400}
	ErrWebrpcBadRoute           = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod          = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest         = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse        = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic        = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
	ErrWebrpcInternalError      = WebRPCError{Code: -7, Name: "WebrpcInternalError", Message: "internal error", HTTPStatus: 500}
	ErrWebrpcClientDisconnected = WebRPCError{Code: -8, Name: "WebrpcClientDisconnected", Message: "client disconnected", HTTPStatus: 400}
	ErrWebrpcStreamLost         = WebRPCError{Code: -9, Name: "WebrpcStreamLost", Message: "stream lost", HTTPStatus: 400}
	ErrWebrpcStreamFinished     = WebRPCError{Code: -10, Name: "WebrpcStreamFinished", Message: "stream finished", HTTPStatus: 200}
)

// Schema errors
var (
	ErrNotFound         = WebRPCError{Code: 1000, Name: "NotFound", Message: "Not found", HTTPStatus: 404}
	ErrUnauthorized     = WebRPCError{Code: 2000, Name: "Unauthorized", Message: "Unauthorized access", HTTPStatus: 401}
	ErrPermissionDenied = WebRPCError{Code: 3000, Name: "PermissionDenied", Message: "Permission denied", HTTPStatus: 403}
)

// bundler v0.1.0 1c22ef30dd0e7a04ab1b19d375c8bd70d5b560d7
// --
// Code generated by webrpc-gen@v0.14.0 with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=rpc.ridl -target=golang -pkg=proto -server -client -out=./rpc.gen.go
package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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
	return "1c22ef30dd0e7a04ab1b19d375c8bd70d5b560d7"
}

//
// Types
//

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

// TODO : we may move the core types to protobuf , and use this just for RPC
// calls. Or we can keep this , rename to bundler.ridl and this for all types..
type Version struct {
	WebrpcVersion string `json:"webrpcVersion"`
	SchemaVersion string `json:"schemaVersion"`
	SchemaHash    string `json:"schemaHash"`
	NodeVersion   string `json:"nodeVersion"`
}

type Status struct {
	// overall status , true/false
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
	// contract address that must be called with callData to execute the operation.
	Entrypoint prototyp.Hash `json:"entrypoint"`
	// data ( in hex ) that must be passed to the entrypoint call to execute the operation.
	Data prototyp.Hash `json:"data"`
	// additional data that must be passed to the endorser when calling isOperationReady ( ) .
	EndorserCallData prototyp.Hash `json:"endorserCallData"`
	// fixed gas to be paid regardless used gas GasLimit
	FixedGas prototyp.BigInt `json:"fixedGas"`
	// minimum gasLimit that must be passed when executing the operation.
	GasLimit prototyp.BigInt `json:"gasLimit"`
	// address of the endorser contract that should be used to validate the operation.
	Endorser prototyp.Hash `json:"endorser"`
	// amount of gas that should be passed to the endorser when validating the operation.
	EndorserGasLimit prototyp.BigInt `json:"endorserGasLimit"`
	// uint256 max amount of basefee that the operation execution is expected to pay. ( Similar to EIP-1559 max_fee_per_gas ) .
	// TODO : can we use BigInt in JS later.. ? check webrpc..
	MaxFeePerGas prototyp.BigInt `json:"maxFeePerGas"`
	// uint256 fixed amount of fees that the operation execution is expected to pay to the bundler. ( Similar to EIP-1559 max_priority_fee_per_gas ) .
	// TODO : can we use BigInt in JS later.. ? check webrpc..
	MaxPriorityFeePerGas prototyp.BigInt `json:"maxPriorityFeePerGas"`
	// contract address of the ERC-20 token used to repay the bundler. ( address ( 0 ) for the native token ) .
	FeeToken prototyp.Hash `json:"feeToken"`
	// Scaling factor to convert block.basefee into the feeToken unit.
	FeeScalingFactor prototyp.BigInt `json:"feeScalingFactor"`
	// Normalization factor to convert block.basefee into the feeToken unit.
	FeeNormalizationFactor prototyp.BigInt `json:"feeNormalizationFactor"`
	// If true , the operation may have untrusted code paths. These should be treated differently by the bundler ( see untrusted environment ) .
	HasUntrustedContext bool `json:"hasUntrustedContext"`
	// Chain ID of the network where the operation is intended to be executed.
	ChainID prototyp.BigInt `json:"chainId"`
	Hash    *string         `json:"hash,omitempty"`
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
// Server types
//

type Bundler interface {
	Ping(ctx context.Context) (bool, error)
	Status(ctx context.Context) (*Status, error)
	Peers(ctx context.Context) ([]string, []string, error)
	Mempool(ctx context.Context) (*MempoolView, error)
	SendOperation(ctx context.Context, operation *Operation) (string, error)
	Operations(ctx context.Context) (*Operations, error)
	FeeAsks(ctx context.Context) (*FeeAsks, error)
}

type Admin interface {
	SendOperation(ctx context.Context, operation *Operation, ignorePayment *bool) (string, error)
	ReserveOperations(ctx context.Context, num int, skip int, strategy *OperationStrategy) ([]*Operation, error)
	ReleaseOperations(ctx context.Context, operations []string, readyAtChange *ReadyAtChange) error
	DiscardOperations(ctx context.Context, operations []string) error
	BanEndorser(ctx context.Context, endorser string, duration int) error
	BannedEndorsers(ctx context.Context) ([]string, error)
}

//
// Client types
//

type BundlerClient interface {
	Ping(ctx context.Context) (bool, error)
	Status(ctx context.Context) (*Status, error)
	Peers(ctx context.Context) ([]string, []string, error)
	Mempool(ctx context.Context) (*MempoolView, error)
	SendOperation(ctx context.Context, operation *Operation) (string, error)
	Operations(ctx context.Context) (*Operations, error)
	FeeAsks(ctx context.Context) (*FeeAsks, error)
}

type AdminClient interface {
	SendOperation(ctx context.Context, operation *Operation, ignorePayment *bool) (string, error)
	ReserveOperations(ctx context.Context, num int, skip int, strategy *OperationStrategy) ([]*Operation, error)
	ReleaseOperations(ctx context.Context, operations []string, readyAtChange *ReadyAtChange) error
	DiscardOperations(ctx context.Context, operations []string) error
	BanEndorser(ctx context.Context, endorser string, duration int) error
	BannedEndorsers(ctx context.Context) ([]string, error)
}

//
// Server
//

type WebRPCServer interface {
	http.Handler
}

type bundlerServer struct {
	Bundler
	OnError func(r *http.Request, rpcErr *WebRPCError)
}

func NewBundlerServer(svc Bundler) *bundlerServer {
	return &bundlerServer{
		Bundler: svc,
	}
}

func (s *bundlerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// In case of a panic, serve a HTTP 500 error and then panic.
		if rr := recover(); rr != nil {
			s.sendErrorJSON(w, r, ErrWebrpcServerPanic.WithCause(fmt.Errorf("%v", rr)))
			panic(rr)
		}
	}()

	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "Bundler")

	var handler func(ctx context.Context, w http.ResponseWriter, r *http.Request)
	switch r.URL.Path {
	case "/rpc/Bundler/Ping":
		handler = s.servePingJSON
	case "/rpc/Bundler/Status":
		handler = s.serveStatusJSON
	case "/rpc/Bundler/Peers":
		handler = s.servePeersJSON
	case "/rpc/Bundler/Mempool":
		handler = s.serveMempoolJSON
	case "/rpc/Bundler/SendOperation":
		handler = s.serveSendOperationJSON
	case "/rpc/Bundler/Operations":
		handler = s.serveOperationsJSON
	case "/rpc/Bundler/FeeAsks":
		handler = s.serveFeeAsksJSON
	default:
		err := ErrWebrpcBadRoute.WithCause(fmt.Errorf("no handler for path %q", r.URL.Path))
		s.sendErrorJSON(w, r, err)
		return
	}

	if r.Method != "POST" {
		w.Header().Add("Allow", "POST") // RFC 9110.
		err := ErrWebrpcBadMethod.WithCause(fmt.Errorf("unsupported method %q (only POST is allowed)", r.Method))
		s.sendErrorJSON(w, r, err)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if i := strings.Index(contentType, ";"); i >= 0 {
		contentType = contentType[:i]
	}
	contentType = strings.TrimSpace(strings.ToLower(contentType))

	switch contentType {
	case "application/json":
		handler(ctx, w, r)
	default:
		err := ErrWebrpcBadRequest.WithCause(fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		s.sendErrorJSON(w, r, err)
	}
}

func (s *bundlerServer) servePingJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Ping")

	// Call service method implementation.
	ret0, err := s.Bundler.Ping(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 bool `json:"status"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *bundlerServer) serveStatusJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Status")

	// Call service method implementation.
	ret0, err := s.Bundler.Status(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 *Status `json:"status"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *bundlerServer) servePeersJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Peers")

	// Call service method implementation.
	ret0, ret1, err := s.Bundler.Peers(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 []string `json:"peers"`
		Ret1 []string `json:"priorityPeers"`
	}{ret0, ret1}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *bundlerServer) serveMempoolJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Mempool")

	// Call service method implementation.
	ret0, err := s.Bundler.Mempool(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 *MempoolView `json:"mempool"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *bundlerServer) serveSendOperationJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "SendOperation")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to read request data: %w", err)))
		return
	}
	defer r.Body.Close()

	reqPayload := struct {
		Arg0 *Operation `json:"operation"`
	}{}
	if err := json.Unmarshal(reqBody, &reqPayload); err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to unmarshal request data: %w", err)))
		return
	}

	// Call service method implementation.
	ret0, err := s.Bundler.SendOperation(ctx, reqPayload.Arg0)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 string `json:"operation"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *bundlerServer) serveOperationsJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Operations")

	// Call service method implementation.
	ret0, err := s.Bundler.Operations(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 *Operations `json:"operations"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *bundlerServer) serveFeeAsksJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "FeeAsks")

	// Call service method implementation.
	ret0, err := s.Bundler.FeeAsks(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 *FeeAsks `json:"feeAsks"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *bundlerServer) sendErrorJSON(w http.ResponseWriter, r *http.Request, rpcErr WebRPCError) {
	if s.OnError != nil {
		s.OnError(r, &rpcErr)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
}

type adminServer struct {
	Admin
	OnError func(r *http.Request, rpcErr *WebRPCError)
}

func NewAdminServer(svc Admin) *adminServer {
	return &adminServer{
		Admin: svc,
	}
}

func (s *adminServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// In case of a panic, serve a HTTP 500 error and then panic.
		if rr := recover(); rr != nil {
			s.sendErrorJSON(w, r, ErrWebrpcServerPanic.WithCause(fmt.Errorf("%v", rr)))
			panic(rr)
		}
	}()

	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "Admin")

	var handler func(ctx context.Context, w http.ResponseWriter, r *http.Request)
	switch r.URL.Path {
	case "/rpc/Admin/SendOperation":
		handler = s.serveSendOperationJSON
	case "/rpc/Admin/ReserveOperations":
		handler = s.serveReserveOperationsJSON
	case "/rpc/Admin/ReleaseOperations":
		handler = s.serveReleaseOperationsJSON
	case "/rpc/Admin/DiscardOperations":
		handler = s.serveDiscardOperationsJSON
	case "/rpc/Admin/BanEndorser":
		handler = s.serveBanEndorserJSON
	case "/rpc/Admin/BannedEndorsers":
		handler = s.serveBannedEndorsersJSON
	default:
		err := ErrWebrpcBadRoute.WithCause(fmt.Errorf("no handler for path %q", r.URL.Path))
		s.sendErrorJSON(w, r, err)
		return
	}

	if r.Method != "POST" {
		w.Header().Add("Allow", "POST") // RFC 9110.
		err := ErrWebrpcBadMethod.WithCause(fmt.Errorf("unsupported method %q (only POST is allowed)", r.Method))
		s.sendErrorJSON(w, r, err)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if i := strings.Index(contentType, ";"); i >= 0 {
		contentType = contentType[:i]
	}
	contentType = strings.TrimSpace(strings.ToLower(contentType))

	switch contentType {
	case "application/json":
		handler(ctx, w, r)
	default:
		err := ErrWebrpcBadRequest.WithCause(fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		s.sendErrorJSON(w, r, err)
	}
}

func (s *adminServer) serveSendOperationJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "SendOperation")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to read request data: %w", err)))
		return
	}
	defer r.Body.Close()

	reqPayload := struct {
		Arg0 *Operation `json:"operation"`
		Arg1 *bool      `json:"ignorePayment"`
	}{}
	if err := json.Unmarshal(reqBody, &reqPayload); err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to unmarshal request data: %w", err)))
		return
	}

	// Call service method implementation.
	ret0, err := s.Admin.SendOperation(ctx, reqPayload.Arg0, reqPayload.Arg1)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 string `json:"operation"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *adminServer) serveReserveOperationsJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "ReserveOperations")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to read request data: %w", err)))
		return
	}
	defer r.Body.Close()

	reqPayload := struct {
		Arg0 int                `json:"num"`
		Arg1 int                `json:"skip"`
		Arg2 *OperationStrategy `json:"strategy"`
	}{}
	if err := json.Unmarshal(reqBody, &reqPayload); err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to unmarshal request data: %w", err)))
		return
	}

	// Call service method implementation.
	ret0, err := s.Admin.ReserveOperations(ctx, reqPayload.Arg0, reqPayload.Arg1, reqPayload.Arg2)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 []*Operation `json:"operations"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *adminServer) serveReleaseOperationsJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "ReleaseOperations")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to read request data: %w", err)))
		return
	}
	defer r.Body.Close()

	reqPayload := struct {
		Arg0 []string       `json:"operations"`
		Arg1 *ReadyAtChange `json:"readyAtChange"`
	}{}
	if err := json.Unmarshal(reqBody, &reqPayload); err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to unmarshal request data: %w", err)))
		return
	}

	// Call service method implementation.
	err = s.Admin.ReleaseOperations(ctx, reqPayload.Arg0, reqPayload.Arg1)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *adminServer) serveDiscardOperationsJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "DiscardOperations")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to read request data: %w", err)))
		return
	}
	defer r.Body.Close()

	reqPayload := struct {
		Arg0 []string `json:"operations"`
	}{}
	if err := json.Unmarshal(reqBody, &reqPayload); err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to unmarshal request data: %w", err)))
		return
	}

	// Call service method implementation.
	err = s.Admin.DiscardOperations(ctx, reqPayload.Arg0)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *adminServer) serveBanEndorserJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "BanEndorser")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to read request data: %w", err)))
		return
	}
	defer r.Body.Close()

	reqPayload := struct {
		Arg0 string `json:"endorser"`
		Arg1 int    `json:"duration"`
	}{}
	if err := json.Unmarshal(reqBody, &reqPayload); err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadRequest.WithCause(fmt.Errorf("failed to unmarshal request data: %w", err)))
		return
	}

	// Call service method implementation.
	err = s.Admin.BanEndorser(ctx, reqPayload.Arg0, reqPayload.Arg1)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *adminServer) serveBannedEndorsersJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "BannedEndorsers")

	// Call service method implementation.
	ret0, err := s.Admin.BannedEndorsers(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 []string `json:"endorser"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *adminServer) sendErrorJSON(w http.ResponseWriter, r *http.Request, rpcErr WebRPCError) {
	if s.OnError != nil {
		s.OnError(r, &rpcErr)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
}

func RespondWithError(w http.ResponseWriter, err error) {
	rpcErr, ok := err.(WebRPCError)
	if !ok {
		rpcErr = ErrWebrpcEndpoint.WithCause(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
}

//
// Client
//

const BundlerPathPrefix = "/rpc/Bundler/"
const AdminPathPrefix = "/rpc/Admin/"

type bundlerClient struct {
	client HTTPClient
	urls   [7]string
}

func NewBundlerClient(addr string, client HTTPClient) BundlerClient {
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

	resp, err := doHTTPRequest(ctx, c.client, c.urls[0], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *bundlerClient) Status(ctx context.Context) (*Status, error) {
	out := struct {
		Ret0 *Status `json:"status"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[1], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *bundlerClient) Peers(ctx context.Context) ([]string, []string, error) {
	out := struct {
		Ret0 []string `json:"peers"`
		Ret1 []string `json:"priorityPeers"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[2], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, out.Ret1, err
}

func (c *bundlerClient) Mempool(ctx context.Context) (*MempoolView, error) {
	out := struct {
		Ret0 *MempoolView `json:"mempool"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[3], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *bundlerClient) SendOperation(ctx context.Context, operation *Operation) (string, error) {
	in := struct {
		Arg0 *Operation `json:"operation"`
	}{operation}
	out := struct {
		Ret0 string `json:"operation"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[4], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *bundlerClient) Operations(ctx context.Context) (*Operations, error) {
	out := struct {
		Ret0 *Operations `json:"operations"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[5], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *bundlerClient) FeeAsks(ctx context.Context) (*FeeAsks, error) {
	out := struct {
		Ret0 *FeeAsks `json:"feeAsks"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[6], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

type adminClient struct {
	client HTTPClient
	urls   [6]string
}

func NewAdminClient(addr string, client HTTPClient) AdminClient {
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

	resp, err := doHTTPRequest(ctx, c.client, c.urls[0], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

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

	resp, err := doHTTPRequest(ctx, c.client, c.urls[1], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *adminClient) ReleaseOperations(ctx context.Context, operations []string, readyAtChange *ReadyAtChange) error {
	in := struct {
		Arg0 []string       `json:"operations"`
		Arg1 *ReadyAtChange `json:"readyAtChange"`
	}{operations, readyAtChange}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[2], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return err
}

func (c *adminClient) DiscardOperations(ctx context.Context, operations []string) error {
	in := struct {
		Arg0 []string `json:"operations"`
	}{operations}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[3], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return err
}

func (c *adminClient) BanEndorser(ctx context.Context, endorser string, duration int) error {
	in := struct {
		Arg0 string `json:"endorser"`
		Arg1 int    `json:"duration"`
	}{endorser, duration}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[4], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return err
}

func (c *adminClient) BannedEndorsers(ctx context.Context) ([]string, error) {
	out := struct {
		Ret0 []string `json:"endorser"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[5], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

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
	req, err := http.NewRequestWithContext(ctx, "POST", url, reqBody)
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

// doHTTPRequest is common code to make a request to the remote service.
func doHTTPRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) (*http.Response, error) {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to marshal JSON body: %w", err))
	}
	if err = ctx.Err(); err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("aborted because context was done: %w", err))
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("could not build request: %w", err))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(err)
	}

	if resp.StatusCode != 200 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to read server error response body: %w", err))
		}

		var rpcErr WebRPCError
		if err := json.Unmarshal(respBody, &rpcErr); err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to unmarshal server error: %w", err))
		}
		if rpcErr.Cause != "" {
			rpcErr.cause = errors.New(rpcErr.Cause)
		}
		return nil, rpcErr
	}

	if out != nil {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to read response body: %w", err))
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to unmarshal JSON response body: %w", err))
		}
	}

	return resp, nil
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
	HTTPResponseWriterCtxKey       = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

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
func ResponseWriterFromContext(ctx context.Context) http.ResponseWriter {
	w, _ := ctx.Value(HTTPResponseWriterCtxKey).(http.ResponseWriter)
	return w
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

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pet/v2/pet.proto

package petv2connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v2 "github.com/kiyani-org/proto-solution/gen/go/pet/v2"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PetStoreServiceName is the fully-qualified name of the PetStoreService service.
	PetStoreServiceName = "pet.v2.PetStoreService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PetStoreServiceGetPetProcedure is the fully-qualified name of the PetStoreService's GetPet RPC.
	PetStoreServiceGetPetProcedure = "/pet.v2.PetStoreService/GetPet"
	// PetStoreServicePutPetProcedure is the fully-qualified name of the PetStoreService's PutPet RPC.
	PetStoreServicePutPetProcedure = "/pet.v2.PetStoreService/PutPet"
	// PetStoreServiceDeletePetProcedure is the fully-qualified name of the PetStoreService's DeletePet
	// RPC.
	PetStoreServiceDeletePetProcedure = "/pet.v2.PetStoreService/DeletePet"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	petStoreServiceServiceDescriptor         = v2.File_pet_v2_pet_proto.Services().ByName("PetStoreService")
	petStoreServiceGetPetMethodDescriptor    = petStoreServiceServiceDescriptor.Methods().ByName("GetPet")
	petStoreServicePutPetMethodDescriptor    = petStoreServiceServiceDescriptor.Methods().ByName("PutPet")
	petStoreServiceDeletePetMethodDescriptor = petStoreServiceServiceDescriptor.Methods().ByName("DeletePet")
)

// PetStoreServiceClient is a client for the pet.v2.PetStoreService service.
type PetStoreServiceClient interface {
	GetPet(context.Context, *connect.Request[v2.GetPetRequest]) (*connect.Response[v2.GetPetResponse], error)
	PutPet(context.Context, *connect.Request[v2.PutPetRequest]) (*connect.Response[v2.PutPetResponse], error)
	DeletePet(context.Context, *connect.Request[v2.DeletePetRequest]) (*connect.Response[v2.DeletePetResponse], error)
}

// NewPetStoreServiceClient constructs a client for the pet.v2.PetStoreService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPetStoreServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PetStoreServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &petStoreServiceClient{
		getPet: connect.NewClient[v2.GetPetRequest, v2.GetPetResponse](
			httpClient,
			baseURL+PetStoreServiceGetPetProcedure,
			connect.WithSchema(petStoreServiceGetPetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		putPet: connect.NewClient[v2.PutPetRequest, v2.PutPetResponse](
			httpClient,
			baseURL+PetStoreServicePutPetProcedure,
			connect.WithSchema(petStoreServicePutPetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deletePet: connect.NewClient[v2.DeletePetRequest, v2.DeletePetResponse](
			httpClient,
			baseURL+PetStoreServiceDeletePetProcedure,
			connect.WithSchema(petStoreServiceDeletePetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// petStoreServiceClient implements PetStoreServiceClient.
type petStoreServiceClient struct {
	getPet    *connect.Client[v2.GetPetRequest, v2.GetPetResponse]
	putPet    *connect.Client[v2.PutPetRequest, v2.PutPetResponse]
	deletePet *connect.Client[v2.DeletePetRequest, v2.DeletePetResponse]
}

// GetPet calls pet.v2.PetStoreService.GetPet.
func (c *petStoreServiceClient) GetPet(ctx context.Context, req *connect.Request[v2.GetPetRequest]) (*connect.Response[v2.GetPetResponse], error) {
	return c.getPet.CallUnary(ctx, req)
}

// PutPet calls pet.v2.PetStoreService.PutPet.
func (c *petStoreServiceClient) PutPet(ctx context.Context, req *connect.Request[v2.PutPetRequest]) (*connect.Response[v2.PutPetResponse], error) {
	return c.putPet.CallUnary(ctx, req)
}

// DeletePet calls pet.v2.PetStoreService.DeletePet.
func (c *petStoreServiceClient) DeletePet(ctx context.Context, req *connect.Request[v2.DeletePetRequest]) (*connect.Response[v2.DeletePetResponse], error) {
	return c.deletePet.CallUnary(ctx, req)
}

// PetStoreServiceHandler is an implementation of the pet.v2.PetStoreService service.
type PetStoreServiceHandler interface {
	GetPet(context.Context, *connect.Request[v2.GetPetRequest]) (*connect.Response[v2.GetPetResponse], error)
	PutPet(context.Context, *connect.Request[v2.PutPetRequest]) (*connect.Response[v2.PutPetResponse], error)
	DeletePet(context.Context, *connect.Request[v2.DeletePetRequest]) (*connect.Response[v2.DeletePetResponse], error)
}

// NewPetStoreServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPetStoreServiceHandler(svc PetStoreServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	petStoreServiceGetPetHandler := connect.NewUnaryHandler(
		PetStoreServiceGetPetProcedure,
		svc.GetPet,
		connect.WithSchema(petStoreServiceGetPetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	petStoreServicePutPetHandler := connect.NewUnaryHandler(
		PetStoreServicePutPetProcedure,
		svc.PutPet,
		connect.WithSchema(petStoreServicePutPetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	petStoreServiceDeletePetHandler := connect.NewUnaryHandler(
		PetStoreServiceDeletePetProcedure,
		svc.DeletePet,
		connect.WithSchema(petStoreServiceDeletePetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/pet.v2.PetStoreService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PetStoreServiceGetPetProcedure:
			petStoreServiceGetPetHandler.ServeHTTP(w, r)
		case PetStoreServicePutPetProcedure:
			petStoreServicePutPetHandler.ServeHTTP(w, r)
		case PetStoreServiceDeletePetProcedure:
			petStoreServiceDeletePetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPetStoreServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPetStoreServiceHandler struct{}

func (UnimplementedPetStoreServiceHandler) GetPet(context.Context, *connect.Request[v2.GetPetRequest]) (*connect.Response[v2.GetPetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pet.v2.PetStoreService.GetPet is not implemented"))
}

func (UnimplementedPetStoreServiceHandler) PutPet(context.Context, *connect.Request[v2.PutPetRequest]) (*connect.Response[v2.PutPetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pet.v2.PetStoreService.PutPet is not implemented"))
}

func (UnimplementedPetStoreServiceHandler) DeletePet(context.Context, *connect.Request[v2.DeletePetRequest]) (*connect.Response[v2.DeletePetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pet.v2.PetStoreService.DeletePet is not implemented"))
}
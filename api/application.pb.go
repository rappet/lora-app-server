// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IntegrationKind int32

const (
	IntegrationKind_HTTP IntegrationKind = 0
)

var IntegrationKind_name = map[int32]string{
	0: "HTTP",
}
var IntegrationKind_value = map[string]int32{
	"HTTP": 0,
}

func (x IntegrationKind) String() string {
	return proto.EnumName(IntegrationKind_name, int32(x))
}
func (IntegrationKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type CreateApplicationRequest struct {
	// Name of the application (must be unique).
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// ID of the organization to which the application belongs.
	OrganizationID int64 `protobuf:"varint,14,opt,name=organizationID" json:"organizationID,omitempty"`
	// ID of the service profile.
	ServiceProfileID string `protobuf:"bytes,15,opt,name=serviceProfileID" json:"serviceProfileID,omitempty"`
}

func (m *CreateApplicationRequest) Reset()                    { *m = CreateApplicationRequest{} }
func (m *CreateApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateApplicationRequest) ProtoMessage()               {}
func (*CreateApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateApplicationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateApplicationRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateApplicationRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

func (m *CreateApplicationRequest) GetServiceProfileID() string {
	if m != nil {
		return m.ServiceProfileID
	}
	return ""
}

type CreateApplicationResponse struct {
	// ID of the application that was created.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateApplicationResponse) Reset()                    { *m = CreateApplicationResponse{} }
func (m *CreateApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateApplicationResponse) ProtoMessage()               {}
func (*CreateApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CreateApplicationResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetApplicationRequest struct {
	// Name of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetApplicationRequest) Reset()                    { *m = GetApplicationRequest{} }
func (m *GetApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetApplicationRequest) ProtoMessage()               {}
func (*GetApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetApplicationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetApplicationResponse struct {
	// ID of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the application.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// ID of the organization to which the application belongs.
	OrganizationID int64 `protobuf:"varint,14,opt,name=organizationID" json:"organizationID,omitempty"`
	// ID of the service profile.
	ServiceProfileID string `protobuf:"bytes,15,opt,name=serviceProfileID" json:"serviceProfileID,omitempty"`
}

func (m *GetApplicationResponse) Reset()                    { *m = GetApplicationResponse{} }
func (m *GetApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*GetApplicationResponse) ProtoMessage()               {}
func (*GetApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetApplicationResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetApplicationResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetApplicationResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetApplicationResponse) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

func (m *GetApplicationResponse) GetServiceProfileID() string {
	if m != nil {
		return m.ServiceProfileID
	}
	return ""
}

type UpdateApplicationRequest struct {
	// ID of the application to update.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the application (must be unique).
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// ID of the service profile.
	ServiceProfileID string `protobuf:"bytes,15,opt,name=serviceProfileID" json:"serviceProfileID,omitempty"`
}

func (m *UpdateApplicationRequest) Reset()                    { *m = UpdateApplicationRequest{} }
func (m *UpdateApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateApplicationRequest) ProtoMessage()               {}
func (*UpdateApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *UpdateApplicationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateApplicationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateApplicationRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateApplicationRequest) GetServiceProfileID() string {
	if m != nil {
		return m.ServiceProfileID
	}
	return ""
}

type UpdateApplicationResponse struct {
}

func (m *UpdateApplicationResponse) Reset()                    { *m = UpdateApplicationResponse{} }
func (m *UpdateApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateApplicationResponse) ProtoMessage()               {}
func (*UpdateApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type DeleteApplicationRequest struct {
	// ID of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteApplicationRequest) Reset()                    { *m = DeleteApplicationRequest{} }
func (m *DeleteApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteApplicationRequest) ProtoMessage()               {}
func (*DeleteApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *DeleteApplicationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteApplicationResponse struct {
}

func (m *DeleteApplicationResponse) Reset()                    { *m = DeleteApplicationResponse{} }
func (m *DeleteApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteApplicationResponse) ProtoMessage()               {}
func (*DeleteApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type ListApplicationRequest struct {
	// Max number of applications to return in the result-test.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	// ID of the organization to filter on.
	OrganizationID int64 `protobuf:"varint,3,opt,name=organizationID" json:"organizationID,omitempty"`
}

func (m *ListApplicationRequest) Reset()                    { *m = ListApplicationRequest{} }
func (m *ListApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*ListApplicationRequest) ProtoMessage()               {}
func (*ListApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ListApplicationRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListApplicationRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListApplicationRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

type ApplicationListItem struct {
	// ID of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the application.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// ID of the organization to which the application belongs.
	OrganizationID int64 `protobuf:"varint,14,opt,name=organizationID" json:"organizationID,omitempty"`
	// ID of the service profile.
	ServiceProfileID string `protobuf:"bytes,15,opt,name=serviceProfileID" json:"serviceProfileID,omitempty"`
	// Service-profile name.
	ServiceProfileName string `protobuf:"bytes,16,opt,name=serviceProfileName" json:"serviceProfileName,omitempty"`
}

func (m *ApplicationListItem) Reset()                    { *m = ApplicationListItem{} }
func (m *ApplicationListItem) String() string            { return proto.CompactTextString(m) }
func (*ApplicationListItem) ProtoMessage()               {}
func (*ApplicationListItem) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ApplicationListItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ApplicationListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplicationListItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ApplicationListItem) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

func (m *ApplicationListItem) GetServiceProfileID() string {
	if m != nil {
		return m.ServiceProfileID
	}
	return ""
}

func (m *ApplicationListItem) GetServiceProfileName() string {
	if m != nil {
		return m.ServiceProfileName
	}
	return ""
}

type ListApplicationResponse struct {
	// Total number of applications available within the result-set.
	TotalCount int64 `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	// Applications within this result-set.
	Result []*ApplicationListItem `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListApplicationResponse) Reset()                    { *m = ListApplicationResponse{} }
func (m *ListApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*ListApplicationResponse) ProtoMessage()               {}
func (*ListApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *ListApplicationResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListApplicationResponse) GetResult() []*ApplicationListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

type HTTPIntegrationHeader struct {
	// Key
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Value
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *HTTPIntegrationHeader) Reset()                    { *m = HTTPIntegrationHeader{} }
func (m *HTTPIntegrationHeader) String() string            { return proto.CompactTextString(m) }
func (*HTTPIntegrationHeader) ProtoMessage()               {}
func (*HTTPIntegrationHeader) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *HTTPIntegrationHeader) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HTTPIntegrationHeader) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HTTPIntegration struct {
	// The id of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The headers to use when making HTTP callbacks.
	Headers []*HTTPIntegrationHeader `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty"`
	// The URL to call for uplink data.
	DataUpURL string `protobuf:"bytes,3,opt,name=dataUpURL" json:"dataUpURL,omitempty"`
	// The URL to call for join notifications.
	JoinNotificationURL string `protobuf:"bytes,4,opt,name=joinNotificationURL" json:"joinNotificationURL,omitempty"`
	// The URL to call for ACK notifications (for confirmed downlink data).
	AckNotificationURL string `protobuf:"bytes,5,opt,name=ackNotificationURL" json:"ackNotificationURL,omitempty"`
	// The URL to call for error notifications.
	ErrorNotificationURL string `protobuf:"bytes,6,opt,name=errorNotificationURL" json:"errorNotificationURL,omitempty"`
}

func (m *HTTPIntegration) Reset()                    { *m = HTTPIntegration{} }
func (m *HTTPIntegration) String() string            { return proto.CompactTextString(m) }
func (*HTTPIntegration) ProtoMessage()               {}
func (*HTTPIntegration) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *HTTPIntegration) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HTTPIntegration) GetHeaders() []*HTTPIntegrationHeader {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPIntegration) GetDataUpURL() string {
	if m != nil {
		return m.DataUpURL
	}
	return ""
}

func (m *HTTPIntegration) GetJoinNotificationURL() string {
	if m != nil {
		return m.JoinNotificationURL
	}
	return ""
}

func (m *HTTPIntegration) GetAckNotificationURL() string {
	if m != nil {
		return m.AckNotificationURL
	}
	return ""
}

func (m *HTTPIntegration) GetErrorNotificationURL() string {
	if m != nil {
		return m.ErrorNotificationURL
	}
	return ""
}

type GetHTTPIntegrationRequest struct {
	// The id of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetHTTPIntegrationRequest) Reset()                    { *m = GetHTTPIntegrationRequest{} }
func (m *GetHTTPIntegrationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetHTTPIntegrationRequest) ProtoMessage()               {}
func (*GetHTTPIntegrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *GetHTTPIntegrationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteIntegrationRequest struct {
	// The id of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteIntegrationRequest) Reset()                    { *m = DeleteIntegrationRequest{} }
func (m *DeleteIntegrationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteIntegrationRequest) ProtoMessage()               {}
func (*DeleteIntegrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *DeleteIntegrationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListIntegrationRequest struct {
	// The id of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ListIntegrationRequest) Reset()                    { *m = ListIntegrationRequest{} }
func (m *ListIntegrationRequest) String() string            { return proto.CompactTextString(m) }
func (*ListIntegrationRequest) ProtoMessage()               {}
func (*ListIntegrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *ListIntegrationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListIntegrationResponse struct {
	// The integration kinds associated with the application.
	Kinds []IntegrationKind `protobuf:"varint,1,rep,packed,name=kinds,enum=api.IntegrationKind" json:"kinds,omitempty"`
}

func (m *ListIntegrationResponse) Reset()                    { *m = ListIntegrationResponse{} }
func (m *ListIntegrationResponse) String() string            { return proto.CompactTextString(m) }
func (*ListIntegrationResponse) ProtoMessage()               {}
func (*ListIntegrationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

func (m *ListIntegrationResponse) GetKinds() []IntegrationKind {
	if m != nil {
		return m.Kinds
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateApplicationRequest)(nil), "api.CreateApplicationRequest")
	proto.RegisterType((*CreateApplicationResponse)(nil), "api.CreateApplicationResponse")
	proto.RegisterType((*GetApplicationRequest)(nil), "api.GetApplicationRequest")
	proto.RegisterType((*GetApplicationResponse)(nil), "api.GetApplicationResponse")
	proto.RegisterType((*UpdateApplicationRequest)(nil), "api.UpdateApplicationRequest")
	proto.RegisterType((*UpdateApplicationResponse)(nil), "api.UpdateApplicationResponse")
	proto.RegisterType((*DeleteApplicationRequest)(nil), "api.DeleteApplicationRequest")
	proto.RegisterType((*DeleteApplicationResponse)(nil), "api.DeleteApplicationResponse")
	proto.RegisterType((*ListApplicationRequest)(nil), "api.ListApplicationRequest")
	proto.RegisterType((*ApplicationListItem)(nil), "api.ApplicationListItem")
	proto.RegisterType((*ListApplicationResponse)(nil), "api.ListApplicationResponse")
	proto.RegisterType((*EmptyResponse)(nil), "api.EmptyResponse")
	proto.RegisterType((*HTTPIntegrationHeader)(nil), "api.HTTPIntegrationHeader")
	proto.RegisterType((*HTTPIntegration)(nil), "api.HTTPIntegration")
	proto.RegisterType((*GetHTTPIntegrationRequest)(nil), "api.GetHTTPIntegrationRequest")
	proto.RegisterType((*DeleteIntegrationRequest)(nil), "api.DeleteIntegrationRequest")
	proto.RegisterType((*ListIntegrationRequest)(nil), "api.ListIntegrationRequest")
	proto.RegisterType((*ListIntegrationResponse)(nil), "api.ListIntegrationResponse")
	proto.RegisterEnum("api.IntegrationKind", IntegrationKind_name, IntegrationKind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Application service

type ApplicationClient interface {
	// Create creates the given application.
	Create(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*CreateApplicationResponse, error)
	// Get returns the requested application.
	Get(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error)
	// Update updates the given application.
	Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*UpdateApplicationResponse, error)
	// Delete deletes the given application.
	Delete(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*DeleteApplicationResponse, error)
	// List lists the available applications.
	List(ctx context.Context, in *ListApplicationRequest, opts ...grpc.CallOption) (*ListApplicationResponse, error)
	// CreateHTTPIntegration creates an HTTP application-integration.
	CreateHTTPIntegration(ctx context.Context, in *HTTPIntegration, opts ...grpc.CallOption) (*EmptyResponse, error)
	// GetHTTPIntegration returns the HTTP application-itegration.
	GetHTTPIntegration(ctx context.Context, in *GetHTTPIntegrationRequest, opts ...grpc.CallOption) (*HTTPIntegration, error)
	// UpdateHTTPIntegration updates the HTTP application-integration.
	UpdateHTTPIntegration(ctx context.Context, in *HTTPIntegration, opts ...grpc.CallOption) (*EmptyResponse, error)
	// DeleteIntegration deletes the application-integration of the given type.
	DeleteHTTPIntegration(ctx context.Context, in *DeleteIntegrationRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// ListIntegrations lists all configured integrations.
	ListIntegrations(ctx context.Context, in *ListIntegrationRequest, opts ...grpc.CallOption) (*ListIntegrationResponse, error)
}

type applicationClient struct {
	cc *grpc.ClientConn
}

func NewApplicationClient(cc *grpc.ClientConn) ApplicationClient {
	return &applicationClient{cc}
}

func (c *applicationClient) Create(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*CreateApplicationResponse, error) {
	out := new(CreateApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Get(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error) {
	out := new(GetApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*UpdateApplicationResponse, error) {
	out := new(UpdateApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Delete(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*DeleteApplicationResponse, error) {
	out := new(DeleteApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) List(ctx context.Context, in *ListApplicationRequest, opts ...grpc.CallOption) (*ListApplicationResponse, error) {
	out := new(ListApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) CreateHTTPIntegration(ctx context.Context, in *HTTPIntegration, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/api.Application/CreateHTTPIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) GetHTTPIntegration(ctx context.Context, in *GetHTTPIntegrationRequest, opts ...grpc.CallOption) (*HTTPIntegration, error) {
	out := new(HTTPIntegration)
	err := grpc.Invoke(ctx, "/api.Application/GetHTTPIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) UpdateHTTPIntegration(ctx context.Context, in *HTTPIntegration, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/api.Application/UpdateHTTPIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) DeleteHTTPIntegration(ctx context.Context, in *DeleteIntegrationRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/api.Application/DeleteHTTPIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) ListIntegrations(ctx context.Context, in *ListIntegrationRequest, opts ...grpc.CallOption) (*ListIntegrationResponse, error) {
	out := new(ListIntegrationResponse)
	err := grpc.Invoke(ctx, "/api.Application/ListIntegrations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Application service

type ApplicationServer interface {
	// Create creates the given application.
	Create(context.Context, *CreateApplicationRequest) (*CreateApplicationResponse, error)
	// Get returns the requested application.
	Get(context.Context, *GetApplicationRequest) (*GetApplicationResponse, error)
	// Update updates the given application.
	Update(context.Context, *UpdateApplicationRequest) (*UpdateApplicationResponse, error)
	// Delete deletes the given application.
	Delete(context.Context, *DeleteApplicationRequest) (*DeleteApplicationResponse, error)
	// List lists the available applications.
	List(context.Context, *ListApplicationRequest) (*ListApplicationResponse, error)
	// CreateHTTPIntegration creates an HTTP application-integration.
	CreateHTTPIntegration(context.Context, *HTTPIntegration) (*EmptyResponse, error)
	// GetHTTPIntegration returns the HTTP application-itegration.
	GetHTTPIntegration(context.Context, *GetHTTPIntegrationRequest) (*HTTPIntegration, error)
	// UpdateHTTPIntegration updates the HTTP application-integration.
	UpdateHTTPIntegration(context.Context, *HTTPIntegration) (*EmptyResponse, error)
	// DeleteIntegration deletes the application-integration of the given type.
	DeleteHTTPIntegration(context.Context, *DeleteIntegrationRequest) (*EmptyResponse, error)
	// ListIntegrations lists all configured integrations.
	ListIntegrations(context.Context, *ListIntegrationRequest) (*ListIntegrationResponse, error)
}

func RegisterApplicationServer(s *grpc.Server, srv ApplicationServer) {
	s.RegisterService(&_Application_serviceDesc, srv)
}

func _Application_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Create(ctx, req.(*CreateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Get(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Update(ctx, req.(*UpdateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Delete(ctx, req.(*DeleteApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).List(ctx, req.(*ListApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_CreateHTTPIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPIntegration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).CreateHTTPIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/CreateHTTPIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).CreateHTTPIntegration(ctx, req.(*HTTPIntegration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_GetHTTPIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHTTPIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).GetHTTPIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/GetHTTPIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).GetHTTPIntegration(ctx, req.(*GetHTTPIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_UpdateHTTPIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPIntegration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).UpdateHTTPIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/UpdateHTTPIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).UpdateHTTPIntegration(ctx, req.(*HTTPIntegration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_DeleteHTTPIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).DeleteHTTPIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/DeleteHTTPIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).DeleteHTTPIntegration(ctx, req.(*DeleteIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_ListIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).ListIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/ListIntegrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).ListIntegrations(ctx, req.(*ListIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Application_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Application",
	HandlerType: (*ApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Application_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Application_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Application_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Application_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Application_List_Handler,
		},
		{
			MethodName: "CreateHTTPIntegration",
			Handler:    _Application_CreateHTTPIntegration_Handler,
		},
		{
			MethodName: "GetHTTPIntegration",
			Handler:    _Application_GetHTTPIntegration_Handler,
		},
		{
			MethodName: "UpdateHTTPIntegration",
			Handler:    _Application_UpdateHTTPIntegration_Handler,
		},
		{
			MethodName: "DeleteHTTPIntegration",
			Handler:    _Application_DeleteHTTPIntegration_Handler,
		},
		{
			MethodName: "ListIntegrations",
			Handler:    _Application_ListIntegrations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application.proto",
}

func init() { proto.RegisterFile("application.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 834 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xfe, 0x1d, 0xa7, 0xf9, 0xdb, 0xa9, 0xfe, 0x24, 0xdd, 0x26, 0xf9, 0x5d, 0x27, 0x44, 0xc1,
	0x08, 0x88, 0x5c, 0x91, 0x54, 0x29, 0x27, 0x2e, 0x08, 0xb5, 0x55, 0x1b, 0x51, 0x55, 0x95, 0xd5,
	0xde, 0x10, 0x92, 0x89, 0x37, 0xe9, 0x36, 0x8e, 0x6d, 0xec, 0x4d, 0xa5, 0x16, 0xb8, 0x70, 0x85,
	0x1b, 0xaf, 0xc0, 0x13, 0xf0, 0x2a, 0x3c, 0x00, 0x17, 0xee, 0xbc, 0x02, 0xf2, 0xee, 0xa6, 0x49,
	0x9d, 0xb5, 0x94, 0x0a, 0x84, 0xb8, 0xc5, 0x33, 0xdf, 0xce, 0xcc, 0x37, 0xdf, 0xec, 0x6c, 0x60,
	0xcd, 0x0e, 0x02, 0x97, 0xf4, 0x6c, 0x4a, 0x7c, 0xaf, 0x15, 0x84, 0x3e, 0xf5, 0x91, 0x6a, 0x07,
	0x44, 0xaf, 0x0d, 0x7c, 0x7f, 0xe0, 0xe2, 0xb6, 0x1d, 0x90, 0xb6, 0xed, 0x79, 0x3e, 0x65, 0x88,
	0x88, 0x43, 0x8c, 0xcf, 0x0a, 0x68, 0x3b, 0x21, 0xb6, 0x29, 0x7e, 0x36, 0x3d, 0x6e, 0xe1, 0xd7,
	0x63, 0x1c, 0x51, 0x84, 0x20, 0xeb, 0xd9, 0x23, 0xac, 0x29, 0x0d, 0xa5, 0xb9, 0x62, 0xb1, 0xdf,
	0xa8, 0x01, 0xab, 0x0e, 0x8e, 0x7a, 0x21, 0x09, 0x62, 0xa4, 0x96, 0x61, 0xae, 0x59, 0x13, 0x7a,
	0x00, 0x79, 0x3f, 0x1c, 0xd8, 0x1e, 0xb9, 0x62, 0xc1, 0xba, 0xbb, 0x5a, 0xbe, 0xa1, 0x34, 0x55,
	0x2b, 0x61, 0x45, 0x26, 0x14, 0x23, 0x1c, 0x5e, 0x90, 0x1e, 0x3e, 0x0e, 0xfd, 0x3e, 0x71, 0x71,
	0x77, 0x57, 0x2b, 0xb0, 0x70, 0x73, 0x76, 0x63, 0x13, 0x36, 0x24, 0x55, 0x46, 0x81, 0xef, 0x45,
	0x18, 0xe5, 0x21, 0x43, 0x1c, 0x56, 0xa4, 0x6a, 0x65, 0x88, 0x63, 0x3c, 0x84, 0xf2, 0x3e, 0xa6,
	0x12, 0x3e, 0x49, 0xe0, 0x17, 0x05, 0x2a, 0x49, 0xa4, 0x3c, 0xe6, 0x75, 0x2b, 0x32, 0xe9, 0xad,
	0x50, 0xff, 0x4c, 0x2b, 0x3e, 0x2a, 0xa0, 0x9d, 0x06, 0x8e, 0x5c, 0xb1, 0xdf, 0x53, 0xf6, 0x6d,
	0xca, 0xa9, 0xc2, 0x86, 0xa4, 0x1a, 0xde, 0x45, 0xc3, 0x04, 0x6d, 0x17, 0xbb, 0x78, 0x91, 0x52,
	0xe3, 0x40, 0x12, 0xac, 0x08, 0xe4, 0x41, 0xe5, 0x90, 0x44, 0x32, 0x4d, 0x4b, 0xb0, 0xe4, 0x92,
	0x11, 0xa1, 0x22, 0x12, 0xff, 0x40, 0x15, 0xc8, 0xf9, 0xfd, 0x7e, 0x84, 0x29, 0x63, 0xae, 0x5a,
	0xe2, 0x4b, 0x22, 0x88, 0x2a, 0x13, 0xc4, 0xf8, 0xa6, 0xc0, 0xfa, 0x4c, 0xb2, 0x38, 0x77, 0x97,
	0xe2, 0xd1, 0xdf, 0x3b, 0x16, 0xa8, 0x05, 0xe8, 0xa6, 0xed, 0x28, 0xae, 0xab, 0xc8, 0xd0, 0x12,
	0x8f, 0x31, 0x84, 0xff, 0xe7, 0x3a, 0x2a, 0x66, 0xbf, 0x0e, 0x40, 0x7d, 0x6a, 0xbb, 0x3b, 0xfe,
	0xd8, 0x9b, 0xf4, 0x75, 0xc6, 0x82, 0xb6, 0x20, 0x17, 0xe2, 0x68, 0xec, 0xc6, 0xcd, 0x55, 0x9b,
	0xab, 0x1d, 0xad, 0x65, 0x07, 0xa4, 0x25, 0x69, 0x97, 0x25, 0x70, 0x46, 0x01, 0xfe, 0xdb, 0x1b,
	0x05, 0xf4, 0xf2, 0x5a, 0xcf, 0xa7, 0x50, 0x3e, 0x38, 0x39, 0x39, 0xee, 0x7a, 0x14, 0x0f, 0x42,
	0x76, 0xe6, 0x00, 0xdb, 0x0e, 0x0e, 0x51, 0x11, 0xd4, 0x21, 0xbe, 0x14, 0x1b, 0x27, 0xfe, 0x19,
	0x0b, 0x7c, 0x61, 0xbb, 0xe3, 0x49, 0x8f, 0xf9, 0x87, 0xf1, 0x21, 0x03, 0x85, 0x44, 0x84, 0x39,
	0x71, 0x1e, 0xc3, 0xbf, 0x67, 0x2c, 0x6a, 0x24, 0x0a, 0xd5, 0x59, 0xa1, 0xd2, 0xc4, 0xd6, 0x04,
	0x8a, 0x6a, 0xb0, 0xe2, 0xd8, 0xd4, 0x3e, 0x0d, 0x4e, 0xad, 0x43, 0x21, 0xde, 0xd4, 0x80, 0xb6,
	0x60, 0xfd, 0xdc, 0x27, 0xde, 0x91, 0x4f, 0x49, 0x5f, 0xb0, 0x8d, 0x71, 0x59, 0x86, 0x93, 0xb9,
	0x62, 0x61, 0xec, 0xde, 0x30, 0x79, 0x60, 0x89, 0x0b, 0x33, 0xef, 0x41, 0x1d, 0x28, 0xe1, 0x30,
	0xf4, 0xc3, 0xe4, 0x89, 0x1c, 0x3b, 0x21, 0xf5, 0xc5, 0xeb, 0x71, 0x1f, 0xd3, 0x04, 0xb1, 0xb4,
	0x8b, 0x76, 0x7d, 0x29, 0x17, 0xc0, 0x36, 0xf9, 0xbd, 0x5b, 0x00, 0xb9, 0xc7, 0xe7, 0xe9, 0x06,
	0x52, 0xcc, 0x93, 0x09, 0x4b, 0x43, 0xe2, 0x39, 0x91, 0xa6, 0x34, 0xd4, 0x66, 0xbe, 0x53, 0x62,
	0x2a, 0xcc, 0x00, 0x9f, 0x13, 0xcf, 0xb1, 0x38, 0xc4, 0xac, 0x42, 0x21, 0xe1, 0x41, 0xcb, 0x90,
	0x8d, 0x99, 0x15, 0xff, 0xe9, 0xfc, 0x58, 0x86, 0xd5, 0x99, 0x31, 0x43, 0x18, 0x72, 0xfc, 0x55,
	0x40, 0x77, 0x58, 0xcc, 0xb4, 0x87, 0x4c, 0xaf, 0xa7, 0xb9, 0xc5, 0x38, 0xd6, 0xde, 0x7f, 0xfd,
	0xfe, 0x29, 0x53, 0x31, 0xd6, 0xf8, 0x2b, 0x39, 0x45, 0x44, 0x4f, 0x14, 0x13, 0xbd, 0x04, 0x75,
	0x1f, 0x53, 0xc4, 0xa7, 0x47, 0xfa, 0xb2, 0xe8, 0x55, 0xa9, 0x4f, 0x44, 0xaf, 0xb3, 0xe8, 0x1a,
	0xaa, 0xcc, 0x45, 0x6f, 0xbf, 0x21, 0xce, 0x3b, 0x74, 0x0e, 0x39, 0xbe, 0x42, 0x05, 0x8d, 0xb4,
	0xed, 0x2e, 0x68, 0xa4, 0xaf, 0xdb, 0xbb, 0x2c, 0x51, 0x55, 0x4f, 0x49, 0x14, 0x73, 0x19, 0x40,
	0x8e, 0x8b, 0x2f, 0x72, 0xa5, 0xad, 0x67, 0x91, 0x2b, 0x7d, 0x23, 0x0b, 0x52, 0x66, 0x1a, 0xa9,
	0x17, 0x90, 0x8d, 0xe7, 0x01, 0xf1, 0xce, 0xc8, 0x97, 0xb7, 0x5e, 0x93, 0x3b, 0x45, 0x8a, 0x0d,
	0x96, 0x62, 0x1d, 0xcd, 0xab, 0x82, 0x2e, 0xa0, 0xcc, 0xd5, 0x4c, 0xee, 0x80, 0x92, 0xec, 0x8a,
	0xeb, 0x88, 0x59, 0x6f, 0xae, 0xa0, 0x6d, 0x16, 0xfd, 0x91, 0xd1, 0x94, 0x13, 0x68, 0x93, 0xe9,
	0xf9, 0xa8, 0x7d, 0x46, 0x69, 0x10, 0xb7, 0xef, 0x2d, 0xa0, 0xf9, 0x8b, 0x86, 0xea, 0x13, 0xf5,
	0xe5, 0x37, 0x50, 0x97, 0x16, 0x65, 0x6c, 0xb1, 0x02, 0x4c, 0xb4, 0x70, 0x01, 0x31, 0x6b, 0x2e,
	0xfe, 0x2f, 0xb3, 0xd6, 0x6f, 0xc9, 0xba, 0xcc, 0x07, 0x21, 0x99, 0x77, 0x76, 0x86, 0x24, 0xbc,
	0x65, 0x05, 0x08, 0xd6, 0xe6, 0xe2, 0xac, 0xaf, 0xa0, 0x98, 0xd8, 0x2c, 0xd1, 0xcc, 0x54, 0x49,
	0xd2, 0xd6, 0xe4, 0x4e, 0x51, 0xc0, 0x26, 0x2b, 0xe0, 0x3e, 0xba, 0xb7, 0x40, 0x01, 0xaf, 0x72,
	0xec, 0x5f, 0xf2, 0xf6, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x22, 0x39, 0x8b, 0x5d, 0x0b,
	0x00, 0x00,
}

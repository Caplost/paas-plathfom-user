// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/role/role.proto

package role

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Role service

func NewRoleEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Role service

type RoleService interface {
	//对外提供添加服务
	AddRole(ctx context.Context, in *RoleInfo, opts ...client.CallOption) (*Response, error)
	DeleteRole(ctx context.Context, in *RoleId, opts ...client.CallOption) (*Response, error)
	UpdateRole(ctx context.Context, in *RoleInfo, opts ...client.CallOption) (*Response, error)
	FindRoleByID(ctx context.Context, in *RoleId, opts ...client.CallOption) (*RoleInfo, error)
	FindAllRole(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllRole, error)
	AddPermission(ctx context.Context, in *RolePermission, opts ...client.CallOption) (*Response, error)
	UpdatePermission(ctx context.Context, in *RolePermission, opts ...client.CallOption) (*Response, error)
	DeletePermission(ctx context.Context, in *RolePermission, opts ...client.CallOption) (*Response, error)
}

type roleService struct {
	c    client.Client
	name string
}

func NewRoleService(name string, c client.Client) RoleService {
	return &roleService{
		c:    c,
		name: name,
	}
}

func (c *roleService) AddRole(ctx context.Context, in *RoleInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Role.AddRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) DeleteRole(ctx context.Context, in *RoleId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Role.DeleteRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) UpdateRole(ctx context.Context, in *RoleInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Role.UpdateRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) FindRoleByID(ctx context.Context, in *RoleId, opts ...client.CallOption) (*RoleInfo, error) {
	req := c.c.NewRequest(c.name, "Role.FindRoleByID", in)
	out := new(RoleInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) FindAllRole(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllRole, error) {
	req := c.c.NewRequest(c.name, "Role.FindAllRole", in)
	out := new(AllRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) AddPermission(ctx context.Context, in *RolePermission, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Role.AddPermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) UpdatePermission(ctx context.Context, in *RolePermission, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Role.UpdatePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) DeletePermission(ctx context.Context, in *RolePermission, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Role.DeletePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Role service

type RoleHandler interface {
	//对外提供添加服务
	AddRole(context.Context, *RoleInfo, *Response) error
	DeleteRole(context.Context, *RoleId, *Response) error
	UpdateRole(context.Context, *RoleInfo, *Response) error
	FindRoleByID(context.Context, *RoleId, *RoleInfo) error
	FindAllRole(context.Context, *FindAll, *AllRole) error
	AddPermission(context.Context, *RolePermission, *Response) error
	UpdatePermission(context.Context, *RolePermission, *Response) error
	DeletePermission(context.Context, *RolePermission, *Response) error
}

func RegisterRoleHandler(s server.Server, hdlr RoleHandler, opts ...server.HandlerOption) error {
	type role interface {
		AddRole(ctx context.Context, in *RoleInfo, out *Response) error
		DeleteRole(ctx context.Context, in *RoleId, out *Response) error
		UpdateRole(ctx context.Context, in *RoleInfo, out *Response) error
		FindRoleByID(ctx context.Context, in *RoleId, out *RoleInfo) error
		FindAllRole(ctx context.Context, in *FindAll, out *AllRole) error
		AddPermission(ctx context.Context, in *RolePermission, out *Response) error
		UpdatePermission(ctx context.Context, in *RolePermission, out *Response) error
		DeletePermission(ctx context.Context, in *RolePermission, out *Response) error
	}
	type Role struct {
		role
	}
	h := &roleHandler{hdlr}
	return s.Handle(s.NewHandler(&Role{h}, opts...))
}

type roleHandler struct {
	RoleHandler
}

func (h *roleHandler) AddRole(ctx context.Context, in *RoleInfo, out *Response) error {
	return h.RoleHandler.AddRole(ctx, in, out)
}

func (h *roleHandler) DeleteRole(ctx context.Context, in *RoleId, out *Response) error {
	return h.RoleHandler.DeleteRole(ctx, in, out)
}

func (h *roleHandler) UpdateRole(ctx context.Context, in *RoleInfo, out *Response) error {
	return h.RoleHandler.UpdateRole(ctx, in, out)
}

func (h *roleHandler) FindRoleByID(ctx context.Context, in *RoleId, out *RoleInfo) error {
	return h.RoleHandler.FindRoleByID(ctx, in, out)
}

func (h *roleHandler) FindAllRole(ctx context.Context, in *FindAll, out *AllRole) error {
	return h.RoleHandler.FindAllRole(ctx, in, out)
}

func (h *roleHandler) AddPermission(ctx context.Context, in *RolePermission, out *Response) error {
	return h.RoleHandler.AddPermission(ctx, in, out)
}

func (h *roleHandler) UpdatePermission(ctx context.Context, in *RolePermission, out *Response) error {
	return h.RoleHandler.UpdatePermission(ctx, in, out)
}

func (h *roleHandler) DeletePermission(ctx context.Context, in *RolePermission, out *Response) error {
	return h.RoleHandler.DeletePermission(ctx, in, out)
}

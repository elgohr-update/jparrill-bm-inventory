// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/filanov/bm-inventory/restapi/operations"
	"github.com/filanov/bm-inventory/restapi/operations/inventory"
)

type contextKey string

const AuthKey contextKey = "Auth"

//go:generate mockery -name InventoryAPI -inpkg

// InventoryAPI
type InventoryAPI interface {
	ListNodes(ctx context.Context, params inventory.ListNodesParams) middleware.Responder
	RegisterNode(ctx context.Context, params inventory.RegisterNodeParams) middleware.Responder
}

// Config is configuration for Handler
type Config struct {
	InventoryAPI
	Logger func(string, ...interface{})
	// InnerMiddleware is for the handler executors. These do not apply to the swagger.json document.
	// The middleware executes after routing but before authentication, binding and validation
	InnerMiddleware func(http.Handler) http.Handler

	// Authorizer is used to authorize a request after the Auth function was called using the "Auth*" functions
	// and the principal was stored in the context in the "AuthKey" context value.
	Authorizer func(*http.Request) error
}

// Handler returns an http.Handler given the handler configuration
// It mounts all the business logic implementers in the right routing.
func Handler(c Config) (http.Handler, error) {
	h, _, err := HandlerAPI(c)
	return h, err
}

// HandlerAPI returns an http.Handler given the handler configuration
// and the corresponding *BMInventory instance.
// It mounts all the business logic implementers in the right routing.
func HandlerAPI(c Config) (http.Handler, *operations.BMInventoryAPI, error) {
	spec, err := loads.Analyzed(swaggerCopy(SwaggerJSON), "")
	if err != nil {
		return nil, nil, fmt.Errorf("analyze swagger: %v", err)
	}
	api := operations.NewBMInventoryAPI(spec)
	api.ServeError = errors.ServeError
	api.Logger = c.Logger

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.InventoryListNodesHandler = inventory.ListNodesHandlerFunc(func(params inventory.ListNodesParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InventoryAPI.ListNodes(ctx, params)
	})
	api.InventoryRegisterNodeHandler = inventory.RegisterNodeHandlerFunc(func(params inventory.RegisterNodeParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InventoryAPI.RegisterNode(ctx, params)
	})
	api.ServerShutdown = func() {}
	return api.Serve(c.InnerMiddleware), api, nil
}

// swaggerCopy copies the swagger json to prevent data races in runtime
func swaggerCopy(orig json.RawMessage) json.RawMessage {
	c := make(json.RawMessage, len(orig))
	copy(c, orig)
	return c
}

// authorizer is a helper function to implement the runtime.Authorizer interface.
type authorizer func(*http.Request) error

func (a authorizer) Authorize(req *http.Request, principal interface{}) error {
	if a == nil {
		return nil
	}
	ctx := storeAuth(req.Context(), principal)
	return a(req.WithContext(ctx))
}

func storeAuth(ctx context.Context, principal interface{}) context.Context {
	return context.WithValue(ctx, AuthKey, principal)
}
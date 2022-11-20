package app

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/go-eagle/eagle/infrastructure/registry"

	"github.com/go-eagle/eagle/infrastructure/logger"
	"github.com/go-eagle/eagle/infrastructure/transport"
)

// Option is func for application
type Option func(o *options)

// options is an application options
type options struct {
	id        string
	name      string
	version   string
	metadata  map[string]string
	endpoints []*url.URL

	sigs []os.Signal
	ctx  context.Context

	logger logger.Logger

	registry        registry.Registry
	registryTimeout time.Duration
	servers         []transport.Server
}

// WithID with app id
func WithID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

// WithName .
func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

// WithVersion with a version
func WithVersion(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

// WithContext with a context
func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

// WithSignal with some system signal
func WithSignal(sigs ...os.Signal) Option {
	return func(o *options) {
		o.sigs = sigs
	}
}

// WithMetadata with service metadata.
func WithMetadata(md map[string]string) Option {
	return func(o *options) { o.metadata = md }
}

// WithEndpoint with service endpoint.
func WithEndpoint(endpoints ...*url.URL) Option {
	return func(o *options) { o.endpoints = endpoints }
}

// WithRegistry with service registry.
func WithRegistry(r registry.Registry) Option {
	return func(o *options) { o.registry = r }
}

// WithLogger .
func WithLogger(logger logger.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// WithServer with a server , http or grpc
func WithServer(srv ...transport.Server) Option {
	return func(o *options) {
		o.servers = srv
	}
}

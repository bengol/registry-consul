package consul

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

// consulHttpResolver is a consul resolver based on HTTP protocol.
type consulHttpResolver struct {
	consulHttpClient *http.Client
}

// NewConsulResolver creates a consul based resolver
func NewConsulResolver(consul_agent_address string) (discovery.Resolver, error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	return &consulHttpResolver{
		consulHttpClient: &http.Client{Transport: tr},
	}, nil
}

// Name implements the Resolver interface.
func (e *consulHttpResolver) Name() string {
	return "consul"
}

// Target implements the Resolver interface.
func (e *consulHttpResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	return target.ServiceName()
}

// Resolve implements the Resolver interface.
func (e *consulHttpResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	var eps []discovery.Instance
	// TODO: fill eps

	return discovery.Result{
		Cacheable: true,
		CacheKey:  "",
		Instances: eps,
	}, nil
}

// Diff implements the Resolver interface.
func (e *consulHttpResolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}

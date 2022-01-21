package consul

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/hashicorp/consul/api"
)

// consulHttpResolver is a consul resolver based on HTTP protocol.
type consulHttpResolver struct {
	consulClient *api.Client
}

// NewConsulResolver creates a consul based resolver
func NewConsulResolver(endpoint string) (discovery.Resolver, error) {
	// Make client config
	conf := api.DefaultConfig()

	conf.Address = endpoint

	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &consulHttpResolver{
		consulClient: client,
	}, nil
}

// Name implements the Resolver interface.
func (r *consulHttpResolver) Name() string {
	return "consul"
}

// Target implements the Resolver interface.
func (r *consulHttpResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	return target.ServiceName()
}

// Resolve implements the Resolver interface.
func (r *consulHttpResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	var eps []discovery.Instance

	agentServices, err := r.consulClient.Agent().Services()
	if err != nil {
		return discovery.Result{
			Cacheable: false,
			CacheKey:  "",
			Instances: eps,
		}, err
	}

	for key, service := range agentServices {
		log.Printf("%s:%v", key, service.ID)
	}

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

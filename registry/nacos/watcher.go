package nacos

import (
	"context"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"package-demo/registry"
)

var _ registry.Watcher = (*watcher)(nil)

type watcher struct {
	serviceName string
	groupName   string
	kind        string
	clusters    []string
	ctx         context.Context
	cancel      context.CancelFunc
	watchChan   chan struct{}
	client      naming_client.INamingClient
}

func newWatcher(ctx context.Context, client naming_client.INamingClient, serviceName string, opts options) (registry.Watcher, error) {
	w := &watcher{
		serviceName: serviceName,
		groupName:   opts.group,
		kind:        opts.kind,
		clusters:    []string{opts.cluster},
		client:      client,
	}
	w.ctx, w.cancel = context.WithCancel(ctx)
	err := w.client.Subscribe(&vo.SubscribeParam{
		ServiceName: w.serviceName,
		GroupName:   w.groupName,
		SubscribeCallback: func(services []model.Instance, err error) {
			w.watchChan <- struct{}{}
		},
	})
	return w, err
}

func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
	select {
	case <-w.ctx.Done():
		return nil, w.ctx.Err()
	case <-w.watchChan:
	}

	service, err := w.client.GetService(vo.GetServiceParam{
		ServiceName: w.serviceName,
		GroupName:   w.groupName,
		Clusters:    w.clusters,
	})
	if err != nil {
		return nil, err
	}
	items := make([]*registry.ServiceInstance, 0, len(service.Hosts))
	for _, in := range service.Hosts {
		kind := w.kind
		if k, ok := in.Metadata["kind"]; ok {
			kind = k
		}
		items = append(items, &registry.ServiceInstance{
			ID:        in.InstanceId,
			Name:      service.Name,
			Version:   in.Metadata["version"],
			Metadata:  in.Metadata,
			Endpoints: []string{fmt.Sprintf("%s://%s:%d", kind, in.Ip, in.Port)},
		})
	}
	return items, nil
}

// Stop close the watcher.
func (w *watcher) Stop() error {
	w.cancel()
	return w.client.Unsubscribe(&vo.SubscribeParam{
		ServiceName: w.serviceName,
		GroupName:   w.groupName,
		Clusters:    w.clusters,
	})
}

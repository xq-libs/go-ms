package registry

import (
	"errors"
	"fmt"
	"github.com/xq-libs/go-utils/fetcher"
	"math/rand"
	"sync"
)

type ServiceFetcher struct {
	GroupId   string
	ServiceId string
	Fetchers  []*InstanceFetcher
}

type InstanceFetcher struct {
	Id      string
	Status  string
	Config  ServiceInstance
	Fetcher *fetcher.Fetcher
}

var (
	lcxFetcher      = sync.Mutex{}
	serviceFetchers = make([]*ServiceFetcher, 0)
)

func (sf *ServiceFetcher) GetAnyActiveInstanceFetcher() (*InstanceFetcher, error) {
	activeFetcher := sf.GetActiveInstanceFetcher()
	if len(activeFetcher) < 1 {
		refreshServiceFetcher(sf.GroupId, sf.ServiceId)
		activeFetcher = sf.GetActiveInstanceFetcher()
	}
	activeLength := len(activeFetcher)
	if activeLength < 1 {
		return nil, errors.New("not found any active instance fetcher")
	}
	return activeFetcher[rand.Intn(activeLength)], nil
}

func (sf *ServiceFetcher) GetActiveInstanceFetcher() []*InstanceFetcher {
	result := make([]*InstanceFetcher, 0)
	for _, instanceFetcher := range sf.Fetchers {
		if instanceFetcher.Status == StatusOk {
			result = append(result, instanceFetcher)
		}
	}
	return result
}

func GetOneServiceFetcher(groupId string, serviceId string) (*fetcher.Fetcher, error) {
	instanceFetcher, err := GetOneInstanceFetcher(groupId, serviceId)
	if err != nil {
		return nil, err
	}
	return instanceFetcher.Fetcher, nil
}

func GetOneInstanceFetcher(groupId string, serviceId string) (*InstanceFetcher, error) {
	sf := findServiceFetcher(groupId, serviceId)
	if sf == nil {
		sf = loadServiceFetcher(groupId, serviceId)
	}
	if sf == nil {
		return nil, errors.New("not found any fetcher from registry")
	}
	return sf.GetAnyActiveInstanceFetcher()
}

func findServiceFetcher(groupId string, serviceId string) *ServiceFetcher {
	lcxFetcher.Lock()
	defer lcxFetcher.Unlock()

	for _, sf := range serviceFetchers {
		if sf.GroupId == groupId && sf.ServiceId == serviceId {
			return sf
		}
	}
	return nil
}

func refreshServiceFetcher(groupId string, serviceId string) {
	removeServiceFetcher(groupId, serviceId)
	loadServiceFetcher(groupId, serviceId)
}

func removeServiceFetcher(groupId string, serviceId string) {
	lcxFetcher.Lock()
	defer lcxFetcher.Unlock()

	index := getIndex(groupId, serviceId)
	if index >= 0 {
		serviceFetchers = append(serviceFetchers[:index], serviceFetchers[index+1:]...)
	}
}

func loadServiceFetcher(groupId string, serviceId string) *ServiceFetcher {
	lcxFetcher.Lock()
	defer lcxFetcher.Unlock()

	// Load all service instance
	instances := FindAllServiceInstances(groupId, serviceId)
	instanceFetchers := make([]*InstanceFetcher, 0)
	for _, instance := range instances {
		instanceFetchers = append(instanceFetchers, buildInstanceFetcher(instance))
	}
	// return a new service fetcher
	serviceFetcher := &ServiceFetcher{
		GroupId:   groupId,
		ServiceId: serviceId,
		Fetchers:  instanceFetchers,
	}
	// append to global fetchers
	serviceFetchers = append(serviceFetchers, serviceFetcher)
	// return
	return serviceFetcher
}

func buildInstanceFetcher(instance ServiceInstance) *InstanceFetcher {
	return &InstanceFetcher{
		Id:      instance.GetInstanceId(),
		Status:  instance.Status,
		Config:  instance,
		Fetcher: buildFetcher(instance),
	}
}

func buildFetcher(instance ServiceInstance) *fetcher.Fetcher {
	return fetcher.NewFetcher(fetcher.Config{
		BaseUrl: fmt.Sprintf("%s:%d", instance.Host, instance.Port),
	})
}

func getIndex(groupId string, serviceId string) int {
	for index, item := range serviceFetchers {
		if item.GroupId == groupId && item.ServiceId == serviceId {
			return index
		}
	}
	return -1
}

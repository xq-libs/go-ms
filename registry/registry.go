package registry

import (
	"fmt"
	"github.com/xq-libs/go-ms"
	"github.com/xq-libs/go-utils/fetcher"
	"log"
	"sync"
)

const (
	StatusOk   = "Ok"
	StatusDown = "Down"
)

type ServiceInstance struct {
	GroupId        string `json:"groupId"`
	ServiceId      string `json:"serviceId"`
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Status         string `json:"status"`
	HomePageUrl    string `json:"homePageUrl"`
	HealthCheckUrl string `json:"healthCheckUrl"`
}

type Response[T any] struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
	Data T      `json:"data,omitempty"`
}

var (
	lcxInstance      = sync.Mutex{}
	serviceInstances []*ServiceInstance
)

func (i *ServiceInstance) GetInstanceId() string {
	return fmt.Sprintf("%s:%s:%s:%d", i.GroupId, i.ServiceId, i.Host, i.Port)
}

func RegisterServiceInstance() bool {
	config := GetConfig()
	req := createRegisterRequest(config.Instance)
	result, _, err := fetcher.PostJson(config.Client.Url+"/register", req, &Response[any]{})
	if err != nil {
		log.Printf("Register service to registry center failure: %v", err)
		return false
	}
	if result.Code != ms.SuccessCode {
		log.Printf("Register service to registry center failure: %s", result.Msg)
		return false
	}
	return true
}

func createRegisterRequest(instance *InstanceConfig) *ServiceInstance {
	return &ServiceInstance{
		GroupId:        instance.Group,
		ServiceId:      instance.Service,
		Host:           instance.Host,
		Port:           instance.Port,
		Status:         StatusOk,
		HomePageUrl:    instance.HomePageUrl,
		HealthCheckUrl: instance.HealthCheckUrl,
	}
}

func UnregisterServiceInstance() bool {
	config := GetConfig()
	req := createUnregisterRequest(config.Instance)
	result, _, err := fetcher.PostJson(config.Client.Url+"/unregister", req, &Response[any]{})
	if err != nil {
		log.Panicf("Unregister service from registry center failure: %v", err)
		return false
	}
	if result.Code != ms.SuccessCode {
		log.Printf("Unregister service to registry center failure: %s", result.Msg)
		return false
	}
	return true
}

func createUnregisterRequest(instance *InstanceConfig) *ServiceInstance {
	return &ServiceInstance{
		GroupId:   instance.Group,
		ServiceId: instance.Service,
		Host:      instance.Host,
	}
}

func LoadAllServiceInstances() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Load all service instances from register centor failure: %v", err)
		}
	}()

	config := GetConfig()
	queryParam := fmt.Sprintf("groupId=%s", config.Instance.Group)
	result, _, err := fetcher.GetJson(config.Client.Url+"/services?"+queryParam, &Response[[]ServiceInstance]{})
	if err != nil {
		log.Panicf("Load all services from registry center failure: %v", err)
		return false
	}
	if result.Code != ms.SuccessCode {
		log.Printf("Load all services from registry center failure: %s", result.Msg)
		return false
	}
	// add lock
	lcxInstance.Lock()
	serviceInstances = convertToServiceInstances(result.Data)
	lcxInstance.Unlock()
	// return
	log.Printf("Registered services are: %v", result.Data)
	return true
}

func convertToServiceInstances(instances []ServiceInstance) []*ServiceInstance {
	result := make([]*ServiceInstance, 0)
	for _, instance := range instances {
		result = append(result, &instance)
	}
	return result
}

func FindAllServiceInstances(groupId string, serviceId string) []*ServiceInstance {
	lcxInstance.Lock()
	defer lcxInstance.Unlock()

	result := make([]*ServiceInstance, 0)
	for _, instance := range serviceInstances {
		if instance.GroupId == groupId && instance.ServiceId == serviceId {
			result = append(result, instance)
		}
	}
	return result
}

func FindAllActiveServiceInstances(groupId string, serviceId string) []*ServiceInstance {
	lcxInstance.Lock()
	defer lcxInstance.Unlock()

	result := make([]*ServiceInstance, 0)
	for _, instance := range serviceInstances {
		if instance.GroupId == groupId && instance.ServiceId == serviceId && instance.Status == StatusOk {
			result = append(result, instance)
		}
	}
	return result
}

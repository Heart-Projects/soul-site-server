package factory

import (
	"fmt"
	"reflect"
	"sync"
)

type Store interface {
	// Store 存储一个实例
	Store(instanceName string, ins interface{}) (bool, error)
	// StoreWithDefaultName 根据类型存储一个实例
	StoreWithDefaultName(ins interface{}) (bool, error)
	// Get 获取一个实例
	Get(instanceName string) interface{}
	// StoreFactory 设置一个对象工厂
	StoreFactory(instanceName string, instanceFactory *Factory[any])
}

type DefaultStore struct {
	instances        map[string]interface{}
	instancesFactory map[string]Factory[any]
}

// CreateInstanceFunc 创建实例的函数
type CreateInstanceFunc func(params ...interface{}) interface{}

var (
	store *DefaultStore
	once  sync.Once
)

func GetOrCreate() Store {
	if store == nil {
		once.Do(func() {
			store = &DefaultStore{
				instances:        make(map[string]interface{}),
				instancesFactory: make(map[string]Factory[any]),
			}
		})
	}

	return store
}

func (s *DefaultStore) Store(instanceName string, ins interface{}) (bool, error) {

	if ins == nil {
		return false, fmt.Errorf("store instance is nil")
	}
	if s.instances[instanceName] != nil {
		return false, fmt.Errorf("instance %s already exists", instanceName)
	}
	s.instances[instanceName] = ins
	return true, nil
}

// StoreInstance 根据实例名存储一个实例
func StoreInstance(instanceName string, ins interface{}) (bool, error) {
	return GetOrCreate().Store(instanceName, ins)

}

func (s *DefaultStore) StoreWithDefaultName(ins interface{}) (bool, error) {
	t := reflect.TypeOf(ins)
	return s.Store(t.String(), ins)
}

func (s *DefaultStore) Get(instanceName string) interface{} {
	if ins, ok := s.instances[instanceName]; ok {
		return ins
	}
	if f, ok := s.instancesFactory[instanceName]; ok {
		_, ins, err := f.GetOrCreate()
		if err != nil {
			return nil
		}
		s.instances[instanceName] = ins
		return ins

	}
	return nil
}

func (s *DefaultStore) StoreFactory(instanceName string, instanceFactory *Factory[any]) {
	s.instancesFactory[instanceName] = *instanceFactory
}

func StoreInstanceWithType[T any](ins T) (bool, error) {
	instanceName := reflect.TypeOf(new(T)).String()
	return StoreInstance(instanceName, ins)
}
func GetInstanceWithType[T any]() T {
	instanceName := reflect.TypeOf(new(T)).String()
	ins := GetOrCreate().Get(instanceName)
	return ins.(T)
}

func GetInstance[T any](instanceName string) T {
	ins := GetOrCreate().Get(instanceName)
	if ins == nil {
		return *new(T)
	}
	return ins.(T)
}

// GetOrCreateIns 范型方法根据实例名存储一个实例
func GetOrCreateIns[T any](instanceFunc CreateInstanceFunc) T {
	instanceName := reflect.TypeOf(new(T)).String()
	ins := GetOrCreate().Get(instanceName)
	if ins == nil {
		ins = instanceFunc()
		StoreInstance(instanceName, ins)
	}
	return ins.(T)
}

// GetOrCreateInsWithName 根据指定的实例名存储一个实例
func GetOrCreateInsWithName(instanceName string, instanceFunc CreateInstanceFunc) any {
	ins := GetOrCreate().Get(instanceName)
	if ins == nil {
		ins = instanceFunc()
		StoreInstance(instanceName, ins)
	}
	return ins
}

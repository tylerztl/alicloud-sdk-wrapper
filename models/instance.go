package models

import (
	"github.com/astaxie/beego/orm"
	"sync"
)

type Instance struct {
	Id               int              `orm:"column(id);auto" description:"id"`
	InstanceId       string           `orm:"column(instanceId)" description:"instanceId"`
	InstanceName     string           `orm:"column(instanceName)" description:"instanceName"`
	InstanceType     string           `orm:"column(instanceType)" description:"instanceType"`
	PublicIpAddress  SliceStringField `orm:"column(publicIpAddress)" description:"publicIpAddress"`
	VSwitchId        string           `orm:"column(vSwitchId)" description:"vSwitchId"`
	SecurityGroupIds SliceStringField `orm:"column(securityGroupIds)" description:"securityGroupIds"`
	CreationTime     string           `orm:"column(creationTime)" description:"creationTime"`
}

type InstanceModel struct {
	mutex  *sync.RWMutex
	mapper orm.Ormer
}

func NewInstanceModel() *InstanceModel {
	return &InstanceModel{
		mutex:  new(sync.RWMutex),
		mapper: orm.NewOrm(),
	}
}

func (m *InstanceModel) AddInstance(v *Instance) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.mapper.Insert(v)
}

func (m *InstanceModel) DeleteInstance(instanceId string) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.mapper.Delete(&Instance{InstanceId: instanceId})
}

func (m *InstanceModel) UpdateInstance(v *Instance) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if err := m.mapper.Read(&Instance{InstanceId: v.InstanceId}); err != nil {
		return 0, err
	}
	return m.mapper.Update(v)
}

func (m *InstanceModel) GetAllInstances() ([]orm.Params, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var maps []orm.Params
	_, err := m.mapper.Raw("SELECT * FROM instance").Values(&maps)
	return maps, err
}

func (m *InstanceModel) GetInstance(instanceId string) (v *Instance, err error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	u := &Instance{InstanceId: instanceId}
	err = m.mapper.Read(u)
	return u, err
}

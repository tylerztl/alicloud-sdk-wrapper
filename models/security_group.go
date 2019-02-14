package models

import (
	"sync"

	"github.com/astaxie/beego/orm"
)

type SecurityGroup struct {
	Id                int    `orm:"column(id);auto" description:"id"`
	SecurityGroupName string `orm:"column(securityGroupName)" description:"securityGroupName"`
	SecurityGroupId   string `orm:"column(securityGroupId)" description:"securityGroupId"`
	VpcId             string `orm:"column(vpcId)" description:"vpcId"`
	RegionId          string `orm:"column(regionId)" description:"regionId"`
}

type SecurityGroupModel struct {
	mutex  *sync.RWMutex
	mapper orm.Ormer
}

func NewSecurityGroupModel() *SecurityGroupModel {
	return &SecurityGroupModel{
		mutex:  new(sync.RWMutex),
		mapper: orm.NewOrm(),
	}
}

func (m *SecurityGroupModel) AddSecurityGroup(v *SecurityGroup) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.mapper.Insert(v)
}

func (m *SecurityGroupModel) DeleteSecurityGroup(sgId string) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.mapper.Delete(&SecurityGroup{SecurityGroupId: sgId})
}

func (m *SecurityGroupModel) UpdateSecurityGroup(v *SecurityGroup) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if err := m.mapper.Read(&SecurityGroup{SecurityGroupId: v.SecurityGroupId}); err != nil {
		return 0, err
	}
	return m.mapper.Update(v)
}

func (m *SecurityGroupModel) GetAllSecurityGroups() ([]orm.Params, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var maps []orm.Params
	_, err := m.mapper.Raw("SELECT * FROM security_group").Values(&maps)
	return maps, err
}

func (m *SecurityGroupModel) GetSecurityGroup(sgId string) (v *SecurityGroup, err error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	u := &SecurityGroup{SecurityGroupId: sgId}
	err = m.mapper.Read(u)
	return u, err
}

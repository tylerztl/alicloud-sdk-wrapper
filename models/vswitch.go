package models

import (
	"sync"

	"github.com/astaxie/beego/orm"
)

type VSwitch struct {
	Id          int    `orm:"column(id);auto" description:"id"`
	VSwitchName string `orm:"column(vSwitchName)" description:"vSwitchName"`
	VSwitchId   string `orm:"column(vSwitchId)" description:"vSwitchId"`
	VpcId       string `orm:"column(vpcId)" description:"vpcId"`
	RegionId    string `orm:"column(regionId)" description:"regionId"`
	ZoneId      string `orm:"column(zoneId)" description:"zoneId"`
	CidrBlock   string `orm:"column(cidrBlock)" description:"cidrBlock"`
}

type VSwitchModel struct {
	mutex  *sync.RWMutex
	mapper orm.Ormer
}

func NewVSwitchModel() *VSwitchModel {
	return &VSwitchModel{
		mutex:  new(sync.RWMutex),
		mapper: orm.NewOrm(),
	}
}

func (m *VSwitchModel) AddVSwitch(v *VSwitch) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.mapper.Insert(v)
}

func (m *VSwitchModel) DeleteVSwicth(vswitchId string) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.mapper.Delete(&VSwitch{VSwitchId: vswitchId})
}

func (m *VSwitchModel) UpdateVSwitch(v *VSwitch) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if err := m.mapper.Read(&VSwitch{VSwitchId: v.VSwitchId}); err != nil {
		return 0, err
	}
	return m.mapper.Update(v)
}

func (m *VSwitchModel) GetAllVSwitchs() ([]orm.Params, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var maps []orm.Params
	_, err := m.mapper.Raw("SELECT * FROM v_switch").Values(&maps)
	return maps, err
}

func (m *VSwitchModel) GetVSwitch(vswitchId string) (v *VSwitch, err error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	u := &VSwitch{VSwitchId: vswitchId}
	err = m.mapper.Read(u)
	return u, err
}

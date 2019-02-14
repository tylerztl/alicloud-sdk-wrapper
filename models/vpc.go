package models

import (
	"sync"

	"github.com/astaxie/beego/orm"
)

type Vpc struct {
	Id              int    `orm:"column(id);auto" description:"id"`
	VpcName         string `orm:"column(vpcName)" description:"vpcName"`
	VpcId           string `orm:"column(vpcId)" description:"vpcId"`
	RegionId        string `orm:"column(regionId)" description:"regionId"`
	CidrBlock       string `orm:"column(cidrBlock)" description:"cidrBlock"`
	VRouterId       string `orm:"column(vRouterId)" description:"vRouterId"`
	RouteTableId    string `orm:"column(routeTableId)" description:"routeTableId"`
	ResourceGroupId string `orm:"column(resourceGroupId)" description:"resourceGroupId"`
}

type VpcModel struct {
	mutex  *sync.RWMutex
	mapper orm.Ormer
}

func NewVpcModel() *VpcModel {
	return &VpcModel{
		mutex:  new(sync.RWMutex),
		mapper: orm.NewOrm(),
	}
}

func (m *VpcModel) AddVpc(v *Vpc) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.mapper.Insert(v)
}

func (m *VpcModel) DeleteVpc(vpcId string) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.mapper.Delete(&Vpc{VpcId: vpcId})
}

func (m *VpcModel) UpdateVpc(v *Vpc) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if err := m.mapper.Read(&Vpc{VpcId: v.VpcId}); err != nil {
		return 0, err
	}
	return m.mapper.Update(v)
}

func (m *VpcModel) GetAllVpcs() ([]orm.Params, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var maps []orm.Params
	_, err := m.mapper.Raw("SELECT * FROM vpc").Values(&maps)
	return maps, err
}

func (m *VpcModel) GetVpc(vpcId string) (v *Vpc, err error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	u := &Vpc{VpcId: vpcId}
	err = m.mapper.Read(u)
	return u, err
}

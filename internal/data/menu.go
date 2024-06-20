package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewMenuData() *MenuData {
	return &MenuData{}
}

type MenuData struct {
}

func (s MenuData) List() ([]*po.SysMenu, error) {
	list := make([]*po.SysMenu, 0)
	err := conn.GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("MenuData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s MenuData) Info(id int64) (*po.SysMenu, error) {
	info := &po.SysMenu{}
	err := conn.GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("MenuData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s MenuData) Add(menu *po.SysMenu) error {
	err := conn.GetMerchantDB().Model(&po.SysMenu{}).Where("id = ?", &menu.Id).Create(&menu).Error
	if err != nil {
		logrus.Errorf("MenuData Add Err: %s", err.Error())
		return err
	}
	return nil
}

func (s MenuData) Edit(menu *po.SysMenu) error {
	err := conn.GetMerchantDB().Model(&po.User{}).Where("id = ?", &menu.Id).Updates(menu).Error
	if err != nil {
		logrus.Errorf("MenuData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s MenuData) Delete(menu *po.SysMenu) error {
	err := conn.GetMerchantDB().Where("id = ?", &menu.Id).Delete(&menu).Error
	if err != nil {
		logrus.Errorf("MenuData Delete Err: %s", err.Error())
		return err
	}
	return nil
}

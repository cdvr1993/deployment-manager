package services

import (
	"github.com/astaxie/beego/orm"
)

type IOrmService interface {
	NewOrm() orm.Ormer
}
type OrmService struct{}

var (
	ormService = OrmService{}
)

func NewOrmService() IOrmService {
	return &ormService
}

func (s *OrmService) NewOrm() orm.Ormer {
	return orm.NewOrm()
}

package mdd

import (
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
)

//Server 多维度数据服务
type Server struct {
	mddDB      *xorm.Engine
	businessDB *xorm.Engine
	dimensions map[string]Dimension //TODO:后面有时间换成安全map
}

//NewServer 创建一个新的服务
func NewServer(mdb, bdb *xorm.Engine) *Server {
	return &Server{
		mddDB:      mdb,
		businessDB: bdb,
		dimensions: make(map[string]Dimension),
	}
}

//Sync 同步表结构
func (s Server) Sync() {
	s.syncTable(new(Dimension), new(StandardMember))
}
func (s Server) syncTable(bean ...interface{}) error {
	err := s.mddDB.Sync2(bean...)
	if err != nil {
		log.Printf("同步表结构失败：%s", err.Error())
	}
	return err
}

//LoadDimensionFromDB 从数据库里加载维度信息
func (s Server) LoadDimensionFromDB() (map[string]Dimension, error) {
	ds := make([]Dimension, 0)
	if err := s.mddDB.Find(&ds); err != nil {
		return nil, err
	}

	if s.dimensions == nil {
		s.dimensions = make(map[string]Dimension)
	}
	for _, v := range ds {
		s.dimensions[v.Name] = v
	}

	return s.dimensions, nil
}

//RegisterDimension 手工注册维度
func (s Server) RegisterDimension(d Dimension) error {
	if ok, msg := d.Check(); !ok {
		return fmt.Errorf(msg)
	}
	s.dimensions[d.Name] = d
	return nil
}

// func (s Server)RegisterDimensionByMember( m Member)error{
// 	d:=Dimension{
// 		Name=
// 	}
// }

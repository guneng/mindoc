package models

import (
	"time"
	"github.com/lifei6671/mindoc/conf"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type MemberGroup struct {
	GroupId int 		`orm:"column(group_id);pk;auto;unique;" json:"group_id"`
	GroupName string	`orm:"column(group_name);size(255);" json:"group_name"`
	GroupNumber int		`orm:"column(group_number);" json:"group_number"`
	CreateTime    time.Time `orm:"type(datetime);column(create_time);auto_now_add" json:"create_time"`
	CreateAt      int       `orm:"type(int);column(create_at)" json:"create_at"`
	ModifyTime time.Time     `orm:"column(modify_time);type(datetime);auto_now" json:"modify_time"`
	ModifyAt   int           `orm:"column(modify_at);type(int)" json:"-"`
}



// TableName 获取对应数据库表名.
func (m *MemberGroup) TableName() string {
	return "member_group"
}

// TableEngine 获取数据使用的引擎.
func (m *MemberGroup) TableEngine() string {
	return "INNODB"
}

func (m *MemberGroup) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + m.TableName()
}

func NewMemberGroup() *MemberGroup {
	return &MemberGroup{}
}

//根据id查询用户组
func (m *MemberGroup) FindFirst(id int) (*MemberGroup,error){
	o := orm.NewOrm()

	if err :=o.QueryTable(m.TableNameWithPrefix()).Filter("group_id",id).One(m); err != nil {
		beego.Error("查询用户组时出错 =>",err)
		return m,err
	}
	return m,nil
}




















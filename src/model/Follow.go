package model

import (
	"hmdp/src/config/mysql"
	"time"

	_ "github.com/jinzhu/gorm"
)

type Follow struct {
	Id  int64  `gorm:"primary;AUTO_INCREMENT;column:id" json:"id"`
	UserId int64  `gorm:"column:user_id" json:"userId"`
	FollowUserId int64  `gorm:"column:follow_user_id" json:"followUserId"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
}

func (*Follow) TableName() string {
	return "tb_follow"
}

func (f *Follow) RemoveUserFollow(id int64 , userId int64) error {
	err := mysql.GetMysqlDB().Table(f.TableName()).Where("user_id = ? and follow_user_id = ?" , userId , id).Delete(nil).Error
	return err
}

func (f *Follow) SaveUserFollow() error {
	err := mysql.GetMysqlDB().Table(f.TableName()).Create(f).Error
	return err
}

func (f *Follow) CountFollow() (int , error) {
	var count int
	err := mysql.GetMysqlDB().Table(f.TableName()).Where("user_id = ? and follow_user_id = ?" , f.UserId , f.Id).Count(&count).Error
	return count , err
}

func (f *Follow) GetFollowsByFollowId(id int64) ([]Follow , error) {
	var follows []Follow
	err := mysql.GetMysqlDB().Table(f.TableName()).Where("follow_user_id = ?" , id).Find(&follows).Error
	return follows , err
}

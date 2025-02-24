package model

import (
	"hmdp/src/config/mysql"
	"hmdp/src/utils"
	"time"

	_ "github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

const BLOG_TABLE_NAME = "tb_blog"

type Blog struct {
	Id         int64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	ShopId     int64     `gorm:"column:shop_id" json:"shopId"`
	UserId     int64     `gorm:"column:user_id" json:"userId"`
	Icon       string    `gorm:"-" json:"icon"`
	Name       string    `gorm:"-" json:"name"`
	IsLike     bool      `gorm:"-" json:"isLike"`
	Title      string    `gorm:"column:title" json:"title"`
	Images     string    `gorm:"column:images" json:"images"`
	Content    string    `gorm:"column:content" json:"content"`
	Liked      int       `gorm:"column:liked" json:"liked"`
	Comments   int       `gorm:"column:comments" json:"comments"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (*Blog) TableName() string {
	return BLOG_TABLE_NAME
}

func (blog *Blog) SaveBlog() (id int64 , err error){
	err = mysql.GetMysqlDB().Table(blog.TableName()).Create(blog).Error
	if err != nil {
		logrus.Error("[Blog model] insert data into database failed")	
		return id , err
	}
	id = blog.Id
	return id , nil
}

func (blog *Blog) IncreseLike() (err error) {
	err = mysql.GetMysqlDB().Table(blog.TableName()).Where("id = ?" , blog.Id).First(blog).Error
	if err != nil {
		return 
	}
	err = mysql.GetMysqlDB().Table(blog.TableName()).Where("id = ?" , blog.Id).Update("liked" , blog.Liked + 1).Error
	return 	
}

func (blog *Blog) QueryBlogs(current int) ([]Blog , error) {
	var blogs []Blog
	err := mysql.GetMysqlDB().Table(blog.TableName()).Where("user_id = ?" , blog.UserId).Offset((current - 1) * utils.MAXPAGESIZE).Limit(utils.MAXPAGESIZE).Find(&blogs).Error
	return blogs , err 
}

func (blog *Blog) QueryHots(current int) ([]Blog , error) {
	var blogs []Blog
	err := mysql.GetMysqlDB().Order("liked desc").Offset((current - 1) * utils.MAXPAGESIZE).Limit(utils.MAXPAGESIZE).Find(&blogs).Error
	return blogs , err
}

func (blog *Blog) GetBlogById(id int64) (error) {
	err := mysql.GetMysqlDB().Where("id = ?" , id).First(blog).Error
	return err
}

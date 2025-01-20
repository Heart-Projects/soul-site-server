package dao

import (
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type IUserDao interface {
	// FindUser 根据用户名查找用户
	FindUser(username string) *entity.SysUser

	UserInfo(userId uint64) *entity.SysUser

	FindUserArticleStatisticsData(userId uint64) *entity.UserArticleStatisticsData
}

type userDao struct {
	db *gorm.DB
}

func newUserDao(db *gorm.DB) IUserDao {
	return &userDao{db: db}
}

func (u *userDao) FindUser(username string) *entity.SysUser {
	var user entity.SysUser
	r := u.db.Where("name = ?", username).Take(&user)
	if r.Error != nil {
		logrus.Errorf("查找用户 %s 失败, error: %s", username, r.Error)
		return nil
	}
	return &user
}

func (u *userDao) UserInfo(userId uint64) *entity.SysUser {
	var user entity.SysUser
	r := u.db.Where("id = ?", userId).Take(&user)
	if r.Error != nil {
		logrus.Errorf("查找用户 %d 失败, error: %s", userId, r.Error)
		return nil
	}
	return &user
}

func (u *userDao) FindUserArticleStatisticsData(userId uint64) *entity.UserArticleStatisticsData {
	var statistics entity.UserArticleStatisticsData
	r := u.db.Where("user_id = ?", userId).Take(&statistics)
	if r.Error != nil {
		// 未找到, 给时间赋予当前时间，避免接口序列化以time 的空值序列化，出现奇怪的值
		statistics.UpdatedAt = types.DateTime{Time: time.Now()}
		statistics.CreatedAt = types.DateTime{Time: time.Now()}
		logrus.Errorf("查找用户 %d 文章统计数据失败, error: %s", userId, r.Error)
	}
	return &statistics
}

package service

import (
	"com.sj/admin/pkg/dao"
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/entity/vo"
	"errors"
)

type IUserService interface {
	CheckLogin(username, password string) (bool, error, *entity.SysUser)
	ArticleSummaryData(userId uint64) *vo.UserArticleData
}

type userService struct {
	userDao dao.IUserDao
}

func NewUserService(userDao dao.IUserDao) IUserService {
	return &userService{
		userDao: userDao,
	}
}

func (u *userService) CheckLogin(username, password string) (bool, error, *entity.SysUser) {
	user := u.userDao.FindUser(username)
	if user == nil {
		return false, errors.New("该用户不存在"), nil
	}
	if password != user.Password {
		return false, errors.New("用户名密码错误"), nil
	}
	return true, nil, user
}

func (u *userService) ArticleSummaryData(userId uint64) *vo.UserArticleData {
	user := u.userDao.UserInfo(userId)
	if user == nil {
		return nil
	}
	return &vo.UserArticleData{
		User: &vo.SimpleUserVo{
			ID:           user.ID,
			UserIdentify: user.UserIdentify,
			Name:         user.Name,
			Email:        user.Email,
			HomeUrl:      user.HomeUrl,
			Avatar:       user.Avatar,
		},
		ArticleData: u.userDao.FindUserArticleStatisticsData(userId),
	}
}

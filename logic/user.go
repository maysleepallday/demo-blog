package logic

import (
	"Blog/dao"
	"Blog/models"
	"Blog/param"
	"Blog/utils/jwt"
	"Blog/utils/snowflake"
	"errors"
)

func RegisterLogic(p *param.UserRegister) error {

	//判断用户名是否重复
	if dao.CheckUserExist(p.Name) {
		return errors.New("用户已存在")
	}

	//生成uid
	identity := snowflake.GenerateID()
	println(identity)

	//将新用户写入数据库
	user := &models.User{
		Identity: identity,
		Name:     p.Name,
		Password: p.Password,
	}
	err := dao.UserCreate(user)
	if err != nil {
		return errors.New("用户创建失败")
	}
	return nil
}

func LoginLogic(p *param.UserLogin) (*models.User, error) {
	// 检查用户名密码是否正确
	user, err := dao.UserLogin(p.Name, p.Password)
	if err != nil {
		return nil, errors.New("用户名和密码不正确")
	}

	//生成token
	tokenStr, err := jwt.GenerateToken(user.Identity, user.Name)
	if err != nil {
		return nil, errors.New("token生成失败")
	}

	user.Token = tokenStr

	return user, nil

}

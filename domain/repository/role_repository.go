package repository

import (
	"git.imooc.com/coding-535/user/domain/model"
	"github.com/jinzhu/gorm"
)
//创建需要实现的接口
type IRoleRepository interface{
	//根据ID查处找数据
	FindRoleByID(int64) (*model.Role, error)
	//创建一条 role 数据
	CreateRole(*model.Role) (int64, error)
	//根据ID删除一条 role 数据
	DeleteRoleByID(int64) error
	//修改更新数据
	UpdateRole(*model.Role) error
	//查找role所有数据
	FindAll()([]model.Role,error)

	//根据ID查找所有角色
	FindAllRoleById([]int64)([]*model.Role,error)

	//添加角色权限
	AddPermission(*model.Role,[]*model.Permission)error
	//更新角色权限
	UpdatePermission(*model.Role,[]*model.Permission)error
	//删除角色权限
	DeletePermission(*model.Role,[]*model.Permission)error
}
//创建RoleRepository
func NewRoleRepository(db *gorm.DB) IRoleRepository  {
	return &RoleRepository{mysqlDb:db}
}

type RoleRepository struct {
	mysqlDb *gorm.DB
}

//根据ID获取素有角色
func (u *RoleRepository) FindAllRoleById(id []int64) (roleAll []*model.Role, err error) {
	return roleAll,u.mysqlDb.Find(&roleAll,id).Error
}

//为角色添加权限
func (u *RoleRepository) AddPermission(role *model.Role,permission []*model.Permission) error {
	return u.mysqlDb.Model(&role).Association("Permission").Append(permission).Error
}

//为角更新加权限
func (u *RoleRepository) UpdatePermission(role *model.Role,permission []*model.Permission) error {
	return u.mysqlDb.Model(&role).Association("Permission").Replace(permission).Error
}

//删除角色权限
func (u *RoleRepository) DeletePermission(role *model.Role,permission []*model.Permission) error {
	return u.mysqlDb.Model(&role).Association("Permission").Delete(permission).Error
}

//根据ID查找Role信息
func (u *RoleRepository)FindRoleByID(roleID int64) (role *model.Role,err error) {
	role = &model.Role{}
	return role, u.mysqlDb.Preload("Permission").First(role,roleID).Error
}
//创建Role信息
func (u *RoleRepository) CreateRole(role *model.Role) (int64, error) {
	return role.ID, u.mysqlDb.Create(role).Error
}

//根据ID删除Role信息
func (u *RoleRepository) DeleteRoleByID(roleID int64) error {
	return u.mysqlDb.Where("id = ?",roleID).Delete(&model.Role{}).Error
}

//更新Role信息
func (u *RoleRepository) UpdateRole(role *model.Role) error {
	return u.mysqlDb.Model(role).Update(role).Error
}

//获取结果集
func (u *RoleRepository) FindAll()(roleAll []model.Role,err error) {
	return roleAll, u.mysqlDb.Find(&roleAll).Error
}




package service

import (
	"git.imooc.com/coding-535/user/domain/model"
	"git.imooc.com/coding-535/user/domain/repository"
	"k8s.io/client-go/kubernetes"
)

//这里是接口类型
type IPermissionDataService interface {
	AddPermission(*model.Permission) (int64 , error)
	DeletePermission(int64) error
	UpdatePermission(*model.Permission) error
	FindPermissionByID(int64) (*model.Permission, error)
	FindAllPermission() ([]model.Permission, error)

	//根据ID查询所有权限
	FindAllPermissionById([]int64)([]*model.Permission,error)

}


//创建
//注意：返回值 IPermissionDataService 接口类型
func NewPermissionDataService(PermissionRepository repository.IPermissionRepository,clientSet *kubernetes.Clientset) IPermissionDataService{
	return &PermissionDataService{ PermissionRepository:PermissionRepository}
}

type PermissionDataService struct {
	//注意：这里是 IPermissionRepository 类型
	PermissionRepository repository.IPermissionRepository

}

func (u *PermissionDataService) FindAllPermissionById(id []int64) ([]*model.Permission, error) {
	return u.PermissionRepository.FindAllPermissionById(id)
}

//插入
func (u *PermissionDataService) AddPermission(Permission *model.Permission) (int64 ,error) {
	return u.PermissionRepository.CreatePermission(Permission)
}

//删除
func (u *PermissionDataService) DeletePermission(PermissionID int64) error {
	return u.PermissionRepository.DeletePermissionByID(PermissionID)
}

//更新
func (u *PermissionDataService) UpdatePermission(Permission *model.Permission) error {
	return u.PermissionRepository.UpdatePermission(Permission)
}

//查找
func (u *PermissionDataService) FindPermissionByID(PermissionID int64) (*model.Permission, error) {
	return u.PermissionRepository.FindPermissionByID(PermissionID)
}

//查找
func (u *PermissionDataService) FindAllPermission() ([]model.Permission, error) {
	return u.PermissionRepository.FindAll()
}


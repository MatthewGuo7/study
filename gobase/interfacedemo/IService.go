package interfacedemo

import (
	"fmt"
	"testing"
)

type IService interface {
	Save()
}

type UserService struct {

}

func (u *UserService) Save() {
	fmt.Println("user service save")
}

func NewUserService() *UserService {
	return &UserService{}
}

type ProductService struct {

}

func (p *ProductService) Save() {
	fmt.Println("product service save")
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func SaveService(service IService)  {
	service.Save()
}

func TestInterface(t *testing.T) {
	SaveService(NewUserService())
}

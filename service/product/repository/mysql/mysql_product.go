package mysql

import (
	"project-version3/superindo-task/service/domain"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	Conn *gorm.DB
}

func NewMysqlProductRepository(conn *gorm.DB) domain.ProductRepository {
	return &mysqlProductRepository{conn}
}

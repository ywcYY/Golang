package dao

import (
	"gorm/entity"

	"github.com/jinzhu/gorm"
)

type OrderDAO interface {
	FindAllOrders(db *gorm.DB) ([]entity.Order, error)
	FindOrderByID(db *gorm.DB, id int) (entity.Order, error)
	CreateOrder(db *gorm.DB, order entity.Order) error
	UpdateOrder(db *gorm.DB, order entity.Order) error
	DeleteOrder(db *gorm.DB, order entity.Order) error
}

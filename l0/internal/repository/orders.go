package repository

import "github.com/jmoiron/sqlx"

type OrdersRepository struct {
	*sqlx.DB
}

func (or *OrdersRepository) CreateOrder() {

}

func (or *OrdersRepository) GetOrder() {

}

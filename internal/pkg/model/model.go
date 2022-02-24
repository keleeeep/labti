/*
 * @Author: Adrian Faisal
 * @Date: 31/08/21 1.03 PM
 */

package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	SKU       string    `json:"sku"`
	Stock     int       `json:"stock"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Product) TableName() string {
	return "products"
}

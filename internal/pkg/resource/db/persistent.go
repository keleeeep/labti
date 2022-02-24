/*
 * @Author: Adrian Faisal
 * @Date: 02/09/21 9.17 PM
 */

package db

import (
	"context"
	"fmt"
	"github.com/apldex/workshop-labti/internal/pkg/model"
)

func (p *persistent) GetUserByUsername(username string) (*model.User, error) {
	panic("not implemented")
}

func (p *persistent) CreateProduct(ctx context.Context, product *model.Product)  error {
	query := fmt.Sprintf("INSERT INTO %s (name, sku, stock, price) VALUES (?, ?, ?, ?)",
		product.TableName())

	tx, err := p.conn.Begin()
	if err != nil {
		return fmt.Errorf("can't start db transaction: %v", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		query, product.Name, product.SKU, product.Stock, product.Price)
	if err != nil {
		return fmt.Errorf("exec query failed: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commiting transaction failed: %v", err)
	}

	return nil
}

func (p *persistent) GetProduct(ctx context.Context, id int) (*model.Product, error) {
	m := &model.Product{} // struct literal

	query := fmt.Sprintf("SELECT id, name, sku, stock, price FROM %s WHERE id = ?",
		m.TableName())

	row := p.conn.QueryRowContext(ctx, query, id)
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("query row error: %v", err)
	}

	err := row.Scan(&m.ID, &m.Name, &m.SKU, &m.Stock, &m.Price)
	if err != nil {
		return nil , fmt.Errorf("failed to scan row result: %v", err)
	}

	return m, nil
}

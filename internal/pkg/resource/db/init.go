/*
 * @Author: Adrian Faisal
 * @Date: 31/08/21 1.02 PM
 */

package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/apldex/workshop-labti/internal/pkg/model"

	_ "github.com/go-sql-driver/mysql"
)

type Persistent interface {
	GetUserByUsername(username string) (*model.User, error)
	CreateProduct(ctx context.Context, product *model.Product) error
	GetProduct(ctx context.Context, id int) (*model.Product, error)
}

type persistent struct {
	conn *sql.DB
}

func NewPersistent(datasource string) (Persistent, error) {
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		return nil, fmt.Errorf("open database connection failed: %v", err)
	}

	// ping a database connection is recommended to verify the connection still alive
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping database failed: %v", err)
	}

	return &persistent{conn: db}, nil
}

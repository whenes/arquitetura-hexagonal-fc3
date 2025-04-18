package db

import (
	"database/sql"

	"github.com/codeedu/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(produto application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	err := p.db.QueryRow("SELECT COUNT(*) FROM products WHERE id = ?", produto.GetID()).Scan(&rows)
	if err != nil {
		_, err = p.create(produto)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(produto)
		if err != nil {
			return nil, err
		}
	}
	return produto, nil
}

func (p *ProductDb) create(produto application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(produto.GetID(), produto.GetName(), produto.GetPrice(), produto.GetStatus())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return produto, nil
}

func (p *ProductDb) update(produto application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec(`UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?`,
		produto.GetName(), produto.GetPrice(), produto.GetStatus(), produto.GetID())
	if err != nil {
		return nil, err
	}
	return produto, nil
}



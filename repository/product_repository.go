package repository

import (
	"database/sql"
	"fmt"
	"game-shop-api/model"
)

type ProductRespository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRespository {
	return ProductRespository{
		connection: connection,
	}
}

func (pr *ProductRespository) GetProducts() []model.Procuct {
	query := "SELECT * FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Procuct{}
	}

	var productList []model.Procuct
	var productObj model.Procuct

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
			&productObj.Description)
	}
	if err != nil {
		fmt.Println(err)
		return []model.Procuct{}
	}

	productList = append(productList, productObj)

	rows.Close()

	return productList
}

func (pr *ProductRespository) CreateProduct(product model.Procuct) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product(name, price, description) VALUES($1, $2, $3) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price, product.Description).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}

func (pr *ProductRespository) GetProductsById(id int) (*model.Procuct, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Procuct

	err = query.QueryRow(id).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
		&produto.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &produto, nil
}

// func (pr *ProductRespository) DeleteProducts(id int) (model.Procuct, error) {
// 	query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")

// }

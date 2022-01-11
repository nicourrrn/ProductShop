package db

import (
	"github.com/nicourrrn/ProductShop/pkg/core"
)

func Get(c Connector, key, value string) (*core.Client, error) {
	row := c.QueryRow("SELECT id, first_name, last_name, address, basket_id, phone, email FROM clients WHERE ? = ?", key, value)
	var client core.Client
	var basketId int
	err := row.Scan(&client.Id, &client.FirstName, &client.LastName,
		&client.Address, &basketId, &client.Phone, &client.Email)
	if err != nil {
		return nil, err
	}
	row = c.QueryRow("SELECT * FROM baskets WHERE id = ?", basketId)
	var backet core.Basket
	err = row.Scan(backet.Id, backet.Address, backet.Paid, backet.Close, backet.Total)
	if err != nil {
		return nil, err
	}
	client.Basket = backet
	return &client, nil
}

func Add(c Connector, client *core.Client, passwordHash string) (int64, error) {
	t, err := c.Begin()
	if err != nil {
		return 0, err
	}
	resultBasket, err := t.Exec("INSERT INTO baskets(address, paid, close, total) VALUE (?, false, false, 0)", client.Address)
	if err != nil {
		return 0, err
	}
	basketID, err := resultBasket.LastInsertId()
	if err != nil {
		return 0, err
	}
	resultClient, err := t.Exec("INSERT INTO clients("+
		"first_name, last_name, address, basket_id, phone, email, password)"+
		"VALUES ?, ?, ?, ?, ?, ?, ?",
		client.FirstName, client.LastName, client.Address, basketID,
		client.Phone, client.Email, passwordHash)
	if err != nil {
		rollError := t.Rollback()
		if rollError != nil {
			return 0, rollError
		}
		return 0, err
	}
	clientId, err := resultClient.LastInsertId()
	if err != nil {
		rollError := t.Rollback()
		if rollError != nil {
			return 0, rollError
		}
		return 0, err
	}
	_, err = t.Exec("INSERT INTO clients_baskets VALUE (?, ?)", clientId, basketID)
	if err != nil {
		rollError := t.Rollback()
		if rollError != nil {
			return 0, rollError
		}
		return 0, err
	}
	return clientId, nil
}

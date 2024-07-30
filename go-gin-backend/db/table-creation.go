package db

import (
	_ "github.com/go-sql-driver/mysql"
)

func InitTables() {
	_ = createUserTable()
	_ = createStoreTable()
	_ = createBillboardTable()
	_ = createCategoryTable()
	_ = createSizeTable()
	_ = createColorTable()
	_ = createProductTable()
	_ = createImageTable()
	_ = createOrderTable()
	_ = createOrderItemTable()
}

func createUserTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS users(
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		hashedPassword VARCHAR(255) NOT NULL,

		PRIMARY KEY(id)
		)
	`
	_, err := DB.Exec(query)

	return err
}

func createStoreTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS Store (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		userId CHAR(36) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (userId) REFERENCES users(id)
		)
	`
	_, err := DB.Exec(query)
	return err
}

func createBillboardTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS Billboard (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		label VARCHAR(50) NOT NULL,
		storeId CHAR(36) NOT NULL,
		imageUrl VARCHAR(500),
		PRIMARY KEY (id),
		FOREIGN KEY (storeId) REFERENCES Store(id)
		)
	`
	_, err := DB.Exec(query)
	return err
}

func createCategoryTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS Category (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		storeId CHAR(36) NOT NULL,
		billboardId CHAR(46) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (storeId) REFERENCES Store(id),
		FOREIGN KEY (billboardId) REFERENCES Billboard(id)
		)
	`
	_, err := DB.Exec(query)
	return err
}

func createSizeTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS Size (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		value VARCHAR(50) NOT NULL,
		storeId CHAR(36) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (storeId) REFERENCES Store(id)
		)
	`
	_, err := DB.Exec(query)
	return err
}

func createColorTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS Color (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		value VARCHAR(50) NOT NULL,
		storeId CHAR(36) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (storeId) REFERENCES Store(id)
		)
	`
	_, err := DB.Exec(query)
	return err
}

func createProductTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS Product (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		price VARCHAR(20) NOT NULL,
		isFeatured BOOLEAN NOT NULL DEFAULT FALSE,
		isArchived BOOLEAN NOT NULL DEFAULT FALSE,
		storeId CHAR(36) NOT NULL,
		categoryId CHAR(36) NOT NULL,
		sizeId CHAR(36) NOT NULL,
		colorId CHAR(36) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (storeId) REFERENCES Store(id),
		FOREIGN KEY (categoryId) REFERENCES Category(id),
		FOREIGN KEY (sizeId) REFERENCES Size(id),
		FOREIGN KEY (colorId) REFERENCES Color(id)
		)
	`
	_, err := DB.Exec(query)
	return err
}

func createImageTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS Image (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		url VARCHAR(50) NOT NULL,
		productId CHAR(36) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (productId) REFERENCES Product(id)
		)
	`
	_, err := DB.Exec(query)
	return err
}

func createOrderTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS Orders (
	id CHAR(36) NOT NULL,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	isPaid BOOLEAN NOT NULL DEFAULT FALSE,
	phone VARCHAR(20) NOT NULL,
	address VARCHAR(100) NOT NULL,
	storeId CHAR(36) NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (storeId) REFERENCES Store(id)
	)
`
	_, err := DB.Exec(query)
	return err
}

func createOrderItemTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS OrderItem (
	id CHAR(36) NOT NULL,
	productId CHAR(36) NOT NULL,
	orderId CHAR(36) NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (productId) REFERENCES Product(id),
	FOREIGN KEY (orderId) REFERENCES Orders(id)
	)
`
	_, err := DB.Exec(query)
	return err
}

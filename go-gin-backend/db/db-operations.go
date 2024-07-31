package db

import (
	"admin-dashboard/backend/golan-gin/types"
	"fmt"

	"github.com/google/uuid"
)

func CreateUser(name string, email string, password string) (uuid.UUID, error) {
	id := uuid.New()

	query := `INSERT INTO users (id,name,email,password) VALUES (? , ?, ?,?)`
	_, err := DB.Exec(query, id, name, email, password)
	return id, err
}

func CreateStore(name string, userId string) (uuid.UUID, error) {
	id := uuid.New()

	query := `INSERT INTO store (id, name,userId) VALUES(?,?,?)`
	_, err := DB.Exec(query, id, name, userId)
	return id, err
}

func CreateBillboard(label string, storeId string, imageUrl string) error {
	id := uuid.New()

	query := `INSERT INTO billboard (id,label,storeId,ImageUrl) VALUES(?,?,?,?)`
	_, err := DB.Exec(query, id, label, storeId, imageUrl)
	return err
}

func CreateCategory(name string, storeId string, billboardId string) error {
	id := uuid.New()

	query := `INSERT INTO category (id,name,storeId,billboardId) VALUES(?,?,?,?)`
	_, err := DB.Exec(query, id, name, storeId, billboardId)
	return err
}

func CreateColor(name string, storeId string, value string) error {
	id := uuid.New()

	query := `INSERT INTO color (id,name,storeId,value) VALUES(?,?,?,?)`
	_, err := DB.Exec(query, id, name, storeId, value)
	return err
}

func CreateImage(url string, productId string) error {
	id := uuid.New()

	query := `INSERT INTO image (id,url,productId) VALUES(?,?,?)`
	_, err := DB.Exec(query, id, url, productId)
	return err
}

func CreateOrder(storeId string, isPaid bool, phone string, address string) error {
	id := uuid.New()

	query := `INSERT INTO orders (id,storeId,isPaid,phone,address) VALUES(?,?,?,?,?)`
	_, err := DB.Exec(query, id, storeId, isPaid, phone, address)
	return err
}

func CreateOrderItem(orderId string, productId string) error {
	id := uuid.New()

	query := `INSERT INTO orderItem (id,orderId,productId) VALUES(?,?,?)`
	_, err := DB.Exec(query, id, orderId, productId)
	return err
}

func CreateSize(name string, value string, storeId string) error {
	id := uuid.New()

	query := `INSERT INTO size (id,name,value,storeId) VALUES(?,?,?,?)`
	_, err := DB.Exec(query, id, name, value, storeId)
	return err
}

func CreateProduct(name string, categoryId string, storeId string, isFeatured bool, isArchived bool, sizeId string, colorId string, price string) (uuid.UUID, error) {
	id := uuid.New()

	query := `INSERT INTO product (id,name,storeId,categoryId,price,isFeatured,isArchived,sizeId,colorId) VALUES(?,?,?,?,?,?,?,?,?)`
	_, err := DB.Exec(query, id, name, storeId, categoryId, price, isFeatured, isArchived, sizeId, colorId)
	return id, err
}

func GetStore(store *types.FullStoreType, filter string, values ...any) error {

	query := fmt.Sprintf("SELECT * FROM store WHERE %s", filter)

	err := DB.QueryRow(query, values...).Scan(&store.ID, &store.CreatedAt, &store.UpdatedAt, &store.Name, &store.UserId)

	return err
}

func GetUser(user *types.User, filter string, values ...any) error {
	query := fmt.Sprintf("SELECT id,name,email,createdAt,updatedAt,password FROM users WHERE %s ", filter)

	err := DB.QueryRow(query, values...).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Password)

	return err
}

func GetBillboard(billboard *types.Billboard, filter string, values ...any) error {
	query := fmt.Sprintf("SELECT * FROM billboard WHERE %s", filter)

	err := DB.QueryRow(query, values...).Scan(&billboard.ID, &billboard.CreatedAt, &billboard.UpdatedAt, &billboard.Label, &billboard.StoreId, &billboard.ImageUrl)
	return err
}

func GetCategory(category *types.Category, filter string, values ...any) error {
	query := fmt.Sprintf("SELECT id,createdAt,updatedAt,storeId,name,billboardId FROM category WHERE %s", filter)
	err := DB.QueryRow(query, values...).Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt, &category.StoreId, &category.Name, &category.BillboardId)
	return err
}

func GetColor(color *types.Color, filter string, values ...any) error {
	query := fmt.Sprintf("SELECT id,name,value,storeId,createdAt,updatedAt FROM color WHERE  %s ", filter)
	err := DB.QueryRow(query, values...).Scan(&color.ID, &color.Name, &color.Value, &color.StoreId, &color.CreatedAt, &color.UpdatedAt)
	return err
}

func GetSize(size *types.Size, filter string, values ...any) error {
	query := fmt.Sprintf("SELECT id,name,value,storeId,createdAt,updatedAt FROM size WHERE %s", filter)
	err := DB.QueryRow(query, values...).Scan(&size.ID, &size.Name, &size.Value, &size.StoreId, &size.CreatedAt, &size.UpdatedAt)
	return err
}

func GetProduct(product *types.FullProduct, filter string, values ...any) error {
	query := fmt.Sprintf(`
			SELECT 
			id,
			createdAt,
			updatedAt,
			name,
			price,
			isFeatured,
			isArchived,
			storeId,
			categoryId,
			sizeId,
			colorId
			FROM product
			WHERE %s`, filter)

	err := DB.QueryRow(query, values...).Scan(
		&product.ID,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.Name,
		&product.Price,
		&product.IsFeatured,
		&product.IsArchived,
		&product.StoreId,
		&product.CategoryId,
		&product.SizeId,
		&product.ColorId)
	return err
}

func GetImage(image *types.Image, filter string, values ...any) error {
	query := fmt.Sprintf("SELECT * FROM image WHERE %s", filter)
	err := DB.QueryRow(query, values...).Scan(&image.ID, &image.CreatedAt, &image.UpdatedAt, &image.ProductId, &image.Url)
	return err
}

func GetOrder(order *types.Order, filter string, values ...any) error {
	query := fmt.Sprintf("SELECT * FROM order WHERE %s", filter)
	err := DB.QueryRow(query, values...).Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt, &order.IsPaid, &order.Phone, &order.Address, &order.StoreId)
	return err
}

func GetOrderItem(orderItem *types.OrderItem, filter string, values ...any) error {
	query := fmt.Sprintf("SELECT * FROM orderitem WHERE %s", filter)
	err := DB.QueryRow(query, values...).Scan(&orderItem.ID, &orderItem.ProductId, &orderItem.OrderId)
	return err
}

func GetFullOrderItem(orderItem *types.FullOrderItem, filter string, values ...any) error {
	query := fmt.Sprintf(`SELECT 
							id,
							orderId,
							productId
						FROM order
						WHERE %s
						`, filter)
	err := DB.QueryRow(query, values...).Scan(
		&orderItem.ID,
		&orderItem.OrderId,
		&orderItem.ProductId,
	)

	GetProduct(&orderItem.Product, " id = ? ", orderItem.ProductId)

	return err
}

func GetCategories(filter string, values ...any) []types.Category {
	var categories []types.Category
	query := fmt.Sprintf("SELECT id,createdAt,updatedAt,storeId,name,billboardId FROM category WHERE  %s", filter)
	rows, _ := DB.Query(query, values...)

	defer rows.Close()

	for rows.Next() {
		var category types.Category
		err := rows.Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt, &category.StoreId, &category.Name, &category.BillboardId)
		if err != nil {
			fmt.Println(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func GetFullCategories(filter string, values ...any) []types.FullCategory {
	var categories []types.FullCategory
	query := fmt.Sprintf("SELECT id,createdAt,updatedAt,storeId,name,billboardId FROM category WHERE %s", filter)
	rows, _ := DB.Query(query, values...)

	defer rows.Close()

	for rows.Next() {
		var category types.FullCategory
		err := rows.Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt, &category.StoreId, &category.Name, &category.BillboardId)
		if err != nil {
			fmt.Println(err)
		}

		err = GetBillboard(&category.Billboard, " id = ? ", category.BillboardId)

		if err != nil {
			fmt.Println(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func GetBillboards(filter string, values ...any) ([]types.Billboard, error) {
	var billboards []types.Billboard
	query := fmt.Sprintf("SELECT * FROM billboard WHERE %s", filter)
	rows, err := DB.Query(query, values...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var billboard types.Billboard
		err := rows.Scan(&billboard.ID, &billboard.CreatedAt, &billboard.UpdatedAt, &billboard.Label, &billboard.StoreId, &billboard.ImageUrl)
		if err != nil {
			fmt.Println(err)
		}
		billboards = append(billboards, billboard)
	}

	return billboards, nil
}

func GetColors(filter string, values ...any) []types.Color {
	var colors []types.Color
	query := fmt.Sprintf("SELECT id,name,value,storeId,createdAt,updatedAt FROM color WHERE  %s", filter)
	rows, _ := DB.Query(query, values...)

	defer rows.Close()

	for rows.Next() {
		var color types.Color
		err := rows.Scan(&color.ID, &color.Name, &color.Value, &color.StoreId, &color.CreatedAt, &color.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}
		colors = append(colors, color)
	}

	return colors
}

func GetSizes(filter string, values ...any) []types.Size {
	var sizes []types.Size
	query := fmt.Sprintf("SELECT id,name,value,storeId,createdAt,updatedAt FROM size WHERE  %s", filter)
	rows, _ := DB.Query(query, values...)

	defer rows.Close()

	for rows.Next() {
		var size types.Size
		err := rows.Scan(&size.ID, &size.Name, &size.Value, &size.StoreId, &size.CreatedAt, &size.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}
		sizes = append(sizes, size)
	}

	return sizes
}

func GetImages(filter string, values ...any) []types.Image {
	var images []types.Image
	query := fmt.Sprintf("SELECT id,createdAt,updatedAt,productId,url FROM image WHERE %s", filter)
	rows, _ := DB.Query(query, values...)

	defer rows.Close()

	for rows.Next() {
		var image types.Image
		err := rows.Scan(&image.ID, &image.CreatedAt, &image.UpdatedAt, &image.ProductId, &image.Url)
		if err != nil {
			fmt.Println(err)
		}
		images = append(images, image)
	}

	return images
}

func GetOrders(filter string, values ...any) []types.FullOrder {
	var orders []types.FullOrder
	query := fmt.Sprintf("SELECT * FROM orders WHERE %s", filter)
	rows, _ := DB.Query(query, values...)

	defer rows.Close()

	for rows.Next() {
		var order types.FullOrder
		err := rows.Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt, &order.IsPaid, &order.Phone, &order.Address, &order.StoreId)
		if err != nil {
			fmt.Println(err)
		}
		GetFullOrderItem(&order.OrderItem, " orderId = ? ", order.ID)
		orders = append(orders, order)
	}

	return orders
}

func GetOrderItems(filter string, values ...any) []types.OrderItem {
	var orderItems []types.OrderItem
	query := fmt.Sprintf("SELECT * FROM orderitem WHERE %s", filter)
	rows, _ := DB.Query(query, values...)

	defer rows.Close()

	for rows.Next() {
		var orderItem types.OrderItem
		err := rows.Scan(&orderItem.ID, &orderItem.ProductId, &orderItem.OrderId)
		if err != nil {
			fmt.Println(err)
		}
		orderItems = append(orderItems, orderItem)
	}

	return orderItems
}

func GetProducts(filter string, values ...any) []types.Product {
	var products []types.Product
	query := fmt.Sprintf("SELECT * FROM product WHERE %s", filter)
	rows, _ := DB.Query(query, values...)

	defer rows.Close()

	for rows.Next() {
		var product types.Product
		err := rows.Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt, &product.StoreId, &product.Name, &product.Price, &product.IsArchived, &product.IsFeatured, &product.CategoryId, &product.SizeId, &product.ColorId)
		if err != nil {
			fmt.Println(err)
		}
		products = append(products, product)
	}

	return products
}
func GetFullProducts(filter string, values ...any) []types.FullProduct {
	var products []types.FullProduct
	query := fmt.Sprintf(`
		SELECT 
			id,
			createdAt,
			updatedAt,
			name,
			price,
			isFeatured,
			isArchived,
			storeId,
			categoryId,
			sizeId,
			colorId 
		FROM product WHERE %s`, filter)

	rows, _ := DB.Query(query, values...)

	defer rows.Close()

	for rows.Next() {
		var product types.FullProduct
		err := rows.Scan(
			&product.ID,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Name,
			&product.Price,
			&product.IsFeatured,
			&product.IsArchived,
			&product.StoreId,
			&product.CategoryId,
			&product.SizeId,
			&product.ColorId)
		if err != nil {
			fmt.Println(err)
		}

		product.Image = GetImages(" productId = ? ", product.ID)
		err = GetColor(&product.Color, "id = ?", product.ColorId)
		if err != nil {
			fmt.Println(err)
		}

		err = GetSize(&product.Size, "id = ?", product.SizeId)
		if err != nil {
			fmt.Println(err)
		}

		err = GetCategory(&product.Category, "id = ?", product.CategoryId)
		if err != nil {
			fmt.Println(err)
		}

		products = append(products, product)
	}

	return products
}

func GetFullStores(filter string, values ...any) []types.FullStoreType {
	var stores []types.FullStoreType
	query := fmt.Sprintf(`
		SELECT 
			id,
			createdAt,
			updatedAt,
			name,
			userId
		FROM store WHERE %s`, filter)
	rows, _ := DB.Query(query, values...)

	defer rows.Close()

	for rows.Next() {
		var store types.FullStoreType
		err := rows.Scan(

			&store.ID,
			&store.CreatedAt,
			&store.UpdatedAt,
			&store.Name,
			&store.UserId)

		if err != nil {
			fmt.Println(err)
		}

		store.Colors = GetColors(" storeId = ? ", store.ID)
		store.Sizes = GetSizes(" storeId = ? ", store.ID)
		store.Categories = GetCategories(" storeId = ? ", store.ID)

		stores = append(stores, store)
	}

	return stores
}

func DeleteItem(tableName string, filter string, values ...any) error {
	query := fmt.Sprintf("DELETE FROM %s  %s", tableName, filter)
	_, err := DB.Exec(query, values...)
	return err
}

func UpdateItem(tableName string, updateString string, values ...any) error {
	query := fmt.Sprintf("UPDATE %s %s ", tableName, updateString)
	_, err := DB.Exec(query, values...)
	return err
}

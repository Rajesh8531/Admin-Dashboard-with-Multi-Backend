package db

import (
	"connection-to-mongo/project/types"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetStoreById(store *types.Store, filter primitive.M) {
	_ = Stores.FindOne(context.Background(), filter).Decode(&store)
}

func GetCategory(filter primitive.M) types.FullCategory {
	var fullCategory types.FullCategory
	var category types.Category
	var billboard types.Billboard

	_ = Categories.FindOne(context.Background(), filter).Decode(&category)
	_ = Billboards.FindOne(context.Background(), bson.M{"_id": category.BillboardId}).Decode(&billboard)
	fullCategory.BillboardId = category.BillboardId
	fullCategory.CreatedAt = category.CreatedAt
	fullCategory.ID = category.ID
	fullCategory.Name = category.Name
	fullCategory.StoreId = category.StoreId
	fullCategory.UpdatedAt = category.UpdatedAt
	fullCategory.Billboard = billboard
	return fullCategory
}

func GetFullCategory(category *types.Category, filter primitive.M) {
	_ = Categories.FindOne(context.Background(), filter).Decode(&category)
}

func GetBillboardById(billboard *types.Billboard, filter primitive.M) {
	_ = Billboards.FindOne(context.Background(), filter).Decode(&billboard)
}

func GetColorById(color *types.Color, filter primitive.M) {
	_ = Colors.FindOne(context.Background(), filter).Decode(&color)
}

func GetImageById(image *types.Image, filter primitive.M) {
	_ = Images.FindOne(context.Background(), filter).Decode(&image)
}

func GetOrderById(order *types.Order, filter primitive.M) {
	_ = Orders.FindOne(context.Background(), filter).Decode(&order)
}

func GetOrderItemById(orderItem *types.OrderItem, filter primitive.M) {
	_ = OrderItems.FindOne(context.Background(), filter).Decode(&orderItem)
}

func GetProduct(product *types.FullProduct, filter primitive.M) {
	var prod types.FullProduct
	_ = Products.FindOne(context.Background(), filter).Decode(&prod)
	product.CategoryId = prod.CategoryId
	product.ColorId = prod.ColorId
	product.SizeId = prod.SizeId
	product.StoreId = prod.StoreId
	product.ID = prod.ID
	product.Name = prod.Name
	product.Price = prod.Price
	product.CreatedAt = prod.CreatedAt
	product.UpdatedAt = prod.UpdatedAt
	product.IsArchived = prod.IsArchived
	product.IsFeatured = prod.IsFeatured

	var color types.Color
	var size types.Size
	var image []types.Image
	image, _ = GetImagesByValue(bson.M{"productId": product.ID})
	product.Image = image
	category := GetCategory(bson.M{"_id": prod.CategoryId})
	product.Category = category
	GetSizeById(&size, bson.M{"_id": prod.SizeId})
	GetColorById(&color, bson.M{"_id": prod.ColorId})
	product.Size = size
	product.Color = color
}

func GetSizeById(size *types.Size, filter primitive.M) {
	_ = Sizes.FindOne(context.Background(), filter).Decode(&size)
}

func GetStoresByValue(filter primitive.M) ([]types.Store, error) {
	var stores []types.Store
	cursor, err := Stores.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var store types.Store
		err := cursor.Decode(&store)
		if err != nil {
			log.Fatal(err)
		}

		stores = append(stores, store)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return stores, nil
}

func GetCategoriesByValue(filter primitive.M) ([]types.FullCategory, error) {
	var categories []types.FullCategory

	cursor, err := Categories.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var category types.FullCategory
		err := cursor.Decode(&category)
		if err != nil {
			log.Fatal(err)
		}
		var billboard types.Billboard
		GetBillboardById(&billboard, bson.M{"_id": category.BillboardId})

		category.Billboard = billboard
		categories = append(categories, category)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return categories, nil
}

func GetCategoriesByValueData(filter primitive.M) ([]types.Category, error) {
	var categories []types.Category
	cursor, err := Categories.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var category types.Category
		err := cursor.Decode(&category)
		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, category)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return categories, nil
}

func GetBillboardsByValue(filter primitive.M) ([]types.Billboard, error) {
	var billboards []types.Billboard
	cursor, err := Billboards.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var billboard types.Billboard
		err := cursor.Decode(&billboard)
		if err != nil {
			log.Fatal(err)
		}

		billboards = append(billboards, billboard)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return billboards, nil
}

func GetColorsByValue(filter primitive.M) ([]types.Color, error) {
	var colors []types.Color
	cursor, err := Colors.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var color types.Color
		err := cursor.Decode(&color)
		if err != nil {
			log.Fatal(err)
		}

		colors = append(colors, color)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return colors, nil
}

func GetImagesByValue(filter primitive.M) ([]types.Image, error) {
	var images []types.Image
	cursor, err := Images.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var image types.Image
		err := cursor.Decode(&image)
		if err != nil {
			log.Fatal(err)
		}

		images = append(images, image)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return images, nil
}

func GetOrdersByValue(filter primitive.M) ([]types.Order, error) {
	var orders []types.Order
	cursor, err := Orders.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var order types.Order
		err := cursor.Decode(&order)
		if err != nil {
			log.Fatal(err)
		}

		orders = append(orders, order)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return orders, nil
}

func GetOrderItemsByValue(filter primitive.M) ([]types.OrderItem, error) {
	var orderItems []types.OrderItem
	cursor, err := OrderItems.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var orderItem types.OrderItem
		err := cursor.Decode(&orderItem)
		if err != nil {
			log.Fatal(err)
		}

		orderItems = append(orderItems, orderItem)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return orderItems, nil
}

func GetProductsByValue(filter primitive.M) ([]types.FullProduct, error) {
	var products []types.FullProduct
	cursor, err := Products.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var product types.FullProduct
		err := cursor.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		var color types.Color
		var size types.Size
		var category types.FullCategory
		var image []types.Image

		GetColorById(&color, bson.M{"_id": product.ColorId})
		GetSizeById(&size, bson.M{"_id": product.SizeId})
		category = GetCategory(bson.M{"_id": product.CategoryId})
		image, _ = GetImagesByValue(bson.M{"productId": product.ID})
		product.Category = category
		product.Size = size
		product.Color = color
		product.Image = image

		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return products, nil
}

func GetProductsByValueData(filter primitive.M) []types.Product {
	var products []types.Product
	cursor, err := Products.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var size types.Product
		err := cursor.Decode(&size)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, size)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return products
}

func GetSizesByValue(filter primitive.M) ([]types.Size, error) {
	var sizes []types.Size
	cursor, err := Sizes.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var size types.Size
		err := cursor.Decode(&size)
		if err != nil {
			log.Fatal(err)
		}

		sizes = append(sizes, size)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return sizes, nil
}

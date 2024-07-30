package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FullStoreType struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
	Name       string             `json:"name" bson:"name"`
	UserId     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Billboards []Billboard        `json:"billboards" bson:"billboards"`
	Categories []Category         `json:"categories" bson:"categories"`
	Colors     []Color            `json:"colors" bson:"colors"`
	Products   []Product          `json:"products" bson:"products"`
	Sizes      []Size             `json:"sizes" bson:"sizes"`
}

type FullBillboardType struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
	Label      string             `json:"label" bson:"label"`
	StoreId    primitive.ObjectID `json:"storeId" bson:"storeId"`
	ImageUrl   string             `json:"imageUrl" bson:"imageUrl"`
	Categories []Category         `json:"categories" bson:"categories"`
}

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Name      string             `json:"name" bson:"name"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type TokenResponse struct {
	Token string             `json:"token"`
	Name  string             `json:"name" `
	Email string             `json:"email" `
	ID    primitive.ObjectID `json:"id"`
}

type Store struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Name      string             `json:"name" bson:"name"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id"`
}

type Billboard struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Label     string             `json:"label" bson:"label"`
	StoreId   primitive.ObjectID `json:"storeId" bson:"storeId"`
	ImageUrl  string             `json:"imageUrl" bson:"imageUrl"`
}

type Category struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
	StoreId     primitive.ObjectID `json:"storeId" bson:"storeId"`
	BillboardId primitive.ObjectID `json:"billboardId" bson:"billboardId"`
	Name        string             `json:"name" bson:"name"`
}

type FullCategory struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
	StoreId     primitive.ObjectID `json:"storeId" bson:"storeId"`
	BillboardId primitive.ObjectID `json:"billboardId" bson:"billboardId"`
	Name        string             `json:"name" bson:"name"`
	Billboard   Billboard          `json:"billboard" bson:"billboard"`
}

type Color struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Name      string             `json:"name" bson:"name"`
	StoreId   primitive.ObjectID `json:"storeId" bson:"storeId"`
	Value     string             `json:"value" bson:"value"`
}

type Image struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Url       string             `json:"url" bson:"url"`
	ProductId primitive.ObjectID `json:"productId" bson:"productId"`
}

type Order struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	StoreId   primitive.ObjectID `json:"storeId" bson:"storeId"`
	IsPaid    bool               `json:"isPaid" bson:"isPaid"`
	Phone     string             `json:"phone" bson:"phone"`
	Address   string             `json:"address" bson:"address"`
}

type OrderItem struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OrderId   primitive.ObjectID `json:"orderId" bson:"orderId"`
	ProductId primitive.ObjectID `json:"productId" bson:"productId"`
}

type Product struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
	Name       string             `json:"name" bson:"name"`
	StoreId    primitive.ObjectID `json:"storeId" bson:"storeId"`
	CategoryId primitive.ObjectID `json:"categoryId" bson:"categoryId"`
	Price      string             `json:"price" bson:"price"`
	IsFeatured bool               `json:"isFeatured" bson:"isFeatured"`
	IsArchived bool               `json:"isArchived" bson:"isArchived"`
	SizeId     primitive.ObjectID `json:"sizeId" bson:"sizeId"`
	ColorId    primitive.ObjectID `json:"colorId" bson:"colorId"`
}

type FullProduct struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
	Name       string             `json:"name" bson:"name"`
	StoreId    primitive.ObjectID `json:"storeId" bson:"storeId"`
	CategoryId primitive.ObjectID `json:"categoryId" bson:"categoryId"`
	Price      string             `json:"price" bson:"price"`
	IsFeatured bool               `json:"isFeatured" bson:"isFeatured"`
	IsArchived bool               `json:"isArchived" bson:"isArchived"`
	SizeId     primitive.ObjectID `json:"sizeId" bson:"sizeId"`
	ColorId    primitive.ObjectID `json:"colorId" bson:"colorId"`
	Category   FullCategory       `json:"category" bson:"category"`
	Size       Size               `json:"size" bson:"size"`
	Color      Color              `json:"color" bson:"color"`
	Image      []Image            `json:"image" bson:"image"`
}

type Size struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Name      string             `json:"name" bson:"name"`
	Value     string             `json:"value" bson:"value"`
	StoreId   primitive.ObjectID `json:"storeId" bson:"storeId"`
}

type ProductWithImageUrl struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
	Name       string             `json:"name" bson:"name"`
	StoreId    primitive.ObjectID `json:"storeId" bson:"storeId"`
	CategoryId primitive.ObjectID `json:"categoryId" bson:"categoryId"`
	Price      string             `json:"price" bson:"price"`
	IsFeatured bool               `json:"isFeatured" bson:"isFeatured"`
	IsArchived bool               `json:"isArchived" bson:"isArchived"`
	SizeId     primitive.ObjectID `json:"sizeId" bson:"sizeId"`
	ColorId    primitive.ObjectID `json:"colorId" bson:"colorId"`
	ImageUrl   []string           `json:"imageUrl" bson:"imageUrl"`
}

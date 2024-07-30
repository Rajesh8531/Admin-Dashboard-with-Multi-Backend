package types

import (
	"time"
)

type FullStoreType struct {
	ID         string      `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt" bson:"updatedAt"`
	Name       string      `json:"name" bson:"name"`
	UserId     string      `json:"user_id" bson:"user_id"`
	Billboards []Billboard `json:"billboards" bson:"billboards"`
	Categories []Category  `json:"categories" bson:"categories"`
	Colors     []Color     `json:"colors" bson:"colors"`
	Products   []Product   `json:"products" bson:"products"`
	Sizes      []Size      `json:"sizes" bson:"sizes"`
}

type FullBillboardType struct {
	ID         string     `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt" bson:"updatedAt"`
	Label      string     `json:"label" bson:"label"`
	StoreId    string     `json:"storeId" bson:"storeId"`
	ImageUrl   string     `json:"imageUrl" bson:"imageUrl"`
	Categories []Category `json:"categories" bson:"categories"`
}

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Store struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	Name      string    `json:"name" bson:"name"`
	UserId    string    `json:"user_id" bson:"user_id"`
}

type Billboard struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	Label     string    `json:"label" bson:"label"`
	StoreId   string    `json:"storeId" bson:"storeId"`
	ImageUrl  string    `json:"imageUrl" bson:"imageUrl"`
}

type Category struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
	StoreId     string    `json:"storeId" bson:"storeId"`
	BillboardId string    `json:"billboardId" bson:"billboardId"`
	Name        string    `json:"name" bson:"name"`
}

type FullCategory struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
	StoreId     string    `json:"storeId" bson:"storeId"`
	BillboardId string    `json:"billboardId" bson:"billboardId"`
	Name        string    `json:"name" bson:"name"`
	Billboard   Billboard `json:"billboard" bson:"billboard"`
	Products    []Product `json:"products" bson:"product"`
}

type Color struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	Name      string    `json:"name" bson:"name"`
	StoreId   string    `json:"storeId" bson:"storeId"`
	Value     string    `json:"value" bson:"value"`
}

type Image struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	Url       string    `json:"url" bson:"url"`
	ProductId string    `json:"productId" bson:"productId"`
}

type Order struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	StoreId   string    `json:"storeId" bson:"storeId"`
	IsPaid    bool      `json:"isPaid" bson:"isPaid"`
	Phone     string    `json:"phone" bson:"phone"`
	Address   string    `json:"address" bson:"address"`
}

type FullOrder struct {
	ID        string        `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
	StoreId   string        `json:"storeId" bson:"storeId"`
	IsPaid    bool          `json:"isPaid" bson:"isPaid"`
	Phone     string        `json:"phone" bson:"phone"`
	Address   string        `json:"address" bson:"address"`
	OrderItem FullOrderItem `json:"orderItem" bson:"orderItem"`
}

type OrderItem struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	OrderId   string `json:"orderId" bson:"orderId"`
	ProductId string `json:"productId" bson:"productId"`
}

type FullOrderItem struct {
	ID        string  `json:"id,omitempty" bson:"_id,omitempty"`
	OrderId   string  `json:"orderId" bson:"orderId"`
	ProductId string  `json:"productId" bson:"productId"`
	Product   FullProduct `json:"product" bson:"product"`
}

type Product struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	Name       string    `json:"name" bson:"name"`
	StoreId    string    `json:"storeId" bson:"storeId"`
	CategoryId string    `json:"categoryId" bson:"categoryId"`
	Price      string    `json:"price" bson:"price"`
	IsFeatured bool      `json:"isFeatured" bson:"isFeatured"`
	IsArchived bool      `json:"isArchived" bson:"isArchived"`
	SizeId     string    `json:"sizeId" bson:"sizeId"`
	ColorId    string    `json:"colorId" bson:"colorId"`
}

type FullProduct struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	Name       string    `json:"name" bson:"name"`
	StoreId    string    `json:"storeId" bson:"storeId"`
	CategoryId string    `json:"categoryId" bson:"categoryId"`
	Price      string    `json:"price" bson:"price"`
	IsFeatured bool      `json:"isFeatured" bson:"isFeatured"`
	IsArchived bool      `json:"isArchived" bson:"isArchived"`
	SizeId     string    `json:"sizeId" bson:"sizeId"`
	ColorId    string    `json:"colorId" bson:"colorId"`
	Category   Category  `json:"category" bson:"category"`
	Size       Size      `json:"size" bson:"size"`
	Color      Color     `json:"color" bson:"color"`
	Image      []Image   `json:"image" bson:"image"`
}

type Size struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	Name      string    `json:"name" bson:"name"`
	Value     string    `json:"value" bson:"value"`
	StoreId   string    `json:"storeId" bson:"storeId"`
}

type ProductWithImageUrl struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	Name       string    `json:"name" bson:"name"`
	StoreId    string    `json:"storeId" bson:"storeId"`
	CategoryId string    `json:"categoryId" bson:"categoryId"`
	Price      string    `json:"price" bson:"price"`
	IsFeatured bool      `json:"isFeatured" bson:"isFeatured"`
	IsArchived bool      `json:"isArchived" bson:"isArchived"`
	SizeId     string    `json:"sizeId" bson:"sizeId"`
	ColorId    string    `json:"colorId" bson:"colorId"`
	ImageUrl   []string  `json:"imageUrl" bson:"imageUrl"`
}

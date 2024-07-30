export type Store = {
    id: string;
    name: string;
    userId: string;
    createdAt: Date;
    updatedAt: Date;
}

export type User = {
    id: string;
    name: string | null;
    email: string | null;
    password: string | null;
    createdAt: Date;
    updatedAt: Date;
}

export type BillBoard = {
    id: string;
    label: string;
    storeId: string;
    createdAt: Date;
    updatedAt: Date;
    imageUrl: string;
}

export type Category = {
    id: string;
    storeId: string;
    billboardId: string;
    name: string;
    createdAt: Date;
    updatedAt: Date;
}

export type Color = {
    id: string;
    storeId: string;
    name: string;
    value: string;
    createdAt: Date;
    updatedAt: Date;
}

export type Image = {
    id: string;
    productId: string;
    url: string;
    createdAt: Date;
    updatedAt: Date;
}

export type Order = {
    id: string;
    storeId: string;
    isPaid: boolean;
    phone: string;
    address: string;
    createdAt: Date;
    updatedAt: Date;
}

export type OrderItem = {
    id: string;
    orderId: string;
    productId: string;
}

export type FullOrderItem = OrderItem & {product : Product}

export type Product = {
    id: string;
    storeId: string;
    categoryId: string;
    name: string;
    price: number;
    isFeatured: boolean;
    isArchived: boolean;
    sizeId: string;
    colorId: string;
    createdAt: Date;
    updatedAt: Date;
}

export type FullOrderType = Order & {
    orderItems : FullOrderItem[]
}

export type Size = {
    id: string;
    storeId: string;
    name: string;
    value: string;
    createdAt: Date;
    updatedAt: Date;
}
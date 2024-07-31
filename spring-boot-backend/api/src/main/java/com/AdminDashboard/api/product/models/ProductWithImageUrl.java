package com.AdminDashboard.api.product.models;

import java.time.Instant;
import java.util.Arrays;

public class ProductWithImageUrl {
    private String id;
    private String name;
    private String price;
    private boolean isFeatured = false;
    private boolean isArchived = false;
    private Instant createdAt;
    private Instant updatedAt;
    private String[] imageUrl;
    private String storeId;
    private String categoryId;
    private String sizeId;
    private String colorId;


    @Override
    public String toString() {
        return "ProductWithImageUrl{" +
                "id='" + id + '\'' +
                ", name='" + name + '\'' +
                ", price='" + price + '\'' +
                ", isFeatured=" + isFeatured +
                ", isArchived=" + isArchived +
                ", createdAt=" + createdAt +
                ", updatedAt=" + updatedAt +
                ", imageUrl=" + Arrays.toString(imageUrl) +
                ", storeId='" + storeId + '\'' +
                ", categoryId='" + categoryId + '\'' +
                ", sizeId='" + sizeId + '\'' +
                ", colorId='" + colorId + '\'' +
                '}';
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getPrice() {
        return price;
    }

    public void setPrice(String price) {
        this.price = price;
    }

    public boolean isFeatured() {
        return isFeatured;
    }

    public void setisFeatured(boolean featured) {
        isFeatured = featured;
    }

    public boolean isArchived() {
        return isArchived;
    }

    public void setisArchived(boolean archived) {
        isArchived = archived;
    }

    public Instant getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Instant createdAt) {
        this.createdAt = createdAt;
    }

    public Instant getUpdatedAt() {
        return updatedAt;
    }

    public void setUpdatedAt(Instant updatedAt) {
        this.updatedAt = updatedAt;
    }

    public String[] getImageUrl() {
        return imageUrl;
    }

    public void setImageUrl(String[] imageUrl) {
        this.imageUrl = imageUrl;
    }

    public String getCategoryId() {
        return categoryId;
    }

    public void setCategoryId(String categoryId) {
        this.categoryId = categoryId;
    }

    public String getStoreId() {
        return storeId;
    }

    public void setStoreId(String storeId) {
        this.storeId = storeId;
    }

    public String getSizeId() {
        return sizeId;
    }

    public void setSizeId(String sizeId) {
        this.sizeId = sizeId;
    }

    public String getColorId() {
        return colorId;
    }

    public void setColorId(String colorId) {
        this.colorId = colorId;
    }
}

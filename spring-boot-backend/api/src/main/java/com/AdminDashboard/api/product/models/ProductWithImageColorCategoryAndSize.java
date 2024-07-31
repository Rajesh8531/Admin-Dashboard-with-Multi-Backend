package com.AdminDashboard.api.product.models;

import org.springframework.stereotype.Component;

import java.time.Instant;
import java.util.ArrayList;
import java.util.List;

@Component
public class ProductWithImageColorCategoryAndSize {
    private String id;
    private String storeId;
    private String categoryId;
    private String sizeId;
    private String name;
    private String price;

    @Override
    public String toString() {
        return "ProductWithImageColorAndSize{" +
                "id='" + id + '\'' +
                ", storeId='" + storeId + '\'' +
                ", categoryId='" + categoryId + '\'' +
                ", sizeId='" + sizeId + '\'' +
                ", name='" + name + '\'' +
                ", price='" + price + '\'' +
                ", isArchived=" + isArchived +
                ", isFeatured=" + isFeatured +
                ", colorId='" + colorId + '\'' +
                ", createdAt=" + createdAt +
                ", updatedAt=" + updatedAt +
                ", image=" + image +
                '}';
    }

    private boolean isArchived;
    private boolean isFeatured;
    private String colorId;
    private Instant createdAt;
    private Instant updatedAt;
    private List<Image> image = new ArrayList<>();
    private UniqueSize size;
    private UniqueColor color;
    private UniqueCategory category;

    public UniqueCategory getCategory() {
        return category;
    }

    public void setCategory(UniqueCategory category) {
        this.category = category;
    }

    public UniqueSize getSize() {
        return size;
    }

    public void setSize(UniqueSize size) {
        this.size = size;
    }

    public UniqueColor getColor() {
        return color;
    }

    public void setColor(UniqueColor color) {
        this.color = color;
    }

    public List<Image> getImage() {
        return image;
    }

    public void setImage(List<Image> image) {
        this.image = image;
    }

    public Instant getUpdatedAt() {
        return updatedAt;
    }

    public void setUpdatedAt(Instant updatedAt) {
        this.updatedAt = updatedAt;
    }

    public Instant getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Instant createdAt) {
        this.createdAt = createdAt;
    }

    public String getColorId() {
        return colorId;
    }

    public void setColorId(String colorId) {
        this.colorId = colorId;
    }

    public boolean isFeatured() {
        return isFeatured;
    }

    public void setFeatured(boolean featured) {
        isFeatured = featured;
    }

    public boolean isArchived() {
        return isArchived;
    }

    public void setArchived(boolean archived) {
        isArchived = archived;
    }

    public String getPrice() {
        return price;
    }

    public void setPrice(String price) {
        this.price = price;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getSizeId() {
        return sizeId;
    }

    public void setSizeId(String sizeId) {
        this.sizeId = sizeId;
    }

    public String getCategoryId() {
        return categoryId;
    }

    public void setCategoryId(String categoryId) {
        this.categoryId = categoryId;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getStoreId() {
        return storeId;
    }

    public void setStoreId(String storeId) {
        this.storeId = storeId;
    }
}

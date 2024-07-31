package com.AdminDashboard.api.product.models;

import com.AdminDashboard.api.category.Category;
import com.AdminDashboard.api.color.Color;
import com.AdminDashboard.api.order.models.OrderItem;
import com.AdminDashboard.api.size.Size;
import com.AdminDashboard.api.store.Store;
import com.fasterxml.jackson.annotation.JsonIgnore;
import jakarta.persistence.*;

import java.time.Instant;
import java.util.ArrayList;
import java.util.List;

@Entity
public class Product {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    private String name;
    private String price;
    @Column(name = "isFeatured")
    private boolean isFeatured = false;
    @Column(name = "isArchived")
    private boolean isArchived = false;
    private Instant createdAt;
    private Instant updatedAt;

    @JsonIgnore
    @ManyToOne
    @JoinColumn(name = "storeId")
    private Store store;
    @Column(name = "storeId",updatable = false,insertable = false)
    private String storeId;

    @JsonIgnore
    @ManyToOne
    @JoinColumn(name = "categoryId")
    private Category category;
    @Column(name = "categoryId",updatable = false,insertable = false)
    private String categoryId;

    @JsonIgnore
    @ManyToOne
    @JoinColumn(name = "sizeId")
    private Size size;
    @Column(name = "sizeId",updatable = false,insertable = false)
    private String sizeId;

    @JsonIgnore
    @ManyToOne
    @JoinColumn(name = "colorId")
    private Color color;
    @Column(name = "colorId",updatable = false,insertable = false)
    private String colorId;

    @OneToMany(mappedBy = "product")
    private List<Image> image = new ArrayList<>();

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getId() {
        return id;
    }

    @Override
    public String toString() {
        return "Product{" +
                "id='" + id + '\'' +
                ", name='" + name + '\'' +
                ", price='" + price + '\'' +
                ", isFeatured=" + isFeatured +
                ", isArchived=" + isArchived +
                ", createdAt=" + createdAt +
                ", updatedAt=" + updatedAt +
                ", store=" + store +
                ", storeId='" + storeId + '\'' +
                ", category=" + category +
                ", categoryId='" + categoryId + '\'' +
                ", size=" + size +
                ", sizeId='" + sizeId + '\'' +
                ", color=" + color +
                ", colorId='" + colorId + '\'' +
                ", image=" + image +
                ", orderItems=" + orderItems +
                '}';
    }

    public void setId(String id) {
        this.id = id;
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

    public Store getStore() {
        return store;
    }

    public void setStore(Store store) {
        this.store = store;
    }

    public String getStoreId() {
        return storeId;
    }

    public void setStoreId(String storeId) {
        this.storeId = storeId;
    }

    public Category getCategory() {
        return category;
    }

    public void setCategory(Category category) {
        this.category = category;
    }

    public String getCategoryId() {
        return categoryId;
    }

    public void setCategoryId(String categoryId) {
        this.categoryId = categoryId;
    }

    public Size getSize() {
        return size;
    }

    public void setSize(Size size) {
        this.size = size;
    }

    public String getSizeId() {
        return sizeId;
    }

    public void setSizeId(String sizeId) {
        this.sizeId = sizeId;
    }

    public Color getColor() {
        return color;
    }

    public void setColor(Color color) {
        this.color = color;
    }

    public String getColorId() {
        return colorId;
    }

    public void setColorId(String colorId) {
        this.colorId = colorId;
    }

    public List<Image> getImage() {
        return image;
    }

    public void setImage(List<Image> image) {
        this.image = image;
    }

    public List<OrderItem> getOrderItems() {
        return orderItems;
    }

    public void setOrderItems(List<OrderItem> orderItems) {
        this.orderItems = orderItems;
    }

    @OneToMany(mappedBy = "product")
    private List<OrderItem> orderItems;
}

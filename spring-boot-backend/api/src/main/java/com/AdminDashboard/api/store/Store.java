package com.AdminDashboard.api.store;

import com.AdminDashboard.api.billboard.Billboard;
import com.AdminDashboard.api.category.Category;
import com.AdminDashboard.api.color.Color;
import com.AdminDashboard.api.product.models.Image;
import com.AdminDashboard.api.product.models.Product;
import com.AdminDashboard.api.size.Size;
import jakarta.persistence.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import java.time.Instant;
import java.util.List;

@Entity
public class Store {
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    private String name;
    private String userId;
    @CreatedDate
    @Column(name = "createdAt")
    private Instant createdAt;

    public List<Billboard> getBillboards() {
        return billboards;
    }

    public void setBillboards(List<Billboard> billboards) {
        this.billboards = billboards;
    }

    @LastModifiedDate
    @Column(name = "updatedAt")
    private Instant updatedAt;

    @OneToMany(mappedBy = "store")
    private List<Billboard> billboards;

    @OneToMany(mappedBy = "store",fetch = FetchType.LAZY)
    private List<Category> categories;

    @OneToMany(mappedBy = "store",fetch = FetchType.LAZY)
    private List<Size> sizes;

    @OneToMany(mappedBy = "store",fetch = FetchType.LAZY)
    private List<Color> colors;

    @OneToMany(mappedBy = "store")
    private List<Product> products;

    public List<Product> getProducts() {
        return products;
    }

    public void setProducts(List<Product> products) {
        this.products = products;
    }

    public List<Color> getColors() {
        return colors;
    }

    public void setColors(List<Color> colors) {
        this.colors = colors;
    }

    public List<Size> getSizes() {
        return sizes;
    }

    public void setSizes(List<Size> sizes) {
        this.sizes = sizes;
    }

    public List<Category> getCategories() {
        return categories;
    }

    public void setCategories(List<Category> categories) {
        this.categories = categories;
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

    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
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
}

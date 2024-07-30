package com.AdminDashboard.api.category;

import com.AdminDashboard.api.billboard.Billboard;
import com.AdminDashboard.api.product.models.Product;
import com.AdminDashboard.api.store.Store;
import com.fasterxml.jackson.annotation.JsonIgnore;
import jakarta.persistence.*;
import org.springframework.data.annotation.CreatedDate;

import java.time.Instant;
import java.util.List;

@Entity
public class Category {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    private String name;

    @CreatedDate
    @Column(name = "createdAt")
    private Instant createdAt;

    @CreatedDate
    @Column(name = "updatedAt")
    private Instant updatedAt;

    @OneToMany(mappedBy = "category")
    private List<Product> products;

    public List<Product> getProducts() {
        return products;
    }

    public void setProducts(List<Product> products) {
        this.products = products;
    }

    public Store getStore() {
        return store;
    }

    public void setStore(Store store) {
        this.store = store;
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

    public String getStoreId() {
        return storeId;
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

    @JsonIgnore
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "storeId")
    private Store store;

    @Column(name = "storeId",updatable = false,insertable = false)
    private String storeId;

    @JsonIgnore
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "billboardId")
    private Billboard billboard;

    @Column(name = "billboardId",updatable = false,insertable = false)
    private String billboardId;

    public Billboard getBillboard() {
        return billboard;
    }

    @Override
    public String toString() {
        return "Category{" +
                "id='" + id + '\'' +
                ", name='" + name + '\'' +
                ", createdAt=" + createdAt +
                ", updatedAt=" + updatedAt +
                ", store=" + store +
                ", storeId='" + storeId + '\'' +
                ", billboard=" + billboard +
                ", billboardId='" + billboardId + '\'' +
                '}';
    }

    public void setBillboard(Billboard billboard) {
        this.billboard = billboard;
    }

    public String getBillboardId() {
        return billboardId;
    }

    public void setBillboardId(String billboardId) {
        this.billboardId = billboardId;
    }

    public void setStoreId(String storeId) {
        this.storeId = storeId;
    }
}

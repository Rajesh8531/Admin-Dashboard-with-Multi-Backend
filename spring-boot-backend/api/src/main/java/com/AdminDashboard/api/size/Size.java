package com.AdminDashboard.api.size;

import com.AdminDashboard.api.billboard.Billboard;
import com.AdminDashboard.api.product.models.Product;
import com.AdminDashboard.api.store.Store;
import com.fasterxml.jackson.annotation.JsonIgnore;
import jakarta.persistence.*;
import org.springframework.data.annotation.CreatedDate;

import java.time.Instant;
import java.util.List;

@Entity
public class Size {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    private String name;
    private String value;

    @OneToMany(mappedBy = "size")
    private List<Product> products;

    public String getValue() {
        return value;
    }

    public List<Product> getProducts() {
        return products;
    }

    public void setProducts(List<Product> products) {
        this.products = products;
    }

    public void setValue(String value) {
        this.value = value;
    }

    @CreatedDate
    @Column(name = "createdAt")
    private Instant createdAt;

    @CreatedDate
    @Column(name = "updatedAt")
    private Instant updatedAt;

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



    @Override
    public String toString() {
        return "Category{" +
                "id='" + id + '\'' +
                ", name='" + name + '\'' +
                ", createdAt=" + createdAt +
                ", updatedAt=" + updatedAt +
                ", store=" + store +
                ", storeId='" + storeId + '\'' +
                 +
                '}';
    }



    public void setStoreId(String storeId) {
        this.storeId = storeId;
    }
}

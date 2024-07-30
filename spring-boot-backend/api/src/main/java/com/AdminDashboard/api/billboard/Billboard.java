package com.AdminDashboard.api.billboard;

import com.AdminDashboard.api.category.Category;
import com.AdminDashboard.api.store.Store;
import com.fasterxml.jackson.annotation.JsonIgnore;
import jakarta.persistence.*;
import org.springframework.data.annotation.CreatedDate;

import java.time.Instant;
import java.util.List;

@Entity
public class Billboard {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    @Override
    public String toString() {
        return "Billboard{" +
                "id='" + id + '\'' +
                ", label='" + label + '\'' +
                ", imageUrl='" + imageUrl + '\'' +
                ", categories=" + categories +
                ", createdAt=" + createdAt +
                ", updatedAt=" + updatedAt +
                ", store=" + store +
                ", storeId='" + storeId + '\'' +
                '}';
    }

    private String label;
    private String imageUrl;

    @OneToMany(mappedBy = "billboard",fetch=FetchType.LAZY)
    private List<Category> categories;

    public void setStoreId(String storeId) {
        this.storeId = storeId;
    }

    public List<Category> getCategories() {
        return categories;
    }

    public void setCategories(List<Category> categories) {
        this.categories = categories;
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

    public String getLabel() {
        return label;
    }

    public void setLabel(String label) {
        this.label = label;
    }

    public String getStoreId() {
        return storeId;
    }

//    public void setStoreId(String storeId) {
//        this.storeId = storeId;
//    }

    public String getImageUrl() {
        return imageUrl;
    }

    public void setImageUrl(String imageUrl) {
        this.imageUrl = imageUrl;
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

}

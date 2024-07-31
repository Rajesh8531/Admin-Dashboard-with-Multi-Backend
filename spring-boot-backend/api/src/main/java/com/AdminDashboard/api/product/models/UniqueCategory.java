package com.AdminDashboard.api.product.models;

import java.time.Instant;

public class UniqueCategory {
    private String id;
    private String storeId;
    private String name;
    private Instant createdAt;
    private Instant updatedAt;
    private String billboardId;

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

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
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

    public String getBillboardId() {
        return billboardId;
    }

    public void setBillboardId(String billboardId) {
        this.billboardId = billboardId;
    }
}

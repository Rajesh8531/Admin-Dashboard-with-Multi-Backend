package com.AdminDashboard.api.store;

import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface StoreRepository extends JpaRepository<Store,String> {
    List<Store> findAllByUserId(String userId);
}

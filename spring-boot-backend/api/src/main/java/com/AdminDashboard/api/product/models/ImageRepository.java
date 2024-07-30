package com.AdminDashboard.api.product.models;

import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;


public interface ImageRepository extends JpaRepository<Image,String> {
    List<Image> getImagesByProductId(String productId);

}

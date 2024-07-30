package com.AdminDashboard.api.product;

import com.AdminDashboard.api.category.Category;
import com.AdminDashboard.api.product.models.Product;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;

public interface ProductRepository extends JpaRepository<Product,String> {
    List<Product> getProductsByStoreId(String storeId);

    @Query(value = "SELECT * FROM product p WHERE p.store_id = :storeId AND p.id = :productId", nativeQuery = true)
    Product findProductByStoreIdAndProductId(@Param("storeId") String storeId, @Param("productId") String productId);
}

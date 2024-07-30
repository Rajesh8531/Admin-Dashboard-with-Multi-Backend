package com.AdminDashboard.api.category;

import com.AdminDashboard.api.category.Category;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;

public interface CategoryRepository extends JpaRepository<Category,String> {
    List<Category> getCategoriesByStoreId(String storeId);

    @Query(value = "SELECT * FROM category c WHERE c.store_id = :storeId AND c.id = :categoryId", nativeQuery = true)
    Category findCategoryByStoreIdAndCategoryId(@Param("storeId") String storeId,@Param("categoryId") String categoryId);
}

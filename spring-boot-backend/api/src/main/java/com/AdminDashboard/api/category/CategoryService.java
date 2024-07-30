package com.AdminDashboard.api.category;

import com.AdminDashboard.api.category.Category;

import java.util.List;

public interface CategoryService {
    List<CategoryWithBillboardDTO> getCategories(String storeId);
    boolean createCategory(String storeId, Category category);
    CategoryWithBillboardDTO getCategory(String storeId, String categoryId);
    boolean updateCategory(String storeId, String categoryId, Category category);
    boolean deleteCategory(String storeId, String categoryId);
}

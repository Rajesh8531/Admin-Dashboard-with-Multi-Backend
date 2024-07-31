package com.AdminDashboard.api.product;

import com.AdminDashboard.api.product.models.Product;
import com.AdminDashboard.api.product.models.ProductWithImageColorCategoryAndSize;
import com.AdminDashboard.api.product.models.ProductWithImageUrl;

import java.util.List;

public interface ProductService {
    List<ProductWithImageColorCategoryAndSize> getProducts(String storeId);
    boolean createProduct(String storeId, ProductWithImageUrl productWithImageUrl);
    Product getProduct(String storeId,String productId);
    boolean updateProduct(String storeId,String productId,ProductWithImageUrl productWithImageUrl);
    boolean deleteProduct(String storeId,String productId);
}

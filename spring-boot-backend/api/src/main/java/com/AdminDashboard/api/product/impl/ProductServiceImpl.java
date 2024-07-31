package com.AdminDashboard.api.product.impl;

import com.AdminDashboard.api.category.Category;
import com.AdminDashboard.api.category.CategoryRepository;
import com.AdminDashboard.api.color.Color;
import com.AdminDashboard.api.color.ColorRepository;
import com.AdminDashboard.api.product.ProductRepository;
import com.AdminDashboard.api.product.ProductService;
import com.AdminDashboard.api.product.models.*;
import com.AdminDashboard.api.size.Size;
import com.AdminDashboard.api.size.SizeRepository;
import com.AdminDashboard.api.store.Store;
import com.AdminDashboard.api.store.StoreRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.Instant;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

@Service
public class ProductServiceImpl implements ProductService {

    ProductRepository productRepository;
    StoreRepository storeRepository;
    CategoryRepository categoryRepository;
    ColorRepository colorRepository;
    SizeRepository sizeRepository;
    ImageRepository imageRepository;

    public ProductServiceImpl(ProductRepository productRepository,
                              StoreRepository storeRepository,
                              ColorRepository colorRepository,
                              CategoryRepository categoryRepository,
                              SizeRepository sizeRepository,
                              ImageRepository imageRepository) {
        this.productRepository = productRepository;
        this.storeRepository = storeRepository;
        this.colorRepository = colorRepository;
        this.categoryRepository = categoryRepository;
        this.sizeRepository = sizeRepository;
        this.imageRepository = imageRepository;
    }

    @Override
    public List<ProductWithImageColorCategoryAndSize> getProducts(String storeId) {
        List<Product> products = productRepository.getProductsByStoreId(storeId);
        return products.stream().map(this::getFullProduct).collect(Collectors.toList());
    }

    @Override
    public boolean createProduct(String storeId, ProductWithImageUrl productWithImageUrl) {

        Store store = storeRepository.findById(storeId).orElse(null);
        Category category = categoryRepository.findById(productWithImageUrl.getCategoryId()).orElse(null);
        Color color = colorRepository.findById(productWithImageUrl.getColorId()).orElse(null);
        Size size = sizeRepository.findById(productWithImageUrl.getSizeId()).orElse(null);
        if(store == null || category == null || color == null || size == null){
            return false;
        }

        Product product = this.productWithImageUrlToProduct(productWithImageUrl);
        product.setSize(size);
        product.setColor(color);
        product.setStore(store);
        product.setCategory(category);
        product.setCreatedAt(Instant.now());
        product.setUpdatedAt(Instant.now());

        store.getProducts().add(product);
        color.getProducts().add(product);
        size.getProducts().add(product);
        category.getProducts().add(product);

        storeRepository.save(store);
        colorRepository.save(color);
        sizeRepository.save(size);
        categoryRepository.save(category);

        Product newProduct = productRepository.save(product);
        String[] imageUrl = productWithImageUrl.getImageUrl();

        for (String url : imageUrl){
            Image image = new Image();
            image.setCreatedAt(Instant.now());
            image.setUpdatedAt(Instant.now());
            image.setUrl(url);
            image.setProduct(newProduct);
            image.setProductId(newProduct.getId());

            newProduct.getImage().add(image);
            imageRepository.save(image);
        }
        return true;
    }

    @Override
    public Product getProduct(String storeId, String productId) {
        return productRepository.findProductByStoreIdAndProductId(storeId,productId);
    }

    @Override
    public boolean updateProduct(String storeId, String productId, ProductWithImageUrl productWithImageUrl) {
        Store store = storeRepository.findById(storeId).orElse(null);
        Category category = categoryRepository.findById(productWithImageUrl.getCategoryId()).orElse(null);
        Color color = colorRepository.findById(productWithImageUrl.getColorId()).orElse(null);
        Size size = sizeRepository.findById(productWithImageUrl.getSizeId()).orElse(null);

        if(store == null || category == null || color == null || size == null){
            return false;
        }

        List<Image> images = imageRepository.getImagesByProductId(productId);
        for (Image image : images){
            imageRepository.delete(image);
        }

        Product product = this.productWithImageUrlToProduct(productWithImageUrl);
        product.setisFeatured(productWithImageUrl.isFeatured());
        product.setisArchived(productWithImageUrl.isArchived());
        product.setId(productId);
        product.setStoreId(storeId);
        product.setSize(size);
        product.setColor(color);
        product.setStore(store);
        product.setCategory(category);
        product.setCreatedAt(Instant.now());
        product.setUpdatedAt(Instant.now());
        store.getProducts().add(product);
        color.getProducts().add(product);
        size.getProducts().add(product);
        category.getProducts().add(product);
        productRepository.save(product);

        String[] imageUrl = productWithImageUrl.getImageUrl();

        for (String url : imageUrl){
            Image image = new Image();
            image.setCreatedAt(Instant.now());
            image.setUpdatedAt(Instant.now());
            image.setUrl(url);
            image.setProduct(product);
            image.setProductId(product.getId());

            product.getImage().add(image);
            imageRepository.save(image);
        }
        return true;
    }


    @Override
    public boolean deleteProduct(String storeId, String productId) {
        Store store = storeRepository.findById(storeId).orElse(null);
        Product existingProduct = productRepository.findProductByStoreIdAndProductId(storeId,productId);
        Category category = categoryRepository.findById(existingProduct.getCategoryId()).orElse(null);
        Color color = colorRepository.findById(existingProduct.getColorId()).orElse(null);
        Size size = sizeRepository.findById(existingProduct.getSizeId()).orElse(null);

        if(store == null || category == null || color == null || size == null){
            return false;
        }

        List<Image> images = imageRepository.getImagesByProductId(productId);
        for (Image image : images){
            imageRepository.delete(image);
        }

        category.getProducts().remove(existingProduct);
        color.getProducts().remove(existingProduct);
        size.getProducts().remove(existingProduct);
        store.getProducts().remove(existingProduct);
        productRepository.delete(existingProduct);

        return true;
    }

    private Product productWithImageUrlToProduct(ProductWithImageUrl productWithImageUrl){
        Product product = new Product();
        product.setId(productWithImageUrl.getId());
        product.setCategoryId(productWithImageUrl.getCategoryId());
        product.setStoreId(productWithImageUrl.getStoreId());
        product.setColorId(productWithImageUrl.getColorId());
        product.setSizeId(productWithImageUrl.getSizeId());
        product.setName(productWithImageUrl.getName());
        product.setPrice(productWithImageUrl.getPrice());
        product.setisArchived(productWithImageUrl.isArchived());
        product.setisFeatured(productWithImageUrl.isFeatured());
        return product;
    }

    private ProductWithImageColorCategoryAndSize getFullProduct(Product product){
        ProductWithImageColorCategoryAndSize fullProduct = new ProductWithImageColorCategoryAndSize();
        fullProduct.setArchived(product.isArchived());
        fullProduct.setFeatured(product.isFeatured());
        fullProduct.setCategoryId(product.getCategoryId());
        fullProduct.setStoreId(product.getStoreId());
        fullProduct.setName(product.getName());
        fullProduct.setPrice(product.getPrice());
        fullProduct.setSizeId(product.getSizeId());
        fullProduct.setColorId(product.getColorId());
        fullProduct.setUpdatedAt(product.getUpdatedAt());
        fullProduct.setCreatedAt(product.getCreatedAt());
        fullProduct.setId(product.getId());
        fullProduct.setColor(colorToUniqueColor(colorRepository.findById(product.getColorId()).orElse(null)));
        fullProduct.setSize(sizeToUniqueSize(sizeRepository.findById(product.getSizeId()).orElse(null)));
        fullProduct.setCategory(categoryToUniqueCategory(categoryRepository.findById(product.getCategoryId()).orElse(null)));
        return fullProduct;
    }

    private UniqueColor colorToUniqueColor(Color color){
        UniqueColor uniqueColor = new UniqueColor();
        uniqueColor.setCreatedAt(color.getCreatedAt());
        uniqueColor.setUpdatedAt(color.getUpdatedAt());
        uniqueColor.setId(color.getId());
        uniqueColor.setName(color.getName());
        uniqueColor.setValue(color.getValue());
        return uniqueColor;
    }

    private UniqueSize sizeToUniqueSize(Size size){
        UniqueSize uniqueSize = new UniqueSize();
        uniqueSize.setCreatedAt(size.getCreatedAt());
        uniqueSize.setUpdatedAt(size.getUpdatedAt());
        uniqueSize.setId(size.getId());
        uniqueSize.setName(size.getName());
        uniqueSize.setValue(size.getValue());
        return uniqueSize;
    }

    private UniqueCategory categoryToUniqueCategory(Category category){
        UniqueCategory uniqueCategory = new UniqueCategory();
        uniqueCategory.setCreatedAt(category.getCreatedAt());
        uniqueCategory.setUpdatedAt(category.getUpdatedAt());
        uniqueCategory.setId(category.getId());
        uniqueCategory.setName(category.getName());
        uniqueCategory.setStoreId(category.getStoreId());
        uniqueCategory.setBillboardId(category.getBillboardId());
        return uniqueCategory;
    }
}


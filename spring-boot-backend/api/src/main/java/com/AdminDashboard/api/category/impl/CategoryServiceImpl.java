package com.AdminDashboard.api.category.impl;

import com.AdminDashboard.api.billboard.Billboard;
import com.AdminDashboard.api.billboard.BillboardRepository;
import com.AdminDashboard.api.category.Category;
import com.AdminDashboard.api.category.CategoryRepository;
import com.AdminDashboard.api.category.CategoryService;
import com.AdminDashboard.api.category.CategoryWithBillboardDTO;
import com.AdminDashboard.api.store.Store;
import com.AdminDashboard.api.store.StoreRepository;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.time.Instant;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

@Service
public class CategoryServiceImpl implements CategoryService {

    BillboardRepository billboardRepository;
    StoreRepository storeRepository;
    CategoryRepository categoryRepository;

    public CategoryServiceImpl(BillboardRepository billboardRepository, StoreRepository storeRepository,CategoryRepository categoryRepository) {
        this.billboardRepository = billboardRepository;
        this.storeRepository = storeRepository;
        this.categoryRepository = categoryRepository;
    }

    @Override
    public List<CategoryWithBillboardDTO> getCategories(String storeId) {
        return categoryRepository.getCategoriesByStoreId(storeId).stream().map(this::getCategoryWithBillboardDTO).collect(Collectors.toList());
    }

    @Override
    public boolean createCategory(String storeId, Category category) {
        Store store = storeRepository.findById(storeId).orElse(null);
        Billboard billboard = billboardRepository.findById(category.getBillboardId()).orElse(null);
        if(store == null || billboard == null){
            return false;
        }
        category.setCreatedAt(Instant.now());
        category.setUpdatedAt(Instant.now());
        category.setStore(store);
        category.setBillboard(billboard);

        store.getCategories().add(category);
        storeRepository.save(store);

        billboard.getCategories().add(category);
        billboardRepository.save(billboard);

        categoryRepository.save(category);
        return true;
    }

    @Override
    public CategoryWithBillboardDTO getCategory(String storeId, String categoryId) {
        return this.getCategoryWithBillboardDTO(categoryRepository.findCategoryByStoreIdAndCategoryId(storeId,categoryId));
    }

    @Override
    public boolean updateCategory(String storeId, String categoryId, Category category) {
        Store store = storeRepository.findById(storeId).orElse(null);
        Billboard billboard = billboardRepository.findById(category.getBillboardId()).orElse(null);
        if(store == null || billboard == null){
            return false;
        }
        System.out.println(category.getBillboardId());
        Category existingCategory = categoryRepository.findCategoryByStoreIdAndCategoryId(storeId,categoryId);
        existingCategory.setName(category.getName());
        existingCategory.setBillboard(billboard);
        existingCategory.setBillboardId(category.getBillboardId());
        categoryRepository.save(category);
        store.getCategories().add(existingCategory);
        billboard.getCategories().add(existingCategory);
        storeRepository.save(store);
        billboardRepository.save(billboard);
        return true;
    }

    @Override
    public boolean deleteCategory(String storeId, String categoryId) {
        Store store = storeRepository.findById(storeId).orElse(null);
        Category existingCategory = categoryRepository.findCategoryByStoreIdAndCategoryId(storeId,categoryId);
        Billboard billboard = billboardRepository.findById(existingCategory.getBillboardId()).orElse(null);
        if(store == null || billboard == null){
            return false;
        }
        billboard.getCategories().remove(existingCategory);
        store.getCategories().remove(existingCategory);
        storeRepository.save(store);
        billboardRepository.save(billboard);
        categoryRepository.delete(existingCategory);
        return true;
    }

    private CategoryWithBillboardDTO getCategoryWithBillboardDTO(Category category){
        if(category == null){
            return null;
        }
        RestTemplate restTemplate = new RestTemplate();
        Billboard billboard = restTemplate.getForObject("http://localhost:8080/store/" + category.getStoreId() +"/billboards/"+category.getBillboardId(),Billboard.class);
        CategoryWithBillboardDTO categoryWithBillboard = new CategoryWithBillboardDTO();
        categoryWithBillboard.setBillboard(billboard);
        categoryWithBillboard.setBillboardId(category.getBillboardId());
        categoryWithBillboard.setName(category.getName());
        categoryWithBillboard.setId(category.getId());
        categoryWithBillboard.setStoreId(category.getStoreId());
        categoryWithBillboard.setCreatedAt(category.getCreatedAt());
        categoryWithBillboard.setUpdatedAt(category.getUpdatedAt());
        return categoryWithBillboard;
    }
}

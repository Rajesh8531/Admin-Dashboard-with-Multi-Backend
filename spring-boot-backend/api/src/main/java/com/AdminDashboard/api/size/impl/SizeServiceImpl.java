package com.AdminDashboard.api.size.impl;

import com.AdminDashboard.api.billboard.Billboard;
import com.AdminDashboard.api.billboard.BillboardRepository;
import com.AdminDashboard.api.category.Category;
import com.AdminDashboard.api.category.CategoryRepository;
import com.AdminDashboard.api.category.CategoryService;
import com.AdminDashboard.api.category.CategoryWithBillboardDTO;
import com.AdminDashboard.api.size.Size;
import com.AdminDashboard.api.size.SizeRepository;
import com.AdminDashboard.api.size.SizeService;
import com.AdminDashboard.api.store.Store;
import com.AdminDashboard.api.store.StoreRepository;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.time.Instant;
import java.util.List;
import java.util.stream.Collectors;

@Service
public class SizeServiceImpl implements SizeService {


    StoreRepository storeRepository;
    SizeRepository sizeRepository;

    public SizeServiceImpl(SizeRepository sizeRepository, StoreRepository storeRepository) {
        this.sizeRepository = sizeRepository;
        this.storeRepository = storeRepository;

    }

    @Override
    public List<Size> getSizes(String storeId) {
        return sizeRepository.getSizesByStoreId(storeId);
    }

    @Override
    public boolean createSize(String storeId, Size size) {
        Store store = storeRepository.findById(storeId).orElse(null);
        if(store == null ){
            return false;
        }
        size.setCreatedAt(Instant.now());
        size.setUpdatedAt(Instant.now());
        size.setStore(store);


        store.getSizes().add(size);
        storeRepository.save(store);

        sizeRepository.save(size);
        return true;
    }

    @Override
    public Size getSize(String storeId, String sizeId) {
    return sizeRepository.findSizeByStoreIdAndSizeId(storeId,sizeId);
}

    @Override
    public boolean updateSize(String storeId, String sizeId, Size size) {
        Store store = storeRepository.findById(storeId).orElse(null);
        if(store == null ){
            return false;
        }
        Size existingSize = sizeRepository.findSizeByStoreIdAndSizeId(storeId,sizeId);
        existingSize.setName(size.getName());
        existingSize.setValue(size.getValue());
        existingSize.setUpdatedAt(Instant.now());
        sizeRepository.save(existingSize);
        store.getSizes().add(existingSize);
        storeRepository.save(store);
        return true;
    }

    @Override
    public boolean deleteSize(String storeId, String sizeId) {
        Store store = storeRepository.findById(storeId).orElse(null);
        Size existingSize = sizeRepository.findSizeByStoreIdAndSizeId(storeId,sizeId);
        if(store == null){
            return false;
        }
        store.getSizes().remove(existingSize);
        storeRepository.save(store);
        sizeRepository.delete(existingSize);
        return true;
    }
}

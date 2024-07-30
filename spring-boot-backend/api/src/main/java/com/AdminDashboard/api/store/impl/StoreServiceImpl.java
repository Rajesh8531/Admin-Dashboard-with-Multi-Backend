package com.AdminDashboard.api.store.impl;

import com.AdminDashboard.api.store.Store;
import com.AdminDashboard.api.store.StoreRepository;
import com.AdminDashboard.api.store.StoreService;
import org.springframework.stereotype.Service;

import java.time.Instant;
import java.util.List;

@Service
public class StoreServiceImpl implements StoreService {

    StoreRepository storeRepository;

    public StoreServiceImpl(StoreRepository storeRepository) {
        this.storeRepository = storeRepository;
    }

    @Override
    public List<Store> getStores() {
        return storeRepository.findAll();
    }

    @Override
    public boolean createStore(Store store) {
        store.setCreatedAt(Instant.now());
        store.setUpdatedAt(Instant.now());
        storeRepository.save(store);
        return true;
    }

    @Override
    public Store getStore(String id) {
        return storeRepository.findById(id).orElse(null);
    }

    @Override
    public boolean updateStore(String id, Store store) {
        Store existingStore = storeRepository.findById(id).orElse(null);
        if(existingStore == null){
            return false;
        }
        existingStore.setName(store.getName());
        existingStore.setUpdatedAt(Instant.now());
        storeRepository.save(existingStore);
        return true;
    }

    @Override
    public boolean deleteStore(String id) {
        Store existingStore = storeRepository.findById(id).orElse(null);
        if(existingStore == null){
            return false;
        }
        storeRepository.delete(existingStore);
        return true;
    }
}

package com.AdminDashboard.api.utils;

import com.AdminDashboard.api.store.Store;
import com.AdminDashboard.api.store.StoreRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class Utils {

    @Autowired
    StoreRepository storeRepository;

    public boolean isAuthorized(String userId, String storeId){
        Store store = storeRepository.findById(storeId).orElse(null);
        if(store == null){
            return false;
        }
        String newUserId = userId.replace("\"","");

        return store.getUserId().equals(newUserId);
    }
}

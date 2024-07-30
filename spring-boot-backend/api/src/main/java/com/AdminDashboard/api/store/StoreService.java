package com.AdminDashboard.api.store;

import java.util.List;

public interface StoreService {
    List<Store> getStores();
    boolean createStore(Store store);
    Store getStore(String id);
    boolean updateStore(String id,Store store);
    boolean deleteStore(String id);
}

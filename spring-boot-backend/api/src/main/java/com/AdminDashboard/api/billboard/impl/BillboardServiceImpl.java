package com.AdminDashboard.api.billboard.impl;

import com.AdminDashboard.api.billboard.Billboard;
import com.AdminDashboard.api.billboard.Billboard;
import com.AdminDashboard.api.billboard.BillboardRepository;
import com.AdminDashboard.api.billboard.BillboardService;
import com.AdminDashboard.api.store.Store;
import com.AdminDashboard.api.store.StoreRepository;
import com.AdminDashboard.api.utils.Utils;
import org.springframework.stereotype.Service;

import java.time.Instant;
import java.util.List;

@Service
public class BillboardServiceImpl implements BillboardService {

    BillboardRepository billboardRepository;
    StoreRepository storeRepository;

    public BillboardServiceImpl(BillboardRepository billboardRepository,StoreRepository storeRepository) {
        this.billboardRepository = billboardRepository;
        this.storeRepository = storeRepository;
    }

    @Override
    public List<Billboard> getBillboards(String storeId) {
        return billboardRepository.getBillboardByStoreId(storeId);
    }

    @Override
    public boolean createBillboard(String storeId, Billboard billboard) {
        Store store = storeRepository.findById(storeId).orElse(null);
        if(store == null){
            return false;
        }
        billboard.setStore(store);
        billboard.setCreatedAt(Instant.now());
        billboard.setUpdatedAt(Instant.now());
        billboardRepository.save(billboard);
        return true;
    }

    @Override
    public Billboard getBillboard(String storeId, String billboardId) {
        return billboardRepository.findBillboardByStoreIdAndBillboardId(storeId,billboardId);
    }

    @Override
    public boolean updateBillboard(String storeId, String billboardId, Billboard billboard) {
        Store store = storeRepository.findById(storeId).orElse(null);
        if(store == null){
            return false;
        }
        Billboard existingBillboard = billboardRepository.findBillboardByStoreIdAndBillboardId(storeId,billboardId);
        existingBillboard.setLabel(billboard.getLabel());
        billboardRepository.save(existingBillboard);
        store.getBillboards().add(existingBillboard);
        storeRepository.save(store);
        return true;
    }

    @Override
    public boolean deleteBillboard(String storeId, String billboardId) {
        Store store = storeRepository.findById(storeId).orElse(null);
        if(store == null){
            return false;
        }
        Billboard existingBillboard = billboardRepository.findBillboardByStoreIdAndBillboardId(storeId,billboardId);
        store.getBillboards().remove(existingBillboard);
        storeRepository.save(store);
        billboardRepository.delete(existingBillboard);
        return true;
    }
}

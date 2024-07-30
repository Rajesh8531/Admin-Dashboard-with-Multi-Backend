package com.AdminDashboard.api.color.impl;

import com.AdminDashboard.api.color.Color;
import com.AdminDashboard.api.color.ColorRepository;
import com.AdminDashboard.api.color.ColorService;
import com.AdminDashboard.api.size.Size;
import com.AdminDashboard.api.size.SizeRepository;
import com.AdminDashboard.api.size.SizeService;
import com.AdminDashboard.api.store.Store;
import com.AdminDashboard.api.store.StoreRepository;
import org.springframework.stereotype.Service;

import java.time.Instant;
import java.util.List;

@Service
public class ColorServiceImpl implements ColorService {


    StoreRepository storeRepository;
    ColorRepository colorRepository;

    public ColorServiceImpl(ColorRepository colorRepository, StoreRepository storeRepository) {
        this.colorRepository = colorRepository;
        this.storeRepository = storeRepository;

    }

    @Override
    public List<Color> getColors(String storeId) {
        return colorRepository.getColorsByStoreId(storeId);
    }

    @Override
    public boolean createColor(String storeId, Color color) {
        Store store = storeRepository.findById(storeId).orElse(null);
        if(store == null ){
            return false;
        }
        color.setCreatedAt(Instant.now());
        color.setUpdatedAt(Instant.now());
        color.setStore(store);


        store.getColors().add(color);
        storeRepository.save(store);

        colorRepository.save(color);
        return true;
    }

    @Override
    public Color getColor(String storeId, String colorId) {
    return colorRepository.findColorByStoreIdAndSizeId(storeId,colorId);
}

    @Override
    public boolean updateColor(String storeId, String colorId, Color color) {
        Store store = storeRepository.findById(storeId).orElse(null);
        if(store == null ){
            return false;
        }
        Color existingColor = colorRepository.findColorByStoreIdAndSizeId(storeId,colorId);
        existingColor.setName(color.getName());
        existingColor.setValue(color.getValue());
        existingColor.setUpdatedAt(Instant.now());
        colorRepository.save(color);
        store.getColors().add(existingColor);
        storeRepository.save(store);
        return true;
    }

    @Override
    public boolean deleteColor(String storeId, String colorId) {
        Store store = storeRepository.findById(storeId).orElse(null);
        Color existingColor = colorRepository.findColorByStoreIdAndSizeId(storeId,colorId);
        if(store == null){
            return false;
        }
        store.getColors().remove(existingColor);
        storeRepository.save(store);
        colorRepository.delete(existingColor);
        return true;
    }
}

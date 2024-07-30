package com.AdminDashboard.api.color;

import com.AdminDashboard.api.size.Size;

import java.util.List;

public interface ColorService {
    List<Color> getColors(String storeId);
    boolean createColor(String storeId, Color color);
    Color getColor(String storeId, String colorId);
    boolean updateColor(String storeId, String colorId, Color color);
    boolean deleteColor(String storeId, String colorId);
}

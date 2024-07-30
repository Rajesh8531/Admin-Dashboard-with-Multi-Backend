package com.AdminDashboard.api.size;

import java.util.List;

public interface SizeService {
    List<Size> getSizes(String storeId);
    boolean createSize(String storeId, Size size);
    Size getSize(String storeId, String sizeId);
    boolean updateSize(String storeId, String sizeId, Size size);
    boolean deleteSize(String storeId, String sizeId);
}

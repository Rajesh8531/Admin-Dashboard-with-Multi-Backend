package com.AdminDashboard.api.billboard;

import java.util.List;

public interface BillboardService {
    List<Billboard> getBillboards(String storeId);
    boolean createBillboard(String storeId, Billboard billboard);
    Billboard getBillboard(String storeId, String billboardId);
    boolean updateBillboard(String storeId, String billboardId, Billboard billboard);
    boolean deleteBillboard(String storeId,String billboardId);
}

package com.AdminDashboard.api.billboard;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;

public interface BillboardRepository extends JpaRepository<Billboard,String> {
    List<Billboard> getBillboardByStoreId(String storeId);

    @Query(value = "SELECT * FROM billboard b WHERE b.store_id = :storeId AND b.id = :billboardId", nativeQuery = true)
    Billboard findBillboardByStoreIdAndBillboardId(@Param("storeId") String storeId, @Param("billboardId") String billboardId);
}

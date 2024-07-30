package com.AdminDashboard.api.size;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;

public interface SizeRepository extends JpaRepository<Size,String> {
    List<Size> getSizesByStoreId(String storeId);

    @Query(value = "SELECT * FROM size s WHERE s.store_id = :storeId AND s.id = :sizeId", nativeQuery = true)
    Size findSizeByStoreIdAndSizeId(@Param("storeId") String storeId,@Param("sizeId") String sizeId);
}

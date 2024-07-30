package com.AdminDashboard.api.color;

import com.AdminDashboard.api.size.Size;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;

public interface ColorRepository extends JpaRepository<Color,String> {
    List<Color> getColorsByStoreId(String storeId);

    @Query(value = "SELECT * FROM color c WHERE c.store_id = :storeId AND c.id = :colorId", nativeQuery = true)
    Color findColorByStoreIdAndSizeId(@Param("storeId") String storeId,@Param("colorId") String colorId);
}

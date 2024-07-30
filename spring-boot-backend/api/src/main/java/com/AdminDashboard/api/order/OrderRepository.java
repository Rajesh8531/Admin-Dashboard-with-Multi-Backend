package com.AdminDashboard.api.order;

import com.AdminDashboard.api.order.models.Order;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface OrderRepository extends JpaRepository<Order,String> {
    List<Order> getOrdersById(String storeId);
}

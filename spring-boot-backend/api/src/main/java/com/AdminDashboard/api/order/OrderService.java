package com.AdminDashboard.api.order;

import com.AdminDashboard.api.order.models.Order;

import java.util.List;

public interface OrderService {
    List<Order> getOrders(String storeId);
}

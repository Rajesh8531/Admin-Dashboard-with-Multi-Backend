package com.AdminDashboard.api.order.impl;

import com.AdminDashboard.api.order.OrderRepository;
import com.AdminDashboard.api.order.OrderService;
import com.AdminDashboard.api.order.models.Order;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class OrderServiceImpl implements OrderService {
    OrderRepository orderRepository;

    public OrderServiceImpl(OrderRepository orderRepository) {
        this.orderRepository = orderRepository;
    }

    @Override
    public List<Order> getOrders(String storeId) {
        return orderRepository.getOrdersById(storeId);
    }
}

package com.AdminDashboard.api.order;

import com.AdminDashboard.api.order.models.Order;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@CrossOrigin(origins = "http://localhost:5173")
@RestController
@RequestMapping("/store/{storeId}/orders")
public class OrderController {
    OrderService orderService;

    public OrderController(OrderService orderService) {
        this.orderService = orderService;
    }

    @GetMapping
    public ResponseEntity<List<Order>> getOrders(@PathVariable String storeId){
        return new ResponseEntity<>(orderService.getOrders(storeId), HttpStatus.OK);
    }
}

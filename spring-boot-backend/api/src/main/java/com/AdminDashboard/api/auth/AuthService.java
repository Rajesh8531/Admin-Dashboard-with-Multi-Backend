package com.AdminDashboard.api.auth;

public interface AuthService {
    User registerUser(User user);
    User getUser(String id);
    User getUser(User user);
}

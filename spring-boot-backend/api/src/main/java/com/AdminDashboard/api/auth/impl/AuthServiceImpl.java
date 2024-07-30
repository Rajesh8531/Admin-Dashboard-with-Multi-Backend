package com.AdminDashboard.api.auth.impl;

import com.AdminDashboard.api.auth.User;
import com.AdminDashboard.api.auth.UserRepository;
import com.AdminDashboard.api.auth.AuthService;
import org.springframework.stereotype.Service;

import java.sql.Time;
import java.time.Instant;
import java.time.LocalDate;

@Service
public class AuthServiceImpl implements AuthService {

    UserRepository userRepository;

    public AuthServiceImpl(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @Override
    public User registerUser(User user) {

        user.setCreatedAt(Instant.now());
        user.setUpdatedAt(Instant.now());
        return userRepository.save(user);
    }

    @Override
    public User getUser(String id) {
        return userRepository.findById(id).orElse(null);
    }

    @Override
    public User getUser(User user) {
        return userRepository.findByEmail(user.getEmail());
    }


}

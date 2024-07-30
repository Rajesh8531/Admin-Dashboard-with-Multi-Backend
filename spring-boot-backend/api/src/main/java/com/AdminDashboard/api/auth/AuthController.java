package com.AdminDashboard.api.auth;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.web.bind.annotation.*;

@CrossOrigin(origins = "http://localhost:5173")
@RestController
@RequestMapping("/auth")
public class AuthController {
    AuthService authService;

    public AuthController(AuthService authService) {
        this.authService = authService;
    }

    @PostMapping("/signup")
    public ResponseEntity<AuthResponse> registerUser(@RequestBody User user){
        BCryptPasswordEncoder passwordEncoder = new BCryptPasswordEncoder();
        String hashedPassword = passwordEncoder.encode(user.getPassword());
        user.setPassword(hashedPassword);
        User savedUser = authService.registerUser(user);
        Algorithm algorithm = Algorithm.HMAC256("secret");
        String token = JWT.create()
                .withClaim("id",savedUser.getId())
                .withClaim("email",savedUser.getEmail())
                .sign(algorithm);
        AuthResponse response = new AuthResponse(savedUser.getId(),token,savedUser.getEmail(),savedUser.getName());
        return new ResponseEntity<>(response, HttpStatus.CREATED);
    }

    @PostMapping("/signin")
    public ResponseEntity<AuthResponse> signInUser(@RequestBody User user){
        User existingUser = authService.getUser(user);
        if(existingUser == null){
            return new ResponseEntity<>(null,HttpStatus.NOT_FOUND);
        }
        BCryptPasswordEncoder passwordEncoder = new BCryptPasswordEncoder();
        boolean isPasswordCorrect = passwordEncoder.matches(user.getPassword(), existingUser.getPassword());
        if(existingUser == null){
            return new ResponseEntity<>(null,HttpStatus.BAD_REQUEST);
        }
        Algorithm algorithm = Algorithm.HMAC256("secret");
        String token = JWT.create()
                .withClaim("id",existingUser.getId())
                .withClaim("email",existingUser.getEmail())
                .sign(algorithm);
        AuthResponse response = new AuthResponse(existingUser.getId(),token,existingUser.getEmail(),existingUser.getName());
        return new ResponseEntity<>(response,HttpStatus.ACCEPTED);
    }
}


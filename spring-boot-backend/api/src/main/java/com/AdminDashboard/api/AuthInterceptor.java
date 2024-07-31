package com.AdminDashboard.api;

import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.interfaces.Claim;
import com.auth0.jwt.interfaces.DecodedJWT;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

import java.util.Arrays;
import java.util.Enumeration;
import java.util.Map;

@Component
public class AuthInterceptor implements HandlerInterceptor {

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response,Object handler) throws Exception {
        String authorizationString = request.getHeader("Authorization");
        if(authorizationString == null){
            request.setAttribute("userId","");
            return true;
        }
        String tokenString = authorizationString.split(" ")[1];

        Algorithm algorithm = Algorithm.HMAC256("secret");
        JWTVerifier jwtVerifier = JWT.require(algorithm).build();
        DecodedJWT jwt = jwtVerifier.verify(tokenString);
        Map<String,Claim> jwtClaims= jwt.getClaims();
        request.setAttribute("userId",jwtClaims.get("id").toString());
        return true;
    }

}

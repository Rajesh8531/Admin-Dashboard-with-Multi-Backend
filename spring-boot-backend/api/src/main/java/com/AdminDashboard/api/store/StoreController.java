package com.AdminDashboard.api.store;

import com.AdminDashboard.api.auth.User;
import com.AdminDashboard.api.auth.UserRepository;
import com.AdminDashboard.api.utils.Utils;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@CrossOrigin(origins = "http://localhost:5173")
@RestController
@RequestMapping("/store")
public class StoreController {
    StoreService storeService;
    Utils utils;
    UserRepository userRepository;

    public StoreController(StoreService storeService,Utils utils,UserRepository userRepository) {
        this.storeService = storeService;
        this.utils = utils;
        this.userRepository = userRepository;
    }

    @GetMapping("/")
    public ResponseEntity<List<Store>> getStores(){
        return new ResponseEntity<>(storeService.getStores(), HttpStatus.OK);
    }

    @PostMapping("/")
    public ResponseEntity<String> createStore(@RequestBody Store store,@RequestAttribute String userId){
        userId = userId.replace("\"","");
        User user = userRepository.findById(userId).orElse(null);
        if(user == null){
            return new ResponseEntity<>("UnAuthenticated",HttpStatus.BAD_REQUEST);
        }
        store.setUserId(userId);
        storeService.createStore(store);
        return new ResponseEntity<>("Store Created",HttpStatus.CREATED);
    }

    @GetMapping("/{storeId}")
    public ResponseEntity<Store> getStore(@PathVariable String storeId){
        Store store = storeService.getStore(storeId);
        if(store == null){
            return new ResponseEntity<>(null,HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>(store,HttpStatus.OK);
    }

    @PatchMapping("/{storeId}")
    public ResponseEntity<String> updateStore(@PathVariable String storeId,@RequestBody Store store,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isUpdated = storeService.updateStore(storeId,store);
        if(!isUpdated){
            return new ResponseEntity<>("Store Not Found",HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>("Store Updated Successfully",HttpStatus.ACCEPTED);
    }

    @DeleteMapping("/{storeId}")
    public ResponseEntity<String> deleteStore(@PathVariable String storeId,@RequestBody Store store,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isDeleted = storeService.deleteStore(storeId);
        if(!isDeleted){
            return new ResponseEntity<>("Store Not Found",HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>("Store Deleted Successfully",HttpStatus.ACCEPTED);
    }
}

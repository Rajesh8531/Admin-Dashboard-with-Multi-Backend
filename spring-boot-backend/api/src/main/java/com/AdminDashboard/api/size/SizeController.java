package com.AdminDashboard.api.size;

import com.AdminDashboard.api.utils.Utils;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@CrossOrigin(origins = "http://localhost:5173")
@RestController
@RequestMapping("/store/{storeId}/sizes")
public class SizeController {

    SizeService sizeService;
    Utils utils;

    public SizeController(SizeService sizeService,Utils utils) {
        this.sizeService = sizeService;this.utils = utils;
    }

    @GetMapping
    public ResponseEntity<List<Size>> getSizes(@PathVariable String storeId){
        return new ResponseEntity<>(sizeService.getSizes(storeId), HttpStatus.OK);
    }

    @PostMapping
    public ResponseEntity<String> createSize(@PathVariable String storeId,@RequestBody Size size,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isSizeCreated = sizeService.createSize(storeId,size);
        if(!isSizeCreated){
            return new ResponseEntity<>("Size not Created",HttpStatus.BAD_REQUEST);
        }
        return new ResponseEntity<>("Size Created Successfully",HttpStatus.CREATED);
    }

    @GetMapping("/{sizeId}")
    public ResponseEntity<Size> getCategory(@PathVariable String storeId, @PathVariable String sizeId){
        if(sizeId.equals("new")) {
            return new ResponseEntity<>(null,HttpStatus.OK);
        }
        Size size = sizeService.getSize(storeId,sizeId);
        if(size == null){
            return new ResponseEntity<>(null,HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>(size,HttpStatus.OK);
    }

    @PatchMapping("/{sizeId}")
    public ResponseEntity<String> updateBillboard(@PathVariable String storeId,@PathVariable String sizeId,@RequestBody Size size,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isUpdated = sizeService.updateSize(storeId,sizeId,size);
        if(!isUpdated){
            return new ResponseEntity<>("Size Not Found",HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>("Size Updated",HttpStatus.OK);
    }

    @DeleteMapping("/{sizeId}")
    public ResponseEntity<String> deleteBillboard(@PathVariable String storeId,@PathVariable String sizeId,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isDeleted = sizeService.deleteSize(storeId,sizeId);
        if(!isDeleted){
            return new ResponseEntity<>("Size Not Found",HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>("Size Deleted",HttpStatus.OK);
    }
}

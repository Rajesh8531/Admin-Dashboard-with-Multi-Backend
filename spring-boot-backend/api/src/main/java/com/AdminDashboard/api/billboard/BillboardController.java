package com.AdminDashboard.api.billboard;

import com.AdminDashboard.api.utils.Utils;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@CrossOrigin(origins = "http://localhost:5173")
@RestController
@RequestMapping("/store/{storeId}/billboards")
public class BillboardController {

    Utils utils;
    BillboardService billboardService;

    public BillboardController(BillboardService billboardService,Utils utils) {
        this.billboardService = billboardService;
        this.utils = utils;
    }

    @GetMapping
    public ResponseEntity<List<Billboard>> getBillboards(@PathVariable String storeId){
        return new ResponseEntity<>(billboardService.getBillboards(storeId),HttpStatus.OK);
    }

    @PostMapping
    public ResponseEntity<String> createBillboard(@PathVariable String storeId,@RequestBody Billboard billboard,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isBillboardCreated = billboardService.createBillboard(storeId,billboard);
        if(!isBillboardCreated){
            return new ResponseEntity<>("Billboard not Created",HttpStatus.BAD_REQUEST);
        }
        return new ResponseEntity<>("Billboard Created Successfully",HttpStatus.CREATED);
    }

    @GetMapping("/{billboardId}")
    public ResponseEntity<Billboard> getBillboard(@PathVariable String storeId, @PathVariable String billboardId){
        if(billboardId.equals("new")) {
            return new ResponseEntity<>(null,HttpStatus.OK);
        }
        Billboard billboard = billboardService.getBillboard(storeId,billboardId);
        if(billboard == null){
            return new ResponseEntity<>(billboard,HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>(billboard,HttpStatus.OK);
    }

    @PatchMapping("/{billboardId}")
    public ResponseEntity<String> updateBillboard(@PathVariable String storeId,@PathVariable String billboardId,@RequestBody Billboard billboard,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isUpdated = billboardService.updateBillboard(storeId,billboardId,billboard);
        if(!isUpdated){
            return new ResponseEntity<>("Billboard Not Found",HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>("Billboard Updated",HttpStatus.OK);
    }

    @DeleteMapping("/{billboardId}")
    public ResponseEntity<String> deleteBillboard(@PathVariable String storeId,@PathVariable String billboardId,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isDeleted = billboardService.deleteBillboard(storeId,billboardId);
        if(!isDeleted){
            return new ResponseEntity<>("Billboard Not Found",HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>("Billboard Deleted",HttpStatus.OK);
    }
}

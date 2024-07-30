package com.AdminDashboard.api.color;

import com.AdminDashboard.api.utils.Utils;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@CrossOrigin(origins = "http://localhost:5173")
@RestController
@RequestMapping("/store/{storeId}/colors")
public class ColorController {

    Utils utils;
    ColorService colorService;

    public ColorController(ColorService colorService,Utils utils) {
        this.colorService = colorService;this.utils = utils;
    }

    @GetMapping
    public ResponseEntity<List<Color>> getSizes(@PathVariable String storeId){
        return new ResponseEntity<>(colorService.getColors(storeId), HttpStatus.OK);
    }

    @PostMapping
    public ResponseEntity<String> createSize(@PathVariable String storeId,@RequestBody Color color,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isSizeCreated = colorService.createColor(storeId,color);
        if(!isSizeCreated){
            return new ResponseEntity<>("Color not Created",HttpStatus.BAD_REQUEST);
        }
        return new ResponseEntity<>("Color Created Successfully",HttpStatus.CREATED);
    }

    @GetMapping("/{colorId}")
    public ResponseEntity<Color> getCategory(@PathVariable String storeId, @PathVariable String colorId){
        if(colorId.equals("new")) {
            return new ResponseEntity<>(null,HttpStatus.OK);
        }
        Color color = colorService.getColor(storeId,colorId);
        if(color == null){
            return new ResponseEntity<>(null,HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>(color,HttpStatus.OK);
    }

    @PatchMapping("/{colorId}")
    public ResponseEntity<String> updateBillboard(@PathVariable String storeId,@PathVariable String colorId,@RequestBody Color color,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isUpdated = colorService.updateColor(storeId,colorId,color);
        if(!isUpdated){
            return new ResponseEntity<>("Color Not Found",HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>("Color Updated",HttpStatus.OK);
    }

    @DeleteMapping("/{colorId}")
    public ResponseEntity<String> deleteBillboard(@PathVariable String storeId,@PathVariable String colorId,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isDeleted = colorService.deleteColor(storeId,colorId);
        if(!isDeleted){
            return new ResponseEntity<>("Color Not Found",HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>("Color Deleted",HttpStatus.OK);
    }
}

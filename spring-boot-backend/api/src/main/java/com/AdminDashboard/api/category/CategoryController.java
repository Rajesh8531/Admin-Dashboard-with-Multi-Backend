package com.AdminDashboard.api.category;

import com.AdminDashboard.api.utils.Utils;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@CrossOrigin(origins = "*")
@RestController
@RequestMapping("/store/{storeId}/categories")
public class CategoryController {

    Utils utils;
    CategoryService categoryService;

    public CategoryController(CategoryService categoryService,Utils utils) {
        this.categoryService = categoryService;
        this.utils = utils;
    }

    @GetMapping
    public ResponseEntity<List<CategoryWithBillboardDTO>> getBillboards(@PathVariable String storeId){
        return new ResponseEntity<>(categoryService.getCategories(storeId), HttpStatus.OK);
    }

    @PostMapping
    public ResponseEntity<String> createBillboard(@PathVariable String storeId,@RequestBody Category category,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isBillboardCreated = categoryService.createCategory(storeId,category);
        if(!isBillboardCreated){
            return new ResponseEntity<>("Category not Created",HttpStatus.BAD_REQUEST);
        }
        return new ResponseEntity<>("Category Created Successfully",HttpStatus.CREATED);
    }

    @GetMapping("/{categoryId}")
    public ResponseEntity<CategoryWithBillboardDTO> getCategory(@PathVariable String storeId, @PathVariable String categoryId){
        if(categoryId.equals("new")) {
            return new ResponseEntity<>(null,HttpStatus.OK);
        }
        CategoryWithBillboardDTO categoryWithBillboardDTO = categoryService.getCategory(storeId,categoryId);
        if(categoryWithBillboardDTO == null){
            return new ResponseEntity<>(null,HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>(categoryWithBillboardDTO,HttpStatus.OK);
    }

    @PatchMapping("/{categoryId}")
    public ResponseEntity<String> updateBillboard(@PathVariable String storeId,@PathVariable String categoryId,@RequestBody Category category,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isUpdated = categoryService.updateCategory(storeId,categoryId,category);
        if(!isUpdated){
            return new ResponseEntity<>("Category Not Found",HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>("Category Updated",HttpStatus.OK);
    }

    @DeleteMapping("/{categoryId}")
    public ResponseEntity<String> deleteBillboard(@PathVariable String storeId,@PathVariable String categoryId,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isDeleted = categoryService.deleteCategory(storeId,categoryId);
        if(!isDeleted){
            return new ResponseEntity<>("Category Not Found",HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>("Category Deleted",HttpStatus.OK);
    }
}

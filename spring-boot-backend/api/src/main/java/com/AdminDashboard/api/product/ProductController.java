package com.AdminDashboard.api.product;

import com.AdminDashboard.api.product.models.Product;
import com.AdminDashboard.api.product.models.ProductWithImageUrl;
import com.AdminDashboard.api.utils.Utils;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@CrossOrigin(origins = "http://localhost:5173")
@RestController
@RequestMapping("/store/{storeId}/products")
public class ProductController {

    ProductService productService;
    Utils utils;

    public ProductController(ProductService productService,Utils utils) {
        this.productService = productService;
        this.utils = utils;
    }

    @GetMapping
    public ResponseEntity<List<Product>> getProducts(@PathVariable String storeId){
        return new ResponseEntity<>(productService.getProducts(storeId), HttpStatus.OK);
    }

    @PostMapping
    public ResponseEntity<String> createProduct(@PathVariable String storeId, @RequestBody ProductWithImageUrl productWithImageUrl,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isCreated = productService.createProduct(storeId,productWithImageUrl);
        if(!isCreated){
            return new ResponseEntity<>("Product Not Created",HttpStatus.BAD_REQUEST);
        }
        return new ResponseEntity<>("Product Created",HttpStatus.CREATED);
    }

    @GetMapping("/{productId}")
    public ResponseEntity<Product> getProduct(@PathVariable String storeId,@PathVariable String productId){
        if(productId.equals("new")) {
            return new ResponseEntity<>(null,HttpStatus.OK);
        }
        Product product = productService.getProduct(storeId,productId);
        if(product == null){
            return  new ResponseEntity<>(null,HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>(product,HttpStatus.OK);
    }

    @PatchMapping("/{productId}")
    public ResponseEntity<String> updateProduct(@PathVariable String storeId,@PathVariable String productId,@RequestBody ProductWithImageUrl productWithImageUrl,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isUpdated = productService.updateProduct(storeId,productId,productWithImageUrl);
        if(!isUpdated){
            return new ResponseEntity<>("Product Not Updated",HttpStatus.BAD_REQUEST);
        }
        return new ResponseEntity<>("Product Updated",HttpStatus.CREATED);
    }

    @DeleteMapping("/{productId}")
    public ResponseEntity<String> deleteProduct(@PathVariable String storeId,@PathVariable String productId,@RequestAttribute String userId){
        boolean isAuthorized = utils.isAuthorized(userId,storeId);
        if(!isAuthorized){
            return new ResponseEntity<>("UnAuthorized",HttpStatus.UNAUTHORIZED);
        }
        boolean isDeleted = productService.deleteProduct(storeId,productId);
        if(!isDeleted){
            return new ResponseEntity<>("Product Not Deleted",HttpStatus.BAD_REQUEST);
        }
        return new ResponseEntity<>("Product Deleted",HttpStatus.CREATED);
    }
}

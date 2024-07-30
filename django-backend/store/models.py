from django.db import models

# Create your models here.
class Store(models.Model):
    id = models.CharField(primary_key=True,max_length=50)
    userId = models.CharField(max_length=50)
    name = models.CharField(max_length=50)
    createdAt = models.DateTimeField(auto_now_add=True)
    updatedAt = models.DateTimeField(auto_now=True)
    
    def __str__(self):
        return f"id -> {self.id} name -> {self.name}"

class Billboard(models.Model):
    id = models.CharField(primary_key=True,max_length=50)
    label = models.CharField(max_length=50)
    storeId = models.ForeignKey(Store,related_name='billboards',db_column='storeId',on_delete=models.PROTECT)
    createdAt = models.DateTimeField(auto_now_add=True)
    updatedAt = models.DateTimeField(auto_now=True)
    imageUrl = models.URLField(max_length=500)    
    
    def __str__(self):
        return f"id -> {self.id} name -> {self.label}"
    
class Category(models.Model):
    id = models.CharField(primary_key=True,max_length=50)
    createdAt = models.DateTimeField(auto_now_add=True)
    updatedAt = models.DateTimeField(auto_now=True)
    billboardId = models.ForeignKey(Billboard,related_name='categories',db_column='billboardId',on_delete=models.CASCADE)
    name = models.CharField(max_length=50)
    storeId = models.ForeignKey(Store,related_name='categories',db_column='storeId',on_delete=models.PROTECT)

class Size(models.Model):
    id = models.CharField(primary_key=True,max_length=50)
    createdAt = models.DateTimeField(auto_now_add=True)
    updatedAt = models.DateTimeField(auto_now=True)
    name = models.CharField(max_length=50)
    value = models.CharField(max_length=50)
    storeId = models.ForeignKey(Store,related_name='sizes',db_column='storeId',on_delete=models.PROTECT)
    
class Color(models.Model):
    id = models.CharField(primary_key=True,max_length=50)
    createdAt = models.DateTimeField(auto_now_add=True)
    updatedAt = models.DateTimeField(auto_now=True)
    name = models.CharField(max_length=50)
    value = models.CharField(max_length=50)
    storeId = models.ForeignKey(Store,related_name='colors',db_column='storeId',on_delete=models.PROTECT)
    
class Product(models.Model):
    id = models.CharField(primary_key=True,max_length=50)
    createdAt = models.DateTimeField(auto_now_add=True)
    updatedAt = models.DateTimeField(auto_now=True)
    name = models.CharField(max_length=50)
    price = models.CharField(max_length=50)
    isFeatured = models.BooleanField(default=False)
    isArchived = models.BooleanField(default=False)
    storeId = models.ForeignKey(Store,related_name='products',db_column='storeId',on_delete=models.PROTECT)
    categoryId = models.ForeignKey(Category,related_name='products',db_column='categoryId',on_delete=models.PROTECT)
    sizeId = models.ForeignKey(Size,related_name='products',db_column='sizeId',on_delete=models.PROTECT)
    colorId = models.ForeignKey(Color,related_name='products',db_column='colorId',on_delete=models.PROTECT)

class Image(models.Model):
    id = models.CharField(primary_key=True,max_length=50)
    productId = models.ForeignKey(Product,related_name='image',db_column='productId',on_delete=models.CASCADE)
    url = models.URLField(max_length=400)
    createdAt = models.DateTimeField(auto_now_add=True)
    updatedAt = models.DateTimeField(auto_now=True)

class Order(models.Model):
    id = models.CharField(primary_key=True,max_length=50)
    createdAt = models.DateTimeField(auto_now_add=True)
    updatedAt = models.DateTimeField(auto_now=True)
    isPaid = models.BooleanField(default=False)
    phone = models.CharField(max_length=20)
    address = models.TextField(max_length=100)
    storeId = models.ForeignKey(Store,related_name='orders',db_column='storeId',on_delete=models.PROTECT)
    
class OrderItem(models.Model):
    id = models.CharField(primary_key=True,max_length=50)
    orderId = models.ForeignKey(Order,related_name='orders',db_column='orderId',on_delete=models.PROTECT)
    productId = models.ForeignKey(Product,related_name='orders',db_column='productId',on_delete=models.PROTECT)
    
    
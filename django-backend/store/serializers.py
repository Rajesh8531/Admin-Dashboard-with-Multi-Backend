from rest_framework import serializers
from store.models import Store,Billboard,Category,Size,Color,Image,Product,Order,OrderItem
from store.utils import helper

class CategorySerializer(serializers.ModelSerializer):
    storeId = serializers.PrimaryKeyRelatedField(queryset=Store.objects.all())
    billboardId = serializers.PrimaryKeyRelatedField(queryset=Billboard.objects.all())
    billboard = serializers.SerializerMethodField()
    
    class Meta:
        model = Category
        fields = ['id','storeId','billboardId','createdAt','updatedAt','name','billboard']
    
    def get_billboard(self,obj):
        data = BillboardSerializer(obj.billboardId)
        return data.data
    

class BillboardSerializer(serializers.ModelSerializer):
    storeId = serializers.PrimaryKeyRelatedField(queryset=Store.objects.all())
    
    class Meta:
        model = Billboard
        fields = '__all__' 

class ImageSerializer(serializers.ModelSerializer):
    productId = serializers.PrimaryKeyRelatedField(queryset=Product.objects.all())
    
    class Meta:
        model = Image
        fields = "__all__"

class OrderItemSerializer(serializers.ModelSerializer):
    productId = serializers.PrimaryKeyRelatedField(queryset=Product.objects.all())
    orderId = serializers.PrimaryKeyRelatedField(queryset=Order.objects.all())
    
    class Meta:
        model = OrderItem
        fields = "__all__"

class OrderSerializer(serializers.ModelSerializer):
    storeId = serializers.PrimaryKeyRelatedField(queryset=Store.objects.all())
    
    class Meta:
        model = Order
        fields = "__all__"

class ProductSerializer(serializers.ModelSerializer):
    storeId = serializers.PrimaryKeyRelatedField(queryset=Store.objects.all())
    colorId = serializers.PrimaryKeyRelatedField(queryset=Color.objects.all())
    sizeId = serializers.PrimaryKeyRelatedField(queryset=Size.objects.all())
    categoryId = serializers.PrimaryKeyRelatedField(queryset=Category.objects.all())
    image = ImageSerializer(many=True,read_only=True)
    orderItems = OrderItemSerializer(many=True,read_only=True)
    
    class Meta:
        model = Product
        fields = '__all__'     

class ColorSerializer(serializers.ModelSerializer):
    storeId = serializers.PrimaryKeyRelatedField(queryset=Store.objects.all())
    products = ProductSerializer(many=True,read_only=True)
    
    class Meta:
        model = Color
        fields = '__all__' 

class CategorySerializer(serializers.ModelSerializer):
    storeId = serializers.PrimaryKeyRelatedField(queryset=Store.objects.all())
    billboardId = serializers.PrimaryKeyRelatedField(queryset=Billboard.objects.all())
    billboard = serializers.SerializerMethodField()
    products = ProductSerializer(many=True,read_only=True)
    
    class Meta:
        model = Category
        fields = ['id','storeId','billboardId','createdAt','updatedAt','name','billboard','products']
    
    def get_billboard(self,obj):
        data = BillboardSerializer(obj.billboardId)
        return data.data

class SizeSerializer(serializers.ModelSerializer):
    storeId = serializers.PrimaryKeyRelatedField(queryset=Store.objects.all())
    products = ProductSerializer(many=True,read_only=True)
    class Meta:
        model = Size
        fields = '__all__' 

class StoreSerializer(serializers.ModelSerializer):
    billboards = BillboardSerializer(many=True,read_only=True)
    categories = CategorySerializer(many=True,read_only=True)
    sizes = SizeSerializer(many=True,read_only=True)
    colors = ColorSerializer(many=True,read_only=True)
    
    class Meta:
        model = Store
        fields = '__all__'
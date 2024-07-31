from store.models import Product,Image
from store.serializers import ProductSerializer,ImageSerializer
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from store.models import Category,Color,Size
from store.serializers import CategorySerializer,SizeSerializer,ColorSerializer

from store.utils import auth_decorator, generateUUID, helper, isAuthorized

class ProductView(APIView):
    
    def get_products(self,storeId):
        objects =  Product.objects.all().filter(storeId=storeId)
        data = ProductSerializer(objects,many=True).data
        fullProducts = []
        for product in data:
            category = Category.objects.get(id=product.get('categoryId',''))
            color = Color.objects.get(id=product.get('colorId'))
            size = Size.objects.get(id=product.get('sizeId'))
            product['category'] = CategorySerializer(instance=category).data
            product['color'] = ColorSerializer(instance=color).data
            product['size'] = SizeSerializer(instance=size).data
            fullProducts.append(product)
        return fullProducts
     
    @auth_decorator()     
    def post(self,request,storeId):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        
        productId = generateUUID()
        data = {
            'id' : productId,
            'storeId' : storeId,
            'name' : request.data['name'],
            'categoryId' : request.data['categoryId'],
            'colorId' : request.data['colorId'],
            'price' : request.data['price'],
            'sizeId' : request.data['sizeId'],
            'isFeatured' : request.data['isFeatured'],
            'isArchived' : request.data['isArchived']
        }
        
        imageUrl : list = request.data['imageUrl']
            
        serializer = ProductSerializer(data=data)
        if serializer.is_valid():
            serializer.save()
            
            for url in imageUrl:
                imageId = generateUUID()
                imageData = {"id":imageId,"productId":productId,"url":url}
                imageSerializer = ImageSerializer(data=imageData)
                if imageSerializer.is_valid():
                    imageSerializer.save()
                else:
                    helper(serializer.errors)
                    raise("Invalid ImageData")
                
            return Response(serializer.data,status.HTTP_201_CREATED)
        return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)
    
    def get(self,request,storeId):
        data = self.get_products(storeId=storeId)
        return Response(data,status.HTTP_202_ACCEPTED)

class ProductDetailView(APIView):
    
    def get(self,request,storeId,productId,formate=None):
        if productId == 'new':
            return Response({},status.HTTP_204_NO_CONTENT)
        instance = Product.objects.get(storeId=storeId,id=productId)
        serializer = ProductSerializer(instance=instance)
        return Response(serializer.data,202)
    
    @auth_decorator()
    def patch(self,request,storeId,productId,formate=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        images = Image.objects.filter(productId=productId)
        for image in images:
            image.delete()
        instance = Product.objects.get(storeId=storeId,id=productId)
        data = {
            'id' : productId,
            'storeId' : storeId,
            'name' : request.data['name'],
            'categoryId' : request.data['categoryId'],
            'colorId' : request.data['colorId'],
            'price' : request.data['price'],
            'sizeId' : request.data['sizeId'],
            'isFeatured' : request.data['isFeatured'],
            'isArchived' : request.data['isArchived']
        }
        
        imageUrl : list = request.data['imageUrl']
        
        serializer = ProductSerializer(instance,data=data)
        if serializer.is_valid():
            serializer.save()
            for url in imageUrl:
                imageId = generateUUID()
                imageData = {"id":imageId,"productId":productId,"url":url}
                imageSerializer = ImageSerializer(data=imageData)
                if imageSerializer.is_valid():
                    imageSerializer.save()
                else:
                    helper(serializer.errors)
                    raise("Invalid ImageData")
            return Response(serializer.data,status.HTTP_202_ACCEPTED)
        return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)
    
    @auth_decorator()
    def delete(self,request,storeId,productId,format=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Product.objects.get(id=productId,storeId=storeId)
        instance.delete()
        return Response(status.HTTP_204_NO_CONTENT)
from store.models import Category
from store.serializers import CategorySerializer
from rest_framework.views import APIView
from store.controllers.utils import createBillboard
from rest_framework.response import Response
from rest_framework import status
from store.utils import generateUUID, isAuthorized,auth_decorator

class CategoryView(APIView):
    
    def get_categories(self,storeId):
        objects =  Category.objects.all()
        data = CategorySerializer(objects,many=True).data
        return filter(lambda x:x['storeId'] == storeId,data)
     
    @auth_decorator()     
    def post(self,request,storeId):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        id = generateUUID()
        data = {
            'id' : id,
            'storeId' : storeId,
            'billboardId' : request.data['billboardId'],
            'name' : request.data['name']
        }
        serializer = CategorySerializer(data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data,status.HTTP_201_CREATED)
        return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)
    
    def get(self,request,storeId):
        data = self.get_categories(storeId=storeId)
        return Response(data,status.HTTP_202_ACCEPTED)

class CategoryDetailView(APIView):
    
    def get(self,request,storeId,categoryId,formate=None):
        if categoryId == 'new':
            return Response({},status.HTTP_204_NO_CONTENT)
        instance = Category.objects.get(storeId=storeId,id=categoryId)
        serializer = CategorySerializer(instance=instance)
        return Response(serializer.data,202)
    
    @auth_decorator()
    def patch(self,request,storeId,categoryId,formate=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Category.objects.get(storeId=storeId,id=categoryId)
        data={'id':categoryId,'storeId':storeId,'name':request.data['name'],'billboardId':request.data['billboardId']}
        serializer = CategorySerializer(instance,data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data,status.HTTP_202_ACCEPTED)
        return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)
    
    @auth_decorator()
    def delete(self,request,storeId,categoryId,format=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Category.objects.get(id=categoryId,storeId=storeId)
        instance.delete()
        return Response(status.HTTP_204_NO_CONTENT)
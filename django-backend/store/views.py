from django.shortcuts import render
from rest_framework.views import APIView
from store.serializers import StoreSerializer
from store.models import Store
from rest_framework.response import Response
from rest_framework.request import Request
from store import utils
from store.utils import auth_decorator, isAuthorized
from rest_framework import status

# Create your views here.

class StoreView(APIView):
    
    @auth_decorator()
    def post(self,request):
        userId = request.META.get("userId")
        id = utils.generateUUID()
        data = {'id':id,'name':request.data['name'],'userId':userId}
        serializer = StoreSerializer(data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data,202)
        return Response(serializer.errors,400)
    
    def get(self,request:Request):
        objects = Store.objects.all()
        serializer = StoreSerializer(objects,many=True)
        return Response(serializer.data,202)

class StoreDetailView(APIView):
    
    def get(self,request,storeId):
        instance = Store.objects.get(id=storeId)
        serializer = StoreSerializer(instance)
        return Response(serializer.data,202)

    @auth_decorator()
    def patch(self,request,storeId):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Store.objects.get(id=storeId)
        data={'id':storeId,'name':request.data['name'],'userId' : userId}
        serializer = StoreSerializer(instance,data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data,status.HTTP_202_ACCEPTED)
        return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)
    
    @auth_decorator()
    def delete(self,request,storeId,format=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Store.objects.get(id=storeId)
        instance.delete()
        return Response(status.HTTP_204_NO_CONTENT)
    

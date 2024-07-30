from store.models import Size
from store.serializers import SizeSerializer
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from store.utils import auth_decorator, generateUUID, isAuthorized

class SizeView(APIView):
    
    def get_sizes(self,storeId):
        objects =  Size.objects.all()
        data = SizeSerializer(objects,many=True).data
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
            'value' : request.data['value'],
            'name' : request.data['name']
        }
        serializer = SizeSerializer(data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data,status.HTTP_201_CREATED)
        return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)
    
    def get(self,request,storeId):
        data = self.get_sizes(storeId=storeId)
        return Response(data,status.HTTP_202_ACCEPTED)

class SizeDetailView(APIView):
    
    def get(self,request,storeId,sizeId,formate=None):
        if sizeId == 'new':
            return Response({},status.HTTP_204_NO_CONTENT)
        instance = Size.objects.get(storeId=storeId,id=sizeId)
        serializer = SizeSerializer(instance=instance)
        return Response(serializer.data,202)
    
    @auth_decorator()
    def patch(self,request,storeId,sizeId,formate=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Size.objects.get(storeId=storeId,id=sizeId)
        data={'id':sizeId,'storeId':storeId,'name':request.data['name'],'value':request.data['value']}
        serializer = SizeSerializer(instance,data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data,status.HTTP_202_ACCEPTED)
        return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)
    
    @auth_decorator()
    def delete(self,request,storeId,sizeId,format=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Size.objects.get(id=sizeId,storeId=storeId)
        instance.delete()
        return Response(status.HTTP_204_NO_CONTENT)
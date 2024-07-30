from store.models import Color
from store.serializers import ColorSerializer
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from store.utils import auth_decorator, generateUUID, isAuthorized

class ColorView(APIView):
    
    def get_colors(self,storeId):
        objects =  Color.objects.all()
        data = ColorSerializer(objects,many=True).data
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
        serializer = ColorSerializer(data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data,status.HTTP_201_CREATED)
        return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)
    
    def get(self,request,storeId):
        data = self.get_colors(storeId=storeId)
        return Response(data,status.HTTP_202_ACCEPTED)

class ColorDetailView(APIView):
    
    def get(self,request,storeId,colorId,formate=None):
        if colorId == 'new':
            return Response({},status.HTTP_204_NO_CONTENT)
        instance = Color.objects.get(storeId=storeId,id=colorId)
        serializer = ColorSerializer(instance=instance)
        return Response(serializer.data,202)
    
    @auth_decorator()
    def patch(self,request,storeId,colorId,formate=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Color.objects.get(storeId=storeId,id=colorId)
        data={'id':colorId,'storeId':storeId,'name':request.data['name'],'value':request.data['value']}
        serializer = ColorSerializer(instance,data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data,status.HTTP_202_ACCEPTED)
        return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)
    
    @auth_decorator()
    def delete(self,request,storeId,colorId,format=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Color.objects.get(id=colorId,storeId=storeId)
        instance.delete()
        return Response(status.HTTP_204_NO_CONTENT)
from store.models import Billboard
from store.serializers import BillboardSerializer
from rest_framework.views import APIView
from store.controllers.utils import createBillboard
from rest_framework.response import Response
from rest_framework import status
from store.utils import auth_decorator,isAuthorized

class BillboardView(APIView):
    
    def get_billboards(self,storeId):
        objects =  Billboard.objects.all()
        data = BillboardSerializer(objects,many=True).data
        return filter(lambda x:x['storeId'] == storeId,data)
     
    @auth_decorator()     
    def post(self,request,storeId):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        return createBillboard(request,storeId)
    
    def get(self,request,storeId):
        data = self.get_billboards(storeId=storeId)
        return Response(data,status.HTTP_202_ACCEPTED)

class BillboardDetailView(APIView):
    
    def get(self,request,storeId,billboardId,formate=None):
        if billboardId == 'new':
            return Response({},status.HTTP_204_NO_CONTENT)
        instance = Billboard.objects.get(storeId=storeId,id=billboardId)
        serializer = BillboardSerializer(instance=instance)
        return Response(serializer.data,202)
    
    @auth_decorator() 
    def patch(self,request,storeId,billboardId,formate=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Billboard.objects.get(storeId=storeId,id=billboardId)
        data={'id':billboardId,'storeId':storeId,'label':request.data['label'],'imageUrl':request.data['imageUrl']}
        serializer = BillboardSerializer(instance,data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data,status.HTTP_202_ACCEPTED)
        return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)
    
    @auth_decorator() 
    def delete(self,request,storeId,billboardId,format=None):
        userId = request.META.get("userId")
        if(isAuthorized(userId,storeId) == False):
            return Response("Unauthorized",status.HTTP_401_UNAUTHORIZED)
        instance = Billboard.objects.get(id=billboardId,storeId=storeId)
        instance.delete()
        return Response(status.HTTP_204_NO_CONTENT)
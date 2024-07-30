from store.models import Order,OrderItem
from store.serializers import OrderItemSerializer,OrderSerializer
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from store.utils import generateUUID

class OrderView(APIView):
    
    def get_orders(self,storeId):
        objects =  Order.objects.all()
        data = OrderSerializer(objects,many=True).data
        return filter(lambda x:x['storeId'] == storeId,data)
          
    def get(self,request,storeId):
        data = self.get_orders(storeId=storeId)
        return Response(data,status.HTTP_202_ACCEPTED)
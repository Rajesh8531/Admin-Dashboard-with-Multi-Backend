from store.models import Billboard
from store.serializers import BillboardSerializer
from rest_framework.response import Response
from store.utils import generateUUID
from rest_framework import status

def createBillboard(request,storeId)->Response:
    id = generateUUID()
    data = {
            'id':id,
            'imageUrl':request.data['imageUrl'],
            'label':request.data['label'],
            'storeId' : storeId
           }
    serializer = BillboardSerializer(data=data)
    
    if serializer.is_valid():
        serializer.save()
        return Response(serializer.data,status.HTTP_201_CREATED)
    return Response(serializer.errors,status.HTTP_400_BAD_REQUEST)

    
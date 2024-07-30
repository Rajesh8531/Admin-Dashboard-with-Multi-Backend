from django.shortcuts import render
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from user.serializers import UserSerializer
from user.models import User
import uuid
import jwt
import bcrypt
from dotenv import load_dotenv
import os

load_dotenv()

# Create your views here.
class UserSignUpView(APIView):
    
    def post(self,request,format=None):
        salt = bytes(os.getenv("HASH_SALT"),encoding="utf-8")        
        hashedPassword = bcrypt.hashpw(bytes(request.data['password'],encoding='utf-8'),salt=salt)      
        id = str(uuid.uuid4())                     
        user = {
                'id' : id,
                'name' : request.data['name'],
                'email':request.data['email'],
                'password' : hashedPassword.decode(encoding="utf-8")
                }
        serializer = UserSerializer(data=user)
        tokenString = jwt.encode({"id":id,"email":user["email"]},"test",algorithm="HS256")
        print(tokenString)
        if serializer.is_valid():
            serializer.save()
            return Response({"id":id,"name":user["name"],"email":user["email"],"token":tokenString},status=status.HTTP_201_CREATED)
        return Response(serializer.errors,status=status.HTTP_400_BAD_REQUEST)
    
class UserSignInView(APIView):
    
    def get_user(self,**kwargs) -> User:
        try:
            return User.objects.get(**kwargs)
        except User.DoesNotExist:
            return None

    
    def post(self,request,format=None):
        
        data = self.get_user(email=request.data["email"])
        
        if data is None:
            return Response("User Not Found",status=status.HTTP_404_NOT_FOUND) 
        
        serializer = UserSerializer(data)
        user = serializer.data
        salt = bytes(os.getenv("HASH_SALT"),encoding="utf-8")
        
        isCorrectPassword : bool = user['password'] == bcrypt.hashpw(bytes(request.data['password'],encoding='utf-8'),salt).decode('utf-8')
        if isCorrectPassword == False:
            return Response("Invalid Credentials",status=status.HTTP_406_NOT_ACCEPTABLE)
        tokenString = jwt.encode({"id":user['id'],"email":user["email"]},"test",algorithm="HS256")

        return Response(
                        {"id" : user['id'],
                         "email" : user["email"],
                         "name" : user["name"],
                         "token" : tokenString
                         },
                        status=status.HTTP_202_ACCEPTED
                        )
    

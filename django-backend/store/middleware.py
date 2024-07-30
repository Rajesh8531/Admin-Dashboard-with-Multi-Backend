import jwt
from rest_framework.response import Response
from rest_framework import status

class AuthMiddleware:
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):
        authorizationHeader = request.headers.get("Authorization",None)
        if authorizationHeader is None:
            return Response("Unauthenticated",status.HTTP_401_UNAUTHORIZED)
        authString = authorizationHeader.split(' ')[1]
        obj = jwt.decode(authString,"test",algorithms='HS256')
        request.META["userId"] = obj['id']
        response = self.get_response(request)
        return response


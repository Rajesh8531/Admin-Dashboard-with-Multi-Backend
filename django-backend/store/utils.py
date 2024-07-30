from termcolor import colored
import uuid
from django.utils.decorators import method_decorator
from store.middleware import AuthMiddleware
from store.models import Store

def generateUUID()->str:
    return str(uuid.uuid4())

def helper(data:any)->None:
    print("\n\n",colored("START"),"\n\n")
    print(colored(data))
    print("\n\n",colored("END"),"\n\n")

def authMiddleware(view_func):
    def _wrapped_view(request,*args,**kwargs):
        middleware = AuthMiddleware(lambda req:view_func(req,*args,**kwargs))
        return middleware(request)
    return _wrapped_view


def auth_decorator():
    return method_decorator(authMiddleware)

def isAuthorized(userId,storeId):
    store = Store.objects.get(id=storeId)
    print(userId,store.userId)
    return userId == str(store.userId)
        

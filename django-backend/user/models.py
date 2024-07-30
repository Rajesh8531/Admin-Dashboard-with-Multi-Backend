from django.db import models

# Create your models here.
class User(models.Model):
    id = models.CharField(max_length=100,primary_key=True)
    name = models.CharField(max_length=50)
    email = models.EmailField(max_length=50,unique=True)
    password = models.TextField(max_length=200)
    createdAt = models.DateTimeField(auto_now_add=True)
    updatedAt = models.DateTimeField(auto_now=True)
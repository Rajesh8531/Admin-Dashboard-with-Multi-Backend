
from django.urls import path,include
from store.views import StoreView,StoreDetailView
from store.category_view import CategoryView,CategoryDetailView
from store.billboard_view import BillboardView,BillboardDetailView
from store.size_view import SizeView,SizeDetailView
from store.color_view import ColorView,ColorDetailView
from store.product_view import ProductView,ProductDetailView
from store.order_view import OrderView

urlpatterns = [
    path('', StoreView.as_view(),name="store"),
    path('<str:storeId>', StoreDetailView.as_view(),name="storeDetail"),
    path('<str:storeId>/',include([
        path('billboards',BillboardView.as_view(),name="billboards"),
        path('billboards/<str:billboardId>',BillboardDetailView.as_view(),name="billboardDetail"),
        path('categories',CategoryView.as_view(),name="categories"),
        path('categories/<str:categoryId>',CategoryDetailView.as_view(),name="categoryDetail"),
        path('sizes',SizeView.as_view(),name="size"),
        path('sizes/<str:sizeId>',SizeDetailView.as_view(),name="sizeDetail"),
        path('colors',ColorView.as_view(),name="color"),
        path('colors/<str:colorId>',ColorDetailView.as_view(),name="colorDetail"),
        path('products',ProductView.as_view(),name="product"),
        path('products/<str:productId>',ProductDetailView.as_view(),name="productDetail"),
        path('orders',OrderView.as_view(),name="order"),
    ]))
]
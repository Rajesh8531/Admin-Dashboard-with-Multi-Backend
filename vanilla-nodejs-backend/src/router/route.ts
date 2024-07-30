import { IncomingMessage, ServerResponse } from "http"
import { signIn, signUp } from "../controllers/auth-controller"
import { HttpStatusCode } from "axios";
import { isUrlMatching } from "../utils/utils";
import url from 'url'
import { storesController, uniqueStoreController } from "../controllers/store-controller";
import { billboardsController, uniqueBillboardController } from "../controllers/billboard-controller";
import { categoriesController, uniqueCategoryController } from "../controllers/category-controller";
import { sizesController, uniqueSizeController } from "../controllers/size-controller";
import { colorsController, uniqueColorController } from "../controllers/color-controller";
import { productsController, uniqueProductController } from "../controllers/product-controller";
import { ordersController } from "../controllers/order-controller";

const routes = [
    {path : '/auth/signup' , handler : signUp},
    {path : '/auth/signin' , handler : signIn},
    {path : '/store/' , handler : storesController},
    {path : '/store/:storeId' , handler : uniqueStoreController},
    {path : '/store/:storeId/billboards' , handler : billboardsController},
    {path : '/store/:storeId/billboards/:billboardId' , handler : uniqueBillboardController},
    {path : '/store/:storeId/categories/' , handler : categoriesController},
    {path : '/store/:storeId/categories/:categoryId' , handler : uniqueCategoryController},
    {path : '/store/:storeId/sizes' , handler : sizesController},
    {path : '/store/:storeId/sizes/:sizeId' , handler : uniqueSizeController},
    {path : '/store/:storeId/colors' , handler : colorsController},
    {path : '/store/:storeId/colors/:colorId' , handler : uniqueColorController},
    {path : '/store/:storeId/products' , handler : productsController},
    {path : '/store/:storeId/products/:productId' , handler : uniqueProductController},
    {path : '/store/:storeId/orders' , handler : ordersController},
]

export const router = async (req:IncomingMessage,res:ServerResponse)=>{
    const parsedUrl = url.parse(req.url as string,true)
    const pathname = parsedUrl.pathname
    const route = routes.find( r => isUrlMatching(pathname as string,r.path))
    if(route) {
        try {
            await route.handler(req,res)
        } catch (error) {
            res.writeHead(HttpStatusCode.InternalServerError)
            res.end("InternalServerError")
        }
    } else {
        res.writeHead(HttpStatusCode.NotFound)
        res.end("Not Found")
    }
}


import { IncomingMessage, ServerResponse } from "http";
import { getParams, getRequestBody, resultHeader } from "../utils/utils";
import { DBDeleteRecord,DBUpdateRecord, DBgetRecord, DBgetRecords, DBcreateSize, DBcreateProduct, DBcreateImage } from "../db/db-operation";
import { HttpStatusCode } from "axios";
import { Product } from "../type";

export const productsController = async ( req:IncomingMessage,res:ServerResponse) => {

    switch (req.method){
        case "POST" : createProduct(req,res)
                        return;
        case "GET" : getProducts(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
    }
}

const createProduct = async (req:IncomingMessage,res:ServerResponse) => {
    getRequestBody(req,async(data)=>{
        try {
            const {storeId} = getParams(req,'/store/:storeId/products')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {name,price,isFeatured,isArchived,sizeId,colorId,imageUrl,categoryId} = data
            const id = await DBcreateProduct(name,price,storeId,categoryId,isFeatured,isArchived,sizeId,colorId) as string
            for (let image of imageUrl){
                await DBcreateImage(image,id);
            }
            res.writeHead(HttpStatusCode.Created);
            res.write("PRODUCT CREATED")
            res.end()
            return
        } catch (error) {
            console.log("CREATE PRODUCT CONTROLLER",error)
            res.writeHead(HttpStatusCode.BadRequest);
            res.end("SOMETHING WENT WRONG")
        }
        
    })
}

const getProducts = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId} = getParams(req,'/store/:storeId/products')
        let fullProducts :any[]= []
        const products = await DBgetRecords('product','storeId=?',storeId) as Product[]
        for (let product of products){
            let category = await DBgetRecord('category','id = ?',product.categoryId)
            let color = await DBgetRecord('color','id = ?',product.colorId)
            let size = await DBgetRecord('size','id = ?',product.sizeId)
            fullProducts.push({...product,category,color,size})
        }
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify(fullProducts))
        res.end()

    } catch (error) {
        console.log("GET PRODUCTS CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const uniqueProductController = async (req:IncomingMessage,res:ServerResponse)=>{
    switch (req.method){
        case "PATCH" : updateProduct(req,res)
                        return;
        case "GET" : getProduct(req,res)
                        return;
        case "DELETE" : deleteProduct(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
            return;
    }
}

export const getProduct = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId,productId} = getParams(req,'/store/:storeId/products/:productId')
        if(productId == 'new'){
            res.writeHead(HttpStatusCode.Ok);
            res.end()
            return;
        }
        const product = await DBgetRecord('product','id = ? AND storeId = ? ',productId,storeId) as Product
        const category = await DBgetRecord('category','id = ? AND storeId = ? ',product.categoryId,storeId)
        const color = await DBgetRecord('color','id = ? AND storeId = ? ',product.colorId,storeId)
        const size = await DBgetRecord('size','id = ? AND storeId = ? ',product.sizeId,storeId)
        const image = await DBgetRecords('image','productId = ? ',product.id)
        if(!product){
            res.writeHead(HttpStatusCode.NotFound)
            res.end("NO PRODUCT FOUND")
            return
        }
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify({...product,category,color,size,image}))
        res.end()
        return;
        
    } catch (error) {
        console.log("GET PRODUCT CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const updateProduct = async (req:IncomingMessage,res:ServerResponse)=>{
    getRequestBody(req,async ( data )=>{
        try {
            const {storeId,productId} = getParams(req,'/store/:storeId/products/:productId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {name,price,isFeatured,isArchived,sizeId,colorId,imageUrl,categoryId} = data
            await DBDeleteRecord('image','WHERE productId = ?',productId)
            await DBUpdateRecord('product','SET name = ?,price=?,isFeatured=?,isArchived=?,sizeId=?,colorId=?,categoryId=? WHERE id = ? AND storeId = ?'
                                ,name,price,isFeatured,isArchived,sizeId,colorId,categoryId,productId,storeId)
            for (let image of imageUrl){
                await DBcreateImage(image,productId);
            }
            res.writeHead(HttpStatusCode.Accepted);
            res.write("PRODUCT UPDATED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("UPDATE PRODUCT CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
        }
    })
}

export const deleteProduct = async (req:IncomingMessage,res:ServerResponse)=>{
        try {
            const {storeId,productId} = getParams(req,'/store/:storeId/products/:productId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            await DBDeleteRecord('image','WHERE productId = ?',productId)
            const requestHeader = await DBDeleteRecord('product',' WHERE id = ? AND storeId = ?',productId,storeId) as resultHeader
            if(requestHeader.affectedRows < 1){
                res.writeHead(HttpStatusCode.NotFound);
                res.write("PRODUCT NOT FOUND")
                res.end()
                return
            }
            res.writeHead(HttpStatusCode.Accepted);
            res.write("PRODUCT DELETED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("DELETE PRODUCT CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
            res.end("Something went wrong")
        }
}



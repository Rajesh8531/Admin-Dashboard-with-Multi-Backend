import { IncomingMessage, ServerResponse } from "http";
import { getParams, getRequestBody, resultHeader } from "../utils/utils";
import { DBDeleteRecord,DBUpdateRecord, DBCreateCategory, DBgetRecord, DBgetRecords } from "../db/db-operation";
import { HttpStatusCode } from "axios";
import { Category } from "../type";

export const categoriesController = async ( req:IncomingMessage,res:ServerResponse) => {

    switch (req.method){
        case "POST" : createCategory(req,res)
                        return;
        case "GET" : getCategories(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
    }
}

const createCategory = async (req:IncomingMessage,res:ServerResponse) => {
    getRequestBody(req,async(data)=>{
        try {
            const {storeId} = getParams(req,'/store/:storeId/categories')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {name,billboardId} = data
            await DBCreateCategory(name,billboardId,storeId) 
            res.writeHead(HttpStatusCode.Created);
            res.write("CATEGORY CREATED")
            res.end()
            return
        } catch (error) {
            console.log("CREATE CATEGORY CONTROLLER",error)
            res.writeHead(HttpStatusCode.BadRequest);
            res.end("SOMETHING WENT WRONG")
        }
        
    })
}

const getCategories = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId} = getParams(req,'/store/:storeId/categories')
        var fullCategories : any[] = [] 
        const categories = await DBgetRecords('category','storeId=?',storeId) as Category[]
        for (let category of categories){
            let billboard = await DBgetRecord('billboard','id = ?',category.billboardId)
            let fullCategory = {...category,billboard}
            fullCategories.push(fullCategory)
        }
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify(fullCategories))
        res.end()

    } catch (error) {
        console.log("GET CATEGORIES CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const uniqueCategoryController = async (req:IncomingMessage,res:ServerResponse)=>{
    switch (req.method){
        case "PATCH" : updateCategory(req,res)
                        return;
        case "GET" : getCategory(req,res)
                        return;
        case "DELETE" : deleteCategory(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
            return;
    }
}

export const getCategory = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId,categoryId} = getParams(req,'/store/:storeId/categories/:categoryId')
        if(categoryId == 'new'){
            res.writeHead(HttpStatusCode.Ok);
            res.end()
            return;
        }
        const category = await DBgetRecord('category','id = ? AND storeId = ? ',categoryId,storeId)
        if(!category){
            res.writeHead(HttpStatusCode.NotFound)
            res.end("NO CATEGORY FOUND")
            return
        }
        res.writeHead(HttpStatusCode.Accepted);
        const products = await DBgetRecords('product','categoryId = ? ',categoryId)
        res.write(JSON.stringify({...category,products}))
        res.end()
        return;
        
    } catch (error) {
        console.log("GET CATEGORY CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const updateCategory = async (req:IncomingMessage,res:ServerResponse)=>{
    getRequestBody(req,async ( data )=>{
        try {
            const {storeId,categoryId} = getParams(req,'/store/:storeId/categories/:categoryId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {name,billboardId} = data
            await DBUpdateRecord('category','SET name = ? , billboardId = ? WHERE id = ? AND storeId = ?',name,billboardId,categoryId,storeId)
            res.writeHead(HttpStatusCode.Accepted);
            res.write("CATEGORY UPDATED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("UPDATE CATEGORY CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
        }
    })
}

export const deleteCategory = async (req:IncomingMessage,res:ServerResponse)=>{
        try {
            const {storeId,categoryId} = getParams(req,'/store/:storeId/categories/:categoryId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const requestHeader = await DBDeleteRecord('category',' WHERE id = ? AND storeId = ?',categoryId,storeId) as resultHeader
            if(requestHeader.affectedRows < 1){
                res.writeHead(HttpStatusCode.NotFound);
                res.write("CATEGORY NOT FOUND")
                res.end()
                return
            }
            res.writeHead(HttpStatusCode.Accepted);
            res.write("CATEGORY DELETED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("DELETE CATEGORY CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
            res.end("Something went wrong")
        }
}



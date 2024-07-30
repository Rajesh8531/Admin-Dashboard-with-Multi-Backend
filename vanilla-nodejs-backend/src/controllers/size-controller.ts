import { IncomingMessage, ServerResponse } from "http";
import { getParams, getRequestBody, resultHeader } from "../utils/utils";
import { DBDeleteRecord,DBUpdateRecord, DBgetRecord, DBgetRecords, DBcreateSize } from "../db/db-operation";
import { HttpStatusCode } from "axios";

export const sizesController = async ( req:IncomingMessage,res:ServerResponse) => {
    switch (req.method){
        case "POST" : createSize(req,res)
                        return;
        case "GET" : getSizes(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
    }
}

const createSize = async (req:IncomingMessage,res:ServerResponse) => {
    getRequestBody(req,async(data)=>{
        try {
            const {storeId} = getParams(req,'/store/:storeId/sizes')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {name,value} = data
            await DBcreateSize(name,value,storeId) 
            res.writeHead(HttpStatusCode.Created);
            res.write("SIZE CREATED")
            res.end()
            return
        } catch (error) {
            console.log("CREATE SIZE CONTROLLER",error)
            res.writeHead(HttpStatusCode.BadRequest);
            res.end("SOMETHING WENT WRONG")
        }
        
    })
}

const getSizes = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId} = getParams(req,'/store/:storeId/sizes')
        const sizes = await DBgetRecords('size','storeId=?',storeId)
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify(sizes))
        res.end()

    } catch (error) {
        console.log("GET SIZES CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const uniqueSizeController = async (req:IncomingMessage,res:ServerResponse)=>{
    switch (req.method){
        case "PATCH" : updateSize(req,res)
                        return;
        case "GET" : getSize(req,res)
                        return;
        case "DELETE" : deleteSize(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
            return;
    }
}

export const getSize = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId,sizeId} = getParams(req,'/store/:storeId/sizes/:sizeId')
        if(sizeId == 'new'){
            res.writeHead(HttpStatusCode.Ok);
            res.end()
            return;
        }
        const size = await DBgetRecord('size','id = ? AND storeId = ? ',sizeId,storeId)
        if(!size){
            res.writeHead(HttpStatusCode.NotFound)
            res.end("NO SIZE FOUND")
            return
        }
        const products = await DBgetRecords('product','sizeId = ? ',sizeId)
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify({...size,products}))
        res.end()
        return;
        
    } catch (error) {
        console.log("GET SIZE CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const updateSize = async (req:IncomingMessage,res:ServerResponse)=>{
    getRequestBody(req,async ( data )=>{
        try {
            const {storeId,sizeId} = getParams(req,'/store/:storeId/sizes/:sizeId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {name,value} = data
            await DBUpdateRecord('size','SET name = ? , value = ? WHERE id = ? AND storeId = ?',name,value,sizeId,storeId)
            res.writeHead(HttpStatusCode.Accepted);
            res.write("SIZE UPDATED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("UPDATE SIZE CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
        }
    })
}

export const deleteSize = async (req:IncomingMessage,res:ServerResponse)=>{
        try {
            const {storeId,sizeId} = getParams(req,'/store/:storeId/sizes/:sizeId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const requestHeader = await DBDeleteRecord('size',' WHERE id = ? AND storeId = ?',sizeId,storeId) as resultHeader
            if(requestHeader.affectedRows < 1){
                res.writeHead(HttpStatusCode.NotFound);
                res.write("SIZE NOT FOUND")
                res.end()
                return
            }
            res.writeHead(HttpStatusCode.Accepted);
            res.write("SIZE DELETED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("DELETE SIZE CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
            res.end("Something went wrong")
        }
}



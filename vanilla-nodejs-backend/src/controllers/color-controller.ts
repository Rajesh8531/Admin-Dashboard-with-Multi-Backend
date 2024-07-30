import { IncomingMessage, ServerResponse } from "http";
import { getParams, getRequestBody, resultHeader } from "../utils/utils";
import { DBDeleteRecord,DBUpdateRecord, DBCreateCategory, DBgetRecord, DBgetRecords, DBcreateSize, DBcreateColor } from "../db/db-operation";
import { HttpStatusCode } from "axios";

export const colorsController = async ( req:IncomingMessage,res:ServerResponse) => {

    switch (req.method){
        case "POST" : createColor(req,res)
                        return;
        case "GET" : getColors(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
    }
}

const createColor = async (req:IncomingMessage,res:ServerResponse) => {
    getRequestBody(req,async(data)=>{
        try {
            const {storeId} = getParams(req,'/store/:storeId/colors')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {name,value} = data
            await DBcreateColor(name,value,storeId) 
            res.writeHead(HttpStatusCode.Created);
            res.write("COLOR CREATED")
            res.end()
            return
        } catch (error) {
            console.log("CREATE COLOR CONTROLLER",error)
            res.writeHead(HttpStatusCode.BadRequest);
            res.end("SOMETHING WENT WRONG")
        }
        
    })
}

const getColors = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId} = getParams(req,'/store/:storeId/colors')
        const color = await DBgetRecords('color','storeId=?',storeId)
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify(color))
        res.end()

    } catch (error) {
        console.log("GET COLORS CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const uniqueColorController = async (req:IncomingMessage,res:ServerResponse)=>{
    switch (req.method){
        case "PATCH" : updateColor(req,res)
                        return;
        case "GET" : getColor(req,res)
                        return;
        case "DELETE" : deleteColor(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
            return;
    }
}

export const getColor = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId,colorId} = getParams(req,'/store/:storeId/colors/:colorId')
        if(colorId == 'new'){
            res.writeHead(HttpStatusCode.Ok);
            res.end()
            return;
        }
        const color = await DBgetRecord('color','id = ? AND storeId = ? ',colorId,storeId)
        if(!color){
            res.writeHead(HttpStatusCode.NotFound)
            res.end("NO COLOR FOUND")
            return
        }
        res.writeHead(HttpStatusCode.Accepted);
        const products = await DBgetRecords('product','colorId = ? ',colorId)
        res.write(JSON.stringify({...color,products}))
        res.end()
        return;
        
    } catch (error) {
        console.log("GET COLOR CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const updateColor = async (req:IncomingMessage,res:ServerResponse)=>{
    getRequestBody(req,async ( data )=>{
        try {
            const {storeId,colorId} = getParams(req,'/store/:storeId/colors/:colorId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {name,value} = data
            await DBUpdateRecord('color','SET name = ? , value = ? WHERE id = ? AND storeId = ?',name,value,colorId,storeId)
            res.writeHead(HttpStatusCode.Accepted);
            res.write("COLOR UPDATED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("UPDATE COLOR CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
        }
    })
}

export const deleteColor = async (req:IncomingMessage,res:ServerResponse)=>{
        try {
            const {storeId,colorId} = getParams(req,'/store/:storeId/colors/:colorId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const requestHeader = await DBDeleteRecord('size',' WHERE id = ? AND storeId = ?',colorId,storeId) as resultHeader
            if(requestHeader.affectedRows < 1){
                res.writeHead(HttpStatusCode.NotFound);
                res.write("COLOR NOT FOUND")
                res.end()
                return
            }
            res.writeHead(HttpStatusCode.Accepted);
            res.write("COLOR DELETED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("DELETE COLOR CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
            res.end("Something went wrong")
        }
}



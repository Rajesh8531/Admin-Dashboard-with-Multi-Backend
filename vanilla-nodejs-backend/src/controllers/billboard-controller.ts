import { IncomingMessage, ServerResponse } from "http";
import { getParams, getRequestBody, resultHeader } from "../utils/utils";
import { DBDeleteRecord,DBUpdateRecord, DBcreateBillboard, DBgetRecords, DBgetRecord } from "../db/db-operation";
import { HttpStatusCode } from "axios";

export const billboardsController = async ( req:IncomingMessage,res:ServerResponse) => {

    switch (req.method){
        case "POST" : createBillboard(req,res)
                        return;
        case "GET" : getBillboards(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
    }
}

const createBillboard = async (req:IncomingMessage,res:ServerResponse) => {
    getRequestBody(req,async(data)=>{
        try {
            const {storeId} = getParams(req,'/store/:storeId/billboards')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {label,imageUrl} = data
            await DBcreateBillboard(label,imageUrl,storeId) 
            
            res.writeHead(HttpStatusCode.Created);
            res.write("BILLBOARD CREATED")
            res.end()
            return
        } catch (error) {
            console.log("CREATE BILLBOARD CONTROLLER",error)
            res.writeHead(HttpStatusCode.BadRequest);
            res.end("SOMETHING WENT WRONG")
        }
        
    })
}

const getBillboards = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId} = getParams(req,'/store/:storeId/billboards')
        const stores = await DBgetRecords('billboard','storeId=?',storeId)
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify(stores))
        res.end()

    } catch (error) {
        console.log("GET BILLBOARDS CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const uniqueBillboardController = async (req:IncomingMessage,res:ServerResponse)=>{
    switch (req.method){
        case "PATCH" : updateBillboard(req,res)
                        return;
        case "GET" : getBillboard(req,res)
                        return;
        case "DELETE" : deleteBillboard(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
            return;
    }
}

export const getBillboard = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId,billboardId} = getParams(req,'/store/:storeId/billboards/:billboardId')
        if(billboardId == 'new'){
            res.writeHead(HttpStatusCode.Ok);
            res.end()
            return;
        }
        const billboard = await DBgetRecord('billboard','id = ? AND storeId = ? ',billboardId,storeId)
        if(!billboard){
            res.writeHead(HttpStatusCode.NotFound)
            res.end("NO BILLBOARD FOUND")
            return
        }
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify(billboard))
        res.end()
        return;
        
    } catch (error) {
        console.log("GET BILLBOARD CONTROLLER")
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const updateBillboard = async (req:IncomingMessage,res:ServerResponse)=>{
    getRequestBody(req,async ( data )=>{
        try {
            const {storeId,billboardId} = getParams(req,'/store/:storeId/billboards/:billboardId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {label,imageUrl} = data
            await DBUpdateRecord('billboard','SET label = ? , imageUrl = ? WHERE id = ? AND storeId = ?',label,imageUrl,billboardId,storeId)
            res.writeHead(HttpStatusCode.Accepted);
            res.write("BILLBOARD UPDATED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("UPDATE BILLBOARD CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
        }
    })
}

export const deleteBillboard = async (req:IncomingMessage,res:ServerResponse)=>{
    getRequestBody(req,async ( data )=>{
        try {
            const {storeId,billboardId} = getParams(req,'/store/:storeId/billboards/:billboardId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const requestHeader = await DBDeleteRecord('billboard',' WHERE id = ? AND storeId = ?',billboardId,storeId) as resultHeader
            if(requestHeader.affectedRows < 1){
                res.writeHead(HttpStatusCode.NotFound);
                res.write("BILLBOARD NOT FOUND")
                res.end()
                return
            }
            res.writeHead(HttpStatusCode.Accepted);
            res.write("BILLBOARD DELETED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("DELETE BILLBOARD CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
            res.end("Something went wrong")
        }
    })
}



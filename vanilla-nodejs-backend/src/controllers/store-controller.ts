import { IncomingMessage, ServerResponse } from "http";
import { getParams, getRequestBody } from "../utils/utils";
import { DBcreateStore, DBDeleteRecord,DBUpdateRecord, DBgetRecord, DBgetRecords } from "../db/db-operation";
import { HttpStatusCode } from "axios";
import { Store } from "../type";

export const storesController = async ( req:IncomingMessage,res:ServerResponse) => {

    switch (req.method){
        case "POST" : createStore(req,res)
                        return;
        case "GET" : getStores(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
    }
}

const createStore = async (req:IncomingMessage,res:ServerResponse) => {
    getRequestBody(req,async(data)=>{
        try {
        const {name} = data
        const userId = (req as any).userId as string
        if(!userId){
            res.writeHead(HttpStatusCode.Unauthorized)
            res.end("UnAuthenticated")
            return
        }
        const store = await DBcreateStore(name,userId);
        res.writeHead(HttpStatusCode.Created);
        res.write("STORE CREATED")
        res.end()
        return
        } catch (error) {
            console.log("CREATE STORE CONTROLLER",error)
            res.writeHead(HttpStatusCode.BadRequest);
            res.end("SOMETHING WENT WRONG")
        }
        
    })
}

const getStores = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const userId = (req as any).userId as string
        const stores = await DBgetRecords('store','userId=?',userId)
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify(stores))
        res.end()

    } catch (error) {
        console.log("GET STORES CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const uniqueStoreController = async (req:IncomingMessage,res:ServerResponse)=>{
    switch (req.method){
        case "PATCH" : updateStore(req,res)
                        return;
        case "GET" : getStore(req,res)
                        return;
        case "DELETE" : deleteStore(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
            return;
    }
}

export const getStore = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId} = getParams(req,'/store/:storeId')
        const store = await DBgetRecord('store','id = ?  ',storeId) as Store
        if(!store){
            res.writeHead(HttpStatusCode.NotFound)
            res.end("NO STORE FOUND")
            return
        }
        const billboards = await DBgetRecords("billboard","storeId = ?",store.id)
        const categories = await DBgetRecords("category",'storeId = ?',store.id)
        const sizes = await DBgetRecords('size','storeId=?',storeId)
        const colors = await DBgetRecords('color','storeId=?',storeId)
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify({...store,billboards,categories,sizes,colors}))
        res.end()
        return;
        
    } catch (error) {
        console.log("GET STORE CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}

export const updateStore = async (req:IncomingMessage,res:ServerResponse)=>{
    getRequestBody(req,async ( data )=>{
        try {
            const {storeId} = getParams(req,'/store/:storeId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            const {name} = data
            const resultHeader = await DBUpdateRecord('store','SET name = ? WHERE id = ? AND userId = ?',name,storeId,userId) as {affectedRows:number}
            if(resultHeader.affectedRows < 1){
                res.writeHead(HttpStatusCode.NotFound)
                res.end("NO STORE FOUND")
                return
            }
            res.writeHead(HttpStatusCode.Accepted);
            res.write("STORE UPDATED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("UPDATE STORE CONTROLLER",error)
            res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
        }
    })
}

export const deleteStore = async (req:IncomingMessage,res:ServerResponse)=>{
    getRequestBody(req,async ( data )=>{
        try {
            const {storeId} = getParams(req,'/store/:storeId')
            const userId = (req as any).userId
            const store = await DBgetRecord("store","id = ? AND userId = ?",storeId,userId)
            if(!store){
                res.writeHead(HttpStatusCode.Unauthorized)
                res.end("UnAuthorized")
                return
            }
            
            const resultHeader = await DBDeleteRecord('store',' WHERE id = ? AND userId = ?',storeId,userId) as {affectedRows:number}
            if(resultHeader.affectedRows < 1){
                res.writeHead(HttpStatusCode.NotFound)
                res.end("NO STORE FOUND")
                return
            }
            res.writeHead(HttpStatusCode.Accepted);
            res.write("STORE DELETED SUCCESSFULLY")
            res.end()
        } catch (error) {
            console.log("DELETE STORE CONTROLLER")
            res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
        }
    })
}



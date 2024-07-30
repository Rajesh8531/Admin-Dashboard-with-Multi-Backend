import { IncomingMessage, ServerResponse } from "http";
import { getParams } from "../utils/utils";
import {  DBgetRecord, DBgetRecords } from "../db/db-operation";
import { HttpStatusCode } from "axios";
import { Order } from "../type";


export const ordersController = async ( req:IncomingMessage,res:ServerResponse) => {

    switch (req.method){
        case "GET" : getOrders(req,res)
                        return;
        default :
            res.writeHead(HttpStatusCode.MethodNotAllowed)
            res.write("METHOD NOT ALLOWED")
            res.end()
    }
}


const getOrders = async (req:IncomingMessage,res:ServerResponse)=>{
    try {
        const {storeId} = getParams(req,'/store/:storeId/orders')
        let fullOrders :any[]= []
        const orders = await DBgetRecords('orders','storeId=?',storeId) as Order[]
        for (let order of orders){
            let orderItems = await DBgetRecord('orderitem','orderId = ?',order.id)
            fullOrders.push({...order,orderItems})
        }
        res.writeHead(HttpStatusCode.Accepted);
        res.write(JSON.stringify(fullOrders))
        res.end()

    } catch (error) {
        console.log("GET ORDERS CONTROLLER",error)
        res.writeHead(HttpStatusCode.InternalServerError)
        res.end("Something went wrong")
    }
}
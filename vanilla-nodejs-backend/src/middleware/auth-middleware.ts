import { IncomingMessage, ServerResponse } from "http";
import jwt from 'jsonwebtoken'
import dotenv from 'dotenv'
dotenv.config()

export const addAuthMiddleware = (req:IncomingMessage,res:ServerResponse)=>{
        const token = req.headers.authorization?.split(' ')[1]
        if(!token){return}        
            const claims = jwt.verify(token as string,process.env.JWT_SECRET as string) as {id:string}
            (req as any).userId = claims?.id
}
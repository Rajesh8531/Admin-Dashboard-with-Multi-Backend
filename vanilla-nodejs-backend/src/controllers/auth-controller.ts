import { IncomingMessage, ServerResponse } from "http";
import { getRequestBody } from "../utils/utils";
import { HttpStatusCode } from "axios";
import { createUser, getUser } from "../db/db-operation";
import { compare,hash } from "bcrypt";
import { User } from "../type";
import {config} from 'dotenv'
import jwt from 'jsonwebtoken'

config()

export const signUp = async (req:IncomingMessage,res:ServerResponse)=>{
    if(req.method=="POST"){
        getRequestBody(req, async(data) => {
           
            try {
                const {name,email,password} =await data
            const existingUser = await getUser(" email = ? ",email) as User
            if(existingUser?.email){
                res.writeHead(HttpStatusCode.NotAcceptable)
                res.end("User Already Existing")
                return
            }
            const hashedPassword = await hash(password,12) as string
            const {id} = await createUser(name,email,hashedPassword)
            const token = jwt.sign({id,email},process.env.JWT_SECRET as string,{})
            res.writeHead(HttpStatusCode.Created)
            res.write(JSON.stringify({id,token,email,name}))
            res.end()
                
            } catch (error) {

                console.log("SIGNUP CONTROLLER",error)
                res.writeHead(HttpStatusCode.BadRequest)
                res.end("INTERNAL SERVER ERROR")    
            }      
            })
    }else {
        res.writeHead(HttpStatusCode.MethodNotAllowed)
        res.end("METHOD NOT ALLOWED")
    }
        
}

export const signIn = async (req:IncomingMessage,res:ServerResponse)=>{
    if(req.method == 'POST'){
        getRequestBody(req,async (data)=>{
            if (req.method == "POST"){
                
            }
            try {
            const {email,password} = await data
            
            const existingUser = await getUser("email = ?",email) as User
            if(!existingUser?.email){
                res.writeHead(HttpStatusCode.BadRequest)
                res.end("User Not Existing")
                return
            }

            const isPasswordCorrect = await compare(password,existingUser.password as string)
            if(!isPasswordCorrect){
                res.writeHead(HttpStatusCode.BadRequest)
                res.end("Invalid Credentials")
                return
            }
            const token = jwt.sign({id:existingUser.id,email},process.env.JWT_SECRET as string,{})
            res.writeHead(HttpStatusCode.Accepted)
            res.write(JSON.stringify({id:existingUser.id,email,token,name:existingUser.name}))
            res.end()
                
            } catch (error) {
                console.log("SIGNIN CONTROLLER",error)
                res.writeHead(HttpStatusCode.BadRequest)
                res.end("INTERNAL SERVER ERROR")
            }
        })
    }else {
        res.writeHead(HttpStatusCode.MethodNotAllowed)
        res.end("METHOD NOT ALLOWED")
    }    
}
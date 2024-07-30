
import { IncomingMessage } from 'http'
import {v4} from 'uuid'
import url from 'url'

export const generateUUID = ()=>{
    return v4()
}

export const getRequestBody = async ( req:IncomingMessage,callback:(res:any)=>void) => {
    let body = '';
    req.on('data', chunk => {
        body += chunk.toString();
    });
    req.on('end', () => {
        callback(JSON.parse(body))
    });     
}

export const getParams = (req:IncomingMessage,pattern:string)=>{
    const parsedUrl = url.parse(req.url as string,true)
    const pathname = parsedUrl.pathname as string
    const params : {[key:string]:string} = {}

    const parts = pathname.split('/')
    const patterns = pattern.split('/')
    for(var i = 0;i< parts.length;i++){
        if(patterns[i].startsWith(':')){
            const str = patterns[i].slice(1) as string
           params[str] = parts[i]
        }
    }
    return params
}

export const isUrlMatching = (pathname:string,pattern:string)=>{
    var pattern1 = ""
    var pattern2 = ""
    const parts = pathname.split('/').filter((pattern:string)=>pattern !== '')
    const patterns = pattern.split('/').filter((pattern:string)=>pattern !== '')
    if(parts.length !== patterns.length){
        return false;
    }

    for (var i = 0 ; i < parts.length;i++){
        if(!patterns[i].startsWith(':')){
            pattern1 += parts[i]
            pattern2 += patterns[i]
        }
    }
    return pattern1 == pattern2
}

export type resultHeader = {
    affectedRows : number
}

export interface customRequest extends IncomingMessage {
    userId : string
}
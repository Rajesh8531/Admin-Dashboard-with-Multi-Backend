import chalk from 'chalk'
import http,{IncomingMessage,ServerResponse} from 'http'
import dotenv from 'dotenv'
import { createTables } from './db/table-creation'
import { router } from './router/route'
import cors from 'cors'
import { addAuthMiddleware } from './middleware/auth-middleware'

dotenv.config()

const port = process.env.PORT || 8000

const corsHandler = cors({
  origin: '*', 
  methods: 'GET,POST,OPTIONS,PUT,PATCH,DELETE', 
  allowedHeaders: 'X-Requested-With,content-type,Authorization',
  credentials: true  
});

const requestListener =async (req:IncomingMessage,res:ServerResponse)=>{

    addAuthMiddleware(req,res)
    corsHandler(req,res,async ()=>{
        await createTables()
        await router(req,res)
    })

}

const server = http.createServer(requestListener)


server.listen(port,()=>{
    console.log(chalk.greenBright(`Server is running on port ${port}`))
})
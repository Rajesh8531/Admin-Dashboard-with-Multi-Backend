import { generateUUID } from "../utils/utils"
import { getConnection } from "./connection"

export const createUser = async (name :string,email:string,password:string) => {
    const connection = await getConnection()
    const id = generateUUID()
    const query = `INSERT INTO users (id,name,email,password) VALUES (? , ?, ?,?)`
    connection.query(query,[id,name,email,password],(error,result)=>{
    })
    connection.end()
    return {id,name,email}
}

export const DBcreateStore = async (name :string,userId:string) => {
    const connection = await getConnection()
    const id = generateUUID()
    const query = `INSERT INTO store (id,name,userId) VALUES (? , ?, ?)`
    return new Promise((resolve,reject)=>{
        connection.query(query,[id,name,userId],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(result)
            }
        })
    })
}

export const DBcreateBillboard = async (label :string,imageUrl:string,storeId:string) => {
    const connection = await getConnection()
    const id = generateUUID()
    const query = `INSERT INTO billboard (id,label,imageUrl,storeId) VALUES (? , ?, ? ,?)`
   
    return new Promise((resolve,reject)=>{
        connection.query(query,[id,label,imageUrl,storeId],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(result)
            }
        })
    })
}

export const DBCreateCategory = async (name:string,billboardId:string,storeId:string)=>{
    const connection = await getConnection()
    const id = generateUUID()
    const query = `INSERT INTO category (id,name,billboardId,storeId) VALUES (? , ?, ? ,?)`
   
    return new Promise((resolve,reject)=>{
        connection.query(query,[id,name,billboardId,storeId],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(result)
            }
        })
    })
}

export const DBcreateSize = async (name:string,value:string,storeId:string)=>{
    const connection = await getConnection()
    const id = generateUUID()
    const query = `INSERT INTO size (id,name,value,storeId) VALUES (? , ?, ? ,?)`
   
    return new Promise((resolve,reject)=>{
        connection.query(query,[id,name,value,storeId],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(result)
            }
        })
    })
}

export const DBcreateColor = async (name:string,value:string,storeId:string)=>{
    const connection = await getConnection()
    const id = generateUUID()
    const query = `INSERT INTO color (id,name,value,storeId) VALUES (? , ?, ? ,?)`
   
    return new Promise((resolve,reject)=>{
        connection.query(query,[id,name,value,storeId],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(result)
            }
        })
    })
}

export const DBcreateProduct = async (name:string,price:string,storeId:string,categoryId:string,isFeatured:boolean,isArchived:boolean,sizeId:string,colorId:string)=>{
    const connection = await getConnection()
    const id = generateUUID()
    const query = `INSERT INTO product (id,name,price,storeId,categoryId,isFeatured,isArchived,sizeId,colorId) VALUES (?,?,?,?,?,?,?,?,?)`
   
    return new Promise((resolve,reject)=>{
        connection.query(query,[id,name,price,storeId,categoryId,isFeatured,isArchived,sizeId,colorId],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(id)
            }
        })
    })
}

export const DBcreateImage = async (url:string,productId:string)=>{
    const connection = await getConnection()
    const id = generateUUID()
    const query = `INSERT INTO image (id,url,productId) VALUES (? , ?, ? )`
   
    return new Promise((resolve,reject)=>{
        connection.query(query,[id,url,productId],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(result)
            }
        })
    })
}

export const getUser = async (filter:string,...values :any[])=>{
    const connection = await getConnection()
    
    return new Promise((resolve,reject)=>{
        const query = `SELECT * from users WHERE ${filter}`
        connection.query(query,[...values],(error,result)=>{
            if(error){
                reject(error)
            }else {
                resolve(result[0])
            }
        })
        connection.end()
    }) 
}


export const DBgetRecord = async (tableName : string,filter:string,...values : any[])=>{
    const connection = await getConnection()

    return new Promise((resolve,reject)=>{
        const query = `SELECT * FROM ${tableName} WHERE ${filter} ORDER BY createdAt DESC`
        connection.query(query,[...values],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(result[0])
            }
        })
        connection.end()
    })
}

export const DBgetRecords = async (tableName : string,filter:string,...values : any[])=>{
    const connection = await getConnection()

    return new Promise((resolve,reject)=>{
        const query = `SELECT * FROM ${tableName} WHERE ${filter} ORDER BY createdAt DESC`
        connection.query(query,[...values],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(result)
            }
        })
        connection.end()
    })
}

export const DBUpdateRecord = async (tableName:string,filter:string,...values:any[])=>{
    const connection = await getConnection()
    const query = `UPDATE ${tableName} ${filter}`
    return new Promise((resolve,reject)=>{
        connection.query(query,[...values],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(result)
            }
        })
        connection.end()
    })
}

export const DBDeleteRecord = async (tableName:string,filter:string,...values :any[])=>{
    const connection = await getConnection()
    const query = `DELETE FROM ${tableName} ${filter}`
    return new Promise((resolve,reject)=>{
        connection.query(query,[...values],(error,result)=>{
            if(error){
                reject(error)
            } else {
                resolve(result)
            }
        })
        connection.end()
    })
}
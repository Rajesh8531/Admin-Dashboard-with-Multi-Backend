import mysql from 'mysql2'
import dotenv from 'dotenv'
dotenv.config()

export const getConnection = async () => {
    const connection = await mysql.createConnection({
        host : process.env.HOST,
        user : process.env.USER,
        password : process.env.PASSWORD,
        database : process.env.DATABASE
    })

    return connection;
}
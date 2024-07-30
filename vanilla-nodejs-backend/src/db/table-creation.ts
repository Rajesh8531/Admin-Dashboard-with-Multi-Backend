import { Connection } from "mysql2";
import { getConnection } from "./connection"
import chalk from "chalk";

export const createTables = async ()=>{
    const connection = await getConnection()
    await createUserTable(connection);
    await createStoreTable(connection);
    await createBillboardTable(connection);
    await createCategoryTable(connection)
    await createSizeTable(connection)
    await createColorTable(connection)
    await createProductTable(connection)
    await createImageTable(connection)
    await createOrderTable(connection)
    await createOrderItemTable(connection)
    connection.end()
}

const createUserTable =async (connection : Connection)=>{

    const query = `
        CREATE TABLE IF NOT EXISTS users(
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,

		PRIMARY KEY(id)
		)`

    connection.query(query,(error,result)=>{
        if(error){
            console.log(chalk.redBright(error.message))
        }
    })
}

async function createStoreTable(connection:Connection) {
	const query = `
		CREATE TABLE IF NOT EXISTS Store (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		userId CHAR(36) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (userId) REFERENCES users(id)
		)
	`
	connection.query(query,(error,result)=>{
        if(error){
            console.log(chalk.redBright(error.message))
        } 
    })
}

async function createBillboardTable(connection:Connection) {
	const query = `
		CREATE TABLE IF NOT EXISTS Billboard (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		label VARCHAR(50) NOT NULL,
		storeId CHAR(36) NOT NULL,
		imageUrl VARCHAR(500),
		PRIMARY KEY (id),
		FOREIGN KEY (storeId) REFERENCES Store(id)
		)
	`
	connection.query(query,(error,result)=>{
        if(error){
            console.log(chalk.redBright(error.message))
        } 
    })
}

async function createCategoryTable(connection:Connection) {
	const query = `
		CREATE TABLE IF NOT EXISTS Category (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		storeId CHAR(36) NOT NULL,
		billboardId CHAR(46) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (storeId) REFERENCES Store(id),
		FOREIGN KEY (billboardId) REFERENCES Billboard(id)
		)
	`
	connection.query(query,(error,result)=>{
        if(error){
            console.log(chalk.redBright(error.message))
        } 
    })
}

async function createSizeTable(connection:Connection) {
	const query = `
		CREATE TABLE IF NOT EXISTS Size (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		value VARCHAR(50) NOT NULL,
		storeId CHAR(36) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (storeId) REFERENCES Store(id)
		)
	`
	connection.query(query,(error,result)=>{
        if(error){
            console.log(chalk.redBright(error.message))
        } 
    })
}

async function createColorTable(connection:Connection) {
	const query = `
		CREATE TABLE IF NOT EXISTS Color (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		value VARCHAR(50) NOT NULL,
		storeId CHAR(36) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (storeId) REFERENCES Store(id)
		)
	`
	connection.query(query,(error,result)=>{
        if(error){
            console.log(chalk.redBright(error.message))
        } 
    })
}

async function createProductTable(connection:Connection) {
	const query = `
		CREATE TABLE IF NOT EXISTS Product (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		name VARCHAR(50) NOT NULL,
		price VARCHAR(20) NOT NULL,
		isFeatured BOOLEAN NOT NULL DEFAULT FALSE,
		isArchived BOOLEAN NOT NULL DEFAULT FALSE,
		storeId CHAR(36) NOT NULL,
		categoryId CHAR(36) NOT NULL,
		sizeId CHAR(36) NOT NULL,
		colorId CHAR(36) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (storeId) REFERENCES Store(id),
		FOREIGN KEY (categoryId) REFERENCES Category(id),
		FOREIGN KEY (sizeId) REFERENCES Size(id),
		FOREIGN KEY (colorId) REFERENCES Color(id)
		)
	`
	connection.query(query,(error,result)=>{
        if(error){
            console.log(chalk.redBright(error.message))
        } 
    })
}

async function createImageTable(connection : Connection)  {
	const query = `
		CREATE TABLE IF NOT EXISTS Image (
		id CHAR(36) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		url VARCHAR(50) NOT NULL,
		productId CHAR(36) NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (productId) REFERENCES Product(id)
		)
	`
	connection.query(query,(error,result)=>{
        if(error){
            console.log(chalk.redBright(error.message))
        } 
    })
}

async function createOrderTable(connection:Connection) {
	const query = `
	CREATE TABLE IF NOT EXISTS Orders (
	id CHAR(36) NOT NULL,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	isPaid BOOLEAN NOT NULL DEFAULT FALSE,
	phone VARCHAR(20) NOT NULL,
	address VARCHAR(100) NOT NULL,
	storeId CHAR(36) NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (storeId) REFERENCES Store(id)
	)
    `
    connection.query(query,(error,result)=>{
    if(error){
        console.log(chalk.redBright(error.message))
    } 
})
}

async function createOrderItemTable(connection:Connection) {
	const query = `
	CREATE TABLE IF NOT EXISTS OrderItem (
	id CHAR(36) NOT NULL,
	productId CHAR(36) NOT NULL,
	orderId CHAR(36) NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (productId) REFERENCES Product(id),
	FOREIGN KEY (orderId) REFERENCES Orders(id)
	)
    `
    connection.query(query,(error,result)=>{
        if(error){
            console.log(chalk.redBright(error.message))
        } 
    })
}
package controller

import (
	"connection-to-mongo/project/db"
	"connection-to-mongo/project/types"
	"connection-to-mongo/project/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v79"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var CheckoutController = http.HandlerFunc(checkoutCreateController)

func checkoutCreateController(w http.ResponseWriter, r *http.Request) {
	var productIds []primitive.ObjectID
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid storeId"))
		return
	}

	utils.ParseJSON(r, &productIds)

	if len(productIds) == 0 {
		utils.ResponseError(w, http.StatusNoContent, fmt.Errorf("no productIds found"))
		return
	}

	var products []types.Product
	filter := bson.M{"_id": bson.M{"$in": productIds}}
	cursor, err := db.Sizes.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var product types.Product
		err := cursor.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	var lineItems []*stripe.CheckoutSessionLineItemParams

	for _, product := range products {
		price, _ := strconv.ParseInt(product.Price, 10, 64)
		lineItem := &stripe.CheckoutSessionLineItemParams{
			Quantity: stripe.Int64(1),
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("usd"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String(product.Name),
				},
				UnitAmount: stripe.Int64(price * 100),
			},
		}

		lineItems = append(lineItems, lineItem)
	}

	fmt.Println(lineItems)

	var order types.Order
	order.ID = primitive.NewObjectID()

	for _, productId := range productIds {
		var orderItem types.OrderItem
		orderItem.OrderId = order.ID
		orderItem.ProductId = productId
		_, err = db.OrderItems.InsertOne(context.Background(), orderItem)
		if err != nil {
			log.Fatal(err)
		}
	}

	order.IsPaid = false
	order.StoreId = storeId

	_, err = db.Orders.InsertOne(context.Background(), order)

	if err != nil {
		log.Fatal(err)
	}

	// envFile, _ := godotenv.Read(".env")
	// front_end_url := []byte(envFile["FRONTEND_STORE_URL"])

	// successUrl := fmt.Sprintf("$s/cart?success=1", front_end_url)
	// cancelUrl := fmt.Sprintf("$s/cart?canceled=1", front_end_url)

	stripe.Key = ""

	// params := &stripe.CheckoutSessionParams{
	// 	Mode:                     stripe.String("payment"),
	// 	BillingAddressCollection: stripe.String("required"),
	// 	PhoneNumberCollection:    &stripe.CheckoutSessionPhoneNumberCollectionParams{Enabled: stripe.Bool(true)},
	// 	LineItems:                lineItems,
	// 	SuccessURL:               stripe.String(successUrl),
	// 	CancelURL:                stripe.String(cancelUrl),
	// }

	// result, _ := session.New(params)

	utils.ResponseJSON(w, http.StatusAccepted, "")
}

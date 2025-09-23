package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct{
	ID					primitive.ObjectID  		`bson:"_id"`
	Invoice_id
	Order_id
	Payment_method
	Payment_status
	Payment_due_date
	Created_at
	updated_at
}
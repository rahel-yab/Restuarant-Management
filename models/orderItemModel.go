package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct{
	ID              primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Quantity        *string              `bson:"quantity" json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	Unit_price      *float64             `json:"unit_price" validated:"required"`
	Food_id         primitive.ObjectID   `bson:"food_id" json:"food_id"`
	Created_at      time.Time            `bson:"created_at" json:"created_at"`
	Updated_at      time.Time            `bson:"updated_at" json:"updated_at"`
	Order_item_id   string               `bson:"order_item_id" json:"order_item_id"`
	Order_id        *string              `bson:"order_id" json:"order_id"`
}
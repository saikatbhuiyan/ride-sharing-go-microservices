package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID            primitive.ObjectID `bson:"user_id" json:"user_id"`
	PackageSlug       string             `bson:"package_slug" json:"package_slug"` // ex: van, sedan, suv, bike
	TotalPriceInCents int64              `bson:"total_price_in_cents" json:"total_price_in_cents"`
}

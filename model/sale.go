package model

// Sale representa a estrutura de um documento de venda no MongoDB
type Sale struct {
    Product  string  `json:"product" bson:"product"`
    Category string  `json:"category" bson:"category"`
    Amount   float64 `json:"amount" bson:"amount"`
    Date     string  `json:"date" bson:"date"`
}
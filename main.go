package main

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/labstack/echo/v4"
    "go-mongodb-aggregation/controller"
    "go-mongodb-aggregation/repository"
    "go-mongodb-aggregation/service"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        panic(err)
    }
    defer client.Disconnect(ctx)

    collection := client.Database("sales_db").Collection("sales")
    
    // Cria um índice único no campo "product"
    indexModel := mongo.IndexModel{
        Keys:    bson.D{{Key: "product", Value: 1}}, // Campo que será único
        Options: options.Index().SetUnique(true),     // Define o índice como único
    }

    _, err = collection.Indexes().CreateOne(ctx, indexModel)
    if err != nil {
        log.Fatal(err)
    }
    
    saleRepo := repository.NewSaleRepository(collection)
    saleService := service.NewSaleService(saleRepo)
    saleController := controller.NewSaleController(saleService)

    e := echo.New()

    e.POST("/sales", saleController.CreateSale)
    e.PUT("/sales/:product", saleController.UpdateSale)
    e.DELETE("/sales/:product", saleController.DeleteSale)
    e.GET("/sales", saleController.ListAllSales)
    e.GET("/sales/aggregate", saleController.AggregateSales)
    e.GET("/sales/aggregateByDate", saleController.AggregateSalesByDate)

    e.Start(":8080")
    
}

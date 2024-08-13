package main

import (
    "context"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
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

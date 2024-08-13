package controller

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go-mongodb-aggregation/model"
    "go-mongodb-aggregation/service"
)

type SaleController struct {
    service *service.SaleService
}

func NewSaleController(service *service.SaleService) *SaleController {
    return &SaleController{service: service}
}

// Criação de uma venda
func (c *SaleController) CreateSale(ctx echo.Context) error {
    var sale model.Sale
    if err := ctx.Bind(&sale); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    if err := c.service.CreateSale(ctx.Request().Context(), sale); err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert sale"})
    }

    return ctx.JSON(http.StatusOK, map[string]string{"status": "sale inserted"})
}

// Atualização de uma venda
func (c *SaleController) UpdateSale(ctx echo.Context) error {
    oldProductID := ctx.Param("product")
    var updatedSale model.Sale
    if err := ctx.Bind(&updatedSale); err != nil {
        return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    // Service method to handle the update
    err := c.service.UpdateSale(ctx.Request().Context(), oldProductID, updatedSale)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return ctx.JSON(http.StatusOK, echo.Map{"status": "sale updated"})
}

// Exclusão de uma venda
func (c *SaleController) DeleteSale(ctx echo.Context) error {
    product := ctx.Param("product")
    if product == "" {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Product parameter is required"})
    }

    err := c.service.DeleteSale(ctx.Request().Context(), product)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Sale not found"})
        }
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete sale"})
    }

    return ctx.JSON(http.StatusOK, map[string]string{"status": "sale deleted"})
}

// Listar todas as vendas
func (c *SaleController) ListAllSales(ctx echo.Context) error {
    sales, err := c.service.GetAllSales(ctx.Request().Context())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve sales"})
    }

    return ctx.JSON(http.StatusOK, sales)
}

// Agregar vendas
func (c *SaleController) AggregateSales(ctx echo.Context) error {
    category := ctx.QueryParam("category")
    if category == "" {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Category query parameter is required"})
    }

    pipeline := mongo.Pipeline{
        {
            {Key: "$match", Value: bson.D{{Key: "category", Value: category}}},
        },
        {
            {Key: "$group", Value: bson.D{
                {Key: "_id", Value: "$category"},
                {Key: "totalAmount", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
                {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
            }},
        },
    }

    results, err := c.service.AggregateSales(ctx.Request().Context(), pipeline)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to aggregate sales"})
    }

    return ctx.JSON(http.StatusOK, results)
}

// Agregar vendas por data
func (c *SaleController) AggregateSalesByDate(ctx echo.Context) error {
    startDate := ctx.QueryParam("startDate")
    endDate := ctx.QueryParam("endDate")

    if startDate == "" || endDate == "" {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "StartDate and endDate query parameters are required"})
    }

    result, err := c.service.AggregateSalesByDate(ctx.Request().Context(), startDate, endDate)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to aggregate sales by date"})
    }

    totalSales := result["totalSales"]

    // Retorna o total de vendas no período especificado
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "startDate":   startDate,
        "endDate":     endDate,
        "totalSales":  totalSales,
    })
}

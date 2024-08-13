package service

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go-mongodb-aggregation/model"
    "go-mongodb-aggregation/repository"
)

type SaleService struct {
    repo *repository.SaleRepository
}

func NewSaleService(repo *repository.SaleRepository) *SaleService {
    return &SaleService{repo: repo}
}

func (s *SaleService) CreateSale(ctx context.Context, sale model.Sale) error {
    return s.repo.CreateSale(ctx, sale)
}

func (s *SaleService) UpdateSale(ctx context.Context, oldProductID string, updatedSale model.Sale) error {
    return s.repo.UpdateSale(ctx, oldProductID, updatedSale)
}

func (s *SaleService) DeleteSale(ctx context.Context, product string) error {
    return s.repo.DeleteSale(ctx, product)
}

func (s *SaleService) GetAllSales(ctx context.Context) ([]model.Sale, error) {
    return s.repo.GetAllSales(ctx)
}

func (s *SaleService) AggregateSales(ctx context.Context, pipeline mongo.Pipeline) ([]bson.M, error) {
    cursor, err := s.repo.AggregateSales(ctx, pipeline)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []bson.M
    if err := cursor.All(ctx, &results); err != nil {
        return nil, err
    }

    return results, nil
}

func (s *SaleService) AggregateSalesByDate(ctx context.Context, startDate string, endDate string) (bson.M, error) {
    pipeline := mongo.Pipeline{
        {
            {Key: "$match", Value: bson.D{
                {Key: "date", Value: bson.D{{Key: "$gte", Value: startDate}}},
                {Key: "date", Value: bson.D{{Key: "$lte", Value: endDate}}},
            }},
        },
        {
            {Key: "$group", Value: bson.D{
                {Key: "_id", Value: nil},
                {Key: "totalSales", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
            }},
        },
    }

    cursor, err := s.repo.AggregateSales(ctx, pipeline)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var result bson.M
    if cursor.Next(ctx) {
        if err := cursor.Decode(&result); err != nil {
            return nil, err
        }
    }

    return result, nil
}

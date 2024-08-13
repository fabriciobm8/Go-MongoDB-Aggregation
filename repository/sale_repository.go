package repository

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go-mongodb-aggregation/model"
)

type SaleRepository struct {
    collection *mongo.Collection
}

func NewSaleRepository(collection *mongo.Collection) *SaleRepository {
    return &SaleRepository{collection: collection}
}

func (r *SaleRepository) CreateSale(ctx context.Context, sale model.Sale) error {
    _, err := r.collection.InsertOne(ctx, sale)
    return err
}

func (r *SaleRepository) UpdateSale(ctx context.Context, product string, updatedSale model.Sale) error {
    filter := bson.D{{Key: "product", Value: product}}
    update := bson.D{{Key: "$set", Value: updatedSale}}

    result, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    if result.MatchedCount == 0 {
        return mongo.ErrNoDocuments
    }

    return nil
}

func (r *SaleRepository) DeleteSale(ctx context.Context, product string) error {
    filter := bson.D{{Key: "product", Value: product}}

    result, err := r.collection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }

    if result.DeletedCount == 0 {
        return mongo.ErrNoDocuments
    }

    return nil
}

func (r *SaleRepository) GetAllSales(ctx context.Context) ([]model.Sale, error) {
    cursor, err := r.collection.Find(ctx, bson.D{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var sales []model.Sale
    if err = cursor.All(ctx, &sales); err != nil {
        return nil, err
    }

    return sales, nil
}

func (r *SaleRepository) AggregateSales(ctx context.Context, pipeline mongo.Pipeline) (*mongo.Cursor, error) {
    return r.collection.Aggregate(ctx, pipeline)
}

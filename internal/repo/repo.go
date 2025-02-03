package repo

import (
	"github.com/AzlinII/receipt-processor-challenge/internal/customerror"
	"github.com/google/uuid"
)

type PointsDB struct {
	db map[string]int
}

func NewPointsDB() PointsDB {
	return PointsDB{
		db: map[string]int{},
	}
}

func (p PointsDB) GetPoints(id string) (int, error) {
	points, ok := p.db[id]
	if !ok {
		return 0, customerror.NewReceiptNotFoundError()
	}
	return points, nil
}

func (p PointsDB) SavePoints(points int) string {
	id := uuid.New().String()
	p.db[id] = points
	return id
}

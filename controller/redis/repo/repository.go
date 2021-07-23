package repo

import "github.com/iikmaulana/gateway/models"

type CompositeRepository interface {
	GetRedisByID(models.CompositeID) (models.Composite, error)
}

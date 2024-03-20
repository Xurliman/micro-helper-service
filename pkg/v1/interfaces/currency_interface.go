package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type CurrencyRepoInterface interface {
	Create(currency models.Currency) (models.Currency, error)
	Get(id int64) (models.Currency, error)
	Update(currency models.Currency) (models.Currency, error)
	GetByCode(code string) (models.Currency, error)
	Delete(id int64) error
}

type CurrencyCaseInterface interface {
	Create(currency models.Currency) (models.Currency, error)
	Get(id int64) (models.Currency, error)
	Update(currency models.Currency) (models.Currency, error)
	Delete(id int64) error
}

package v1

import (
	"github.com/Xurliman/banking-microservice/internal/models"
)

type BankInterface interface {
	Create(bank models.Bank) error
	Get(id string) error
	Update(bank models.Bank) error
	Delete(id string) error
}

type BankBranchInterface interface {
	Create(branch models.BankBranch) error
	Get(id string) error
	Update(branch models.BankBranch) error
	Delete(id string) error
}

type ClientTypeClassifierInterface interface {
	Create(classifier models.ClientTypeClassifier) error
	Get(id string) error
	Update(classifier models.ClientTypeClassifier) error
	Delete(id string) error
}

type CountryInterface interface {
	Create(country models.Country) error
	Get(id string) error
	Update(country models.Country) error
	Delete(id string) error
}

type CurrencyInterface interface {
	Create(currency models.Currency) error
	Get(id string) error
	Update(currency models.Currency) error
	Delete(id string) error
}

type DirectOrganInterface interface {
	Create(organ models.DirectOrgan) error
	Get(id string) error
	Update(organ models.DirectOrgan) error
	Delete(id string) error
}

type DistrictInterface interface {
	Create(district models.District) error
	Get(id string) error
	Update(district models.District) error
	Delete(id string) error
}

type NationalEconomySectorInterface interface {
	Create(sector models.NationalEconomySector) error
	Get(id string) error
	Update(sector models.NationalEconomySector) error
	Delete(id string) error
}

type PaymentTypeInterface interface {
	Create(paymentType models.PaymentType) error
	Get(id string) error
	Update(paymentType models.PaymentType) error
	Delete(id string) error
}

type RegionInterface interface {
	Create(region models.Region) error
	Get(id string) error
	Update(region models.Region) error
	Delete(id string) error
}
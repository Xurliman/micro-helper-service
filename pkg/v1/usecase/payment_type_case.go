package usecase

import (
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"gorm.io/gorm"
	"strconv"
)

type PaymentTypeCase struct {
	repo interfaces.PaymentTypeRepoInterface
}

func (paymentTypeCase PaymentTypeCase) Create(paymentType models.PaymentType) (models.PaymentType, error) {
	if _, err := paymentTypeCase.repo.GetByCode(strconv.FormatInt(paymentType.Code, 10)); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.PaymentType{}, errors.New("the code has already been taken")
	}

	return paymentTypeCase.repo.Create(paymentType)
}

func (paymentTypeCase PaymentTypeCase) Get(id int64) (models.PaymentType, error) {
	var paymentType models.PaymentType
	var err error

	paymentType, err = paymentTypeCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.PaymentType{}, errors.New("no such paymentType with the id supplied")
		}

		return models.PaymentType{}, err
	}
	return paymentType, nil
}

func (paymentTypeCase PaymentTypeCase) Update(updatePaymentType models.PaymentType) (models.PaymentType, error) {
	var paymentType models.PaymentType
	var err error

	paymentType, err = paymentTypeCase.repo.Get(int64(updatePaymentType.ID))
	if err != nil {
		return models.PaymentType{}, err
	}

	if paymentType.Code != updatePaymentType.Code {
		return models.PaymentType{}, errors.New("code cannot be changed")
	}

	paymentType, err = paymentTypeCase.repo.Update(updatePaymentType)
	if err != nil {
		return models.PaymentType{}, err
	}
	return paymentType, err
}

func (paymentTypeCase PaymentTypeCase) Delete(id int64) error {
	var err error

	_, err = paymentTypeCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = paymentTypeCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"shellrean.id/belajar-golang-rest-api/domain"
	"shellrean.id/belajar-golang-rest-api/dto"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

// Index implements domain.CustomerService.

func NewCustomer(customerRespository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRespository,
	}
}

func (c *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll((ctx))
	if err != nil {
		return nil, err
	}
	var customerData []dto.CustomerData
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}
	return customerData, nil
}

func (c *customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:        uuid.NewString(),
		Name:      req.Name,
		Code:      req.Code,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	return c.customerRepository.Save(ctx, &customer)
}

func (c *customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	// konversi ID ke UUID

	uuidID, err := uuid.Parse(req.ID)
	if err != nil {
		return errors.New("ID tidak valid")
	}
	persisted, err := c.customerRepository.FindByID(ctx, uuidID.String())
	if err != nil {
		return err
	}

	// ID customer tidak ditemukan
	if persisted.ID == "" {
		return errors.New("data customers tidak ditemukan")
	}

	persisted.Code = req.Code
	persisted.Name = req.Name
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return c.customerRepository.Update(ctx, &persisted)
}

func (c *customerService) Delete(ctx context.Context, id string) error {

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID tidak valid")
	}

	exist, err := c.customerRepository.FindByID(ctx, uuidID.String())
	if err != nil {
		return err

	}

	if exist.ID == "" {
		return errors.New("data customers tidak ditemukan")
	}
	return c.customerRepository.Delete(ctx, uuidID.String())
}

func (c *customerService) ShowByID(ctx context.Context, id string) (dto.CustomerData, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return dto.CustomerData{}, errors.New("ID tidak valid")
	}
	exist, err := c.customerRepository.FindByID(ctx, uuidID.String())
	if err != nil {
		return dto.CustomerData{}, fmt.Errorf("gagal mengambil data customer : %w", err)
	}

	if exist.ID == "" {
		return dto.CustomerData{}, errors.New("data customers tidak ditemukan")
	}
	return dto.CustomerData{
		ID:   exist.ID,
		Code: exist.Code,
		Name: exist.Name,
	}, nil
}

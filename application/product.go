package application

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"log"

	"github.com/asaskevich/govalidator"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"in(enabled|disabled)"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
	IsValid() (bool, error)
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

func (p *Product) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		log.Printf("Product validation error: %v", err)
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("product price must be greater than 0 to enable")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New("product price must be 0 to disable")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func NewProduct() *Product {
	return &Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}
}

func NewProductWithStatus(id string, name string, price float64, status string) *Product {
	return &Product{
		ID:     id,
		Name:   name,
		Price:  price,
		Status: status,
	}
}

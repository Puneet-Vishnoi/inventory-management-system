package store

import (
	"errors"
	"fmt"

	"github.com/Puneet-Vishnoi/advance-smart-inventory-management-system/product"
)

type Manager struct {
	Name string
	Age  int
	MId  string
}
type Store struct {
	StoreID  string
	Location string
	Products map[string]*ProductInventory
	Managers []Manager
}

type ProductInventory struct {
	Product  *product.Product
	Quantity int
}

func NewStore(str map[string]interface{}) (*Store, error) {
	s := &Store{
		Products: make(map[string]*ProductInventory),
		Managers: make([]Manager, 0),
	}

	if id, ok := str["StoreID"]; !ok {
		return nil, errors.New("store id is not defined")
	} else {
		s.StoreID = id.(string)
	}
	if loc, ok := str["Location"]; ok {
		s.Location = loc.(string)
	}
	return s, nil
}

func (s *Store) UpdateInventoryInStore(qnt int, prodId string) error {
	if _, ok := s.Products[prodId]; ok {
		if qnt>=0{
			s.Products[prodId].Quantity += qnt
		}else if (-qnt)<=s.Products[prodId].Quantity{
			s.Products[prodId].Quantity += qnt
		}else if (-qnt) > s.Products[prodId].Quantity{
			return errors.New("product quantity is not available")
		}
		fmt.Println("product inventory updated")
		return nil
	} else {
		return errors.New("product not found in the store")
	}
}

func (s *Store) AddProductInStore(prod *product.Product) error {
	if _, ok := s.Products[prod.ProductID]; !ok {
		s.Products[prod.ProductID] = &ProductInventory{
			Product:  prod,
			Quantity: 0,
		}
		fmt.Println("product added")
		return nil
	} else {
		return errors.New("product alreay available in the store")
	}
}

func (s *Store) RemoveProductInStore(prod *product.Product) error {
	if _, ok := s.Products[prod.ProductID]; ok {
		delete(s.Products, prod.ProductID)
		fmt.Println("product remove")
		return nil
	} else {
		return errors.New("product not available in the store")
	}
}


func (s *Store)ViewInventoryInStore(){
	for _, iProd := range s.Products{
		fmt.Println("quantity: ",iProd.Quantity, " ", "product: ", iProd.Product)
	}
}
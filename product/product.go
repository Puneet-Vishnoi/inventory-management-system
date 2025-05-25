// product/product.go
package product

import (
	"errors"
	"fmt"
)

type Product struct {
	ProductID       string  `json:"product_id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	UnitPrice       float64 `json:"unit_price"`
	SupplierDetails string  `json:"supplier_details"`
}

func NewProduct(prod map[string]interface{}) (*Product, error) {
	p := &Product{}

	// ProductID is required
	if id, ok := prod["ProductId"]; !ok {
		return nil, errors.New("product id is required")
	} else {
		if idStr, ok := id.(string); ok && idStr != "" {
			p.ProductID = idStr
		} else {
			return nil, errors.New("product id must be a non-empty string")
		}
	}

	// Name is required
	if name, ok := prod["Name"]; ok {
		if nameStr, ok := name.(string); ok && nameStr != "" {
			p.Name = nameStr
		} else {
			return nil, errors.New("name must be a non-empty string")
		}
	} else {
		return nil, errors.New("product name is required")
	}

	// Optional fields with validation
	if description, ok := prod["Description"]; ok {
		if descStr, ok := description.(string); ok {
			p.Description = descStr
		}
	}

	if unitPrice, ok := prod["UnitPrice"]; ok {
		if price, ok := unitPrice.(float64); ok && price >= 0 {
			p.UnitPrice = price
		} else {
			return nil, errors.New("unit price must be a non-negative number")
		}
	}

	if supplierDetails, ok := prod["SupplierDetails"]; ok {
		if supplierStr, ok := supplierDetails.(string); ok {
			p.SupplierDetails = supplierStr
		}
	}

	return p, nil
}

func (p *Product) String() string {
	return fmt.Sprintf("Product{ID: %s, Name: %s, Price: %.2f}", 
		p.ProductID, p.Name, p.UnitPrice)
}
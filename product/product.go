package product

import "errors"

type Product struct {
	ProductID       string
	Name            string
	Description     string
	UnitPrice       float64
	SupplierDetails string
}

func NewProduct(prod map[string]interface{}) (*Product, error) {
	p := &Product{}

	if id, ok := prod["ProductId"]; !ok {
		return nil, errors.New("product id is not defined")
	}else{
		p.ProductID = id.(string)
	}
	if name, ok := prod["Name"]; ok {
		p.Name = name.(string)
	}
	if description, ok := prod["Description"]; ok {
		p.Description = description.(string)
	}
	if unitPrice, ok := prod["UnitPrice"]; ok {
		p.UnitPrice = unitPrice.(float64)
	}
	if supplierDetails, ok := prod["SupplierDetails"]; ok {
		p.SupplierDetails = supplierDetails.(string)
	}

	return p, nil
}
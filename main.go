package main

import (
	"fmt"

	inventory "github.com/Puneet-Vishnoi/advance-smart-inventory-management-system/inventory-management-system"
	"github.com/Puneet-Vishnoi/advance-smart-inventory-management-system/product"
	"github.com/Puneet-Vishnoi/advance-smart-inventory-management-system/store"
)

func main() {
	instance := inventory.NewInventoryManagementSystem()
	p1 := make(map[string]interface{})
	p1["ProductId"] = "p1"
	p1["Name"] = "product1"
	p1["Description"] = "this is product 1"
	p1["UnitPrice"] = 800.1
	p1["SupplierDetails"] = "xyz" 

	prod1, err := product.NewProduct(p1)
	if err != nil{
		fmt.Println(err)
	}

	s1 := make(map[string]interface{})
	s1["StoreID"] = "s1"
	s1["Location"] = "abc"

	str, err := store.NewStore(s1)
	if err != nil{
		fmt.Println(err)
	}

	instance.AddStore(str)
	instance.AddProduct(str.StoreID, prod1)
	instance.ViewInventory()
	instance.UpdateInventory(str.StoreID, prod1.ProductID, 5)
	instance.ViewInventory()
	err =instance.UpdateInventory(str.StoreID, prod1.ProductID, -7)
	if err != nil{
		fmt.Println(err)
	}
	instance.ViewInventory()
	instance.RemoveProduct(str.StoreID, prod1)
	instance.ViewInventory()
}

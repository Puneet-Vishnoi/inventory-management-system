package inventorymanagementsystem

import (
	"errors"
	"sync"

	"github.com/Puneet-Vishnoi/advance-smart-inventory-management-system/product"
	"github.com/Puneet-Vishnoi/advance-smart-inventory-management-system/store"
)

type InventoryManagement struct {
	mu sync.Mutex
	Stores map[string]*store.Store
}

var instance *InventoryManagement
var once sync.Once

func NewInventoryManagementSystem()*InventoryManagement{
	once.Do(func() {
		instance = &InventoryManagement{
			Stores: make(map[string]*store.Store),
		}
	})
	return instance
}


func (i *InventoryManagement)UpdateInventory(storeId string, prodId string, qnt int)error{
	defer i.mu.Unlock()
	i.mu.Lock()

	if str, ok := i.Stores[storeId]; ok{
		return str.UpdateInventoryInStore(qnt, prodId)
	}else{
		return errors.New("please add store info first")
	}
}

func (i *InventoryManagement)AddStore(str *store.Store)error{
	defer i.mu.Unlock()
	i.mu.Lock()
	
	if _, ok := i.Stores[str.StoreID]; !ok{
		i.Stores[str.StoreID] = str
		return nil
	}else{
		return errors.New("store alreay there")
	}
}

func (i *InventoryManagement)AddProduct(storeId string, prod *product.Product)error{
	defer i.mu.Unlock()
	i.mu.Lock()

	if str, ok := i.Stores[storeId]; ok{
		return str.AddProductInStore(prod)
	}else{
		return errors.New("please add store info first")
	}
}

func (i *InventoryManagement)RemoveProduct(storeId string, prod *product.Product)error{
	defer i.mu.Unlock()
	i.mu.Lock()

	if str, ok := i.Stores[storeId]; ok{
		return str.RemoveProductInStore(prod)
	}else{
		return errors.New("please add store info first")
	}
}


func (i *InventoryManagement)ViewInventory(){
	for _, iStr := range i.Stores{
		iStr.ViewInventoryInStore()
	}
}
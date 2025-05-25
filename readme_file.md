# Advanced Smart Inventory Management System

A robust, thread-safe inventory management system built in Go that allows you to manage multiple stores, products, and their inventories efficiently.

## 🚀 Features

- **Multi-Store Management**: Manage multiple stores with different locations
- **Product Management**: Add, remove, and track products across stores
- **Inventory Control**: Update inventory quantities with automatic validation
- **Thread Safety**: Concurrent-safe operations using mutex locks
- **Manager Support**: Add and manage store managers
- **Real-time Inventory Views**: Display current inventory status across all stores
- **Singleton Pattern**: Single instance of inventory management system
- **Input Validation**: Comprehensive validation for all operations

## 🏗️ Architecture

The system follows a modular architecture with clear separation of concerns:

```
├── inventorymanagementsystem/
│   └── inventory.go          # Core inventory management logic
├── product/
│   └── product.go           # Product entity and operations
├── store/
│   └── store.go            # Store entity and inventory operations
└── main.go                 # Application entry point and demo
```

## 📦 Installation

1. Clone the repository:
```bash
git clone https://github.com/Puneet-Vishnoi/advance-smart-inventory-management-system.git
cd advance-smart-inventory-management-system
```

2. Initialize Go module:
```bash
go mod init advance-smart-inventory-management-system
```

3. Run the application:
```bash
go run main.go
```

## 🔧 Usage

### Basic Operations

#### 1. Initialize the System
```go
inventorySystem := inventory.NewInventoryManagementSystem()
```

#### 2. Create and Add a Store
```go
storeData := map[string]interface{}{
    "StoreID":  "store1",
    "Location": "New York, NY",
}
store, err := store.NewStore(storeData)
if err != nil {
    log.Fatal(err)
}
inventorySystem.AddStore(store)
```

#### 3. Create and Add a Product
```go
productData := map[string]interface{}{
    "ProductId":       "prod1",
    "Name":           "Laptop",
    "Description":    "High-performance laptop",
    "UnitPrice":      999.99,
    "SupplierDetails": "TechCorp Inc.",
}
product, err := product.NewProduct(productData)
if err != nil {
    log.Fatal(err)
}
inventorySystem.AddProduct("store1", product)
```

#### 4. Update Inventory
```go
// Add 50 units to inventory
err := inventorySystem.UpdateInventory("store1", "prod1", 50)

// Remove 10 units from inventory
err := inventorySystem.UpdateInventory("store1", "prod1", -10)
```

#### 5. View Inventory
```go
inventorySystem.ViewInventory()
```

## 📋 API Reference

### InventoryManagement

| Method | Description | Parameters | Returns |
|--------|-------------|------------|---------|
| `NewInventoryManagementSystem()` | Creates singleton instance | None | `*InventoryManagement` |
| `AddStore(store)` | Adds a new store | `*store.Store` | `error` |
| `AddProduct(storeId, product)` | Adds product to store | `string, *product.Product` | `error` |
| `UpdateInventory(storeId, prodId, quantity)` | Updates inventory quantity | `string, string, int` | `error` |
| `RemoveProduct(storeId, product)` | Removes product from store | `string, *product.Product` | `error` |
| `ViewInventory()` | Displays all store inventories | None | None |
| `GetStore(storeId)` | Retrieves store by ID | `string` | `*store.Store, error` |
| `RemoveStore(storeId)` | Removes store | `string` | `error` |
| `GetStoreCount()` | Returns total stores | None | `int` |

### Product

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `ProductID` | `string` | Yes | Unique product identifier |
| `Name` | `string` | Yes | Product name |
| `Description` | `string` | No | Product description |
| `UnitPrice` | `float64` | No | Price per unit |
| `SupplierDetails` | `string` | No | Supplier information |

### Store

| Field | Type | Description |
|-------|------|-------------|
| `StoreID` | `string` | Unique store identifier |
| `Location` | `string` | Store location |
| `Products` | `map[string]*ProductInventory` | Product inventory mapping |
| `Managers` | `[]Manager` | Store managers |

## 🛡️ Thread Safety

The system is designed to be thread-safe with:
- **Mutex locks** for write operations
- **Read-write locks** for optimized concurrent reads
- **Proper lock ordering** to prevent deadlocks
- **Deferred unlocks** for guaranteed cleanup

## ✅ Error Handling

Comprehensive error handling includes:
- Input validation (nil checks, empty strings)
- Type assertion safety
- Inventory availability checks
- Duplicate prevention
- Descriptive error messages

## 📝 Example Output

```
=== Advanced Smart Inventory Management System ===

Store store1 added successfully
Store store2 added successfully
Product prod1 added to store store1
Product prod2 added to store store1

=== Initial Inventory ===

========== INVENTORY OVERVIEW ==========

=== Store store1 Inventory ===
Location: New York, NY
Product: Laptop | Quantity: 0 | Price: $999.99
Product: Wireless Mouse | Quantity: 0 | Price: $29.99
========================

=== Updating Inventory ===
Product prod1 inventory updated. New quantity: 50
Product prod2 inventory updated. New quantity: 30

=== After Inventory Updates ===

========== INVENTORY OVERVIEW ==========

=== Store store1 Inventory ===
Location: New York, NY
Product: Laptop | Quantity: 50 | Price: $999.99
Product: Wireless Mouse | Quantity: 30 | Price: $29.99
========================
```

## 🧪 Testing

Run the application to see a comprehensive demo:
- Creates sample products and stores
- Demonstrates all major operations
- Shows error handling for edge cases
- Displays formatted inventory reports

## 🔮 Future Enhancements

- [ ] Database integration for persistence
- [ ] REST API endpoints
- [ ] Product categories and filtering
- [ ] Low stock alerts
- [ ] Inventory history tracking
- [ ] Multi-location transfers
- [ ] Barcode scanning support
- [ ] Reporting and analytics
- [ ] User authentication and authorization

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 👨‍💻 Author

**Puneet Vishnoi**
- GitHub: [@Puneet-Vishnoi](https://github.com/Puneet-Vishnoi)

## 📞 Support

For support, please open an issue in the GitHub repository or contact the maintainer.

---

⭐ Star this repository if you find it helpful!
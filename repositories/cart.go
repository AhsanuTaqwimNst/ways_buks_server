package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCarts() ([]models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	CreateCart(Cart models.Cart) (models.Cart, error)
	UpdateCart(Cart models.Cart) (models.Cart, error)
	DeleteCart(Cart models.Cart) (models.Cart, error)
	CreateTransactionID(transaction models.Transaction) (models.Transaction, error)
	FindToppingsID(ToppingID []int) ([]models.Topping, error)
	FindCartsTransaction(TrxID int) ([]models.Cart, error)
	GetTransactionID() (models.Transaction, error)
	GetIDTransaction() (models.Transaction, error)
}

// type repository struct {
// 	db *gorm.DB
// }

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Preload("Topping").Preload("Transaction").Find(&carts).Error

	return carts, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").Preload("Topping").Preload("Transaction").First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Preload("Product").Preload("Topping").Preload("Transaction").Create(&cart).Error

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

func (r *repository) CreateTransactionID(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) FindToppingsID(ToppingID []int) ([]models.Topping, error) {
	var toppings []models.Topping
	err := r.db.Find(&toppings, ToppingID).Error

	return toppings, err
}

func (r *repository) FindTransactionID(TransactionID []int) ([]models.Topping, error) {
	var toppings []models.Topping
	err := r.db.Find(&toppings, TransactionID).Error

	return toppings, err
}

func (r *repository) FindCartsTransaction(TrxID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Preload("Topping").Find(&carts, "transaction_id = ?", TrxID).Error

	return carts, err
}

func (r *repository) GetTransactionID() (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Carts").Preload("Carts.Product").Preload("Carts.Topping").Find(&transaction, "status = ?", "waiting").Error
	return transaction, err
}

func (r *repository) GetIDTransaction() (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Carts").Preload("Carts.Product").Preload("Carts.Topping").Find(&transaction, "status = ?", "waiting").Error
	return transaction, err
}

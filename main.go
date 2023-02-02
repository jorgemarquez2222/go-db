package main

import (
	"fmt"
	"log"

	product "github.com/jorgemarquez2222/storage_GO/pkg/product"
	storage "github.com/jorgemarquez2222/storage_GO/storage"
)

func main() {

	fmt.Println("arr")
	storage.NewPostgresDB()

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)
	// if err := serviceProduct.Migrate(); err != nil {
	// 	log.Fatalf("product.Migrate: %v", err)
	// }

	// storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	// serviceInvoiceHeader := product.NewService(storageInvoiceHeader)
	// if err := serviceInvoiceHeader.Migrate(); err != nil {
	// 	log.Fatalf("invoiceHeader.Migrate: %v", err)
	// }

	// storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	// serviceInvoiceItem := product.NewService(storageInvoiceItem)

	// if err := serviceInvoiceItem.Migrate(); err != nil {
	// 	log.Fatalf("invoiceItem.Migrate: %v", err)
	// }

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	// m := &product.Model{
	// 	Name:         "Curso de db con GO",
	// 	Price:        50,
	// 	Observations: "on fire",
	// }
	// if err := serviceProduct.Create(m); err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(ms)
	for _, v := range ms {
		fmt.Println(*v)
	}

}

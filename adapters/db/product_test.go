package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
			"id" string, 
			"name" string, 
			"price" float, 
			"status" string
			);`
	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values ("abc", "Product test", 0.0, "disabled")`
	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	// stetament que faz com que a função passada seja a última ser rodada, após tudo acabar
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")

	require.Nil(t, err)

	require.Equal(t, "Product test", product.GetName())
	require.Equal(t, 00.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.Name = "Product test"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)

	require.Equal(t, "Product test", productResult.GetName())
	require.Equal(t, float64(25), productResult.GetPrice())
	require.Equal(t, "disabled", productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "enabled", productResult.GetStatus())

}

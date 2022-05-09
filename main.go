package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/codenotary/immudb/pkg/client"
// 	"github.com/codenotary/immudb/pkg/stdlib"
// )

// func main() {
// 	opts := client.DefaultOptions()
// 	opts.Username = "immudb"
// 	opts.Password = "immudb"
// 	opts.Database = "defaultdb"

// 	db := stdlib.OpenDB(opts)
// 	defer db.Close()

// 	_, err := db.ExecContext(context.TODO(), "CREATE TABLE myTable(id INTEGER, name VARCHAR, PRIMARY KEY id)")
// 	_, err = db.ExecContext(context.TODO(), "INSERT INTO myTable (id, name) VALUES (1, 'immu1')")
// 	rows, err := db.QueryContext(context.TODO(), "SELECT * FROM myTable")

// 	var id uint64
// 	var name string
// 	defer rows.Close()
// 	rows.Next()
// 	err = rows.Scan(&id, &name)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("id: %d\n", id)
// 	fmt.Printf("name: %s\n", name)
// }

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	immudb "github.com/codenotary/immudb/pkg/client"
	"github.com/eminetto/immudb/entity"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	opts := immudb.DefaultOptions().WithAddress("localhost").WithPort(3322)
	client := immudb.NewClient().WithOptions(opts)

	err := client.OpenSession(ctx, []byte(`immudb`), []byte(`immudb`), "defaultdb")
	if err != nil {
		log.Fatal(err)
	}

	defer client.CloseSession(ctx)

	tr := entity.Transaction{
		ID:    uuid.New(),
		Type:  entity.TransactionCredit,
		Value: 10.0,
	}
	j, err := json.Marshal(tr)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := client.Set(ctx, []byte(`credito`), j)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully committed tx %d\n", tx.Id)

	entry, err := client.Get(ctx, []byte(`credito`))
	if err != nil {
		log.Fatal(err)
	}
	var r entity.Transaction
	_ = json.Unmarshal(entry.GetValue(), &r)
	fmt.Printf("Successfully retrieved entry: %v\n", r)

}

package main

import (
	"context"
	"fmt"
	"gol/domain"
	"gol/utils"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func updateContact(db*pgx.Conn,data domain.Contact) error {
  _, err := db.Query(
    context.Background(),`
      UPDATE contact SET name = $2, email = $3, phone = $4, address = $5 WHERE id = $1
    `, data.Id, data.Name, data.Email, data.Phone, data.Address,
  )
  return err 
}
func deleteContact(db*pgx.Conn, data domain.Contact) error { 
  fmt.Print(data.Id)
  _, err := db.Query(
    context.Background(),`
    DELETE FROM contact WHERE id = $1
    `, data.Id,
  )
  if err != nil { 
    log.Fatal(err)
  }
  return err 
}

func createContact (db*pgx.Conn, data domain.Contact) error{ 
  _, err := db.Query(
    context.Background(),`
      INSERT INTO contact (name, email, phone, address) VALUES 
      ($1, $2, $3, $4)
    `, data.Name, data.Email, data.Phone, data.Address,
  )

  if err != nil { 
    log.Fatal()
  }
  return err
}

func getAllContact(db *pgx.Conn) ([]domain.Contact,error) { 
  rows, err := db.Query(
    context.Background(),`
      SELECT id, name, email, phone, address FROM contact
    `,
  )
  var result []domain.Contact
  for rows.Next() { 
    var p domain.Contact 
    err = rows.Scan(&p.Id, &p.Name, &p.Email, &p.Phone, &p.Address)
    result = append(result, p)
  }
  return result, err
}

func addContact (db *pgx.Conn) { 
  defer main()
  fmt.Print("Tambah Contact\n\n")
  var cont domain.Contact
  fmt.Print("Masukan Nama: ")
  fmt.Scan(&cont.Name)
  fmt.Print("Masukan Email: ")
  fmt.Scan(&cont.Email)
  fmt.Print("Masukan Phone: ")
  fmt.Scan(&cont.Phone)
  fmt.Print("Masukan Adress: ")
  fmt.Scan(&cont.Address)
  createContact(db, cont)  
}

func showList (db* pgx.Conn) { 
  defer main()
  fmt.Print("List Contact\n\n")
  contact,_ := getAllContact(db)
  for _, user := range contact {
    fmt.Printf("Id: %d \n",user.Id)
    fmt.Printf("  Name    : %s\n",user.Name)
    fmt.Printf("  Email   : %s\n",user.Email)
    fmt.Printf("  Phone   : %s\n",user.Phone)
    fmt.Printf("  Address : %s\n",user.Address)
  }
  fmt.Print("Enter untuk kembali...")
  fmt.Scanln()
}

func deleteList(db *pgx.Conn){ 
  defer main()
  
  fmt.Print("Delete Contact\n\n")
  contact,_ := getAllContact(db)
  for _, user := range contact {
    fmt.Printf("Id: %d \n", user.Id)
    fmt.Printf("  Name    : %s\n",user.Name)
    fmt.Printf("  Email   : %s\n",user.Email)
    fmt.Printf("  Phone   : %s\n",user.Phone)
    fmt.Printf("  Address : %s\n",user.Address)
  }
  var opt domain.Contact 
  fmt.Print("\nSelect Id: ")
  fmt.Scan(&opt.Id)

  deleteContact(db, opt)
}

func updateList(db *pgx.Conn){ 
  defer main()
  fmt.Print("Update Contact\n\n")
  contact,_ := getAllContact(db)
  for _, user := range contact {
    fmt.Printf("Id: %d \n", user.Id)
    fmt.Printf("  Name    : %s\n",user.Name)
    fmt.Printf("  Email   : %s\n",user.Email)
    fmt.Printf("  Phone   : %s\n",user.Phone)
    fmt.Printf("  Address : %s\n",user.Address)
  }
  var opt domain.Contact 
  fmt.Print("Select Id: ")
  fmt.Scan(&opt.Id)
  fmt.Print("Change Name: ")
  fmt.Scan(&opt.Name)
  fmt.Print("Change Email: ")
  fmt.Scan(&opt.Email)
  fmt.Print("Change Phone: ")
  fmt.Scan(&opt.Phone)
  fmt.Print("Change Adress: ")
  fmt.Scan(&opt.Address)
  updateContact(db, opt)
}

func main () { 
  err := godotenv.Load()
  if err != nil { 
    log.Fatal(err.Error())
  }
  db := utils.Conn()

  var opt string 
  utils.Clear()
  fmt.Print("\n1. Add Contact\n2. List Contact\n3. Delete Contact\n4. Update Contact")
  fmt.Print("\n\nChoose Options: ")
  fmt.Scan(&opt)
  switch opt { 
    case "1": 
      addContact(db)
    case "2": 
      showList(db)
    case "3": 
      deleteList(db)
    case "4":
      updateList(db)
  }
}
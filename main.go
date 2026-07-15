package main

import (
	"fmt"
	"gol/domain"
	"gol/repository"
	"gol/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func addContact (repo *repository.ContactRepository) error { 
  defer main()
  fmt.Print("Tambah Contact\n\n")
  var cont domain.Contact
  cont.Name = utils.Input("Masukan Nama : ")
  cont.Email = utils.Input("Masukan Email : ")
  cont.Phone = utils.Input("Masukan Phone : ")
  cont.Address = utils.Input("Masukan Address : ")
  return repo.Create(cont)
}

func showList (repo *repository.ContactRepository) error { 
  fmt.Print("List Contact\n\n")
  contact, _ := repo.FindAll()
  for _, user := range contact {
    fmt.Printf("Id: %d \n",user.Id)
    fmt.Printf("  Name    : %s\n",user.Name)
    fmt.Printf("  Email   : %s\n",user.Email)
    fmt.Printf("  Phone   : %s\n",user.Phone)
    fmt.Printf("  Address : %s\n",user.Address)
  }
  return nil
}

func deleteList(repo *repository.ContactRepository) error { 
  fmt.Print("Delete Contact\n\n")
  if err := showList(repo); err != nil { 
    return err 
  }
  id, err := utils.InputInt("\nSelect Id: ")
  if err != nil { 
    return err 
  }
  return repo.Delete(domain.Contact{
    Id: id,
  })
}

func updateList(repo *repository.ContactRepository) error { 
  if err := showList(repo); err != nil { 
    return err 
  }
  var cont domain.Contact 
  id, err := utils.InputInt("Select Id: ")
  if err != nil {
      return err
  }
  cont.Id = id
  cont.Name = utils.Input("Change Name : ")
  cont.Email = utils.Input("Change Email : ")
  cont.Phone = utils.Input("Change Phone : ")
  cont.Address = utils.Input("Change Address : ")
  return repo.Update(cont)
}

func main () { 
  err := godotenv.Load()
  if err != nil { 
    log.Fatal(err.Error())
  }
  db := utils.Conn()
  repo := repository.NewContactRepository(db)

  for { 
    var opt string 
    utils.Clear()
    fmt.Print("Simple CRUD CLI\n")
    fmt.Print("\n1. Add Contact\n2. List Contact\n3. Delete Contact\n4. Update Contact")
    fmt.Print("\n\nChoose Options: ")
    fmt.Scan(&opt)
    switch opt { 
      case "1": 
        addContact(repo)
      case "2": 
        showList(repo)
      case "3": 
        deleteList(repo)
      case "4":
        updateList(repo)
      case "0": 
        os.Exit(0)
    }
    fmt.Print("\nEnter...")
	  fmt.Scanln()
  }
}
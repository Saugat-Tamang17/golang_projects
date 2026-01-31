package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"net/http"
)


type Book struct{
	Author string  `json:"author"`
	Title string   `json:"title"`
	Publisher string `json:"publisher"`
}
type Repositaries struct{
	DB *gorm.DB
}

func( r *Repositaries)SetupRoutes(app *fiber.App){
	api:=app.Group(("/api"))
	api.Post("/create_books",r.createBook)
	api.Delete("/delete_book/:id",r.DeleteBook)
	api.Get("/get_books/:id",r.GetBookByID)
	api.Get("/books",r.GetBooks)
}


func (r *Repositaries)createBook(c *fiber.Ctx) error{
	book:=Book{};
	err:=c.BodyParser(&book)
	 // what this does is that it converts the json data into go structs //
if err!=nil{
	c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message":"request failed","error":err.Error()})
}
	err=r.DB.Create(&book).Error
	if err!=nil{
c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message":"couldnt create the book","error":er.Error()})
return err;
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message":"Book is created sucessfully",
		"data": book,
	})
}

func ( r *Repositaries)GetBooks(c *fiber.Ctx)error{
	bookModels:=&[]models.Books{}
	err:=r.DB.Find(bookModels).Error
	if err !=nil{
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message":"could not find the searched book"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"messaage":"the searched book has been found","data":bookModels
	})
return nil;
}

func main(){
	fmt.Println("start of the golang+postgresql brochahooooooooooooooooooo\n")
	err:=godotenv.load(".env")
	if err!=nil{
		log.Fatal(err)
	}

	db,err :=storage.NewConnection(config)
	if err!=nil{
		log.Fatal("couldnt load the data-base")
	}

	r:=Repositaries{
		DB:db,
	}
app :=Fiber.New()
r.SetupRoutes(app)
app.Listen(":8080")
}
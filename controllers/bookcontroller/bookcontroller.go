package bookcontroller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go-api-native-basic/config"
	"go-api-native-basic/helper"
	"go-api-native-basic/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	var booksResponse []models.BookResponse

	if err := config.DB.Joins("Author").Joins("Category").Find(&books).Find(&booksResponse).Error; err != nil {
		helper.Response(w, 500, "Error Table Not Found", nil)
	}
	helper.Response(w, 200, "List Books", booksResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	var author models.Author
	if err := config.DB.First(&author, book.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var category models.Category
	if err := config.DB.First(&category, book.CategoryID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Category Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success create book", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book
	var bookResponse models.BookResponse

	if err := config.DB.Joins("Author").Joins("Category").First(&book, id).First(&bookResponse, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Books Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail Book", bookResponse)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Book Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var bookPayload models.Book
	if err := json.NewDecoder(r.Body).Decode(&bookPayload); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	var author models.Author
	if bookPayload.AuthorID != 0 {
		if err := config.DB.First(&author, bookPayload.AuthorID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Author Not Found", nil)
				return
			}
			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	var category models.Category
	if bookPayload.CategoryID != 0 {
		if err := config.DB.First(&category, bookPayload.CategoryID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Category Not Found", nil)
				return
			}
			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := config.DB.Where("id = ?", id).Updates(&bookPayload).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Update Book", nil)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)
	var book models.Book
	res := config.DB.Delete(&book, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Book Not Found", nil)
		return
	}

	helper.Response(w, 200, "Success Delete Book", nil)

}

func SearchBook(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	var booksResponse []models.BookResponse

	title := r.URL.Query().Get("title")

	if title == "" {
		helper.Response(w, 404, "Please add title params!!", nil)
		return
	}

	titlePattern := "%" + title + "%"
	if err := config.DB.Where("title LIKE ?", titlePattern).Joins("Author").Joins("Category").Find(&books).Find(&booksResponse).Error; err != nil {
		helper.Response(w, 404, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "book found", booksResponse)
}

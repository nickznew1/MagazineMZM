package service

import (
	"encoding/json"
	"fmt"
	"io"
	"myWebApp/backend/internal/domain/model"
	"myWebApp/backend/internal/domain/usecase"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ItemService struct {
	useCase usecase.ItemUseCase
}

func NewItemService(c *usecase.ItemUseCase) *ItemService {
	return &ItemService{
		useCase: *c}
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(w, statusCode, map[string]string{"error": message})
}

func (h *ItemService) CreateItem(w http.ResponseWriter, r *http.Request) {
	var input model.Item
	var documents model.ItemSpecFiles
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}
	Img, imgHandler, err := r.FormFile("imgFile")

	if err != nil {
		fmt.Println("err: ", err)
		respondWithError(w, http.StatusInternalServerError, "Ошибка при получении картинки")
		return
	}
	defer Img.Close()
	imgDestination, err := os.Create("./public/images/" + imgHandler.Filename)

	if err != nil {
		fmt.Println("err: ", err)
		respondWithError(w, http.StatusInternalServerError, "Ошибка при копировании файла")
		return
	}
	defer imgDestination.Close()

	if _, err = io.Copy(imgDestination, Img); err != nil {
		fmt.Println("err: ", err)
		respondWithError(w, http.StatusInternalServerError, "Ошибка при копировании файла")
		return
	}

	Pdf, pdfHandler, err := r.FormFile("pdfFile")
	if err != nil {
		fmt.Println("err: ", err)
		respondWithError(w, http.StatusInternalServerError, "Ошибка при получении картинки")
		return
	}

	defer Img.Close()
	pdfDestination, err := os.Create("./public/documents/" + pdfHandler.Filename)
	if err != nil {
		fmt.Println("err: ", err)
		respondWithError(w, http.StatusInternalServerError, "Ошибка при копировании файла")
		return
	}
	defer pdfDestination.Close()

	if _, err = io.Copy(pdfDestination, Pdf); err != nil {
		fmt.Println("err: ", err)
		respondWithError(w, http.StatusInternalServerError, "Ошибка при копировании файла")
		return
	}

	input.Name = r.FormValue("name")
	input.Price = r.FormValue("price")
	input.ItemType = r.FormValue("item_type")
	input.ItemSecondaryType = r.FormValue("secondary_type")
	input.ItemDescription = r.FormValue("item_description")
	input.ItemShortDescription = r.FormValue("item_short_description")
	input.Article = r.FormValue("article")
	input.ItemPicture = imgHandler.Filename
	documents.SpecFileLink = pdfHandler.Filename
	documents.SpecFileName = r.FormValue("document_name")
	pdfFormat := filepath.Ext(pdfHandler.Filename)
	fmt.Println(pdfFormat)
	if pdfFormat == ".pdf" {
		documents.SpecFilePic = "logos/pdf_icon.png"
	}

	fmt.Println("CreateItem")

	newItem, _, err := h.useCase.CreateItem(input, documents)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "oshibka")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newItem)
}

func (h *ItemService) GetItemById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	queryItem, err := strconv.Atoi(idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "cant parse query to int")
		return
	}
	itemId, err := h.useCase.GetItemById(queryItem)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "id doesnt find")
		return
	}
	RespondWithJSON(w, http.StatusOK, itemId)

}

func (h *ItemService) GetItemSpecById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	specQuery, err := strconv.Atoi(idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "cant parse query to int")
		return
	}
	specId, err := h.useCase.GetSpecById(specQuery)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "id doesnt find")
		return
	}
	RespondWithJSON(w, http.StatusOK, specId)
}

func (h *ItemService) GetItemId(w http.ResponseWriter, r *http.Request) {
	var input model.Item
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "error111")
		return
	}
	itemId, err := h.useCase.GetItemId(input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}

	RespondWithJSON(w, http.StatusOK, itemId)

}

func (h *ItemService) DeleteItem(w http.ResponseWriter, r *http.Request) {
	var input model.Item
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "error")
		return
	}
	deleteItem, err := h.useCase.DeleteItem(input)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "oshibka")
		return
	}
	RespondWithJSON(w, http.StatusOK, deleteItem)
}

func (h *ItemService) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.useCase.GetAllItems()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error getting all items")
		return
	}
	respondWithJSON(w, http.StatusOK, items)
}

func (h *ItemService) ChangeVisible(w http.ResponseWriter, r *http.Request) {
	var input model.Item
	idStr := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "error")
		return
	}
	visible, err := h.useCase.ChangeVisible(input.Visible, idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}

	RespondWithJSON(w, http.StatusOK, visible.Visible)

}

func (h *ItemService) GetAllPropsName(w http.ResponseWriter, r *http.Request) {
	props, err := h.useCase.GetAllPropsName()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error getting all props")
		return
	}
	respondWithJSON(w, http.StatusOK, props.PropNameA)
}

func (h *ItemService) SetNewProps(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	idStr := chi.URLParam(r, "id")
	var input []model.ItemProp
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "error")
		return
	}
	newProps, err := h.useCase.SetPropsForItem(input, idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error")
		return
	}
	respondWithJSON(w, http.StatusCreated, newProps)

}

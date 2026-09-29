package item_transport_http

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *ItemHTTPHandler) CreateItem(
	rw http.ResponseWriter,
	r *http.Request) {

	var input model.Item
	var documents model.ItemSpecFiles

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when parse multipart form from request")
		return
	}

	Img, imgHandler, err := r.FormFile("imgFile")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"getting file using key")
		return
	}
	defer func(Img multipart.File) {
		err := Img.Close()
		if err != nil {
			responseHandler.ErrorResponse(
				err,
				"closing img file after try")
			return
		}
	}(Img)

	imgDestination, err := os.Create("./public/images/" + imgHandler.Filename)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"creating new img file on destination")
		return
	}
	defer func(imgDestination *os.File) {
		err := imgDestination.Close()
		if err != nil {
			responseHandler.ErrorResponse(
				err,
				"close destination img file")
			return
		}
	}(imgDestination)

	if _, err = io.Copy(imgDestination, Img); err != nil {
		responseHandler.ErrorResponse(
			err,
			"copy new img file to destination")
		return
	}

	Pdf, pdfHandler, err := r.FormFile("pdfFile")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"getting file from pdf key")
		return
	}

	defer func(Pdf multipart.File) {
		err := Pdf.Close()
		if err != nil {
			responseHandler.ErrorResponse(
				err,
				"closing pdf file")
			return
		}
	}(Pdf)

	pdfDestination, err := os.Create("./public/documents/" + pdfHandler.Filename)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"creating pdf file destination")
		return
	}
	defer func(pdfDestination *os.File) {
		err := pdfDestination.Close()
		if err != nil {
			responseHandler.ErrorResponse(
				err,
				"closing pdf file destination")
			return
		}
	}(pdfDestination)

	if _, err = io.Copy(pdfDestination, Pdf); err != nil {
		responseHandler.ErrorResponse(
			err,
			"copy pdf file to destination")
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

	if pdfFormat == ".pdf" {
		documents.SpecFilePic = "logos/pdf_icon.png"
	}

	newItem, _, err := h.itemService.CreateItem(ctx, input, documents)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"creating pdf file destination")
		return
	}
	responseHandler.ResponseWithJSON(http.StatusCreated, newItem)
}

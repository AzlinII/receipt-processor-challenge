package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AzlinII/receipt-processor-challenge/internal/customerror"
	"github.com/AzlinII/receipt-processor-challenge/internal/model"
)

const INVALID_RECEIPT_ERROR = "The receipt is invalid."

type PointsProcessor interface {
	Process(receipt model.Receipt) (string, error)
	GetPoints(id string) (int, error)
}

type Handler struct {
	pointsProcessor PointsProcessor
}

func NewHandler(pointsProcessor PointsProcessor) Handler {
	return Handler{
		pointsProcessor: pointsProcessor,
	}
}

func (h Handler) Init(router *http.ServeMux) {
	base := "/api/v1/receipts/"
	fmt.Println(base)
	router.HandleFunc("POST "+base+"process", h.Process)
	router.HandleFunc("GET "+base+"{id}/points", h.Points)
}

func (h Handler) Process(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, customerror.NewInvalidReceiptError())
		return
	}
	var receipt model.Receipt
	err = json.Unmarshal(body, &receipt)
	if err != nil {
		writeError(w, http.StatusBadRequest, customerror.NewInvalidReceiptError())
		return
	}
	id, err := h.pointsProcessor.Process(receipt)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	response := model.ProcessResponse{Id: id}
	writeResponse(w, response)
}

func (h Handler) Points(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	points, err := h.pointsProcessor.GetPoints(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}
	writeResponse(w, points)
}

func writeResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}

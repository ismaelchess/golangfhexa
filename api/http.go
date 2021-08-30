package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ismaelchess/golangfhexa/domain"
)

type handler struct {
	employeeService domain.Service
}

//NewHandler ...
func NewHandler(employeeService domain.Service) EmployeeHandler {

	return &handler{employeeService: employeeService}

}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := chi.URLParam(r, "code")
	p, err := h.employeeService.Find(code)

	if err != nil {

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	json.NewEncoder(w).Encode(&p)

}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	//requestBody, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	p := &domain.Employee{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = h.employeeService.Store(p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&p)

}
func (h *handler) Put(w http.ResponseWriter, r *http.Request) {
	p := &domain.Employee{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = h.employeeService.Update(p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&p)

}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	code := chi.URLParam(r, "code")
	err := h.employeeService.Delete(code)
	if err != nil {
		log.Fatal(err)
	}

}
func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	p, err := h.employeeService.FindAll()

	if err != nil {

		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&p)

}

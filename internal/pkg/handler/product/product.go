/*
 * @Author: Adrian Faisal
 * @Date: 04/09/21 1.35 PM
 */

package product

import (
	"encoding/json"
	"fmt"
	"github.com/apldex/workshop-labti/internal/pkg/httputils"
	"github.com/apldex/workshop-labti/internal/pkg/model"
	usecaseProduct "github.com/apldex/workshop-labti/internal/pkg/usecase/product"
	"io/ioutil"
	"net/http"
)

type Handler interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	uc usecaseProduct.Usecase
}

func NewHandler(uc usecaseProduct.Usecase) Handler {
	return &handler{uc: uc}
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("read request body failed: %v", err)

		httputils.RespondErrWithJSON(w, http.StatusInternalServerError, err)
		return
	}

	var product model.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		err = fmt.Errorf("unmarshal json failed: %v", err)

		httputils.RespondErrWithJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = h.uc.CreateProduct(r.Context(), &product)
	if err != nil {
		httputils.RespondErrWithJSON(w, http.StatusInternalServerError, err)
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "OK"})
}

func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {}

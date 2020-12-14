/*
 * Desafio conductor
 *
 * Api para controle de trasacoes de contas
 *
 * API version: 1.0.0
 * Contact: thiagorodriguescamara@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Account retune only one account
func (c *Controller) Account(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	acc, err := c.Db.FindAccountByID(id)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"Error": err.Error()})
	}

	json.NewEncoder(w).Encode(&acc)
}

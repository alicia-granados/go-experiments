package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWrite   http.ResponseWriter
}

// Crear respuesta por defecto
func CreateDefaulResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWrite:   rw,
		contentType: "application/json",
	}
}

// rsponder al cliente
func (resp *Response) Send() {
	//modificar el encabezado
	resp.respWrite.Header().Set("Content-Type", resp.contentType)
	resp.respWrite.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWrite, string(output))
}

// cambios para devolver la data  al cliente
func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaulResponse(rw)
	response.Data = data
	response.Send()
}

// errores en listar, eliminar o obtener un dato //metodo para responder un error
func (resp *Response) NoFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource no Found"
}

// responder error al cliente
func SendNoFound(rw http.ResponseWriter) {
	response := CreateDefaulResponse(rw)
	response.NoFound()
	response.Send()
}

// errores para ingresar o actualizar
func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity no Found"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaulResponse(rw)
	response.UnprocessableEntity()
	response.Send()
}

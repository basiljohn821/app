package controllers

import (
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/app/api/models"
	"github.com/app/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetPostExcel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	post := models.Post{}
	postReceived, err := post.FindPostByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, postReceived)
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "A2", postReceived.ID)
	f.SetCellValue("Sheet1", "B1", "Title")
	f.SetCellValue("Sheet1", "B2", postReceived.Title)
	f.SetCellValue("Sheet1", "C1", "Content")
	f.SetCellValue("Sheet1", "C2", postReceived.Content)
	f.SetActiveSheet(index)
	if err := f.SaveAs("Myexcel.xlsx"); err != nil {
		println(err.Error())
	}
}

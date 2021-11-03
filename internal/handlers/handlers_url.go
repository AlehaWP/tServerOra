package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"tServerOra/internal/models"
)

var Repo models.Repository
var tpl *template.Template

func HandlerTCPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(models.UserKey).(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// text, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	cTC := new(models.CardTC)
	cTC.DriverName = r.FormValue("driver_name")
	cTC.ModelTC = r.FormValue("model_tc")
	cTC.NumTC = r.FormValue("num_tc")
	cTC.NumPric = r.FormValue("num_pric")
	cTC.NumPlomb = r.FormValue("num_plomb")
	cTC.ContNum = r.FormValue("cont_num")
	cTC.Remark = r.FormValue("remark")
	cTC.TypeTC = r.FormValue("type_tc")
	cTC.SizeTC = r.FormValue("size_tc")

	// err = json.Unmarshal(text, cTC)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	cTC.UserID = userID
	fmt.Println(cTC)
	Repo.SaveCard(ctx, cTC)

	//	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Zaebok")
}

func HandlerCheckDBConnect(w http.ResponseWriter, r *http.Request) {
	if err := Repo.CheckDBConnection(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Connection: OK")
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//tpl.Execute(w, nil)
}

func NewHandlers(repo models.Repository) {
	//t, err := template.ParseFiles("../../html/index.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//tpl = template.Must(t, err)
	fmt.Println(tpl)
	Repo = repo
}

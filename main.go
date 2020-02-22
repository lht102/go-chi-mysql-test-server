package main

import (
    "log"
	"net/http"
    "encoding/json"

    "github.com/lht102/go-chi-mysql-test-server/model"
    "github.com/lht102/go-chi-mysql-test-server/common"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func ObjectList(w http.ResponseWriter, r *http.Request) {
    objs, err := model.GetAllObjects()
    if err != nil {
        render.Status(r, http.StatusBadRequest)
        render.JSON(w, r, common.Response{"fail to get list of Object"})
        return
    }
    render.JSON(w, r, objs)
}

func ObjectCreate(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var obj model.Object
    if err := decoder.Decode(&obj); err != nil {
        log.Println(err)
        return
    }
    err := model.CreateObject(&obj)
    if err != nil {
        log.Println(err)
        render.Status(r, http.StatusBadRequest)
        render.JSON(w, r, common.Response{"fail to create Object"})
        return
    }
    render.JSON(w, r, common.Response{"Object created"})
}

func main() {
	r := chi.NewRouter()
	r.Get("/list", ObjectList)
	r.Post("/add", ObjectCreate)
	http.ListenAndServe(":80", r)
}

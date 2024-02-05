package main

import (
	"html/template"
	"main/pkg/models"
	"main/pkg/services"
	"main/pkg/utils/logger"
	"net/http"
	"strconv"
)

type DepHanlder struct {
	Service *services.Service
}

func NewDepHandler(Service *services.Service) *DepHanlder {
	return &DepHanlder{
		Service: Service,
	}
}

func (h *DepHanlder) GetAllDeps(w http.ResponseWriter, r *http.Request) {
	deps, err := h.Service.DepService.GetAllDeps()
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return
	}

	file := "./ui/templates/deps.html"
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		http.Error(w, "Error parsing templates", 500)
		return
	}

	err = tmpl.Execute(w, deps)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		http.Error(w, "Error executing template", 500)
		return
	}
}

func (h *DepHanlder) CreateDep(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("name")
		quantity := r.FormValue("quantity")

		quantityINT, _ := strconv.Atoi(quantity)

		dep := models.Department{
			Dep_name:       name,
			Staff_quantity: quantityINT,
		}

		err := h.Service.DepService.CreateDep(dep)
		if err != nil {
			return
		}

		http.Redirect(w, r, "/deps", http.StatusSeeOther)
	} else if r.Method == "GET" {
		file := "./ui/templates/depCreate.html"
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			http.Error(w, "Error parsing templates", 500)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			logger.GetLogger().Warn(err.Error())
			http.Error(w, "Error executing template", 500)
			return
		}
	}

}

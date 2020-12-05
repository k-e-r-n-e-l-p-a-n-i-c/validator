package controllers

import (
	"fmt"

	"github.com/revel/revel"
	"k8s.io/api/admission/v1beta1"
)

type App struct {
	*revel.Controller
}

func (c App) Validate() revel.Result {
	var request v1beta1.AdmissionReview
	c.params.BindJSON(request)

	fmt.Println(request)
	return c.RenderText("ok")
}

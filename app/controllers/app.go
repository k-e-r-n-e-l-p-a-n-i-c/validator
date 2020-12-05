package controllers

import (
	"fmt"

	"github.com/revel/revel"
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type App struct {
	*revel.Controller
}

func (c App) Validate() revel.Result {
	var request v1beta1.AdmissionReview
	c.Params.BindJSON(&request)

	fmt.Println(request)

	arResponse := v1beta1.AdmissionReview{
		Response: &v1beta1.AdmissionResponse{
			Allowed: false,
			Result: &metav1.Status{
				Message: "Keep calm and not add more crap in the cluster!",
			},
		},
	}
	return c.RenderText(arResponse)
}

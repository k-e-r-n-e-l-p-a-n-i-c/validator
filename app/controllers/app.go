package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/revel/revel"
	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type App struct {
	*revel.Controller
}

func (c App) Validate() revel.Result {
	var request v1beta1.AdmissionReview
	var obj v1.Pod
	c.Params.BindJSON(&request)

	rawObject := request.Request.Object.Raw

	err := json.Unmarshal(rawObject, &obj)
	if err != nil {
		fmt.Println("Error occurred while deserializing request")
		fmt.Println("Policy check will be skipped")
		response := v1beta1.AdmissionReview{
			Response: &v1beta1.AdmissionResponse{
				Allowed: true,
			},
		}
		return c.RenderJSON(response)
	}

	if obj.Namespace == "default" {
		response := v1beta1.AdmissionReview{
			Response: &v1beta1.AdmissionResponse{
				Allowed: false,
				Result: &metav1.Status{
					Message: "Deployments in 'default' namespace is restricted by cluster admin!",
				},
			},
		}
		return c.RenderJSON(response)
	}

	fmt.Println("Pod " + obj.Name + " was allowed by validator")
	response := v1beta1.AdmissionReview{
		Response: &v1beta1.AdmissionResponse{
			Allowed: true,
		},
	}
	return c.RenderJSON(response)
}

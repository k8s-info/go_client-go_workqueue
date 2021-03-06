package main

import (
	"fmt"
	"time"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	testresourcev1beta1 "ycliu912.github.io/kubernetes-test-controller/lib/testresource/v1beta1"
)

func (c *Controller) doesCRDExist() (bool, error) {
	crd, err := c.apiextenssionsclientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(testresourcev1beta1.Name, metav1.GetOptions{})

	if err != nil {
		return false, err
	}

	// Check whether the CRD is appepted.
	for _, condition := range crd.Status.Conditions {
		if condition.Type == apiextensions.Established &&
			condition.Status == apiextensions.ConditionTrue {
			return true, nil
		}
	}

	return false, fmt.Errorf("CRD is not appepted")
}

func (c *Controller) waitCRDAccepted() error {
	err := wait.Poll(1*time.Second, 10*time.Second, func() (bool, error) {
		return c.doesCRDExist()
	})

	return err
}

// CreateCRD cretes a custom resource definition,
// named TestResource.
func (c *Controller) CreateCRD() error {
	if result, _ := c.doesCRDExist(); result {
		return nil
	}
	crd := &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: testresourcev1beta1.Name,
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group:   testresourcev1beta1.GroupName,
			Version: testresourcev1beta1.Version,
			Scope:   apiextensions.NamespaceScoped,
			Names: apiextensions.CustomResourceDefinitionNames{
				Plural:     testresourcev1beta1.Plural,
				Singular:   testresourcev1beta1.Singluar,
				Kind:       testresourcev1beta1.Kind,
				ShortNames: []string{testresourcev1beta1.ShortName},
			},
			Validation: &apiextensions.CustomResourceValidation{
				OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
					Type: "object",
					Properties: map[string]apiextensions.JSONSchemaProps{
						"spec": {
							Type: "object",
							Properties: map[string]apiextensions.JSONSchemaProps{
								"command":        {Type: "string", Pattern: "^(echo).*"},
								"customProperty": {Type: "string"},
							},
							Required: []string{"command", "customProperty"},
						},
					},
				},
			},
			AdditionalPrinterColumns: []apiextensions.CustomResourceColumnDefinition{
				{
					Name:     "command",
					Type:     "string",
					JSONPath: ".spec.command",
				},
			},
		},
	}

	_, err := c.apiextenssionsclientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)

	if err != nil {
		klog.Fatal(err.Error())
	}

	return c.waitCRDAccepted()

}

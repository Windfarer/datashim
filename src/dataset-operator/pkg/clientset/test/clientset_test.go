package test

import (
	"testing"

	"github.com/datashim-io/datashim/src/dataset-operator/pkg/apis/com/v1alpha1"
	"github.com/datashim-io/datashim/src/dataset-operator/pkg/clientset/versioned/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestFakeClient(t *testing.T) {

	client := fake.NewSimpleClientset()

	dataset := v1alpha1.Dataset{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "dataset1",
			Namespace: "default",
		},
		Spec: v1alpha1.DatasetSpec{
			Local: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			Type: "COS",
		},
		Status: v1alpha1.DatasetStatus{
			Caching: v1alpha1.DatasetStatusCondition{
				Status: "Disabled",
				Info:   "",
			},
			Provision: v1alpha1.DatasetStatusCondition{
				Status: "OK",
				Info:   "",
			},
		},
	}
	_, err := client.ComV1alpha1().Datasets("default").Create(&dataset)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	ds_list, err := client.ComV1alpha1().Datasets("default").List(metav1.ListOptions{})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(ds_list.Items) != 1 {
		t.Errorf("Unexpected List size: %d", len(ds_list.Items))
	}
}

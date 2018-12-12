package win95

import (
	win95v1alpha1 "win95-op/win95-operator/pkg/apis/win95/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func createAppLabels(cr *win95v1alpha1.Win95) map[string]string {
	return map[string]string{
		"app": cr.Name,
		"user": cr.Spec.Username,
	}
}

func createObjectMeta(cr *win95v1alpha1.Win95, res_type string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name: cr.Spec.Username + "-" + res_type,
		Namespace: cr.Namespace,
		Labels: createAppLabels(cr),
	}
}
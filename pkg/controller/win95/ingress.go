package win95

import (
	"context"
	win95v1alpha1 "win95-op/win95-operator/pkg/apis/win95/v1alpha1"

	extensions "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)


func createHostName(cr *win95v1alpha1.Win95) string {
	return cr.Spec.Username + "." + cr.Spec.IngressDomain
}

func createWin95Ingress(cr *win95v1alpha1.Win95) *extensions.Ingress {
	return &extensions.Ingress{
		ObjectMeta: createObjectMeta(cr, "ingress"),
		Spec: extensions.IngressSpec{
			Rules: []extensions.IngressRule{{
				Host: createHostName(cr),
				IngressRuleValue: extensions.IngressRuleValue{
					HTTP: &extensions.HTTPIngressRuleValue{
						Paths: []extensions.HTTPIngressPath{{
							Path: "/",
							Backend: extensions.IngressBackend{
								ServiceName: cr.Spec.Username + "-service",
								ServicePort: intstr.FromInt(6080),
							}},
						},
					},
				},
			}},
		},
	}
}

func (r *ReconcileWin95) syncWin95Ingress(cr *win95v1alpha1.Win95) error {
	ingress := createWin95Ingress(cr)

	if err := controllerutil.SetControllerReference(cr, ingress, r.scheme); err != nil {
		return err
	}

	found_ingress := &extensions.Ingress{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
			Name: ingress.Name,
			Namespace: ingress.Namespace,
		}, found_ingress)

	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), ingress)
		if err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	}

	return nil
}
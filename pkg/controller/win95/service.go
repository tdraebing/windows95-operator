package win95

import (
	"context"
	win95v1alpha1 "win95-op/win95-operator/pkg/apis/win95/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func createWin95Service(cr *win95v1alpha1.Win95) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: createObjectMeta(cr, "service"),
		Spec: corev1.ServiceSpec{
			Selector: createAppLabels(cr),
			Type: corev1.ServiceTypeNodePort,
			Ports: []corev1.ServicePort{
				{
					Name: "novnc",
					Port: 6080,
					Protocol: corev1.ProtocolTCP,
				},
			},
		},
	}
}

func (r *ReconcileWin95) syncWin95Service(cr *win95v1alpha1.Win95) error {
	service := createWin95Service(cr)

	if err := controllerutil.SetControllerReference(cr, service, r.scheme); err != nil {
		return err
	}

	found_service := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
			Name: service.Name,
			Namespace: service.Namespace,
		}, found_service)
	
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), service)
		if err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	}

	return nil
}
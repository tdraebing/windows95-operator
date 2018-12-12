package win95

import (
	"context"
	win95v1alpha1 "win95-op/win95-operator/pkg/apis/win95/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func createWin95Secret(cr *win95v1alpha1.Win95) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: createObjectMeta(cr, "secret"),
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"VNC_PWD": []byte(cr.Spec.Password),
		},
	}
}

func (r *ReconcileWin95) syncWin95Secret(cr *win95v1alpha1.Win95) error {
	secret := createWin95Secret(cr)

	if err := controllerutil.SetControllerReference(cr, secret, r.scheme); err != nil {
		return err
	}

	found_secret := &corev1.Secret{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
			Name: secret.Name,
			Namespace: secret.Namespace,
		}, found_secret)

	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), secret)
		if err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	} else {
		err = r.client.Update(context.TODO(), secret)
		if err != nil {
			return err
		}

		return nil
	}
}

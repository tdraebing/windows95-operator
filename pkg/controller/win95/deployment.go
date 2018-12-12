package win95

import (
	"context"
	win95v1alpha1 "win95-op/win95-operator/pkg/apis/win95/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func createResourceRequirements() corev1.ResourceRequirements {
	return corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			corev1.ResourceCPU: resource.MustParse("1"),
			corev1.ResourceMemory: resource.MustParse("3Gi"),
		},
	}
}

func createWin95Container(cr *win95v1alpha1.Win95) corev1.Container {
	return corev1.Container{
		Image: "tdwin/win95:latest",
		Name: "win95",
		Ports: []corev1.ContainerPort{
			{
				Name: "novnc-port",
				ContainerPort: 6080,
			},
		},
		EnvFrom: []corev1.EnvFromSource{{
			SecretRef: &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: cr.Spec.Username + "-secret",
				},
			},
		}},
		Resources: createResourceRequirements(),
	}
}

func createWin95Pod(cr *win95v1alpha1.Win95) corev1.PodTemplateSpec {
	return corev1.PodTemplateSpec{
		ObjectMeta: createObjectMeta(cr, "pod"),
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				createWin95Container(cr),
			},
			RestartPolicy: corev1.RestartPolicyAlways,
		},
	}
}

func createWin95Deployment(cr *win95v1alpha1.Win95) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: createObjectMeta(cr, "deployment"),
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: createAppLabels(cr),
			},
			Template: createWin95Pod(cr),
		},
	}
}

func (r *ReconcileWin95) syncWin95Deployment(cr *win95v1alpha1.Win95) error {
	deployment := createWin95Deployment(cr)

	if err := controllerutil.SetControllerReference(cr, deployment, r.scheme); err != nil {
		return err
	}

	found_deployment := &appsv1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
			Name: deployment.Name,
			Namespace: deployment.Namespace,
		}, found_deployment)

	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), deployment)
		if err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	}

	return nil
}
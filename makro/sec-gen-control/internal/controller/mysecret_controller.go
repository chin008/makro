/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	log "sigs.k8s.io/controller-runtime/pkg/log"

	secretsv1 "github.com/chin008/makro/sec-gen-control/api/v1"
)

// MySecretReconciler reconciles a MySecret object
type MySecretReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=secrets.chinsecretgen.com,resources=mysecrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=secrets.chinsecretgen.com,resources=mysecrets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=secrets.chinsecretgen.com,resources=mysecrets/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=create;update;patch;get;list;delete
// +kubebuilder:rbac:groups=core,resources=secrets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;create;update;patch;delete

// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.4/pkg/reconcile
func (r *MySecretReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var mySecret secretsv1.MySecret
	if err := r.Get(ctx, req.NamespacedName, &mySecret); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	secretName := fmt.Sprintf("%s-generated", mySecret.Name)
	shouldRotate := false

	// Check if secret exists
	var secret corev1.Secret
	err := r.Get(ctx, types.NamespacedName{Name: secretName, Namespace: mySecret.Namespace}, &secret)
	if err != nil {
		shouldRotate = true
	} else {
		last := mySecret.Status.LastRotated.Time
		if time.Since(last) > mySecret.Spec.RotationPeriod.Duration {
			shouldRotate = true
		}
	}

	if shouldRotate {
		password, err := generatePassword(mySecret.Spec.PasswordLength)
		if err != nil {
			return ctrl.Result{}, err
		}

		secret := corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secretName,
				Namespace: mySecret.Namespace,
				Labels:    map[string]string{"generated-by": "mysecret-controller"},
			},
			Type: corev1.SecretTypeBasicAuth,
			StringData: map[string]string{
				"username": mySecret.Spec.Username,
				"password": password,
			},
		}

		_, err = ctrl.CreateOrUpdate(ctx, r.Client, &secret, func() error {
			secret.StringData = map[string]string{
				"username": mySecret.Spec.Username,
				"password": password,
			}
			return nil
		})
		if err != nil {
			return ctrl.Result{}, err
		}

		mySecret.Status.SecretName = secretName
		mySecret.Status.LastRotated = metav1.Now()
		if err := r.Status().Update(ctx, &mySecret); err != nil {
			return ctrl.Result{}, err
		}

		logger.Info("Secret created or updated", "name", secretName)
	}

	return ctrl.Result{RequeueAfter: mySecret.Spec.RotationPeriod.Duration}, nil
}

func (r *MySecretReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&secretsv1.MySecret{}).
		Owns(&corev1.Secret{}).
		Complete(r)
}

func generatePassword(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

/*
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

package cilium

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	networkingv1alpha1 "github.com/anfernee/k8s-egress-gateway/api/v1alpha1"
	"github.com/anfernee/k8s-egress-gateway/pkg/cilium/ipmasq"
)

// EgressGatewayReconciler reconciles a EgressGateway object
type EgressGatewayReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	IPMasqAgent *ipmasq.IPMasqAgent
}

// +kubebuilder:rbac:groups=networking.x-k8s.io,resources=egressgateways,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=networking.x-k8s.io,resources=egressgateways/status,verbs=get;update;patch

func (r *EgressGatewayReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	var (
		ctx           = context.Background()
		log           = r.Log.WithValues("egressgateway", req.NamespacedName)
		egressGateway networkingv1alpha1.EgressGateway
		err           error
	)

	err = r.Get(ctx, req.NamespacedName, &egressGateway)
	if err != nil {
		return ctrl.Result{}, err
	}

	log.Info(fmt.Sprintf("%+v", egressGateway))

	return ctrl.Result{}, r.IPMasqAgent.Update(&ipmasq.Config{
		NonMasqCIDRs: egressGateway.Spec.NonMasqueradeCIDRs,
	})
}

func (r *EgressGatewayReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkingv1alpha1.EgressGateway{}).
		Complete(r)
}

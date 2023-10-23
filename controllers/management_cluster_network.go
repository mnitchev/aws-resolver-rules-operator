package controllers

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	capa "sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
	capi "sigs.k8s.io/cluster-api/api/v1beta1"
	capiannotations "sigs.k8s.io/cluster-api/util/annotations"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/aws-resolver-rules-operator/pkg/conditions"
	"github.com/aws-resolver-rules-operator/pkg/resolver"
	"github.com/aws-resolver-rules-operator/pkg/util/annotations"
)

const FinalizerManagementCluster = "network-topology.finalizers.giantswarm.io/management-cluster"

type ManagementClusterNetworkReconciler struct {
	managementCluster types.NamespacedName
	clusterClient     AWSClusterClient
	transitGateways   resolver.TransitGatewayClient
	prefixLists       resolver.PrefixListClient
}

func NewManagementClusterTransitGateway(
	managementCluster types.NamespacedName,
	client AWSClusterClient,
	transitGatewayClient resolver.TransitGatewayClient,
	prefixListClient resolver.PrefixListClient,
) *ManagementClusterNetworkReconciler {
	return &ManagementClusterNetworkReconciler{
		managementCluster: managementCluster,
		clusterClient:     client,
		transitGateways:   transitGatewayClient,
		prefixLists:       prefixListClient,
	}
}

func (r *ManagementClusterNetworkReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&capa.AWSCluster{}).
		Named("mc-transit-gateway").
		Complete(r)
}

func (r *ManagementClusterNetworkReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Reconciling")
	defer logger.Info("Done reconciling")

	cluster, err := r.clusterClient.GetAWSCluster(ctx, req.NamespacedName)
	if err != nil {
		return ctrl.Result{}, errors.WithStack(client.IgnoreNotFound(err))
	}

	if !r.isManagementCluster(cluster) {
		logger.Info("Cluster not management cluster. Skipping...")
		return ctrl.Result{}, nil
	}

	if !annotations.IsNetworkTopologyModeGiantSwarmManaged(cluster) {
		logger.Info("Cluster not using GiantSwarmManaged network topology mode. Skipping...")
		return ctrl.Result{}, nil
	}

	if capiannotations.HasPaused(cluster) {
		logger.Info("Cluster is marked as paused. Won't reconcile")
		return ctrl.Result{}, nil
	}

	defer func() {
		_ = r.clusterClient.UpdateStatus(ctx, cluster)
	}()

	if !capiconditions.Has(cluster, conditions.NetworkTopologyCondition) {
		capiconditions.MarkFalse(
			cluster,
			conditions.NetworkTopologyCondition,
			"InProgress",
			capi.ConditionSeverityInfo, "")
	}

	if !cluster.DeletionTimestamp.IsZero() {
		logger.Info("Reconciling delete")
		return r.reconcileDelete(ctx, cluster)
	}

	return r.reconcileNormal(ctx, cluster)
}

func (r *ManagementClusterNetworkReconciler) reconcileNormal(ctx context.Context, cluster *capa.AWSCluster) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	err := r.clusterClient.AddFinalizer(ctx, cluster, FinalizerManagementCluster)
	if err != nil {
		logger.Error(err, "Failed to add finalizer")
		return ctrl.Result{}, errors.WithStack(err)
	}

	err = r.applyTransitGateway(ctx, cluster)
	if err != nil {
		return ctrl.Result{}, errors.WithStack(err)
	}

	err = r.applyPrefixList(ctx, cluster)
	if err != nil {
		return ctrl.Result{}, errors.WithStack(err)
	}

	return ctrl.Result{}, nil
}

func (r *ManagementClusterNetworkReconciler) applyTransitGateway(ctx context.Context, cluster *capa.AWSCluster) error {
	logger := log.FromContext(ctx)

	id, err := r.transitGateways.Apply(ctx, cluster.Name)
	if err != nil {
		logger.Error(err, "Failed to create transit gateway")
		return errors.WithStack(err)
	}

	baseCluster := cluster.DeepCopy()
	annotations.SetNetworkTopologyTransitGateway(cluster, id)
	if cluster, err = r.clusterClient.PatchCluster(ctx, cluster, client.MergeFrom(baseCluster)); err != nil {
		logger.Error(err, "Failed to patch cluster resource with TGW ID")
		return errors.WithStack(err)
	}

	conditions.MarkReady(cluster, conditions.TransitGatewayCreated)
	return nil
}

func (r *ManagementClusterNetworkReconciler) applyPrefixList(ctx context.Context, cluster *capa.AWSCluster) error {
	logger := log.FromContext(ctx)

	id, err := r.prefixLists.Apply(ctx, cluster.Name)
	if err != nil {
		logger.Error(err, "Failed to create prefix list")
		return errors.WithStack(err)
	}

	baseCluster := cluster.DeepCopy()
	annotations.SetNetworkTopologyPrefixList(cluster, id)
	if cluster, err = r.clusterClient.PatchCluster(ctx, cluster, client.MergeFrom(baseCluster)); err != nil {
		logger.Error(err, "Failed to patch cluster resource with prefix list ID")
		return errors.WithStack(err)
	}

	conditions.MarkReady(cluster, conditions.TransitGatewayCreated)
	return nil
}

func (r *ManagementClusterNetworkReconciler) reconcileDelete(ctx context.Context, cluster *capa.AWSCluster) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	err := r.transitGateways.Delete(ctx, cluster.Name)
	if err != nil {
		logger.Error(err, "Failed to delete transit gateway")
		return ctrl.Result{}, errors.WithStack(err)
	}

	err = r.prefixLists.Delete(ctx, cluster.Name)
	if err != nil {
		logger.Error(err, "Failed to delete transit gateway")
		return ctrl.Result{}, errors.WithStack(err)
	}

	err = r.clusterClient.RemoveFinalizer(ctx, cluster, FinalizerManagementCluster)
	if err != nil {
		logger.Error(err, "Failed to delete finalizer")
		return ctrl.Result{}, errors.WithStack(err)
	}
	return ctrl.Result{}, nil
}

func (r *ManagementClusterNetworkReconciler) isManagementCluster(cluster *capa.AWSCluster) bool {
	return cluster.Name == r.managementCluster.Name &&
		cluster.Namespace == r.managementCluster.Namespace
}
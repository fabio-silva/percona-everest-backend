// Code generated by ifacemaker; DO NOT EDIT.

package client

import (
	"context"

	everestv1alpha1 "github.com/percona/everest-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// KubeClientConnector ...
type KubeClientConnector interface {
	// ClusterName returns the name of the k8s cluster.
	ClusterName() string
	// GetServerVersion returns server version.
	GetServerVersion() (*version.Info, error)
	// ListDatabaseClusters returns list of managed PCX clusters.
	ListDatabaseClusters(ctx context.Context) (*everestv1alpha1.DatabaseClusterList, error)
	// GetDatabaseCluster returns PXC clusters by provided name.
	GetDatabaseCluster(ctx context.Context, name string) (*everestv1alpha1.DatabaseCluster, error)
	// GetNodes returns list of nodes.
	GetNodes(ctx context.Context) (*corev1.NodeList, error)
	// GetPods returns list of pods.
	GetPods(ctx context.Context, namespace string, labelSelector *metav1.LabelSelector) (*corev1.PodList, error)
	// GetSecret returns secret by name.
	GetSecret(ctx context.Context, name, namespace string) (*corev1.Secret, error)
	// GetStorageClasses returns all storage classes available in the cluster.
	GetStorageClasses(ctx context.Context) (*storagev1.StorageClassList, error)
	// GetPersistentVolumes returns Persistent Volumes available in the cluster.
	GetPersistentVolumes(ctx context.Context) (*corev1.PersistentVolumeList, error)
}

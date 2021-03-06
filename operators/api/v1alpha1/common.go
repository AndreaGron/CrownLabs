package v1alpha1

// NameCreated contains information about the status of a resource created in
// the cluster (e.g. a namespace). Specifically, it contains the name of the
// resource and a flag indicating whether the creation succeeded.
type NameCreated struct {
	// The name of the considered resource.
	Name string `json:"name,omitempty"`

	// Whether the creation succeeded or not.
	Created bool `json:"created"`
}

// +kubebuilder:validation:Enum=Ok;Failed

// SubscriptionStatus is an enumeration of the different states that can be
// assumed by the subscription to a service (e.g. successful or failing).
type SubscriptionStatus string

const (
	// SubscrOk -> the subscription was successful.
	SubscrOk SubscriptionStatus = "Ok"
	// SubscrFailed -> the subscription has failed.
	SubscrFailed SubscriptionStatus = "Failed"
)

// GenericRef represents a reference to a generic Kubernetes resource,
// and it is composed of the resource name and (optionally) its namespace.
type GenericRef struct {
	// The name of the resource to be referenced.
	Name string `json:"name"`

	// The namespace containing the resource to be referenced. It should be left
	// empty in case of cluster-wide resources.
	Namespace string `json:"namespace,omitempty"`
}

// WorkspaceLabelPrefix is the prefix of a label assigned to a tenant indicating it is subscribed to a workspace.
const WorkspaceLabelPrefix = "crownlabs.polito.it/workspace-"

// TnOperatorFinalizerName is the name of the finalizer corresponding to the tenant operator.
const TnOperatorFinalizerName = "crownlabs.polito.it/tenant-operator"

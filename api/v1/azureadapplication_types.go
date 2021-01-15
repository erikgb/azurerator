package v1

// +groupName="nais.io"

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Machine readable event "Reason" fields, used for determining synchronization state.
const (
	EventSynchronized          = "Synchronized"
	EventFailedSynchronization = "FailedSynchronization"
	EventFailedStatusUpdate    = "FailedStatusUpdate"
	EventAddedFinalizer        = "AddedFinalizer"
	EventDeletedFinalizer      = "DeletedFinalizer"
	EventCreatedInAzure        = "CreatedInAzure"
	EventUpdatedInAzure        = "UpdatedInAzure"
	EventRotatedInAzure        = "RotatedInAzure"
	EventDeletedInAzure        = "DeletedInAzure"
	EventNotInTeamNamespace    = "NotInTeamNamespace"
	EventSkipped               = "Skipped"
	EventRetrying              = "Retrying"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=azureapp

// AzureAdApplication is the Schema for the AzureAdApplications API
// +kubebuilder:printcolumn:name="Secret",type=string,JSONPath=`.spec.secretName`
// +kubebuilder:printcolumn:name="ClientId",type=string,JSONPath=`.status.clientId`
// +kubebuilder:printcolumn:name="Tenant",type=string,JSONPath=`.spec.tenant`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type AzureAdApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureAdApplicationSpec   `json:"spec,omitempty"`
	Status AzureAdApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureAdApplicationList contains a list of AzureAdApplication
type AzureAdApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureAdApplication `json:"items"`
}

// AzureAdApplicationSpec defines the desired state of AzureAdApplication
type AzureAdApplicationSpec struct {
	ReplyUrls                 []AzureAdReplyUrl                 `json:"replyUrls,omitempty"`
	PreAuthorizedApplications []AzureAdPreAuthorizedApplication `json:"preAuthorizedApplications,omitempty"`
	// LogoutUrl is the URL where Azure AD sends a request to have the application clear the user's session data.
	// This is required if single sign-out should work correctly. Must start with 'https'
	LogoutUrl string `json:"logoutUrl,omitempty"`
	// SecretName is the name of the resulting Secret resource to be created
	SecretName string `json:"secretName"`
	// Tenant is an optional alias for targeting a tenant that an instance of Azurerator that processes resources for said tenant.
	// Can be omitted if only running a single instance or targeting the default tenant.
	Tenant string `json:"tenant,omitempty"`
	// Claims defines additional configuration of the emitted claims in tokens returned to the AzureAdApplication
	Claims *AzureAdClaims `json:"claims,omitempty"`
}

// AzureAdApplicationStatus defines the observed state of AzureAdApplication
type AzureAdApplicationStatus struct {
	// SynchronizationState denotes whether the provisioning of the AzureAdApplication has been successfully completed or not
	SynchronizationState string `json:"synchronizationState,omitempty"`
	// SynchronizationTime is the last time the Status subresource was updated
	SynchronizationTime *metav1.Time `json:"synchronizationTime,omitempty"`
	// SynchronizationHash is the hash of the AzureAdApplication object
	SynchronizationHash string `json:"synchronizationHash,omitempty"`
	// CorrelationId is the ID referencing the processing transaction last performed on this resource
	CorrelationId string `json:"correlationId,omitempty"`
	// PasswordKeyIds is the list of key IDs for the latest valid password credentials in use
	PasswordKeyIds []string `json:"passwordKeyIds,omitempty"`
	// CertificateKeyIds is the list of key IDs for the latest valid certificate credentials in use
	CertificateKeyIds []string `json:"certificateKeyIds,omitempty"`
	// ClientId is the Azure application client ID
	ClientId string `json:"clientId,omitempty"`
	// ObjectId is the Azure AD Application object ID
	ObjectId string `json:"objectId,omitempty"`
	// ServicePrincipalId is the Azure applications service principal object ID
	ServicePrincipalId string `json:"servicePrincipalId,omitempty"`
}

// AzureAdReplyUrl defines the valid reply URLs for callbacks after OIDC flows for this application
type AzureAdReplyUrl struct {
	Url string `json:"url,omitempty"`
}

// AzureAdPreAuthorizedApplication describes an application that are allowed to request an on-behalf-of token for this application
type AzureAdPreAuthorizedApplication struct {
	Application string `json:"application"`
	Namespace   string `json:"namespace"`
	Cluster     string `json:"cluster"`
}

type AzureAdClaims struct {
	// Extra is a list of additional claims to be mapped from an associated claim-mapping policy.
	Extra []AzureAdExtraClaim `json:"extra,omitempty"`
	// Groups is a list of Azure AD group IDs to be emitted in the 'Groups' claim.
	Groups []AzureAdGroup `json:"groups,omitempty"`
}

// +kubebuilder:validation:Enum=NAVident
type AzureAdExtraClaim string

type AzureAdGroup struct {
	ID string `json:"id,omitempty"`
}

func init() {
	SchemeBuilder.Register(&AzureAdApplication{}, &AzureAdApplicationList{})
}

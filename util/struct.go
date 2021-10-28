package util

type Config struct {
	Properties struct {
		TimeInWeek []struct {
			Day       string `json:"day"`
			HourSlots []int  `json:"hourSlots"`
		} `json:"timeInWeek"`
		NotAllowedTime []struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"notAllowedTime"`
	} `json:"properties"`
}

type EKSAPIParameter struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourceName      string
	ApiVersion        string
	Location          string
	ConfigName        string
	ConfigFile        Config
}

type AKSAPIParameter struct {
	Acr             string
	AksCluster      string
	BranchName      string
	DoNotWait       string
	Port            string
	BindingSelector string
	Repository      string
	Name            string
	ResourceGroup   string
	DisableBrowser  bool
	ListenAddress   string
	ListenPort      string
	Subscription    string
	Location        string
	NodepoolName    string
	Features        []string
}

type AKSAddon struct {
	ResourceGroupName          string
	ClusterName                string
	Addon                      string
	AppgwID                    string
	AppgwName                  string
	AppgwSubnetCidr            string
	AppgwSubnetID              string
	AppgwSubnetPrefix          string
	AppgwWatchNamespace        string
	EnableMsiAuthForMonitoring bool
	EnableSecretRotation       bool
	EnableSgxquotehelper       bool
	SubnetName                 string
	WorkspaceResourceID        string
}

type AKSPodIdentity struct {
	ResourceGroupName  string
	ClusterName        string
	Namespace          string
	IdentityResourceID string
	Name               string
	BindingSelector    string
	PodLabels          string
}

type AKSInstallCLI struct {
	BaseSrcURL               string
	ClientVersion            string
	InstallLocation          string
	KubeloginBaseSrcURL      string
	KubeloginInstallLocation string
	KubeloginVersion         string
	Subscription             string
}

type AKSk8sConfiguration struct {
	ClusterName   string
	ClusterType   string
	Name          string
	RepositoryURL string
	ResourceGroup string
	Scope         string
}

// error struct

type EKSAddonError struct {
	AddonName          string
	ClusterName        string
	FargateProfileName string
	Message_           string
	NodegroupName      string
}

type EKSDeleteAddonError struct {
	AddonName          string
	ClusterName        string
	FargateProfileName string
	Message_           string
	NodegroupName      string
}

// type EKSCreateAddonError struct {
// 	RespMetadata         subStruct
// 	AddonVersion       string
// 	ClientRequestToken string
// 	ClusterName        string
// 	Message_           string
// 	NodegroupName      string
// }

// type

type CloudError struct {
	// Error - Details about the error.
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

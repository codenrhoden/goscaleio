package goscaleio

import "github.com/vmware/govcloudair/types/v56"

// Type: ErrorType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: The standard error message type used in the vCloud REST API.
// Since: 0.9
type Error struct {
	Message                 string `xml:"message,attr"`
	MajorErrorCode          int    `xml:"majorErrorCode,attr"`
	MinorErrorCode          string `xml:"minorErrorCode,attr"`
	VendorSpecificErrorCode string `xml:"vendorSpecificErrorCode,attr,omitempty"`
	StackTrace              string `xml:"stackTrace,attr,omitempty"`
}

type session struct {
	Link []*types.Link `xml:"Link"`
}

type System struct {
	MdmMode                               string   `json:"mdmMode"`
	MdmClusterState                       string   `json:"mdmClusterState"`
	SecondaryMdmActorIPList               []string `json:"secondaryMdmActorIpList"`
	InstallID                             string   `json:"installId"`
	PrimaryActorIPList                    []string `json:"primaryMdmActorIpList"`
	SystemVersionName                     string   `json:"systemVersionName"`
	CapacityAlertHighThresholdPercent     int      `json:"capacityAlertHighThresholdPercent"`
	CapacityAlertCriticalThresholdPercent int      `json:"capacityAlertCriticalThresholdPercent"`
	RemoteReadOnlyLimitState              bool     `json:"remoteReadOnlyLimitState"`
	PrimaryMdmActorPort                   int      `json:"primaryMdmActorPort"`
	SecondaryMdmActorPort                 int      `json:"secondaryMdmActorPort"`
	TiebreakerMdmActorPort                int      `json:"tiebreakerMdmActorPort"`
	MdmManagementPort                     int      `json:"mdmManagementPort"`
	TiebreakerMdmIPList                   []string `json:"tiebreakerMdmIpList"`
	MdmManagementIPList                   []string `json:"mdmManagementIPList"`
	DefaultIsVolumeObfuscated             bool     `json:"defaultIsVolumeObfuscated"`
	RestrictedSdcModeEnabled              bool     `json:"restrictedSdcModeEnabled"`
	Swid                                  string   `json:"swid"`
	DaysInstalled                         int      `json:"daysInstalled"`
	MaxCapacityInGb                       string   `json:"maxCapacityInGb"`
	CapacityTimeLeftInDays                string   `json:"capacityTimeLeftInDays"`
	EnterpriseFeaturesEnabled             bool     `json:"enterpriseFeaturesEnabled"`
	IsInitialLicense                      bool     `json:"isInitialLicense"`
	Name                                  string   `json:"name"`
	ID                                    string   `json:"id"`
	Links                                 []*Link  `json:"links"`
}

type Link struct {
	Rel  string `json:"rel"`
	HREF string `json:"href"`
}

type BWC struct {
	TotalWeightInKb int `json:"totalWeightInKb"`
	NumOccured      int `json:"numOccured"`
	NumSeconds      int `json:"numSeconds"`
}

type Statistics struct {
	PrimaryReadFromDevBwc                    BWC `json:"primaryReadFromDevBwc"`
	NumOfStoragePools                        int `json:"numOfStoragePools"`
	ProtectedCapacityInKb                    int `json:"protectedCapacityInKb"`
	MovingCapacityInKb                       int `json:"movingCapacityInKb"`
	SnapCapacityInUseOccupiedInKb            int `json:"snapCapacityInUseOccupiedInKb"`
	SnapCapacityInUseInKb                    int `json:"snapCapacityInUseInKb"`
	ActiveFwdRebuildCapacityInKb             int `json:"activeFwdRebuildCapacityInKb"`
	DegradedHealthyVacInKb                   int `json:"degradedHealthyVacInKb"`
	ActiveMovingRebalanceJobs                int `json:"activeMovingRebalanceJobs"`
	TotalReadBwc                             BWC `json:"totalReadBwc"`
	MaxCapacityInKb                          int `json:"maxCapacityInKb"`
	PendingBckRebuildCapacityInKb            int `json:"pendingBckRebuildCapacityInKb"`
	ActiveMovingOutFwdRebuildJobs            int `json:"activeMovingOutFwdRebuildJobs"`
	CapacityLimitInKb                        int `json:"capacityLimitInKb"`
	SecondaryVacInKb                         int `json:"secondaryVacInKb"`
	PendingFwdRebuildCapacityInKb            int `json:"pendingFwdRebuildCapacityInKb"`
	ThinCapacityInUseInKb                    int `json:"thinCapacityInUseInKb"`
	AtRestCapacityInKb                       int `json:"atRestCapacityInKb"`
	ActiveMovingInBckRebuildJobs             int `json:"activeMovingInBckRebuildJobs"`
	DegradedHealthyCapacityInKb              int `json:"degradedHealthyCapacityInKb"`
	NumOfScsiInitiators                      int `json:"numOfScsiInitiators"`
	NumOfUnmappedVolumes                     int `json:"numOfUnmappedVolumes"`
	FailedCapacityInKb                       int `json:"failedCapacityInKb"`
	SecondaryReadFromDevBwc                  BWC `json:"secondaryReadFromDevBwc"`
	NumOfVolumes                             int `json:"numOfVolumes"`
	SecondaryWriteBwc                        BWC `json:"secondaryWriteBwc"`
	ActiveBckRebuildCapacityInKb             int `json:"activeBckRebuildCapacityInKb"`
	FailedVacInKb                            int `json:"failedVacInKb"`
	PendingMovingCapacityInKb                int `json:"pendingMovingCapacityInKb"`
	ActiveMovingInRebalanceJobs              int `json:"activeMovingInRebalanceJobs"`
	PendingMovingInRebalanceJobs             int `json:"pendingMovingInRebalanceJobs"`
	BckRebuildReadBwc                        BWC `json:"bckRebuildReadBwc"`
	DegradedFailedVacInKb                    int `json:"degradedFailedVacInKb"`
	NumOfSnapshots                           int `json:"numOfSnapshots"`
	RebalanceCapacityInKb                    int `json:"rebalanceCapacityInKb"`
	fwdRebuildReadBwc                        BWC `json:"fwdRebuildReadBwc"`
	NumOfSdc                                 int `json:"numOfSdc"`
	ActiveMovingInFwdRebuildJobs             int `json:"activeMovingInFwdRebuildJobs"`
	NumOfVtrees                              int `json:"numOfVtrees"`
	ThickCapacityInUseInKb                   int `json:"thickCapacityInUseInKb"`
	ProtectedVacInKb                         int `json:"protectedVacInKb"`
	PendingMovingInBckRebuildJobs            int `json:"pendingMovingInBckRebuildJobs"`
	CapacityAvailableForVolumeAllocationInKb int `json:"capacityAvailableForVolumeAllocationInKb"`
	PendingRebalanceCapacityInKb             int `json:"pendingRebalanceCapacityInKb"`
	PendingMovingRebalanceJobs               int `json:"pendingMovingRebalanceJobs"`
	NumOfProtectionDomains                   int `json:"numOfProtectionDomains"`
	NumOfSds                                 int `json:"numOfSds"`
	CapacityInUseInKb                        int `json:"capacityInUseInKb"`
	BckRebuildWriteBwc                       BWC `json:"bckRebuildWriteBwc"`
	DegradedFailedCapacityInKb               int `json:"degradedFailedCapacityInKb"`
	NumOfThinBaseVolumes                     int `json:"numOfThinBaseVolumes"`
	PendingMovingOutFwdRebuildJobs           int `json:"pendingMovingOutFwdRebuildJobs"`
	SecondaryReadBwc                         BWC `json:"secondaryReadBwc"`
	PendingMovingOutBckRebuildJobs           int `json:"pendingMovingOutBckRebuildJobs"`
	RebalanceWriteBwc                        BWC `json:"rebalanceWriteBwc"`
	PrimaryReadBwc                           BWC `json:"primaryReadBwc"`
	NumOfVolumesInDeletion                   int `json:"numOfVolumesInDeletion"`
	NumOfDevices                             int `json:"numOfDevices"`
	RebalanceReadBwc                         BWC `json:"rebalanceReadBwc"`
	InUseVacInKb                             int `json:"inUseVacInKb"`
	UnreachableUnusedCapacityInKb            int `json:"unreachableUnusedCapacityInKb"`
	TotalWriteBwc                            BWC `json:"totalWriteBwc"`
	SpareCapacityInKb                        int `json:"spareCapacityInKb"`
	ActiveMovingOutBckRebuildJobs            int `json:"activeMovingOutBckRebuildJobs"`
	PrimaryVacInKb                           int `json:"primaryVacInKb"`
	NumOfThickBaseVolumes                    int `json:"numOfThickBaseVolumes"`
	BckRebuildCapacityInKb                   int `json:"bckRebuildCapacityInKb"`
	NumOfMappedToAllVolumes                  int `json:"numOfMappedToAllVolumes"`
	ActiveMovingCapacityInKb                 int `json:"activeMovingCapacityInKb"`
	PendingMovingInFwdRebuildJobs            int `json:"pendingMovingInFwdRebuildJobs"`
	ActiveRebalanceCapacityInKb              int `json:"activeRebalanceCapacityInKb"`
	RmcacheSizeInKb                          int `json:"rmcacheSizeInKb"`
	FwdRebuildCapacityInKb                   int `json:"fwdRebuildCapacityInKb"`
	FwdRebuildWriteBwc                       BWC `json:"fwdRebuildWriteBwc"`
	PrimaryWriteBwc                          BWC `json:"primaryWriteBwc"`
}

type User struct {
	SystemID              string  `json:"systemId"`
	UserRole              string  `json:"userRole"`
	PasswordChangeRequire bool    `json:"passwordChangeRequired"`
	Name                  string  `json:"name"`
	ID                    string  `json:"id"`
	Links                 []*Link `json:"links"`
}

type ScsiInitiator struct {
	Name     string  `json:"name"`
	IQN      string  `json:"iqn"`
	SystemID string  `json:"systemID"`
	Links    []*Link `json:"links"`
}

type ProtectionDomain struct {
	SystemID                          string  `json:"systemId"`
	RebuildNetworkThrottlingInKbps    int     `json:"rebuildNetworkThrottlingInKbps"`
	RebalanceNetworkThrottlingInKbps  int     `json:"rebalanceNetworkThrottlingInKbps"`
	OverallIoNetworkThrottlingInKbps  int     `json:"overallIoNetworkThrottlingInKbps"`
	OverallIoNetworkThrottlingEnabled bool    `json:"overallIoNetworkThrottlingEnabled"`
	RebuildNetworkThrottlingEnabled   bool    `json:"rebuildNetworkThrottlingEnabled"`
	RebalanceNetworkThrottlingEnabled bool    `json:"rebalanceNetworkThrottlingEnabled"`
	ProtectionDomainState             string  `json:"protectionDomainState"`
	Name                              string  `json:"name"`
	ID                                string  `json:"id"`
	Links                             []*Link `json:"links"`
}

type Sdc struct {
	SystemID           string  `json:"systemId"`
	SdcApproved        bool    `json:"sdcApproved"`
	SdcIp              string  `json:"SdcIp"`
	OnVmWare           bool    `json:"onVmWare"`
	SdcGuid            string  `json:"sdcGuid"`
	MdmConnectionState string  `json:"mdmConnectionState"`
	Name               string  `json:"name"`
	ID                 string  `json:"id"`
	Links              []*Link `json:"links"`
}

type StoragePool struct {
	ProtectionDomainID                               string `json:"protectionDomainId"`
	RebalanceioPriorityPolicy                        string `json:"rebalanceIoPriorityPolicy"`
	RebuildioPriorityPolicy                          string `json:"rebuildIoPriorityPolicy"`
	RebuildioPriorityBwLimitPerDeviceInKbps          int    `json:"rebuildIoPriorityBwLimitPerDeviceInKbps"`
	RebuildioPriorityNumOfConcurrentIosPerDevice     int    `json:"rebuildIoPriorityNumOfConcurrentIosPerDevice"`
	RebalanceioPriorityNumOfConcurrentIosPerDevice   int    `json:"rebalanceIoPriorityNumOfConcurrentIosPerDevice"`
	RebalanceioPriorityBwLimitPerDeviceInKbps        int    `json:"rebalanceIoPriorityBwLimitPerDeviceInKbps"`
	RebuildioPriorityAppIopsPerDeviceThreshold       int    `json:"rebuildIoPriorityAppIopsPerDeviceThreshold"`
	RebalanceioPriorityAppIopsPerDeviceThreshold     int    `json:"rebalanceIoPriorityAppIopsPerDeviceThreshold"`
	RebuildioPriorityAppBwPerDeviceThresholdInKbps   int    `json:"rebuildIoPriorityAppBwPerDeviceThresholdInKbps"`
	RebalanceioPriorityAppBwPerDeviceThresholdInKbps int    `json:"rebalanceIoPriorityAppBwPerDeviceThresholdInKbps"`
	RebuildioPriorityQuietPeriodInMsec               int    `json:"rebuildIoPriorityQuietPeriodInMsec"`
	RebalanceioPriorityQuietPeriodInMsec             int    `json:"rebalanceIoPriorityQuietPeriodInMsec"`
	ZeroPaddingEnabled                               bool   `json:"zeroPaddingEnabled"`
	UseRmcache                                       bool   `json:"useRmcache"`
	SparePercentage                                  int    `json:"sparePercentage"`
	RmCacheWriteHandlingMode                         string `json:"rmcacheWriteHandlingMode"`
	RebuildEnabled                                   bool   `json:"rebuildEnabled"`
	RebalanceEnabled                                 bool   `json:"rebalanceEnabled"`
	NumofParallelRebuildRebalanceJobsPerDevice       int    `json:"numOfParallelRebuildRebalanceJobsPerDevice"`
	Name                                             string `json:"name"`
	ID                                               string `json:"id"`
	Links                                            []*Link
}

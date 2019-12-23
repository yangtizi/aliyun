package instances

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// type Instance struct {
// 	ImageId                    string                               `json:"ImageId" xml:"ImageId"`
// 	InstanceType               string                               `json:"InstanceType" xml:"InstanceType"`
// 	AutoReleaseTime            string                               `json:"AutoReleaseTime" xml:"AutoReleaseTime"`
// 	OsType                     string                               `json:"OsType" xml:"OsType"`
// 	DeviceAvailable            bool                                 `json:"DeviceAvailable" xml:"DeviceAvailable"`
// 	InstanceNetworkType        string                               `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
// 	LocalStorageAmount         int                                  `json:"LocalStorageAmount" xml:"LocalStorageAmount"`
// 	NetworkType                string                               `json:"NetworkType" xml:"NetworkType"`
// 	IsSpot                     bool                                 `json:"IsSpot" xml:"IsSpot"`
// 	InstanceChargeType         string                               `json:"InstanceChargeType" xml:"InstanceChargeType"`
// 	ClusterId                  string                               `json:"ClusterId" xml:"ClusterId"`
// 	InstanceName               string                               `json:"InstanceName" xml:"InstanceName"`
// 	CreditSpecification        string                               `json:"CreditSpecification" xml:"CreditSpecification"`
// 	GPUAmount                  int                                  `json:"GPUAmount" xml:"GPUAmount"`
// 	StartTime                  string                               `json:"StartTime" xml:"StartTime"`
// 	ZoneId                     string                               `json:"ZoneId" xml:"ZoneId"`
// 	InternetChargeType         string                               `json:"InternetChargeType" xml:"InternetChargeType"`
// 	InternetMaxBandwidthIn     int                                  `json:"InternetMaxBandwidthIn" xml:"InternetMaxBandwidthIn"`
// 	HostName                   string                               `json:"HostName" xml:"HostName"`
// 	Status                     string                               `json:"Status" xml:"Status"`
// 	CPU                        int                                  `json:"CPU" xml:"CPU"`
// 	Cpu                        int                                  `json:"Cpu" xml:"Cpu"`
// 	SpotPriceLimit             float64                              `json:"SpotPriceLimit" xml:"SpotPriceLimit"`
// 	OSName                     string                               `json:"OSName" xml:"OSName"`
// 	OSNameEn                   string                               `json:"OSNameEn" xml:"OSNameEn"`
// 	SerialNumber               string                               `json:"SerialNumber" xml:"SerialNumber"`
// 	RegionId                   string                               `json:"RegionId" xml:"RegionId"`
// 	IoOptimized                bool                                 `json:"IoOptimized" xml:"IoOptimized"`
// 	InternetMaxBandwidthOut    int                                  `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
// 	ResourceGroupId            string                               `json:"ResourceGroupId" xml:"ResourceGroupId"`
// 	InstanceTypeFamily         string                               `json:"InstanceTypeFamily" xml:"InstanceTypeFamily"`
// 	InstanceId                 string                               `json:"InstanceId" xml:"InstanceId"`
// 	DeploymentSetId            string                               `json:"DeploymentSetId" xml:"DeploymentSetId"`
// 	GPUSpec                    string                               `json:"GPUSpec" xml:"GPUSpec"`
// 	Description                string                               `json:"Description" xml:"Description"`
// 	Recyclable                 bool                                 `json:"Recyclable" xml:"Recyclable"`
// 	SaleCycle                  string                               `json:"SaleCycle" xml:"SaleCycle"`
// 	ExpiredTime                string                               `json:"ExpiredTime" xml:"ExpiredTime"`
// 	OSType                     string                               `json:"OSType" xml:"OSType"`
// 	Memory                     int                                  `json:"Memory" xml:"Memory"`
// 	CreationTime               string                               `json:"CreationTime" xml:"CreationTime"`
// 	KeyPairName                string                               `json:"KeyPairName" xml:"KeyPairName"`
// 	HpcClusterId               string                               `json:"HpcClusterId" xml:"HpcClusterId"`
// 	LocalStorageCapacity       int64                                `json:"LocalStorageCapacity" xml:"LocalStorageCapacity"`
// 	VlanId                     string                               `json:"VlanId" xml:"VlanId"`
// 	StoppedMode                string                               `json:"StoppedMode" xml:"StoppedMode"`
// 	SpotStrategy               string                               `json:"SpotStrategy" xml:"SpotStrategy"`
// 	DeletionProtection         bool                                 `json:"DeletionProtection" xml:"DeletionProtection"`
// 	SecurityGroupIds           SecurityGroupIdsInDescribeInstances  `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
// 	InnerIpAddress             InnerIpAddressInDescribeInstances    `json:"InnerIpAddress" xml:"InnerIpAddress"`
// 	PublicIpAddress            PublicIpAddressInDescribeInstances   `json:"PublicIpAddress" xml:"PublicIpAddress"`
// 	RdmaIpAddress              RdmaIpAddress                        `json:"RdmaIpAddress" xml:"RdmaIpAddress"`
// 	EipAddress                 EipAddressInDescribeInstances        `json:"EipAddress" xml:"EipAddress"`
// 	CpuOptions                 CpuOptions                           `json:"CpuOptions" xml:"CpuOptions"`
// 	EcsCapacityReservationAttr EcsCapacityReservationAttr           `json:"EcsCapacityReservationAttr" xml:"EcsCapacityReservationAttr"`
// 	DedicatedHostAttribute     DedicatedHostAttribute               `json:"DedicatedHostAttribute" xml:"DedicatedHostAttribute"`
// 	DedicatedInstanceAttribute DedicatedInstanceAttribute           `json:"DedicatedInstanceAttribute" xml:"DedicatedInstanceAttribute"`
// 	VpcAttributes              VpcAttributes                        `json:"VpcAttributes" xml:"VpcAttributes"`
// 	NetworkInterfaces          NetworkInterfacesInDescribeInstances `json:"NetworkInterfaces" xml:"NetworkInterfaces"`
// 	OperationLocks             OperationLocksInDescribeInstances    `json:"OperationLocks" xml:"OperationLocks"`
// 	Tags                       TagsInDescribeInstances              `json:"Tags" xml:"Tags"`
// }

func DescribeInstances() {
	client, err := ecs.NewClientWithAccessKey("cn-hongkong", "LTAI4FjV1o5Jir3wtUKMLL2h", "VVVe6PvArNftuhjv0KEFyhIRyMMfQV")

	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"
	request.InstanceName = "sql*"

	response, err := client.DescribeInstances(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	response.TotalCount

	for i, v := range response.Instances.Instance {
		// fmt.Printf("response is %#v\n", v)
		fmt.Println("第", i, "  -->>", v.InstanceId)
	}
}

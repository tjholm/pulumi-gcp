// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get info about a Google Cloud Redis instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/redis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myInstance, err := redis.LookupInstance(ctx, &redis.LookupInstanceArgs{
//				Name: "my-redis-instance",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("instanceMemorySizeGb", myInstance.MemorySizeGb)
//			ctx.Export("instanceConnectMode", myInstance.ConnectMode)
//			ctx.Export("instanceAuthorizedNetwork", myInstance.AuthorizedNetwork)
//			return nil
//		})
//	}
//
// ```
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("gcp:redis/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	// The name of a Redis instance.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the resource belongs. If it
	// is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	AlternativeLocationId string `pulumi:"alternativeLocationId"`
	AuthEnabled           bool   `pulumi:"authEnabled"`
	AuthString            string `pulumi:"authString"`
	AuthorizedNetwork     string `pulumi:"authorizedNetwork"`
	ConnectMode           string `pulumi:"connectMode"`
	CreateTime            string `pulumi:"createTime"`
	CurrentLocationId     string `pulumi:"currentLocationId"`
	CustomerManagedKey    string `pulumi:"customerManagedKey"`
	DisplayName           string `pulumi:"displayName"`
	Host                  string `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string                           `pulumi:"id"`
	Labels                 map[string]string                `pulumi:"labels"`
	LocationId             string                           `pulumi:"locationId"`
	MaintenancePolicies    []GetInstanceMaintenancePolicy   `pulumi:"maintenancePolicies"`
	MaintenanceSchedules   []GetInstanceMaintenanceSchedule `pulumi:"maintenanceSchedules"`
	MemorySizeGb           int                              `pulumi:"memorySizeGb"`
	Name                   string                           `pulumi:"name"`
	Nodes                  []GetInstanceNode                `pulumi:"nodes"`
	PersistenceIamIdentity string                           `pulumi:"persistenceIamIdentity"`
	Port                   int                              `pulumi:"port"`
	Project                *string                          `pulumi:"project"`
	ReadEndpoint           string                           `pulumi:"readEndpoint"`
	ReadEndpointPort       int                              `pulumi:"readEndpointPort"`
	ReadReplicasMode       string                           `pulumi:"readReplicasMode"`
	RedisConfigs           map[string]string                `pulumi:"redisConfigs"`
	RedisVersion           string                           `pulumi:"redisVersion"`
	Region                 *string                          `pulumi:"region"`
	ReplicaCount           int                              `pulumi:"replicaCount"`
	ReservedIpRange        string                           `pulumi:"reservedIpRange"`
	SecondaryIpRange       string                           `pulumi:"secondaryIpRange"`
	ServerCaCerts          []GetInstanceServerCaCert        `pulumi:"serverCaCerts"`
	Tier                   string                           `pulumi:"tier"`
	TransitEncryptionMode  string                           `pulumi:"transitEncryptionMode"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			var s LookupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceResultOutput)
}

// A collection of arguments for invoking getInstance.
type LookupInstanceOutputArgs struct {
	// The name of a Redis instance.
	Name pulumi.StringInput `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The region in which the resource belongs. If it
	// is not provided, the provider region is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getInstance.
type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) AlternativeLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.AlternativeLocationId }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) AuthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.AuthEnabled }).(pulumi.BoolOutput)
}

func (o LookupInstanceResultOutput) AuthString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.AuthString }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) AuthorizedNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.AuthorizedNetwork }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) ConnectMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ConnectMode }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) CurrentLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CurrentLocationId }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) CustomerManagedKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CustomerManagedKey }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Host }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupInstanceResultOutput) LocationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.LocationId }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) MaintenancePolicies() GetInstanceMaintenancePolicyArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceMaintenancePolicy { return v.MaintenancePolicies }).(GetInstanceMaintenancePolicyArrayOutput)
}

func (o LookupInstanceResultOutput) MaintenanceSchedules() GetInstanceMaintenanceScheduleArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceMaintenanceSchedule { return v.MaintenanceSchedules }).(GetInstanceMaintenanceScheduleArrayOutput)
}

func (o LookupInstanceResultOutput) MemorySizeGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.MemorySizeGb }).(pulumi.IntOutput)
}

func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Nodes() GetInstanceNodeArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceNode { return v.Nodes }).(GetInstanceNodeArrayOutput)
}

func (o LookupInstanceResultOutput) PersistenceIamIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.PersistenceIamIdentity }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o LookupInstanceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) ReadEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ReadEndpoint }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) ReadEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.ReadEndpointPort }).(pulumi.IntOutput)
}

func (o LookupInstanceResultOutput) ReadReplicasMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ReadReplicasMode }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) RedisConfigs() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.RedisConfigs }).(pulumi.StringMapOutput)
}

func (o LookupInstanceResultOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.RedisVersion }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) ReplicaCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.ReplicaCount }).(pulumi.IntOutput)
}

func (o LookupInstanceResultOutput) ReservedIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ReservedIpRange }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) SecondaryIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.SecondaryIpRange }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) ServerCaCerts() GetInstanceServerCaCertArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceServerCaCert { return v.ServerCaCerts }).(GetInstanceServerCaCertArrayOutput)
}

func (o LookupInstanceResultOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Tier }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) TransitEncryptionMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.TransitEncryptionMode }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}

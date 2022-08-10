// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a VM instance resource within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/instances)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.LookupInstance(ctx, &compute.LookupInstanceArgs{
//				Name: pulumi.StringRef("primary-application-server"),
//				Zone: pulumi.StringRef("us-central1-a"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("gcp:compute/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	// The name of the instance. One of `name` or `selfLink` must be provided.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If `selfLink` is provided, this value is ignored.  If neither `selfLink`
	// nor `project` are provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The self link of the instance. One of `name` or `selfLink` must be provided.
	SelfLink *string `pulumi:"selfLink"`
	// The zone of the instance. If `selfLink` is provided, this
	// value is ignored.  If neither `selfLink` nor `zone` are provided, the
	// provider zone is used.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	AdvancedMachineFeatures []GetInstanceAdvancedMachineFeature `pulumi:"advancedMachineFeatures"`
	AllowStoppingForUpdate  bool                                `pulumi:"allowStoppingForUpdate"`
	// List of disks attached to the instance. Structure is documented below.
	AttachedDisks []GetInstanceAttachedDisk `pulumi:"attachedDisks"`
	// The boot disk for the instance. Structure is documented below.
	BootDisks []GetInstanceBootDisk `pulumi:"bootDisks"`
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	CanIpForward                bool                                    `pulumi:"canIpForward"`
	ConfidentialInstanceConfigs []GetInstanceConfidentialInstanceConfig `pulumi:"confidentialInstanceConfigs"`
	// The CPU platform used by this instance.
	CpuPlatform   string `pulumi:"cpuPlatform"`
	CurrentStatus string `pulumi:"currentStatus"`
	// Whether deletion protection is enabled on this instance.
	DeletionProtection bool `pulumi:"deletionProtection"`
	// A brief description of the resource.
	Description   string `pulumi:"description"`
	DesiredStatus string `pulumi:"desiredStatus"`
	EnableDisplay bool   `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance. Structure is documented below.
	GuestAccelerators []GetInstanceGuestAccelerator `pulumi:"guestAccelerators"`
	Hostname          string                        `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The server-assigned unique identifier of this instance.
	InstanceId string `pulumi:"instanceId"`
	// The unique fingerprint of the labels.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// A set of key/value label pairs assigned to the instance.
	Labels map[string]string `pulumi:"labels"`
	// The machine type to create.
	MachineType string `pulumi:"machineType"`
	// Metadata key/value pairs made available within the instance.
	Metadata map[string]string `pulumi:"metadata"`
	// The unique fingerprint of the metadata.
	MetadataFingerprint   string `pulumi:"metadataFingerprint"`
	MetadataStartupScript string `pulumi:"metadataStartupScript"`
	// The minimum CPU platform specified for the VM instance.
	MinCpuPlatform string  `pulumi:"minCpuPlatform"`
	Name           *string `pulumi:"name"`
	// The networks attached to the instance. Structure is documented below.
	NetworkInterfaces []GetInstanceNetworkInterface `pulumi:"networkInterfaces"`
	// The network performance configuration setting for the instance, if set. Structure is documented below.
	NetworkPerformanceConfigs []GetInstanceNetworkPerformanceConfig `pulumi:"networkPerformanceConfigs"`
	Project                   *string                               `pulumi:"project"`
	ReservationAffinities     []GetInstanceReservationAffinity      `pulumi:"reservationAffinities"`
	ResourcePolicies          []string                              `pulumi:"resourcePolicies"`
	// The scheduling strategy being used by the instance. Structure is documented below
	Schedulings []GetInstanceScheduling `pulumi:"schedulings"`
	// The scratch disks attached to the instance. Structure is documented below.
	ScratchDisks []GetInstanceScratchDisk `pulumi:"scratchDisks"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The service account to attach to the instance. Structure is documented below.
	ServiceAccounts []GetInstanceServiceAccount `pulumi:"serviceAccounts"`
	// The shielded vm config being used by the instance. Structure is documented below.
	ShieldedInstanceConfigs []GetInstanceShieldedInstanceConfig `pulumi:"shieldedInstanceConfigs"`
	// The list of tags attached to the instance.
	Tags []string `pulumi:"tags"`
	// The unique fingerprint of the tags.
	TagsFingerprint string  `pulumi:"tagsFingerprint"`
	Zone            *string `pulumi:"zone"`
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
	// The name of the instance. One of `name` or `selfLink` must be provided.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If `selfLink` is provided, this value is ignored.  If neither `selfLink`
	// nor `project` are provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The self link of the instance. One of `name` or `selfLink` must be provided.
	SelfLink pulumi.StringPtrInput `pulumi:"selfLink"`
	// The zone of the instance. If `selfLink` is provided, this
	// value is ignored.  If neither `selfLink` nor `zone` are provided, the
	// provider zone is used.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
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

func (o LookupInstanceResultOutput) AdvancedMachineFeatures() GetInstanceAdvancedMachineFeatureArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceAdvancedMachineFeature { return v.AdvancedMachineFeatures }).(GetInstanceAdvancedMachineFeatureArrayOutput)
}

func (o LookupInstanceResultOutput) AllowStoppingForUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.AllowStoppingForUpdate }).(pulumi.BoolOutput)
}

// List of disks attached to the instance. Structure is documented below.
func (o LookupInstanceResultOutput) AttachedDisks() GetInstanceAttachedDiskArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceAttachedDisk { return v.AttachedDisks }).(GetInstanceAttachedDiskArrayOutput)
}

// The boot disk for the instance. Structure is documented below.
func (o LookupInstanceResultOutput) BootDisks() GetInstanceBootDiskArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceBootDisk { return v.BootDisks }).(GetInstanceBootDiskArrayOutput)
}

// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
func (o LookupInstanceResultOutput) CanIpForward() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.CanIpForward }).(pulumi.BoolOutput)
}

func (o LookupInstanceResultOutput) ConfidentialInstanceConfigs() GetInstanceConfidentialInstanceConfigArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceConfidentialInstanceConfig {
		return v.ConfidentialInstanceConfigs
	}).(GetInstanceConfidentialInstanceConfigArrayOutput)
}

// The CPU platform used by this instance.
func (o LookupInstanceResultOutput) CpuPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CpuPlatform }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) CurrentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CurrentStatus }).(pulumi.StringOutput)
}

// Whether deletion protection is enabled on this instance.
func (o LookupInstanceResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// A brief description of the resource.
func (o LookupInstanceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) DesiredStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.DesiredStatus }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) EnableDisplay() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.EnableDisplay }).(pulumi.BoolOutput)
}

// List of the type and count of accelerator cards attached to the instance. Structure is documented below.
func (o LookupInstanceResultOutput) GuestAccelerators() GetInstanceGuestAcceleratorArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceGuestAccelerator { return v.GuestAccelerators }).(GetInstanceGuestAcceleratorArrayOutput)
}

func (o LookupInstanceResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The server-assigned unique identifier of this instance.
func (o LookupInstanceResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The unique fingerprint of the labels.
func (o LookupInstanceResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// A set of key/value label pairs assigned to the instance.
func (o LookupInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The machine type to create.
func (o LookupInstanceResultOutput) MachineType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.MachineType }).(pulumi.StringOutput)
}

// Metadata key/value pairs made available within the instance.
func (o LookupInstanceResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The unique fingerprint of the metadata.
func (o LookupInstanceResultOutput) MetadataFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.MetadataFingerprint }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) MetadataStartupScript() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.MetadataStartupScript }).(pulumi.StringOutput)
}

// The minimum CPU platform specified for the VM instance.
func (o LookupInstanceResultOutput) MinCpuPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.MinCpuPlatform }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The networks attached to the instance. Structure is documented below.
func (o LookupInstanceResultOutput) NetworkInterfaces() GetInstanceNetworkInterfaceArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceNetworkInterface { return v.NetworkInterfaces }).(GetInstanceNetworkInterfaceArrayOutput)
}

// The network performance configuration setting for the instance, if set. Structure is documented below.
func (o LookupInstanceResultOutput) NetworkPerformanceConfigs() GetInstanceNetworkPerformanceConfigArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceNetworkPerformanceConfig { return v.NetworkPerformanceConfigs }).(GetInstanceNetworkPerformanceConfigArrayOutput)
}

func (o LookupInstanceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) ReservationAffinities() GetInstanceReservationAffinityArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceReservationAffinity { return v.ReservationAffinities }).(GetInstanceReservationAffinityArrayOutput)
}

func (o LookupInstanceResultOutput) ResourcePolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.ResourcePolicies }).(pulumi.StringArrayOutput)
}

// The scheduling strategy being used by the instance. Structure is documented below
func (o LookupInstanceResultOutput) Schedulings() GetInstanceSchedulingArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceScheduling { return v.Schedulings }).(GetInstanceSchedulingArrayOutput)
}

// The scratch disks attached to the instance. Structure is documented below.
func (o LookupInstanceResultOutput) ScratchDisks() GetInstanceScratchDiskArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceScratchDisk { return v.ScratchDisks }).(GetInstanceScratchDiskArrayOutput)
}

// The URI of the created resource.
func (o LookupInstanceResultOutput) SelfLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.SelfLink }).(pulumi.StringPtrOutput)
}

// The service account to attach to the instance. Structure is documented below.
func (o LookupInstanceResultOutput) ServiceAccounts() GetInstanceServiceAccountArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceServiceAccount { return v.ServiceAccounts }).(GetInstanceServiceAccountArrayOutput)
}

// The shielded vm config being used by the instance. Structure is documented below.
func (o LookupInstanceResultOutput) ShieldedInstanceConfigs() GetInstanceShieldedInstanceConfigArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceShieldedInstanceConfig { return v.ShieldedInstanceConfigs }).(GetInstanceShieldedInstanceConfigArrayOutput)
}

// The list of tags attached to the instance.
func (o LookupInstanceResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The unique fingerprint of the tags.
func (o LookupInstanceResultOutput) TagsFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.TagsFingerprint }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}

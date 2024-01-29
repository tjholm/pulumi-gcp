// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **Note**: Global instance templates can be used in any region. To lower the impact of outages outside your region and gain data residency within your region, use google_compute_region_instance_template.
//
// Get information about a VM instance template resource within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/instance-templates)
// and
// [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceTemplates).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.LookupInstanceTemplate(ctx, &compute.LookupInstanceTemplateArgs{
//				Filter:     pulumi.StringRef("name != generic-tpl-20200107"),
//				MostRecent: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.LookupInstanceTemplate(ctx, &compute.LookupInstanceTemplateArgs{
//				SelfLinkUnique: pulumi.StringRef("https://www.googleapis.com/compute/v1/projects/your-project-name/global/instanceTemplates/example-template-custom?uniqueId=1234"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupInstanceTemplate(ctx *pulumi.Context, args *LookupInstanceTemplateArgs, opts ...pulumi.InvokeOption) (*LookupInstanceTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceTemplateResult
	err := ctx.Invoke("gcp:compute/getInstanceTemplate:getInstanceTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTemplate.
type LookupInstanceTemplateArgs struct {
	// A filter to retrieve the instance templates.
	// See [gcloud topic filters](https://cloud.google.com/sdk/gcloud/reference/topic/filters) for reference.
	// If multiple instance templates match, either adjust the filter or specify `mostRecent`.
	// One of `name`, `filter` or `selfLinkUnique` must be provided.
	Filter *string `pulumi:"filter"`
	// If `filter` is provided, ensures the most recent template is returned when multiple instance templates match. One of `name`, `filter` or `selfLinkUnique` must be provided.
	MostRecent *bool `pulumi:"mostRecent"`
	// The name of the instance template. One of `name`, `filter` or `selfLinkUnique` must be provided.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If `project` is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The selfLinkUnique URI of the instance template. One of `name`, `filter` or `selfLinkUnique` must be provided.
	SelfLinkUnique *string `pulumi:"selfLinkUnique"`
}

// A collection of values returned by getInstanceTemplate.
type LookupInstanceTemplateResult struct {
	AdvancedMachineFeatures []GetInstanceTemplateAdvancedMachineFeature `pulumi:"advancedMachineFeatures"`
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs. This defaults to false.
	CanIpForward bool `pulumi:"canIpForward"`
	// Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. Structure is documented below
	ConfidentialInstanceConfigs []GetInstanceTemplateConfidentialInstanceConfig `pulumi:"confidentialInstanceConfigs"`
	// A brief description of this resource.
	Description string `pulumi:"description"`
	// Disks to attach to instances created from this template.
	// This can be specified multiple times for multiple disks. Structure is
	// documented below.
	Disks           []GetInstanceTemplateDisk `pulumi:"disks"`
	EffectiveLabels map[string]string         `pulumi:"effectiveLabels"`
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true in order to update this field.
	EnableDisplay bool    `pulumi:"enableDisplay"`
	Filter        *string `pulumi:"filter"`
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	GuestAccelerators []GetInstanceTemplateGuestAccelerator `pulumi:"guestAccelerators"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A brief description to use for instances
	// created from this template.
	InstanceDescription string `pulumi:"instanceDescription"`
	// (Optional) A set of ket/value label pairs to assign to disk created from
	// this template
	Labels map[string]string `pulumi:"labels"`
	// The machine type to create.
	MachineType string `pulumi:"machineType"`
	// Metadata key/value pairs to make available from
	// within instances created from this template.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The unique fingerprint of the metadata.
	MetadataFingerprint string `pulumi:"metadataFingerprint"`
	// An alternative to using the
	// startup-script metadata key, mostly to match the computeInstance resource.
	// This replaces the startup-script metadata key on the created instance and
	// thus the two mechanisms are not allowed to be used simultaneously.
	MetadataStartupScript string `pulumi:"metadataStartupScript"`
	// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	MinCpuPlatform string `pulumi:"minCpuPlatform"`
	MostRecent     *bool  `pulumi:"mostRecent"`
	// The name of the instance template. If you leave
	// this blank, the provider will auto-generate a unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix string `pulumi:"namePrefix"`
	// The URL of the network attachment that this interface should connect to in the following format: projects/{projectNumber}/regions/{region_name}/networkAttachments/{network_attachment_name}.  s
	NetworkInterfaces []GetInstanceTemplateNetworkInterface `pulumi:"networkInterfaces"`
	// The network performance configuration setting
	// for the instance, if set. Structure is documented below.
	NetworkPerformanceConfigs []GetInstanceTemplateNetworkPerformanceConfig `pulumi:"networkPerformanceConfigs"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project      *string           `pulumi:"project"`
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// An instance template is a global resource that is not
	// bound to a zone or a region. However, you can still specify some regional
	// resources in an instance template, which restricts the template to the
	// region where that resource resides. For example, a custom `subnetwork`
	// resource is tied to a specific region. Defaults to the region of the
	// Provider if no value is given.
	Region                string                                   `pulumi:"region"`
	ReservationAffinities []GetInstanceTemplateReservationAffinity `pulumi:"reservationAffinities"`
	ResourceManagerTags   map[string]string                        `pulumi:"resourceManagerTags"`
	// (Optional) -- A list of short names of resource policies to attach to this disk for automatic snapshot creations. Currently a max of 1 resource policy is supported.
	ResourcePolicies []string `pulumi:"resourcePolicies"`
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Schedulings []GetInstanceTemplateScheduling `pulumi:"schedulings"`
	// The URI of the created resource.
	SelfLink string `pulumi:"selfLink"`
	// A special URI of the created resource that uniquely identifies this instance template with the following format: `projects/{{project}}/global/instanceTemplates/{{name}}?uniqueId={{uniqueId}}`
	// Referencing an instance template via this attribute prevents Time of Check to Time of Use attacks when the instance template resides in a shared/untrusted environment.
	SelfLinkUnique *string `pulumi:"selfLinkUnique"`
	// Service account to attach to the instance. Structure is documented below.
	ServiceAccounts []GetInstanceTemplateServiceAccount `pulumi:"serviceAccounts"`
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	ShieldedInstanceConfigs []GetInstanceTemplateShieldedInstanceConfig `pulumi:"shieldedInstanceConfigs"`
	// Tags to attach to the instance.
	Tags []string `pulumi:"tags"`
	// The unique fingerprint of the tags.
	TagsFingerprint string `pulumi:"tagsFingerprint"`
}

func LookupInstanceTemplateOutput(ctx *pulumi.Context, args LookupInstanceTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceTemplateResult, error) {
			args := v.(LookupInstanceTemplateArgs)
			r, err := LookupInstanceTemplate(ctx, &args, opts...)
			var s LookupInstanceTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceTemplateResultOutput)
}

// A collection of arguments for invoking getInstanceTemplate.
type LookupInstanceTemplateOutputArgs struct {
	// A filter to retrieve the instance templates.
	// See [gcloud topic filters](https://cloud.google.com/sdk/gcloud/reference/topic/filters) for reference.
	// If multiple instance templates match, either adjust the filter or specify `mostRecent`.
	// One of `name`, `filter` or `selfLinkUnique` must be provided.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// If `filter` is provided, ensures the most recent template is returned when multiple instance templates match. One of `name`, `filter` or `selfLinkUnique` must be provided.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// The name of the instance template. One of `name`, `filter` or `selfLinkUnique` must be provided.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If `project` is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The selfLinkUnique URI of the instance template. One of `name`, `filter` or `selfLinkUnique` must be provided.
	SelfLinkUnique pulumi.StringPtrInput `pulumi:"selfLinkUnique"`
}

func (LookupInstanceTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceTemplate.
type LookupInstanceTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceTemplateResult)(nil)).Elem()
}

func (o LookupInstanceTemplateResultOutput) ToLookupInstanceTemplateResultOutput() LookupInstanceTemplateResultOutput {
	return o
}

func (o LookupInstanceTemplateResultOutput) ToLookupInstanceTemplateResultOutputWithContext(ctx context.Context) LookupInstanceTemplateResultOutput {
	return o
}

func (o LookupInstanceTemplateResultOutput) AdvancedMachineFeatures() GetInstanceTemplateAdvancedMachineFeatureArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []GetInstanceTemplateAdvancedMachineFeature {
		return v.AdvancedMachineFeatures
	}).(GetInstanceTemplateAdvancedMachineFeatureArrayOutput)
}

// Whether to allow sending and receiving of
// packets with non-matching source or destination IPs. This defaults to false.
func (o LookupInstanceTemplateResultOutput) CanIpForward() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) bool { return v.CanIpForward }).(pulumi.BoolOutput)
}

// Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. Structure is documented below
func (o LookupInstanceTemplateResultOutput) ConfidentialInstanceConfigs() GetInstanceTemplateConfidentialInstanceConfigArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []GetInstanceTemplateConfidentialInstanceConfig {
		return v.ConfidentialInstanceConfigs
	}).(GetInstanceTemplateConfidentialInstanceConfigArrayOutput)
}

// A brief description of this resource.
func (o LookupInstanceTemplateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.Description }).(pulumi.StringOutput)
}

// Disks to attach to instances created from this template.
// This can be specified multiple times for multiple disks. Structure is
// documented below.
func (o LookupInstanceTemplateResultOutput) Disks() GetInstanceTemplateDiskArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []GetInstanceTemplateDisk { return v.Disks }).(GetInstanceTemplateDiskArrayOutput)
}

func (o LookupInstanceTemplateResultOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) map[string]string { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
// **Note**: `allowStoppingForUpdate` must be set to true in order to update this field.
func (o LookupInstanceTemplateResultOutput) EnableDisplay() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) bool { return v.EnableDisplay }).(pulumi.BoolOutput)
}

func (o LookupInstanceTemplateResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// List of the type and count of accelerator cards attached to the instance. Structure documented below.
func (o LookupInstanceTemplateResultOutput) GuestAccelerators() GetInstanceTemplateGuestAcceleratorArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []GetInstanceTemplateGuestAccelerator { return v.GuestAccelerators }).(GetInstanceTemplateGuestAcceleratorArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

// A brief description to use for instances
// created from this template.
func (o LookupInstanceTemplateResultOutput) InstanceDescription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.InstanceDescription }).(pulumi.StringOutput)
}

// (Optional) A set of ket/value label pairs to assign to disk created from
// this template
func (o LookupInstanceTemplateResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The machine type to create.
func (o LookupInstanceTemplateResultOutput) MachineType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.MachineType }).(pulumi.StringOutput)
}

// Metadata key/value pairs to make available from
// within instances created from this template.
func (o LookupInstanceTemplateResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// The unique fingerprint of the metadata.
func (o LookupInstanceTemplateResultOutput) MetadataFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.MetadataFingerprint }).(pulumi.StringOutput)
}

// An alternative to using the
// startup-script metadata key, mostly to match the computeInstance resource.
// This replaces the startup-script metadata key on the created instance and
// thus the two mechanisms are not allowed to be used simultaneously.
func (o LookupInstanceTemplateResultOutput) MetadataStartupScript() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.MetadataStartupScript }).(pulumi.StringOutput)
}

// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
func (o LookupInstanceTemplateResultOutput) MinCpuPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.MinCpuPlatform }).(pulumi.StringOutput)
}

func (o LookupInstanceTemplateResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

// The name of the instance template. If you leave
// this blank, the provider will auto-generate a unique name.
func (o LookupInstanceTemplateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Creates a unique name beginning with the specified
// prefix. Conflicts with `name`.
func (o LookupInstanceTemplateResultOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.NamePrefix }).(pulumi.StringOutput)
}

// The URL of the network attachment that this interface should connect to in the following format: projects/{projectNumber}/regions/{region_name}/networkAttachments/{network_attachment_name}.  s
func (o LookupInstanceTemplateResultOutput) NetworkInterfaces() GetInstanceTemplateNetworkInterfaceArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []GetInstanceTemplateNetworkInterface { return v.NetworkInterfaces }).(GetInstanceTemplateNetworkInterfaceArrayOutput)
}

// The network performance configuration setting
// for the instance, if set. Structure is documented below.
func (o LookupInstanceTemplateResultOutput) NetworkPerformanceConfigs() GetInstanceTemplateNetworkPerformanceConfigArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []GetInstanceTemplateNetworkPerformanceConfig {
		return v.NetworkPerformanceConfigs
	}).(GetInstanceTemplateNetworkPerformanceConfigArrayOutput)
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (o LookupInstanceTemplateResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceTemplateResultOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) map[string]string { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// An instance template is a global resource that is not
// bound to a zone or a region. However, you can still specify some regional
// resources in an instance template, which restricts the template to the
// region where that resource resides. For example, a custom `subnetwork`
// resource is tied to a specific region. Defaults to the region of the
// Provider if no value is given.
func (o LookupInstanceTemplateResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupInstanceTemplateResultOutput) ReservationAffinities() GetInstanceTemplateReservationAffinityArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []GetInstanceTemplateReservationAffinity {
		return v.ReservationAffinities
	}).(GetInstanceTemplateReservationAffinityArrayOutput)
}

func (o LookupInstanceTemplateResultOutput) ResourceManagerTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) map[string]string { return v.ResourceManagerTags }).(pulumi.StringMapOutput)
}

// (Optional) -- A list of short names of resource policies to attach to this disk for automatic snapshot creations. Currently a max of 1 resource policy is supported.
func (o LookupInstanceTemplateResultOutput) ResourcePolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []string { return v.ResourcePolicies }).(pulumi.StringArrayOutput)
}

// The scheduling strategy to use. More details about
// this configuration option are detailed below.
func (o LookupInstanceTemplateResultOutput) Schedulings() GetInstanceTemplateSchedulingArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []GetInstanceTemplateScheduling { return v.Schedulings }).(GetInstanceTemplateSchedulingArrayOutput)
}

// The URI of the created resource.
func (o LookupInstanceTemplateResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// A special URI of the created resource that uniquely identifies this instance template with the following format: `projects/{{project}}/global/instanceTemplates/{{name}}?uniqueId={{uniqueId}}`
// Referencing an instance template via this attribute prevents Time of Check to Time of Use attacks when the instance template resides in a shared/untrusted environment.
func (o LookupInstanceTemplateResultOutput) SelfLinkUnique() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) *string { return v.SelfLinkUnique }).(pulumi.StringPtrOutput)
}

// Service account to attach to the instance. Structure is documented below.
func (o LookupInstanceTemplateResultOutput) ServiceAccounts() GetInstanceTemplateServiceAccountArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []GetInstanceTemplateServiceAccount { return v.ServiceAccounts }).(GetInstanceTemplateServiceAccountArrayOutput)
}

// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
func (o LookupInstanceTemplateResultOutput) ShieldedInstanceConfigs() GetInstanceTemplateShieldedInstanceConfigArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []GetInstanceTemplateShieldedInstanceConfig {
		return v.ShieldedInstanceConfigs
	}).(GetInstanceTemplateShieldedInstanceConfigArrayOutput)
}

// Tags to attach to the instance.
func (o LookupInstanceTemplateResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The unique fingerprint of the tags.
func (o LookupInstanceTemplateResultOutput) TagsFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceTemplateResult) string { return v.TagsFingerprint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceTemplateResultOutput{})
}

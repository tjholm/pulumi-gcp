// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkconnectivity

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The NetworkConnectivity Spoke resource
//
// ## Example Usage
// ### Linked_vpc_network
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/networkconnectivity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			network, err := compute.NewNetwork(ctx, "network", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			basicHub, err := networkconnectivity.NewHub(ctx, "basicHub", &networkconnectivity.HubArgs{
//				Description: pulumi.String("A sample hub"),
//				Labels: pulumi.StringMap{
//					"label-two": pulumi.String("value-one"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networkconnectivity.NewSpoke(ctx, "primary", &networkconnectivity.SpokeArgs{
//				Location:    pulumi.String("global"),
//				Description: pulumi.String("A sample spoke with a linked routher appliance instance"),
//				Labels: pulumi.StringMap{
//					"label-one": pulumi.String("value-one"),
//				},
//				Hub: basicHub.ID(),
//				LinkedVpcNetwork: &networkconnectivity.SpokeLinkedVpcNetworkArgs{
//					ExcludeExportRanges: pulumi.StringArray{
//						pulumi.String("198.51.100.0/24"),
//						pulumi.String("10.10.0.0/16"),
//					},
//					Uri: network.SelfLink,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Router_appliance
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/networkconnectivity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			network, err := compute.NewNetwork(ctx, "network", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			subnetwork, err := compute.NewSubnetwork(ctx, "subnetwork", &compute.SubnetworkArgs{
//				IpCidrRange: pulumi.String("10.0.0.0/28"),
//				Region:      pulumi.String("us-west1"),
//				Network:     network.SelfLink,
//			})
//			if err != nil {
//				return err
//			}
//			instance, err := compute.NewInstance(ctx, "instance", &compute.InstanceArgs{
//				MachineType:  pulumi.String("e2-medium"),
//				CanIpForward: pulumi.Bool(true),
//				Zone:         pulumi.String("us-west1-a"),
//				BootDisk: &compute.InstanceBootDiskArgs{
//					InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
//						Image: pulumi.String("projects/debian-cloud/global/images/debian-10-buster-v20210817"),
//					},
//				},
//				NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
//					&compute.InstanceNetworkInterfaceArgs{
//						Subnetwork: subnetwork.Name,
//						NetworkIp:  pulumi.String("10.0.0.2"),
//						AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
//							&compute.InstanceNetworkInterfaceAccessConfigArgs{
//								NetworkTier: pulumi.String("PREMIUM"),
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			basicHub, err := networkconnectivity.NewHub(ctx, "basicHub", &networkconnectivity.HubArgs{
//				Description: pulumi.String("A sample hub"),
//				Labels: pulumi.StringMap{
//					"label-two": pulumi.String("value-one"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networkconnectivity.NewSpoke(ctx, "primary", &networkconnectivity.SpokeArgs{
//				Location:    pulumi.String("us-west1"),
//				Description: pulumi.String("A sample spoke with a linked routher appliance instance"),
//				Labels: pulumi.StringMap{
//					"label-one": pulumi.String("value-one"),
//				},
//				Hub: basicHub.ID(),
//				LinkedRouterApplianceInstances: &networkconnectivity.SpokeLinkedRouterApplianceInstancesArgs{
//					Instances: networkconnectivity.SpokeLinkedRouterApplianceInstancesInstanceArray{
//						&networkconnectivity.SpokeLinkedRouterApplianceInstancesInstanceArgs{
//							VirtualMachine: instance.SelfLink,
//							IpAddress:      pulumi.String("10.0.0.2"),
//						},
//					},
//					SiteToSiteDataTransfer: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Spoke can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:networkconnectivity/spoke:Spoke default projects/{{project}}/locations/{{location}}/spokes/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networkconnectivity/spoke:Spoke default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networkconnectivity/spoke:Spoke default {{location}}/{{name}}
//
// ```
type Spoke struct {
	pulumi.CustomResourceState

	// Output only. The time the spoke was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// An optional description of the spoke.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.MapOutput `pulumi:"effectiveLabels"`
	// Immutable. The URI of the hub that this spoke is attached to.
	Hub pulumi.StringOutput `pulumi:"hub"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
	LinkedInterconnectAttachments SpokeLinkedInterconnectAttachmentsPtrOutput `pulumi:"linkedInterconnectAttachments"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances SpokeLinkedRouterApplianceInstancesPtrOutput `pulumi:"linkedRouterApplianceInstances"`
	// VPC network that is associated with the spoke.
	LinkedVpcNetwork SpokeLinkedVpcNetworkPtrOutput `pulumi:"linkedVpcNetwork"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels SpokeLinkedVpnTunnelsPtrOutput `pulumi:"linkedVpnTunnels"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Immutable. The name of the spoke. Spoke names must be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels pulumi.MapOutput `pulumi:"pulumiLabels"`
	// Output only. The current lifecycle state of this spoke. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING
	State pulumi.StringOutput `pulumi:"state"`
	// Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// Output only. The time the spoke was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSpoke registers a new resource with the given unique name, arguments, and options.
func NewSpoke(ctx *pulumi.Context,
	name string, args *SpokeArgs, opts ...pulumi.ResourceOption) (*Spoke, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Hub == nil {
		return nil, errors.New("invalid value for required argument 'Hub'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Spoke
	err := ctx.RegisterResource("gcp:networkconnectivity/spoke:Spoke", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpoke gets an existing Spoke resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpoke(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpokeState, opts ...pulumi.ResourceOption) (*Spoke, error) {
	var resource Spoke
	err := ctx.ReadResource("gcp:networkconnectivity/spoke:Spoke", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Spoke resources.
type spokeState struct {
	// Output only. The time the spoke was created.
	CreateTime *string `pulumi:"createTime"`
	// An optional description of the spoke.
	Description *string `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels map[string]interface{} `pulumi:"effectiveLabels"`
	// Immutable. The URI of the hub that this spoke is attached to.
	Hub *string `pulumi:"hub"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
	LinkedInterconnectAttachments *SpokeLinkedInterconnectAttachments `pulumi:"linkedInterconnectAttachments"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances *SpokeLinkedRouterApplianceInstances `pulumi:"linkedRouterApplianceInstances"`
	// VPC network that is associated with the spoke.
	LinkedVpcNetwork *SpokeLinkedVpcNetwork `pulumi:"linkedVpcNetwork"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels *SpokeLinkedVpnTunnels `pulumi:"linkedVpnTunnels"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Immutable. The name of the spoke. Spoke names must be unique.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels map[string]interface{} `pulumi:"pulumiLabels"`
	// Output only. The current lifecycle state of this spoke. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING
	State *string `pulumi:"state"`
	// Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.
	UniqueId *string `pulumi:"uniqueId"`
	// Output only. The time the spoke was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type SpokeState struct {
	// Output only. The time the spoke was created.
	CreateTime pulumi.StringPtrInput
	// An optional description of the spoke.
	Description pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.MapInput
	// Immutable. The URI of the hub that this spoke is attached to.
	Hub pulumi.StringPtrInput
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
	LinkedInterconnectAttachments SpokeLinkedInterconnectAttachmentsPtrInput
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances SpokeLinkedRouterApplianceInstancesPtrInput
	// VPC network that is associated with the spoke.
	LinkedVpcNetwork SpokeLinkedVpcNetworkPtrInput
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels SpokeLinkedVpnTunnelsPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Immutable. The name of the spoke. Spoke names must be unique.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels pulumi.MapInput
	// Output only. The current lifecycle state of this spoke. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING
	State pulumi.StringPtrInput
	// Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.
	UniqueId pulumi.StringPtrInput
	// Output only. The time the spoke was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (SpokeState) ElementType() reflect.Type {
	return reflect.TypeOf((*spokeState)(nil)).Elem()
}

type spokeArgs struct {
	// An optional description of the spoke.
	Description *string `pulumi:"description"`
	// Immutable. The URI of the hub that this spoke is attached to.
	Hub string `pulumi:"hub"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
	LinkedInterconnectAttachments *SpokeLinkedInterconnectAttachments `pulumi:"linkedInterconnectAttachments"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances *SpokeLinkedRouterApplianceInstances `pulumi:"linkedRouterApplianceInstances"`
	// VPC network that is associated with the spoke.
	LinkedVpcNetwork *SpokeLinkedVpcNetwork `pulumi:"linkedVpcNetwork"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels *SpokeLinkedVpnTunnels `pulumi:"linkedVpnTunnels"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Immutable. The name of the spoke. Spoke names must be unique.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Spoke resource.
type SpokeArgs struct {
	// An optional description of the spoke.
	Description pulumi.StringPtrInput
	// Immutable. The URI of the hub that this spoke is attached to.
	Hub pulumi.StringInput
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
	LinkedInterconnectAttachments SpokeLinkedInterconnectAttachmentsPtrInput
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances SpokeLinkedRouterApplianceInstancesPtrInput
	// VPC network that is associated with the spoke.
	LinkedVpcNetwork SpokeLinkedVpcNetworkPtrInput
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels SpokeLinkedVpnTunnelsPtrInput
	// The location for the resource
	Location pulumi.StringInput
	// Immutable. The name of the spoke. Spoke names must be unique.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
}

func (SpokeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spokeArgs)(nil)).Elem()
}

type SpokeInput interface {
	pulumi.Input

	ToSpokeOutput() SpokeOutput
	ToSpokeOutputWithContext(ctx context.Context) SpokeOutput
}

func (*Spoke) ElementType() reflect.Type {
	return reflect.TypeOf((**Spoke)(nil)).Elem()
}

func (i *Spoke) ToSpokeOutput() SpokeOutput {
	return i.ToSpokeOutputWithContext(context.Background())
}

func (i *Spoke) ToSpokeOutputWithContext(ctx context.Context) SpokeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpokeOutput)
}

func (i *Spoke) ToOutput(ctx context.Context) pulumix.Output[*Spoke] {
	return pulumix.Output[*Spoke]{
		OutputState: i.ToSpokeOutputWithContext(ctx).OutputState,
	}
}

// SpokeArrayInput is an input type that accepts SpokeArray and SpokeArrayOutput values.
// You can construct a concrete instance of `SpokeArrayInput` via:
//
//	SpokeArray{ SpokeArgs{...} }
type SpokeArrayInput interface {
	pulumi.Input

	ToSpokeArrayOutput() SpokeArrayOutput
	ToSpokeArrayOutputWithContext(context.Context) SpokeArrayOutput
}

type SpokeArray []SpokeInput

func (SpokeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Spoke)(nil)).Elem()
}

func (i SpokeArray) ToSpokeArrayOutput() SpokeArrayOutput {
	return i.ToSpokeArrayOutputWithContext(context.Background())
}

func (i SpokeArray) ToSpokeArrayOutputWithContext(ctx context.Context) SpokeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpokeArrayOutput)
}

func (i SpokeArray) ToOutput(ctx context.Context) pulumix.Output[[]*Spoke] {
	return pulumix.Output[[]*Spoke]{
		OutputState: i.ToSpokeArrayOutputWithContext(ctx).OutputState,
	}
}

// SpokeMapInput is an input type that accepts SpokeMap and SpokeMapOutput values.
// You can construct a concrete instance of `SpokeMapInput` via:
//
//	SpokeMap{ "key": SpokeArgs{...} }
type SpokeMapInput interface {
	pulumi.Input

	ToSpokeMapOutput() SpokeMapOutput
	ToSpokeMapOutputWithContext(context.Context) SpokeMapOutput
}

type SpokeMap map[string]SpokeInput

func (SpokeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Spoke)(nil)).Elem()
}

func (i SpokeMap) ToSpokeMapOutput() SpokeMapOutput {
	return i.ToSpokeMapOutputWithContext(context.Background())
}

func (i SpokeMap) ToSpokeMapOutputWithContext(ctx context.Context) SpokeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpokeMapOutput)
}

func (i SpokeMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Spoke] {
	return pulumix.Output[map[string]*Spoke]{
		OutputState: i.ToSpokeMapOutputWithContext(ctx).OutputState,
	}
}

type SpokeOutput struct{ *pulumi.OutputState }

func (SpokeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Spoke)(nil)).Elem()
}

func (o SpokeOutput) ToSpokeOutput() SpokeOutput {
	return o
}

func (o SpokeOutput) ToSpokeOutputWithContext(ctx context.Context) SpokeOutput {
	return o
}

func (o SpokeOutput) ToOutput(ctx context.Context) pulumix.Output[*Spoke] {
	return pulumix.Output[*Spoke]{
		OutputState: o.OutputState,
	}
}

// Output only. The time the spoke was created.
func (o SpokeOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// An optional description of the spoke.
func (o SpokeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
// clients and services.
func (o SpokeOutput) EffectiveLabels() pulumi.MapOutput {
	return o.ApplyT(func(v *Spoke) pulumi.MapOutput { return v.EffectiveLabels }).(pulumi.MapOutput)
}

// Immutable. The URI of the hub that this spoke is attached to.
func (o SpokeOutput) Hub() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.Hub }).(pulumi.StringOutput)
}

// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o SpokeOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
func (o SpokeOutput) LinkedInterconnectAttachments() SpokeLinkedInterconnectAttachmentsPtrOutput {
	return o.ApplyT(func(v *Spoke) SpokeLinkedInterconnectAttachmentsPtrOutput { return v.LinkedInterconnectAttachments }).(SpokeLinkedInterconnectAttachmentsPtrOutput)
}

// The URIs of linked Router appliance resources
func (o SpokeOutput) LinkedRouterApplianceInstances() SpokeLinkedRouterApplianceInstancesPtrOutput {
	return o.ApplyT(func(v *Spoke) SpokeLinkedRouterApplianceInstancesPtrOutput { return v.LinkedRouterApplianceInstances }).(SpokeLinkedRouterApplianceInstancesPtrOutput)
}

// VPC network that is associated with the spoke.
func (o SpokeOutput) LinkedVpcNetwork() SpokeLinkedVpcNetworkPtrOutput {
	return o.ApplyT(func(v *Spoke) SpokeLinkedVpcNetworkPtrOutput { return v.LinkedVpcNetwork }).(SpokeLinkedVpcNetworkPtrOutput)
}

// The URIs of linked VPN tunnel resources
func (o SpokeOutput) LinkedVpnTunnels() SpokeLinkedVpnTunnelsPtrOutput {
	return o.ApplyT(func(v *Spoke) SpokeLinkedVpnTunnelsPtrOutput { return v.LinkedVpnTunnels }).(SpokeLinkedVpnTunnelsPtrOutput)
}

// The location for the resource
func (o SpokeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Immutable. The name of the spoke. Spoke names must be unique.
func (o SpokeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o SpokeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o SpokeOutput) PulumiLabels() pulumi.MapOutput {
	return o.ApplyT(func(v *Spoke) pulumi.MapOutput { return v.PulumiLabels }).(pulumi.MapOutput)
}

// Output only. The current lifecycle state of this spoke. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING
func (o SpokeOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.
func (o SpokeOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

// Output only. The time the spoke was last updated.
func (o SpokeOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type SpokeArrayOutput struct{ *pulumi.OutputState }

func (SpokeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Spoke)(nil)).Elem()
}

func (o SpokeArrayOutput) ToSpokeArrayOutput() SpokeArrayOutput {
	return o
}

func (o SpokeArrayOutput) ToSpokeArrayOutputWithContext(ctx context.Context) SpokeArrayOutput {
	return o
}

func (o SpokeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Spoke] {
	return pulumix.Output[[]*Spoke]{
		OutputState: o.OutputState,
	}
}

func (o SpokeArrayOutput) Index(i pulumi.IntInput) SpokeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Spoke {
		return vs[0].([]*Spoke)[vs[1].(int)]
	}).(SpokeOutput)
}

type SpokeMapOutput struct{ *pulumi.OutputState }

func (SpokeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Spoke)(nil)).Elem()
}

func (o SpokeMapOutput) ToSpokeMapOutput() SpokeMapOutput {
	return o
}

func (o SpokeMapOutput) ToSpokeMapOutputWithContext(ctx context.Context) SpokeMapOutput {
	return o
}

func (o SpokeMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Spoke] {
	return pulumix.Output[map[string]*Spoke]{
		OutputState: o.OutputState,
	}
}

func (o SpokeMapOutput) MapIndex(k pulumi.StringInput) SpokeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Spoke {
		return vs[0].(map[string]*Spoke)[vs[1].(string)]
	}).(SpokeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpokeInput)(nil)).Elem(), &Spoke{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpokeArrayInput)(nil)).Elem(), SpokeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpokeMapInput)(nil)).Elem(), SpokeMap{})
	pulumi.RegisterOutputType(SpokeOutput{})
	pulumi.RegisterOutputType(SpokeArrayOutput{})
	pulumi.RegisterOutputType(SpokeMapOutput{})
}

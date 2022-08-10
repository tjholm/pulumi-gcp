// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A zone is a subtree of the DNS namespace under one administrative
// responsibility. A ManagedZone is a resource that represents a DNS zone
// hosted by the Cloud DNS service.
//
// To get more information about ManagedZone, see:
//
// * [API documentation](https://cloud.google.com/dns/api/v1/managedZones)
// * How-to Guides
//   - [Managing Zones](https://cloud.google.com/dns/zones/)
//
// ## Example Usage
// ### Dns Managed Zone Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dns.NewManagedZone(ctx, "example-zone", &dns.ManagedZoneArgs{
//				Description: pulumi.String("Example DNS zone"),
//				DnsName:     pulumi.String("my-domain.com."),
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
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
// ### Dns Managed Zone Private
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewNetwork(ctx, "network-1", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewNetwork(ctx, "network-2", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dns.NewManagedZone(ctx, "private-zone", &dns.ManagedZoneArgs{
//				DnsName:     pulumi.String("private.example.com."),
//				Description: pulumi.String("Example private DNS zone"),
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Visibility: pulumi.String("private"),
//				PrivateVisibilityConfig: &dns.ManagedZonePrivateVisibilityConfigArgs{
//					Networks: dns.ManagedZonePrivateVisibilityConfigNetworkArray{
//						&dns.ManagedZonePrivateVisibilityConfigNetworkArgs{
//							NetworkUrl: network_1.ID(),
//						},
//						&dns.ManagedZonePrivateVisibilityConfigNetworkArgs{
//							NetworkUrl: network_2.ID(),
//						},
//					},
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
// ### Dns Managed Zone Private Forwarding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewNetwork(ctx, "network-1", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewNetwork(ctx, "network-2", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dns.NewManagedZone(ctx, "private-zone", &dns.ManagedZoneArgs{
//				DnsName:     pulumi.String("private.example.com."),
//				Description: pulumi.String("Example private DNS zone"),
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Visibility: pulumi.String("private"),
//				PrivateVisibilityConfig: &dns.ManagedZonePrivateVisibilityConfigArgs{
//					Networks: dns.ManagedZonePrivateVisibilityConfigNetworkArray{
//						&dns.ManagedZonePrivateVisibilityConfigNetworkArgs{
//							NetworkUrl: network_1.ID(),
//						},
//						&dns.ManagedZonePrivateVisibilityConfigNetworkArgs{
//							NetworkUrl: network_2.ID(),
//						},
//					},
//				},
//				ForwardingConfig: &dns.ManagedZoneForwardingConfigArgs{
//					TargetNameServers: dns.ManagedZoneForwardingConfigTargetNameServerArray{
//						&dns.ManagedZoneForwardingConfigTargetNameServerArgs{
//							Ipv4Address: pulumi.String("172.16.1.10"),
//						},
//						&dns.ManagedZoneForwardingConfigTargetNameServerArgs{
//							Ipv4Address: pulumi.String("172.16.1.20"),
//						},
//					},
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
// ### Dns Managed Zone Private Peering
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewNetwork(ctx, "network-source", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewNetwork(ctx, "network-target", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dns.NewManagedZone(ctx, "peering-zone", &dns.ManagedZoneArgs{
//				DnsName:     pulumi.String("peering.example.com."),
//				Description: pulumi.String("Example private DNS peering zone"),
//				Visibility:  pulumi.String("private"),
//				PrivateVisibilityConfig: &dns.ManagedZonePrivateVisibilityConfigArgs{
//					Networks: dns.ManagedZonePrivateVisibilityConfigNetworkArray{
//						&dns.ManagedZonePrivateVisibilityConfigNetworkArgs{
//							NetworkUrl: network_source.ID(),
//						},
//					},
//				},
//				PeeringConfig: &dns.ManagedZonePeeringConfigArgs{
//					TargetNetwork: &dns.ManagedZonePeeringConfigTargetNetworkArgs{
//						NetworkUrl: network_target.ID(),
//					},
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
// ### Dns Managed Zone Service Directory
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dns"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicedirectory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := servicedirectory.NewNamespace(ctx, "example", &servicedirectory.NamespaceArgs{
//				NamespaceId: pulumi.String("example"),
//				Location:    pulumi.String("us-central1"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = dns.NewManagedZone(ctx, "sd-zone", &dns.ManagedZoneArgs{
//				DnsName:     pulumi.String("services.example.com."),
//				Description: pulumi.String("Example private DNS Service Directory zone"),
//				Visibility:  pulumi.String("private"),
//				ServiceDirectoryConfig: &dns.ManagedZoneServiceDirectoryConfigArgs{
//					Namespace: &dns.ManagedZoneServiceDirectoryConfigNamespaceArgs{
//						NamespaceUrl: example.ID(),
//					},
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewNetwork(ctx, "network", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			}, pulumi.Provider(google_beta))
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
// # ManagedZone can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:dns/managedZone:ManagedZone default projects/{{project}}/managedZones/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dns/managedZone:ManagedZone default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dns/managedZone:ManagedZone default {{name}}
//
// ```
type ManagedZone struct {
	pulumi.CustomResourceState

	// The time that this resource was created on the server. This is in RFC3339 text format.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringOutput `pulumi:"description"`
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// DNSSEC configuration
	// Structure is documented below.
	DnssecConfig ManagedZoneDnssecConfigPtrOutput `pulumi:"dnssecConfig"`
	// Set this true to delete all records in the zone.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	// Structure is documented below.
	ForwardingConfig ManagedZoneForwardingConfigPtrOutput `pulumi:"forwardingConfig"`
	// A set of key/value label pairs to assign to this ManagedZone.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Unique identifier for the resource; defined by the server.
	ManagedZoneId pulumi.IntOutput `pulumi:"managedZoneId"`
	// User assigned name for this resource.
	// Must be unique within the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// Delegate your managed_zone to these virtual name servers; defined by the server
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	// Structure is documented below.
	PeeringConfig ManagedZonePeeringConfigPtrOutput `pulumi:"peeringConfig"`
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.
	// Structure is documented below.
	PrivateVisibilityConfig ManagedZonePrivateVisibilityConfigPtrOutput `pulumi:"privateVisibilityConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under `privateVisibilityConfig`.
	ReverseLookup pulumi.BoolPtrOutput `pulumi:"reverseLookup"`
	// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.
	// Structure is documented below.
	ServiceDirectoryConfig ManagedZoneServiceDirectoryConfigPtrOutput `pulumi:"serviceDirectoryConfig"`
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	// Default value is `public`.
	// Possible values are `private` and `public`.
	Visibility pulumi.StringPtrOutput `pulumi:"visibility"`
}

// NewManagedZone registers a new resource with the given unique name, arguments, and options.
func NewManagedZone(ctx *pulumi.Context,
	name string, args *ManagedZoneArgs, opts ...pulumi.ResourceOption) (*ManagedZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsName == nil {
		return nil, errors.New("invalid value for required argument 'DnsName'")
	}
	if isZero(args.Description) {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource ManagedZone
	err := ctx.RegisterResource("gcp:dns/managedZone:ManagedZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedZone gets an existing ManagedZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedZoneState, opts ...pulumi.ResourceOption) (*ManagedZone, error) {
	var resource ManagedZone
	err := ctx.ReadResource("gcp:dns/managedZone:ManagedZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedZone resources.
type managedZoneState struct {
	// The time that this resource was created on the server. This is in RFC3339 text format.
	CreationTime *string `pulumi:"creationTime"`
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description *string `pulumi:"description"`
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName *string `pulumi:"dnsName"`
	// DNSSEC configuration
	// Structure is documented below.
	DnssecConfig *ManagedZoneDnssecConfig `pulumi:"dnssecConfig"`
	// Set this true to delete all records in the zone.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	// Structure is documented below.
	ForwardingConfig *ManagedZoneForwardingConfig `pulumi:"forwardingConfig"`
	// A set of key/value label pairs to assign to this ManagedZone.
	Labels map[string]string `pulumi:"labels"`
	// Unique identifier for the resource; defined by the server.
	ManagedZoneId *int `pulumi:"managedZoneId"`
	// User assigned name for this resource.
	// Must be unique within the project.
	Name *string `pulumi:"name"`
	// Delegate your managed_zone to these virtual name servers; defined by the server
	NameServers []string `pulumi:"nameServers"`
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	// Structure is documented below.
	PeeringConfig *ManagedZonePeeringConfig `pulumi:"peeringConfig"`
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.
	// Structure is documented below.
	PrivateVisibilityConfig *ManagedZonePrivateVisibilityConfig `pulumi:"privateVisibilityConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under `privateVisibilityConfig`.
	ReverseLookup *bool `pulumi:"reverseLookup"`
	// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.
	// Structure is documented below.
	ServiceDirectoryConfig *ManagedZoneServiceDirectoryConfig `pulumi:"serviceDirectoryConfig"`
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	// Default value is `public`.
	// Possible values are `private` and `public`.
	Visibility *string `pulumi:"visibility"`
}

type ManagedZoneState struct {
	// The time that this resource was created on the server. This is in RFC3339 text format.
	CreationTime pulumi.StringPtrInput
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringPtrInput
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName pulumi.StringPtrInput
	// DNSSEC configuration
	// Structure is documented below.
	DnssecConfig ManagedZoneDnssecConfigPtrInput
	// Set this true to delete all records in the zone.
	ForceDestroy pulumi.BoolPtrInput
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	// Structure is documented below.
	ForwardingConfig ManagedZoneForwardingConfigPtrInput
	// A set of key/value label pairs to assign to this ManagedZone.
	Labels pulumi.StringMapInput
	// Unique identifier for the resource; defined by the server.
	ManagedZoneId pulumi.IntPtrInput
	// User assigned name for this resource.
	// Must be unique within the project.
	Name pulumi.StringPtrInput
	// Delegate your managed_zone to these virtual name servers; defined by the server
	NameServers pulumi.StringArrayInput
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	// Structure is documented below.
	PeeringConfig ManagedZonePeeringConfigPtrInput
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.
	// Structure is documented below.
	PrivateVisibilityConfig ManagedZonePrivateVisibilityConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under `privateVisibilityConfig`.
	ReverseLookup pulumi.BoolPtrInput
	// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.
	// Structure is documented below.
	ServiceDirectoryConfig ManagedZoneServiceDirectoryConfigPtrInput
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	// Default value is `public`.
	// Possible values are `private` and `public`.
	Visibility pulumi.StringPtrInput
}

func (ManagedZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedZoneState)(nil)).Elem()
}

type managedZoneArgs struct {
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description *string `pulumi:"description"`
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName string `pulumi:"dnsName"`
	// DNSSEC configuration
	// Structure is documented below.
	DnssecConfig *ManagedZoneDnssecConfig `pulumi:"dnssecConfig"`
	// Set this true to delete all records in the zone.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	// Structure is documented below.
	ForwardingConfig *ManagedZoneForwardingConfig `pulumi:"forwardingConfig"`
	// A set of key/value label pairs to assign to this ManagedZone.
	Labels map[string]string `pulumi:"labels"`
	// User assigned name for this resource.
	// Must be unique within the project.
	Name *string `pulumi:"name"`
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	// Structure is documented below.
	PeeringConfig *ManagedZonePeeringConfig `pulumi:"peeringConfig"`
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.
	// Structure is documented below.
	PrivateVisibilityConfig *ManagedZonePrivateVisibilityConfig `pulumi:"privateVisibilityConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under `privateVisibilityConfig`.
	ReverseLookup *bool `pulumi:"reverseLookup"`
	// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.
	// Structure is documented below.
	ServiceDirectoryConfig *ManagedZoneServiceDirectoryConfig `pulumi:"serviceDirectoryConfig"`
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	// Default value is `public`.
	// Possible values are `private` and `public`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a ManagedZone resource.
type ManagedZoneArgs struct {
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringPtrInput
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName pulumi.StringInput
	// DNSSEC configuration
	// Structure is documented below.
	DnssecConfig ManagedZoneDnssecConfigPtrInput
	// Set this true to delete all records in the zone.
	ForceDestroy pulumi.BoolPtrInput
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	// Structure is documented below.
	ForwardingConfig ManagedZoneForwardingConfigPtrInput
	// A set of key/value label pairs to assign to this ManagedZone.
	Labels pulumi.StringMapInput
	// User assigned name for this resource.
	// Must be unique within the project.
	Name pulumi.StringPtrInput
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	// Structure is documented below.
	PeeringConfig ManagedZonePeeringConfigPtrInput
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.
	// Structure is documented below.
	PrivateVisibilityConfig ManagedZonePrivateVisibilityConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under `privateVisibilityConfig`.
	ReverseLookup pulumi.BoolPtrInput
	// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.
	// Structure is documented below.
	ServiceDirectoryConfig ManagedZoneServiceDirectoryConfigPtrInput
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	// Default value is `public`.
	// Possible values are `private` and `public`.
	Visibility pulumi.StringPtrInput
}

func (ManagedZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedZoneArgs)(nil)).Elem()
}

type ManagedZoneInput interface {
	pulumi.Input

	ToManagedZoneOutput() ManagedZoneOutput
	ToManagedZoneOutputWithContext(ctx context.Context) ManagedZoneOutput
}

func (*ManagedZone) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedZone)(nil)).Elem()
}

func (i *ManagedZone) ToManagedZoneOutput() ManagedZoneOutput {
	return i.ToManagedZoneOutputWithContext(context.Background())
}

func (i *ManagedZone) ToManagedZoneOutputWithContext(ctx context.Context) ManagedZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedZoneOutput)
}

// ManagedZoneArrayInput is an input type that accepts ManagedZoneArray and ManagedZoneArrayOutput values.
// You can construct a concrete instance of `ManagedZoneArrayInput` via:
//
//	ManagedZoneArray{ ManagedZoneArgs{...} }
type ManagedZoneArrayInput interface {
	pulumi.Input

	ToManagedZoneArrayOutput() ManagedZoneArrayOutput
	ToManagedZoneArrayOutputWithContext(context.Context) ManagedZoneArrayOutput
}

type ManagedZoneArray []ManagedZoneInput

func (ManagedZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedZone)(nil)).Elem()
}

func (i ManagedZoneArray) ToManagedZoneArrayOutput() ManagedZoneArrayOutput {
	return i.ToManagedZoneArrayOutputWithContext(context.Background())
}

func (i ManagedZoneArray) ToManagedZoneArrayOutputWithContext(ctx context.Context) ManagedZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedZoneArrayOutput)
}

// ManagedZoneMapInput is an input type that accepts ManagedZoneMap and ManagedZoneMapOutput values.
// You can construct a concrete instance of `ManagedZoneMapInput` via:
//
//	ManagedZoneMap{ "key": ManagedZoneArgs{...} }
type ManagedZoneMapInput interface {
	pulumi.Input

	ToManagedZoneMapOutput() ManagedZoneMapOutput
	ToManagedZoneMapOutputWithContext(context.Context) ManagedZoneMapOutput
}

type ManagedZoneMap map[string]ManagedZoneInput

func (ManagedZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedZone)(nil)).Elem()
}

func (i ManagedZoneMap) ToManagedZoneMapOutput() ManagedZoneMapOutput {
	return i.ToManagedZoneMapOutputWithContext(context.Background())
}

func (i ManagedZoneMap) ToManagedZoneMapOutputWithContext(ctx context.Context) ManagedZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedZoneMapOutput)
}

type ManagedZoneOutput struct{ *pulumi.OutputState }

func (ManagedZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedZone)(nil)).Elem()
}

func (o ManagedZoneOutput) ToManagedZoneOutput() ManagedZoneOutput {
	return o
}

func (o ManagedZoneOutput) ToManagedZoneOutputWithContext(ctx context.Context) ManagedZoneOutput {
	return o
}

// The time that this resource was created on the server. This is in RFC3339 text format.
func (o ManagedZoneOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// A textual description field. Defaults to 'Managed by Pulumi'.
func (o ManagedZoneOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The DNS name of this managed zone, for instance "example.com.".
func (o ManagedZoneOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.StringOutput { return v.DnsName }).(pulumi.StringOutput)
}

// DNSSEC configuration
// Structure is documented below.
func (o ManagedZoneOutput) DnssecConfig() ManagedZoneDnssecConfigPtrOutput {
	return o.ApplyT(func(v *ManagedZone) ManagedZoneDnssecConfigPtrOutput { return v.DnssecConfig }).(ManagedZoneDnssecConfigPtrOutput)
}

// Set this true to delete all records in the zone.
func (o ManagedZoneOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// The presence for this field indicates that outbound forwarding is enabled
// for this zone. The value of this field contains the set of destinations
// to forward to.
// Structure is documented below.
func (o ManagedZoneOutput) ForwardingConfig() ManagedZoneForwardingConfigPtrOutput {
	return o.ApplyT(func(v *ManagedZone) ManagedZoneForwardingConfigPtrOutput { return v.ForwardingConfig }).(ManagedZoneForwardingConfigPtrOutput)
}

// A set of key/value label pairs to assign to this ManagedZone.
func (o ManagedZoneOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Unique identifier for the resource; defined by the server.
func (o ManagedZoneOutput) ManagedZoneId() pulumi.IntOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.IntOutput { return v.ManagedZoneId }).(pulumi.IntOutput)
}

// User assigned name for this resource.
// Must be unique within the project.
func (o ManagedZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Delegate your managed_zone to these virtual name servers; defined by the server
func (o ManagedZoneOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.StringArrayOutput { return v.NameServers }).(pulumi.StringArrayOutput)
}

// The presence of this field indicates that DNS Peering is enabled for this
// zone. The value of this field contains the network to peer with.
// Structure is documented below.
func (o ManagedZoneOutput) PeeringConfig() ManagedZonePeeringConfigPtrOutput {
	return o.ApplyT(func(v *ManagedZone) ManagedZonePeeringConfigPtrOutput { return v.PeeringConfig }).(ManagedZonePeeringConfigPtrOutput)
}

// For privately visible zones, the set of Virtual Private Cloud
// resources that the zone is visible from.
// Structure is documented below.
func (o ManagedZoneOutput) PrivateVisibilityConfig() ManagedZonePrivateVisibilityConfigPtrOutput {
	return o.ApplyT(func(v *ManagedZone) ManagedZonePrivateVisibilityConfigPtrOutput { return v.PrivateVisibilityConfig }).(ManagedZonePrivateVisibilityConfigPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o ManagedZoneOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
// lookup queries using automatically configured records for VPC resources. This only applies
// to networks listed under `privateVisibilityConfig`.
func (o ManagedZoneOutput) ReverseLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.BoolPtrOutput { return v.ReverseLookup }).(pulumi.BoolPtrOutput)
}

// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.
// Structure is documented below.
func (o ManagedZoneOutput) ServiceDirectoryConfig() ManagedZoneServiceDirectoryConfigPtrOutput {
	return o.ApplyT(func(v *ManagedZone) ManagedZoneServiceDirectoryConfigPtrOutput { return v.ServiceDirectoryConfig }).(ManagedZoneServiceDirectoryConfigPtrOutput)
}

// The zone's visibility: public zones are exposed to the Internet,
// while private zones are visible only to Virtual Private Cloud resources.
// Default value is `public`.
// Possible values are `private` and `public`.
func (o ManagedZoneOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedZone) pulumi.StringPtrOutput { return v.Visibility }).(pulumi.StringPtrOutput)
}

type ManagedZoneArrayOutput struct{ *pulumi.OutputState }

func (ManagedZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedZone)(nil)).Elem()
}

func (o ManagedZoneArrayOutput) ToManagedZoneArrayOutput() ManagedZoneArrayOutput {
	return o
}

func (o ManagedZoneArrayOutput) ToManagedZoneArrayOutputWithContext(ctx context.Context) ManagedZoneArrayOutput {
	return o
}

func (o ManagedZoneArrayOutput) Index(i pulumi.IntInput) ManagedZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManagedZone {
		return vs[0].([]*ManagedZone)[vs[1].(int)]
	}).(ManagedZoneOutput)
}

type ManagedZoneMapOutput struct{ *pulumi.OutputState }

func (ManagedZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedZone)(nil)).Elem()
}

func (o ManagedZoneMapOutput) ToManagedZoneMapOutput() ManagedZoneMapOutput {
	return o
}

func (o ManagedZoneMapOutput) ToManagedZoneMapOutputWithContext(ctx context.Context) ManagedZoneMapOutput {
	return o
}

func (o ManagedZoneMapOutput) MapIndex(k pulumi.StringInput) ManagedZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManagedZone {
		return vs[0].(map[string]*ManagedZone)[vs[1].(string)]
	}).(ManagedZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedZoneInput)(nil)).Elem(), &ManagedZone{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedZoneArrayInput)(nil)).Elem(), ManagedZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedZoneMapInput)(nil)).Elem(), ManagedZoneMap{})
	pulumi.RegisterOutputType(ManagedZoneOutput{})
	pulumi.RegisterOutputType(ManagedZoneArrayOutput{})
	pulumi.RegisterOutputType(ManagedZoneMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package certificatemanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CertificateMapEntry is a list of certificate configurations,
// that have been issued for a particular hostname
//
// ## Example Usage
// ### Certificate Manager Certificate Map Entry Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificatemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			certificateMap, err := certificatemanager.NewCertificateMap(ctx, "certificateMap", &certificatemanager.CertificateMapArgs{
//				Description: pulumi.String("My acceptance test certificate map"),
//				Labels: pulumi.StringMap{
//					"terraform": pulumi.String("true"),
//					"acc-test":  pulumi.String("true"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			instance, err := certificatemanager.NewDnsAuthorization(ctx, "instance", &certificatemanager.DnsAuthorizationArgs{
//				Description: pulumi.String("The default dnss"),
//				Domain:      pulumi.String("subdomain.hashicorptest.com"),
//			})
//			if err != nil {
//				return err
//			}
//			instance2, err := certificatemanager.NewDnsAuthorization(ctx, "instance2", &certificatemanager.DnsAuthorizationArgs{
//				Description: pulumi.String("The default dnss"),
//				Domain:      pulumi.String("subdomain2.hashicorptest.com"),
//			})
//			if err != nil {
//				return err
//			}
//			certificate, err := certificatemanager.NewCertificate(ctx, "certificate", &certificatemanager.CertificateArgs{
//				Description: pulumi.String("The default cert"),
//				Scope:       pulumi.String("DEFAULT"),
//				Managed: &certificatemanager.CertificateManagedArgs{
//					Domains: pulumi.StringArray{
//						instance.Domain,
//						instance2.Domain,
//					},
//					DnsAuthorizations: pulumi.StringArray{
//						instance.ID(),
//						instance2.ID(),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = certificatemanager.NewCertificateMapEntry(ctx, "default", &certificatemanager.CertificateMapEntryArgs{
//				Description: pulumi.String("My acceptance test certificate map entry"),
//				Map:         certificateMap.Name,
//				Labels: pulumi.StringMap{
//					"terraform": pulumi.String("true"),
//					"acc-test":  pulumi.String("true"),
//				},
//				Certificates: pulumi.StringArray{
//					certificate.ID(),
//				},
//				Matcher: pulumi.String("PRIMARY"),
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
// # CertificateMapEntry can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/certificateMapEntry:CertificateMapEntry default projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/certificateMapEntry:CertificateMapEntry default {{project}}/{{map}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/certificateMapEntry:CertificateMapEntry default {{map}}/{{name}}
//
// ```
type CertificateMapEntry struct {
	pulumi.CustomResourceState

	// A set of Certificates defines for the given hostname.
	// There can be defined up to fifteen certificates in each Certificate Map Entry.
	// Each certificate must match pattern projects/*/locations/*/certificates/*.
	Certificates pulumi.StringArrayOutput `pulumi:"certificates"`
	// Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A human-readable description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	// selecting a proper certificate.
	Hostname pulumi.StringPtrOutput `pulumi:"hostname"`
	// Set of labels associated with a Certificate Map Entry.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A map entry that is inputted into the cetrificate map
	Map pulumi.StringOutput `pulumi:"map"`
	// A predefined matcher for particular cases, other than SNI selection
	Matcher pulumi.StringPtrOutput `pulumi:"matcher"`
	// A user-defined name of the Certificate Map Entry. Certificate Map Entry
	// names must be unique globally and match pattern
	// 'projects/*/locations/*/certificateMaps/*/certificateMapEntries/*'
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A serving state of this Certificate Map Entry.
	State pulumi.StringOutput `pulumi:"state"`
	// Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCertificateMapEntry registers a new resource with the given unique name, arguments, and options.
func NewCertificateMapEntry(ctx *pulumi.Context,
	name string, args *CertificateMapEntryArgs, opts ...pulumi.ResourceOption) (*CertificateMapEntry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Certificates == nil {
		return nil, errors.New("invalid value for required argument 'Certificates'")
	}
	if args.Map == nil {
		return nil, errors.New("invalid value for required argument 'Map'")
	}
	var resource CertificateMapEntry
	err := ctx.RegisterResource("gcp:certificatemanager/certificateMapEntry:CertificateMapEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateMapEntry gets an existing CertificateMapEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateMapEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateMapEntryState, opts ...pulumi.ResourceOption) (*CertificateMapEntry, error) {
	var resource CertificateMapEntry
	err := ctx.ReadResource("gcp:certificatemanager/certificateMapEntry:CertificateMapEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateMapEntry resources.
type certificateMapEntryState struct {
	// A set of Certificates defines for the given hostname.
	// There can be defined up to fifteen certificates in each Certificate Map Entry.
	// Each certificate must match pattern projects/*/locations/*/certificates/*.
	Certificates []string `pulumi:"certificates"`
	// Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	// selecting a proper certificate.
	Hostname *string `pulumi:"hostname"`
	// Set of labels associated with a Certificate Map Entry.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// A map entry that is inputted into the cetrificate map
	Map *string `pulumi:"map"`
	// A predefined matcher for particular cases, other than SNI selection
	Matcher *string `pulumi:"matcher"`
	// A user-defined name of the Certificate Map Entry. Certificate Map Entry
	// names must be unique globally and match pattern
	// 'projects/*/locations/*/certificateMaps/*/certificateMapEntries/*'
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A serving state of this Certificate Map Entry.
	State *string `pulumi:"state"`
	// Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type CertificateMapEntryState struct {
	// A set of Certificates defines for the given hostname.
	// There can be defined up to fifteen certificates in each Certificate Map Entry.
	// Each certificate must match pattern projects/*/locations/*/certificates/*.
	Certificates pulumi.StringArrayInput
	// Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	// selecting a proper certificate.
	Hostname pulumi.StringPtrInput
	// Set of labels associated with a Certificate Map Entry.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// A map entry that is inputted into the cetrificate map
	Map pulumi.StringPtrInput
	// A predefined matcher for particular cases, other than SNI selection
	Matcher pulumi.StringPtrInput
	// A user-defined name of the Certificate Map Entry. Certificate Map Entry
	// names must be unique globally and match pattern
	// 'projects/*/locations/*/certificateMaps/*/certificateMapEntries/*'
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A serving state of this Certificate Map Entry.
	State pulumi.StringPtrInput
	// Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (CertificateMapEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateMapEntryState)(nil)).Elem()
}

type certificateMapEntryArgs struct {
	// A set of Certificates defines for the given hostname.
	// There can be defined up to fifteen certificates in each Certificate Map Entry.
	// Each certificate must match pattern projects/*/locations/*/certificates/*.
	Certificates []string `pulumi:"certificates"`
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	// selecting a proper certificate.
	Hostname *string `pulumi:"hostname"`
	// Set of labels associated with a Certificate Map Entry.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// A map entry that is inputted into the cetrificate map
	Map string `pulumi:"map"`
	// A predefined matcher for particular cases, other than SNI selection
	Matcher *string `pulumi:"matcher"`
	// A user-defined name of the Certificate Map Entry. Certificate Map Entry
	// names must be unique globally and match pattern
	// 'projects/*/locations/*/certificateMaps/*/certificateMapEntries/*'
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a CertificateMapEntry resource.
type CertificateMapEntryArgs struct {
	// A set of Certificates defines for the given hostname.
	// There can be defined up to fifteen certificates in each Certificate Map Entry.
	// Each certificate must match pattern projects/*/locations/*/certificates/*.
	Certificates pulumi.StringArrayInput
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	// selecting a proper certificate.
	Hostname pulumi.StringPtrInput
	// Set of labels associated with a Certificate Map Entry.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// A map entry that is inputted into the cetrificate map
	Map pulumi.StringInput
	// A predefined matcher for particular cases, other than SNI selection
	Matcher pulumi.StringPtrInput
	// A user-defined name of the Certificate Map Entry. Certificate Map Entry
	// names must be unique globally and match pattern
	// 'projects/*/locations/*/certificateMaps/*/certificateMapEntries/*'
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (CertificateMapEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateMapEntryArgs)(nil)).Elem()
}

type CertificateMapEntryInput interface {
	pulumi.Input

	ToCertificateMapEntryOutput() CertificateMapEntryOutput
	ToCertificateMapEntryOutputWithContext(ctx context.Context) CertificateMapEntryOutput
}

func (*CertificateMapEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateMapEntry)(nil)).Elem()
}

func (i *CertificateMapEntry) ToCertificateMapEntryOutput() CertificateMapEntryOutput {
	return i.ToCertificateMapEntryOutputWithContext(context.Background())
}

func (i *CertificateMapEntry) ToCertificateMapEntryOutputWithContext(ctx context.Context) CertificateMapEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapEntryOutput)
}

// CertificateMapEntryArrayInput is an input type that accepts CertificateMapEntryArray and CertificateMapEntryArrayOutput values.
// You can construct a concrete instance of `CertificateMapEntryArrayInput` via:
//
//	CertificateMapEntryArray{ CertificateMapEntryArgs{...} }
type CertificateMapEntryArrayInput interface {
	pulumi.Input

	ToCertificateMapEntryArrayOutput() CertificateMapEntryArrayOutput
	ToCertificateMapEntryArrayOutputWithContext(context.Context) CertificateMapEntryArrayOutput
}

type CertificateMapEntryArray []CertificateMapEntryInput

func (CertificateMapEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateMapEntry)(nil)).Elem()
}

func (i CertificateMapEntryArray) ToCertificateMapEntryArrayOutput() CertificateMapEntryArrayOutput {
	return i.ToCertificateMapEntryArrayOutputWithContext(context.Background())
}

func (i CertificateMapEntryArray) ToCertificateMapEntryArrayOutputWithContext(ctx context.Context) CertificateMapEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapEntryArrayOutput)
}

// CertificateMapEntryMapInput is an input type that accepts CertificateMapEntryMap and CertificateMapEntryMapOutput values.
// You can construct a concrete instance of `CertificateMapEntryMapInput` via:
//
//	CertificateMapEntryMap{ "key": CertificateMapEntryArgs{...} }
type CertificateMapEntryMapInput interface {
	pulumi.Input

	ToCertificateMapEntryMapOutput() CertificateMapEntryMapOutput
	ToCertificateMapEntryMapOutputWithContext(context.Context) CertificateMapEntryMapOutput
}

type CertificateMapEntryMap map[string]CertificateMapEntryInput

func (CertificateMapEntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateMapEntry)(nil)).Elem()
}

func (i CertificateMapEntryMap) ToCertificateMapEntryMapOutput() CertificateMapEntryMapOutput {
	return i.ToCertificateMapEntryMapOutputWithContext(context.Background())
}

func (i CertificateMapEntryMap) ToCertificateMapEntryMapOutputWithContext(ctx context.Context) CertificateMapEntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapEntryMapOutput)
}

type CertificateMapEntryOutput struct{ *pulumi.OutputState }

func (CertificateMapEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateMapEntry)(nil)).Elem()
}

func (o CertificateMapEntryOutput) ToCertificateMapEntryOutput() CertificateMapEntryOutput {
	return o
}

func (o CertificateMapEntryOutput) ToCertificateMapEntryOutputWithContext(ctx context.Context) CertificateMapEntryOutput {
	return o
}

// A set of Certificates defines for the given hostname.
// There can be defined up to fifteen certificates in each Certificate Map Entry.
// Each certificate must match pattern projects/*/locations/*/certificates/*.
func (o CertificateMapEntryOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringArrayOutput { return v.Certificates }).(pulumi.StringArrayOutput)
}

// Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
// with nanosecond resolution and up to nine fractional digits.
// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o CertificateMapEntryOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A human-readable description of the resource.
func (o CertificateMapEntryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
// selecting a proper certificate.
func (o CertificateMapEntryOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

// Set of labels associated with a Certificate Map Entry.
// An object containing a list of "key": value pairs.
// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
func (o CertificateMapEntryOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// A map entry that is inputted into the cetrificate map
func (o CertificateMapEntryOutput) Map() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringOutput { return v.Map }).(pulumi.StringOutput)
}

// A predefined matcher for particular cases, other than SNI selection
func (o CertificateMapEntryOutput) Matcher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringPtrOutput { return v.Matcher }).(pulumi.StringPtrOutput)
}

// A user-defined name of the Certificate Map Entry. Certificate Map Entry
// names must be unique globally and match pattern
// 'projects/*/locations/*/certificateMaps/*/certificateMapEntries/*'
func (o CertificateMapEntryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o CertificateMapEntryOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A serving state of this Certificate Map Entry.
func (o CertificateMapEntryOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
// with nanosecond resolution and up to nine fractional digits.
// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o CertificateMapEntryOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateMapEntry) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type CertificateMapEntryArrayOutput struct{ *pulumi.OutputState }

func (CertificateMapEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateMapEntry)(nil)).Elem()
}

func (o CertificateMapEntryArrayOutput) ToCertificateMapEntryArrayOutput() CertificateMapEntryArrayOutput {
	return o
}

func (o CertificateMapEntryArrayOutput) ToCertificateMapEntryArrayOutputWithContext(ctx context.Context) CertificateMapEntryArrayOutput {
	return o
}

func (o CertificateMapEntryArrayOutput) Index(i pulumi.IntInput) CertificateMapEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertificateMapEntry {
		return vs[0].([]*CertificateMapEntry)[vs[1].(int)]
	}).(CertificateMapEntryOutput)
}

type CertificateMapEntryMapOutput struct{ *pulumi.OutputState }

func (CertificateMapEntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateMapEntry)(nil)).Elem()
}

func (o CertificateMapEntryMapOutput) ToCertificateMapEntryMapOutput() CertificateMapEntryMapOutput {
	return o
}

func (o CertificateMapEntryMapOutput) ToCertificateMapEntryMapOutputWithContext(ctx context.Context) CertificateMapEntryMapOutput {
	return o
}

func (o CertificateMapEntryMapOutput) MapIndex(k pulumi.StringInput) CertificateMapEntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertificateMapEntry {
		return vs[0].(map[string]*CertificateMapEntry)[vs[1].(string)]
	}).(CertificateMapEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapEntryInput)(nil)).Elem(), &CertificateMapEntry{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapEntryArrayInput)(nil)).Elem(), CertificateMapEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapEntryMapInput)(nil)).Elem(), CertificateMapEntryMap{})
	pulumi.RegisterOutputType(CertificateMapEntryOutput{})
	pulumi.RegisterOutputType(CertificateMapEntryArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapEntryMapOutput{})
}

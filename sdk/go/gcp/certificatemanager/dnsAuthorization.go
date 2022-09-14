// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package certificatemanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DnsAuthorization represents a HTTP-reachable backend for a DnsAuthorization.
//
// ## Example Usage
//
// ## Import
//
// # DnsAuthorization can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/dnsAuthorization:DnsAuthorization default projects/{{project}}/locations/global/dnsAuthorizations/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/dnsAuthorization:DnsAuthorization default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/dnsAuthorization:DnsAuthorization default {{name}}
//
// ```
type DnsAuthorization struct {
	pulumi.CustomResourceState

	// A human-readable description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The structure describing the DNS Resource Record that needs to be added to DNS configuration for the authorization to be
	// usable by certificate.
	DnsResourceRecords DnsAuthorizationDnsResourceRecordArrayOutput `pulumi:"dnsResourceRecords"`
	// A domain which is being authorized. A DnsAuthorization resource covers a
	// single domain and its wildcard, e.g. authorization for "example.com" can
	// be used to issue certificates for "example.com" and "*.example.com".
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Set of label tags associated with the DNS Authorization resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewDnsAuthorization registers a new resource with the given unique name, arguments, and options.
func NewDnsAuthorization(ctx *pulumi.Context,
	name string, args *DnsAuthorizationArgs, opts ...pulumi.ResourceOption) (*DnsAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	var resource DnsAuthorization
	err := ctx.RegisterResource("gcp:certificatemanager/dnsAuthorization:DnsAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsAuthorization gets an existing DnsAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsAuthorizationState, opts ...pulumi.ResourceOption) (*DnsAuthorization, error) {
	var resource DnsAuthorization
	err := ctx.ReadResource("gcp:certificatemanager/dnsAuthorization:DnsAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsAuthorization resources.
type dnsAuthorizationState struct {
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// The structure describing the DNS Resource Record that needs to be added to DNS configuration for the authorization to be
	// usable by certificate.
	DnsResourceRecords []DnsAuthorizationDnsResourceRecord `pulumi:"dnsResourceRecords"`
	// A domain which is being authorized. A DnsAuthorization resource covers a
	// single domain and its wildcard, e.g. authorization for "example.com" can
	// be used to issue certificates for "example.com" and "*.example.com".
	Domain *string `pulumi:"domain"`
	// Set of label tags associated with the DNS Authorization resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type DnsAuthorizationState struct {
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// The structure describing the DNS Resource Record that needs to be added to DNS configuration for the authorization to be
	// usable by certificate.
	DnsResourceRecords DnsAuthorizationDnsResourceRecordArrayInput
	// A domain which is being authorized. A DnsAuthorization resource covers a
	// single domain and its wildcard, e.g. authorization for "example.com" can
	// be used to issue certificates for "example.com" and "*.example.com".
	Domain pulumi.StringPtrInput
	// Set of label tags associated with the DNS Authorization resource.
	Labels pulumi.StringMapInput
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DnsAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsAuthorizationState)(nil)).Elem()
}

type dnsAuthorizationArgs struct {
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// A domain which is being authorized. A DnsAuthorization resource covers a
	// single domain and its wildcard, e.g. authorization for "example.com" can
	// be used to issue certificates for "example.com" and "*.example.com".
	Domain string `pulumi:"domain"`
	// Set of label tags associated with the DNS Authorization resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a DnsAuthorization resource.
type DnsAuthorizationArgs struct {
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// A domain which is being authorized. A DnsAuthorization resource covers a
	// single domain and its wildcard, e.g. authorization for "example.com" can
	// be used to issue certificates for "example.com" and "*.example.com".
	Domain pulumi.StringInput
	// Set of label tags associated with the DNS Authorization resource.
	Labels pulumi.StringMapInput
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DnsAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsAuthorizationArgs)(nil)).Elem()
}

type DnsAuthorizationInput interface {
	pulumi.Input

	ToDnsAuthorizationOutput() DnsAuthorizationOutput
	ToDnsAuthorizationOutputWithContext(ctx context.Context) DnsAuthorizationOutput
}

func (*DnsAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsAuthorization)(nil)).Elem()
}

func (i *DnsAuthorization) ToDnsAuthorizationOutput() DnsAuthorizationOutput {
	return i.ToDnsAuthorizationOutputWithContext(context.Background())
}

func (i *DnsAuthorization) ToDnsAuthorizationOutputWithContext(ctx context.Context) DnsAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsAuthorizationOutput)
}

// DnsAuthorizationArrayInput is an input type that accepts DnsAuthorizationArray and DnsAuthorizationArrayOutput values.
// You can construct a concrete instance of `DnsAuthorizationArrayInput` via:
//
//	DnsAuthorizationArray{ DnsAuthorizationArgs{...} }
type DnsAuthorizationArrayInput interface {
	pulumi.Input

	ToDnsAuthorizationArrayOutput() DnsAuthorizationArrayOutput
	ToDnsAuthorizationArrayOutputWithContext(context.Context) DnsAuthorizationArrayOutput
}

type DnsAuthorizationArray []DnsAuthorizationInput

func (DnsAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsAuthorization)(nil)).Elem()
}

func (i DnsAuthorizationArray) ToDnsAuthorizationArrayOutput() DnsAuthorizationArrayOutput {
	return i.ToDnsAuthorizationArrayOutputWithContext(context.Background())
}

func (i DnsAuthorizationArray) ToDnsAuthorizationArrayOutputWithContext(ctx context.Context) DnsAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsAuthorizationArrayOutput)
}

// DnsAuthorizationMapInput is an input type that accepts DnsAuthorizationMap and DnsAuthorizationMapOutput values.
// You can construct a concrete instance of `DnsAuthorizationMapInput` via:
//
//	DnsAuthorizationMap{ "key": DnsAuthorizationArgs{...} }
type DnsAuthorizationMapInput interface {
	pulumi.Input

	ToDnsAuthorizationMapOutput() DnsAuthorizationMapOutput
	ToDnsAuthorizationMapOutputWithContext(context.Context) DnsAuthorizationMapOutput
}

type DnsAuthorizationMap map[string]DnsAuthorizationInput

func (DnsAuthorizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsAuthorization)(nil)).Elem()
}

func (i DnsAuthorizationMap) ToDnsAuthorizationMapOutput() DnsAuthorizationMapOutput {
	return i.ToDnsAuthorizationMapOutputWithContext(context.Background())
}

func (i DnsAuthorizationMap) ToDnsAuthorizationMapOutputWithContext(ctx context.Context) DnsAuthorizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsAuthorizationMapOutput)
}

type DnsAuthorizationOutput struct{ *pulumi.OutputState }

func (DnsAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsAuthorization)(nil)).Elem()
}

func (o DnsAuthorizationOutput) ToDnsAuthorizationOutput() DnsAuthorizationOutput {
	return o
}

func (o DnsAuthorizationOutput) ToDnsAuthorizationOutputWithContext(ctx context.Context) DnsAuthorizationOutput {
	return o
}

// A human-readable description of the resource.
func (o DnsAuthorizationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsAuthorization) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The structure describing the DNS Resource Record that needs to be added to DNS configuration for the authorization to be
// usable by certificate.
func (o DnsAuthorizationOutput) DnsResourceRecords() DnsAuthorizationDnsResourceRecordArrayOutput {
	return o.ApplyT(func(v *DnsAuthorization) DnsAuthorizationDnsResourceRecordArrayOutput { return v.DnsResourceRecords }).(DnsAuthorizationDnsResourceRecordArrayOutput)
}

// A domain which is being authorized. A DnsAuthorization resource covers a
// single domain and its wildcard, e.g. authorization for "example.com" can
// be used to issue certificates for "example.com" and "*.example.com".
func (o DnsAuthorizationOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsAuthorization) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Set of label tags associated with the DNS Authorization resource.
func (o DnsAuthorizationOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DnsAuthorization) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource; provided by the client when the resource is created.
// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
// and all following characters must be a dash, underscore, letter or digit.
func (o DnsAuthorizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsAuthorization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o DnsAuthorizationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsAuthorization) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type DnsAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (DnsAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsAuthorization)(nil)).Elem()
}

func (o DnsAuthorizationArrayOutput) ToDnsAuthorizationArrayOutput() DnsAuthorizationArrayOutput {
	return o
}

func (o DnsAuthorizationArrayOutput) ToDnsAuthorizationArrayOutputWithContext(ctx context.Context) DnsAuthorizationArrayOutput {
	return o
}

func (o DnsAuthorizationArrayOutput) Index(i pulumi.IntInput) DnsAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsAuthorization {
		return vs[0].([]*DnsAuthorization)[vs[1].(int)]
	}).(DnsAuthorizationOutput)
}

type DnsAuthorizationMapOutput struct{ *pulumi.OutputState }

func (DnsAuthorizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsAuthorization)(nil)).Elem()
}

func (o DnsAuthorizationMapOutput) ToDnsAuthorizationMapOutput() DnsAuthorizationMapOutput {
	return o
}

func (o DnsAuthorizationMapOutput) ToDnsAuthorizationMapOutputWithContext(ctx context.Context) DnsAuthorizationMapOutput {
	return o
}

func (o DnsAuthorizationMapOutput) MapIndex(k pulumi.StringInput) DnsAuthorizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsAuthorization {
		return vs[0].(map[string]*DnsAuthorization)[vs[1].(string)]
	}).(DnsAuthorizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsAuthorizationInput)(nil)).Elem(), &DnsAuthorization{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsAuthorizationArrayInput)(nil)).Elem(), DnsAuthorizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsAuthorizationMapInput)(nil)).Elem(), DnsAuthorizationMap{})
	pulumi.RegisterOutputType(DnsAuthorizationOutput{})
	pulumi.RegisterOutputType(DnsAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(DnsAuthorizationMapOutput{})
}

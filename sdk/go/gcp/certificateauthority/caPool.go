// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package certificateauthority

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A CaPool represents a group of CertificateAuthorities that form a trust anchor. A CaPool can be used to manage
// issuance policies for one or more CertificateAuthority resources and to rotate CA certificates in and out of the
// trust anchor.
//
// ## Example Usage
// ### Privateca Capool Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/certificateauthority"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := certificateauthority.NewCaPool(ctx, "default", &certificateauthority.CaPoolArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Location: pulumi.String("us-central1"),
//				PublishingOptions: &certificateauthority.CaPoolPublishingOptionsArgs{
//					PublishCaCert: pulumi.Bool(true),
//					PublishCrl:    pulumi.Bool(true),
//				},
//				Tier: pulumi.String("ENTERPRISE"),
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
// # CaPool can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:certificateauthority/caPool:CaPool default projects/{{project}}/locations/{{location}}/caPools/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:certificateauthority/caPool:CaPool default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:certificateauthority/caPool:CaPool default {{location}}/{{name}}
//
// ```
type CaPool struct {
	pulumi.CustomResourceState

	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The IssuancePolicy to control how Certificates will be issued from this CaPool.
	// Structure is documented below.
	IssuancePolicy CaPoolIssuancePolicyPtrOutput `pulumi:"issuancePolicy"`
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	//
	// ***
	Location pulumi.StringOutput `pulumi:"location"`
	// The name for this CaPool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	// Structure is documented below.
	PublishingOptions CaPoolPublishingOptionsPtrOutput `pulumi:"publishingOptions"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// The Tier of this CaPool.
	// Possible values are: `ENTERPRISE`, `DEVOPS`.
	Tier pulumi.StringOutput `pulumi:"tier"`
}

// NewCaPool registers a new resource with the given unique name, arguments, and options.
func NewCaPool(ctx *pulumi.Context,
	name string, args *CaPoolArgs, opts ...pulumi.ResourceOption) (*CaPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CaPool
	err := ctx.RegisterResource("gcp:certificateauthority/caPool:CaPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaPool gets an existing CaPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaPoolState, opts ...pulumi.ResourceOption) (*CaPool, error) {
	var resource CaPool
	err := ctx.ReadResource("gcp:certificateauthority/caPool:CaPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaPool resources.
type caPoolState struct {
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The IssuancePolicy to control how Certificates will be issued from this CaPool.
	// Structure is documented below.
	IssuancePolicy *CaPoolIssuancePolicy `pulumi:"issuancePolicy"`
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	//
	// ***
	Location *string `pulumi:"location"`
	// The name for this CaPool.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	// Structure is documented below.
	PublishingOptions *CaPoolPublishingOptions `pulumi:"publishingOptions"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// The Tier of this CaPool.
	// Possible values are: `ENTERPRISE`, `DEVOPS`.
	Tier *string `pulumi:"tier"`
}

type CaPoolState struct {
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapInput
	// The IssuancePolicy to control how Certificates will be issued from this CaPool.
	// Structure is documented below.
	IssuancePolicy CaPoolIssuancePolicyPtrInput
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	//
	// ***
	Location pulumi.StringPtrInput
	// The name for this CaPool.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	// Structure is documented below.
	PublishingOptions CaPoolPublishingOptionsPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// The Tier of this CaPool.
	// Possible values are: `ENTERPRISE`, `DEVOPS`.
	Tier pulumi.StringPtrInput
}

func (CaPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolState)(nil)).Elem()
}

type caPoolArgs struct {
	// The IssuancePolicy to control how Certificates will be issued from this CaPool.
	// Structure is documented below.
	IssuancePolicy *CaPoolIssuancePolicy `pulumi:"issuancePolicy"`
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	//
	// ***
	Location string `pulumi:"location"`
	// The name for this CaPool.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	// Structure is documented below.
	PublishingOptions *CaPoolPublishingOptions `pulumi:"publishingOptions"`
	// The Tier of this CaPool.
	// Possible values are: `ENTERPRISE`, `DEVOPS`.
	Tier string `pulumi:"tier"`
}

// The set of arguments for constructing a CaPool resource.
type CaPoolArgs struct {
	// The IssuancePolicy to control how Certificates will be issued from this CaPool.
	// Structure is documented below.
	IssuancePolicy CaPoolIssuancePolicyPtrInput
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	//
	// ***
	Location pulumi.StringInput
	// The name for this CaPool.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	// Structure is documented below.
	PublishingOptions CaPoolPublishingOptionsPtrInput
	// The Tier of this CaPool.
	// Possible values are: `ENTERPRISE`, `DEVOPS`.
	Tier pulumi.StringInput
}

func (CaPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolArgs)(nil)).Elem()
}

type CaPoolInput interface {
	pulumi.Input

	ToCaPoolOutput() CaPoolOutput
	ToCaPoolOutputWithContext(ctx context.Context) CaPoolOutput
}

func (*CaPool) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPool)(nil)).Elem()
}

func (i *CaPool) ToCaPoolOutput() CaPoolOutput {
	return i.ToCaPoolOutputWithContext(context.Background())
}

func (i *CaPool) ToCaPoolOutputWithContext(ctx context.Context) CaPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolOutput)
}

func (i *CaPool) ToOutput(ctx context.Context) pulumix.Output[*CaPool] {
	return pulumix.Output[*CaPool]{
		OutputState: i.ToCaPoolOutputWithContext(ctx).OutputState,
	}
}

// CaPoolArrayInput is an input type that accepts CaPoolArray and CaPoolArrayOutput values.
// You can construct a concrete instance of `CaPoolArrayInput` via:
//
//	CaPoolArray{ CaPoolArgs{...} }
type CaPoolArrayInput interface {
	pulumi.Input

	ToCaPoolArrayOutput() CaPoolArrayOutput
	ToCaPoolArrayOutputWithContext(context.Context) CaPoolArrayOutput
}

type CaPoolArray []CaPoolInput

func (CaPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CaPool)(nil)).Elem()
}

func (i CaPoolArray) ToCaPoolArrayOutput() CaPoolArrayOutput {
	return i.ToCaPoolArrayOutputWithContext(context.Background())
}

func (i CaPoolArray) ToCaPoolArrayOutputWithContext(ctx context.Context) CaPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolArrayOutput)
}

func (i CaPoolArray) ToOutput(ctx context.Context) pulumix.Output[[]*CaPool] {
	return pulumix.Output[[]*CaPool]{
		OutputState: i.ToCaPoolArrayOutputWithContext(ctx).OutputState,
	}
}

// CaPoolMapInput is an input type that accepts CaPoolMap and CaPoolMapOutput values.
// You can construct a concrete instance of `CaPoolMapInput` via:
//
//	CaPoolMap{ "key": CaPoolArgs{...} }
type CaPoolMapInput interface {
	pulumi.Input

	ToCaPoolMapOutput() CaPoolMapOutput
	ToCaPoolMapOutputWithContext(context.Context) CaPoolMapOutput
}

type CaPoolMap map[string]CaPoolInput

func (CaPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CaPool)(nil)).Elem()
}

func (i CaPoolMap) ToCaPoolMapOutput() CaPoolMapOutput {
	return i.ToCaPoolMapOutputWithContext(context.Background())
}

func (i CaPoolMap) ToCaPoolMapOutputWithContext(ctx context.Context) CaPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolMapOutput)
}

func (i CaPoolMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CaPool] {
	return pulumix.Output[map[string]*CaPool]{
		OutputState: i.ToCaPoolMapOutputWithContext(ctx).OutputState,
	}
}

type CaPoolOutput struct{ *pulumi.OutputState }

func (CaPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPool)(nil)).Elem()
}

func (o CaPoolOutput) ToCaPoolOutput() CaPoolOutput {
	return o
}

func (o CaPoolOutput) ToCaPoolOutputWithContext(ctx context.Context) CaPoolOutput {
	return o
}

func (o CaPoolOutput) ToOutput(ctx context.Context) pulumix.Output[*CaPool] {
	return pulumix.Output[*CaPool]{
		OutputState: o.OutputState,
	}
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
// clients and services.
func (o CaPoolOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CaPool) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The IssuancePolicy to control how Certificates will be issued from this CaPool.
// Structure is documented below.
func (o CaPoolOutput) IssuancePolicy() CaPoolIssuancePolicyPtrOutput {
	return o.ApplyT(func(v *CaPool) CaPoolIssuancePolicyPtrOutput { return v.IssuancePolicy }).(CaPoolIssuancePolicyPtrOutput)
}

// Labels with user-defined metadata.
// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
// "1.3kg", "count": "3" }.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o CaPoolOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CaPool) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Location of the CaPool. A full list of valid locations can be found by
// running `gcloud privateca locations list`.
//
// ***
func (o CaPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name for this CaPool.
func (o CaPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o CaPoolOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPool) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
// Structure is documented below.
func (o CaPoolOutput) PublishingOptions() CaPoolPublishingOptionsPtrOutput {
	return o.ApplyT(func(v *CaPool) CaPoolPublishingOptionsPtrOutput { return v.PublishingOptions }).(CaPoolPublishingOptionsPtrOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o CaPoolOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CaPool) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// The Tier of this CaPool.
// Possible values are: `ENTERPRISE`, `DEVOPS`.
func (o CaPoolOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPool) pulumi.StringOutput { return v.Tier }).(pulumi.StringOutput)
}

type CaPoolArrayOutput struct{ *pulumi.OutputState }

func (CaPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CaPool)(nil)).Elem()
}

func (o CaPoolArrayOutput) ToCaPoolArrayOutput() CaPoolArrayOutput {
	return o
}

func (o CaPoolArrayOutput) ToCaPoolArrayOutputWithContext(ctx context.Context) CaPoolArrayOutput {
	return o
}

func (o CaPoolArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CaPool] {
	return pulumix.Output[[]*CaPool]{
		OutputState: o.OutputState,
	}
}

func (o CaPoolArrayOutput) Index(i pulumi.IntInput) CaPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CaPool {
		return vs[0].([]*CaPool)[vs[1].(int)]
	}).(CaPoolOutput)
}

type CaPoolMapOutput struct{ *pulumi.OutputState }

func (CaPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CaPool)(nil)).Elem()
}

func (o CaPoolMapOutput) ToCaPoolMapOutput() CaPoolMapOutput {
	return o
}

func (o CaPoolMapOutput) ToCaPoolMapOutputWithContext(ctx context.Context) CaPoolMapOutput {
	return o
}

func (o CaPoolMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CaPool] {
	return pulumix.Output[map[string]*CaPool]{
		OutputState: o.OutputState,
	}
}

func (o CaPoolMapOutput) MapIndex(k pulumi.StringInput) CaPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CaPool {
		return vs[0].(map[string]*CaPool)[vs[1].(string)]
	}).(CaPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolInput)(nil)).Elem(), &CaPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolArrayInput)(nil)).Elem(), CaPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolMapInput)(nil)).Elem(), CaPoolMap{})
	pulumi.RegisterOutputType(CaPoolOutput{})
	pulumi.RegisterOutputType(CaPoolArrayOutput{})
	pulumi.RegisterOutputType(CaPoolMapOutput{})
}

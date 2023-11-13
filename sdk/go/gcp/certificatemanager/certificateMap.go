// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package certificatemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// CertificateMap defines a collection of certificate configurations,
// which are usable by any associated target proxies
//
// ## Example Usage
// ### Certificate Manager Certificate Map Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/certificatemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := certificatemanager.NewCertificateMap(ctx, "default", &certificatemanager.CertificateMapArgs{
//				Description: pulumi.String("My acceptance test certificate map"),
//				Labels: pulumi.StringMap{
//					"terraform": pulumi.String("true"),
//					"acc-test":  pulumi.String("true"),
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
// # CertificateMap can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/certificateMap:CertificateMap default projects/{{project}}/locations/global/certificateMaps/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/certificateMap:CertificateMap default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/certificateMap:CertificateMap default {{name}}
//
// ```
type CertificateMapResource struct {
	pulumi.CustomResourceState

	// Creation timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds with up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A human-readable description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// A list of target proxies that use this Certificate Map
	// Structure is documented below.
	GclbTargets CertificateMapGclbTargetArrayOutput `pulumi:"gclbTargets"`
	// Set of labels associated with a Certificate Map resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A user-defined name of the Certificate Map. Certificate Map names must be unique
	// globally and match the pattern `projects/*/locations/*/certificateMaps/*`.
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// Update timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds with up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCertificateMapResource registers a new resource with the given unique name, arguments, and options.
func NewCertificateMapResource(ctx *pulumi.Context,
	name string, args *CertificateMapResourceArgs, opts ...pulumi.ResourceOption) (*CertificateMapResource, error) {
	if args == nil {
		args = &CertificateMapResourceArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CertificateMapResource
	err := ctx.RegisterResource("gcp:certificatemanager/certificateMap:CertificateMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateMapResource gets an existing CertificateMapResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateMapResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateMapResourceState, opts ...pulumi.ResourceOption) (*CertificateMapResource, error) {
	var resource CertificateMapResource
	err := ctx.ReadResource("gcp:certificatemanager/certificateMap:CertificateMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateMapResource resources.
type certificateMapResourceState struct {
	// Creation timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds with up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// A list of target proxies that use this Certificate Map
	// Structure is documented below.
	GclbTargets []CertificateMapGclbTarget `pulumi:"gclbTargets"`
	// Set of labels associated with a Certificate Map resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// A user-defined name of the Certificate Map. Certificate Map names must be unique
	// globally and match the pattern `projects/*/locations/*/certificateMaps/*`.
	//
	// ***
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// Update timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds with up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type CertificateMapResourceState struct {
	// Creation timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds with up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapInput
	// A list of target proxies that use this Certificate Map
	// Structure is documented below.
	GclbTargets CertificateMapGclbTargetArrayInput
	// Set of labels associated with a Certificate Map resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// A user-defined name of the Certificate Map. Certificate Map names must be unique
	// globally and match the pattern `projects/*/locations/*/certificateMaps/*`.
	//
	// ***
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// Update timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds with up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (CertificateMapResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateMapResourceState)(nil)).Elem()
}

type certificateMapResourceArgs struct {
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// Set of labels associated with a Certificate Map resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// A user-defined name of the Certificate Map. Certificate Map names must be unique
	// globally and match the pattern `projects/*/locations/*/certificateMaps/*`.
	//
	// ***
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a CertificateMapResource resource.
type CertificateMapResourceArgs struct {
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// Set of labels associated with a Certificate Map resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// A user-defined name of the Certificate Map. Certificate Map names must be unique
	// globally and match the pattern `projects/*/locations/*/certificateMaps/*`.
	//
	// ***
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (CertificateMapResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateMapResourceArgs)(nil)).Elem()
}

type CertificateMapResourceInput interface {
	pulumi.Input

	ToCertificateMapResourceOutput() CertificateMapResourceOutput
	ToCertificateMapResourceOutputWithContext(ctx context.Context) CertificateMapResourceOutput
}

func (*CertificateMapResource) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateMapResource)(nil)).Elem()
}

func (i *CertificateMapResource) ToCertificateMapResourceOutput() CertificateMapResourceOutput {
	return i.ToCertificateMapResourceOutputWithContext(context.Background())
}

func (i *CertificateMapResource) ToCertificateMapResourceOutputWithContext(ctx context.Context) CertificateMapResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapResourceOutput)
}

func (i *CertificateMapResource) ToOutput(ctx context.Context) pulumix.Output[*CertificateMapResource] {
	return pulumix.Output[*CertificateMapResource]{
		OutputState: i.ToCertificateMapResourceOutputWithContext(ctx).OutputState,
	}
}

// CertificateMapResourceArrayInput is an input type that accepts CertificateMapResourceArray and CertificateMapResourceArrayOutput values.
// You can construct a concrete instance of `CertificateMapResourceArrayInput` via:
//
//	CertificateMapResourceArray{ CertificateMapResourceArgs{...} }
type CertificateMapResourceArrayInput interface {
	pulumi.Input

	ToCertificateMapResourceArrayOutput() CertificateMapResourceArrayOutput
	ToCertificateMapResourceArrayOutputWithContext(context.Context) CertificateMapResourceArrayOutput
}

type CertificateMapResourceArray []CertificateMapResourceInput

func (CertificateMapResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateMapResource)(nil)).Elem()
}

func (i CertificateMapResourceArray) ToCertificateMapResourceArrayOutput() CertificateMapResourceArrayOutput {
	return i.ToCertificateMapResourceArrayOutputWithContext(context.Background())
}

func (i CertificateMapResourceArray) ToCertificateMapResourceArrayOutputWithContext(ctx context.Context) CertificateMapResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapResourceArrayOutput)
}

func (i CertificateMapResourceArray) ToOutput(ctx context.Context) pulumix.Output[[]*CertificateMapResource] {
	return pulumix.Output[[]*CertificateMapResource]{
		OutputState: i.ToCertificateMapResourceArrayOutputWithContext(ctx).OutputState,
	}
}

// CertificateMapResourceMapInput is an input type that accepts CertificateMapResourceMap and CertificateMapResourceMapOutput values.
// You can construct a concrete instance of `CertificateMapResourceMapInput` via:
//
//	CertificateMapResourceMap{ "key": CertificateMapResourceArgs{...} }
type CertificateMapResourceMapInput interface {
	pulumi.Input

	ToCertificateMapResourceMapOutput() CertificateMapResourceMapOutput
	ToCertificateMapResourceMapOutputWithContext(context.Context) CertificateMapResourceMapOutput
}

type CertificateMapResourceMap map[string]CertificateMapResourceInput

func (CertificateMapResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateMapResource)(nil)).Elem()
}

func (i CertificateMapResourceMap) ToCertificateMapResourceMapOutput() CertificateMapResourceMapOutput {
	return i.ToCertificateMapResourceMapOutputWithContext(context.Background())
}

func (i CertificateMapResourceMap) ToCertificateMapResourceMapOutputWithContext(ctx context.Context) CertificateMapResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapResourceMapOutput)
}

func (i CertificateMapResourceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CertificateMapResource] {
	return pulumix.Output[map[string]*CertificateMapResource]{
		OutputState: i.ToCertificateMapResourceMapOutputWithContext(ctx).OutputState,
	}
}

type CertificateMapResourceOutput struct{ *pulumi.OutputState }

func (CertificateMapResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateMapResource)(nil)).Elem()
}

func (o CertificateMapResourceOutput) ToCertificateMapResourceOutput() CertificateMapResourceOutput {
	return o
}

func (o CertificateMapResourceOutput) ToCertificateMapResourceOutputWithContext(ctx context.Context) CertificateMapResourceOutput {
	return o
}

func (o CertificateMapResourceOutput) ToOutput(ctx context.Context) pulumix.Output[*CertificateMapResource] {
	return pulumix.Output[*CertificateMapResource]{
		OutputState: o.OutputState,
	}
}

// Creation timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
// accurate to nanoseconds with up to nine fractional digits.
// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o CertificateMapResourceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateMapResource) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A human-readable description of the resource.
func (o CertificateMapResourceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateMapResource) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
func (o CertificateMapResourceOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateMapResource) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// A list of target proxies that use this Certificate Map
// Structure is documented below.
func (o CertificateMapResourceOutput) GclbTargets() CertificateMapGclbTargetArrayOutput {
	return o.ApplyT(func(v *CertificateMapResource) CertificateMapGclbTargetArrayOutput { return v.GclbTargets }).(CertificateMapGclbTargetArrayOutput)
}

// Set of labels associated with a Certificate Map resource.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o CertificateMapResourceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateMapResource) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// A user-defined name of the Certificate Map. Certificate Map names must be unique
// globally and match the pattern `projects/*/locations/*/certificateMaps/*`.
//
// ***
func (o CertificateMapResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateMapResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o CertificateMapResourceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateMapResource) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o CertificateMapResourceOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateMapResource) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// Update timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
// accurate to nanoseconds with up to nine fractional digits.
// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o CertificateMapResourceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateMapResource) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type CertificateMapResourceArrayOutput struct{ *pulumi.OutputState }

func (CertificateMapResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateMapResource)(nil)).Elem()
}

func (o CertificateMapResourceArrayOutput) ToCertificateMapResourceArrayOutput() CertificateMapResourceArrayOutput {
	return o
}

func (o CertificateMapResourceArrayOutput) ToCertificateMapResourceArrayOutputWithContext(ctx context.Context) CertificateMapResourceArrayOutput {
	return o
}

func (o CertificateMapResourceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CertificateMapResource] {
	return pulumix.Output[[]*CertificateMapResource]{
		OutputState: o.OutputState,
	}
}

func (o CertificateMapResourceArrayOutput) Index(i pulumi.IntInput) CertificateMapResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertificateMapResource {
		return vs[0].([]*CertificateMapResource)[vs[1].(int)]
	}).(CertificateMapResourceOutput)
}

type CertificateMapResourceMapOutput struct{ *pulumi.OutputState }

func (CertificateMapResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateMapResource)(nil)).Elem()
}

func (o CertificateMapResourceMapOutput) ToCertificateMapResourceMapOutput() CertificateMapResourceMapOutput {
	return o
}

func (o CertificateMapResourceMapOutput) ToCertificateMapResourceMapOutputWithContext(ctx context.Context) CertificateMapResourceMapOutput {
	return o
}

func (o CertificateMapResourceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CertificateMapResource] {
	return pulumix.Output[map[string]*CertificateMapResource]{
		OutputState: o.OutputState,
	}
}

func (o CertificateMapResourceMapOutput) MapIndex(k pulumi.StringInput) CertificateMapResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertificateMapResource {
		return vs[0].(map[string]*CertificateMapResource)[vs[1].(string)]
	}).(CertificateMapResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapResourceInput)(nil)).Elem(), &CertificateMapResource{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapResourceArrayInput)(nil)).Elem(), CertificateMapResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapResourceMapInput)(nil)).Elem(), CertificateMapResourceMap{})
	pulumi.RegisterOutputType(CertificateMapResourceOutput{})
	pulumi.RegisterOutputType(CertificateMapResourceArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapResourceMapOutput{})
}

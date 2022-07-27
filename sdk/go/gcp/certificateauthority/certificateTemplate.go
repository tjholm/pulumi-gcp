// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package certificateauthority

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Certificate Authority Service provides reusable and parameterized templates that you can use for common certificate issuance scenarios. A certificate template represents a relatively static and well-defined certificate issuance schema within an organization.  A certificate template can essentially become a full-fledged vertical certificate issuance framework.
//
// For more information, see:
// * [Understanding Certificate Templates](https://cloud.google.com/certificate-authority-service/docs/certificate-template)
// * [Common configurations and Certificate Profiles](https://cloud.google.com/certificate-authority-service/docs/certificate-profile)
// ## Example Usage
//
// ## Import
//
// CertificateTemplate can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default {{location}}/{{name}}
// ```
type CertificateTemplate struct {
	pulumi.CustomResourceState

	// Output only. The time at which this CertificateTemplate was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	IdentityConstraints CertificateTemplateIdentityConstraintsPtrOutput `pulumi:"identityConstraints"`
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baselineValues that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	PassthroughExtensions CertificateTemplatePassthroughExtensionsPtrOutput `pulumi:"passthroughExtensions"`
	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baselineValues for the same properties, the certificate issuance request will fail.
	PredefinedValues CertificateTemplatePredefinedValuesPtrOutput `pulumi:"predefinedValues"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. The time at which this CertificateTemplate was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCertificateTemplate registers a new resource with the given unique name, arguments, and options.
func NewCertificateTemplate(ctx *pulumi.Context,
	name string, args *CertificateTemplateArgs, opts ...pulumi.ResourceOption) (*CertificateTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource CertificateTemplate
	err := ctx.RegisterResource("gcp:certificateauthority/certificateTemplate:CertificateTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateTemplate gets an existing CertificateTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateTemplateState, opts ...pulumi.ResourceOption) (*CertificateTemplate, error) {
	var resource CertificateTemplate
	err := ctx.ReadResource("gcp:certificateauthority/certificateTemplate:CertificateTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateTemplate resources.
type certificateTemplateState struct {
	// Output only. The time at which this CertificateTemplate was created.
	CreateTime *string `pulumi:"createTime"`
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description *string `pulumi:"description"`
	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	IdentityConstraints *CertificateTemplateIdentityConstraints `pulumi:"identityConstraints"`
	// Optional. Labels with user-defined metadata.
	Labels map[string]string `pulumi:"labels"`
	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location *string `pulumi:"location"`
	// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
	Name *string `pulumi:"name"`
	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baselineValues that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	PassthroughExtensions *CertificateTemplatePassthroughExtensions `pulumi:"passthroughExtensions"`
	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baselineValues for the same properties, the certificate issuance request will fail.
	PredefinedValues *CertificateTemplatePredefinedValues `pulumi:"predefinedValues"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Output only. The time at which this CertificateTemplate was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type CertificateTemplateState struct {
	// Output only. The time at which this CertificateTemplate was created.
	CreateTime pulumi.StringPtrInput
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description pulumi.StringPtrInput
	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	IdentityConstraints CertificateTemplateIdentityConstraintsPtrInput
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapInput
	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location pulumi.StringPtrInput
	// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
	Name pulumi.StringPtrInput
	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baselineValues that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	PassthroughExtensions CertificateTemplatePassthroughExtensionsPtrInput
	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baselineValues for the same properties, the certificate issuance request will fail.
	PredefinedValues CertificateTemplatePredefinedValuesPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Output only. The time at which this CertificateTemplate was updated.
	UpdateTime pulumi.StringPtrInput
}

func (CertificateTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateTemplateState)(nil)).Elem()
}

type certificateTemplateArgs struct {
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description *string `pulumi:"description"`
	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	IdentityConstraints *CertificateTemplateIdentityConstraints `pulumi:"identityConstraints"`
	// Optional. Labels with user-defined metadata.
	Labels map[string]string `pulumi:"labels"`
	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location string `pulumi:"location"`
	// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
	Name *string `pulumi:"name"`
	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baselineValues that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	PassthroughExtensions *CertificateTemplatePassthroughExtensions `pulumi:"passthroughExtensions"`
	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baselineValues for the same properties, the certificate issuance request will fail.
	PredefinedValues *CertificateTemplatePredefinedValues `pulumi:"predefinedValues"`
	// The project for the resource
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a CertificateTemplate resource.
type CertificateTemplateArgs struct {
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description pulumi.StringPtrInput
	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	IdentityConstraints CertificateTemplateIdentityConstraintsPtrInput
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapInput
	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location pulumi.StringInput
	// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
	Name pulumi.StringPtrInput
	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baselineValues that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	PassthroughExtensions CertificateTemplatePassthroughExtensionsPtrInput
	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baselineValues for the same properties, the certificate issuance request will fail.
	PredefinedValues CertificateTemplatePredefinedValuesPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
}

func (CertificateTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateTemplateArgs)(nil)).Elem()
}

type CertificateTemplateInput interface {
	pulumi.Input

	ToCertificateTemplateOutput() CertificateTemplateOutput
	ToCertificateTemplateOutputWithContext(ctx context.Context) CertificateTemplateOutput
}

func (*CertificateTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateTemplate)(nil)).Elem()
}

func (i *CertificateTemplate) ToCertificateTemplateOutput() CertificateTemplateOutput {
	return i.ToCertificateTemplateOutputWithContext(context.Background())
}

func (i *CertificateTemplate) ToCertificateTemplateOutputWithContext(ctx context.Context) CertificateTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateTemplateOutput)
}

// CertificateTemplateArrayInput is an input type that accepts CertificateTemplateArray and CertificateTemplateArrayOutput values.
// You can construct a concrete instance of `CertificateTemplateArrayInput` via:
//
//          CertificateTemplateArray{ CertificateTemplateArgs{...} }
type CertificateTemplateArrayInput interface {
	pulumi.Input

	ToCertificateTemplateArrayOutput() CertificateTemplateArrayOutput
	ToCertificateTemplateArrayOutputWithContext(context.Context) CertificateTemplateArrayOutput
}

type CertificateTemplateArray []CertificateTemplateInput

func (CertificateTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateTemplate)(nil)).Elem()
}

func (i CertificateTemplateArray) ToCertificateTemplateArrayOutput() CertificateTemplateArrayOutput {
	return i.ToCertificateTemplateArrayOutputWithContext(context.Background())
}

func (i CertificateTemplateArray) ToCertificateTemplateArrayOutputWithContext(ctx context.Context) CertificateTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateTemplateArrayOutput)
}

// CertificateTemplateMapInput is an input type that accepts CertificateTemplateMap and CertificateTemplateMapOutput values.
// You can construct a concrete instance of `CertificateTemplateMapInput` via:
//
//          CertificateTemplateMap{ "key": CertificateTemplateArgs{...} }
type CertificateTemplateMapInput interface {
	pulumi.Input

	ToCertificateTemplateMapOutput() CertificateTemplateMapOutput
	ToCertificateTemplateMapOutputWithContext(context.Context) CertificateTemplateMapOutput
}

type CertificateTemplateMap map[string]CertificateTemplateInput

func (CertificateTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateTemplate)(nil)).Elem()
}

func (i CertificateTemplateMap) ToCertificateTemplateMapOutput() CertificateTemplateMapOutput {
	return i.ToCertificateTemplateMapOutputWithContext(context.Background())
}

func (i CertificateTemplateMap) ToCertificateTemplateMapOutputWithContext(ctx context.Context) CertificateTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateTemplateMapOutput)
}

type CertificateTemplateOutput struct{ *pulumi.OutputState }

func (CertificateTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateTemplate)(nil)).Elem()
}

func (o CertificateTemplateOutput) ToCertificateTemplateOutput() CertificateTemplateOutput {
	return o
}

func (o CertificateTemplateOutput) ToCertificateTemplateOutputWithContext(ctx context.Context) CertificateTemplateOutput {
	return o
}

// Output only. The time at which this CertificateTemplate was created.
func (o CertificateTemplateOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateTemplate) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
func (o CertificateTemplateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateTemplate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
func (o CertificateTemplateOutput) IdentityConstraints() CertificateTemplateIdentityConstraintsPtrOutput {
	return o.ApplyT(func(v *CertificateTemplate) CertificateTemplateIdentityConstraintsPtrOutput {
		return v.IdentityConstraints
	}).(CertificateTemplateIdentityConstraintsPtrOutput)
}

// Optional. Labels with user-defined metadata.
func (o CertificateTemplateOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateTemplate) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
func (o CertificateTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
func (o CertificateTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baselineValues that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
func (o CertificateTemplateOutput) PassthroughExtensions() CertificateTemplatePassthroughExtensionsPtrOutput {
	return o.ApplyT(func(v *CertificateTemplate) CertificateTemplatePassthroughExtensionsPtrOutput {
		return v.PassthroughExtensions
	}).(CertificateTemplatePassthroughExtensionsPtrOutput)
}

// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baselineValues for the same properties, the certificate issuance request will fail.
func (o CertificateTemplateOutput) PredefinedValues() CertificateTemplatePredefinedValuesPtrOutput {
	return o.ApplyT(func(v *CertificateTemplate) CertificateTemplatePredefinedValuesPtrOutput { return v.PredefinedValues }).(CertificateTemplatePredefinedValuesPtrOutput)
}

// The project for the resource
func (o CertificateTemplateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateTemplate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. The time at which this CertificateTemplate was updated.
func (o CertificateTemplateOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateTemplate) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type CertificateTemplateArrayOutput struct{ *pulumi.OutputState }

func (CertificateTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateTemplate)(nil)).Elem()
}

func (o CertificateTemplateArrayOutput) ToCertificateTemplateArrayOutput() CertificateTemplateArrayOutput {
	return o
}

func (o CertificateTemplateArrayOutput) ToCertificateTemplateArrayOutputWithContext(ctx context.Context) CertificateTemplateArrayOutput {
	return o
}

func (o CertificateTemplateArrayOutput) Index(i pulumi.IntInput) CertificateTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertificateTemplate {
		return vs[0].([]*CertificateTemplate)[vs[1].(int)]
	}).(CertificateTemplateOutput)
}

type CertificateTemplateMapOutput struct{ *pulumi.OutputState }

func (CertificateTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateTemplate)(nil)).Elem()
}

func (o CertificateTemplateMapOutput) ToCertificateTemplateMapOutput() CertificateTemplateMapOutput {
	return o
}

func (o CertificateTemplateMapOutput) ToCertificateTemplateMapOutputWithContext(ctx context.Context) CertificateTemplateMapOutput {
	return o
}

func (o CertificateTemplateMapOutput) MapIndex(k pulumi.StringInput) CertificateTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertificateTemplate {
		return vs[0].(map[string]*CertificateTemplate)[vs[1].(string)]
	}).(CertificateTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateTemplateInput)(nil)).Elem(), &CertificateTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateTemplateArrayInput)(nil)).Elem(), CertificateTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateTemplateMapInput)(nil)).Elem(), CertificateTemplateMap{})
	pulumi.RegisterOutputType(CertificateTemplateOutput{})
	pulumi.RegisterOutputType(CertificateTemplateArrayOutput{})
	pulumi.RegisterOutputType(CertificateTemplateMapOutput{})
}

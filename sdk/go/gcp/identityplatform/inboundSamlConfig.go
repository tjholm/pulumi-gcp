// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identityplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Inbound SAML configuration for a Identity Toolkit project.
//
// You must enable the
// [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
// the marketplace prior to using this resource.
//
// ## Example Usage
//
// ## Import
//
// InboundSamlConfig can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:identityplatform/inboundSamlConfig:InboundSamlConfig default projects/{{project}}/inboundSamlConfigs/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:identityplatform/inboundSamlConfig:InboundSamlConfig default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:identityplatform/inboundSamlConfig:InboundSamlConfig default {{name}}
// ```
type InboundSamlConfig struct {
	pulumi.CustomResourceState

	// Human friendly display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdpConfig InboundSamlConfigIdpConfigOutput `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig InboundSamlConfigSpConfigOutput `pulumi:"spConfig"`
}

// NewInboundSamlConfig registers a new resource with the given unique name, arguments, and options.
func NewInboundSamlConfig(ctx *pulumi.Context,
	name string, args *InboundSamlConfigArgs, opts ...pulumi.ResourceOption) (*InboundSamlConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.IdpConfig == nil {
		return nil, errors.New("invalid value for required argument 'IdpConfig'")
	}
	if args.SpConfig == nil {
		return nil, errors.New("invalid value for required argument 'SpConfig'")
	}
	var resource InboundSamlConfig
	err := ctx.RegisterResource("gcp:identityplatform/inboundSamlConfig:InboundSamlConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInboundSamlConfig gets an existing InboundSamlConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInboundSamlConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InboundSamlConfigState, opts ...pulumi.ResourceOption) (*InboundSamlConfig, error) {
	var resource InboundSamlConfig
	err := ctx.ReadResource("gcp:identityplatform/inboundSamlConfig:InboundSamlConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InboundSamlConfig resources.
type inboundSamlConfigState struct {
	// Human friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled *bool `pulumi:"enabled"`
	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdpConfig *InboundSamlConfigIdpConfig `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig *InboundSamlConfigSpConfig `pulumi:"spConfig"`
}

type InboundSamlConfigState struct {
	// Human friendly display name.
	DisplayName pulumi.StringPtrInput
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrInput
	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdpConfig InboundSamlConfigIdpConfigPtrInput
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig InboundSamlConfigSpConfigPtrInput
}

func (InboundSamlConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundSamlConfigState)(nil)).Elem()
}

type inboundSamlConfigArgs struct {
	// Human friendly display name.
	DisplayName string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled *bool `pulumi:"enabled"`
	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdpConfig InboundSamlConfigIdpConfig `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig InboundSamlConfigSpConfig `pulumi:"spConfig"`
}

// The set of arguments for constructing a InboundSamlConfig resource.
type InboundSamlConfigArgs struct {
	// Human friendly display name.
	DisplayName pulumi.StringInput
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrInput
	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdpConfig InboundSamlConfigIdpConfigInput
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig InboundSamlConfigSpConfigInput
}

func (InboundSamlConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundSamlConfigArgs)(nil)).Elem()
}

type InboundSamlConfigInput interface {
	pulumi.Input

	ToInboundSamlConfigOutput() InboundSamlConfigOutput
	ToInboundSamlConfigOutputWithContext(ctx context.Context) InboundSamlConfigOutput
}

func (*InboundSamlConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundSamlConfig)(nil))
}

func (i *InboundSamlConfig) ToInboundSamlConfigOutput() InboundSamlConfigOutput {
	return i.ToInboundSamlConfigOutputWithContext(context.Background())
}

func (i *InboundSamlConfig) ToInboundSamlConfigOutputWithContext(ctx context.Context) InboundSamlConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundSamlConfigOutput)
}

func (i *InboundSamlConfig) ToInboundSamlConfigPtrOutput() InboundSamlConfigPtrOutput {
	return i.ToInboundSamlConfigPtrOutputWithContext(context.Background())
}

func (i *InboundSamlConfig) ToInboundSamlConfigPtrOutputWithContext(ctx context.Context) InboundSamlConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundSamlConfigPtrOutput)
}

type InboundSamlConfigPtrInput interface {
	pulumi.Input

	ToInboundSamlConfigPtrOutput() InboundSamlConfigPtrOutput
	ToInboundSamlConfigPtrOutputWithContext(ctx context.Context) InboundSamlConfigPtrOutput
}

type inboundSamlConfigPtrType InboundSamlConfigArgs

func (*inboundSamlConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundSamlConfig)(nil))
}

func (i *inboundSamlConfigPtrType) ToInboundSamlConfigPtrOutput() InboundSamlConfigPtrOutput {
	return i.ToInboundSamlConfigPtrOutputWithContext(context.Background())
}

func (i *inboundSamlConfigPtrType) ToInboundSamlConfigPtrOutputWithContext(ctx context.Context) InboundSamlConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundSamlConfigPtrOutput)
}

// InboundSamlConfigArrayInput is an input type that accepts InboundSamlConfigArray and InboundSamlConfigArrayOutput values.
// You can construct a concrete instance of `InboundSamlConfigArrayInput` via:
//
//          InboundSamlConfigArray{ InboundSamlConfigArgs{...} }
type InboundSamlConfigArrayInput interface {
	pulumi.Input

	ToInboundSamlConfigArrayOutput() InboundSamlConfigArrayOutput
	ToInboundSamlConfigArrayOutputWithContext(context.Context) InboundSamlConfigArrayOutput
}

type InboundSamlConfigArray []InboundSamlConfigInput

func (InboundSamlConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*InboundSamlConfig)(nil))
}

func (i InboundSamlConfigArray) ToInboundSamlConfigArrayOutput() InboundSamlConfigArrayOutput {
	return i.ToInboundSamlConfigArrayOutputWithContext(context.Background())
}

func (i InboundSamlConfigArray) ToInboundSamlConfigArrayOutputWithContext(ctx context.Context) InboundSamlConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundSamlConfigArrayOutput)
}

// InboundSamlConfigMapInput is an input type that accepts InboundSamlConfigMap and InboundSamlConfigMapOutput values.
// You can construct a concrete instance of `InboundSamlConfigMapInput` via:
//
//          InboundSamlConfigMap{ "key": InboundSamlConfigArgs{...} }
type InboundSamlConfigMapInput interface {
	pulumi.Input

	ToInboundSamlConfigMapOutput() InboundSamlConfigMapOutput
	ToInboundSamlConfigMapOutputWithContext(context.Context) InboundSamlConfigMapOutput
}

type InboundSamlConfigMap map[string]InboundSamlConfigInput

func (InboundSamlConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*InboundSamlConfig)(nil))
}

func (i InboundSamlConfigMap) ToInboundSamlConfigMapOutput() InboundSamlConfigMapOutput {
	return i.ToInboundSamlConfigMapOutputWithContext(context.Background())
}

func (i InboundSamlConfigMap) ToInboundSamlConfigMapOutputWithContext(ctx context.Context) InboundSamlConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundSamlConfigMapOutput)
}

type InboundSamlConfigOutput struct {
	*pulumi.OutputState
}

func (InboundSamlConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundSamlConfig)(nil))
}

func (o InboundSamlConfigOutput) ToInboundSamlConfigOutput() InboundSamlConfigOutput {
	return o
}

func (o InboundSamlConfigOutput) ToInboundSamlConfigOutputWithContext(ctx context.Context) InboundSamlConfigOutput {
	return o
}

func (o InboundSamlConfigOutput) ToInboundSamlConfigPtrOutput() InboundSamlConfigPtrOutput {
	return o.ToInboundSamlConfigPtrOutputWithContext(context.Background())
}

func (o InboundSamlConfigOutput) ToInboundSamlConfigPtrOutputWithContext(ctx context.Context) InboundSamlConfigPtrOutput {
	return o.ApplyT(func(v InboundSamlConfig) *InboundSamlConfig {
		return &v
	}).(InboundSamlConfigPtrOutput)
}

type InboundSamlConfigPtrOutput struct {
	*pulumi.OutputState
}

func (InboundSamlConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundSamlConfig)(nil))
}

func (o InboundSamlConfigPtrOutput) ToInboundSamlConfigPtrOutput() InboundSamlConfigPtrOutput {
	return o
}

func (o InboundSamlConfigPtrOutput) ToInboundSamlConfigPtrOutputWithContext(ctx context.Context) InboundSamlConfigPtrOutput {
	return o
}

type InboundSamlConfigArrayOutput struct{ *pulumi.OutputState }

func (InboundSamlConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundSamlConfig)(nil))
}

func (o InboundSamlConfigArrayOutput) ToInboundSamlConfigArrayOutput() InboundSamlConfigArrayOutput {
	return o
}

func (o InboundSamlConfigArrayOutput) ToInboundSamlConfigArrayOutputWithContext(ctx context.Context) InboundSamlConfigArrayOutput {
	return o
}

func (o InboundSamlConfigArrayOutput) Index(i pulumi.IntInput) InboundSamlConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundSamlConfig {
		return vs[0].([]InboundSamlConfig)[vs[1].(int)]
	}).(InboundSamlConfigOutput)
}

type InboundSamlConfigMapOutput struct{ *pulumi.OutputState }

func (InboundSamlConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InboundSamlConfig)(nil))
}

func (o InboundSamlConfigMapOutput) ToInboundSamlConfigMapOutput() InboundSamlConfigMapOutput {
	return o
}

func (o InboundSamlConfigMapOutput) ToInboundSamlConfigMapOutputWithContext(ctx context.Context) InboundSamlConfigMapOutput {
	return o
}

func (o InboundSamlConfigMapOutput) MapIndex(k pulumi.StringInput) InboundSamlConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InboundSamlConfig {
		return vs[0].(map[string]InboundSamlConfig)[vs[1].(string)]
	}).(InboundSamlConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(InboundSamlConfigOutput{})
	pulumi.RegisterOutputType(InboundSamlConfigPtrOutput{})
	pulumi.RegisterOutputType(InboundSamlConfigArrayOutput{})
	pulumi.RegisterOutputType(InboundSamlConfigMapOutput{})
}

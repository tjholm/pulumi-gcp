// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identityplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OIDC IdP configuration for a Identity Toolkit project.
//
// You must enable the
// [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
// the marketplace prior to using this resource.
//
// ## Example Usage
// ### Identity Platform Oauth Idp Config Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/identityplatform"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identityplatform.NewOauthIdpConfig(ctx, "oauthIdpConfig", &identityplatform.OauthIdpConfigArgs{
//				ClientId:     pulumi.String("client-id"),
//				ClientSecret: pulumi.String("secret"),
//				DisplayName:  pulumi.String("Display Name"),
//				Enabled:      pulumi.Bool(true),
//				Issuer:       pulumi.String("issuer"),
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
// # OauthIdpConfig can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:identityplatform/oauthIdpConfig:OauthIdpConfig default projects/{{project}}/oauthIdpConfigs/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:identityplatform/oauthIdpConfig:OauthIdpConfig default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:identityplatform/oauthIdpConfig:OauthIdpConfig default {{name}}
//
// ```
type OauthIdpConfig struct {
	pulumi.CustomResourceState

	// The client id of an OAuth client.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// Human friendly display name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// For OIDC Idps, the issuer identifier.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// The name of the OauthIdpConfig. Must start with `oidc.`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewOauthIdpConfig registers a new resource with the given unique name, arguments, and options.
func NewOauthIdpConfig(ctx *pulumi.Context,
	name string, args *OauthIdpConfigArgs, opts ...pulumi.ResourceOption) (*OauthIdpConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	var resource OauthIdpConfig
	err := ctx.RegisterResource("gcp:identityplatform/oauthIdpConfig:OauthIdpConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOauthIdpConfig gets an existing OauthIdpConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOauthIdpConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OauthIdpConfigState, opts ...pulumi.ResourceOption) (*OauthIdpConfig, error) {
	var resource OauthIdpConfig
	err := ctx.ReadResource("gcp:identityplatform/oauthIdpConfig:OauthIdpConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OauthIdpConfig resources.
type oauthIdpConfigState struct {
	// The client id of an OAuth client.
	ClientId *string `pulumi:"clientId"`
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret *string `pulumi:"clientSecret"`
	// Human friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled *bool `pulumi:"enabled"`
	// For OIDC Idps, the issuer identifier.
	Issuer *string `pulumi:"issuer"`
	// The name of the OauthIdpConfig. Must start with `oidc.`.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type OauthIdpConfigState struct {
	// The client id of an OAuth client.
	ClientId pulumi.StringPtrInput
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret pulumi.StringPtrInput
	// Human friendly display name.
	DisplayName pulumi.StringPtrInput
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrInput
	// For OIDC Idps, the issuer identifier.
	Issuer pulumi.StringPtrInput
	// The name of the OauthIdpConfig. Must start with `oidc.`.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (OauthIdpConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*oauthIdpConfigState)(nil)).Elem()
}

type oauthIdpConfigArgs struct {
	// The client id of an OAuth client.
	ClientId string `pulumi:"clientId"`
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret *string `pulumi:"clientSecret"`
	// Human friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled *bool `pulumi:"enabled"`
	// For OIDC Idps, the issuer identifier.
	Issuer string `pulumi:"issuer"`
	// The name of the OauthIdpConfig. Must start with `oidc.`.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a OauthIdpConfig resource.
type OauthIdpConfigArgs struct {
	// The client id of an OAuth client.
	ClientId pulumi.StringInput
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret pulumi.StringPtrInput
	// Human friendly display name.
	DisplayName pulumi.StringPtrInput
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrInput
	// For OIDC Idps, the issuer identifier.
	Issuer pulumi.StringInput
	// The name of the OauthIdpConfig. Must start with `oidc.`.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (OauthIdpConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oauthIdpConfigArgs)(nil)).Elem()
}

type OauthIdpConfigInput interface {
	pulumi.Input

	ToOauthIdpConfigOutput() OauthIdpConfigOutput
	ToOauthIdpConfigOutputWithContext(ctx context.Context) OauthIdpConfigOutput
}

func (*OauthIdpConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**OauthIdpConfig)(nil)).Elem()
}

func (i *OauthIdpConfig) ToOauthIdpConfigOutput() OauthIdpConfigOutput {
	return i.ToOauthIdpConfigOutputWithContext(context.Background())
}

func (i *OauthIdpConfig) ToOauthIdpConfigOutputWithContext(ctx context.Context) OauthIdpConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OauthIdpConfigOutput)
}

// OauthIdpConfigArrayInput is an input type that accepts OauthIdpConfigArray and OauthIdpConfigArrayOutput values.
// You can construct a concrete instance of `OauthIdpConfigArrayInput` via:
//
//	OauthIdpConfigArray{ OauthIdpConfigArgs{...} }
type OauthIdpConfigArrayInput interface {
	pulumi.Input

	ToOauthIdpConfigArrayOutput() OauthIdpConfigArrayOutput
	ToOauthIdpConfigArrayOutputWithContext(context.Context) OauthIdpConfigArrayOutput
}

type OauthIdpConfigArray []OauthIdpConfigInput

func (OauthIdpConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OauthIdpConfig)(nil)).Elem()
}

func (i OauthIdpConfigArray) ToOauthIdpConfigArrayOutput() OauthIdpConfigArrayOutput {
	return i.ToOauthIdpConfigArrayOutputWithContext(context.Background())
}

func (i OauthIdpConfigArray) ToOauthIdpConfigArrayOutputWithContext(ctx context.Context) OauthIdpConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OauthIdpConfigArrayOutput)
}

// OauthIdpConfigMapInput is an input type that accepts OauthIdpConfigMap and OauthIdpConfigMapOutput values.
// You can construct a concrete instance of `OauthIdpConfigMapInput` via:
//
//	OauthIdpConfigMap{ "key": OauthIdpConfigArgs{...} }
type OauthIdpConfigMapInput interface {
	pulumi.Input

	ToOauthIdpConfigMapOutput() OauthIdpConfigMapOutput
	ToOauthIdpConfigMapOutputWithContext(context.Context) OauthIdpConfigMapOutput
}

type OauthIdpConfigMap map[string]OauthIdpConfigInput

func (OauthIdpConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OauthIdpConfig)(nil)).Elem()
}

func (i OauthIdpConfigMap) ToOauthIdpConfigMapOutput() OauthIdpConfigMapOutput {
	return i.ToOauthIdpConfigMapOutputWithContext(context.Background())
}

func (i OauthIdpConfigMap) ToOauthIdpConfigMapOutputWithContext(ctx context.Context) OauthIdpConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OauthIdpConfigMapOutput)
}

type OauthIdpConfigOutput struct{ *pulumi.OutputState }

func (OauthIdpConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OauthIdpConfig)(nil)).Elem()
}

func (o OauthIdpConfigOutput) ToOauthIdpConfigOutput() OauthIdpConfigOutput {
	return o
}

func (o OauthIdpConfigOutput) ToOauthIdpConfigOutputWithContext(ctx context.Context) OauthIdpConfigOutput {
	return o
}

// The client id of an OAuth client.
func (o OauthIdpConfigOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret of the OAuth client, to enable OIDC code flow.
func (o OauthIdpConfigOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// Human friendly display name.
func (o OauthIdpConfigOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// If this config allows users to sign in with the provider.
func (o OauthIdpConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// For OIDC Idps, the issuer identifier.
func (o OauthIdpConfigOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// The name of the OauthIdpConfig. Must start with `oidc.`.
func (o OauthIdpConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o OauthIdpConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type OauthIdpConfigArrayOutput struct{ *pulumi.OutputState }

func (OauthIdpConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OauthIdpConfig)(nil)).Elem()
}

func (o OauthIdpConfigArrayOutput) ToOauthIdpConfigArrayOutput() OauthIdpConfigArrayOutput {
	return o
}

func (o OauthIdpConfigArrayOutput) ToOauthIdpConfigArrayOutputWithContext(ctx context.Context) OauthIdpConfigArrayOutput {
	return o
}

func (o OauthIdpConfigArrayOutput) Index(i pulumi.IntInput) OauthIdpConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OauthIdpConfig {
		return vs[0].([]*OauthIdpConfig)[vs[1].(int)]
	}).(OauthIdpConfigOutput)
}

type OauthIdpConfigMapOutput struct{ *pulumi.OutputState }

func (OauthIdpConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OauthIdpConfig)(nil)).Elem()
}

func (o OauthIdpConfigMapOutput) ToOauthIdpConfigMapOutput() OauthIdpConfigMapOutput {
	return o
}

func (o OauthIdpConfigMapOutput) ToOauthIdpConfigMapOutputWithContext(ctx context.Context) OauthIdpConfigMapOutput {
	return o
}

func (o OauthIdpConfigMapOutput) MapIndex(k pulumi.StringInput) OauthIdpConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OauthIdpConfig {
		return vs[0].(map[string]*OauthIdpConfig)[vs[1].(string)]
	}).(OauthIdpConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OauthIdpConfigInput)(nil)).Elem(), &OauthIdpConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*OauthIdpConfigArrayInput)(nil)).Elem(), OauthIdpConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OauthIdpConfigMapInput)(nil)).Elem(), OauthIdpConfigMap{})
	pulumi.RegisterOutputType(OauthIdpConfigOutput{})
	pulumi.RegisterOutputType(OauthIdpConfigArrayOutput{})
	pulumi.RegisterOutputType(OauthIdpConfigMapOutput{})
}

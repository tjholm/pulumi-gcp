// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identityplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OIDC IdP configuration for a Identity Toolkit project within a tenant.
//
// You must enable the
// [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
// the marketplace prior to using this resource.
//
// ## Example Usage
// ### Identity Platform Tenant Oauth Idp Config Basic
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
//			tenant, err := identityplatform.NewTenant(ctx, "tenant", &identityplatform.TenantArgs{
//				DisplayName: pulumi.String("tenant"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identityplatform.NewTenantOauthIdpConfig(ctx, "tenantOauthIdpConfig", &identityplatform.TenantOauthIdpConfigArgs{
//				Tenant:       tenant.Name,
//				DisplayName:  pulumi.String("Display Name"),
//				ClientId:     pulumi.String("client-id"),
//				Issuer:       pulumi.String("issuer"),
//				Enabled:      pulumi.Bool(true),
//				ClientSecret: pulumi.String("secret"),
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
// # TenantOauthIdpConfig can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig default projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig default {{project}}/{{tenant}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig default {{tenant}}/{{name}}
//
// ```
type TenantOauthIdpConfig struct {
	pulumi.CustomResourceState

	// The client id of an OAuth client.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// Human friendly display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// For OIDC Idps, the issuer identifier.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// The name of the OauthIdpConfig. Must start with `oidc.`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The name of the tenant where this OIDC IDP configuration resource exists
	Tenant pulumi.StringOutput `pulumi:"tenant"`
}

// NewTenantOauthIdpConfig registers a new resource with the given unique name, arguments, and options.
func NewTenantOauthIdpConfig(ctx *pulumi.Context,
	name string, args *TenantOauthIdpConfigArgs, opts ...pulumi.ResourceOption) (*TenantOauthIdpConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	if args.Tenant == nil {
		return nil, errors.New("invalid value for required argument 'Tenant'")
	}
	var resource TenantOauthIdpConfig
	err := ctx.RegisterResource("gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenantOauthIdpConfig gets an existing TenantOauthIdpConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenantOauthIdpConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantOauthIdpConfigState, opts ...pulumi.ResourceOption) (*TenantOauthIdpConfig, error) {
	var resource TenantOauthIdpConfig
	err := ctx.ReadResource("gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TenantOauthIdpConfig resources.
type tenantOauthIdpConfigState struct {
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
	// The name of the tenant where this OIDC IDP configuration resource exists
	Tenant *string `pulumi:"tenant"`
}

type TenantOauthIdpConfigState struct {
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
	// The name of the tenant where this OIDC IDP configuration resource exists
	Tenant pulumi.StringPtrInput
}

func (TenantOauthIdpConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantOauthIdpConfigState)(nil)).Elem()
}

type tenantOauthIdpConfigArgs struct {
	// The client id of an OAuth client.
	ClientId string `pulumi:"clientId"`
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret *string `pulumi:"clientSecret"`
	// Human friendly display name.
	DisplayName string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled *bool `pulumi:"enabled"`
	// For OIDC Idps, the issuer identifier.
	Issuer string `pulumi:"issuer"`
	// The name of the OauthIdpConfig. Must start with `oidc.`.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The name of the tenant where this OIDC IDP configuration resource exists
	Tenant string `pulumi:"tenant"`
}

// The set of arguments for constructing a TenantOauthIdpConfig resource.
type TenantOauthIdpConfigArgs struct {
	// The client id of an OAuth client.
	ClientId pulumi.StringInput
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret pulumi.StringPtrInput
	// Human friendly display name.
	DisplayName pulumi.StringInput
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrInput
	// For OIDC Idps, the issuer identifier.
	Issuer pulumi.StringInput
	// The name of the OauthIdpConfig. Must start with `oidc.`.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The name of the tenant where this OIDC IDP configuration resource exists
	Tenant pulumi.StringInput
}

func (TenantOauthIdpConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantOauthIdpConfigArgs)(nil)).Elem()
}

type TenantOauthIdpConfigInput interface {
	pulumi.Input

	ToTenantOauthIdpConfigOutput() TenantOauthIdpConfigOutput
	ToTenantOauthIdpConfigOutputWithContext(ctx context.Context) TenantOauthIdpConfigOutput
}

func (*TenantOauthIdpConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantOauthIdpConfig)(nil)).Elem()
}

func (i *TenantOauthIdpConfig) ToTenantOauthIdpConfigOutput() TenantOauthIdpConfigOutput {
	return i.ToTenantOauthIdpConfigOutputWithContext(context.Background())
}

func (i *TenantOauthIdpConfig) ToTenantOauthIdpConfigOutputWithContext(ctx context.Context) TenantOauthIdpConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantOauthIdpConfigOutput)
}

// TenantOauthIdpConfigArrayInput is an input type that accepts TenantOauthIdpConfigArray and TenantOauthIdpConfigArrayOutput values.
// You can construct a concrete instance of `TenantOauthIdpConfigArrayInput` via:
//
//	TenantOauthIdpConfigArray{ TenantOauthIdpConfigArgs{...} }
type TenantOauthIdpConfigArrayInput interface {
	pulumi.Input

	ToTenantOauthIdpConfigArrayOutput() TenantOauthIdpConfigArrayOutput
	ToTenantOauthIdpConfigArrayOutputWithContext(context.Context) TenantOauthIdpConfigArrayOutput
}

type TenantOauthIdpConfigArray []TenantOauthIdpConfigInput

func (TenantOauthIdpConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TenantOauthIdpConfig)(nil)).Elem()
}

func (i TenantOauthIdpConfigArray) ToTenantOauthIdpConfigArrayOutput() TenantOauthIdpConfigArrayOutput {
	return i.ToTenantOauthIdpConfigArrayOutputWithContext(context.Background())
}

func (i TenantOauthIdpConfigArray) ToTenantOauthIdpConfigArrayOutputWithContext(ctx context.Context) TenantOauthIdpConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantOauthIdpConfigArrayOutput)
}

// TenantOauthIdpConfigMapInput is an input type that accepts TenantOauthIdpConfigMap and TenantOauthIdpConfigMapOutput values.
// You can construct a concrete instance of `TenantOauthIdpConfigMapInput` via:
//
//	TenantOauthIdpConfigMap{ "key": TenantOauthIdpConfigArgs{...} }
type TenantOauthIdpConfigMapInput interface {
	pulumi.Input

	ToTenantOauthIdpConfigMapOutput() TenantOauthIdpConfigMapOutput
	ToTenantOauthIdpConfigMapOutputWithContext(context.Context) TenantOauthIdpConfigMapOutput
}

type TenantOauthIdpConfigMap map[string]TenantOauthIdpConfigInput

func (TenantOauthIdpConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TenantOauthIdpConfig)(nil)).Elem()
}

func (i TenantOauthIdpConfigMap) ToTenantOauthIdpConfigMapOutput() TenantOauthIdpConfigMapOutput {
	return i.ToTenantOauthIdpConfigMapOutputWithContext(context.Background())
}

func (i TenantOauthIdpConfigMap) ToTenantOauthIdpConfigMapOutputWithContext(ctx context.Context) TenantOauthIdpConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantOauthIdpConfigMapOutput)
}

type TenantOauthIdpConfigOutput struct{ *pulumi.OutputState }

func (TenantOauthIdpConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantOauthIdpConfig)(nil)).Elem()
}

func (o TenantOauthIdpConfigOutput) ToTenantOauthIdpConfigOutput() TenantOauthIdpConfigOutput {
	return o
}

func (o TenantOauthIdpConfigOutput) ToTenantOauthIdpConfigOutputWithContext(ctx context.Context) TenantOauthIdpConfigOutput {
	return o
}

// The client id of an OAuth client.
func (o TenantOauthIdpConfigOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantOauthIdpConfig) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret of the OAuth client, to enable OIDC code flow.
func (o TenantOauthIdpConfigOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TenantOauthIdpConfig) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// Human friendly display name.
func (o TenantOauthIdpConfigOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantOauthIdpConfig) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// If this config allows users to sign in with the provider.
func (o TenantOauthIdpConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TenantOauthIdpConfig) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// For OIDC Idps, the issuer identifier.
func (o TenantOauthIdpConfigOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantOauthIdpConfig) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// The name of the OauthIdpConfig. Must start with `oidc.`.
func (o TenantOauthIdpConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantOauthIdpConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o TenantOauthIdpConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantOauthIdpConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name of the tenant where this OIDC IDP configuration resource exists
func (o TenantOauthIdpConfigOutput) Tenant() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantOauthIdpConfig) pulumi.StringOutput { return v.Tenant }).(pulumi.StringOutput)
}

type TenantOauthIdpConfigArrayOutput struct{ *pulumi.OutputState }

func (TenantOauthIdpConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TenantOauthIdpConfig)(nil)).Elem()
}

func (o TenantOauthIdpConfigArrayOutput) ToTenantOauthIdpConfigArrayOutput() TenantOauthIdpConfigArrayOutput {
	return o
}

func (o TenantOauthIdpConfigArrayOutput) ToTenantOauthIdpConfigArrayOutputWithContext(ctx context.Context) TenantOauthIdpConfigArrayOutput {
	return o
}

func (o TenantOauthIdpConfigArrayOutput) Index(i pulumi.IntInput) TenantOauthIdpConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TenantOauthIdpConfig {
		return vs[0].([]*TenantOauthIdpConfig)[vs[1].(int)]
	}).(TenantOauthIdpConfigOutput)
}

type TenantOauthIdpConfigMapOutput struct{ *pulumi.OutputState }

func (TenantOauthIdpConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TenantOauthIdpConfig)(nil)).Elem()
}

func (o TenantOauthIdpConfigMapOutput) ToTenantOauthIdpConfigMapOutput() TenantOauthIdpConfigMapOutput {
	return o
}

func (o TenantOauthIdpConfigMapOutput) ToTenantOauthIdpConfigMapOutputWithContext(ctx context.Context) TenantOauthIdpConfigMapOutput {
	return o
}

func (o TenantOauthIdpConfigMapOutput) MapIndex(k pulumi.StringInput) TenantOauthIdpConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TenantOauthIdpConfig {
		return vs[0].(map[string]*TenantOauthIdpConfig)[vs[1].(string)]
	}).(TenantOauthIdpConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TenantOauthIdpConfigInput)(nil)).Elem(), &TenantOauthIdpConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantOauthIdpConfigArrayInput)(nil)).Elem(), TenantOauthIdpConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantOauthIdpConfigMapInput)(nil)).Elem(), TenantOauthIdpConfigMap{})
	pulumi.RegisterOutputType(TenantOauthIdpConfigOutput{})
	pulumi.RegisterOutputType(TenantOauthIdpConfigArrayOutput{})
	pulumi.RegisterOutputType(TenantOauthIdpConfigMapOutput{})
}

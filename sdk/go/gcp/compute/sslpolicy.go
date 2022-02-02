// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a SSL policy. SSL policies give you the ability to control the
// features of SSL that your SSL proxy or HTTPS load balancer negotiates.
//
// To get more information about SslPolicy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslPolicies)
// * How-to Guides
//     * [Using SSL Policies](https://cloud.google.com/compute/docs/load-balancing/ssl-policies)
//
// ## Example Usage
// ### Ssl Policy Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewSSLPolicy(ctx, "prod-ssl-policy", &compute.SSLPolicyArgs{
// 			Profile: pulumi.String("MODERN"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSSLPolicy(ctx, "nonprod-ssl-policy", &compute.SSLPolicyArgs{
// 			MinTlsVersion: pulumi.String("TLS_1_2"),
// 			Profile:       pulumi.String("MODERN"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSSLPolicy(ctx, "custom-ssl-policy", &compute.SSLPolicyArgs{
// 			CustomFeatures: pulumi.StringArray{
// 				pulumi.String("TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"),
// 				pulumi.String("TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"),
// 			},
// 			MinTlsVersion: pulumi.String("TLS_1_2"),
// 			Profile:       pulumi.String("CUSTOM"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// SslPolicy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default projects/{{project}}/global/sslPolicies/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default {{name}}
// ```
type SSLPolicy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the `CUSTOM` profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures pulumi.StringArrayOutput `pulumi:"customFeatures"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The list of features enabled in the SSL policy.
	EnabledFeatures pulumi.StringArrayOutput `pulumi:"enabledFeatures"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer.
	// Default value is `TLS_1_0`.
	// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
	MinTlsVersion pulumi.StringPtrOutput `pulumi:"minTlsVersion"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// `CUSTOM` is used, the `customFeatures` attribute **must be set**.
	// Default value is `COMPATIBLE`.
	// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
	Profile pulumi.StringPtrOutput `pulumi:"profile"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewSSLPolicy registers a new resource with the given unique name, arguments, and options.
func NewSSLPolicy(ctx *pulumi.Context,
	name string, args *SSLPolicyArgs, opts ...pulumi.ResourceOption) (*SSLPolicy, error) {
	if args == nil {
		args = &SSLPolicyArgs{}
	}

	var resource SSLPolicy
	err := ctx.RegisterResource("gcp:compute/sSLPolicy:SSLPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSSLPolicy gets an existing SSLPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSSLPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SSLPolicyState, opts ...pulumi.ResourceOption) (*SSLPolicy, error) {
	var resource SSLPolicy
	err := ctx.ReadResource("gcp:compute/sSLPolicy:SSLPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SSLPolicy resources.
type sslpolicyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the `CUSTOM` profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures []string `pulumi:"customFeatures"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The list of features enabled in the SSL policy.
	EnabledFeatures []string `pulumi:"enabledFeatures"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint *string `pulumi:"fingerprint"`
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer.
	// Default value is `TLS_1_0`.
	// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// `CUSTOM` is used, the `customFeatures` attribute **must be set**.
	// Default value is `COMPATIBLE`.
	// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
	Profile *string `pulumi:"profile"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type SSLPolicyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the `CUSTOM` profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures pulumi.StringArrayInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The list of features enabled in the SSL policy.
	EnabledFeatures pulumi.StringArrayInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringPtrInput
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer.
	// Default value is `TLS_1_0`.
	// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
	MinTlsVersion pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// `CUSTOM` is used, the `customFeatures` attribute **must be set**.
	// Default value is `COMPATIBLE`.
	// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
	Profile pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (SSLPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslpolicyState)(nil)).Elem()
}

type sslpolicyArgs struct {
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the `CUSTOM` profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures []string `pulumi:"customFeatures"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer.
	// Default value is `TLS_1_0`.
	// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// `CUSTOM` is used, the `customFeatures` attribute **must be set**.
	// Default value is `COMPATIBLE`.
	// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
	Profile *string `pulumi:"profile"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a SSLPolicy resource.
type SSLPolicyArgs struct {
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the `CUSTOM` profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures pulumi.StringArrayInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer.
	// Default value is `TLS_1_0`.
	// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
	MinTlsVersion pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// `CUSTOM` is used, the `customFeatures` attribute **must be set**.
	// Default value is `COMPATIBLE`.
	// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
	Profile pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (SSLPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslpolicyArgs)(nil)).Elem()
}

type SSLPolicyInput interface {
	pulumi.Input

	ToSSLPolicyOutput() SSLPolicyOutput
	ToSSLPolicyOutputWithContext(ctx context.Context) SSLPolicyOutput
}

func (*SSLPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SSLPolicy)(nil)).Elem()
}

func (i *SSLPolicy) ToSSLPolicyOutput() SSLPolicyOutput {
	return i.ToSSLPolicyOutputWithContext(context.Background())
}

func (i *SSLPolicy) ToSSLPolicyOutputWithContext(ctx context.Context) SSLPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSLPolicyOutput)
}

// SSLPolicyArrayInput is an input type that accepts SSLPolicyArray and SSLPolicyArrayOutput values.
// You can construct a concrete instance of `SSLPolicyArrayInput` via:
//
//          SSLPolicyArray{ SSLPolicyArgs{...} }
type SSLPolicyArrayInput interface {
	pulumi.Input

	ToSSLPolicyArrayOutput() SSLPolicyArrayOutput
	ToSSLPolicyArrayOutputWithContext(context.Context) SSLPolicyArrayOutput
}

type SSLPolicyArray []SSLPolicyInput

func (SSLPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SSLPolicy)(nil)).Elem()
}

func (i SSLPolicyArray) ToSSLPolicyArrayOutput() SSLPolicyArrayOutput {
	return i.ToSSLPolicyArrayOutputWithContext(context.Background())
}

func (i SSLPolicyArray) ToSSLPolicyArrayOutputWithContext(ctx context.Context) SSLPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSLPolicyArrayOutput)
}

// SSLPolicyMapInput is an input type that accepts SSLPolicyMap and SSLPolicyMapOutput values.
// You can construct a concrete instance of `SSLPolicyMapInput` via:
//
//          SSLPolicyMap{ "key": SSLPolicyArgs{...} }
type SSLPolicyMapInput interface {
	pulumi.Input

	ToSSLPolicyMapOutput() SSLPolicyMapOutput
	ToSSLPolicyMapOutputWithContext(context.Context) SSLPolicyMapOutput
}

type SSLPolicyMap map[string]SSLPolicyInput

func (SSLPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SSLPolicy)(nil)).Elem()
}

func (i SSLPolicyMap) ToSSLPolicyMapOutput() SSLPolicyMapOutput {
	return i.ToSSLPolicyMapOutputWithContext(context.Background())
}

func (i SSLPolicyMap) ToSSLPolicyMapOutputWithContext(ctx context.Context) SSLPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSLPolicyMapOutput)
}

type SSLPolicyOutput struct{ *pulumi.OutputState }

func (SSLPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SSLPolicy)(nil)).Elem()
}

func (o SSLPolicyOutput) ToSSLPolicyOutput() SSLPolicyOutput {
	return o
}

func (o SSLPolicyOutput) ToSSLPolicyOutputWithContext(ctx context.Context) SSLPolicyOutput {
	return o
}

type SSLPolicyArrayOutput struct{ *pulumi.OutputState }

func (SSLPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SSLPolicy)(nil)).Elem()
}

func (o SSLPolicyArrayOutput) ToSSLPolicyArrayOutput() SSLPolicyArrayOutput {
	return o
}

func (o SSLPolicyArrayOutput) ToSSLPolicyArrayOutputWithContext(ctx context.Context) SSLPolicyArrayOutput {
	return o
}

func (o SSLPolicyArrayOutput) Index(i pulumi.IntInput) SSLPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SSLPolicy {
		return vs[0].([]*SSLPolicy)[vs[1].(int)]
	}).(SSLPolicyOutput)
}

type SSLPolicyMapOutput struct{ *pulumi.OutputState }

func (SSLPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SSLPolicy)(nil)).Elem()
}

func (o SSLPolicyMapOutput) ToSSLPolicyMapOutput() SSLPolicyMapOutput {
	return o
}

func (o SSLPolicyMapOutput) ToSSLPolicyMapOutputWithContext(ctx context.Context) SSLPolicyMapOutput {
	return o
}

func (o SSLPolicyMapOutput) MapIndex(k pulumi.StringInput) SSLPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SSLPolicy {
		return vs[0].(map[string]*SSLPolicy)[vs[1].(string)]
	}).(SSLPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SSLPolicyInput)(nil)).Elem(), &SSLPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SSLPolicyArrayInput)(nil)).Elem(), SSLPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SSLPolicyMapInput)(nil)).Elem(), SSLPolicyMap{})
	pulumi.RegisterOutputType(SSLPolicyOutput{})
	pulumi.RegisterOutputType(SSLPolicyArrayOutput{})
	pulumi.RegisterOutputType(SSLPolicyMapOutput{})
}

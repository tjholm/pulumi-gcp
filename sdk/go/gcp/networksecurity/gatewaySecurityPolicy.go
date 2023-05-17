// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networksecurity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Network Security Gateway Security Policy Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/networksecurity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networksecurity.NewGatewaySecurityPolicy(ctx, "default", &networksecurity.GatewaySecurityPolicyArgs{
//				Location:    pulumi.String("us-central1"),
//				Description: pulumi.String("my description"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Network Security Gateway Security Policy Tls Inspection Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/networksecurity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultCaPool, err := certificateauthority.NewCaPool(ctx, "defaultCaPool", &certificateauthority.CaPoolArgs{
//				Location: pulumi.String("us-central1"),
//				Tier:     pulumi.String("DEVOPS"),
//				PublishingOptions: &certificateauthority.CaPoolPublishingOptionsArgs{
//					PublishCaCert: pulumi.Bool(false),
//					PublishCrl:    pulumi.Bool(false),
//				},
//				IssuancePolicy: &certificateauthority.CaPoolIssuancePolicyArgs{
//					MaximumLifetime: pulumi.String("1209600s"),
//					BaselineValues: &certificateauthority.CaPoolIssuancePolicyBaselineValuesArgs{
//						CaOptions: &certificateauthority.CaPoolIssuancePolicyBaselineValuesCaOptionsArgs{
//							IsCa: pulumi.Bool(false),
//						},
//						KeyUsage: &certificateauthority.CaPoolIssuancePolicyBaselineValuesKeyUsageArgs{
//							BaseKeyUsage: nil,
//							ExtendedKeyUsage: &certificateauthority.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageArgs{
//								ServerAuth: pulumi.Bool(true),
//							},
//						},
//					},
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			defaultAuthority, err := certificateauthority.NewAuthority(ctx, "defaultAuthority", &certificateauthority.AuthorityArgs{
//				Pool:                               defaultCaPool.Name,
//				CertificateAuthorityId:             pulumi.String("my-basic-certificate-authority"),
//				Location:                           pulumi.String("us-central1"),
//				Lifetime:                           pulumi.String("86400s"),
//				Type:                               pulumi.String("SELF_SIGNED"),
//				DeletionProtection:                 pulumi.Bool(false),
//				SkipGracePeriod:                    pulumi.Bool(true),
//				IgnoreActiveCertificatesOnDeletion: pulumi.Bool(true),
//				Config: &certificateauthority.AuthorityConfigArgs{
//					SubjectConfig: &certificateauthority.AuthorityConfigSubjectConfigArgs{
//						Subject: &certificateauthority.AuthorityConfigSubjectConfigSubjectArgs{
//							Organization: pulumi.String("Test LLC"),
//							CommonName:   pulumi.String("my-ca"),
//						},
//					},
//					X509Config: &certificateauthority.AuthorityConfigX509ConfigArgs{
//						CaOptions: &certificateauthority.AuthorityConfigX509ConfigCaOptionsArgs{
//							IsCa: pulumi.Bool(true),
//						},
//						KeyUsage: &certificateauthority.AuthorityConfigX509ConfigKeyUsageArgs{
//							BaseKeyUsage: &certificateauthority.AuthorityConfigX509ConfigKeyUsageBaseKeyUsageArgs{
//								CertSign: pulumi.Bool(true),
//								CrlSign:  pulumi.Bool(true),
//							},
//							ExtendedKeyUsage: &certificateauthority.AuthorityConfigX509ConfigKeyUsageExtendedKeyUsageArgs{
//								ServerAuth: pulumi.Bool(false),
//							},
//						},
//					},
//				},
//				KeySpec: &certificateauthority.AuthorityKeySpecArgs{
//					Algorithm: pulumi.String("RSA_PKCS1_4096_SHA256"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			defaultTlsInspectionPolicy, err := networksecurity.NewTlsInspectionPolicy(ctx, "defaultTlsInspectionPolicy", &networksecurity.TlsInspectionPolicyArgs{
//				Location: pulumi.String("us-central1"),
//				CaPool:   defaultCaPool.ID(),
//			}, pulumi.Provider(google_beta), pulumi.DependsOn([]pulumi.Resource{
//				defaultCaPool,
//				defaultAuthority,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = networksecurity.NewGatewaySecurityPolicy(ctx, "defaultGatewaySecurityPolicy", &networksecurity.GatewaySecurityPolicyArgs{
//				Location:            pulumi.String("us-central1"),
//				Description:         pulumi.String("my description"),
//				TlsInspectionPolicy: defaultTlsInspectionPolicy.ID(),
//			}, pulumi.Provider(google_beta), pulumi.DependsOn([]pulumi.Resource{
//				defaultTlsInspectionPolicy,
//			}))
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
// # GatewaySecurityPolicy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy default projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy default {{location}}/{{name}}
//
// ```
type GatewaySecurityPolicy struct {
	pulumi.CustomResourceState

	// The timestamp when the resource was created.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The location of the gateway security policy.
	// The default value is `global`.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
	// gatewaySecurityPolicy should match the pattern:(^a-z?$).
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Server-defined URL of this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
	TlsInspectionPolicy pulumi.StringPtrOutput `pulumi:"tlsInspectionPolicy"`
	// The timestamp when the resource was updated.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGatewaySecurityPolicy registers a new resource with the given unique name, arguments, and options.
func NewGatewaySecurityPolicy(ctx *pulumi.Context,
	name string, args *GatewaySecurityPolicyArgs, opts ...pulumi.ResourceOption) (*GatewaySecurityPolicy, error) {
	if args == nil {
		args = &GatewaySecurityPolicyArgs{}
	}

	var resource GatewaySecurityPolicy
	err := ctx.RegisterResource("gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewaySecurityPolicy gets an existing GatewaySecurityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewaySecurityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewaySecurityPolicyState, opts ...pulumi.ResourceOption) (*GatewaySecurityPolicy, error) {
	var resource GatewaySecurityPolicy
	err := ctx.ReadResource("gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewaySecurityPolicy resources.
type gatewaySecurityPolicyState struct {
	// The timestamp when the resource was created.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
	CreateTime *string `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// The location of the gateway security policy.
	// The default value is `global`.
	Location *string `pulumi:"location"`
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
	// gatewaySecurityPolicy should match the pattern:(^a-z?$).
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Server-defined URL of this resource.
	SelfLink *string `pulumi:"selfLink"`
	// Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
	TlsInspectionPolicy *string `pulumi:"tlsInspectionPolicy"`
	// The timestamp when the resource was updated.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type GatewaySecurityPolicyState struct {
	// The timestamp when the resource was created.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
	CreateTime pulumi.StringPtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// The location of the gateway security policy.
	// The default value is `global`.
	Location pulumi.StringPtrInput
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
	// gatewaySecurityPolicy should match the pattern:(^a-z?$).
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Server-defined URL of this resource.
	SelfLink pulumi.StringPtrInput
	// Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
	TlsInspectionPolicy pulumi.StringPtrInput
	// The timestamp when the resource was updated.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (GatewaySecurityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewaySecurityPolicyState)(nil)).Elem()
}

type gatewaySecurityPolicyArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// The location of the gateway security policy.
	// The default value is `global`.
	Location *string `pulumi:"location"`
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
	// gatewaySecurityPolicy should match the pattern:(^a-z?$).
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
	TlsInspectionPolicy *string `pulumi:"tlsInspectionPolicy"`
}

// The set of arguments for constructing a GatewaySecurityPolicy resource.
type GatewaySecurityPolicyArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// The location of the gateway security policy.
	// The default value is `global`.
	Location pulumi.StringPtrInput
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
	// gatewaySecurityPolicy should match the pattern:(^a-z?$).
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
	TlsInspectionPolicy pulumi.StringPtrInput
}

func (GatewaySecurityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewaySecurityPolicyArgs)(nil)).Elem()
}

type GatewaySecurityPolicyInput interface {
	pulumi.Input

	ToGatewaySecurityPolicyOutput() GatewaySecurityPolicyOutput
	ToGatewaySecurityPolicyOutputWithContext(ctx context.Context) GatewaySecurityPolicyOutput
}

func (*GatewaySecurityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewaySecurityPolicy)(nil)).Elem()
}

func (i *GatewaySecurityPolicy) ToGatewaySecurityPolicyOutput() GatewaySecurityPolicyOutput {
	return i.ToGatewaySecurityPolicyOutputWithContext(context.Background())
}

func (i *GatewaySecurityPolicy) ToGatewaySecurityPolicyOutputWithContext(ctx context.Context) GatewaySecurityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewaySecurityPolicyOutput)
}

// GatewaySecurityPolicyArrayInput is an input type that accepts GatewaySecurityPolicyArray and GatewaySecurityPolicyArrayOutput values.
// You can construct a concrete instance of `GatewaySecurityPolicyArrayInput` via:
//
//	GatewaySecurityPolicyArray{ GatewaySecurityPolicyArgs{...} }
type GatewaySecurityPolicyArrayInput interface {
	pulumi.Input

	ToGatewaySecurityPolicyArrayOutput() GatewaySecurityPolicyArrayOutput
	ToGatewaySecurityPolicyArrayOutputWithContext(context.Context) GatewaySecurityPolicyArrayOutput
}

type GatewaySecurityPolicyArray []GatewaySecurityPolicyInput

func (GatewaySecurityPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewaySecurityPolicy)(nil)).Elem()
}

func (i GatewaySecurityPolicyArray) ToGatewaySecurityPolicyArrayOutput() GatewaySecurityPolicyArrayOutput {
	return i.ToGatewaySecurityPolicyArrayOutputWithContext(context.Background())
}

func (i GatewaySecurityPolicyArray) ToGatewaySecurityPolicyArrayOutputWithContext(ctx context.Context) GatewaySecurityPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewaySecurityPolicyArrayOutput)
}

// GatewaySecurityPolicyMapInput is an input type that accepts GatewaySecurityPolicyMap and GatewaySecurityPolicyMapOutput values.
// You can construct a concrete instance of `GatewaySecurityPolicyMapInput` via:
//
//	GatewaySecurityPolicyMap{ "key": GatewaySecurityPolicyArgs{...} }
type GatewaySecurityPolicyMapInput interface {
	pulumi.Input

	ToGatewaySecurityPolicyMapOutput() GatewaySecurityPolicyMapOutput
	ToGatewaySecurityPolicyMapOutputWithContext(context.Context) GatewaySecurityPolicyMapOutput
}

type GatewaySecurityPolicyMap map[string]GatewaySecurityPolicyInput

func (GatewaySecurityPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewaySecurityPolicy)(nil)).Elem()
}

func (i GatewaySecurityPolicyMap) ToGatewaySecurityPolicyMapOutput() GatewaySecurityPolicyMapOutput {
	return i.ToGatewaySecurityPolicyMapOutputWithContext(context.Background())
}

func (i GatewaySecurityPolicyMap) ToGatewaySecurityPolicyMapOutputWithContext(ctx context.Context) GatewaySecurityPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewaySecurityPolicyMapOutput)
}

type GatewaySecurityPolicyOutput struct{ *pulumi.OutputState }

func (GatewaySecurityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewaySecurityPolicy)(nil)).Elem()
}

func (o GatewaySecurityPolicyOutput) ToGatewaySecurityPolicyOutput() GatewaySecurityPolicyOutput {
	return o
}

func (o GatewaySecurityPolicyOutput) ToGatewaySecurityPolicyOutputWithContext(ctx context.Context) GatewaySecurityPolicyOutput {
	return o
}

// The timestamp when the resource was created.
// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
func (o GatewaySecurityPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewaySecurityPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A free-text description of the resource. Max length 1024 characters.
func (o GatewaySecurityPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewaySecurityPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The location of the gateway security policy.
// The default value is `global`.
func (o GatewaySecurityPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewaySecurityPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
// gatewaySecurityPolicy should match the pattern:(^a-z?$).
func (o GatewaySecurityPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewaySecurityPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o GatewaySecurityPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewaySecurityPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Server-defined URL of this resource.
func (o GatewaySecurityPolicyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewaySecurityPolicy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
func (o GatewaySecurityPolicyOutput) TlsInspectionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewaySecurityPolicy) pulumi.StringPtrOutput { return v.TlsInspectionPolicy }).(pulumi.StringPtrOutput)
}

// The timestamp when the resource was updated.
// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o GatewaySecurityPolicyOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewaySecurityPolicy) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type GatewaySecurityPolicyArrayOutput struct{ *pulumi.OutputState }

func (GatewaySecurityPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewaySecurityPolicy)(nil)).Elem()
}

func (o GatewaySecurityPolicyArrayOutput) ToGatewaySecurityPolicyArrayOutput() GatewaySecurityPolicyArrayOutput {
	return o
}

func (o GatewaySecurityPolicyArrayOutput) ToGatewaySecurityPolicyArrayOutputWithContext(ctx context.Context) GatewaySecurityPolicyArrayOutput {
	return o
}

func (o GatewaySecurityPolicyArrayOutput) Index(i pulumi.IntInput) GatewaySecurityPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewaySecurityPolicy {
		return vs[0].([]*GatewaySecurityPolicy)[vs[1].(int)]
	}).(GatewaySecurityPolicyOutput)
}

type GatewaySecurityPolicyMapOutput struct{ *pulumi.OutputState }

func (GatewaySecurityPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewaySecurityPolicy)(nil)).Elem()
}

func (o GatewaySecurityPolicyMapOutput) ToGatewaySecurityPolicyMapOutput() GatewaySecurityPolicyMapOutput {
	return o
}

func (o GatewaySecurityPolicyMapOutput) ToGatewaySecurityPolicyMapOutputWithContext(ctx context.Context) GatewaySecurityPolicyMapOutput {
	return o
}

func (o GatewaySecurityPolicyMapOutput) MapIndex(k pulumi.StringInput) GatewaySecurityPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewaySecurityPolicy {
		return vs[0].(map[string]*GatewaySecurityPolicy)[vs[1].(string)]
	}).(GatewaySecurityPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewaySecurityPolicyInput)(nil)).Elem(), &GatewaySecurityPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewaySecurityPolicyArrayInput)(nil)).Elem(), GatewaySecurityPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewaySecurityPolicyMapInput)(nil)).Elem(), GatewaySecurityPolicyMap{})
	pulumi.RegisterOutputType(GatewaySecurityPolicyOutput{})
	pulumi.RegisterOutputType(GatewaySecurityPolicyArrayOutput{})
	pulumi.RegisterOutputType(GatewaySecurityPolicyMapOutput{})
}

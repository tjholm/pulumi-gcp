// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networksecurity

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Network Security Tls Inspection Policy Basic
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
//			_, err = networksecurity.NewTlsInspectionPolicy(ctx, "defaultTlsInspectionPolicy", &networksecurity.TlsInspectionPolicyArgs{
//				Location:           pulumi.String("us-central1"),
//				CaPool:             defaultCaPool.ID(),
//				ExcludePublicCaSet: pulumi.Bool(false),
//			}, pulumi.Provider(google_beta), pulumi.DependsOn([]pulumi.Resource{
//				defaultCaPool,
//				defaultAuthority,
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
// # TlsInspectionPolicy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:networksecurity/tlsInspectionPolicy:TlsInspectionPolicy default projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networksecurity/tlsInspectionPolicy:TlsInspectionPolicy default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networksecurity/tlsInspectionPolicy:TlsInspectionPolicy default {{location}}/{{name}}
//
// ```
type TlsInspectionPolicy struct {
	pulumi.CustomResourceState

	// A CA pool resource used to issue interception certificates.
	CaPool pulumi.StringOutput `pulumi:"caPool"`
	// The timestamp when the resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Free-text description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
	ExcludePublicCaSet pulumi.BoolPtrOutput `pulumi:"excludePublicCaSet"`
	// The location of the tls inspection policy.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Short name of the TlsInspectionPolicy resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The timestamp when the resource was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewTlsInspectionPolicy registers a new resource with the given unique name, arguments, and options.
func NewTlsInspectionPolicy(ctx *pulumi.Context,
	name string, args *TlsInspectionPolicyArgs, opts ...pulumi.ResourceOption) (*TlsInspectionPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPool == nil {
		return nil, errors.New("invalid value for required argument 'CaPool'")
	}
	var resource TlsInspectionPolicy
	err := ctx.RegisterResource("gcp:networksecurity/tlsInspectionPolicy:TlsInspectionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTlsInspectionPolicy gets an existing TlsInspectionPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTlsInspectionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TlsInspectionPolicyState, opts ...pulumi.ResourceOption) (*TlsInspectionPolicy, error) {
	var resource TlsInspectionPolicy
	err := ctx.ReadResource("gcp:networksecurity/tlsInspectionPolicy:TlsInspectionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TlsInspectionPolicy resources.
type tlsInspectionPolicyState struct {
	// A CA pool resource used to issue interception certificates.
	CaPool *string `pulumi:"caPool"`
	// The timestamp when the resource was created.
	CreateTime *string `pulumi:"createTime"`
	// Free-text description of the resource.
	Description *string `pulumi:"description"`
	// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
	ExcludePublicCaSet *bool `pulumi:"excludePublicCaSet"`
	// The location of the tls inspection policy.
	Location *string `pulumi:"location"`
	// Short name of the TlsInspectionPolicy resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The timestamp when the resource was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type TlsInspectionPolicyState struct {
	// A CA pool resource used to issue interception certificates.
	CaPool pulumi.StringPtrInput
	// The timestamp when the resource was created.
	CreateTime pulumi.StringPtrInput
	// Free-text description of the resource.
	Description pulumi.StringPtrInput
	// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
	ExcludePublicCaSet pulumi.BoolPtrInput
	// The location of the tls inspection policy.
	Location pulumi.StringPtrInput
	// Short name of the TlsInspectionPolicy resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The timestamp when the resource was updated.
	UpdateTime pulumi.StringPtrInput
}

func (TlsInspectionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tlsInspectionPolicyState)(nil)).Elem()
}

type tlsInspectionPolicyArgs struct {
	// A CA pool resource used to issue interception certificates.
	CaPool string `pulumi:"caPool"`
	// Free-text description of the resource.
	Description *string `pulumi:"description"`
	// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
	ExcludePublicCaSet *bool `pulumi:"excludePublicCaSet"`
	// The location of the tls inspection policy.
	Location *string `pulumi:"location"`
	// Short name of the TlsInspectionPolicy resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a TlsInspectionPolicy resource.
type TlsInspectionPolicyArgs struct {
	// A CA pool resource used to issue interception certificates.
	CaPool pulumi.StringInput
	// Free-text description of the resource.
	Description pulumi.StringPtrInput
	// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
	ExcludePublicCaSet pulumi.BoolPtrInput
	// The location of the tls inspection policy.
	Location pulumi.StringPtrInput
	// Short name of the TlsInspectionPolicy resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (TlsInspectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tlsInspectionPolicyArgs)(nil)).Elem()
}

type TlsInspectionPolicyInput interface {
	pulumi.Input

	ToTlsInspectionPolicyOutput() TlsInspectionPolicyOutput
	ToTlsInspectionPolicyOutputWithContext(ctx context.Context) TlsInspectionPolicyOutput
}

func (*TlsInspectionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsInspectionPolicy)(nil)).Elem()
}

func (i *TlsInspectionPolicy) ToTlsInspectionPolicyOutput() TlsInspectionPolicyOutput {
	return i.ToTlsInspectionPolicyOutputWithContext(context.Background())
}

func (i *TlsInspectionPolicy) ToTlsInspectionPolicyOutputWithContext(ctx context.Context) TlsInspectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsInspectionPolicyOutput)
}

// TlsInspectionPolicyArrayInput is an input type that accepts TlsInspectionPolicyArray and TlsInspectionPolicyArrayOutput values.
// You can construct a concrete instance of `TlsInspectionPolicyArrayInput` via:
//
//	TlsInspectionPolicyArray{ TlsInspectionPolicyArgs{...} }
type TlsInspectionPolicyArrayInput interface {
	pulumi.Input

	ToTlsInspectionPolicyArrayOutput() TlsInspectionPolicyArrayOutput
	ToTlsInspectionPolicyArrayOutputWithContext(context.Context) TlsInspectionPolicyArrayOutput
}

type TlsInspectionPolicyArray []TlsInspectionPolicyInput

func (TlsInspectionPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TlsInspectionPolicy)(nil)).Elem()
}

func (i TlsInspectionPolicyArray) ToTlsInspectionPolicyArrayOutput() TlsInspectionPolicyArrayOutput {
	return i.ToTlsInspectionPolicyArrayOutputWithContext(context.Background())
}

func (i TlsInspectionPolicyArray) ToTlsInspectionPolicyArrayOutputWithContext(ctx context.Context) TlsInspectionPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsInspectionPolicyArrayOutput)
}

// TlsInspectionPolicyMapInput is an input type that accepts TlsInspectionPolicyMap and TlsInspectionPolicyMapOutput values.
// You can construct a concrete instance of `TlsInspectionPolicyMapInput` via:
//
//	TlsInspectionPolicyMap{ "key": TlsInspectionPolicyArgs{...} }
type TlsInspectionPolicyMapInput interface {
	pulumi.Input

	ToTlsInspectionPolicyMapOutput() TlsInspectionPolicyMapOutput
	ToTlsInspectionPolicyMapOutputWithContext(context.Context) TlsInspectionPolicyMapOutput
}

type TlsInspectionPolicyMap map[string]TlsInspectionPolicyInput

func (TlsInspectionPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TlsInspectionPolicy)(nil)).Elem()
}

func (i TlsInspectionPolicyMap) ToTlsInspectionPolicyMapOutput() TlsInspectionPolicyMapOutput {
	return i.ToTlsInspectionPolicyMapOutputWithContext(context.Background())
}

func (i TlsInspectionPolicyMap) ToTlsInspectionPolicyMapOutputWithContext(ctx context.Context) TlsInspectionPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsInspectionPolicyMapOutput)
}

type TlsInspectionPolicyOutput struct{ *pulumi.OutputState }

func (TlsInspectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsInspectionPolicy)(nil)).Elem()
}

func (o TlsInspectionPolicyOutput) ToTlsInspectionPolicyOutput() TlsInspectionPolicyOutput {
	return o
}

func (o TlsInspectionPolicyOutput) ToTlsInspectionPolicyOutputWithContext(ctx context.Context) TlsInspectionPolicyOutput {
	return o
}

// A CA pool resource used to issue interception certificates.
func (o TlsInspectionPolicyOutput) CaPool() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.CaPool }).(pulumi.StringOutput)
}

// The timestamp when the resource was created.
func (o TlsInspectionPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Free-text description of the resource.
func (o TlsInspectionPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
func (o TlsInspectionPolicyOutput) ExcludePublicCaSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.BoolPtrOutput { return v.ExcludePublicCaSet }).(pulumi.BoolPtrOutput)
}

// The location of the tls inspection policy.
func (o TlsInspectionPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Short name of the TlsInspectionPolicy resource to be created.
func (o TlsInspectionPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o TlsInspectionPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o TlsInspectionPolicyOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type TlsInspectionPolicyArrayOutput struct{ *pulumi.OutputState }

func (TlsInspectionPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TlsInspectionPolicy)(nil)).Elem()
}

func (o TlsInspectionPolicyArrayOutput) ToTlsInspectionPolicyArrayOutput() TlsInspectionPolicyArrayOutput {
	return o
}

func (o TlsInspectionPolicyArrayOutput) ToTlsInspectionPolicyArrayOutputWithContext(ctx context.Context) TlsInspectionPolicyArrayOutput {
	return o
}

func (o TlsInspectionPolicyArrayOutput) Index(i pulumi.IntInput) TlsInspectionPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TlsInspectionPolicy {
		return vs[0].([]*TlsInspectionPolicy)[vs[1].(int)]
	}).(TlsInspectionPolicyOutput)
}

type TlsInspectionPolicyMapOutput struct{ *pulumi.OutputState }

func (TlsInspectionPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TlsInspectionPolicy)(nil)).Elem()
}

func (o TlsInspectionPolicyMapOutput) ToTlsInspectionPolicyMapOutput() TlsInspectionPolicyMapOutput {
	return o
}

func (o TlsInspectionPolicyMapOutput) ToTlsInspectionPolicyMapOutputWithContext(ctx context.Context) TlsInspectionPolicyMapOutput {
	return o
}

func (o TlsInspectionPolicyMapOutput) MapIndex(k pulumi.StringInput) TlsInspectionPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TlsInspectionPolicy {
		return vs[0].(map[string]*TlsInspectionPolicy)[vs[1].(string)]
	}).(TlsInspectionPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TlsInspectionPolicyInput)(nil)).Elem(), &TlsInspectionPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsInspectionPolicyArrayInput)(nil)).Elem(), TlsInspectionPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsInspectionPolicyMapInput)(nil)).Elem(), TlsInspectionPolicyMap{})
	pulumi.RegisterOutputType(TlsInspectionPolicyOutput{})
	pulumi.RegisterOutputType(TlsInspectionPolicyArrayOutput{})
	pulumi.RegisterOutputType(TlsInspectionPolicyMapOutput{})
}

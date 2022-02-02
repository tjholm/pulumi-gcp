// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package certificateauthority

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := certificateauthority.NewCaPool(ctx, "default", &certificateauthority.CaPoolArgs{
// 			Labels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			Location: pulumi.String("us-central1"),
// 			PublishingOptions: &certificateauthority.CaPoolPublishingOptionsArgs{
// 				PublishCaCert: pulumi.Bool(true),
// 				PublishCrl:    pulumi.Bool(true),
// 			},
// 			Tier: pulumi.String("ENTERPRISE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Privateca Capool All Fields
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := certificateauthority.NewCaPool(ctx, "default", &certificateauthority.CaPoolArgs{
// 			IssuancePolicy: &certificateauthority.CaPoolIssuancePolicyArgs{
// 				AllowedIssuanceModes: &certificateauthority.CaPoolIssuancePolicyAllowedIssuanceModesArgs{
// 					AllowConfigBasedIssuance: pulumi.Bool(true),
// 					AllowCsrBasedIssuance:    pulumi.Bool(true),
// 				},
// 				AllowedKeyTypes: certificateauthority.CaPoolIssuancePolicyAllowedKeyTypeArray{
// 					&certificateauthority.CaPoolIssuancePolicyAllowedKeyTypeArgs{
// 						EllipticCurve: &certificateauthority.CaPoolIssuancePolicyAllowedKeyTypeEllipticCurveArgs{
// 							SignatureAlgorithm: pulumi.String("ECDSA_P256"),
// 						},
// 					},
// 					&certificateauthority.CaPoolIssuancePolicyAllowedKeyTypeArgs{
// 						Rsa: &certificateauthority.CaPoolIssuancePolicyAllowedKeyTypeRsaArgs{
// 							MaxModulusSize: pulumi.String("10"),
// 							MinModulusSize: pulumi.String("5"),
// 						},
// 					},
// 				},
// 				BaselineValues: &certificateauthority.CaPoolIssuancePolicyBaselineValuesArgs{
// 					AdditionalExtensions: certificateauthority.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionArray{
// 						&certificateauthority.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionArgs{
// 							Critical: pulumi.Bool(true),
// 							ObjectId: &certificateauthority.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionObjectIdArgs{
// 								ObjectIdPath: []float64{
// 									1,
// 									7,
// 								},
// 							},
// 							Value: pulumi.String("asdf"),
// 						},
// 					},
// 					AiaOcspServers: pulumi.StringArray{
// 						pulumi.String("example.com"),
// 					},
// 					CaOptions: &certificateauthority.CaPoolIssuancePolicyBaselineValuesCaOptionsArgs{
// 						IsCa:                pulumi.Bool(true),
// 						MaxIssuerPathLength: pulumi.Int(10),
// 					},
// 					KeyUsage: &certificateauthority.CaPoolIssuancePolicyBaselineValuesKeyUsageArgs{
// 						BaseKeyUsage: &certificateauthority.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageArgs{
// 							CertSign:          pulumi.Bool(false),
// 							ContentCommitment: pulumi.Bool(true),
// 							CrlSign:           pulumi.Bool(true),
// 							DataEncipherment:  pulumi.Bool(true),
// 							DecipherOnly:      pulumi.Bool(true),
// 							DigitalSignature:  pulumi.Bool(true),
// 							KeyAgreement:      pulumi.Bool(true),
// 							KeyEncipherment:   pulumi.Bool(false),
// 						},
// 						ExtendedKeyUsage: &certificateauthority.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageArgs{
// 							ClientAuth:      pulumi.Bool(false),
// 							CodeSigning:     pulumi.Bool(true),
// 							EmailProtection: pulumi.Bool(true),
// 							ServerAuth:      pulumi.Bool(true),
// 							TimeStamping:    pulumi.Bool(true),
// 						},
// 					},
// 					PolicyIds: certificateauthority.CaPoolIssuancePolicyBaselineValuesPolicyIdArray{
// 						&certificateauthority.CaPoolIssuancePolicyBaselineValuesPolicyIdArgs{
// 							ObjectIdPath: []float64{
// 								1,
// 								5,
// 							},
// 						},
// 						&certificateauthority.CaPoolIssuancePolicyBaselineValuesPolicyIdArgs{
// 							ObjectIdPath: []float64{
// 								1,
// 								5,
// 								7,
// 							},
// 						},
// 					},
// 				},
// 				IdentityConstraints: &certificateauthority.CaPoolIssuancePolicyIdentityConstraintsArgs{
// 					AllowSubjectAltNamesPassthrough: pulumi.Bool(true),
// 					AllowSubjectPassthrough:         pulumi.Bool(true),
// 					CelExpression: &certificateauthority.CaPoolIssuancePolicyIdentityConstraintsCelExpressionArgs{
// 						Expression: pulumi.String("subject_alt_names.all(san, san.type == DNS || san.type == EMAIL )"),
// 						Title:      pulumi.String("My title"),
// 					},
// 				},
// 				MaximumLifetime: pulumi.String("50000s"),
// 			},
// 			Labels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			Location: pulumi.String("us-central1"),
// 			PublishingOptions: &certificateauthority.CaPoolPublishingOptionsArgs{
// 				PublishCaCert: pulumi.Bool(false),
// 				PublishCrl:    pulumi.Bool(true),
// 			},
// 			Tier: pulumi.String("ENTERPRISE"),
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
// CaPool can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPool:CaPool default projects/{{project}}/locations/{{location}}/caPools/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPool:CaPool default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPool:CaPool default {{location}}/{{name}}
// ```
type CaPool struct {
	pulumi.CustomResourceState

	// The IssuancePolicy to control how Certificates will be issued from this CaPool.
	// Structure is documented below.
	IssuancePolicy CaPoolIssuancePolicyPtrOutput `pulumi:"issuancePolicy"`
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name for this CaPool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	// Structure is documented below.
	PublishingOptions CaPoolPublishingOptionsPtrOutput `pulumi:"publishingOptions"`
	// The Tier of this CaPool.
	// Possible values are `ENTERPRISE` and `DEVOPS`.
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
	// The IssuancePolicy to control how Certificates will be issued from this CaPool.
	// Structure is documented below.
	IssuancePolicy *CaPoolIssuancePolicy `pulumi:"issuancePolicy"`
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location *string `pulumi:"location"`
	// The name for this CaPool.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	// Structure is documented below.
	PublishingOptions *CaPoolPublishingOptions `pulumi:"publishingOptions"`
	// The Tier of this CaPool.
	// Possible values are `ENTERPRISE` and `DEVOPS`.
	Tier *string `pulumi:"tier"`
}

type CaPoolState struct {
	// The IssuancePolicy to control how Certificates will be issued from this CaPool.
	// Structure is documented below.
	IssuancePolicy CaPoolIssuancePolicyPtrInput
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location pulumi.StringPtrInput
	// The name for this CaPool.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	// Structure is documented below.
	PublishingOptions CaPoolPublishingOptionsPtrInput
	// The Tier of this CaPool.
	// Possible values are `ENTERPRISE` and `DEVOPS`.
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
	Labels map[string]string `pulumi:"labels"`
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
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
	// Possible values are `ENTERPRISE` and `DEVOPS`.
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
	Labels pulumi.StringMapInput
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
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
	// Possible values are `ENTERPRISE` and `DEVOPS`.
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

// CaPoolArrayInput is an input type that accepts CaPoolArray and CaPoolArrayOutput values.
// You can construct a concrete instance of `CaPoolArrayInput` via:
//
//          CaPoolArray{ CaPoolArgs{...} }
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

// CaPoolMapInput is an input type that accepts CaPoolMap and CaPoolMapOutput values.
// You can construct a concrete instance of `CaPoolMapInput` via:
//
//          CaPoolMap{ "key": CaPoolArgs{...} }
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

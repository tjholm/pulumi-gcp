// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for KMS crypto key. Each of these resources serves a different use case:
//
// * `kms.CryptoKeyIAMPolicy`: Authoritative. Sets the IAM policy for the crypto key and replaces any existing policy already attached.
// * `kms.CryptoKeyIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the crypto key are preserved.
// * `kms.CryptoKeyIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the crypto key are preserved.
//
// > **Note:** `kms.CryptoKeyIAMPolicy` **cannot** be used in conjunction with `kms.CryptoKeyIAMBinding` and `kms.CryptoKeyIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `kms.CryptoKeyIAMBinding` resources **can be** used in conjunction with `kms.CryptoKeyIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		keyring, err := kms.NewKeyRing(ctx, "keyring", &kms.KeyRingArgs{
// 			Location: pulumi.String("global"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		key, err := kms.NewCryptoKey(ctx, "key", &kms.CryptoKeyArgs{
// 			KeyRing:        keyring.ID(),
// 			RotationPeriod: pulumi.String("100000s"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/cloudkms.cryptoKeyEncrypter",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewCryptoKeyIAMPolicy(ctx, "cryptoKey", &kms.CryptoKeyIAMPolicyArgs{
// 			CryptoKeyId: key.ID(),
// 			PolicyData:  pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 						Title:       "expires_after_2019_12_31",
// 					},
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Role: "roles/cloudkms.cryptoKeyEncrypter",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMBinding(ctx, "cryptoKey", &kms.CryptoKeyIAMBindingArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMBinding(ctx, "cryptoKey", &kms.CryptoKeyIAMBindingArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &kms.CryptoKeyIAMBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMMember(ctx, "cryptoKey", &kms.CryptoKeyIAMMemberArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Member:      pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMMember(ctx, "cryptoKey", &kms.CryptoKeyIAMMemberArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Member:      pulumi.String("user:jane@example.com"),
// 			Condition: &kms.CryptoKeyIAMMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `crypto_key_id`, role, and member identity e.g.
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember crypto_key "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; first the resource in question and then the role.
//
// These bindings can be imported using the `crypto_key_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember crypto_key "your-project-id/location-name/key-ring-name/key-name roles/editor"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `crypto_key_id`, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember crypto_key your-project-id/location-name/key-ring-name/key-name
// ```
type CryptoKeyIAMMember struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CryptoKeyIAMMemberConditionPtrOutput `pulumi:"condition"`
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId pulumi.StringOutput `pulumi:"cryptoKeyId"`
	// (Computed) The etag of the project's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewCryptoKeyIAMMember registers a new resource with the given unique name, arguments, and options.
func NewCryptoKeyIAMMember(ctx *pulumi.Context,
	name string, args *CryptoKeyIAMMemberArgs, opts ...pulumi.ResourceOption) (*CryptoKeyIAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CryptoKeyId == nil {
		return nil, errors.New("invalid value for required argument 'CryptoKeyId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource CryptoKeyIAMMember
	err := ctx.RegisterResource("gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCryptoKeyIAMMember gets an existing CryptoKeyIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCryptoKeyIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CryptoKeyIAMMemberState, opts ...pulumi.ResourceOption) (*CryptoKeyIAMMember, error) {
	var resource CryptoKeyIAMMember
	err := ctx.ReadResource("gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CryptoKeyIAMMember resources.
type cryptoKeyIAMMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *CryptoKeyIAMMemberCondition `pulumi:"condition"`
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId *string `pulumi:"cryptoKeyId"`
	// (Computed) The etag of the project's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type CryptoKeyIAMMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CryptoKeyIAMMemberConditionPtrInput
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId pulumi.StringPtrInput
	// (Computed) The etag of the project's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (CryptoKeyIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyIAMMemberState)(nil)).Elem()
}

type cryptoKeyIAMMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *CryptoKeyIAMMemberCondition `pulumi:"condition"`
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId string `pulumi:"cryptoKeyId"`
	Member      string `pulumi:"member"`
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a CryptoKeyIAMMember resource.
type CryptoKeyIAMMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CryptoKeyIAMMemberConditionPtrInput
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId pulumi.StringInput
	Member      pulumi.StringInput
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (CryptoKeyIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyIAMMemberArgs)(nil)).Elem()
}

type CryptoKeyIAMMemberInput interface {
	pulumi.Input

	ToCryptoKeyIAMMemberOutput() CryptoKeyIAMMemberOutput
	ToCryptoKeyIAMMemberOutputWithContext(ctx context.Context) CryptoKeyIAMMemberOutput
}

func (*CryptoKeyIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((*CryptoKeyIAMMember)(nil))
}

func (i *CryptoKeyIAMMember) ToCryptoKeyIAMMemberOutput() CryptoKeyIAMMemberOutput {
	return i.ToCryptoKeyIAMMemberOutputWithContext(context.Background())
}

func (i *CryptoKeyIAMMember) ToCryptoKeyIAMMemberOutputWithContext(ctx context.Context) CryptoKeyIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyIAMMemberOutput)
}

func (i *CryptoKeyIAMMember) ToCryptoKeyIAMMemberPtrOutput() CryptoKeyIAMMemberPtrOutput {
	return i.ToCryptoKeyIAMMemberPtrOutputWithContext(context.Background())
}

func (i *CryptoKeyIAMMember) ToCryptoKeyIAMMemberPtrOutputWithContext(ctx context.Context) CryptoKeyIAMMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyIAMMemberPtrOutput)
}

type CryptoKeyIAMMemberPtrInput interface {
	pulumi.Input

	ToCryptoKeyIAMMemberPtrOutput() CryptoKeyIAMMemberPtrOutput
	ToCryptoKeyIAMMemberPtrOutputWithContext(ctx context.Context) CryptoKeyIAMMemberPtrOutput
}

type cryptoKeyIAMMemberPtrType CryptoKeyIAMMemberArgs

func (*cryptoKeyIAMMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CryptoKeyIAMMember)(nil))
}

func (i *cryptoKeyIAMMemberPtrType) ToCryptoKeyIAMMemberPtrOutput() CryptoKeyIAMMemberPtrOutput {
	return i.ToCryptoKeyIAMMemberPtrOutputWithContext(context.Background())
}

func (i *cryptoKeyIAMMemberPtrType) ToCryptoKeyIAMMemberPtrOutputWithContext(ctx context.Context) CryptoKeyIAMMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyIAMMemberPtrOutput)
}

// CryptoKeyIAMMemberArrayInput is an input type that accepts CryptoKeyIAMMemberArray and CryptoKeyIAMMemberArrayOutput values.
// You can construct a concrete instance of `CryptoKeyIAMMemberArrayInput` via:
//
//          CryptoKeyIAMMemberArray{ CryptoKeyIAMMemberArgs{...} }
type CryptoKeyIAMMemberArrayInput interface {
	pulumi.Input

	ToCryptoKeyIAMMemberArrayOutput() CryptoKeyIAMMemberArrayOutput
	ToCryptoKeyIAMMemberArrayOutputWithContext(context.Context) CryptoKeyIAMMemberArrayOutput
}

type CryptoKeyIAMMemberArray []CryptoKeyIAMMemberInput

func (CryptoKeyIAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CryptoKeyIAMMember)(nil)).Elem()
}

func (i CryptoKeyIAMMemberArray) ToCryptoKeyIAMMemberArrayOutput() CryptoKeyIAMMemberArrayOutput {
	return i.ToCryptoKeyIAMMemberArrayOutputWithContext(context.Background())
}

func (i CryptoKeyIAMMemberArray) ToCryptoKeyIAMMemberArrayOutputWithContext(ctx context.Context) CryptoKeyIAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyIAMMemberArrayOutput)
}

// CryptoKeyIAMMemberMapInput is an input type that accepts CryptoKeyIAMMemberMap and CryptoKeyIAMMemberMapOutput values.
// You can construct a concrete instance of `CryptoKeyIAMMemberMapInput` via:
//
//          CryptoKeyIAMMemberMap{ "key": CryptoKeyIAMMemberArgs{...} }
type CryptoKeyIAMMemberMapInput interface {
	pulumi.Input

	ToCryptoKeyIAMMemberMapOutput() CryptoKeyIAMMemberMapOutput
	ToCryptoKeyIAMMemberMapOutputWithContext(context.Context) CryptoKeyIAMMemberMapOutput
}

type CryptoKeyIAMMemberMap map[string]CryptoKeyIAMMemberInput

func (CryptoKeyIAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CryptoKeyIAMMember)(nil)).Elem()
}

func (i CryptoKeyIAMMemberMap) ToCryptoKeyIAMMemberMapOutput() CryptoKeyIAMMemberMapOutput {
	return i.ToCryptoKeyIAMMemberMapOutputWithContext(context.Background())
}

func (i CryptoKeyIAMMemberMap) ToCryptoKeyIAMMemberMapOutputWithContext(ctx context.Context) CryptoKeyIAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyIAMMemberMapOutput)
}

type CryptoKeyIAMMemberOutput struct{ *pulumi.OutputState }

func (CryptoKeyIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CryptoKeyIAMMember)(nil))
}

func (o CryptoKeyIAMMemberOutput) ToCryptoKeyIAMMemberOutput() CryptoKeyIAMMemberOutput {
	return o
}

func (o CryptoKeyIAMMemberOutput) ToCryptoKeyIAMMemberOutputWithContext(ctx context.Context) CryptoKeyIAMMemberOutput {
	return o
}

func (o CryptoKeyIAMMemberOutput) ToCryptoKeyIAMMemberPtrOutput() CryptoKeyIAMMemberPtrOutput {
	return o.ToCryptoKeyIAMMemberPtrOutputWithContext(context.Background())
}

func (o CryptoKeyIAMMemberOutput) ToCryptoKeyIAMMemberPtrOutputWithContext(ctx context.Context) CryptoKeyIAMMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CryptoKeyIAMMember) *CryptoKeyIAMMember {
		return &v
	}).(CryptoKeyIAMMemberPtrOutput)
}

type CryptoKeyIAMMemberPtrOutput struct{ *pulumi.OutputState }

func (CryptoKeyIAMMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CryptoKeyIAMMember)(nil))
}

func (o CryptoKeyIAMMemberPtrOutput) ToCryptoKeyIAMMemberPtrOutput() CryptoKeyIAMMemberPtrOutput {
	return o
}

func (o CryptoKeyIAMMemberPtrOutput) ToCryptoKeyIAMMemberPtrOutputWithContext(ctx context.Context) CryptoKeyIAMMemberPtrOutput {
	return o
}

func (o CryptoKeyIAMMemberPtrOutput) Elem() CryptoKeyIAMMemberOutput {
	return o.ApplyT(func(v *CryptoKeyIAMMember) CryptoKeyIAMMember {
		if v != nil {
			return *v
		}
		var ret CryptoKeyIAMMember
		return ret
	}).(CryptoKeyIAMMemberOutput)
}

type CryptoKeyIAMMemberArrayOutput struct{ *pulumi.OutputState }

func (CryptoKeyIAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CryptoKeyIAMMember)(nil))
}

func (o CryptoKeyIAMMemberArrayOutput) ToCryptoKeyIAMMemberArrayOutput() CryptoKeyIAMMemberArrayOutput {
	return o
}

func (o CryptoKeyIAMMemberArrayOutput) ToCryptoKeyIAMMemberArrayOutputWithContext(ctx context.Context) CryptoKeyIAMMemberArrayOutput {
	return o
}

func (o CryptoKeyIAMMemberArrayOutput) Index(i pulumi.IntInput) CryptoKeyIAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CryptoKeyIAMMember {
		return vs[0].([]CryptoKeyIAMMember)[vs[1].(int)]
	}).(CryptoKeyIAMMemberOutput)
}

type CryptoKeyIAMMemberMapOutput struct{ *pulumi.OutputState }

func (CryptoKeyIAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CryptoKeyIAMMember)(nil))
}

func (o CryptoKeyIAMMemberMapOutput) ToCryptoKeyIAMMemberMapOutput() CryptoKeyIAMMemberMapOutput {
	return o
}

func (o CryptoKeyIAMMemberMapOutput) ToCryptoKeyIAMMemberMapOutputWithContext(ctx context.Context) CryptoKeyIAMMemberMapOutput {
	return o
}

func (o CryptoKeyIAMMemberMapOutput) MapIndex(k pulumi.StringInput) CryptoKeyIAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CryptoKeyIAMMember {
		return vs[0].(map[string]CryptoKeyIAMMember)[vs[1].(string)]
	}).(CryptoKeyIAMMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(CryptoKeyIAMMemberOutput{})
	pulumi.RegisterOutputType(CryptoKeyIAMMemberPtrOutput{})
	pulumi.RegisterOutputType(CryptoKeyIAMMemberArrayOutput{})
	pulumi.RegisterOutputType(CryptoKeyIAMMemberMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for KMS key ring. Each of these resources serves a different use case:
//
// * `kms.KeyRingIAMPolicy`: Authoritative. Sets the IAM policy for the key ring and replaces any existing policy already attached.
// * `kms.KeyRingIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the key ring are preserved.
// * `kms.KeyRingIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the key ring are preserved.
//
// > **Note:** `kms.KeyRingIAMPolicy` **cannot** be used in conjunction with `kms.KeyRingIAMBinding` and `kms.KeyRingIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `kms.KeyRingIAMBinding` resources **can be** used in conjunction with `kms.KeyRingIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_kms\_key\_ring\_iam\_policy
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
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewKeyRingIAMPolicy(ctx, "keyRing", &kms.KeyRingIAMPolicyArgs{
// 			KeyRingId:  keyring.ID(),
// 			PolicyData: pulumi.String(admin.PolicyData),
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
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewKeyRingIAMPolicy(ctx, "keyRing", &kms.KeyRingIAMPolicyArgs{
// 			KeyRingId:  keyring.ID(),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_kms\_key\_ring\_iam\_binding
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
// 		_, err := kms.NewKeyRingIAMBinding(ctx, "keyRing", &kms.KeyRingIAMBindingArgs{
// 			KeyRingId: pulumi.String("your-key-ring-id"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/editor"),
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
// 		_, err := kms.NewKeyRingIAMBinding(ctx, "keyRing", &kms.KeyRingIAMBindingArgs{
// 			Condition: &kms.KeyRingIAMBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			KeyRingId: pulumi.String("your-key-ring-id"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_kms\_key\_ring\_iam\_member
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
// 		_, err := kms.NewKeyRingIAMMember(ctx, "keyRing", &kms.KeyRingIAMMemberArgs{
// 			KeyRingId: pulumi.String("your-key-ring-id"),
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Role:      pulumi.String("roles/editor"),
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
// 		_, err := kms.NewKeyRingIAMMember(ctx, "keyRing", &kms.KeyRingIAMMemberArgs{
// 			Condition: &kms.KeyRingIAMMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			KeyRingId: pulumi.String("your-key-ring-id"),
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Role:      pulumi.String("roles/editor"),
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
// This member resource can be imported using the `key_ring_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:kms/keyRingIAMBinding:KeyRingIAMBinding key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `key_ring_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/keyRingIAMBinding:KeyRingIAMBinding key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `key_ring_id`, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/keyRingIAMBinding:KeyRingIAMBinding key_ring_iam your-project-id/location-name/key-ring-name
// ```
type KeyRingIAMBinding struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition KeyRingIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the key ring's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringOutput      `pulumi:"keyRingId"`
	Members   pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewKeyRingIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewKeyRingIAMBinding(ctx *pulumi.Context,
	name string, args *KeyRingIAMBindingArgs, opts ...pulumi.ResourceOption) (*KeyRingIAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyRingId == nil {
		return nil, errors.New("invalid value for required argument 'KeyRingId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource KeyRingIAMBinding
	err := ctx.RegisterResource("gcp:kms/keyRingIAMBinding:KeyRingIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyRingIAMBinding gets an existing KeyRingIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRingIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyRingIAMBindingState, opts ...pulumi.ResourceOption) (*KeyRingIAMBinding, error) {
	var resource KeyRingIAMBinding
	err := ctx.ReadResource("gcp:kms/keyRingIAMBinding:KeyRingIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyRingIAMBinding resources.
type keyRingIAMBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *KeyRingIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the key ring's IAM policy.
	Etag *string `pulumi:"etag"`
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId *string  `pulumi:"keyRingId"`
	Members   []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type KeyRingIAMBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition KeyRingIAMBindingConditionPtrInput
	// (Computed) The etag of the key ring's IAM policy.
	Etag pulumi.StringPtrInput
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringPtrInput
	Members   pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (KeyRingIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingIAMBindingState)(nil)).Elem()
}

type keyRingIAMBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *KeyRingIAMBindingCondition `pulumi:"condition"`
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId string   `pulumi:"keyRingId"`
	Members   []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a KeyRingIAMBinding resource.
type KeyRingIAMBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition KeyRingIAMBindingConditionPtrInput
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringInput
	Members   pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (KeyRingIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingIAMBindingArgs)(nil)).Elem()
}

type KeyRingIAMBindingInput interface {
	pulumi.Input

	ToKeyRingIAMBindingOutput() KeyRingIAMBindingOutput
	ToKeyRingIAMBindingOutputWithContext(ctx context.Context) KeyRingIAMBindingOutput
}

func (*KeyRingIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRingIAMBinding)(nil)).Elem()
}

func (i *KeyRingIAMBinding) ToKeyRingIAMBindingOutput() KeyRingIAMBindingOutput {
	return i.ToKeyRingIAMBindingOutputWithContext(context.Background())
}

func (i *KeyRingIAMBinding) ToKeyRingIAMBindingOutputWithContext(ctx context.Context) KeyRingIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingIAMBindingOutput)
}

// KeyRingIAMBindingArrayInput is an input type that accepts KeyRingIAMBindingArray and KeyRingIAMBindingArrayOutput values.
// You can construct a concrete instance of `KeyRingIAMBindingArrayInput` via:
//
//          KeyRingIAMBindingArray{ KeyRingIAMBindingArgs{...} }
type KeyRingIAMBindingArrayInput interface {
	pulumi.Input

	ToKeyRingIAMBindingArrayOutput() KeyRingIAMBindingArrayOutput
	ToKeyRingIAMBindingArrayOutputWithContext(context.Context) KeyRingIAMBindingArrayOutput
}

type KeyRingIAMBindingArray []KeyRingIAMBindingInput

func (KeyRingIAMBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyRingIAMBinding)(nil)).Elem()
}

func (i KeyRingIAMBindingArray) ToKeyRingIAMBindingArrayOutput() KeyRingIAMBindingArrayOutput {
	return i.ToKeyRingIAMBindingArrayOutputWithContext(context.Background())
}

func (i KeyRingIAMBindingArray) ToKeyRingIAMBindingArrayOutputWithContext(ctx context.Context) KeyRingIAMBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingIAMBindingArrayOutput)
}

// KeyRingIAMBindingMapInput is an input type that accepts KeyRingIAMBindingMap and KeyRingIAMBindingMapOutput values.
// You can construct a concrete instance of `KeyRingIAMBindingMapInput` via:
//
//          KeyRingIAMBindingMap{ "key": KeyRingIAMBindingArgs{...} }
type KeyRingIAMBindingMapInput interface {
	pulumi.Input

	ToKeyRingIAMBindingMapOutput() KeyRingIAMBindingMapOutput
	ToKeyRingIAMBindingMapOutputWithContext(context.Context) KeyRingIAMBindingMapOutput
}

type KeyRingIAMBindingMap map[string]KeyRingIAMBindingInput

func (KeyRingIAMBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyRingIAMBinding)(nil)).Elem()
}

func (i KeyRingIAMBindingMap) ToKeyRingIAMBindingMapOutput() KeyRingIAMBindingMapOutput {
	return i.ToKeyRingIAMBindingMapOutputWithContext(context.Background())
}

func (i KeyRingIAMBindingMap) ToKeyRingIAMBindingMapOutputWithContext(ctx context.Context) KeyRingIAMBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingIAMBindingMapOutput)
}

type KeyRingIAMBindingOutput struct{ *pulumi.OutputState }

func (KeyRingIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRingIAMBinding)(nil)).Elem()
}

func (o KeyRingIAMBindingOutput) ToKeyRingIAMBindingOutput() KeyRingIAMBindingOutput {
	return o
}

func (o KeyRingIAMBindingOutput) ToKeyRingIAMBindingOutputWithContext(ctx context.Context) KeyRingIAMBindingOutput {
	return o
}

type KeyRingIAMBindingArrayOutput struct{ *pulumi.OutputState }

func (KeyRingIAMBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyRingIAMBinding)(nil)).Elem()
}

func (o KeyRingIAMBindingArrayOutput) ToKeyRingIAMBindingArrayOutput() KeyRingIAMBindingArrayOutput {
	return o
}

func (o KeyRingIAMBindingArrayOutput) ToKeyRingIAMBindingArrayOutputWithContext(ctx context.Context) KeyRingIAMBindingArrayOutput {
	return o
}

func (o KeyRingIAMBindingArrayOutput) Index(i pulumi.IntInput) KeyRingIAMBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KeyRingIAMBinding {
		return vs[0].([]*KeyRingIAMBinding)[vs[1].(int)]
	}).(KeyRingIAMBindingOutput)
}

type KeyRingIAMBindingMapOutput struct{ *pulumi.OutputState }

func (KeyRingIAMBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyRingIAMBinding)(nil)).Elem()
}

func (o KeyRingIAMBindingMapOutput) ToKeyRingIAMBindingMapOutput() KeyRingIAMBindingMapOutput {
	return o
}

func (o KeyRingIAMBindingMapOutput) ToKeyRingIAMBindingMapOutputWithContext(ctx context.Context) KeyRingIAMBindingMapOutput {
	return o
}

func (o KeyRingIAMBindingMapOutput) MapIndex(k pulumi.StringInput) KeyRingIAMBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KeyRingIAMBinding {
		return vs[0].(map[string]*KeyRingIAMBinding)[vs[1].(string)]
	}).(KeyRingIAMBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRingIAMBindingInput)(nil)).Elem(), &KeyRingIAMBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRingIAMBindingArrayInput)(nil)).Elem(), KeyRingIAMBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRingIAMBindingMapInput)(nil)).Elem(), KeyRingIAMBindingMap{})
	pulumi.RegisterOutputType(KeyRingIAMBindingOutput{})
	pulumi.RegisterOutputType(KeyRingIAMBindingArrayOutput{})
	pulumi.RegisterOutputType(KeyRingIAMBindingMapOutput{})
}

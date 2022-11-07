// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Healthcare ConsentStore. Each of these resources serves a different use case:
//
// * `healthcare.ConsentStoreIamPolicy`: Authoritative. Sets the IAM policy for the consentstore and replaces any existing policy already attached.
// * `healthcare.ConsentStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the consentstore are preserved.
// * `healthcare.ConsentStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the consentstore are preserved.
//
// > **Note:** `healthcare.ConsentStoreIamPolicy` **cannot** be used in conjunction with `healthcare.ConsentStoreIamBinding` and `healthcare.ConsentStoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.ConsentStoreIamBinding` resources **can be** used in conjunction with `healthcare.ConsentStoreIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_healthcare\_consent\_store\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					organizations.GetIAMPolicyBinding{
//						Role: "roles/viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = healthcare.NewConsentStoreIamPolicy(ctx, "policy", &healthcare.ConsentStoreIamPolicyArgs{
//				Dataset:        pulumi.Any(google_healthcare_consent_store.MyConsent.Dataset),
//				ConsentStoreId: pulumi.Any(google_healthcare_consent_store.MyConsent.Name),
//				PolicyData:     pulumi.String(admin.PolicyData),
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
// ## google\_healthcare\_consent\_store\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := healthcare.NewConsentStoreIamBinding(ctx, "binding", &healthcare.ConsentStoreIamBindingArgs{
//				Dataset:        pulumi.Any(google_healthcare_consent_store.MyConsent.Dataset),
//				ConsentStoreId: pulumi.Any(google_healthcare_consent_store.MyConsent.Name),
//				Role:           pulumi.String("roles/viewer"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
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
// ## google\_healthcare\_consent\_store\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := healthcare.NewConsentStoreIamMember(ctx, "member", &healthcare.ConsentStoreIamMemberArgs{
//				Dataset:        pulumi.Any(google_healthcare_consent_store.MyConsent.Dataset),
//				ConsentStoreId: pulumi.Any(google_healthcare_consent_store.MyConsent.Name),
//				Role:           pulumi.String("roles/viewer"),
//				Member:         pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* {{dataset}}/consentStores/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Healthcare consentstore IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:healthcare/consentStoreIamBinding:ConsentStoreIamBinding editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:healthcare/consentStoreIamBinding:ConsentStoreIamBinding editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:healthcare/consentStoreIamBinding:ConsentStoreIamBinding editor {{dataset}}/consentStores/{{consent_store}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ConsentStoreIamBinding struct {
	pulumi.CustomResourceState

	Condition ConsentStoreIamBindingConditionPtrOutput `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId pulumi.StringOutput `pulumi:"consentStoreId"`
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewConsentStoreIamBinding registers a new resource with the given unique name, arguments, and options.
func NewConsentStoreIamBinding(ctx *pulumi.Context,
	name string, args *ConsentStoreIamBindingArgs, opts ...pulumi.ResourceOption) (*ConsentStoreIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsentStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConsentStoreId'")
	}
	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ConsentStoreIamBinding
	err := ctx.RegisterResource("gcp:healthcare/consentStoreIamBinding:ConsentStoreIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsentStoreIamBinding gets an existing ConsentStoreIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsentStoreIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsentStoreIamBindingState, opts ...pulumi.ResourceOption) (*ConsentStoreIamBinding, error) {
	var resource ConsentStoreIamBinding
	err := ctx.ReadResource("gcp:healthcare/consentStoreIamBinding:ConsentStoreIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsentStoreIamBinding resources.
type consentStoreIamBindingState struct {
	Condition *ConsentStoreIamBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId *string `pulumi:"consentStoreId"`
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset *string `pulumi:"dataset"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ConsentStoreIamBindingState struct {
	Condition ConsentStoreIamBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId pulumi.StringPtrInput
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ConsentStoreIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*consentStoreIamBindingState)(nil)).Elem()
}

type consentStoreIamBindingArgs struct {
	Condition *ConsentStoreIamBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId string `pulumi:"consentStoreId"`
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset string   `pulumi:"dataset"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ConsentStoreIamBinding resource.
type ConsentStoreIamBindingArgs struct {
	Condition ConsentStoreIamBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId pulumi.StringInput
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset pulumi.StringInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ConsentStoreIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consentStoreIamBindingArgs)(nil)).Elem()
}

type ConsentStoreIamBindingInput interface {
	pulumi.Input

	ToConsentStoreIamBindingOutput() ConsentStoreIamBindingOutput
	ToConsentStoreIamBindingOutputWithContext(ctx context.Context) ConsentStoreIamBindingOutput
}

func (*ConsentStoreIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsentStoreIamBinding)(nil)).Elem()
}

func (i *ConsentStoreIamBinding) ToConsentStoreIamBindingOutput() ConsentStoreIamBindingOutput {
	return i.ToConsentStoreIamBindingOutputWithContext(context.Background())
}

func (i *ConsentStoreIamBinding) ToConsentStoreIamBindingOutputWithContext(ctx context.Context) ConsentStoreIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentStoreIamBindingOutput)
}

// ConsentStoreIamBindingArrayInput is an input type that accepts ConsentStoreIamBindingArray and ConsentStoreIamBindingArrayOutput values.
// You can construct a concrete instance of `ConsentStoreIamBindingArrayInput` via:
//
//	ConsentStoreIamBindingArray{ ConsentStoreIamBindingArgs{...} }
type ConsentStoreIamBindingArrayInput interface {
	pulumi.Input

	ToConsentStoreIamBindingArrayOutput() ConsentStoreIamBindingArrayOutput
	ToConsentStoreIamBindingArrayOutputWithContext(context.Context) ConsentStoreIamBindingArrayOutput
}

type ConsentStoreIamBindingArray []ConsentStoreIamBindingInput

func (ConsentStoreIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsentStoreIamBinding)(nil)).Elem()
}

func (i ConsentStoreIamBindingArray) ToConsentStoreIamBindingArrayOutput() ConsentStoreIamBindingArrayOutput {
	return i.ToConsentStoreIamBindingArrayOutputWithContext(context.Background())
}

func (i ConsentStoreIamBindingArray) ToConsentStoreIamBindingArrayOutputWithContext(ctx context.Context) ConsentStoreIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentStoreIamBindingArrayOutput)
}

// ConsentStoreIamBindingMapInput is an input type that accepts ConsentStoreIamBindingMap and ConsentStoreIamBindingMapOutput values.
// You can construct a concrete instance of `ConsentStoreIamBindingMapInput` via:
//
//	ConsentStoreIamBindingMap{ "key": ConsentStoreIamBindingArgs{...} }
type ConsentStoreIamBindingMapInput interface {
	pulumi.Input

	ToConsentStoreIamBindingMapOutput() ConsentStoreIamBindingMapOutput
	ToConsentStoreIamBindingMapOutputWithContext(context.Context) ConsentStoreIamBindingMapOutput
}

type ConsentStoreIamBindingMap map[string]ConsentStoreIamBindingInput

func (ConsentStoreIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsentStoreIamBinding)(nil)).Elem()
}

func (i ConsentStoreIamBindingMap) ToConsentStoreIamBindingMapOutput() ConsentStoreIamBindingMapOutput {
	return i.ToConsentStoreIamBindingMapOutputWithContext(context.Background())
}

func (i ConsentStoreIamBindingMap) ToConsentStoreIamBindingMapOutputWithContext(ctx context.Context) ConsentStoreIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentStoreIamBindingMapOutput)
}

type ConsentStoreIamBindingOutput struct{ *pulumi.OutputState }

func (ConsentStoreIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsentStoreIamBinding)(nil)).Elem()
}

func (o ConsentStoreIamBindingOutput) ToConsentStoreIamBindingOutput() ConsentStoreIamBindingOutput {
	return o
}

func (o ConsentStoreIamBindingOutput) ToConsentStoreIamBindingOutputWithContext(ctx context.Context) ConsentStoreIamBindingOutput {
	return o
}

func (o ConsentStoreIamBindingOutput) Condition() ConsentStoreIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *ConsentStoreIamBinding) ConsentStoreIamBindingConditionPtrOutput { return v.Condition }).(ConsentStoreIamBindingConditionPtrOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o ConsentStoreIamBindingOutput) ConsentStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsentStoreIamBinding) pulumi.StringOutput { return v.ConsentStoreId }).(pulumi.StringOutput)
}

// Identifies the dataset addressed by this request. Must be in the format
// 'projects/{project}/locations/{location}/datasets/{dataset}'
// Used to find the parent resource to bind the IAM policy to
func (o ConsentStoreIamBindingOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsentStoreIamBinding) pulumi.StringOutput { return v.Dataset }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o ConsentStoreIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsentStoreIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ConsentStoreIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConsentStoreIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The role that should be applied. Only one
// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o ConsentStoreIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsentStoreIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type ConsentStoreIamBindingArrayOutput struct{ *pulumi.OutputState }

func (ConsentStoreIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsentStoreIamBinding)(nil)).Elem()
}

func (o ConsentStoreIamBindingArrayOutput) ToConsentStoreIamBindingArrayOutput() ConsentStoreIamBindingArrayOutput {
	return o
}

func (o ConsentStoreIamBindingArrayOutput) ToConsentStoreIamBindingArrayOutputWithContext(ctx context.Context) ConsentStoreIamBindingArrayOutput {
	return o
}

func (o ConsentStoreIamBindingArrayOutput) Index(i pulumi.IntInput) ConsentStoreIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConsentStoreIamBinding {
		return vs[0].([]*ConsentStoreIamBinding)[vs[1].(int)]
	}).(ConsentStoreIamBindingOutput)
}

type ConsentStoreIamBindingMapOutput struct{ *pulumi.OutputState }

func (ConsentStoreIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsentStoreIamBinding)(nil)).Elem()
}

func (o ConsentStoreIamBindingMapOutput) ToConsentStoreIamBindingMapOutput() ConsentStoreIamBindingMapOutput {
	return o
}

func (o ConsentStoreIamBindingMapOutput) ToConsentStoreIamBindingMapOutputWithContext(ctx context.Context) ConsentStoreIamBindingMapOutput {
	return o
}

func (o ConsentStoreIamBindingMapOutput) MapIndex(k pulumi.StringInput) ConsentStoreIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConsentStoreIamBinding {
		return vs[0].(map[string]*ConsentStoreIamBinding)[vs[1].(string)]
	}).(ConsentStoreIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentStoreIamBindingInput)(nil)).Elem(), &ConsentStoreIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentStoreIamBindingArrayInput)(nil)).Elem(), ConsentStoreIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentStoreIamBindingMapInput)(nil)).Elem(), ConsentStoreIamBindingMap{})
	pulumi.RegisterOutputType(ConsentStoreIamBindingOutput{})
	pulumi.RegisterOutputType(ConsentStoreIamBindingArrayOutput{})
	pulumi.RegisterOutputType(ConsentStoreIamBindingMapOutput{})
}

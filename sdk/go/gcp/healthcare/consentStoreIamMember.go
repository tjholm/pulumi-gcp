// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = healthcare.NewConsentStoreIamPolicy(ctx, "policy", &healthcare.ConsentStoreIamPolicyArgs{
// 			Dataset:        pulumi.Any(google_healthcare_consent_store.My - consent.Dataset),
// 			ConsentStoreId: pulumi.Any(google_healthcare_consent_store.My - consent.Name),
// 			PolicyData:     pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_healthcare\_consent\_store\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewConsentStoreIamBinding(ctx, "binding", &healthcare.ConsentStoreIamBindingArgs{
// 			Dataset:        pulumi.Any(google_healthcare_consent_store.My - consent.Dataset),
// 			ConsentStoreId: pulumi.Any(google_healthcare_consent_store.My - consent.Name),
// 			Role:           pulumi.String("roles/viewer"),
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
// ## google\_healthcare\_consent\_store\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewConsentStoreIamMember(ctx, "member", &healthcare.ConsentStoreIamMemberArgs{
// 			Dataset:        pulumi.Any(google_healthcare_consent_store.My - consent.Dataset),
// 			ConsentStoreId: pulumi.Any(google_healthcare_consent_store.My - consent.Name),
// 			Role:           pulumi.String("roles/viewer"),
// 			Member:         pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* {{dataset}}/consentStores/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Healthcare consentstore IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/consentStoreIamMember:ConsentStoreIamMember editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/consentStoreIamMember:ConsentStoreIamMember editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/consentStoreIamMember:ConsentStoreIamMember editor {{dataset}}/consentStores/{{consent_store}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ConsentStoreIamMember struct {
	pulumi.CustomResourceState

	Condition ConsentStoreIamMemberConditionPtrOutput `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId pulumi.StringOutput `pulumi:"consentStoreId"`
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewConsentStoreIamMember registers a new resource with the given unique name, arguments, and options.
func NewConsentStoreIamMember(ctx *pulumi.Context,
	name string, args *ConsentStoreIamMemberArgs, opts ...pulumi.ResourceOption) (*ConsentStoreIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsentStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConsentStoreId'")
	}
	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ConsentStoreIamMember
	err := ctx.RegisterResource("gcp:healthcare/consentStoreIamMember:ConsentStoreIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsentStoreIamMember gets an existing ConsentStoreIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsentStoreIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsentStoreIamMemberState, opts ...pulumi.ResourceOption) (*ConsentStoreIamMember, error) {
	var resource ConsentStoreIamMember
	err := ctx.ReadResource("gcp:healthcare/consentStoreIamMember:ConsentStoreIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsentStoreIamMember resources.
type consentStoreIamMemberState struct {
	Condition *ConsentStoreIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId *string `pulumi:"consentStoreId"`
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset *string `pulumi:"dataset"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ConsentStoreIamMemberState struct {
	Condition ConsentStoreIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId pulumi.StringPtrInput
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ConsentStoreIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*consentStoreIamMemberState)(nil)).Elem()
}

type consentStoreIamMemberArgs struct {
	Condition *ConsentStoreIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId string `pulumi:"consentStoreId"`
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset string `pulumi:"dataset"`
	Member  string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ConsentStoreIamMember resource.
type ConsentStoreIamMemberArgs struct {
	Condition ConsentStoreIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId pulumi.StringInput
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset pulumi.StringInput
	Member  pulumi.StringInput
	// The role that should be applied. Only one
	// `healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ConsentStoreIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consentStoreIamMemberArgs)(nil)).Elem()
}

type ConsentStoreIamMemberInput interface {
	pulumi.Input

	ToConsentStoreIamMemberOutput() ConsentStoreIamMemberOutput
	ToConsentStoreIamMemberOutputWithContext(ctx context.Context) ConsentStoreIamMemberOutput
}

func (*ConsentStoreIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsentStoreIamMember)(nil)).Elem()
}

func (i *ConsentStoreIamMember) ToConsentStoreIamMemberOutput() ConsentStoreIamMemberOutput {
	return i.ToConsentStoreIamMemberOutputWithContext(context.Background())
}

func (i *ConsentStoreIamMember) ToConsentStoreIamMemberOutputWithContext(ctx context.Context) ConsentStoreIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentStoreIamMemberOutput)
}

// ConsentStoreIamMemberArrayInput is an input type that accepts ConsentStoreIamMemberArray and ConsentStoreIamMemberArrayOutput values.
// You can construct a concrete instance of `ConsentStoreIamMemberArrayInput` via:
//
//          ConsentStoreIamMemberArray{ ConsentStoreIamMemberArgs{...} }
type ConsentStoreIamMemberArrayInput interface {
	pulumi.Input

	ToConsentStoreIamMemberArrayOutput() ConsentStoreIamMemberArrayOutput
	ToConsentStoreIamMemberArrayOutputWithContext(context.Context) ConsentStoreIamMemberArrayOutput
}

type ConsentStoreIamMemberArray []ConsentStoreIamMemberInput

func (ConsentStoreIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsentStoreIamMember)(nil)).Elem()
}

func (i ConsentStoreIamMemberArray) ToConsentStoreIamMemberArrayOutput() ConsentStoreIamMemberArrayOutput {
	return i.ToConsentStoreIamMemberArrayOutputWithContext(context.Background())
}

func (i ConsentStoreIamMemberArray) ToConsentStoreIamMemberArrayOutputWithContext(ctx context.Context) ConsentStoreIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentStoreIamMemberArrayOutput)
}

// ConsentStoreIamMemberMapInput is an input type that accepts ConsentStoreIamMemberMap and ConsentStoreIamMemberMapOutput values.
// You can construct a concrete instance of `ConsentStoreIamMemberMapInput` via:
//
//          ConsentStoreIamMemberMap{ "key": ConsentStoreIamMemberArgs{...} }
type ConsentStoreIamMemberMapInput interface {
	pulumi.Input

	ToConsentStoreIamMemberMapOutput() ConsentStoreIamMemberMapOutput
	ToConsentStoreIamMemberMapOutputWithContext(context.Context) ConsentStoreIamMemberMapOutput
}

type ConsentStoreIamMemberMap map[string]ConsentStoreIamMemberInput

func (ConsentStoreIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsentStoreIamMember)(nil)).Elem()
}

func (i ConsentStoreIamMemberMap) ToConsentStoreIamMemberMapOutput() ConsentStoreIamMemberMapOutput {
	return i.ToConsentStoreIamMemberMapOutputWithContext(context.Background())
}

func (i ConsentStoreIamMemberMap) ToConsentStoreIamMemberMapOutputWithContext(ctx context.Context) ConsentStoreIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentStoreIamMemberMapOutput)
}

type ConsentStoreIamMemberOutput struct{ *pulumi.OutputState }

func (ConsentStoreIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsentStoreIamMember)(nil)).Elem()
}

func (o ConsentStoreIamMemberOutput) ToConsentStoreIamMemberOutput() ConsentStoreIamMemberOutput {
	return o
}

func (o ConsentStoreIamMemberOutput) ToConsentStoreIamMemberOutputWithContext(ctx context.Context) ConsentStoreIamMemberOutput {
	return o
}

type ConsentStoreIamMemberArrayOutput struct{ *pulumi.OutputState }

func (ConsentStoreIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsentStoreIamMember)(nil)).Elem()
}

func (o ConsentStoreIamMemberArrayOutput) ToConsentStoreIamMemberArrayOutput() ConsentStoreIamMemberArrayOutput {
	return o
}

func (o ConsentStoreIamMemberArrayOutput) ToConsentStoreIamMemberArrayOutputWithContext(ctx context.Context) ConsentStoreIamMemberArrayOutput {
	return o
}

func (o ConsentStoreIamMemberArrayOutput) Index(i pulumi.IntInput) ConsentStoreIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConsentStoreIamMember {
		return vs[0].([]*ConsentStoreIamMember)[vs[1].(int)]
	}).(ConsentStoreIamMemberOutput)
}

type ConsentStoreIamMemberMapOutput struct{ *pulumi.OutputState }

func (ConsentStoreIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsentStoreIamMember)(nil)).Elem()
}

func (o ConsentStoreIamMemberMapOutput) ToConsentStoreIamMemberMapOutput() ConsentStoreIamMemberMapOutput {
	return o
}

func (o ConsentStoreIamMemberMapOutput) ToConsentStoreIamMemberMapOutputWithContext(ctx context.Context) ConsentStoreIamMemberMapOutput {
	return o
}

func (o ConsentStoreIamMemberMapOutput) MapIndex(k pulumi.StringInput) ConsentStoreIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConsentStoreIamMember {
		return vs[0].(map[string]*ConsentStoreIamMember)[vs[1].(string)]
	}).(ConsentStoreIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentStoreIamMemberInput)(nil)).Elem(), &ConsentStoreIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentStoreIamMemberArrayInput)(nil)).Elem(), ConsentStoreIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentStoreIamMemberMapInput)(nil)).Elem(), ConsentStoreIamMemberMap{})
	pulumi.RegisterOutputType(ConsentStoreIamMemberOutput{})
	pulumi.RegisterOutputType(ConsentStoreIamMemberArrayOutput{})
	pulumi.RegisterOutputType(ConsentStoreIamMemberMapOutput{})
}

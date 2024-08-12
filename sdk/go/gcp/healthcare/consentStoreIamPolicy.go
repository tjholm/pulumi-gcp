// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Healthcare ConsentStore. Each of these resources serves a different use case:
//
// * `healthcare.ConsentStoreIamPolicy`: Authoritative. Sets the IAM policy for the consentstore and replaces any existing policy already attached.
// * `healthcare.ConsentStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the consentstore are preserved.
// * `healthcare.ConsentStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the consentstore are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `healthcare.ConsentStoreIamPolicy`: Retrieves the IAM policy for the consentstore
//
// > **Note:** `healthcare.ConsentStoreIamPolicy` **cannot** be used in conjunction with `healthcare.ConsentStoreIamBinding` and `healthcare.ConsentStoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.ConsentStoreIamBinding` resources **can be** used in conjunction with `healthcare.ConsentStoreIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## healthcare.ConsentStoreIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/healthcare"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
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
//				Dataset:        pulumi.Any(my_consent.Dataset),
//				ConsentStoreId: pulumi.Any(my_consent.Name),
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
// ## healthcare.ConsentStoreIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/healthcare"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := healthcare.NewConsentStoreIamBinding(ctx, "binding", &healthcare.ConsentStoreIamBindingArgs{
//				Dataset:        pulumi.Any(my_consent.Dataset),
//				ConsentStoreId: pulumi.Any(my_consent.Name),
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
// ## healthcare.ConsentStoreIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/healthcare"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := healthcare.NewConsentStoreIamMember(ctx, "member", &healthcare.ConsentStoreIamMemberArgs{
//				Dataset:        pulumi.Any(my_consent.Dataset),
//				ConsentStoreId: pulumi.Any(my_consent.Name),
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
// ## > **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
// ---
//
// # IAM policy for Cloud Healthcare ConsentStore
// Three different resources help you manage your IAM policy for Cloud Healthcare ConsentStore. Each of these resources serves a different use case:
//
// * `healthcare.ConsentStoreIamPolicy`: Authoritative. Sets the IAM policy for the consentstore and replaces any existing policy already attached.
// * `healthcare.ConsentStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the consentstore are preserved.
// * `healthcare.ConsentStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the consentstore are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `healthcare.ConsentStoreIamPolicy`: Retrieves the IAM policy for the consentstore
//
// > **Note:** `healthcare.ConsentStoreIamPolicy` **cannot** be used in conjunction with `healthcare.ConsentStoreIamBinding` and `healthcare.ConsentStoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.ConsentStoreIamBinding` resources **can be** used in conjunction with `healthcare.ConsentStoreIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## healthcare.ConsentStoreIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/healthcare"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
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
//				Dataset:        pulumi.Any(my_consent.Dataset),
//				ConsentStoreId: pulumi.Any(my_consent.Name),
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
// ## healthcare.ConsentStoreIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/healthcare"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := healthcare.NewConsentStoreIamBinding(ctx, "binding", &healthcare.ConsentStoreIamBindingArgs{
//				Dataset:        pulumi.Any(my_consent.Dataset),
//				ConsentStoreId: pulumi.Any(my_consent.Name),
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
// ## healthcare.ConsentStoreIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/healthcare"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := healthcare.NewConsentStoreIamMember(ctx, "member", &healthcare.ConsentStoreIamMemberArgs{
//				Dataset:        pulumi.Any(my_consent.Dataset),
//				ConsentStoreId: pulumi.Any(my_consent.Name),
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
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * {{dataset}}/consentStores/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Cloud Healthcare consentstore IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy editor {{dataset}}/consentStores/{{consent_store}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ConsentStoreIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId pulumi.StringOutput `pulumi:"consentStoreId"`
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewConsentStoreIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewConsentStoreIamPolicy(ctx *pulumi.Context,
	name string, args *ConsentStoreIamPolicyArgs, opts ...pulumi.ResourceOption) (*ConsentStoreIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsentStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConsentStoreId'")
	}
	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConsentStoreIamPolicy
	err := ctx.RegisterResource("gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsentStoreIamPolicy gets an existing ConsentStoreIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsentStoreIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsentStoreIamPolicyState, opts ...pulumi.ResourceOption) (*ConsentStoreIamPolicy, error) {
	var resource ConsentStoreIamPolicy
	err := ctx.ReadResource("gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsentStoreIamPolicy resources.
type consentStoreIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId *string `pulumi:"consentStoreId"`
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset *string `pulumi:"dataset"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type ConsentStoreIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId pulumi.StringPtrInput
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (ConsentStoreIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*consentStoreIamPolicyState)(nil)).Elem()
}

type consentStoreIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId string `pulumi:"consentStoreId"`
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset string `pulumi:"dataset"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a ConsentStoreIamPolicy resource.
type ConsentStoreIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId pulumi.StringInput
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// Used to find the parent resource to bind the IAM policy to
	Dataset pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (ConsentStoreIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consentStoreIamPolicyArgs)(nil)).Elem()
}

type ConsentStoreIamPolicyInput interface {
	pulumi.Input

	ToConsentStoreIamPolicyOutput() ConsentStoreIamPolicyOutput
	ToConsentStoreIamPolicyOutputWithContext(ctx context.Context) ConsentStoreIamPolicyOutput
}

func (*ConsentStoreIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsentStoreIamPolicy)(nil)).Elem()
}

func (i *ConsentStoreIamPolicy) ToConsentStoreIamPolicyOutput() ConsentStoreIamPolicyOutput {
	return i.ToConsentStoreIamPolicyOutputWithContext(context.Background())
}

func (i *ConsentStoreIamPolicy) ToConsentStoreIamPolicyOutputWithContext(ctx context.Context) ConsentStoreIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentStoreIamPolicyOutput)
}

// ConsentStoreIamPolicyArrayInput is an input type that accepts ConsentStoreIamPolicyArray and ConsentStoreIamPolicyArrayOutput values.
// You can construct a concrete instance of `ConsentStoreIamPolicyArrayInput` via:
//
//	ConsentStoreIamPolicyArray{ ConsentStoreIamPolicyArgs{...} }
type ConsentStoreIamPolicyArrayInput interface {
	pulumi.Input

	ToConsentStoreIamPolicyArrayOutput() ConsentStoreIamPolicyArrayOutput
	ToConsentStoreIamPolicyArrayOutputWithContext(context.Context) ConsentStoreIamPolicyArrayOutput
}

type ConsentStoreIamPolicyArray []ConsentStoreIamPolicyInput

func (ConsentStoreIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsentStoreIamPolicy)(nil)).Elem()
}

func (i ConsentStoreIamPolicyArray) ToConsentStoreIamPolicyArrayOutput() ConsentStoreIamPolicyArrayOutput {
	return i.ToConsentStoreIamPolicyArrayOutputWithContext(context.Background())
}

func (i ConsentStoreIamPolicyArray) ToConsentStoreIamPolicyArrayOutputWithContext(ctx context.Context) ConsentStoreIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentStoreIamPolicyArrayOutput)
}

// ConsentStoreIamPolicyMapInput is an input type that accepts ConsentStoreIamPolicyMap and ConsentStoreIamPolicyMapOutput values.
// You can construct a concrete instance of `ConsentStoreIamPolicyMapInput` via:
//
//	ConsentStoreIamPolicyMap{ "key": ConsentStoreIamPolicyArgs{...} }
type ConsentStoreIamPolicyMapInput interface {
	pulumi.Input

	ToConsentStoreIamPolicyMapOutput() ConsentStoreIamPolicyMapOutput
	ToConsentStoreIamPolicyMapOutputWithContext(context.Context) ConsentStoreIamPolicyMapOutput
}

type ConsentStoreIamPolicyMap map[string]ConsentStoreIamPolicyInput

func (ConsentStoreIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsentStoreIamPolicy)(nil)).Elem()
}

func (i ConsentStoreIamPolicyMap) ToConsentStoreIamPolicyMapOutput() ConsentStoreIamPolicyMapOutput {
	return i.ToConsentStoreIamPolicyMapOutputWithContext(context.Background())
}

func (i ConsentStoreIamPolicyMap) ToConsentStoreIamPolicyMapOutputWithContext(ctx context.Context) ConsentStoreIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentStoreIamPolicyMapOutput)
}

type ConsentStoreIamPolicyOutput struct{ *pulumi.OutputState }

func (ConsentStoreIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsentStoreIamPolicy)(nil)).Elem()
}

func (o ConsentStoreIamPolicyOutput) ToConsentStoreIamPolicyOutput() ConsentStoreIamPolicyOutput {
	return o
}

func (o ConsentStoreIamPolicyOutput) ToConsentStoreIamPolicyOutputWithContext(ctx context.Context) ConsentStoreIamPolicyOutput {
	return o
}

// Used to find the parent resource to bind the IAM policy to
func (o ConsentStoreIamPolicyOutput) ConsentStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsentStoreIamPolicy) pulumi.StringOutput { return v.ConsentStoreId }).(pulumi.StringOutput)
}

// Identifies the dataset addressed by this request. Must be in the format
// 'projects/{project}/locations/{location}/datasets/{dataset}'
// Used to find the parent resource to bind the IAM policy to
func (o ConsentStoreIamPolicyOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsentStoreIamPolicy) pulumi.StringOutput { return v.Dataset }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o ConsentStoreIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsentStoreIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o ConsentStoreIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsentStoreIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

type ConsentStoreIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (ConsentStoreIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsentStoreIamPolicy)(nil)).Elem()
}

func (o ConsentStoreIamPolicyArrayOutput) ToConsentStoreIamPolicyArrayOutput() ConsentStoreIamPolicyArrayOutput {
	return o
}

func (o ConsentStoreIamPolicyArrayOutput) ToConsentStoreIamPolicyArrayOutputWithContext(ctx context.Context) ConsentStoreIamPolicyArrayOutput {
	return o
}

func (o ConsentStoreIamPolicyArrayOutput) Index(i pulumi.IntInput) ConsentStoreIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConsentStoreIamPolicy {
		return vs[0].([]*ConsentStoreIamPolicy)[vs[1].(int)]
	}).(ConsentStoreIamPolicyOutput)
}

type ConsentStoreIamPolicyMapOutput struct{ *pulumi.OutputState }

func (ConsentStoreIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsentStoreIamPolicy)(nil)).Elem()
}

func (o ConsentStoreIamPolicyMapOutput) ToConsentStoreIamPolicyMapOutput() ConsentStoreIamPolicyMapOutput {
	return o
}

func (o ConsentStoreIamPolicyMapOutput) ToConsentStoreIamPolicyMapOutputWithContext(ctx context.Context) ConsentStoreIamPolicyMapOutput {
	return o
}

func (o ConsentStoreIamPolicyMapOutput) MapIndex(k pulumi.StringInput) ConsentStoreIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConsentStoreIamPolicy {
		return vs[0].(map[string]*ConsentStoreIamPolicy)[vs[1].(string)]
	}).(ConsentStoreIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentStoreIamPolicyInput)(nil)).Elem(), &ConsentStoreIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentStoreIamPolicyArrayInput)(nil)).Elem(), ConsentStoreIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentStoreIamPolicyMapInput)(nil)).Elem(), ConsentStoreIamPolicyMap{})
	pulumi.RegisterOutputType(ConsentStoreIamPolicyOutput{})
	pulumi.RegisterOutputType(ConsentStoreIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(ConsentStoreIamPolicyMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Three different resources help you manage your IAM policy for Healthcare FHIR store. Each of these resources serves a different use case:
//
// * `healthcare.FhirStoreIamPolicy`: Authoritative. Sets the IAM policy for the FHIR store and replaces any existing policy already attached.
// * `healthcare.FhirStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the FHIR store are preserved.
// * `healthcare.FhirStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the FHIR store are preserved.
//
// > **Note:** `healthcare.FhirStoreIamPolicy` **cannot** be used in conjunction with `healthcare.FhirStoreIamBinding` and `healthcare.FhirStoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.FhirStoreIamBinding` resources **can be** used in conjunction with `healthcare.FhirStoreIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_healthcare\_fhir\_store\_iam\_policy
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
//					{
//						Role: "roles/editor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = healthcare.NewFhirStoreIamPolicy(ctx, "fhirStore", &healthcare.FhirStoreIamPolicyArgs{
//				FhirStoreId: pulumi.String("your-fhir-store-id"),
//				PolicyData:  *pulumi.String(admin.PolicyData),
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
// ## google\_healthcare\_fhir\_store\_iam\_binding
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
//			_, err := healthcare.NewFhirStoreIamBinding(ctx, "fhirStore", &healthcare.FhirStoreIamBindingArgs{
//				FhirStoreId: pulumi.String("your-fhir-store-id"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Role: pulumi.String("roles/editor"),
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
// ## google\_healthcare\_fhir\_store\_iam\_member
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
//			_, err := healthcare.NewFhirStoreIamMember(ctx, "fhirStore", &healthcare.FhirStoreIamMemberArgs{
//				FhirStoreId: pulumi.String("your-fhir-store-id"),
//				Member:      pulumi.String("user:jane@example.com"),
//				Role:        pulumi.String("roles/editor"),
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `fhir_store_id`, role, and account e.g.
//
// ```sh
//
//	$ pulumi import gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer user:foo@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `fhir_store_id` and role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `fhir_store_id`, role, and account e.g.
//
// ```sh
//
//	$ pulumi import gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy fhir_store_iam your-project-id/location-name/dataset-name/fhir-store-name
//
// ```
type FhirStoreIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	FhirStoreId pulumi.StringOutput `pulumi:"fhirStoreId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewFhirStoreIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewFhirStoreIamPolicy(ctx *pulumi.Context,
	name string, args *FhirStoreIamPolicyArgs, opts ...pulumi.ResourceOption) (*FhirStoreIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FhirStoreId == nil {
		return nil, errors.New("invalid value for required argument 'FhirStoreId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FhirStoreIamPolicy
	err := ctx.RegisterResource("gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFhirStoreIamPolicy gets an existing FhirStoreIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirStoreIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FhirStoreIamPolicyState, opts ...pulumi.ResourceOption) (*FhirStoreIamPolicy, error) {
	var resource FhirStoreIamPolicy
	err := ctx.ReadResource("gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FhirStoreIamPolicy resources.
type fhirStoreIamPolicyState struct {
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag *string `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	FhirStoreId *string `pulumi:"fhirStoreId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type FhirStoreIamPolicyState struct {
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringPtrInput
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	FhirStoreId pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (FhirStoreIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreIamPolicyState)(nil)).Elem()
}

type fhirStoreIamPolicyArgs struct {
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	FhirStoreId string `pulumi:"fhirStoreId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a FhirStoreIamPolicy resource.
type FhirStoreIamPolicyArgs struct {
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	FhirStoreId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (FhirStoreIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreIamPolicyArgs)(nil)).Elem()
}

type FhirStoreIamPolicyInput interface {
	pulumi.Input

	ToFhirStoreIamPolicyOutput() FhirStoreIamPolicyOutput
	ToFhirStoreIamPolicyOutputWithContext(ctx context.Context) FhirStoreIamPolicyOutput
}

func (*FhirStoreIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirStoreIamPolicy)(nil)).Elem()
}

func (i *FhirStoreIamPolicy) ToFhirStoreIamPolicyOutput() FhirStoreIamPolicyOutput {
	return i.ToFhirStoreIamPolicyOutputWithContext(context.Background())
}

func (i *FhirStoreIamPolicy) ToFhirStoreIamPolicyOutputWithContext(ctx context.Context) FhirStoreIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamPolicyOutput)
}

func (i *FhirStoreIamPolicy) ToOutput(ctx context.Context) pulumix.Output[*FhirStoreIamPolicy] {
	return pulumix.Output[*FhirStoreIamPolicy]{
		OutputState: i.ToFhirStoreIamPolicyOutputWithContext(ctx).OutputState,
	}
}

// FhirStoreIamPolicyArrayInput is an input type that accepts FhirStoreIamPolicyArray and FhirStoreIamPolicyArrayOutput values.
// You can construct a concrete instance of `FhirStoreIamPolicyArrayInput` via:
//
//	FhirStoreIamPolicyArray{ FhirStoreIamPolicyArgs{...} }
type FhirStoreIamPolicyArrayInput interface {
	pulumi.Input

	ToFhirStoreIamPolicyArrayOutput() FhirStoreIamPolicyArrayOutput
	ToFhirStoreIamPolicyArrayOutputWithContext(context.Context) FhirStoreIamPolicyArrayOutput
}

type FhirStoreIamPolicyArray []FhirStoreIamPolicyInput

func (FhirStoreIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FhirStoreIamPolicy)(nil)).Elem()
}

func (i FhirStoreIamPolicyArray) ToFhirStoreIamPolicyArrayOutput() FhirStoreIamPolicyArrayOutput {
	return i.ToFhirStoreIamPolicyArrayOutputWithContext(context.Background())
}

func (i FhirStoreIamPolicyArray) ToFhirStoreIamPolicyArrayOutputWithContext(ctx context.Context) FhirStoreIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamPolicyArrayOutput)
}

func (i FhirStoreIamPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*FhirStoreIamPolicy] {
	return pulumix.Output[[]*FhirStoreIamPolicy]{
		OutputState: i.ToFhirStoreIamPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// FhirStoreIamPolicyMapInput is an input type that accepts FhirStoreIamPolicyMap and FhirStoreIamPolicyMapOutput values.
// You can construct a concrete instance of `FhirStoreIamPolicyMapInput` via:
//
//	FhirStoreIamPolicyMap{ "key": FhirStoreIamPolicyArgs{...} }
type FhirStoreIamPolicyMapInput interface {
	pulumi.Input

	ToFhirStoreIamPolicyMapOutput() FhirStoreIamPolicyMapOutput
	ToFhirStoreIamPolicyMapOutputWithContext(context.Context) FhirStoreIamPolicyMapOutput
}

type FhirStoreIamPolicyMap map[string]FhirStoreIamPolicyInput

func (FhirStoreIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FhirStoreIamPolicy)(nil)).Elem()
}

func (i FhirStoreIamPolicyMap) ToFhirStoreIamPolicyMapOutput() FhirStoreIamPolicyMapOutput {
	return i.ToFhirStoreIamPolicyMapOutputWithContext(context.Background())
}

func (i FhirStoreIamPolicyMap) ToFhirStoreIamPolicyMapOutputWithContext(ctx context.Context) FhirStoreIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamPolicyMapOutput)
}

func (i FhirStoreIamPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FhirStoreIamPolicy] {
	return pulumix.Output[map[string]*FhirStoreIamPolicy]{
		OutputState: i.ToFhirStoreIamPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type FhirStoreIamPolicyOutput struct{ *pulumi.OutputState }

func (FhirStoreIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirStoreIamPolicy)(nil)).Elem()
}

func (o FhirStoreIamPolicyOutput) ToFhirStoreIamPolicyOutput() FhirStoreIamPolicyOutput {
	return o
}

func (o FhirStoreIamPolicyOutput) ToFhirStoreIamPolicyOutputWithContext(ctx context.Context) FhirStoreIamPolicyOutput {
	return o
}

func (o FhirStoreIamPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*FhirStoreIamPolicy] {
	return pulumix.Output[*FhirStoreIamPolicy]{
		OutputState: o.OutputState,
	}
}

// (Computed) The etag of the FHIR store's IAM policy.
func (o FhirStoreIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirStoreIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The FHIR store ID, in the form
// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
// project setting will be used as a fallback.
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o FhirStoreIamPolicyOutput) FhirStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirStoreIamPolicy) pulumi.StringOutput { return v.FhirStoreId }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o FhirStoreIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirStoreIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

type FhirStoreIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (FhirStoreIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FhirStoreIamPolicy)(nil)).Elem()
}

func (o FhirStoreIamPolicyArrayOutput) ToFhirStoreIamPolicyArrayOutput() FhirStoreIamPolicyArrayOutput {
	return o
}

func (o FhirStoreIamPolicyArrayOutput) ToFhirStoreIamPolicyArrayOutputWithContext(ctx context.Context) FhirStoreIamPolicyArrayOutput {
	return o
}

func (o FhirStoreIamPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FhirStoreIamPolicy] {
	return pulumix.Output[[]*FhirStoreIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o FhirStoreIamPolicyArrayOutput) Index(i pulumi.IntInput) FhirStoreIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FhirStoreIamPolicy {
		return vs[0].([]*FhirStoreIamPolicy)[vs[1].(int)]
	}).(FhirStoreIamPolicyOutput)
}

type FhirStoreIamPolicyMapOutput struct{ *pulumi.OutputState }

func (FhirStoreIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FhirStoreIamPolicy)(nil)).Elem()
}

func (o FhirStoreIamPolicyMapOutput) ToFhirStoreIamPolicyMapOutput() FhirStoreIamPolicyMapOutput {
	return o
}

func (o FhirStoreIamPolicyMapOutput) ToFhirStoreIamPolicyMapOutputWithContext(ctx context.Context) FhirStoreIamPolicyMapOutput {
	return o
}

func (o FhirStoreIamPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FhirStoreIamPolicy] {
	return pulumix.Output[map[string]*FhirStoreIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o FhirStoreIamPolicyMapOutput) MapIndex(k pulumi.StringInput) FhirStoreIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FhirStoreIamPolicy {
		return vs[0].(map[string]*FhirStoreIamPolicy)[vs[1].(string)]
	}).(FhirStoreIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FhirStoreIamPolicyInput)(nil)).Elem(), &FhirStoreIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*FhirStoreIamPolicyArrayInput)(nil)).Elem(), FhirStoreIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FhirStoreIamPolicyMapInput)(nil)).Elem(), FhirStoreIamPolicyMap{})
	pulumi.RegisterOutputType(FhirStoreIamPolicyOutput{})
	pulumi.RegisterOutputType(FhirStoreIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(FhirStoreIamPolicyMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
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
// 		_, err = healthcare.NewFhirStoreIamPolicy(ctx, "fhirStore", &healthcare.FhirStoreIamPolicyArgs{
// 			FhirStoreId: pulumi.String("your-fhir-store-id"),
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
// ## google\_healthcare\_fhir\_store\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewFhirStoreIamBinding(ctx, "fhirStore", &healthcare.FhirStoreIamBindingArgs{
// 			FhirStoreId: pulumi.String("your-fhir-store-id"),
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
// ## google\_healthcare\_fhir\_store\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewFhirStoreIamMember(ctx, "fhirStore", &healthcare.FhirStoreIamMemberArgs{
// 			FhirStoreId: pulumi.String("your-fhir-store-id"),
// 			Member:      pulumi.String("user:jane@example.com"),
// 			Role:        pulumi.String("roles/editor"),
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
// This member resource can be imported using the `fhir_store_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `fhir_store_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `fhir_store_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy fhir_store_iam your-project-id/location-name/dataset-name/fhir-store-name
// ```
type FhirStoreIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
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
	return reflect.TypeOf((*FhirStoreIamPolicy)(nil))
}

func (i *FhirStoreIamPolicy) ToFhirStoreIamPolicyOutput() FhirStoreIamPolicyOutput {
	return i.ToFhirStoreIamPolicyOutputWithContext(context.Background())
}

func (i *FhirStoreIamPolicy) ToFhirStoreIamPolicyOutputWithContext(ctx context.Context) FhirStoreIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamPolicyOutput)
}

func (i *FhirStoreIamPolicy) ToFhirStoreIamPolicyPtrOutput() FhirStoreIamPolicyPtrOutput {
	return i.ToFhirStoreIamPolicyPtrOutputWithContext(context.Background())
}

func (i *FhirStoreIamPolicy) ToFhirStoreIamPolicyPtrOutputWithContext(ctx context.Context) FhirStoreIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamPolicyPtrOutput)
}

type FhirStoreIamPolicyPtrInput interface {
	pulumi.Input

	ToFhirStoreIamPolicyPtrOutput() FhirStoreIamPolicyPtrOutput
	ToFhirStoreIamPolicyPtrOutputWithContext(ctx context.Context) FhirStoreIamPolicyPtrOutput
}

type fhirStoreIamPolicyPtrType FhirStoreIamPolicyArgs

func (*fhirStoreIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirStoreIamPolicy)(nil))
}

func (i *fhirStoreIamPolicyPtrType) ToFhirStoreIamPolicyPtrOutput() FhirStoreIamPolicyPtrOutput {
	return i.ToFhirStoreIamPolicyPtrOutputWithContext(context.Background())
}

func (i *fhirStoreIamPolicyPtrType) ToFhirStoreIamPolicyPtrOutputWithContext(ctx context.Context) FhirStoreIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamPolicyPtrOutput)
}

// FhirStoreIamPolicyArrayInput is an input type that accepts FhirStoreIamPolicyArray and FhirStoreIamPolicyArrayOutput values.
// You can construct a concrete instance of `FhirStoreIamPolicyArrayInput` via:
//
//          FhirStoreIamPolicyArray{ FhirStoreIamPolicyArgs{...} }
type FhirStoreIamPolicyArrayInput interface {
	pulumi.Input

	ToFhirStoreIamPolicyArrayOutput() FhirStoreIamPolicyArrayOutput
	ToFhirStoreIamPolicyArrayOutputWithContext(context.Context) FhirStoreIamPolicyArrayOutput
}

type FhirStoreIamPolicyArray []FhirStoreIamPolicyInput

func (FhirStoreIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FhirStoreIamPolicy)(nil))
}

func (i FhirStoreIamPolicyArray) ToFhirStoreIamPolicyArrayOutput() FhirStoreIamPolicyArrayOutput {
	return i.ToFhirStoreIamPolicyArrayOutputWithContext(context.Background())
}

func (i FhirStoreIamPolicyArray) ToFhirStoreIamPolicyArrayOutputWithContext(ctx context.Context) FhirStoreIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamPolicyArrayOutput)
}

// FhirStoreIamPolicyMapInput is an input type that accepts FhirStoreIamPolicyMap and FhirStoreIamPolicyMapOutput values.
// You can construct a concrete instance of `FhirStoreIamPolicyMapInput` via:
//
//          FhirStoreIamPolicyMap{ "key": FhirStoreIamPolicyArgs{...} }
type FhirStoreIamPolicyMapInput interface {
	pulumi.Input

	ToFhirStoreIamPolicyMapOutput() FhirStoreIamPolicyMapOutput
	ToFhirStoreIamPolicyMapOutputWithContext(context.Context) FhirStoreIamPolicyMapOutput
}

type FhirStoreIamPolicyMap map[string]FhirStoreIamPolicyInput

func (FhirStoreIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FhirStoreIamPolicy)(nil))
}

func (i FhirStoreIamPolicyMap) ToFhirStoreIamPolicyMapOutput() FhirStoreIamPolicyMapOutput {
	return i.ToFhirStoreIamPolicyMapOutputWithContext(context.Background())
}

func (i FhirStoreIamPolicyMap) ToFhirStoreIamPolicyMapOutputWithContext(ctx context.Context) FhirStoreIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamPolicyMapOutput)
}

type FhirStoreIamPolicyOutput struct {
	*pulumi.OutputState
}

func (FhirStoreIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirStoreIamPolicy)(nil))
}

func (o FhirStoreIamPolicyOutput) ToFhirStoreIamPolicyOutput() FhirStoreIamPolicyOutput {
	return o
}

func (o FhirStoreIamPolicyOutput) ToFhirStoreIamPolicyOutputWithContext(ctx context.Context) FhirStoreIamPolicyOutput {
	return o
}

func (o FhirStoreIamPolicyOutput) ToFhirStoreIamPolicyPtrOutput() FhirStoreIamPolicyPtrOutput {
	return o.ToFhirStoreIamPolicyPtrOutputWithContext(context.Background())
}

func (o FhirStoreIamPolicyOutput) ToFhirStoreIamPolicyPtrOutputWithContext(ctx context.Context) FhirStoreIamPolicyPtrOutput {
	return o.ApplyT(func(v FhirStoreIamPolicy) *FhirStoreIamPolicy {
		return &v
	}).(FhirStoreIamPolicyPtrOutput)
}

type FhirStoreIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (FhirStoreIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirStoreIamPolicy)(nil))
}

func (o FhirStoreIamPolicyPtrOutput) ToFhirStoreIamPolicyPtrOutput() FhirStoreIamPolicyPtrOutput {
	return o
}

func (o FhirStoreIamPolicyPtrOutput) ToFhirStoreIamPolicyPtrOutputWithContext(ctx context.Context) FhirStoreIamPolicyPtrOutput {
	return o
}

type FhirStoreIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (FhirStoreIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FhirStoreIamPolicy)(nil))
}

func (o FhirStoreIamPolicyArrayOutput) ToFhirStoreIamPolicyArrayOutput() FhirStoreIamPolicyArrayOutput {
	return o
}

func (o FhirStoreIamPolicyArrayOutput) ToFhirStoreIamPolicyArrayOutputWithContext(ctx context.Context) FhirStoreIamPolicyArrayOutput {
	return o
}

func (o FhirStoreIamPolicyArrayOutput) Index(i pulumi.IntInput) FhirStoreIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FhirStoreIamPolicy {
		return vs[0].([]FhirStoreIamPolicy)[vs[1].(int)]
	}).(FhirStoreIamPolicyOutput)
}

type FhirStoreIamPolicyMapOutput struct{ *pulumi.OutputState }

func (FhirStoreIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FhirStoreIamPolicy)(nil))
}

func (o FhirStoreIamPolicyMapOutput) ToFhirStoreIamPolicyMapOutput() FhirStoreIamPolicyMapOutput {
	return o
}

func (o FhirStoreIamPolicyMapOutput) ToFhirStoreIamPolicyMapOutputWithContext(ctx context.Context) FhirStoreIamPolicyMapOutput {
	return o
}

func (o FhirStoreIamPolicyMapOutput) MapIndex(k pulumi.StringInput) FhirStoreIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FhirStoreIamPolicy {
		return vs[0].(map[string]FhirStoreIamPolicy)[vs[1].(string)]
	}).(FhirStoreIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(FhirStoreIamPolicyOutput{})
	pulumi.RegisterOutputType(FhirStoreIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(FhirStoreIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(FhirStoreIamPolicyMapOutput{})
}

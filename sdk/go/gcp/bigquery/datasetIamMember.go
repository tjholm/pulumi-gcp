// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for BigQuery dataset. Each of these resources serves a different use case:
//
// * `bigquery.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
// * `bigquery.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
// * `bigquery.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.
//
// These resources are intended to convert the permissions system for BigQuery datasets to the standard IAM interface. For advanced usages, including [creating authorized views](https://cloud.google.com/bigquery/docs/share-access-views), please use either `bigquery.DatasetAccess` or the `access` field on `bigquery.Dataset`.
//
// > **Note:** These resources **cannot** be used with `bigquery.DatasetAccess` resources or the `access` field on `bigquery.Dataset` or they will fight over what the policy should be.
//
// > **Note:** Using any of these resources will remove any authorized view permissions from the dataset. To assign and preserve authorized view permissions use the `bigquery.DatasetAccess` instead.
//
// > **Note:** Legacy BigQuery roles `OWNER` `WRITER` and `READER` **cannot** be used with any of these IAM resources. Instead use the full role form of: `roles/bigquery.dataOwner` `roles/bigquery.dataEditor` and `roles/bigquery.dataViewer`.
//
// > **Note:** `bigquery.DatasetIamPolicy` **cannot** be used in conjunction with `bigquery.DatasetIamBinding` and `bigquery.DatasetIamMember` or they will fight over what your policy should be.
//
// > **Note:** `bigquery.DatasetIamBinding` resources **can be** used in conjunction with `bigquery.DatasetIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_bigquery\_dataset\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			owner, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					organizations.GetIAMPolicyBinding{
//						Role: "roles/bigquery.dataOwner",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			datasetDataset, err := bigquery.NewDataset(ctx, "datasetDataset", &bigquery.DatasetArgs{
//				DatasetId: pulumi.String("example_dataset"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewDatasetIamPolicy(ctx, "datasetDatasetIamPolicy", &bigquery.DatasetIamPolicyArgs{
//				DatasetId:  datasetDataset.DatasetId,
//				PolicyData: pulumi.String(owner.PolicyData),
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
// ## google\_bigquery\_dataset\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			dataset, err := bigquery.NewDataset(ctx, "dataset", &bigquery.DatasetArgs{
//				DatasetId: pulumi.String("example_dataset"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewDatasetIamBinding(ctx, "reader", &bigquery.DatasetIamBindingArgs{
//				DatasetId: dataset.DatasetId,
//				Role:      pulumi.String("roles/bigquery.dataViewer"),
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
// ## google\_bigquery\_dataset\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			dataset, err := bigquery.NewDataset(ctx, "dataset", &bigquery.DatasetArgs{
//				DatasetId: pulumi.String("example_dataset"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewDatasetIamMember(ctx, "editor", &bigquery.DatasetIamMemberArgs{
//				DatasetId: dataset.DatasetId,
//				Role:      pulumi.String("roles/bigquery.dataEditor"),
//				Member:    pulumi.String("user:jane@example.com"),
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
// This member resource can be imported using the `dataset_id`, role, and account e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigquery/datasetIamMember:DatasetIamMember dataset_iam "projects/your-project-id/datasets/dataset-id roles/viewer user:foo@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `dataset_id` and role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigquery/datasetIamMember:DatasetIamMember dataset_iam "projects/your-project-id/datasets/dataset-id roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `dataset_id`, role, and account e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigquery/datasetIamMember:DatasetIamMember dataset_iam projects/your-project-id/datasets/dataset-id
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DatasetIamMember struct {
	pulumi.CustomResourceState

	Condition DatasetIamMemberConditionPtrOutput `pulumi:"condition"`
	// The dataset ID.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDatasetIamMember registers a new resource with the given unique name, arguments, and options.
func NewDatasetIamMember(ctx *pulumi.Context,
	name string, args *DatasetIamMemberArgs, opts ...pulumi.ResourceOption) (*DatasetIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource DatasetIamMember
	err := ctx.RegisterResource("gcp:bigquery/datasetIamMember:DatasetIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetIamMember gets an existing DatasetIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetIamMemberState, opts ...pulumi.ResourceOption) (*DatasetIamMember, error) {
	var resource DatasetIamMember
	err := ctx.ReadResource("gcp:bigquery/datasetIamMember:DatasetIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetIamMember resources.
type datasetIamMemberState struct {
	Condition *DatasetIamMemberCondition `pulumi:"condition"`
	// The dataset ID.
	DatasetId *string `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DatasetIamMemberState struct {
	Condition DatasetIamMemberConditionPtrInput
	// The dataset ID.
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the dataset's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DatasetIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamMemberState)(nil)).Elem()
}

type datasetIamMemberArgs struct {
	Condition *DatasetIamMemberCondition `pulumi:"condition"`
	// The dataset ID.
	DatasetId string `pulumi:"datasetId"`
	Member    string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DatasetIamMember resource.
type DatasetIamMemberArgs struct {
	Condition DatasetIamMemberConditionPtrInput
	// The dataset ID.
	DatasetId pulumi.StringInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DatasetIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamMemberArgs)(nil)).Elem()
}

type DatasetIamMemberInput interface {
	pulumi.Input

	ToDatasetIamMemberOutput() DatasetIamMemberOutput
	ToDatasetIamMemberOutputWithContext(ctx context.Context) DatasetIamMemberOutput
}

func (*DatasetIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetIamMember)(nil)).Elem()
}

func (i *DatasetIamMember) ToDatasetIamMemberOutput() DatasetIamMemberOutput {
	return i.ToDatasetIamMemberOutputWithContext(context.Background())
}

func (i *DatasetIamMember) ToDatasetIamMemberOutputWithContext(ctx context.Context) DatasetIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamMemberOutput)
}

// DatasetIamMemberArrayInput is an input type that accepts DatasetIamMemberArray and DatasetIamMemberArrayOutput values.
// You can construct a concrete instance of `DatasetIamMemberArrayInput` via:
//
//	DatasetIamMemberArray{ DatasetIamMemberArgs{...} }
type DatasetIamMemberArrayInput interface {
	pulumi.Input

	ToDatasetIamMemberArrayOutput() DatasetIamMemberArrayOutput
	ToDatasetIamMemberArrayOutputWithContext(context.Context) DatasetIamMemberArrayOutput
}

type DatasetIamMemberArray []DatasetIamMemberInput

func (DatasetIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetIamMember)(nil)).Elem()
}

func (i DatasetIamMemberArray) ToDatasetIamMemberArrayOutput() DatasetIamMemberArrayOutput {
	return i.ToDatasetIamMemberArrayOutputWithContext(context.Background())
}

func (i DatasetIamMemberArray) ToDatasetIamMemberArrayOutputWithContext(ctx context.Context) DatasetIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamMemberArrayOutput)
}

// DatasetIamMemberMapInput is an input type that accepts DatasetIamMemberMap and DatasetIamMemberMapOutput values.
// You can construct a concrete instance of `DatasetIamMemberMapInput` via:
//
//	DatasetIamMemberMap{ "key": DatasetIamMemberArgs{...} }
type DatasetIamMemberMapInput interface {
	pulumi.Input

	ToDatasetIamMemberMapOutput() DatasetIamMemberMapOutput
	ToDatasetIamMemberMapOutputWithContext(context.Context) DatasetIamMemberMapOutput
}

type DatasetIamMemberMap map[string]DatasetIamMemberInput

func (DatasetIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetIamMember)(nil)).Elem()
}

func (i DatasetIamMemberMap) ToDatasetIamMemberMapOutput() DatasetIamMemberMapOutput {
	return i.ToDatasetIamMemberMapOutputWithContext(context.Background())
}

func (i DatasetIamMemberMap) ToDatasetIamMemberMapOutputWithContext(ctx context.Context) DatasetIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamMemberMapOutput)
}

type DatasetIamMemberOutput struct{ *pulumi.OutputState }

func (DatasetIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetIamMember)(nil)).Elem()
}

func (o DatasetIamMemberOutput) ToDatasetIamMemberOutput() DatasetIamMemberOutput {
	return o
}

func (o DatasetIamMemberOutput) ToDatasetIamMemberOutputWithContext(ctx context.Context) DatasetIamMemberOutput {
	return o
}

func (o DatasetIamMemberOutput) Condition() DatasetIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *DatasetIamMember) DatasetIamMemberConditionPtrOutput { return v.Condition }).(DatasetIamMemberConditionPtrOutput)
}

// The dataset ID.
func (o DatasetIamMemberOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamMember) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

// (Computed) The etag of the dataset's IAM policy.
func (o DatasetIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DatasetIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o DatasetIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o DatasetIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type DatasetIamMemberArrayOutput struct{ *pulumi.OutputState }

func (DatasetIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetIamMember)(nil)).Elem()
}

func (o DatasetIamMemberArrayOutput) ToDatasetIamMemberArrayOutput() DatasetIamMemberArrayOutput {
	return o
}

func (o DatasetIamMemberArrayOutput) ToDatasetIamMemberArrayOutputWithContext(ctx context.Context) DatasetIamMemberArrayOutput {
	return o
}

func (o DatasetIamMemberArrayOutput) Index(i pulumi.IntInput) DatasetIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatasetIamMember {
		return vs[0].([]*DatasetIamMember)[vs[1].(int)]
	}).(DatasetIamMemberOutput)
}

type DatasetIamMemberMapOutput struct{ *pulumi.OutputState }

func (DatasetIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetIamMember)(nil)).Elem()
}

func (o DatasetIamMemberMapOutput) ToDatasetIamMemberMapOutput() DatasetIamMemberMapOutput {
	return o
}

func (o DatasetIamMemberMapOutput) ToDatasetIamMemberMapOutputWithContext(ctx context.Context) DatasetIamMemberMapOutput {
	return o
}

func (o DatasetIamMemberMapOutput) MapIndex(k pulumi.StringInput) DatasetIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatasetIamMember {
		return vs[0].(map[string]*DatasetIamMember)[vs[1].(string)]
	}).(DatasetIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetIamMemberInput)(nil)).Elem(), &DatasetIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetIamMemberArrayInput)(nil)).Elem(), DatasetIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetIamMemberMapInput)(nil)).Elem(), DatasetIamMemberMap{})
	pulumi.RegisterOutputType(DatasetIamMemberOutput{})
	pulumi.RegisterOutputType(DatasetIamMemberArrayOutput{})
	pulumi.RegisterOutputType(DatasetIamMemberMapOutput{})
}

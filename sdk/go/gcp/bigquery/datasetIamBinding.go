// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
//					{
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
//				PolicyData: *pulumi.String(owner.PolicyData),
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
//	$ pulumi import gcp:bigquery/datasetIamBinding:DatasetIamBinding dataset_iam "projects/your-project-id/datasets/dataset-id roles/viewer user:foo@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `dataset_id` and role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigquery/datasetIamBinding:DatasetIamBinding dataset_iam "projects/your-project-id/datasets/dataset-id roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `dataset_id`, role, and account e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigquery/datasetIamBinding:DatasetIamBinding dataset_iam projects/your-project-id/datasets/dataset-id
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DatasetIamBinding struct {
	pulumi.CustomResourceState

	Condition DatasetIamBindingConditionPtrOutput `pulumi:"condition"`
	// The dataset ID.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDatasetIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDatasetIamBinding(ctx *pulumi.Context,
	name string, args *DatasetIamBindingArgs, opts ...pulumi.ResourceOption) (*DatasetIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatasetIamBinding
	err := ctx.RegisterResource("gcp:bigquery/datasetIamBinding:DatasetIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetIamBinding gets an existing DatasetIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetIamBindingState, opts ...pulumi.ResourceOption) (*DatasetIamBinding, error) {
	var resource DatasetIamBinding
	err := ctx.ReadResource("gcp:bigquery/datasetIamBinding:DatasetIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetIamBinding resources.
type datasetIamBindingState struct {
	Condition *DatasetIamBindingCondition `pulumi:"condition"`
	// The dataset ID.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	DatasetId *string `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DatasetIamBindingState struct {
	Condition DatasetIamBindingConditionPtrInput
	// The dataset ID.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the dataset's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DatasetIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamBindingState)(nil)).Elem()
}

type datasetIamBindingArgs struct {
	Condition *DatasetIamBindingCondition `pulumi:"condition"`
	// The dataset ID.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	DatasetId string   `pulumi:"datasetId"`
	Members   []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DatasetIamBinding resource.
type DatasetIamBindingArgs struct {
	Condition DatasetIamBindingConditionPtrInput
	// The dataset ID.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	DatasetId pulumi.StringInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DatasetIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamBindingArgs)(nil)).Elem()
}

type DatasetIamBindingInput interface {
	pulumi.Input

	ToDatasetIamBindingOutput() DatasetIamBindingOutput
	ToDatasetIamBindingOutputWithContext(ctx context.Context) DatasetIamBindingOutput
}

func (*DatasetIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetIamBinding)(nil)).Elem()
}

func (i *DatasetIamBinding) ToDatasetIamBindingOutput() DatasetIamBindingOutput {
	return i.ToDatasetIamBindingOutputWithContext(context.Background())
}

func (i *DatasetIamBinding) ToDatasetIamBindingOutputWithContext(ctx context.Context) DatasetIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamBindingOutput)
}

func (i *DatasetIamBinding) ToOutput(ctx context.Context) pulumix.Output[*DatasetIamBinding] {
	return pulumix.Output[*DatasetIamBinding]{
		OutputState: i.ToDatasetIamBindingOutputWithContext(ctx).OutputState,
	}
}

// DatasetIamBindingArrayInput is an input type that accepts DatasetIamBindingArray and DatasetIamBindingArrayOutput values.
// You can construct a concrete instance of `DatasetIamBindingArrayInput` via:
//
//	DatasetIamBindingArray{ DatasetIamBindingArgs{...} }
type DatasetIamBindingArrayInput interface {
	pulumi.Input

	ToDatasetIamBindingArrayOutput() DatasetIamBindingArrayOutput
	ToDatasetIamBindingArrayOutputWithContext(context.Context) DatasetIamBindingArrayOutput
}

type DatasetIamBindingArray []DatasetIamBindingInput

func (DatasetIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetIamBinding)(nil)).Elem()
}

func (i DatasetIamBindingArray) ToDatasetIamBindingArrayOutput() DatasetIamBindingArrayOutput {
	return i.ToDatasetIamBindingArrayOutputWithContext(context.Background())
}

func (i DatasetIamBindingArray) ToDatasetIamBindingArrayOutputWithContext(ctx context.Context) DatasetIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamBindingArrayOutput)
}

func (i DatasetIamBindingArray) ToOutput(ctx context.Context) pulumix.Output[[]*DatasetIamBinding] {
	return pulumix.Output[[]*DatasetIamBinding]{
		OutputState: i.ToDatasetIamBindingArrayOutputWithContext(ctx).OutputState,
	}
}

// DatasetIamBindingMapInput is an input type that accepts DatasetIamBindingMap and DatasetIamBindingMapOutput values.
// You can construct a concrete instance of `DatasetIamBindingMapInput` via:
//
//	DatasetIamBindingMap{ "key": DatasetIamBindingArgs{...} }
type DatasetIamBindingMapInput interface {
	pulumi.Input

	ToDatasetIamBindingMapOutput() DatasetIamBindingMapOutput
	ToDatasetIamBindingMapOutputWithContext(context.Context) DatasetIamBindingMapOutput
}

type DatasetIamBindingMap map[string]DatasetIamBindingInput

func (DatasetIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetIamBinding)(nil)).Elem()
}

func (i DatasetIamBindingMap) ToDatasetIamBindingMapOutput() DatasetIamBindingMapOutput {
	return i.ToDatasetIamBindingMapOutputWithContext(context.Background())
}

func (i DatasetIamBindingMap) ToDatasetIamBindingMapOutputWithContext(ctx context.Context) DatasetIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamBindingMapOutput)
}

func (i DatasetIamBindingMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DatasetIamBinding] {
	return pulumix.Output[map[string]*DatasetIamBinding]{
		OutputState: i.ToDatasetIamBindingMapOutputWithContext(ctx).OutputState,
	}
}

type DatasetIamBindingOutput struct{ *pulumi.OutputState }

func (DatasetIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetIamBinding)(nil)).Elem()
}

func (o DatasetIamBindingOutput) ToDatasetIamBindingOutput() DatasetIamBindingOutput {
	return o
}

func (o DatasetIamBindingOutput) ToDatasetIamBindingOutputWithContext(ctx context.Context) DatasetIamBindingOutput {
	return o
}

func (o DatasetIamBindingOutput) ToOutput(ctx context.Context) pulumix.Output[*DatasetIamBinding] {
	return pulumix.Output[*DatasetIamBinding]{
		OutputState: o.OutputState,
	}
}

func (o DatasetIamBindingOutput) Condition() DatasetIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *DatasetIamBinding) DatasetIamBindingConditionPtrOutput { return v.Condition }).(DatasetIamBindingConditionPtrOutput)
}

// The dataset ID.
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o DatasetIamBindingOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamBinding) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

// (Computed) The etag of the dataset's IAM policy.
func (o DatasetIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DatasetIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatasetIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o DatasetIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o DatasetIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type DatasetIamBindingArrayOutput struct{ *pulumi.OutputState }

func (DatasetIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetIamBinding)(nil)).Elem()
}

func (o DatasetIamBindingArrayOutput) ToDatasetIamBindingArrayOutput() DatasetIamBindingArrayOutput {
	return o
}

func (o DatasetIamBindingArrayOutput) ToDatasetIamBindingArrayOutputWithContext(ctx context.Context) DatasetIamBindingArrayOutput {
	return o
}

func (o DatasetIamBindingArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DatasetIamBinding] {
	return pulumix.Output[[]*DatasetIamBinding]{
		OutputState: o.OutputState,
	}
}

func (o DatasetIamBindingArrayOutput) Index(i pulumi.IntInput) DatasetIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatasetIamBinding {
		return vs[0].([]*DatasetIamBinding)[vs[1].(int)]
	}).(DatasetIamBindingOutput)
}

type DatasetIamBindingMapOutput struct{ *pulumi.OutputState }

func (DatasetIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetIamBinding)(nil)).Elem()
}

func (o DatasetIamBindingMapOutput) ToDatasetIamBindingMapOutput() DatasetIamBindingMapOutput {
	return o
}

func (o DatasetIamBindingMapOutput) ToDatasetIamBindingMapOutputWithContext(ctx context.Context) DatasetIamBindingMapOutput {
	return o
}

func (o DatasetIamBindingMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DatasetIamBinding] {
	return pulumix.Output[map[string]*DatasetIamBinding]{
		OutputState: o.OutputState,
	}
}

func (o DatasetIamBindingMapOutput) MapIndex(k pulumi.StringInput) DatasetIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatasetIamBinding {
		return vs[0].(map[string]*DatasetIamBinding)[vs[1].(string)]
	}).(DatasetIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetIamBindingInput)(nil)).Elem(), &DatasetIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetIamBindingArrayInput)(nil)).Elem(), DatasetIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetIamBindingMapInput)(nil)).Elem(), DatasetIamBindingMap{})
	pulumi.RegisterOutputType(DatasetIamBindingOutput{})
	pulumi.RegisterOutputType(DatasetIamBindingArrayOutput{})
	pulumi.RegisterOutputType(DatasetIamBindingMapOutput{})
}

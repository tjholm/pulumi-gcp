// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquery"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquery"
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquery"
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
// ### Importing IAM policies IAM policy imports use the identifier of the BigQuery Dataset resource. For example* `projects/{{project_id}}/datasets/{{dataset_id}}` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {
//
//	id = projects/{{project_id}}/datasets/{{dataset_id}}
//
//	to = google_bigquery_dataset_iam_policy.default } The `pulumi import` command can also be used
//
// ```sh
//
//	$ pulumi import gcp:bigquery/datasetIamPolicy:DatasetIamPolicy default projects/{{project_id}}/datasets/{{dataset_id}}
//
// ```
type DatasetIamPolicy struct {
	pulumi.CustomResourceState

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
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewDatasetIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatasetIamPolicy(ctx *pulumi.Context,
	name string, args *DatasetIamPolicyArgs, opts ...pulumi.ResourceOption) (*DatasetIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatasetIamPolicy
	err := ctx.RegisterResource("gcp:bigquery/datasetIamPolicy:DatasetIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetIamPolicy gets an existing DatasetIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetIamPolicyState, opts ...pulumi.ResourceOption) (*DatasetIamPolicy, error) {
	var resource DatasetIamPolicy
	err := ctx.ReadResource("gcp:bigquery/datasetIamPolicy:DatasetIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetIamPolicy resources.
type datasetIamPolicyState struct {
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
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type DatasetIamPolicyState struct {
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
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DatasetIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamPolicyState)(nil)).Elem()
}

type datasetIamPolicyArgs struct {
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
	DatasetId string `pulumi:"datasetId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a DatasetIamPolicy resource.
type DatasetIamPolicyArgs struct {
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
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DatasetIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamPolicyArgs)(nil)).Elem()
}

type DatasetIamPolicyInput interface {
	pulumi.Input

	ToDatasetIamPolicyOutput() DatasetIamPolicyOutput
	ToDatasetIamPolicyOutputWithContext(ctx context.Context) DatasetIamPolicyOutput
}

func (*DatasetIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetIamPolicy)(nil)).Elem()
}

func (i *DatasetIamPolicy) ToDatasetIamPolicyOutput() DatasetIamPolicyOutput {
	return i.ToDatasetIamPolicyOutputWithContext(context.Background())
}

func (i *DatasetIamPolicy) ToDatasetIamPolicyOutputWithContext(ctx context.Context) DatasetIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamPolicyOutput)
}

// DatasetIamPolicyArrayInput is an input type that accepts DatasetIamPolicyArray and DatasetIamPolicyArrayOutput values.
// You can construct a concrete instance of `DatasetIamPolicyArrayInput` via:
//
//	DatasetIamPolicyArray{ DatasetIamPolicyArgs{...} }
type DatasetIamPolicyArrayInput interface {
	pulumi.Input

	ToDatasetIamPolicyArrayOutput() DatasetIamPolicyArrayOutput
	ToDatasetIamPolicyArrayOutputWithContext(context.Context) DatasetIamPolicyArrayOutput
}

type DatasetIamPolicyArray []DatasetIamPolicyInput

func (DatasetIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetIamPolicy)(nil)).Elem()
}

func (i DatasetIamPolicyArray) ToDatasetIamPolicyArrayOutput() DatasetIamPolicyArrayOutput {
	return i.ToDatasetIamPolicyArrayOutputWithContext(context.Background())
}

func (i DatasetIamPolicyArray) ToDatasetIamPolicyArrayOutputWithContext(ctx context.Context) DatasetIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamPolicyArrayOutput)
}

// DatasetIamPolicyMapInput is an input type that accepts DatasetIamPolicyMap and DatasetIamPolicyMapOutput values.
// You can construct a concrete instance of `DatasetIamPolicyMapInput` via:
//
//	DatasetIamPolicyMap{ "key": DatasetIamPolicyArgs{...} }
type DatasetIamPolicyMapInput interface {
	pulumi.Input

	ToDatasetIamPolicyMapOutput() DatasetIamPolicyMapOutput
	ToDatasetIamPolicyMapOutputWithContext(context.Context) DatasetIamPolicyMapOutput
}

type DatasetIamPolicyMap map[string]DatasetIamPolicyInput

func (DatasetIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetIamPolicy)(nil)).Elem()
}

func (i DatasetIamPolicyMap) ToDatasetIamPolicyMapOutput() DatasetIamPolicyMapOutput {
	return i.ToDatasetIamPolicyMapOutputWithContext(context.Background())
}

func (i DatasetIamPolicyMap) ToDatasetIamPolicyMapOutputWithContext(ctx context.Context) DatasetIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamPolicyMapOutput)
}

type DatasetIamPolicyOutput struct{ *pulumi.OutputState }

func (DatasetIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetIamPolicy)(nil)).Elem()
}

func (o DatasetIamPolicyOutput) ToDatasetIamPolicyOutput() DatasetIamPolicyOutput {
	return o
}

func (o DatasetIamPolicyOutput) ToDatasetIamPolicyOutputWithContext(ctx context.Context) DatasetIamPolicyOutput {
	return o
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
func (o DatasetIamPolicyOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamPolicy) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

// (Computed) The etag of the dataset's IAM policy.
func (o DatasetIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o DatasetIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o DatasetIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DatasetIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type DatasetIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (DatasetIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetIamPolicy)(nil)).Elem()
}

func (o DatasetIamPolicyArrayOutput) ToDatasetIamPolicyArrayOutput() DatasetIamPolicyArrayOutput {
	return o
}

func (o DatasetIamPolicyArrayOutput) ToDatasetIamPolicyArrayOutputWithContext(ctx context.Context) DatasetIamPolicyArrayOutput {
	return o
}

func (o DatasetIamPolicyArrayOutput) Index(i pulumi.IntInput) DatasetIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatasetIamPolicy {
		return vs[0].([]*DatasetIamPolicy)[vs[1].(int)]
	}).(DatasetIamPolicyOutput)
}

type DatasetIamPolicyMapOutput struct{ *pulumi.OutputState }

func (DatasetIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetIamPolicy)(nil)).Elem()
}

func (o DatasetIamPolicyMapOutput) ToDatasetIamPolicyMapOutput() DatasetIamPolicyMapOutput {
	return o
}

func (o DatasetIamPolicyMapOutput) ToDatasetIamPolicyMapOutputWithContext(ctx context.Context) DatasetIamPolicyMapOutput {
	return o
}

func (o DatasetIamPolicyMapOutput) MapIndex(k pulumi.StringInput) DatasetIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatasetIamPolicy {
		return vs[0].(map[string]*DatasetIamPolicy)[vs[1].(string)]
	}).(DatasetIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetIamPolicyInput)(nil)).Elem(), &DatasetIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetIamPolicyArrayInput)(nil)).Elem(), DatasetIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetIamPolicyMapInput)(nil)).Elem(), DatasetIamPolicyMap{})
	pulumi.RegisterOutputType(DatasetIamPolicyOutput{})
	pulumi.RegisterOutputType(DatasetIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(DatasetIamPolicyMapOutput{})
}

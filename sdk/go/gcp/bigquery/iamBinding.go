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

// Three different resources help you manage your IAM policy for BigQuery Table. Each of these resources serves a different use case:
//
// * `bigquery.IamPolicy`: Authoritative. Sets the IAM policy for the table and replaces any existing policy already attached.
// * `bigquery.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
// * `bigquery.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `bigquery.IamPolicy`: Retrieves the IAM policy for the table
//
// > **Note:** `bigquery.IamPolicy` **cannot** be used in conjunction with `bigquery.IamBinding` and `bigquery.IamMember` or they will fight over what your policy should be.
//
// > **Note:** `bigquery.IamBinding` resources **can be** used in conjunction with `bigquery.IamMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## google\_bigquery\_table\_iam\_policy
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
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
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
//			_, err = bigquery.NewIamPolicy(ctx, "policy", &bigquery.IamPolicyArgs{
//				Project:    pulumi.Any(google_bigquery_table.Test.Project),
//				DatasetId:  pulumi.Any(google_bigquery_table.Test.Dataset_id),
//				TableId:    pulumi.Any(google_bigquery_table.Test.Table_id),
//				PolicyData: *pulumi.String(admin.PolicyData),
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
// With IAM Conditions:
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
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/bigquery.dataOwner",
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Condition: {
//							Title:       "expires_after_2019_12_31",
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewIamPolicy(ctx, "policy", &bigquery.IamPolicyArgs{
//				Project:    pulumi.Any(google_bigquery_table.Test.Project),
//				DatasetId:  pulumi.Any(google_bigquery_table.Test.Dataset_id),
//				TableId:    pulumi.Any(google_bigquery_table.Test.Table_id),
//				PolicyData: *pulumi.String(admin.PolicyData),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## google\_bigquery\_table\_iam\_binding
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
//			_, err := bigquery.NewIamBinding(ctx, "binding", &bigquery.IamBindingArgs{
//				Project:   pulumi.Any(google_bigquery_table.Test.Project),
//				DatasetId: pulumi.Any(google_bigquery_table.Test.Dataset_id),
//				TableId:   pulumi.Any(google_bigquery_table.Test.Table_id),
//				Role:      pulumi.String("roles/bigquery.dataOwner"),
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
// With IAM Conditions:
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
//			_, err := bigquery.NewIamBinding(ctx, "binding", &bigquery.IamBindingArgs{
//				Project:   pulumi.Any(google_bigquery_table.Test.Project),
//				DatasetId: pulumi.Any(google_bigquery_table.Test.Dataset_id),
//				TableId:   pulumi.Any(google_bigquery_table.Test.Table_id),
//				Role:      pulumi.String("roles/bigquery.dataOwner"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &bigquery.IamBindingConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## google\_bigquery\_table\_iam\_member
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
//			_, err := bigquery.NewIamMember(ctx, "member", &bigquery.IamMemberArgs{
//				Project:   pulumi.Any(google_bigquery_table.Test.Project),
//				DatasetId: pulumi.Any(google_bigquery_table.Test.Dataset_id),
//				TableId:   pulumi.Any(google_bigquery_table.Test.Table_id),
//				Role:      pulumi.String("roles/bigquery.dataOwner"),
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
// With IAM Conditions:
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
//			_, err := bigquery.NewIamMember(ctx, "member", &bigquery.IamMemberArgs{
//				Project:   pulumi.Any(google_bigquery_table.Test.Project),
//				DatasetId: pulumi.Any(google_bigquery_table.Test.Dataset_id),
//				TableId:   pulumi.Any(google_bigquery_table.Test.Table_id),
//				Role:      pulumi.String("roles/bigquery.dataOwner"),
//				Member:    pulumi.String("user:jane@example.com"),
//				Condition: &bigquery.IamMemberConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} * {{project}}/{{dataset_id}}/{{table_id}} * {{dataset_id}}/{{table_id}} * {{table_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery table IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigquery/iamBinding:IamBinding editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigquery/iamBinding:IamBinding editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigquery/iamBinding:IamBinding editor projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type IamBinding struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IamBindingConditionPtrOutput `pulumi:"condition"`
	DatasetId pulumi.StringOutput          `pulumi:"datasetId"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    pulumi.StringOutput `pulumi:"role"`
	TableId pulumi.StringOutput `pulumi:"tableId"`
}

// NewIamBinding registers a new resource with the given unique name, arguments, and options.
func NewIamBinding(ctx *pulumi.Context,
	name string, args *IamBindingArgs, opts ...pulumi.ResourceOption) (*IamBinding, error) {
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
	if args.TableId == nil {
		return nil, errors.New("invalid value for required argument 'TableId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamBinding
	err := ctx.RegisterResource("gcp:bigquery/iamBinding:IamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamBinding gets an existing IamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamBindingState, opts ...pulumi.ResourceOption) (*IamBinding, error) {
	var resource IamBinding
	err := ctx.ReadResource("gcp:bigquery/iamBinding:IamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamBinding resources.
type iamBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IamBindingCondition `pulumi:"condition"`
	DatasetId *string              `pulumi:"datasetId"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    *string `pulumi:"role"`
	TableId *string `pulumi:"tableId"`
}

type IamBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IamBindingConditionPtrInput
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    pulumi.StringPtrInput
	TableId pulumi.StringPtrInput
}

func (IamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamBindingState)(nil)).Elem()
}

type iamBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IamBindingCondition `pulumi:"condition"`
	DatasetId string               `pulumi:"datasetId"`
	Members   []string             `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    string `pulumi:"role"`
	TableId string `pulumi:"tableId"`
}

// The set of arguments for constructing a IamBinding resource.
type IamBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IamBindingConditionPtrInput
	DatasetId pulumi.StringInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    pulumi.StringInput
	TableId pulumi.StringInput
}

func (IamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamBindingArgs)(nil)).Elem()
}

type IamBindingInput interface {
	pulumi.Input

	ToIamBindingOutput() IamBindingOutput
	ToIamBindingOutputWithContext(ctx context.Context) IamBindingOutput
}

func (*IamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**IamBinding)(nil)).Elem()
}

func (i *IamBinding) ToIamBindingOutput() IamBindingOutput {
	return i.ToIamBindingOutputWithContext(context.Background())
}

func (i *IamBinding) ToIamBindingOutputWithContext(ctx context.Context) IamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamBindingOutput)
}

func (i *IamBinding) ToOutput(ctx context.Context) pulumix.Output[*IamBinding] {
	return pulumix.Output[*IamBinding]{
		OutputState: i.ToIamBindingOutputWithContext(ctx).OutputState,
	}
}

// IamBindingArrayInput is an input type that accepts IamBindingArray and IamBindingArrayOutput values.
// You can construct a concrete instance of `IamBindingArrayInput` via:
//
//	IamBindingArray{ IamBindingArgs{...} }
type IamBindingArrayInput interface {
	pulumi.Input

	ToIamBindingArrayOutput() IamBindingArrayOutput
	ToIamBindingArrayOutputWithContext(context.Context) IamBindingArrayOutput
}

type IamBindingArray []IamBindingInput

func (IamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamBinding)(nil)).Elem()
}

func (i IamBindingArray) ToIamBindingArrayOutput() IamBindingArrayOutput {
	return i.ToIamBindingArrayOutputWithContext(context.Background())
}

func (i IamBindingArray) ToIamBindingArrayOutputWithContext(ctx context.Context) IamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamBindingArrayOutput)
}

func (i IamBindingArray) ToOutput(ctx context.Context) pulumix.Output[[]*IamBinding] {
	return pulumix.Output[[]*IamBinding]{
		OutputState: i.ToIamBindingArrayOutputWithContext(ctx).OutputState,
	}
}

// IamBindingMapInput is an input type that accepts IamBindingMap and IamBindingMapOutput values.
// You can construct a concrete instance of `IamBindingMapInput` via:
//
//	IamBindingMap{ "key": IamBindingArgs{...} }
type IamBindingMapInput interface {
	pulumi.Input

	ToIamBindingMapOutput() IamBindingMapOutput
	ToIamBindingMapOutputWithContext(context.Context) IamBindingMapOutput
}

type IamBindingMap map[string]IamBindingInput

func (IamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamBinding)(nil)).Elem()
}

func (i IamBindingMap) ToIamBindingMapOutput() IamBindingMapOutput {
	return i.ToIamBindingMapOutputWithContext(context.Background())
}

func (i IamBindingMap) ToIamBindingMapOutputWithContext(ctx context.Context) IamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamBindingMapOutput)
}

func (i IamBindingMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IamBinding] {
	return pulumix.Output[map[string]*IamBinding]{
		OutputState: i.ToIamBindingMapOutputWithContext(ctx).OutputState,
	}
}

type IamBindingOutput struct{ *pulumi.OutputState }

func (IamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamBinding)(nil)).Elem()
}

func (o IamBindingOutput) ToIamBindingOutput() IamBindingOutput {
	return o
}

func (o IamBindingOutput) ToIamBindingOutputWithContext(ctx context.Context) IamBindingOutput {
	return o
}

func (o IamBindingOutput) ToOutput(ctx context.Context) pulumix.Output[*IamBinding] {
	return pulumix.Output[*IamBinding]{
		OutputState: o.OutputState,
	}
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o IamBindingOutput) Condition() IamBindingConditionPtrOutput {
	return o.ApplyT(func(v *IamBinding) IamBindingConditionPtrOutput { return v.Condition }).(IamBindingConditionPtrOutput)
}

func (o IamBindingOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamBinding) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o IamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
//   - **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
//   - **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
//   - **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
func (o IamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o IamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o IamBindingOutput) TableId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamBinding) pulumi.StringOutput { return v.TableId }).(pulumi.StringOutput)
}

type IamBindingArrayOutput struct{ *pulumi.OutputState }

func (IamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamBinding)(nil)).Elem()
}

func (o IamBindingArrayOutput) ToIamBindingArrayOutput() IamBindingArrayOutput {
	return o
}

func (o IamBindingArrayOutput) ToIamBindingArrayOutputWithContext(ctx context.Context) IamBindingArrayOutput {
	return o
}

func (o IamBindingArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IamBinding] {
	return pulumix.Output[[]*IamBinding]{
		OutputState: o.OutputState,
	}
}

func (o IamBindingArrayOutput) Index(i pulumi.IntInput) IamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamBinding {
		return vs[0].([]*IamBinding)[vs[1].(int)]
	}).(IamBindingOutput)
}

type IamBindingMapOutput struct{ *pulumi.OutputState }

func (IamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamBinding)(nil)).Elem()
}

func (o IamBindingMapOutput) ToIamBindingMapOutput() IamBindingMapOutput {
	return o
}

func (o IamBindingMapOutput) ToIamBindingMapOutputWithContext(ctx context.Context) IamBindingMapOutput {
	return o
}

func (o IamBindingMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IamBinding] {
	return pulumix.Output[map[string]*IamBinding]{
		OutputState: o.OutputState,
	}
}

func (o IamBindingMapOutput) MapIndex(k pulumi.StringInput) IamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamBinding {
		return vs[0].(map[string]*IamBinding)[vs[1].(string)]
	}).(IamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamBindingInput)(nil)).Elem(), &IamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamBindingArrayInput)(nil)).Elem(), IamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamBindingMapInput)(nil)).Elem(), IamBindingMap{})
	pulumi.RegisterOutputType(IamBindingOutput{})
	pulumi.RegisterOutputType(IamBindingArrayOutput{})
	pulumi.RegisterOutputType(IamBindingMapOutput{})
}

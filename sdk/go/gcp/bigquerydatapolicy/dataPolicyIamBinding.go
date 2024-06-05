// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigquerydatapolicy

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for BigQuery Data Policy DataPolicy. Each of these resources serves a different use case:
//
// * `bigquerydatapolicy.DataPolicyIamPolicy`: Authoritative. Sets the IAM policy for the datapolicy and replaces any existing policy already attached.
// * `bigquerydatapolicy.DataPolicyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the datapolicy are preserved.
// * `bigquerydatapolicy.DataPolicyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the datapolicy are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `bigquerydatapolicy.DataPolicyIamPolicy`: Retrieves the IAM policy for the datapolicy
//
// > **Note:** `bigquerydatapolicy.DataPolicyIamPolicy` **cannot** be used in conjunction with `bigquerydatapolicy.DataPolicyIamBinding` and `bigquerydatapolicy.DataPolicyIamMember` or they will fight over what your policy should be.
//
// > **Note:** `bigquerydatapolicy.DataPolicyIamBinding` resources **can be** used in conjunction with `bigquerydatapolicy.DataPolicyIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## bigquerydatapolicy.DataPolicyIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
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
//			_, err = bigquerydatapolicy.NewDataPolicyIamPolicy(ctx, "policy", &bigquerydatapolicy.DataPolicyIamPolicyArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				PolicyData:   pulumi.String(admin.PolicyData),
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
// ## bigquerydatapolicy.DataPolicyIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigquerydatapolicy.NewDataPolicyIamBinding(ctx, "binding", &bigquerydatapolicy.DataPolicyIamBindingArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				Role:         pulumi.String("roles/viewer"),
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
// ## bigquerydatapolicy.DataPolicyIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigquerydatapolicy.NewDataPolicyIamMember(ctx, "member", &bigquerydatapolicy.DataPolicyIamMemberArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				Role:         pulumi.String("roles/viewer"),
//				Member:       pulumi.String("user:jane@example.com"),
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
// ## bigquerydatapolicy.DataPolicyIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
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
//			_, err = bigquerydatapolicy.NewDataPolicyIamPolicy(ctx, "policy", &bigquerydatapolicy.DataPolicyIamPolicyArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				PolicyData:   pulumi.String(admin.PolicyData),
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
// ## bigquerydatapolicy.DataPolicyIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigquerydatapolicy.NewDataPolicyIamBinding(ctx, "binding", &bigquerydatapolicy.DataPolicyIamBindingArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				Role:         pulumi.String("roles/viewer"),
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
// ## bigquerydatapolicy.DataPolicyIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigquerydatapolicy.NewDataPolicyIamMember(ctx, "member", &bigquerydatapolicy.DataPolicyIamMemberArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				Role:         pulumi.String("roles/viewer"),
//				Member:       pulumi.String("user:jane@example.com"),
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
// * projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}
//
// * {{project}}/{{location}}/{{data_policy_id}}
//
// * {{location}}/{{data_policy_id}}
//
// * {{data_policy_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// BigQuery Data Policy datapolicy IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor "projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor "projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DataPolicyIamBinding struct {
	pulumi.CustomResourceState

	Condition    DataPolicyIamBindingConditionPtrOutput `pulumi:"condition"`
	DataPolicyId pulumi.StringOutput                    `pulumi:"dataPolicyId"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the location of the data policy.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringOutput `pulumi:"location"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDataPolicyIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDataPolicyIamBinding(ctx *pulumi.Context,
	name string, args *DataPolicyIamBindingArgs, opts ...pulumi.ResourceOption) (*DataPolicyIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'DataPolicyId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataPolicyIamBinding
	err := ctx.RegisterResource("gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataPolicyIamBinding gets an existing DataPolicyIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataPolicyIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataPolicyIamBindingState, opts ...pulumi.ResourceOption) (*DataPolicyIamBinding, error) {
	var resource DataPolicyIamBinding
	err := ctx.ReadResource("gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataPolicyIamBinding resources.
type dataPolicyIamBindingState struct {
	Condition    *DataPolicyIamBindingCondition `pulumi:"condition"`
	DataPolicyId *string                        `pulumi:"dataPolicyId"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the location of the data policy.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DataPolicyIamBindingState struct {
	Condition    DataPolicyIamBindingConditionPtrInput
	DataPolicyId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the location of the data policy.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DataPolicyIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataPolicyIamBindingState)(nil)).Elem()
}

type dataPolicyIamBindingArgs struct {
	Condition    *DataPolicyIamBindingCondition `pulumi:"condition"`
	DataPolicyId string                         `pulumi:"dataPolicyId"`
	// The name of the location of the data policy.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DataPolicyIamBinding resource.
type DataPolicyIamBindingArgs struct {
	Condition    DataPolicyIamBindingConditionPtrInput
	DataPolicyId pulumi.StringInput
	// The name of the location of the data policy.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DataPolicyIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataPolicyIamBindingArgs)(nil)).Elem()
}

type DataPolicyIamBindingInput interface {
	pulumi.Input

	ToDataPolicyIamBindingOutput() DataPolicyIamBindingOutput
	ToDataPolicyIamBindingOutputWithContext(ctx context.Context) DataPolicyIamBindingOutput
}

func (*DataPolicyIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPolicyIamBinding)(nil)).Elem()
}

func (i *DataPolicyIamBinding) ToDataPolicyIamBindingOutput() DataPolicyIamBindingOutput {
	return i.ToDataPolicyIamBindingOutputWithContext(context.Background())
}

func (i *DataPolicyIamBinding) ToDataPolicyIamBindingOutputWithContext(ctx context.Context) DataPolicyIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPolicyIamBindingOutput)
}

// DataPolicyIamBindingArrayInput is an input type that accepts DataPolicyIamBindingArray and DataPolicyIamBindingArrayOutput values.
// You can construct a concrete instance of `DataPolicyIamBindingArrayInput` via:
//
//	DataPolicyIamBindingArray{ DataPolicyIamBindingArgs{...} }
type DataPolicyIamBindingArrayInput interface {
	pulumi.Input

	ToDataPolicyIamBindingArrayOutput() DataPolicyIamBindingArrayOutput
	ToDataPolicyIamBindingArrayOutputWithContext(context.Context) DataPolicyIamBindingArrayOutput
}

type DataPolicyIamBindingArray []DataPolicyIamBindingInput

func (DataPolicyIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataPolicyIamBinding)(nil)).Elem()
}

func (i DataPolicyIamBindingArray) ToDataPolicyIamBindingArrayOutput() DataPolicyIamBindingArrayOutput {
	return i.ToDataPolicyIamBindingArrayOutputWithContext(context.Background())
}

func (i DataPolicyIamBindingArray) ToDataPolicyIamBindingArrayOutputWithContext(ctx context.Context) DataPolicyIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPolicyIamBindingArrayOutput)
}

// DataPolicyIamBindingMapInput is an input type that accepts DataPolicyIamBindingMap and DataPolicyIamBindingMapOutput values.
// You can construct a concrete instance of `DataPolicyIamBindingMapInput` via:
//
//	DataPolicyIamBindingMap{ "key": DataPolicyIamBindingArgs{...} }
type DataPolicyIamBindingMapInput interface {
	pulumi.Input

	ToDataPolicyIamBindingMapOutput() DataPolicyIamBindingMapOutput
	ToDataPolicyIamBindingMapOutputWithContext(context.Context) DataPolicyIamBindingMapOutput
}

type DataPolicyIamBindingMap map[string]DataPolicyIamBindingInput

func (DataPolicyIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataPolicyIamBinding)(nil)).Elem()
}

func (i DataPolicyIamBindingMap) ToDataPolicyIamBindingMapOutput() DataPolicyIamBindingMapOutput {
	return i.ToDataPolicyIamBindingMapOutputWithContext(context.Background())
}

func (i DataPolicyIamBindingMap) ToDataPolicyIamBindingMapOutputWithContext(ctx context.Context) DataPolicyIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPolicyIamBindingMapOutput)
}

type DataPolicyIamBindingOutput struct{ *pulumi.OutputState }

func (DataPolicyIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPolicyIamBinding)(nil)).Elem()
}

func (o DataPolicyIamBindingOutput) ToDataPolicyIamBindingOutput() DataPolicyIamBindingOutput {
	return o
}

func (o DataPolicyIamBindingOutput) ToDataPolicyIamBindingOutputWithContext(ctx context.Context) DataPolicyIamBindingOutput {
	return o
}

func (o DataPolicyIamBindingOutput) Condition() DataPolicyIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *DataPolicyIamBinding) DataPolicyIamBindingConditionPtrOutput { return v.Condition }).(DataPolicyIamBindingConditionPtrOutput)
}

func (o DataPolicyIamBindingOutput) DataPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataPolicyIamBinding) pulumi.StringOutput { return v.DataPolicyId }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o DataPolicyIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataPolicyIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the location of the data policy.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o DataPolicyIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataPolicyIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in `role`.
// Each entry can have one of the following values:
// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
func (o DataPolicyIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataPolicyIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o DataPolicyIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataPolicyIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o DataPolicyIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataPolicyIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type DataPolicyIamBindingArrayOutput struct{ *pulumi.OutputState }

func (DataPolicyIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataPolicyIamBinding)(nil)).Elem()
}

func (o DataPolicyIamBindingArrayOutput) ToDataPolicyIamBindingArrayOutput() DataPolicyIamBindingArrayOutput {
	return o
}

func (o DataPolicyIamBindingArrayOutput) ToDataPolicyIamBindingArrayOutputWithContext(ctx context.Context) DataPolicyIamBindingArrayOutput {
	return o
}

func (o DataPolicyIamBindingArrayOutput) Index(i pulumi.IntInput) DataPolicyIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataPolicyIamBinding {
		return vs[0].([]*DataPolicyIamBinding)[vs[1].(int)]
	}).(DataPolicyIamBindingOutput)
}

type DataPolicyIamBindingMapOutput struct{ *pulumi.OutputState }

func (DataPolicyIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataPolicyIamBinding)(nil)).Elem()
}

func (o DataPolicyIamBindingMapOutput) ToDataPolicyIamBindingMapOutput() DataPolicyIamBindingMapOutput {
	return o
}

func (o DataPolicyIamBindingMapOutput) ToDataPolicyIamBindingMapOutputWithContext(ctx context.Context) DataPolicyIamBindingMapOutput {
	return o
}

func (o DataPolicyIamBindingMapOutput) MapIndex(k pulumi.StringInput) DataPolicyIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataPolicyIamBinding {
		return vs[0].(map[string]*DataPolicyIamBinding)[vs[1].(string)]
	}).(DataPolicyIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataPolicyIamBindingInput)(nil)).Elem(), &DataPolicyIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPolicyIamBindingArrayInput)(nil)).Elem(), DataPolicyIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPolicyIamBindingMapInput)(nil)).Elem(), DataPolicyIamBindingMap{})
	pulumi.RegisterOutputType(DataPolicyIamBindingOutput{})
	pulumi.RegisterOutputType(DataPolicyIamBindingArrayOutput{})
	pulumi.RegisterOutputType(DataPolicyIamBindingMapOutput{})
}

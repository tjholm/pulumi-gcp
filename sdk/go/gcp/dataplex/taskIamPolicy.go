// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataplex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Dataplex Task. Each of these resources serves a different use case:
//
// * `dataplex.TaskIamPolicy`: Authoritative. Sets the IAM policy for the task and replaces any existing policy already attached.
// * `dataplex.TaskIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the task are preserved.
// * `dataplex.TaskIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the task are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dataplex.TaskIamPolicy`: Retrieves the IAM policy for the task
//
// > **Note:** `dataplex.TaskIamPolicy` **cannot** be used in conjunction with `dataplex.TaskIamBinding` and `dataplex.TaskIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dataplex.TaskIamBinding` resources **can be** used in conjunction with `dataplex.TaskIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## dataplex.TaskIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
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
//			_, err = dataplex.NewTaskIamPolicy(ctx, "policy", &dataplex.TaskIamPolicyArgs{
//				Project:    pulumi.Any(example.Project),
//				Location:   pulumi.Any(example.Location),
//				Lake:       pulumi.Any(example.Lake),
//				TaskId:     pulumi.Any(example.TaskId),
//				PolicyData: pulumi.String(admin.PolicyData),
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
// ## dataplex.TaskIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewTaskIamBinding(ctx, "binding", &dataplex.TaskIamBindingArgs{
//				Project:  pulumi.Any(example.Project),
//				Location: pulumi.Any(example.Location),
//				Lake:     pulumi.Any(example.Lake),
//				TaskId:   pulumi.Any(example.TaskId),
//				Role:     pulumi.String("roles/viewer"),
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
// ## dataplex.TaskIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewTaskIamMember(ctx, "member", &dataplex.TaskIamMemberArgs{
//				Project:  pulumi.Any(example.Project),
//				Location: pulumi.Any(example.Location),
//				Lake:     pulumi.Any(example.Lake),
//				TaskId:   pulumi.Any(example.TaskId),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
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
// ## This resource supports User Project Overrides.
//
// -
//
// # IAM policy for Dataplex Task
// Three different resources help you manage your IAM policy for Dataplex Task. Each of these resources serves a different use case:
//
// * `dataplex.TaskIamPolicy`: Authoritative. Sets the IAM policy for the task and replaces any existing policy already attached.
// * `dataplex.TaskIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the task are preserved.
// * `dataplex.TaskIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the task are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dataplex.TaskIamPolicy`: Retrieves the IAM policy for the task
//
// > **Note:** `dataplex.TaskIamPolicy` **cannot** be used in conjunction with `dataplex.TaskIamBinding` and `dataplex.TaskIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dataplex.TaskIamBinding` resources **can be** used in conjunction with `dataplex.TaskIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## dataplex.TaskIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
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
//			_, err = dataplex.NewTaskIamPolicy(ctx, "policy", &dataplex.TaskIamPolicyArgs{
//				Project:    pulumi.Any(example.Project),
//				Location:   pulumi.Any(example.Location),
//				Lake:       pulumi.Any(example.Lake),
//				TaskId:     pulumi.Any(example.TaskId),
//				PolicyData: pulumi.String(admin.PolicyData),
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
// ## dataplex.TaskIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewTaskIamBinding(ctx, "binding", &dataplex.TaskIamBindingArgs{
//				Project:  pulumi.Any(example.Project),
//				Location: pulumi.Any(example.Location),
//				Lake:     pulumi.Any(example.Lake),
//				TaskId:   pulumi.Any(example.TaskId),
//				Role:     pulumi.String("roles/viewer"),
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
// ## dataplex.TaskIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewTaskIamMember(ctx, "member", &dataplex.TaskIamMemberArgs{
//				Project:  pulumi.Any(example.Project),
//				Location: pulumi.Any(example.Location),
//				Lake:     pulumi.Any(example.Lake),
//				TaskId:   pulumi.Any(example.TaskId),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
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
// * projects/{{project}}/locations/{{location}}/lakes/{{lake}}/tasks/{{task_id}}
//
// * {{project}}/{{location}}/{{lake}}/{{task_id}}
//
// * {{location}}/{{lake}}/{{task_id}}
//
// * {{task_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Dataplex task IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/taskIamPolicy:TaskIamPolicy editor "projects/{{project}}/locations/{{location}}/lakes/{{lake}}/tasks/{{task_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/taskIamPolicy:TaskIamPolicy editor "projects/{{project}}/locations/{{location}}/lakes/{{lake}}/tasks/{{task_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/taskIamPolicy:TaskIamPolicy editor projects/{{project}}/locations/{{location}}/lakes/{{lake}}/tasks/{{task_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TaskIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The lake in which the task will be created in.
	// Used to find the parent resource to bind the IAM policy to
	Lake pulumi.StringOutput `pulumi:"lake"`
	// The location in which the task will be created in.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringOutput `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	TaskId  pulumi.StringOutput `pulumi:"taskId"`
}

// NewTaskIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewTaskIamPolicy(ctx *pulumi.Context,
	name string, args *TaskIamPolicyArgs, opts ...pulumi.ResourceOption) (*TaskIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Lake == nil {
		return nil, errors.New("invalid value for required argument 'Lake'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.TaskId == nil {
		return nil, errors.New("invalid value for required argument 'TaskId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TaskIamPolicy
	err := ctx.RegisterResource("gcp:dataplex/taskIamPolicy:TaskIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaskIamPolicy gets an existing TaskIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaskIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskIamPolicyState, opts ...pulumi.ResourceOption) (*TaskIamPolicy, error) {
	var resource TaskIamPolicy
	err := ctx.ReadResource("gcp:dataplex/taskIamPolicy:TaskIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaskIamPolicy resources.
type taskIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The lake in which the task will be created in.
	// Used to find the parent resource to bind the IAM policy to
	Lake *string `pulumi:"lake"`
	// The location in which the task will be created in.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	TaskId  *string `pulumi:"taskId"`
}

type TaskIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The lake in which the task will be created in.
	// Used to find the parent resource to bind the IAM policy to
	Lake pulumi.StringPtrInput
	// The location in which the task will be created in.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	TaskId  pulumi.StringPtrInput
}

func (TaskIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskIamPolicyState)(nil)).Elem()
}

type taskIamPolicyArgs struct {
	// The lake in which the task will be created in.
	// Used to find the parent resource to bind the IAM policy to
	Lake string `pulumi:"lake"`
	// The location in which the task will be created in.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	TaskId  string  `pulumi:"taskId"`
}

// The set of arguments for constructing a TaskIamPolicy resource.
type TaskIamPolicyArgs struct {
	// The lake in which the task will be created in.
	// Used to find the parent resource to bind the IAM policy to
	Lake pulumi.StringInput
	// The location in which the task will be created in.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	TaskId  pulumi.StringInput
}

func (TaskIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskIamPolicyArgs)(nil)).Elem()
}

type TaskIamPolicyInput interface {
	pulumi.Input

	ToTaskIamPolicyOutput() TaskIamPolicyOutput
	ToTaskIamPolicyOutputWithContext(ctx context.Context) TaskIamPolicyOutput
}

func (*TaskIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskIamPolicy)(nil)).Elem()
}

func (i *TaskIamPolicy) ToTaskIamPolicyOutput() TaskIamPolicyOutput {
	return i.ToTaskIamPolicyOutputWithContext(context.Background())
}

func (i *TaskIamPolicy) ToTaskIamPolicyOutputWithContext(ctx context.Context) TaskIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskIamPolicyOutput)
}

// TaskIamPolicyArrayInput is an input type that accepts TaskIamPolicyArray and TaskIamPolicyArrayOutput values.
// You can construct a concrete instance of `TaskIamPolicyArrayInput` via:
//
//	TaskIamPolicyArray{ TaskIamPolicyArgs{...} }
type TaskIamPolicyArrayInput interface {
	pulumi.Input

	ToTaskIamPolicyArrayOutput() TaskIamPolicyArrayOutput
	ToTaskIamPolicyArrayOutputWithContext(context.Context) TaskIamPolicyArrayOutput
}

type TaskIamPolicyArray []TaskIamPolicyInput

func (TaskIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaskIamPolicy)(nil)).Elem()
}

func (i TaskIamPolicyArray) ToTaskIamPolicyArrayOutput() TaskIamPolicyArrayOutput {
	return i.ToTaskIamPolicyArrayOutputWithContext(context.Background())
}

func (i TaskIamPolicyArray) ToTaskIamPolicyArrayOutputWithContext(ctx context.Context) TaskIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskIamPolicyArrayOutput)
}

// TaskIamPolicyMapInput is an input type that accepts TaskIamPolicyMap and TaskIamPolicyMapOutput values.
// You can construct a concrete instance of `TaskIamPolicyMapInput` via:
//
//	TaskIamPolicyMap{ "key": TaskIamPolicyArgs{...} }
type TaskIamPolicyMapInput interface {
	pulumi.Input

	ToTaskIamPolicyMapOutput() TaskIamPolicyMapOutput
	ToTaskIamPolicyMapOutputWithContext(context.Context) TaskIamPolicyMapOutput
}

type TaskIamPolicyMap map[string]TaskIamPolicyInput

func (TaskIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaskIamPolicy)(nil)).Elem()
}

func (i TaskIamPolicyMap) ToTaskIamPolicyMapOutput() TaskIamPolicyMapOutput {
	return i.ToTaskIamPolicyMapOutputWithContext(context.Background())
}

func (i TaskIamPolicyMap) ToTaskIamPolicyMapOutputWithContext(ctx context.Context) TaskIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskIamPolicyMapOutput)
}

type TaskIamPolicyOutput struct{ *pulumi.OutputState }

func (TaskIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskIamPolicy)(nil)).Elem()
}

func (o TaskIamPolicyOutput) ToTaskIamPolicyOutput() TaskIamPolicyOutput {
	return o
}

func (o TaskIamPolicyOutput) ToTaskIamPolicyOutputWithContext(ctx context.Context) TaskIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o TaskIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The lake in which the task will be created in.
// Used to find the parent resource to bind the IAM policy to
func (o TaskIamPolicyOutput) Lake() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskIamPolicy) pulumi.StringOutput { return v.Lake }).(pulumi.StringOutput)
}

// The location in which the task will be created in.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o TaskIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o TaskIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o TaskIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o TaskIamPolicyOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskIamPolicy) pulumi.StringOutput { return v.TaskId }).(pulumi.StringOutput)
}

type TaskIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (TaskIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaskIamPolicy)(nil)).Elem()
}

func (o TaskIamPolicyArrayOutput) ToTaskIamPolicyArrayOutput() TaskIamPolicyArrayOutput {
	return o
}

func (o TaskIamPolicyArrayOutput) ToTaskIamPolicyArrayOutputWithContext(ctx context.Context) TaskIamPolicyArrayOutput {
	return o
}

func (o TaskIamPolicyArrayOutput) Index(i pulumi.IntInput) TaskIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TaskIamPolicy {
		return vs[0].([]*TaskIamPolicy)[vs[1].(int)]
	}).(TaskIamPolicyOutput)
}

type TaskIamPolicyMapOutput struct{ *pulumi.OutputState }

func (TaskIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaskIamPolicy)(nil)).Elem()
}

func (o TaskIamPolicyMapOutput) ToTaskIamPolicyMapOutput() TaskIamPolicyMapOutput {
	return o
}

func (o TaskIamPolicyMapOutput) ToTaskIamPolicyMapOutputWithContext(ctx context.Context) TaskIamPolicyMapOutput {
	return o
}

func (o TaskIamPolicyMapOutput) MapIndex(k pulumi.StringInput) TaskIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TaskIamPolicy {
		return vs[0].(map[string]*TaskIamPolicy)[vs[1].(string)]
	}).(TaskIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaskIamPolicyInput)(nil)).Elem(), &TaskIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskIamPolicyArrayInput)(nil)).Elem(), TaskIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskIamPolicyMapInput)(nil)).Elem(), TaskIamPolicyMap{})
	pulumi.RegisterOutputType(TaskIamPolicyOutput{})
	pulumi.RegisterOutputType(TaskIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(TaskIamPolicyMapOutput{})
}

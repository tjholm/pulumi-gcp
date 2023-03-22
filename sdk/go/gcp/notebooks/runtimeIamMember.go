// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud AI Notebooks Runtime. Each of these resources serves a different use case:
//
// * `notebooks.RuntimeIamPolicy`: Authoritative. Sets the IAM policy for the runtime and replaces any existing policy already attached.
// * `notebooks.RuntimeIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the runtime are preserved.
// * `notebooks.RuntimeIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the runtime are preserved.
//
// > **Note:** `notebooks.RuntimeIamPolicy` **cannot** be used in conjunction with `notebooks.RuntimeIamBinding` and `notebooks.RuntimeIamMember` or they will fight over what your policy should be.
//
// > **Note:** `notebooks.RuntimeIamBinding` resources **can be** used in conjunction with `notebooks.RuntimeIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_notebooks\_runtime\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/notebooks"
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
//			_, err = notebooks.NewRuntimeIamPolicy(ctx, "policy", &notebooks.RuntimeIamPolicyArgs{
//				Project:     pulumi.Any(google_notebooks_runtime.Runtime.Project),
//				Location:    pulumi.Any(google_notebooks_runtime.Runtime.Location),
//				RuntimeName: pulumi.Any(google_notebooks_runtime.Runtime.Name),
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
// ## google\_notebooks\_runtime\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/notebooks"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := notebooks.NewRuntimeIamBinding(ctx, "binding", &notebooks.RuntimeIamBindingArgs{
//				Project:     pulumi.Any(google_notebooks_runtime.Runtime.Project),
//				Location:    pulumi.Any(google_notebooks_runtime.Runtime.Location),
//				RuntimeName: pulumi.Any(google_notebooks_runtime.Runtime.Name),
//				Role:        pulumi.String("roles/viewer"),
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
// ## google\_notebooks\_runtime\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/notebooks"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := notebooks.NewRuntimeIamMember(ctx, "member", &notebooks.RuntimeIamMemberArgs{
//				Project:     pulumi.Any(google_notebooks_runtime.Runtime.Project),
//				Location:    pulumi.Any(google_notebooks_runtime.Runtime.Location),
//				RuntimeName: pulumi.Any(google_notebooks_runtime.Runtime.Name),
//				Role:        pulumi.String("roles/viewer"),
//				Member:      pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/runtimes/{{runtime_name}} * {{project}}/{{location}}/{{runtime_name}} * {{location}}/{{runtime_name}} * {{runtime_name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud AI Notebooks runtime IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:notebooks/runtimeIamMember:RuntimeIamMember editor "projects/{{project}}/locations/{{location}}/runtimes/{{runtime_name}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:notebooks/runtimeIamMember:RuntimeIamMember editor "projects/{{project}}/locations/{{location}}/runtimes/{{runtime_name}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:notebooks/runtimeIamMember:RuntimeIamMember editor projects/{{project}}/locations/{{location}}/runtimes/{{runtime_name}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RuntimeIamMember struct {
	pulumi.CustomResourceState

	Condition RuntimeIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	Member   pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `notebooks.RuntimeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	RuntimeName pulumi.StringOutput `pulumi:"runtimeName"`
}

// NewRuntimeIamMember registers a new resource with the given unique name, arguments, and options.
func NewRuntimeIamMember(ctx *pulumi.Context,
	name string, args *RuntimeIamMemberArgs, opts ...pulumi.ResourceOption) (*RuntimeIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.RuntimeName == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeName'")
	}
	var resource RuntimeIamMember
	err := ctx.RegisterResource("gcp:notebooks/runtimeIamMember:RuntimeIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntimeIamMember gets an existing RuntimeIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntimeIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeIamMemberState, opts ...pulumi.ResourceOption) (*RuntimeIamMember, error) {
	var resource RuntimeIamMember
	err := ctx.ReadResource("gcp:notebooks/runtimeIamMember:RuntimeIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuntimeIamMember resources.
type runtimeIamMemberState struct {
	Condition *RuntimeIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `notebooks.RuntimeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	RuntimeName *string `pulumi:"runtimeName"`
}

type RuntimeIamMemberState struct {
	Condition RuntimeIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `notebooks.RuntimeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	RuntimeName pulumi.StringPtrInput
}

func (RuntimeIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeIamMemberState)(nil)).Elem()
}

type runtimeIamMemberArgs struct {
	Condition *RuntimeIamMemberCondition `pulumi:"condition"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   string  `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `notebooks.RuntimeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	RuntimeName string `pulumi:"runtimeName"`
}

// The set of arguments for constructing a RuntimeIamMember resource.
type RuntimeIamMemberArgs struct {
	Condition RuntimeIamMemberConditionPtrInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `notebooks.RuntimeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	RuntimeName pulumi.StringInput
}

func (RuntimeIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeIamMemberArgs)(nil)).Elem()
}

type RuntimeIamMemberInput interface {
	pulumi.Input

	ToRuntimeIamMemberOutput() RuntimeIamMemberOutput
	ToRuntimeIamMemberOutputWithContext(ctx context.Context) RuntimeIamMemberOutput
}

func (*RuntimeIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeIamMember)(nil)).Elem()
}

func (i *RuntimeIamMember) ToRuntimeIamMemberOutput() RuntimeIamMemberOutput {
	return i.ToRuntimeIamMemberOutputWithContext(context.Background())
}

func (i *RuntimeIamMember) ToRuntimeIamMemberOutputWithContext(ctx context.Context) RuntimeIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeIamMemberOutput)
}

// RuntimeIamMemberArrayInput is an input type that accepts RuntimeIamMemberArray and RuntimeIamMemberArrayOutput values.
// You can construct a concrete instance of `RuntimeIamMemberArrayInput` via:
//
//	RuntimeIamMemberArray{ RuntimeIamMemberArgs{...} }
type RuntimeIamMemberArrayInput interface {
	pulumi.Input

	ToRuntimeIamMemberArrayOutput() RuntimeIamMemberArrayOutput
	ToRuntimeIamMemberArrayOutputWithContext(context.Context) RuntimeIamMemberArrayOutput
}

type RuntimeIamMemberArray []RuntimeIamMemberInput

func (RuntimeIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuntimeIamMember)(nil)).Elem()
}

func (i RuntimeIamMemberArray) ToRuntimeIamMemberArrayOutput() RuntimeIamMemberArrayOutput {
	return i.ToRuntimeIamMemberArrayOutputWithContext(context.Background())
}

func (i RuntimeIamMemberArray) ToRuntimeIamMemberArrayOutputWithContext(ctx context.Context) RuntimeIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeIamMemberArrayOutput)
}

// RuntimeIamMemberMapInput is an input type that accepts RuntimeIamMemberMap and RuntimeIamMemberMapOutput values.
// You can construct a concrete instance of `RuntimeIamMemberMapInput` via:
//
//	RuntimeIamMemberMap{ "key": RuntimeIamMemberArgs{...} }
type RuntimeIamMemberMapInput interface {
	pulumi.Input

	ToRuntimeIamMemberMapOutput() RuntimeIamMemberMapOutput
	ToRuntimeIamMemberMapOutputWithContext(context.Context) RuntimeIamMemberMapOutput
}

type RuntimeIamMemberMap map[string]RuntimeIamMemberInput

func (RuntimeIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuntimeIamMember)(nil)).Elem()
}

func (i RuntimeIamMemberMap) ToRuntimeIamMemberMapOutput() RuntimeIamMemberMapOutput {
	return i.ToRuntimeIamMemberMapOutputWithContext(context.Background())
}

func (i RuntimeIamMemberMap) ToRuntimeIamMemberMapOutputWithContext(ctx context.Context) RuntimeIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeIamMemberMapOutput)
}

type RuntimeIamMemberOutput struct{ *pulumi.OutputState }

func (RuntimeIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeIamMember)(nil)).Elem()
}

func (o RuntimeIamMemberOutput) ToRuntimeIamMemberOutput() RuntimeIamMemberOutput {
	return o
}

func (o RuntimeIamMemberOutput) ToRuntimeIamMemberOutputWithContext(ctx context.Context) RuntimeIamMemberOutput {
	return o
}

func (o RuntimeIamMemberOutput) Condition() RuntimeIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *RuntimeIamMember) RuntimeIamMemberConditionPtrOutput { return v.Condition }).(RuntimeIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o RuntimeIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
func (o RuntimeIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o RuntimeIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o RuntimeIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `notebooks.RuntimeIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o RuntimeIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o RuntimeIamMemberOutput) RuntimeName() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeIamMember) pulumi.StringOutput { return v.RuntimeName }).(pulumi.StringOutput)
}

type RuntimeIamMemberArrayOutput struct{ *pulumi.OutputState }

func (RuntimeIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuntimeIamMember)(nil)).Elem()
}

func (o RuntimeIamMemberArrayOutput) ToRuntimeIamMemberArrayOutput() RuntimeIamMemberArrayOutput {
	return o
}

func (o RuntimeIamMemberArrayOutput) ToRuntimeIamMemberArrayOutputWithContext(ctx context.Context) RuntimeIamMemberArrayOutput {
	return o
}

func (o RuntimeIamMemberArrayOutput) Index(i pulumi.IntInput) RuntimeIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RuntimeIamMember {
		return vs[0].([]*RuntimeIamMember)[vs[1].(int)]
	}).(RuntimeIamMemberOutput)
}

type RuntimeIamMemberMapOutput struct{ *pulumi.OutputState }

func (RuntimeIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuntimeIamMember)(nil)).Elem()
}

func (o RuntimeIamMemberMapOutput) ToRuntimeIamMemberMapOutput() RuntimeIamMemberMapOutput {
	return o
}

func (o RuntimeIamMemberMapOutput) ToRuntimeIamMemberMapOutputWithContext(ctx context.Context) RuntimeIamMemberMapOutput {
	return o
}

func (o RuntimeIamMemberMapOutput) MapIndex(k pulumi.StringInput) RuntimeIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RuntimeIamMember {
		return vs[0].(map[string]*RuntimeIamMember)[vs[1].(string)]
	}).(RuntimeIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeIamMemberInput)(nil)).Elem(), &RuntimeIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeIamMemberArrayInput)(nil)).Elem(), RuntimeIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeIamMemberMapInput)(nil)).Elem(), RuntimeIamMemberMap{})
	pulumi.RegisterOutputType(RuntimeIamMemberOutput{})
	pulumi.RegisterOutputType(RuntimeIamMemberArrayOutput{})
	pulumi.RegisterOutputType(RuntimeIamMemberMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud AI Notebooks Instance. Each of these resources serves a different use case:
//
// * `notebooks.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
// * `notebooks.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `notebooks.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// > **Note:** `notebooks.InstanceIamPolicy` **cannot** be used in conjunction with `notebooks.InstanceIamBinding` and `notebooks.InstanceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `notebooks.InstanceIamBinding` resources **can be** used in conjunction with `notebooks.InstanceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_notebooks\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = notebooks.NewInstanceIamPolicy(ctx, "policy", &notebooks.InstanceIamPolicyArgs{
// 			Project:      pulumi.Any(google_notebooks_instance.Instance.Project),
// 			Location:     pulumi.Any(google_notebooks_instance.Instance.Location),
// 			InstanceName: pulumi.Any(google_notebooks_instance.Instance.Name),
// 			PolicyData:   pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_notebooks\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewInstanceIamBinding(ctx, "binding", &notebooks.InstanceIamBindingArgs{
// 			Project:      pulumi.Any(google_notebooks_instance.Instance.Project),
// 			Location:     pulumi.Any(google_notebooks_instance.Instance.Location),
// 			InstanceName: pulumi.Any(google_notebooks_instance.Instance.Name),
// 			Role:         pulumi.String("roles/viewer"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_notebooks\_instance\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewInstanceIamMember(ctx, "member", &notebooks.InstanceIamMemberArgs{
// 			Project:      pulumi.Any(google_notebooks_instance.Instance.Project),
// 			Location:     pulumi.Any(google_notebooks_instance.Instance.Location),
// 			InstanceName: pulumi.Any(google_notebooks_instance.Instance.Name),
// 			Role:         pulumi.String("roles/viewer"),
// 			Member:       pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/instances/{{instance_name}} * {{project}}/{{location}}/{{instance_name}} * {{location}}/{{instance_name}} * {{instance_name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud AI Notebooks instance IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/instanceIamBinding:InstanceIamBinding editor "projects/{{project}}/locations/{{location}}/instances/{{instance_name}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/instanceIamBinding:InstanceIamBinding editor "projects/{{project}}/locations/{{location}}/instances/{{instance_name}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/instanceIamBinding:InstanceIamBinding editor projects/{{project}}/locations/{{location}}/instances/{{instance_name}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIamBinding struct {
	pulumi.CustomResourceState

	Condition InstanceIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput      `pulumi:"location"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `notebooks.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewInstanceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewInstanceIamBinding(ctx *pulumi.Context,
	name string, args *InstanceIamBindingArgs, opts ...pulumi.ResourceOption) (*InstanceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource InstanceIamBinding
	err := ctx.RegisterResource("gcp:notebooks/instanceIamBinding:InstanceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIamBinding gets an existing InstanceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIamBindingState, opts ...pulumi.ResourceOption) (*InstanceIamBinding, error) {
	var resource InstanceIamBinding
	err := ctx.ReadResource("gcp:notebooks/instanceIamBinding:InstanceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIamBinding resources.
type instanceIamBindingState struct {
	Condition *InstanceIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName *string `pulumi:"instanceName"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `notebooks.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type InstanceIamBindingState struct {
	Condition InstanceIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringPtrInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `notebooks.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (InstanceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamBindingState)(nil)).Elem()
}

type instanceIamBindingArgs struct {
	Condition *InstanceIamBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName string `pulumi:"instanceName"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `notebooks.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a InstanceIamBinding resource.
type InstanceIamBindingArgs struct {
	Condition InstanceIamBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `notebooks.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (InstanceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamBindingArgs)(nil)).Elem()
}

type InstanceIamBindingInput interface {
	pulumi.Input

	ToInstanceIamBindingOutput() InstanceIamBindingOutput
	ToInstanceIamBindingOutputWithContext(ctx context.Context) InstanceIamBindingOutput
}

func (*InstanceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIamBinding)(nil))
}

func (i *InstanceIamBinding) ToInstanceIamBindingOutput() InstanceIamBindingOutput {
	return i.ToInstanceIamBindingOutputWithContext(context.Background())
}

func (i *InstanceIamBinding) ToInstanceIamBindingOutputWithContext(ctx context.Context) InstanceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamBindingOutput)
}

func (i *InstanceIamBinding) ToInstanceIamBindingPtrOutput() InstanceIamBindingPtrOutput {
	return i.ToInstanceIamBindingPtrOutputWithContext(context.Background())
}

func (i *InstanceIamBinding) ToInstanceIamBindingPtrOutputWithContext(ctx context.Context) InstanceIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamBindingPtrOutput)
}

type InstanceIamBindingPtrInput interface {
	pulumi.Input

	ToInstanceIamBindingPtrOutput() InstanceIamBindingPtrOutput
	ToInstanceIamBindingPtrOutputWithContext(ctx context.Context) InstanceIamBindingPtrOutput
}

type instanceIamBindingPtrType InstanceIamBindingArgs

func (*instanceIamBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIamBinding)(nil))
}

func (i *instanceIamBindingPtrType) ToInstanceIamBindingPtrOutput() InstanceIamBindingPtrOutput {
	return i.ToInstanceIamBindingPtrOutputWithContext(context.Background())
}

func (i *instanceIamBindingPtrType) ToInstanceIamBindingPtrOutputWithContext(ctx context.Context) InstanceIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamBindingPtrOutput)
}

// InstanceIamBindingArrayInput is an input type that accepts InstanceIamBindingArray and InstanceIamBindingArrayOutput values.
// You can construct a concrete instance of `InstanceIamBindingArrayInput` via:
//
//          InstanceIamBindingArray{ InstanceIamBindingArgs{...} }
type InstanceIamBindingArrayInput interface {
	pulumi.Input

	ToInstanceIamBindingArrayOutput() InstanceIamBindingArrayOutput
	ToInstanceIamBindingArrayOutputWithContext(context.Context) InstanceIamBindingArrayOutput
}

type InstanceIamBindingArray []InstanceIamBindingInput

func (InstanceIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIamBinding)(nil)).Elem()
}

func (i InstanceIamBindingArray) ToInstanceIamBindingArrayOutput() InstanceIamBindingArrayOutput {
	return i.ToInstanceIamBindingArrayOutputWithContext(context.Background())
}

func (i InstanceIamBindingArray) ToInstanceIamBindingArrayOutputWithContext(ctx context.Context) InstanceIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamBindingArrayOutput)
}

// InstanceIamBindingMapInput is an input type that accepts InstanceIamBindingMap and InstanceIamBindingMapOutput values.
// You can construct a concrete instance of `InstanceIamBindingMapInput` via:
//
//          InstanceIamBindingMap{ "key": InstanceIamBindingArgs{...} }
type InstanceIamBindingMapInput interface {
	pulumi.Input

	ToInstanceIamBindingMapOutput() InstanceIamBindingMapOutput
	ToInstanceIamBindingMapOutputWithContext(context.Context) InstanceIamBindingMapOutput
}

type InstanceIamBindingMap map[string]InstanceIamBindingInput

func (InstanceIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIamBinding)(nil)).Elem()
}

func (i InstanceIamBindingMap) ToInstanceIamBindingMapOutput() InstanceIamBindingMapOutput {
	return i.ToInstanceIamBindingMapOutputWithContext(context.Background())
}

func (i InstanceIamBindingMap) ToInstanceIamBindingMapOutputWithContext(ctx context.Context) InstanceIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamBindingMapOutput)
}

type InstanceIamBindingOutput struct{ *pulumi.OutputState }

func (InstanceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIamBinding)(nil))
}

func (o InstanceIamBindingOutput) ToInstanceIamBindingOutput() InstanceIamBindingOutput {
	return o
}

func (o InstanceIamBindingOutput) ToInstanceIamBindingOutputWithContext(ctx context.Context) InstanceIamBindingOutput {
	return o
}

func (o InstanceIamBindingOutput) ToInstanceIamBindingPtrOutput() InstanceIamBindingPtrOutput {
	return o.ToInstanceIamBindingPtrOutputWithContext(context.Background())
}

func (o InstanceIamBindingOutput) ToInstanceIamBindingPtrOutputWithContext(ctx context.Context) InstanceIamBindingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceIamBinding) *InstanceIamBinding {
		return &v
	}).(InstanceIamBindingPtrOutput)
}

type InstanceIamBindingPtrOutput struct{ *pulumi.OutputState }

func (InstanceIamBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIamBinding)(nil))
}

func (o InstanceIamBindingPtrOutput) ToInstanceIamBindingPtrOutput() InstanceIamBindingPtrOutput {
	return o
}

func (o InstanceIamBindingPtrOutput) ToInstanceIamBindingPtrOutputWithContext(ctx context.Context) InstanceIamBindingPtrOutput {
	return o
}

func (o InstanceIamBindingPtrOutput) Elem() InstanceIamBindingOutput {
	return o.ApplyT(func(v *InstanceIamBinding) InstanceIamBinding {
		if v != nil {
			return *v
		}
		var ret InstanceIamBinding
		return ret
	}).(InstanceIamBindingOutput)
}

type InstanceIamBindingArrayOutput struct{ *pulumi.OutputState }

func (InstanceIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceIamBinding)(nil))
}

func (o InstanceIamBindingArrayOutput) ToInstanceIamBindingArrayOutput() InstanceIamBindingArrayOutput {
	return o
}

func (o InstanceIamBindingArrayOutput) ToInstanceIamBindingArrayOutputWithContext(ctx context.Context) InstanceIamBindingArrayOutput {
	return o
}

func (o InstanceIamBindingArrayOutput) Index(i pulumi.IntInput) InstanceIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceIamBinding {
		return vs[0].([]InstanceIamBinding)[vs[1].(int)]
	}).(InstanceIamBindingOutput)
}

type InstanceIamBindingMapOutput struct{ *pulumi.OutputState }

func (InstanceIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceIamBinding)(nil))
}

func (o InstanceIamBindingMapOutput) ToInstanceIamBindingMapOutput() InstanceIamBindingMapOutput {
	return o
}

func (o InstanceIamBindingMapOutput) ToInstanceIamBindingMapOutputWithContext(ctx context.Context) InstanceIamBindingMapOutput {
	return o
}

func (o InstanceIamBindingMapOutput) MapIndex(k pulumi.StringInput) InstanceIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceIamBinding {
		return vs[0].(map[string]InstanceIamBinding)[vs[1].(string)]
	}).(InstanceIamBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceIamBindingOutput{})
	pulumi.RegisterOutputType(InstanceIamBindingPtrOutput{})
	pulumi.RegisterOutputType(InstanceIamBindingArrayOutput{})
	pulumi.RegisterOutputType(InstanceIamBindingMapOutput{})
}

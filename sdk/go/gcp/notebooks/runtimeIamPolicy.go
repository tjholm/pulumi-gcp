// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
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
// 		_, err = notebooks.NewRuntimeIamPolicy(ctx, "policy", &notebooks.RuntimeIamPolicyArgs{
// 			Project:     pulumi.Any(google_notebooks_runtime.Runtime.Project),
// 			Location:    pulumi.Any(google_notebooks_runtime.Runtime.Location),
// 			RuntimeName: pulumi.Any(google_notebooks_runtime.Runtime.Name),
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
// ## google\_notebooks\_runtime\_iam\_binding
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
// 		_, err := notebooks.NewRuntimeIamBinding(ctx, "binding", &notebooks.RuntimeIamBindingArgs{
// 			Project:     pulumi.Any(google_notebooks_runtime.Runtime.Project),
// 			Location:    pulumi.Any(google_notebooks_runtime.Runtime.Location),
// 			RuntimeName: pulumi.Any(google_notebooks_runtime.Runtime.Name),
// 			Role:        pulumi.String("roles/viewer"),
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
// ## google\_notebooks\_runtime\_iam\_member
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
// 		_, err := notebooks.NewRuntimeIamMember(ctx, "member", &notebooks.RuntimeIamMemberArgs{
// 			Project:     pulumi.Any(google_notebooks_runtime.Runtime.Project),
// 			Location:    pulumi.Any(google_notebooks_runtime.Runtime.Location),
// 			RuntimeName: pulumi.Any(google_notebooks_runtime.Runtime.Name),
// 			Role:        pulumi.String("roles/viewer"),
// 			Member:      pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/runtimes/{{runtime_name}} * {{project}}/{{location}}/{{runtime_name}} * {{location}}/{{runtime_name}} * {{runtime_name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud AI Notebooks runtime IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/runtimeIamPolicy:RuntimeIamPolicy editor "projects/{{project}}/locations/{{location}}/runtimes/{{runtime_name}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/runtimeIamPolicy:RuntimeIamPolicy editor "projects/{{project}}/locations/{{location}}/runtimes/{{runtime_name}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/runtimeIamPolicy:RuntimeIamPolicy editor projects/{{project}}/locations/{{location}}/runtimes/{{runtime_name}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RuntimeIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	RuntimeName pulumi.StringOutput `pulumi:"runtimeName"`
}

// NewRuntimeIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewRuntimeIamPolicy(ctx *pulumi.Context,
	name string, args *RuntimeIamPolicyArgs, opts ...pulumi.ResourceOption) (*RuntimeIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.RuntimeName == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeName'")
	}
	var resource RuntimeIamPolicy
	err := ctx.RegisterResource("gcp:notebooks/runtimeIamPolicy:RuntimeIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntimeIamPolicy gets an existing RuntimeIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntimeIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeIamPolicyState, opts ...pulumi.ResourceOption) (*RuntimeIamPolicy, error) {
	var resource RuntimeIamPolicy
	err := ctx.ReadResource("gcp:notebooks/runtimeIamPolicy:RuntimeIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuntimeIamPolicy resources.
type runtimeIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	RuntimeName *string `pulumi:"runtimeName"`
}

type RuntimeIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	RuntimeName pulumi.StringPtrInput
}

func (RuntimeIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeIamPolicyState)(nil)).Elem()
}

type runtimeIamPolicyArgs struct {
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	RuntimeName string `pulumi:"runtimeName"`
}

// The set of arguments for constructing a RuntimeIamPolicy resource.
type RuntimeIamPolicyArgs struct {
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	RuntimeName pulumi.StringInput
}

func (RuntimeIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeIamPolicyArgs)(nil)).Elem()
}

type RuntimeIamPolicyInput interface {
	pulumi.Input

	ToRuntimeIamPolicyOutput() RuntimeIamPolicyOutput
	ToRuntimeIamPolicyOutputWithContext(ctx context.Context) RuntimeIamPolicyOutput
}

func (*RuntimeIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeIamPolicy)(nil)).Elem()
}

func (i *RuntimeIamPolicy) ToRuntimeIamPolicyOutput() RuntimeIamPolicyOutput {
	return i.ToRuntimeIamPolicyOutputWithContext(context.Background())
}

func (i *RuntimeIamPolicy) ToRuntimeIamPolicyOutputWithContext(ctx context.Context) RuntimeIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeIamPolicyOutput)
}

// RuntimeIamPolicyArrayInput is an input type that accepts RuntimeIamPolicyArray and RuntimeIamPolicyArrayOutput values.
// You can construct a concrete instance of `RuntimeIamPolicyArrayInput` via:
//
//          RuntimeIamPolicyArray{ RuntimeIamPolicyArgs{...} }
type RuntimeIamPolicyArrayInput interface {
	pulumi.Input

	ToRuntimeIamPolicyArrayOutput() RuntimeIamPolicyArrayOutput
	ToRuntimeIamPolicyArrayOutputWithContext(context.Context) RuntimeIamPolicyArrayOutput
}

type RuntimeIamPolicyArray []RuntimeIamPolicyInput

func (RuntimeIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuntimeIamPolicy)(nil)).Elem()
}

func (i RuntimeIamPolicyArray) ToRuntimeIamPolicyArrayOutput() RuntimeIamPolicyArrayOutput {
	return i.ToRuntimeIamPolicyArrayOutputWithContext(context.Background())
}

func (i RuntimeIamPolicyArray) ToRuntimeIamPolicyArrayOutputWithContext(ctx context.Context) RuntimeIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeIamPolicyArrayOutput)
}

// RuntimeIamPolicyMapInput is an input type that accepts RuntimeIamPolicyMap and RuntimeIamPolicyMapOutput values.
// You can construct a concrete instance of `RuntimeIamPolicyMapInput` via:
//
//          RuntimeIamPolicyMap{ "key": RuntimeIamPolicyArgs{...} }
type RuntimeIamPolicyMapInput interface {
	pulumi.Input

	ToRuntimeIamPolicyMapOutput() RuntimeIamPolicyMapOutput
	ToRuntimeIamPolicyMapOutputWithContext(context.Context) RuntimeIamPolicyMapOutput
}

type RuntimeIamPolicyMap map[string]RuntimeIamPolicyInput

func (RuntimeIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuntimeIamPolicy)(nil)).Elem()
}

func (i RuntimeIamPolicyMap) ToRuntimeIamPolicyMapOutput() RuntimeIamPolicyMapOutput {
	return i.ToRuntimeIamPolicyMapOutputWithContext(context.Background())
}

func (i RuntimeIamPolicyMap) ToRuntimeIamPolicyMapOutputWithContext(ctx context.Context) RuntimeIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeIamPolicyMapOutput)
}

type RuntimeIamPolicyOutput struct{ *pulumi.OutputState }

func (RuntimeIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeIamPolicy)(nil)).Elem()
}

func (o RuntimeIamPolicyOutput) ToRuntimeIamPolicyOutput() RuntimeIamPolicyOutput {
	return o
}

func (o RuntimeIamPolicyOutput) ToRuntimeIamPolicyOutputWithContext(ctx context.Context) RuntimeIamPolicyOutput {
	return o
}

type RuntimeIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (RuntimeIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuntimeIamPolicy)(nil)).Elem()
}

func (o RuntimeIamPolicyArrayOutput) ToRuntimeIamPolicyArrayOutput() RuntimeIamPolicyArrayOutput {
	return o
}

func (o RuntimeIamPolicyArrayOutput) ToRuntimeIamPolicyArrayOutputWithContext(ctx context.Context) RuntimeIamPolicyArrayOutput {
	return o
}

func (o RuntimeIamPolicyArrayOutput) Index(i pulumi.IntInput) RuntimeIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RuntimeIamPolicy {
		return vs[0].([]*RuntimeIamPolicy)[vs[1].(int)]
	}).(RuntimeIamPolicyOutput)
}

type RuntimeIamPolicyMapOutput struct{ *pulumi.OutputState }

func (RuntimeIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuntimeIamPolicy)(nil)).Elem()
}

func (o RuntimeIamPolicyMapOutput) ToRuntimeIamPolicyMapOutput() RuntimeIamPolicyMapOutput {
	return o
}

func (o RuntimeIamPolicyMapOutput) ToRuntimeIamPolicyMapOutputWithContext(ctx context.Context) RuntimeIamPolicyMapOutput {
	return o
}

func (o RuntimeIamPolicyMapOutput) MapIndex(k pulumi.StringInput) RuntimeIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RuntimeIamPolicy {
		return vs[0].(map[string]*RuntimeIamPolicy)[vs[1].(string)]
	}).(RuntimeIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeIamPolicyInput)(nil)).Elem(), &RuntimeIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeIamPolicyArrayInput)(nil)).Elem(), RuntimeIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeIamPolicyMapInput)(nil)).Elem(), RuntimeIamPolicyMap{})
	pulumi.RegisterOutputType(RuntimeIamPolicyOutput{})
	pulumi.RegisterOutputType(RuntimeIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(RuntimeIamPolicyMapOutput{})
}

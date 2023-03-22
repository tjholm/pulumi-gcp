// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Functions CloudFunction. Each of these resources serves a different use case:
//
// * `cloudfunctions.FunctionIamPolicy`: Authoritative. Sets the IAM policy for the cloudfunction and replaces any existing policy already attached.
// * `cloudfunctions.FunctionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cloudfunction are preserved.
// * `cloudfunctions.FunctionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cloudfunction are preserved.
//
// > **Note:** `cloudfunctions.FunctionIamPolicy` **cannot** be used in conjunction with `cloudfunctions.FunctionIamBinding` and `cloudfunctions.FunctionIamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudfunctions.FunctionIamBinding` resources **can be** used in conjunction with `cloudfunctions.FunctionIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_cloudfunctions\_function\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
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
//			_, err = cloudfunctions.NewFunctionIamPolicy(ctx, "policy", &cloudfunctions.FunctionIamPolicyArgs{
//				Project:       pulumi.Any(google_cloudfunctions_function.Function.Project),
//				Region:        pulumi.Any(google_cloudfunctions_function.Function.Region),
//				CloudFunction: pulumi.Any(google_cloudfunctions_function.Function.Name),
//				PolicyData:    *pulumi.String(admin.PolicyData),
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
// ## google\_cloudfunctions\_function\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfunctions.NewFunctionIamBinding(ctx, "binding", &cloudfunctions.FunctionIamBindingArgs{
//				Project:       pulumi.Any(google_cloudfunctions_function.Function.Project),
//				Region:        pulumi.Any(google_cloudfunctions_function.Function.Region),
//				CloudFunction: pulumi.Any(google_cloudfunctions_function.Function.Name),
//				Role:          pulumi.String("roles/viewer"),
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
// ## google\_cloudfunctions\_function\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfunctions.NewFunctionIamMember(ctx, "member", &cloudfunctions.FunctionIamMemberArgs{
//				Project:       pulumi.Any(google_cloudfunctions_function.Function.Project),
//				Region:        pulumi.Any(google_cloudfunctions_function.Function.Region),
//				CloudFunction: pulumi.Any(google_cloudfunctions_function.Function.Name),
//				Role:          pulumi.String("roles/viewer"),
//				Member:        pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} * {{project}}/{{region}}/{{cloud_function}} * {{region}}/{{cloud_function}} * {{cloud_function}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Functions cloudfunction IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudfunctions/functionIamBinding:FunctionIamBinding editor "projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudfunctions/functionIamBinding:FunctionIamBinding editor "projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudfunctions/functionIamBinding:FunctionIamBinding editor projects/{{project}}/locations/{{region}}/functions/{{cloud_function}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type FunctionIamBinding struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringOutput                  `pulumi:"cloudFunction"`
	Condition     FunctionIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewFunctionIamBinding registers a new resource with the given unique name, arguments, and options.
func NewFunctionIamBinding(ctx *pulumi.Context,
	name string, args *FunctionIamBindingArgs, opts ...pulumi.ResourceOption) (*FunctionIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudFunction == nil {
		return nil, errors.New("invalid value for required argument 'CloudFunction'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource FunctionIamBinding
	err := ctx.RegisterResource("gcp:cloudfunctions/functionIamBinding:FunctionIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionIamBinding gets an existing FunctionIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionIamBindingState, opts ...pulumi.ResourceOption) (*FunctionIamBinding, error) {
	var resource FunctionIamBinding
	err := ctx.ReadResource("gcp:cloudfunctions/functionIamBinding:FunctionIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionIamBinding resources.
type functionIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction *string                      `pulumi:"cloudFunction"`
	Condition     *FunctionIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type FunctionIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringPtrInput
	Condition     FunctionIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (FunctionIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamBindingState)(nil)).Elem()
}

type functionIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction string                       `pulumi:"cloudFunction"`
	Condition     *FunctionIamBindingCondition `pulumi:"condition"`
	Members       []string                     `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a FunctionIamBinding resource.
type FunctionIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringInput
	Condition     FunctionIamBindingConditionPtrInput
	Members       pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (FunctionIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamBindingArgs)(nil)).Elem()
}

type FunctionIamBindingInput interface {
	pulumi.Input

	ToFunctionIamBindingOutput() FunctionIamBindingOutput
	ToFunctionIamBindingOutputWithContext(ctx context.Context) FunctionIamBindingOutput
}

func (*FunctionIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionIamBinding)(nil)).Elem()
}

func (i *FunctionIamBinding) ToFunctionIamBindingOutput() FunctionIamBindingOutput {
	return i.ToFunctionIamBindingOutputWithContext(context.Background())
}

func (i *FunctionIamBinding) ToFunctionIamBindingOutputWithContext(ctx context.Context) FunctionIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamBindingOutput)
}

// FunctionIamBindingArrayInput is an input type that accepts FunctionIamBindingArray and FunctionIamBindingArrayOutput values.
// You can construct a concrete instance of `FunctionIamBindingArrayInput` via:
//
//	FunctionIamBindingArray{ FunctionIamBindingArgs{...} }
type FunctionIamBindingArrayInput interface {
	pulumi.Input

	ToFunctionIamBindingArrayOutput() FunctionIamBindingArrayOutput
	ToFunctionIamBindingArrayOutputWithContext(context.Context) FunctionIamBindingArrayOutput
}

type FunctionIamBindingArray []FunctionIamBindingInput

func (FunctionIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionIamBinding)(nil)).Elem()
}

func (i FunctionIamBindingArray) ToFunctionIamBindingArrayOutput() FunctionIamBindingArrayOutput {
	return i.ToFunctionIamBindingArrayOutputWithContext(context.Background())
}

func (i FunctionIamBindingArray) ToFunctionIamBindingArrayOutputWithContext(ctx context.Context) FunctionIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamBindingArrayOutput)
}

// FunctionIamBindingMapInput is an input type that accepts FunctionIamBindingMap and FunctionIamBindingMapOutput values.
// You can construct a concrete instance of `FunctionIamBindingMapInput` via:
//
//	FunctionIamBindingMap{ "key": FunctionIamBindingArgs{...} }
type FunctionIamBindingMapInput interface {
	pulumi.Input

	ToFunctionIamBindingMapOutput() FunctionIamBindingMapOutput
	ToFunctionIamBindingMapOutputWithContext(context.Context) FunctionIamBindingMapOutput
}

type FunctionIamBindingMap map[string]FunctionIamBindingInput

func (FunctionIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionIamBinding)(nil)).Elem()
}

func (i FunctionIamBindingMap) ToFunctionIamBindingMapOutput() FunctionIamBindingMapOutput {
	return i.ToFunctionIamBindingMapOutputWithContext(context.Background())
}

func (i FunctionIamBindingMap) ToFunctionIamBindingMapOutputWithContext(ctx context.Context) FunctionIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamBindingMapOutput)
}

type FunctionIamBindingOutput struct{ *pulumi.OutputState }

func (FunctionIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionIamBinding)(nil)).Elem()
}

func (o FunctionIamBindingOutput) ToFunctionIamBindingOutput() FunctionIamBindingOutput {
	return o
}

func (o FunctionIamBindingOutput) ToFunctionIamBindingOutputWithContext(ctx context.Context) FunctionIamBindingOutput {
	return o
}

// Used to find the parent resource to bind the IAM policy to
func (o FunctionIamBindingOutput) CloudFunction() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamBinding) pulumi.StringOutput { return v.CloudFunction }).(pulumi.StringOutput)
}

func (o FunctionIamBindingOutput) Condition() FunctionIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *FunctionIamBinding) FunctionIamBindingConditionPtrOutput { return v.Condition }).(FunctionIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o FunctionIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o FunctionIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FunctionIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o FunctionIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
// region is specified, it is taken from the provider configuration.
func (o FunctionIamBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o FunctionIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type FunctionIamBindingArrayOutput struct{ *pulumi.OutputState }

func (FunctionIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionIamBinding)(nil)).Elem()
}

func (o FunctionIamBindingArrayOutput) ToFunctionIamBindingArrayOutput() FunctionIamBindingArrayOutput {
	return o
}

func (o FunctionIamBindingArrayOutput) ToFunctionIamBindingArrayOutputWithContext(ctx context.Context) FunctionIamBindingArrayOutput {
	return o
}

func (o FunctionIamBindingArrayOutput) Index(i pulumi.IntInput) FunctionIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionIamBinding {
		return vs[0].([]*FunctionIamBinding)[vs[1].(int)]
	}).(FunctionIamBindingOutput)
}

type FunctionIamBindingMapOutput struct{ *pulumi.OutputState }

func (FunctionIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionIamBinding)(nil)).Elem()
}

func (o FunctionIamBindingMapOutput) ToFunctionIamBindingMapOutput() FunctionIamBindingMapOutput {
	return o
}

func (o FunctionIamBindingMapOutput) ToFunctionIamBindingMapOutputWithContext(ctx context.Context) FunctionIamBindingMapOutput {
	return o
}

func (o FunctionIamBindingMapOutput) MapIndex(k pulumi.StringInput) FunctionIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionIamBinding {
		return vs[0].(map[string]*FunctionIamBinding)[vs[1].(string)]
	}).(FunctionIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamBindingInput)(nil)).Elem(), &FunctionIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamBindingArrayInput)(nil)).Elem(), FunctionIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamBindingMapInput)(nil)).Elem(), FunctionIamBindingMap{})
	pulumi.RegisterOutputType(FunctionIamBindingOutput{})
	pulumi.RegisterOutputType(FunctionIamBindingArrayOutput{})
	pulumi.RegisterOutputType(FunctionIamBindingMapOutput{})
}

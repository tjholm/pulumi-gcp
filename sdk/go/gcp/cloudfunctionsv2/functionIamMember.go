// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctionsv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Functions (2nd gen) function. Each of these resources serves a different use case:
//
// * `cloudfunctionsv2.FunctionIamPolicy`: Authoritative. Sets the IAM policy for the function and replaces any existing policy already attached.
// * `cloudfunctionsv2.FunctionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the function are preserved.
// * `cloudfunctionsv2.FunctionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the function are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `cloudfunctionsv2.FunctionIamPolicy`: Retrieves the IAM policy for the function
//
// > **Note:** `cloudfunctionsv2.FunctionIamPolicy` **cannot** be used in conjunction with `cloudfunctionsv2.FunctionIamBinding` and `cloudfunctionsv2.FunctionIamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudfunctionsv2.FunctionIamBinding` resources **can be** used in conjunction with `cloudfunctionsv2.FunctionIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## cloudfunctionsv2.FunctionIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudfunctionsv2"
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
//			_, err = cloudfunctionsv2.NewFunctionIamPolicy(ctx, "policy", &cloudfunctionsv2.FunctionIamPolicyArgs{
//				Project:       pulumi.Any(function.Project),
//				Location:      pulumi.Any(function.Location),
//				CloudFunction: pulumi.Any(function.Name),
//				PolicyData:    pulumi.String(admin.PolicyData),
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
// ## cloudfunctionsv2.FunctionIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudfunctionsv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfunctionsv2.NewFunctionIamBinding(ctx, "binding", &cloudfunctionsv2.FunctionIamBindingArgs{
//				Project:       pulumi.Any(function.Project),
//				Location:      pulumi.Any(function.Location),
//				CloudFunction: pulumi.Any(function.Name),
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
// ## cloudfunctionsv2.FunctionIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudfunctionsv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfunctionsv2.NewFunctionIamMember(ctx, "member", &cloudfunctionsv2.FunctionIamMemberArgs{
//				Project:       pulumi.Any(function.Project),
//				Location:      pulumi.Any(function.Location),
//				CloudFunction: pulumi.Any(function.Name),
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
// ## cloudfunctionsv2.FunctionIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudfunctionsv2"
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
//			_, err = cloudfunctionsv2.NewFunctionIamPolicy(ctx, "policy", &cloudfunctionsv2.FunctionIamPolicyArgs{
//				Project:       pulumi.Any(function.Project),
//				Location:      pulumi.Any(function.Location),
//				CloudFunction: pulumi.Any(function.Name),
//				PolicyData:    pulumi.String(admin.PolicyData),
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
// ## cloudfunctionsv2.FunctionIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudfunctionsv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfunctionsv2.NewFunctionIamBinding(ctx, "binding", &cloudfunctionsv2.FunctionIamBindingArgs{
//				Project:       pulumi.Any(function.Project),
//				Location:      pulumi.Any(function.Location),
//				CloudFunction: pulumi.Any(function.Name),
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
// ## cloudfunctionsv2.FunctionIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudfunctionsv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfunctionsv2.NewFunctionIamMember(ctx, "member", &cloudfunctionsv2.FunctionIamMemberArgs{
//				Project:       pulumi.Any(function.Project),
//				Location:      pulumi.Any(function.Location),
//				CloudFunction: pulumi.Any(function.Name),
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
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/locations/{{location}}/functions/{{cloud_function}}
//
// * {{project}}/{{location}}/{{cloud_function}}
//
// * {{location}}/{{cloud_function}}
//
// * {{cloud_function}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Cloud Functions (2nd gen) function IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:cloudfunctionsv2/functionIamMember:FunctionIamMember editor "projects/{{project}}/locations/{{location}}/functions/{{cloud_function}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:cloudfunctionsv2/functionIamMember:FunctionIamMember editor "projects/{{project}}/locations/{{location}}/functions/{{cloud_function}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:cloudfunctionsv2/functionIamMember:FunctionIamMember editor projects/{{project}}/locations/{{location}}/functions/{{cloud_function}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type FunctionIamMember struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringOutput                 `pulumi:"cloudFunction"`
	Condition     FunctionIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
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
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudfunctionsv2.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewFunctionIamMember registers a new resource with the given unique name, arguments, and options.
func NewFunctionIamMember(ctx *pulumi.Context,
	name string, args *FunctionIamMemberArgs, opts ...pulumi.ResourceOption) (*FunctionIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudFunction == nil {
		return nil, errors.New("invalid value for required argument 'CloudFunction'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FunctionIamMember
	err := ctx.RegisterResource("gcp:cloudfunctionsv2/functionIamMember:FunctionIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionIamMember gets an existing FunctionIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionIamMemberState, opts ...pulumi.ResourceOption) (*FunctionIamMember, error) {
	var resource FunctionIamMember
	err := ctx.ReadResource("gcp:cloudfunctionsv2/functionIamMember:FunctionIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionIamMember resources.
type functionIamMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction *string                     `pulumi:"cloudFunction"`
	Condition     *FunctionIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
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
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudfunctionsv2.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type FunctionIamMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringPtrInput
	Condition     FunctionIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
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
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudfunctionsv2.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (FunctionIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamMemberState)(nil)).Elem()
}

type functionIamMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction string                      `pulumi:"cloudFunction"`
	Condition     *FunctionIamMemberCondition `pulumi:"condition"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
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
	Member string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudfunctionsv2.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a FunctionIamMember resource.
type FunctionIamMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringInput
	Condition     FunctionIamMemberConditionPtrInput
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
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
	Member pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudfunctionsv2.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (FunctionIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamMemberArgs)(nil)).Elem()
}

type FunctionIamMemberInput interface {
	pulumi.Input

	ToFunctionIamMemberOutput() FunctionIamMemberOutput
	ToFunctionIamMemberOutputWithContext(ctx context.Context) FunctionIamMemberOutput
}

func (*FunctionIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionIamMember)(nil)).Elem()
}

func (i *FunctionIamMember) ToFunctionIamMemberOutput() FunctionIamMemberOutput {
	return i.ToFunctionIamMemberOutputWithContext(context.Background())
}

func (i *FunctionIamMember) ToFunctionIamMemberOutputWithContext(ctx context.Context) FunctionIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamMemberOutput)
}

// FunctionIamMemberArrayInput is an input type that accepts FunctionIamMemberArray and FunctionIamMemberArrayOutput values.
// You can construct a concrete instance of `FunctionIamMemberArrayInput` via:
//
//	FunctionIamMemberArray{ FunctionIamMemberArgs{...} }
type FunctionIamMemberArrayInput interface {
	pulumi.Input

	ToFunctionIamMemberArrayOutput() FunctionIamMemberArrayOutput
	ToFunctionIamMemberArrayOutputWithContext(context.Context) FunctionIamMemberArrayOutput
}

type FunctionIamMemberArray []FunctionIamMemberInput

func (FunctionIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionIamMember)(nil)).Elem()
}

func (i FunctionIamMemberArray) ToFunctionIamMemberArrayOutput() FunctionIamMemberArrayOutput {
	return i.ToFunctionIamMemberArrayOutputWithContext(context.Background())
}

func (i FunctionIamMemberArray) ToFunctionIamMemberArrayOutputWithContext(ctx context.Context) FunctionIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamMemberArrayOutput)
}

// FunctionIamMemberMapInput is an input type that accepts FunctionIamMemberMap and FunctionIamMemberMapOutput values.
// You can construct a concrete instance of `FunctionIamMemberMapInput` via:
//
//	FunctionIamMemberMap{ "key": FunctionIamMemberArgs{...} }
type FunctionIamMemberMapInput interface {
	pulumi.Input

	ToFunctionIamMemberMapOutput() FunctionIamMemberMapOutput
	ToFunctionIamMemberMapOutputWithContext(context.Context) FunctionIamMemberMapOutput
}

type FunctionIamMemberMap map[string]FunctionIamMemberInput

func (FunctionIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionIamMember)(nil)).Elem()
}

func (i FunctionIamMemberMap) ToFunctionIamMemberMapOutput() FunctionIamMemberMapOutput {
	return i.ToFunctionIamMemberMapOutputWithContext(context.Background())
}

func (i FunctionIamMemberMap) ToFunctionIamMemberMapOutputWithContext(ctx context.Context) FunctionIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamMemberMapOutput)
}

type FunctionIamMemberOutput struct{ *pulumi.OutputState }

func (FunctionIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionIamMember)(nil)).Elem()
}

func (o FunctionIamMemberOutput) ToFunctionIamMemberOutput() FunctionIamMemberOutput {
	return o
}

func (o FunctionIamMemberOutput) ToFunctionIamMemberOutputWithContext(ctx context.Context) FunctionIamMemberOutput {
	return o
}

// Used to find the parent resource to bind the IAM policy to
func (o FunctionIamMemberOutput) CloudFunction() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamMember) pulumi.StringOutput { return v.CloudFunction }).(pulumi.StringOutput)
}

func (o FunctionIamMemberOutput) Condition() FunctionIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *FunctionIamMember) FunctionIamMemberConditionPtrOutput { return v.Condition }).(FunctionIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o FunctionIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o FunctionIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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
func (o FunctionIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o FunctionIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `cloudfunctionsv2.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o FunctionIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type FunctionIamMemberArrayOutput struct{ *pulumi.OutputState }

func (FunctionIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionIamMember)(nil)).Elem()
}

func (o FunctionIamMemberArrayOutput) ToFunctionIamMemberArrayOutput() FunctionIamMemberArrayOutput {
	return o
}

func (o FunctionIamMemberArrayOutput) ToFunctionIamMemberArrayOutputWithContext(ctx context.Context) FunctionIamMemberArrayOutput {
	return o
}

func (o FunctionIamMemberArrayOutput) Index(i pulumi.IntInput) FunctionIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionIamMember {
		return vs[0].([]*FunctionIamMember)[vs[1].(int)]
	}).(FunctionIamMemberOutput)
}

type FunctionIamMemberMapOutput struct{ *pulumi.OutputState }

func (FunctionIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionIamMember)(nil)).Elem()
}

func (o FunctionIamMemberMapOutput) ToFunctionIamMemberMapOutput() FunctionIamMemberMapOutput {
	return o
}

func (o FunctionIamMemberMapOutput) ToFunctionIamMemberMapOutputWithContext(ctx context.Context) FunctionIamMemberMapOutput {
	return o
}

func (o FunctionIamMemberMapOutput) MapIndex(k pulumi.StringInput) FunctionIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionIamMember {
		return vs[0].(map[string]*FunctionIamMember)[vs[1].(string)]
	}).(FunctionIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamMemberInput)(nil)).Elem(), &FunctionIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamMemberArrayInput)(nil)).Elem(), FunctionIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamMemberMapInput)(nil)).Elem(), FunctionIamMemberMap{})
	pulumi.RegisterOutputType(FunctionIamMemberOutput{})
	pulumi.RegisterOutputType(FunctionIamMemberArrayOutput{})
	pulumi.RegisterOutputType(FunctionIamMemberMapOutput{})
}

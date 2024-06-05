// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrun

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Run Service. Each of these resources serves a different use case:
//
// * `cloudrun.IamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
// * `cloudrun.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
// * `cloudrun.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `cloudrun.IamPolicy`: Retrieves the IAM policy for the service
//
// > **Note:** `cloudrun.IamPolicy` **cannot** be used in conjunction with `cloudrun.IamBinding` and `cloudrun.IamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudrun.IamBinding` resources **can be** used in conjunction with `cloudrun.IamMember` resources **only if** they do not grant privilege to the same role.
//
// ## cloudrun.IamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudrun"
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
//			_, err = cloudrun.NewIamPolicy(ctx, "policy", &cloudrun.IamPolicyArgs{
//				Location:   pulumi.Any(_default.Location),
//				Project:    pulumi.Any(_default.Project),
//				Service:    pulumi.Any(_default.Name),
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
// ## cloudrun.IamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudrun"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudrun.NewIamBinding(ctx, "binding", &cloudrun.IamBindingArgs{
//				Location: pulumi.Any(_default.Location),
//				Project:  pulumi.Any(_default.Project),
//				Service:  pulumi.Any(_default.Name),
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
// ## cloudrun.IamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudrun"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudrun.NewIamMember(ctx, "member", &cloudrun.IamMemberArgs{
//				Location: pulumi.Any(_default.Location),
//				Project:  pulumi.Any(_default.Project),
//				Service:  pulumi.Any(_default.Name),
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
// ## cloudrun.IamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudrun"
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
//			_, err = cloudrun.NewIamPolicy(ctx, "policy", &cloudrun.IamPolicyArgs{
//				Location:   pulumi.Any(_default.Location),
//				Project:    pulumi.Any(_default.Project),
//				Service:    pulumi.Any(_default.Name),
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
// ## cloudrun.IamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudrun"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudrun.NewIamBinding(ctx, "binding", &cloudrun.IamBindingArgs{
//				Location: pulumi.Any(_default.Location),
//				Project:  pulumi.Any(_default.Project),
//				Service:  pulumi.Any(_default.Name),
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
// ## cloudrun.IamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudrun"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudrun.NewIamMember(ctx, "member", &cloudrun.IamMemberArgs{
//				Location: pulumi.Any(_default.Location),
//				Project:  pulumi.Any(_default.Project),
//				Service:  pulumi.Any(_default.Name),
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
// * projects/{{project}}/locations/{{location}}/services/{{service}}
//
// * {{project}}/{{location}}/{{service}}
//
// * {{location}}/{{service}}
//
// * {{service}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Cloud Run service IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:cloudrun/iamMember:IamMember editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:cloudrun/iamMember:IamMember editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:cloudrun/iamMember:IamMember editor projects/{{project}}/locations/{{location}}/services/{{service}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type IamMember struct {
	pulumi.CustomResourceState

	Condition IamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
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
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewIamMember registers a new resource with the given unique name, arguments, and options.
func NewIamMember(ctx *pulumi.Context,
	name string, args *IamMemberArgs, opts ...pulumi.ResourceOption) (*IamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamMember
	err := ctx.RegisterResource("gcp:cloudrun/iamMember:IamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamMember gets an existing IamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamMemberState, opts ...pulumi.ResourceOption) (*IamMember, error) {
	var resource IamMember
	err := ctx.ReadResource("gcp:cloudrun/iamMember:IamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamMember resources.
type iamMemberState struct {
	Condition *IamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
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
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Service *string `pulumi:"service"`
}

type IamMemberState struct {
	Condition IamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
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
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringPtrInput
}

func (IamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamMemberState)(nil)).Elem()
}

type iamMemberArgs struct {
	Condition *IamMemberCondition `pulumi:"condition"`
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
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
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a IamMember resource.
type IamMemberArgs struct {
	Condition IamMemberConditionPtrInput
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
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
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringInput
}

func (IamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamMemberArgs)(nil)).Elem()
}

type IamMemberInput interface {
	pulumi.Input

	ToIamMemberOutput() IamMemberOutput
	ToIamMemberOutputWithContext(ctx context.Context) IamMemberOutput
}

func (*IamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**IamMember)(nil)).Elem()
}

func (i *IamMember) ToIamMemberOutput() IamMemberOutput {
	return i.ToIamMemberOutputWithContext(context.Background())
}

func (i *IamMember) ToIamMemberOutputWithContext(ctx context.Context) IamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamMemberOutput)
}

// IamMemberArrayInput is an input type that accepts IamMemberArray and IamMemberArrayOutput values.
// You can construct a concrete instance of `IamMemberArrayInput` via:
//
//	IamMemberArray{ IamMemberArgs{...} }
type IamMemberArrayInput interface {
	pulumi.Input

	ToIamMemberArrayOutput() IamMemberArrayOutput
	ToIamMemberArrayOutputWithContext(context.Context) IamMemberArrayOutput
}

type IamMemberArray []IamMemberInput

func (IamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamMember)(nil)).Elem()
}

func (i IamMemberArray) ToIamMemberArrayOutput() IamMemberArrayOutput {
	return i.ToIamMemberArrayOutputWithContext(context.Background())
}

func (i IamMemberArray) ToIamMemberArrayOutputWithContext(ctx context.Context) IamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamMemberArrayOutput)
}

// IamMemberMapInput is an input type that accepts IamMemberMap and IamMemberMapOutput values.
// You can construct a concrete instance of `IamMemberMapInput` via:
//
//	IamMemberMap{ "key": IamMemberArgs{...} }
type IamMemberMapInput interface {
	pulumi.Input

	ToIamMemberMapOutput() IamMemberMapOutput
	ToIamMemberMapOutputWithContext(context.Context) IamMemberMapOutput
}

type IamMemberMap map[string]IamMemberInput

func (IamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamMember)(nil)).Elem()
}

func (i IamMemberMap) ToIamMemberMapOutput() IamMemberMapOutput {
	return i.ToIamMemberMapOutputWithContext(context.Background())
}

func (i IamMemberMap) ToIamMemberMapOutputWithContext(ctx context.Context) IamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamMemberMapOutput)
}

type IamMemberOutput struct{ *pulumi.OutputState }

func (IamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamMember)(nil)).Elem()
}

func (o IamMemberOutput) ToIamMemberOutput() IamMemberOutput {
	return o
}

func (o IamMemberOutput) ToIamMemberOutputWithContext(ctx context.Context) IamMemberOutput {
	return o
}

func (o IamMemberOutput) Condition() IamMemberConditionPtrOutput {
	return o.ApplyT(func(v *IamMember) IamMemberConditionPtrOutput { return v.Condition }).(IamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o IamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o IamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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
func (o IamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *IamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o IamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o IamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o IamMemberOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *IamMember) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

type IamMemberArrayOutput struct{ *pulumi.OutputState }

func (IamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamMember)(nil)).Elem()
}

func (o IamMemberArrayOutput) ToIamMemberArrayOutput() IamMemberArrayOutput {
	return o
}

func (o IamMemberArrayOutput) ToIamMemberArrayOutputWithContext(ctx context.Context) IamMemberArrayOutput {
	return o
}

func (o IamMemberArrayOutput) Index(i pulumi.IntInput) IamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamMember {
		return vs[0].([]*IamMember)[vs[1].(int)]
	}).(IamMemberOutput)
}

type IamMemberMapOutput struct{ *pulumi.OutputState }

func (IamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamMember)(nil)).Elem()
}

func (o IamMemberMapOutput) ToIamMemberMapOutput() IamMemberMapOutput {
	return o
}

func (o IamMemberMapOutput) ToIamMemberMapOutputWithContext(ctx context.Context) IamMemberMapOutput {
	return o
}

func (o IamMemberMapOutput) MapIndex(k pulumi.StringInput) IamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamMember {
		return vs[0].(map[string]*IamMember)[vs[1].(string)]
	}).(IamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamMemberInput)(nil)).Elem(), &IamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamMemberArrayInput)(nil)).Elem(), IamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamMemberMapInput)(nil)).Elem(), IamMemberMap{})
	pulumi.RegisterOutputType(IamMemberOutput{})
	pulumi.RegisterOutputType(IamMemberArrayOutput{})
	pulumi.RegisterOutputType(IamMemberMapOutput{})
}

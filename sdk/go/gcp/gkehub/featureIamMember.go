// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gkehub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for GKEHub Feature. Each of these resources serves a different use case:
//
// * `gkehub.FeatureIamPolicy`: Authoritative. Sets the IAM policy for the feature and replaces any existing policy already attached.
// * `gkehub.FeatureIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the feature are preserved.
// * `gkehub.FeatureIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the feature are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `gkehub.FeatureIamPolicy`: Retrieves the IAM policy for the feature
//
// > **Note:** `gkehub.FeatureIamPolicy` **cannot** be used in conjunction with `gkehub.FeatureIamBinding` and `gkehub.FeatureIamMember` or they will fight over what your policy should be.
//
// > **Note:** `gkehub.FeatureIamBinding` resources **can be** used in conjunction with `gkehub.FeatureIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## gkehub.FeatureIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
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
//			_, err = gkehub.NewFeatureIamPolicy(ctx, "policy", &gkehub.FeatureIamPolicyArgs{
//				Project:    pulumi.Any(feature.Project),
//				Location:   pulumi.Any(feature.Location),
//				Name:       pulumi.Any(feature.Name),
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
// ## gkehub.FeatureIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeatureIamBinding(ctx, "binding", &gkehub.FeatureIamBindingArgs{
//				Project:  pulumi.Any(feature.Project),
//				Location: pulumi.Any(feature.Location),
//				Name:     pulumi.Any(feature.Name),
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
// ## gkehub.FeatureIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeatureIamMember(ctx, "member", &gkehub.FeatureIamMemberArgs{
//				Project:  pulumi.Any(feature.Project),
//				Location: pulumi.Any(feature.Location),
//				Name:     pulumi.Any(feature.Name),
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
// ## gkehub.FeatureIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
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
//			_, err = gkehub.NewFeatureIamPolicy(ctx, "policy", &gkehub.FeatureIamPolicyArgs{
//				Project:    pulumi.Any(feature.Project),
//				Location:   pulumi.Any(feature.Location),
//				Name:       pulumi.Any(feature.Name),
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
// ## gkehub.FeatureIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeatureIamBinding(ctx, "binding", &gkehub.FeatureIamBindingArgs{
//				Project:  pulumi.Any(feature.Project),
//				Location: pulumi.Any(feature.Location),
//				Name:     pulumi.Any(feature.Name),
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
// ## gkehub.FeatureIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeatureIamMember(ctx, "member", &gkehub.FeatureIamMemberArgs{
//				Project:  pulumi.Any(feature.Project),
//				Location: pulumi.Any(feature.Location),
//				Name:     pulumi.Any(feature.Name),
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
// * projects/{{project}}/locations/{{location}}/features/{{name}}
//
// * {{project}}/{{location}}/{{name}}
//
// * {{location}}/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// GKEHub feature IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:gkehub/featureIamMember:FeatureIamMember editor "projects/{{project}}/locations/{{location}}/features/{{feature}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:gkehub/featureIamMember:FeatureIamMember editor "projects/{{project}}/locations/{{location}}/features/{{feature}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:gkehub/featureIamMember:FeatureIamMember editor projects/{{project}}/locations/{{location}}/features/{{feature}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type FeatureIamMember struct {
	pulumi.CustomResourceState

	Condition FeatureIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
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
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `gkehub.FeatureIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewFeatureIamMember registers a new resource with the given unique name, arguments, and options.
func NewFeatureIamMember(ctx *pulumi.Context,
	name string, args *FeatureIamMemberArgs, opts ...pulumi.ResourceOption) (*FeatureIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeatureIamMember
	err := ctx.RegisterResource("gcp:gkehub/featureIamMember:FeatureIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeatureIamMember gets an existing FeatureIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeatureIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeatureIamMemberState, opts ...pulumi.ResourceOption) (*FeatureIamMember, error) {
	var resource FeatureIamMember
	err := ctx.ReadResource("gcp:gkehub/featureIamMember:FeatureIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeatureIamMember resources.
type featureIamMemberState struct {
	Condition *FeatureIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
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
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `gkehub.FeatureIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type FeatureIamMemberState struct {
	Condition FeatureIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
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
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `gkehub.FeatureIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (FeatureIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*featureIamMemberState)(nil)).Elem()
}

type featureIamMemberArgs struct {
	Condition *FeatureIamMemberCondition `pulumi:"condition"`
	// The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
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
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `gkehub.FeatureIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a FeatureIamMember resource.
type FeatureIamMemberArgs struct {
	Condition FeatureIamMemberConditionPtrInput
	// The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
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
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `gkehub.FeatureIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (FeatureIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featureIamMemberArgs)(nil)).Elem()
}

type FeatureIamMemberInput interface {
	pulumi.Input

	ToFeatureIamMemberOutput() FeatureIamMemberOutput
	ToFeatureIamMemberOutputWithContext(ctx context.Context) FeatureIamMemberOutput
}

func (*FeatureIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureIamMember)(nil)).Elem()
}

func (i *FeatureIamMember) ToFeatureIamMemberOutput() FeatureIamMemberOutput {
	return i.ToFeatureIamMemberOutputWithContext(context.Background())
}

func (i *FeatureIamMember) ToFeatureIamMemberOutputWithContext(ctx context.Context) FeatureIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureIamMemberOutput)
}

// FeatureIamMemberArrayInput is an input type that accepts FeatureIamMemberArray and FeatureIamMemberArrayOutput values.
// You can construct a concrete instance of `FeatureIamMemberArrayInput` via:
//
//	FeatureIamMemberArray{ FeatureIamMemberArgs{...} }
type FeatureIamMemberArrayInput interface {
	pulumi.Input

	ToFeatureIamMemberArrayOutput() FeatureIamMemberArrayOutput
	ToFeatureIamMemberArrayOutputWithContext(context.Context) FeatureIamMemberArrayOutput
}

type FeatureIamMemberArray []FeatureIamMemberInput

func (FeatureIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureIamMember)(nil)).Elem()
}

func (i FeatureIamMemberArray) ToFeatureIamMemberArrayOutput() FeatureIamMemberArrayOutput {
	return i.ToFeatureIamMemberArrayOutputWithContext(context.Background())
}

func (i FeatureIamMemberArray) ToFeatureIamMemberArrayOutputWithContext(ctx context.Context) FeatureIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureIamMemberArrayOutput)
}

// FeatureIamMemberMapInput is an input type that accepts FeatureIamMemberMap and FeatureIamMemberMapOutput values.
// You can construct a concrete instance of `FeatureIamMemberMapInput` via:
//
//	FeatureIamMemberMap{ "key": FeatureIamMemberArgs{...} }
type FeatureIamMemberMapInput interface {
	pulumi.Input

	ToFeatureIamMemberMapOutput() FeatureIamMemberMapOutput
	ToFeatureIamMemberMapOutputWithContext(context.Context) FeatureIamMemberMapOutput
}

type FeatureIamMemberMap map[string]FeatureIamMemberInput

func (FeatureIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureIamMember)(nil)).Elem()
}

func (i FeatureIamMemberMap) ToFeatureIamMemberMapOutput() FeatureIamMemberMapOutput {
	return i.ToFeatureIamMemberMapOutputWithContext(context.Background())
}

func (i FeatureIamMemberMap) ToFeatureIamMemberMapOutputWithContext(ctx context.Context) FeatureIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureIamMemberMapOutput)
}

type FeatureIamMemberOutput struct{ *pulumi.OutputState }

func (FeatureIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureIamMember)(nil)).Elem()
}

func (o FeatureIamMemberOutput) ToFeatureIamMemberOutput() FeatureIamMemberOutput {
	return o
}

func (o FeatureIamMemberOutput) ToFeatureIamMemberOutputWithContext(ctx context.Context) FeatureIamMemberOutput {
	return o
}

func (o FeatureIamMemberOutput) Condition() FeatureIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *FeatureIamMember) FeatureIamMemberConditionPtrOutput { return v.Condition }).(FeatureIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o FeatureIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o FeatureIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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
func (o FeatureIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o FeatureIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o FeatureIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `gkehub.FeatureIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o FeatureIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type FeatureIamMemberArrayOutput struct{ *pulumi.OutputState }

func (FeatureIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureIamMember)(nil)).Elem()
}

func (o FeatureIamMemberArrayOutput) ToFeatureIamMemberArrayOutput() FeatureIamMemberArrayOutput {
	return o
}

func (o FeatureIamMemberArrayOutput) ToFeatureIamMemberArrayOutputWithContext(ctx context.Context) FeatureIamMemberArrayOutput {
	return o
}

func (o FeatureIamMemberArrayOutput) Index(i pulumi.IntInput) FeatureIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FeatureIamMember {
		return vs[0].([]*FeatureIamMember)[vs[1].(int)]
	}).(FeatureIamMemberOutput)
}

type FeatureIamMemberMapOutput struct{ *pulumi.OutputState }

func (FeatureIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureIamMember)(nil)).Elem()
}

func (o FeatureIamMemberMapOutput) ToFeatureIamMemberMapOutput() FeatureIamMemberMapOutput {
	return o
}

func (o FeatureIamMemberMapOutput) ToFeatureIamMemberMapOutputWithContext(ctx context.Context) FeatureIamMemberMapOutput {
	return o
}

func (o FeatureIamMemberMapOutput) MapIndex(k pulumi.StringInput) FeatureIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FeatureIamMember {
		return vs[0].(map[string]*FeatureIamMember)[vs[1].(string)]
	}).(FeatureIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureIamMemberInput)(nil)).Elem(), &FeatureIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureIamMemberArrayInput)(nil)).Elem(), FeatureIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureIamMemberMapInput)(nil)).Elem(), FeatureIamMemberMap{})
	pulumi.RegisterOutputType(FeatureIamMemberOutput{})
	pulumi.RegisterOutputType(FeatureIamMemberArrayOutput{})
	pulumi.RegisterOutputType(FeatureIamMemberMapOutput{})
}

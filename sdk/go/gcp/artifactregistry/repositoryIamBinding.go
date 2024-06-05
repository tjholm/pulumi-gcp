// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactregistry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Artifact Registry Repository. Each of these resources serves a different use case:
//
// * `artifactregistry.RepositoryIamPolicy`: Authoritative. Sets the IAM policy for the repository and replaces any existing policy already attached.
// * `artifactregistry.RepositoryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the repository are preserved.
// * `artifactregistry.RepositoryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the repository are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `artifactregistry.RepositoryIamPolicy`: Retrieves the IAM policy for the repository
//
// > **Note:** `artifactregistry.RepositoryIamPolicy` **cannot** be used in conjunction with `artifactregistry.RepositoryIamBinding` and `artifactregistry.RepositoryIamMember` or they will fight over what your policy should be.
//
// > **Note:** `artifactregistry.RepositoryIamBinding` resources **can be** used in conjunction with `artifactregistry.RepositoryIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## artifactregistry.RepositoryIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/artifactregistry"
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
//						Role: "roles/artifactregistry.reader",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = artifactregistry.NewRepositoryIamPolicy(ctx, "policy", &artifactregistry.RepositoryIamPolicyArgs{
//				Project:    pulumi.Any(my_repo.Project),
//				Location:   pulumi.Any(my_repo.Location),
//				Repository: pulumi.Any(my_repo.Name),
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
// ## artifactregistry.RepositoryIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/artifactregistry"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactregistry.NewRepositoryIamBinding(ctx, "binding", &artifactregistry.RepositoryIamBindingArgs{
//				Project:    pulumi.Any(my_repo.Project),
//				Location:   pulumi.Any(my_repo.Location),
//				Repository: pulumi.Any(my_repo.Name),
//				Role:       pulumi.String("roles/artifactregistry.reader"),
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
// ## artifactregistry.RepositoryIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/artifactregistry"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactregistry.NewRepositoryIamMember(ctx, "member", &artifactregistry.RepositoryIamMemberArgs{
//				Project:    pulumi.Any(my_repo.Project),
//				Location:   pulumi.Any(my_repo.Location),
//				Repository: pulumi.Any(my_repo.Name),
//				Role:       pulumi.String("roles/artifactregistry.reader"),
//				Member:     pulumi.String("user:jane@example.com"),
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
// ## artifactregistry.RepositoryIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/artifactregistry"
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
//						Role: "roles/artifactregistry.reader",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = artifactregistry.NewRepositoryIamPolicy(ctx, "policy", &artifactregistry.RepositoryIamPolicyArgs{
//				Project:    pulumi.Any(my_repo.Project),
//				Location:   pulumi.Any(my_repo.Location),
//				Repository: pulumi.Any(my_repo.Name),
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
// ## artifactregistry.RepositoryIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/artifactregistry"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactregistry.NewRepositoryIamBinding(ctx, "binding", &artifactregistry.RepositoryIamBindingArgs{
//				Project:    pulumi.Any(my_repo.Project),
//				Location:   pulumi.Any(my_repo.Location),
//				Repository: pulumi.Any(my_repo.Name),
//				Role:       pulumi.String("roles/artifactregistry.reader"),
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
// ## artifactregistry.RepositoryIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/artifactregistry"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactregistry.NewRepositoryIamMember(ctx, "member", &artifactregistry.RepositoryIamMemberArgs{
//				Project:    pulumi.Any(my_repo.Project),
//				Location:   pulumi.Any(my_repo.Location),
//				Repository: pulumi.Any(my_repo.Name),
//				Role:       pulumi.String("roles/artifactregistry.reader"),
//				Member:     pulumi.String("user:jane@example.com"),
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
// * projects/{{project}}/locations/{{location}}/repositories/{{repository}}
//
// * {{project}}/{{location}}/{{repository}}
//
// * {{location}}/{{repository}}
//
// * {{repository}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Artifact Registry repository IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding editor "projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/artifactregistry.reader user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding editor "projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/artifactregistry.reader"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding editor projects/{{project}}/locations/{{location}}/repositories/{{repository}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RepositoryIamBinding struct {
	pulumi.CustomResourceState

	Condition RepositoryIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the location this repository is located in.
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
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The role that should be applied. Only one
	// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRepositoryIamBinding registers a new resource with the given unique name, arguments, and options.
func NewRepositoryIamBinding(ctx *pulumi.Context,
	name string, args *RepositoryIamBindingArgs, opts ...pulumi.ResourceOption) (*RepositoryIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryIamBinding
	err := ctx.RegisterResource("gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryIamBinding gets an existing RepositoryIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryIamBindingState, opts ...pulumi.ResourceOption) (*RepositoryIamBinding, error) {
	var resource RepositoryIamBinding
	err := ctx.ReadResource("gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryIamBinding resources.
type repositoryIamBindingState struct {
	Condition *RepositoryIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the location this repository is located in.
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
	// Used to find the parent resource to bind the IAM policy to
	Repository *string `pulumi:"repository"`
	// The role that should be applied. Only one
	// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type RepositoryIamBindingState struct {
	Condition RepositoryIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the location this repository is located in.
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
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (RepositoryIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamBindingState)(nil)).Elem()
}

type repositoryIamBindingArgs struct {
	Condition *RepositoryIamBindingCondition `pulumi:"condition"`
	// The name of the location this repository is located in.
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
	// Used to find the parent resource to bind the IAM policy to
	Repository string `pulumi:"repository"`
	// The role that should be applied. Only one
	// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a RepositoryIamBinding resource.
type RepositoryIamBindingArgs struct {
	Condition RepositoryIamBindingConditionPtrInput
	// The name of the location this repository is located in.
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
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringInput
	// The role that should be applied. Only one
	// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (RepositoryIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamBindingArgs)(nil)).Elem()
}

type RepositoryIamBindingInput interface {
	pulumi.Input

	ToRepositoryIamBindingOutput() RepositoryIamBindingOutput
	ToRepositoryIamBindingOutputWithContext(ctx context.Context) RepositoryIamBindingOutput
}

func (*RepositoryIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamBinding)(nil)).Elem()
}

func (i *RepositoryIamBinding) ToRepositoryIamBindingOutput() RepositoryIamBindingOutput {
	return i.ToRepositoryIamBindingOutputWithContext(context.Background())
}

func (i *RepositoryIamBinding) ToRepositoryIamBindingOutputWithContext(ctx context.Context) RepositoryIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamBindingOutput)
}

// RepositoryIamBindingArrayInput is an input type that accepts RepositoryIamBindingArray and RepositoryIamBindingArrayOutput values.
// You can construct a concrete instance of `RepositoryIamBindingArrayInput` via:
//
//	RepositoryIamBindingArray{ RepositoryIamBindingArgs{...} }
type RepositoryIamBindingArrayInput interface {
	pulumi.Input

	ToRepositoryIamBindingArrayOutput() RepositoryIamBindingArrayOutput
	ToRepositoryIamBindingArrayOutputWithContext(context.Context) RepositoryIamBindingArrayOutput
}

type RepositoryIamBindingArray []RepositoryIamBindingInput

func (RepositoryIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryIamBinding)(nil)).Elem()
}

func (i RepositoryIamBindingArray) ToRepositoryIamBindingArrayOutput() RepositoryIamBindingArrayOutput {
	return i.ToRepositoryIamBindingArrayOutputWithContext(context.Background())
}

func (i RepositoryIamBindingArray) ToRepositoryIamBindingArrayOutputWithContext(ctx context.Context) RepositoryIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamBindingArrayOutput)
}

// RepositoryIamBindingMapInput is an input type that accepts RepositoryIamBindingMap and RepositoryIamBindingMapOutput values.
// You can construct a concrete instance of `RepositoryIamBindingMapInput` via:
//
//	RepositoryIamBindingMap{ "key": RepositoryIamBindingArgs{...} }
type RepositoryIamBindingMapInput interface {
	pulumi.Input

	ToRepositoryIamBindingMapOutput() RepositoryIamBindingMapOutput
	ToRepositoryIamBindingMapOutputWithContext(context.Context) RepositoryIamBindingMapOutput
}

type RepositoryIamBindingMap map[string]RepositoryIamBindingInput

func (RepositoryIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryIamBinding)(nil)).Elem()
}

func (i RepositoryIamBindingMap) ToRepositoryIamBindingMapOutput() RepositoryIamBindingMapOutput {
	return i.ToRepositoryIamBindingMapOutputWithContext(context.Background())
}

func (i RepositoryIamBindingMap) ToRepositoryIamBindingMapOutputWithContext(ctx context.Context) RepositoryIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamBindingMapOutput)
}

type RepositoryIamBindingOutput struct{ *pulumi.OutputState }

func (RepositoryIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamBinding)(nil)).Elem()
}

func (o RepositoryIamBindingOutput) ToRepositoryIamBindingOutput() RepositoryIamBindingOutput {
	return o
}

func (o RepositoryIamBindingOutput) ToRepositoryIamBindingOutputWithContext(ctx context.Context) RepositoryIamBindingOutput {
	return o
}

func (o RepositoryIamBindingOutput) Condition() RepositoryIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *RepositoryIamBinding) RepositoryIamBindingConditionPtrOutput { return v.Condition }).(RepositoryIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o RepositoryIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the location this repository is located in.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o RepositoryIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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
func (o RepositoryIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o RepositoryIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o RepositoryIamBindingOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamBinding) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o RepositoryIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type RepositoryIamBindingArrayOutput struct{ *pulumi.OutputState }

func (RepositoryIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryIamBinding)(nil)).Elem()
}

func (o RepositoryIamBindingArrayOutput) ToRepositoryIamBindingArrayOutput() RepositoryIamBindingArrayOutput {
	return o
}

func (o RepositoryIamBindingArrayOutput) ToRepositoryIamBindingArrayOutputWithContext(ctx context.Context) RepositoryIamBindingArrayOutput {
	return o
}

func (o RepositoryIamBindingArrayOutput) Index(i pulumi.IntInput) RepositoryIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryIamBinding {
		return vs[0].([]*RepositoryIamBinding)[vs[1].(int)]
	}).(RepositoryIamBindingOutput)
}

type RepositoryIamBindingMapOutput struct{ *pulumi.OutputState }

func (RepositoryIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryIamBinding)(nil)).Elem()
}

func (o RepositoryIamBindingMapOutput) ToRepositoryIamBindingMapOutput() RepositoryIamBindingMapOutput {
	return o
}

func (o RepositoryIamBindingMapOutput) ToRepositoryIamBindingMapOutputWithContext(ctx context.Context) RepositoryIamBindingMapOutput {
	return o
}

func (o RepositoryIamBindingMapOutput) MapIndex(k pulumi.StringInput) RepositoryIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryIamBinding {
		return vs[0].(map[string]*RepositoryIamBinding)[vs[1].(string)]
	}).(RepositoryIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryIamBindingInput)(nil)).Elem(), &RepositoryIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryIamBindingArrayInput)(nil)).Elem(), RepositoryIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryIamBindingMapInput)(nil)).Elem(), RepositoryIamBindingMap{})
	pulumi.RegisterOutputType(RepositoryIamBindingOutput{})
	pulumi.RegisterOutputType(RepositoryIamBindingArrayOutput{})
	pulumi.RegisterOutputType(RepositoryIamBindingMapOutput{})
}

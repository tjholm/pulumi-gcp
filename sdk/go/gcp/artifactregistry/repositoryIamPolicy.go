// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactregistry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
// ## google\_artifact\_registry\_repository\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/artifactregistry"
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
//				Project:    pulumi.Any(google_artifact_registry_repository.MyRepo.Project),
//				Location:   pulumi.Any(google_artifact_registry_repository.MyRepo.Location),
//				Repository: pulumi.Any(google_artifact_registry_repository.MyRepo.Name),
//				PolicyData: *pulumi.String(admin.PolicyData),
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
// ## google\_artifact\_registry\_repository\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/artifactregistry"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactregistry.NewRepositoryIamBinding(ctx, "binding", &artifactregistry.RepositoryIamBindingArgs{
//				Project:    pulumi.Any(google_artifact_registry_repository.MyRepo.Project),
//				Location:   pulumi.Any(google_artifact_registry_repository.MyRepo.Location),
//				Repository: pulumi.Any(google_artifact_registry_repository.MyRepo.Name),
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
// ## google\_artifact\_registry\_repository\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/artifactregistry"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactregistry.NewRepositoryIamMember(ctx, "member", &artifactregistry.RepositoryIamMemberArgs{
//				Project:    pulumi.Any(google_artifact_registry_repository.MyRepo.Project),
//				Location:   pulumi.Any(google_artifact_registry_repository.MyRepo.Location),
//				Repository: pulumi.Any(google_artifact_registry_repository.MyRepo.Name),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/repositories/{{repository}} * {{project}}/{{location}}/{{repository}} * {{location}}/{{repository}} * {{repository}} Any variables not passed in the import command will be taken from the provider configuration. Artifact Registry repository IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy editor "projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/artifactregistry.reader user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy editor "projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/artifactregistry.reader"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy editor projects/{{project}}/locations/{{location}}/repositories/{{repository}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RepositoryIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringOutput `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringOutput `pulumi:"repository"`
}

// NewRepositoryIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewRepositoryIamPolicy(ctx *pulumi.Context,
	name string, args *RepositoryIamPolicyArgs, opts ...pulumi.ResourceOption) (*RepositoryIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryIamPolicy
	err := ctx.RegisterResource("gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryIamPolicy gets an existing RepositoryIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryIamPolicyState, opts ...pulumi.ResourceOption) (*RepositoryIamPolicy, error) {
	var resource RepositoryIamPolicy
	err := ctx.ReadResource("gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryIamPolicy resources.
type repositoryIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository *string `pulumi:"repository"`
}

type RepositoryIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringPtrInput
}

func (RepositoryIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamPolicyState)(nil)).Elem()
}

type repositoryIamPolicyArgs struct {
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryIamPolicy resource.
type RepositoryIamPolicyArgs struct {
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringInput
}

func (RepositoryIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamPolicyArgs)(nil)).Elem()
}

type RepositoryIamPolicyInput interface {
	pulumi.Input

	ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput
	ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput
}

func (*RepositoryIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamPolicy)(nil)).Elem()
}

func (i *RepositoryIamPolicy) ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput {
	return i.ToRepositoryIamPolicyOutputWithContext(context.Background())
}

func (i *RepositoryIamPolicy) ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyOutput)
}

func (i *RepositoryIamPolicy) ToOutput(ctx context.Context) pulumix.Output[*RepositoryIamPolicy] {
	return pulumix.Output[*RepositoryIamPolicy]{
		OutputState: i.ToRepositoryIamPolicyOutputWithContext(ctx).OutputState,
	}
}

// RepositoryIamPolicyArrayInput is an input type that accepts RepositoryIamPolicyArray and RepositoryIamPolicyArrayOutput values.
// You can construct a concrete instance of `RepositoryIamPolicyArrayInput` via:
//
//	RepositoryIamPolicyArray{ RepositoryIamPolicyArgs{...} }
type RepositoryIamPolicyArrayInput interface {
	pulumi.Input

	ToRepositoryIamPolicyArrayOutput() RepositoryIamPolicyArrayOutput
	ToRepositoryIamPolicyArrayOutputWithContext(context.Context) RepositoryIamPolicyArrayOutput
}

type RepositoryIamPolicyArray []RepositoryIamPolicyInput

func (RepositoryIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryIamPolicy)(nil)).Elem()
}

func (i RepositoryIamPolicyArray) ToRepositoryIamPolicyArrayOutput() RepositoryIamPolicyArrayOutput {
	return i.ToRepositoryIamPolicyArrayOutputWithContext(context.Background())
}

func (i RepositoryIamPolicyArray) ToRepositoryIamPolicyArrayOutputWithContext(ctx context.Context) RepositoryIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyArrayOutput)
}

func (i RepositoryIamPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*RepositoryIamPolicy] {
	return pulumix.Output[[]*RepositoryIamPolicy]{
		OutputState: i.ToRepositoryIamPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// RepositoryIamPolicyMapInput is an input type that accepts RepositoryIamPolicyMap and RepositoryIamPolicyMapOutput values.
// You can construct a concrete instance of `RepositoryIamPolicyMapInput` via:
//
//	RepositoryIamPolicyMap{ "key": RepositoryIamPolicyArgs{...} }
type RepositoryIamPolicyMapInput interface {
	pulumi.Input

	ToRepositoryIamPolicyMapOutput() RepositoryIamPolicyMapOutput
	ToRepositoryIamPolicyMapOutputWithContext(context.Context) RepositoryIamPolicyMapOutput
}

type RepositoryIamPolicyMap map[string]RepositoryIamPolicyInput

func (RepositoryIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryIamPolicy)(nil)).Elem()
}

func (i RepositoryIamPolicyMap) ToRepositoryIamPolicyMapOutput() RepositoryIamPolicyMapOutput {
	return i.ToRepositoryIamPolicyMapOutputWithContext(context.Background())
}

func (i RepositoryIamPolicyMap) ToRepositoryIamPolicyMapOutputWithContext(ctx context.Context) RepositoryIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyMapOutput)
}

func (i RepositoryIamPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RepositoryIamPolicy] {
	return pulumix.Output[map[string]*RepositoryIamPolicy]{
		OutputState: i.ToRepositoryIamPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type RepositoryIamPolicyOutput struct{ *pulumi.OutputState }

func (RepositoryIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamPolicy)(nil)).Elem()
}

func (o RepositoryIamPolicyOutput) ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput {
	return o
}

func (o RepositoryIamPolicyOutput) ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput {
	return o
}

func (o RepositoryIamPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*RepositoryIamPolicy] {
	return pulumix.Output[*RepositoryIamPolicy]{
		OutputState: o.OutputState,
	}
}

// (Computed) The etag of the IAM policy.
func (o RepositoryIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the location this repository is located in.
// Used to find the parent resource to bind the IAM policy to
func (o RepositoryIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o RepositoryIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
//   - **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
//   - **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
//   - **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
func (o RepositoryIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o RepositoryIamPolicyOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

type RepositoryIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (RepositoryIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryIamPolicy)(nil)).Elem()
}

func (o RepositoryIamPolicyArrayOutput) ToRepositoryIamPolicyArrayOutput() RepositoryIamPolicyArrayOutput {
	return o
}

func (o RepositoryIamPolicyArrayOutput) ToRepositoryIamPolicyArrayOutputWithContext(ctx context.Context) RepositoryIamPolicyArrayOutput {
	return o
}

func (o RepositoryIamPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RepositoryIamPolicy] {
	return pulumix.Output[[]*RepositoryIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o RepositoryIamPolicyArrayOutput) Index(i pulumi.IntInput) RepositoryIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryIamPolicy {
		return vs[0].([]*RepositoryIamPolicy)[vs[1].(int)]
	}).(RepositoryIamPolicyOutput)
}

type RepositoryIamPolicyMapOutput struct{ *pulumi.OutputState }

func (RepositoryIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryIamPolicy)(nil)).Elem()
}

func (o RepositoryIamPolicyMapOutput) ToRepositoryIamPolicyMapOutput() RepositoryIamPolicyMapOutput {
	return o
}

func (o RepositoryIamPolicyMapOutput) ToRepositoryIamPolicyMapOutputWithContext(ctx context.Context) RepositoryIamPolicyMapOutput {
	return o
}

func (o RepositoryIamPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RepositoryIamPolicy] {
	return pulumix.Output[map[string]*RepositoryIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o RepositoryIamPolicyMapOutput) MapIndex(k pulumi.StringInput) RepositoryIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryIamPolicy {
		return vs[0].(map[string]*RepositoryIamPolicy)[vs[1].(string)]
	}).(RepositoryIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryIamPolicyInput)(nil)).Elem(), &RepositoryIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryIamPolicyArrayInput)(nil)).Elem(), RepositoryIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryIamPolicyMapInput)(nil)).Elem(), RepositoryIamPolicyMap{})
	pulumi.RegisterOutputType(RepositoryIamPolicyOutput{})
	pulumi.RegisterOutputType(RepositoryIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(RepositoryIamPolicyMapOutput{})
}

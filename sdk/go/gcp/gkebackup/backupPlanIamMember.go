// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gkebackup

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Backup for GKE BackupPlan. Each of these resources serves a different use case:
//
// * `gkebackup.BackupPlanIamPolicy`: Authoritative. Sets the IAM policy for the backupplan and replaces any existing policy already attached.
// * `gkebackup.BackupPlanIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the backupplan are preserved.
// * `gkebackup.BackupPlanIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the backupplan are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `gkebackup.BackupPlanIamPolicy`: Retrieves the IAM policy for the backupplan
//
// > **Note:** `gkebackup.BackupPlanIamPolicy` **cannot** be used in conjunction with `gkebackup.BackupPlanIamBinding` and `gkebackup.BackupPlanIamMember` or they will fight over what your policy should be.
//
// > **Note:** `gkebackup.BackupPlanIamBinding` resources **can be** used in conjunction with `gkebackup.BackupPlanIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_gke\_backup\_backup\_plan\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkebackup"
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
//			_, err = gkebackup.NewBackupPlanIamPolicy(ctx, "policy", &gkebackup.BackupPlanIamPolicyArgs{
//				Project:    pulumi.Any(basic.Project),
//				Location:   pulumi.Any(basic.Location),
//				Name:       pulumi.Any(basic.Name),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_gke\_backup\_backup\_plan\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkebackup"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkebackup.NewBackupPlanIamBinding(ctx, "binding", &gkebackup.BackupPlanIamBindingArgs{
//				Project:  pulumi.Any(basic.Project),
//				Location: pulumi.Any(basic.Location),
//				Name:     pulumi.Any(basic.Name),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_gke\_backup\_backup\_plan\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkebackup"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkebackup.NewBackupPlanIamMember(ctx, "member", &gkebackup.BackupPlanIamMemberArgs{
//				Project:  pulumi.Any(basic.Project),
//				Location: pulumi.Any(basic.Location),
//				Name:     pulumi.Any(basic.Name),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_gke\_backup\_backup\_plan\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkebackup"
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
//			_, err = gkebackup.NewBackupPlanIamPolicy(ctx, "policy", &gkebackup.BackupPlanIamPolicyArgs{
//				Project:    pulumi.Any(basic.Project),
//				Location:   pulumi.Any(basic.Location),
//				Name:       pulumi.Any(basic.Name),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_gke\_backup\_backup\_plan\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkebackup"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkebackup.NewBackupPlanIamBinding(ctx, "binding", &gkebackup.BackupPlanIamBindingArgs{
//				Project:  pulumi.Any(basic.Project),
//				Location: pulumi.Any(basic.Location),
//				Name:     pulumi.Any(basic.Name),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_gke\_backup\_backup\_plan\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkebackup"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkebackup.NewBackupPlanIamMember(ctx, "member", &gkebackup.BackupPlanIamMemberArgs{
//				Project:  pulumi.Any(basic.Project),
//				Location: pulumi.Any(basic.Location),
//				Name:     pulumi.Any(basic.Name),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/locations/{{location}}/backupPlans/{{name}}
//
// * {{project}}/{{location}}/{{name}}
//
// * {{location}}/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Backup for GKE backupplan IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:gkebackup/backupPlanIamMember:BackupPlanIamMember editor "projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:gkebackup/backupPlanIamMember:BackupPlanIamMember editor "projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:gkebackup/backupPlanIamMember:BackupPlanIamMember editor projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type BackupPlanIamMember struct {
	pulumi.CustomResourceState

	Condition BackupPlanIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The region of the Backup Plan.
	// Used to find the parent resource to bind the IAM policy to
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
	// `gkebackup.BackupPlanIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewBackupPlanIamMember registers a new resource with the given unique name, arguments, and options.
func NewBackupPlanIamMember(ctx *pulumi.Context,
	name string, args *BackupPlanIamMemberArgs, opts ...pulumi.ResourceOption) (*BackupPlanIamMember, error) {
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
	var resource BackupPlanIamMember
	err := ctx.RegisterResource("gcp:gkebackup/backupPlanIamMember:BackupPlanIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPlanIamMember gets an existing BackupPlanIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPlanIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPlanIamMemberState, opts ...pulumi.ResourceOption) (*BackupPlanIamMember, error) {
	var resource BackupPlanIamMember
	err := ctx.ReadResource("gcp:gkebackup/backupPlanIamMember:BackupPlanIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPlanIamMember resources.
type backupPlanIamMemberState struct {
	Condition *BackupPlanIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The region of the Backup Plan.
	// Used to find the parent resource to bind the IAM policy to
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
	// `gkebackup.BackupPlanIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type BackupPlanIamMemberState struct {
	Condition BackupPlanIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The region of the Backup Plan.
	// Used to find the parent resource to bind the IAM policy to
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
	// `gkebackup.BackupPlanIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (BackupPlanIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPlanIamMemberState)(nil)).Elem()
}

type backupPlanIamMemberArgs struct {
	Condition *BackupPlanIamMemberCondition `pulumi:"condition"`
	// The region of the Backup Plan.
	// Used to find the parent resource to bind the IAM policy to
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
	// `gkebackup.BackupPlanIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a BackupPlanIamMember resource.
type BackupPlanIamMemberArgs struct {
	Condition BackupPlanIamMemberConditionPtrInput
	// The region of the Backup Plan.
	// Used to find the parent resource to bind the IAM policy to
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
	// `gkebackup.BackupPlanIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (BackupPlanIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPlanIamMemberArgs)(nil)).Elem()
}

type BackupPlanIamMemberInput interface {
	pulumi.Input

	ToBackupPlanIamMemberOutput() BackupPlanIamMemberOutput
	ToBackupPlanIamMemberOutputWithContext(ctx context.Context) BackupPlanIamMemberOutput
}

func (*BackupPlanIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPlanIamMember)(nil)).Elem()
}

func (i *BackupPlanIamMember) ToBackupPlanIamMemberOutput() BackupPlanIamMemberOutput {
	return i.ToBackupPlanIamMemberOutputWithContext(context.Background())
}

func (i *BackupPlanIamMember) ToBackupPlanIamMemberOutputWithContext(ctx context.Context) BackupPlanIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPlanIamMemberOutput)
}

// BackupPlanIamMemberArrayInput is an input type that accepts BackupPlanIamMemberArray and BackupPlanIamMemberArrayOutput values.
// You can construct a concrete instance of `BackupPlanIamMemberArrayInput` via:
//
//	BackupPlanIamMemberArray{ BackupPlanIamMemberArgs{...} }
type BackupPlanIamMemberArrayInput interface {
	pulumi.Input

	ToBackupPlanIamMemberArrayOutput() BackupPlanIamMemberArrayOutput
	ToBackupPlanIamMemberArrayOutputWithContext(context.Context) BackupPlanIamMemberArrayOutput
}

type BackupPlanIamMemberArray []BackupPlanIamMemberInput

func (BackupPlanIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPlanIamMember)(nil)).Elem()
}

func (i BackupPlanIamMemberArray) ToBackupPlanIamMemberArrayOutput() BackupPlanIamMemberArrayOutput {
	return i.ToBackupPlanIamMemberArrayOutputWithContext(context.Background())
}

func (i BackupPlanIamMemberArray) ToBackupPlanIamMemberArrayOutputWithContext(ctx context.Context) BackupPlanIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPlanIamMemberArrayOutput)
}

// BackupPlanIamMemberMapInput is an input type that accepts BackupPlanIamMemberMap and BackupPlanIamMemberMapOutput values.
// You can construct a concrete instance of `BackupPlanIamMemberMapInput` via:
//
//	BackupPlanIamMemberMap{ "key": BackupPlanIamMemberArgs{...} }
type BackupPlanIamMemberMapInput interface {
	pulumi.Input

	ToBackupPlanIamMemberMapOutput() BackupPlanIamMemberMapOutput
	ToBackupPlanIamMemberMapOutputWithContext(context.Context) BackupPlanIamMemberMapOutput
}

type BackupPlanIamMemberMap map[string]BackupPlanIamMemberInput

func (BackupPlanIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPlanIamMember)(nil)).Elem()
}

func (i BackupPlanIamMemberMap) ToBackupPlanIamMemberMapOutput() BackupPlanIamMemberMapOutput {
	return i.ToBackupPlanIamMemberMapOutputWithContext(context.Background())
}

func (i BackupPlanIamMemberMap) ToBackupPlanIamMemberMapOutputWithContext(ctx context.Context) BackupPlanIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPlanIamMemberMapOutput)
}

type BackupPlanIamMemberOutput struct{ *pulumi.OutputState }

func (BackupPlanIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPlanIamMember)(nil)).Elem()
}

func (o BackupPlanIamMemberOutput) ToBackupPlanIamMemberOutput() BackupPlanIamMemberOutput {
	return o
}

func (o BackupPlanIamMemberOutput) ToBackupPlanIamMemberOutputWithContext(ctx context.Context) BackupPlanIamMemberOutput {
	return o
}

func (o BackupPlanIamMemberOutput) Condition() BackupPlanIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *BackupPlanIamMember) BackupPlanIamMemberConditionPtrOutput { return v.Condition }).(BackupPlanIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o BackupPlanIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The region of the Backup Plan.
// Used to find the parent resource to bind the IAM policy to
func (o BackupPlanIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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
func (o BackupPlanIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o BackupPlanIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o BackupPlanIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `gkebackup.BackupPlanIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o BackupPlanIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type BackupPlanIamMemberArrayOutput struct{ *pulumi.OutputState }

func (BackupPlanIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPlanIamMember)(nil)).Elem()
}

func (o BackupPlanIamMemberArrayOutput) ToBackupPlanIamMemberArrayOutput() BackupPlanIamMemberArrayOutput {
	return o
}

func (o BackupPlanIamMemberArrayOutput) ToBackupPlanIamMemberArrayOutputWithContext(ctx context.Context) BackupPlanIamMemberArrayOutput {
	return o
}

func (o BackupPlanIamMemberArrayOutput) Index(i pulumi.IntInput) BackupPlanIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupPlanIamMember {
		return vs[0].([]*BackupPlanIamMember)[vs[1].(int)]
	}).(BackupPlanIamMemberOutput)
}

type BackupPlanIamMemberMapOutput struct{ *pulumi.OutputState }

func (BackupPlanIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPlanIamMember)(nil)).Elem()
}

func (o BackupPlanIamMemberMapOutput) ToBackupPlanIamMemberMapOutput() BackupPlanIamMemberMapOutput {
	return o
}

func (o BackupPlanIamMemberMapOutput) ToBackupPlanIamMemberMapOutputWithContext(ctx context.Context) BackupPlanIamMemberMapOutput {
	return o
}

func (o BackupPlanIamMemberMapOutput) MapIndex(k pulumi.StringInput) BackupPlanIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupPlanIamMember {
		return vs[0].(map[string]*BackupPlanIamMember)[vs[1].(string)]
	}).(BackupPlanIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPlanIamMemberInput)(nil)).Elem(), &BackupPlanIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPlanIamMemberArrayInput)(nil)).Elem(), BackupPlanIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPlanIamMemberMapInput)(nil)).Elem(), BackupPlanIamMemberMap{})
	pulumi.RegisterOutputType(BackupPlanIamMemberOutput{})
	pulumi.RegisterOutputType(BackupPlanIamMemberArrayOutput{})
	pulumi.RegisterOutputType(BackupPlanIamMemberMapOutput{})
}

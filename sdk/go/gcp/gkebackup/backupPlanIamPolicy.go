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
// ## gkebackup.BackupPlanIamPolicy
//
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
//
// ## gkebackup.BackupPlanIamBinding
//
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
//
// ## gkebackup.BackupPlanIamMember
//
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
//
// ## This resource supports User Project Overrides.
//
// -
//
// # IAM policy for Backup for GKE BackupPlan
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
// ## gkebackup.BackupPlanIamPolicy
//
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
//
// ## gkebackup.BackupPlanIamBinding
//
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
//
// ## gkebackup.BackupPlanIamMember
//
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
// $ pulumi import gcp:gkebackup/backupPlanIamPolicy:BackupPlanIamPolicy editor "projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:gkebackup/backupPlanIamPolicy:BackupPlanIamPolicy editor "projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:gkebackup/backupPlanIamPolicy:BackupPlanIamPolicy editor projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type BackupPlanIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The region of the Backup Plan.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringOutput `pulumi:"location"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewBackupPlanIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPlanIamPolicy(ctx *pulumi.Context,
	name string, args *BackupPlanIamPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPlanIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupPlanIamPolicy
	err := ctx.RegisterResource("gcp:gkebackup/backupPlanIamPolicy:BackupPlanIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPlanIamPolicy gets an existing BackupPlanIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPlanIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPlanIamPolicyState, opts ...pulumi.ResourceOption) (*BackupPlanIamPolicy, error) {
	var resource BackupPlanIamPolicy
	err := ctx.ReadResource("gcp:gkebackup/backupPlanIamPolicy:BackupPlanIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPlanIamPolicy resources.
type backupPlanIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The region of the Backup Plan.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type BackupPlanIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The region of the Backup Plan.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackupPlanIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPlanIamPolicyState)(nil)).Elem()
}

type backupPlanIamPolicyArgs struct {
	// The region of the Backup Plan.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a BackupPlanIamPolicy resource.
type BackupPlanIamPolicyArgs struct {
	// The region of the Backup Plan.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackupPlanIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPlanIamPolicyArgs)(nil)).Elem()
}

type BackupPlanIamPolicyInput interface {
	pulumi.Input

	ToBackupPlanIamPolicyOutput() BackupPlanIamPolicyOutput
	ToBackupPlanIamPolicyOutputWithContext(ctx context.Context) BackupPlanIamPolicyOutput
}

func (*BackupPlanIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPlanIamPolicy)(nil)).Elem()
}

func (i *BackupPlanIamPolicy) ToBackupPlanIamPolicyOutput() BackupPlanIamPolicyOutput {
	return i.ToBackupPlanIamPolicyOutputWithContext(context.Background())
}

func (i *BackupPlanIamPolicy) ToBackupPlanIamPolicyOutputWithContext(ctx context.Context) BackupPlanIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPlanIamPolicyOutput)
}

// BackupPlanIamPolicyArrayInput is an input type that accepts BackupPlanIamPolicyArray and BackupPlanIamPolicyArrayOutput values.
// You can construct a concrete instance of `BackupPlanIamPolicyArrayInput` via:
//
//	BackupPlanIamPolicyArray{ BackupPlanIamPolicyArgs{...} }
type BackupPlanIamPolicyArrayInput interface {
	pulumi.Input

	ToBackupPlanIamPolicyArrayOutput() BackupPlanIamPolicyArrayOutput
	ToBackupPlanIamPolicyArrayOutputWithContext(context.Context) BackupPlanIamPolicyArrayOutput
}

type BackupPlanIamPolicyArray []BackupPlanIamPolicyInput

func (BackupPlanIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPlanIamPolicy)(nil)).Elem()
}

func (i BackupPlanIamPolicyArray) ToBackupPlanIamPolicyArrayOutput() BackupPlanIamPolicyArrayOutput {
	return i.ToBackupPlanIamPolicyArrayOutputWithContext(context.Background())
}

func (i BackupPlanIamPolicyArray) ToBackupPlanIamPolicyArrayOutputWithContext(ctx context.Context) BackupPlanIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPlanIamPolicyArrayOutput)
}

// BackupPlanIamPolicyMapInput is an input type that accepts BackupPlanIamPolicyMap and BackupPlanIamPolicyMapOutput values.
// You can construct a concrete instance of `BackupPlanIamPolicyMapInput` via:
//
//	BackupPlanIamPolicyMap{ "key": BackupPlanIamPolicyArgs{...} }
type BackupPlanIamPolicyMapInput interface {
	pulumi.Input

	ToBackupPlanIamPolicyMapOutput() BackupPlanIamPolicyMapOutput
	ToBackupPlanIamPolicyMapOutputWithContext(context.Context) BackupPlanIamPolicyMapOutput
}

type BackupPlanIamPolicyMap map[string]BackupPlanIamPolicyInput

func (BackupPlanIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPlanIamPolicy)(nil)).Elem()
}

func (i BackupPlanIamPolicyMap) ToBackupPlanIamPolicyMapOutput() BackupPlanIamPolicyMapOutput {
	return i.ToBackupPlanIamPolicyMapOutputWithContext(context.Background())
}

func (i BackupPlanIamPolicyMap) ToBackupPlanIamPolicyMapOutputWithContext(ctx context.Context) BackupPlanIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPlanIamPolicyMapOutput)
}

type BackupPlanIamPolicyOutput struct{ *pulumi.OutputState }

func (BackupPlanIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPlanIamPolicy)(nil)).Elem()
}

func (o BackupPlanIamPolicyOutput) ToBackupPlanIamPolicyOutput() BackupPlanIamPolicyOutput {
	return o
}

func (o BackupPlanIamPolicyOutput) ToBackupPlanIamPolicyOutputWithContext(ctx context.Context) BackupPlanIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o BackupPlanIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The region of the Backup Plan.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o BackupPlanIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o BackupPlanIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o BackupPlanIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o BackupPlanIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlanIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type BackupPlanIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (BackupPlanIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPlanIamPolicy)(nil)).Elem()
}

func (o BackupPlanIamPolicyArrayOutput) ToBackupPlanIamPolicyArrayOutput() BackupPlanIamPolicyArrayOutput {
	return o
}

func (o BackupPlanIamPolicyArrayOutput) ToBackupPlanIamPolicyArrayOutputWithContext(ctx context.Context) BackupPlanIamPolicyArrayOutput {
	return o
}

func (o BackupPlanIamPolicyArrayOutput) Index(i pulumi.IntInput) BackupPlanIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupPlanIamPolicy {
		return vs[0].([]*BackupPlanIamPolicy)[vs[1].(int)]
	}).(BackupPlanIamPolicyOutput)
}

type BackupPlanIamPolicyMapOutput struct{ *pulumi.OutputState }

func (BackupPlanIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPlanIamPolicy)(nil)).Elem()
}

func (o BackupPlanIamPolicyMapOutput) ToBackupPlanIamPolicyMapOutput() BackupPlanIamPolicyMapOutput {
	return o
}

func (o BackupPlanIamPolicyMapOutput) ToBackupPlanIamPolicyMapOutputWithContext(ctx context.Context) BackupPlanIamPolicyMapOutput {
	return o
}

func (o BackupPlanIamPolicyMapOutput) MapIndex(k pulumi.StringInput) BackupPlanIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupPlanIamPolicy {
		return vs[0].(map[string]*BackupPlanIamPolicy)[vs[1].(string)]
	}).(BackupPlanIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPlanIamPolicyInput)(nil)).Elem(), &BackupPlanIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPlanIamPolicyArrayInput)(nil)).Elem(), BackupPlanIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPlanIamPolicyMapInput)(nil)).Elem(), BackupPlanIamPolicyMap{})
	pulumi.RegisterOutputType(BackupPlanIamPolicyOutput{})
	pulumi.RegisterOutputType(BackupPlanIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(BackupPlanIamPolicyMapOutput{})
}

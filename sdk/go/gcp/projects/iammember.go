// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
//
// * `projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
// * `projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
// * `projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
// * `projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
//
// > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.
//
// > **Note:** `projects.IAMBinding` resources **can be** used in conjunction with `projects.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:** The underlying API method `projects.setIamPolicy` has a lot of constraints which are documented [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy). In addition to these constraints,
//
//	IAM Conditions cannot be used with Basic Roles such as Owner. Violating these constraints will result in the API returning 400 error code so please review these if you encounter errors with this resource.
//
// ## google\_project\_iam\_policy
//
// !> **Be careful!** You can accidentally lock yourself out of your project
//
//	using this resource. Deleting a `projects.IAMPolicy` removes access
//	from anyone without organization-level access to the project. Proceed with caution.
//	It's not recommended to use `projects.IAMPolicy` with your provider project
//	to avoid locking yourself out, and it should generally only be used with projects
//	fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
//	applying the change.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/editor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
//				Project:    pulumi.String("your-project-id"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Condition: {
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//							Title:       "expires_after_2019_12_31",
//						},
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Role: "roles/compute.admin",
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
//				PolicyData: *pulumi.String(admin.PolicyData),
//				Project:    pulumi.String("your-project-id"),
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
// ## google\_project\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/editor"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
//				Condition: &projects.IAMBindingConditionArgs{
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
//					Title:       pulumi.String("expires_after_2019_12_31"),
//				},
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/container.admin"),
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
// ## google\_project\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
//				Member:  pulumi.String("user:jane@example.com"),
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/editor"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
//				Condition: &projects.IAMMemberConditionArgs{
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
//					Title:       pulumi.String("expires_after_2019_12_31"),
//				},
//				Member:  pulumi.String("user:jane@example.com"),
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/firebase.admin"),
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
// ## google\_project\_iam\_audit\_config
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMAuditConfig(ctx, "project", &projects.IAMAuditConfigArgs{
//				AuditLogConfigs: projects.IAMAuditConfigAuditLogConfigArray{
//					&projects.IAMAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("ADMIN_READ"),
//					},
//					&projects.IAMAuditConfigAuditLogConfigArgs{
//						ExemptedMembers: pulumi.StringArray{
//							pulumi.String("user:joebloggs@hashicorp.com"),
//						},
//						LogType: pulumi.String("DATA_READ"),
//					},
//				},
//				Project: pulumi.String("your-project-id"),
//				Service: pulumi.String("allServices"),
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
// ### Importing Audit Configs An audit config can be imported into a `google_project_iam_audit_config` resource using the resource's `project_id` and the `service`, e.g* `"{{project_id}} foo.googleapis.com"` An `import` block (Terraform v1.5.0 and later) can be used to import audit configstf import {
//
//	id = "{{project_id}} foo.googleapis.com"
//
//	to = google_project_iam_audit_config.default } The `pulumi import` command can also be used
//
// ```sh
//
//	$ pulumi import gcp:projects/iAMMember:IAMMember default "{{project_id}} foo.googleapis.com"
//
// ```
type IAMMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the project's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewIAMMember registers a new resource with the given unique name, arguments, and options.
func NewIAMMember(ctx *pulumi.Context,
	name string, args *IAMMemberArgs, opts ...pulumi.ResourceOption) (*IAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IAMMember
	err := ctx.RegisterResource("gcp:projects/iAMMember:IAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMMember gets an existing IAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMMemberState, opts ...pulumi.ResourceOption) (*IAMMember, error) {
	var resource IAMMember
	err := ctx.ReadResource("gcp:projects/iAMMember:IAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMMember resources.
type iammemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the project's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type IAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrInput
	// (Computed) The etag of the project's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (IAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iammemberState)(nil)).Elem()
}

type iammemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMMemberCondition `pulumi:"condition"`
	Member    string              `pulumi:"member"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project string `pulumi:"project"`
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a IAMMember resource.
type IAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrInput
	Member    pulumi.StringInput
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringInput
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (IAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iammemberArgs)(nil)).Elem()
}

type IAMMemberInput interface {
	pulumi.Input

	ToIAMMemberOutput() IAMMemberOutput
	ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput
}

func (*IAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMMember)(nil)).Elem()
}

func (i *IAMMember) ToIAMMemberOutput() IAMMemberOutput {
	return i.ToIAMMemberOutputWithContext(context.Background())
}

func (i *IAMMember) ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberOutput)
}

// IAMMemberArrayInput is an input type that accepts IAMMemberArray and IAMMemberArrayOutput values.
// You can construct a concrete instance of `IAMMemberArrayInput` via:
//
//	IAMMemberArray{ IAMMemberArgs{...} }
type IAMMemberArrayInput interface {
	pulumi.Input

	ToIAMMemberArrayOutput() IAMMemberArrayOutput
	ToIAMMemberArrayOutputWithContext(context.Context) IAMMemberArrayOutput
}

type IAMMemberArray []IAMMemberInput

func (IAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMMember)(nil)).Elem()
}

func (i IAMMemberArray) ToIAMMemberArrayOutput() IAMMemberArrayOutput {
	return i.ToIAMMemberArrayOutputWithContext(context.Background())
}

func (i IAMMemberArray) ToIAMMemberArrayOutputWithContext(ctx context.Context) IAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberArrayOutput)
}

// IAMMemberMapInput is an input type that accepts IAMMemberMap and IAMMemberMapOutput values.
// You can construct a concrete instance of `IAMMemberMapInput` via:
//
//	IAMMemberMap{ "key": IAMMemberArgs{...} }
type IAMMemberMapInput interface {
	pulumi.Input

	ToIAMMemberMapOutput() IAMMemberMapOutput
	ToIAMMemberMapOutputWithContext(context.Context) IAMMemberMapOutput
}

type IAMMemberMap map[string]IAMMemberInput

func (IAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMMember)(nil)).Elem()
}

func (i IAMMemberMap) ToIAMMemberMapOutput() IAMMemberMapOutput {
	return i.ToIAMMemberMapOutputWithContext(context.Background())
}

func (i IAMMemberMap) ToIAMMemberMapOutputWithContext(ctx context.Context) IAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberMapOutput)
}

type IAMMemberOutput struct{ *pulumi.OutputState }

func (IAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMMember)(nil)).Elem()
}

func (o IAMMemberOutput) ToIAMMemberOutput() IAMMemberOutput {
	return o
}

func (o IAMMemberOutput) ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o IAMMemberOutput) Condition() IAMMemberConditionPtrOutput {
	return o.ApplyT(func(v *IAMMember) IAMMemberConditionPtrOutput { return v.Condition }).(IAMMemberConditionPtrOutput)
}

// (Computed) The etag of the project's IAM policy.
func (o IAMMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IAMMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The project id of the target project. This is not
// inferred from the provider.
func (o IAMMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o IAMMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type IAMMemberArrayOutput struct{ *pulumi.OutputState }

func (IAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMMember)(nil)).Elem()
}

func (o IAMMemberArrayOutput) ToIAMMemberArrayOutput() IAMMemberArrayOutput {
	return o
}

func (o IAMMemberArrayOutput) ToIAMMemberArrayOutputWithContext(ctx context.Context) IAMMemberArrayOutput {
	return o
}

func (o IAMMemberArrayOutput) Index(i pulumi.IntInput) IAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IAMMember {
		return vs[0].([]*IAMMember)[vs[1].(int)]
	}).(IAMMemberOutput)
}

type IAMMemberMapOutput struct{ *pulumi.OutputState }

func (IAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMMember)(nil)).Elem()
}

func (o IAMMemberMapOutput) ToIAMMemberMapOutput() IAMMemberMapOutput {
	return o
}

func (o IAMMemberMapOutput) ToIAMMemberMapOutputWithContext(ctx context.Context) IAMMemberMapOutput {
	return o
}

func (o IAMMemberMapOutput) MapIndex(k pulumi.StringInput) IAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IAMMember {
		return vs[0].(map[string]*IAMMember)[vs[1].(string)]
	}).(IAMMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberInput)(nil)).Elem(), &IAMMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberArrayInput)(nil)).Elem(), IAMMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberMapInput)(nil)).Elem(), IAMMemberMap{})
	pulumi.RegisterOutputType(IAMMemberOutput{})
	pulumi.RegisterOutputType(IAMMemberArrayOutput{})
	pulumi.RegisterOutputType(IAMMemberMapOutput{})
}

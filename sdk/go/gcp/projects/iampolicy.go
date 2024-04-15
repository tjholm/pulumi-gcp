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
// <!--Start PulumiCodeChooser -->
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
// With IAM Conditions:
//
// <!--Start PulumiCodeChooser -->
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
//						Role: "roles/compute.admin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Condition: {
//							Title:       "expires_after_2019_12_31",
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
//				Project:    pulumi.String("your-project-id"),
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
// ## google\_project\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
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
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/editor"),
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
// With IAM Conditions:
//
// <!--Start PulumiCodeChooser -->
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
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/container.admin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &projects.IAMBindingConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## google\_project\_iam\_member
//
// <!--Start PulumiCodeChooser -->
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
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/editor"),
//				Member:  pulumi.String("user:jane@example.com"),
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
// With IAM Conditions:
//
// <!--Start PulumiCodeChooser -->
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
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/firebase.admin"),
//				Member:  pulumi.String("user:jane@example.com"),
//				Condition: &projects.IAMMemberConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## google\_project\_iam\_audit\_config
//
// <!--Start PulumiCodeChooser -->
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
//				Project: pulumi.String("your-project-id"),
//				Service: pulumi.String("allServices"),
//				AuditLogConfigs: projects.IAMAuditConfigAuditLogConfigArray{
//					&projects.IAMAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("ADMIN_READ"),
//					},
//					&projects.IAMAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("DATA_READ"),
//						ExemptedMembers: pulumi.StringArray{
//							pulumi.String("user:joebloggs@example.com"),
//						},
//					},
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
// <!--Start PulumiCodeChooser -->
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
// With IAM Conditions:
//
// <!--Start PulumiCodeChooser -->
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
//						Role: "roles/compute.admin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Condition: {
//							Title:       "expires_after_2019_12_31",
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
//				Project:    pulumi.String("your-project-id"),
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
// ## google\_project\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
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
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/editor"),
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
// With IAM Conditions:
//
// <!--Start PulumiCodeChooser -->
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
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/container.admin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &projects.IAMBindingConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## google\_project\_iam\_member
//
// <!--Start PulumiCodeChooser -->
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
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/editor"),
//				Member:  pulumi.String("user:jane@example.com"),
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
// With IAM Conditions:
//
// <!--Start PulumiCodeChooser -->
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
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/firebase.admin"),
//				Member:  pulumi.String("user:jane@example.com"),
//				Condition: &projects.IAMMemberConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## google\_project\_iam\_audit\_config
//
// <!--Start PulumiCodeChooser -->
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
//				Project: pulumi.String("your-project-id"),
//				Service: pulumi.String("allServices"),
//				AuditLogConfigs: projects.IAMAuditConfigAuditLogConfigArray{
//					&projects.IAMAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("ADMIN_READ"),
//					},
//					&projects.IAMAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("DATA_READ"),
//						ExemptedMembers: pulumi.StringArray{
//							pulumi.String("user:joebloggs@example.com"),
//						},
//					},
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
// ## Import
//
// ### Importing Audit Configs
//
// An audit config can be imported into a `google_project_iam_audit_config` resource using the resource's `project_id` and the `service`, e.g:
//
// * `"{{project_id}} foo.googleapis.com"`
//
// An `import` block (Terraform v1.5.0 and later) can be used to import audit configs:
//
// tf
//
// import {
//
//	id = "{{project_id}} foo.googleapis.com"
//
//	to = google_project_iam_audit_config.default
//
// }
//
// The `pulumi import` command can also be used:
//
// ```sh
// $ pulumi import gcp:projects/iAMPolicy:IAMPolicy default "{{project_id}} foo.googleapis.com"
// ```
type IAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the project's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	//
	// Changing this updates the policy.
	//
	// Deleting this removes all policies from the project, locking out users without
	// organization-level access.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewIAMPolicy(ctx *pulumi.Context,
	name string, args *IAMPolicyArgs, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IAMPolicy
	err := ctx.RegisterResource("gcp:projects/iAMPolicy:IAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMPolicy gets an existing IAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMPolicyState, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	var resource IAMPolicy
	err := ctx.ReadResource("gcp:projects/iAMPolicy:IAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMPolicy resources.
type iampolicyState struct {
	// (Computed) The etag of the project's IAM policy.
	Etag *string `pulumi:"etag"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	//
	// Changing this updates the policy.
	//
	// Deleting this removes all policies from the project, locking out users without
	// organization-level access.
	PolicyData *string `pulumi:"policyData"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project *string `pulumi:"project"`
}

type IAMPolicyState struct {
	// (Computed) The etag of the project's IAM policy.
	Etag pulumi.StringPtrInput
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	//
	// Changing this updates the policy.
	//
	// Deleting this removes all policies from the project, locking out users without
	// organization-level access.
	PolicyData pulumi.StringPtrInput
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringPtrInput
}

func (IAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyState)(nil)).Elem()
}

type iampolicyArgs struct {
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	//
	// Changing this updates the policy.
	//
	// Deleting this removes all policies from the project, locking out users without
	// organization-level access.
	PolicyData string `pulumi:"policyData"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a IAMPolicy resource.
type IAMPolicyArgs struct {
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	//
	// Changing this updates the policy.
	//
	// Deleting this removes all policies from the project, locking out users without
	// organization-level access.
	PolicyData pulumi.StringInput
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringInput
}

func (IAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyArgs)(nil)).Elem()
}

type IAMPolicyInput interface {
	pulumi.Input

	ToIAMPolicyOutput() IAMPolicyOutput
	ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput
}

func (*IAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMPolicy)(nil)).Elem()
}

func (i *IAMPolicy) ToIAMPolicyOutput() IAMPolicyOutput {
	return i.ToIAMPolicyOutputWithContext(context.Background())
}

func (i *IAMPolicy) ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyOutput)
}

// IAMPolicyArrayInput is an input type that accepts IAMPolicyArray and IAMPolicyArrayOutput values.
// You can construct a concrete instance of `IAMPolicyArrayInput` via:
//
//	IAMPolicyArray{ IAMPolicyArgs{...} }
type IAMPolicyArrayInput interface {
	pulumi.Input

	ToIAMPolicyArrayOutput() IAMPolicyArrayOutput
	ToIAMPolicyArrayOutputWithContext(context.Context) IAMPolicyArrayOutput
}

type IAMPolicyArray []IAMPolicyInput

func (IAMPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMPolicy)(nil)).Elem()
}

func (i IAMPolicyArray) ToIAMPolicyArrayOutput() IAMPolicyArrayOutput {
	return i.ToIAMPolicyArrayOutputWithContext(context.Background())
}

func (i IAMPolicyArray) ToIAMPolicyArrayOutputWithContext(ctx context.Context) IAMPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyArrayOutput)
}

// IAMPolicyMapInput is an input type that accepts IAMPolicyMap and IAMPolicyMapOutput values.
// You can construct a concrete instance of `IAMPolicyMapInput` via:
//
//	IAMPolicyMap{ "key": IAMPolicyArgs{...} }
type IAMPolicyMapInput interface {
	pulumi.Input

	ToIAMPolicyMapOutput() IAMPolicyMapOutput
	ToIAMPolicyMapOutputWithContext(context.Context) IAMPolicyMapOutput
}

type IAMPolicyMap map[string]IAMPolicyInput

func (IAMPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMPolicy)(nil)).Elem()
}

func (i IAMPolicyMap) ToIAMPolicyMapOutput() IAMPolicyMapOutput {
	return i.ToIAMPolicyMapOutputWithContext(context.Background())
}

func (i IAMPolicyMap) ToIAMPolicyMapOutputWithContext(ctx context.Context) IAMPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyMapOutput)
}

type IAMPolicyOutput struct{ *pulumi.OutputState }

func (IAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMPolicy)(nil)).Elem()
}

func (o IAMPolicyOutput) ToIAMPolicyOutput() IAMPolicyOutput {
	return o
}

func (o IAMPolicyOutput) ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput {
	return o
}

// (Computed) The etag of the project's IAM policy.
func (o IAMPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The `organizations.getIAMPolicy` data source that represents
// the IAM policy that will be applied to the project. The policy will be
// merged with any existing policy applied to the project.
//
// Changing this updates the policy.
//
// Deleting this removes all policies from the project, locking out users without
// organization-level access.
func (o IAMPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The project id of the target project. This is not
// inferred from the provider.
func (o IAMPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type IAMPolicyArrayOutput struct{ *pulumi.OutputState }

func (IAMPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMPolicy)(nil)).Elem()
}

func (o IAMPolicyArrayOutput) ToIAMPolicyArrayOutput() IAMPolicyArrayOutput {
	return o
}

func (o IAMPolicyArrayOutput) ToIAMPolicyArrayOutputWithContext(ctx context.Context) IAMPolicyArrayOutput {
	return o
}

func (o IAMPolicyArrayOutput) Index(i pulumi.IntInput) IAMPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IAMPolicy {
		return vs[0].([]*IAMPolicy)[vs[1].(int)]
	}).(IAMPolicyOutput)
}

type IAMPolicyMapOutput struct{ *pulumi.OutputState }

func (IAMPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMPolicy)(nil)).Elem()
}

func (o IAMPolicyMapOutput) ToIAMPolicyMapOutput() IAMPolicyMapOutput {
	return o
}

func (o IAMPolicyMapOutput) ToIAMPolicyMapOutputWithContext(ctx context.Context) IAMPolicyMapOutput {
	return o
}

func (o IAMPolicyMapOutput) MapIndex(k pulumi.StringInput) IAMPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IAMPolicy {
		return vs[0].(map[string]*IAMPolicy)[vs[1].(string)]
	}).(IAMPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAMPolicyInput)(nil)).Elem(), &IAMPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMPolicyArrayInput)(nil)).Elem(), IAMPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMPolicyMapInput)(nil)).Elem(), IAMPolicyMap{})
	pulumi.RegisterOutputType(IAMPolicyOutput{})
	pulumi.RegisterOutputType(IAMPolicyArrayOutput{})
	pulumi.RegisterOutputType(IAMPolicyMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Compute Engine MachineImage. Each of these resources serves a different use case:
//
// * `compute.MachineImageIamPolicy`: Authoritative. Sets the IAM policy for the machineimage and replaces any existing policy already attached.
// * `compute.MachineImageIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the machineimage are preserved.
// * `compute.MachineImageIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the machineimage are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `compute.MachineImageIamPolicy`: Retrieves the IAM policy for the machineimage
//
// > **Note:** `compute.MachineImageIamPolicy` **cannot** be used in conjunction with `compute.MachineImageIamBinding` and `compute.MachineImageIamMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.MachineImageIamBinding` resources **can be** used in conjunction with `compute.MachineImageIamMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
// ## google\_compute\_machine\_image\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
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
//						Role: "roles/compute.admin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewMachineImageIamPolicy(ctx, "policy", &compute.MachineImageIamPolicyArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				PolicyData:   pulumi.String(admin.PolicyData),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
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
//			_, err = compute.NewMachineImageIamPolicy(ctx, "policy", &compute.MachineImageIamPolicyArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				PolicyData:   pulumi.String(admin.PolicyData),
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
// ## google\_compute\_machine\_image\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewMachineImageIamBinding(ctx, "binding", &compute.MachineImageIamBindingArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				Role:         pulumi.String("roles/compute.admin"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewMachineImageIamBinding(ctx, "binding", &compute.MachineImageIamBindingArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				Role:         pulumi.String("roles/compute.admin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &compute.MachineImageIamBindingConditionArgs{
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
// ## google\_compute\_machine\_image\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewMachineImageIamMember(ctx, "member", &compute.MachineImageIamMemberArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				Role:         pulumi.String("roles/compute.admin"),
//				Member:       pulumi.String("user:jane@example.com"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewMachineImageIamMember(ctx, "member", &compute.MachineImageIamMemberArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				Role:         pulumi.String("roles/compute.admin"),
//				Member:       pulumi.String("user:jane@example.com"),
//				Condition: &compute.MachineImageIamMemberConditionArgs{
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
// ## google\_compute\_machine\_image\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
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
//						Role: "roles/compute.admin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewMachineImageIamPolicy(ctx, "policy", &compute.MachineImageIamPolicyArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				PolicyData:   pulumi.String(admin.PolicyData),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
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
//			_, err = compute.NewMachineImageIamPolicy(ctx, "policy", &compute.MachineImageIamPolicyArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				PolicyData:   pulumi.String(admin.PolicyData),
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
// ## google\_compute\_machine\_image\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewMachineImageIamBinding(ctx, "binding", &compute.MachineImageIamBindingArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				Role:         pulumi.String("roles/compute.admin"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewMachineImageIamBinding(ctx, "binding", &compute.MachineImageIamBindingArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				Role:         pulumi.String("roles/compute.admin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &compute.MachineImageIamBindingConditionArgs{
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
// ## google\_compute\_machine\_image\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewMachineImageIamMember(ctx, "member", &compute.MachineImageIamMemberArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				Role:         pulumi.String("roles/compute.admin"),
//				Member:       pulumi.String("user:jane@example.com"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewMachineImageIamMember(ctx, "member", &compute.MachineImageIamMemberArgs{
//				Project:      pulumi.Any(image.Project),
//				MachineImage: pulumi.Any(image.Name),
//				Role:         pulumi.String("roles/compute.admin"),
//				Member:       pulumi.String("user:jane@example.com"),
//				Condition: &compute.MachineImageIamMemberConditionArgs{
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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/global/machineImages/{{name}}
//
// * {{project}}/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Compute Engine machineimage IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:compute/machineImageIamMember:MachineImageIamMember editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:compute/machineImageIamMember:MachineImageIamMember editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:compute/machineImageIamMember:MachineImageIamMember editor projects/{{project}}/global/machineImages/{{machine_image}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type MachineImageIamMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition MachineImageIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringOutput `pulumi:"machineImage"`
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
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewMachineImageIamMember registers a new resource with the given unique name, arguments, and options.
func NewMachineImageIamMember(ctx *pulumi.Context,
	name string, args *MachineImageIamMemberArgs, opts ...pulumi.ResourceOption) (*MachineImageIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineImage == nil {
		return nil, errors.New("invalid value for required argument 'MachineImage'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MachineImageIamMember
	err := ctx.RegisterResource("gcp:compute/machineImageIamMember:MachineImageIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineImageIamMember gets an existing MachineImageIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineImageIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineImageIamMemberState, opts ...pulumi.ResourceOption) (*MachineImageIamMember, error) {
	var resource MachineImageIamMember
	err := ctx.ReadResource("gcp:compute/machineImageIamMember:MachineImageIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineImageIamMember resources.
type machineImageIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *MachineImageIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage *string `pulumi:"machineImage"`
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
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type MachineImageIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition MachineImageIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringPtrInput
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
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (MachineImageIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageIamMemberState)(nil)).Elem()
}

type machineImageIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *MachineImageIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage string `pulumi:"machineImage"`
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
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a MachineImageIamMember resource.
type MachineImageIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition MachineImageIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringInput
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
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (MachineImageIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageIamMemberArgs)(nil)).Elem()
}

type MachineImageIamMemberInput interface {
	pulumi.Input

	ToMachineImageIamMemberOutput() MachineImageIamMemberOutput
	ToMachineImageIamMemberOutputWithContext(ctx context.Context) MachineImageIamMemberOutput
}

func (*MachineImageIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineImageIamMember)(nil)).Elem()
}

func (i *MachineImageIamMember) ToMachineImageIamMemberOutput() MachineImageIamMemberOutput {
	return i.ToMachineImageIamMemberOutputWithContext(context.Background())
}

func (i *MachineImageIamMember) ToMachineImageIamMemberOutputWithContext(ctx context.Context) MachineImageIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageIamMemberOutput)
}

// MachineImageIamMemberArrayInput is an input type that accepts MachineImageIamMemberArray and MachineImageIamMemberArrayOutput values.
// You can construct a concrete instance of `MachineImageIamMemberArrayInput` via:
//
//	MachineImageIamMemberArray{ MachineImageIamMemberArgs{...} }
type MachineImageIamMemberArrayInput interface {
	pulumi.Input

	ToMachineImageIamMemberArrayOutput() MachineImageIamMemberArrayOutput
	ToMachineImageIamMemberArrayOutputWithContext(context.Context) MachineImageIamMemberArrayOutput
}

type MachineImageIamMemberArray []MachineImageIamMemberInput

func (MachineImageIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineImageIamMember)(nil)).Elem()
}

func (i MachineImageIamMemberArray) ToMachineImageIamMemberArrayOutput() MachineImageIamMemberArrayOutput {
	return i.ToMachineImageIamMemberArrayOutputWithContext(context.Background())
}

func (i MachineImageIamMemberArray) ToMachineImageIamMemberArrayOutputWithContext(ctx context.Context) MachineImageIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageIamMemberArrayOutput)
}

// MachineImageIamMemberMapInput is an input type that accepts MachineImageIamMemberMap and MachineImageIamMemberMapOutput values.
// You can construct a concrete instance of `MachineImageIamMemberMapInput` via:
//
//	MachineImageIamMemberMap{ "key": MachineImageIamMemberArgs{...} }
type MachineImageIamMemberMapInput interface {
	pulumi.Input

	ToMachineImageIamMemberMapOutput() MachineImageIamMemberMapOutput
	ToMachineImageIamMemberMapOutputWithContext(context.Context) MachineImageIamMemberMapOutput
}

type MachineImageIamMemberMap map[string]MachineImageIamMemberInput

func (MachineImageIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineImageIamMember)(nil)).Elem()
}

func (i MachineImageIamMemberMap) ToMachineImageIamMemberMapOutput() MachineImageIamMemberMapOutput {
	return i.ToMachineImageIamMemberMapOutputWithContext(context.Background())
}

func (i MachineImageIamMemberMap) ToMachineImageIamMemberMapOutputWithContext(ctx context.Context) MachineImageIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageIamMemberMapOutput)
}

type MachineImageIamMemberOutput struct{ *pulumi.OutputState }

func (MachineImageIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineImageIamMember)(nil)).Elem()
}

func (o MachineImageIamMemberOutput) ToMachineImageIamMemberOutput() MachineImageIamMemberOutput {
	return o
}

func (o MachineImageIamMemberOutput) ToMachineImageIamMemberOutputWithContext(ctx context.Context) MachineImageIamMemberOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o MachineImageIamMemberOutput) Condition() MachineImageIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *MachineImageIamMember) MachineImageIamMemberConditionPtrOutput { return v.Condition }).(MachineImageIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o MachineImageIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineImageIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o MachineImageIamMemberOutput) MachineImage() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineImageIamMember) pulumi.StringOutput { return v.MachineImage }).(pulumi.StringOutput)
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
func (o MachineImageIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineImageIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o MachineImageIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineImageIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o MachineImageIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineImageIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type MachineImageIamMemberArrayOutput struct{ *pulumi.OutputState }

func (MachineImageIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineImageIamMember)(nil)).Elem()
}

func (o MachineImageIamMemberArrayOutput) ToMachineImageIamMemberArrayOutput() MachineImageIamMemberArrayOutput {
	return o
}

func (o MachineImageIamMemberArrayOutput) ToMachineImageIamMemberArrayOutputWithContext(ctx context.Context) MachineImageIamMemberArrayOutput {
	return o
}

func (o MachineImageIamMemberArrayOutput) Index(i pulumi.IntInput) MachineImageIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MachineImageIamMember {
		return vs[0].([]*MachineImageIamMember)[vs[1].(int)]
	}).(MachineImageIamMemberOutput)
}

type MachineImageIamMemberMapOutput struct{ *pulumi.OutputState }

func (MachineImageIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineImageIamMember)(nil)).Elem()
}

func (o MachineImageIamMemberMapOutput) ToMachineImageIamMemberMapOutput() MachineImageIamMemberMapOutput {
	return o
}

func (o MachineImageIamMemberMapOutput) ToMachineImageIamMemberMapOutputWithContext(ctx context.Context) MachineImageIamMemberMapOutput {
	return o
}

func (o MachineImageIamMemberMapOutput) MapIndex(k pulumi.StringInput) MachineImageIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MachineImageIamMember {
		return vs[0].(map[string]*MachineImageIamMember)[vs[1].(string)]
	}).(MachineImageIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineImageIamMemberInput)(nil)).Elem(), &MachineImageIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineImageIamMemberArrayInput)(nil)).Elem(), MachineImageIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineImageIamMemberMapInput)(nil)).Elem(), MachineImageIamMemberMap{})
	pulumi.RegisterOutputType(MachineImageIamMemberOutput{})
	pulumi.RegisterOutputType(MachineImageIamMemberArrayOutput{})
	pulumi.RegisterOutputType(MachineImageIamMemberMapOutput{})
}

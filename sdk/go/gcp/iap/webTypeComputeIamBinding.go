// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeCompute. Each of these resources serves a different use case:
//
// * `iap.WebTypeComputeIamPolicy`: Authoritative. Sets the IAM policy for the webtypecompute and replaces any existing policy already attached.
// * `iap.WebTypeComputeIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypecompute are preserved.
// * `iap.WebTypeComputeIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypecompute are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `iap.WebTypeComputeIamPolicy`: Retrieves the IAM policy for the webtypecompute
//
// > **Note:** `iap.WebTypeComputeIamPolicy` **cannot** be used in conjunction with `iap.WebTypeComputeIamBinding` and `iap.WebTypeComputeIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebTypeComputeIamBinding` resources **can be** used in conjunction with `iap.WebTypeComputeIamMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## iap.WebTypeComputeIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
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
//						Role: "roles/iap.httpsResourceAccessor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iap.NewWebTypeComputeIamPolicy(ctx, "policy", &iap.WebTypeComputeIamPolicyArgs{
//				Project:    pulumi.Any(projectService.Project),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
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
//						Role: "roles/iap.httpsResourceAccessor",
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
//			_, err = iap.NewWebTypeComputeIamPolicy(ctx, "policy", &iap.WebTypeComputeIamPolicyArgs{
//				Project:    pulumi.Any(projectService.Project),
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
// ## iap.WebTypeComputeIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebTypeComputeIamBinding(ctx, "binding", &iap.WebTypeComputeIamBindingArgs{
//				Project: pulumi.Any(projectService.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebTypeComputeIamBinding(ctx, "binding", &iap.WebTypeComputeIamBindingArgs{
//				Project: pulumi.Any(projectService.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &iap.WebTypeComputeIamBindingConditionArgs{
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
// ## iap.WebTypeComputeIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebTypeComputeIamMember(ctx, "member", &iap.WebTypeComputeIamMemberArgs{
//				Project: pulumi.Any(projectService.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
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
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebTypeComputeIamMember(ctx, "member", &iap.WebTypeComputeIamMemberArgs{
//				Project: pulumi.Any(projectService.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
//				Member:  pulumi.String("user:jane@example.com"),
//				Condition: &iap.WebTypeComputeIamMemberConditionArgs{
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
//
// ## This resource supports User Project Overrides.
//
// -
//
// # IAM policy for Identity-Aware Proxy WebTypeCompute
// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeCompute. Each of these resources serves a different use case:
//
// * `iap.WebTypeComputeIamPolicy`: Authoritative. Sets the IAM policy for the webtypecompute and replaces any existing policy already attached.
// * `iap.WebTypeComputeIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypecompute are preserved.
// * `iap.WebTypeComputeIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypecompute are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `iap.WebTypeComputeIamPolicy`: Retrieves the IAM policy for the webtypecompute
//
// > **Note:** `iap.WebTypeComputeIamPolicy` **cannot** be used in conjunction with `iap.WebTypeComputeIamBinding` and `iap.WebTypeComputeIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebTypeComputeIamBinding` resources **can be** used in conjunction with `iap.WebTypeComputeIamMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## iap.WebTypeComputeIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
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
//						Role: "roles/iap.httpsResourceAccessor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iap.NewWebTypeComputeIamPolicy(ctx, "policy", &iap.WebTypeComputeIamPolicyArgs{
//				Project:    pulumi.Any(projectService.Project),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
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
//						Role: "roles/iap.httpsResourceAccessor",
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
//			_, err = iap.NewWebTypeComputeIamPolicy(ctx, "policy", &iap.WebTypeComputeIamPolicyArgs{
//				Project:    pulumi.Any(projectService.Project),
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
// ## iap.WebTypeComputeIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebTypeComputeIamBinding(ctx, "binding", &iap.WebTypeComputeIamBindingArgs{
//				Project: pulumi.Any(projectService.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebTypeComputeIamBinding(ctx, "binding", &iap.WebTypeComputeIamBindingArgs{
//				Project: pulumi.Any(projectService.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &iap.WebTypeComputeIamBindingConditionArgs{
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
// ## iap.WebTypeComputeIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebTypeComputeIamMember(ctx, "member", &iap.WebTypeComputeIamMemberArgs{
//				Project: pulumi.Any(projectService.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
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
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebTypeComputeIamMember(ctx, "member", &iap.WebTypeComputeIamMemberArgs{
//				Project: pulumi.Any(projectService.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
//				Member:  pulumi.String("user:jane@example.com"),
//				Condition: &iap.WebTypeComputeIamMemberConditionArgs{
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
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/iap_web/compute
//
// * {{project}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Identity-Aware Proxy webtypecompute IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:iap/webTypeComputeIamBinding:WebTypeComputeIamBinding editor "projects/{{project}}/iap_web/compute roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:iap/webTypeComputeIamBinding:WebTypeComputeIamBinding editor "projects/{{project}}/iap_web/compute roles/iap.httpsResourceAccessor"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:iap/webTypeComputeIamBinding:WebTypeComputeIamBinding editor projects/{{project}}/iap_web/compute
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebTypeComputeIamBinding struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeComputeIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
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
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewWebTypeComputeIamBinding registers a new resource with the given unique name, arguments, and options.
func NewWebTypeComputeIamBinding(ctx *pulumi.Context,
	name string, args *WebTypeComputeIamBindingArgs, opts ...pulumi.ResourceOption) (*WebTypeComputeIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebTypeComputeIamBinding
	err := ctx.RegisterResource("gcp:iap/webTypeComputeIamBinding:WebTypeComputeIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTypeComputeIamBinding gets an existing WebTypeComputeIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTypeComputeIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTypeComputeIamBindingState, opts ...pulumi.ResourceOption) (*WebTypeComputeIamBinding, error) {
	var resource WebTypeComputeIamBinding
	err := ctx.ReadResource("gcp:iap/webTypeComputeIamBinding:WebTypeComputeIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTypeComputeIamBinding resources.
type webTypeComputeIamBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebTypeComputeIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
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
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type WebTypeComputeIamBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeComputeIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
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
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (WebTypeComputeIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeComputeIamBindingState)(nil)).Elem()
}

type webTypeComputeIamBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebTypeComputeIamBindingCondition `pulumi:"condition"`
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
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a WebTypeComputeIamBinding resource.
type WebTypeComputeIamBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeComputeIamBindingConditionPtrInput
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
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (WebTypeComputeIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeComputeIamBindingArgs)(nil)).Elem()
}

type WebTypeComputeIamBindingInput interface {
	pulumi.Input

	ToWebTypeComputeIamBindingOutput() WebTypeComputeIamBindingOutput
	ToWebTypeComputeIamBindingOutputWithContext(ctx context.Context) WebTypeComputeIamBindingOutput
}

func (*WebTypeComputeIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTypeComputeIamBinding)(nil)).Elem()
}

func (i *WebTypeComputeIamBinding) ToWebTypeComputeIamBindingOutput() WebTypeComputeIamBindingOutput {
	return i.ToWebTypeComputeIamBindingOutputWithContext(context.Background())
}

func (i *WebTypeComputeIamBinding) ToWebTypeComputeIamBindingOutputWithContext(ctx context.Context) WebTypeComputeIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamBindingOutput)
}

// WebTypeComputeIamBindingArrayInput is an input type that accepts WebTypeComputeIamBindingArray and WebTypeComputeIamBindingArrayOutput values.
// You can construct a concrete instance of `WebTypeComputeIamBindingArrayInput` via:
//
//	WebTypeComputeIamBindingArray{ WebTypeComputeIamBindingArgs{...} }
type WebTypeComputeIamBindingArrayInput interface {
	pulumi.Input

	ToWebTypeComputeIamBindingArrayOutput() WebTypeComputeIamBindingArrayOutput
	ToWebTypeComputeIamBindingArrayOutputWithContext(context.Context) WebTypeComputeIamBindingArrayOutput
}

type WebTypeComputeIamBindingArray []WebTypeComputeIamBindingInput

func (WebTypeComputeIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebTypeComputeIamBinding)(nil)).Elem()
}

func (i WebTypeComputeIamBindingArray) ToWebTypeComputeIamBindingArrayOutput() WebTypeComputeIamBindingArrayOutput {
	return i.ToWebTypeComputeIamBindingArrayOutputWithContext(context.Background())
}

func (i WebTypeComputeIamBindingArray) ToWebTypeComputeIamBindingArrayOutputWithContext(ctx context.Context) WebTypeComputeIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamBindingArrayOutput)
}

// WebTypeComputeIamBindingMapInput is an input type that accepts WebTypeComputeIamBindingMap and WebTypeComputeIamBindingMapOutput values.
// You can construct a concrete instance of `WebTypeComputeIamBindingMapInput` via:
//
//	WebTypeComputeIamBindingMap{ "key": WebTypeComputeIamBindingArgs{...} }
type WebTypeComputeIamBindingMapInput interface {
	pulumi.Input

	ToWebTypeComputeIamBindingMapOutput() WebTypeComputeIamBindingMapOutput
	ToWebTypeComputeIamBindingMapOutputWithContext(context.Context) WebTypeComputeIamBindingMapOutput
}

type WebTypeComputeIamBindingMap map[string]WebTypeComputeIamBindingInput

func (WebTypeComputeIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebTypeComputeIamBinding)(nil)).Elem()
}

func (i WebTypeComputeIamBindingMap) ToWebTypeComputeIamBindingMapOutput() WebTypeComputeIamBindingMapOutput {
	return i.ToWebTypeComputeIamBindingMapOutputWithContext(context.Background())
}

func (i WebTypeComputeIamBindingMap) ToWebTypeComputeIamBindingMapOutputWithContext(ctx context.Context) WebTypeComputeIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamBindingMapOutput)
}

type WebTypeComputeIamBindingOutput struct{ *pulumi.OutputState }

func (WebTypeComputeIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTypeComputeIamBinding)(nil)).Elem()
}

func (o WebTypeComputeIamBindingOutput) ToWebTypeComputeIamBindingOutput() WebTypeComputeIamBindingOutput {
	return o
}

func (o WebTypeComputeIamBindingOutput) ToWebTypeComputeIamBindingOutputWithContext(ctx context.Context) WebTypeComputeIamBindingOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o WebTypeComputeIamBindingOutput) Condition() WebTypeComputeIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *WebTypeComputeIamBinding) WebTypeComputeIamBindingConditionPtrOutput { return v.Condition }).(WebTypeComputeIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o WebTypeComputeIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTypeComputeIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
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
func (o WebTypeComputeIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebTypeComputeIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o WebTypeComputeIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTypeComputeIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o WebTypeComputeIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTypeComputeIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type WebTypeComputeIamBindingArrayOutput struct{ *pulumi.OutputState }

func (WebTypeComputeIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebTypeComputeIamBinding)(nil)).Elem()
}

func (o WebTypeComputeIamBindingArrayOutput) ToWebTypeComputeIamBindingArrayOutput() WebTypeComputeIamBindingArrayOutput {
	return o
}

func (o WebTypeComputeIamBindingArrayOutput) ToWebTypeComputeIamBindingArrayOutputWithContext(ctx context.Context) WebTypeComputeIamBindingArrayOutput {
	return o
}

func (o WebTypeComputeIamBindingArrayOutput) Index(i pulumi.IntInput) WebTypeComputeIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebTypeComputeIamBinding {
		return vs[0].([]*WebTypeComputeIamBinding)[vs[1].(int)]
	}).(WebTypeComputeIamBindingOutput)
}

type WebTypeComputeIamBindingMapOutput struct{ *pulumi.OutputState }

func (WebTypeComputeIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebTypeComputeIamBinding)(nil)).Elem()
}

func (o WebTypeComputeIamBindingMapOutput) ToWebTypeComputeIamBindingMapOutput() WebTypeComputeIamBindingMapOutput {
	return o
}

func (o WebTypeComputeIamBindingMapOutput) ToWebTypeComputeIamBindingMapOutputWithContext(ctx context.Context) WebTypeComputeIamBindingMapOutput {
	return o
}

func (o WebTypeComputeIamBindingMapOutput) MapIndex(k pulumi.StringInput) WebTypeComputeIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebTypeComputeIamBinding {
		return vs[0].(map[string]*WebTypeComputeIamBinding)[vs[1].(string)]
	}).(WebTypeComputeIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebTypeComputeIamBindingInput)(nil)).Elem(), &WebTypeComputeIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebTypeComputeIamBindingArrayInput)(nil)).Elem(), WebTypeComputeIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebTypeComputeIamBindingMapInput)(nil)).Elem(), WebTypeComputeIamBindingMap{})
	pulumi.RegisterOutputType(WebTypeComputeIamBindingOutput{})
	pulumi.RegisterOutputType(WebTypeComputeIamBindingArrayOutput{})
	pulumi.RegisterOutputType(WebTypeComputeIamBindingMapOutput{})
}

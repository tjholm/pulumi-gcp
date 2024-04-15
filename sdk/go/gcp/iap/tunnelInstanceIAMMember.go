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

// Three different resources help you manage your IAM policy for Identity-Aware Proxy TunnelInstance. Each of these resources serves a different use case:
//
// * `iap.TunnelInstanceIAMPolicy`: Authoritative. Sets the IAM policy for the tunnelinstance and replaces any existing policy already attached.
// * `iap.TunnelInstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tunnelinstance are preserved.
// * `iap.TunnelInstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tunnelinstance are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `iap.TunnelInstanceIAMPolicy`: Retrieves the IAM policy for the tunnelinstance
//
// > **Note:** `iap.TunnelInstanceIAMPolicy` **cannot** be used in conjunction with `iap.TunnelInstanceIAMBinding` and `iap.TunnelInstanceIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.TunnelInstanceIAMBinding` resources **can be** used in conjunction with `iap.TunnelInstanceIAMMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## google\_iap\_tunnel\_instance\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
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
//						Role: "roles/iap.tunnelResourceAccessor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iap.NewTunnelInstanceIAMPolicy(ctx, "policy", &iap.TunnelInstanceIAMPolicyArgs{
//				Project:    pulumi.Any(tunnelvm.Project),
//				Zone:       pulumi.Any(tunnelvm.Zone),
//				Instance:   pulumi.Any(tunnelvm.Name),
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
//						Role: "roles/iap.tunnelResourceAccessor",
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
//			_, err = iap.NewTunnelInstanceIAMPolicy(ctx, "policy", &iap.TunnelInstanceIAMPolicyArgs{
//				Project:    pulumi.Any(tunnelvm.Project),
//				Zone:       pulumi.Any(tunnelvm.Zone),
//				Instance:   pulumi.Any(tunnelvm.Name),
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
// ## google\_iap\_tunnel\_instance\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := iap.NewTunnelInstanceIAMBinding(ctx, "binding", &iap.TunnelInstanceIAMBindingArgs{
//				Project:  pulumi.Any(tunnelvm.Project),
//				Zone:     pulumi.Any(tunnelvm.Zone),
//				Instance: pulumi.Any(tunnelvm.Name),
//				Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewTunnelInstanceIAMBinding(ctx, "binding", &iap.TunnelInstanceIAMBindingArgs{
//				Project:  pulumi.Any(tunnelvm.Project),
//				Zone:     pulumi.Any(tunnelvm.Zone),
//				Instance: pulumi.Any(tunnelvm.Name),
//				Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &iap.TunnelInstanceIAMBindingConditionArgs{
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
// ## google\_iap\_tunnel\_instance\_iam\_member
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := iap.NewTunnelInstanceIAMMember(ctx, "member", &iap.TunnelInstanceIAMMemberArgs{
//				Project:  pulumi.Any(tunnelvm.Project),
//				Zone:     pulumi.Any(tunnelvm.Zone),
//				Instance: pulumi.Any(tunnelvm.Name),
//				Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
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
// With IAM Conditions:
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := iap.NewTunnelInstanceIAMMember(ctx, "member", &iap.TunnelInstanceIAMMemberArgs{
//				Project:  pulumi.Any(tunnelvm.Project),
//				Zone:     pulumi.Any(tunnelvm.Zone),
//				Instance: pulumi.Any(tunnelvm.Name),
//				Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
//				Member:   pulumi.String("user:jane@example.com"),
//				Condition: &iap.TunnelInstanceIAMMemberConditionArgs{
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
// ## google\_iap\_tunnel\_instance\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
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
//						Role: "roles/iap.tunnelResourceAccessor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iap.NewTunnelInstanceIAMPolicy(ctx, "policy", &iap.TunnelInstanceIAMPolicyArgs{
//				Project:    pulumi.Any(tunnelvm.Project),
//				Zone:       pulumi.Any(tunnelvm.Zone),
//				Instance:   pulumi.Any(tunnelvm.Name),
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
//						Role: "roles/iap.tunnelResourceAccessor",
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
//			_, err = iap.NewTunnelInstanceIAMPolicy(ctx, "policy", &iap.TunnelInstanceIAMPolicyArgs{
//				Project:    pulumi.Any(tunnelvm.Project),
//				Zone:       pulumi.Any(tunnelvm.Zone),
//				Instance:   pulumi.Any(tunnelvm.Name),
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
// ## google\_iap\_tunnel\_instance\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := iap.NewTunnelInstanceIAMBinding(ctx, "binding", &iap.TunnelInstanceIAMBindingArgs{
//				Project:  pulumi.Any(tunnelvm.Project),
//				Zone:     pulumi.Any(tunnelvm.Zone),
//				Instance: pulumi.Any(tunnelvm.Name),
//				Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewTunnelInstanceIAMBinding(ctx, "binding", &iap.TunnelInstanceIAMBindingArgs{
//				Project:  pulumi.Any(tunnelvm.Project),
//				Zone:     pulumi.Any(tunnelvm.Zone),
//				Instance: pulumi.Any(tunnelvm.Name),
//				Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &iap.TunnelInstanceIAMBindingConditionArgs{
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
// ## google\_iap\_tunnel\_instance\_iam\_member
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := iap.NewTunnelInstanceIAMMember(ctx, "member", &iap.TunnelInstanceIAMMemberArgs{
//				Project:  pulumi.Any(tunnelvm.Project),
//				Zone:     pulumi.Any(tunnelvm.Zone),
//				Instance: pulumi.Any(tunnelvm.Name),
//				Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
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
// With IAM Conditions:
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := iap.NewTunnelInstanceIAMMember(ctx, "member", &iap.TunnelInstanceIAMMemberArgs{
//				Project:  pulumi.Any(tunnelvm.Project),
//				Zone:     pulumi.Any(tunnelvm.Zone),
//				Instance: pulumi.Any(tunnelvm.Name),
//				Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
//				Member:   pulumi.String("user:jane@example.com"),
//				Condition: &iap.TunnelInstanceIAMMemberConditionArgs{
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
// * projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{name}}
//
// * projects/{{project}}/zones/{{zone}}/instances/{{name}}
//
// * {{project}}/{{zone}}/{{name}}
//
// * {{zone}}/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Identity-Aware Proxy tunnelinstance IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember editor "projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}} roles/iap.tunnelResourceAccessor user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember editor "projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}} roles/iap.tunnelResourceAccessor"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember editor projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TunnelInstanceIAMMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition TunnelInstanceIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Instance pulumi.StringOutput `pulumi:"instance"`
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
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewTunnelInstanceIAMMember registers a new resource with the given unique name, arguments, and options.
func NewTunnelInstanceIAMMember(ctx *pulumi.Context,
	name string, args *TunnelInstanceIAMMemberArgs, opts ...pulumi.ResourceOption) (*TunnelInstanceIAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TunnelInstanceIAMMember
	err := ctx.RegisterResource("gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTunnelInstanceIAMMember gets an existing TunnelInstanceIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTunnelInstanceIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TunnelInstanceIAMMemberState, opts ...pulumi.ResourceOption) (*TunnelInstanceIAMMember, error) {
	var resource TunnelInstanceIAMMember
	err := ctx.ReadResource("gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TunnelInstanceIAMMember resources.
type tunnelInstanceIAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *TunnelInstanceIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Instance *string `pulumi:"instance"`
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
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	Zone *string `pulumi:"zone"`
}

type TunnelInstanceIAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition TunnelInstanceIAMMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Instance pulumi.StringPtrInput
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
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	Zone pulumi.StringPtrInput
}

func (TunnelInstanceIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelInstanceIAMMemberState)(nil)).Elem()
}

type tunnelInstanceIAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *TunnelInstanceIAMMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Instance string `pulumi:"instance"`
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
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string  `pulumi:"role"`
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a TunnelInstanceIAMMember resource.
type TunnelInstanceIAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition TunnelInstanceIAMMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Instance pulumi.StringInput
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
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	Zone pulumi.StringPtrInput
}

func (TunnelInstanceIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelInstanceIAMMemberArgs)(nil)).Elem()
}

type TunnelInstanceIAMMemberInput interface {
	pulumi.Input

	ToTunnelInstanceIAMMemberOutput() TunnelInstanceIAMMemberOutput
	ToTunnelInstanceIAMMemberOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberOutput
}

func (*TunnelInstanceIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((**TunnelInstanceIAMMember)(nil)).Elem()
}

func (i *TunnelInstanceIAMMember) ToTunnelInstanceIAMMemberOutput() TunnelInstanceIAMMemberOutput {
	return i.ToTunnelInstanceIAMMemberOutputWithContext(context.Background())
}

func (i *TunnelInstanceIAMMember) ToTunnelInstanceIAMMemberOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMMemberOutput)
}

// TunnelInstanceIAMMemberArrayInput is an input type that accepts TunnelInstanceIAMMemberArray and TunnelInstanceIAMMemberArrayOutput values.
// You can construct a concrete instance of `TunnelInstanceIAMMemberArrayInput` via:
//
//	TunnelInstanceIAMMemberArray{ TunnelInstanceIAMMemberArgs{...} }
type TunnelInstanceIAMMemberArrayInput interface {
	pulumi.Input

	ToTunnelInstanceIAMMemberArrayOutput() TunnelInstanceIAMMemberArrayOutput
	ToTunnelInstanceIAMMemberArrayOutputWithContext(context.Context) TunnelInstanceIAMMemberArrayOutput
}

type TunnelInstanceIAMMemberArray []TunnelInstanceIAMMemberInput

func (TunnelInstanceIAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TunnelInstanceIAMMember)(nil)).Elem()
}

func (i TunnelInstanceIAMMemberArray) ToTunnelInstanceIAMMemberArrayOutput() TunnelInstanceIAMMemberArrayOutput {
	return i.ToTunnelInstanceIAMMemberArrayOutputWithContext(context.Background())
}

func (i TunnelInstanceIAMMemberArray) ToTunnelInstanceIAMMemberArrayOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMMemberArrayOutput)
}

// TunnelInstanceIAMMemberMapInput is an input type that accepts TunnelInstanceIAMMemberMap and TunnelInstanceIAMMemberMapOutput values.
// You can construct a concrete instance of `TunnelInstanceIAMMemberMapInput` via:
//
//	TunnelInstanceIAMMemberMap{ "key": TunnelInstanceIAMMemberArgs{...} }
type TunnelInstanceIAMMemberMapInput interface {
	pulumi.Input

	ToTunnelInstanceIAMMemberMapOutput() TunnelInstanceIAMMemberMapOutput
	ToTunnelInstanceIAMMemberMapOutputWithContext(context.Context) TunnelInstanceIAMMemberMapOutput
}

type TunnelInstanceIAMMemberMap map[string]TunnelInstanceIAMMemberInput

func (TunnelInstanceIAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TunnelInstanceIAMMember)(nil)).Elem()
}

func (i TunnelInstanceIAMMemberMap) ToTunnelInstanceIAMMemberMapOutput() TunnelInstanceIAMMemberMapOutput {
	return i.ToTunnelInstanceIAMMemberMapOutputWithContext(context.Background())
}

func (i TunnelInstanceIAMMemberMap) ToTunnelInstanceIAMMemberMapOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMMemberMapOutput)
}

type TunnelInstanceIAMMemberOutput struct{ *pulumi.OutputState }

func (TunnelInstanceIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TunnelInstanceIAMMember)(nil)).Elem()
}

func (o TunnelInstanceIAMMemberOutput) ToTunnelInstanceIAMMemberOutput() TunnelInstanceIAMMemberOutput {
	return o
}

func (o TunnelInstanceIAMMemberOutput) ToTunnelInstanceIAMMemberOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o TunnelInstanceIAMMemberOutput) Condition() TunnelInstanceIAMMemberConditionPtrOutput {
	return o.ApplyT(func(v *TunnelInstanceIAMMember) TunnelInstanceIAMMemberConditionPtrOutput { return v.Condition }).(TunnelInstanceIAMMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o TunnelInstanceIAMMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelInstanceIAMMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o TunnelInstanceIAMMemberOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelInstanceIAMMember) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
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
func (o TunnelInstanceIAMMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelInstanceIAMMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o TunnelInstanceIAMMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelInstanceIAMMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o TunnelInstanceIAMMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelInstanceIAMMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o TunnelInstanceIAMMemberOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelInstanceIAMMember) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type TunnelInstanceIAMMemberArrayOutput struct{ *pulumi.OutputState }

func (TunnelInstanceIAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TunnelInstanceIAMMember)(nil)).Elem()
}

func (o TunnelInstanceIAMMemberArrayOutput) ToTunnelInstanceIAMMemberArrayOutput() TunnelInstanceIAMMemberArrayOutput {
	return o
}

func (o TunnelInstanceIAMMemberArrayOutput) ToTunnelInstanceIAMMemberArrayOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberArrayOutput {
	return o
}

func (o TunnelInstanceIAMMemberArrayOutput) Index(i pulumi.IntInput) TunnelInstanceIAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TunnelInstanceIAMMember {
		return vs[0].([]*TunnelInstanceIAMMember)[vs[1].(int)]
	}).(TunnelInstanceIAMMemberOutput)
}

type TunnelInstanceIAMMemberMapOutput struct{ *pulumi.OutputState }

func (TunnelInstanceIAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TunnelInstanceIAMMember)(nil)).Elem()
}

func (o TunnelInstanceIAMMemberMapOutput) ToTunnelInstanceIAMMemberMapOutput() TunnelInstanceIAMMemberMapOutput {
	return o
}

func (o TunnelInstanceIAMMemberMapOutput) ToTunnelInstanceIAMMemberMapOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberMapOutput {
	return o
}

func (o TunnelInstanceIAMMemberMapOutput) MapIndex(k pulumi.StringInput) TunnelInstanceIAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TunnelInstanceIAMMember {
		return vs[0].(map[string]*TunnelInstanceIAMMember)[vs[1].(string)]
	}).(TunnelInstanceIAMMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TunnelInstanceIAMMemberInput)(nil)).Elem(), &TunnelInstanceIAMMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*TunnelInstanceIAMMemberArrayInput)(nil)).Elem(), TunnelInstanceIAMMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TunnelInstanceIAMMemberMapInput)(nil)).Elem(), TunnelInstanceIAMMemberMap{})
	pulumi.RegisterOutputType(TunnelInstanceIAMMemberOutput{})
	pulumi.RegisterOutputType(TunnelInstanceIAMMemberArrayOutput{})
	pulumi.RegisterOutputType(TunnelInstanceIAMMemberMapOutput{})
}

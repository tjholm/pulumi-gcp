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

// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebBackendService. Each of these resources serves a different use case:
//
// * `iap.WebBackendServiceIamPolicy`: Authoritative. Sets the IAM policy for the webbackendservice and replaces any existing policy already attached.
// * `iap.WebBackendServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webbackendservice are preserved.
// * `iap.WebBackendServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webbackendservice are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `iap.WebBackendServiceIamPolicy`: Retrieves the IAM policy for the webbackendservice
//
// > **Note:** `iap.WebBackendServiceIamPolicy` **cannot** be used in conjunction with `iap.WebBackendServiceIamBinding` and `iap.WebBackendServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebBackendServiceIamBinding` resources **can be** used in conjunction with `iap.WebBackendServiceIamMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## google\_iap\_web\_backend\_service\_iam\_policy
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
//			_, err = iap.NewWebBackendServiceIamPolicy(ctx, "policy", &iap.WebBackendServiceIamPolicyArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				PolicyData:        pulumi.String(admin.PolicyData),
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
//			_, err = iap.NewWebBackendServiceIamPolicy(ctx, "policy", &iap.WebBackendServiceIamPolicyArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				PolicyData:        pulumi.String(admin.PolicyData),
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
// ## google\_iap\_web\_backend\_service\_iam\_binding
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
//			_, err := iap.NewWebBackendServiceIamBinding(ctx, "binding", &iap.WebBackendServiceIamBindingArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
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
//			_, err := iap.NewWebBackendServiceIamBinding(ctx, "binding", &iap.WebBackendServiceIamBindingArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &iap.WebBackendServiceIamBindingConditionArgs{
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
// ## google\_iap\_web\_backend\_service\_iam\_member
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
//			_, err := iap.NewWebBackendServiceIamMember(ctx, "member", &iap.WebBackendServiceIamMemberArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
//				Member:            pulumi.String("user:jane@example.com"),
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
//			_, err := iap.NewWebBackendServiceIamMember(ctx, "member", &iap.WebBackendServiceIamMemberArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
//				Member:            pulumi.String("user:jane@example.com"),
//				Condition: &iap.WebBackendServiceIamMemberConditionArgs{
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
// ## google\_iap\_web\_backend\_service\_iam\_policy
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
//			_, err = iap.NewWebBackendServiceIamPolicy(ctx, "policy", &iap.WebBackendServiceIamPolicyArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				PolicyData:        pulumi.String(admin.PolicyData),
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
//			_, err = iap.NewWebBackendServiceIamPolicy(ctx, "policy", &iap.WebBackendServiceIamPolicyArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				PolicyData:        pulumi.String(admin.PolicyData),
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
// ## google\_iap\_web\_backend\_service\_iam\_binding
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
//			_, err := iap.NewWebBackendServiceIamBinding(ctx, "binding", &iap.WebBackendServiceIamBindingArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
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
//			_, err := iap.NewWebBackendServiceIamBinding(ctx, "binding", &iap.WebBackendServiceIamBindingArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &iap.WebBackendServiceIamBindingConditionArgs{
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
// ## google\_iap\_web\_backend\_service\_iam\_member
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
//			_, err := iap.NewWebBackendServiceIamMember(ctx, "member", &iap.WebBackendServiceIamMemberArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
//				Member:            pulumi.String("user:jane@example.com"),
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
//			_, err := iap.NewWebBackendServiceIamMember(ctx, "member", &iap.WebBackendServiceIamMemberArgs{
//				Project:           pulumi.Any(_default.Project),
//				WebBackendService: pulumi.Any(_default.Name),
//				Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
//				Member:            pulumi.String("user:jane@example.com"),
//				Condition: &iap.WebBackendServiceIamMemberConditionArgs{
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
// * projects/{{project}}/iap_web/compute/services/{{name}}
//
// * {{project}}/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Identity-Aware Proxy webbackendservice IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor projects/{{project}}/iap_web/compute/services/{{web_backend_service}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebBackendServiceIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringOutput `pulumi:"webBackendService"`
}

// NewWebBackendServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, args *WebBackendServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*WebBackendServiceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.WebBackendService == nil {
		return nil, errors.New("invalid value for required argument 'WebBackendService'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebBackendServiceIamPolicy
	err := ctx.RegisterResource("gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebBackendServiceIamPolicy gets an existing WebBackendServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebBackendServiceIamPolicyState, opts ...pulumi.ResourceOption) (*WebBackendServiceIamPolicy, error) {
	var resource WebBackendServiceIamPolicy
	err := ctx.ReadResource("gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebBackendServiceIamPolicy resources.
type webBackendServiceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService *string `pulumi:"webBackendService"`
}

type WebBackendServiceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringPtrInput
}

func (WebBackendServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamPolicyState)(nil)).Elem()
}

type webBackendServiceIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService string `pulumi:"webBackendService"`
}

// The set of arguments for constructing a WebBackendServiceIamPolicy resource.
type WebBackendServiceIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringInput
}

func (WebBackendServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamPolicyArgs)(nil)).Elem()
}

type WebBackendServiceIamPolicyInput interface {
	pulumi.Input

	ToWebBackendServiceIamPolicyOutput() WebBackendServiceIamPolicyOutput
	ToWebBackendServiceIamPolicyOutputWithContext(ctx context.Context) WebBackendServiceIamPolicyOutput
}

func (*WebBackendServiceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**WebBackendServiceIamPolicy)(nil)).Elem()
}

func (i *WebBackendServiceIamPolicy) ToWebBackendServiceIamPolicyOutput() WebBackendServiceIamPolicyOutput {
	return i.ToWebBackendServiceIamPolicyOutputWithContext(context.Background())
}

func (i *WebBackendServiceIamPolicy) ToWebBackendServiceIamPolicyOutputWithContext(ctx context.Context) WebBackendServiceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamPolicyOutput)
}

// WebBackendServiceIamPolicyArrayInput is an input type that accepts WebBackendServiceIamPolicyArray and WebBackendServiceIamPolicyArrayOutput values.
// You can construct a concrete instance of `WebBackendServiceIamPolicyArrayInput` via:
//
//	WebBackendServiceIamPolicyArray{ WebBackendServiceIamPolicyArgs{...} }
type WebBackendServiceIamPolicyArrayInput interface {
	pulumi.Input

	ToWebBackendServiceIamPolicyArrayOutput() WebBackendServiceIamPolicyArrayOutput
	ToWebBackendServiceIamPolicyArrayOutputWithContext(context.Context) WebBackendServiceIamPolicyArrayOutput
}

type WebBackendServiceIamPolicyArray []WebBackendServiceIamPolicyInput

func (WebBackendServiceIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebBackendServiceIamPolicy)(nil)).Elem()
}

func (i WebBackendServiceIamPolicyArray) ToWebBackendServiceIamPolicyArrayOutput() WebBackendServiceIamPolicyArrayOutput {
	return i.ToWebBackendServiceIamPolicyArrayOutputWithContext(context.Background())
}

func (i WebBackendServiceIamPolicyArray) ToWebBackendServiceIamPolicyArrayOutputWithContext(ctx context.Context) WebBackendServiceIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamPolicyArrayOutput)
}

// WebBackendServiceIamPolicyMapInput is an input type that accepts WebBackendServiceIamPolicyMap and WebBackendServiceIamPolicyMapOutput values.
// You can construct a concrete instance of `WebBackendServiceIamPolicyMapInput` via:
//
//	WebBackendServiceIamPolicyMap{ "key": WebBackendServiceIamPolicyArgs{...} }
type WebBackendServiceIamPolicyMapInput interface {
	pulumi.Input

	ToWebBackendServiceIamPolicyMapOutput() WebBackendServiceIamPolicyMapOutput
	ToWebBackendServiceIamPolicyMapOutputWithContext(context.Context) WebBackendServiceIamPolicyMapOutput
}

type WebBackendServiceIamPolicyMap map[string]WebBackendServiceIamPolicyInput

func (WebBackendServiceIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebBackendServiceIamPolicy)(nil)).Elem()
}

func (i WebBackendServiceIamPolicyMap) ToWebBackendServiceIamPolicyMapOutput() WebBackendServiceIamPolicyMapOutput {
	return i.ToWebBackendServiceIamPolicyMapOutputWithContext(context.Background())
}

func (i WebBackendServiceIamPolicyMap) ToWebBackendServiceIamPolicyMapOutputWithContext(ctx context.Context) WebBackendServiceIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamPolicyMapOutput)
}

type WebBackendServiceIamPolicyOutput struct{ *pulumi.OutputState }

func (WebBackendServiceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebBackendServiceIamPolicy)(nil)).Elem()
}

func (o WebBackendServiceIamPolicyOutput) ToWebBackendServiceIamPolicyOutput() WebBackendServiceIamPolicyOutput {
	return o
}

func (o WebBackendServiceIamPolicyOutput) ToWebBackendServiceIamPolicyOutputWithContext(ctx context.Context) WebBackendServiceIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o WebBackendServiceIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WebBackendServiceIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o WebBackendServiceIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *WebBackendServiceIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o WebBackendServiceIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WebBackendServiceIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o WebBackendServiceIamPolicyOutput) WebBackendService() pulumi.StringOutput {
	return o.ApplyT(func(v *WebBackendServiceIamPolicy) pulumi.StringOutput { return v.WebBackendService }).(pulumi.StringOutput)
}

type WebBackendServiceIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (WebBackendServiceIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebBackendServiceIamPolicy)(nil)).Elem()
}

func (o WebBackendServiceIamPolicyArrayOutput) ToWebBackendServiceIamPolicyArrayOutput() WebBackendServiceIamPolicyArrayOutput {
	return o
}

func (o WebBackendServiceIamPolicyArrayOutput) ToWebBackendServiceIamPolicyArrayOutputWithContext(ctx context.Context) WebBackendServiceIamPolicyArrayOutput {
	return o
}

func (o WebBackendServiceIamPolicyArrayOutput) Index(i pulumi.IntInput) WebBackendServiceIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebBackendServiceIamPolicy {
		return vs[0].([]*WebBackendServiceIamPolicy)[vs[1].(int)]
	}).(WebBackendServiceIamPolicyOutput)
}

type WebBackendServiceIamPolicyMapOutput struct{ *pulumi.OutputState }

func (WebBackendServiceIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebBackendServiceIamPolicy)(nil)).Elem()
}

func (o WebBackendServiceIamPolicyMapOutput) ToWebBackendServiceIamPolicyMapOutput() WebBackendServiceIamPolicyMapOutput {
	return o
}

func (o WebBackendServiceIamPolicyMapOutput) ToWebBackendServiceIamPolicyMapOutputWithContext(ctx context.Context) WebBackendServiceIamPolicyMapOutput {
	return o
}

func (o WebBackendServiceIamPolicyMapOutput) MapIndex(k pulumi.StringInput) WebBackendServiceIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebBackendServiceIamPolicy {
		return vs[0].(map[string]*WebBackendServiceIamPolicy)[vs[1].(string)]
	}).(WebBackendServiceIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebBackendServiceIamPolicyInput)(nil)).Elem(), &WebBackendServiceIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebBackendServiceIamPolicyArrayInput)(nil)).Elem(), WebBackendServiceIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebBackendServiceIamPolicyMapInput)(nil)).Elem(), WebBackendServiceIamPolicyMap{})
	pulumi.RegisterOutputType(WebBackendServiceIamPolicyOutput{})
	pulumi.RegisterOutputType(WebBackendServiceIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(WebBackendServiceIamPolicyMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
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
//				Project:           pulumi.Any(google_compute_backend_service.Default.Project),
//				WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
//				PolicyData:        *pulumi.String(admin.PolicyData),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
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
//				Project:           pulumi.Any(google_compute_backend_service.Default.Project),
//				WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
//				PolicyData:        *pulumi.String(admin.PolicyData),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## google\_iap\_web\_backend\_service\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebBackendServiceIamBinding(ctx, "binding", &iap.WebBackendServiceIamBindingArgs{
//				Project:           pulumi.Any(google_compute_backend_service.Default.Project),
//				WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
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
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebBackendServiceIamBinding(ctx, "binding", &iap.WebBackendServiceIamBindingArgs{
//				Project:           pulumi.Any(google_compute_backend_service.Default.Project),
//				WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
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
// ## google\_iap\_web\_backend\_service\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebBackendServiceIamMember(ctx, "member", &iap.WebBackendServiceIamMemberArgs{
//				Project:           pulumi.Any(google_compute_backend_service.Default.Project),
//				WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
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
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebBackendServiceIamMember(ctx, "member", &iap.WebBackendServiceIamMemberArgs{
//				Project:           pulumi.Any(google_compute_backend_service.Default.Project),
//				WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
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
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/compute/services/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webbackendservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember editor projects/{{project}}/iap_web/compute/services/{{web_backend_service}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebBackendServiceIamMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebBackendServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
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
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringOutput `pulumi:"webBackendService"`
}

// NewWebBackendServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewWebBackendServiceIamMember(ctx *pulumi.Context,
	name string, args *WebBackendServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*WebBackendServiceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.WebBackendService == nil {
		return nil, errors.New("invalid value for required argument 'WebBackendService'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebBackendServiceIamMember
	err := ctx.RegisterResource("gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebBackendServiceIamMember gets an existing WebBackendServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebBackendServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebBackendServiceIamMemberState, opts ...pulumi.ResourceOption) (*WebBackendServiceIamMember, error) {
	var resource WebBackendServiceIamMember
	err := ctx.ReadResource("gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebBackendServiceIamMember resources.
type webBackendServiceIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebBackendServiceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
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
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService *string `pulumi:"webBackendService"`
}

type WebBackendServiceIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebBackendServiceIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
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
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringPtrInput
}

func (WebBackendServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamMemberState)(nil)).Elem()
}

type webBackendServiceIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebBackendServiceIamMemberCondition `pulumi:"condition"`
	Member    string                               `pulumi:"member"`
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
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService string `pulumi:"webBackendService"`
}

// The set of arguments for constructing a WebBackendServiceIamMember resource.
type WebBackendServiceIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebBackendServiceIamMemberConditionPtrInput
	Member    pulumi.StringInput
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
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringInput
}

func (WebBackendServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamMemberArgs)(nil)).Elem()
}

type WebBackendServiceIamMemberInput interface {
	pulumi.Input

	ToWebBackendServiceIamMemberOutput() WebBackendServiceIamMemberOutput
	ToWebBackendServiceIamMemberOutputWithContext(ctx context.Context) WebBackendServiceIamMemberOutput
}

func (*WebBackendServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**WebBackendServiceIamMember)(nil)).Elem()
}

func (i *WebBackendServiceIamMember) ToWebBackendServiceIamMemberOutput() WebBackendServiceIamMemberOutput {
	return i.ToWebBackendServiceIamMemberOutputWithContext(context.Background())
}

func (i *WebBackendServiceIamMember) ToWebBackendServiceIamMemberOutputWithContext(ctx context.Context) WebBackendServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamMemberOutput)
}

func (i *WebBackendServiceIamMember) ToOutput(ctx context.Context) pulumix.Output[*WebBackendServiceIamMember] {
	return pulumix.Output[*WebBackendServiceIamMember]{
		OutputState: i.ToWebBackendServiceIamMemberOutputWithContext(ctx).OutputState,
	}
}

// WebBackendServiceIamMemberArrayInput is an input type that accepts WebBackendServiceIamMemberArray and WebBackendServiceIamMemberArrayOutput values.
// You can construct a concrete instance of `WebBackendServiceIamMemberArrayInput` via:
//
//	WebBackendServiceIamMemberArray{ WebBackendServiceIamMemberArgs{...} }
type WebBackendServiceIamMemberArrayInput interface {
	pulumi.Input

	ToWebBackendServiceIamMemberArrayOutput() WebBackendServiceIamMemberArrayOutput
	ToWebBackendServiceIamMemberArrayOutputWithContext(context.Context) WebBackendServiceIamMemberArrayOutput
}

type WebBackendServiceIamMemberArray []WebBackendServiceIamMemberInput

func (WebBackendServiceIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebBackendServiceIamMember)(nil)).Elem()
}

func (i WebBackendServiceIamMemberArray) ToWebBackendServiceIamMemberArrayOutput() WebBackendServiceIamMemberArrayOutput {
	return i.ToWebBackendServiceIamMemberArrayOutputWithContext(context.Background())
}

func (i WebBackendServiceIamMemberArray) ToWebBackendServiceIamMemberArrayOutputWithContext(ctx context.Context) WebBackendServiceIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamMemberArrayOutput)
}

func (i WebBackendServiceIamMemberArray) ToOutput(ctx context.Context) pulumix.Output[[]*WebBackendServiceIamMember] {
	return pulumix.Output[[]*WebBackendServiceIamMember]{
		OutputState: i.ToWebBackendServiceIamMemberArrayOutputWithContext(ctx).OutputState,
	}
}

// WebBackendServiceIamMemberMapInput is an input type that accepts WebBackendServiceIamMemberMap and WebBackendServiceIamMemberMapOutput values.
// You can construct a concrete instance of `WebBackendServiceIamMemberMapInput` via:
//
//	WebBackendServiceIamMemberMap{ "key": WebBackendServiceIamMemberArgs{...} }
type WebBackendServiceIamMemberMapInput interface {
	pulumi.Input

	ToWebBackendServiceIamMemberMapOutput() WebBackendServiceIamMemberMapOutput
	ToWebBackendServiceIamMemberMapOutputWithContext(context.Context) WebBackendServiceIamMemberMapOutput
}

type WebBackendServiceIamMemberMap map[string]WebBackendServiceIamMemberInput

func (WebBackendServiceIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebBackendServiceIamMember)(nil)).Elem()
}

func (i WebBackendServiceIamMemberMap) ToWebBackendServiceIamMemberMapOutput() WebBackendServiceIamMemberMapOutput {
	return i.ToWebBackendServiceIamMemberMapOutputWithContext(context.Background())
}

func (i WebBackendServiceIamMemberMap) ToWebBackendServiceIamMemberMapOutputWithContext(ctx context.Context) WebBackendServiceIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamMemberMapOutput)
}

func (i WebBackendServiceIamMemberMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WebBackendServiceIamMember] {
	return pulumix.Output[map[string]*WebBackendServiceIamMember]{
		OutputState: i.ToWebBackendServiceIamMemberMapOutputWithContext(ctx).OutputState,
	}
}

type WebBackendServiceIamMemberOutput struct{ *pulumi.OutputState }

func (WebBackendServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebBackendServiceIamMember)(nil)).Elem()
}

func (o WebBackendServiceIamMemberOutput) ToWebBackendServiceIamMemberOutput() WebBackendServiceIamMemberOutput {
	return o
}

func (o WebBackendServiceIamMemberOutput) ToWebBackendServiceIamMemberOutputWithContext(ctx context.Context) WebBackendServiceIamMemberOutput {
	return o
}

func (o WebBackendServiceIamMemberOutput) ToOutput(ctx context.Context) pulumix.Output[*WebBackendServiceIamMember] {
	return pulumix.Output[*WebBackendServiceIamMember]{
		OutputState: o.OutputState,
	}
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o WebBackendServiceIamMemberOutput) Condition() WebBackendServiceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *WebBackendServiceIamMember) WebBackendServiceIamMemberConditionPtrOutput { return v.Condition }).(WebBackendServiceIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o WebBackendServiceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WebBackendServiceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o WebBackendServiceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *WebBackendServiceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
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
func (o WebBackendServiceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WebBackendServiceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o WebBackendServiceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *WebBackendServiceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o WebBackendServiceIamMemberOutput) WebBackendService() pulumi.StringOutput {
	return o.ApplyT(func(v *WebBackendServiceIamMember) pulumi.StringOutput { return v.WebBackendService }).(pulumi.StringOutput)
}

type WebBackendServiceIamMemberArrayOutput struct{ *pulumi.OutputState }

func (WebBackendServiceIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebBackendServiceIamMember)(nil)).Elem()
}

func (o WebBackendServiceIamMemberArrayOutput) ToWebBackendServiceIamMemberArrayOutput() WebBackendServiceIamMemberArrayOutput {
	return o
}

func (o WebBackendServiceIamMemberArrayOutput) ToWebBackendServiceIamMemberArrayOutputWithContext(ctx context.Context) WebBackendServiceIamMemberArrayOutput {
	return o
}

func (o WebBackendServiceIamMemberArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WebBackendServiceIamMember] {
	return pulumix.Output[[]*WebBackendServiceIamMember]{
		OutputState: o.OutputState,
	}
}

func (o WebBackendServiceIamMemberArrayOutput) Index(i pulumi.IntInput) WebBackendServiceIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebBackendServiceIamMember {
		return vs[0].([]*WebBackendServiceIamMember)[vs[1].(int)]
	}).(WebBackendServiceIamMemberOutput)
}

type WebBackendServiceIamMemberMapOutput struct{ *pulumi.OutputState }

func (WebBackendServiceIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebBackendServiceIamMember)(nil)).Elem()
}

func (o WebBackendServiceIamMemberMapOutput) ToWebBackendServiceIamMemberMapOutput() WebBackendServiceIamMemberMapOutput {
	return o
}

func (o WebBackendServiceIamMemberMapOutput) ToWebBackendServiceIamMemberMapOutputWithContext(ctx context.Context) WebBackendServiceIamMemberMapOutput {
	return o
}

func (o WebBackendServiceIamMemberMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WebBackendServiceIamMember] {
	return pulumix.Output[map[string]*WebBackendServiceIamMember]{
		OutputState: o.OutputState,
	}
}

func (o WebBackendServiceIamMemberMapOutput) MapIndex(k pulumi.StringInput) WebBackendServiceIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebBackendServiceIamMember {
		return vs[0].(map[string]*WebBackendServiceIamMember)[vs[1].(string)]
	}).(WebBackendServiceIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebBackendServiceIamMemberInput)(nil)).Elem(), &WebBackendServiceIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebBackendServiceIamMemberArrayInput)(nil)).Elem(), WebBackendServiceIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebBackendServiceIamMemberMapInput)(nil)).Elem(), WebBackendServiceIamMemberMap{})
	pulumi.RegisterOutputType(WebBackendServiceIamMemberOutput{})
	pulumi.RegisterOutputType(WebBackendServiceIamMemberArrayOutput{})
	pulumi.RegisterOutputType(WebBackendServiceIamMemberMapOutput{})
}

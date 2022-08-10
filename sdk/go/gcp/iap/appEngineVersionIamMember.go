// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy AppEngineVersion. Each of these resources serves a different use case:
//
// * `iap.AppEngineVersionIamPolicy`: Authoritative. Sets the IAM policy for the appengineversion and replaces any existing policy already attached.
// * `iap.AppEngineVersionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the appengineversion are preserved.
// * `iap.AppEngineVersionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the appengineversion are preserved.
//
// > **Note:** `iap.AppEngineVersionIamPolicy` **cannot** be used in conjunction with `iap.AppEngineVersionIamBinding` and `iap.AppEngineVersionIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.AppEngineVersionIamBinding` resources **can be** used in conjunction with `iap.AppEngineVersionIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_app\_engine\_version\_iam\_policy
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
//					organizations.GetIAMPolicyBinding{
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
//			_, err = iap.NewAppEngineVersionIamPolicy(ctx, "policy", &iap.AppEngineVersionIamPolicyArgs{
//				Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
//				VersionId:  pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
//					organizations.GetIAMPolicyBinding{
//						Role: "roles/iap.httpsResourceAccessor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Condition: organizations.GetIAMPolicyBindingCondition{
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
//			_, err = iap.NewAppEngineVersionIamPolicy(ctx, "policy", &iap.AppEngineVersionIamPolicyArgs{
//				Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
//				VersionId:  pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
// ## google\_iap\_app\_engine\_version\_iam\_binding
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
//			_, err := iap.NewAppEngineVersionIamBinding(ctx, "binding", &iap.AppEngineVersionIamBindingArgs{
//				AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
//				Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
//				VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
//			_, err := iap.NewAppEngineVersionIamBinding(ctx, "binding", &iap.AppEngineVersionIamBindingArgs{
//				AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Condition: &iap.AppEngineVersionIamBindingConditionArgs{
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
//					Title:       pulumi.String("expires_after_2019_12_31"),
//				},
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
//				Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
//				VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## google\_iap\_app\_engine\_version\_iam\_member
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
//			_, err := iap.NewAppEngineVersionIamMember(ctx, "member", &iap.AppEngineVersionIamMemberArgs{
//				AppId:     pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Member:    pulumi.String("user:jane@example.com"),
//				Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
//				Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
//				VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
//			_, err := iap.NewAppEngineVersionIamMember(ctx, "member", &iap.AppEngineVersionIamMemberArgs{
//				AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Condition: &iap.AppEngineVersionIamMemberConditionArgs{
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
//					Title:       pulumi.String("expires_after_2019_12_31"),
//				},
//				Member:    pulumi.String("user:jane@example.com"),
//				Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
//				Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
//				VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} * {{project}}/{{appId}}/{{service}}/{{versionId}} * {{appId}}/{{service}}/{{versionId}} * {{version}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy appengineversion IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/appEngineVersionIamMember:AppEngineVersionIamMember editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/appEngineVersionIamMember:AppEngineVersionIamMember editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/appEngineVersionIamMember:AppEngineVersionIamMember editor projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AppEngineVersionIamMember struct {
	pulumi.CustomResourceState

	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineVersionIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringOutput `pulumi:"service"`
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId pulumi.StringOutput `pulumi:"versionId"`
}

// NewAppEngineVersionIamMember registers a new resource with the given unique name, arguments, and options.
func NewAppEngineVersionIamMember(ctx *pulumi.Context,
	name string, args *AppEngineVersionIamMemberArgs, opts ...pulumi.ResourceOption) (*AppEngineVersionIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	if args.VersionId == nil {
		return nil, errors.New("invalid value for required argument 'VersionId'")
	}
	var resource AppEngineVersionIamMember
	err := ctx.RegisterResource("gcp:iap/appEngineVersionIamMember:AppEngineVersionIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppEngineVersionIamMember gets an existing AppEngineVersionIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppEngineVersionIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppEngineVersionIamMemberState, opts ...pulumi.ResourceOption) (*AppEngineVersionIamMember, error) {
	var resource AppEngineVersionIamMember
	err := ctx.ReadResource("gcp:iap/appEngineVersionIamMember:AppEngineVersionIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppEngineVersionIamMember resources.
type appEngineVersionIamMemberState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId *string `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *AppEngineVersionIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service *string `pulumi:"service"`
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId *string `pulumi:"versionId"`
}

type AppEngineVersionIamMemberState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringPtrInput
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineVersionIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringPtrInput
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId pulumi.StringPtrInput
}

func (AppEngineVersionIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineVersionIamMemberState)(nil)).Elem()
}

type appEngineVersionIamMemberArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId string `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *AppEngineVersionIamMemberCondition `pulumi:"condition"`
	Member    string                              `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service string `pulumi:"service"`
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId string `pulumi:"versionId"`
}

// The set of arguments for constructing a AppEngineVersionIamMember resource.
type AppEngineVersionIamMemberArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringInput
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineVersionIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringInput
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId pulumi.StringInput
}

func (AppEngineVersionIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineVersionIamMemberArgs)(nil)).Elem()
}

type AppEngineVersionIamMemberInput interface {
	pulumi.Input

	ToAppEngineVersionIamMemberOutput() AppEngineVersionIamMemberOutput
	ToAppEngineVersionIamMemberOutputWithContext(ctx context.Context) AppEngineVersionIamMemberOutput
}

func (*AppEngineVersionIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineVersionIamMember)(nil)).Elem()
}

func (i *AppEngineVersionIamMember) ToAppEngineVersionIamMemberOutput() AppEngineVersionIamMemberOutput {
	return i.ToAppEngineVersionIamMemberOutputWithContext(context.Background())
}

func (i *AppEngineVersionIamMember) ToAppEngineVersionIamMemberOutputWithContext(ctx context.Context) AppEngineVersionIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamMemberOutput)
}

// AppEngineVersionIamMemberArrayInput is an input type that accepts AppEngineVersionIamMemberArray and AppEngineVersionIamMemberArrayOutput values.
// You can construct a concrete instance of `AppEngineVersionIamMemberArrayInput` via:
//
//	AppEngineVersionIamMemberArray{ AppEngineVersionIamMemberArgs{...} }
type AppEngineVersionIamMemberArrayInput interface {
	pulumi.Input

	ToAppEngineVersionIamMemberArrayOutput() AppEngineVersionIamMemberArrayOutput
	ToAppEngineVersionIamMemberArrayOutputWithContext(context.Context) AppEngineVersionIamMemberArrayOutput
}

type AppEngineVersionIamMemberArray []AppEngineVersionIamMemberInput

func (AppEngineVersionIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppEngineVersionIamMember)(nil)).Elem()
}

func (i AppEngineVersionIamMemberArray) ToAppEngineVersionIamMemberArrayOutput() AppEngineVersionIamMemberArrayOutput {
	return i.ToAppEngineVersionIamMemberArrayOutputWithContext(context.Background())
}

func (i AppEngineVersionIamMemberArray) ToAppEngineVersionIamMemberArrayOutputWithContext(ctx context.Context) AppEngineVersionIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamMemberArrayOutput)
}

// AppEngineVersionIamMemberMapInput is an input type that accepts AppEngineVersionIamMemberMap and AppEngineVersionIamMemberMapOutput values.
// You can construct a concrete instance of `AppEngineVersionIamMemberMapInput` via:
//
//	AppEngineVersionIamMemberMap{ "key": AppEngineVersionIamMemberArgs{...} }
type AppEngineVersionIamMemberMapInput interface {
	pulumi.Input

	ToAppEngineVersionIamMemberMapOutput() AppEngineVersionIamMemberMapOutput
	ToAppEngineVersionIamMemberMapOutputWithContext(context.Context) AppEngineVersionIamMemberMapOutput
}

type AppEngineVersionIamMemberMap map[string]AppEngineVersionIamMemberInput

func (AppEngineVersionIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppEngineVersionIamMember)(nil)).Elem()
}

func (i AppEngineVersionIamMemberMap) ToAppEngineVersionIamMemberMapOutput() AppEngineVersionIamMemberMapOutput {
	return i.ToAppEngineVersionIamMemberMapOutputWithContext(context.Background())
}

func (i AppEngineVersionIamMemberMap) ToAppEngineVersionIamMemberMapOutputWithContext(ctx context.Context) AppEngineVersionIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamMemberMapOutput)
}

type AppEngineVersionIamMemberOutput struct{ *pulumi.OutputState }

func (AppEngineVersionIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineVersionIamMember)(nil)).Elem()
}

func (o AppEngineVersionIamMemberOutput) ToAppEngineVersionIamMemberOutput() AppEngineVersionIamMemberOutput {
	return o
}

func (o AppEngineVersionIamMemberOutput) ToAppEngineVersionIamMemberOutputWithContext(ctx context.Context) AppEngineVersionIamMemberOutput {
	return o
}

// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
func (o AppEngineVersionIamMemberOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineVersionIamMember) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o AppEngineVersionIamMemberOutput) Condition() AppEngineVersionIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *AppEngineVersionIamMember) AppEngineVersionIamMemberConditionPtrOutput { return v.Condition }).(AppEngineVersionIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o AppEngineVersionIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineVersionIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AppEngineVersionIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineVersionIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o AppEngineVersionIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineVersionIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o AppEngineVersionIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineVersionIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
func (o AppEngineVersionIamMemberOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineVersionIamMember) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
func (o AppEngineVersionIamMemberOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineVersionIamMember) pulumi.StringOutput { return v.VersionId }).(pulumi.StringOutput)
}

type AppEngineVersionIamMemberArrayOutput struct{ *pulumi.OutputState }

func (AppEngineVersionIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppEngineVersionIamMember)(nil)).Elem()
}

func (o AppEngineVersionIamMemberArrayOutput) ToAppEngineVersionIamMemberArrayOutput() AppEngineVersionIamMemberArrayOutput {
	return o
}

func (o AppEngineVersionIamMemberArrayOutput) ToAppEngineVersionIamMemberArrayOutputWithContext(ctx context.Context) AppEngineVersionIamMemberArrayOutput {
	return o
}

func (o AppEngineVersionIamMemberArrayOutput) Index(i pulumi.IntInput) AppEngineVersionIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppEngineVersionIamMember {
		return vs[0].([]*AppEngineVersionIamMember)[vs[1].(int)]
	}).(AppEngineVersionIamMemberOutput)
}

type AppEngineVersionIamMemberMapOutput struct{ *pulumi.OutputState }

func (AppEngineVersionIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppEngineVersionIamMember)(nil)).Elem()
}

func (o AppEngineVersionIamMemberMapOutput) ToAppEngineVersionIamMemberMapOutput() AppEngineVersionIamMemberMapOutput {
	return o
}

func (o AppEngineVersionIamMemberMapOutput) ToAppEngineVersionIamMemberMapOutputWithContext(ctx context.Context) AppEngineVersionIamMemberMapOutput {
	return o
}

func (o AppEngineVersionIamMemberMapOutput) MapIndex(k pulumi.StringInput) AppEngineVersionIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppEngineVersionIamMember {
		return vs[0].(map[string]*AppEngineVersionIamMember)[vs[1].(string)]
	}).(AppEngineVersionIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineVersionIamMemberInput)(nil)).Elem(), &AppEngineVersionIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineVersionIamMemberArrayInput)(nil)).Elem(), AppEngineVersionIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineVersionIamMemberMapInput)(nil)).Elem(), AppEngineVersionIamMemberMap{})
	pulumi.RegisterOutputType(AppEngineVersionIamMemberOutput{})
	pulumi.RegisterOutputType(AppEngineVersionIamMemberArrayOutput{})
	pulumi.RegisterOutputType(AppEngineVersionIamMemberMapOutput{})
}

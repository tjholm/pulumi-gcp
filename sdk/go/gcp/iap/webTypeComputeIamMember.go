// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeCompute. Each of these resources serves a different use case:
//
// * `iap.WebTypeComputeIamPolicy`: Authoritative. Sets the IAM policy for the webtypecompute and replaces any existing policy already attached.
// * `iap.WebTypeComputeIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypecompute are preserved.
// * `iap.WebTypeComputeIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypecompute are preserved.
//
// > **Note:** `iap.WebTypeComputeIamPolicy` **cannot** be used in conjunction with `iap.WebTypeComputeIamBinding` and `iap.WebTypeComputeIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebTypeComputeIamBinding` resources **can be** used in conjunction with `iap.WebTypeComputeIamMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## google\_iap\_web\_type\_compute\_iam\_policy
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
//			_, err = iap.NewWebTypeComputeIamPolicy(ctx, "policy", &iap.WebTypeComputeIamPolicyArgs{
//				Project:    pulumi.Any(google_project_service.Project_service.Project),
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
//			_, err = iap.NewWebTypeComputeIamPolicy(ctx, "policy", &iap.WebTypeComputeIamPolicyArgs{
//				Project:    pulumi.Any(google_project_service.Project_service.Project),
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
// ## google\_iap\_web\_type\_compute\_iam\_binding
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
//			_, err := iap.NewWebTypeComputeIamBinding(ctx, "binding", &iap.WebTypeComputeIamBindingArgs{
//				Project: pulumi.Any(google_project_service.Project_service.Project),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebTypeComputeIamBinding(ctx, "binding", &iap.WebTypeComputeIamBindingArgs{
//				Project: pulumi.Any(google_project_service.Project_service.Project),
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
// ## google\_iap\_web\_type\_compute\_iam\_member
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
//			_, err := iap.NewWebTypeComputeIamMember(ctx, "member", &iap.WebTypeComputeIamMemberArgs{
//				Project: pulumi.Any(google_project_service.Project_service.Project),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewWebTypeComputeIamMember(ctx, "member", &iap.WebTypeComputeIamMemberArgs{
//				Project: pulumi.Any(google_project_service.Project_service.Project),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/compute * {{project}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webtypecompute IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/webTypeComputeIamMember:WebTypeComputeIamMember editor "projects/{{project}}/iap_web/compute roles/iap.httpsResourceAccessor user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/webTypeComputeIamMember:WebTypeComputeIamMember editor "projects/{{project}}/iap_web/compute roles/iap.httpsResourceAccessor"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/webTypeComputeIamMember:WebTypeComputeIamMember editor projects/{{project}}/iap_web/compute
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebTypeComputeIamMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeComputeIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewWebTypeComputeIamMember registers a new resource with the given unique name, arguments, and options.
func NewWebTypeComputeIamMember(ctx *pulumi.Context,
	name string, args *WebTypeComputeIamMemberArgs, opts ...pulumi.ResourceOption) (*WebTypeComputeIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource WebTypeComputeIamMember
	err := ctx.RegisterResource("gcp:iap/webTypeComputeIamMember:WebTypeComputeIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTypeComputeIamMember gets an existing WebTypeComputeIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTypeComputeIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTypeComputeIamMemberState, opts ...pulumi.ResourceOption) (*WebTypeComputeIamMember, error) {
	var resource WebTypeComputeIamMember
	err := ctx.ReadResource("gcp:iap/webTypeComputeIamMember:WebTypeComputeIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTypeComputeIamMember resources.
type webTypeComputeIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebTypeComputeIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type WebTypeComputeIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeComputeIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (WebTypeComputeIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeComputeIamMemberState)(nil)).Elem()
}

type webTypeComputeIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebTypeComputeIamMemberCondition `pulumi:"condition"`
	Member    string                            `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a WebTypeComputeIamMember resource.
type WebTypeComputeIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeComputeIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (WebTypeComputeIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeComputeIamMemberArgs)(nil)).Elem()
}

type WebTypeComputeIamMemberInput interface {
	pulumi.Input

	ToWebTypeComputeIamMemberOutput() WebTypeComputeIamMemberOutput
	ToWebTypeComputeIamMemberOutputWithContext(ctx context.Context) WebTypeComputeIamMemberOutput
}

func (*WebTypeComputeIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTypeComputeIamMember)(nil)).Elem()
}

func (i *WebTypeComputeIamMember) ToWebTypeComputeIamMemberOutput() WebTypeComputeIamMemberOutput {
	return i.ToWebTypeComputeIamMemberOutputWithContext(context.Background())
}

func (i *WebTypeComputeIamMember) ToWebTypeComputeIamMemberOutputWithContext(ctx context.Context) WebTypeComputeIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamMemberOutput)
}

// WebTypeComputeIamMemberArrayInput is an input type that accepts WebTypeComputeIamMemberArray and WebTypeComputeIamMemberArrayOutput values.
// You can construct a concrete instance of `WebTypeComputeIamMemberArrayInput` via:
//
//	WebTypeComputeIamMemberArray{ WebTypeComputeIamMemberArgs{...} }
type WebTypeComputeIamMemberArrayInput interface {
	pulumi.Input

	ToWebTypeComputeIamMemberArrayOutput() WebTypeComputeIamMemberArrayOutput
	ToWebTypeComputeIamMemberArrayOutputWithContext(context.Context) WebTypeComputeIamMemberArrayOutput
}

type WebTypeComputeIamMemberArray []WebTypeComputeIamMemberInput

func (WebTypeComputeIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebTypeComputeIamMember)(nil)).Elem()
}

func (i WebTypeComputeIamMemberArray) ToWebTypeComputeIamMemberArrayOutput() WebTypeComputeIamMemberArrayOutput {
	return i.ToWebTypeComputeIamMemberArrayOutputWithContext(context.Background())
}

func (i WebTypeComputeIamMemberArray) ToWebTypeComputeIamMemberArrayOutputWithContext(ctx context.Context) WebTypeComputeIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamMemberArrayOutput)
}

// WebTypeComputeIamMemberMapInput is an input type that accepts WebTypeComputeIamMemberMap and WebTypeComputeIamMemberMapOutput values.
// You can construct a concrete instance of `WebTypeComputeIamMemberMapInput` via:
//
//	WebTypeComputeIamMemberMap{ "key": WebTypeComputeIamMemberArgs{...} }
type WebTypeComputeIamMemberMapInput interface {
	pulumi.Input

	ToWebTypeComputeIamMemberMapOutput() WebTypeComputeIamMemberMapOutput
	ToWebTypeComputeIamMemberMapOutputWithContext(context.Context) WebTypeComputeIamMemberMapOutput
}

type WebTypeComputeIamMemberMap map[string]WebTypeComputeIamMemberInput

func (WebTypeComputeIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebTypeComputeIamMember)(nil)).Elem()
}

func (i WebTypeComputeIamMemberMap) ToWebTypeComputeIamMemberMapOutput() WebTypeComputeIamMemberMapOutput {
	return i.ToWebTypeComputeIamMemberMapOutputWithContext(context.Background())
}

func (i WebTypeComputeIamMemberMap) ToWebTypeComputeIamMemberMapOutputWithContext(ctx context.Context) WebTypeComputeIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamMemberMapOutput)
}

type WebTypeComputeIamMemberOutput struct{ *pulumi.OutputState }

func (WebTypeComputeIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTypeComputeIamMember)(nil)).Elem()
}

func (o WebTypeComputeIamMemberOutput) ToWebTypeComputeIamMemberOutput() WebTypeComputeIamMemberOutput {
	return o
}

func (o WebTypeComputeIamMemberOutput) ToWebTypeComputeIamMemberOutputWithContext(ctx context.Context) WebTypeComputeIamMemberOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o WebTypeComputeIamMemberOutput) Condition() WebTypeComputeIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *WebTypeComputeIamMember) WebTypeComputeIamMemberConditionPtrOutput { return v.Condition }).(WebTypeComputeIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o WebTypeComputeIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTypeComputeIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o WebTypeComputeIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTypeComputeIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o WebTypeComputeIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTypeComputeIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o WebTypeComputeIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTypeComputeIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type WebTypeComputeIamMemberArrayOutput struct{ *pulumi.OutputState }

func (WebTypeComputeIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebTypeComputeIamMember)(nil)).Elem()
}

func (o WebTypeComputeIamMemberArrayOutput) ToWebTypeComputeIamMemberArrayOutput() WebTypeComputeIamMemberArrayOutput {
	return o
}

func (o WebTypeComputeIamMemberArrayOutput) ToWebTypeComputeIamMemberArrayOutputWithContext(ctx context.Context) WebTypeComputeIamMemberArrayOutput {
	return o
}

func (o WebTypeComputeIamMemberArrayOutput) Index(i pulumi.IntInput) WebTypeComputeIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebTypeComputeIamMember {
		return vs[0].([]*WebTypeComputeIamMember)[vs[1].(int)]
	}).(WebTypeComputeIamMemberOutput)
}

type WebTypeComputeIamMemberMapOutput struct{ *pulumi.OutputState }

func (WebTypeComputeIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebTypeComputeIamMember)(nil)).Elem()
}

func (o WebTypeComputeIamMemberMapOutput) ToWebTypeComputeIamMemberMapOutput() WebTypeComputeIamMemberMapOutput {
	return o
}

func (o WebTypeComputeIamMemberMapOutput) ToWebTypeComputeIamMemberMapOutputWithContext(ctx context.Context) WebTypeComputeIamMemberMapOutput {
	return o
}

func (o WebTypeComputeIamMemberMapOutput) MapIndex(k pulumi.StringInput) WebTypeComputeIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebTypeComputeIamMember {
		return vs[0].(map[string]*WebTypeComputeIamMember)[vs[1].(string)]
	}).(WebTypeComputeIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebTypeComputeIamMemberInput)(nil)).Elem(), &WebTypeComputeIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebTypeComputeIamMemberArrayInput)(nil)).Elem(), WebTypeComputeIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebTypeComputeIamMemberMapInput)(nil)).Elem(), WebTypeComputeIamMemberMap{})
	pulumi.RegisterOutputType(WebTypeComputeIamMemberOutput{})
	pulumi.RegisterOutputType(WebTypeComputeIamMemberArrayOutput{})
	pulumi.RegisterOutputType(WebTypeComputeIamMemberMapOutput{})
}

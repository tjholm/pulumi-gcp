// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeAppEngine. Each of these resources serves a different use case:
//
// * `iap.WebTypeAppEngingIamPolicy`: Authoritative. Sets the IAM policy for the webtypeappengine and replaces any existing policy already attached.
// * `iap.WebTypeAppEngingIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypeappengine are preserved.
// * `iap.WebTypeAppEngingIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypeappengine are preserved.
//
// > **Note:** `iap.WebTypeAppEngingIamPolicy` **cannot** be used in conjunction with `iap.WebTypeAppEngingIamBinding` and `iap.WebTypeAppEngingIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebTypeAppEngingIamBinding` resources **can be** used in conjunction with `iap.WebTypeAppEngingIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_web\_type\_app\_engine\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iap.httpsResourceAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewWebTypeAppEngingIamPolicy(ctx, "policy", &iap.WebTypeAppEngingIamPolicyArgs{
// 			Project:    pulumi.Any(google_app_engine_application.App.Project),
// 			AppId:      pulumi.Any(google_app_engine_application.App.App_id),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iap.httpsResourceAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewWebTypeAppEngingIamPolicy(ctx, "policy", &iap.WebTypeAppEngingIamPolicyArgs{
// 			Project:    pulumi.Any(google_app_engine_application.App.Project),
// 			AppId:      pulumi.Any(google_app_engine_application.App.App_id),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_web\_type\_app\_engine\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebTypeAppEngingIamBinding(ctx, "binding", &iap.WebTypeAppEngingIamBindingArgs{
// 			Project: pulumi.Any(google_app_engine_application.App.Project),
// 			AppId:   pulumi.Any(google_app_engine_application.App.App_id),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebTypeAppEngingIamBinding(ctx, "binding", &iap.WebTypeAppEngingIamBindingArgs{
// 			Project: pulumi.Any(google_app_engine_application.App.Project),
// 			AppId:   pulumi.Any(google_app_engine_application.App.App_id),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &iap.WebTypeAppEngingIamBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_web\_type\_app\_engine\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebTypeAppEngingIamMember(ctx, "member", &iap.WebTypeAppEngingIamMemberArgs{
// 			Project: pulumi.Any(google_app_engine_application.App.Project),
// 			AppId:   pulumi.Any(google_app_engine_application.App.App_id),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebTypeAppEngingIamMember(ctx, "member", &iap.WebTypeAppEngingIamMemberArgs{
// 			Project: pulumi.Any(google_app_engine_application.App.Project),
// 			AppId:   pulumi.Any(google_app_engine_application.App.App_id),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Condition: &iap.WebTypeAppEngingIamMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}} * {{project}}/{{appId}} * {{appId}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webtypeappengine IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember editor "projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember editor "projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember editor projects/{{project}}/iap_web/appengine-{{appId}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebTypeAppEngingIamMember struct {
	pulumi.CustomResourceState

	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeAppEngingIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewWebTypeAppEngingIamMember registers a new resource with the given unique name, arguments, and options.
func NewWebTypeAppEngingIamMember(ctx *pulumi.Context,
	name string, args *WebTypeAppEngingIamMemberArgs, opts ...pulumi.ResourceOption) (*WebTypeAppEngingIamMember, error) {
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
	var resource WebTypeAppEngingIamMember
	err := ctx.RegisterResource("gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTypeAppEngingIamMember gets an existing WebTypeAppEngingIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTypeAppEngingIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTypeAppEngingIamMemberState, opts ...pulumi.ResourceOption) (*WebTypeAppEngingIamMember, error) {
	var resource WebTypeAppEngingIamMember
	err := ctx.ReadResource("gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTypeAppEngingIamMember resources.
type webTypeAppEngingIamMemberState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId *string `pulumi:"appId"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebTypeAppEngingIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type WebTypeAppEngingIamMemberState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringPtrInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeAppEngingIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (WebTypeAppEngingIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeAppEngingIamMemberState)(nil)).Elem()
}

type webTypeAppEngingIamMemberArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId string `pulumi:"appId"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebTypeAppEngingIamMemberCondition `pulumi:"condition"`
	Member    string                              `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a WebTypeAppEngingIamMember resource.
type WebTypeAppEngingIamMemberArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeAppEngingIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (WebTypeAppEngingIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeAppEngingIamMemberArgs)(nil)).Elem()
}

type WebTypeAppEngingIamMemberInput interface {
	pulumi.Input

	ToWebTypeAppEngingIamMemberOutput() WebTypeAppEngingIamMemberOutput
	ToWebTypeAppEngingIamMemberOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberOutput
}

func (*WebTypeAppEngingIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTypeAppEngingIamMember)(nil))
}

func (i *WebTypeAppEngingIamMember) ToWebTypeAppEngingIamMemberOutput() WebTypeAppEngingIamMemberOutput {
	return i.ToWebTypeAppEngingIamMemberOutputWithContext(context.Background())
}

func (i *WebTypeAppEngingIamMember) ToWebTypeAppEngingIamMemberOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamMemberOutput)
}

func (i *WebTypeAppEngingIamMember) ToWebTypeAppEngingIamMemberPtrOutput() WebTypeAppEngingIamMemberPtrOutput {
	return i.ToWebTypeAppEngingIamMemberPtrOutputWithContext(context.Background())
}

func (i *WebTypeAppEngingIamMember) ToWebTypeAppEngingIamMemberPtrOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamMemberPtrOutput)
}

type WebTypeAppEngingIamMemberPtrInput interface {
	pulumi.Input

	ToWebTypeAppEngingIamMemberPtrOutput() WebTypeAppEngingIamMemberPtrOutput
	ToWebTypeAppEngingIamMemberPtrOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberPtrOutput
}

type webTypeAppEngingIamMemberPtrType WebTypeAppEngingIamMemberArgs

func (*webTypeAppEngingIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTypeAppEngingIamMember)(nil))
}

func (i *webTypeAppEngingIamMemberPtrType) ToWebTypeAppEngingIamMemberPtrOutput() WebTypeAppEngingIamMemberPtrOutput {
	return i.ToWebTypeAppEngingIamMemberPtrOutputWithContext(context.Background())
}

func (i *webTypeAppEngingIamMemberPtrType) ToWebTypeAppEngingIamMemberPtrOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamMemberPtrOutput)
}

// WebTypeAppEngingIamMemberArrayInput is an input type that accepts WebTypeAppEngingIamMemberArray and WebTypeAppEngingIamMemberArrayOutput values.
// You can construct a concrete instance of `WebTypeAppEngingIamMemberArrayInput` via:
//
//          WebTypeAppEngingIamMemberArray{ WebTypeAppEngingIamMemberArgs{...} }
type WebTypeAppEngingIamMemberArrayInput interface {
	pulumi.Input

	ToWebTypeAppEngingIamMemberArrayOutput() WebTypeAppEngingIamMemberArrayOutput
	ToWebTypeAppEngingIamMemberArrayOutputWithContext(context.Context) WebTypeAppEngingIamMemberArrayOutput
}

type WebTypeAppEngingIamMemberArray []WebTypeAppEngingIamMemberInput

func (WebTypeAppEngingIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WebTypeAppEngingIamMember)(nil))
}

func (i WebTypeAppEngingIamMemberArray) ToWebTypeAppEngingIamMemberArrayOutput() WebTypeAppEngingIamMemberArrayOutput {
	return i.ToWebTypeAppEngingIamMemberArrayOutputWithContext(context.Background())
}

func (i WebTypeAppEngingIamMemberArray) ToWebTypeAppEngingIamMemberArrayOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamMemberArrayOutput)
}

// WebTypeAppEngingIamMemberMapInput is an input type that accepts WebTypeAppEngingIamMemberMap and WebTypeAppEngingIamMemberMapOutput values.
// You can construct a concrete instance of `WebTypeAppEngingIamMemberMapInput` via:
//
//          WebTypeAppEngingIamMemberMap{ "key": WebTypeAppEngingIamMemberArgs{...} }
type WebTypeAppEngingIamMemberMapInput interface {
	pulumi.Input

	ToWebTypeAppEngingIamMemberMapOutput() WebTypeAppEngingIamMemberMapOutput
	ToWebTypeAppEngingIamMemberMapOutputWithContext(context.Context) WebTypeAppEngingIamMemberMapOutput
}

type WebTypeAppEngingIamMemberMap map[string]WebTypeAppEngingIamMemberInput

func (WebTypeAppEngingIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WebTypeAppEngingIamMember)(nil))
}

func (i WebTypeAppEngingIamMemberMap) ToWebTypeAppEngingIamMemberMapOutput() WebTypeAppEngingIamMemberMapOutput {
	return i.ToWebTypeAppEngingIamMemberMapOutputWithContext(context.Background())
}

func (i WebTypeAppEngingIamMemberMap) ToWebTypeAppEngingIamMemberMapOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamMemberMapOutput)
}

type WebTypeAppEngingIamMemberOutput struct {
	*pulumi.OutputState
}

func (WebTypeAppEngingIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTypeAppEngingIamMember)(nil))
}

func (o WebTypeAppEngingIamMemberOutput) ToWebTypeAppEngingIamMemberOutput() WebTypeAppEngingIamMemberOutput {
	return o
}

func (o WebTypeAppEngingIamMemberOutput) ToWebTypeAppEngingIamMemberOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberOutput {
	return o
}

func (o WebTypeAppEngingIamMemberOutput) ToWebTypeAppEngingIamMemberPtrOutput() WebTypeAppEngingIamMemberPtrOutput {
	return o.ToWebTypeAppEngingIamMemberPtrOutputWithContext(context.Background())
}

func (o WebTypeAppEngingIamMemberOutput) ToWebTypeAppEngingIamMemberPtrOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberPtrOutput {
	return o.ApplyT(func(v WebTypeAppEngingIamMember) *WebTypeAppEngingIamMember {
		return &v
	}).(WebTypeAppEngingIamMemberPtrOutput)
}

type WebTypeAppEngingIamMemberPtrOutput struct {
	*pulumi.OutputState
}

func (WebTypeAppEngingIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTypeAppEngingIamMember)(nil))
}

func (o WebTypeAppEngingIamMemberPtrOutput) ToWebTypeAppEngingIamMemberPtrOutput() WebTypeAppEngingIamMemberPtrOutput {
	return o
}

func (o WebTypeAppEngingIamMemberPtrOutput) ToWebTypeAppEngingIamMemberPtrOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberPtrOutput {
	return o
}

type WebTypeAppEngingIamMemberArrayOutput struct{ *pulumi.OutputState }

func (WebTypeAppEngingIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTypeAppEngingIamMember)(nil))
}

func (o WebTypeAppEngingIamMemberArrayOutput) ToWebTypeAppEngingIamMemberArrayOutput() WebTypeAppEngingIamMemberArrayOutput {
	return o
}

func (o WebTypeAppEngingIamMemberArrayOutput) ToWebTypeAppEngingIamMemberArrayOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberArrayOutput {
	return o
}

func (o WebTypeAppEngingIamMemberArrayOutput) Index(i pulumi.IntInput) WebTypeAppEngingIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebTypeAppEngingIamMember {
		return vs[0].([]WebTypeAppEngingIamMember)[vs[1].(int)]
	}).(WebTypeAppEngingIamMemberOutput)
}

type WebTypeAppEngingIamMemberMapOutput struct{ *pulumi.OutputState }

func (WebTypeAppEngingIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebTypeAppEngingIamMember)(nil))
}

func (o WebTypeAppEngingIamMemberMapOutput) ToWebTypeAppEngingIamMemberMapOutput() WebTypeAppEngingIamMemberMapOutput {
	return o
}

func (o WebTypeAppEngingIamMemberMapOutput) ToWebTypeAppEngingIamMemberMapOutputWithContext(ctx context.Context) WebTypeAppEngingIamMemberMapOutput {
	return o
}

func (o WebTypeAppEngingIamMemberMapOutput) MapIndex(k pulumi.StringInput) WebTypeAppEngingIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebTypeAppEngingIamMember {
		return vs[0].(map[string]WebTypeAppEngingIamMember)[vs[1].(string)]
	}).(WebTypeAppEngingIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(WebTypeAppEngingIamMemberOutput{})
	pulumi.RegisterOutputType(WebTypeAppEngingIamMemberPtrOutput{})
	pulumi.RegisterOutputType(WebTypeAppEngingIamMemberArrayOutput{})
	pulumi.RegisterOutputType(WebTypeAppEngingIamMemberMapOutput{})
}

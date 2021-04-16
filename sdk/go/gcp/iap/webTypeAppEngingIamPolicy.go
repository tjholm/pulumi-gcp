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
//  $ pulumi import gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy editor projects/{{project}}/iap_web/appengine-{{appId}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebTypeAppEngingIamPolicy struct {
	pulumi.CustomResourceState

	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewWebTypeAppEngingIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebTypeAppEngingIamPolicy(ctx *pulumi.Context,
	name string, args *WebTypeAppEngingIamPolicyArgs, opts ...pulumi.ResourceOption) (*WebTypeAppEngingIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource WebTypeAppEngingIamPolicy
	err := ctx.RegisterResource("gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTypeAppEngingIamPolicy gets an existing WebTypeAppEngingIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTypeAppEngingIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTypeAppEngingIamPolicyState, opts ...pulumi.ResourceOption) (*WebTypeAppEngingIamPolicy, error) {
	var resource WebTypeAppEngingIamPolicy
	err := ctx.ReadResource("gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTypeAppEngingIamPolicy resources.
type webTypeAppEngingIamPolicyState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId *string `pulumi:"appId"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type WebTypeAppEngingIamPolicyState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebTypeAppEngingIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeAppEngingIamPolicyState)(nil)).Elem()
}

type webTypeAppEngingIamPolicyArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId string `pulumi:"appId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a WebTypeAppEngingIamPolicy resource.
type WebTypeAppEngingIamPolicyArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebTypeAppEngingIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeAppEngingIamPolicyArgs)(nil)).Elem()
}

type WebTypeAppEngingIamPolicyInput interface {
	pulumi.Input

	ToWebTypeAppEngingIamPolicyOutput() WebTypeAppEngingIamPolicyOutput
	ToWebTypeAppEngingIamPolicyOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyOutput
}

func (*WebTypeAppEngingIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTypeAppEngingIamPolicy)(nil))
}

func (i *WebTypeAppEngingIamPolicy) ToWebTypeAppEngingIamPolicyOutput() WebTypeAppEngingIamPolicyOutput {
	return i.ToWebTypeAppEngingIamPolicyOutputWithContext(context.Background())
}

func (i *WebTypeAppEngingIamPolicy) ToWebTypeAppEngingIamPolicyOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamPolicyOutput)
}

func (i *WebTypeAppEngingIamPolicy) ToWebTypeAppEngingIamPolicyPtrOutput() WebTypeAppEngingIamPolicyPtrOutput {
	return i.ToWebTypeAppEngingIamPolicyPtrOutputWithContext(context.Background())
}

func (i *WebTypeAppEngingIamPolicy) ToWebTypeAppEngingIamPolicyPtrOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamPolicyPtrOutput)
}

type WebTypeAppEngingIamPolicyPtrInput interface {
	pulumi.Input

	ToWebTypeAppEngingIamPolicyPtrOutput() WebTypeAppEngingIamPolicyPtrOutput
	ToWebTypeAppEngingIamPolicyPtrOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyPtrOutput
}

type webTypeAppEngingIamPolicyPtrType WebTypeAppEngingIamPolicyArgs

func (*webTypeAppEngingIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTypeAppEngingIamPolicy)(nil))
}

func (i *webTypeAppEngingIamPolicyPtrType) ToWebTypeAppEngingIamPolicyPtrOutput() WebTypeAppEngingIamPolicyPtrOutput {
	return i.ToWebTypeAppEngingIamPolicyPtrOutputWithContext(context.Background())
}

func (i *webTypeAppEngingIamPolicyPtrType) ToWebTypeAppEngingIamPolicyPtrOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamPolicyPtrOutput)
}

// WebTypeAppEngingIamPolicyArrayInput is an input type that accepts WebTypeAppEngingIamPolicyArray and WebTypeAppEngingIamPolicyArrayOutput values.
// You can construct a concrete instance of `WebTypeAppEngingIamPolicyArrayInput` via:
//
//          WebTypeAppEngingIamPolicyArray{ WebTypeAppEngingIamPolicyArgs{...} }
type WebTypeAppEngingIamPolicyArrayInput interface {
	pulumi.Input

	ToWebTypeAppEngingIamPolicyArrayOutput() WebTypeAppEngingIamPolicyArrayOutput
	ToWebTypeAppEngingIamPolicyArrayOutputWithContext(context.Context) WebTypeAppEngingIamPolicyArrayOutput
}

type WebTypeAppEngingIamPolicyArray []WebTypeAppEngingIamPolicyInput

func (WebTypeAppEngingIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WebTypeAppEngingIamPolicy)(nil))
}

func (i WebTypeAppEngingIamPolicyArray) ToWebTypeAppEngingIamPolicyArrayOutput() WebTypeAppEngingIamPolicyArrayOutput {
	return i.ToWebTypeAppEngingIamPolicyArrayOutputWithContext(context.Background())
}

func (i WebTypeAppEngingIamPolicyArray) ToWebTypeAppEngingIamPolicyArrayOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamPolicyArrayOutput)
}

// WebTypeAppEngingIamPolicyMapInput is an input type that accepts WebTypeAppEngingIamPolicyMap and WebTypeAppEngingIamPolicyMapOutput values.
// You can construct a concrete instance of `WebTypeAppEngingIamPolicyMapInput` via:
//
//          WebTypeAppEngingIamPolicyMap{ "key": WebTypeAppEngingIamPolicyArgs{...} }
type WebTypeAppEngingIamPolicyMapInput interface {
	pulumi.Input

	ToWebTypeAppEngingIamPolicyMapOutput() WebTypeAppEngingIamPolicyMapOutput
	ToWebTypeAppEngingIamPolicyMapOutputWithContext(context.Context) WebTypeAppEngingIamPolicyMapOutput
}

type WebTypeAppEngingIamPolicyMap map[string]WebTypeAppEngingIamPolicyInput

func (WebTypeAppEngingIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WebTypeAppEngingIamPolicy)(nil))
}

func (i WebTypeAppEngingIamPolicyMap) ToWebTypeAppEngingIamPolicyMapOutput() WebTypeAppEngingIamPolicyMapOutput {
	return i.ToWebTypeAppEngingIamPolicyMapOutputWithContext(context.Background())
}

func (i WebTypeAppEngingIamPolicyMap) ToWebTypeAppEngingIamPolicyMapOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamPolicyMapOutput)
}

type WebTypeAppEngingIamPolicyOutput struct {
	*pulumi.OutputState
}

func (WebTypeAppEngingIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTypeAppEngingIamPolicy)(nil))
}

func (o WebTypeAppEngingIamPolicyOutput) ToWebTypeAppEngingIamPolicyOutput() WebTypeAppEngingIamPolicyOutput {
	return o
}

func (o WebTypeAppEngingIamPolicyOutput) ToWebTypeAppEngingIamPolicyOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyOutput {
	return o
}

func (o WebTypeAppEngingIamPolicyOutput) ToWebTypeAppEngingIamPolicyPtrOutput() WebTypeAppEngingIamPolicyPtrOutput {
	return o.ToWebTypeAppEngingIamPolicyPtrOutputWithContext(context.Background())
}

func (o WebTypeAppEngingIamPolicyOutput) ToWebTypeAppEngingIamPolicyPtrOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyPtrOutput {
	return o.ApplyT(func(v WebTypeAppEngingIamPolicy) *WebTypeAppEngingIamPolicy {
		return &v
	}).(WebTypeAppEngingIamPolicyPtrOutput)
}

type WebTypeAppEngingIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (WebTypeAppEngingIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTypeAppEngingIamPolicy)(nil))
}

func (o WebTypeAppEngingIamPolicyPtrOutput) ToWebTypeAppEngingIamPolicyPtrOutput() WebTypeAppEngingIamPolicyPtrOutput {
	return o
}

func (o WebTypeAppEngingIamPolicyPtrOutput) ToWebTypeAppEngingIamPolicyPtrOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyPtrOutput {
	return o
}

type WebTypeAppEngingIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (WebTypeAppEngingIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTypeAppEngingIamPolicy)(nil))
}

func (o WebTypeAppEngingIamPolicyArrayOutput) ToWebTypeAppEngingIamPolicyArrayOutput() WebTypeAppEngingIamPolicyArrayOutput {
	return o
}

func (o WebTypeAppEngingIamPolicyArrayOutput) ToWebTypeAppEngingIamPolicyArrayOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyArrayOutput {
	return o
}

func (o WebTypeAppEngingIamPolicyArrayOutput) Index(i pulumi.IntInput) WebTypeAppEngingIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebTypeAppEngingIamPolicy {
		return vs[0].([]WebTypeAppEngingIamPolicy)[vs[1].(int)]
	}).(WebTypeAppEngingIamPolicyOutput)
}

type WebTypeAppEngingIamPolicyMapOutput struct{ *pulumi.OutputState }

func (WebTypeAppEngingIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebTypeAppEngingIamPolicy)(nil))
}

func (o WebTypeAppEngingIamPolicyMapOutput) ToWebTypeAppEngingIamPolicyMapOutput() WebTypeAppEngingIamPolicyMapOutput {
	return o
}

func (o WebTypeAppEngingIamPolicyMapOutput) ToWebTypeAppEngingIamPolicyMapOutputWithContext(ctx context.Context) WebTypeAppEngingIamPolicyMapOutput {
	return o
}

func (o WebTypeAppEngingIamPolicyMapOutput) MapIndex(k pulumi.StringInput) WebTypeAppEngingIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebTypeAppEngingIamPolicy {
		return vs[0].(map[string]WebTypeAppEngingIamPolicy)[vs[1].(string)]
	}).(WebTypeAppEngingIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(WebTypeAppEngingIamPolicyOutput{})
	pulumi.RegisterOutputType(WebTypeAppEngingIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(WebTypeAppEngingIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(WebTypeAppEngingIamPolicyMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy AppEngineService. Each of these resources serves a different use case:
//
// * `iap.AppEngineServiceIamPolicy`: Authoritative. Sets the IAM policy for the appengineservice and replaces any existing policy already attached.
// * `iap.AppEngineServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the appengineservice are preserved.
// * `iap.AppEngineServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the appengineservice are preserved.
//
// > **Note:** `iap.AppEngineServiceIamPolicy` **cannot** be used in conjunction with `iap.AppEngineServiceIamBinding` and `iap.AppEngineServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.AppEngineServiceIamBinding` resources **can be** used in conjunction with `iap.AppEngineServiceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_app\_engine\_service\_iam\_policy
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
//			_, err = iap.NewAppEngineServiceIamPolicy(ctx, "policy", &iap.AppEngineServiceIamPolicyArgs{
//				Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
//			_, err = iap.NewAppEngineServiceIamPolicy(ctx, "policy", &iap.AppEngineServiceIamPolicyArgs{
//				Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
// ## google\_iap\_app\_engine\_service\_iam\_binding
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
//			_, err := iap.NewAppEngineServiceIamBinding(ctx, "binding", &iap.AppEngineServiceIamBindingArgs{
//				AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
//				Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
//			_, err := iap.NewAppEngineServiceIamBinding(ctx, "binding", &iap.AppEngineServiceIamBindingArgs{
//				AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Condition: &iap.AppEngineServiceIamBindingConditionArgs{
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
//					Title:       pulumi.String("expires_after_2019_12_31"),
//				},
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
//				Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## google\_iap\_app\_engine\_service\_iam\_member
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
//			_, err := iap.NewAppEngineServiceIamMember(ctx, "member", &iap.AppEngineServiceIamMemberArgs{
//				AppId:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Member:  pulumi.String("user:jane@example.com"),
//				Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
//				Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
//			_, err := iap.NewAppEngineServiceIamMember(ctx, "member", &iap.AppEngineServiceIamMemberArgs{
//				AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Condition: &iap.AppEngineServiceIamMemberConditionArgs{
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
//					Title:       pulumi.String("expires_after_2019_12_31"),
//				},
//				Member:  pulumi.String("user:jane@example.com"),
//				Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
//				Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
//				Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} * {{project}}/{{appId}}/{{service}} * {{appId}}/{{service}} * {{service}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy appengineservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} roles/iap.httpsResourceAccessor user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} roles/iap.httpsResourceAccessor"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy editor projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AppEngineServiceIamPolicy struct {
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
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewAppEngineServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewAppEngineServiceIamPolicy(ctx *pulumi.Context,
	name string, args *AppEngineServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*AppEngineServiceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource AppEngineServiceIamPolicy
	err := ctx.RegisterResource("gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppEngineServiceIamPolicy gets an existing AppEngineServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppEngineServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppEngineServiceIamPolicyState, opts ...pulumi.ResourceOption) (*AppEngineServiceIamPolicy, error) {
	var resource AppEngineServiceIamPolicy
	err := ctx.ReadResource("gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppEngineServiceIamPolicy resources.
type appEngineServiceIamPolicyState struct {
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
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service *string `pulumi:"service"`
}

type AppEngineServiceIamPolicyState struct {
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
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringPtrInput
}

func (AppEngineServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineServiceIamPolicyState)(nil)).Elem()
}

type appEngineServiceIamPolicyArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId string `pulumi:"appId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a AppEngineServiceIamPolicy resource.
type AppEngineServiceIamPolicyArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringInput
}

func (AppEngineServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineServiceIamPolicyArgs)(nil)).Elem()
}

type AppEngineServiceIamPolicyInput interface {
	pulumi.Input

	ToAppEngineServiceIamPolicyOutput() AppEngineServiceIamPolicyOutput
	ToAppEngineServiceIamPolicyOutputWithContext(ctx context.Context) AppEngineServiceIamPolicyOutput
}

func (*AppEngineServiceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineServiceIamPolicy)(nil)).Elem()
}

func (i *AppEngineServiceIamPolicy) ToAppEngineServiceIamPolicyOutput() AppEngineServiceIamPolicyOutput {
	return i.ToAppEngineServiceIamPolicyOutputWithContext(context.Background())
}

func (i *AppEngineServiceIamPolicy) ToAppEngineServiceIamPolicyOutputWithContext(ctx context.Context) AppEngineServiceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineServiceIamPolicyOutput)
}

// AppEngineServiceIamPolicyArrayInput is an input type that accepts AppEngineServiceIamPolicyArray and AppEngineServiceIamPolicyArrayOutput values.
// You can construct a concrete instance of `AppEngineServiceIamPolicyArrayInput` via:
//
//	AppEngineServiceIamPolicyArray{ AppEngineServiceIamPolicyArgs{...} }
type AppEngineServiceIamPolicyArrayInput interface {
	pulumi.Input

	ToAppEngineServiceIamPolicyArrayOutput() AppEngineServiceIamPolicyArrayOutput
	ToAppEngineServiceIamPolicyArrayOutputWithContext(context.Context) AppEngineServiceIamPolicyArrayOutput
}

type AppEngineServiceIamPolicyArray []AppEngineServiceIamPolicyInput

func (AppEngineServiceIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppEngineServiceIamPolicy)(nil)).Elem()
}

func (i AppEngineServiceIamPolicyArray) ToAppEngineServiceIamPolicyArrayOutput() AppEngineServiceIamPolicyArrayOutput {
	return i.ToAppEngineServiceIamPolicyArrayOutputWithContext(context.Background())
}

func (i AppEngineServiceIamPolicyArray) ToAppEngineServiceIamPolicyArrayOutputWithContext(ctx context.Context) AppEngineServiceIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineServiceIamPolicyArrayOutput)
}

// AppEngineServiceIamPolicyMapInput is an input type that accepts AppEngineServiceIamPolicyMap and AppEngineServiceIamPolicyMapOutput values.
// You can construct a concrete instance of `AppEngineServiceIamPolicyMapInput` via:
//
//	AppEngineServiceIamPolicyMap{ "key": AppEngineServiceIamPolicyArgs{...} }
type AppEngineServiceIamPolicyMapInput interface {
	pulumi.Input

	ToAppEngineServiceIamPolicyMapOutput() AppEngineServiceIamPolicyMapOutput
	ToAppEngineServiceIamPolicyMapOutputWithContext(context.Context) AppEngineServiceIamPolicyMapOutput
}

type AppEngineServiceIamPolicyMap map[string]AppEngineServiceIamPolicyInput

func (AppEngineServiceIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppEngineServiceIamPolicy)(nil)).Elem()
}

func (i AppEngineServiceIamPolicyMap) ToAppEngineServiceIamPolicyMapOutput() AppEngineServiceIamPolicyMapOutput {
	return i.ToAppEngineServiceIamPolicyMapOutputWithContext(context.Background())
}

func (i AppEngineServiceIamPolicyMap) ToAppEngineServiceIamPolicyMapOutputWithContext(ctx context.Context) AppEngineServiceIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineServiceIamPolicyMapOutput)
}

type AppEngineServiceIamPolicyOutput struct{ *pulumi.OutputState }

func (AppEngineServiceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineServiceIamPolicy)(nil)).Elem()
}

func (o AppEngineServiceIamPolicyOutput) ToAppEngineServiceIamPolicyOutput() AppEngineServiceIamPolicyOutput {
	return o
}

func (o AppEngineServiceIamPolicyOutput) ToAppEngineServiceIamPolicyOutputWithContext(ctx context.Context) AppEngineServiceIamPolicyOutput {
	return o
}

// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
func (o AppEngineServiceIamPolicyOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineServiceIamPolicy) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o AppEngineServiceIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineServiceIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o AppEngineServiceIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineServiceIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o AppEngineServiceIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineServiceIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
func (o AppEngineServiceIamPolicyOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineServiceIamPolicy) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

type AppEngineServiceIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (AppEngineServiceIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppEngineServiceIamPolicy)(nil)).Elem()
}

func (o AppEngineServiceIamPolicyArrayOutput) ToAppEngineServiceIamPolicyArrayOutput() AppEngineServiceIamPolicyArrayOutput {
	return o
}

func (o AppEngineServiceIamPolicyArrayOutput) ToAppEngineServiceIamPolicyArrayOutputWithContext(ctx context.Context) AppEngineServiceIamPolicyArrayOutput {
	return o
}

func (o AppEngineServiceIamPolicyArrayOutput) Index(i pulumi.IntInput) AppEngineServiceIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppEngineServiceIamPolicy {
		return vs[0].([]*AppEngineServiceIamPolicy)[vs[1].(int)]
	}).(AppEngineServiceIamPolicyOutput)
}

type AppEngineServiceIamPolicyMapOutput struct{ *pulumi.OutputState }

func (AppEngineServiceIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppEngineServiceIamPolicy)(nil)).Elem()
}

func (o AppEngineServiceIamPolicyMapOutput) ToAppEngineServiceIamPolicyMapOutput() AppEngineServiceIamPolicyMapOutput {
	return o
}

func (o AppEngineServiceIamPolicyMapOutput) ToAppEngineServiceIamPolicyMapOutputWithContext(ctx context.Context) AppEngineServiceIamPolicyMapOutput {
	return o
}

func (o AppEngineServiceIamPolicyMapOutput) MapIndex(k pulumi.StringInput) AppEngineServiceIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppEngineServiceIamPolicy {
		return vs[0].(map[string]*AppEngineServiceIamPolicy)[vs[1].(string)]
	}).(AppEngineServiceIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineServiceIamPolicyInput)(nil)).Elem(), &AppEngineServiceIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineServiceIamPolicyArrayInput)(nil)).Elem(), AppEngineServiceIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineServiceIamPolicyMapInput)(nil)).Elem(), AppEngineServiceIamPolicyMap{})
	pulumi.RegisterOutputType(AppEngineServiceIamPolicyOutput{})
	pulumi.RegisterOutputType(AppEngineServiceIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(AppEngineServiceIamPolicyMapOutput{})
}

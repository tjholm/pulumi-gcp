// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
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
// 		_, err = iap.NewAppEngineServiceIamPolicy(ctx, "policy", &iap.AppEngineServiceIamPolicyArgs{
// 			Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
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
// 		_, err = iap.NewAppEngineServiceIamPolicy(ctx, "policy", &iap.AppEngineServiceIamPolicyArgs{
// 			Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_app\_engine\_service\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineServiceIamBinding(ctx, "binding", &iap.AppEngineServiceIamBindingArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineServiceIamBinding(ctx, "binding", &iap.AppEngineServiceIamBindingArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Condition: &iap.AppEngineServiceIamBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_app\_engine\_service\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineServiceIamMember(ctx, "member", &iap.AppEngineServiceIamMemberArgs{
// 			AppId:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineServiceIamMember(ctx, "member", &iap.AppEngineServiceIamMemberArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Condition: &iap.AppEngineServiceIamMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} * {{project}}/{{appId}}/{{service}} * {{appId}}/{{service}} * {{service}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy appengineservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineServiceIamBinding:AppEngineServiceIamBinding editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineServiceIamBinding:AppEngineServiceIamBinding editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineServiceIamBinding:AppEngineServiceIamBinding editor projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AppEngineServiceIamBinding struct {
	pulumi.CustomResourceState

	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineServiceIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewAppEngineServiceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAppEngineServiceIamBinding(ctx *pulumi.Context,
	name string, args *AppEngineServiceIamBindingArgs, opts ...pulumi.ResourceOption) (*AppEngineServiceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource AppEngineServiceIamBinding
	err := ctx.RegisterResource("gcp:iap/appEngineServiceIamBinding:AppEngineServiceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppEngineServiceIamBinding gets an existing AppEngineServiceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppEngineServiceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppEngineServiceIamBindingState, opts ...pulumi.ResourceOption) (*AppEngineServiceIamBinding, error) {
	var resource AppEngineServiceIamBinding
	err := ctx.ReadResource("gcp:iap/appEngineServiceIamBinding:AppEngineServiceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppEngineServiceIamBinding resources.
type appEngineServiceIamBindingState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId *string `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *AppEngineServiceIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service *string `pulumi:"service"`
}

type AppEngineServiceIamBindingState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringPtrInput
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineServiceIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringPtrInput
}

func (AppEngineServiceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineServiceIamBindingState)(nil)).Elem()
}

type appEngineServiceIamBindingArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId string `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *AppEngineServiceIamBindingCondition `pulumi:"condition"`
	Members   []string                             `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a AppEngineServiceIamBinding resource.
type AppEngineServiceIamBindingArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringInput
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineServiceIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringInput
}

func (AppEngineServiceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineServiceIamBindingArgs)(nil)).Elem()
}

type AppEngineServiceIamBindingInput interface {
	pulumi.Input

	ToAppEngineServiceIamBindingOutput() AppEngineServiceIamBindingOutput
	ToAppEngineServiceIamBindingOutputWithContext(ctx context.Context) AppEngineServiceIamBindingOutput
}

func (*AppEngineServiceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineServiceIamBinding)(nil)).Elem()
}

func (i *AppEngineServiceIamBinding) ToAppEngineServiceIamBindingOutput() AppEngineServiceIamBindingOutput {
	return i.ToAppEngineServiceIamBindingOutputWithContext(context.Background())
}

func (i *AppEngineServiceIamBinding) ToAppEngineServiceIamBindingOutputWithContext(ctx context.Context) AppEngineServiceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineServiceIamBindingOutput)
}

// AppEngineServiceIamBindingArrayInput is an input type that accepts AppEngineServiceIamBindingArray and AppEngineServiceIamBindingArrayOutput values.
// You can construct a concrete instance of `AppEngineServiceIamBindingArrayInput` via:
//
//          AppEngineServiceIamBindingArray{ AppEngineServiceIamBindingArgs{...} }
type AppEngineServiceIamBindingArrayInput interface {
	pulumi.Input

	ToAppEngineServiceIamBindingArrayOutput() AppEngineServiceIamBindingArrayOutput
	ToAppEngineServiceIamBindingArrayOutputWithContext(context.Context) AppEngineServiceIamBindingArrayOutput
}

type AppEngineServiceIamBindingArray []AppEngineServiceIamBindingInput

func (AppEngineServiceIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppEngineServiceIamBinding)(nil)).Elem()
}

func (i AppEngineServiceIamBindingArray) ToAppEngineServiceIamBindingArrayOutput() AppEngineServiceIamBindingArrayOutput {
	return i.ToAppEngineServiceIamBindingArrayOutputWithContext(context.Background())
}

func (i AppEngineServiceIamBindingArray) ToAppEngineServiceIamBindingArrayOutputWithContext(ctx context.Context) AppEngineServiceIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineServiceIamBindingArrayOutput)
}

// AppEngineServiceIamBindingMapInput is an input type that accepts AppEngineServiceIamBindingMap and AppEngineServiceIamBindingMapOutput values.
// You can construct a concrete instance of `AppEngineServiceIamBindingMapInput` via:
//
//          AppEngineServiceIamBindingMap{ "key": AppEngineServiceIamBindingArgs{...} }
type AppEngineServiceIamBindingMapInput interface {
	pulumi.Input

	ToAppEngineServiceIamBindingMapOutput() AppEngineServiceIamBindingMapOutput
	ToAppEngineServiceIamBindingMapOutputWithContext(context.Context) AppEngineServiceIamBindingMapOutput
}

type AppEngineServiceIamBindingMap map[string]AppEngineServiceIamBindingInput

func (AppEngineServiceIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppEngineServiceIamBinding)(nil)).Elem()
}

func (i AppEngineServiceIamBindingMap) ToAppEngineServiceIamBindingMapOutput() AppEngineServiceIamBindingMapOutput {
	return i.ToAppEngineServiceIamBindingMapOutputWithContext(context.Background())
}

func (i AppEngineServiceIamBindingMap) ToAppEngineServiceIamBindingMapOutputWithContext(ctx context.Context) AppEngineServiceIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineServiceIamBindingMapOutput)
}

type AppEngineServiceIamBindingOutput struct{ *pulumi.OutputState }

func (AppEngineServiceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineServiceIamBinding)(nil)).Elem()
}

func (o AppEngineServiceIamBindingOutput) ToAppEngineServiceIamBindingOutput() AppEngineServiceIamBindingOutput {
	return o
}

func (o AppEngineServiceIamBindingOutput) ToAppEngineServiceIamBindingOutputWithContext(ctx context.Context) AppEngineServiceIamBindingOutput {
	return o
}

type AppEngineServiceIamBindingArrayOutput struct{ *pulumi.OutputState }

func (AppEngineServiceIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppEngineServiceIamBinding)(nil)).Elem()
}

func (o AppEngineServiceIamBindingArrayOutput) ToAppEngineServiceIamBindingArrayOutput() AppEngineServiceIamBindingArrayOutput {
	return o
}

func (o AppEngineServiceIamBindingArrayOutput) ToAppEngineServiceIamBindingArrayOutputWithContext(ctx context.Context) AppEngineServiceIamBindingArrayOutput {
	return o
}

func (o AppEngineServiceIamBindingArrayOutput) Index(i pulumi.IntInput) AppEngineServiceIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppEngineServiceIamBinding {
		return vs[0].([]*AppEngineServiceIamBinding)[vs[1].(int)]
	}).(AppEngineServiceIamBindingOutput)
}

type AppEngineServiceIamBindingMapOutput struct{ *pulumi.OutputState }

func (AppEngineServiceIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppEngineServiceIamBinding)(nil)).Elem()
}

func (o AppEngineServiceIamBindingMapOutput) ToAppEngineServiceIamBindingMapOutput() AppEngineServiceIamBindingMapOutput {
	return o
}

func (o AppEngineServiceIamBindingMapOutput) ToAppEngineServiceIamBindingMapOutputWithContext(ctx context.Context) AppEngineServiceIamBindingMapOutput {
	return o
}

func (o AppEngineServiceIamBindingMapOutput) MapIndex(k pulumi.StringInput) AppEngineServiceIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppEngineServiceIamBinding {
		return vs[0].(map[string]*AppEngineServiceIamBinding)[vs[1].(string)]
	}).(AppEngineServiceIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineServiceIamBindingInput)(nil)).Elem(), &AppEngineServiceIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineServiceIamBindingArrayInput)(nil)).Elem(), AppEngineServiceIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineServiceIamBindingMapInput)(nil)).Elem(), AppEngineServiceIamBindingMap{})
	pulumi.RegisterOutputType(AppEngineServiceIamBindingOutput{})
	pulumi.RegisterOutputType(AppEngineServiceIamBindingArrayOutput{})
	pulumi.RegisterOutputType(AppEngineServiceIamBindingMapOutput{})
}

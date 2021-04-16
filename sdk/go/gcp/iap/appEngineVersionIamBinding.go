// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
// 		_, err = iap.NewAppEngineVersionIamPolicy(ctx, "policy", &iap.AppEngineVersionIamPolicyArgs{
// 			Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId:  pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
// 		_, err = iap.NewAppEngineVersionIamPolicy(ctx, "policy", &iap.AppEngineVersionIamPolicyArgs{
// 			Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId:  pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_app\_engine\_version\_iam\_binding
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
// 		_, err := iap.NewAppEngineVersionIamBinding(ctx, "binding", &iap.AppEngineVersionIamBindingArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
// 		_, err := iap.NewAppEngineVersionIamBinding(ctx, "binding", &iap.AppEngineVersionIamBindingArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Condition: &iap.AppEngineVersionIamBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_app\_engine\_version\_iam\_member
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
// 		_, err := iap.NewAppEngineVersionIamMember(ctx, "member", &iap.AppEngineVersionIamMemberArgs{
// 			AppId:     pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
// 		_, err := iap.NewAppEngineVersionIamMember(ctx, "member", &iap.AppEngineVersionIamMemberArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Condition: &iap.AppEngineVersionIamMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} * {{project}}/{{appId}}/{{service}}/{{versionId}} * {{appId}}/{{service}}/{{versionId}} * {{version}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy appengineversion IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding editor projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AppEngineVersionIamBinding struct {
	pulumi.CustomResourceState

	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineVersionIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
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

// NewAppEngineVersionIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAppEngineVersionIamBinding(ctx *pulumi.Context,
	name string, args *AppEngineVersionIamBindingArgs, opts ...pulumi.ResourceOption) (*AppEngineVersionIamBinding, error) {
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
	if args.VersionId == nil {
		return nil, errors.New("invalid value for required argument 'VersionId'")
	}
	var resource AppEngineVersionIamBinding
	err := ctx.RegisterResource("gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppEngineVersionIamBinding gets an existing AppEngineVersionIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppEngineVersionIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppEngineVersionIamBindingState, opts ...pulumi.ResourceOption) (*AppEngineVersionIamBinding, error) {
	var resource AppEngineVersionIamBinding
	err := ctx.ReadResource("gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppEngineVersionIamBinding resources.
type appEngineVersionIamBindingState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId *string `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *AppEngineVersionIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
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

type AppEngineVersionIamBindingState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringPtrInput
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineVersionIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
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

func (AppEngineVersionIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineVersionIamBindingState)(nil)).Elem()
}

type appEngineVersionIamBindingArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId string `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *AppEngineVersionIamBindingCondition `pulumi:"condition"`
	Members   []string                             `pulumi:"members"`
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

// The set of arguments for constructing a AppEngineVersionIamBinding resource.
type AppEngineVersionIamBindingArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringInput
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineVersionIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
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

func (AppEngineVersionIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineVersionIamBindingArgs)(nil)).Elem()
}

type AppEngineVersionIamBindingInput interface {
	pulumi.Input

	ToAppEngineVersionIamBindingOutput() AppEngineVersionIamBindingOutput
	ToAppEngineVersionIamBindingOutputWithContext(ctx context.Context) AppEngineVersionIamBindingOutput
}

func (*AppEngineVersionIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*AppEngineVersionIamBinding)(nil))
}

func (i *AppEngineVersionIamBinding) ToAppEngineVersionIamBindingOutput() AppEngineVersionIamBindingOutput {
	return i.ToAppEngineVersionIamBindingOutputWithContext(context.Background())
}

func (i *AppEngineVersionIamBinding) ToAppEngineVersionIamBindingOutputWithContext(ctx context.Context) AppEngineVersionIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamBindingOutput)
}

func (i *AppEngineVersionIamBinding) ToAppEngineVersionIamBindingPtrOutput() AppEngineVersionIamBindingPtrOutput {
	return i.ToAppEngineVersionIamBindingPtrOutputWithContext(context.Background())
}

func (i *AppEngineVersionIamBinding) ToAppEngineVersionIamBindingPtrOutputWithContext(ctx context.Context) AppEngineVersionIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamBindingPtrOutput)
}

type AppEngineVersionIamBindingPtrInput interface {
	pulumi.Input

	ToAppEngineVersionIamBindingPtrOutput() AppEngineVersionIamBindingPtrOutput
	ToAppEngineVersionIamBindingPtrOutputWithContext(ctx context.Context) AppEngineVersionIamBindingPtrOutput
}

type appEngineVersionIamBindingPtrType AppEngineVersionIamBindingArgs

func (*appEngineVersionIamBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineVersionIamBinding)(nil))
}

func (i *appEngineVersionIamBindingPtrType) ToAppEngineVersionIamBindingPtrOutput() AppEngineVersionIamBindingPtrOutput {
	return i.ToAppEngineVersionIamBindingPtrOutputWithContext(context.Background())
}

func (i *appEngineVersionIamBindingPtrType) ToAppEngineVersionIamBindingPtrOutputWithContext(ctx context.Context) AppEngineVersionIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamBindingPtrOutput)
}

// AppEngineVersionIamBindingArrayInput is an input type that accepts AppEngineVersionIamBindingArray and AppEngineVersionIamBindingArrayOutput values.
// You can construct a concrete instance of `AppEngineVersionIamBindingArrayInput` via:
//
//          AppEngineVersionIamBindingArray{ AppEngineVersionIamBindingArgs{...} }
type AppEngineVersionIamBindingArrayInput interface {
	pulumi.Input

	ToAppEngineVersionIamBindingArrayOutput() AppEngineVersionIamBindingArrayOutput
	ToAppEngineVersionIamBindingArrayOutputWithContext(context.Context) AppEngineVersionIamBindingArrayOutput
}

type AppEngineVersionIamBindingArray []AppEngineVersionIamBindingInput

func (AppEngineVersionIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppEngineVersionIamBinding)(nil))
}

func (i AppEngineVersionIamBindingArray) ToAppEngineVersionIamBindingArrayOutput() AppEngineVersionIamBindingArrayOutput {
	return i.ToAppEngineVersionIamBindingArrayOutputWithContext(context.Background())
}

func (i AppEngineVersionIamBindingArray) ToAppEngineVersionIamBindingArrayOutputWithContext(ctx context.Context) AppEngineVersionIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamBindingArrayOutput)
}

// AppEngineVersionIamBindingMapInput is an input type that accepts AppEngineVersionIamBindingMap and AppEngineVersionIamBindingMapOutput values.
// You can construct a concrete instance of `AppEngineVersionIamBindingMapInput` via:
//
//          AppEngineVersionIamBindingMap{ "key": AppEngineVersionIamBindingArgs{...} }
type AppEngineVersionIamBindingMapInput interface {
	pulumi.Input

	ToAppEngineVersionIamBindingMapOutput() AppEngineVersionIamBindingMapOutput
	ToAppEngineVersionIamBindingMapOutputWithContext(context.Context) AppEngineVersionIamBindingMapOutput
}

type AppEngineVersionIamBindingMap map[string]AppEngineVersionIamBindingInput

func (AppEngineVersionIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppEngineVersionIamBinding)(nil))
}

func (i AppEngineVersionIamBindingMap) ToAppEngineVersionIamBindingMapOutput() AppEngineVersionIamBindingMapOutput {
	return i.ToAppEngineVersionIamBindingMapOutputWithContext(context.Background())
}

func (i AppEngineVersionIamBindingMap) ToAppEngineVersionIamBindingMapOutputWithContext(ctx context.Context) AppEngineVersionIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamBindingMapOutput)
}

type AppEngineVersionIamBindingOutput struct {
	*pulumi.OutputState
}

func (AppEngineVersionIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppEngineVersionIamBinding)(nil))
}

func (o AppEngineVersionIamBindingOutput) ToAppEngineVersionIamBindingOutput() AppEngineVersionIamBindingOutput {
	return o
}

func (o AppEngineVersionIamBindingOutput) ToAppEngineVersionIamBindingOutputWithContext(ctx context.Context) AppEngineVersionIamBindingOutput {
	return o
}

func (o AppEngineVersionIamBindingOutput) ToAppEngineVersionIamBindingPtrOutput() AppEngineVersionIamBindingPtrOutput {
	return o.ToAppEngineVersionIamBindingPtrOutputWithContext(context.Background())
}

func (o AppEngineVersionIamBindingOutput) ToAppEngineVersionIamBindingPtrOutputWithContext(ctx context.Context) AppEngineVersionIamBindingPtrOutput {
	return o.ApplyT(func(v AppEngineVersionIamBinding) *AppEngineVersionIamBinding {
		return &v
	}).(AppEngineVersionIamBindingPtrOutput)
}

type AppEngineVersionIamBindingPtrOutput struct {
	*pulumi.OutputState
}

func (AppEngineVersionIamBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineVersionIamBinding)(nil))
}

func (o AppEngineVersionIamBindingPtrOutput) ToAppEngineVersionIamBindingPtrOutput() AppEngineVersionIamBindingPtrOutput {
	return o
}

func (o AppEngineVersionIamBindingPtrOutput) ToAppEngineVersionIamBindingPtrOutputWithContext(ctx context.Context) AppEngineVersionIamBindingPtrOutput {
	return o
}

type AppEngineVersionIamBindingArrayOutput struct{ *pulumi.OutputState }

func (AppEngineVersionIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppEngineVersionIamBinding)(nil))
}

func (o AppEngineVersionIamBindingArrayOutput) ToAppEngineVersionIamBindingArrayOutput() AppEngineVersionIamBindingArrayOutput {
	return o
}

func (o AppEngineVersionIamBindingArrayOutput) ToAppEngineVersionIamBindingArrayOutputWithContext(ctx context.Context) AppEngineVersionIamBindingArrayOutput {
	return o
}

func (o AppEngineVersionIamBindingArrayOutput) Index(i pulumi.IntInput) AppEngineVersionIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppEngineVersionIamBinding {
		return vs[0].([]AppEngineVersionIamBinding)[vs[1].(int)]
	}).(AppEngineVersionIamBindingOutput)
}

type AppEngineVersionIamBindingMapOutput struct{ *pulumi.OutputState }

func (AppEngineVersionIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppEngineVersionIamBinding)(nil))
}

func (o AppEngineVersionIamBindingMapOutput) ToAppEngineVersionIamBindingMapOutput() AppEngineVersionIamBindingMapOutput {
	return o
}

func (o AppEngineVersionIamBindingMapOutput) ToAppEngineVersionIamBindingMapOutputWithContext(ctx context.Context) AppEngineVersionIamBindingMapOutput {
	return o
}

func (o AppEngineVersionIamBindingMapOutput) MapIndex(k pulumi.StringInput) AppEngineVersionIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppEngineVersionIamBinding {
		return vs[0].(map[string]AppEngineVersionIamBinding)[vs[1].(string)]
	}).(AppEngineVersionIamBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(AppEngineVersionIamBindingOutput{})
	pulumi.RegisterOutputType(AppEngineVersionIamBindingPtrOutput{})
	pulumi.RegisterOutputType(AppEngineVersionIamBindingArrayOutput{})
	pulumi.RegisterOutputType(AppEngineVersionIamBindingMapOutput{})
}

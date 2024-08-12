// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for API Gateway ApiConfig. Each of these resources serves a different use case:
//
// * `apigateway.ApiConfigIamPolicy`: Authoritative. Sets the IAM policy for the apiconfig and replaces any existing policy already attached.
// * `apigateway.ApiConfigIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the apiconfig are preserved.
// * `apigateway.ApiConfigIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the apiconfig are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `apigateway.ApiConfigIamPolicy`: Retrieves the IAM policy for the apiconfig
//
// > **Note:** `apigateway.ApiConfigIamPolicy` **cannot** be used in conjunction with `apigateway.ApiConfigIamBinding` and `apigateway.ApiConfigIamMember` or they will fight over what your policy should be.
//
// > **Note:** `apigateway.ApiConfigIamBinding` resources **can be** used in conjunction with `apigateway.ApiConfigIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_api\_gateway\_api\_config\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigateway"
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
//						Role: "roles/apigateway.viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = apigateway.NewApiConfigIamPolicy(ctx, "policy", &apigateway.ApiConfigIamPolicyArgs{
//				Api:        pulumi.Any(apiCfg.Api),
//				ApiConfig:  pulumi.Any(apiCfg.ApiConfigId),
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
// ## apigateway.ApiConfigIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigateway.NewApiConfigIamBinding(ctx, "binding", &apigateway.ApiConfigIamBindingArgs{
//				Api:       pulumi.Any(apiCfg.Api),
//				ApiConfig: pulumi.Any(apiCfg.ApiConfigId),
//				Role:      pulumi.String("roles/apigateway.viewer"),
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
// ## apigateway.ApiConfigIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigateway.NewApiConfigIamMember(ctx, "member", &apigateway.ApiConfigIamMemberArgs{
//				Api:       pulumi.Any(apiCfg.Api),
//				ApiConfig: pulumi.Any(apiCfg.ApiConfigId),
//				Role:      pulumi.String("roles/apigateway.viewer"),
//				Member:    pulumi.String("user:jane@example.com"),
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
// ## This resource supports User Project Overrides.
//
// -
//
// # IAM policy for API Gateway ApiConfig
// Three different resources help you manage your IAM policy for API Gateway ApiConfig. Each of these resources serves a different use case:
//
// * `apigateway.ApiConfigIamPolicy`: Authoritative. Sets the IAM policy for the apiconfig and replaces any existing policy already attached.
// * `apigateway.ApiConfigIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the apiconfig are preserved.
// * `apigateway.ApiConfigIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the apiconfig are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `apigateway.ApiConfigIamPolicy`: Retrieves the IAM policy for the apiconfig
//
// > **Note:** `apigateway.ApiConfigIamPolicy` **cannot** be used in conjunction with `apigateway.ApiConfigIamBinding` and `apigateway.ApiConfigIamMember` or they will fight over what your policy should be.
//
// > **Note:** `apigateway.ApiConfigIamBinding` resources **can be** used in conjunction with `apigateway.ApiConfigIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_api\_gateway\_api\_config\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigateway"
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
//						Role: "roles/apigateway.viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = apigateway.NewApiConfigIamPolicy(ctx, "policy", &apigateway.ApiConfigIamPolicyArgs{
//				Api:        pulumi.Any(apiCfg.Api),
//				ApiConfig:  pulumi.Any(apiCfg.ApiConfigId),
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
// ## apigateway.ApiConfigIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigateway.NewApiConfigIamBinding(ctx, "binding", &apigateway.ApiConfigIamBindingArgs{
//				Api:       pulumi.Any(apiCfg.Api),
//				ApiConfig: pulumi.Any(apiCfg.ApiConfigId),
//				Role:      pulumi.String("roles/apigateway.viewer"),
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
// ## apigateway.ApiConfigIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigateway.NewApiConfigIamMember(ctx, "member", &apigateway.ApiConfigIamMemberArgs{
//				Api:       pulumi.Any(apiCfg.Api),
//				ApiConfig: pulumi.Any(apiCfg.ApiConfigId),
//				Role:      pulumi.String("roles/apigateway.viewer"),
//				Member:    pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}}
//
// * {{project}}/{{api}}/{{api_config}}
//
// * {{api}}/{{api_config}}
//
// * {{api_config}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// API Gateway apiconfig IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/apiConfigIamMember:ApiConfigIamMember editor "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/apiConfigIamMember:ApiConfigIamMember editor "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/apiConfigIamMember:ApiConfigIamMember editor projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ApiConfigIamMember struct {
	pulumi.CustomResourceState

	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringOutput                  `pulumi:"api"`
	ApiConfig pulumi.StringOutput                  `pulumi:"apiConfig"`
	Condition ApiConfigIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
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
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewApiConfigIamMember registers a new resource with the given unique name, arguments, and options.
func NewApiConfigIamMember(ctx *pulumi.Context,
	name string, args *ApiConfigIamMemberArgs, opts ...pulumi.ResourceOption) (*ApiConfigIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	if args.ApiConfig == nil {
		return nil, errors.New("invalid value for required argument 'ApiConfig'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiConfigIamMember
	err := ctx.RegisterResource("gcp:apigateway/apiConfigIamMember:ApiConfigIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiConfigIamMember gets an existing ApiConfigIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiConfigIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiConfigIamMemberState, opts ...pulumi.ResourceOption) (*ApiConfigIamMember, error) {
	var resource ApiConfigIamMember
	err := ctx.ReadResource("gcp:apigateway/apiConfigIamMember:ApiConfigIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiConfigIamMember resources.
type apiConfigIamMemberState struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       *string                      `pulumi:"api"`
	ApiConfig *string                      `pulumi:"apiConfig"`
	Condition *ApiConfigIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
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
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ApiConfigIamMemberState struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringPtrInput
	ApiConfig pulumi.StringPtrInput
	Condition ApiConfigIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
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
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ApiConfigIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamMemberState)(nil)).Elem()
}

type apiConfigIamMemberArgs struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       string                       `pulumi:"api"`
	ApiConfig string                       `pulumi:"apiConfig"`
	Condition *ApiConfigIamMemberCondition `pulumi:"condition"`
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
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ApiConfigIamMember resource.
type ApiConfigIamMemberArgs struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringInput
	ApiConfig pulumi.StringInput
	Condition ApiConfigIamMemberConditionPtrInput
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
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ApiConfigIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamMemberArgs)(nil)).Elem()
}

type ApiConfigIamMemberInput interface {
	pulumi.Input

	ToApiConfigIamMemberOutput() ApiConfigIamMemberOutput
	ToApiConfigIamMemberOutputWithContext(ctx context.Context) ApiConfigIamMemberOutput
}

func (*ApiConfigIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConfigIamMember)(nil)).Elem()
}

func (i *ApiConfigIamMember) ToApiConfigIamMemberOutput() ApiConfigIamMemberOutput {
	return i.ToApiConfigIamMemberOutputWithContext(context.Background())
}

func (i *ApiConfigIamMember) ToApiConfigIamMemberOutputWithContext(ctx context.Context) ApiConfigIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamMemberOutput)
}

// ApiConfigIamMemberArrayInput is an input type that accepts ApiConfigIamMemberArray and ApiConfigIamMemberArrayOutput values.
// You can construct a concrete instance of `ApiConfigIamMemberArrayInput` via:
//
//	ApiConfigIamMemberArray{ ApiConfigIamMemberArgs{...} }
type ApiConfigIamMemberArrayInput interface {
	pulumi.Input

	ToApiConfigIamMemberArrayOutput() ApiConfigIamMemberArrayOutput
	ToApiConfigIamMemberArrayOutputWithContext(context.Context) ApiConfigIamMemberArrayOutput
}

type ApiConfigIamMemberArray []ApiConfigIamMemberInput

func (ApiConfigIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiConfigIamMember)(nil)).Elem()
}

func (i ApiConfigIamMemberArray) ToApiConfigIamMemberArrayOutput() ApiConfigIamMemberArrayOutput {
	return i.ToApiConfigIamMemberArrayOutputWithContext(context.Background())
}

func (i ApiConfigIamMemberArray) ToApiConfigIamMemberArrayOutputWithContext(ctx context.Context) ApiConfigIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamMemberArrayOutput)
}

// ApiConfigIamMemberMapInput is an input type that accepts ApiConfigIamMemberMap and ApiConfigIamMemberMapOutput values.
// You can construct a concrete instance of `ApiConfigIamMemberMapInput` via:
//
//	ApiConfigIamMemberMap{ "key": ApiConfigIamMemberArgs{...} }
type ApiConfigIamMemberMapInput interface {
	pulumi.Input

	ToApiConfigIamMemberMapOutput() ApiConfigIamMemberMapOutput
	ToApiConfigIamMemberMapOutputWithContext(context.Context) ApiConfigIamMemberMapOutput
}

type ApiConfigIamMemberMap map[string]ApiConfigIamMemberInput

func (ApiConfigIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiConfigIamMember)(nil)).Elem()
}

func (i ApiConfigIamMemberMap) ToApiConfigIamMemberMapOutput() ApiConfigIamMemberMapOutput {
	return i.ToApiConfigIamMemberMapOutputWithContext(context.Background())
}

func (i ApiConfigIamMemberMap) ToApiConfigIamMemberMapOutputWithContext(ctx context.Context) ApiConfigIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamMemberMapOutput)
}

type ApiConfigIamMemberOutput struct{ *pulumi.OutputState }

func (ApiConfigIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConfigIamMember)(nil)).Elem()
}

func (o ApiConfigIamMemberOutput) ToApiConfigIamMemberOutput() ApiConfigIamMemberOutput {
	return o
}

func (o ApiConfigIamMemberOutput) ToApiConfigIamMemberOutputWithContext(ctx context.Context) ApiConfigIamMemberOutput {
	return o
}

// The API to attach the config to.
// Used to find the parent resource to bind the IAM policy to
func (o ApiConfigIamMemberOutput) Api() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamMember) pulumi.StringOutput { return v.Api }).(pulumi.StringOutput)
}

func (o ApiConfigIamMemberOutput) ApiConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamMember) pulumi.StringOutput { return v.ApiConfig }).(pulumi.StringOutput)
}

func (o ApiConfigIamMemberOutput) Condition() ApiConfigIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *ApiConfigIamMember) ApiConfigIamMemberConditionPtrOutput { return v.Condition }).(ApiConfigIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o ApiConfigIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
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
func (o ApiConfigIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o ApiConfigIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o ApiConfigIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type ApiConfigIamMemberArrayOutput struct{ *pulumi.OutputState }

func (ApiConfigIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiConfigIamMember)(nil)).Elem()
}

func (o ApiConfigIamMemberArrayOutput) ToApiConfigIamMemberArrayOutput() ApiConfigIamMemberArrayOutput {
	return o
}

func (o ApiConfigIamMemberArrayOutput) ToApiConfigIamMemberArrayOutputWithContext(ctx context.Context) ApiConfigIamMemberArrayOutput {
	return o
}

func (o ApiConfigIamMemberArrayOutput) Index(i pulumi.IntInput) ApiConfigIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiConfigIamMember {
		return vs[0].([]*ApiConfigIamMember)[vs[1].(int)]
	}).(ApiConfigIamMemberOutput)
}

type ApiConfigIamMemberMapOutput struct{ *pulumi.OutputState }

func (ApiConfigIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiConfigIamMember)(nil)).Elem()
}

func (o ApiConfigIamMemberMapOutput) ToApiConfigIamMemberMapOutput() ApiConfigIamMemberMapOutput {
	return o
}

func (o ApiConfigIamMemberMapOutput) ToApiConfigIamMemberMapOutputWithContext(ctx context.Context) ApiConfigIamMemberMapOutput {
	return o
}

func (o ApiConfigIamMemberMapOutput) MapIndex(k pulumi.StringInput) ApiConfigIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiConfigIamMember {
		return vs[0].(map[string]*ApiConfigIamMember)[vs[1].(string)]
	}).(ApiConfigIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiConfigIamMemberInput)(nil)).Elem(), &ApiConfigIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiConfigIamMemberArrayInput)(nil)).Elem(), ApiConfigIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiConfigIamMemberMapInput)(nil)).Elem(), ApiConfigIamMemberMap{})
	pulumi.RegisterOutputType(ApiConfigIamMemberOutput{})
	pulumi.RegisterOutputType(ApiConfigIamMemberArrayOutput{})
	pulumi.RegisterOutputType(ApiConfigIamMemberMapOutput{})
}

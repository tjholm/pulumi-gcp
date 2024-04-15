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

// Three different resources help you manage your IAM policy for API Gateway Api. Each of these resources serves a different use case:
//
// * `apigateway.ApiIamPolicy`: Authoritative. Sets the IAM policy for the api and replaces any existing policy already attached.
// * `apigateway.ApiIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the api are preserved.
// * `apigateway.ApiIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the api are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `apigateway.ApiIamPolicy`: Retrieves the IAM policy for the api
//
// > **Note:** `apigateway.ApiIamPolicy` **cannot** be used in conjunction with `apigateway.ApiIamBinding` and `apigateway.ApiIamMember` or they will fight over what your policy should be.
//
// > **Note:** `apigateway.ApiIamBinding` resources **can be** used in conjunction with `apigateway.ApiIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_api\_gateway\_api\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
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
//			_, err = apigateway.NewApiIamPolicy(ctx, "policy", &apigateway.ApiIamPolicyArgs{
//				Project:    pulumi.Any(api.Project),
//				Api:        pulumi.Any(api.ApiId),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_api\_gateway\_api\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := apigateway.NewApiIamBinding(ctx, "binding", &apigateway.ApiIamBindingArgs{
//				Project: pulumi.Any(api.Project),
//				Api:     pulumi.Any(api.ApiId),
//				Role:    pulumi.String("roles/apigateway.viewer"),
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
// ## google\_api\_gateway\_api\_iam\_member
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := apigateway.NewApiIamMember(ctx, "member", &apigateway.ApiIamMemberArgs{
//				Project: pulumi.Any(api.Project),
//				Api:     pulumi.Any(api.ApiId),
//				Role:    pulumi.String("roles/apigateway.viewer"),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_api\_gateway\_api\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
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
//			_, err = apigateway.NewApiIamPolicy(ctx, "policy", &apigateway.ApiIamPolicyArgs{
//				Project:    pulumi.Any(api.Project),
//				Api:        pulumi.Any(api.ApiId),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_api\_gateway\_api\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := apigateway.NewApiIamBinding(ctx, "binding", &apigateway.ApiIamBindingArgs{
//				Project: pulumi.Any(api.Project),
//				Api:     pulumi.Any(api.ApiId),
//				Role:    pulumi.String("roles/apigateway.viewer"),
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
// ## google\_api\_gateway\_api\_iam\_member
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := apigateway.NewApiIamMember(ctx, "member", &apigateway.ApiIamMemberArgs{
//				Project: pulumi.Any(api.Project),
//				Api:     pulumi.Any(api.ApiId),
//				Role:    pulumi.String("roles/apigateway.viewer"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/locations/global/apis/{{api}}
//
// * {{project}}/{{api}}
//
// * {{api}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// API Gateway api IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/apiIamBinding:ApiIamBinding editor "projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/apiIamBinding:ApiIamBinding editor "projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/apiIamBinding:ApiIamBinding editor projects/{{project}}/locations/global/apis/{{api}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ApiIamBinding struct {
	pulumi.CustomResourceState

	Api       pulumi.StringOutput             `pulumi:"api"`
	Condition ApiIamBindingConditionPtrOutput `pulumi:"condition"`
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
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewApiIamBinding registers a new resource with the given unique name, arguments, and options.
func NewApiIamBinding(ctx *pulumi.Context,
	name string, args *ApiIamBindingArgs, opts ...pulumi.ResourceOption) (*ApiIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiIamBinding
	err := ctx.RegisterResource("gcp:apigateway/apiIamBinding:ApiIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiIamBinding gets an existing ApiIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIamBindingState, opts ...pulumi.ResourceOption) (*ApiIamBinding, error) {
	var resource ApiIamBinding
	err := ctx.ReadResource("gcp:apigateway/apiIamBinding:ApiIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiIamBinding resources.
type apiIamBindingState struct {
	Api       *string                 `pulumi:"api"`
	Condition *ApiIamBindingCondition `pulumi:"condition"`
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
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ApiIamBindingState struct {
	Api       pulumi.StringPtrInput
	Condition ApiIamBindingConditionPtrInput
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
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ApiIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIamBindingState)(nil)).Elem()
}

type apiIamBindingArgs struct {
	Api       string                  `pulumi:"api"`
	Condition *ApiIamBindingCondition `pulumi:"condition"`
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
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ApiIamBinding resource.
type ApiIamBindingArgs struct {
	Api       pulumi.StringInput
	Condition ApiIamBindingConditionPtrInput
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
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ApiIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIamBindingArgs)(nil)).Elem()
}

type ApiIamBindingInput interface {
	pulumi.Input

	ToApiIamBindingOutput() ApiIamBindingOutput
	ToApiIamBindingOutputWithContext(ctx context.Context) ApiIamBindingOutput
}

func (*ApiIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIamBinding)(nil)).Elem()
}

func (i *ApiIamBinding) ToApiIamBindingOutput() ApiIamBindingOutput {
	return i.ToApiIamBindingOutputWithContext(context.Background())
}

func (i *ApiIamBinding) ToApiIamBindingOutputWithContext(ctx context.Context) ApiIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamBindingOutput)
}

// ApiIamBindingArrayInput is an input type that accepts ApiIamBindingArray and ApiIamBindingArrayOutput values.
// You can construct a concrete instance of `ApiIamBindingArrayInput` via:
//
//	ApiIamBindingArray{ ApiIamBindingArgs{...} }
type ApiIamBindingArrayInput interface {
	pulumi.Input

	ToApiIamBindingArrayOutput() ApiIamBindingArrayOutput
	ToApiIamBindingArrayOutputWithContext(context.Context) ApiIamBindingArrayOutput
}

type ApiIamBindingArray []ApiIamBindingInput

func (ApiIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiIamBinding)(nil)).Elem()
}

func (i ApiIamBindingArray) ToApiIamBindingArrayOutput() ApiIamBindingArrayOutput {
	return i.ToApiIamBindingArrayOutputWithContext(context.Background())
}

func (i ApiIamBindingArray) ToApiIamBindingArrayOutputWithContext(ctx context.Context) ApiIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamBindingArrayOutput)
}

// ApiIamBindingMapInput is an input type that accepts ApiIamBindingMap and ApiIamBindingMapOutput values.
// You can construct a concrete instance of `ApiIamBindingMapInput` via:
//
//	ApiIamBindingMap{ "key": ApiIamBindingArgs{...} }
type ApiIamBindingMapInput interface {
	pulumi.Input

	ToApiIamBindingMapOutput() ApiIamBindingMapOutput
	ToApiIamBindingMapOutputWithContext(context.Context) ApiIamBindingMapOutput
}

type ApiIamBindingMap map[string]ApiIamBindingInput

func (ApiIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiIamBinding)(nil)).Elem()
}

func (i ApiIamBindingMap) ToApiIamBindingMapOutput() ApiIamBindingMapOutput {
	return i.ToApiIamBindingMapOutputWithContext(context.Background())
}

func (i ApiIamBindingMap) ToApiIamBindingMapOutputWithContext(ctx context.Context) ApiIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamBindingMapOutput)
}

type ApiIamBindingOutput struct{ *pulumi.OutputState }

func (ApiIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIamBinding)(nil)).Elem()
}

func (o ApiIamBindingOutput) ToApiIamBindingOutput() ApiIamBindingOutput {
	return o
}

func (o ApiIamBindingOutput) ToApiIamBindingOutputWithContext(ctx context.Context) ApiIamBindingOutput {
	return o
}

func (o ApiIamBindingOutput) Api() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamBinding) pulumi.StringOutput { return v.Api }).(pulumi.StringOutput)
}

func (o ApiIamBindingOutput) Condition() ApiIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *ApiIamBinding) ApiIamBindingConditionPtrOutput { return v.Condition }).(ApiIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o ApiIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
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
func (o ApiIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o ApiIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o ApiIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type ApiIamBindingArrayOutput struct{ *pulumi.OutputState }

func (ApiIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiIamBinding)(nil)).Elem()
}

func (o ApiIamBindingArrayOutput) ToApiIamBindingArrayOutput() ApiIamBindingArrayOutput {
	return o
}

func (o ApiIamBindingArrayOutput) ToApiIamBindingArrayOutputWithContext(ctx context.Context) ApiIamBindingArrayOutput {
	return o
}

func (o ApiIamBindingArrayOutput) Index(i pulumi.IntInput) ApiIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiIamBinding {
		return vs[0].([]*ApiIamBinding)[vs[1].(int)]
	}).(ApiIamBindingOutput)
}

type ApiIamBindingMapOutput struct{ *pulumi.OutputState }

func (ApiIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiIamBinding)(nil)).Elem()
}

func (o ApiIamBindingMapOutput) ToApiIamBindingMapOutput() ApiIamBindingMapOutput {
	return o
}

func (o ApiIamBindingMapOutput) ToApiIamBindingMapOutputWithContext(ctx context.Context) ApiIamBindingMapOutput {
	return o
}

func (o ApiIamBindingMapOutput) MapIndex(k pulumi.StringInput) ApiIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiIamBinding {
		return vs[0].(map[string]*ApiIamBinding)[vs[1].(string)]
	}).(ApiIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiIamBindingInput)(nil)).Elem(), &ApiIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiIamBindingArrayInput)(nil)).Elem(), ApiIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiIamBindingMapInput)(nil)).Elem(), ApiIamBindingMap{})
	pulumi.RegisterOutputType(ApiIamBindingOutput{})
	pulumi.RegisterOutputType(ApiIamBindingArrayOutput{})
	pulumi.RegisterOutputType(ApiIamBindingMapOutput{})
}

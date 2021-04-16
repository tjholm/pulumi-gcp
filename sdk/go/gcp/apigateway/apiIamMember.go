// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for API Gateway Api. Each of these resources serves a different use case:
//
// * `apigateway.ApiIamPolicy`: Authoritative. Sets the IAM policy for the api and replaces any existing policy already attached.
// * `apigateway.ApiIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the api are preserved.
// * `apigateway.ApiIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the api are preserved.
//
// > **Note:** `apigateway.ApiIamPolicy` **cannot** be used in conjunction with `apigateway.ApiIamBinding` and `apigateway.ApiIamMember` or they will fight over what your policy should be.
//
// > **Note:** `apigateway.ApiIamBinding` resources **can be** used in conjunction with `apigateway.ApiIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_api\_gateway\_api\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/apigateway"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/apigateway.viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewApiIamPolicy(ctx, "policy", &apigateway.ApiIamPolicyArgs{
// 			Project:    pulumi.Any(google_api_gateway_api.Api.Project),
// 			Api:        pulumi.Any(google_api_gateway_api.Api.Api_id),
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
// ## google\_api\_gateway\_api\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigateway.NewApiIamBinding(ctx, "binding", &apigateway.ApiIamBindingArgs{
// 			Project: pulumi.Any(google_api_gateway_api.Api.Project),
// 			Api:     pulumi.Any(google_api_gateway_api.Api.Api_id),
// 			Role:    pulumi.String("roles/apigateway.viewer"),
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
// ## google\_api\_gateway\_api\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigateway.NewApiIamMember(ctx, "member", &apigateway.ApiIamMemberArgs{
// 			Project: pulumi.Any(google_api_gateway_api.Api.Project),
// 			Api:     pulumi.Any(google_api_gateway_api.Api.Api_id),
// 			Role:    pulumi.String("roles/apigateway.viewer"),
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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/global/apis/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway api IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/apiIamMember:ApiIamMember editor "projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/apiIamMember:ApiIamMember editor "projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/apiIamMember:ApiIamMember editor projects/{{project}}/locations/global/apis/{{api}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ApiIamMember struct {
	pulumi.CustomResourceState

	Api       pulumi.StringOutput            `pulumi:"api"`
	Condition ApiIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewApiIamMember registers a new resource with the given unique name, arguments, and options.
func NewApiIamMember(ctx *pulumi.Context,
	name string, args *ApiIamMemberArgs, opts ...pulumi.ResourceOption) (*ApiIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ApiIamMember
	err := ctx.RegisterResource("gcp:apigateway/apiIamMember:ApiIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiIamMember gets an existing ApiIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIamMemberState, opts ...pulumi.ResourceOption) (*ApiIamMember, error) {
	var resource ApiIamMember
	err := ctx.ReadResource("gcp:apigateway/apiIamMember:ApiIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiIamMember resources.
type apiIamMemberState struct {
	Api       *string                `pulumi:"api"`
	Condition *ApiIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ApiIamMemberState struct {
	Api       pulumi.StringPtrInput
	Condition ApiIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ApiIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIamMemberState)(nil)).Elem()
}

type apiIamMemberArgs struct {
	Api       string                 `pulumi:"api"`
	Condition *ApiIamMemberCondition `pulumi:"condition"`
	Member    string                 `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ApiIamMember resource.
type ApiIamMemberArgs struct {
	Api       pulumi.StringInput
	Condition ApiIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ApiIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIamMemberArgs)(nil)).Elem()
}

type ApiIamMemberInput interface {
	pulumi.Input

	ToApiIamMemberOutput() ApiIamMemberOutput
	ToApiIamMemberOutputWithContext(ctx context.Context) ApiIamMemberOutput
}

func (*ApiIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiIamMember)(nil))
}

func (i *ApiIamMember) ToApiIamMemberOutput() ApiIamMemberOutput {
	return i.ToApiIamMemberOutputWithContext(context.Background())
}

func (i *ApiIamMember) ToApiIamMemberOutputWithContext(ctx context.Context) ApiIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamMemberOutput)
}

func (i *ApiIamMember) ToApiIamMemberPtrOutput() ApiIamMemberPtrOutput {
	return i.ToApiIamMemberPtrOutputWithContext(context.Background())
}

func (i *ApiIamMember) ToApiIamMemberPtrOutputWithContext(ctx context.Context) ApiIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamMemberPtrOutput)
}

type ApiIamMemberPtrInput interface {
	pulumi.Input

	ToApiIamMemberPtrOutput() ApiIamMemberPtrOutput
	ToApiIamMemberPtrOutputWithContext(ctx context.Context) ApiIamMemberPtrOutput
}

type apiIamMemberPtrType ApiIamMemberArgs

func (*apiIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIamMember)(nil))
}

func (i *apiIamMemberPtrType) ToApiIamMemberPtrOutput() ApiIamMemberPtrOutput {
	return i.ToApiIamMemberPtrOutputWithContext(context.Background())
}

func (i *apiIamMemberPtrType) ToApiIamMemberPtrOutputWithContext(ctx context.Context) ApiIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamMemberPtrOutput)
}

// ApiIamMemberArrayInput is an input type that accepts ApiIamMemberArray and ApiIamMemberArrayOutput values.
// You can construct a concrete instance of `ApiIamMemberArrayInput` via:
//
//          ApiIamMemberArray{ ApiIamMemberArgs{...} }
type ApiIamMemberArrayInput interface {
	pulumi.Input

	ToApiIamMemberArrayOutput() ApiIamMemberArrayOutput
	ToApiIamMemberArrayOutputWithContext(context.Context) ApiIamMemberArrayOutput
}

type ApiIamMemberArray []ApiIamMemberInput

func (ApiIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ApiIamMember)(nil))
}

func (i ApiIamMemberArray) ToApiIamMemberArrayOutput() ApiIamMemberArrayOutput {
	return i.ToApiIamMemberArrayOutputWithContext(context.Background())
}

func (i ApiIamMemberArray) ToApiIamMemberArrayOutputWithContext(ctx context.Context) ApiIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamMemberArrayOutput)
}

// ApiIamMemberMapInput is an input type that accepts ApiIamMemberMap and ApiIamMemberMapOutput values.
// You can construct a concrete instance of `ApiIamMemberMapInput` via:
//
//          ApiIamMemberMap{ "key": ApiIamMemberArgs{...} }
type ApiIamMemberMapInput interface {
	pulumi.Input

	ToApiIamMemberMapOutput() ApiIamMemberMapOutput
	ToApiIamMemberMapOutputWithContext(context.Context) ApiIamMemberMapOutput
}

type ApiIamMemberMap map[string]ApiIamMemberInput

func (ApiIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ApiIamMember)(nil))
}

func (i ApiIamMemberMap) ToApiIamMemberMapOutput() ApiIamMemberMapOutput {
	return i.ToApiIamMemberMapOutputWithContext(context.Background())
}

func (i ApiIamMemberMap) ToApiIamMemberMapOutputWithContext(ctx context.Context) ApiIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamMemberMapOutput)
}

type ApiIamMemberOutput struct {
	*pulumi.OutputState
}

func (ApiIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiIamMember)(nil))
}

func (o ApiIamMemberOutput) ToApiIamMemberOutput() ApiIamMemberOutput {
	return o
}

func (o ApiIamMemberOutput) ToApiIamMemberOutputWithContext(ctx context.Context) ApiIamMemberOutput {
	return o
}

func (o ApiIamMemberOutput) ToApiIamMemberPtrOutput() ApiIamMemberPtrOutput {
	return o.ToApiIamMemberPtrOutputWithContext(context.Background())
}

func (o ApiIamMemberOutput) ToApiIamMemberPtrOutputWithContext(ctx context.Context) ApiIamMemberPtrOutput {
	return o.ApplyT(func(v ApiIamMember) *ApiIamMember {
		return &v
	}).(ApiIamMemberPtrOutput)
}

type ApiIamMemberPtrOutput struct {
	*pulumi.OutputState
}

func (ApiIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIamMember)(nil))
}

func (o ApiIamMemberPtrOutput) ToApiIamMemberPtrOutput() ApiIamMemberPtrOutput {
	return o
}

func (o ApiIamMemberPtrOutput) ToApiIamMemberPtrOutputWithContext(ctx context.Context) ApiIamMemberPtrOutput {
	return o
}

type ApiIamMemberArrayOutput struct{ *pulumi.OutputState }

func (ApiIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiIamMember)(nil))
}

func (o ApiIamMemberArrayOutput) ToApiIamMemberArrayOutput() ApiIamMemberArrayOutput {
	return o
}

func (o ApiIamMemberArrayOutput) ToApiIamMemberArrayOutputWithContext(ctx context.Context) ApiIamMemberArrayOutput {
	return o
}

func (o ApiIamMemberArrayOutput) Index(i pulumi.IntInput) ApiIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiIamMember {
		return vs[0].([]ApiIamMember)[vs[1].(int)]
	}).(ApiIamMemberOutput)
}

type ApiIamMemberMapOutput struct{ *pulumi.OutputState }

func (ApiIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiIamMember)(nil))
}

func (o ApiIamMemberMapOutput) ToApiIamMemberMapOutput() ApiIamMemberMapOutput {
	return o
}

func (o ApiIamMemberMapOutput) ToApiIamMemberMapOutputWithContext(ctx context.Context) ApiIamMemberMapOutput {
	return o
}

func (o ApiIamMemberMapOutput) MapIndex(k pulumi.StringInput) ApiIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiIamMember {
		return vs[0].(map[string]ApiIamMember)[vs[1].(string)]
	}).(ApiIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiIamMemberOutput{})
	pulumi.RegisterOutputType(ApiIamMemberPtrOutput{})
	pulumi.RegisterOutputType(ApiIamMemberArrayOutput{})
	pulumi.RegisterOutputType(ApiIamMemberMapOutput{})
}

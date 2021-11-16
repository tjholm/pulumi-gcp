// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for API Gateway ApiConfig. Each of these resources serves a different use case:
//
// * `apigateway.ApiConfigIamPolicy`: Authoritative. Sets the IAM policy for the apiconfig and replaces any existing policy already attached.
// * `apigateway.ApiConfigIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the apiconfig are preserved.
// * `apigateway.ApiConfigIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the apiconfig are preserved.
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigateway"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
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
// 		_, err = apigateway.NewApiConfigIamPolicy(ctx, "policy", &apigateway.ApiConfigIamPolicyArgs{
// 			Api:        pulumi.Any(google_api_gateway_api_config.Api_cfg.Api),
// 			ApiConfig:  pulumi.Any(google_api_gateway_api_config.Api_cfg.Api_config_id),
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
// ## google\_api\_gateway\_api\_config\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigateway.NewApiConfigIamBinding(ctx, "binding", &apigateway.ApiConfigIamBindingArgs{
// 			Api:       pulumi.Any(google_api_gateway_api_config.Api_cfg.Api),
// 			ApiConfig: pulumi.Any(google_api_gateway_api_config.Api_cfg.Api_config_id),
// 			Role:      pulumi.String("roles/apigateway.viewer"),
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
// ## google\_api\_gateway\_api\_config\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigateway.NewApiConfigIamMember(ctx, "member", &apigateway.ApiConfigIamMemberArgs{
// 			Api:       pulumi.Any(google_api_gateway_api_config.Api_cfg.Api),
// 			ApiConfig: pulumi.Any(google_api_gateway_api_config.Api_cfg.Api_config_id),
// 			Role:      pulumi.String("roles/apigateway.viewer"),
// 			Member:    pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} * {{project}}/{{api}}/{{api_config}} * {{api}}/{{api_config}} * {{api_config}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway apiconfig IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding editor "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding editor "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding editor projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ApiConfigIamBinding struct {
	pulumi.CustomResourceState

	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringOutput                   `pulumi:"api"`
	ApiConfig pulumi.StringOutput                   `pulumi:"apiConfig"`
	Condition ApiConfigIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewApiConfigIamBinding registers a new resource with the given unique name, arguments, and options.
func NewApiConfigIamBinding(ctx *pulumi.Context,
	name string, args *ApiConfigIamBindingArgs, opts ...pulumi.ResourceOption) (*ApiConfigIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	if args.ApiConfig == nil {
		return nil, errors.New("invalid value for required argument 'ApiConfig'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ApiConfigIamBinding
	err := ctx.RegisterResource("gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiConfigIamBinding gets an existing ApiConfigIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiConfigIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiConfigIamBindingState, opts ...pulumi.ResourceOption) (*ApiConfigIamBinding, error) {
	var resource ApiConfigIamBinding
	err := ctx.ReadResource("gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiConfigIamBinding resources.
type apiConfigIamBindingState struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       *string                       `pulumi:"api"`
	ApiConfig *string                       `pulumi:"apiConfig"`
	Condition *ApiConfigIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ApiConfigIamBindingState struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringPtrInput
	ApiConfig pulumi.StringPtrInput
	Condition ApiConfigIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ApiConfigIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamBindingState)(nil)).Elem()
}

type apiConfigIamBindingArgs struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       string                        `pulumi:"api"`
	ApiConfig string                        `pulumi:"apiConfig"`
	Condition *ApiConfigIamBindingCondition `pulumi:"condition"`
	Members   []string                      `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ApiConfigIamBinding resource.
type ApiConfigIamBindingArgs struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringInput
	ApiConfig pulumi.StringInput
	Condition ApiConfigIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ApiConfigIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamBindingArgs)(nil)).Elem()
}

type ApiConfigIamBindingInput interface {
	pulumi.Input

	ToApiConfigIamBindingOutput() ApiConfigIamBindingOutput
	ToApiConfigIamBindingOutputWithContext(ctx context.Context) ApiConfigIamBindingOutput
}

func (*ApiConfigIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConfigIamBinding)(nil))
}

func (i *ApiConfigIamBinding) ToApiConfigIamBindingOutput() ApiConfigIamBindingOutput {
	return i.ToApiConfigIamBindingOutputWithContext(context.Background())
}

func (i *ApiConfigIamBinding) ToApiConfigIamBindingOutputWithContext(ctx context.Context) ApiConfigIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamBindingOutput)
}

func (i *ApiConfigIamBinding) ToApiConfigIamBindingPtrOutput() ApiConfigIamBindingPtrOutput {
	return i.ToApiConfigIamBindingPtrOutputWithContext(context.Background())
}

func (i *ApiConfigIamBinding) ToApiConfigIamBindingPtrOutputWithContext(ctx context.Context) ApiConfigIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamBindingPtrOutput)
}

type ApiConfigIamBindingPtrInput interface {
	pulumi.Input

	ToApiConfigIamBindingPtrOutput() ApiConfigIamBindingPtrOutput
	ToApiConfigIamBindingPtrOutputWithContext(ctx context.Context) ApiConfigIamBindingPtrOutput
}

type apiConfigIamBindingPtrType ApiConfigIamBindingArgs

func (*apiConfigIamBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConfigIamBinding)(nil))
}

func (i *apiConfigIamBindingPtrType) ToApiConfigIamBindingPtrOutput() ApiConfigIamBindingPtrOutput {
	return i.ToApiConfigIamBindingPtrOutputWithContext(context.Background())
}

func (i *apiConfigIamBindingPtrType) ToApiConfigIamBindingPtrOutputWithContext(ctx context.Context) ApiConfigIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamBindingPtrOutput)
}

// ApiConfigIamBindingArrayInput is an input type that accepts ApiConfigIamBindingArray and ApiConfigIamBindingArrayOutput values.
// You can construct a concrete instance of `ApiConfigIamBindingArrayInput` via:
//
//          ApiConfigIamBindingArray{ ApiConfigIamBindingArgs{...} }
type ApiConfigIamBindingArrayInput interface {
	pulumi.Input

	ToApiConfigIamBindingArrayOutput() ApiConfigIamBindingArrayOutput
	ToApiConfigIamBindingArrayOutputWithContext(context.Context) ApiConfigIamBindingArrayOutput
}

type ApiConfigIamBindingArray []ApiConfigIamBindingInput

func (ApiConfigIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiConfigIamBinding)(nil)).Elem()
}

func (i ApiConfigIamBindingArray) ToApiConfigIamBindingArrayOutput() ApiConfigIamBindingArrayOutput {
	return i.ToApiConfigIamBindingArrayOutputWithContext(context.Background())
}

func (i ApiConfigIamBindingArray) ToApiConfigIamBindingArrayOutputWithContext(ctx context.Context) ApiConfigIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamBindingArrayOutput)
}

// ApiConfigIamBindingMapInput is an input type that accepts ApiConfigIamBindingMap and ApiConfigIamBindingMapOutput values.
// You can construct a concrete instance of `ApiConfigIamBindingMapInput` via:
//
//          ApiConfigIamBindingMap{ "key": ApiConfigIamBindingArgs{...} }
type ApiConfigIamBindingMapInput interface {
	pulumi.Input

	ToApiConfigIamBindingMapOutput() ApiConfigIamBindingMapOutput
	ToApiConfigIamBindingMapOutputWithContext(context.Context) ApiConfigIamBindingMapOutput
}

type ApiConfigIamBindingMap map[string]ApiConfigIamBindingInput

func (ApiConfigIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiConfigIamBinding)(nil)).Elem()
}

func (i ApiConfigIamBindingMap) ToApiConfigIamBindingMapOutput() ApiConfigIamBindingMapOutput {
	return i.ToApiConfigIamBindingMapOutputWithContext(context.Background())
}

func (i ApiConfigIamBindingMap) ToApiConfigIamBindingMapOutputWithContext(ctx context.Context) ApiConfigIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamBindingMapOutput)
}

type ApiConfigIamBindingOutput struct{ *pulumi.OutputState }

func (ApiConfigIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConfigIamBinding)(nil))
}

func (o ApiConfigIamBindingOutput) ToApiConfigIamBindingOutput() ApiConfigIamBindingOutput {
	return o
}

func (o ApiConfigIamBindingOutput) ToApiConfigIamBindingOutputWithContext(ctx context.Context) ApiConfigIamBindingOutput {
	return o
}

func (o ApiConfigIamBindingOutput) ToApiConfigIamBindingPtrOutput() ApiConfigIamBindingPtrOutput {
	return o.ToApiConfigIamBindingPtrOutputWithContext(context.Background())
}

func (o ApiConfigIamBindingOutput) ToApiConfigIamBindingPtrOutputWithContext(ctx context.Context) ApiConfigIamBindingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiConfigIamBinding) *ApiConfigIamBinding {
		return &v
	}).(ApiConfigIamBindingPtrOutput)
}

type ApiConfigIamBindingPtrOutput struct{ *pulumi.OutputState }

func (ApiConfigIamBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConfigIamBinding)(nil))
}

func (o ApiConfigIamBindingPtrOutput) ToApiConfigIamBindingPtrOutput() ApiConfigIamBindingPtrOutput {
	return o
}

func (o ApiConfigIamBindingPtrOutput) ToApiConfigIamBindingPtrOutputWithContext(ctx context.Context) ApiConfigIamBindingPtrOutput {
	return o
}

func (o ApiConfigIamBindingPtrOutput) Elem() ApiConfigIamBindingOutput {
	return o.ApplyT(func(v *ApiConfigIamBinding) ApiConfigIamBinding {
		if v != nil {
			return *v
		}
		var ret ApiConfigIamBinding
		return ret
	}).(ApiConfigIamBindingOutput)
}

type ApiConfigIamBindingArrayOutput struct{ *pulumi.OutputState }

func (ApiConfigIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiConfigIamBinding)(nil))
}

func (o ApiConfigIamBindingArrayOutput) ToApiConfigIamBindingArrayOutput() ApiConfigIamBindingArrayOutput {
	return o
}

func (o ApiConfigIamBindingArrayOutput) ToApiConfigIamBindingArrayOutputWithContext(ctx context.Context) ApiConfigIamBindingArrayOutput {
	return o
}

func (o ApiConfigIamBindingArrayOutput) Index(i pulumi.IntInput) ApiConfigIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiConfigIamBinding {
		return vs[0].([]ApiConfigIamBinding)[vs[1].(int)]
	}).(ApiConfigIamBindingOutput)
}

type ApiConfigIamBindingMapOutput struct{ *pulumi.OutputState }

func (ApiConfigIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiConfigIamBinding)(nil))
}

func (o ApiConfigIamBindingMapOutput) ToApiConfigIamBindingMapOutput() ApiConfigIamBindingMapOutput {
	return o
}

func (o ApiConfigIamBindingMapOutput) ToApiConfigIamBindingMapOutputWithContext(ctx context.Context) ApiConfigIamBindingMapOutput {
	return o
}

func (o ApiConfigIamBindingMapOutput) MapIndex(k pulumi.StringInput) ApiConfigIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiConfigIamBinding {
		return vs[0].(map[string]ApiConfigIamBinding)[vs[1].(string)]
	}).(ApiConfigIamBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiConfigIamBindingOutput{})
	pulumi.RegisterOutputType(ApiConfigIamBindingPtrOutput{})
	pulumi.RegisterOutputType(ApiConfigIamBindingArrayOutput{})
	pulumi.RegisterOutputType(ApiConfigIamBindingMapOutput{})
}

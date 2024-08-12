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
// $ pulumi import gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy editor "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy editor "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy editor projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ApiConfigIamPolicy struct {
	pulumi.CustomResourceState

	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringOutput `pulumi:"api"`
	ApiConfig pulumi.StringOutput `pulumi:"apiConfig"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewApiConfigIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiConfigIamPolicy(ctx *pulumi.Context,
	name string, args *ApiConfigIamPolicyArgs, opts ...pulumi.ResourceOption) (*ApiConfigIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	if args.ApiConfig == nil {
		return nil, errors.New("invalid value for required argument 'ApiConfig'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiConfigIamPolicy
	err := ctx.RegisterResource("gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiConfigIamPolicy gets an existing ApiConfigIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiConfigIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiConfigIamPolicyState, opts ...pulumi.ResourceOption) (*ApiConfigIamPolicy, error) {
	var resource ApiConfigIamPolicy
	err := ctx.ReadResource("gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiConfigIamPolicy resources.
type apiConfigIamPolicyState struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       *string `pulumi:"api"`
	ApiConfig *string `pulumi:"apiConfig"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type ApiConfigIamPolicyState struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringPtrInput
	ApiConfig pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApiConfigIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamPolicyState)(nil)).Elem()
}

type apiConfigIamPolicyArgs struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       string `pulumi:"api"`
	ApiConfig string `pulumi:"apiConfig"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ApiConfigIamPolicy resource.
type ApiConfigIamPolicyArgs struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringInput
	ApiConfig pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApiConfigIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamPolicyArgs)(nil)).Elem()
}

type ApiConfigIamPolicyInput interface {
	pulumi.Input

	ToApiConfigIamPolicyOutput() ApiConfigIamPolicyOutput
	ToApiConfigIamPolicyOutputWithContext(ctx context.Context) ApiConfigIamPolicyOutput
}

func (*ApiConfigIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConfigIamPolicy)(nil)).Elem()
}

func (i *ApiConfigIamPolicy) ToApiConfigIamPolicyOutput() ApiConfigIamPolicyOutput {
	return i.ToApiConfigIamPolicyOutputWithContext(context.Background())
}

func (i *ApiConfigIamPolicy) ToApiConfigIamPolicyOutputWithContext(ctx context.Context) ApiConfigIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamPolicyOutput)
}

// ApiConfigIamPolicyArrayInput is an input type that accepts ApiConfigIamPolicyArray and ApiConfigIamPolicyArrayOutput values.
// You can construct a concrete instance of `ApiConfigIamPolicyArrayInput` via:
//
//	ApiConfigIamPolicyArray{ ApiConfigIamPolicyArgs{...} }
type ApiConfigIamPolicyArrayInput interface {
	pulumi.Input

	ToApiConfigIamPolicyArrayOutput() ApiConfigIamPolicyArrayOutput
	ToApiConfigIamPolicyArrayOutputWithContext(context.Context) ApiConfigIamPolicyArrayOutput
}

type ApiConfigIamPolicyArray []ApiConfigIamPolicyInput

func (ApiConfigIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiConfigIamPolicy)(nil)).Elem()
}

func (i ApiConfigIamPolicyArray) ToApiConfigIamPolicyArrayOutput() ApiConfigIamPolicyArrayOutput {
	return i.ToApiConfigIamPolicyArrayOutputWithContext(context.Background())
}

func (i ApiConfigIamPolicyArray) ToApiConfigIamPolicyArrayOutputWithContext(ctx context.Context) ApiConfigIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamPolicyArrayOutput)
}

// ApiConfigIamPolicyMapInput is an input type that accepts ApiConfigIamPolicyMap and ApiConfigIamPolicyMapOutput values.
// You can construct a concrete instance of `ApiConfigIamPolicyMapInput` via:
//
//	ApiConfigIamPolicyMap{ "key": ApiConfigIamPolicyArgs{...} }
type ApiConfigIamPolicyMapInput interface {
	pulumi.Input

	ToApiConfigIamPolicyMapOutput() ApiConfigIamPolicyMapOutput
	ToApiConfigIamPolicyMapOutputWithContext(context.Context) ApiConfigIamPolicyMapOutput
}

type ApiConfigIamPolicyMap map[string]ApiConfigIamPolicyInput

func (ApiConfigIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiConfigIamPolicy)(nil)).Elem()
}

func (i ApiConfigIamPolicyMap) ToApiConfigIamPolicyMapOutput() ApiConfigIamPolicyMapOutput {
	return i.ToApiConfigIamPolicyMapOutputWithContext(context.Background())
}

func (i ApiConfigIamPolicyMap) ToApiConfigIamPolicyMapOutputWithContext(ctx context.Context) ApiConfigIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamPolicyMapOutput)
}

type ApiConfigIamPolicyOutput struct{ *pulumi.OutputState }

func (ApiConfigIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConfigIamPolicy)(nil)).Elem()
}

func (o ApiConfigIamPolicyOutput) ToApiConfigIamPolicyOutput() ApiConfigIamPolicyOutput {
	return o
}

func (o ApiConfigIamPolicyOutput) ToApiConfigIamPolicyOutputWithContext(ctx context.Context) ApiConfigIamPolicyOutput {
	return o
}

// The API to attach the config to.
// Used to find the parent resource to bind the IAM policy to
func (o ApiConfigIamPolicyOutput) Api() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.StringOutput { return v.Api }).(pulumi.StringOutput)
}

func (o ApiConfigIamPolicyOutput) ApiConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.StringOutput { return v.ApiConfig }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o ApiConfigIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o ApiConfigIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o ApiConfigIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ApiConfigIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (ApiConfigIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiConfigIamPolicy)(nil)).Elem()
}

func (o ApiConfigIamPolicyArrayOutput) ToApiConfigIamPolicyArrayOutput() ApiConfigIamPolicyArrayOutput {
	return o
}

func (o ApiConfigIamPolicyArrayOutput) ToApiConfigIamPolicyArrayOutputWithContext(ctx context.Context) ApiConfigIamPolicyArrayOutput {
	return o
}

func (o ApiConfigIamPolicyArrayOutput) Index(i pulumi.IntInput) ApiConfigIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiConfigIamPolicy {
		return vs[0].([]*ApiConfigIamPolicy)[vs[1].(int)]
	}).(ApiConfigIamPolicyOutput)
}

type ApiConfigIamPolicyMapOutput struct{ *pulumi.OutputState }

func (ApiConfigIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiConfigIamPolicy)(nil)).Elem()
}

func (o ApiConfigIamPolicyMapOutput) ToApiConfigIamPolicyMapOutput() ApiConfigIamPolicyMapOutput {
	return o
}

func (o ApiConfigIamPolicyMapOutput) ToApiConfigIamPolicyMapOutputWithContext(ctx context.Context) ApiConfigIamPolicyMapOutput {
	return o
}

func (o ApiConfigIamPolicyMapOutput) MapIndex(k pulumi.StringInput) ApiConfigIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiConfigIamPolicy {
		return vs[0].(map[string]*ApiConfigIamPolicy)[vs[1].(string)]
	}).(ApiConfigIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiConfigIamPolicyInput)(nil)).Elem(), &ApiConfigIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiConfigIamPolicyArrayInput)(nil)).Elem(), ApiConfigIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiConfigIamPolicyMapInput)(nil)).Elem(), ApiConfigIamPolicyMap{})
	pulumi.RegisterOutputType(ApiConfigIamPolicyOutput{})
	pulumi.RegisterOutputType(ApiConfigIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(ApiConfigIamPolicyMapOutput{})
}

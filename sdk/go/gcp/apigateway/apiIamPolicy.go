// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigateway"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
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
//				Project:    pulumi.Any(google_api_gateway_api.Api.Project),
//				Api:        pulumi.Any(google_api_gateway_api.Api.Api_id),
//				PolicyData: *pulumi.String(admin.PolicyData),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## google\_api\_gateway\_api\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigateway.NewApiIamBinding(ctx, "binding", &apigateway.ApiIamBindingArgs{
//				Project: pulumi.Any(google_api_gateway_api.Api.Project),
//				Api:     pulumi.Any(google_api_gateway_api.Api.Api_id),
//				Role:    pulumi.String("roles/apigateway.viewer"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## google\_api\_gateway\_api\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigateway.NewApiIamMember(ctx, "member", &apigateway.ApiIamMemberArgs{
//				Project: pulumi.Any(google_api_gateway_api.Api.Project),
//				Api:     pulumi.Any(google_api_gateway_api.Api.Api_id),
//				Role:    pulumi.String("roles/apigateway.viewer"),
//				Member:  pulumi.String("user:jane@example.com"),
//			}, pulumi.Provider(google_beta))
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/global/apis/{{api}} * {{project}}/{{api}} * {{api}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway api IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:apigateway/apiIamPolicy:ApiIamPolicy editor "projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:apigateway/apiIamPolicy:ApiIamPolicy editor "projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:apigateway/apiIamPolicy:ApiIamPolicy editor projects/{{project}}/locations/global/apis/{{api}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ApiIamPolicy struct {
	pulumi.CustomResourceState

	Api pulumi.StringOutput `pulumi:"api"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewApiIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiIamPolicy(ctx *pulumi.Context,
	name string, args *ApiIamPolicyArgs, opts ...pulumi.ResourceOption) (*ApiIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiIamPolicy
	err := ctx.RegisterResource("gcp:apigateway/apiIamPolicy:ApiIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiIamPolicy gets an existing ApiIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIamPolicyState, opts ...pulumi.ResourceOption) (*ApiIamPolicy, error) {
	var resource ApiIamPolicy
	err := ctx.ReadResource("gcp:apigateway/apiIamPolicy:ApiIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiIamPolicy resources.
type apiIamPolicyState struct {
	Api *string `pulumi:"api"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
}

type ApiIamPolicyState struct {
	Api pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
}

func (ApiIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIamPolicyState)(nil)).Elem()
}

type apiIamPolicyArgs struct {
	Api string `pulumi:"api"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ApiIamPolicy resource.
type ApiIamPolicyArgs struct {
	Api pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
}

func (ApiIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIamPolicyArgs)(nil)).Elem()
}

type ApiIamPolicyInput interface {
	pulumi.Input

	ToApiIamPolicyOutput() ApiIamPolicyOutput
	ToApiIamPolicyOutputWithContext(ctx context.Context) ApiIamPolicyOutput
}

func (*ApiIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIamPolicy)(nil)).Elem()
}

func (i *ApiIamPolicy) ToApiIamPolicyOutput() ApiIamPolicyOutput {
	return i.ToApiIamPolicyOutputWithContext(context.Background())
}

func (i *ApiIamPolicy) ToApiIamPolicyOutputWithContext(ctx context.Context) ApiIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamPolicyOutput)
}

func (i *ApiIamPolicy) ToOutput(ctx context.Context) pulumix.Output[*ApiIamPolicy] {
	return pulumix.Output[*ApiIamPolicy]{
		OutputState: i.ToApiIamPolicyOutputWithContext(ctx).OutputState,
	}
}

// ApiIamPolicyArrayInput is an input type that accepts ApiIamPolicyArray and ApiIamPolicyArrayOutput values.
// You can construct a concrete instance of `ApiIamPolicyArrayInput` via:
//
//	ApiIamPolicyArray{ ApiIamPolicyArgs{...} }
type ApiIamPolicyArrayInput interface {
	pulumi.Input

	ToApiIamPolicyArrayOutput() ApiIamPolicyArrayOutput
	ToApiIamPolicyArrayOutputWithContext(context.Context) ApiIamPolicyArrayOutput
}

type ApiIamPolicyArray []ApiIamPolicyInput

func (ApiIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiIamPolicy)(nil)).Elem()
}

func (i ApiIamPolicyArray) ToApiIamPolicyArrayOutput() ApiIamPolicyArrayOutput {
	return i.ToApiIamPolicyArrayOutputWithContext(context.Background())
}

func (i ApiIamPolicyArray) ToApiIamPolicyArrayOutputWithContext(ctx context.Context) ApiIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamPolicyArrayOutput)
}

func (i ApiIamPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*ApiIamPolicy] {
	return pulumix.Output[[]*ApiIamPolicy]{
		OutputState: i.ToApiIamPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// ApiIamPolicyMapInput is an input type that accepts ApiIamPolicyMap and ApiIamPolicyMapOutput values.
// You can construct a concrete instance of `ApiIamPolicyMapInput` via:
//
//	ApiIamPolicyMap{ "key": ApiIamPolicyArgs{...} }
type ApiIamPolicyMapInput interface {
	pulumi.Input

	ToApiIamPolicyMapOutput() ApiIamPolicyMapOutput
	ToApiIamPolicyMapOutputWithContext(context.Context) ApiIamPolicyMapOutput
}

type ApiIamPolicyMap map[string]ApiIamPolicyInput

func (ApiIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiIamPolicy)(nil)).Elem()
}

func (i ApiIamPolicyMap) ToApiIamPolicyMapOutput() ApiIamPolicyMapOutput {
	return i.ToApiIamPolicyMapOutputWithContext(context.Background())
}

func (i ApiIamPolicyMap) ToApiIamPolicyMapOutputWithContext(ctx context.Context) ApiIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamPolicyMapOutput)
}

func (i ApiIamPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApiIamPolicy] {
	return pulumix.Output[map[string]*ApiIamPolicy]{
		OutputState: i.ToApiIamPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type ApiIamPolicyOutput struct{ *pulumi.OutputState }

func (ApiIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIamPolicy)(nil)).Elem()
}

func (o ApiIamPolicyOutput) ToApiIamPolicyOutput() ApiIamPolicyOutput {
	return o
}

func (o ApiIamPolicyOutput) ToApiIamPolicyOutputWithContext(ctx context.Context) ApiIamPolicyOutput {
	return o
}

func (o ApiIamPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*ApiIamPolicy] {
	return pulumix.Output[*ApiIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o ApiIamPolicyOutput) Api() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamPolicy) pulumi.StringOutput { return v.Api }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o ApiIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o ApiIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
//   - **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
//   - **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
//   - **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
func (o ApiIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ApiIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (ApiIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiIamPolicy)(nil)).Elem()
}

func (o ApiIamPolicyArrayOutput) ToApiIamPolicyArrayOutput() ApiIamPolicyArrayOutput {
	return o
}

func (o ApiIamPolicyArrayOutput) ToApiIamPolicyArrayOutputWithContext(ctx context.Context) ApiIamPolicyArrayOutput {
	return o
}

func (o ApiIamPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ApiIamPolicy] {
	return pulumix.Output[[]*ApiIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o ApiIamPolicyArrayOutput) Index(i pulumi.IntInput) ApiIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiIamPolicy {
		return vs[0].([]*ApiIamPolicy)[vs[1].(int)]
	}).(ApiIamPolicyOutput)
}

type ApiIamPolicyMapOutput struct{ *pulumi.OutputState }

func (ApiIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiIamPolicy)(nil)).Elem()
}

func (o ApiIamPolicyMapOutput) ToApiIamPolicyMapOutput() ApiIamPolicyMapOutput {
	return o
}

func (o ApiIamPolicyMapOutput) ToApiIamPolicyMapOutputWithContext(ctx context.Context) ApiIamPolicyMapOutput {
	return o
}

func (o ApiIamPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApiIamPolicy] {
	return pulumix.Output[map[string]*ApiIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o ApiIamPolicyMapOutput) MapIndex(k pulumi.StringInput) ApiIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiIamPolicy {
		return vs[0].(map[string]*ApiIamPolicy)[vs[1].(string)]
	}).(ApiIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiIamPolicyInput)(nil)).Elem(), &ApiIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiIamPolicyArrayInput)(nil)).Elem(), ApiIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiIamPolicyMapInput)(nil)).Elem(), ApiIamPolicyMap{})
	pulumi.RegisterOutputType(ApiIamPolicyOutput{})
	pulumi.RegisterOutputType(ApiIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(ApiIamPolicyMapOutput{})
}

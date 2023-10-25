// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Three different resources help you manage your IAM policy for Cloud Functions CloudFunction. Each of these resources serves a different use case:
//
// * `cloudfunctions.FunctionIamPolicy`: Authoritative. Sets the IAM policy for the cloudfunction and replaces any existing policy already attached.
// * `cloudfunctions.FunctionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cloudfunction are preserved.
// * `cloudfunctions.FunctionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cloudfunction are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `cloudfunctions.FunctionIamPolicy`: Retrieves the IAM policy for the cloudfunction
//
// > **Note:** `cloudfunctions.FunctionIamPolicy` **cannot** be used in conjunction with `cloudfunctions.FunctionIamBinding` and `cloudfunctions.FunctionIamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudfunctions.FunctionIamBinding` resources **can be** used in conjunction with `cloudfunctions.FunctionIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_cloudfunctions\_function\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
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
//						Role: "roles/viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudfunctions.NewFunctionIamPolicy(ctx, "policy", &cloudfunctions.FunctionIamPolicyArgs{
//				Project:       pulumi.Any(google_cloudfunctions_function.Function.Project),
//				Region:        pulumi.Any(google_cloudfunctions_function.Function.Region),
//				CloudFunction: pulumi.Any(google_cloudfunctions_function.Function.Name),
//				PolicyData:    *pulumi.String(admin.PolicyData),
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
// ## google\_cloudfunctions\_function\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfunctions.NewFunctionIamBinding(ctx, "binding", &cloudfunctions.FunctionIamBindingArgs{
//				Project:       pulumi.Any(google_cloudfunctions_function.Function.Project),
//				Region:        pulumi.Any(google_cloudfunctions_function.Function.Region),
//				CloudFunction: pulumi.Any(google_cloudfunctions_function.Function.Name),
//				Role:          pulumi.String("roles/viewer"),
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
// ## google\_cloudfunctions\_function\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfunctions.NewFunctionIamMember(ctx, "member", &cloudfunctions.FunctionIamMemberArgs{
//				Project:       pulumi.Any(google_cloudfunctions_function.Function.Project),
//				Region:        pulumi.Any(google_cloudfunctions_function.Function.Region),
//				CloudFunction: pulumi.Any(google_cloudfunctions_function.Function.Name),
//				Role:          pulumi.String("roles/viewer"),
//				Member:        pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} * {{project}}/{{region}}/{{cloud_function}} * {{region}}/{{cloud_function}} * {{cloud_function}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Functions cloudfunction IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy editor "projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy editor "projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy editor projects/{{project}}/locations/{{region}}/functions/{{cloud_function}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type FunctionIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringOutput `pulumi:"cloudFunction"`
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
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewFunctionIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewFunctionIamPolicy(ctx *pulumi.Context,
	name string, args *FunctionIamPolicyArgs, opts ...pulumi.ResourceOption) (*FunctionIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudFunction == nil {
		return nil, errors.New("invalid value for required argument 'CloudFunction'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FunctionIamPolicy
	err := ctx.RegisterResource("gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionIamPolicy gets an existing FunctionIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionIamPolicyState, opts ...pulumi.ResourceOption) (*FunctionIamPolicy, error) {
	var resource FunctionIamPolicy
	err := ctx.ReadResource("gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionIamPolicy resources.
type functionIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction *string `pulumi:"cloudFunction"`
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
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
}

type FunctionIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringPtrInput
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
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
}

func (FunctionIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamPolicyState)(nil)).Elem()
}

type functionIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction string `pulumi:"cloudFunction"`
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
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a FunctionIamPolicy resource.
type FunctionIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringInput
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
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
}

func (FunctionIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamPolicyArgs)(nil)).Elem()
}

type FunctionIamPolicyInput interface {
	pulumi.Input

	ToFunctionIamPolicyOutput() FunctionIamPolicyOutput
	ToFunctionIamPolicyOutputWithContext(ctx context.Context) FunctionIamPolicyOutput
}

func (*FunctionIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionIamPolicy)(nil)).Elem()
}

func (i *FunctionIamPolicy) ToFunctionIamPolicyOutput() FunctionIamPolicyOutput {
	return i.ToFunctionIamPolicyOutputWithContext(context.Background())
}

func (i *FunctionIamPolicy) ToFunctionIamPolicyOutputWithContext(ctx context.Context) FunctionIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamPolicyOutput)
}

func (i *FunctionIamPolicy) ToOutput(ctx context.Context) pulumix.Output[*FunctionIamPolicy] {
	return pulumix.Output[*FunctionIamPolicy]{
		OutputState: i.ToFunctionIamPolicyOutputWithContext(ctx).OutputState,
	}
}

// FunctionIamPolicyArrayInput is an input type that accepts FunctionIamPolicyArray and FunctionIamPolicyArrayOutput values.
// You can construct a concrete instance of `FunctionIamPolicyArrayInput` via:
//
//	FunctionIamPolicyArray{ FunctionIamPolicyArgs{...} }
type FunctionIamPolicyArrayInput interface {
	pulumi.Input

	ToFunctionIamPolicyArrayOutput() FunctionIamPolicyArrayOutput
	ToFunctionIamPolicyArrayOutputWithContext(context.Context) FunctionIamPolicyArrayOutput
}

type FunctionIamPolicyArray []FunctionIamPolicyInput

func (FunctionIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionIamPolicy)(nil)).Elem()
}

func (i FunctionIamPolicyArray) ToFunctionIamPolicyArrayOutput() FunctionIamPolicyArrayOutput {
	return i.ToFunctionIamPolicyArrayOutputWithContext(context.Background())
}

func (i FunctionIamPolicyArray) ToFunctionIamPolicyArrayOutputWithContext(ctx context.Context) FunctionIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamPolicyArrayOutput)
}

func (i FunctionIamPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*FunctionIamPolicy] {
	return pulumix.Output[[]*FunctionIamPolicy]{
		OutputState: i.ToFunctionIamPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// FunctionIamPolicyMapInput is an input type that accepts FunctionIamPolicyMap and FunctionIamPolicyMapOutput values.
// You can construct a concrete instance of `FunctionIamPolicyMapInput` via:
//
//	FunctionIamPolicyMap{ "key": FunctionIamPolicyArgs{...} }
type FunctionIamPolicyMapInput interface {
	pulumi.Input

	ToFunctionIamPolicyMapOutput() FunctionIamPolicyMapOutput
	ToFunctionIamPolicyMapOutputWithContext(context.Context) FunctionIamPolicyMapOutput
}

type FunctionIamPolicyMap map[string]FunctionIamPolicyInput

func (FunctionIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionIamPolicy)(nil)).Elem()
}

func (i FunctionIamPolicyMap) ToFunctionIamPolicyMapOutput() FunctionIamPolicyMapOutput {
	return i.ToFunctionIamPolicyMapOutputWithContext(context.Background())
}

func (i FunctionIamPolicyMap) ToFunctionIamPolicyMapOutputWithContext(ctx context.Context) FunctionIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamPolicyMapOutput)
}

func (i FunctionIamPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FunctionIamPolicy] {
	return pulumix.Output[map[string]*FunctionIamPolicy]{
		OutputState: i.ToFunctionIamPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type FunctionIamPolicyOutput struct{ *pulumi.OutputState }

func (FunctionIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionIamPolicy)(nil)).Elem()
}

func (o FunctionIamPolicyOutput) ToFunctionIamPolicyOutput() FunctionIamPolicyOutput {
	return o
}

func (o FunctionIamPolicyOutput) ToFunctionIamPolicyOutputWithContext(ctx context.Context) FunctionIamPolicyOutput {
	return o
}

func (o FunctionIamPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*FunctionIamPolicy] {
	return pulumix.Output[*FunctionIamPolicy]{
		OutputState: o.OutputState,
	}
}

// Used to find the parent resource to bind the IAM policy to
func (o FunctionIamPolicyOutput) CloudFunction() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamPolicy) pulumi.StringOutput { return v.CloudFunction }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o FunctionIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o FunctionIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
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
func (o FunctionIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
// region is specified, it is taken from the provider configuration.
func (o FunctionIamPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionIamPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type FunctionIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (FunctionIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionIamPolicy)(nil)).Elem()
}

func (o FunctionIamPolicyArrayOutput) ToFunctionIamPolicyArrayOutput() FunctionIamPolicyArrayOutput {
	return o
}

func (o FunctionIamPolicyArrayOutput) ToFunctionIamPolicyArrayOutputWithContext(ctx context.Context) FunctionIamPolicyArrayOutput {
	return o
}

func (o FunctionIamPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FunctionIamPolicy] {
	return pulumix.Output[[]*FunctionIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o FunctionIamPolicyArrayOutput) Index(i pulumi.IntInput) FunctionIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionIamPolicy {
		return vs[0].([]*FunctionIamPolicy)[vs[1].(int)]
	}).(FunctionIamPolicyOutput)
}

type FunctionIamPolicyMapOutput struct{ *pulumi.OutputState }

func (FunctionIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionIamPolicy)(nil)).Elem()
}

func (o FunctionIamPolicyMapOutput) ToFunctionIamPolicyMapOutput() FunctionIamPolicyMapOutput {
	return o
}

func (o FunctionIamPolicyMapOutput) ToFunctionIamPolicyMapOutputWithContext(ctx context.Context) FunctionIamPolicyMapOutput {
	return o
}

func (o FunctionIamPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FunctionIamPolicy] {
	return pulumix.Output[map[string]*FunctionIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o FunctionIamPolicyMapOutput) MapIndex(k pulumi.StringInput) FunctionIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionIamPolicy {
		return vs[0].(map[string]*FunctionIamPolicy)[vs[1].(string)]
	}).(FunctionIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamPolicyInput)(nil)).Elem(), &FunctionIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamPolicyArrayInput)(nil)).Elem(), FunctionIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamPolicyMapInput)(nil)).Elem(), FunctionIamPolicyMap{})
	pulumi.RegisterOutputType(FunctionIamPolicyOutput{})
	pulumi.RegisterOutputType(FunctionIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(FunctionIamPolicyMapOutput{})
}

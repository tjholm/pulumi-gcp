// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Three different resources help you manage your IAM policy for Data catalog PolicyTag. Each of these resources serves a different use case:
//
// * `datacatalog.PolicyTagIamPolicy`: Authoritative. Sets the IAM policy for the policytag and replaces any existing policy already attached.
// * `datacatalog.PolicyTagIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the policytag are preserved.
// * `datacatalog.PolicyTagIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the policytag are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `datacatalog.PolicyTagIamPolicy`: Retrieves the IAM policy for the policytag
//
// > **Note:** `datacatalog.PolicyTagIamPolicy` **cannot** be used in conjunction with `datacatalog.PolicyTagIamBinding` and `datacatalog.PolicyTagIamMember` or they will fight over what your policy should be.
//
// > **Note:** `datacatalog.PolicyTagIamBinding` resources **can be** used in conjunction with `datacatalog.PolicyTagIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_data\_catalog\_policy\_tag\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datacatalog"
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
//			_, err = datacatalog.NewPolicyTagIamPolicy(ctx, "policy", &datacatalog.PolicyTagIamPolicyArgs{
//				PolicyTag:  pulumi.Any(google_data_catalog_policy_tag.Basic_policy_tag.Name),
//				PolicyData: *pulumi.String(admin.PolicyData),
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
// ## google\_data\_catalog\_policy\_tag\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datacatalog.NewPolicyTagIamBinding(ctx, "binding", &datacatalog.PolicyTagIamBindingArgs{
//				PolicyTag: pulumi.Any(google_data_catalog_policy_tag.Basic_policy_tag.Name),
//				Role:      pulumi.String("roles/viewer"),
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
// ## google\_data\_catalog\_policy\_tag\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datacatalog.NewPolicyTagIamMember(ctx, "member", &datacatalog.PolicyTagIamMemberArgs{
//				PolicyTag: pulumi.Any(google_data_catalog_policy_tag.Basic_policy_tag.Name),
//				Role:      pulumi.String("roles/viewer"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* {{policy_tag}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog policytag IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor "{{policy_tag}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor "{{policy_tag}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor {{policy_tag}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type PolicyTagIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
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
	PolicyTag pulumi.StringOutput `pulumi:"policyTag"`
}

// NewPolicyTagIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicyTagIamPolicy(ctx *pulumi.Context,
	name string, args *PolicyTagIamPolicyArgs, opts ...pulumi.ResourceOption) (*PolicyTagIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.PolicyTag == nil {
		return nil, errors.New("invalid value for required argument 'PolicyTag'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyTagIamPolicy
	err := ctx.RegisterResource("gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyTagIamPolicy gets an existing PolicyTagIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyTagIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyTagIamPolicyState, opts ...pulumi.ResourceOption) (*PolicyTagIamPolicy, error) {
	var resource PolicyTagIamPolicy
	err := ctx.ReadResource("gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyTagIamPolicy resources.
type policyTagIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
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
	PolicyTag *string `pulumi:"policyTag"`
}

type PolicyTagIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
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
	PolicyTag pulumi.StringPtrInput
}

func (PolicyTagIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagIamPolicyState)(nil)).Elem()
}

type policyTagIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
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
	PolicyTag string `pulumi:"policyTag"`
}

// The set of arguments for constructing a PolicyTagIamPolicy resource.
type PolicyTagIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
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
	PolicyTag pulumi.StringInput
}

func (PolicyTagIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagIamPolicyArgs)(nil)).Elem()
}

type PolicyTagIamPolicyInput interface {
	pulumi.Input

	ToPolicyTagIamPolicyOutput() PolicyTagIamPolicyOutput
	ToPolicyTagIamPolicyOutputWithContext(ctx context.Context) PolicyTagIamPolicyOutput
}

func (*PolicyTagIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTagIamPolicy)(nil)).Elem()
}

func (i *PolicyTagIamPolicy) ToPolicyTagIamPolicyOutput() PolicyTagIamPolicyOutput {
	return i.ToPolicyTagIamPolicyOutputWithContext(context.Background())
}

func (i *PolicyTagIamPolicy) ToPolicyTagIamPolicyOutputWithContext(ctx context.Context) PolicyTagIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamPolicyOutput)
}

func (i *PolicyTagIamPolicy) ToOutput(ctx context.Context) pulumix.Output[*PolicyTagIamPolicy] {
	return pulumix.Output[*PolicyTagIamPolicy]{
		OutputState: i.ToPolicyTagIamPolicyOutputWithContext(ctx).OutputState,
	}
}

// PolicyTagIamPolicyArrayInput is an input type that accepts PolicyTagIamPolicyArray and PolicyTagIamPolicyArrayOutput values.
// You can construct a concrete instance of `PolicyTagIamPolicyArrayInput` via:
//
//	PolicyTagIamPolicyArray{ PolicyTagIamPolicyArgs{...} }
type PolicyTagIamPolicyArrayInput interface {
	pulumi.Input

	ToPolicyTagIamPolicyArrayOutput() PolicyTagIamPolicyArrayOutput
	ToPolicyTagIamPolicyArrayOutputWithContext(context.Context) PolicyTagIamPolicyArrayOutput
}

type PolicyTagIamPolicyArray []PolicyTagIamPolicyInput

func (PolicyTagIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyTagIamPolicy)(nil)).Elem()
}

func (i PolicyTagIamPolicyArray) ToPolicyTagIamPolicyArrayOutput() PolicyTagIamPolicyArrayOutput {
	return i.ToPolicyTagIamPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyTagIamPolicyArray) ToPolicyTagIamPolicyArrayOutputWithContext(ctx context.Context) PolicyTagIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamPolicyArrayOutput)
}

func (i PolicyTagIamPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*PolicyTagIamPolicy] {
	return pulumix.Output[[]*PolicyTagIamPolicy]{
		OutputState: i.ToPolicyTagIamPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// PolicyTagIamPolicyMapInput is an input type that accepts PolicyTagIamPolicyMap and PolicyTagIamPolicyMapOutput values.
// You can construct a concrete instance of `PolicyTagIamPolicyMapInput` via:
//
//	PolicyTagIamPolicyMap{ "key": PolicyTagIamPolicyArgs{...} }
type PolicyTagIamPolicyMapInput interface {
	pulumi.Input

	ToPolicyTagIamPolicyMapOutput() PolicyTagIamPolicyMapOutput
	ToPolicyTagIamPolicyMapOutputWithContext(context.Context) PolicyTagIamPolicyMapOutput
}

type PolicyTagIamPolicyMap map[string]PolicyTagIamPolicyInput

func (PolicyTagIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyTagIamPolicy)(nil)).Elem()
}

func (i PolicyTagIamPolicyMap) ToPolicyTagIamPolicyMapOutput() PolicyTagIamPolicyMapOutput {
	return i.ToPolicyTagIamPolicyMapOutputWithContext(context.Background())
}

func (i PolicyTagIamPolicyMap) ToPolicyTagIamPolicyMapOutputWithContext(ctx context.Context) PolicyTagIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamPolicyMapOutput)
}

func (i PolicyTagIamPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PolicyTagIamPolicy] {
	return pulumix.Output[map[string]*PolicyTagIamPolicy]{
		OutputState: i.ToPolicyTagIamPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type PolicyTagIamPolicyOutput struct{ *pulumi.OutputState }

func (PolicyTagIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTagIamPolicy)(nil)).Elem()
}

func (o PolicyTagIamPolicyOutput) ToPolicyTagIamPolicyOutput() PolicyTagIamPolicyOutput {
	return o
}

func (o PolicyTagIamPolicyOutput) ToPolicyTagIamPolicyOutputWithContext(ctx context.Context) PolicyTagIamPolicyOutput {
	return o
}

func (o PolicyTagIamPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*PolicyTagIamPolicy] {
	return pulumix.Output[*PolicyTagIamPolicy]{
		OutputState: o.OutputState,
	}
}

// (Computed) The etag of the IAM policy.
func (o PolicyTagIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyTagIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o PolicyTagIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyTagIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
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
func (o PolicyTagIamPolicyOutput) PolicyTag() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyTagIamPolicy) pulumi.StringOutput { return v.PolicyTag }).(pulumi.StringOutput)
}

type PolicyTagIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyTagIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyTagIamPolicy)(nil)).Elem()
}

func (o PolicyTagIamPolicyArrayOutput) ToPolicyTagIamPolicyArrayOutput() PolicyTagIamPolicyArrayOutput {
	return o
}

func (o PolicyTagIamPolicyArrayOutput) ToPolicyTagIamPolicyArrayOutputWithContext(ctx context.Context) PolicyTagIamPolicyArrayOutput {
	return o
}

func (o PolicyTagIamPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PolicyTagIamPolicy] {
	return pulumix.Output[[]*PolicyTagIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o PolicyTagIamPolicyArrayOutput) Index(i pulumi.IntInput) PolicyTagIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyTagIamPolicy {
		return vs[0].([]*PolicyTagIamPolicy)[vs[1].(int)]
	}).(PolicyTagIamPolicyOutput)
}

type PolicyTagIamPolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyTagIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyTagIamPolicy)(nil)).Elem()
}

func (o PolicyTagIamPolicyMapOutput) ToPolicyTagIamPolicyMapOutput() PolicyTagIamPolicyMapOutput {
	return o
}

func (o PolicyTagIamPolicyMapOutput) ToPolicyTagIamPolicyMapOutputWithContext(ctx context.Context) PolicyTagIamPolicyMapOutput {
	return o
}

func (o PolicyTagIamPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PolicyTagIamPolicy] {
	return pulumix.Output[map[string]*PolicyTagIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o PolicyTagIamPolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyTagIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyTagIamPolicy {
		return vs[0].(map[string]*PolicyTagIamPolicy)[vs[1].(string)]
	}).(PolicyTagIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTagIamPolicyInput)(nil)).Elem(), &PolicyTagIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTagIamPolicyArrayInput)(nil)).Elem(), PolicyTagIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTagIamPolicyMapInput)(nil)).Elem(), PolicyTagIamPolicyMap{})
	pulumi.RegisterOutputType(PolicyTagIamPolicyOutput{})
	pulumi.RegisterOutputType(PolicyTagIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyTagIamPolicyMapOutput{})
}

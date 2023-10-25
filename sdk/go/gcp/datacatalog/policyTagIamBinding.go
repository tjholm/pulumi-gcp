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
//	$ pulumi import gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding editor "{{policy_tag}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding editor "{{policy_tag}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding editor {{policy_tag}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type PolicyTagIamBinding struct {
	pulumi.CustomResourceState

	Condition PolicyTagIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
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
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewPolicyTagIamBinding registers a new resource with the given unique name, arguments, and options.
func NewPolicyTagIamBinding(ctx *pulumi.Context,
	name string, args *PolicyTagIamBindingArgs, opts ...pulumi.ResourceOption) (*PolicyTagIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.PolicyTag == nil {
		return nil, errors.New("invalid value for required argument 'PolicyTag'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyTagIamBinding
	err := ctx.RegisterResource("gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyTagIamBinding gets an existing PolicyTagIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyTagIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyTagIamBindingState, opts ...pulumi.ResourceOption) (*PolicyTagIamBinding, error) {
	var resource PolicyTagIamBinding
	err := ctx.ReadResource("gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyTagIamBinding resources.
type policyTagIamBindingState struct {
	Condition *PolicyTagIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
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
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type PolicyTagIamBindingState struct {
	Condition PolicyTagIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
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
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (PolicyTagIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagIamBindingState)(nil)).Elem()
}

type policyTagIamBindingArgs struct {
	Condition *PolicyTagIamBindingCondition `pulumi:"condition"`
	Members   []string                      `pulumi:"members"`
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
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a PolicyTagIamBinding resource.
type PolicyTagIamBindingArgs struct {
	Condition PolicyTagIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
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
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (PolicyTagIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagIamBindingArgs)(nil)).Elem()
}

type PolicyTagIamBindingInput interface {
	pulumi.Input

	ToPolicyTagIamBindingOutput() PolicyTagIamBindingOutput
	ToPolicyTagIamBindingOutputWithContext(ctx context.Context) PolicyTagIamBindingOutput
}

func (*PolicyTagIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTagIamBinding)(nil)).Elem()
}

func (i *PolicyTagIamBinding) ToPolicyTagIamBindingOutput() PolicyTagIamBindingOutput {
	return i.ToPolicyTagIamBindingOutputWithContext(context.Background())
}

func (i *PolicyTagIamBinding) ToPolicyTagIamBindingOutputWithContext(ctx context.Context) PolicyTagIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamBindingOutput)
}

func (i *PolicyTagIamBinding) ToOutput(ctx context.Context) pulumix.Output[*PolicyTagIamBinding] {
	return pulumix.Output[*PolicyTagIamBinding]{
		OutputState: i.ToPolicyTagIamBindingOutputWithContext(ctx).OutputState,
	}
}

// PolicyTagIamBindingArrayInput is an input type that accepts PolicyTagIamBindingArray and PolicyTagIamBindingArrayOutput values.
// You can construct a concrete instance of `PolicyTagIamBindingArrayInput` via:
//
//	PolicyTagIamBindingArray{ PolicyTagIamBindingArgs{...} }
type PolicyTagIamBindingArrayInput interface {
	pulumi.Input

	ToPolicyTagIamBindingArrayOutput() PolicyTagIamBindingArrayOutput
	ToPolicyTagIamBindingArrayOutputWithContext(context.Context) PolicyTagIamBindingArrayOutput
}

type PolicyTagIamBindingArray []PolicyTagIamBindingInput

func (PolicyTagIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyTagIamBinding)(nil)).Elem()
}

func (i PolicyTagIamBindingArray) ToPolicyTagIamBindingArrayOutput() PolicyTagIamBindingArrayOutput {
	return i.ToPolicyTagIamBindingArrayOutputWithContext(context.Background())
}

func (i PolicyTagIamBindingArray) ToPolicyTagIamBindingArrayOutputWithContext(ctx context.Context) PolicyTagIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamBindingArrayOutput)
}

func (i PolicyTagIamBindingArray) ToOutput(ctx context.Context) pulumix.Output[[]*PolicyTagIamBinding] {
	return pulumix.Output[[]*PolicyTagIamBinding]{
		OutputState: i.ToPolicyTagIamBindingArrayOutputWithContext(ctx).OutputState,
	}
}

// PolicyTagIamBindingMapInput is an input type that accepts PolicyTagIamBindingMap and PolicyTagIamBindingMapOutput values.
// You can construct a concrete instance of `PolicyTagIamBindingMapInput` via:
//
//	PolicyTagIamBindingMap{ "key": PolicyTagIamBindingArgs{...} }
type PolicyTagIamBindingMapInput interface {
	pulumi.Input

	ToPolicyTagIamBindingMapOutput() PolicyTagIamBindingMapOutput
	ToPolicyTagIamBindingMapOutputWithContext(context.Context) PolicyTagIamBindingMapOutput
}

type PolicyTagIamBindingMap map[string]PolicyTagIamBindingInput

func (PolicyTagIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyTagIamBinding)(nil)).Elem()
}

func (i PolicyTagIamBindingMap) ToPolicyTagIamBindingMapOutput() PolicyTagIamBindingMapOutput {
	return i.ToPolicyTagIamBindingMapOutputWithContext(context.Background())
}

func (i PolicyTagIamBindingMap) ToPolicyTagIamBindingMapOutputWithContext(ctx context.Context) PolicyTagIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamBindingMapOutput)
}

func (i PolicyTagIamBindingMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PolicyTagIamBinding] {
	return pulumix.Output[map[string]*PolicyTagIamBinding]{
		OutputState: i.ToPolicyTagIamBindingMapOutputWithContext(ctx).OutputState,
	}
}

type PolicyTagIamBindingOutput struct{ *pulumi.OutputState }

func (PolicyTagIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTagIamBinding)(nil)).Elem()
}

func (o PolicyTagIamBindingOutput) ToPolicyTagIamBindingOutput() PolicyTagIamBindingOutput {
	return o
}

func (o PolicyTagIamBindingOutput) ToPolicyTagIamBindingOutputWithContext(ctx context.Context) PolicyTagIamBindingOutput {
	return o
}

func (o PolicyTagIamBindingOutput) ToOutput(ctx context.Context) pulumix.Output[*PolicyTagIamBinding] {
	return pulumix.Output[*PolicyTagIamBinding]{
		OutputState: o.OutputState,
	}
}

func (o PolicyTagIamBindingOutput) Condition() PolicyTagIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *PolicyTagIamBinding) PolicyTagIamBindingConditionPtrOutput { return v.Condition }).(PolicyTagIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o PolicyTagIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyTagIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PolicyTagIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyTagIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
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
func (o PolicyTagIamBindingOutput) PolicyTag() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyTagIamBinding) pulumi.StringOutput { return v.PolicyTag }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o PolicyTagIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyTagIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type PolicyTagIamBindingArrayOutput struct{ *pulumi.OutputState }

func (PolicyTagIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyTagIamBinding)(nil)).Elem()
}

func (o PolicyTagIamBindingArrayOutput) ToPolicyTagIamBindingArrayOutput() PolicyTagIamBindingArrayOutput {
	return o
}

func (o PolicyTagIamBindingArrayOutput) ToPolicyTagIamBindingArrayOutputWithContext(ctx context.Context) PolicyTagIamBindingArrayOutput {
	return o
}

func (o PolicyTagIamBindingArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PolicyTagIamBinding] {
	return pulumix.Output[[]*PolicyTagIamBinding]{
		OutputState: o.OutputState,
	}
}

func (o PolicyTagIamBindingArrayOutput) Index(i pulumi.IntInput) PolicyTagIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyTagIamBinding {
		return vs[0].([]*PolicyTagIamBinding)[vs[1].(int)]
	}).(PolicyTagIamBindingOutput)
}

type PolicyTagIamBindingMapOutput struct{ *pulumi.OutputState }

func (PolicyTagIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyTagIamBinding)(nil)).Elem()
}

func (o PolicyTagIamBindingMapOutput) ToPolicyTagIamBindingMapOutput() PolicyTagIamBindingMapOutput {
	return o
}

func (o PolicyTagIamBindingMapOutput) ToPolicyTagIamBindingMapOutputWithContext(ctx context.Context) PolicyTagIamBindingMapOutput {
	return o
}

func (o PolicyTagIamBindingMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PolicyTagIamBinding] {
	return pulumix.Output[map[string]*PolicyTagIamBinding]{
		OutputState: o.OutputState,
	}
}

func (o PolicyTagIamBindingMapOutput) MapIndex(k pulumi.StringInput) PolicyTagIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyTagIamBinding {
		return vs[0].(map[string]*PolicyTagIamBinding)[vs[1].(string)]
	}).(PolicyTagIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTagIamBindingInput)(nil)).Elem(), &PolicyTagIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTagIamBindingArrayInput)(nil)).Elem(), PolicyTagIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTagIamBindingMapInput)(nil)).Elem(), PolicyTagIamBindingMap{})
	pulumi.RegisterOutputType(PolicyTagIamBindingOutput{})
	pulumi.RegisterOutputType(PolicyTagIamBindingArrayOutput{})
	pulumi.RegisterOutputType(PolicyTagIamBindingMapOutput{})
}

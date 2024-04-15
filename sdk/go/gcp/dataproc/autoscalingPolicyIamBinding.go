// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Dataproc AutoscalingPolicy. Each of these resources serves a different use case:
//
// * `dataproc.AutoscalingPolicyIamPolicy`: Authoritative. Sets the IAM policy for the autoscalingpolicy and replaces any existing policy already attached.
// * `dataproc.AutoscalingPolicyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the autoscalingpolicy are preserved.
// * `dataproc.AutoscalingPolicyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the autoscalingpolicy are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dataproc.AutoscalingPolicyIamPolicy`: Retrieves the IAM policy for the autoscalingpolicy
//
// > **Note:** `dataproc.AutoscalingPolicyIamPolicy` **cannot** be used in conjunction with `dataproc.AutoscalingPolicyIamBinding` and `dataproc.AutoscalingPolicyIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dataproc.AutoscalingPolicyIamBinding` resources **can be** used in conjunction with `dataproc.AutoscalingPolicyIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_dataproc\_autoscaling\_policy\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
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
//			_, err = dataproc.NewAutoscalingPolicyIamPolicy(ctx, "policy", &dataproc.AutoscalingPolicyIamPolicyArgs{
//				Project:    pulumi.Any(basic.Project),
//				Location:   pulumi.Any(basic.Location),
//				PolicyId:   pulumi.Any(basic.PolicyId),
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
// ## google\_dataproc\_autoscaling\_policy\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataproc.NewAutoscalingPolicyIamBinding(ctx, "binding", &dataproc.AutoscalingPolicyIamBindingArgs{
//				Project:  pulumi.Any(basic.Project),
//				Location: pulumi.Any(basic.Location),
//				PolicyId: pulumi.Any(basic.PolicyId),
//				Role:     pulumi.String("roles/viewer"),
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
// ## google\_dataproc\_autoscaling\_policy\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataproc.NewAutoscalingPolicyIamMember(ctx, "member", &dataproc.AutoscalingPolicyIamMemberArgs{
//				Project:  pulumi.Any(basic.Project),
//				Location: pulumi.Any(basic.Location),
//				PolicyId: pulumi.Any(basic.PolicyId),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
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
// ## google\_dataproc\_autoscaling\_policy\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
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
//			_, err = dataproc.NewAutoscalingPolicyIamPolicy(ctx, "policy", &dataproc.AutoscalingPolicyIamPolicyArgs{
//				Project:    pulumi.Any(basic.Project),
//				Location:   pulumi.Any(basic.Location),
//				PolicyId:   pulumi.Any(basic.PolicyId),
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
// ## google\_dataproc\_autoscaling\_policy\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataproc.NewAutoscalingPolicyIamBinding(ctx, "binding", &dataproc.AutoscalingPolicyIamBindingArgs{
//				Project:  pulumi.Any(basic.Project),
//				Location: pulumi.Any(basic.Location),
//				PolicyId: pulumi.Any(basic.PolicyId),
//				Role:     pulumi.String("roles/viewer"),
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
// ## google\_dataproc\_autoscaling\_policy\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataproc.NewAutoscalingPolicyIamMember(ctx, "member", &dataproc.AutoscalingPolicyIamMemberArgs{
//				Project:  pulumi.Any(basic.Project),
//				Location: pulumi.Any(basic.Location),
//				PolicyId: pulumi.Any(basic.PolicyId),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
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
// * projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}
//
// * {{project}}/{{location}}/{{policy_id}}
//
// * {{location}}/{{policy_id}}
//
// * {{policy_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Dataproc autoscalingpolicy IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:dataproc/autoscalingPolicyIamBinding:AutoscalingPolicyIamBinding editor "projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:dataproc/autoscalingPolicyIamBinding:AutoscalingPolicyIamBinding editor "projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:dataproc/autoscalingPolicyIamBinding:AutoscalingPolicyIamBinding editor projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AutoscalingPolicyIamBinding struct {
	pulumi.CustomResourceState

	Condition AutoscalingPolicyIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
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
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 50 characters.
	// Used to find the parent resource to bind the IAM policy to
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.AutoscalingPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAutoscalingPolicyIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAutoscalingPolicyIamBinding(ctx *pulumi.Context,
	name string, args *AutoscalingPolicyIamBindingArgs, opts ...pulumi.ResourceOption) (*AutoscalingPolicyIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AutoscalingPolicyIamBinding
	err := ctx.RegisterResource("gcp:dataproc/autoscalingPolicyIamBinding:AutoscalingPolicyIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscalingPolicyIamBinding gets an existing AutoscalingPolicyIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscalingPolicyIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscalingPolicyIamBindingState, opts ...pulumi.ResourceOption) (*AutoscalingPolicyIamBinding, error) {
	var resource AutoscalingPolicyIamBinding
	err := ctx.ReadResource("gcp:dataproc/autoscalingPolicyIamBinding:AutoscalingPolicyIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoscalingPolicyIamBinding resources.
type autoscalingPolicyIamBindingState struct {
	Condition *AutoscalingPolicyIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
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
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 50 characters.
	// Used to find the parent resource to bind the IAM policy to
	PolicyId *string `pulumi:"policyId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.AutoscalingPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type AutoscalingPolicyIamBindingState struct {
	Condition AutoscalingPolicyIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
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
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 50 characters.
	// Used to find the parent resource to bind the IAM policy to
	PolicyId pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.AutoscalingPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (AutoscalingPolicyIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalingPolicyIamBindingState)(nil)).Elem()
}

type autoscalingPolicyIamBindingArgs struct {
	Condition *AutoscalingPolicyIamBindingCondition `pulumi:"condition"`
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
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
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 50 characters.
	// Used to find the parent resource to bind the IAM policy to
	PolicyId string `pulumi:"policyId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.AutoscalingPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AutoscalingPolicyIamBinding resource.
type AutoscalingPolicyIamBindingArgs struct {
	Condition AutoscalingPolicyIamBindingConditionPtrInput
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
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
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 50 characters.
	// Used to find the parent resource to bind the IAM policy to
	PolicyId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.AutoscalingPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (AutoscalingPolicyIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalingPolicyIamBindingArgs)(nil)).Elem()
}

type AutoscalingPolicyIamBindingInput interface {
	pulumi.Input

	ToAutoscalingPolicyIamBindingOutput() AutoscalingPolicyIamBindingOutput
	ToAutoscalingPolicyIamBindingOutputWithContext(ctx context.Context) AutoscalingPolicyIamBindingOutput
}

func (*AutoscalingPolicyIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscalingPolicyIamBinding)(nil)).Elem()
}

func (i *AutoscalingPolicyIamBinding) ToAutoscalingPolicyIamBindingOutput() AutoscalingPolicyIamBindingOutput {
	return i.ToAutoscalingPolicyIamBindingOutputWithContext(context.Background())
}

func (i *AutoscalingPolicyIamBinding) ToAutoscalingPolicyIamBindingOutputWithContext(ctx context.Context) AutoscalingPolicyIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingPolicyIamBindingOutput)
}

// AutoscalingPolicyIamBindingArrayInput is an input type that accepts AutoscalingPolicyIamBindingArray and AutoscalingPolicyIamBindingArrayOutput values.
// You can construct a concrete instance of `AutoscalingPolicyIamBindingArrayInput` via:
//
//	AutoscalingPolicyIamBindingArray{ AutoscalingPolicyIamBindingArgs{...} }
type AutoscalingPolicyIamBindingArrayInput interface {
	pulumi.Input

	ToAutoscalingPolicyIamBindingArrayOutput() AutoscalingPolicyIamBindingArrayOutput
	ToAutoscalingPolicyIamBindingArrayOutputWithContext(context.Context) AutoscalingPolicyIamBindingArrayOutput
}

type AutoscalingPolicyIamBindingArray []AutoscalingPolicyIamBindingInput

func (AutoscalingPolicyIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoscalingPolicyIamBinding)(nil)).Elem()
}

func (i AutoscalingPolicyIamBindingArray) ToAutoscalingPolicyIamBindingArrayOutput() AutoscalingPolicyIamBindingArrayOutput {
	return i.ToAutoscalingPolicyIamBindingArrayOutputWithContext(context.Background())
}

func (i AutoscalingPolicyIamBindingArray) ToAutoscalingPolicyIamBindingArrayOutputWithContext(ctx context.Context) AutoscalingPolicyIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingPolicyIamBindingArrayOutput)
}

// AutoscalingPolicyIamBindingMapInput is an input type that accepts AutoscalingPolicyIamBindingMap and AutoscalingPolicyIamBindingMapOutput values.
// You can construct a concrete instance of `AutoscalingPolicyIamBindingMapInput` via:
//
//	AutoscalingPolicyIamBindingMap{ "key": AutoscalingPolicyIamBindingArgs{...} }
type AutoscalingPolicyIamBindingMapInput interface {
	pulumi.Input

	ToAutoscalingPolicyIamBindingMapOutput() AutoscalingPolicyIamBindingMapOutput
	ToAutoscalingPolicyIamBindingMapOutputWithContext(context.Context) AutoscalingPolicyIamBindingMapOutput
}

type AutoscalingPolicyIamBindingMap map[string]AutoscalingPolicyIamBindingInput

func (AutoscalingPolicyIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoscalingPolicyIamBinding)(nil)).Elem()
}

func (i AutoscalingPolicyIamBindingMap) ToAutoscalingPolicyIamBindingMapOutput() AutoscalingPolicyIamBindingMapOutput {
	return i.ToAutoscalingPolicyIamBindingMapOutputWithContext(context.Background())
}

func (i AutoscalingPolicyIamBindingMap) ToAutoscalingPolicyIamBindingMapOutputWithContext(ctx context.Context) AutoscalingPolicyIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingPolicyIamBindingMapOutput)
}

type AutoscalingPolicyIamBindingOutput struct{ *pulumi.OutputState }

func (AutoscalingPolicyIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscalingPolicyIamBinding)(nil)).Elem()
}

func (o AutoscalingPolicyIamBindingOutput) ToAutoscalingPolicyIamBindingOutput() AutoscalingPolicyIamBindingOutput {
	return o
}

func (o AutoscalingPolicyIamBindingOutput) ToAutoscalingPolicyIamBindingOutputWithContext(ctx context.Context) AutoscalingPolicyIamBindingOutput {
	return o
}

func (o AutoscalingPolicyIamBindingOutput) Condition() AutoscalingPolicyIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *AutoscalingPolicyIamBinding) AutoscalingPolicyIamBindingConditionPtrOutput { return v.Condition }).(AutoscalingPolicyIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o AutoscalingPolicyIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoscalingPolicyIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The  location where the autoscaling policy should reside.
// The default value is `global`.
// Used to find the parent resource to bind the IAM policy to
func (o AutoscalingPolicyIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoscalingPolicyIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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
func (o AutoscalingPolicyIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoscalingPolicyIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
// 3 and 50 characters.
// Used to find the parent resource to bind the IAM policy to
func (o AutoscalingPolicyIamBindingOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoscalingPolicyIamBinding) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o AutoscalingPolicyIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoscalingPolicyIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dataproc.AutoscalingPolicyIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o AutoscalingPolicyIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoscalingPolicyIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type AutoscalingPolicyIamBindingArrayOutput struct{ *pulumi.OutputState }

func (AutoscalingPolicyIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoscalingPolicyIamBinding)(nil)).Elem()
}

func (o AutoscalingPolicyIamBindingArrayOutput) ToAutoscalingPolicyIamBindingArrayOutput() AutoscalingPolicyIamBindingArrayOutput {
	return o
}

func (o AutoscalingPolicyIamBindingArrayOutput) ToAutoscalingPolicyIamBindingArrayOutputWithContext(ctx context.Context) AutoscalingPolicyIamBindingArrayOutput {
	return o
}

func (o AutoscalingPolicyIamBindingArrayOutput) Index(i pulumi.IntInput) AutoscalingPolicyIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AutoscalingPolicyIamBinding {
		return vs[0].([]*AutoscalingPolicyIamBinding)[vs[1].(int)]
	}).(AutoscalingPolicyIamBindingOutput)
}

type AutoscalingPolicyIamBindingMapOutput struct{ *pulumi.OutputState }

func (AutoscalingPolicyIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoscalingPolicyIamBinding)(nil)).Elem()
}

func (o AutoscalingPolicyIamBindingMapOutput) ToAutoscalingPolicyIamBindingMapOutput() AutoscalingPolicyIamBindingMapOutput {
	return o
}

func (o AutoscalingPolicyIamBindingMapOutput) ToAutoscalingPolicyIamBindingMapOutputWithContext(ctx context.Context) AutoscalingPolicyIamBindingMapOutput {
	return o
}

func (o AutoscalingPolicyIamBindingMapOutput) MapIndex(k pulumi.StringInput) AutoscalingPolicyIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AutoscalingPolicyIamBinding {
		return vs[0].(map[string]*AutoscalingPolicyIamBinding)[vs[1].(string)]
	}).(AutoscalingPolicyIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscalingPolicyIamBindingInput)(nil)).Elem(), &AutoscalingPolicyIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscalingPolicyIamBindingArrayInput)(nil)).Elem(), AutoscalingPolicyIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscalingPolicyIamBindingMapInput)(nil)).Elem(), AutoscalingPolicyIamBindingMap{})
	pulumi.RegisterOutputType(AutoscalingPolicyIamBindingOutput{})
	pulumi.RegisterOutputType(AutoscalingPolicyIamBindingArrayOutput{})
	pulumi.RegisterOutputType(AutoscalingPolicyIamBindingMapOutput{})
}

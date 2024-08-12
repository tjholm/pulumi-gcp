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

// Three different resources help you manage your IAM policy for API Gateway Gateway. Each of these resources serves a different use case:
//
// * `apigateway.GatewayIamPolicy`: Authoritative. Sets the IAM policy for the gateway and replaces any existing policy already attached.
// * `apigateway.GatewayIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the gateway are preserved.
// * `apigateway.GatewayIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the gateway are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `apigateway.GatewayIamPolicy`: Retrieves the IAM policy for the gateway
//
// > **Note:** `apigateway.GatewayIamPolicy` **cannot** be used in conjunction with `apigateway.GatewayIamBinding` and `apigateway.GatewayIamMember` or they will fight over what your policy should be.
//
// > **Note:** `apigateway.GatewayIamBinding` resources **can be** used in conjunction with `apigateway.GatewayIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_api\_gateway\_gateway\_iam\_policy
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
//			_, err = apigateway.NewGatewayIamPolicy(ctx, "policy", &apigateway.GatewayIamPolicyArgs{
//				Project:    pulumi.Any(apiGw.Project),
//				Region:     pulumi.Any(apiGw.Region),
//				Gateway:    pulumi.Any(apiGw.GatewayId),
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
// ## apigateway.GatewayIamBinding
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
//			_, err := apigateway.NewGatewayIamBinding(ctx, "binding", &apigateway.GatewayIamBindingArgs{
//				Project: pulumi.Any(apiGw.Project),
//				Region:  pulumi.Any(apiGw.Region),
//				Gateway: pulumi.Any(apiGw.GatewayId),
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
//
// ## apigateway.GatewayIamMember
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
//			_, err := apigateway.NewGatewayIamMember(ctx, "member", &apigateway.GatewayIamMemberArgs{
//				Project: pulumi.Any(apiGw.Project),
//				Region:  pulumi.Any(apiGw.Region),
//				Gateway: pulumi.Any(apiGw.GatewayId),
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
//
// ## This resource supports User Project Overrides.
//
// -
//
// # IAM policy for API Gateway Gateway
// Three different resources help you manage your IAM policy for API Gateway Gateway. Each of these resources serves a different use case:
//
// * `apigateway.GatewayIamPolicy`: Authoritative. Sets the IAM policy for the gateway and replaces any existing policy already attached.
// * `apigateway.GatewayIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the gateway are preserved.
// * `apigateway.GatewayIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the gateway are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `apigateway.GatewayIamPolicy`: Retrieves the IAM policy for the gateway
//
// > **Note:** `apigateway.GatewayIamPolicy` **cannot** be used in conjunction with `apigateway.GatewayIamBinding` and `apigateway.GatewayIamMember` or they will fight over what your policy should be.
//
// > **Note:** `apigateway.GatewayIamBinding` resources **can be** used in conjunction with `apigateway.GatewayIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_api\_gateway\_gateway\_iam\_policy
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
//			_, err = apigateway.NewGatewayIamPolicy(ctx, "policy", &apigateway.GatewayIamPolicyArgs{
//				Project:    pulumi.Any(apiGw.Project),
//				Region:     pulumi.Any(apiGw.Region),
//				Gateway:    pulumi.Any(apiGw.GatewayId),
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
// ## apigateway.GatewayIamBinding
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
//			_, err := apigateway.NewGatewayIamBinding(ctx, "binding", &apigateway.GatewayIamBindingArgs{
//				Project: pulumi.Any(apiGw.Project),
//				Region:  pulumi.Any(apiGw.Region),
//				Gateway: pulumi.Any(apiGw.GatewayId),
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
//
// ## apigateway.GatewayIamMember
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
//			_, err := apigateway.NewGatewayIamMember(ctx, "member", &apigateway.GatewayIamMemberArgs{
//				Project: pulumi.Any(apiGw.Project),
//				Region:  pulumi.Any(apiGw.Region),
//				Gateway: pulumi.Any(apiGw.GatewayId),
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
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/locations/{{region}}/gateways/{{gateway}}
//
// * {{project}}/{{region}}/{{gateway}}
//
// * {{region}}/{{gateway}}
//
// * {{gateway}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// API Gateway gateway IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/gatewayIamBinding:GatewayIamBinding editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/gatewayIamBinding:GatewayIamBinding editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:apigateway/gatewayIamBinding:GatewayIamBinding editor projects/{{project}}/locations/{{region}}/gateways/{{gateway}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type GatewayIamBinding struct {
	pulumi.CustomResourceState

	Condition GatewayIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput `pulumi:"etag"`
	Gateway pulumi.StringOutput `pulumi:"gateway"`
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
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewGatewayIamBinding registers a new resource with the given unique name, arguments, and options.
func NewGatewayIamBinding(ctx *pulumi.Context,
	name string, args *GatewayIamBindingArgs, opts ...pulumi.ResourceOption) (*GatewayIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayIamBinding
	err := ctx.RegisterResource("gcp:apigateway/gatewayIamBinding:GatewayIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayIamBinding gets an existing GatewayIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayIamBindingState, opts ...pulumi.ResourceOption) (*GatewayIamBinding, error) {
	var resource GatewayIamBinding
	err := ctx.ReadResource("gcp:apigateway/gatewayIamBinding:GatewayIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayIamBinding resources.
type gatewayIamBindingState struct {
	Condition *GatewayIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string `pulumi:"etag"`
	Gateway *string `pulumi:"gateway"`
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
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type GatewayIamBindingState struct {
	Condition GatewayIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Gateway pulumi.StringPtrInput
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
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (GatewayIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayIamBindingState)(nil)).Elem()
}

type gatewayIamBindingArgs struct {
	Condition *GatewayIamBindingCondition `pulumi:"condition"`
	Gateway   string                      `pulumi:"gateway"`
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
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a GatewayIamBinding resource.
type GatewayIamBindingArgs struct {
	Condition GatewayIamBindingConditionPtrInput
	Gateway   pulumi.StringInput
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
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (GatewayIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayIamBindingArgs)(nil)).Elem()
}

type GatewayIamBindingInput interface {
	pulumi.Input

	ToGatewayIamBindingOutput() GatewayIamBindingOutput
	ToGatewayIamBindingOutputWithContext(ctx context.Context) GatewayIamBindingOutput
}

func (*GatewayIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayIamBinding)(nil)).Elem()
}

func (i *GatewayIamBinding) ToGatewayIamBindingOutput() GatewayIamBindingOutput {
	return i.ToGatewayIamBindingOutputWithContext(context.Background())
}

func (i *GatewayIamBinding) ToGatewayIamBindingOutputWithContext(ctx context.Context) GatewayIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamBindingOutput)
}

// GatewayIamBindingArrayInput is an input type that accepts GatewayIamBindingArray and GatewayIamBindingArrayOutput values.
// You can construct a concrete instance of `GatewayIamBindingArrayInput` via:
//
//	GatewayIamBindingArray{ GatewayIamBindingArgs{...} }
type GatewayIamBindingArrayInput interface {
	pulumi.Input

	ToGatewayIamBindingArrayOutput() GatewayIamBindingArrayOutput
	ToGatewayIamBindingArrayOutputWithContext(context.Context) GatewayIamBindingArrayOutput
}

type GatewayIamBindingArray []GatewayIamBindingInput

func (GatewayIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayIamBinding)(nil)).Elem()
}

func (i GatewayIamBindingArray) ToGatewayIamBindingArrayOutput() GatewayIamBindingArrayOutput {
	return i.ToGatewayIamBindingArrayOutputWithContext(context.Background())
}

func (i GatewayIamBindingArray) ToGatewayIamBindingArrayOutputWithContext(ctx context.Context) GatewayIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamBindingArrayOutput)
}

// GatewayIamBindingMapInput is an input type that accepts GatewayIamBindingMap and GatewayIamBindingMapOutput values.
// You can construct a concrete instance of `GatewayIamBindingMapInput` via:
//
//	GatewayIamBindingMap{ "key": GatewayIamBindingArgs{...} }
type GatewayIamBindingMapInput interface {
	pulumi.Input

	ToGatewayIamBindingMapOutput() GatewayIamBindingMapOutput
	ToGatewayIamBindingMapOutputWithContext(context.Context) GatewayIamBindingMapOutput
}

type GatewayIamBindingMap map[string]GatewayIamBindingInput

func (GatewayIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayIamBinding)(nil)).Elem()
}

func (i GatewayIamBindingMap) ToGatewayIamBindingMapOutput() GatewayIamBindingMapOutput {
	return i.ToGatewayIamBindingMapOutputWithContext(context.Background())
}

func (i GatewayIamBindingMap) ToGatewayIamBindingMapOutputWithContext(ctx context.Context) GatewayIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamBindingMapOutput)
}

type GatewayIamBindingOutput struct{ *pulumi.OutputState }

func (GatewayIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayIamBinding)(nil)).Elem()
}

func (o GatewayIamBindingOutput) ToGatewayIamBindingOutput() GatewayIamBindingOutput {
	return o
}

func (o GatewayIamBindingOutput) ToGatewayIamBindingOutputWithContext(ctx context.Context) GatewayIamBindingOutput {
	return o
}

func (o GatewayIamBindingOutput) Condition() GatewayIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *GatewayIamBinding) GatewayIamBindingConditionPtrOutput { return v.Condition }).(GatewayIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o GatewayIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o GatewayIamBindingOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayIamBinding) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
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
func (o GatewayIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o GatewayIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region of the gateway for the API.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
// region is specified, it is taken from the provider configuration.
func (o GatewayIamBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayIamBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o GatewayIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type GatewayIamBindingArrayOutput struct{ *pulumi.OutputState }

func (GatewayIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayIamBinding)(nil)).Elem()
}

func (o GatewayIamBindingArrayOutput) ToGatewayIamBindingArrayOutput() GatewayIamBindingArrayOutput {
	return o
}

func (o GatewayIamBindingArrayOutput) ToGatewayIamBindingArrayOutputWithContext(ctx context.Context) GatewayIamBindingArrayOutput {
	return o
}

func (o GatewayIamBindingArrayOutput) Index(i pulumi.IntInput) GatewayIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayIamBinding {
		return vs[0].([]*GatewayIamBinding)[vs[1].(int)]
	}).(GatewayIamBindingOutput)
}

type GatewayIamBindingMapOutput struct{ *pulumi.OutputState }

func (GatewayIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayIamBinding)(nil)).Elem()
}

func (o GatewayIamBindingMapOutput) ToGatewayIamBindingMapOutput() GatewayIamBindingMapOutput {
	return o
}

func (o GatewayIamBindingMapOutput) ToGatewayIamBindingMapOutputWithContext(ctx context.Context) GatewayIamBindingMapOutput {
	return o
}

func (o GatewayIamBindingMapOutput) MapIndex(k pulumi.StringInput) GatewayIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayIamBinding {
		return vs[0].(map[string]*GatewayIamBinding)[vs[1].(string)]
	}).(GatewayIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayIamBindingInput)(nil)).Elem(), &GatewayIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayIamBindingArrayInput)(nil)).Elem(), GatewayIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayIamBindingMapInput)(nil)).Elem(), GatewayIamBindingMap{})
	pulumi.RegisterOutputType(GatewayIamBindingOutput{})
	pulumi.RegisterOutputType(GatewayIamBindingArrayOutput{})
	pulumi.RegisterOutputType(GatewayIamBindingMapOutput{})
}

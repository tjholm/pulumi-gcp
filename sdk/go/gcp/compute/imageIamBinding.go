// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Compute Engine Image. Each of these resources serves a different use case:
//
// * `compute.ImageIamPolicy`: Authoritative. Sets the IAM policy for the image and replaces any existing policy already attached.
// * `compute.ImageIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the image are preserved.
// * `compute.ImageIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the image are preserved.
//
// > **Note:** `compute.ImageIamPolicy` **cannot** be used in conjunction with `compute.ImageIamBinding` and `compute.ImageIamMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.ImageIamBinding` resources **can be** used in conjunction with `compute.ImageIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_compute\_image\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/compute.imageUser",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewImageIamPolicy(ctx, "policy", &compute.ImageIamPolicyArgs{
// 			Project:    pulumi.Any(google_compute_image.Example.Project),
// 			Image:      pulumi.Any(google_compute_image.Example.Name),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/compute.imageUser",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewImageIamPolicy(ctx, "policy", &compute.ImageIamPolicyArgs{
// 			Project:    pulumi.Any(google_compute_image.Example.Project),
// 			Image:      pulumi.Any(google_compute_image.Example.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_compute\_image\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImageIamBinding(ctx, "binding", &compute.ImageIamBindingArgs{
// 			Project: pulumi.Any(google_compute_image.Example.Project),
// 			Image:   pulumi.Any(google_compute_image.Example.Name),
// 			Role:    pulumi.String("roles/compute.imageUser"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImageIamBinding(ctx, "binding", &compute.ImageIamBindingArgs{
// 			Project: pulumi.Any(google_compute_image.Example.Project),
// 			Image:   pulumi.Any(google_compute_image.Example.Name),
// 			Role:    pulumi.String("roles/compute.imageUser"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &compute.ImageIamBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_compute\_image\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImageIamMember(ctx, "member", &compute.ImageIamMemberArgs{
// 			Project: pulumi.Any(google_compute_image.Example.Project),
// 			Image:   pulumi.Any(google_compute_image.Example.Name),
// 			Role:    pulumi.String("roles/compute.imageUser"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImageIamMember(ctx, "member", &compute.ImageIamMemberArgs{
// 			Project: pulumi.Any(google_compute_image.Example.Project),
// 			Image:   pulumi.Any(google_compute_image.Example.Name),
// 			Role:    pulumi.String("roles/compute.imageUser"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Condition: &compute.ImageIamMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/global/images/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine image IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/imageIamBinding:ImageIamBinding editor "projects/{{project}}/global/images/{{image}} roles/compute.imageUser user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/imageIamBinding:ImageIamBinding editor "projects/{{project}}/global/images/{{image}} roles/compute.imageUser"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/imageIamBinding:ImageIamBinding editor projects/{{project}}/global/images/{{image}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ImageIamBinding struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition ImageIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Image   pulumi.StringOutput      `pulumi:"image"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.ImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewImageIamBinding registers a new resource with the given unique name, arguments, and options.
func NewImageIamBinding(ctx *pulumi.Context,
	name string, args *ImageIamBindingArgs, opts ...pulumi.ResourceOption) (*ImageIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Image == nil {
		return nil, errors.New("invalid value for required argument 'Image'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ImageIamBinding
	err := ctx.RegisterResource("gcp:compute/imageIamBinding:ImageIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageIamBinding gets an existing ImageIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageIamBindingState, opts ...pulumi.ResourceOption) (*ImageIamBinding, error) {
	var resource ImageIamBinding
	err := ctx.ReadResource("gcp:compute/imageIamBinding:ImageIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageIamBinding resources.
type imageIamBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *ImageIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Image   *string  `pulumi:"image"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.ImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ImageIamBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition ImageIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Image   pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.ImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ImageIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageIamBindingState)(nil)).Elem()
}

type imageIamBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *ImageIamBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Image   string   `pulumi:"image"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.ImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ImageIamBinding resource.
type ImageIamBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition ImageIamBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Image   pulumi.StringInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.ImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ImageIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageIamBindingArgs)(nil)).Elem()
}

type ImageIamBindingInput interface {
	pulumi.Input

	ToImageIamBindingOutput() ImageIamBindingOutput
	ToImageIamBindingOutputWithContext(ctx context.Context) ImageIamBindingOutput
}

func (*ImageIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageIamBinding)(nil))
}

func (i *ImageIamBinding) ToImageIamBindingOutput() ImageIamBindingOutput {
	return i.ToImageIamBindingOutputWithContext(context.Background())
}

func (i *ImageIamBinding) ToImageIamBindingOutputWithContext(ctx context.Context) ImageIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageIamBindingOutput)
}

func (i *ImageIamBinding) ToImageIamBindingPtrOutput() ImageIamBindingPtrOutput {
	return i.ToImageIamBindingPtrOutputWithContext(context.Background())
}

func (i *ImageIamBinding) ToImageIamBindingPtrOutputWithContext(ctx context.Context) ImageIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageIamBindingPtrOutput)
}

type ImageIamBindingPtrInput interface {
	pulumi.Input

	ToImageIamBindingPtrOutput() ImageIamBindingPtrOutput
	ToImageIamBindingPtrOutputWithContext(ctx context.Context) ImageIamBindingPtrOutput
}

type imageIamBindingPtrType ImageIamBindingArgs

func (*imageIamBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageIamBinding)(nil))
}

func (i *imageIamBindingPtrType) ToImageIamBindingPtrOutput() ImageIamBindingPtrOutput {
	return i.ToImageIamBindingPtrOutputWithContext(context.Background())
}

func (i *imageIamBindingPtrType) ToImageIamBindingPtrOutputWithContext(ctx context.Context) ImageIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageIamBindingPtrOutput)
}

// ImageIamBindingArrayInput is an input type that accepts ImageIamBindingArray and ImageIamBindingArrayOutput values.
// You can construct a concrete instance of `ImageIamBindingArrayInput` via:
//
//          ImageIamBindingArray{ ImageIamBindingArgs{...} }
type ImageIamBindingArrayInput interface {
	pulumi.Input

	ToImageIamBindingArrayOutput() ImageIamBindingArrayOutput
	ToImageIamBindingArrayOutputWithContext(context.Context) ImageIamBindingArrayOutput
}

type ImageIamBindingArray []ImageIamBindingInput

func (ImageIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ImageIamBinding)(nil))
}

func (i ImageIamBindingArray) ToImageIamBindingArrayOutput() ImageIamBindingArrayOutput {
	return i.ToImageIamBindingArrayOutputWithContext(context.Background())
}

func (i ImageIamBindingArray) ToImageIamBindingArrayOutputWithContext(ctx context.Context) ImageIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageIamBindingArrayOutput)
}

// ImageIamBindingMapInput is an input type that accepts ImageIamBindingMap and ImageIamBindingMapOutput values.
// You can construct a concrete instance of `ImageIamBindingMapInput` via:
//
//          ImageIamBindingMap{ "key": ImageIamBindingArgs{...} }
type ImageIamBindingMapInput interface {
	pulumi.Input

	ToImageIamBindingMapOutput() ImageIamBindingMapOutput
	ToImageIamBindingMapOutputWithContext(context.Context) ImageIamBindingMapOutput
}

type ImageIamBindingMap map[string]ImageIamBindingInput

func (ImageIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ImageIamBinding)(nil))
}

func (i ImageIamBindingMap) ToImageIamBindingMapOutput() ImageIamBindingMapOutput {
	return i.ToImageIamBindingMapOutputWithContext(context.Background())
}

func (i ImageIamBindingMap) ToImageIamBindingMapOutputWithContext(ctx context.Context) ImageIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageIamBindingMapOutput)
}

type ImageIamBindingOutput struct {
	*pulumi.OutputState
}

func (ImageIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageIamBinding)(nil))
}

func (o ImageIamBindingOutput) ToImageIamBindingOutput() ImageIamBindingOutput {
	return o
}

func (o ImageIamBindingOutput) ToImageIamBindingOutputWithContext(ctx context.Context) ImageIamBindingOutput {
	return o
}

func (o ImageIamBindingOutput) ToImageIamBindingPtrOutput() ImageIamBindingPtrOutput {
	return o.ToImageIamBindingPtrOutputWithContext(context.Background())
}

func (o ImageIamBindingOutput) ToImageIamBindingPtrOutputWithContext(ctx context.Context) ImageIamBindingPtrOutput {
	return o.ApplyT(func(v ImageIamBinding) *ImageIamBinding {
		return &v
	}).(ImageIamBindingPtrOutput)
}

type ImageIamBindingPtrOutput struct {
	*pulumi.OutputState
}

func (ImageIamBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageIamBinding)(nil))
}

func (o ImageIamBindingPtrOutput) ToImageIamBindingPtrOutput() ImageIamBindingPtrOutput {
	return o
}

func (o ImageIamBindingPtrOutput) ToImageIamBindingPtrOutputWithContext(ctx context.Context) ImageIamBindingPtrOutput {
	return o
}

type ImageIamBindingArrayOutput struct{ *pulumi.OutputState }

func (ImageIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageIamBinding)(nil))
}

func (o ImageIamBindingArrayOutput) ToImageIamBindingArrayOutput() ImageIamBindingArrayOutput {
	return o
}

func (o ImageIamBindingArrayOutput) ToImageIamBindingArrayOutputWithContext(ctx context.Context) ImageIamBindingArrayOutput {
	return o
}

func (o ImageIamBindingArrayOutput) Index(i pulumi.IntInput) ImageIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageIamBinding {
		return vs[0].([]ImageIamBinding)[vs[1].(int)]
	}).(ImageIamBindingOutput)
}

type ImageIamBindingMapOutput struct{ *pulumi.OutputState }

func (ImageIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImageIamBinding)(nil))
}

func (o ImageIamBindingMapOutput) ToImageIamBindingMapOutput() ImageIamBindingMapOutput {
	return o
}

func (o ImageIamBindingMapOutput) ToImageIamBindingMapOutputWithContext(ctx context.Context) ImageIamBindingMapOutput {
	return o
}

func (o ImageIamBindingMapOutput) MapIndex(k pulumi.StringInput) ImageIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ImageIamBinding {
		return vs[0].(map[string]ImageIamBinding)[vs[1].(string)]
	}).(ImageIamBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageIamBindingOutput{})
	pulumi.RegisterOutputType(ImageIamBindingPtrOutput{})
	pulumi.RegisterOutputType(ImageIamBindingArrayOutput{})
	pulumi.RegisterOutputType(ImageIamBindingMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Compute Engine MachineImage. Each of these resources serves a different use case:
//
// * `compute.MachineImageIamPolicy`: Authoritative. Sets the IAM policy for the machineimage and replaces any existing policy already attached.
// * `compute.MachineImageIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the machineimage are preserved.
// * `compute.MachineImageIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the machineimage are preserved.
//
// > **Note:** `compute.MachineImageIamPolicy` **cannot** be used in conjunction with `compute.MachineImageIamBinding` and `compute.MachineImageIamMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.MachineImageIamBinding` resources **can be** used in conjunction with `compute.MachineImageIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_compute\_machine\_image\_iam\_policy
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
// 					Role: "roles/compute.admin",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewMachineImageIamPolicy(ctx, "policy", &compute.MachineImageIamPolicyArgs{
// 			Project:      pulumi.Any(google_compute_machine_image.Image.Project),
// 			MachineImage: pulumi.Any(google_compute_machine_image.Image.Name),
// 			PolicyData:   pulumi.String(admin.PolicyData),
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
// 					Role: "roles/compute.admin",
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
// 		_, err = compute.NewMachineImageIamPolicy(ctx, "policy", &compute.MachineImageIamPolicyArgs{
// 			Project:      pulumi.Any(google_compute_machine_image.Image.Project),
// 			MachineImage: pulumi.Any(google_compute_machine_image.Image.Name),
// 			PolicyData:   pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_compute\_machine\_image\_iam\_binding
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
// 		_, err := compute.NewMachineImageIamBinding(ctx, "binding", &compute.MachineImageIamBindingArgs{
// 			Project:      pulumi.Any(google_compute_machine_image.Image.Project),
// 			MachineImage: pulumi.Any(google_compute_machine_image.Image.Name),
// 			Role:         pulumi.String("roles/compute.admin"),
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
// 		_, err := compute.NewMachineImageIamBinding(ctx, "binding", &compute.MachineImageIamBindingArgs{
// 			Project:      pulumi.Any(google_compute_machine_image.Image.Project),
// 			MachineImage: pulumi.Any(google_compute_machine_image.Image.Name),
// 			Role:         pulumi.String("roles/compute.admin"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &compute.MachineImageIamBindingConditionArgs{
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
// ## google\_compute\_machine\_image\_iam\_member
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
// 		_, err := compute.NewMachineImageIamMember(ctx, "member", &compute.MachineImageIamMemberArgs{
// 			Project:      pulumi.Any(google_compute_machine_image.Image.Project),
// 			MachineImage: pulumi.Any(google_compute_machine_image.Image.Name),
// 			Role:         pulumi.String("roles/compute.admin"),
// 			Member:       pulumi.String("user:jane@example.com"),
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
// 		_, err := compute.NewMachineImageIamMember(ctx, "member", &compute.MachineImageIamMemberArgs{
// 			Project:      pulumi.Any(google_compute_machine_image.Image.Project),
// 			MachineImage: pulumi.Any(google_compute_machine_image.Image.Name),
// 			Role:         pulumi.String("roles/compute.admin"),
// 			Member:       pulumi.String("user:jane@example.com"),
// 			Condition: &compute.MachineImageIamMemberConditionArgs{
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/global/machineImages/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine machineimage IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/machineImageIamPolicy:MachineImageIamPolicy editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/machineImageIamPolicy:MachineImageIamPolicy editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/machineImageIamPolicy:MachineImageIamPolicy editor projects/{{project}}/global/machineImages/{{machine_image}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type MachineImageIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringOutput `pulumi:"machineImage"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewMachineImageIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewMachineImageIamPolicy(ctx *pulumi.Context,
	name string, args *MachineImageIamPolicyArgs, opts ...pulumi.ResourceOption) (*MachineImageIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineImage == nil {
		return nil, errors.New("invalid value for required argument 'MachineImage'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource MachineImageIamPolicy
	err := ctx.RegisterResource("gcp:compute/machineImageIamPolicy:MachineImageIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineImageIamPolicy gets an existing MachineImageIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineImageIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineImageIamPolicyState, opts ...pulumi.ResourceOption) (*MachineImageIamPolicy, error) {
	var resource MachineImageIamPolicy
	err := ctx.ReadResource("gcp:compute/machineImageIamPolicy:MachineImageIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineImageIamPolicy resources.
type machineImageIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage *string `pulumi:"machineImage"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type MachineImageIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (MachineImageIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageIamPolicyState)(nil)).Elem()
}

type machineImageIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	MachineImage string `pulumi:"machineImage"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a MachineImageIamPolicy resource.
type MachineImageIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (MachineImageIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageIamPolicyArgs)(nil)).Elem()
}

type MachineImageIamPolicyInput interface {
	pulumi.Input

	ToMachineImageIamPolicyOutput() MachineImageIamPolicyOutput
	ToMachineImageIamPolicyOutputWithContext(ctx context.Context) MachineImageIamPolicyOutput
}

func (*MachineImageIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineImageIamPolicy)(nil))
}

func (i *MachineImageIamPolicy) ToMachineImageIamPolicyOutput() MachineImageIamPolicyOutput {
	return i.ToMachineImageIamPolicyOutputWithContext(context.Background())
}

func (i *MachineImageIamPolicy) ToMachineImageIamPolicyOutputWithContext(ctx context.Context) MachineImageIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageIamPolicyOutput)
}

func (i *MachineImageIamPolicy) ToMachineImageIamPolicyPtrOutput() MachineImageIamPolicyPtrOutput {
	return i.ToMachineImageIamPolicyPtrOutputWithContext(context.Background())
}

func (i *MachineImageIamPolicy) ToMachineImageIamPolicyPtrOutputWithContext(ctx context.Context) MachineImageIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageIamPolicyPtrOutput)
}

type MachineImageIamPolicyPtrInput interface {
	pulumi.Input

	ToMachineImageIamPolicyPtrOutput() MachineImageIamPolicyPtrOutput
	ToMachineImageIamPolicyPtrOutputWithContext(ctx context.Context) MachineImageIamPolicyPtrOutput
}

type machineImageIamPolicyPtrType MachineImageIamPolicyArgs

func (*machineImageIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineImageIamPolicy)(nil))
}

func (i *machineImageIamPolicyPtrType) ToMachineImageIamPolicyPtrOutput() MachineImageIamPolicyPtrOutput {
	return i.ToMachineImageIamPolicyPtrOutputWithContext(context.Background())
}

func (i *machineImageIamPolicyPtrType) ToMachineImageIamPolicyPtrOutputWithContext(ctx context.Context) MachineImageIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageIamPolicyPtrOutput)
}

// MachineImageIamPolicyArrayInput is an input type that accepts MachineImageIamPolicyArray and MachineImageIamPolicyArrayOutput values.
// You can construct a concrete instance of `MachineImageIamPolicyArrayInput` via:
//
//          MachineImageIamPolicyArray{ MachineImageIamPolicyArgs{...} }
type MachineImageIamPolicyArrayInput interface {
	pulumi.Input

	ToMachineImageIamPolicyArrayOutput() MachineImageIamPolicyArrayOutput
	ToMachineImageIamPolicyArrayOutputWithContext(context.Context) MachineImageIamPolicyArrayOutput
}

type MachineImageIamPolicyArray []MachineImageIamPolicyInput

func (MachineImageIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MachineImageIamPolicy)(nil))
}

func (i MachineImageIamPolicyArray) ToMachineImageIamPolicyArrayOutput() MachineImageIamPolicyArrayOutput {
	return i.ToMachineImageIamPolicyArrayOutputWithContext(context.Background())
}

func (i MachineImageIamPolicyArray) ToMachineImageIamPolicyArrayOutputWithContext(ctx context.Context) MachineImageIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageIamPolicyArrayOutput)
}

// MachineImageIamPolicyMapInput is an input type that accepts MachineImageIamPolicyMap and MachineImageIamPolicyMapOutput values.
// You can construct a concrete instance of `MachineImageIamPolicyMapInput` via:
//
//          MachineImageIamPolicyMap{ "key": MachineImageIamPolicyArgs{...} }
type MachineImageIamPolicyMapInput interface {
	pulumi.Input

	ToMachineImageIamPolicyMapOutput() MachineImageIamPolicyMapOutput
	ToMachineImageIamPolicyMapOutputWithContext(context.Context) MachineImageIamPolicyMapOutput
}

type MachineImageIamPolicyMap map[string]MachineImageIamPolicyInput

func (MachineImageIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MachineImageIamPolicy)(nil))
}

func (i MachineImageIamPolicyMap) ToMachineImageIamPolicyMapOutput() MachineImageIamPolicyMapOutput {
	return i.ToMachineImageIamPolicyMapOutputWithContext(context.Background())
}

func (i MachineImageIamPolicyMap) ToMachineImageIamPolicyMapOutputWithContext(ctx context.Context) MachineImageIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageIamPolicyMapOutput)
}

type MachineImageIamPolicyOutput struct {
	*pulumi.OutputState
}

func (MachineImageIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineImageIamPolicy)(nil))
}

func (o MachineImageIamPolicyOutput) ToMachineImageIamPolicyOutput() MachineImageIamPolicyOutput {
	return o
}

func (o MachineImageIamPolicyOutput) ToMachineImageIamPolicyOutputWithContext(ctx context.Context) MachineImageIamPolicyOutput {
	return o
}

func (o MachineImageIamPolicyOutput) ToMachineImageIamPolicyPtrOutput() MachineImageIamPolicyPtrOutput {
	return o.ToMachineImageIamPolicyPtrOutputWithContext(context.Background())
}

func (o MachineImageIamPolicyOutput) ToMachineImageIamPolicyPtrOutputWithContext(ctx context.Context) MachineImageIamPolicyPtrOutput {
	return o.ApplyT(func(v MachineImageIamPolicy) *MachineImageIamPolicy {
		return &v
	}).(MachineImageIamPolicyPtrOutput)
}

type MachineImageIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (MachineImageIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineImageIamPolicy)(nil))
}

func (o MachineImageIamPolicyPtrOutput) ToMachineImageIamPolicyPtrOutput() MachineImageIamPolicyPtrOutput {
	return o
}

func (o MachineImageIamPolicyPtrOutput) ToMachineImageIamPolicyPtrOutputWithContext(ctx context.Context) MachineImageIamPolicyPtrOutput {
	return o
}

type MachineImageIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (MachineImageIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineImageIamPolicy)(nil))
}

func (o MachineImageIamPolicyArrayOutput) ToMachineImageIamPolicyArrayOutput() MachineImageIamPolicyArrayOutput {
	return o
}

func (o MachineImageIamPolicyArrayOutput) ToMachineImageIamPolicyArrayOutputWithContext(ctx context.Context) MachineImageIamPolicyArrayOutput {
	return o
}

func (o MachineImageIamPolicyArrayOutput) Index(i pulumi.IntInput) MachineImageIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineImageIamPolicy {
		return vs[0].([]MachineImageIamPolicy)[vs[1].(int)]
	}).(MachineImageIamPolicyOutput)
}

type MachineImageIamPolicyMapOutput struct{ *pulumi.OutputState }

func (MachineImageIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MachineImageIamPolicy)(nil))
}

func (o MachineImageIamPolicyMapOutput) ToMachineImageIamPolicyMapOutput() MachineImageIamPolicyMapOutput {
	return o
}

func (o MachineImageIamPolicyMapOutput) ToMachineImageIamPolicyMapOutputWithContext(ctx context.Context) MachineImageIamPolicyMapOutput {
	return o
}

func (o MachineImageIamPolicyMapOutput) MapIndex(k pulumi.StringInput) MachineImageIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MachineImageIamPolicy {
		return vs[0].(map[string]MachineImageIamPolicy)[vs[1].(string)]
	}).(MachineImageIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineImageIamPolicyOutput{})
	pulumi.RegisterOutputType(MachineImageIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(MachineImageIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(MachineImageIamPolicyMapOutput{})
}

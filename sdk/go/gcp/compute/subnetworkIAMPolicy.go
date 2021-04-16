// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Compute Engine Subnetwork. Each of these resources serves a different use case:
//
// * `compute.SubnetworkIAMPolicy`: Authoritative. Sets the IAM policy for the subnetwork and replaces any existing policy already attached.
// * `compute.SubnetworkIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subnetwork are preserved.
// * `compute.SubnetworkIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subnetwork are preserved.
//
// > **Note:** `compute.SubnetworkIAMPolicy` **cannot** be used in conjunction with `compute.SubnetworkIAMBinding` and `compute.SubnetworkIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.SubnetworkIAMBinding` resources **can be** used in conjunction with `compute.SubnetworkIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_compute\_subnetwork\_iam\_policy
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
// 					Role: "roles/compute.networkUser",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetworkIAMPolicy(ctx, "policy", &compute.SubnetworkIAMPolicyArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
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
// 					Role: "roles/compute.networkUser",
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
// 		_, err = compute.NewSubnetworkIAMPolicy(ctx, "policy", &compute.SubnetworkIAMPolicyArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_compute\_subnetwork\_iam\_binding
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
// 		_, err := compute.NewSubnetworkIAMBinding(ctx, "binding", &compute.SubnetworkIAMBindingArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
// 			Role:       pulumi.String("roles/compute.networkUser"),
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
// 		_, err := compute.NewSubnetworkIAMBinding(ctx, "binding", &compute.SubnetworkIAMBindingArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
// 			Role:       pulumi.String("roles/compute.networkUser"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &compute.SubnetworkIAMBindingConditionArgs{
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
// ## google\_compute\_subnetwork\_iam\_member
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
// 		_, err := compute.NewSubnetworkIAMMember(ctx, "member", &compute.SubnetworkIAMMemberArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
// 			Role:       pulumi.String("roles/compute.networkUser"),
// 			Member:     pulumi.String("user:jane@example.com"),
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
// 		_, err := compute.NewSubnetworkIAMMember(ctx, "member", &compute.SubnetworkIAMMemberArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
// 			Role:       pulumi.String("roles/compute.networkUser"),
// 			Member:     pulumi.String("user:jane@example.com"),
// 			Condition: &compute.SubnetworkIAMMemberConditionArgs{
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/regions/{{region}}/subnetworks/{{name}} * {{project}}/{{region}}/{{name}} * {{region}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine subnetwork IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy editor "projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy editor "projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy editor projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type SubnetworkIAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
}

// NewSubnetworkIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewSubnetworkIAMPolicy(ctx *pulumi.Context,
	name string, args *SubnetworkIAMPolicyArgs, opts ...pulumi.ResourceOption) (*SubnetworkIAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Subnetwork == nil {
		return nil, errors.New("invalid value for required argument 'Subnetwork'")
	}
	var resource SubnetworkIAMPolicy
	err := ctx.RegisterResource("gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetworkIAMPolicy gets an existing SubnetworkIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetworkIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetworkIAMPolicyState, opts ...pulumi.ResourceOption) (*SubnetworkIAMPolicy, error) {
	var resource SubnetworkIAMPolicy
	err := ctx.ReadResource("gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetworkIAMPolicy resources.
type subnetworkIAMPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork *string `pulumi:"subnetwork"`
}

type SubnetworkIAMPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringPtrInput
}

func (SubnetworkIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkIAMPolicyState)(nil)).Elem()
}

type subnetworkIAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork string `pulumi:"subnetwork"`
}

// The set of arguments for constructing a SubnetworkIAMPolicy resource.
type SubnetworkIAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringInput
}

func (SubnetworkIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkIAMPolicyArgs)(nil)).Elem()
}

type SubnetworkIAMPolicyInput interface {
	pulumi.Input

	ToSubnetworkIAMPolicyOutput() SubnetworkIAMPolicyOutput
	ToSubnetworkIAMPolicyOutputWithContext(ctx context.Context) SubnetworkIAMPolicyOutput
}

func (*SubnetworkIAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetworkIAMPolicy)(nil))
}

func (i *SubnetworkIAMPolicy) ToSubnetworkIAMPolicyOutput() SubnetworkIAMPolicyOutput {
	return i.ToSubnetworkIAMPolicyOutputWithContext(context.Background())
}

func (i *SubnetworkIAMPolicy) ToSubnetworkIAMPolicyOutputWithContext(ctx context.Context) SubnetworkIAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkIAMPolicyOutput)
}

func (i *SubnetworkIAMPolicy) ToSubnetworkIAMPolicyPtrOutput() SubnetworkIAMPolicyPtrOutput {
	return i.ToSubnetworkIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *SubnetworkIAMPolicy) ToSubnetworkIAMPolicyPtrOutputWithContext(ctx context.Context) SubnetworkIAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkIAMPolicyPtrOutput)
}

type SubnetworkIAMPolicyPtrInput interface {
	pulumi.Input

	ToSubnetworkIAMPolicyPtrOutput() SubnetworkIAMPolicyPtrOutput
	ToSubnetworkIAMPolicyPtrOutputWithContext(ctx context.Context) SubnetworkIAMPolicyPtrOutput
}

type subnetworkIAMPolicyPtrType SubnetworkIAMPolicyArgs

func (*subnetworkIAMPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetworkIAMPolicy)(nil))
}

func (i *subnetworkIAMPolicyPtrType) ToSubnetworkIAMPolicyPtrOutput() SubnetworkIAMPolicyPtrOutput {
	return i.ToSubnetworkIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *subnetworkIAMPolicyPtrType) ToSubnetworkIAMPolicyPtrOutputWithContext(ctx context.Context) SubnetworkIAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkIAMPolicyPtrOutput)
}

// SubnetworkIAMPolicyArrayInput is an input type that accepts SubnetworkIAMPolicyArray and SubnetworkIAMPolicyArrayOutput values.
// You can construct a concrete instance of `SubnetworkIAMPolicyArrayInput` via:
//
//          SubnetworkIAMPolicyArray{ SubnetworkIAMPolicyArgs{...} }
type SubnetworkIAMPolicyArrayInput interface {
	pulumi.Input

	ToSubnetworkIAMPolicyArrayOutput() SubnetworkIAMPolicyArrayOutput
	ToSubnetworkIAMPolicyArrayOutputWithContext(context.Context) SubnetworkIAMPolicyArrayOutput
}

type SubnetworkIAMPolicyArray []SubnetworkIAMPolicyInput

func (SubnetworkIAMPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SubnetworkIAMPolicy)(nil))
}

func (i SubnetworkIAMPolicyArray) ToSubnetworkIAMPolicyArrayOutput() SubnetworkIAMPolicyArrayOutput {
	return i.ToSubnetworkIAMPolicyArrayOutputWithContext(context.Background())
}

func (i SubnetworkIAMPolicyArray) ToSubnetworkIAMPolicyArrayOutputWithContext(ctx context.Context) SubnetworkIAMPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkIAMPolicyArrayOutput)
}

// SubnetworkIAMPolicyMapInput is an input type that accepts SubnetworkIAMPolicyMap and SubnetworkIAMPolicyMapOutput values.
// You can construct a concrete instance of `SubnetworkIAMPolicyMapInput` via:
//
//          SubnetworkIAMPolicyMap{ "key": SubnetworkIAMPolicyArgs{...} }
type SubnetworkIAMPolicyMapInput interface {
	pulumi.Input

	ToSubnetworkIAMPolicyMapOutput() SubnetworkIAMPolicyMapOutput
	ToSubnetworkIAMPolicyMapOutputWithContext(context.Context) SubnetworkIAMPolicyMapOutput
}

type SubnetworkIAMPolicyMap map[string]SubnetworkIAMPolicyInput

func (SubnetworkIAMPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SubnetworkIAMPolicy)(nil))
}

func (i SubnetworkIAMPolicyMap) ToSubnetworkIAMPolicyMapOutput() SubnetworkIAMPolicyMapOutput {
	return i.ToSubnetworkIAMPolicyMapOutputWithContext(context.Background())
}

func (i SubnetworkIAMPolicyMap) ToSubnetworkIAMPolicyMapOutputWithContext(ctx context.Context) SubnetworkIAMPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkIAMPolicyMapOutput)
}

type SubnetworkIAMPolicyOutput struct {
	*pulumi.OutputState
}

func (SubnetworkIAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetworkIAMPolicy)(nil))
}

func (o SubnetworkIAMPolicyOutput) ToSubnetworkIAMPolicyOutput() SubnetworkIAMPolicyOutput {
	return o
}

func (o SubnetworkIAMPolicyOutput) ToSubnetworkIAMPolicyOutputWithContext(ctx context.Context) SubnetworkIAMPolicyOutput {
	return o
}

func (o SubnetworkIAMPolicyOutput) ToSubnetworkIAMPolicyPtrOutput() SubnetworkIAMPolicyPtrOutput {
	return o.ToSubnetworkIAMPolicyPtrOutputWithContext(context.Background())
}

func (o SubnetworkIAMPolicyOutput) ToSubnetworkIAMPolicyPtrOutputWithContext(ctx context.Context) SubnetworkIAMPolicyPtrOutput {
	return o.ApplyT(func(v SubnetworkIAMPolicy) *SubnetworkIAMPolicy {
		return &v
	}).(SubnetworkIAMPolicyPtrOutput)
}

type SubnetworkIAMPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (SubnetworkIAMPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetworkIAMPolicy)(nil))
}

func (o SubnetworkIAMPolicyPtrOutput) ToSubnetworkIAMPolicyPtrOutput() SubnetworkIAMPolicyPtrOutput {
	return o
}

func (o SubnetworkIAMPolicyPtrOutput) ToSubnetworkIAMPolicyPtrOutputWithContext(ctx context.Context) SubnetworkIAMPolicyPtrOutput {
	return o
}

type SubnetworkIAMPolicyArrayOutput struct{ *pulumi.OutputState }

func (SubnetworkIAMPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetworkIAMPolicy)(nil))
}

func (o SubnetworkIAMPolicyArrayOutput) ToSubnetworkIAMPolicyArrayOutput() SubnetworkIAMPolicyArrayOutput {
	return o
}

func (o SubnetworkIAMPolicyArrayOutput) ToSubnetworkIAMPolicyArrayOutputWithContext(ctx context.Context) SubnetworkIAMPolicyArrayOutput {
	return o
}

func (o SubnetworkIAMPolicyArrayOutput) Index(i pulumi.IntInput) SubnetworkIAMPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetworkIAMPolicy {
		return vs[0].([]SubnetworkIAMPolicy)[vs[1].(int)]
	}).(SubnetworkIAMPolicyOutput)
}

type SubnetworkIAMPolicyMapOutput struct{ *pulumi.OutputState }

func (SubnetworkIAMPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SubnetworkIAMPolicy)(nil))
}

func (o SubnetworkIAMPolicyMapOutput) ToSubnetworkIAMPolicyMapOutput() SubnetworkIAMPolicyMapOutput {
	return o
}

func (o SubnetworkIAMPolicyMapOutput) ToSubnetworkIAMPolicyMapOutputWithContext(ctx context.Context) SubnetworkIAMPolicyMapOutput {
	return o
}

func (o SubnetworkIAMPolicyMapOutput) MapIndex(k pulumi.StringInput) SubnetworkIAMPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SubnetworkIAMPolicy {
		return vs[0].(map[string]SubnetworkIAMPolicy)[vs[1].(string)]
	}).(SubnetworkIAMPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(SubnetworkIAMPolicyOutput{})
	pulumi.RegisterOutputType(SubnetworkIAMPolicyPtrOutput{})
	pulumi.RegisterOutputType(SubnetworkIAMPolicyArrayOutput{})
	pulumi.RegisterOutputType(SubnetworkIAMPolicyMapOutput{})
}

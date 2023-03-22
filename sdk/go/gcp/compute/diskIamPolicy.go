// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Compute Engine Disk. Each of these resources serves a different use case:
//
// * `compute.DiskIamPolicy`: Authoritative. Sets the IAM policy for the disk and replaces any existing policy already attached.
// * `compute.DiskIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the disk are preserved.
// * `compute.DiskIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the disk are preserved.
//
// > **Note:** `compute.DiskIamPolicy` **cannot** be used in conjunction with `compute.DiskIamBinding` and `compute.DiskIamMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.DiskIamBinding` resources **can be** used in conjunction with `compute.DiskIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_compute\_disk\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
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
//			_, err = compute.NewDiskIamPolicy(ctx, "policy", &compute.DiskIamPolicyArgs{
//				Project:    pulumi.Any(google_compute_disk.Default.Project),
//				Zone:       pulumi.Any(google_compute_disk.Default.Zone),
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
// ## google\_compute\_disk\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewDiskIamBinding(ctx, "binding", &compute.DiskIamBindingArgs{
//				Project: pulumi.Any(google_compute_disk.Default.Project),
//				Zone:    pulumi.Any(google_compute_disk.Default.Zone),
//				Role:    pulumi.String("roles/viewer"),
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
// ## google\_compute\_disk\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewDiskIamMember(ctx, "member", &compute.DiskIamMemberArgs{
//				Project: pulumi.Any(google_compute_disk.Default.Project),
//				Zone:    pulumi.Any(google_compute_disk.Default.Zone),
//				Role:    pulumi.String("roles/viewer"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/zones/{{zone}}/disks/{{name}} * {{project}}/{{zone}}/{{name}} * {{zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine disk IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/diskIamPolicy:DiskIamPolicy editor "projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/diskIamPolicy:DiskIamPolicy editor "projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/diskIamPolicy:DiskIamPolicy editor projects/{{project}}/zones/{{zone}}/disks/{{disk}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DiskIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A reference to the zone where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewDiskIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDiskIamPolicy(ctx *pulumi.Context,
	name string, args *DiskIamPolicyArgs, opts ...pulumi.ResourceOption) (*DiskIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource DiskIamPolicy
	err := ctx.RegisterResource("gcp:compute/diskIamPolicy:DiskIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskIamPolicy gets an existing DiskIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskIamPolicyState, opts ...pulumi.ResourceOption) (*DiskIamPolicy, error) {
	var resource DiskIamPolicy
	err := ctx.ReadResource("gcp:compute/diskIamPolicy:DiskIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskIamPolicy resources.
type diskIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the zone where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone *string `pulumi:"zone"`
}

type DiskIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the zone where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringPtrInput
}

func (DiskIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskIamPolicyState)(nil)).Elem()
}

type diskIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the zone where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a DiskIamPolicy resource.
type DiskIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the zone where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringPtrInput
}

func (DiskIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskIamPolicyArgs)(nil)).Elem()
}

type DiskIamPolicyInput interface {
	pulumi.Input

	ToDiskIamPolicyOutput() DiskIamPolicyOutput
	ToDiskIamPolicyOutputWithContext(ctx context.Context) DiskIamPolicyOutput
}

func (*DiskIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskIamPolicy)(nil)).Elem()
}

func (i *DiskIamPolicy) ToDiskIamPolicyOutput() DiskIamPolicyOutput {
	return i.ToDiskIamPolicyOutputWithContext(context.Background())
}

func (i *DiskIamPolicy) ToDiskIamPolicyOutputWithContext(ctx context.Context) DiskIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskIamPolicyOutput)
}

// DiskIamPolicyArrayInput is an input type that accepts DiskIamPolicyArray and DiskIamPolicyArrayOutput values.
// You can construct a concrete instance of `DiskIamPolicyArrayInput` via:
//
//	DiskIamPolicyArray{ DiskIamPolicyArgs{...} }
type DiskIamPolicyArrayInput interface {
	pulumi.Input

	ToDiskIamPolicyArrayOutput() DiskIamPolicyArrayOutput
	ToDiskIamPolicyArrayOutputWithContext(context.Context) DiskIamPolicyArrayOutput
}

type DiskIamPolicyArray []DiskIamPolicyInput

func (DiskIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DiskIamPolicy)(nil)).Elem()
}

func (i DiskIamPolicyArray) ToDiskIamPolicyArrayOutput() DiskIamPolicyArrayOutput {
	return i.ToDiskIamPolicyArrayOutputWithContext(context.Background())
}

func (i DiskIamPolicyArray) ToDiskIamPolicyArrayOutputWithContext(ctx context.Context) DiskIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskIamPolicyArrayOutput)
}

// DiskIamPolicyMapInput is an input type that accepts DiskIamPolicyMap and DiskIamPolicyMapOutput values.
// You can construct a concrete instance of `DiskIamPolicyMapInput` via:
//
//	DiskIamPolicyMap{ "key": DiskIamPolicyArgs{...} }
type DiskIamPolicyMapInput interface {
	pulumi.Input

	ToDiskIamPolicyMapOutput() DiskIamPolicyMapOutput
	ToDiskIamPolicyMapOutputWithContext(context.Context) DiskIamPolicyMapOutput
}

type DiskIamPolicyMap map[string]DiskIamPolicyInput

func (DiskIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DiskIamPolicy)(nil)).Elem()
}

func (i DiskIamPolicyMap) ToDiskIamPolicyMapOutput() DiskIamPolicyMapOutput {
	return i.ToDiskIamPolicyMapOutputWithContext(context.Background())
}

func (i DiskIamPolicyMap) ToDiskIamPolicyMapOutputWithContext(ctx context.Context) DiskIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskIamPolicyMapOutput)
}

type DiskIamPolicyOutput struct{ *pulumi.OutputState }

func (DiskIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskIamPolicy)(nil)).Elem()
}

func (o DiskIamPolicyOutput) ToDiskIamPolicyOutput() DiskIamPolicyOutput {
	return o
}

func (o DiskIamPolicyOutput) ToDiskIamPolicyOutputWithContext(ctx context.Context) DiskIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o DiskIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o DiskIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o DiskIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o DiskIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A reference to the zone where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
// zone is specified, it is taken from the provider configuration.
func (o DiskIamPolicyOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskIamPolicy) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type DiskIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (DiskIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DiskIamPolicy)(nil)).Elem()
}

func (o DiskIamPolicyArrayOutput) ToDiskIamPolicyArrayOutput() DiskIamPolicyArrayOutput {
	return o
}

func (o DiskIamPolicyArrayOutput) ToDiskIamPolicyArrayOutputWithContext(ctx context.Context) DiskIamPolicyArrayOutput {
	return o
}

func (o DiskIamPolicyArrayOutput) Index(i pulumi.IntInput) DiskIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DiskIamPolicy {
		return vs[0].([]*DiskIamPolicy)[vs[1].(int)]
	}).(DiskIamPolicyOutput)
}

type DiskIamPolicyMapOutput struct{ *pulumi.OutputState }

func (DiskIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DiskIamPolicy)(nil)).Elem()
}

func (o DiskIamPolicyMapOutput) ToDiskIamPolicyMapOutput() DiskIamPolicyMapOutput {
	return o
}

func (o DiskIamPolicyMapOutput) ToDiskIamPolicyMapOutputWithContext(ctx context.Context) DiskIamPolicyMapOutput {
	return o
}

func (o DiskIamPolicyMapOutput) MapIndex(k pulumi.StringInput) DiskIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DiskIamPolicy {
		return vs[0].(map[string]*DiskIamPolicy)[vs[1].(string)]
	}).(DiskIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiskIamPolicyInput)(nil)).Elem(), &DiskIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskIamPolicyArrayInput)(nil)).Elem(), DiskIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskIamPolicyMapInput)(nil)).Elem(), DiskIamPolicyMap{})
	pulumi.RegisterOutputType(DiskIamPolicyOutput{})
	pulumi.RegisterOutputType(DiskIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(DiskIamPolicyMapOutput{})
}

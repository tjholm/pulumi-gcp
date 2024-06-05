// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigquerydatapolicy

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for BigQuery Data Policy DataPolicy. Each of these resources serves a different use case:
//
// * `bigquerydatapolicy.DataPolicyIamPolicy`: Authoritative. Sets the IAM policy for the datapolicy and replaces any existing policy already attached.
// * `bigquerydatapolicy.DataPolicyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the datapolicy are preserved.
// * `bigquerydatapolicy.DataPolicyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the datapolicy are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `bigquerydatapolicy.DataPolicyIamPolicy`: Retrieves the IAM policy for the datapolicy
//
// > **Note:** `bigquerydatapolicy.DataPolicyIamPolicy` **cannot** be used in conjunction with `bigquerydatapolicy.DataPolicyIamBinding` and `bigquerydatapolicy.DataPolicyIamMember` or they will fight over what your policy should be.
//
// > **Note:** `bigquerydatapolicy.DataPolicyIamBinding` resources **can be** used in conjunction with `bigquerydatapolicy.DataPolicyIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## bigquerydatapolicy.DataPolicyIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
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
//			_, err = bigquerydatapolicy.NewDataPolicyIamPolicy(ctx, "policy", &bigquerydatapolicy.DataPolicyIamPolicyArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				PolicyData:   pulumi.String(admin.PolicyData),
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
// ## bigquerydatapolicy.DataPolicyIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigquerydatapolicy.NewDataPolicyIamBinding(ctx, "binding", &bigquerydatapolicy.DataPolicyIamBindingArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				Role:         pulumi.String("roles/viewer"),
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
// ## bigquerydatapolicy.DataPolicyIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigquerydatapolicy.NewDataPolicyIamMember(ctx, "member", &bigquerydatapolicy.DataPolicyIamMemberArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				Role:         pulumi.String("roles/viewer"),
//				Member:       pulumi.String("user:jane@example.com"),
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
// ## bigquerydatapolicy.DataPolicyIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
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
//			_, err = bigquerydatapolicy.NewDataPolicyIamPolicy(ctx, "policy", &bigquerydatapolicy.DataPolicyIamPolicyArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				PolicyData:   pulumi.String(admin.PolicyData),
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
// ## bigquerydatapolicy.DataPolicyIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigquerydatapolicy.NewDataPolicyIamBinding(ctx, "binding", &bigquerydatapolicy.DataPolicyIamBindingArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				Role:         pulumi.String("roles/viewer"),
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
// ## bigquerydatapolicy.DataPolicyIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquerydatapolicy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigquerydatapolicy.NewDataPolicyIamMember(ctx, "member", &bigquerydatapolicy.DataPolicyIamMemberArgs{
//				Project:      pulumi.Any(dataPolicy.Project),
//				Location:     pulumi.Any(dataPolicy.Location),
//				DataPolicyId: pulumi.Any(dataPolicy.DataPolicyId),
//				Role:         pulumi.String("roles/viewer"),
//				Member:       pulumi.String("user:jane@example.com"),
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
// * projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}
//
// * {{project}}/{{location}}/{{data_policy_id}}
//
// * {{location}}/{{data_policy_id}}
//
// * {{data_policy_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// BigQuery Data Policy datapolicy IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamPolicy:DataPolicyIamPolicy editor "projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamPolicy:DataPolicyIamPolicy editor "projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamPolicy:DataPolicyIamPolicy editor projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DataPolicyIamPolicy struct {
	pulumi.CustomResourceState

	DataPolicyId pulumi.StringOutput `pulumi:"dataPolicyId"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the location of the data policy.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringOutput `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewDataPolicyIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDataPolicyIamPolicy(ctx *pulumi.Context,
	name string, args *DataPolicyIamPolicyArgs, opts ...pulumi.ResourceOption) (*DataPolicyIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'DataPolicyId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataPolicyIamPolicy
	err := ctx.RegisterResource("gcp:bigquerydatapolicy/dataPolicyIamPolicy:DataPolicyIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataPolicyIamPolicy gets an existing DataPolicyIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataPolicyIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataPolicyIamPolicyState, opts ...pulumi.ResourceOption) (*DataPolicyIamPolicy, error) {
	var resource DataPolicyIamPolicy
	err := ctx.ReadResource("gcp:bigquerydatapolicy/dataPolicyIamPolicy:DataPolicyIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataPolicyIamPolicy resources.
type dataPolicyIamPolicyState struct {
	DataPolicyId *string `pulumi:"dataPolicyId"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the location of the data policy.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type DataPolicyIamPolicyState struct {
	DataPolicyId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the location of the data policy.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DataPolicyIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataPolicyIamPolicyState)(nil)).Elem()
}

type dataPolicyIamPolicyArgs struct {
	DataPolicyId string `pulumi:"dataPolicyId"`
	// The name of the location of the data policy.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a DataPolicyIamPolicy resource.
type DataPolicyIamPolicyArgs struct {
	DataPolicyId pulumi.StringInput
	// The name of the location of the data policy.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DataPolicyIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataPolicyIamPolicyArgs)(nil)).Elem()
}

type DataPolicyIamPolicyInput interface {
	pulumi.Input

	ToDataPolicyIamPolicyOutput() DataPolicyIamPolicyOutput
	ToDataPolicyIamPolicyOutputWithContext(ctx context.Context) DataPolicyIamPolicyOutput
}

func (*DataPolicyIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPolicyIamPolicy)(nil)).Elem()
}

func (i *DataPolicyIamPolicy) ToDataPolicyIamPolicyOutput() DataPolicyIamPolicyOutput {
	return i.ToDataPolicyIamPolicyOutputWithContext(context.Background())
}

func (i *DataPolicyIamPolicy) ToDataPolicyIamPolicyOutputWithContext(ctx context.Context) DataPolicyIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPolicyIamPolicyOutput)
}

// DataPolicyIamPolicyArrayInput is an input type that accepts DataPolicyIamPolicyArray and DataPolicyIamPolicyArrayOutput values.
// You can construct a concrete instance of `DataPolicyIamPolicyArrayInput` via:
//
//	DataPolicyIamPolicyArray{ DataPolicyIamPolicyArgs{...} }
type DataPolicyIamPolicyArrayInput interface {
	pulumi.Input

	ToDataPolicyIamPolicyArrayOutput() DataPolicyIamPolicyArrayOutput
	ToDataPolicyIamPolicyArrayOutputWithContext(context.Context) DataPolicyIamPolicyArrayOutput
}

type DataPolicyIamPolicyArray []DataPolicyIamPolicyInput

func (DataPolicyIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataPolicyIamPolicy)(nil)).Elem()
}

func (i DataPolicyIamPolicyArray) ToDataPolicyIamPolicyArrayOutput() DataPolicyIamPolicyArrayOutput {
	return i.ToDataPolicyIamPolicyArrayOutputWithContext(context.Background())
}

func (i DataPolicyIamPolicyArray) ToDataPolicyIamPolicyArrayOutputWithContext(ctx context.Context) DataPolicyIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPolicyIamPolicyArrayOutput)
}

// DataPolicyIamPolicyMapInput is an input type that accepts DataPolicyIamPolicyMap and DataPolicyIamPolicyMapOutput values.
// You can construct a concrete instance of `DataPolicyIamPolicyMapInput` via:
//
//	DataPolicyIamPolicyMap{ "key": DataPolicyIamPolicyArgs{...} }
type DataPolicyIamPolicyMapInput interface {
	pulumi.Input

	ToDataPolicyIamPolicyMapOutput() DataPolicyIamPolicyMapOutput
	ToDataPolicyIamPolicyMapOutputWithContext(context.Context) DataPolicyIamPolicyMapOutput
}

type DataPolicyIamPolicyMap map[string]DataPolicyIamPolicyInput

func (DataPolicyIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataPolicyIamPolicy)(nil)).Elem()
}

func (i DataPolicyIamPolicyMap) ToDataPolicyIamPolicyMapOutput() DataPolicyIamPolicyMapOutput {
	return i.ToDataPolicyIamPolicyMapOutputWithContext(context.Background())
}

func (i DataPolicyIamPolicyMap) ToDataPolicyIamPolicyMapOutputWithContext(ctx context.Context) DataPolicyIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPolicyIamPolicyMapOutput)
}

type DataPolicyIamPolicyOutput struct{ *pulumi.OutputState }

func (DataPolicyIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPolicyIamPolicy)(nil)).Elem()
}

func (o DataPolicyIamPolicyOutput) ToDataPolicyIamPolicyOutput() DataPolicyIamPolicyOutput {
	return o
}

func (o DataPolicyIamPolicyOutput) ToDataPolicyIamPolicyOutputWithContext(ctx context.Context) DataPolicyIamPolicyOutput {
	return o
}

func (o DataPolicyIamPolicyOutput) DataPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataPolicyIamPolicy) pulumi.StringOutput { return v.DataPolicyId }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o DataPolicyIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataPolicyIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the location of the data policy.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o DataPolicyIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataPolicyIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o DataPolicyIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DataPolicyIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o DataPolicyIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataPolicyIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type DataPolicyIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (DataPolicyIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataPolicyIamPolicy)(nil)).Elem()
}

func (o DataPolicyIamPolicyArrayOutput) ToDataPolicyIamPolicyArrayOutput() DataPolicyIamPolicyArrayOutput {
	return o
}

func (o DataPolicyIamPolicyArrayOutput) ToDataPolicyIamPolicyArrayOutputWithContext(ctx context.Context) DataPolicyIamPolicyArrayOutput {
	return o
}

func (o DataPolicyIamPolicyArrayOutput) Index(i pulumi.IntInput) DataPolicyIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataPolicyIamPolicy {
		return vs[0].([]*DataPolicyIamPolicy)[vs[1].(int)]
	}).(DataPolicyIamPolicyOutput)
}

type DataPolicyIamPolicyMapOutput struct{ *pulumi.OutputState }

func (DataPolicyIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataPolicyIamPolicy)(nil)).Elem()
}

func (o DataPolicyIamPolicyMapOutput) ToDataPolicyIamPolicyMapOutput() DataPolicyIamPolicyMapOutput {
	return o
}

func (o DataPolicyIamPolicyMapOutput) ToDataPolicyIamPolicyMapOutputWithContext(ctx context.Context) DataPolicyIamPolicyMapOutput {
	return o
}

func (o DataPolicyIamPolicyMapOutput) MapIndex(k pulumi.StringInput) DataPolicyIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataPolicyIamPolicy {
		return vs[0].(map[string]*DataPolicyIamPolicy)[vs[1].(string)]
	}).(DataPolicyIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataPolicyIamPolicyInput)(nil)).Elem(), &DataPolicyIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPolicyIamPolicyArrayInput)(nil)).Elem(), DataPolicyIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPolicyIamPolicyMapInput)(nil)).Elem(), DataPolicyIamPolicyMap{})
	pulumi.RegisterOutputType(DataPolicyIamPolicyOutput{})
	pulumi.RegisterOutputType(DataPolicyIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(DataPolicyIamPolicyMapOutput{})
}

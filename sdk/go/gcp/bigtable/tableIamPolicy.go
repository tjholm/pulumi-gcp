// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage IAM policies on bigtable tables. Each of these resources serves a different use case:
//
// * `bigtable.TableIamPolicy`: Authoritative. Sets the IAM policy for the tables and replaces any existing policy already attached.
// * `bigtable.TableIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
// * `bigtable.TableIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.
//
// > **Note:** `bigtable.TableIamPolicy` **cannot** be used in conjunction with `bigtable.TableIamBinding` and `bigtable.TableIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the table as `bigtable.TableIamPolicy` replaces the entire policy.
//
// > **Note:** `bigtable.TableIamBinding` resources **can be** used in conjunction with `bigtable.TableIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_bigtable\_table\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigtable"
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
//						Role: "roles/bigtable.user",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = bigtable.NewTableIamPolicy(ctx, "editor", &bigtable.TableIamPolicyArgs{
//				Project:    pulumi.String("your-project"),
//				Instance:   pulumi.String("your-bigtable-instance"),
//				Table:      pulumi.String("your-bigtable-table"),
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
// ## google\_bigtable\_table\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigtable"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigtable.NewTableIamBinding(ctx, "editor", &bigtable.TableIamBindingArgs{
//				Instance: pulumi.String("your-bigtable-instance"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Role:  pulumi.String("roles/bigtable.user"),
//				Table: pulumi.String("your-bigtable-table"),
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
// ## google\_bigtable\_table\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigtable"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigtable.NewTableIamMember(ctx, "editor", &bigtable.TableIamMemberArgs{
//				Instance: pulumi.String("your-bigtable-instance"),
//				Member:   pulumi.String("user:jane@example.com"),
//				Role:     pulumi.String("roles/bigtable.user"),
//				Table:    pulumi.String("your-bigtable-table"),
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
// Table IAM resources can be imported using the project, table name, role and/or member.
//
// ```sh
//
//	$ pulumi import gcp:bigtable/tableIamPolicy:TableIamPolicy editor "projects/{project}/tables/{table}"
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigtable/tableIamPolicy:TableIamPolicy editor "projects/{project}/tables/{table} roles/editor"
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigtable/tableIamPolicy:TableIamPolicy editor "projects/{project}/tables/{table} roles/editor user:jane@example.com"
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TableIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the tables's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	Project    pulumi.StringOutput `pulumi:"project"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringOutput `pulumi:"table"`
}

// NewTableIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewTableIamPolicy(ctx *pulumi.Context,
	name string, args *TableIamPolicyArgs, opts ...pulumi.ResourceOption) (*TableIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Table == nil {
		return nil, errors.New("invalid value for required argument 'Table'")
	}
	var resource TableIamPolicy
	err := ctx.RegisterResource("gcp:bigtable/tableIamPolicy:TableIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTableIamPolicy gets an existing TableIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableIamPolicyState, opts ...pulumi.ResourceOption) (*TableIamPolicy, error) {
	var resource TableIamPolicy
	err := ctx.ReadResource("gcp:bigtable/tableIamPolicy:TableIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TableIamPolicy resources.
type tableIamPolicyState struct {
	// (Computed) The etag of the tables's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name or relative resource id of the instance that owns the table.
	Instance *string `pulumi:"instance"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	Project    *string `pulumi:"project"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table *string `pulumi:"table"`
}

type TableIamPolicyState struct {
	// (Computed) The etag of the tables's IAM policy.
	Etag pulumi.StringPtrInput
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringPtrInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	Project    pulumi.StringPtrInput
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringPtrInput
}

func (TableIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableIamPolicyState)(nil)).Elem()
}

type tableIamPolicyArgs struct {
	// The name or relative resource id of the instance that owns the table.
	Instance string `pulumi:"instance"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData string  `pulumi:"policyData"`
	Project    *string `pulumi:"project"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table string `pulumi:"table"`
}

// The set of arguments for constructing a TableIamPolicy resource.
type TableIamPolicyArgs struct {
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	Project    pulumi.StringPtrInput
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringInput
}

func (TableIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableIamPolicyArgs)(nil)).Elem()
}

type TableIamPolicyInput interface {
	pulumi.Input

	ToTableIamPolicyOutput() TableIamPolicyOutput
	ToTableIamPolicyOutputWithContext(ctx context.Context) TableIamPolicyOutput
}

func (*TableIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TableIamPolicy)(nil)).Elem()
}

func (i *TableIamPolicy) ToTableIamPolicyOutput() TableIamPolicyOutput {
	return i.ToTableIamPolicyOutputWithContext(context.Background())
}

func (i *TableIamPolicy) ToTableIamPolicyOutputWithContext(ctx context.Context) TableIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableIamPolicyOutput)
}

// TableIamPolicyArrayInput is an input type that accepts TableIamPolicyArray and TableIamPolicyArrayOutput values.
// You can construct a concrete instance of `TableIamPolicyArrayInput` via:
//
//	TableIamPolicyArray{ TableIamPolicyArgs{...} }
type TableIamPolicyArrayInput interface {
	pulumi.Input

	ToTableIamPolicyArrayOutput() TableIamPolicyArrayOutput
	ToTableIamPolicyArrayOutputWithContext(context.Context) TableIamPolicyArrayOutput
}

type TableIamPolicyArray []TableIamPolicyInput

func (TableIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TableIamPolicy)(nil)).Elem()
}

func (i TableIamPolicyArray) ToTableIamPolicyArrayOutput() TableIamPolicyArrayOutput {
	return i.ToTableIamPolicyArrayOutputWithContext(context.Background())
}

func (i TableIamPolicyArray) ToTableIamPolicyArrayOutputWithContext(ctx context.Context) TableIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableIamPolicyArrayOutput)
}

// TableIamPolicyMapInput is an input type that accepts TableIamPolicyMap and TableIamPolicyMapOutput values.
// You can construct a concrete instance of `TableIamPolicyMapInput` via:
//
//	TableIamPolicyMap{ "key": TableIamPolicyArgs{...} }
type TableIamPolicyMapInput interface {
	pulumi.Input

	ToTableIamPolicyMapOutput() TableIamPolicyMapOutput
	ToTableIamPolicyMapOutputWithContext(context.Context) TableIamPolicyMapOutput
}

type TableIamPolicyMap map[string]TableIamPolicyInput

func (TableIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TableIamPolicy)(nil)).Elem()
}

func (i TableIamPolicyMap) ToTableIamPolicyMapOutput() TableIamPolicyMapOutput {
	return i.ToTableIamPolicyMapOutputWithContext(context.Background())
}

func (i TableIamPolicyMap) ToTableIamPolicyMapOutputWithContext(ctx context.Context) TableIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableIamPolicyMapOutput)
}

type TableIamPolicyOutput struct{ *pulumi.OutputState }

func (TableIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableIamPolicy)(nil)).Elem()
}

func (o TableIamPolicyOutput) ToTableIamPolicyOutput() TableIamPolicyOutput {
	return o
}

func (o TableIamPolicyOutput) ToTableIamPolicyOutputWithContext(ctx context.Context) TableIamPolicyOutput {
	return o
}

// (Computed) The etag of the tables's IAM policy.
func (o TableIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TableIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name or relative resource id of the instance that owns the table.
func (o TableIamPolicyOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *TableIamPolicy) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// The policy data generated by a `organizations.getIAMPolicy` data source.
func (o TableIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *TableIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o TableIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TableIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name or relative resource id of the table to manage IAM policies for.
func (o TableIamPolicyOutput) Table() pulumi.StringOutput {
	return o.ApplyT(func(v *TableIamPolicy) pulumi.StringOutput { return v.Table }).(pulumi.StringOutput)
}

type TableIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (TableIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TableIamPolicy)(nil)).Elem()
}

func (o TableIamPolicyArrayOutput) ToTableIamPolicyArrayOutput() TableIamPolicyArrayOutput {
	return o
}

func (o TableIamPolicyArrayOutput) ToTableIamPolicyArrayOutputWithContext(ctx context.Context) TableIamPolicyArrayOutput {
	return o
}

func (o TableIamPolicyArrayOutput) Index(i pulumi.IntInput) TableIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TableIamPolicy {
		return vs[0].([]*TableIamPolicy)[vs[1].(int)]
	}).(TableIamPolicyOutput)
}

type TableIamPolicyMapOutput struct{ *pulumi.OutputState }

func (TableIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TableIamPolicy)(nil)).Elem()
}

func (o TableIamPolicyMapOutput) ToTableIamPolicyMapOutput() TableIamPolicyMapOutput {
	return o
}

func (o TableIamPolicyMapOutput) ToTableIamPolicyMapOutputWithContext(ctx context.Context) TableIamPolicyMapOutput {
	return o
}

func (o TableIamPolicyMapOutput) MapIndex(k pulumi.StringInput) TableIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TableIamPolicy {
		return vs[0].(map[string]*TableIamPolicy)[vs[1].(string)]
	}).(TableIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableIamPolicyInput)(nil)).Elem(), &TableIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableIamPolicyArrayInput)(nil)).Elem(), TableIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableIamPolicyMapInput)(nil)).Elem(), TableIamPolicyMap{})
	pulumi.RegisterOutputType(TableIamPolicyOutput{})
	pulumi.RegisterOutputType(TableIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(TableIamPolicyMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Google Cloud Bigtable table inside an instance. For more information see
// [the official documentation](https://cloud.google.com/bigtable/) and
// [API](https://cloud.google.com/bigtable/docs/go/reference).
//
// ## Example Usage
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
//			instance, err := bigtable.NewInstance(ctx, "instance", &bigtable.InstanceArgs{
//				Clusters: bigtable.InstanceClusterArray{
//					&bigtable.InstanceClusterArgs{
//						ClusterId:   pulumi.String("tf-instance-cluster"),
//						Zone:        pulumi.String("us-central1-b"),
//						NumNodes:    pulumi.Int(3),
//						StorageType: pulumi.String("HDD"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigtable.NewTable(ctx, "table", &bigtable.TableArgs{
//				InstanceName: instance.Name,
//				SplitKeys: pulumi.StringArray{
//					pulumi.String("a"),
//					pulumi.String("b"),
//					pulumi.String("c"),
//				},
//				ColumnFamilies: bigtable.TableColumnFamilyArray{
//					&bigtable.TableColumnFamilyArgs{
//						Family: pulumi.String("family-first"),
//					},
//					&bigtable.TableColumnFamilyArgs{
//						Family: pulumi.String("family-second"),
//					},
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
// ## Import
//
// # Bigtable Tables can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:bigtable/table:Table default projects/{{project}}/instances/{{instance_name}}/tables/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigtable/table:Table default {{project}}/{{instance_name}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigtable/table:Table default {{instance_name}}/{{name}}
//
// ```
//
//	The following fields can't be read and will show diffs if set in config when imported- `split_keys`
type Table struct {
	pulumi.CustomResourceState

	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies TableColumnFamilyArrayOutput `pulumi:"columnFamilies"`
	// A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
	DeletionProtection pulumi.StringOutput `pulumi:"deletionProtection"`
	// The name of the Bigtable instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A list of predefined keys to split the table on.
	// !> **Warning:** Modifying the `splitKeys` of an existing table will cause the provider
	// to delete/recreate the entire `bigtable.Table` resource.
	SplitKeys pulumi.StringArrayOutput `pulumi:"splitKeys"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	var resource Table
	err := ctx.RegisterResource("gcp:bigtable/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("gcp:bigtable/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies []TableColumnFamily `pulumi:"columnFamilies"`
	// A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
	DeletionProtection *string `pulumi:"deletionProtection"`
	// The name of the Bigtable instance.
	InstanceName *string `pulumi:"instanceName"`
	// The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A list of predefined keys to split the table on.
	// !> **Warning:** Modifying the `splitKeys` of an existing table will cause the provider
	// to delete/recreate the entire `bigtable.Table` resource.
	SplitKeys []string `pulumi:"splitKeys"`
}

type TableState struct {
	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies TableColumnFamilyArrayInput
	// A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
	DeletionProtection pulumi.StringPtrInput
	// The name of the Bigtable instance.
	InstanceName pulumi.StringPtrInput
	// The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A list of predefined keys to split the table on.
	// !> **Warning:** Modifying the `splitKeys` of an existing table will cause the provider
	// to delete/recreate the entire `bigtable.Table` resource.
	SplitKeys pulumi.StringArrayInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies []TableColumnFamily `pulumi:"columnFamilies"`
	// A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
	DeletionProtection *string `pulumi:"deletionProtection"`
	// The name of the Bigtable instance.
	InstanceName string `pulumi:"instanceName"`
	// The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A list of predefined keys to split the table on.
	// !> **Warning:** Modifying the `splitKeys` of an existing table will cause the provider
	// to delete/recreate the entire `bigtable.Table` resource.
	SplitKeys []string `pulumi:"splitKeys"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies TableColumnFamilyArrayInput
	// A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
	DeletionProtection pulumi.StringPtrInput
	// The name of the Bigtable instance.
	InstanceName pulumi.StringInput
	// The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A list of predefined keys to split the table on.
	// !> **Warning:** Modifying the `splitKeys` of an existing table will cause the provider
	// to delete/recreate the entire `bigtable.Table` resource.
	SplitKeys pulumi.StringArrayInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

// TableArrayInput is an input type that accepts TableArray and TableArrayOutput values.
// You can construct a concrete instance of `TableArrayInput` via:
//
//	TableArray{ TableArgs{...} }
type TableArrayInput interface {
	pulumi.Input

	ToTableArrayOutput() TableArrayOutput
	ToTableArrayOutputWithContext(context.Context) TableArrayOutput
}

type TableArray []TableInput

func (TableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Table)(nil)).Elem()
}

func (i TableArray) ToTableArrayOutput() TableArrayOutput {
	return i.ToTableArrayOutputWithContext(context.Background())
}

func (i TableArray) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableArrayOutput)
}

// TableMapInput is an input type that accepts TableMap and TableMapOutput values.
// You can construct a concrete instance of `TableMapInput` via:
//
//	TableMap{ "key": TableArgs{...} }
type TableMapInput interface {
	pulumi.Input

	ToTableMapOutput() TableMapOutput
	ToTableMapOutputWithContext(context.Context) TableMapOutput
}

type TableMap map[string]TableInput

func (TableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Table)(nil)).Elem()
}

func (i TableMap) ToTableMapOutput() TableMapOutput {
	return i.ToTableMapOutputWithContext(context.Background())
}

func (i TableMap) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableMapOutput)
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
func (o TableOutput) ColumnFamilies() TableColumnFamilyArrayOutput {
	return o.ApplyT(func(v *Table) TableColumnFamilyArrayOutput { return v.ColumnFamilies }).(TableColumnFamilyArrayOutput)
}

// A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
func (o TableOutput) DeletionProtection() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.DeletionProtection }).(pulumi.StringOutput)
}

// The name of the Bigtable instance.
func (o TableOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
func (o TableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (o TableOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A list of predefined keys to split the table on.
// !> **Warning:** Modifying the `splitKeys` of an existing table will cause the provider
// to delete/recreate the entire `bigtable.Table` resource.
func (o TableOutput) SplitKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Table) pulumi.StringArrayOutput { return v.SplitKeys }).(pulumi.StringArrayOutput)
}

type TableArrayOutput struct{ *pulumi.OutputState }

func (TableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Table)(nil)).Elem()
}

func (o TableArrayOutput) ToTableArrayOutput() TableArrayOutput {
	return o
}

func (o TableArrayOutput) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return o
}

func (o TableArrayOutput) Index(i pulumi.IntInput) TableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Table {
		return vs[0].([]*Table)[vs[1].(int)]
	}).(TableOutput)
}

type TableMapOutput struct{ *pulumi.OutputState }

func (TableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Table)(nil)).Elem()
}

func (o TableMapOutput) ToTableMapOutput() TableMapOutput {
	return o
}

func (o TableMapOutput) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return o
}

func (o TableMapOutput) MapIndex(k pulumi.StringInput) TableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Table {
		return vs[0].(map[string]*Table)[vs[1].(string)]
	}).(TableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableInput)(nil)).Elem(), &Table{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableArrayInput)(nil)).Elem(), TableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableMapInput)(nil)).Elem(), TableMap{})
	pulumi.RegisterOutputType(TableOutput{})
	pulumi.RegisterOutputType(TableArrayOutput{})
	pulumi.RegisterOutputType(TableMapOutput{})
}

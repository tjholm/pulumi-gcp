// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package biglake

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a table.
//
// To get more information about Table, see:
//
// * [API documentation](https://cloud.google.com/bigquery/docs/reference/biglake/rest/v1/projects.locations.catalogs.databases.tables)
// * How-to Guides
//   - [Manage open source metadata with BigLake Metastore](https://cloud.google.com/bigquery/docs/manage-open-source-metadata#create_tables)
//
// ## Example Usage
// ### Biglake Table
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/biglake"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			catalog, err := biglake.NewCatalog(ctx, "catalog", &biglake.CatalogArgs{
//				Location: pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
//				Location:                 pulumi.String("US"),
//				ForceDestroy:             pulumi.Bool(true),
//				UniformBucketLevelAccess: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			metadataFolder, err := storage.NewBucketObject(ctx, "metadataFolder", &storage.BucketObjectArgs{
//				Content: pulumi.String(" "),
//				Bucket:  bucket.Name,
//			})
//			if err != nil {
//				return err
//			}
//			dataFolder, err := storage.NewBucketObject(ctx, "dataFolder", &storage.BucketObjectArgs{
//				Content: pulumi.String(" "),
//				Bucket:  bucket.Name,
//			})
//			if err != nil {
//				return err
//			}
//			database, err := biglake.NewDatabase(ctx, "database", &biglake.DatabaseArgs{
//				Catalog: catalog.ID(),
//				Type:    pulumi.String("HIVE"),
//				HiveOptions: &biglake.DatabaseHiveOptionsArgs{
//					LocationUri: pulumi.All(bucket.Name, metadataFolder.Name).ApplyT(func(_args []interface{}) (string, error) {
//						bucketName := _args[0].(string)
//						metadataFolderName := _args[1].(string)
//						return fmt.Sprintf("gs://%v/%v", bucketName, metadataFolderName), nil
//					}).(pulumi.StringOutput),
//					Parameters: pulumi.StringMap{
//						"owner": pulumi.String("Alex"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = biglake.NewTable(ctx, "table", &biglake.TableArgs{
//				Database: database.ID(),
//				Type:     pulumi.String("HIVE"),
//				HiveOptions: &biglake.TableHiveOptionsArgs{
//					TableType: pulumi.String("MANAGED_TABLE"),
//					StorageDescriptor: &biglake.TableHiveOptionsStorageDescriptorArgs{
//						LocationUri: pulumi.All(bucket.Name, dataFolder.Name).ApplyT(func(_args []interface{}) (string, error) {
//							bucketName := _args[0].(string)
//							dataFolderName := _args[1].(string)
//							return fmt.Sprintf("gs://%v/%v", bucketName, dataFolderName), nil
//						}).(pulumi.StringOutput),
//						InputFormat:  pulumi.String("org.apache.hadoop.mapred.SequenceFileInputFormat"),
//						OutputFormat: pulumi.String("org.apache.hadoop.hive.ql.io.HiveSequenceFileOutputFormat"),
//					},
//					Parameters: pulumi.StringMap{
//						"spark.sql.create.version":          pulumi.String("3.1.3"),
//						"spark.sql.sources.schema.numParts": pulumi.String("1"),
//						"transient_lastDdlTime":             pulumi.String("1680894197"),
//						"spark.sql.partitionProvider":       pulumi.String("catalog"),
//						"owner":                             pulumi.String("John Doe"),
//						"spark.sql.sources.schema.part.0":   pulumi.String("{\"type\":\"struct\",\"fields\":[{\"name\":\"id\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}},{\"name\":\"name\",\"type\":\"string\",\"nullable\":true,\"metadata\":{}},{\"name\":\"age\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}}]}"),
//						"spark.sql.sources.provider":        pulumi.String("iceberg"),
//						"provider":                          pulumi.String("iceberg"),
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
// Table can be imported using any of these accepted formats* `{{database}}/tables/{{name}}` When using the `pulumi import` command, Table can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:biglake/table:Table default {{database}}/tables/{{name}}
//
// ```
type Table struct {
	pulumi.CustomResourceState

	// Output only. The creation time of the table. A timestamp in RFC3339 UTC
	// "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and
	// "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The id of the parent database.
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// Output only. The deletion time of the table. Only set after the
	// table is deleted. A timestamp in RFC3339 UTC "Zulu" format, with
	// nanosecond resolution and up to nine fractional digits. Examples:
	// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// The checksum of a table object computed by the server based on the value
	// of other fields. It may be sent on update requests to ensure the client
	// has an up-to-date value before proceeding. It is only checked for update
	// table operations.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Output only. The time when this table is considered expired. Only set
	// after the table is deleted. A timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits. Examples:
	// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Options of a Hive table.
	// Structure is documented below.
	HiveOptions TableHiveOptionsPtrOutput `pulumi:"hiveOptions"`
	// Output only. The name of the Table. Format:
	// projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}/databases/{databaseId}/tables/{tableId}
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// The database type.
	// Possible values are: `HIVE`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Output only. The last modification time of the table. A timestamp in
	// RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and
	// "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		args = &TableArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Table
	err := ctx.RegisterResource("gcp:biglake/table:Table", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:biglake/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// Output only. The creation time of the table. A timestamp in RFC3339 UTC
	// "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and
	// "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// The id of the parent database.
	Database *string `pulumi:"database"`
	// Output only. The deletion time of the table. Only set after the
	// table is deleted. A timestamp in RFC3339 UTC "Zulu" format, with
	// nanosecond resolution and up to nine fractional digits. Examples:
	// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	DeleteTime *string `pulumi:"deleteTime"`
	// The checksum of a table object computed by the server based on the value
	// of other fields. It may be sent on update requests to ensure the client
	// has an up-to-date value before proceeding. It is only checked for update
	// table operations.
	Etag *string `pulumi:"etag"`
	// Output only. The time when this table is considered expired. Only set
	// after the table is deleted. A timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits. Examples:
	// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	ExpireTime *string `pulumi:"expireTime"`
	// Options of a Hive table.
	// Structure is documented below.
	HiveOptions *TableHiveOptions `pulumi:"hiveOptions"`
	// Output only. The name of the Table. Format:
	// projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}/databases/{databaseId}/tables/{tableId}
	//
	// ***
	Name *string `pulumi:"name"`
	// The database type.
	// Possible values are: `HIVE`.
	Type *string `pulumi:"type"`
	// Output only. The last modification time of the table. A timestamp in
	// RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and
	// "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type TableState struct {
	// Output only. The creation time of the table. A timestamp in RFC3339 UTC
	// "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and
	// "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// The id of the parent database.
	Database pulumi.StringPtrInput
	// Output only. The deletion time of the table. Only set after the
	// table is deleted. A timestamp in RFC3339 UTC "Zulu" format, with
	// nanosecond resolution and up to nine fractional digits. Examples:
	// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	DeleteTime pulumi.StringPtrInput
	// The checksum of a table object computed by the server based on the value
	// of other fields. It may be sent on update requests to ensure the client
	// has an up-to-date value before proceeding. It is only checked for update
	// table operations.
	Etag pulumi.StringPtrInput
	// Output only. The time when this table is considered expired. Only set
	// after the table is deleted. A timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits. Examples:
	// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	ExpireTime pulumi.StringPtrInput
	// Options of a Hive table.
	// Structure is documented below.
	HiveOptions TableHiveOptionsPtrInput
	// Output only. The name of the Table. Format:
	// projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}/databases/{databaseId}/tables/{tableId}
	//
	// ***
	Name pulumi.StringPtrInput
	// The database type.
	// Possible values are: `HIVE`.
	Type pulumi.StringPtrInput
	// Output only. The last modification time of the table. A timestamp in
	// RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and
	// "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// The id of the parent database.
	Database *string `pulumi:"database"`
	// Options of a Hive table.
	// Structure is documented below.
	HiveOptions *TableHiveOptions `pulumi:"hiveOptions"`
	// Output only. The name of the Table. Format:
	// projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}/databases/{databaseId}/tables/{tableId}
	//
	// ***
	Name *string `pulumi:"name"`
	// The database type.
	// Possible values are: `HIVE`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// The id of the parent database.
	Database pulumi.StringPtrInput
	// Options of a Hive table.
	// Structure is documented below.
	HiveOptions TableHiveOptionsPtrInput
	// Output only. The name of the Table. Format:
	// projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}/databases/{databaseId}/tables/{tableId}
	//
	// ***
	Name pulumi.StringPtrInput
	// The database type.
	// Possible values are: `HIVE`.
	Type pulumi.StringPtrInput
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

// Output only. The creation time of the table. A timestamp in RFC3339 UTC
// "Zulu" format, with nanosecond resolution and up to nine fractional
// digits. Examples: "2014-10-02T15:01:23Z" and
// "2014-10-02T15:01:23.045123456Z".
func (o TableOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The id of the parent database.
func (o TableOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.Database }).(pulumi.StringPtrOutput)
}

// Output only. The deletion time of the table. Only set after the
// table is deleted. A timestamp in RFC3339 UTC "Zulu" format, with
// nanosecond resolution and up to nine fractional digits. Examples:
// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o TableOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// The checksum of a table object computed by the server based on the value
// of other fields. It may be sent on update requests to ensure the client
// has an up-to-date value before proceeding. It is only checked for update
// table operations.
func (o TableOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Output only. The time when this table is considered expired. Only set
// after the table is deleted. A timestamp in RFC3339 UTC "Zulu" format,
// with nanosecond resolution and up to nine fractional digits. Examples:
// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o TableOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// Options of a Hive table.
// Structure is documented below.
func (o TableOutput) HiveOptions() TableHiveOptionsPtrOutput {
	return o.ApplyT(func(v *Table) TableHiveOptionsPtrOutput { return v.HiveOptions }).(TableHiveOptionsPtrOutput)
}

// Output only. The name of the Table. Format:
// projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}/databases/{databaseId}/tables/{tableId}
//
// ***
func (o TableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The database type.
// Possible values are: `HIVE`.
func (o TableOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// Output only. The last modification time of the table. A timestamp in
// RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
// fractional digits. Examples: "2014-10-02T15:01:23Z" and
// "2014-10-02T15:01:23.045123456Z".
func (o TableOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
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

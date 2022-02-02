// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A source representation instance is a Cloud SQL instance that represents
// the source database server to the Cloud SQL replica. It is visible in the
// Cloud Console and appears the same as a regular Cloud SQL instance, but it
// contains no data, requires no configuration or maintenance, and does not
// affect billing. You cannot update the source representation instance.
//
// ## Example Usage
// ### Sql Source Representation Instance Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewSourceRepresentationInstance(ctx, "instance", &sql.SourceRepresentationInstanceArgs{
// 			DatabaseVersion: pulumi.String("MYSQL_8_0"),
// 			Host:            pulumi.String("10.20.30.40"),
// 			Port:            pulumi.Int(3306),
// 			Region:          pulumi.String("us-central1"),
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
// SourceRepresentationInstance can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance default projects/{{project}}/instances/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance default {{name}}
// ```
type SourceRepresentationInstance struct {
	pulumi.CustomResourceState

	// The MySQL version running on your source database server.
	// Possible values are `MYSQL_5_5`, `MYSQL_5_6`, `MYSQL_5_7`, and `MYSQL_8_0`.
	DatabaseVersion pulumi.StringOutput `pulumi:"databaseVersion"`
	// The externally accessible IPv4 address for the source database server.
	Host pulumi.StringOutput `pulumi:"host"`
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewSourceRepresentationInstance registers a new resource with the given unique name, arguments, and options.
func NewSourceRepresentationInstance(ctx *pulumi.Context,
	name string, args *SourceRepresentationInstanceArgs, opts ...pulumi.ResourceOption) (*SourceRepresentationInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseVersion == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseVersion'")
	}
	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	var resource SourceRepresentationInstance
	err := ctx.RegisterResource("gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceRepresentationInstance gets an existing SourceRepresentationInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceRepresentationInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceRepresentationInstanceState, opts ...pulumi.ResourceOption) (*SourceRepresentationInstance, error) {
	var resource SourceRepresentationInstance
	err := ctx.ReadResource("gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceRepresentationInstance resources.
type sourceRepresentationInstanceState struct {
	// The MySQL version running on your source database server.
	// Possible values are `MYSQL_5_5`, `MYSQL_5_6`, `MYSQL_5_7`, and `MYSQL_8_0`.
	DatabaseVersion *string `pulumi:"databaseVersion"`
	// The externally accessible IPv4 address for the source database server.
	Host *string `pulumi:"host"`
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name *string `pulumi:"name"`
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}

type SourceRepresentationInstanceState struct {
	// The MySQL version running on your source database server.
	// Possible values are `MYSQL_5_5`, `MYSQL_5_6`, `MYSQL_5_7`, and `MYSQL_8_0`.
	DatabaseVersion pulumi.StringPtrInput
	// The externally accessible IPv4 address for the source database server.
	Host pulumi.StringPtrInput
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name pulumi.StringPtrInput
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
}

func (SourceRepresentationInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRepresentationInstanceState)(nil)).Elem()
}

type sourceRepresentationInstanceArgs struct {
	// The MySQL version running on your source database server.
	// Possible values are `MYSQL_5_5`, `MYSQL_5_6`, `MYSQL_5_7`, and `MYSQL_8_0`.
	DatabaseVersion string `pulumi:"databaseVersion"`
	// The externally accessible IPv4 address for the source database server.
	Host string `pulumi:"host"`
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name *string `pulumi:"name"`
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a SourceRepresentationInstance resource.
type SourceRepresentationInstanceArgs struct {
	// The MySQL version running on your source database server.
	// Possible values are `MYSQL_5_5`, `MYSQL_5_6`, `MYSQL_5_7`, and `MYSQL_8_0`.
	DatabaseVersion pulumi.StringInput
	// The externally accessible IPv4 address for the source database server.
	Host pulumi.StringInput
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name pulumi.StringPtrInput
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
}

func (SourceRepresentationInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRepresentationInstanceArgs)(nil)).Elem()
}

type SourceRepresentationInstanceInput interface {
	pulumi.Input

	ToSourceRepresentationInstanceOutput() SourceRepresentationInstanceOutput
	ToSourceRepresentationInstanceOutputWithContext(ctx context.Context) SourceRepresentationInstanceOutput
}

func (*SourceRepresentationInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRepresentationInstance)(nil)).Elem()
}

func (i *SourceRepresentationInstance) ToSourceRepresentationInstanceOutput() SourceRepresentationInstanceOutput {
	return i.ToSourceRepresentationInstanceOutputWithContext(context.Background())
}

func (i *SourceRepresentationInstance) ToSourceRepresentationInstanceOutputWithContext(ctx context.Context) SourceRepresentationInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRepresentationInstanceOutput)
}

// SourceRepresentationInstanceArrayInput is an input type that accepts SourceRepresentationInstanceArray and SourceRepresentationInstanceArrayOutput values.
// You can construct a concrete instance of `SourceRepresentationInstanceArrayInput` via:
//
//          SourceRepresentationInstanceArray{ SourceRepresentationInstanceArgs{...} }
type SourceRepresentationInstanceArrayInput interface {
	pulumi.Input

	ToSourceRepresentationInstanceArrayOutput() SourceRepresentationInstanceArrayOutput
	ToSourceRepresentationInstanceArrayOutputWithContext(context.Context) SourceRepresentationInstanceArrayOutput
}

type SourceRepresentationInstanceArray []SourceRepresentationInstanceInput

func (SourceRepresentationInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceRepresentationInstance)(nil)).Elem()
}

func (i SourceRepresentationInstanceArray) ToSourceRepresentationInstanceArrayOutput() SourceRepresentationInstanceArrayOutput {
	return i.ToSourceRepresentationInstanceArrayOutputWithContext(context.Background())
}

func (i SourceRepresentationInstanceArray) ToSourceRepresentationInstanceArrayOutputWithContext(ctx context.Context) SourceRepresentationInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRepresentationInstanceArrayOutput)
}

// SourceRepresentationInstanceMapInput is an input type that accepts SourceRepresentationInstanceMap and SourceRepresentationInstanceMapOutput values.
// You can construct a concrete instance of `SourceRepresentationInstanceMapInput` via:
//
//          SourceRepresentationInstanceMap{ "key": SourceRepresentationInstanceArgs{...} }
type SourceRepresentationInstanceMapInput interface {
	pulumi.Input

	ToSourceRepresentationInstanceMapOutput() SourceRepresentationInstanceMapOutput
	ToSourceRepresentationInstanceMapOutputWithContext(context.Context) SourceRepresentationInstanceMapOutput
}

type SourceRepresentationInstanceMap map[string]SourceRepresentationInstanceInput

func (SourceRepresentationInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceRepresentationInstance)(nil)).Elem()
}

func (i SourceRepresentationInstanceMap) ToSourceRepresentationInstanceMapOutput() SourceRepresentationInstanceMapOutput {
	return i.ToSourceRepresentationInstanceMapOutputWithContext(context.Background())
}

func (i SourceRepresentationInstanceMap) ToSourceRepresentationInstanceMapOutputWithContext(ctx context.Context) SourceRepresentationInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRepresentationInstanceMapOutput)
}

type SourceRepresentationInstanceOutput struct{ *pulumi.OutputState }

func (SourceRepresentationInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRepresentationInstance)(nil)).Elem()
}

func (o SourceRepresentationInstanceOutput) ToSourceRepresentationInstanceOutput() SourceRepresentationInstanceOutput {
	return o
}

func (o SourceRepresentationInstanceOutput) ToSourceRepresentationInstanceOutputWithContext(ctx context.Context) SourceRepresentationInstanceOutput {
	return o
}

type SourceRepresentationInstanceArrayOutput struct{ *pulumi.OutputState }

func (SourceRepresentationInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceRepresentationInstance)(nil)).Elem()
}

func (o SourceRepresentationInstanceArrayOutput) ToSourceRepresentationInstanceArrayOutput() SourceRepresentationInstanceArrayOutput {
	return o
}

func (o SourceRepresentationInstanceArrayOutput) ToSourceRepresentationInstanceArrayOutputWithContext(ctx context.Context) SourceRepresentationInstanceArrayOutput {
	return o
}

func (o SourceRepresentationInstanceArrayOutput) Index(i pulumi.IntInput) SourceRepresentationInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceRepresentationInstance {
		return vs[0].([]*SourceRepresentationInstance)[vs[1].(int)]
	}).(SourceRepresentationInstanceOutput)
}

type SourceRepresentationInstanceMapOutput struct{ *pulumi.OutputState }

func (SourceRepresentationInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceRepresentationInstance)(nil)).Elem()
}

func (o SourceRepresentationInstanceMapOutput) ToSourceRepresentationInstanceMapOutput() SourceRepresentationInstanceMapOutput {
	return o
}

func (o SourceRepresentationInstanceMapOutput) ToSourceRepresentationInstanceMapOutputWithContext(ctx context.Context) SourceRepresentationInstanceMapOutput {
	return o
}

func (o SourceRepresentationInstanceMapOutput) MapIndex(k pulumi.StringInput) SourceRepresentationInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceRepresentationInstance {
		return vs[0].(map[string]*SourceRepresentationInstance)[vs[1].(string)]
	}).(SourceRepresentationInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRepresentationInstanceInput)(nil)).Elem(), &SourceRepresentationInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRepresentationInstanceArrayInput)(nil)).Elem(), SourceRepresentationInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRepresentationInstanceMapInput)(nil)).Elem(), SourceRepresentationInstanceMap{})
	pulumi.RegisterOutputType(SourceRepresentationInstanceOutput{})
	pulumi.RegisterOutputType(SourceRepresentationInstanceArrayOutput{})
	pulumi.RegisterOutputType(SourceRepresentationInstanceMapOutput{})
}

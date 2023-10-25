// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to get information about a list of databases in a Cloud SQL instance.
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sql.GetDatabases(ctx, &sql.GetDatabasesArgs{
//				Instance: google_sql_database_instance.Main.Name,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDatabases(ctx *pulumi.Context, args *GetDatabasesArgs, opts ...pulumi.InvokeOption) (*GetDatabasesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasesResult
	err := ctx.Invoke("gcp:sql/getDatabases:getDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesArgs struct {
	// The name of the Cloud SQL database instance in which the database belongs.
	Instance string `pulumi:"instance"`
	// The ID of the project in which the instance belongs.
	//
	// > **Note** This datasource performs client-side sorting to provide consistent ordering of the databases.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getDatabases.
type GetDatabasesResult struct {
	Databases []GetDatabasesDatabase `pulumi:"databases"`
	// The provider-assigned unique ID for this managed resource.
	Id       string  `pulumi:"id"`
	Instance string  `pulumi:"instance"`
	Project  *string `pulumi:"project"`
}

func GetDatabasesOutput(ctx *pulumi.Context, args GetDatabasesOutputArgs, opts ...pulumi.InvokeOption) GetDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabasesResult, error) {
			args := v.(GetDatabasesArgs)
			r, err := GetDatabases(ctx, &args, opts...)
			var s GetDatabasesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabasesResultOutput)
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesOutputArgs struct {
	// The name of the Cloud SQL database instance in which the database belongs.
	Instance pulumi.StringInput `pulumi:"instance"`
	// The ID of the project in which the instance belongs.
	//
	// > **Note** This datasource performs client-side sorting to provide consistent ordering of the databases.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getDatabases.
type GetDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesResult)(nil)).Elem()
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutput() GetDatabasesResultOutput {
	return o
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutputWithContext(ctx context.Context) GetDatabasesResultOutput {
	return o
}

func (o GetDatabasesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDatabasesResult] {
	return pulumix.Output[GetDatabasesResult]{
		OutputState: o.OutputState,
	}
}

func (o GetDatabasesResultOutput) Databases() GetDatabasesDatabaseArrayOutput {
	return o.ApplyT(func(v GetDatabasesResult) []GetDatabasesDatabase { return v.Databases }).(GetDatabasesDatabaseArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDatabasesResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.Instance }).(pulumi.StringOutput)
}

func (o GetDatabasesResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabasesResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasesResultOutput{})
}

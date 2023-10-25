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

// Use this data source to get information about a Cloud SQL instance backup run.
//
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
//			_, err := sql.GetBackupRun(ctx, &sql.GetBackupRunArgs{
//				Instance:   google_sql_database_instance.Main.Name,
//				MostRecent: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetBackupRun(ctx *pulumi.Context, args *GetBackupRunArgs, opts ...pulumi.InvokeOption) (*GetBackupRunResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBackupRunResult
	err := ctx.Invoke("gcp:sql/getBackupRun:getBackupRun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackupRun.
type GetBackupRunArgs struct {
	// The identifier for this backup run. Unique only for a specific Cloud SQL instance.
	// If left empty and multiple backups exist for the instance, `mostRecent` must be set to `true`.
	BackupId *int `pulumi:"backupId"`
	// The name of the instance the backup is taken from.
	Instance string `pulumi:"instance"`
	// Toggles use of the most recent backup run if multiple backups exist for a
	// Cloud SQL instance.
	MostRecent *bool `pulumi:"mostRecent"`
	// The project to list instances for. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getBackupRun.
type GetBackupRunResult struct {
	BackupId int `pulumi:"backupId"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Instance string `pulumi:"instance"`
	// Location of the backups.
	Location   string `pulumi:"location"`
	MostRecent *bool  `pulumi:"mostRecent"`
	Project    string `pulumi:"project"`
	// The time the backup operation actually started in UTC timezone in RFC 3339 format, for
	// example 2012-11-15T16:19:00.094Z.
	StartTime string `pulumi:"startTime"`
	// The status of this run. Refer to [API reference](https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/backupRuns#SqlBackupRunStatus) for possible status values.
	Status string `pulumi:"status"`
}

func GetBackupRunOutput(ctx *pulumi.Context, args GetBackupRunOutputArgs, opts ...pulumi.InvokeOption) GetBackupRunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackupRunResult, error) {
			args := v.(GetBackupRunArgs)
			r, err := GetBackupRun(ctx, &args, opts...)
			var s GetBackupRunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackupRunResultOutput)
}

// A collection of arguments for invoking getBackupRun.
type GetBackupRunOutputArgs struct {
	// The identifier for this backup run. Unique only for a specific Cloud SQL instance.
	// If left empty and multiple backups exist for the instance, `mostRecent` must be set to `true`.
	BackupId pulumi.IntPtrInput `pulumi:"backupId"`
	// The name of the instance the backup is taken from.
	Instance pulumi.StringInput `pulumi:"instance"`
	// Toggles use of the most recent backup run if multiple backups exist for a
	// Cloud SQL instance.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// The project to list instances for. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetBackupRunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupRunArgs)(nil)).Elem()
}

// A collection of values returned by getBackupRun.
type GetBackupRunResultOutput struct{ *pulumi.OutputState }

func (GetBackupRunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupRunResult)(nil)).Elem()
}

func (o GetBackupRunResultOutput) ToGetBackupRunResultOutput() GetBackupRunResultOutput {
	return o
}

func (o GetBackupRunResultOutput) ToGetBackupRunResultOutputWithContext(ctx context.Context) GetBackupRunResultOutput {
	return o
}

func (o GetBackupRunResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetBackupRunResult] {
	return pulumix.Output[GetBackupRunResult]{
		OutputState: o.OutputState,
	}
}

func (o GetBackupRunResultOutput) BackupId() pulumi.IntOutput {
	return o.ApplyT(func(v GetBackupRunResult) int { return v.BackupId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackupRunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupRunResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBackupRunResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupRunResult) string { return v.Instance }).(pulumi.StringOutput)
}

// Location of the backups.
func (o GetBackupRunResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupRunResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetBackupRunResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetBackupRunResult) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

func (o GetBackupRunResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupRunResult) string { return v.Project }).(pulumi.StringOutput)
}

// The time the backup operation actually started in UTC timezone in RFC 3339 format, for
// example 2012-11-15T16:19:00.094Z.
func (o GetBackupRunResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupRunResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// The status of this run. Refer to [API reference](https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/backupRuns#SqlBackupRunStatus) for possible status values.
func (o GetBackupRunResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupRunResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackupRunResultOutput{})
}

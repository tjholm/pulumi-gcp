// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a Cloud SQL instance.
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
//			_, err := sql.LookupDatabaseInstance(ctx, &sql.LookupDatabaseInstanceArgs{
//				Name: "test-sql-instance",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDatabaseInstance(ctx *pulumi.Context, args *LookupDatabaseInstanceArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseInstanceResult, error) {
	var rv LookupDatabaseInstanceResult
	err := ctx.Invoke("gcp:sql/getDatabaseInstance:getDatabaseInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseInstance.
type LookupDatabaseInstanceArgs struct {
	// The name of the instance.
	Name string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getDatabaseInstance.
type LookupDatabaseInstanceResult struct {
	Clones             []GetDatabaseInstanceClone `pulumi:"clones"`
	ConnectionName     string                     `pulumi:"connectionName"`
	DatabaseVersion    string                     `pulumi:"databaseVersion"`
	DeletionProtection bool                       `pulumi:"deletionProtection"`
	EncryptionKeyName  string                     `pulumi:"encryptionKeyName"`
	FirstIpAddress     string                     `pulumi:"firstIpAddress"`
	// The provider-assigned unique ID for this managed resource.
	Id                         string                                    `pulumi:"id"`
	IpAddresses                []GetDatabaseInstanceIpAddress            `pulumi:"ipAddresses"`
	MasterInstanceName         string                                    `pulumi:"masterInstanceName"`
	Name                       string                                    `pulumi:"name"`
	PrivateIpAddress           string                                    `pulumi:"privateIpAddress"`
	Project                    *string                                   `pulumi:"project"`
	PublicIpAddress            string                                    `pulumi:"publicIpAddress"`
	Region                     string                                    `pulumi:"region"`
	ReplicaConfigurations      []GetDatabaseInstanceReplicaConfiguration `pulumi:"replicaConfigurations"`
	RestoreBackupContexts      []GetDatabaseInstanceRestoreBackupContext `pulumi:"restoreBackupContexts"`
	RootPassword               string                                    `pulumi:"rootPassword"`
	SelfLink                   string                                    `pulumi:"selfLink"`
	ServerCaCerts              []GetDatabaseInstanceServerCaCert         `pulumi:"serverCaCerts"`
	ServiceAccountEmailAddress string                                    `pulumi:"serviceAccountEmailAddress"`
	Settings                   []GetDatabaseInstanceSetting              `pulumi:"settings"`
}

func LookupDatabaseInstanceOutput(ctx *pulumi.Context, args LookupDatabaseInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseInstanceResult, error) {
			args := v.(LookupDatabaseInstanceArgs)
			r, err := LookupDatabaseInstance(ctx, &args, opts...)
			var s LookupDatabaseInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseInstanceResultOutput)
}

// A collection of arguments for invoking getDatabaseInstance.
type LookupDatabaseInstanceOutputArgs struct {
	// The name of the instance.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDatabaseInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseInstance.
type LookupDatabaseInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseInstanceResult)(nil)).Elem()
}

func (o LookupDatabaseInstanceResultOutput) ToLookupDatabaseInstanceResultOutput() LookupDatabaseInstanceResultOutput {
	return o
}

func (o LookupDatabaseInstanceResultOutput) ToLookupDatabaseInstanceResultOutputWithContext(ctx context.Context) LookupDatabaseInstanceResultOutput {
	return o
}

func (o LookupDatabaseInstanceResultOutput) Clones() GetDatabaseInstanceCloneArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []GetDatabaseInstanceClone { return v.Clones }).(GetDatabaseInstanceCloneArrayOutput)
}

func (o LookupDatabaseInstanceResultOutput) ConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.ConnectionName }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) DatabaseVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.DatabaseVersion }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

func (o LookupDatabaseInstanceResultOutput) EncryptionKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.EncryptionKeyName }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) FirstIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.FirstIpAddress }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatabaseInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) IpAddresses() GetDatabaseInstanceIpAddressArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []GetDatabaseInstanceIpAddress { return v.IpAddresses }).(GetDatabaseInstanceIpAddressArrayOutput)
}

func (o LookupDatabaseInstanceResultOutput) MasterInstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.MasterInstanceName }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseInstanceResultOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.PublicIpAddress }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) ReplicaConfigurations() GetDatabaseInstanceReplicaConfigurationArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []GetDatabaseInstanceReplicaConfiguration {
		return v.ReplicaConfigurations
	}).(GetDatabaseInstanceReplicaConfigurationArrayOutput)
}

func (o LookupDatabaseInstanceResultOutput) RestoreBackupContexts() GetDatabaseInstanceRestoreBackupContextArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []GetDatabaseInstanceRestoreBackupContext {
		return v.RestoreBackupContexts
	}).(GetDatabaseInstanceRestoreBackupContextArrayOutput)
}

func (o LookupDatabaseInstanceResultOutput) RootPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.RootPassword }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) ServerCaCerts() GetDatabaseInstanceServerCaCertArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []GetDatabaseInstanceServerCaCert { return v.ServerCaCerts }).(GetDatabaseInstanceServerCaCertArrayOutput)
}

func (o LookupDatabaseInstanceResultOutput) ServiceAccountEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.ServiceAccountEmailAddress }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) Settings() GetDatabaseInstanceSettingArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []GetDatabaseInstanceSetting { return v.Settings }).(GetDatabaseInstanceSettingArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseInstanceResultOutput{})
}

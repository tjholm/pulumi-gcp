// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alloydb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Alloydb Cluster Basic
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/alloydb"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := compute.NewNetwork(ctx, "defaultNetwork", nil)
//			if err != nil {
//				return err
//			}
//			_, err = alloydb.NewCluster(ctx, "defaultCluster", &alloydb.ClusterArgs{
//				ClusterId: pulumi.String("alloydb-cluster"),
//				Location:  pulumi.String("us-central1"),
//				Network: defaultNetwork.Name.ApplyT(func(name string) (string, error) {
//					return fmt.Sprintf("projects/%v/global/networks/%v", project.Number, name), nil
//				}).(pulumi.StringOutput),
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
// # Cluster can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:alloydb/cluster:Cluster default projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:alloydb/cluster:Cluster default {{project}}/{{location}}/{{cluster_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:alloydb/cluster:Cluster default {{location}}/{{cluster_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:alloydb/cluster:Cluster default {{cluster_id}}
//
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// The automated backup policy for this cluster.
	// If no policy is provided then the default policy will be used. The default policy takes one backup a day, has a backup window of 1 hour, and retains backups for 14 days.
	// Structure is documented below.
	AutomatedBackupPolicy ClusterAutomatedBackupPolicyPtrOutput `pulumi:"automatedBackupPolicy"`
	// Cluster created from backup.
	// Structure is documented below.
	BackupSources ClusterBackupSourceArrayOutput `pulumi:"backupSources"`
	// The ID of the alloydb cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The database engine major version. This is an output-only field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.
	DatabaseVersion pulumi.StringOutput `pulumi:"databaseVersion"`
	// User-settable and human-readable display name for the Cluster.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Initial user to setup during cluster creation.
	// Structure is documented below.
	InitialUser ClusterInitialUserPtrOutput `pulumi:"initialUser"`
	// User-defined labels for the alloydb cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location where the alloydb cluster should reside.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Cluster created via DMS migration.
	// Structure is documented below.
	MigrationSources ClusterMigrationSourceArrayOutput `pulumi:"migrationSources"`
	// The name of the cluster resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
	// "projects/{projectNumber}/global/networks/{network_id}".
	Network pulumi.StringOutput `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The system-generated UID of the resource.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	var resource Cluster
	err := ctx.RegisterResource("gcp:alloydb/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("gcp:alloydb/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The automated backup policy for this cluster.
	// If no policy is provided then the default policy will be used. The default policy takes one backup a day, has a backup window of 1 hour, and retains backups for 14 days.
	// Structure is documented below.
	AutomatedBackupPolicy *ClusterAutomatedBackupPolicy `pulumi:"automatedBackupPolicy"`
	// Cluster created from backup.
	// Structure is documented below.
	BackupSources []ClusterBackupSource `pulumi:"backupSources"`
	// The ID of the alloydb cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The database engine major version. This is an output-only field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.
	DatabaseVersion *string `pulumi:"databaseVersion"`
	// User-settable and human-readable display name for the Cluster.
	DisplayName *string `pulumi:"displayName"`
	// Initial user to setup during cluster creation.
	// Structure is documented below.
	InitialUser *ClusterInitialUser `pulumi:"initialUser"`
	// User-defined labels for the alloydb cluster.
	Labels map[string]string `pulumi:"labels"`
	// The location where the alloydb cluster should reside.
	Location *string `pulumi:"location"`
	// Cluster created via DMS migration.
	// Structure is documented below.
	MigrationSources []ClusterMigrationSource `pulumi:"migrationSources"`
	// The name of the cluster resource.
	Name *string `pulumi:"name"`
	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
	// "projects/{projectNumber}/global/networks/{network_id}".
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The system-generated UID of the resource.
	Uid *string `pulumi:"uid"`
}

type ClusterState struct {
	// The automated backup policy for this cluster.
	// If no policy is provided then the default policy will be used. The default policy takes one backup a day, has a backup window of 1 hour, and retains backups for 14 days.
	// Structure is documented below.
	AutomatedBackupPolicy ClusterAutomatedBackupPolicyPtrInput
	// Cluster created from backup.
	// Structure is documented below.
	BackupSources ClusterBackupSourceArrayInput
	// The ID of the alloydb cluster.
	ClusterId pulumi.StringPtrInput
	// The database engine major version. This is an output-only field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.
	DatabaseVersion pulumi.StringPtrInput
	// User-settable and human-readable display name for the Cluster.
	DisplayName pulumi.StringPtrInput
	// Initial user to setup during cluster creation.
	// Structure is documented below.
	InitialUser ClusterInitialUserPtrInput
	// User-defined labels for the alloydb cluster.
	Labels pulumi.StringMapInput
	// The location where the alloydb cluster should reside.
	Location pulumi.StringPtrInput
	// Cluster created via DMS migration.
	// Structure is documented below.
	MigrationSources ClusterMigrationSourceArrayInput
	// The name of the cluster resource.
	Name pulumi.StringPtrInput
	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
	// "projects/{projectNumber}/global/networks/{network_id}".
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The system-generated UID of the resource.
	Uid pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The automated backup policy for this cluster.
	// If no policy is provided then the default policy will be used. The default policy takes one backup a day, has a backup window of 1 hour, and retains backups for 14 days.
	// Structure is documented below.
	AutomatedBackupPolicy *ClusterAutomatedBackupPolicy `pulumi:"automatedBackupPolicy"`
	// The ID of the alloydb cluster.
	ClusterId string `pulumi:"clusterId"`
	// User-settable and human-readable display name for the Cluster.
	DisplayName *string `pulumi:"displayName"`
	// Initial user to setup during cluster creation.
	// Structure is documented below.
	InitialUser *ClusterInitialUser `pulumi:"initialUser"`
	// User-defined labels for the alloydb cluster.
	Labels map[string]string `pulumi:"labels"`
	// The location where the alloydb cluster should reside.
	Location *string `pulumi:"location"`
	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
	// "projects/{projectNumber}/global/networks/{network_id}".
	Network string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The automated backup policy for this cluster.
	// If no policy is provided then the default policy will be used. The default policy takes one backup a day, has a backup window of 1 hour, and retains backups for 14 days.
	// Structure is documented below.
	AutomatedBackupPolicy ClusterAutomatedBackupPolicyPtrInput
	// The ID of the alloydb cluster.
	ClusterId pulumi.StringInput
	// User-settable and human-readable display name for the Cluster.
	DisplayName pulumi.StringPtrInput
	// Initial user to setup during cluster creation.
	// Structure is documented below.
	InitialUser ClusterInitialUserPtrInput
	// User-defined labels for the alloydb cluster.
	Labels pulumi.StringMapInput
	// The location where the alloydb cluster should reside.
	Location pulumi.StringPtrInput
	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
	// "projects/{projectNumber}/global/networks/{network_id}".
	Network pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//	ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//	ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// The automated backup policy for this cluster.
// If no policy is provided then the default policy will be used. The default policy takes one backup a day, has a backup window of 1 hour, and retains backups for 14 days.
// Structure is documented below.
func (o ClusterOutput) AutomatedBackupPolicy() ClusterAutomatedBackupPolicyPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterAutomatedBackupPolicyPtrOutput { return v.AutomatedBackupPolicy }).(ClusterAutomatedBackupPolicyPtrOutput)
}

// Cluster created from backup.
// Structure is documented below.
func (o ClusterOutput) BackupSources() ClusterBackupSourceArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterBackupSourceArrayOutput { return v.BackupSources }).(ClusterBackupSourceArrayOutput)
}

// The ID of the alloydb cluster.
func (o ClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The database engine major version. This is an output-only field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.
func (o ClusterOutput) DatabaseVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.DatabaseVersion }).(pulumi.StringOutput)
}

// User-settable and human-readable display name for the Cluster.
func (o ClusterOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Initial user to setup during cluster creation.
// Structure is documented below.
func (o ClusterOutput) InitialUser() ClusterInitialUserPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterInitialUserPtrOutput { return v.InitialUser }).(ClusterInitialUserPtrOutput)
}

// User-defined labels for the alloydb cluster.
func (o ClusterOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location where the alloydb cluster should reside.
func (o ClusterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Cluster created via DMS migration.
// Structure is documented below.
func (o ClusterOutput) MigrationSources() ClusterMigrationSourceArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterMigrationSourceArrayOutput { return v.MigrationSources }).(ClusterMigrationSourceArrayOutput)
}

// The name of the cluster resource.
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
// "projects/{projectNumber}/global/networks/{network_id}".
func (o ClusterOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o ClusterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The system-generated UID of the resource.
func (o ClusterOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}

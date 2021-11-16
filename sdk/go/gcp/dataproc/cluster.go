// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Cloud Dataproc cluster resource within GCP.
//
// * [API documentation](https://cloud.google.com/dataproc/docs/reference/rest/v1/projects.regions.clusters)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/dataproc/docs)
//
// !> **Warning:** Due to limitations of the API, all arguments except
// `labels`,`cluster_config.worker_config.num_instances` and `cluster_config.preemptible_worker_config.num_instances` are non-updatable. Changing others will cause recreation of the
// whole cluster!
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewCluster(ctx, "simplecluster", &dataproc.ClusterArgs{
// 			Region: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Advanced
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := serviceAccount.NewAccount(ctx, "_default", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("service-account-id"),
// 			DisplayName: pulumi.String("Service Account"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dataproc.NewCluster(ctx, "mycluster", &dataproc.ClusterArgs{
// 			Region:                      pulumi.String("us-central1"),
// 			GracefulDecommissionTimeout: pulumi.String("120s"),
// 			Labels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			ClusterConfig: &dataproc.ClusterClusterConfigArgs{
// 				StagingBucket: pulumi.String("dataproc-staging-bucket"),
// 				MasterConfig: &dataproc.ClusterClusterConfigMasterConfigArgs{
// 					NumInstances: pulumi.Int(1),
// 					MachineType:  pulumi.String("e2-medium"),
// 					DiskConfig: &dataproc.ClusterClusterConfigMasterConfigDiskConfigArgs{
// 						BootDiskType:   pulumi.String("pd-ssd"),
// 						BootDiskSizeGb: pulumi.Int(30),
// 					},
// 				},
// 				WorkerConfig: &dataproc.ClusterClusterConfigWorkerConfigArgs{
// 					NumInstances:   pulumi.Int(2),
// 					MachineType:    pulumi.String("e2-medium"),
// 					MinCpuPlatform: pulumi.String("Intel Skylake"),
// 					DiskConfig: &dataproc.ClusterClusterConfigWorkerConfigDiskConfigArgs{
// 						BootDiskSizeGb: pulumi.Int(30),
// 						NumLocalSsds:   pulumi.Int(1),
// 					},
// 				},
// 				PreemptibleWorkerConfig: &dataproc.ClusterClusterConfigPreemptibleWorkerConfigArgs{
// 					NumInstances: pulumi.Int(0),
// 				},
// 				SoftwareConfig: &dataproc.ClusterClusterConfigSoftwareConfigArgs{
// 					ImageVersion: pulumi.String("1.3.7-deb9"),
// 					OverrideProperties: pulumi.StringMap{
// 						"dataproc:dataproc.allow.zero.workers": pulumi.String("true"),
// 					},
// 				},
// 				GceClusterConfig: &dataproc.ClusterClusterConfigGceClusterConfigArgs{
// 					Tags: pulumi.StringArray{
// 						pulumi.String("foo"),
// 						pulumi.String("bar"),
// 					},
// 					ServiceAccount: _default.Email,
// 					ServiceAccountScopes: pulumi.StringArray{
// 						pulumi.String("cloud-platform"),
// 					},
// 				},
// 				InitializationActions: dataproc.ClusterClusterConfigInitializationActionArray{
// 					&dataproc.ClusterClusterConfigInitializationActionArgs{
// 						Script:     pulumi.String("gs://dataproc-initialization-actions/stackdriver/stackdriver.sh"),
// 						TimeoutSec: pulumi.Int(500),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using A GPU Accelerator
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewCluster(ctx, "acceleratedCluster", &dataproc.ClusterArgs{
// 			ClusterConfig: &dataproc.ClusterClusterConfigArgs{
// 				GceClusterConfig: &dataproc.ClusterClusterConfigGceClusterConfigArgs{
// 					Zone: pulumi.String("us-central1-a"),
// 				},
// 				MasterConfig: &dataproc.ClusterClusterConfigMasterConfigArgs{
// 					Accelerators: dataproc.ClusterClusterConfigMasterConfigAcceleratorArray{
// 						&dataproc.ClusterClusterConfigMasterConfigAcceleratorArgs{
// 							AcceleratorCount: pulumi.Int(1),
// 							AcceleratorType:  pulumi.String("nvidia-tesla-k80"),
// 						},
// 					},
// 				},
// 			},
// 			Region: pulumi.String("us-central1"),
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
// This resource does not support import.
type Cluster struct {
	pulumi.CustomResourceState

	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig ClusterClusterConfigOutput `pulumi:"clusterConfig"`
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	// terraform apply
	GracefulDecommissionTimeout pulumi.StringPtrOutput `pulumi:"gracefulDecommissionTimeout"`
	// The list of labels (key/value pairs) to be applied to
	// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
	// which is the name of the cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the cluster, unique within the project and
	// zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the `cluster` will exist. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the cluster and associated nodes will be created in.
	// Defaults to `global`.
	Region pulumi.StringPtrOutput `pulumi:"region"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		args = &ClusterArgs{}
	}

	var resource Cluster
	err := ctx.RegisterResource("gcp:dataproc/cluster:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:dataproc/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig *ClusterClusterConfig `pulumi:"clusterConfig"`
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	// terraform apply
	GracefulDecommissionTimeout *string `pulumi:"gracefulDecommissionTimeout"`
	// The list of labels (key/value pairs) to be applied to
	// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
	// which is the name of the cluster.
	Labels map[string]string `pulumi:"labels"`
	// The name of the cluster, unique within the project and
	// zone.
	Name *string `pulumi:"name"`
	// The ID of the project in which the `cluster` will exist. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the cluster and associated nodes will be created in.
	// Defaults to `global`.
	Region *string `pulumi:"region"`
}

type ClusterState struct {
	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig ClusterClusterConfigPtrInput
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	// terraform apply
	GracefulDecommissionTimeout pulumi.StringPtrInput
	// The list of labels (key/value pairs) to be applied to
	// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
	// which is the name of the cluster.
	Labels pulumi.StringMapInput
	// The name of the cluster, unique within the project and
	// zone.
	Name pulumi.StringPtrInput
	// The ID of the project in which the `cluster` will exist. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the cluster and associated nodes will be created in.
	// Defaults to `global`.
	Region pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig *ClusterClusterConfig `pulumi:"clusterConfig"`
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	// terraform apply
	GracefulDecommissionTimeout *string `pulumi:"gracefulDecommissionTimeout"`
	// The list of labels (key/value pairs) to be applied to
	// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
	// which is the name of the cluster.
	Labels map[string]string `pulumi:"labels"`
	// The name of the cluster, unique within the project and
	// zone.
	Name *string `pulumi:"name"`
	// The ID of the project in which the `cluster` will exist. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the cluster and associated nodes will be created in.
	// Defaults to `global`.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig ClusterClusterConfigPtrInput
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	// terraform apply
	GracefulDecommissionTimeout pulumi.StringPtrInput
	// The list of labels (key/value pairs) to be applied to
	// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
	// which is the name of the cluster.
	Labels pulumi.StringMapInput
	// The name of the cluster, unique within the project and
	// zone.
	Name pulumi.StringPtrInput
	// The ID of the project in which the `cluster` will exist. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the cluster and associated nodes will be created in.
	// Defaults to `global`.
	Region pulumi.StringPtrInput
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
	return reflect.TypeOf((*Cluster)(nil))
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

func (i *Cluster) ToClusterPtrOutput() ClusterPtrOutput {
	return i.ToClusterPtrOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPtrOutput)
}

type ClusterPtrInput interface {
	pulumi.Input

	ToClusterPtrOutput() ClusterPtrOutput
	ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput
}

type clusterPtrType ClusterArgs

func (*clusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (i *clusterPtrType) ToClusterPtrOutput() ClusterPtrOutput {
	return i.ToClusterPtrOutputWithContext(context.Background())
}

func (i *clusterPtrType) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPtrOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//          ClusterArray{ ClusterArgs{...} }
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
//          ClusterMap{ "key": ClusterArgs{...} }
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
	return reflect.TypeOf((*Cluster)(nil))
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o.ToClusterPtrOutputWithContext(context.Background())
}

func (o ClusterOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Cluster) *Cluster {
		return &v
	}).(ClusterPtrOutput)
}

type ClusterPtrOutput struct{ *pulumi.OutputState }

func (ClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (o ClusterPtrOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o
}

func (o ClusterPtrOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o
}

func (o ClusterPtrOutput) Elem() ClusterOutput {
	return o.ApplyT(func(v *Cluster) Cluster {
		if v != nil {
			return *v
		}
		var ret Cluster
		return ret
	}).(ClusterOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Cluster)(nil))
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Cluster {
		return vs[0].([]Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Cluster)(nil))
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Cluster {
		return vs[0].(map[string]Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterPtrOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}

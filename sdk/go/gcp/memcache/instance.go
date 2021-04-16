// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package memcache

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Google Cloud Memcache instance.
//
// To get more information about Instance, see:
//
// * [API documentation](https://cloud.google.com/memorystore/docs/memcached/reference/rest/v1beta2/projects.locations.instances)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/memcache/docs/creating-instances)
//
// ## Example Usage
// ### Memcache Instance Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/memcache"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/servicenetworking"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		memcacheNetwork, err := compute.LookupNetwork(ctx, &compute.LookupNetworkArgs{
// 			Name: "test-network",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		serviceRange, err := compute.NewGlobalAddress(ctx, "serviceRange", &compute.GlobalAddressArgs{
// 			Purpose:      pulumi.String("VPC_PEERING"),
// 			AddressType:  pulumi.String("INTERNAL"),
// 			PrefixLength: pulumi.Int(16),
// 			Network:      pulumi.String(memcacheNetwork.Id),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		privateServiceConnection, err := servicenetworking.NewConnection(ctx, "privateServiceConnection", &servicenetworking.ConnectionArgs{
// 			Network: pulumi.String(memcacheNetwork.Id),
// 			Service: pulumi.String("servicenetworking.googleapis.com"),
// 			ReservedPeeringRanges: pulumi.StringArray{
// 				serviceRange.Name,
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = memcache.NewInstance(ctx, "instance", &memcache.InstanceArgs{
// 			AuthorizedNetwork: privateServiceConnection.Network,
// 			NodeConfig: &memcache.InstanceNodeConfigArgs{
// 				CpuCount:     pulumi.Int(1),
// 				MemorySizeMb: pulumi.Int(1024),
// 			},
// 			NodeCount:       pulumi.Int(1),
// 			MemcacheVersion: pulumi.String("MEMCACHE_1_5"),
// 		}, pulumi.Provider(google_beta))
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
// Instance can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:memcache/instance:Instance default projects/{{project}}/locations/{{region}}/instances/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:memcache/instance:Instance default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:memcache/instance:Instance default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:memcache/instance:Instance default {{name}}
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The full name of the GCE network to connect the instance to.  If not provided,
	// 'default' will be used.
	AuthorizedNetwork pulumi.StringOutput `pulumi:"authorizedNetwork"`
	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Endpoint for Discovery API
	DiscoveryEndpoint pulumi.StringOutput `pulumi:"discoveryEndpoint"`
	// A user-visible name for the instance.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The full version of memcached server running on this instance.
	MemcacheFullVersion pulumi.StringOutput `pulumi:"memcacheFullVersion"`
	// Additional information about the instance state, if available.
	MemcacheNodes InstanceMemcacheNodeArrayOutput `pulumi:"memcacheNodes"`
	// User-specified parameters for this memcache instance.
	// Structure is documented below.
	MemcacheParameters InstanceMemcacheParametersPtrOutput `pulumi:"memcacheParameters"`
	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version.
	// Default value is `MEMCACHE_1_5`.
	// Possible values are `MEMCACHE_1_5`.
	MemcacheVersion pulumi.StringPtrOutput `pulumi:"memcacheVersion"`
	// The resource name of the instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration for memcache nodes.
	// Structure is documented below.
	NodeConfig InstanceNodeConfigOutput `pulumi:"nodeConfig"`
	// Number of nodes in the memcache instance.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the Memcache instance. If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NodeConfig == nil {
		return nil, errors.New("invalid value for required argument 'NodeConfig'")
	}
	if args.NodeCount == nil {
		return nil, errors.New("invalid value for required argument 'NodeCount'")
	}
	var resource Instance
	err := ctx.RegisterResource("gcp:memcache/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("gcp:memcache/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The full name of the GCE network to connect the instance to.  If not provided,
	// 'default' will be used.
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// Creation timestamp in RFC3339 text format.
	CreateTime *string `pulumi:"createTime"`
	// Endpoint for Discovery API
	DiscoveryEndpoint *string `pulumi:"discoveryEndpoint"`
	// A user-visible name for the instance.
	DisplayName *string `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The full version of memcached server running on this instance.
	MemcacheFullVersion *string `pulumi:"memcacheFullVersion"`
	// Additional information about the instance state, if available.
	MemcacheNodes []InstanceMemcacheNode `pulumi:"memcacheNodes"`
	// User-specified parameters for this memcache instance.
	// Structure is documented below.
	MemcacheParameters *InstanceMemcacheParameters `pulumi:"memcacheParameters"`
	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version.
	// Default value is `MEMCACHE_1_5`.
	// Possible values are `MEMCACHE_1_5`.
	MemcacheVersion *string `pulumi:"memcacheVersion"`
	// The resource name of the instance.
	Name *string `pulumi:"name"`
	// Configuration for memcache nodes.
	// Structure is documented below.
	NodeConfig *InstanceNodeConfig `pulumi:"nodeConfig"`
	// Number of nodes in the memcache instance.
	NodeCount *int `pulumi:"nodeCount"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the Memcache instance. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	Zones []string `pulumi:"zones"`
}

type InstanceState struct {
	// The full name of the GCE network to connect the instance to.  If not provided,
	// 'default' will be used.
	AuthorizedNetwork pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringPtrInput
	// Endpoint for Discovery API
	DiscoveryEndpoint pulumi.StringPtrInput
	// A user-visible name for the instance.
	DisplayName pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// The full version of memcached server running on this instance.
	MemcacheFullVersion pulumi.StringPtrInput
	// Additional information about the instance state, if available.
	MemcacheNodes InstanceMemcacheNodeArrayInput
	// User-specified parameters for this memcache instance.
	// Structure is documented below.
	MemcacheParameters InstanceMemcacheParametersPtrInput
	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version.
	// Default value is `MEMCACHE_1_5`.
	// Possible values are `MEMCACHE_1_5`.
	MemcacheVersion pulumi.StringPtrInput
	// The resource name of the instance.
	Name pulumi.StringPtrInput
	// Configuration for memcache nodes.
	// Structure is documented below.
	NodeConfig InstanceNodeConfigPtrInput
	// Number of nodes in the memcache instance.
	NodeCount pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the Memcache instance. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	Zones pulumi.StringArrayInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The full name of the GCE network to connect the instance to.  If not provided,
	// 'default' will be used.
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// A user-visible name for the instance.
	DisplayName *string `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// User-specified parameters for this memcache instance.
	// Structure is documented below.
	MemcacheParameters *InstanceMemcacheParameters `pulumi:"memcacheParameters"`
	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version.
	// Default value is `MEMCACHE_1_5`.
	// Possible values are `MEMCACHE_1_5`.
	MemcacheVersion *string `pulumi:"memcacheVersion"`
	// The resource name of the instance.
	Name *string `pulumi:"name"`
	// Configuration for memcache nodes.
	// Structure is documented below.
	NodeConfig InstanceNodeConfig `pulumi:"nodeConfig"`
	// Number of nodes in the memcache instance.
	NodeCount int `pulumi:"nodeCount"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the Memcache instance. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The full name of the GCE network to connect the instance to.  If not provided,
	// 'default' will be used.
	AuthorizedNetwork pulumi.StringPtrInput
	// A user-visible name for the instance.
	DisplayName pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// User-specified parameters for this memcache instance.
	// Structure is documented below.
	MemcacheParameters InstanceMemcacheParametersPtrInput
	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version.
	// Default value is `MEMCACHE_1_5`.
	// Possible values are `MEMCACHE_1_5`.
	MemcacheVersion pulumi.StringPtrInput
	// The resource name of the instance.
	Name pulumi.StringPtrInput
	// Configuration for memcache nodes.
	// Structure is documented below.
	NodeConfig InstanceNodeConfigInput
	// Number of nodes in the memcache instance.
	NodeCount pulumi.IntInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the Memcache instance. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	Zones pulumi.StringArrayInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

func (i *Instance) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *Instance) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

type InstancePtrInput interface {
	pulumi.Input

	ToInstancePtrOutput() InstancePtrOutput
	ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput
}

type instancePtrType InstanceArgs

func (*instancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (i *instancePtrType) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *instancePtrType) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Instance)(nil))
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Instance)(nil))
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct {
	*pulumi.OutputState
}

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o.ToInstancePtrOutputWithContext(context.Background())
}

func (o InstanceOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o.ApplyT(func(v Instance) *Instance {
		return &v
	}).(InstancePtrOutput)
}

type InstancePtrOutput struct {
	*pulumi.OutputState
}

func (InstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (o InstancePtrOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o
}

func (o InstancePtrOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Instance)(nil))
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Instance {
		return vs[0].([]Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Instance)(nil))
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Instance {
		return vs[0].(map[string]Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstancePtrOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}

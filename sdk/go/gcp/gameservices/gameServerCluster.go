// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gameservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A game server cluster resource.
//
// To get more information about GameServerCluster, see:
//
// * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.realms.gameServerClusters)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/game-servers/docs)
//
// ## Example Usage
//
// ## Import
//
// GameServerCluster can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default {{project}}/{{location}}/{{realm_id}}/{{cluster_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default {{location}}/{{realm_id}}/{{cluster_id}}
// ```
type GameServerCluster struct {
	pulumi.CustomResourceState

	// Required. The resource name of the game server cluster
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// Structure is documented below.
	ConnectionInfo GameServerClusterConnectionInfoOutput `pulumi:"connectionInfo"`
	// Human readable description of the cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location of the Cluster.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource id of the game server cluster, eg:
	// 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
	// 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The realm id of the game server realm.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewGameServerCluster registers a new resource with the given unique name, arguments, and options.
func NewGameServerCluster(ctx *pulumi.Context,
	name string, args *GameServerClusterArgs, opts ...pulumi.ResourceOption) (*GameServerCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.ConnectionInfo == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionInfo'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource GameServerCluster
	err := ctx.RegisterResource("gcp:gameservices/gameServerCluster:GameServerCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGameServerCluster gets an existing GameServerCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGameServerCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GameServerClusterState, opts ...pulumi.ResourceOption) (*GameServerCluster, error) {
	var resource GameServerCluster
	err := ctx.ReadResource("gcp:gameservices/gameServerCluster:GameServerCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GameServerCluster resources.
type gameServerClusterState struct {
	// Required. The resource name of the game server cluster
	ClusterId *string `pulumi:"clusterId"`
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// Structure is documented below.
	ConnectionInfo *GameServerClusterConnectionInfo `pulumi:"connectionInfo"`
	// Human readable description of the cluster.
	Description *string `pulumi:"description"`
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// Location of the Cluster.
	Location *string `pulumi:"location"`
	// The resource id of the game server cluster, eg:
	// 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
	// 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The realm id of the game server realm.
	RealmId *string `pulumi:"realmId"`
}

type GameServerClusterState struct {
	// Required. The resource name of the game server cluster
	ClusterId pulumi.StringPtrInput
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// Structure is documented below.
	ConnectionInfo GameServerClusterConnectionInfoPtrInput
	// Human readable description of the cluster.
	Description pulumi.StringPtrInput
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	Labels pulumi.StringMapInput
	// Location of the Cluster.
	Location pulumi.StringPtrInput
	// The resource id of the game server cluster, eg:
	// 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
	// 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The realm id of the game server realm.
	RealmId pulumi.StringPtrInput
}

func (GameServerClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerClusterState)(nil)).Elem()
}

type gameServerClusterArgs struct {
	// Required. The resource name of the game server cluster
	ClusterId string `pulumi:"clusterId"`
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// Structure is documented below.
	ConnectionInfo GameServerClusterConnectionInfo `pulumi:"connectionInfo"`
	// Human readable description of the cluster.
	Description *string `pulumi:"description"`
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// Location of the Cluster.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The realm id of the game server realm.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a GameServerCluster resource.
type GameServerClusterArgs struct {
	// Required. The resource name of the game server cluster
	ClusterId pulumi.StringInput
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// Structure is documented below.
	ConnectionInfo GameServerClusterConnectionInfoInput
	// Human readable description of the cluster.
	Description pulumi.StringPtrInput
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	Labels pulumi.StringMapInput
	// Location of the Cluster.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The realm id of the game server realm.
	RealmId pulumi.StringInput
}

func (GameServerClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerClusterArgs)(nil)).Elem()
}

type GameServerClusterInput interface {
	pulumi.Input

	ToGameServerClusterOutput() GameServerClusterOutput
	ToGameServerClusterOutputWithContext(ctx context.Context) GameServerClusterOutput
}

func (*GameServerCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**GameServerCluster)(nil)).Elem()
}

func (i *GameServerCluster) ToGameServerClusterOutput() GameServerClusterOutput {
	return i.ToGameServerClusterOutputWithContext(context.Background())
}

func (i *GameServerCluster) ToGameServerClusterOutputWithContext(ctx context.Context) GameServerClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerClusterOutput)
}

// GameServerClusterArrayInput is an input type that accepts GameServerClusterArray and GameServerClusterArrayOutput values.
// You can construct a concrete instance of `GameServerClusterArrayInput` via:
//
//          GameServerClusterArray{ GameServerClusterArgs{...} }
type GameServerClusterArrayInput interface {
	pulumi.Input

	ToGameServerClusterArrayOutput() GameServerClusterArrayOutput
	ToGameServerClusterArrayOutputWithContext(context.Context) GameServerClusterArrayOutput
}

type GameServerClusterArray []GameServerClusterInput

func (GameServerClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GameServerCluster)(nil)).Elem()
}

func (i GameServerClusterArray) ToGameServerClusterArrayOutput() GameServerClusterArrayOutput {
	return i.ToGameServerClusterArrayOutputWithContext(context.Background())
}

func (i GameServerClusterArray) ToGameServerClusterArrayOutputWithContext(ctx context.Context) GameServerClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerClusterArrayOutput)
}

// GameServerClusterMapInput is an input type that accepts GameServerClusterMap and GameServerClusterMapOutput values.
// You can construct a concrete instance of `GameServerClusterMapInput` via:
//
//          GameServerClusterMap{ "key": GameServerClusterArgs{...} }
type GameServerClusterMapInput interface {
	pulumi.Input

	ToGameServerClusterMapOutput() GameServerClusterMapOutput
	ToGameServerClusterMapOutputWithContext(context.Context) GameServerClusterMapOutput
}

type GameServerClusterMap map[string]GameServerClusterInput

func (GameServerClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GameServerCluster)(nil)).Elem()
}

func (i GameServerClusterMap) ToGameServerClusterMapOutput() GameServerClusterMapOutput {
	return i.ToGameServerClusterMapOutputWithContext(context.Background())
}

func (i GameServerClusterMap) ToGameServerClusterMapOutputWithContext(ctx context.Context) GameServerClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerClusterMapOutput)
}

type GameServerClusterOutput struct{ *pulumi.OutputState }

func (GameServerClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GameServerCluster)(nil)).Elem()
}

func (o GameServerClusterOutput) ToGameServerClusterOutput() GameServerClusterOutput {
	return o
}

func (o GameServerClusterOutput) ToGameServerClusterOutputWithContext(ctx context.Context) GameServerClusterOutput {
	return o
}

type GameServerClusterArrayOutput struct{ *pulumi.OutputState }

func (GameServerClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GameServerCluster)(nil)).Elem()
}

func (o GameServerClusterArrayOutput) ToGameServerClusterArrayOutput() GameServerClusterArrayOutput {
	return o
}

func (o GameServerClusterArrayOutput) ToGameServerClusterArrayOutputWithContext(ctx context.Context) GameServerClusterArrayOutput {
	return o
}

func (o GameServerClusterArrayOutput) Index(i pulumi.IntInput) GameServerClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GameServerCluster {
		return vs[0].([]*GameServerCluster)[vs[1].(int)]
	}).(GameServerClusterOutput)
}

type GameServerClusterMapOutput struct{ *pulumi.OutputState }

func (GameServerClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GameServerCluster)(nil)).Elem()
}

func (o GameServerClusterMapOutput) ToGameServerClusterMapOutput() GameServerClusterMapOutput {
	return o
}

func (o GameServerClusterMapOutput) ToGameServerClusterMapOutputWithContext(ctx context.Context) GameServerClusterMapOutput {
	return o
}

func (o GameServerClusterMapOutput) MapIndex(k pulumi.StringInput) GameServerClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GameServerCluster {
		return vs[0].(map[string]*GameServerCluster)[vs[1].(string)]
	}).(GameServerClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GameServerClusterInput)(nil)).Elem(), &GameServerCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*GameServerClusterArrayInput)(nil)).Elem(), GameServerClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GameServerClusterMapInput)(nil)).Elem(), GameServerClusterMap{})
	pulumi.RegisterOutputType(GameServerClusterOutput{})
	pulumi.RegisterOutputType(GameServerClusterArrayOutput{})
	pulumi.RegisterOutputType(GameServerClusterMapOutput{})
}

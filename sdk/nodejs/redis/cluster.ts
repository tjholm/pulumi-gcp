// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Google Cloud Redis Cluster instance.
 *
 * To get more information about Cluster, see:
 *
 * * [API documentation](https://cloud.google.com/memorystore/docs/cluster/reference/rest/v1/projects.locations.clusters)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/memorystore/docs/cluster/)
 *
 * ## Example Usage
 * ### Redis Cluster Ha
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const producerNet = new gcp.compute.Network("producerNet", {autoCreateSubnetworks: false});
 * const producerSubnet = new gcp.compute.Subnetwork("producerSubnet", {
 *     ipCidrRange: "10.0.0.248/29",
 *     region: "us-central1",
 *     network: producerNet.id,
 * });
 * const _default = new gcp.networkconnectivity.ServiceConnectionPolicy("default", {
 *     location: "us-central1",
 *     serviceClass: "gcp-memorystore-redis",
 *     description: "my basic service connection policy",
 *     network: producerNet.id,
 *     pscConfig: {
 *         subnetworks: [producerSubnet.id],
 *     },
 * });
 * const cluster_ha = new gcp.redis.Cluster("cluster-ha", {
 *     shardCount: 3,
 *     pscConfigs: [{
 *         network: producerNet.id,
 *     }],
 *     region: "us-central1",
 *     replicaCount: 1,
 *     transitEncryptionMode: "TRANSIT_ENCRYPTION_MODE_DISABLED",
 *     authorizationMode: "AUTH_MODE_DISABLED",
 * }, {
 *     dependsOn: [_default],
 * });
 * ```
 *
 * ## Import
 *
 * Cluster can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/clusters/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Cluster can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:redis/cluster:Cluster default projects/{{project}}/locations/{{region}}/clusters/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:redis/cluster:Cluster default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:redis/cluster:Cluster default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:redis/cluster:Cluster default {{name}}
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:redis/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster.
     * Default value is `AUTH_MODE_DISABLED`.
     * Possible values are: `AUTH_MODE_UNSPECIFIED`, `AUTH_MODE_IAM_AUTH`, `AUTH_MODE_DISABLED`.
     */
    public readonly authorizationMode!: pulumi.Output<string | undefined>;
    /**
     * The timestamp associated with the cluster creation request. A timestamp in
     * RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
     * digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Output only. Endpoints created on each given network,
     * for Redis clients to connect to the cluster.
     * Currently only one endpoint is supported.
     * Structure is documented below.
     */
    public /*out*/ readonly discoveryEndpoints!: pulumi.Output<outputs.redis.ClusterDiscoveryEndpoint[]>;
    /**
     * Unique name of the resource in this scope including project and location using the form:
     * projects/{projectId}/locations/{locationId}/clusters/{clusterId}
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Required. Each PscConfig configures the consumer network where two
     * network addresses will be designated to the cluster for client access.
     * Currently, only one PscConfig is supported.
     * Structure is documented below.
     */
    public readonly pscConfigs!: pulumi.Output<outputs.redis.ClusterPscConfig[]>;
    /**
     * Output only. PSC connections for discovery of the cluster topology and accessing the cluster.
     * Structure is documented below.
     */
    public /*out*/ readonly pscConnections!: pulumi.Output<outputs.redis.ClusterPscConnection[]>;
    /**
     * The name of the region of the Redis cluster.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Optional. The number of replica nodes per shard.
     */
    public readonly replicaCount!: pulumi.Output<number | undefined>;
    /**
     * Required. Number of shards for the Redis cluster.
     */
    public readonly shardCount!: pulumi.Output<number>;
    /**
     * Output only. Redis memory size in GB for the entire cluster.
     */
    public /*out*/ readonly sizeGb!: pulumi.Output<number>;
    /**
     * The current state of this cluster. Can be CREATING, READY, UPDATING, DELETING and SUSPENDED
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Output only. Additional information about the current state of the cluster.
     * Structure is documented below.
     */
    public /*out*/ readonly stateInfos!: pulumi.Output<outputs.redis.ClusterStateInfo[]>;
    /**
     * Optional. The in-transit encryption for the Redis cluster.
     * If not provided, encryption is disabled for the cluster.
     * Default value is `TRANSIT_ENCRYPTION_MODE_DISABLED`.
     * Possible values are: `TRANSIT_ENCRYPTION_MODE_UNSPECIFIED`, `TRANSIT_ENCRYPTION_MODE_DISABLED`, `TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION`.
     */
    public readonly transitEncryptionMode!: pulumi.Output<string | undefined>;
    /**
     * System assigned, unique identifier for the cluster.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["authorizationMode"] = state ? state.authorizationMode : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["discoveryEndpoints"] = state ? state.discoveryEndpoints : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pscConfigs"] = state ? state.pscConfigs : undefined;
            resourceInputs["pscConnections"] = state ? state.pscConnections : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["replicaCount"] = state ? state.replicaCount : undefined;
            resourceInputs["shardCount"] = state ? state.shardCount : undefined;
            resourceInputs["sizeGb"] = state ? state.sizeGb : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["stateInfos"] = state ? state.stateInfos : undefined;
            resourceInputs["transitEncryptionMode"] = state ? state.transitEncryptionMode : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.pscConfigs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pscConfigs'");
            }
            if ((!args || args.shardCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shardCount'");
            }
            resourceInputs["authorizationMode"] = args ? args.authorizationMode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pscConfigs"] = args ? args.pscConfigs : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["replicaCount"] = args ? args.replicaCount : undefined;
            resourceInputs["shardCount"] = args ? args.shardCount : undefined;
            resourceInputs["transitEncryptionMode"] = args ? args.transitEncryptionMode : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["discoveryEndpoints"] = undefined /*out*/;
            resourceInputs["pscConnections"] = undefined /*out*/;
            resourceInputs["sizeGb"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateInfos"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster.
     * Default value is `AUTH_MODE_DISABLED`.
     * Possible values are: `AUTH_MODE_UNSPECIFIED`, `AUTH_MODE_IAM_AUTH`, `AUTH_MODE_DISABLED`.
     */
    authorizationMode?: pulumi.Input<string>;
    /**
     * The timestamp associated with the cluster creation request. A timestamp in
     * RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
     * digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    createTime?: pulumi.Input<string>;
    /**
     * Output only. Endpoints created on each given network,
     * for Redis clients to connect to the cluster.
     * Currently only one endpoint is supported.
     * Structure is documented below.
     */
    discoveryEndpoints?: pulumi.Input<pulumi.Input<inputs.redis.ClusterDiscoveryEndpoint>[]>;
    /**
     * Unique name of the resource in this scope including project and location using the form:
     * projects/{projectId}/locations/{locationId}/clusters/{clusterId}
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Required. Each PscConfig configures the consumer network where two
     * network addresses will be designated to the cluster for client access.
     * Currently, only one PscConfig is supported.
     * Structure is documented below.
     */
    pscConfigs?: pulumi.Input<pulumi.Input<inputs.redis.ClusterPscConfig>[]>;
    /**
     * Output only. PSC connections for discovery of the cluster topology and accessing the cluster.
     * Structure is documented below.
     */
    pscConnections?: pulumi.Input<pulumi.Input<inputs.redis.ClusterPscConnection>[]>;
    /**
     * The name of the region of the Redis cluster.
     */
    region?: pulumi.Input<string>;
    /**
     * Optional. The number of replica nodes per shard.
     */
    replicaCount?: pulumi.Input<number>;
    /**
     * Required. Number of shards for the Redis cluster.
     */
    shardCount?: pulumi.Input<number>;
    /**
     * Output only. Redis memory size in GB for the entire cluster.
     */
    sizeGb?: pulumi.Input<number>;
    /**
     * The current state of this cluster. Can be CREATING, READY, UPDATING, DELETING and SUSPENDED
     */
    state?: pulumi.Input<string>;
    /**
     * Output only. Additional information about the current state of the cluster.
     * Structure is documented below.
     */
    stateInfos?: pulumi.Input<pulumi.Input<inputs.redis.ClusterStateInfo>[]>;
    /**
     * Optional. The in-transit encryption for the Redis cluster.
     * If not provided, encryption is disabled for the cluster.
     * Default value is `TRANSIT_ENCRYPTION_MODE_DISABLED`.
     * Possible values are: `TRANSIT_ENCRYPTION_MODE_UNSPECIFIED`, `TRANSIT_ENCRYPTION_MODE_DISABLED`, `TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION`.
     */
    transitEncryptionMode?: pulumi.Input<string>;
    /**
     * System assigned, unique identifier for the cluster.
     */
    uid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster.
     * Default value is `AUTH_MODE_DISABLED`.
     * Possible values are: `AUTH_MODE_UNSPECIFIED`, `AUTH_MODE_IAM_AUTH`, `AUTH_MODE_DISABLED`.
     */
    authorizationMode?: pulumi.Input<string>;
    /**
     * Unique name of the resource in this scope including project and location using the form:
     * projects/{projectId}/locations/{locationId}/clusters/{clusterId}
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Required. Each PscConfig configures the consumer network where two
     * network addresses will be designated to the cluster for client access.
     * Currently, only one PscConfig is supported.
     * Structure is documented below.
     */
    pscConfigs: pulumi.Input<pulumi.Input<inputs.redis.ClusterPscConfig>[]>;
    /**
     * The name of the region of the Redis cluster.
     */
    region?: pulumi.Input<string>;
    /**
     * Optional. The number of replica nodes per shard.
     */
    replicaCount?: pulumi.Input<number>;
    /**
     * Required. Number of shards for the Redis cluster.
     */
    shardCount: pulumi.Input<number>;
    /**
     * Optional. The in-transit encryption for the Redis cluster.
     * If not provided, encryption is disabled for the cluster.
     * Default value is `TRANSIT_ENCRYPTION_MODE_DISABLED`.
     * Possible values are: `TRANSIT_ENCRYPTION_MODE_UNSPECIFIED`, `TRANSIT_ENCRYPTION_MODE_DISABLED`, `TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION`.
     */
    transitEncryptionMode?: pulumi.Input<string>;
}

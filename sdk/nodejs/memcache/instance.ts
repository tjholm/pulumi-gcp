// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Google Cloud Memcache instance.
 *
 * To get more information about Instance, see:
 *
 * * [API documentation](https://cloud.google.com/memorystore/docs/memcached/reference/rest/v1/projects.locations.instances)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/memcache/docs/creating-instances)
 *
 * ## Example Usage
 * ### Memcache Instance Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * // This example assumes this network already exists.
 * // The API creates a tenant network per network authorized for a
 * // Memcache instance and that network is not deleted when the user-created
 * // network (authorized_network) is deleted, so this prevents issues
 * // with tenant network quota.
 * // If this network hasn't been created and you are using this example in your
 * // config, add an additional network resource or change
 * // this from "data"to "resource"
 * const memcacheNetwork = new gcp.compute.Network("memcacheNetwork", {});
 * const serviceRange = new gcp.compute.GlobalAddress("serviceRange", {
 *     purpose: "VPC_PEERING",
 *     addressType: "INTERNAL",
 *     prefixLength: 16,
 *     network: memcacheNetwork.id,
 * });
 * const privateServiceConnection = new gcp.servicenetworking.Connection("privateServiceConnection", {
 *     network: memcacheNetwork.id,
 *     service: "servicenetworking.googleapis.com",
 *     reservedPeeringRanges: [serviceRange.name],
 * });
 * const instance = new gcp.memcache.Instance("instance", {
 *     authorizedNetwork: privateServiceConnection.network,
 *     labels: {
 *         env: "test",
 *     },
 *     nodeConfig: {
 *         cpuCount: 1,
 *         memorySizeMb: 1024,
 *     },
 *     nodeCount: 1,
 *     memcacheVersion: "MEMCACHE_1_5",
 *     maintenancePolicy: {
 *         weeklyMaintenanceWindows: [{
 *             day: "SATURDAY",
 *             duration: "14400s",
 *             startTime: {
 *                 hours: 0,
 *                 minutes: 30,
 *                 seconds: 0,
 *                 nanos: 0,
 *             },
 *         }],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Instance can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/instances/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Instance can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:memcache/instance:Instance default projects/{{project}}/locations/{{region}}/instances/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:memcache/instance:Instance default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:memcache/instance:Instance default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:memcache/instance:Instance default {{name}}
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:memcache/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The full name of the GCE network to connect the instance to.  If not provided,
     * 'default' will be used.
     */
    public readonly authorizedNetwork!: pulumi.Output<string>;
    /**
     * (Output)
     * Output only. The time when the policy was created.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
     * resolution and up to nine fractional digits
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Endpoint for Discovery API
     */
    public /*out*/ readonly discoveryEndpoint!: pulumi.Output<string>;
    /**
     * A user-visible name for the instance.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource labels to represent user-provided metadata.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Maintenance policy for an instance.
     * Structure is documented below.
     */
    public readonly maintenancePolicy!: pulumi.Output<outputs.memcache.InstanceMaintenancePolicy | undefined>;
    /**
     * Output only. Published maintenance schedule.
     * Structure is documented below.
     */
    public /*out*/ readonly maintenanceSchedules!: pulumi.Output<outputs.memcache.InstanceMaintenanceSchedule[]>;
    /**
     * The full version of memcached server running on this instance.
     */
    public /*out*/ readonly memcacheFullVersion!: pulumi.Output<string>;
    /**
     * Additional information about the instance state, if available.
     * Structure is documented below.
     */
    public /*out*/ readonly memcacheNodes!: pulumi.Output<outputs.memcache.InstanceMemcacheNode[]>;
    /**
     * User-specified parameters for this memcache instance.
     * Structure is documented below.
     */
    public readonly memcacheParameters!: pulumi.Output<outputs.memcache.InstanceMemcacheParameters | undefined>;
    /**
     * The major version of Memcached software. If not provided, latest supported version will be used.
     * Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
     * determined by our system based on the latest supported minor version.
     * Default value is `MEMCACHE_1_5`.
     * Possible values are: `MEMCACHE_1_5`, `MEMCACHE_1_6_15`.
     */
    public readonly memcacheVersion!: pulumi.Output<string | undefined>;
    /**
     * The resource name of the instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration for memcache nodes.
     * Structure is documented below.
     */
    public readonly nodeConfig!: pulumi.Output<outputs.memcache.InstanceNodeConfig>;
    /**
     * Number of nodes in the memcache instance.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The region of the Memcache instance. If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Zones where memcache nodes should be provisioned.  If not
     * provided, all zones will be used.
     */
    public readonly zones!: pulumi.Output<string[]>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["authorizedNetwork"] = state ? state.authorizedNetwork : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["discoveryEndpoint"] = state ? state.discoveryEndpoint : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["maintenancePolicy"] = state ? state.maintenancePolicy : undefined;
            resourceInputs["maintenanceSchedules"] = state ? state.maintenanceSchedules : undefined;
            resourceInputs["memcacheFullVersion"] = state ? state.memcacheFullVersion : undefined;
            resourceInputs["memcacheNodes"] = state ? state.memcacheNodes : undefined;
            resourceInputs["memcacheParameters"] = state ? state.memcacheParameters : undefined;
            resourceInputs["memcacheVersion"] = state ? state.memcacheVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeConfig"] = state ? state.nodeConfig : undefined;
            resourceInputs["nodeCount"] = state ? state.nodeCount : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.nodeConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeConfig'");
            }
            if ((!args || args.nodeCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeCount'");
            }
            resourceInputs["authorizedNetwork"] = args ? args.authorizedNetwork : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["maintenancePolicy"] = args ? args.maintenancePolicy : undefined;
            resourceInputs["memcacheParameters"] = args ? args.memcacheParameters : undefined;
            resourceInputs["memcacheVersion"] = args ? args.memcacheVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeConfig"] = args ? args.nodeConfig : undefined;
            resourceInputs["nodeCount"] = args ? args.nodeCount : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["discoveryEndpoint"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["maintenanceSchedules"] = undefined /*out*/;
            resourceInputs["memcacheFullVersion"] = undefined /*out*/;
            resourceInputs["memcacheNodes"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The full name of the GCE network to connect the instance to.  If not provided,
     * 'default' will be used.
     */
    authorizedNetwork?: pulumi.Input<string>;
    /**
     * (Output)
     * Output only. The time when the policy was created.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
     * resolution and up to nine fractional digits
     */
    createTime?: pulumi.Input<string>;
    /**
     * Endpoint for Discovery API
     */
    discoveryEndpoint?: pulumi.Input<string>;
    /**
     * A user-visible name for the instance.
     */
    displayName?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource labels to represent user-provided metadata.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Maintenance policy for an instance.
     * Structure is documented below.
     */
    maintenancePolicy?: pulumi.Input<inputs.memcache.InstanceMaintenancePolicy>;
    /**
     * Output only. Published maintenance schedule.
     * Structure is documented below.
     */
    maintenanceSchedules?: pulumi.Input<pulumi.Input<inputs.memcache.InstanceMaintenanceSchedule>[]>;
    /**
     * The full version of memcached server running on this instance.
     */
    memcacheFullVersion?: pulumi.Input<string>;
    /**
     * Additional information about the instance state, if available.
     * Structure is documented below.
     */
    memcacheNodes?: pulumi.Input<pulumi.Input<inputs.memcache.InstanceMemcacheNode>[]>;
    /**
     * User-specified parameters for this memcache instance.
     * Structure is documented below.
     */
    memcacheParameters?: pulumi.Input<inputs.memcache.InstanceMemcacheParameters>;
    /**
     * The major version of Memcached software. If not provided, latest supported version will be used.
     * Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
     * determined by our system based on the latest supported minor version.
     * Default value is `MEMCACHE_1_5`.
     * Possible values are: `MEMCACHE_1_5`, `MEMCACHE_1_6_15`.
     */
    memcacheVersion?: pulumi.Input<string>;
    /**
     * The resource name of the instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration for memcache nodes.
     * Structure is documented below.
     */
    nodeConfig?: pulumi.Input<inputs.memcache.InstanceNodeConfig>;
    /**
     * Number of nodes in the memcache instance.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The region of the Memcache instance. If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * Zones where memcache nodes should be provisioned.  If not
     * provided, all zones will be used.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The full name of the GCE network to connect the instance to.  If not provided,
     * 'default' will be used.
     */
    authorizedNetwork?: pulumi.Input<string>;
    /**
     * A user-visible name for the instance.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Resource labels to represent user-provided metadata.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Maintenance policy for an instance.
     * Structure is documented below.
     */
    maintenancePolicy?: pulumi.Input<inputs.memcache.InstanceMaintenancePolicy>;
    /**
     * User-specified parameters for this memcache instance.
     * Structure is documented below.
     */
    memcacheParameters?: pulumi.Input<inputs.memcache.InstanceMemcacheParameters>;
    /**
     * The major version of Memcached software. If not provided, latest supported version will be used.
     * Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
     * determined by our system based on the latest supported minor version.
     * Default value is `MEMCACHE_1_5`.
     * Possible values are: `MEMCACHE_1_5`, `MEMCACHE_1_6_15`.
     */
    memcacheVersion?: pulumi.Input<string>;
    /**
     * The resource name of the instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration for memcache nodes.
     * Structure is documented below.
     */
    nodeConfig: pulumi.Input<inputs.memcache.InstanceNodeConfig>;
    /**
     * Number of nodes in the memcache instance.
     */
    nodeCount: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region of the Memcache instance. If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * Zones where memcache nodes should be provisioned.  If not
     * provided, all zones will be used.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}

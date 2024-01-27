// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A cluster in a private cloud.
 *
 * To get more information about Cluster, see:
 *
 * * [API documentation](https://cloud.google.com/vmware-engine/docs/reference/rest/v1/projects.locations.privateClouds.clusters)
 *
 * ## Example Usage
 * ### Vmware Engine Cluster Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cluster_nw = new gcp.vmwareengine.Network("cluster-nw", {
 *     type: "STANDARD",
 *     location: "global",
 *     description: "PC network description.",
 * });
 * const cluster_pc = new gcp.vmwareengine.PrivateCloud("cluster-pc", {
 *     location: "us-west1-a",
 *     description: "Sample test PC.",
 *     networkConfig: {
 *         managementCidr: "192.168.30.0/24",
 *         vmwareEngineNetwork: cluster_nw.id,
 *     },
 *     managementCluster: {
 *         clusterId: "sample-mgmt-cluster",
 *         nodeTypeConfigs: [{
 *             nodeTypeId: "standard-72",
 *             nodeCount: 3,
 *         }],
 *     },
 * });
 * const vmw_engine_ext_cluster = new gcp.vmwareengine.Cluster("vmw-engine-ext-cluster", {
 *     parent: cluster_pc.id,
 *     nodeTypeConfigs: [{
 *         nodeTypeId: "standard-72",
 *         nodeCount: 3,
 *     }],
 * });
 * ```
 * ### Vmware Engine Cluster Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cluster_nw = new gcp.vmwareengine.Network("cluster-nw", {
 *     type: "STANDARD",
 *     location: "global",
 *     description: "PC network description.",
 * });
 * const cluster_pc = new gcp.vmwareengine.PrivateCloud("cluster-pc", {
 *     location: "us-west1-a",
 *     description: "Sample test PC.",
 *     networkConfig: {
 *         managementCidr: "192.168.30.0/24",
 *         vmwareEngineNetwork: cluster_nw.id,
 *     },
 *     managementCluster: {
 *         clusterId: "sample-mgmt-cluster",
 *         nodeTypeConfigs: [{
 *             nodeTypeId: "standard-72",
 *             nodeCount: 3,
 *             customCoreCount: 32,
 *         }],
 *     },
 * });
 * const vmw_ext_cluster = new gcp.vmwareengine.Cluster("vmw-ext-cluster", {
 *     parent: cluster_pc.id,
 *     nodeTypeConfigs: [{
 *         nodeTypeId: "standard-72",
 *         nodeCount: 3,
 *         customCoreCount: 32,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Cluster can be imported using any of these accepted formats* `{{parent}}/clusters/{{name}}` When using the `pulumi import` command, Cluster can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:vmwareengine/cluster:Cluster default {{parent}}/clusters/{{name}}
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
    public static readonly __pulumiType = 'gcp:vmwareengine/cluster:Cluster';

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
     * True if the cluster is a management cluster; false otherwise.
     * There can only be one management cluster in a private cloud and it has to be the first one.
     */
    public /*out*/ readonly management!: pulumi.Output<boolean>;
    /**
     * The ID of the Cluster.
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The map of cluster node types in this cluster,
     * where the key is canonical identifier of the node type (corresponds to the NodeType).
     * Structure is documented below.
     */
    public readonly nodeTypeConfigs!: pulumi.Output<outputs.vmwareengine.ClusterNodeTypeConfig[] | undefined>;
    /**
     * The resource name of the private cloud to create a new cluster in.
     * Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
     * For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * State of the Cluster.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * System-generated unique identifier for the resource.
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
            resourceInputs["management"] = state ? state.management : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeTypeConfigs"] = state ? state.nodeTypeConfigs : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeTypeConfigs"] = args ? args.nodeTypeConfigs : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["management"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
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
     * True if the cluster is a management cluster; false otherwise.
     * There can only be one management cluster in a private cloud and it has to be the first one.
     */
    management?: pulumi.Input<boolean>;
    /**
     * The ID of the Cluster.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The map of cluster node types in this cluster,
     * where the key is canonical identifier of the node type (corresponds to the NodeType).
     * Structure is documented below.
     */
    nodeTypeConfigs?: pulumi.Input<pulumi.Input<inputs.vmwareengine.ClusterNodeTypeConfig>[]>;
    /**
     * The resource name of the private cloud to create a new cluster in.
     * Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
     * For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
     */
    parent?: pulumi.Input<string>;
    /**
     * State of the Cluster.
     */
    state?: pulumi.Input<string>;
    /**
     * System-generated unique identifier for the resource.
     */
    uid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The ID of the Cluster.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The map of cluster node types in this cluster,
     * where the key is canonical identifier of the node type (corresponds to the NodeType).
     * Structure is documented below.
     */
    nodeTypeConfigs?: pulumi.Input<pulumi.Input<inputs.vmwareengine.ClusterNodeTypeConfig>[]>;
    /**
     * The resource name of the private cloud to create a new cluster in.
     * Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
     * For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
     */
    parent: pulumi.Input<string>;
}

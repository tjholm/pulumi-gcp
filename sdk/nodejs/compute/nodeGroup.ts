// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Represents a NodeGroup resource to manage a group of sole-tenant nodes.
 *
 * To get more information about NodeGroup, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups)
 * * How-to Guides
 *     * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
 *
 * > **Warning:** Due to limitations of the API, this provider cannot update the
 * number of nodes in a node group and changes to node group size either
 * through provider config or through external changes will cause
 * the provider to delete and recreate the node group.
 *
 * ## Example Usage
 * ### Node Group Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const soletenant_tmpl = new gcp.compute.NodeTemplate("soletenant-tmpl", {
 *     region: "us-central1",
 *     nodeType: "n1-node-96-624",
 * });
 * const nodes = new gcp.compute.NodeGroup("nodes", {
 *     zone: "us-central1-a",
 *     description: "example google_compute_node_group for the Google Provider",
 *     size: 1,
 *     nodeTemplate: soletenant_tmpl.id,
 * });
 * ```
 * ### Node Group Autoscaling Policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const soletenant_tmpl = new gcp.compute.NodeTemplate("soletenant-tmpl", {
 *     region: "us-central1",
 *     nodeType: "n1-node-96-624",
 * });
 * const nodes = new gcp.compute.NodeGroup("nodes", {
 *     zone: "us-central1-a",
 *     description: "example google_compute_node_group for Google Provider",
 *     maintenancePolicy: "RESTART_IN_PLACE",
 *     maintenanceWindow: {
 *         startTime: "08:00",
 *     },
 *     initialSize: 1,
 *     nodeTemplate: soletenant_tmpl.id,
 *     autoscalingPolicy: {
 *         mode: "ONLY_SCALE_OUT",
 *         minNodes: 1,
 *         maxNodes: 10,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * NodeGroup can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/nodeGroup:NodeGroup default projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{project}}/{{zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{name}}
 * ```
 */
export class NodeGroup extends pulumi.CustomResource {
    /**
     * Get an existing NodeGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodeGroupState, opts?: pulumi.CustomResourceOptions): NodeGroup {
        return new NodeGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/nodeGroup:NodeGroup';

    /**
     * Returns true if the given object is an instance of NodeGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodeGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodeGroup.__pulumiType;
    }

    /**
     * If you use sole-tenant nodes for your workloads, you can use the node
     * group autoscaler to automatically manage the sizes of your node groups.
     * Structure is documented below.
     */
    public readonly autoscalingPolicy!: pulumi.Output<outputs.compute.NodeGroupAutoscalingPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional textual description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The initial number of nodes in the node group. One of `initialSize` or `size` must be specified.
     */
    public readonly initialSize!: pulumi.Output<number | undefined>;
    /**
     * Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
     */
    public readonly maintenancePolicy!: pulumi.Output<string | undefined>;
    /**
     * contains properties for the timeframe of maintenance
     * Structure is documented below.
     */
    public readonly maintenanceWindow!: pulumi.Output<outputs.compute.NodeGroupMaintenanceWindow | undefined>;
    /**
     * Name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL of the node template to which this node group belongs.
     */
    public readonly nodeTemplate!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The total number of nodes in the node group. One of `initialSize` or `size` must be specified.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * Zone where this node group is located
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a NodeGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodeGroupArgs | NodeGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodeGroupState | undefined;
            resourceInputs["autoscalingPolicy"] = state ? state.autoscalingPolicy : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["initialSize"] = state ? state.initialSize : undefined;
            resourceInputs["maintenancePolicy"] = state ? state.maintenancePolicy : undefined;
            resourceInputs["maintenanceWindow"] = state ? state.maintenanceWindow : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeTemplate"] = state ? state.nodeTemplate : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as NodeGroupArgs | undefined;
            if ((!args || args.nodeTemplate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeTemplate'");
            }
            resourceInputs["autoscalingPolicy"] = args ? args.autoscalingPolicy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["initialSize"] = args ? args.initialSize : undefined;
            resourceInputs["maintenancePolicy"] = args ? args.maintenancePolicy : undefined;
            resourceInputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeTemplate"] = args ? args.nodeTemplate : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NodeGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodeGroup resources.
 */
export interface NodeGroupState {
    /**
     * If you use sole-tenant nodes for your workloads, you can use the node
     * group autoscaler to automatically manage the sizes of your node groups.
     * Structure is documented below.
     */
    autoscalingPolicy?: pulumi.Input<inputs.compute.NodeGroupAutoscalingPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional textual description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The initial number of nodes in the node group. One of `initialSize` or `size` must be specified.
     */
    initialSize?: pulumi.Input<number>;
    /**
     * Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
     */
    maintenancePolicy?: pulumi.Input<string>;
    /**
     * contains properties for the timeframe of maintenance
     * Structure is documented below.
     */
    maintenanceWindow?: pulumi.Input<inputs.compute.NodeGroupMaintenanceWindow>;
    /**
     * Name of the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL of the node template to which this node group belongs.
     */
    nodeTemplate?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The total number of nodes in the node group. One of `initialSize` or `size` must be specified.
     */
    size?: pulumi.Input<number>;
    /**
     * Zone where this node group is located
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NodeGroup resource.
 */
export interface NodeGroupArgs {
    /**
     * If you use sole-tenant nodes for your workloads, you can use the node
     * group autoscaler to automatically manage the sizes of your node groups.
     * Structure is documented below.
     */
    autoscalingPolicy?: pulumi.Input<inputs.compute.NodeGroupAutoscalingPolicy>;
    /**
     * An optional textual description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The initial number of nodes in the node group. One of `initialSize` or `size` must be specified.
     */
    initialSize?: pulumi.Input<number>;
    /**
     * Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
     */
    maintenancePolicy?: pulumi.Input<string>;
    /**
     * contains properties for the timeframe of maintenance
     * Structure is documented below.
     */
    maintenanceWindow?: pulumi.Input<inputs.compute.NodeGroupMaintenanceWindow>;
    /**
     * Name of the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL of the node template to which this node group belongs.
     */
    nodeTemplate: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The total number of nodes in the node group. One of `initialSize` or `size` must be specified.
     */
    size?: pulumi.Input<number>;
    /**
     * Zone where this node group is located
     */
    zone?: pulumi.Input<string>;
}

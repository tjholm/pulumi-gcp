// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * "A set of Kubernetes nodes in a cluster with common configuration and specification."
 *
 * To get more information about NodePool, see:
 *
 * * [API documentation](https://cloud.google.com/distributed-cloud/edge/latest/docs/reference/container/rest/v1/projects.locations.clusters.nodePools)
 * * How-to Guides
 *     * [Google Distributed Cloud Edge](https://cloud.google.com/distributed-cloud/edge/latest/docs)
 *
 * ## Example Usage
 * ### Edgecontainer Node Pool
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = gcp.organizations.getProject({});
 * const cluster = new gcp.edgecontainer.Cluster("cluster", {
 *     location: "us-central1",
 *     authorization: {
 *         adminUsers: {
 *             username: "admin@hashicorptest.com",
 *         },
 *     },
 *     networking: {
 *         clusterIpv4CidrBlocks: ["10.0.0.0/16"],
 *         servicesIpv4CidrBlocks: ["10.1.0.0/16"],
 *     },
 *     fleet: {
 *         project: project.then(project => `projects/${project.number}`),
 *     },
 * });
 * const _default = new gcp.edgecontainer.NodePool("default", {
 *     cluster: cluster.name,
 *     location: "us-central1",
 *     nodeLocation: "us-central1-edge-example-edgesite",
 *     nodeCount: 3,
 *     labels: {
 *         my_key: "my_val",
 *         other_key: "other_val",
 *     },
 * });
 * ```
 * ### Edgecontainer Node Pool With Cmek
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = gcp.organizations.getProject({});
 * const cluster = new gcp.edgecontainer.Cluster("cluster", {
 *     location: "us-central1",
 *     authorization: {
 *         adminUsers: {
 *             username: "admin@hashicorptest.com",
 *         },
 *     },
 *     networking: {
 *         clusterIpv4CidrBlocks: ["10.0.0.0/16"],
 *         servicesIpv4CidrBlocks: ["10.1.0.0/16"],
 *     },
 *     fleet: {
 *         project: project.then(project => `projects/${project.number}`),
 *     },
 * });
 * const keyRing = new gcp.kms.KeyRing("keyRing", {location: "us-central1"});
 * const cryptoKeyCryptoKey = new gcp.kms.CryptoKey("cryptoKeyCryptoKey", {keyRing: keyRing.id});
 * const cryptoKeyCryptoKeyIAMMember = new gcp.kms.CryptoKeyIAMMember("cryptoKeyCryptoKeyIAMMember", {
 *     cryptoKeyId: cryptoKeyCryptoKey.id,
 *     role: "roles/cloudkms.cryptoKeyEncrypterDecrypter",
 *     member: project.then(project => `serviceAccount:service-${project.number}@gcp-sa-edgecontainer.iam.gserviceaccount.com`),
 * });
 * const _default = new gcp.edgecontainer.NodePool("default", {
 *     cluster: cluster.name,
 *     location: "us-central1",
 *     nodeLocation: "us-central1-edge-example-edgesite",
 *     nodeCount: 3,
 *     localDiskEncryption: {
 *         kmsKey: cryptoKeyCryptoKey.id,
 *     },
 * }, {
 *     dependsOn: [cryptoKeyCryptoKeyIAMMember],
 * });
 * ```
 * ### Edgecontainer Local Control Plane Node Pool
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = gcp.organizations.getProject({});
 * const defaultCluster = new gcp.edgecontainer.Cluster("defaultCluster", {
 *     location: "us-central1",
 *     authorization: {
 *         adminUsers: {
 *             username: "admin@hashicorptest.com",
 *         },
 *     },
 *     networking: {
 *         clusterIpv4CidrBlocks: ["10.0.0.0/16"],
 *         servicesIpv4CidrBlocks: ["10.1.0.0/16"],
 *     },
 *     fleet: {
 *         project: project.then(project => `projects/${project.number}`),
 *     },
 *     externalLoadBalancerIpv4AddressPools: ["10.100.0.0-10.100.0.10"],
 *     controlPlane: {
 *         local: {
 *             nodeLocation: "us-central1-edge-example-edgesite",
 *             nodeCount: 1,
 *             machineFilter: "machine-name",
 *             sharedDeploymentPolicy: "ALLOWED",
 *         },
 *     },
 * });
 * const defaultNodePool = new gcp.edgecontainer.NodePool("defaultNodePool", {
 *     cluster: google_edgecontainer_cluster.cluster.name,
 *     location: "us-central1",
 *     nodeLocation: "us-central1-edge-example-edgesite",
 *     nodeCount: 3,
 * });
 * ```
 *
 * ## Import
 *
 * NodePool can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}` * `{{project}}/{{location}}/{{cluster}}/{{name}}` * `{{location}}/{{cluster}}/{{name}}` When using the `pulumi import` command, NodePool can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:edgecontainer/nodePool:NodePool default projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:edgecontainer/nodePool:NodePool default {{project}}/{{location}}/{{cluster}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:edgecontainer/nodePool:NodePool default {{location}}/{{cluster}}/{{name}}
 * ```
 */
export class NodePool extends pulumi.CustomResource {
    /**
     * Get an existing NodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodePoolState, opts?: pulumi.CustomResourceOptions): NodePool {
        return new NodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:edgecontainer/nodePool:NodePool';

    /**
     * Returns true if the given object is an instance of NodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodePool.__pulumiType;
    }

    /**
     * The name of the target Distributed Cloud Edge Cluster.
     *
     *
     * - - -
     */
    public readonly cluster!: pulumi.Output<string>;
    /**
     * The time when the node pool was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Labels associated with this resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Local disk encryption options. This field is only used when enabling CMEK support.
     * Structure is documented below.
     */
    public readonly localDiskEncryption!: pulumi.Output<outputs.edgecontainer.NodePoolLocalDiskEncryption | undefined>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Only machines matching this filter will be allowed to join the node pool.
     * The filtering language accepts strings like "name=<name>", and is
     * documented in more detail in [AIP-160](https://google.aip.dev/160).
     */
    public readonly machineFilter!: pulumi.Output<string>;
    /**
     * The resource name of the node pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration for each node in the NodePool
     * Structure is documented below.
     */
    public readonly nodeConfig!: pulumi.Output<outputs.edgecontainer.NodePoolNodeConfig>;
    /**
     * The number of nodes in the pool.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * Name of the Google Distributed Cloud Edge zone where this node pool will be created. For example: `us-central1-edge-customer-a`.
     */
    public readonly nodeLocation!: pulumi.Output<string>;
    /**
     * The lowest release version among all worker nodes.
     */
    public /*out*/ readonly nodeVersion!: pulumi.Output<string>;
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
     * The time when the node pool was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a NodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodePoolArgs | NodePoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodePoolState | undefined;
            resourceInputs["cluster"] = state ? state.cluster : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["localDiskEncryption"] = state ? state.localDiskEncryption : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["machineFilter"] = state ? state.machineFilter : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeConfig"] = state ? state.nodeConfig : undefined;
            resourceInputs["nodeCount"] = state ? state.nodeCount : undefined;
            resourceInputs["nodeLocation"] = state ? state.nodeLocation : undefined;
            resourceInputs["nodeVersion"] = state ? state.nodeVersion : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as NodePoolArgs | undefined;
            if ((!args || args.cluster === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cluster'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.nodeCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeCount'");
            }
            if ((!args || args.nodeLocation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeLocation'");
            }
            resourceInputs["cluster"] = args ? args.cluster : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["localDiskEncryption"] = args ? args.localDiskEncryption : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["machineFilter"] = args ? args.machineFilter : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeConfig"] = args ? args.nodeConfig : undefined;
            resourceInputs["nodeCount"] = args ? args.nodeCount : undefined;
            resourceInputs["nodeLocation"] = args ? args.nodeLocation : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["nodeVersion"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(NodePool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodePool resources.
 */
export interface NodePoolState {
    /**
     * The name of the target Distributed Cloud Edge Cluster.
     *
     *
     * - - -
     */
    cluster?: pulumi.Input<string>;
    /**
     * The time when the node pool was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Labels associated with this resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Local disk encryption options. This field is only used when enabling CMEK support.
     * Structure is documented below.
     */
    localDiskEncryption?: pulumi.Input<inputs.edgecontainer.NodePoolLocalDiskEncryption>;
    /**
     * The location of the resource.
     */
    location?: pulumi.Input<string>;
    /**
     * Only machines matching this filter will be allowed to join the node pool.
     * The filtering language accepts strings like "name=<name>", and is
     * documented in more detail in [AIP-160](https://google.aip.dev/160).
     */
    machineFilter?: pulumi.Input<string>;
    /**
     * The resource name of the node pool.
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration for each node in the NodePool
     * Structure is documented below.
     */
    nodeConfig?: pulumi.Input<inputs.edgecontainer.NodePoolNodeConfig>;
    /**
     * The number of nodes in the pool.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * Name of the Google Distributed Cloud Edge zone where this node pool will be created. For example: `us-central1-edge-customer-a`.
     */
    nodeLocation?: pulumi.Input<string>;
    /**
     * The lowest release version among all worker nodes.
     */
    nodeVersion?: pulumi.Input<string>;
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
     * The time when the node pool was last updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NodePool resource.
 */
export interface NodePoolArgs {
    /**
     * The name of the target Distributed Cloud Edge Cluster.
     *
     *
     * - - -
     */
    cluster: pulumi.Input<string>;
    /**
     * Labels associated with this resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Local disk encryption options. This field is only used when enabling CMEK support.
     * Structure is documented below.
     */
    localDiskEncryption?: pulumi.Input<inputs.edgecontainer.NodePoolLocalDiskEncryption>;
    /**
     * The location of the resource.
     */
    location: pulumi.Input<string>;
    /**
     * Only machines matching this filter will be allowed to join the node pool.
     * The filtering language accepts strings like "name=<name>", and is
     * documented in more detail in [AIP-160](https://google.aip.dev/160).
     */
    machineFilter?: pulumi.Input<string>;
    /**
     * The resource name of the node pool.
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration for each node in the NodePool
     * Structure is documented below.
     */
    nodeConfig?: pulumi.Input<inputs.edgecontainer.NodePoolNodeConfig>;
    /**
     * The number of nodes in the pool.
     */
    nodeCount: pulumi.Input<number>;
    /**
     * Name of the Google Distributed Cloud Edge zone where this node pool will be created. For example: `us-central1-edge-customer-a`.
     */
    nodeLocation: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

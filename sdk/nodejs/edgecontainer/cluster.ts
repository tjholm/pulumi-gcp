// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Cluster contains information about a Google Distributed Cloud Edge Kubernetes cluster.
 *
 * To get more information about Cluster, see:
 *
 * * [API documentation](https://cloud.google.com/distributed-cloud/edge/latest/docs/reference/container/rest/v1/projects.locations.clusters)
 * * How-to Guides
 *     * [Create and manage clusters](https://cloud.google.com/distributed-cloud/edge/latest/docs/clusters)
 *
 * ## Example Usage
 * ### Edgecontainer Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = gcp.organizations.getProject({});
 * const _default = new gcp.edgecontainer.Cluster("default", {
 *     authorization: {
 *         adminUsers: {
 *             username: "admin@hashicorptest.com",
 *         },
 *     },
 *     fleet: {
 *         project: project.then(project => `projects/${project.number}`),
 *     },
 *     labels: {
 *         my_key: "my_val",
 *         other_key: "other_val",
 *     },
 *     location: "us-central1",
 *     networking: {
 *         clusterIpv4CidrBlocks: ["10.0.0.0/16"],
 *         servicesIpv4CidrBlocks: ["10.1.0.0/16"],
 *     },
 * });
 * ```
 * ### Edgecontainer Cluster With Maintenance Window
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = gcp.organizations.getProject({});
 * const _default = new gcp.edgecontainer.Cluster("default", {
 *     authorization: {
 *         adminUsers: {
 *             username: "admin@hashicorptest.com",
 *         },
 *     },
 *     fleet: {
 *         project: project.then(project => `projects/${project.number}`),
 *     },
 *     location: "us-central1",
 *     maintenancePolicy: {
 *         window: {
 *             recurringWindow: {
 *                 recurrence: "FREQ=WEEKLY;BYDAY=SA",
 *                 window: {
 *                     endTime: "2023-01-01T17:00:00Z",
 *                     startTime: "2023-01-01T08:00:00Z",
 *                 },
 *             },
 *         },
 *     },
 *     networking: {
 *         clusterIpv4CidrBlocks: ["10.0.0.0/16"],
 *         servicesIpv4CidrBlocks: ["10.1.0.0/16"],
 *     },
 * });
 * ```
 * ### Edgecontainer Local Control Plane Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = gcp.organizations.getProject({});
 * const _default = new gcp.edgecontainer.Cluster("default", {
 *     authorization: {
 *         adminUsers: {
 *             username: "admin@hashicorptest.com",
 *         },
 *     },
 *     controlPlane: {
 *         local: {
 *             machineFilter: "machine-name",
 *             nodeCount: 1,
 *             nodeLocation: "us-central1-edge-example-edgesite",
 *             sharedDeploymentPolicy: "ALLOWED",
 *         },
 *     },
 *     externalLoadBalancerIpv4AddressPools: ["10.100.0.0-10.100.0.10"],
 *     fleet: {
 *         project: project.then(project => `projects/${project.number}`),
 *     },
 *     location: "us-central1",
 *     networking: {
 *         clusterIpv4CidrBlocks: ["10.0.0.0/16"],
 *         servicesIpv4CidrBlocks: ["10.1.0.0/16"],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Cluster can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/clusters/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Cluster can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:edgecontainer/cluster:Cluster default projects/{{project}}/locations/{{location}}/clusters/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:edgecontainer/cluster:Cluster default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:edgecontainer/cluster:Cluster default {{location}}/{{name}}
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
    public static readonly __pulumiType = 'gcp:edgecontainer/cluster:Cluster';

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
     * RBAC policy that will be applied and managed by GEC.
     * Structure is documented below.
     */
    public readonly authorization!: pulumi.Output<outputs.edgecontainer.ClusterAuthorization>;
    /**
     * The PEM-encoded public certificate of the cluster's CA.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    public /*out*/ readonly clusterCaCertificate!: pulumi.Output<string>;
    /**
     * The configuration of the cluster control plane.
     * Structure is documented below.
     */
    public readonly controlPlane!: pulumi.Output<outputs.edgecontainer.ClusterControlPlane | undefined>;
    /**
     * Remote control plane disk encryption options. This field is only used when
     * enabling CMEK support.
     * Structure is documented below.
     */
    public readonly controlPlaneEncryption!: pulumi.Output<outputs.edgecontainer.ClusterControlPlaneEncryption>;
    /**
     * The control plane release version.
     */
    public /*out*/ readonly controlPlaneVersion!: pulumi.Output<string>;
    /**
     * (Output)
     * The time when the maintenance event request was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The default maximum number of pods per node used if a maximum value is not
     * specified explicitly for a node pool in this cluster. If unspecified, the
     * Kubernetes default value will be used.
     */
    public readonly defaultMaxPodsPerNode!: pulumi.Output<number>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The IP address of the Kubernetes API server.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Address pools for cluster data plane external load balancing.
     */
    public readonly externalLoadBalancerIpv4AddressPools!: pulumi.Output<string[]>;
    /**
     * Fleet related configuration.
     * Fleets are a Google Cloud concept for logically organizing clusters,
     * letting you use and manage multi-cluster capabilities and apply
     * consistent policies across your systems.
     * Structure is documented below.
     */
    public readonly fleet!: pulumi.Output<outputs.edgecontainer.ClusterFleet>;
    /**
     * User-defined labels for the edgecloud cluster.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * All the maintenance events scheduled for the cluster, including the ones
     * ongoing, planned for the future and done in the past (up to 90 days).
     * Structure is documented below.
     */
    public /*out*/ readonly maintenanceEvents!: pulumi.Output<outputs.edgecontainer.ClusterMaintenanceEvent[]>;
    /**
     * Cluster-wide maintenance policy configuration.
     * Structure is documented below.
     */
    public readonly maintenancePolicy!: pulumi.Output<outputs.edgecontainer.ClusterMaintenancePolicy>;
    /**
     * The GDCE cluster name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Fleet related configuration.
     * Fleets are a Google Cloud concept for logically organizing clusters,
     * letting you use and manage multi-cluster capabilities and apply
     * consistent policies across your systems.
     * Structure is documented below.
     */
    public readonly networking!: pulumi.Output<outputs.edgecontainer.ClusterNetworking>;
    /**
     * The lowest release version among all worker nodes. This field can be empty
     * if the cluster does not have any worker nodes.
     */
    public /*out*/ readonly nodeVersion!: pulumi.Output<string>;
    /**
     * The port number of the Kubernetes API server.
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
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
     * The release channel a cluster is subscribed to.
     * Possible values are: `RELEASE_CHANNEL_UNSPECIFIED`, `NONE`, `REGULAR`.
     */
    public readonly releaseChannel!: pulumi.Output<string>;
    /**
     * Indicates the status of the cluster.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Config that customers are allowed to define for GDCE system add-ons.
     * Structure is documented below.
     */
    public readonly systemAddonsConfig!: pulumi.Output<outputs.edgecontainer.ClusterSystemAddonsConfig>;
    /**
     * The target cluster version. For example: "1.5.0".
     */
    public readonly targetVersion!: pulumi.Output<string>;
    /**
     * (Output)
     * The time when the maintenance event message was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

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
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["clusterCaCertificate"] = state ? state.clusterCaCertificate : undefined;
            resourceInputs["controlPlane"] = state ? state.controlPlane : undefined;
            resourceInputs["controlPlaneEncryption"] = state ? state.controlPlaneEncryption : undefined;
            resourceInputs["controlPlaneVersion"] = state ? state.controlPlaneVersion : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["defaultMaxPodsPerNode"] = state ? state.defaultMaxPodsPerNode : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["externalLoadBalancerIpv4AddressPools"] = state ? state.externalLoadBalancerIpv4AddressPools : undefined;
            resourceInputs["fleet"] = state ? state.fleet : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["maintenanceEvents"] = state ? state.maintenanceEvents : undefined;
            resourceInputs["maintenancePolicy"] = state ? state.maintenancePolicy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networking"] = state ? state.networking : undefined;
            resourceInputs["nodeVersion"] = state ? state.nodeVersion : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["releaseChannel"] = state ? state.releaseChannel : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["systemAddonsConfig"] = state ? state.systemAddonsConfig : undefined;
            resourceInputs["targetVersion"] = state ? state.targetVersion : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.authorization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorization'");
            }
            if ((!args || args.fleet === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fleet'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.networking === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networking'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["controlPlane"] = args ? args.controlPlane : undefined;
            resourceInputs["controlPlaneEncryption"] = args ? args.controlPlaneEncryption : undefined;
            resourceInputs["defaultMaxPodsPerNode"] = args ? args.defaultMaxPodsPerNode : undefined;
            resourceInputs["externalLoadBalancerIpv4AddressPools"] = args ? args.externalLoadBalancerIpv4AddressPools : undefined;
            resourceInputs["fleet"] = args ? args.fleet : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["maintenancePolicy"] = args ? args.maintenancePolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networking"] = args ? args.networking : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["releaseChannel"] = args ? args.releaseChannel : undefined;
            resourceInputs["systemAddonsConfig"] = args ? args.systemAddonsConfig : undefined;
            resourceInputs["targetVersion"] = args ? args.targetVersion : undefined;
            resourceInputs["clusterCaCertificate"] = undefined /*out*/;
            resourceInputs["controlPlaneVersion"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["maintenanceEvents"] = undefined /*out*/;
            resourceInputs["nodeVersion"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clusterCaCertificate", "effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * RBAC policy that will be applied and managed by GEC.
     * Structure is documented below.
     */
    authorization?: pulumi.Input<inputs.edgecontainer.ClusterAuthorization>;
    /**
     * The PEM-encoded public certificate of the cluster's CA.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    clusterCaCertificate?: pulumi.Input<string>;
    /**
     * The configuration of the cluster control plane.
     * Structure is documented below.
     */
    controlPlane?: pulumi.Input<inputs.edgecontainer.ClusterControlPlane>;
    /**
     * Remote control plane disk encryption options. This field is only used when
     * enabling CMEK support.
     * Structure is documented below.
     */
    controlPlaneEncryption?: pulumi.Input<inputs.edgecontainer.ClusterControlPlaneEncryption>;
    /**
     * The control plane release version.
     */
    controlPlaneVersion?: pulumi.Input<string>;
    /**
     * (Output)
     * The time when the maintenance event request was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The default maximum number of pods per node used if a maximum value is not
     * specified explicitly for a node pool in this cluster. If unspecified, the
     * Kubernetes default value will be used.
     */
    defaultMaxPodsPerNode?: pulumi.Input<number>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The IP address of the Kubernetes API server.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Address pools for cluster data plane external load balancing.
     */
    externalLoadBalancerIpv4AddressPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Fleet related configuration.
     * Fleets are a Google Cloud concept for logically organizing clusters,
     * letting you use and manage multi-cluster capabilities and apply
     * consistent policies across your systems.
     * Structure is documented below.
     */
    fleet?: pulumi.Input<inputs.edgecontainer.ClusterFleet>;
    /**
     * User-defined labels for the edgecloud cluster.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location of the resource.
     */
    location?: pulumi.Input<string>;
    /**
     * All the maintenance events scheduled for the cluster, including the ones
     * ongoing, planned for the future and done in the past (up to 90 days).
     * Structure is documented below.
     */
    maintenanceEvents?: pulumi.Input<pulumi.Input<inputs.edgecontainer.ClusterMaintenanceEvent>[]>;
    /**
     * Cluster-wide maintenance policy configuration.
     * Structure is documented below.
     */
    maintenancePolicy?: pulumi.Input<inputs.edgecontainer.ClusterMaintenancePolicy>;
    /**
     * The GDCE cluster name.
     */
    name?: pulumi.Input<string>;
    /**
     * Fleet related configuration.
     * Fleets are a Google Cloud concept for logically organizing clusters,
     * letting you use and manage multi-cluster capabilities and apply
     * consistent policies across your systems.
     * Structure is documented below.
     */
    networking?: pulumi.Input<inputs.edgecontainer.ClusterNetworking>;
    /**
     * The lowest release version among all worker nodes. This field can be empty
     * if the cluster does not have any worker nodes.
     */
    nodeVersion?: pulumi.Input<string>;
    /**
     * The port number of the Kubernetes API server.
     */
    port?: pulumi.Input<number>;
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
     * The release channel a cluster is subscribed to.
     * Possible values are: `RELEASE_CHANNEL_UNSPECIFIED`, `NONE`, `REGULAR`.
     */
    releaseChannel?: pulumi.Input<string>;
    /**
     * Indicates the status of the cluster.
     */
    status?: pulumi.Input<string>;
    /**
     * Config that customers are allowed to define for GDCE system add-ons.
     * Structure is documented below.
     */
    systemAddonsConfig?: pulumi.Input<inputs.edgecontainer.ClusterSystemAddonsConfig>;
    /**
     * The target cluster version. For example: "1.5.0".
     */
    targetVersion?: pulumi.Input<string>;
    /**
     * (Output)
     * The time when the maintenance event message was updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * RBAC policy that will be applied and managed by GEC.
     * Structure is documented below.
     */
    authorization: pulumi.Input<inputs.edgecontainer.ClusterAuthorization>;
    /**
     * The configuration of the cluster control plane.
     * Structure is documented below.
     */
    controlPlane?: pulumi.Input<inputs.edgecontainer.ClusterControlPlane>;
    /**
     * Remote control plane disk encryption options. This field is only used when
     * enabling CMEK support.
     * Structure is documented below.
     */
    controlPlaneEncryption?: pulumi.Input<inputs.edgecontainer.ClusterControlPlaneEncryption>;
    /**
     * The default maximum number of pods per node used if a maximum value is not
     * specified explicitly for a node pool in this cluster. If unspecified, the
     * Kubernetes default value will be used.
     */
    defaultMaxPodsPerNode?: pulumi.Input<number>;
    /**
     * Address pools for cluster data plane external load balancing.
     */
    externalLoadBalancerIpv4AddressPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Fleet related configuration.
     * Fleets are a Google Cloud concept for logically organizing clusters,
     * letting you use and manage multi-cluster capabilities and apply
     * consistent policies across your systems.
     * Structure is documented below.
     */
    fleet: pulumi.Input<inputs.edgecontainer.ClusterFleet>;
    /**
     * User-defined labels for the edgecloud cluster.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location of the resource.
     */
    location: pulumi.Input<string>;
    /**
     * Cluster-wide maintenance policy configuration.
     * Structure is documented below.
     */
    maintenancePolicy?: pulumi.Input<inputs.edgecontainer.ClusterMaintenancePolicy>;
    /**
     * The GDCE cluster name.
     */
    name?: pulumi.Input<string>;
    /**
     * Fleet related configuration.
     * Fleets are a Google Cloud concept for logically organizing clusters,
     * letting you use and manage multi-cluster capabilities and apply
     * consistent policies across your systems.
     * Structure is documented below.
     */
    networking: pulumi.Input<inputs.edgecontainer.ClusterNetworking>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The release channel a cluster is subscribed to.
     * Possible values are: `RELEASE_CHANNEL_UNSPECIFIED`, `NONE`, `REGULAR`.
     */
    releaseChannel?: pulumi.Input<string>;
    /**
     * Config that customers are allowed to define for GDCE system add-ons.
     * Structure is documented below.
     */
    systemAddonsConfig?: pulumi.Input<inputs.edgecontainer.ClusterSystemAddonsConfig>;
    /**
     * The target cluster version. For example: "1.5.0".
     */
    targetVersion?: pulumi.Input<string>;
}

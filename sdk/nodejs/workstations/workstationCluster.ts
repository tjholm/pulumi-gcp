// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Workstation Cluster Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {autoCreateSubnetworks: false}, {
 *     provider: google_beta,
 * });
 * const defaultSubnetwork = new gcp.compute.Subnetwork("defaultSubnetwork", {
 *     ipCidrRange: "10.0.0.0/24",
 *     region: "us-central1",
 *     network: defaultNetwork.name,
 * }, {
 *     provider: google_beta,
 * });
 * const defaultWorkstationCluster = new gcp.workstations.WorkstationCluster("defaultWorkstationCluster", {
 *     workstationClusterId: "workstation-cluster",
 *     network: defaultNetwork.id,
 *     subnetwork: defaultSubnetwork.id,
 *     location: "us-central1",
 *     labels: {
 *         label: "key",
 *     },
 *     annotations: {
 *         "label-one": "value-one",
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const project = gcp.organizations.getProject({});
 * ```
 * ### Workstation Cluster Private
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {autoCreateSubnetworks: false}, {
 *     provider: google_beta,
 * });
 * const defaultSubnetwork = new gcp.compute.Subnetwork("defaultSubnetwork", {
 *     ipCidrRange: "10.0.0.0/24",
 *     region: "us-central1",
 *     network: defaultNetwork.name,
 * }, {
 *     provider: google_beta,
 * });
 * const defaultWorkstationCluster = new gcp.workstations.WorkstationCluster("defaultWorkstationCluster", {
 *     workstationClusterId: "workstation-cluster-private",
 *     network: defaultNetwork.id,
 *     subnetwork: defaultSubnetwork.id,
 *     location: "us-central1",
 *     privateClusterConfig: {
 *         enablePrivateEndpoint: true,
 *     },
 *     labels: {
 *         label: "key",
 *     },
 *     annotations: {
 *         "label-one": "value-one",
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const project = gcp.organizations.getProject({});
 * ```
 *
 * ## Import
 *
 * WorkstationCluster can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:workstations/workstationCluster:WorkstationCluster default projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:workstations/workstationCluster:WorkstationCluster default {{project}}/{{location}}/{{workstation_cluster_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:workstations/workstationCluster:WorkstationCluster default {{location}}/{{workstation_cluster_id}}
 * ```
 */
export class WorkstationCluster extends pulumi.CustomResource {
    /**
     * Get an existing WorkstationCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkstationClusterState, opts?: pulumi.CustomResourceOptions): WorkstationCluster {
        return new WorkstationCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:workstations/workstationCluster:WorkstationCluster';

    /**
     * Returns true if the given object is an instance of WorkstationCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkstationCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkstationCluster.__pulumiType;
    }

    /**
     * Client-specified annotations. This is distinct from labels.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Status conditions describing the current resource state.
     * Structure is documented below.
     */
    public /*out*/ readonly conditions!: pulumi.Output<outputs.workstations.WorkstationClusterCondition[]>;
    /**
     * Time when this resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Whether this resource is in degraded mode, in which case it may require user action to restore full functionality.
     * Details can be found in the conditions field.
     */
    public /*out*/ readonly degraded!: pulumi.Output<boolean>;
    /**
     * Human-readable name for this resource.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Checksum computed by the server.
     * May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location where the workstation cluster should reside.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the cluster resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The relative resource name of the VPC network on which the instance can be accessed.
     * It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Configuration for private cluster.
     * Structure is documented below.
     */
    public readonly privateClusterConfig!: pulumi.Output<outputs.workstations.WorkstationClusterPrivateClusterConfig | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
     * Must be part of the subnetwork specified for this cluster.
     */
    public readonly subnetwork!: pulumi.Output<string>;
    /**
     * The system-generated UID of the resource.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * ID to use for the workstation cluster.
     *
     *
     * - - -
     */
    public readonly workstationClusterId!: pulumi.Output<string>;

    /**
     * Create a WorkstationCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkstationClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkstationClusterArgs | WorkstationClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkstationClusterState | undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["degraded"] = state ? state.degraded : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["privateClusterConfig"] = state ? state.privateClusterConfig : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["subnetwork"] = state ? state.subnetwork : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["workstationClusterId"] = state ? state.workstationClusterId : undefined;
        } else {
            const args = argsOrState as WorkstationClusterArgs | undefined;
            if ((!args || args.network === undefined) && !opts.urn) {
                throw new Error("Missing required property 'network'");
            }
            if ((!args || args.subnetwork === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetwork'");
            }
            if ((!args || args.workstationClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workstationClusterId'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["privateClusterConfig"] = args ? args.privateClusterConfig : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["subnetwork"] = args ? args.subnetwork : undefined;
            resourceInputs["workstationClusterId"] = args ? args.workstationClusterId : undefined;
            resourceInputs["conditions"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["degraded"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WorkstationCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkstationCluster resources.
 */
export interface WorkstationClusterState {
    /**
     * Client-specified annotations. This is distinct from labels.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Status conditions describing the current resource state.
     * Structure is documented below.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.workstations.WorkstationClusterCondition>[]>;
    /**
     * Time when this resource was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Whether this resource is in degraded mode, in which case it may require user action to restore full functionality.
     * Details can be found in the conditions field.
     */
    degraded?: pulumi.Input<boolean>;
    /**
     * Human-readable name for this resource.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Checksum computed by the server.
     * May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location where the workstation cluster should reside.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the cluster resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The relative resource name of the VPC network on which the instance can be accessed.
     * It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
     */
    network?: pulumi.Input<string>;
    /**
     * Configuration for private cluster.
     * Structure is documented below.
     */
    privateClusterConfig?: pulumi.Input<inputs.workstations.WorkstationClusterPrivateClusterConfig>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
     * Must be part of the subnetwork specified for this cluster.
     */
    subnetwork?: pulumi.Input<string>;
    /**
     * The system-generated UID of the resource.
     */
    uid?: pulumi.Input<string>;
    /**
     * ID to use for the workstation cluster.
     *
     *
     * - - -
     */
    workstationClusterId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkstationCluster resource.
 */
export interface WorkstationClusterArgs {
    /**
     * Client-specified annotations. This is distinct from labels.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Human-readable name for this resource.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location where the workstation cluster should reside.
     */
    location?: pulumi.Input<string>;
    /**
     * The relative resource name of the VPC network on which the instance can be accessed.
     * It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
     */
    network: pulumi.Input<string>;
    /**
     * Configuration for private cluster.
     * Structure is documented below.
     */
    privateClusterConfig?: pulumi.Input<inputs.workstations.WorkstationClusterPrivateClusterConfig>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
     * Must be part of the subnetwork specified for this cluster.
     */
    subnetwork: pulumi.Input<string>;
    /**
     * ID to use for the workstation cluster.
     *
     *
     * - - -
     */
    workstationClusterId: pulumi.Input<string>;
}

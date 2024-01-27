// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * An Anthos node pool running on Azure.
 *
 * For more information, see:
 * * [Multicloud overview](https://cloud.google.com/anthos/clusters/docs/multi-cloud)
 * ## Example Usage
 * ### Basic_azure_node_pool
 * A basic example of a containerazure azure node pool
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const versions = gcp.container.getAzureVersions({
 *     project: "my-project-name",
 *     location: "us-west1",
 * });
 * const basic = new gcp.container.AzureClient("basic", {
 *     applicationId: "12345678-1234-1234-1234-123456789111",
 *     location: "us-west1",
 *     tenantId: "12345678-1234-1234-1234-123456789111",
 *     project: "my-project-name",
 * });
 * const primaryAzureCluster = new gcp.container.AzureCluster("primaryAzureCluster", {
 *     authorization: {
 *         adminUsers: [{
 *             username: "mmv2@google.com",
 *         }],
 *     },
 *     azureRegion: "westus2",
 *     client: pulumi.interpolate`projects/my-project-number/locations/us-west1/azureClients/${basic.name}`,
 *     controlPlane: {
 *         sshConfig: {
 *             authorizedKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8yaayO6lnb2v+SedxUMa2c8vtIEzCzBjM3EJJsv8Vm9zUDWR7dXWKoNGARUb2mNGXASvI6mFIDXTIlkQ0poDEPpMaXR0g2cb5xT8jAAJq7fqXL3+0rcJhY/uigQ+MrT6s+ub0BFVbsmGHNrMQttXX9gtmwkeAEvj3mra9e5pkNf90qlKnZz6U0SVArxVsLx07vHPHDIYrl0OPG4zUREF52igbBPiNrHJFDQJT/4YlDMJmo/QT/A1D6n9ocemvZSzhRx15/Arjowhr+VVKSbaxzPtEfY0oIg2SrqJnnr/l3Du5qIefwh5VmCZe4xopPUaDDoOIEFriZ88sB+3zz8ib8sk8zJJQCgeP78tQvXCgS+4e5W3TUg9mxjB6KjXTyHIVhDZqhqde0OI3Fy1UuVzRUwnBaLjBnAwP5EoFQGRmDYk/rEYe7HTmovLeEBUDQocBQKT4Ripm/xJkkWY7B07K/tfo56dGUCkvyIVXKBInCh+dLK7gZapnd4UWkY0xBYcwo1geMLRq58iFTLA2j/JmpmHXp7m0l7jJii7d44uD3tTIFYThn7NlOnvhLim/YcBK07GMGIN7XwrrKZKmxXaspw6KBWVhzuw1UPxctxshYEaMLfFg/bwOw8HvMPr9VtrElpSB7oiOh91PDIPdPBgHCi7N2QgQ5l/ZDBHieSpNrQ== thomasrodgers",
 *         },
 *         subnetId: "/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-byo/providers/Microsoft.Network/virtualNetworks/my--dev-vnet/subnets/default",
 *         version: versions.then(versions => versions.validVersions?.[0]),
 *     },
 *     fleet: {
 *         project: "my-project-number",
 *     },
 *     location: "us-west1",
 *     networking: {
 *         podAddressCidrBlocks: ["10.200.0.0/16"],
 *         serviceAddressCidrBlocks: ["10.32.0.0/24"],
 *         virtualNetworkId: "/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-byo/providers/Microsoft.Network/virtualNetworks/my--dev-vnet",
 *     },
 *     resourceGroupId: "/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-cluster",
 *     project: "my-project-name",
 * });
 * const primaryAzureNodePool = new gcp.container.AzureNodePool("primaryAzureNodePool", {
 *     autoscaling: {
 *         maxNodeCount: 3,
 *         minNodeCount: 2,
 *     },
 *     cluster: primaryAzureCluster.name,
 *     config: {
 *         sshConfig: {
 *             authorizedKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8yaayO6lnb2v+SedxUMa2c8vtIEzCzBjM3EJJsv8Vm9zUDWR7dXWKoNGARUb2mNGXASvI6mFIDXTIlkQ0poDEPpMaXR0g2cb5xT8jAAJq7fqXL3+0rcJhY/uigQ+MrT6s+ub0BFVbsmGHNrMQttXX9gtmwkeAEvj3mra9e5pkNf90qlKnZz6U0SVArxVsLx07vHPHDIYrl0OPG4zUREF52igbBPiNrHJFDQJT/4YlDMJmo/QT/A1D6n9ocemvZSzhRx15/Arjowhr+VVKSbaxzPtEfY0oIg2SrqJnnr/l3Du5qIefwh5VmCZe4xopPUaDDoOIEFriZ88sB+3zz8ib8sk8zJJQCgeP78tQvXCgS+4e5W3TUg9mxjB6KjXTyHIVhDZqhqde0OI3Fy1UuVzRUwnBaLjBnAwP5EoFQGRmDYk/rEYe7HTmovLeEBUDQocBQKT4Ripm/xJkkWY7B07K/tfo56dGUCkvyIVXKBInCh+dLK7gZapnd4UWkY0xBYcwo1geMLRq58iFTLA2j/JmpmHXp7m0l7jJii7d44uD3tTIFYThn7NlOnvhLim/YcBK07GMGIN7XwrrKZKmxXaspw6KBWVhzuw1UPxctxshYEaMLfFg/bwOw8HvMPr9VtrElpSB7oiOh91PDIPdPBgHCi7N2QgQ5l/ZDBHieSpNrQ== thomasrodgers",
 *         },
 *         proxyConfig: {
 *             resourceGroupId: "/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-cluster",
 *             secretId: "https://my--dev-keyvault.vault.azure.net/secrets/my--dev-secret/0000000000000000000000000000000000",
 *         },
 *         rootVolume: {
 *             sizeGib: 32,
 *         },
 *         tags: {
 *             owner: "mmv2",
 *         },
 *         labels: {
 *             key_one: "label_one",
 *         },
 *         vmSize: "Standard_DS2_v2",
 *     },
 *     location: "us-west1",
 *     maxPodsConstraint: {
 *         maxPodsPerNode: 110,
 *     },
 *     subnetId: "/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-byo/providers/Microsoft.Network/virtualNetworks/my--dev-vnet/subnets/default",
 *     version: versions.then(versions => versions.validVersions?.[0]),
 *     annotations: {
 *         "annotation-one": "value-one",
 *     },
 *     management: {
 *         autoRepair: true,
 *     },
 *     project: "my-project-name",
 * });
 * ```
 *
 * ## Import
 *
 * NodePool can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/azureClusters/{{cluster}}/azureNodePools/{{name}}` * `{{project}}/{{location}}/{{cluster}}/{{name}}` * `{{location}}/{{cluster}}/{{name}}` When using the `pulumi import` command, NodePool can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:container/azureNodePool:AzureNodePool default projects/{{project}}/locations/{{location}}/azureClusters/{{cluster}}/azureNodePools/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:container/azureNodePool:AzureNodePool default {{project}}/{{location}}/{{cluster}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:container/azureNodePool:AzureNodePool default {{location}}/{{cluster}}/{{name}}
 * ```
 */
export class AzureNodePool extends pulumi.CustomResource {
    /**
     * Get an existing AzureNodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureNodePoolState, opts?: pulumi.CustomResourceOptions): AzureNodePool {
        return new AzureNodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:container/azureNodePool:AzureNodePool';

    /**
     * Returns true if the given object is an instance of AzureNodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureNodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureNodePool.__pulumiType;
    }

    /**
     * Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
     *
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Autoscaler configuration for this node pool.
     */
    public readonly autoscaling!: pulumi.Output<outputs.container.AzureNodePoolAutoscaling>;
    /**
     * Optional. The Azure availability zone of the nodes in this nodepool. When unspecified, it defaults to `1`.
     */
    public readonly azureAvailabilityZone!: pulumi.Output<string>;
    /**
     * The azureCluster for the resource
     */
    public readonly cluster!: pulumi.Output<string>;
    /**
     * The node configuration of the node pool.
     */
    public readonly config!: pulumi.Output<outputs.container.AzureNodePoolConfig>;
    /**
     * Output only. The time at which this node pool was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
     * Terraform, other clients and services.
     */
    public /*out*/ readonly effectiveAnnotations!: pulumi.Output<{[key: string]: any}>;
    /**
     * Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The location for the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The Management configuration for this node pool.
     */
    public readonly management!: pulumi.Output<outputs.container.AzureNodePoolManagement>;
    /**
     * The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
     */
    public readonly maxPodsConstraint!: pulumi.Output<outputs.container.AzureNodePoolMaxPodsConstraint>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Output only. If set, there are currently pending changes to the node pool.
     */
    public /*out*/ readonly reconciling!: pulumi.Output<boolean>;
    /**
     * Output only. The current state of the node pool. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The ARM ID of the subnet where the node pool VMs run. Make sure it's a subnet under the virtual network in the cluster configuration.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Output only. A globally unique identifier for the node pool.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Output only. The time at which this node pool was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The Kubernetes version (e.g. `1.19.10-gke.1000`) running on this node pool.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a AzureNodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureNodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureNodePoolArgs | AzureNodePoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureNodePoolState | undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["autoscaling"] = state ? state.autoscaling : undefined;
            resourceInputs["azureAvailabilityZone"] = state ? state.azureAvailabilityZone : undefined;
            resourceInputs["cluster"] = state ? state.cluster : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["effectiveAnnotations"] = state ? state.effectiveAnnotations : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["management"] = state ? state.management : undefined;
            resourceInputs["maxPodsConstraint"] = state ? state.maxPodsConstraint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["reconciling"] = state ? state.reconciling : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AzureNodePoolArgs | undefined;
            if ((!args || args.autoscaling === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoscaling'");
            }
            if ((!args || args.cluster === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cluster'");
            }
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.maxPodsConstraint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxPodsConstraint'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["autoscaling"] = args ? args.autoscaling : undefined;
            resourceInputs["azureAvailabilityZone"] = args ? args.azureAvailabilityZone : undefined;
            resourceInputs["cluster"] = args ? args.cluster : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["management"] = args ? args.management : undefined;
            resourceInputs["maxPodsConstraint"] = args ? args.maxPodsConstraint : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveAnnotations"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["reconciling"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AzureNodePool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureNodePool resources.
 */
export interface AzureNodePoolState {
    /**
     * Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
     *
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Autoscaler configuration for this node pool.
     */
    autoscaling?: pulumi.Input<inputs.container.AzureNodePoolAutoscaling>;
    /**
     * Optional. The Azure availability zone of the nodes in this nodepool. When unspecified, it defaults to `1`.
     */
    azureAvailabilityZone?: pulumi.Input<string>;
    /**
     * The azureCluster for the resource
     */
    cluster?: pulumi.Input<string>;
    /**
     * The node configuration of the node pool.
     */
    config?: pulumi.Input<inputs.container.AzureNodePoolConfig>;
    /**
     * Output only. The time at which this node pool was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
     * Terraform, other clients and services.
     */
    effectiveAnnotations?: pulumi.Input<{[key: string]: any}>;
    /**
     * Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * The location for the resource
     */
    location?: pulumi.Input<string>;
    /**
     * The Management configuration for this node pool.
     */
    management?: pulumi.Input<inputs.container.AzureNodePoolManagement>;
    /**
     * The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
     */
    maxPodsConstraint?: pulumi.Input<inputs.container.AzureNodePoolMaxPodsConstraint>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Output only. If set, there are currently pending changes to the node pool.
     */
    reconciling?: pulumi.Input<boolean>;
    /**
     * Output only. The current state of the node pool. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
     */
    state?: pulumi.Input<string>;
    /**
     * The ARM ID of the subnet where the node pool VMs run. Make sure it's a subnet under the virtual network in the cluster configuration.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Output only. A globally unique identifier for the node pool.
     */
    uid?: pulumi.Input<string>;
    /**
     * Output only. The time at which this node pool was last updated.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * The Kubernetes version (e.g. `1.19.10-gke.1000`) running on this node pool.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AzureNodePool resource.
 */
export interface AzureNodePoolArgs {
    /**
     * Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
     *
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Autoscaler configuration for this node pool.
     */
    autoscaling: pulumi.Input<inputs.container.AzureNodePoolAutoscaling>;
    /**
     * Optional. The Azure availability zone of the nodes in this nodepool. When unspecified, it defaults to `1`.
     */
    azureAvailabilityZone?: pulumi.Input<string>;
    /**
     * The azureCluster for the resource
     */
    cluster: pulumi.Input<string>;
    /**
     * The node configuration of the node pool.
     */
    config: pulumi.Input<inputs.container.AzureNodePoolConfig>;
    /**
     * The location for the resource
     */
    location: pulumi.Input<string>;
    /**
     * The Management configuration for this node pool.
     */
    management?: pulumi.Input<inputs.container.AzureNodePoolManagement>;
    /**
     * The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
     */
    maxPodsConstraint: pulumi.Input<inputs.container.AzureNodePoolMaxPodsConstraint>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * The ARM ID of the subnet where the node pool VMs run. Make sure it's a subnet under the virtual network in the cluster configuration.
     */
    subnetId: pulumi.Input<string>;
    /**
     * The Kubernetes version (e.g. `1.19.10-gke.1000`) running on this node pool.
     */
    version: pulumi.Input<string>;
}

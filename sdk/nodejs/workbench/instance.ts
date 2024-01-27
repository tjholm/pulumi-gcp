// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Workbench instance.
 *
 * ## Example Usage
 * ### Workbench Instance Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.workbench.Instance("instance", {location: "us-west1-a"});
 * ```
 * ### Workbench Instance Basic Gpu
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.workbench.Instance("instance", {
 *     gceSetup: {
 *         acceleratorConfigs: [{
 *             coreCount: "1",
 *             type: "NVIDIA_TESLA_T4",
 *         }],
 *         machineType: "n1-standard-1",
 *         vmImage: {
 *             family: "tf-latest-gpu",
 *             project: "deeplearning-platform-release",
 *         },
 *     },
 *     location: "us-central1-a",
 * });
 * ```
 * ### Workbench Instance Labels
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.workbench.Instance("instance", {
 *     gceSetup: {
 *         machineType: "e2-standard-4",
 *         metadata: {
 *             terraform: "true",
 *         },
 *         serviceAccounts: [{
 *             email: "my@service-account.com",
 *         }],
 *     },
 *     instanceOwners: ["my@service-account.com"],
 *     labels: {
 *         k: "val",
 *     },
 *     location: "us-central1-a",
 * });
 * ```
 * ### Workbench Instance Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myNetwork = new gcp.compute.Network("myNetwork", {autoCreateSubnetworks: false});
 * const mySubnetwork = new gcp.compute.Subnetwork("mySubnetwork", {
 *     network: myNetwork.id,
 *     region: "us-central1",
 *     ipCidrRange: "10.0.1.0/24",
 * });
 * const keyring = new gcp.kms.KeyRing("keyring", {location: "global"});
 * const crypto_key = new gcp.kms.CryptoKey("crypto-key", {keyRing: keyring.id});
 * const instance = new gcp.workbench.Instance("instance", {
 *     location: "us-central1-a",
 *     gceSetup: {
 *         machineType: "n1-standard-4",
 *         acceleratorConfigs: [{
 *             type: "NVIDIA_TESLA_T4",
 *             coreCount: "1",
 *         }],
 *         disablePublicIp: false,
 *         serviceAccounts: [{
 *             email: "my@service-account.com",
 *         }],
 *         bootDisk: {
 *             diskSizeGb: "310",
 *             diskType: "PD_SSD",
 *             diskEncryption: "GMEK",
 *             kmsKey: crypto_key.id,
 *         },
 *         dataDisks: {
 *             diskSizeGb: "330",
 *             diskType: "PD_SSD",
 *             diskEncryption: "GMEK",
 *             kmsKey: crypto_key.id,
 *         },
 *         networkInterfaces: [{
 *             network: myNetwork.id,
 *             subnet: mySubnetwork.id,
 *             nicType: "GVNIC",
 *         }],
 *         metadata: {
 *             terraform: "true",
 *         },
 *         enableIpForwarding: true,
 *         tags: [
 *             "abc",
 *             "def",
 *         ],
 *     },
 *     disableProxyAccess: true,
 *     instanceOwners: ["my@service-account.com"],
 *     labels: {
 *         k: "val",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Instance can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/instances/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Instance can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:workbench/instance:Instance default projects/{{project}}/locations/{{location}}/instances/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:workbench/instance:Instance default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:workbench/instance:Instance default {{location}}/{{name}}
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
    public static readonly __pulumiType = 'gcp:workbench/instance:Instance';

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
     * An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ.
     * The milliseconds portion (".SSS") is optional.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Output only. Email address of entity that sent original CreateInstance request.
     */
    public /*out*/ readonly creator!: pulumi.Output<string>;
    /**
     * Optional. If true, the workbench instance will not register with the proxy.
     */
    public readonly disableProxyAccess!: pulumi.Output<boolean | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The definition of how to configure a VM instance outside of Resources and Identity.
     * Structure is documented below.
     */
    public readonly gceSetup!: pulumi.Output<outputs.workbench.InstanceGceSetup>;
    /**
     * 'Output only. Additional information about instance health. Example:
     * healthInfo": { "dockerProxyAgentStatus": "1", "dockerStatus": "1", "jupyterlabApiStatus":
     * "-1", "jupyterlabStatus": "-1", "updated": "2020-10-18 09:40:03.573409" }'
     */
    public /*out*/ readonly healthInfos!: pulumi.Output<outputs.workbench.InstanceHealthInfo[]>;
    /**
     * Output only. Instance health_state.
     */
    public /*out*/ readonly healthState!: pulumi.Output<string>;
    /**
     * Required. User-defined unique ID of this instance.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;
    /**
     * 'Optional. Input only. The owner of this instance after creation. Format:
     * `alias@example.com` Currently supports one owner only. If not specified, all of
     * the service account users of your VM instance''s service account can use the instance.'
     */
    public readonly instanceOwners!: pulumi.Output<string[] | undefined>;
    /**
     * Optional. Labels to apply to this instance. These can be later modified
     * by the UpdateInstance method.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Part of `parent`. See documentation of `projectsId`.
     *
     *
     * - - -
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of this workbench instance. Format: `projects/{project_id}/locations/{location}/instances/{instance_id}`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Output only. The proxy endpoint that is used to access the Jupyter notebook.
     */
    public /*out*/ readonly proxyUri!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * (Output)
     * Output only. The state of this instance upgrade history entry.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ.
     * The milliseconds portion (".SSS") is optional.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Output only. The upgrade history of this instance.
     * Structure is documented below.
     */
    public /*out*/ readonly upgradeHistories!: pulumi.Output<outputs.workbench.InstanceUpgradeHistory[]>;

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
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["creator"] = state ? state.creator : undefined;
            resourceInputs["disableProxyAccess"] = state ? state.disableProxyAccess : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["gceSetup"] = state ? state.gceSetup : undefined;
            resourceInputs["healthInfos"] = state ? state.healthInfos : undefined;
            resourceInputs["healthState"] = state ? state.healthState : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceOwners"] = state ? state.instanceOwners : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["proxyUri"] = state ? state.proxyUri : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["upgradeHistories"] = state ? state.upgradeHistories : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["disableProxyAccess"] = args ? args.disableProxyAccess : undefined;
            resourceInputs["gceSetup"] = args ? args.gceSetup : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["instanceOwners"] = args ? args.instanceOwners : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["creator"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["healthInfos"] = undefined /*out*/;
            resourceInputs["healthState"] = undefined /*out*/;
            resourceInputs["proxyUri"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["upgradeHistories"] = undefined /*out*/;
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
     * An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ.
     * The milliseconds portion (".SSS") is optional.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Output only. Email address of entity that sent original CreateInstance request.
     */
    creator?: pulumi.Input<string>;
    /**
     * Optional. If true, the workbench instance will not register with the proxy.
     */
    disableProxyAccess?: pulumi.Input<boolean>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The definition of how to configure a VM instance outside of Resources and Identity.
     * Structure is documented below.
     */
    gceSetup?: pulumi.Input<inputs.workbench.InstanceGceSetup>;
    /**
     * 'Output only. Additional information about instance health. Example:
     * healthInfo": { "dockerProxyAgentStatus": "1", "dockerStatus": "1", "jupyterlabApiStatus":
     * "-1", "jupyterlabStatus": "-1", "updated": "2020-10-18 09:40:03.573409" }'
     */
    healthInfos?: pulumi.Input<pulumi.Input<inputs.workbench.InstanceHealthInfo>[]>;
    /**
     * Output only. Instance health_state.
     */
    healthState?: pulumi.Input<string>;
    /**
     * Required. User-defined unique ID of this instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * 'Optional. Input only. The owner of this instance after creation. Format:
     * `alias@example.com` Currently supports one owner only. If not specified, all of
     * the service account users of your VM instance''s service account can use the instance.'
     */
    instanceOwners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. Labels to apply to this instance. These can be later modified
     * by the UpdateInstance method.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Part of `parent`. See documentation of `projectsId`.
     *
     *
     * - - -
     */
    location?: pulumi.Input<string>;
    /**
     * The name of this workbench instance. Format: `projects/{project_id}/locations/{location}/instances/{instance_id}`
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Output only. The proxy endpoint that is used to access the Jupyter notebook.
     */
    proxyUri?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * (Output)
     * Output only. The state of this instance upgrade history entry.
     */
    state?: pulumi.Input<string>;
    /**
     * An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ.
     * The milliseconds portion (".SSS") is optional.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * Output only. The upgrade history of this instance.
     * Structure is documented below.
     */
    upgradeHistories?: pulumi.Input<pulumi.Input<inputs.workbench.InstanceUpgradeHistory>[]>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Optional. If true, the workbench instance will not register with the proxy.
     */
    disableProxyAccess?: pulumi.Input<boolean>;
    /**
     * The definition of how to configure a VM instance outside of Resources and Identity.
     * Structure is documented below.
     */
    gceSetup?: pulumi.Input<inputs.workbench.InstanceGceSetup>;
    /**
     * Required. User-defined unique ID of this instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * 'Optional. Input only. The owner of this instance after creation. Format:
     * `alias@example.com` Currently supports one owner only. If not specified, all of
     * the service account users of your VM instance''s service account can use the instance.'
     */
    instanceOwners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. Labels to apply to this instance. These can be later modified
     * by the UpdateInstance method.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Part of `parent`. See documentation of `projectsId`.
     *
     *
     * - - -
     */
    location: pulumi.Input<string>;
    /**
     * The name of this workbench instance. Format: `projects/{project_id}/locations/{location}/instances/{instance_id}`
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

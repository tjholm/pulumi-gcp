// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Workstation Config Basic
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
 * const defaultWorkstationConfig = new gcp.workstations.WorkstationConfig("defaultWorkstationConfig", {
 *     workstationConfigId: "workstation-config",
 *     workstationClusterId: defaultWorkstationCluster.workstationClusterId,
 *     location: "us-central1",
 *     idleTimeout: "600s",
 *     runningTimeout: "21600s",
 *     replicaZones: [
 *         "us-central1-a",
 *         "us-central1-b",
 *     ],
 *     annotations: {
 *         "label-one": "value-one",
 *     },
 *     labels: {
 *         label: "key",
 *     },
 *     host: {
 *         gceInstance: {
 *             machineType: "e2-standard-4",
 *             bootDiskSizeGb: 35,
 *             disablePublicIpAddresses: true,
 *         },
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Workstation Config Container
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
 * const defaultWorkstationConfig = new gcp.workstations.WorkstationConfig("defaultWorkstationConfig", {
 *     workstationConfigId: "workstation-config",
 *     workstationClusterId: defaultWorkstationCluster.workstationClusterId,
 *     location: "us-central1",
 *     host: {
 *         gceInstance: {
 *             machineType: "n1-standard-4",
 *             bootDiskSizeGb: 35,
 *             disablePublicIpAddresses: true,
 *             enableNestedVirtualization: true,
 *         },
 *     },
 *     container: {
 *         image: "intellij",
 *         env: {
 *             NAME: "FOO",
 *             BABE: "bar",
 *         },
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Workstation Config Persistent Directories
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
 * const defaultWorkstationConfig = new gcp.workstations.WorkstationConfig("defaultWorkstationConfig", {
 *     workstationConfigId: "workstation-config",
 *     workstationClusterId: defaultWorkstationCluster.workstationClusterId,
 *     location: "us-central1",
 *     host: {
 *         gceInstance: {
 *             machineType: "e2-standard-4",
 *             bootDiskSizeGb: 35,
 *             disablePublicIpAddresses: true,
 *             shieldedInstanceConfig: {
 *                 enableSecureBoot: true,
 *                 enableVtpm: true,
 *             },
 *         },
 *     },
 *     persistentDirectories: [{
 *         mountPath: "/home",
 *         gcePd: {
 *             sizeGb: 200,
 *             fsType: "ext4",
 *             diskType: "pd-standard",
 *             reclaimPolicy: "DELETE",
 *         },
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Workstation Config Source Snapshot
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
 * const mySourceDisk = new gcp.compute.Disk("mySourceDisk", {
 *     size: 10,
 *     type: "pd-ssd",
 *     zone: "us-central1-a",
 * }, {
 *     provider: google_beta,
 * });
 * const mySourceSnapshot = new gcp.compute.Snapshot("mySourceSnapshot", {
 *     sourceDisk: mySourceDisk.name,
 *     zone: "us-central1-a",
 * }, {
 *     provider: google_beta,
 * });
 * const defaultWorkstationCluster = new gcp.workstations.WorkstationCluster("defaultWorkstationCluster", {
 *     workstationClusterId: "workstation-cluster",
 *     network: defaultNetwork.id,
 *     subnetwork: defaultSubnetwork.id,
 *     location: "us-central1",
 * }, {
 *     provider: google_beta,
 * });
 * const defaultWorkstationConfig = new gcp.workstations.WorkstationConfig("defaultWorkstationConfig", {
 *     workstationConfigId: "workstation-config",
 *     workstationClusterId: defaultWorkstationCluster.workstationClusterId,
 *     location: defaultWorkstationCluster.location,
 *     persistentDirectories: [{
 *         mountPath: "/home",
 *         gcePd: {
 *             sourceSnapshot: mySourceSnapshot.id,
 *             reclaimPolicy: "DELETE",
 *         },
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Workstation Config Shielded Instance Config
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
 * const defaultWorkstationConfig = new gcp.workstations.WorkstationConfig("defaultWorkstationConfig", {
 *     workstationConfigId: "workstation-config",
 *     workstationClusterId: defaultWorkstationCluster.workstationClusterId,
 *     location: "us-central1",
 *     host: {
 *         gceInstance: {
 *             machineType: "e2-standard-4",
 *             bootDiskSizeGb: 35,
 *             disablePublicIpAddresses: true,
 *             shieldedInstanceConfig: {
 *                 enableSecureBoot: true,
 *                 enableVtpm: true,
 *             },
 *         },
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Workstation Config Accelerators
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
 * const defaultWorkstationConfig = new gcp.workstations.WorkstationConfig("defaultWorkstationConfig", {
 *     workstationConfigId: "workstation-config",
 *     workstationClusterId: defaultWorkstationCluster.workstationClusterId,
 *     location: "us-central1",
 *     host: {
 *         gceInstance: {
 *             machineType: "n1-standard-2",
 *             bootDiskSizeGb: 35,
 *             disablePublicIpAddresses: true,
 *             accelerators: [{
 *                 type: "nvidia-tesla-p100",
 *                 count: 1,
 *             }],
 *         },
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Workstation Config Encryption Key
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
 * const defaultKeyRing = new gcp.kms.KeyRing("defaultKeyRing", {location: "us-central1"}, {
 *     provider: google_beta,
 * });
 * const defaultCryptoKey = new gcp.kms.CryptoKey("defaultCryptoKey", {keyRing: defaultKeyRing.id}, {
 *     provider: google_beta,
 * });
 * const defaultAccount = new gcp.serviceaccount.Account("defaultAccount", {
 *     accountId: "my-account",
 *     displayName: "Service Account",
 * }, {
 *     provider: google_beta,
 * });
 * const defaultWorkstationConfig = new gcp.workstations.WorkstationConfig("defaultWorkstationConfig", {
 *     workstationConfigId: "workstation-config",
 *     workstationClusterId: defaultWorkstationCluster.workstationClusterId,
 *     location: "us-central1",
 *     host: {
 *         gceInstance: {
 *             machineType: "e2-standard-4",
 *             bootDiskSizeGb: 35,
 *             disablePublicIpAddresses: true,
 *             shieldedInstanceConfig: {
 *                 enableSecureBoot: true,
 *                 enableVtpm: true,
 *             },
 *         },
 *     },
 *     encryptionKey: {
 *         kmsKey: defaultCryptoKey.id,
 *         kmsKeyServiceAccount: defaultAccount.email,
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * WorkstationConfig can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:workstations/workstationConfig:WorkstationConfig default projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:workstations/workstationConfig:WorkstationConfig default {{project}}/{{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:workstations/workstationConfig:WorkstationConfig default {{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}
 * ```
 */
export class WorkstationConfig extends pulumi.CustomResource {
    /**
     * Get an existing WorkstationConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkstationConfigState, opts?: pulumi.CustomResourceOptions): WorkstationConfig {
        return new WorkstationConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:workstations/workstationConfig:WorkstationConfig';

    /**
     * Returns true if the given object is an instance of WorkstationConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkstationConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkstationConfig.__pulumiType;
    }

    /**
     * Client-specified annotations. This is distinct from labels.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Status conditions describing the current resource state.
     * Structure is documented below.
     */
    public /*out*/ readonly conditions!: pulumi.Output<outputs.workstations.WorkstationConfigCondition[]>;
    /**
     * Container that will be run for each workstation using this configuration when that workstation is started.
     * Structure is documented below.
     */
    public readonly container!: pulumi.Output<outputs.workstations.WorkstationConfigContainer>;
    /**
     * Time when this resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Whether this resource is in degraded mode, in which case it may require user action to restore full functionality. Details can be found in the conditions field.
     */
    public /*out*/ readonly degraded!: pulumi.Output<boolean>;
    /**
     * Human-readable name for this resource.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
     * Terraform, other clients and services.
     */
    public /*out*/ readonly effectiveAnnotations!: pulumi.Output<{[key: string]: string}>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Whether to enable Linux `auditd` logging on the workstation. When enabled, a service account must also be specified that has `logging.buckets.write` permission on the project. Operating system audit logging is distinct from Cloud Audit Logs.
     */
    public readonly enableAuditAgent!: pulumi.Output<boolean | undefined>;
    /**
     * Encrypts resources of this workstation configuration using a customer-managed encryption key.
     * If specified, the boot disk of the Compute Engine instance and the persistent disk are encrypted using this encryption key. If this field is not set, the disks are encrypted using a generated key. Customer-managed encryption keys do not protect disk metadata.
     * If the customer-managed encryption key is rotated, when the workstation instance is stopped, the system attempts to recreate the persistent disk with the new version of the key. Be sure to keep older versions of the key until the persistent disk is recreated. Otherwise, data on the persistent disk will be lost.
     * If the encryption key is revoked, the workstation session will automatically be stopped within 7 hours.
     * Structure is documented below.
     */
    public readonly encryptionKey!: pulumi.Output<outputs.workstations.WorkstationConfigEncryptionKey | undefined>;
    /**
     * Checksum computed by the server.
     * May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Runtime host for a workstation.
     * Structure is documented below.
     */
    public readonly host!: pulumi.Output<outputs.workstations.WorkstationConfigHost>;
    /**
     * How long to wait before automatically stopping an instance that hasn't recently received any user traffic. A value of 0 indicates that this instance should never time out from idleness. Defaults to 20 minutes.
     * A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
     */
    public readonly idleTimeout!: pulumi.Output<string | undefined>;
    /**
     * Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location where the workstation cluster config should reside.
     *
     *
     * - - -
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Full name of this resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Directories to persist across workstation sessions.
     * Structure is documented below.
     */
    public readonly persistentDirectories!: pulumi.Output<outputs.workstations.WorkstationConfigPersistentDirectory[]>;
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
     * Specifies the zones used to replicate the VM and disk resources within the region. If set, exactly two zones within the workstation cluster's region must be specified—for example, `['us-central1-a', 'us-central1-f']`.
     * If this field is empty, two default zones within the region are used. Immutable after the workstation configuration is created.
     */
    public readonly replicaZones!: pulumi.Output<string[]>;
    /**
     * How long to wait before automatically stopping a workstation after it was started. A value of 0 indicates that workstations using this configuration should never time out from running duration. Must be greater than 0 and less than 24 hours if `encryptionKey` is set. Defaults to 12 hours.
     * A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
     */
    public readonly runningTimeout!: pulumi.Output<string | undefined>;
    /**
     * The system-generated UID of the resource.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The ID of the parent workstation cluster.
     */
    public readonly workstationClusterId!: pulumi.Output<string>;
    /**
     * The ID to be assigned to the workstation cluster config.
     */
    public readonly workstationConfigId!: pulumi.Output<string>;

    /**
     * Create a WorkstationConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkstationConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkstationConfigArgs | WorkstationConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkstationConfigState | undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["container"] = state ? state.container : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["degraded"] = state ? state.degraded : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["effectiveAnnotations"] = state ? state.effectiveAnnotations : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["enableAuditAgent"] = state ? state.enableAuditAgent : undefined;
            resourceInputs["encryptionKey"] = state ? state.encryptionKey : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["persistentDirectories"] = state ? state.persistentDirectories : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["replicaZones"] = state ? state.replicaZones : undefined;
            resourceInputs["runningTimeout"] = state ? state.runningTimeout : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["workstationClusterId"] = state ? state.workstationClusterId : undefined;
            resourceInputs["workstationConfigId"] = state ? state.workstationConfigId : undefined;
        } else {
            const args = argsOrState as WorkstationConfigArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.workstationClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workstationClusterId'");
            }
            if ((!args || args.workstationConfigId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workstationConfigId'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["container"] = args ? args.container : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enableAuditAgent"] = args ? args.enableAuditAgent : undefined;
            resourceInputs["encryptionKey"] = args ? args.encryptionKey : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["persistentDirectories"] = args ? args.persistentDirectories : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["replicaZones"] = args ? args.replicaZones : undefined;
            resourceInputs["runningTimeout"] = args ? args.runningTimeout : undefined;
            resourceInputs["workstationClusterId"] = args ? args.workstationClusterId : undefined;
            resourceInputs["workstationConfigId"] = args ? args.workstationConfigId : undefined;
            resourceInputs["conditions"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["degraded"] = undefined /*out*/;
            resourceInputs["effectiveAnnotations"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(WorkstationConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkstationConfig resources.
 */
export interface WorkstationConfigState {
    /**
     * Client-specified annotations. This is distinct from labels.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Status conditions describing the current resource state.
     * Structure is documented below.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.workstations.WorkstationConfigCondition>[]>;
    /**
     * Container that will be run for each workstation using this configuration when that workstation is started.
     * Structure is documented below.
     */
    container?: pulumi.Input<inputs.workstations.WorkstationConfigContainer>;
    /**
     * Time when this resource was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Whether this resource is in degraded mode, in which case it may require user action to restore full functionality. Details can be found in the conditions field.
     */
    degraded?: pulumi.Input<boolean>;
    /**
     * Human-readable name for this resource.
     */
    displayName?: pulumi.Input<string>;
    /**
     * All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
     * Terraform, other clients and services.
     */
    effectiveAnnotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to enable Linux `auditd` logging on the workstation. When enabled, a service account must also be specified that has `logging.buckets.write` permission on the project. Operating system audit logging is distinct from Cloud Audit Logs.
     */
    enableAuditAgent?: pulumi.Input<boolean>;
    /**
     * Encrypts resources of this workstation configuration using a customer-managed encryption key.
     * If specified, the boot disk of the Compute Engine instance and the persistent disk are encrypted using this encryption key. If this field is not set, the disks are encrypted using a generated key. Customer-managed encryption keys do not protect disk metadata.
     * If the customer-managed encryption key is rotated, when the workstation instance is stopped, the system attempts to recreate the persistent disk with the new version of the key. Be sure to keep older versions of the key until the persistent disk is recreated. Otherwise, data on the persistent disk will be lost.
     * If the encryption key is revoked, the workstation session will automatically be stopped within 7 hours.
     * Structure is documented below.
     */
    encryptionKey?: pulumi.Input<inputs.workstations.WorkstationConfigEncryptionKey>;
    /**
     * Checksum computed by the server.
     * May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * Runtime host for a workstation.
     * Structure is documented below.
     */
    host?: pulumi.Input<inputs.workstations.WorkstationConfigHost>;
    /**
     * How long to wait before automatically stopping an instance that hasn't recently received any user traffic. A value of 0 indicates that this instance should never time out from idleness. Defaults to 20 minutes.
     * A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
     */
    idleTimeout?: pulumi.Input<string>;
    /**
     * Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location where the workstation cluster config should reside.
     *
     *
     * - - -
     */
    location?: pulumi.Input<string>;
    /**
     * Full name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Directories to persist across workstation sessions.
     * Structure is documented below.
     */
    persistentDirectories?: pulumi.Input<pulumi.Input<inputs.workstations.WorkstationConfigPersistentDirectory>[]>;
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
     * Specifies the zones used to replicate the VM and disk resources within the region. If set, exactly two zones within the workstation cluster's region must be specified—for example, `['us-central1-a', 'us-central1-f']`.
     * If this field is empty, two default zones within the region are used. Immutable after the workstation configuration is created.
     */
    replicaZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * How long to wait before automatically stopping a workstation after it was started. A value of 0 indicates that workstations using this configuration should never time out from running duration. Must be greater than 0 and less than 24 hours if `encryptionKey` is set. Defaults to 12 hours.
     * A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
     */
    runningTimeout?: pulumi.Input<string>;
    /**
     * The system-generated UID of the resource.
     */
    uid?: pulumi.Input<string>;
    /**
     * The ID of the parent workstation cluster.
     */
    workstationClusterId?: pulumi.Input<string>;
    /**
     * The ID to be assigned to the workstation cluster config.
     */
    workstationConfigId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkstationConfig resource.
 */
export interface WorkstationConfigArgs {
    /**
     * Client-specified annotations. This is distinct from labels.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Container that will be run for each workstation using this configuration when that workstation is started.
     * Structure is documented below.
     */
    container?: pulumi.Input<inputs.workstations.WorkstationConfigContainer>;
    /**
     * Human-readable name for this resource.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Whether to enable Linux `auditd` logging on the workstation. When enabled, a service account must also be specified that has `logging.buckets.write` permission on the project. Operating system audit logging is distinct from Cloud Audit Logs.
     */
    enableAuditAgent?: pulumi.Input<boolean>;
    /**
     * Encrypts resources of this workstation configuration using a customer-managed encryption key.
     * If specified, the boot disk of the Compute Engine instance and the persistent disk are encrypted using this encryption key. If this field is not set, the disks are encrypted using a generated key. Customer-managed encryption keys do not protect disk metadata.
     * If the customer-managed encryption key is rotated, when the workstation instance is stopped, the system attempts to recreate the persistent disk with the new version of the key. Be sure to keep older versions of the key until the persistent disk is recreated. Otherwise, data on the persistent disk will be lost.
     * If the encryption key is revoked, the workstation session will automatically be stopped within 7 hours.
     * Structure is documented below.
     */
    encryptionKey?: pulumi.Input<inputs.workstations.WorkstationConfigEncryptionKey>;
    /**
     * Runtime host for a workstation.
     * Structure is documented below.
     */
    host?: pulumi.Input<inputs.workstations.WorkstationConfigHost>;
    /**
     * How long to wait before automatically stopping an instance that hasn't recently received any user traffic. A value of 0 indicates that this instance should never time out from idleness. Defaults to 20 minutes.
     * A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
     */
    idleTimeout?: pulumi.Input<string>;
    /**
     * Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location where the workstation cluster config should reside.
     *
     *
     * - - -
     */
    location: pulumi.Input<string>;
    /**
     * Directories to persist across workstation sessions.
     * Structure is documented below.
     */
    persistentDirectories?: pulumi.Input<pulumi.Input<inputs.workstations.WorkstationConfigPersistentDirectory>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Specifies the zones used to replicate the VM and disk resources within the region. If set, exactly two zones within the workstation cluster's region must be specified—for example, `['us-central1-a', 'us-central1-f']`.
     * If this field is empty, two default zones within the region are used. Immutable after the workstation configuration is created.
     */
    replicaZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * How long to wait before automatically stopping a workstation after it was started. A value of 0 indicates that workstations using this configuration should never time out from running duration. Must be greater than 0 and less than 24 hours if `encryptionKey` is set. Defaults to 12 hours.
     * A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
     */
    runningTimeout?: pulumi.Input<string>;
    /**
     * The ID of the parent workstation cluster.
     */
    workstationClusterId: pulumi.Input<string>;
    /**
     * The ID to be assigned to the workstation cluster config.
     */
    workstationConfigId: pulumi.Input<string>;
}

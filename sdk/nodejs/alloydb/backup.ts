// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * An AlloyDB Backup.
 *
 * To get more information about Backup, see:
 *
 * * [API documentation](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations.backups/create)
 * * How-to Guides
 *     * [AlloyDB](https://cloud.google.com/alloydb/docs/)
 *
 * ## Example Usage
 * ### Alloydb Backup Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {});
 * const defaultCluster = new gcp.alloydb.Cluster("defaultCluster", {
 *     clusterId: "alloydb-cluster",
 *     location: "us-central1",
 *     network: defaultNetwork.id,
 * });
 * const privateIpAlloc = new gcp.compute.GlobalAddress("privateIpAlloc", {
 *     addressType: "INTERNAL",
 *     purpose: "VPC_PEERING",
 *     prefixLength: 16,
 *     network: defaultNetwork.id,
 * });
 * const vpcConnection = new gcp.servicenetworking.Connection("vpcConnection", {
 *     network: defaultNetwork.id,
 *     service: "servicenetworking.googleapis.com",
 *     reservedPeeringRanges: [privateIpAlloc.name],
 * });
 * const defaultInstance = new gcp.alloydb.Instance("defaultInstance", {
 *     cluster: defaultCluster.name,
 *     instanceId: "alloydb-instance",
 *     instanceType: "PRIMARY",
 * }, {
 *     dependsOn: [vpcConnection],
 * });
 * const defaultBackup = new gcp.alloydb.Backup("defaultBackup", {
 *     location: "us-central1",
 *     backupId: "alloydb-backup",
 *     clusterName: defaultCluster.name,
 * }, {
 *     dependsOn: [defaultInstance],
 * });
 * ```
 * ### Alloydb Backup Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {});
 * const defaultCluster = new gcp.alloydb.Cluster("defaultCluster", {
 *     clusterId: "alloydb-cluster",
 *     location: "us-central1",
 *     network: defaultNetwork.id,
 * });
 * const privateIpAlloc = new gcp.compute.GlobalAddress("privateIpAlloc", {
 *     addressType: "INTERNAL",
 *     purpose: "VPC_PEERING",
 *     prefixLength: 16,
 *     network: defaultNetwork.id,
 * });
 * const vpcConnection = new gcp.servicenetworking.Connection("vpcConnection", {
 *     network: defaultNetwork.id,
 *     service: "servicenetworking.googleapis.com",
 *     reservedPeeringRanges: [privateIpAlloc.name],
 * });
 * const defaultInstance = new gcp.alloydb.Instance("defaultInstance", {
 *     cluster: defaultCluster.name,
 *     instanceId: "alloydb-instance",
 *     instanceType: "PRIMARY",
 * }, {
 *     dependsOn: [vpcConnection],
 * });
 * const defaultBackup = new gcp.alloydb.Backup("defaultBackup", {
 *     location: "us-central1",
 *     backupId: "alloydb-backup",
 *     clusterName: defaultCluster.name,
 *     description: "example description",
 *     type: "ON_DEMAND",
 *     labels: {
 *         label: "key",
 *     },
 * }, {
 *     dependsOn: [defaultInstance],
 * });
 * ```
 *
 * ## Import
 *
 * Backup can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:alloydb/backup:Backup default projects/{{project}}/locations/{{location}}/backups/{{backup_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:alloydb/backup:Backup default {{project}}/{{location}}/{{backup_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:alloydb/backup:Backup default {{location}}/{{backup_id}}
 * ```
 */
export class Backup extends pulumi.CustomResource {
    /**
     * Get an existing Backup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupState, opts?: pulumi.CustomResourceOptions): Backup {
        return new Backup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:alloydb/backup:Backup';

    /**
     * Returns true if the given object is an instance of Backup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backup.__pulumiType;
    }

    /**
     * Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the alloydb backup.
     */
    public readonly backupId!: pulumi.Output<string>;
    /**
     * The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Output only. The system-generated UID of the cluster which was used to create this resource.
     */
    public /*out*/ readonly clusterUid!: pulumi.Output<string>;
    /**
     * Output only. Create time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Output only. Delete time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * User-provided description of the backup.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User-settable and human-readable display name for the Backup.
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
     * EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
     * Structure is documented below.
     */
    public readonly encryptionConfig!: pulumi.Output<outputs.alloydb.BackupEncryptionConfig | undefined>;
    /**
     * EncryptionInfo describes the encryption information of a cluster or a backup.
     * Structure is documented below.
     */
    public /*out*/ readonly encryptionInfos!: pulumi.Output<outputs.alloydb.BackupEncryptionInfo[]>;
    /**
     * For Resource freshness validation (https://google.aip.dev/154)
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Output only. The QuantityBasedExpiry of the backup, specified by the backup's retention policy.
     * Once the expiry quantity is over retention, the backup is eligible to be garbage collected.
     * Structure is documented below.
     */
    public /*out*/ readonly expiryQuantities!: pulumi.Output<outputs.alloydb.BackupExpiryQuantity[]>;
    /**
     * Output only. The time at which after the backup is eligible to be garbage collected.
     * It is the duration specified by the backup's retention policy, added to the backup's createTime.
     */
    public /*out*/ readonly expiryTime!: pulumi.Output<string>;
    /**
     * User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location where the alloydb backup should reside.
     *
     *
     * - - -
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
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
     * Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively updating the resource.
     * This can happen due to user-triggered updates or system actions like failover or maintenance.
     */
    public /*out*/ readonly reconciling!: pulumi.Output<boolean>;
    /**
     * Output only. The size of the backup in bytes.
     */
    public /*out*/ readonly sizeBytes!: pulumi.Output<string>;
    /**
     * Output only. The current state of the backup.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The backup type, which suggests the trigger for the backup.
     * Possible values are: `TYPE_UNSPECIFIED`, `ON_DEMAND`, `AUTOMATED`, `CONTINUOUS`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Output only. Update time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Backup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupArgs | BackupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupState | undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["backupId"] = state ? state.backupId : undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["clusterUid"] = state ? state.clusterUid : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["deleteTime"] = state ? state.deleteTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["effectiveAnnotations"] = state ? state.effectiveAnnotations : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["encryptionConfig"] = state ? state.encryptionConfig : undefined;
            resourceInputs["encryptionInfos"] = state ? state.encryptionInfos : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["expiryQuantities"] = state ? state.expiryQuantities : undefined;
            resourceInputs["expiryTime"] = state ? state.expiryTime : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["reconciling"] = state ? state.reconciling : undefined;
            resourceInputs["sizeBytes"] = state ? state.sizeBytes : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as BackupArgs | undefined;
            if ((!args || args.backupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupId'");
            }
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["backupId"] = args ? args.backupId : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["encryptionConfig"] = args ? args.encryptionConfig : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["clusterUid"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["effectiveAnnotations"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["encryptionInfos"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["expiryQuantities"] = undefined /*out*/;
            resourceInputs["expiryTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["reconciling"] = undefined /*out*/;
            resourceInputs["sizeBytes"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Backup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Backup resources.
 */
export interface BackupState {
    /**
     * Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the alloydb backup.
     */
    backupId?: pulumi.Input<string>;
    /**
     * The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Output only. The system-generated UID of the cluster which was used to create this resource.
     */
    clusterUid?: pulumi.Input<string>;
    /**
     * Output only. Create time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    createTime?: pulumi.Input<string>;
    /**
     * Output only. Delete time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    deleteTime?: pulumi.Input<string>;
    /**
     * User-provided description of the backup.
     */
    description?: pulumi.Input<string>;
    /**
     * User-settable and human-readable display name for the Backup.
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
     * EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
     * Structure is documented below.
     */
    encryptionConfig?: pulumi.Input<inputs.alloydb.BackupEncryptionConfig>;
    /**
     * EncryptionInfo describes the encryption information of a cluster or a backup.
     * Structure is documented below.
     */
    encryptionInfos?: pulumi.Input<pulumi.Input<inputs.alloydb.BackupEncryptionInfo>[]>;
    /**
     * For Resource freshness validation (https://google.aip.dev/154)
     */
    etag?: pulumi.Input<string>;
    /**
     * Output only. The QuantityBasedExpiry of the backup, specified by the backup's retention policy.
     * Once the expiry quantity is over retention, the backup is eligible to be garbage collected.
     * Structure is documented below.
     */
    expiryQuantities?: pulumi.Input<pulumi.Input<inputs.alloydb.BackupExpiryQuantity>[]>;
    /**
     * Output only. The time at which after the backup is eligible to be garbage collected.
     * It is the duration specified by the backup's retention policy, added to the backup's createTime.
     */
    expiryTime?: pulumi.Input<string>;
    /**
     * User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location where the alloydb backup should reside.
     *
     *
     * - - -
     */
    location?: pulumi.Input<string>;
    /**
     * Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
     */
    name?: pulumi.Input<string>;
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
     * Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively updating the resource.
     * This can happen due to user-triggered updates or system actions like failover or maintenance.
     */
    reconciling?: pulumi.Input<boolean>;
    /**
     * Output only. The size of the backup in bytes.
     */
    sizeBytes?: pulumi.Input<string>;
    /**
     * Output only. The current state of the backup.
     */
    state?: pulumi.Input<string>;
    /**
     * The backup type, which suggests the trigger for the backup.
     * Possible values are: `TYPE_UNSPECIFIED`, `ON_DEMAND`, `AUTOMATED`, `CONTINUOUS`.
     */
    type?: pulumi.Input<string>;
    /**
     * Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
     */
    uid?: pulumi.Input<string>;
    /**
     * Output only. Update time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Backup resource.
 */
export interface BackupArgs {
    /**
     * Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the alloydb backup.
     */
    backupId: pulumi.Input<string>;
    /**
     * The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
     */
    clusterName: pulumi.Input<string>;
    /**
     * User-provided description of the backup.
     */
    description?: pulumi.Input<string>;
    /**
     * User-settable and human-readable display name for the Backup.
     */
    displayName?: pulumi.Input<string>;
    /**
     * EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
     * Structure is documented below.
     */
    encryptionConfig?: pulumi.Input<inputs.alloydb.BackupEncryptionConfig>;
    /**
     * User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location where the alloydb backup should reside.
     *
     *
     * - - -
     */
    location: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The backup type, which suggests the trigger for the backup.
     * Possible values are: `TYPE_UNSPECIFIED`, `ON_DEMAND`, `AUTOMATED`, `CONTINUOUS`.
     */
    type?: pulumi.Input<string>;
}

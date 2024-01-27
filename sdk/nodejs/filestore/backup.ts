// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Google Cloud Filestore backup.
 *
 * To get more information about Backup, see:
 *
 * * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1/projects.locations.instances.backups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/filestore/docs/backups)
 *     * [Creating Backups](https://cloud.google.com/filestore/docs/create-backups)
 *
 * ## Example Usage
 * ### Filestore Backup Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.filestore.Instance("instance", {
 *     location: "us-central1-b",
 *     tier: "BASIC_HDD",
 *     fileShares: {
 *         capacityGb: 1024,
 *         name: "share1",
 *     },
 *     networks: [{
 *         network: "default",
 *         modes: ["MODE_IPV4"],
 *         connectMode: "DIRECT_PEERING",
 *     }],
 * });
 * const backup = new gcp.filestore.Backup("backup", {
 *     location: "us-central1",
 *     description: "This is a filestore backup for the test instance",
 *     sourceInstance: instance.id,
 *     sourceFileShare: "share1",
 *     labels: {
 *         files: "label1",
 *         "other-label": "label2",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Backup can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/backups/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Backup can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:filestore/backup:Backup default projects/{{project}}/locations/{{location}}/backups/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:filestore/backup:Backup default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:filestore/backup:Backup default {{location}}/{{name}}
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
    public static readonly __pulumiType = 'gcp:filestore/backup:Backup';

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
     * The amount of bytes needed to allocate a full copy of the snapshot content.
     */
    public /*out*/ readonly capacityGb!: pulumi.Output<string>;
    /**
     * The time when the snapshot was created in RFC3339 text format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Amount of bytes that will be downloaded if the backup is restored.
     */
    public /*out*/ readonly downloadBytes!: pulumi.Output<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * KMS key name used for data encryption.
     */
    public /*out*/ readonly kmsKeyName!: pulumi.Output<string>;
    /**
     * Resource labels to represent user-provided metadata.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     *
     *
     * - - -
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name of the backup. The name must be unique within the specified instance.
     * The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
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
     * Name of the file share in the source Cloud Filestore instance that the backup is created from.
     */
    public readonly sourceFileShare!: pulumi.Output<string>;
    /**
     * The resource name of the source Cloud Filestore instance, in the format projects/{projectId}/locations/{locationId}/instances/{instanceId}, used to create this backup.
     */
    public readonly sourceInstance!: pulumi.Output<string>;
    /**
     * The service tier of the source Cloud Filestore instance that this backup is created from.
     */
    public /*out*/ readonly sourceInstanceTier!: pulumi.Output<string>;
    /**
     * The backup state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.
     */
    public /*out*/ readonly storageBytes!: pulumi.Output<string>;

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
            resourceInputs["capacityGb"] = state ? state.capacityGb : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["downloadBytes"] = state ? state.downloadBytes : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["kmsKeyName"] = state ? state.kmsKeyName : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["sourceFileShare"] = state ? state.sourceFileShare : undefined;
            resourceInputs["sourceInstance"] = state ? state.sourceInstance : undefined;
            resourceInputs["sourceInstanceTier"] = state ? state.sourceInstanceTier : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["storageBytes"] = state ? state.storageBytes : undefined;
        } else {
            const args = argsOrState as BackupArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.sourceFileShare === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceFileShare'");
            }
            if ((!args || args.sourceInstance === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceInstance'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["sourceFileShare"] = args ? args.sourceFileShare : undefined;
            resourceInputs["sourceInstance"] = args ? args.sourceInstance : undefined;
            resourceInputs["capacityGb"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["downloadBytes"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["kmsKeyName"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["sourceInstanceTier"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["storageBytes"] = undefined /*out*/;
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
     * The amount of bytes needed to allocate a full copy of the snapshot content.
     */
    capacityGb?: pulumi.Input<string>;
    /**
     * The time when the snapshot was created in RFC3339 text format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
     */
    description?: pulumi.Input<string>;
    /**
     * Amount of bytes that will be downloaded if the backup is restored.
     */
    downloadBytes?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * KMS key name used for data encryption.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * Resource labels to represent user-provided metadata.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     *
     *
     * - - -
     */
    location?: pulumi.Input<string>;
    /**
     * The resource name of the backup. The name must be unique within the specified instance.
     * The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
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
     * Name of the file share in the source Cloud Filestore instance that the backup is created from.
     */
    sourceFileShare?: pulumi.Input<string>;
    /**
     * The resource name of the source Cloud Filestore instance, in the format projects/{projectId}/locations/{locationId}/instances/{instanceId}, used to create this backup.
     */
    sourceInstance?: pulumi.Input<string>;
    /**
     * The service tier of the source Cloud Filestore instance that this backup is created from.
     */
    sourceInstanceTier?: pulumi.Input<string>;
    /**
     * The backup state.
     */
    state?: pulumi.Input<string>;
    /**
     * The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.
     */
    storageBytes?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Backup resource.
 */
export interface BackupArgs {
    /**
     * A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
     */
    description?: pulumi.Input<string>;
    /**
     * Resource labels to represent user-provided metadata.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     *
     *
     * - - -
     */
    location: pulumi.Input<string>;
    /**
     * The resource name of the backup. The name must be unique within the specified instance.
     * The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Name of the file share in the source Cloud Filestore instance that the backup is created from.
     */
    sourceFileShare: pulumi.Input<string>;
    /**
     * The resource name of the source Cloud Filestore instance, in the format projects/{projectId}/locations/{locationId}/instances/{instanceId}, used to create this backup.
     */
    sourceInstance: pulumi.Input<string>;
}

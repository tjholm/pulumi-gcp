// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The Dataplex Lake resource
 *
 * ## Example Usage
 * ### Basic_lake
 * A basic example of a dataplex lake
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const primary = new gcp.dataplex.Lake("primary", {
 *     description: "Lake for DCL",
 *     displayName: "Lake for DCL",
 *     labels: {
 *         "my-lake": "exists",
 *     },
 *     location: "us-west1",
 *     project: "my-project-name",
 * });
 * ```
 *
 * ## Import
 *
 * Lake can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:dataplex/lake:Lake default projects/{{project}}/locations/{{location}}/lakes/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataplex/lake:Lake default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataplex/lake:Lake default {{location}}/{{name}}
 * ```
 */
export class Lake extends pulumi.CustomResource {
    /**
     * Get an existing Lake resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LakeState, opts?: pulumi.CustomResourceOptions): Lake {
        return new Lake(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataplex/lake:Lake';

    /**
     * Returns true if the given object is an instance of Lake.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Lake {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Lake.__pulumiType;
    }

    /**
     * Output only. Aggregated status of the underlying assets of the lake.
     */
    public /*out*/ readonly assetStatuses!: pulumi.Output<outputs.dataplex.LakeAssetStatus[]>;
    /**
     * Output only. The time when the lake was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Description of the lake.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Optional. User friendly display name.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: any}>;
    /**
     * Optional. User-defined labels for the lake.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location for the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Optional. Settings to manage lake and Dataproc Metastore service instance association.
     */
    public readonly metastore!: pulumi.Output<outputs.dataplex.LakeMetastore | undefined>;
    /**
     * Output only. Metastore status of the lake.
     */
    public /*out*/ readonly metastoreStatuses!: pulumi.Output<outputs.dataplex.LakeMetastoreStatus[]>;
    /**
     * The name of the lake.
     *
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: any}>;
    /**
     * Output only. Service account associated with this lake. This service account must be authorized to access or operate on resources managed by the lake.
     */
    public /*out*/ readonly serviceAccount!: pulumi.Output<string>;
    /**
     * Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Output only. System generated globally unique ID for the lake. This ID will be different if the lake is deleted and re-created with the same name.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Output only. The time when the lake was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Lake resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LakeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LakeArgs | LakeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LakeState | undefined;
            resourceInputs["assetStatuses"] = state ? state.assetStatuses : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["metastore"] = state ? state.metastore : undefined;
            resourceInputs["metastoreStatuses"] = state ? state.metastoreStatuses : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as LakeArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["metastore"] = args ? args.metastore : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["assetStatuses"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["metastoreStatuses"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Lake.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Lake resources.
 */
export interface LakeState {
    /**
     * Output only. Aggregated status of the underlying assets of the lake.
     */
    assetStatuses?: pulumi.Input<pulumi.Input<inputs.dataplex.LakeAssetStatus>[]>;
    /**
     * Output only. The time when the lake was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Optional. Description of the lake.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. User friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Optional. User-defined labels for the lake.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location?: pulumi.Input<string>;
    /**
     * Optional. Settings to manage lake and Dataproc Metastore service instance association.
     */
    metastore?: pulumi.Input<inputs.dataplex.LakeMetastore>;
    /**
     * Output only. Metastore status of the lake.
     */
    metastoreStatuses?: pulumi.Input<pulumi.Input<inputs.dataplex.LakeMetastoreStatus>[]>;
    /**
     * The name of the lake.
     *
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Output only. Service account associated with this lake. This service account must be authorized to access or operate on resources managed by the lake.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
     */
    state?: pulumi.Input<string>;
    /**
     * Output only. System generated globally unique ID for the lake. This ID will be different if the lake is deleted and re-created with the same name.
     */
    uid?: pulumi.Input<string>;
    /**
     * Output only. The time when the lake was last updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Lake resource.
 */
export interface LakeArgs {
    /**
     * Optional. Description of the lake.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. User friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Optional. User-defined labels for the lake.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location: pulumi.Input<string>;
    /**
     * Optional. Settings to manage lake and Dataproc Metastore service instance association.
     */
    metastore?: pulumi.Input<inputs.dataplex.LakeMetastore>;
    /**
     * The name of the lake.
     *
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
}

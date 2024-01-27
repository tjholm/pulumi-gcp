// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A collection of DataItems and Annotations on them.
 *
 * To get more information about Featurestore, see:
 *
 * * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/vertex-ai/docs)
 *
 * ## Example Usage
 * ### Vertex Ai Featurestore
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const featurestore = new gcp.vertex.AiFeatureStore("featurestore", {
 *     encryptionSpec: {
 *         kmsKeyName: "kms-name",
 *     },
 *     forceDestroy: true,
 *     labels: {
 *         foo: "bar",
 *     },
 *     onlineServingConfig: {
 *         fixedNodeCount: 2,
 *     },
 *     region: "us-central1",
 * });
 * ```
 * ### Vertex Ai Featurestore With Beta Fields
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const featurestore = new gcp.vertex.AiFeatureStore("featurestore", {
 *     labels: {
 *         foo: "bar",
 *     },
 *     region: "us-central1",
 *     onlineServingConfig: {
 *         fixedNodeCount: 2,
 *     },
 *     encryptionSpec: {
 *         kmsKeyName: "kms-name",
 *     },
 *     onlineStorageTtlDays: 30,
 *     forceDestroy: true,
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Vertex Ai Featurestore Scaling
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const featurestore = new gcp.vertex.AiFeatureStore("featurestore", {
 *     encryptionSpec: {
 *         kmsKeyName: "kms-name",
 *     },
 *     forceDestroy: true,
 *     labels: {
 *         foo: "bar",
 *     },
 *     onlineServingConfig: {
 *         scaling: {
 *             maxNodeCount: 10,
 *             minNodeCount: 2,
 *         },
 *     },
 *     region: "us-central1",
 * });
 * ```
 *
 * ## Import
 *
 * Featurestore can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/featurestores/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Featurestore can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureStore:AiFeatureStore default projects/{{project}}/locations/{{region}}/featurestores/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureStore:AiFeatureStore default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureStore:AiFeatureStore default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureStore:AiFeatureStore default {{name}}
 * ```
 */
export class AiFeatureStore extends pulumi.CustomResource {
    /**
     * Get an existing AiFeatureStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AiFeatureStoreState, opts?: pulumi.CustomResourceOptions): AiFeatureStore {
        return new AiFeatureStore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:vertex/aiFeatureStore:AiFeatureStore';

    /**
     * Returns true if the given object is an instance of AiFeatureStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AiFeatureStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AiFeatureStore.__pulumiType;
    }

    /**
     * The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * If set, both of the online and offline data storage will be secured by this key.
     * Structure is documented below.
     */
    public readonly encryptionSpec!: pulumi.Output<outputs.vertex.AiFeatureStoreEncryptionSpec | undefined>;
    /**
     * Used to perform consistent read-modify-write updates.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * If set to true, any EntityTypes and Features for this Featurestore will also be deleted
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * A set of key/value label pairs to assign to this Featurestore.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Config for online serving resources.
     * Structure is documented below.
     */
    public readonly onlineServingConfig!: pulumi.Output<outputs.vertex.AiFeatureStoreOnlineServingConfig | undefined>;
    /**
     * (Optional, Beta)
     * TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a featurestore. If not set, default to 4000 days
     */
    public readonly onlineStorageTtlDays!: pulumi.Output<number | undefined>;
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
     * The region of the dataset. eg us-central1
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a AiFeatureStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AiFeatureStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AiFeatureStoreArgs | AiFeatureStoreState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AiFeatureStoreState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["encryptionSpec"] = state ? state.encryptionSpec : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["onlineServingConfig"] = state ? state.onlineServingConfig : undefined;
            resourceInputs["onlineStorageTtlDays"] = state ? state.onlineStorageTtlDays : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as AiFeatureStoreArgs | undefined;
            resourceInputs["encryptionSpec"] = args ? args.encryptionSpec : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["onlineServingConfig"] = args ? args.onlineServingConfig : undefined;
            resourceInputs["onlineStorageTtlDays"] = args ? args.onlineStorageTtlDays : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AiFeatureStore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AiFeatureStore resources.
 */
export interface AiFeatureStoreState {
    /**
     * The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     */
    createTime?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If set, both of the online and offline data storage will be secured by this key.
     * Structure is documented below.
     */
    encryptionSpec?: pulumi.Input<inputs.vertex.AiFeatureStoreEncryptionSpec>;
    /**
     * Used to perform consistent read-modify-write updates.
     */
    etag?: pulumi.Input<string>;
    /**
     * If set to true, any EntityTypes and Features for this Featurestore will also be deleted
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * A set of key/value label pairs to assign to this Featurestore.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     */
    name?: pulumi.Input<string>;
    /**
     * Config for online serving resources.
     * Structure is documented below.
     */
    onlineServingConfig?: pulumi.Input<inputs.vertex.AiFeatureStoreOnlineServingConfig>;
    /**
     * (Optional, Beta)
     * TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a featurestore. If not set, default to 4000 days
     */
    onlineStorageTtlDays?: pulumi.Input<number>;
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
     * The region of the dataset. eg us-central1
     */
    region?: pulumi.Input<string>;
    /**
     * The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AiFeatureStore resource.
 */
export interface AiFeatureStoreArgs {
    /**
     * If set, both of the online and offline data storage will be secured by this key.
     * Structure is documented below.
     */
    encryptionSpec?: pulumi.Input<inputs.vertex.AiFeatureStoreEncryptionSpec>;
    /**
     * If set to true, any EntityTypes and Features for this Featurestore will also be deleted
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * A set of key/value label pairs to assign to this Featurestore.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     */
    name?: pulumi.Input<string>;
    /**
     * Config for online serving resources.
     * Structure is documented below.
     */
    onlineServingConfig?: pulumi.Input<inputs.vertex.AiFeatureStoreOnlineServingConfig>;
    /**
     * (Optional, Beta)
     * TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a featurestore. If not set, default to 4000 days
     */
    onlineStorageTtlDays?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region of the dataset. eg us-central1
     */
    region?: pulumi.Input<string>;
}

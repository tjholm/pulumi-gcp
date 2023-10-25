// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a new bucket in Google cloud storage service (GCS).
 * Once a bucket has been created, its location can't be changed.
 *
 * For more information see
 * [the official documentation](https://cloud.google.com/storage/docs/overview)
 * and
 * [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
 *
 * **Note**: If the project id is not set on the resource or in the provider block it will be dynamically
 * determined which will require enabling the compute api.
 *
 * ## Example Usage
 * ### Creating A Private Bucket In Standard Storage, In The EU Region. Bucket Configured As Static Website And CORS Configurations
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const static_site = new gcp.storage.Bucket("static-site", {
 *     cors: [{
 *         maxAgeSeconds: 3600,
 *         methods: [
 *             "GET",
 *             "HEAD",
 *             "PUT",
 *             "POST",
 *             "DELETE",
 *         ],
 *         origins: ["http://image-store.com"],
 *         responseHeaders: ["*"],
 *     }],
 *     forceDestroy: true,
 *     location: "EU",
 *     uniformBucketLevelAccess: true,
 *     website: {
 *         mainPageSuffix: "index.html",
 *         notFoundPage: "404.html",
 *     },
 * });
 * ```
 * ### Life Cycle Settings For Storage Bucket Objects
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const auto_expire = new gcp.storage.Bucket("auto-expire", {
 *     forceDestroy: true,
 *     lifecycleRules: [
 *         {
 *             action: {
 *                 type: "Delete",
 *             },
 *             condition: {
 *                 age: 3,
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "AbortIncompleteMultipartUpload",
 *             },
 *             condition: {
 *                 age: 1,
 *             },
 *         },
 *     ],
 *     location: "US",
 * });
 * ```
 * ### Enabling Public Access Prevention
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const auto_expire = new gcp.storage.Bucket("auto-expire", {
 *     forceDestroy: true,
 *     location: "US",
 *     publicAccessPrevention: "enforced",
 * });
 * ```
 *
 * ## Import
 *
 * Storage buckets can be imported using the `name` or
 *
 * `project/name`. If the project is not passed to the import command it will be inferred from the provider block or environment variables. If it cannot be inferred it will be queried from the Compute API (this will fail if the API is not enabled). e.g.
 *
 * ```sh
 *  $ pulumi import gcp:storage/bucket:Bucket image-store image-store-bucket
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:storage/bucket:Bucket image-store tf-test-project/image-store-bucket
 * ```
 *
 *  `false` in state. If you've set it to `true` in config, run `pulumi up` to update the value set in state. If you delete this resource before updating the value, objects in the bucket will not be destroyed.
 */
export class Bucket extends pulumi.CustomResource {
    /**
     * Get an existing Bucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketState, opts?: pulumi.CustomResourceOptions): Bucket {
        return new Bucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:storage/bucket:Bucket';

    /**
     * Returns true if the given object is an instance of Bucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bucket.__pulumiType;
    }

    /**
     * The bucket's [Autoclass](https://cloud.google.com/storage/docs/autoclass) configuration.  Structure is documented below.
     */
    public readonly autoclass!: pulumi.Output<outputs.storage.BucketAutoclass | undefined>;
    /**
     * The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
     */
    public readonly cors!: pulumi.Output<outputs.storage.BucketCor[] | undefined>;
    /**
     * The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty. Structure is documented below.
     */
    public readonly customPlacementConfig!: pulumi.Output<outputs.storage.BucketCustomPlacementConfig | undefined>;
    /**
     * Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.
     */
    public readonly defaultEventBasedHold!: pulumi.Output<boolean | undefined>;
    /**
     * The bucket's encryption configuration. Structure is documented below.
     */
    public readonly encryption!: pulumi.Output<outputs.storage.BucketEncryption | undefined>;
    /**
     * When deleting a bucket, this
     * boolean option will delete all contained objects. If you try to delete a
     * bucket that contains objects, the provider will fail that run.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * A map of key/value label pairs to assign to the bucket.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
     */
    public readonly lifecycleRules!: pulumi.Output<outputs.storage.BucketLifecycleRule[] | undefined>;
    /**
     * The [GCS location](https://cloud.google.com/storage/docs/bucket-locations).
     *
     * - - -
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration. Structure is documented below.
     */
    public readonly logging!: pulumi.Output<outputs.storage.BucketLogging | undefined>;
    /**
     * The name of the bucket.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Prevents public access to a bucket. Acceptable values are "inherited" or "enforced". If "inherited", the bucket uses [public access prevention](https://cloud.google.com/storage/docs/public-access-prevention). only if the bucket is subject to the public access prevention organization policy constraint. Defaults to "inherited".
     */
    public readonly publicAccessPrevention!: pulumi.Output<string>;
    /**
     * Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
     */
    public readonly requesterPays!: pulumi.Output<boolean | undefined>;
    /**
     * Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
     */
    public readonly retentionPolicy!: pulumi.Output<outputs.storage.BucketRetentionPolicy | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
     */
    public readonly storageClass!: pulumi.Output<string | undefined>;
    /**
     * Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
     */
    public readonly uniformBucketLevelAccess!: pulumi.Output<boolean>;
    /**
     * The base URL of the bucket, in the format `gs://<bucket-name>`.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.  Structure is documented below.
     */
    public readonly versioning!: pulumi.Output<outputs.storage.BucketVersioning>;
    /**
     * Configuration if the bucket acts as a website. Structure is documented below.
     */
    public readonly website!: pulumi.Output<outputs.storage.BucketWebsite>;

    /**
     * Create a Bucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketArgs | BucketState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketState | undefined;
            resourceInputs["autoclass"] = state ? state.autoclass : undefined;
            resourceInputs["cors"] = state ? state.cors : undefined;
            resourceInputs["customPlacementConfig"] = state ? state.customPlacementConfig : undefined;
            resourceInputs["defaultEventBasedHold"] = state ? state.defaultEventBasedHold : undefined;
            resourceInputs["encryption"] = state ? state.encryption : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["lifecycleRules"] = state ? state.lifecycleRules : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["logging"] = state ? state.logging : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["publicAccessPrevention"] = state ? state.publicAccessPrevention : undefined;
            resourceInputs["requesterPays"] = state ? state.requesterPays : undefined;
            resourceInputs["retentionPolicy"] = state ? state.retentionPolicy : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["storageClass"] = state ? state.storageClass : undefined;
            resourceInputs["uniformBucketLevelAccess"] = state ? state.uniformBucketLevelAccess : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["versioning"] = state ? state.versioning : undefined;
            resourceInputs["website"] = state ? state.website : undefined;
        } else {
            const args = argsOrState as BucketArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["autoclass"] = args ? args.autoclass : undefined;
            resourceInputs["cors"] = args ? args.cors : undefined;
            resourceInputs["customPlacementConfig"] = args ? args.customPlacementConfig : undefined;
            resourceInputs["defaultEventBasedHold"] = args ? args.defaultEventBasedHold : undefined;
            resourceInputs["encryption"] = args ? args.encryption : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["lifecycleRules"] = args ? args.lifecycleRules : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["logging"] = args ? args.logging : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["publicAccessPrevention"] = args ? args.publicAccessPrevention : undefined;
            resourceInputs["requesterPays"] = args ? args.requesterPays : undefined;
            resourceInputs["retentionPolicy"] = args ? args.retentionPolicy : undefined;
            resourceInputs["storageClass"] = args ? args.storageClass : undefined;
            resourceInputs["uniformBucketLevelAccess"] = args ? args.uniformBucketLevelAccess : undefined;
            resourceInputs["versioning"] = args ? args.versioning : undefined;
            resourceInputs["website"] = args ? args.website : undefined;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Bucket.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Bucket resources.
 */
export interface BucketState {
    /**
     * The bucket's [Autoclass](https://cloud.google.com/storage/docs/autoclass) configuration.  Structure is documented below.
     */
    autoclass?: pulumi.Input<inputs.storage.BucketAutoclass>;
    /**
     * The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
     */
    cors?: pulumi.Input<pulumi.Input<inputs.storage.BucketCor>[]>;
    /**
     * The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty. Structure is documented below.
     */
    customPlacementConfig?: pulumi.Input<inputs.storage.BucketCustomPlacementConfig>;
    /**
     * Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.
     */
    defaultEventBasedHold?: pulumi.Input<boolean>;
    /**
     * The bucket's encryption configuration. Structure is documented below.
     */
    encryption?: pulumi.Input<inputs.storage.BucketEncryption>;
    /**
     * When deleting a bucket, this
     * boolean option will delete all contained objects. If you try to delete a
     * bucket that contains objects, the provider will fail that run.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * A map of key/value label pairs to assign to the bucket.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
     */
    lifecycleRules?: pulumi.Input<pulumi.Input<inputs.storage.BucketLifecycleRule>[]>;
    /**
     * The [GCS location](https://cloud.google.com/storage/docs/bucket-locations).
     *
     * - - -
     */
    location?: pulumi.Input<string>;
    /**
     * The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration. Structure is documented below.
     */
    logging?: pulumi.Input<inputs.storage.BucketLogging>;
    /**
     * The name of the bucket.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Prevents public access to a bucket. Acceptable values are "inherited" or "enforced". If "inherited", the bucket uses [public access prevention](https://cloud.google.com/storage/docs/public-access-prevention). only if the bucket is subject to the public access prevention organization policy constraint. Defaults to "inherited".
     */
    publicAccessPrevention?: pulumi.Input<string>;
    /**
     * Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
     */
    requesterPays?: pulumi.Input<boolean>;
    /**
     * Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
     */
    retentionPolicy?: pulumi.Input<inputs.storage.BucketRetentionPolicy>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
     */
    storageClass?: pulumi.Input<string>;
    /**
     * Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
     */
    uniformBucketLevelAccess?: pulumi.Input<boolean>;
    /**
     * The base URL of the bucket, in the format `gs://<bucket-name>`.
     */
    url?: pulumi.Input<string>;
    /**
     * The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.  Structure is documented below.
     */
    versioning?: pulumi.Input<inputs.storage.BucketVersioning>;
    /**
     * Configuration if the bucket acts as a website. Structure is documented below.
     */
    website?: pulumi.Input<inputs.storage.BucketWebsite>;
}

/**
 * The set of arguments for constructing a Bucket resource.
 */
export interface BucketArgs {
    /**
     * The bucket's [Autoclass](https://cloud.google.com/storage/docs/autoclass) configuration.  Structure is documented below.
     */
    autoclass?: pulumi.Input<inputs.storage.BucketAutoclass>;
    /**
     * The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
     */
    cors?: pulumi.Input<pulumi.Input<inputs.storage.BucketCor>[]>;
    /**
     * The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty. Structure is documented below.
     */
    customPlacementConfig?: pulumi.Input<inputs.storage.BucketCustomPlacementConfig>;
    /**
     * Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.
     */
    defaultEventBasedHold?: pulumi.Input<boolean>;
    /**
     * The bucket's encryption configuration. Structure is documented below.
     */
    encryption?: pulumi.Input<inputs.storage.BucketEncryption>;
    /**
     * When deleting a bucket, this
     * boolean option will delete all contained objects. If you try to delete a
     * bucket that contains objects, the provider will fail that run.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * A map of key/value label pairs to assign to the bucket.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
     */
    lifecycleRules?: pulumi.Input<pulumi.Input<inputs.storage.BucketLifecycleRule>[]>;
    /**
     * The [GCS location](https://cloud.google.com/storage/docs/bucket-locations).
     *
     * - - -
     */
    location: pulumi.Input<string>;
    /**
     * The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration. Structure is documented below.
     */
    logging?: pulumi.Input<inputs.storage.BucketLogging>;
    /**
     * The name of the bucket.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Prevents public access to a bucket. Acceptable values are "inherited" or "enforced". If "inherited", the bucket uses [public access prevention](https://cloud.google.com/storage/docs/public-access-prevention). only if the bucket is subject to the public access prevention organization policy constraint. Defaults to "inherited".
     */
    publicAccessPrevention?: pulumi.Input<string>;
    /**
     * Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
     */
    requesterPays?: pulumi.Input<boolean>;
    /**
     * Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
     */
    retentionPolicy?: pulumi.Input<inputs.storage.BucketRetentionPolicy>;
    /**
     * The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
     */
    storageClass?: pulumi.Input<string>;
    /**
     * Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
     */
    uniformBucketLevelAccess?: pulumi.Input<boolean>;
    /**
     * The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.  Structure is documented below.
     */
    versioning?: pulumi.Input<inputs.storage.BucketVersioning>;
    /**
     * Configuration if the bucket acts as a website. Structure is documented below.
     */
    website?: pulumi.Input<inputs.storage.BucketWebsite>;
}

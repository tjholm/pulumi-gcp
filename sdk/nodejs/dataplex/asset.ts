// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The Dataplex Asset resource
 *
 * ## Example Usage
 * ### Basic_asset
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basicBucket = new gcp.storage.Bucket("basicBucket", {
 *     location: "us-west1",
 *     uniformBucketLevelAccess: true,
 *     project: "my-project-name",
 * });
 * const basicLake = new gcp.dataplex.Lake("basicLake", {
 *     location: "us-west1",
 *     project: "my-project-name",
 * });
 * const basicZone = new gcp.dataplex.Zone("basicZone", {
 *     location: "us-west1",
 *     lake: basicLake.name,
 *     type: "RAW",
 *     discoverySpec: {
 *         enabled: false,
 *     },
 *     resourceSpec: {
 *         locationType: "SINGLE_REGION",
 *     },
 *     project: "my-project-name",
 * });
 * const primary = new gcp.dataplex.Asset("primary", {
 *     location: "us-west1",
 *     lake: basicLake.name,
 *     dataplexZone: basicZone.name,
 *     discoverySpec: {
 *         enabled: false,
 *     },
 *     resourceSpec: {
 *         name: "projects/my-project-name/buckets/bucket",
 *         type: "STORAGE_BUCKET",
 *     },
 *     project: "my-project-name",
 * }, {
 *     dependsOn: [basicBucket],
 * });
 * ```
 *
 * ## Import
 *
 * Asset can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:dataplex/asset:Asset default projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}/assets/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataplex/asset:Asset default {{project}}/{{location}}/{{lake}}/{{dataplex_zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataplex/asset:Asset default {{location}}/{{lake}}/{{dataplex_zone}}/{{name}}
 * ```
 */
export class Asset extends pulumi.CustomResource {
    /**
     * Get an existing Asset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssetState, opts?: pulumi.CustomResourceOptions): Asset {
        return new Asset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataplex/asset:Asset';

    /**
     * Returns true if the given object is an instance of Asset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Asset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Asset.__pulumiType;
    }

    /**
     * Output only. The time when the asset was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The zone for the resource
     */
    public readonly dataplexZone!: pulumi.Output<string>;
    /**
     * Optional. Description of the asset.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
     */
    public readonly discoverySpec!: pulumi.Output<outputs.dataplex.AssetDiscoverySpec>;
    /**
     * Output only. Status of the discovery feature applied to data referenced by this asset.
     */
    public /*out*/ readonly discoveryStatuses!: pulumi.Output<outputs.dataplex.AssetDiscoveryStatus[]>;
    /**
     * Optional. User friendly display name.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Optional. User defined labels for the asset.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The lake for the resource
     */
    public readonly lake!: pulumi.Output<string>;
    /**
     * The location for the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the asset.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Required. Immutable. Specification of the resource that is referenced by this asset.
     */
    public readonly resourceSpec!: pulumi.Output<outputs.dataplex.AssetResourceSpec>;
    /**
     * Output only. Status of the resource referenced by this asset.
     */
    public /*out*/ readonly resourceStatuses!: pulumi.Output<outputs.dataplex.AssetResourceStatus[]>;
    /**
     * Output only. Status of the security policy applied to resource referenced by this asset.
     */
    public /*out*/ readonly securityStatuses!: pulumi.Output<outputs.dataplex.AssetSecurityStatus[]>;
    /**
     * Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Output only. System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Output only. The time when the asset was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Asset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssetArgs | AssetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AssetState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["dataplexZone"] = state ? state.dataplexZone : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["discoverySpec"] = state ? state.discoverySpec : undefined;
            resourceInputs["discoveryStatuses"] = state ? state.discoveryStatuses : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["lake"] = state ? state.lake : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["resourceSpec"] = state ? state.resourceSpec : undefined;
            resourceInputs["resourceStatuses"] = state ? state.resourceStatuses : undefined;
            resourceInputs["securityStatuses"] = state ? state.securityStatuses : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as AssetArgs | undefined;
            if ((!args || args.dataplexZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataplexZone'");
            }
            if ((!args || args.discoverySpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'discoverySpec'");
            }
            if ((!args || args.lake === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lake'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.resourceSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceSpec'");
            }
            resourceInputs["dataplexZone"] = args ? args.dataplexZone : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["discoverySpec"] = args ? args.discoverySpec : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["lake"] = args ? args.lake : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["resourceSpec"] = args ? args.resourceSpec : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["discoveryStatuses"] = undefined /*out*/;
            resourceInputs["resourceStatuses"] = undefined /*out*/;
            resourceInputs["securityStatuses"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Asset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Asset resources.
 */
export interface AssetState {
    /**
     * Output only. The time when the asset was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The zone for the resource
     */
    dataplexZone?: pulumi.Input<string>;
    /**
     * Optional. Description of the asset.
     */
    description?: pulumi.Input<string>;
    /**
     * Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
     */
    discoverySpec?: pulumi.Input<inputs.dataplex.AssetDiscoverySpec>;
    /**
     * Output only. Status of the discovery feature applied to data referenced by this asset.
     */
    discoveryStatuses?: pulumi.Input<pulumi.Input<inputs.dataplex.AssetDiscoveryStatus>[]>;
    /**
     * Optional. User friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Optional. User defined labels for the asset.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The lake for the resource
     */
    lake?: pulumi.Input<string>;
    /**
     * The location for the resource
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the asset.
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Required. Immutable. Specification of the resource that is referenced by this asset.
     */
    resourceSpec?: pulumi.Input<inputs.dataplex.AssetResourceSpec>;
    /**
     * Output only. Status of the resource referenced by this asset.
     */
    resourceStatuses?: pulumi.Input<pulumi.Input<inputs.dataplex.AssetResourceStatus>[]>;
    /**
     * Output only. Status of the security policy applied to resource referenced by this asset.
     */
    securityStatuses?: pulumi.Input<pulumi.Input<inputs.dataplex.AssetSecurityStatus>[]>;
    /**
     * Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
     */
    state?: pulumi.Input<string>;
    /**
     * Output only. System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.
     */
    uid?: pulumi.Input<string>;
    /**
     * Output only. The time when the asset was last updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Asset resource.
 */
export interface AssetArgs {
    /**
     * The zone for the resource
     */
    dataplexZone: pulumi.Input<string>;
    /**
     * Optional. Description of the asset.
     */
    description?: pulumi.Input<string>;
    /**
     * Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
     */
    discoverySpec: pulumi.Input<inputs.dataplex.AssetDiscoverySpec>;
    /**
     * Optional. User friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Optional. User defined labels for the asset.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The lake for the resource
     */
    lake: pulumi.Input<string>;
    /**
     * The location for the resource
     */
    location: pulumi.Input<string>;
    /**
     * The name of the asset.
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Required. Immutable. Specification of the resource that is referenced by this asset.
     */
    resourceSpec: pulumi.Input<inputs.dataplex.AssetResourceSpec>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Catalogs are top-level containers for Databases and Tables.
 *
 * To get more information about Catalog, see:
 *
 * * [API documentation](https://cloud.google.com/bigquery/docs/reference/biglake/rest/v1/projects.locations.catalogs)
 * * How-to Guides
 *     * [Manage open source metadata with BigLake Metastore](https://cloud.google.com/bigquery/docs/manage-open-source-metadata#create_catalogs)
 *
 * ## Example Usage
 * ### Bigquery Biglake Catalog
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.biglake.Catalog("default", {location: "US"});
 * ```
 *
 * ## Import
 *
 * Catalog can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/catalogs/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Catalog can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:biglake/catalog:Catalog default projects/{{project}}/locations/{{location}}/catalogs/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:biglake/catalog:Catalog default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:biglake/catalog:Catalog default {{location}}/{{name}}
 * ```
 */
export class Catalog extends pulumi.CustomResource {
    /**
     * Get an existing Catalog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CatalogState, opts?: pulumi.CustomResourceOptions): Catalog {
        return new Catalog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:biglake/catalog:Catalog';

    /**
     * Returns true if the given object is an instance of Catalog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Catalog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Catalog.__pulumiType;
    }

    /**
     * Output only. The creation time of the catalog. A timestamp in RFC3339 UTC
     * "Zulu" format, with nanosecond resolution and up to nine fractional
     * digits.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Output only. The deletion time of the catalog. Only set after the catalog
     * is deleted. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
     * resolution and up to nine fractional digits.
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * Output only. The time when this catalog is considered expired. Only set
     * after the catalog is deleted. Only set after the catalog is deleted.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
     * up to nine fractional digits.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * The geographic location where the Catalog should reside.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Catalog. Format:
     * projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Output only. The last modification time of the catalog. A timestamp in
     * RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
     * fractional digits.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Catalog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CatalogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CatalogArgs | CatalogState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CatalogState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["deleteTime"] = state ? state.deleteTime : undefined;
            resourceInputs["expireTime"] = state ? state.expireTime : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as CatalogArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Catalog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Catalog resources.
 */
export interface CatalogState {
    /**
     * Output only. The creation time of the catalog. A timestamp in RFC3339 UTC
     * "Zulu" format, with nanosecond resolution and up to nine fractional
     * digits.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Output only. The deletion time of the catalog. Only set after the catalog
     * is deleted. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
     * resolution and up to nine fractional digits.
     */
    deleteTime?: pulumi.Input<string>;
    /**
     * Output only. The time when this catalog is considered expired. Only set
     * after the catalog is deleted. Only set after the catalog is deleted.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
     * up to nine fractional digits.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * The geographic location where the Catalog should reside.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Catalog. Format:
     * projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Output only. The last modification time of the catalog. A timestamp in
     * RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
     * fractional digits.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Catalog resource.
 */
export interface CatalogArgs {
    /**
     * The geographic location where the Catalog should reside.
     */
    location: pulumi.Input<string>;
    /**
     * The name of the Catalog. Format:
     * projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

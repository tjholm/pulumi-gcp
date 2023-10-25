// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A collection of policy tags that classify data along a common axis.
 *
 * To get more information about Taxonomy, see:
 *
 * * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.taxonomies)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
 *
 * ## Example Usage
 * ### Data Catalog Taxonomy Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basicTaxonomy = new gcp.datacatalog.Taxonomy("basicTaxonomy", {
 *     activatedPolicyTypes: ["FINE_GRAINED_ACCESS_CONTROL"],
 *     description: "A collection of policy tags",
 *     displayName: "my_taxonomy",
 * });
 * ```
 *
 * ## Import
 *
 * Taxonomy can be imported using any of these accepted formats:
 *
 * ```sh
 *  $ pulumi import gcp:datacatalog/taxonomy:Taxonomy default {{name}}
 * ```
 */
export class Taxonomy extends pulumi.CustomResource {
    /**
     * Get an existing Taxonomy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TaxonomyState, opts?: pulumi.CustomResourceOptions): Taxonomy {
        return new Taxonomy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:datacatalog/taxonomy:Taxonomy';

    /**
     * Returns true if the given object is an instance of Taxonomy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Taxonomy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Taxonomy.__pulumiType;
    }

    /**
     * A list of policy types that are activated for this taxonomy. If not set,
     * defaults to an empty list.
     * Each value may be one of: `POLICY_TYPE_UNSPECIFIED`, `FINE_GRAINED_ACCESS_CONTROL`.
     */
    public readonly activatedPolicyTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Description of this taxonomy. It must: contain only unicode characters,
     * tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
     * long when encoded in UTF-8. If not set, defaults to an empty description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User defined name of this taxonomy.
     * It must: contain only unicode letters, numbers, underscores, dashes
     * and spaces; not start or end with spaces; and be at most 200 bytes
     * long when encoded in UTF-8.
     *
     *
     * - - -
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Resource name of this taxonomy, whose format is:
     * "projects/{project}/locations/{region}/taxonomies/{taxonomy}".
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Taxonomy location region.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a Taxonomy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TaxonomyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TaxonomyArgs | TaxonomyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TaxonomyState | undefined;
            resourceInputs["activatedPolicyTypes"] = state ? state.activatedPolicyTypes : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as TaxonomyArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["activatedPolicyTypes"] = args ? args.activatedPolicyTypes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Taxonomy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Taxonomy resources.
 */
export interface TaxonomyState {
    /**
     * A list of policy types that are activated for this taxonomy. If not set,
     * defaults to an empty list.
     * Each value may be one of: `POLICY_TYPE_UNSPECIFIED`, `FINE_GRAINED_ACCESS_CONTROL`.
     */
    activatedPolicyTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of this taxonomy. It must: contain only unicode characters,
     * tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
     * long when encoded in UTF-8. If not set, defaults to an empty description.
     */
    description?: pulumi.Input<string>;
    /**
     * User defined name of this taxonomy.
     * It must: contain only unicode letters, numbers, underscores, dashes
     * and spaces; not start or end with spaces; and be at most 200 bytes
     * long when encoded in UTF-8.
     *
     *
     * - - -
     */
    displayName?: pulumi.Input<string>;
    /**
     * Resource name of this taxonomy, whose format is:
     * "projects/{project}/locations/{region}/taxonomies/{taxonomy}".
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Taxonomy location region.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Taxonomy resource.
 */
export interface TaxonomyArgs {
    /**
     * A list of policy types that are activated for this taxonomy. If not set,
     * defaults to an empty list.
     * Each value may be one of: `POLICY_TYPE_UNSPECIFIED`, `FINE_GRAINED_ACCESS_CONTROL`.
     */
    activatedPolicyTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of this taxonomy. It must: contain only unicode characters,
     * tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
     * long when encoded in UTF-8. If not set, defaults to an empty description.
     */
    description?: pulumi.Input<string>;
    /**
     * User defined name of this taxonomy.
     * It must: contain only unicode letters, numbers, underscores, dashes
     * and spaces; not start or end with spaces; and be at most 200 bytes
     * long when encoded in UTF-8.
     *
     *
     * - - -
     */
    displayName: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Taxonomy location region.
     */
    region?: pulumi.Input<string>;
}

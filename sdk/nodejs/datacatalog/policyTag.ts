// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Denotes one policy tag in a taxonomy.
 *
 * To get more information about PolicyTag, see:
 *
 * * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1beta1/projects.locations.taxonomies.policyTags)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
 *
 * ## Example Usage
 * ### Data Catalog Taxonomies Policy Tag Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myTaxonomy = new gcp.datacatalog.Taxonomy("myTaxonomy", {
 *     region: "us",
 *     displayName: "taxonomy_display_name",
 *     description: "A collection of policy tags",
 *     activatedPolicyTypes: ["FINE_GRAINED_ACCESS_CONTROL"],
 * }, {
 *     provider: google_beta,
 * });
 * const basicPolicyTag = new gcp.datacatalog.PolicyTag("basicPolicyTag", {
 *     taxonomy: myTaxonomy.id,
 *     displayName: "Low security",
 *     description: "A policy tag normally associated with low security items",
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Data Catalog Taxonomies Policy Tag Child Policies
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myTaxonomy = new gcp.datacatalog.Taxonomy("myTaxonomy", {
 *     region: "us",
 *     displayName: "taxonomy_display_name",
 *     description: "A collection of policy tags",
 *     activatedPolicyTypes: ["FINE_GRAINED_ACCESS_CONTROL"],
 * }, {
 *     provider: google_beta,
 * });
 * const parentPolicy = new gcp.datacatalog.PolicyTag("parentPolicy", {
 *     taxonomy: myTaxonomy.id,
 *     displayName: "High",
 *     description: "A policy tag category used for high security access",
 * }, {
 *     provider: google_beta,
 * });
 * const childPolicy = new gcp.datacatalog.PolicyTag("childPolicy", {
 *     taxonomy: myTaxonomy.id,
 *     displayName: "ssn",
 *     description: "A hash of the users ssn",
 *     parentPolicyTag: parentPolicy.id,
 * }, {
 *     provider: google_beta,
 * });
 * const childPolicy2 = new gcp.datacatalog.PolicyTag("childPolicy2", {
 *     taxonomy: myTaxonomy.id,
 *     displayName: "dob",
 *     description: "The users date of birth",
 *     parentPolicyTag: parentPolicy.id,
 * }, {
 *     provider: google_beta,
 *     dependsOn: [childPolicy],
 * });
 * ```
 *
 * ## Import
 *
 * PolicyTag can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:datacatalog/policyTag:PolicyTag default {{name}}
 * ```
 */
export class PolicyTag extends pulumi.CustomResource {
    /**
     * Get an existing PolicyTag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyTagState, opts?: pulumi.CustomResourceOptions): PolicyTag {
        return new PolicyTag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:datacatalog/policyTag:PolicyTag';

    /**
     * Returns true if the given object is an instance of PolicyTag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyTag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyTag.__pulumiType;
    }

    /**
     * Resource names of child policy tags of this policy tag.
     */
    public /*out*/ readonly childPolicyTags!: pulumi.Output<string[]>;
    /**
     * Description of this policy tag. It must: contain only unicode characters, tabs,
     * newlines, carriage returns and page breaks; and be at most 2000 bytes long when
     * encoded in UTF-8. If not set, defaults to an empty description.
     * If not set, defaults to an empty description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User defined name of this policy tag. It must: be unique within the parent
     * taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
     * not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Resource name of this policy tag, whose format is:
     * "projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}"
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource name of this policy tag's parent policy tag.
     * If empty, it means this policy tag is a top level policy tag.
     * If not set, defaults to an empty string.
     */
    public readonly parentPolicyTag!: pulumi.Output<string | undefined>;
    /**
     * Taxonomy the policy tag is associated with
     */
    public readonly taxonomy!: pulumi.Output<string>;

    /**
     * Create a PolicyTag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyTagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyTagArgs | PolicyTagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyTagState | undefined;
            resourceInputs["childPolicyTags"] = state ? state.childPolicyTags : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentPolicyTag"] = state ? state.parentPolicyTag : undefined;
            resourceInputs["taxonomy"] = state ? state.taxonomy : undefined;
        } else {
            const args = argsOrState as PolicyTagArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.taxonomy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taxonomy'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["parentPolicyTag"] = args ? args.parentPolicyTag : undefined;
            resourceInputs["taxonomy"] = args ? args.taxonomy : undefined;
            resourceInputs["childPolicyTags"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PolicyTag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyTag resources.
 */
export interface PolicyTagState {
    /**
     * Resource names of child policy tags of this policy tag.
     */
    childPolicyTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of this policy tag. It must: contain only unicode characters, tabs,
     * newlines, carriage returns and page breaks; and be at most 2000 bytes long when
     * encoded in UTF-8. If not set, defaults to an empty description.
     * If not set, defaults to an empty description.
     */
    description?: pulumi.Input<string>;
    /**
     * User defined name of this policy tag. It must: be unique within the parent
     * taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
     * not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Resource name of this policy tag, whose format is:
     * "projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}"
     */
    name?: pulumi.Input<string>;
    /**
     * Resource name of this policy tag's parent policy tag.
     * If empty, it means this policy tag is a top level policy tag.
     * If not set, defaults to an empty string.
     */
    parentPolicyTag?: pulumi.Input<string>;
    /**
     * Taxonomy the policy tag is associated with
     */
    taxonomy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyTag resource.
 */
export interface PolicyTagArgs {
    /**
     * Description of this policy tag. It must: contain only unicode characters, tabs,
     * newlines, carriage returns and page breaks; and be at most 2000 bytes long when
     * encoded in UTF-8. If not set, defaults to an empty description.
     * If not set, defaults to an empty description.
     */
    description?: pulumi.Input<string>;
    /**
     * User defined name of this policy tag. It must: be unique within the parent
     * taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
     * not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
     */
    displayName: pulumi.Input<string>;
    /**
     * Resource name of this policy tag's parent policy tag.
     * If empty, it means this policy tag is a top level policy tag.
     * If not set, defaults to an empty string.
     */
    parentPolicyTag?: pulumi.Input<string>;
    /**
     * Taxonomy the policy tag is associated with
     */
    taxonomy: pulumi.Input<string>;
}

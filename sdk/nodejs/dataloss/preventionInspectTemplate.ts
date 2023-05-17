// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * An inspect job template.
 *
 * To get more information about InspectTemplate, see:
 *
 * * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.inspectTemplates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dlp/docs/creating-templates-inspect)
 *
 * ## Example Usage
 * ### Dlp Inspect Template Custom Type
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const custom = new gcp.dataloss.PreventionInspectTemplate("custom", {
 *     description: "My description",
 *     displayName: "display_name",
 *     inspectConfig: {
 *         customInfoTypes: [{
 *             infoType: {
 *                 name: "MY_CUSTOM_TYPE",
 *             },
 *             likelihood: "UNLIKELY",
 *             regex: {
 *                 pattern: "test*",
 *             },
 *         }],
 *         infoTypes: [{
 *             name: "EMAIL_ADDRESS",
 *         }],
 *         limits: {
 *             maxFindingsPerItem: 10,
 *             maxFindingsPerRequest: 50,
 *         },
 *         minLikelihood: "UNLIKELY",
 *         ruleSets: [
 *             {
 *                 infoTypes: [{
 *                     name: "EMAIL_ADDRESS",
 *                 }],
 *                 rules: [{
 *                     exclusionRule: {
 *                         matchingType: "MATCHING_TYPE_FULL_MATCH",
 *                         regex: {
 *                             pattern: ".+@example.com",
 *                         },
 *                     },
 *                 }],
 *             },
 *             {
 *                 infoTypes: [{
 *                     name: "MY_CUSTOM_TYPE",
 *                 }],
 *                 rules: [{
 *                     hotwordRule: {
 *                         hotwordRegex: {
 *                             pattern: "example*",
 *                         },
 *                         likelihoodAdjustment: {
 *                             fixedLikelihood: "VERY_LIKELY",
 *                         },
 *                         proximity: {
 *                             windowBefore: 50,
 *                         },
 *                     },
 *                 }],
 *             },
 *         ],
 *     },
 *     parent: "projects/my-project-name",
 * });
 * ```
 * ### Dlp Inspect Template Custom Type Surrogate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const customTypeSurrogate = new gcp.dataloss.PreventionInspectTemplate("customTypeSurrogate", {
 *     description: "My description",
 *     displayName: "display_name",
 *     inspectConfig: {
 *         customInfoTypes: [{
 *             infoType: {
 *                 name: "MY_CUSTOM_TYPE",
 *             },
 *             likelihood: "UNLIKELY",
 *             surrogateType: {},
 *         }],
 *         infoTypes: [{
 *             name: "EMAIL_ADDRESS",
 *         }],
 *         limits: {
 *             maxFindingsPerItem: 10,
 *             maxFindingsPerRequest: 50,
 *         },
 *         minLikelihood: "UNLIKELY",
 *         ruleSets: [
 *             {
 *                 infoTypes: [{
 *                     name: "EMAIL_ADDRESS",
 *                 }],
 *                 rules: [{
 *                     exclusionRule: {
 *                         matchingType: "MATCHING_TYPE_FULL_MATCH",
 *                         regex: {
 *                             pattern: ".+@example.com",
 *                         },
 *                     },
 *                 }],
 *             },
 *             {
 *                 infoTypes: [{
 *                     name: "MY_CUSTOM_TYPE",
 *                 }],
 *                 rules: [{
 *                     hotwordRule: {
 *                         hotwordRegex: {
 *                             pattern: "example*",
 *                         },
 *                         likelihoodAdjustment: {
 *                             fixedLikelihood: "VERY_LIKELY",
 *                         },
 *                         proximity: {
 *                             windowBefore: 50,
 *                         },
 *                     },
 *                 }],
 *             },
 *         ],
 *     },
 *     parent: "projects/my-project-name",
 * });
 * ```
 *
 * ## Import
 *
 * InspectTemplate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:dataloss/preventionInspectTemplate:PreventionInspectTemplate default {{parent}}/inspectTemplates/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataloss/preventionInspectTemplate:PreventionInspectTemplate default {{parent}}/{{name}}
 * ```
 */
export class PreventionInspectTemplate extends pulumi.CustomResource {
    /**
     * Get an existing PreventionInspectTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PreventionInspectTemplateState, opts?: pulumi.CustomResourceOptions): PreventionInspectTemplate {
        return new PreventionInspectTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataloss/preventionInspectTemplate:PreventionInspectTemplate';

    /**
     * Returns true if the given object is an instance of PreventionInspectTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PreventionInspectTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PreventionInspectTemplate.__pulumiType;
    }

    /**
     * A description of the inspect template.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User set display name of the inspect template.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The core content of the template.
     * Structure is documented below.
     */
    public readonly inspectConfig!: pulumi.Output<outputs.dataloss.PreventionInspectTemplateInspectConfig | undefined>;
    /**
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * (Required)
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * (Required)
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * (Required)
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * (Required)
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names
     * listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * (Required)
     * Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
     * or `projects/project-id/storedInfoTypes/432452342`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent of the inspect template in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     */
    public readonly parent!: pulumi.Output<string>;

    /**
     * Create a PreventionInspectTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PreventionInspectTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PreventionInspectTemplateArgs | PreventionInspectTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PreventionInspectTemplateState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["inspectConfig"] = state ? state.inspectConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
        } else {
            const args = argsOrState as PreventionInspectTemplateArgs | undefined;
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["inspectConfig"] = args ? args.inspectConfig : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PreventionInspectTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PreventionInspectTemplate resources.
 */
export interface PreventionInspectTemplateState {
    /**
     * A description of the inspect template.
     */
    description?: pulumi.Input<string>;
    /**
     * User set display name of the inspect template.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The core content of the template.
     * Structure is documented below.
     */
    inspectConfig?: pulumi.Input<inputs.dataloss.PreventionInspectTemplateInspectConfig>;
    /**
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * (Required)
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * (Required)
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * (Required)
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * (Required)
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names
     * listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * (Required)
     * Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
     * or `projects/project-id/storedInfoTypes/432452342`.
     */
    name?: pulumi.Input<string>;
    /**
     * The parent of the inspect template in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     */
    parent?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PreventionInspectTemplate resource.
 */
export interface PreventionInspectTemplateArgs {
    /**
     * A description of the inspect template.
     */
    description?: pulumi.Input<string>;
    /**
     * User set display name of the inspect template.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The core content of the template.
     * Structure is documented below.
     */
    inspectConfig?: pulumi.Input<inputs.dataloss.PreventionInspectTemplateInspectConfig>;
    /**
     * The parent of the inspect template in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     */
    parent: pulumi.Input<string>;
}

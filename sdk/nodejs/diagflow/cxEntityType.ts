// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Entities are extracted from user input and represent parameters that are meaningful to your application.
 * For example, a date range, a proper name such as a geographic location or landmark, and so on. Entities represent actionable data for your application.
 *
 * To get more information about EntityType, see:
 *
 * * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.entityTypes)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
 *
 * ## Example Usage
 * ### Dialogflowcx Entity Type Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const agent = new gcp.diagflow.CxAgent("agent", {
 *     displayName: "dialogflowcx-agent",
 *     location: "global",
 *     defaultLanguageCode: "en",
 *     supportedLanguageCodes: [
 *         "fr",
 *         "de",
 *         "es",
 *     ],
 *     timeZone: "America/New_York",
 *     description: "Example description.",
 *     avatarUri: "https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png",
 *     enableStackdriverLogging: true,
 *     enableSpellCorrection: true,
 *     speechToTextSettings: {
 *         enableSpeechAdaptation: true,
 *     },
 * });
 * const basicEntityType = new gcp.diagflow.CxEntityType("basicEntityType", {
 *     parent: agent.id,
 *     displayName: "MyEntity",
 *     kind: "KIND_MAP",
 *     entities: [
 *         {
 *             value: "value1",
 *             synonyms: [
 *                 "synonym1",
 *                 "synonym2",
 *             ],
 *         },
 *         {
 *             value: "value2",
 *             synonyms: [
 *                 "synonym3",
 *                 "synonym4",
 *             ],
 *         },
 *     ],
 *     enableFuzzyExtraction: false,
 * });
 * ```
 *
 * ## Import
 *
 * EntityType can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:diagflow/cxEntityType:CxEntityType default {{parent}}/entityTypes/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:diagflow/cxEntityType:CxEntityType default {{parent}}/{{name}}
 * ```
 */
export class CxEntityType extends pulumi.CustomResource {
    /**
     * Get an existing CxEntityType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CxEntityTypeState, opts?: pulumi.CustomResourceOptions): CxEntityType {
        return new CxEntityType(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:diagflow/cxEntityType:CxEntityType';

    /**
     * Returns true if the given object is an instance of CxEntityType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CxEntityType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CxEntityType.__pulumiType;
    }

    /**
     * Represents kinds of entities.
     * * AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
     * * AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity.
     * Possible values are: `AUTO_EXPANSION_MODE_DEFAULT`, `AUTO_EXPANSION_MODE_UNSPECIFIED`.
     */
    public readonly autoExpansionMode!: pulumi.Output<string | undefined>;
    /**
     * The human-readable name of the entity type, unique within the agent.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Enables fuzzy entity extraction during classification.
     */
    public readonly enableFuzzyExtraction!: pulumi.Output<boolean | undefined>;
    /**
     * The collection of entity entries associated with the entity type.
     * Structure is documented below.
     */
    public readonly entities!: pulumi.Output<outputs.diagflow.CxEntityTypeEntity[]>;
    /**
     * Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
     * If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
     * Structure is documented below.
     */
    public readonly excludedPhrases!: pulumi.Output<outputs.diagflow.CxEntityTypeExcludedPhrase[] | undefined>;
    /**
     * Indicates whether the entity type can be automatically expanded.
     * * KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
     * * KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
     * * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
     * Possible values are: `KIND_MAP`, `KIND_LIST`, `KIND_REGEXP`.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The language of the following fields in entityType:
     * EntityType.entities.value
     * EntityType.entities.synonyms
     * EntityType.excluded_phrases.value
     * If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
     */
    public readonly languageCode!: pulumi.Output<string | undefined>;
    /**
     * The unique identifier of the entity type.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID>.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The agent to create a entity type for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
     */
    public readonly parent!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
     */
    public readonly redact!: pulumi.Output<boolean | undefined>;

    /**
     * Create a CxEntityType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CxEntityTypeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CxEntityTypeArgs | CxEntityTypeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CxEntityTypeState | undefined;
            resourceInputs["autoExpansionMode"] = state ? state.autoExpansionMode : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["enableFuzzyExtraction"] = state ? state.enableFuzzyExtraction : undefined;
            resourceInputs["entities"] = state ? state.entities : undefined;
            resourceInputs["excludedPhrases"] = state ? state.excludedPhrases : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["languageCode"] = state ? state.languageCode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["redact"] = state ? state.redact : undefined;
        } else {
            const args = argsOrState as CxEntityTypeArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.entities === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entities'");
            }
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            resourceInputs["autoExpansionMode"] = args ? args.autoExpansionMode : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enableFuzzyExtraction"] = args ? args.enableFuzzyExtraction : undefined;
            resourceInputs["entities"] = args ? args.entities : undefined;
            resourceInputs["excludedPhrases"] = args ? args.excludedPhrases : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["languageCode"] = args ? args.languageCode : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["redact"] = args ? args.redact : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CxEntityType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CxEntityType resources.
 */
export interface CxEntityTypeState {
    /**
     * Represents kinds of entities.
     * * AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
     * * AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity.
     * Possible values are: `AUTO_EXPANSION_MODE_DEFAULT`, `AUTO_EXPANSION_MODE_UNSPECIFIED`.
     */
    autoExpansionMode?: pulumi.Input<string>;
    /**
     * The human-readable name of the entity type, unique within the agent.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Enables fuzzy entity extraction during classification.
     */
    enableFuzzyExtraction?: pulumi.Input<boolean>;
    /**
     * The collection of entity entries associated with the entity type.
     * Structure is documented below.
     */
    entities?: pulumi.Input<pulumi.Input<inputs.diagflow.CxEntityTypeEntity>[]>;
    /**
     * Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
     * If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
     * Structure is documented below.
     */
    excludedPhrases?: pulumi.Input<pulumi.Input<inputs.diagflow.CxEntityTypeExcludedPhrase>[]>;
    /**
     * Indicates whether the entity type can be automatically expanded.
     * * KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
     * * KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
     * * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
     * Possible values are: `KIND_MAP`, `KIND_LIST`, `KIND_REGEXP`.
     */
    kind?: pulumi.Input<string>;
    /**
     * The language of the following fields in entityType:
     * EntityType.entities.value
     * EntityType.entities.synonyms
     * EntityType.excluded_phrases.value
     * If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
     */
    languageCode?: pulumi.Input<string>;
    /**
     * The unique identifier of the entity type.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID>.
     */
    name?: pulumi.Input<string>;
    /**
     * The agent to create a entity type for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
     */
    parent?: pulumi.Input<string>;
    /**
     * Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
     */
    redact?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a CxEntityType resource.
 */
export interface CxEntityTypeArgs {
    /**
     * Represents kinds of entities.
     * * AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
     * * AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity.
     * Possible values are: `AUTO_EXPANSION_MODE_DEFAULT`, `AUTO_EXPANSION_MODE_UNSPECIFIED`.
     */
    autoExpansionMode?: pulumi.Input<string>;
    /**
     * The human-readable name of the entity type, unique within the agent.
     */
    displayName: pulumi.Input<string>;
    /**
     * Enables fuzzy entity extraction during classification.
     */
    enableFuzzyExtraction?: pulumi.Input<boolean>;
    /**
     * The collection of entity entries associated with the entity type.
     * Structure is documented below.
     */
    entities: pulumi.Input<pulumi.Input<inputs.diagflow.CxEntityTypeEntity>[]>;
    /**
     * Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
     * If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
     * Structure is documented below.
     */
    excludedPhrases?: pulumi.Input<pulumi.Input<inputs.diagflow.CxEntityTypeExcludedPhrase>[]>;
    /**
     * Indicates whether the entity type can be automatically expanded.
     * * KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
     * * KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
     * * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
     * Possible values are: `KIND_MAP`, `KIND_LIST`, `KIND_REGEXP`.
     */
    kind: pulumi.Input<string>;
    /**
     * The language of the following fields in entityType:
     * EntityType.entities.value
     * EntityType.entities.synonyms
     * EntityType.excluded_phrases.value
     * If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
     */
    languageCode?: pulumi.Input<string>;
    /**
     * The agent to create a entity type for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
     */
    parent?: pulumi.Input<string>;
    /**
     * Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
     */
    redact?: pulumi.Input<boolean>;
}

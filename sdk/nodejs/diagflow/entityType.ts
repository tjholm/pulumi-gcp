// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Represents an entity type. Entity types serve as a tool for extracting parameter values from natural language queries.
 *
 * To get more information about EntityType, see:
 *
 * * [API documentation](https://cloud.google.com/dialogflow/docs/reference/rest/v2/projects.agent.entityTypes)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dialogflow/docs/)
 *
 * ## Example Usage
 * ### Dialogflow Entity Type Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basicAgent = new gcp.diagflow.Agent("basicAgent", {
 *     displayName: "example_agent",
 *     defaultLanguageCode: "en",
 *     timeZone: "America/New_York",
 * });
 * const basicEntityType = new gcp.diagflow.EntityType("basicEntityType", {
 *     displayName: "",
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
 * }, {
 *     dependsOn: [basicAgent],
 * });
 * ```
 *
 * ## Import
 *
 * EntityType can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:diagflow/entityType:EntityType default {{name}}
 * ```
 */
export class EntityType extends pulumi.CustomResource {
    /**
     * Get an existing EntityType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EntityTypeState, opts?: pulumi.CustomResourceOptions): EntityType {
        return new EntityType(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:diagflow/entityType:EntityType';

    /**
     * Returns true if the given object is an instance of EntityType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EntityType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EntityType.__pulumiType;
    }

    /**
     * The name of this entity type to be displayed on the console.
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
    public readonly entities!: pulumi.Output<outputs.diagflow.EntityTypeEntity[] | undefined>;
    /**
     * Indicates the kind of entity type.
     * * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
     * * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
     * types can contain references to other entity types (with or without aliases).
     * * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
     * Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a EntityType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntityTypeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EntityTypeArgs | EntityTypeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EntityTypeState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["enableFuzzyExtraction"] = state ? state.enableFuzzyExtraction : undefined;
            resourceInputs["entities"] = state ? state.entities : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as EntityTypeArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enableFuzzyExtraction"] = args ? args.enableFuzzyExtraction : undefined;
            resourceInputs["entities"] = args ? args.entities : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EntityType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EntityType resources.
 */
export interface EntityTypeState {
    /**
     * The name of this entity type to be displayed on the console.
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
    entities?: pulumi.Input<pulumi.Input<inputs.diagflow.EntityTypeEntity>[]>;
    /**
     * Indicates the kind of entity type.
     * * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
     * * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
     * types can contain references to other entity types (with or without aliases).
     * * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
     * Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
     */
    kind?: pulumi.Input<string>;
    /**
     * The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EntityType resource.
 */
export interface EntityTypeArgs {
    /**
     * The name of this entity type to be displayed on the console.
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
    entities?: pulumi.Input<pulumi.Input<inputs.diagflow.EntityTypeEntity>[]>;
    /**
     * Indicates the kind of entity type.
     * * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
     * * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
     * types can contain references to other entity types (with or without aliases).
     * * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
     * Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
     */
    kind: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

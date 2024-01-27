// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * An intent represents a user's intent to interact with a conversational agent.
 *
 * To get more information about Intent, see:
 *
 * * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.intents)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
 *
 * ## Example Usage
 * ### Dialogflowcx Intent Full
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
 * const basicIntent = new gcp.diagflow.CxIntent("basicIntent", {
 *     parent: agent.id,
 *     displayName: "Example",
 *     priority: 1,
 *     description: "Intent example",
 *     trainingPhrases: [{
 *         parts: [
 *             {
 *                 text: "training",
 *             },
 *             {
 *                 text: "phrase",
 *             },
 *             {
 *                 text: "example",
 *             },
 *         ],
 *         repeatCount: 1,
 *     }],
 *     parameters: [{
 *         id: "param1",
 *         entityType: "projects/-/locations/-/agents/-/entityTypes/sys.date",
 *     }],
 *     labels: {
 *         label1: "value1",
 *         label2: "value2",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Intent can be imported using any of these accepted formats* `{{parent}}/intents/{{name}}` * `{{parent}}/{{name}}` When using the `pulumi import` command, Intent can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:diagflow/cxIntent:CxIntent default {{parent}}/intents/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:diagflow/cxIntent:CxIntent default {{parent}}/{{name}}
 * ```
 */
export class CxIntent extends pulumi.CustomResource {
    /**
     * Get an existing CxIntent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CxIntentState, opts?: pulumi.CustomResourceOptions): CxIntent {
        return new CxIntent(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:diagflow/cxIntent:CxIntent';

    /**
     * Returns true if the given object is an instance of CxIntent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CxIntent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CxIntent.__pulumiType;
    }

    /**
     * Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The human-readable name of the intent, unique within the agent.
     *
     *
     * - - -
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Marks this as the [Default Negative Intent](https://cloud.google.com/dialogflow/cx/docs/concept/intent#negative) for an agent. When you create an agent, a Default Negative Intent is created automatically.
     * The Default Negative Intent cannot be deleted; deleting the `gcp.diagflow.CxIntent` resource does nothing to the underlying GCP resources.
     *
     * > Avoid having multiple `gcp.diagflow.CxIntent` resources linked to the same agent with `isDefaultNegativeIntent = true` because they will compete to control a single Default Negative Intent resource in GCP.
     */
    public readonly isDefaultNegativeIntent!: pulumi.Output<boolean | undefined>;
    /**
     * Marks this as the [Default Welcome Intent](https://cloud.google.com/dialogflow/cx/docs/concept/intent#welcome) for an agent. When you create an agent, a Default Welcome Intent is created automatically.
     * The Default Welcome Intent cannot be deleted; deleting the `gcp.diagflow.CxIntent` resource does nothing to the underlying GCP resources.
     *
     * > Avoid having multiple `gcp.diagflow.CxIntent` resources linked to the same agent with `isDefaultWelcomeIntent = true` because they will compete to control a single Default Welcome Intent resource in GCP.
     */
    public readonly isDefaultWelcomeIntent!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
     * Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
     * To manage the fallback intent, set `isDefaultNegativeIntent = true`
     */
    public readonly isFallback!: pulumi.Output<boolean | undefined>;
    /**
     * The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
     * Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The language of the following fields in intent:
     * Intent.training_phrases.parts.text
     * If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
     */
    public readonly languageCode!: pulumi.Output<string | undefined>;
    /**
     * The unique identifier of the intent.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The collection of parameters associated with the intent.
     * Structure is documented below.
     */
    public readonly parameters!: pulumi.Output<outputs.diagflow.CxIntentParameter[] | undefined>;
    /**
     * The agent to create an intent for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
     */
    public readonly parent!: pulumi.Output<string | undefined>;
    /**
     * The priority of this intent. Higher numbers represent higher priorities.
     * If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
     * If the supplied value is negative, the intent is ignored in runtime detect intent requests.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The collection of training phrases the agent is trained on to identify the intent.
     * Structure is documented below.
     */
    public readonly trainingPhrases!: pulumi.Output<outputs.diagflow.CxIntentTrainingPhrase[] | undefined>;

    /**
     * Create a CxIntent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CxIntentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CxIntentArgs | CxIntentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CxIntentState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["isDefaultNegativeIntent"] = state ? state.isDefaultNegativeIntent : undefined;
            resourceInputs["isDefaultWelcomeIntent"] = state ? state.isDefaultWelcomeIntent : undefined;
            resourceInputs["isFallback"] = state ? state.isFallback : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["languageCode"] = state ? state.languageCode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["trainingPhrases"] = state ? state.trainingPhrases : undefined;
        } else {
            const args = argsOrState as CxIntentArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["isDefaultNegativeIntent"] = args ? args.isDefaultNegativeIntent : undefined;
            resourceInputs["isDefaultWelcomeIntent"] = args ? args.isDefaultWelcomeIntent : undefined;
            resourceInputs["isFallback"] = args ? args.isFallback : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["languageCode"] = args ? args.languageCode : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["trainingPhrases"] = args ? args.trainingPhrases : undefined;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(CxIntent.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CxIntent resources.
 */
export interface CxIntentState {
    /**
     * Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name of the intent, unique within the agent.
     *
     *
     * - - -
     */
    displayName?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Marks this as the [Default Negative Intent](https://cloud.google.com/dialogflow/cx/docs/concept/intent#negative) for an agent. When you create an agent, a Default Negative Intent is created automatically.
     * The Default Negative Intent cannot be deleted; deleting the `gcp.diagflow.CxIntent` resource does nothing to the underlying GCP resources.
     *
     * > Avoid having multiple `gcp.diagflow.CxIntent` resources linked to the same agent with `isDefaultNegativeIntent = true` because they will compete to control a single Default Negative Intent resource in GCP.
     */
    isDefaultNegativeIntent?: pulumi.Input<boolean>;
    /**
     * Marks this as the [Default Welcome Intent](https://cloud.google.com/dialogflow/cx/docs/concept/intent#welcome) for an agent. When you create an agent, a Default Welcome Intent is created automatically.
     * The Default Welcome Intent cannot be deleted; deleting the `gcp.diagflow.CxIntent` resource does nothing to the underlying GCP resources.
     *
     * > Avoid having multiple `gcp.diagflow.CxIntent` resources linked to the same agent with `isDefaultWelcomeIntent = true` because they will compete to control a single Default Welcome Intent resource in GCP.
     */
    isDefaultWelcomeIntent?: pulumi.Input<boolean>;
    /**
     * Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
     * Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
     * To manage the fallback intent, set `isDefaultNegativeIntent = true`
     */
    isFallback?: pulumi.Input<boolean>;
    /**
     * The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
     * Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The language of the following fields in intent:
     * Intent.training_phrases.parts.text
     * If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
     */
    languageCode?: pulumi.Input<string>;
    /**
     * The unique identifier of the intent.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>.
     */
    name?: pulumi.Input<string>;
    /**
     * The collection of parameters associated with the intent.
     * Structure is documented below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.diagflow.CxIntentParameter>[]>;
    /**
     * The agent to create an intent for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
     */
    parent?: pulumi.Input<string>;
    /**
     * The priority of this intent. Higher numbers represent higher priorities.
     * If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
     * If the supplied value is negative, the intent is ignored in runtime detect intent requests.
     */
    priority?: pulumi.Input<number>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The collection of training phrases the agent is trained on to identify the intent.
     * Structure is documented below.
     */
    trainingPhrases?: pulumi.Input<pulumi.Input<inputs.diagflow.CxIntentTrainingPhrase>[]>;
}

/**
 * The set of arguments for constructing a CxIntent resource.
 */
export interface CxIntentArgs {
    /**
     * Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name of the intent, unique within the agent.
     *
     *
     * - - -
     */
    displayName: pulumi.Input<string>;
    /**
     * Marks this as the [Default Negative Intent](https://cloud.google.com/dialogflow/cx/docs/concept/intent#negative) for an agent. When you create an agent, a Default Negative Intent is created automatically.
     * The Default Negative Intent cannot be deleted; deleting the `gcp.diagflow.CxIntent` resource does nothing to the underlying GCP resources.
     *
     * > Avoid having multiple `gcp.diagflow.CxIntent` resources linked to the same agent with `isDefaultNegativeIntent = true` because they will compete to control a single Default Negative Intent resource in GCP.
     */
    isDefaultNegativeIntent?: pulumi.Input<boolean>;
    /**
     * Marks this as the [Default Welcome Intent](https://cloud.google.com/dialogflow/cx/docs/concept/intent#welcome) for an agent. When you create an agent, a Default Welcome Intent is created automatically.
     * The Default Welcome Intent cannot be deleted; deleting the `gcp.diagflow.CxIntent` resource does nothing to the underlying GCP resources.
     *
     * > Avoid having multiple `gcp.diagflow.CxIntent` resources linked to the same agent with `isDefaultWelcomeIntent = true` because they will compete to control a single Default Welcome Intent resource in GCP.
     */
    isDefaultWelcomeIntent?: pulumi.Input<boolean>;
    /**
     * Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
     * Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
     * To manage the fallback intent, set `isDefaultNegativeIntent = true`
     */
    isFallback?: pulumi.Input<boolean>;
    /**
     * The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
     * Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The language of the following fields in intent:
     * Intent.training_phrases.parts.text
     * If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
     */
    languageCode?: pulumi.Input<string>;
    /**
     * The collection of parameters associated with the intent.
     * Structure is documented below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.diagflow.CxIntentParameter>[]>;
    /**
     * The agent to create an intent for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
     */
    parent?: pulumi.Input<string>;
    /**
     * The priority of this intent. Higher numbers represent higher priorities.
     * If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
     * If the supplied value is negative, the intent is ignored in runtime detect intent requests.
     */
    priority?: pulumi.Input<number>;
    /**
     * The collection of training phrases the agent is trained on to identify the intent.
     * Structure is documented below.
     */
    trainingPhrases?: pulumi.Input<pulumi.Input<inputs.diagflow.CxIntentTrainingPhrase>[]>;
}

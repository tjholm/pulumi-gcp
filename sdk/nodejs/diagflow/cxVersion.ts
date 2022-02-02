// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * You can create multiple versions of your agent flows and deploy them to separate serving environments.
 * When you edit a flow, you are editing a draft flow. At any point, you can save a draft flow as a flow version. A flow version is an immutable snapshot of your flow data and associated agent data like intents, entities, webhooks, pages, route groups, etc.
 *
 * To get more information about Version, see:
 *
 * * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows.versions)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
 *
 * ## Example Usage
 * ### Dialogflowcx Version Full
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
 * const version1 = new gcp.diagflow.CxVersion("version1", {
 *     parent: agent.startFlow,
 *     displayName: "1.0.0",
 *     description: "version 1.0.0",
 * });
 * ```
 *
 * ## Import
 *
 * Version can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:diagflow/cxVersion:CxVersion default {{parent}}/versions/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:diagflow/cxVersion:CxVersion default {{parent}}/{{name}}
 * ```
 */
export class CxVersion extends pulumi.CustomResource {
    /**
     * Get an existing CxVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CxVersionState, opts?: pulumi.CustomResourceOptions): CxVersion {
        return new CxVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:diagflow/cxVersion:CxVersion';

    /**
     * Returns true if the given object is an instance of CxVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CxVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CxVersion.__pulumiType;
    }

    /**
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples:
     * "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The human-readable name of the version. Limit of 64 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow
     * upon version creation.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The NLU settings of the flow at version creation.
     */
    public /*out*/ readonly nluSettings!: pulumi.Output<outputs.diagflow.CxVersionNluSetting[]>;
    /**
     * The Flow to create an Version for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
     */
    public readonly parent!: pulumi.Output<string | undefined>;
    /**
     * The state of this version. * RUNNING: Version is not ready to serve (e.g. training is running). * SUCCEEDED: Training
     * has succeeded and this version is ready to serve. * FAILED: Version training failed.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a CxVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CxVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CxVersionArgs | CxVersionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CxVersionState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nluSettings"] = state ? state.nluSettings : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as CxVersionArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nluSettings"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CxVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CxVersion resources.
 */
export interface CxVersionState {
    /**
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples:
     * "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name of the version. Limit of 64 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow
     * upon version creation.
     */
    name?: pulumi.Input<string>;
    /**
     * The NLU settings of the flow at version creation.
     */
    nluSettings?: pulumi.Input<pulumi.Input<inputs.diagflow.CxVersionNluSetting>[]>;
    /**
     * The Flow to create an Version for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
     */
    parent?: pulumi.Input<string>;
    /**
     * The state of this version. * RUNNING: Version is not ready to serve (e.g. training is running). * SUCCEEDED: Training
     * has succeeded and this version is ready to serve. * FAILED: Version training failed.
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CxVersion resource.
 */
export interface CxVersionArgs {
    /**
     * The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name of the version. Limit of 64 characters.
     */
    displayName: pulumi.Input<string>;
    /**
     * The Flow to create an Version for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
     */
    parent?: pulumi.Input<string>;
}

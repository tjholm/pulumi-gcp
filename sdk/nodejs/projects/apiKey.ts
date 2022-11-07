// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The Apikeys Key resource
 *
 * ## Example Usage
 * ### Android_key
 * A basic example of a android api keys key
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.organizations.Project("basic", {
 *     projectId: "app",
 *     orgId: "123456789",
 * });
 * const primary = new gcp.projects.ApiKey("primary", {
 *     displayName: "sample-key",
 *     project: basic.name,
 *     restrictions: {
 *         androidKeyRestrictions: {
 *             allowedApplications: [{
 *                 packageName: "com.example.app123",
 *                 sha1Fingerprint: "1699466a142d4682a5f91b50fdf400f2358e2b0b",
 *             }],
 *         },
 *         apiTargets: [{
 *             service: "translate.googleapis.com",
 *             methods: ["GET*"],
 *         }],
 *     },
 * });
 * ```
 * ### Basic_key
 * A basic example of a api keys key
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.organizations.Project("basic", {
 *     projectId: "app",
 *     orgId: "123456789",
 * });
 * const primary = new gcp.projects.ApiKey("primary", {
 *     displayName: "sample-key",
 *     project: basic.name,
 *     restrictions: {
 *         apiTargets: [{
 *             service: "translate.googleapis.com",
 *             methods: ["GET*"],
 *         }],
 *         browserKeyRestrictions: {
 *             allowedReferrers: [".*"],
 *         },
 *     },
 * });
 * ```
 * ### Ios_key
 * A basic example of a ios api keys key
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.organizations.Project("basic", {
 *     projectId: "app",
 *     orgId: "123456789",
 * });
 * const primary = new gcp.projects.ApiKey("primary", {
 *     displayName: "sample-key",
 *     project: basic.name,
 *     restrictions: {
 *         apiTargets: [{
 *             service: "translate.googleapis.com",
 *             methods: ["GET*"],
 *         }],
 *         iosKeyRestrictions: {
 *             allowedBundleIds: ["com.google.app.macos"],
 *         },
 *     },
 * });
 * ```
 * ### Minimal_key
 * A minimal example of a api keys key
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.organizations.Project("basic", {
 *     projectId: "app",
 *     orgId: "123456789",
 * });
 * const primary = new gcp.projects.ApiKey("primary", {
 *     displayName: "sample-key",
 *     project: basic.name,
 * });
 * ```
 * ### Server_key
 * A basic example of a server api keys key
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.organizations.Project("basic", {
 *     projectId: "app",
 *     orgId: "123456789",
 * });
 * const primary = new gcp.projects.ApiKey("primary", {
 *     displayName: "sample-key",
 *     project: basic.name,
 *     restrictions: {
 *         apiTargets: [{
 *             service: "translate.googleapis.com",
 *             methods: ["GET*"],
 *         }],
 *         serverKeyRestrictions: {
 *             allowedIps: ["127.0.0.1"],
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Key can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:projects/apiKey:ApiKey default projects/{{project}}/locations/global/keys/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:projects/apiKey:ApiKey default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:projects/apiKey:ApiKey default {{name}}
 * ```
 */
export class ApiKey extends pulumi.CustomResource {
    /**
     * Get an existing ApiKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiKeyState, opts?: pulumi.CustomResourceOptions): ApiKey {
        return new ApiKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:projects/apiKey:ApiKey';

    /**
     * Returns true if the given object is an instance of ApiKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiKey.__pulumiType;
    }

    /**
     * Human-readable display name of this API key. Modifiable by user.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Output only. An encrypted and signed value held by this key. This field can be accessed only through the `GetKeyString`
     * method.
     */
    public /*out*/ readonly keyString!: pulumi.Output<string>;
    /**
     * The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular expression: `a-z?`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Key restrictions.
     */
    public readonly restrictions!: pulumi.Output<outputs.projects.ApiKeyRestrictions | undefined>;
    /**
     * Output only. Unique id in UUID4 format.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;

    /**
     * Create a ApiKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ApiKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiKeyArgs | ApiKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiKeyState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["keyString"] = state ? state.keyString : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["restrictions"] = state ? state.restrictions : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
        } else {
            const args = argsOrState as ApiKeyArgs | undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["restrictions"] = args ? args.restrictions : undefined;
            resourceInputs["keyString"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["keyString"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ApiKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiKey resources.
 */
export interface ApiKeyState {
    /**
     * Human-readable display name of this API key. Modifiable by user.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Output only. An encrypted and signed value held by this key. This field can be accessed only through the `GetKeyString`
     * method.
     */
    keyString?: pulumi.Input<string>;
    /**
     * The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular expression: `a-z?`.
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Key restrictions.
     */
    restrictions?: pulumi.Input<inputs.projects.ApiKeyRestrictions>;
    /**
     * Output only. Unique id in UUID4 format.
     */
    uid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiKey resource.
 */
export interface ApiKeyArgs {
    /**
     * Human-readable display name of this API key. Modifiable by user.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular expression: `a-z?`.
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Key restrictions.
     */
    restrictions?: pulumi.Input<inputs.projects.ApiKeyRestrictions>;
}

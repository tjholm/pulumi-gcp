// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * OIDC IdP configuration for a Identity Toolkit project.
 *
 * You must enable the
 * [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
 * the marketplace prior to using this resource.
 *
 * ## Example Usage
 * ### Identity Platform Oauth Idp Config Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const oauthIdpConfig = new gcp.identityplatform.OauthIdpConfig("oauthIdpConfig", {
 *     clientId: "client-id",
 *     clientSecret: "secret",
 *     displayName: "Display Name",
 *     enabled: true,
 *     issuer: "issuer",
 * });
 * ```
 *
 * ## Import
 *
 * OauthIdpConfig can be imported using any of these accepted formats* `projects/{{project}}/oauthIdpConfigs/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, OauthIdpConfig can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/oauthIdpConfig:OauthIdpConfig default projects/{{project}}/oauthIdpConfigs/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/oauthIdpConfig:OauthIdpConfig default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/oauthIdpConfig:OauthIdpConfig default {{name}}
 * ```
 */
export class OauthIdpConfig extends pulumi.CustomResource {
    /**
     * Get an existing OauthIdpConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OauthIdpConfigState, opts?: pulumi.CustomResourceOptions): OauthIdpConfig {
        return new OauthIdpConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:identityplatform/oauthIdpConfig:OauthIdpConfig';

    /**
     * Returns true if the given object is an instance of OauthIdpConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OauthIdpConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OauthIdpConfig.__pulumiType;
    }

    /**
     * The client id of an OAuth client.
     *
     *
     * - - -
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The client secret of the OAuth client, to enable OIDC code flow.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * Human friendly display name.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * If this config allows users to sign in with the provider.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * For OIDC Idps, the issuer identifier.
     */
    public readonly issuer!: pulumi.Output<string>;
    /**
     * The name of the OauthIdpConfig. Must start with `oidc.`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a OauthIdpConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OauthIdpConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OauthIdpConfigArgs | OauthIdpConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OauthIdpConfigState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as OauthIdpConfigArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.issuer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issuer'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args ? args.clientSecret : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["issuer"] = args ? args.issuer : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OauthIdpConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OauthIdpConfig resources.
 */
export interface OauthIdpConfigState {
    /**
     * The client id of an OAuth client.
     *
     *
     * - - -
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client secret of the OAuth client, to enable OIDC code flow.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Human friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * For OIDC Idps, the issuer identifier.
     */
    issuer?: pulumi.Input<string>;
    /**
     * The name of the OauthIdpConfig. Must start with `oidc.`.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OauthIdpConfig resource.
 */
export interface OauthIdpConfigArgs {
    /**
     * The client id of an OAuth client.
     *
     *
     * - - -
     */
    clientId: pulumi.Input<string>;
    /**
     * The client secret of the OAuth client, to enable OIDC code flow.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Human friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * For OIDC Idps, the issuer identifier.
     */
    issuer: pulumi.Input<string>;
    /**
     * The name of the OauthIdpConfig. Must start with `oidc.`.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

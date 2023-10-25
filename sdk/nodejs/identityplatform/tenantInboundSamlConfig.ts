// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Inbound SAML configuration for a Identity Toolkit tenant.
 *
 * You must enable the
 * [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
 * the marketplace prior to using this resource.
 *
 * ## Example Usage
 * ### Identity Platform Tenant Inbound Saml Config Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 *
 * const tenant = new gcp.identityplatform.Tenant("tenant", {displayName: "tenant"});
 * const tenantSamlConfig = new gcp.identityplatform.TenantInboundSamlConfig("tenantSamlConfig", {
 *     displayName: "Display Name",
 *     tenant: tenant.name,
 *     idpConfig: {
 *         idpEntityId: "tf-idp",
 *         signRequest: true,
 *         ssoUrl: "https://example.com",
 *         idpCertificates: [{
 *             x509Certificate: fs.readFileSync("test-fixtures/rsa_cert.pem"),
 *         }],
 *     },
 *     spConfig: {
 *         spEntityId: "tf-sp",
 *         callbackUri: "https://example.com",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * TenantInboundSamlConfig can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig default projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig default {{project}}/{{tenant}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig default {{tenant}}/{{name}}
 * ```
 */
export class TenantInboundSamlConfig extends pulumi.CustomResource {
    /**
     * Get an existing TenantInboundSamlConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TenantInboundSamlConfigState, opts?: pulumi.CustomResourceOptions): TenantInboundSamlConfig {
        return new TenantInboundSamlConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig';

    /**
     * Returns true if the given object is an instance of TenantInboundSamlConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TenantInboundSamlConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TenantInboundSamlConfig.__pulumiType;
    }

    /**
     * Human friendly display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * SAML IdP configuration when the project acts as the relying party
     * Structure is documented below.
     */
    public readonly idpConfig!: pulumi.Output<outputs.identityplatform.TenantInboundSamlConfigIdpConfig>;
    /**
     * The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
     * hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
     * alphanumeric character, and have at least 2 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * SAML SP (Service Provider) configuration when the project acts as the relying party to receive
     * and accept an authentication assertion issued by a SAML identity provider.
     * Structure is documented below.
     */
    public readonly spConfig!: pulumi.Output<outputs.identityplatform.TenantInboundSamlConfigSpConfig>;
    /**
     * The name of the tenant where this inbound SAML config resource exists
     */
    public readonly tenant!: pulumi.Output<string>;

    /**
     * Create a TenantInboundSamlConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TenantInboundSamlConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TenantInboundSamlConfigArgs | TenantInboundSamlConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TenantInboundSamlConfigState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["idpConfig"] = state ? state.idpConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["spConfig"] = state ? state.spConfig : undefined;
            resourceInputs["tenant"] = state ? state.tenant : undefined;
        } else {
            const args = argsOrState as TenantInboundSamlConfigArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.idpConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'idpConfig'");
            }
            if ((!args || args.spConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spConfig'");
            }
            if ((!args || args.tenant === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenant'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["idpConfig"] = args ? args.idpConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["spConfig"] = args ? args.spConfig : undefined;
            resourceInputs["tenant"] = args ? args.tenant : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TenantInboundSamlConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TenantInboundSamlConfig resources.
 */
export interface TenantInboundSamlConfigState {
    /**
     * Human friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * SAML IdP configuration when the project acts as the relying party
     * Structure is documented below.
     */
    idpConfig?: pulumi.Input<inputs.identityplatform.TenantInboundSamlConfigIdpConfig>;
    /**
     * The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
     * hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
     * alphanumeric character, and have at least 2 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * SAML SP (Service Provider) configuration when the project acts as the relying party to receive
     * and accept an authentication assertion issued by a SAML identity provider.
     * Structure is documented below.
     */
    spConfig?: pulumi.Input<inputs.identityplatform.TenantInboundSamlConfigSpConfig>;
    /**
     * The name of the tenant where this inbound SAML config resource exists
     */
    tenant?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TenantInboundSamlConfig resource.
 */
export interface TenantInboundSamlConfigArgs {
    /**
     * Human friendly display name.
     */
    displayName: pulumi.Input<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * SAML IdP configuration when the project acts as the relying party
     * Structure is documented below.
     */
    idpConfig: pulumi.Input<inputs.identityplatform.TenantInboundSamlConfigIdpConfig>;
    /**
     * The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
     * hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
     * alphanumeric character, and have at least 2 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * SAML SP (Service Provider) configuration when the project acts as the relying party to receive
     * and accept an authentication assertion issued by a SAML identity provider.
     * Structure is documented below.
     */
    spConfig: pulumi.Input<inputs.identityplatform.TenantInboundSamlConfigSpConfig>;
    /**
     * The name of the tenant where this inbound SAML config resource exists
     */
    tenant: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * An alias from a key/certificate pair.
 *
 * To get more information about KeysotresAliasesKeyCertFile, see:
 *
 * * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.keystores.aliases)
 * * How-to Guides
 *     * [Keystores Aliases](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.keystores.aliases)
 *
 * ## Import
 *
 * KeysotresAliasesKeyCertFile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:apigee/keystoresAliasesKeyCertFile:KeystoresAliasesKeyCertFile default organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:apigee/keystoresAliasesKeyCertFile:KeystoresAliasesKeyCertFile default {{org_id}}/{{environment}}/{{keystore}}/{{alias}}
 * ```
 */
export class KeystoresAliasesKeyCertFile extends pulumi.CustomResource {
    /**
     * Get an existing KeystoresAliasesKeyCertFile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeystoresAliasesKeyCertFileState, opts?: pulumi.CustomResourceOptions): KeystoresAliasesKeyCertFile {
        return new KeystoresAliasesKeyCertFile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:apigee/keystoresAliasesKeyCertFile:KeystoresAliasesKeyCertFile';

    /**
     * Returns true if the given object is an instance of KeystoresAliasesKeyCertFile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeystoresAliasesKeyCertFile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeystoresAliasesKeyCertFile.__pulumiType;
    }

    /**
     * Alias Name
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * Cert content
     */
    public readonly cert!: pulumi.Output<string>;
    /**
     * Chain of certificates under this alias.
     * Structure is documented below.
     */
    public readonly certsInfo!: pulumi.Output<outputs.apigee.KeystoresAliasesKeyCertFileCertsInfo>;
    /**
     * Environment associated with the alias
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * Private Key content, omit if uploading to truststore
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * Keystore Name
     */
    public readonly keystore!: pulumi.Output<string>;
    /**
     * Organization ID associated with the alias, without organization/ prefix
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * Password for the Private Key if it's encrypted
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Optional.Type of Alias
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a KeystoresAliasesKeyCertFile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeystoresAliasesKeyCertFileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeystoresAliasesKeyCertFileArgs | KeystoresAliasesKeyCertFileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeystoresAliasesKeyCertFileState | undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["cert"] = state ? state.cert : undefined;
            resourceInputs["certsInfo"] = state ? state.certsInfo : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["keystore"] = state ? state.keystore : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as KeystoresAliasesKeyCertFileArgs | undefined;
            if ((!args || args.alias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alias'");
            }
            if ((!args || args.cert === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cert'");
            }
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.keystore === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keystore'");
            }
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["cert"] = args ? args.cert : undefined;
            resourceInputs["certsInfo"] = args ? args.certsInfo : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["key"] = args?.key ? pulumi.secret(args.key) : undefined;
            resourceInputs["keystore"] = args ? args.keystore : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["key", "password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(KeystoresAliasesKeyCertFile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeystoresAliasesKeyCertFile resources.
 */
export interface KeystoresAliasesKeyCertFileState {
    /**
     * Alias Name
     */
    alias?: pulumi.Input<string>;
    /**
     * Cert content
     */
    cert?: pulumi.Input<string>;
    /**
     * Chain of certificates under this alias.
     * Structure is documented below.
     */
    certsInfo?: pulumi.Input<inputs.apigee.KeystoresAliasesKeyCertFileCertsInfo>;
    /**
     * Environment associated with the alias
     */
    environment?: pulumi.Input<string>;
    /**
     * Private Key content, omit if uploading to truststore
     */
    key?: pulumi.Input<string>;
    /**
     * Keystore Name
     */
    keystore?: pulumi.Input<string>;
    /**
     * Organization ID associated with the alias, without organization/ prefix
     */
    orgId?: pulumi.Input<string>;
    /**
     * Password for the Private Key if it's encrypted
     */
    password?: pulumi.Input<string>;
    /**
     * Optional.Type of Alias
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeystoresAliasesKeyCertFile resource.
 */
export interface KeystoresAliasesKeyCertFileArgs {
    /**
     * Alias Name
     */
    alias: pulumi.Input<string>;
    /**
     * Cert content
     */
    cert: pulumi.Input<string>;
    /**
     * Chain of certificates under this alias.
     * Structure is documented below.
     */
    certsInfo?: pulumi.Input<inputs.apigee.KeystoresAliasesKeyCertFileCertsInfo>;
    /**
     * Environment associated with the alias
     */
    environment: pulumi.Input<string>;
    /**
     * Private Key content, omit if uploading to truststore
     */
    key?: pulumi.Input<string>;
    /**
     * Keystore Name
     */
    keystore: pulumi.Input<string>;
    /**
     * Organization ID associated with the alias, without organization/ prefix
     */
    orgId: pulumi.Input<string>;
    /**
     * Password for the Private Key if it's encrypted
     */
    password?: pulumi.Input<string>;
}

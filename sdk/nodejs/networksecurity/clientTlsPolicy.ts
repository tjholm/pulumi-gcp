// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Network Security Client Tls Policy Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.networksecurity.ClientTlsPolicy("default", {
 *     labels: {
 *         foo: "bar",
 *     },
 *     description: "my description",
 *     sni: "secure.example.com",
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Network Security Client Tls Policy Advanced
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.networksecurity.ClientTlsPolicy("default", {
 *     labels: {
 *         foo: "bar",
 *     },
 *     description: "my description",
 *     clientCertificate: {
 *         certificateProviderInstance: {
 *             pluginInstance: "google_cloud_private_spiffe",
 *         },
 *     },
 *     serverValidationCas: [
 *         {
 *             grpcEndpoint: {
 *                 targetUri: "unix:mypath",
 *             },
 *         },
 *         {
 *             grpcEndpoint: {
 *                 targetUri: "unix:mypath1",
 *             },
 *         },
 *     ],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * ClientTlsPolicy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:networksecurity/clientTlsPolicy:ClientTlsPolicy default projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networksecurity/clientTlsPolicy:ClientTlsPolicy default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networksecurity/clientTlsPolicy:ClientTlsPolicy default {{location}}/{{name}}
 * ```
 */
export class ClientTlsPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ClientTlsPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientTlsPolicyState, opts?: pulumi.CustomResourceOptions): ClientTlsPolicy {
        return new ClientTlsPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:networksecurity/clientTlsPolicy:ClientTlsPolicy';

    /**
     * Returns true if the given object is an instance of ClientTlsPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientTlsPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientTlsPolicy.__pulumiType;
    }

    /**
     * Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
     * Structure is documented below.
     */
    public readonly clientCertificate!: pulumi.Output<outputs.networksecurity.ClientTlsPolicyClientCertificate | undefined>;
    /**
     * Time the ClientTlsPolicy was created in UTC.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A free-text description of the resource. Max length 1024 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Set of label tags associated with the ClientTlsPolicy resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location of the client tls policy.
     * The default value is `global`.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Name of the ClientTlsPolicy resource.
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
     * Structure is documented below.
     */
    public readonly serverValidationCas!: pulumi.Output<outputs.networksecurity.ClientTlsPolicyServerValidationCa[] | undefined>;
    /**
     * Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
     */
    public readonly sni!: pulumi.Output<string | undefined>;
    /**
     * Time the ClientTlsPolicy was updated in UTC.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ClientTlsPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClientTlsPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientTlsPolicyArgs | ClientTlsPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientTlsPolicyState | undefined;
            resourceInputs["clientCertificate"] = state ? state.clientCertificate : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["serverValidationCas"] = state ? state.serverValidationCas : undefined;
            resourceInputs["sni"] = state ? state.sni : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as ClientTlsPolicyArgs | undefined;
            resourceInputs["clientCertificate"] = args ? args.clientCertificate : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serverValidationCas"] = args ? args.serverValidationCas : undefined;
            resourceInputs["sni"] = args ? args.sni : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ClientTlsPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientTlsPolicy resources.
 */
export interface ClientTlsPolicyState {
    /**
     * Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
     * Structure is documented below.
     */
    clientCertificate?: pulumi.Input<inputs.networksecurity.ClientTlsPolicyClientCertificate>;
    /**
     * Time the ClientTlsPolicy was created in UTC.
     */
    createTime?: pulumi.Input<string>;
    /**
     * A free-text description of the resource. Max length 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Set of label tags associated with the ClientTlsPolicy resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location of the client tls policy.
     * The default value is `global`.
     */
    location?: pulumi.Input<string>;
    /**
     * Name of the ClientTlsPolicy resource.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
     * Structure is documented below.
     */
    serverValidationCas?: pulumi.Input<pulumi.Input<inputs.networksecurity.ClientTlsPolicyServerValidationCa>[]>;
    /**
     * Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
     */
    sni?: pulumi.Input<string>;
    /**
     * Time the ClientTlsPolicy was updated in UTC.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientTlsPolicy resource.
 */
export interface ClientTlsPolicyArgs {
    /**
     * Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
     * Structure is documented below.
     */
    clientCertificate?: pulumi.Input<inputs.networksecurity.ClientTlsPolicyClientCertificate>;
    /**
     * A free-text description of the resource. Max length 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Set of label tags associated with the ClientTlsPolicy resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location of the client tls policy.
     * The default value is `global`.
     */
    location?: pulumi.Input<string>;
    /**
     * Name of the ClientTlsPolicy resource.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
     * Structure is documented below.
     */
    serverValidationCas?: pulumi.Input<pulumi.Input<inputs.networksecurity.ClientTlsPolicyServerValidationCa>[]>;
    /**
     * Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
     */
    sni?: pulumi.Input<string>;
}

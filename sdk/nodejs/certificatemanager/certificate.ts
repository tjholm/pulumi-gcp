// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Certificate Manager Google Managed Certificate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.certificatemanager.DnsAuthorization("instance", {
 *     description: "The default dnss",
 *     domain: "subdomain.hashicorptest.com",
 * });
 * const instance2 = new gcp.certificatemanager.DnsAuthorization("instance2", {
 *     description: "The default dnss",
 *     domain: "subdomain2.hashicorptest.com",
 * });
 * const _default = new gcp.certificatemanager.Certificate("default", {
 *     description: "The default cert",
 *     scope: "EDGE_CACHE",
 *     managed: {
 *         domains: [
 *             instance.domain,
 *             instance2.domain,
 *         ],
 *         dnsAuthorizations: [
 *             instance.id,
 *             instance2.id,
 *         ],
 *     },
 * });
 * ```
 * ### Certificate Manager Certificate Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.certificatemanager.DnsAuthorization("instance", {
 *     description: "The default dnss",
 *     domain: "subdomain.hashicorptest.com",
 * });
 * const instance2 = new gcp.certificatemanager.DnsAuthorization("instance2", {
 *     description: "The default dnss",
 *     domain: "subdomain2.hashicorptest.com",
 * });
 * const _default = new gcp.certificatemanager.Certificate("default", {
 *     description: "The default cert",
 *     scope: "EDGE_CACHE",
 *     managed: {
 *         domains: [
 *             instance.domain,
 *             instance2.domain,
 *         ],
 *         dnsAuthorizations: [
 *             instance.id,
 *             instance2.id,
 *         ],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Certificate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:certificatemanager/certificate:Certificate default projects/{{project}}/locations/global/certificates/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:certificatemanager/certificate:Certificate default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:certificatemanager/certificate:Certificate default {{name}}
 * ```
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateState, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:certificatemanager/certificate:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * A human-readable description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Set of label tags associated with the Certificate resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Configuration and state of a Managed Certificate.
     * Certificate Manager provisions and renews Managed Certificates
     * automatically, for as long as it's authorized to do so.
     * Structure is documented below.
     */
    public readonly managed!: pulumi.Output<outputs.certificatemanager.CertificateManaged | undefined>;
    /**
     * A user-defined name of the certificate. Certificate names must be unique
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The scope of the certificate.
     * DEFAULT: Certificates with default scope are served from core Google data centers.
     * If unsure, choose this option.
     * EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
     * served from non-core Google data centers.
     * Currently allowed only for managed certificates.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * Certificate data for a SelfManaged Certificate.
     * SelfManaged Certificates are uploaded by the user. Updating such
     * certificates before they expire remains the user's responsibility.
     * Structure is documented below.
     */
    public readonly selfManaged!: pulumi.Output<outputs.certificatemanager.CertificateSelfManaged | undefined>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs | CertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["managed"] = state ? state.managed : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["selfManaged"] = state ? state.selfManaged : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["managed"] = args ? args.managed : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["selfManaged"] = args ? args.selfManaged : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Certificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certificate resources.
 */
export interface CertificateState {
    /**
     * A human-readable description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Set of label tags associated with the Certificate resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration and state of a Managed Certificate.
     * Certificate Manager provisions and renews Managed Certificates
     * automatically, for as long as it's authorized to do so.
     * Structure is documented below.
     */
    managed?: pulumi.Input<inputs.certificatemanager.CertificateManaged>;
    /**
     * A user-defined name of the certificate. Certificate names must be unique
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The scope of the certificate.
     * DEFAULT: Certificates with default scope are served from core Google data centers.
     * If unsure, choose this option.
     * EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
     * served from non-core Google data centers.
     * Currently allowed only for managed certificates.
     */
    scope?: pulumi.Input<string>;
    /**
     * Certificate data for a SelfManaged Certificate.
     * SelfManaged Certificates are uploaded by the user. Updating such
     * certificates before they expire remains the user's responsibility.
     * Structure is documented below.
     */
    selfManaged?: pulumi.Input<inputs.certificatemanager.CertificateSelfManaged>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * A human-readable description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Set of label tags associated with the Certificate resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration and state of a Managed Certificate.
     * Certificate Manager provisions and renews Managed Certificates
     * automatically, for as long as it's authorized to do so.
     * Structure is documented below.
     */
    managed?: pulumi.Input<inputs.certificatemanager.CertificateManaged>;
    /**
     * A user-defined name of the certificate. Certificate names must be unique
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The scope of the certificate.
     * DEFAULT: Certificates with default scope are served from core Google data centers.
     * If unsure, choose this option.
     * EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
     * served from non-core Google data centers.
     * Currently allowed only for managed certificates.
     */
    scope?: pulumi.Input<string>;
    /**
     * Certificate data for a SelfManaged Certificate.
     * SelfManaged Certificates are uploaded by the user. Updating such
     * certificates before they expire remains the user's responsibility.
     * Structure is documented below.
     */
    selfManaged?: pulumi.Input<inputs.certificatemanager.CertificateSelfManaged>;
}

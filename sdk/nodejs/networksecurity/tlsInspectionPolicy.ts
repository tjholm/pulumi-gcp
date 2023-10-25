// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Network Security Tls Inspection Policy Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultCaPool = new gcp.certificateauthority.CaPool("defaultCaPool", {
 *     location: "us-central1",
 *     tier: "DEVOPS",
 *     publishingOptions: {
 *         publishCaCert: false,
 *         publishCrl: false,
 *     },
 *     issuancePolicy: {
 *         maximumLifetime: "1209600s",
 *         baselineValues: {
 *             caOptions: {
 *                 isCa: false,
 *             },
 *             keyUsage: {
 *                 baseKeyUsage: {},
 *                 extendedKeyUsage: {
 *                     serverAuth: true,
 *                 },
 *             },
 *         },
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const defaultAuthority = new gcp.certificateauthority.Authority("defaultAuthority", {
 *     pool: defaultCaPool.name,
 *     certificateAuthorityId: "my-basic-certificate-authority",
 *     location: "us-central1",
 *     lifetime: "86400s",
 *     type: "SELF_SIGNED",
 *     deletionProtection: false,
 *     skipGracePeriod: true,
 *     ignoreActiveCertificatesOnDeletion: true,
 *     config: {
 *         subjectConfig: {
 *             subject: {
 *                 organization: "Test LLC",
 *                 commonName: "my-ca",
 *             },
 *         },
 *         x509Config: {
 *             caOptions: {
 *                 isCa: true,
 *             },
 *             keyUsage: {
 *                 baseKeyUsage: {
 *                     certSign: true,
 *                     crlSign: true,
 *                 },
 *                 extendedKeyUsage: {
 *                     serverAuth: false,
 *                 },
 *             },
 *         },
 *     },
 *     keySpec: {
 *         algorithm: "RSA_PKCS1_4096_SHA256",
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const nsSa = new gcp.projects.ServiceIdentity("nsSa", {service: "networksecurity.googleapis.com"}, {
 *     provider: google_beta,
 * });
 * const tlsInspectionPermission = new gcp.certificateauthority.CaPoolIamMember("tlsInspectionPermission", {
 *     caPool: defaultCaPool.id,
 *     role: "roles/privateca.certificateManager",
 *     member: pulumi.interpolate`serviceAccount:${nsSa.email}`,
 * }, {
 *     provider: google_beta,
 * });
 * const defaultTlsInspectionPolicy = new gcp.networksecurity.TlsInspectionPolicy("defaultTlsInspectionPolicy", {
 *     location: "us-central1",
 *     caPool: defaultCaPool.id,
 *     excludePublicCaSet: false,
 * }, {
 *     provider: google_beta,
 *     dependsOn: [
 *         defaultCaPool,
 *         defaultAuthority,
 *         tlsInspectionPermission,
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * TlsInspectionPolicy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:networksecurity/tlsInspectionPolicy:TlsInspectionPolicy default projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networksecurity/tlsInspectionPolicy:TlsInspectionPolicy default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networksecurity/tlsInspectionPolicy:TlsInspectionPolicy default {{location}}/{{name}}
 * ```
 */
export class TlsInspectionPolicy extends pulumi.CustomResource {
    /**
     * Get an existing TlsInspectionPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TlsInspectionPolicyState, opts?: pulumi.CustomResourceOptions): TlsInspectionPolicy {
        return new TlsInspectionPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:networksecurity/tlsInspectionPolicy:TlsInspectionPolicy';

    /**
     * Returns true if the given object is an instance of TlsInspectionPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TlsInspectionPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TlsInspectionPolicy.__pulumiType;
    }

    /**
     * A CA pool resource used to issue interception certificates.
     */
    public readonly caPool!: pulumi.Output<string>;
    /**
     * The timestamp when the resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Free-text description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
     */
    public readonly excludePublicCaSet!: pulumi.Output<boolean | undefined>;
    /**
     * The location of the tls inspection policy.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Short name of the TlsInspectionPolicy resource to be created.
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
     * The timestamp when the resource was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a TlsInspectionPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TlsInspectionPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TlsInspectionPolicyArgs | TlsInspectionPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TlsInspectionPolicyState | undefined;
            resourceInputs["caPool"] = state ? state.caPool : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["excludePublicCaSet"] = state ? state.excludePublicCaSet : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as TlsInspectionPolicyArgs | undefined;
            if ((!args || args.caPool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'caPool'");
            }
            resourceInputs["caPool"] = args ? args.caPool : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["excludePublicCaSet"] = args ? args.excludePublicCaSet : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TlsInspectionPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TlsInspectionPolicy resources.
 */
export interface TlsInspectionPolicyState {
    /**
     * A CA pool resource used to issue interception certificates.
     */
    caPool?: pulumi.Input<string>;
    /**
     * The timestamp when the resource was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Free-text description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
     */
    excludePublicCaSet?: pulumi.Input<boolean>;
    /**
     * The location of the tls inspection policy.
     */
    location?: pulumi.Input<string>;
    /**
     * Short name of the TlsInspectionPolicy resource to be created.
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
     * The timestamp when the resource was updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TlsInspectionPolicy resource.
 */
export interface TlsInspectionPolicyArgs {
    /**
     * A CA pool resource used to issue interception certificates.
     */
    caPool: pulumi.Input<string>;
    /**
     * Free-text description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
     */
    excludePublicCaSet?: pulumi.Input<boolean>;
    /**
     * The location of the tls inspection policy.
     */
    location?: pulumi.Input<string>;
    /**
     * Short name of the TlsInspectionPolicy resource to be created.
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
}

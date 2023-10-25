// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Certificate corresponds to a signed X.509 certificate issued by a Certificate.
 *
 * > **Note:** The Certificate Authority that is referenced by this resource **must** be
 * `tier = "ENTERPRISE"`
 *
 * ## Example Usage
 * ### Privateca Certificate With Template
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultCaPool = new gcp.certificateauthority.CaPool("defaultCaPool", {
 *     location: "us-central1",
 *     tier: "ENTERPRISE",
 * });
 * const defaultCertificateTemplate = new gcp.certificateauthority.CertificateTemplate("defaultCertificateTemplate", {
 *     location: "us-central1",
 *     description: "An updated sample certificate template",
 *     identityConstraints: {
 *         allowSubjectAltNamesPassthrough: true,
 *         allowSubjectPassthrough: true,
 *         celExpression: {
 *             description: "Always true",
 *             expression: "true",
 *             location: "any.file.anywhere",
 *             title: "Sample expression",
 *         },
 *     },
 *     passthroughExtensions: {
 *         additionalExtensions: [{
 *             objectIdPaths: [
 *                 1,
 *                 6,
 *             ],
 *         }],
 *         knownExtensions: ["EXTENDED_KEY_USAGE"],
 *     },
 *     predefinedValues: {
 *         additionalExtensions: [{
 *             objectId: {
 *                 objectIdPaths: [
 *                     1,
 *                     6,
 *                 ],
 *             },
 *             value: "c3RyaW5nCg==",
 *             critical: true,
 *         }],
 *         aiaOcspServers: ["string"],
 *         caOptions: {
 *             isCa: false,
 *             maxIssuerPathLength: 6,
 *         },
 *         keyUsage: {
 *             baseKeyUsage: {
 *                 certSign: false,
 *                 contentCommitment: true,
 *                 crlSign: false,
 *                 dataEncipherment: true,
 *                 decipherOnly: true,
 *                 digitalSignature: true,
 *                 encipherOnly: true,
 *                 keyAgreement: true,
 *                 keyEncipherment: true,
 *             },
 *             extendedKeyUsage: {
 *                 clientAuth: true,
 *                 codeSigning: true,
 *                 emailProtection: true,
 *                 ocspSigning: true,
 *                 serverAuth: true,
 *                 timeStamping: true,
 *             },
 *             unknownExtendedKeyUsages: [{
 *                 objectIdPaths: [
 *                     1,
 *                     6,
 *                 ],
 *             }],
 *         },
 *         policyIds: [{
 *             objectIdPaths: [
 *                 1,
 *                 6,
 *             ],
 *         }],
 *     },
 * });
 * const defaultAuthority = new gcp.certificateauthority.Authority("defaultAuthority", {
 *     location: "us-central1",
 *     pool: defaultCaPool.name,
 *     certificateAuthorityId: "my-authority",
 *     config: {
 *         subjectConfig: {
 *             subject: {
 *                 organization: "HashiCorp",
 *                 commonName: "my-certificate-authority",
 *             },
 *             subjectAltName: {
 *                 dnsNames: ["hashicorp.com"],
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
 *     deletionProtection: false,
 *     skipGracePeriod: true,
 *     ignoreActiveCertificatesOnDeletion: true,
 * });
 * const defaultCertificate = new gcp.certificateauthority.Certificate("defaultCertificate", {
 *     location: "us-central1",
 *     pool: defaultCaPool.name,
 *     certificateAuthority: defaultAuthority.certificateAuthorityId,
 *     lifetime: "860s",
 *     pemCsr: fs.readFileSync("test-fixtures/rsa_csr.pem"),
 *     certificateTemplate: defaultCertificateTemplate.id,
 * });
 * ```
 * ### Privateca Certificate Csr
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultCaPool = new gcp.certificateauthority.CaPool("defaultCaPool", {
 *     location: "us-central1",
 *     tier: "ENTERPRISE",
 * });
 * const defaultAuthority = new gcp.certificateauthority.Authority("defaultAuthority", {
 *     location: "us-central1",
 *     pool: defaultCaPool.name,
 *     certificateAuthorityId: "my-authority",
 *     config: {
 *         subjectConfig: {
 *             subject: {
 *                 organization: "HashiCorp",
 *                 commonName: "my-certificate-authority",
 *             },
 *             subjectAltName: {
 *                 dnsNames: ["hashicorp.com"],
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
 *     deletionProtection: false,
 *     skipGracePeriod: true,
 *     ignoreActiveCertificatesOnDeletion: true,
 * });
 * const defaultCertificate = new gcp.certificateauthority.Certificate("defaultCertificate", {
 *     location: "us-central1",
 *     pool: defaultCaPool.name,
 *     certificateAuthority: defaultAuthority.certificateAuthorityId,
 *     lifetime: "860s",
 *     pemCsr: fs.readFileSync("test-fixtures/rsa_csr.pem"),
 * });
 * ```
 * ### Privateca Certificate No Authority
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultCaPool = new gcp.certificateauthority.CaPool("defaultCaPool", {
 *     location: "us-central1",
 *     tier: "ENTERPRISE",
 * });
 * const defaultAuthority = new gcp.certificateauthority.Authority("defaultAuthority", {
 *     location: "us-central1",
 *     pool: defaultCaPool.name,
 *     certificateAuthorityId: "my-authority",
 *     config: {
 *         subjectConfig: {
 *             subject: {
 *                 organization: "HashiCorp",
 *                 commonName: "my-certificate-authority",
 *             },
 *             subjectAltName: {
 *                 dnsNames: ["hashicorp.com"],
 *             },
 *         },
 *         x509Config: {
 *             caOptions: {
 *                 isCa: true,
 *             },
 *             keyUsage: {
 *                 baseKeyUsage: {
 *                     digitalSignature: true,
 *                     certSign: true,
 *                     crlSign: true,
 *                 },
 *                 extendedKeyUsage: {
 *                     serverAuth: true,
 *                 },
 *             },
 *         },
 *     },
 *     lifetime: "86400s",
 *     keySpec: {
 *         algorithm: "RSA_PKCS1_4096_SHA256",
 *     },
 *     deletionProtection: false,
 *     skipGracePeriod: true,
 *     ignoreActiveCertificatesOnDeletion: true,
 * });
 * const defaultCertificate = new gcp.certificateauthority.Certificate("defaultCertificate", {
 *     location: "us-central1",
 *     pool: defaultCaPool.name,
 *     lifetime: "860s",
 *     config: {
 *         subjectConfig: {
 *             subject: {
 *                 commonName: "san1.example.com",
 *                 countryCode: "us",
 *                 organization: "google",
 *                 organizationalUnit: "enterprise",
 *                 locality: "mountain view",
 *                 province: "california",
 *                 streetAddress: "1600 amphitheatre parkway",
 *                 postalCode: "94109",
 *             },
 *         },
 *         x509Config: {
 *             caOptions: {
 *                 isCa: false,
 *             },
 *             keyUsage: {
 *                 baseKeyUsage: {
 *                     crlSign: true,
 *                 },
 *                 extendedKeyUsage: {
 *                     serverAuth: true,
 *                 },
 *             },
 *         },
 *         publicKey: {
 *             format: "PEM",
 *             key: Buffer.from(fs.readFileSync("test-fixtures/rsa_public.pem"), 'binary').toString('base64'),
 *         },
 *     },
 * }, {
 *     dependsOn: [defaultAuthority],
 * });
 * ```
 *
 * ## Import
 *
 * Certificate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/certificate:Certificate default projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificates/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/certificate:Certificate default {{project}}/{{location}}/{{pool}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/certificate:Certificate default {{location}}/{{pool}}/{{name}}
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
    public static readonly __pulumiType = 'gcp:certificateauthority/certificate:Certificate';

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
     * The Certificate Authority ID that should issue the certificate. For example, to issue a Certificate from
     * a Certificate Authority with resource name `projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca`,
     * argument `pool` should be set to `projects/my-project/locations/us-central1/caPools/my-pool`, argument `certificateAuthority`
     * should be set to `my-ca`.
     */
    public readonly certificateAuthority!: pulumi.Output<string | undefined>;
    /**
     * Output only. Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
     * Structure is documented below.
     */
    public /*out*/ readonly certificateDescriptions!: pulumi.Output<outputs.certificateauthority.CertificateCertificateDescription[]>;
    /**
     * The resource name for a CertificateTemplate used to issue this certificate,
     * in the format `projects/*&#47;locations/*&#47;certificateTemplates/*`. If this is specified,
     * the caller must have the necessary permission to use this template. If this is
     * omitted, no template will be used. This template must be in the same location
     * as the Certificate.
     */
    public readonly certificateTemplate!: pulumi.Output<string | undefined>;
    /**
     * The config used to create a self-signed X.509 certificate or CSR.
     * Structure is documented below.
     */
    public readonly config!: pulumi.Output<outputs.certificateauthority.CertificateConfig | undefined>;
    /**
     * The time that this resource was created on the server.
     * This is in RFC3339 text format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The resource name of the issuing CertificateAuthority in the format `projects/*&#47;locations/*&#47;caPools/*&#47;certificateAuthorities/*`.
     */
    public /*out*/ readonly issuerCertificateAuthority!: pulumi.Output<string>;
    /**
     * Labels with user-defined metadata to apply to this resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
     * "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
     * fractional digits, terminated by 's'. Example: "3.5s".
     */
    public readonly lifetime!: pulumi.Output<string | undefined>;
    /**
     * Location of the Certificate. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     *
     *
     * - - -
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name for this Certificate.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Output only. The pem-encoded, signed X.509 certificate.
     */
    public /*out*/ readonly pemCertificate!: pulumi.Output<string>;
    /**
     * The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
     */
    public /*out*/ readonly pemCertificateChains!: pulumi.Output<string[]>;
    /**
     * (Deprecated)
     * Required. Expected to be in leaf-to-root order according to RFC 5246.
     *
     * @deprecated `pem_certificates` is deprecated and will be removed in a future major release. Use `pem_certificate_chain` instead.
     */
    public /*out*/ readonly pemCertificates!: pulumi.Output<string[]>;
    /**
     * Immutable. A pem-encoded X.509 certificate signing request (CSR).
     */
    public readonly pemCsr!: pulumi.Output<string | undefined>;
    /**
     * The name of the CaPool this Certificate belongs to.
     */
    public readonly pool!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Output only. Details regarding the revocation of this Certificate. This Certificate is
     * considered revoked if and only if this field is present.
     * Structure is documented below.
     */
    public /*out*/ readonly revocationDetails!: pulumi.Output<outputs.certificateauthority.CertificateRevocationDetail[]>;
    /**
     * Output only. The time at which this CertificateAuthority was updated.
     * This is in RFC3339 text format.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs | CertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateState | undefined;
            resourceInputs["certificateAuthority"] = state ? state.certificateAuthority : undefined;
            resourceInputs["certificateDescriptions"] = state ? state.certificateDescriptions : undefined;
            resourceInputs["certificateTemplate"] = state ? state.certificateTemplate : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["issuerCertificateAuthority"] = state ? state.issuerCertificateAuthority : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["lifetime"] = state ? state.lifetime : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pemCertificate"] = state ? state.pemCertificate : undefined;
            resourceInputs["pemCertificateChains"] = state ? state.pemCertificateChains : undefined;
            resourceInputs["pemCertificates"] = state ? state.pemCertificates : undefined;
            resourceInputs["pemCsr"] = state ? state.pemCsr : undefined;
            resourceInputs["pool"] = state ? state.pool : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["revocationDetails"] = state ? state.revocationDetails : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.pool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pool'");
            }
            resourceInputs["certificateAuthority"] = args ? args.certificateAuthority : undefined;
            resourceInputs["certificateTemplate"] = args ? args.certificateTemplate : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["lifetime"] = args ? args.lifetime : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pemCsr"] = args ? args.pemCsr : undefined;
            resourceInputs["pool"] = args ? args.pool : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["certificateDescriptions"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["issuerCertificateAuthority"] = undefined /*out*/;
            resourceInputs["pemCertificate"] = undefined /*out*/;
            resourceInputs["pemCertificateChains"] = undefined /*out*/;
            resourceInputs["pemCertificates"] = undefined /*out*/;
            resourceInputs["revocationDetails"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
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
     * The Certificate Authority ID that should issue the certificate. For example, to issue a Certificate from
     * a Certificate Authority with resource name `projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca`,
     * argument `pool` should be set to `projects/my-project/locations/us-central1/caPools/my-pool`, argument `certificateAuthority`
     * should be set to `my-ca`.
     */
    certificateAuthority?: pulumi.Input<string>;
    /**
     * Output only. Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
     * Structure is documented below.
     */
    certificateDescriptions?: pulumi.Input<pulumi.Input<inputs.certificateauthority.CertificateCertificateDescription>[]>;
    /**
     * The resource name for a CertificateTemplate used to issue this certificate,
     * in the format `projects/*&#47;locations/*&#47;certificateTemplates/*`. If this is specified,
     * the caller must have the necessary permission to use this template. If this is
     * omitted, no template will be used. This template must be in the same location
     * as the Certificate.
     */
    certificateTemplate?: pulumi.Input<string>;
    /**
     * The config used to create a self-signed X.509 certificate or CSR.
     * Structure is documented below.
     */
    config?: pulumi.Input<inputs.certificateauthority.CertificateConfig>;
    /**
     * The time that this resource was created on the server.
     * This is in RFC3339 text format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The resource name of the issuing CertificateAuthority in the format `projects/*&#47;locations/*&#47;caPools/*&#47;certificateAuthorities/*`.
     */
    issuerCertificateAuthority?: pulumi.Input<string>;
    /**
     * Labels with user-defined metadata to apply to this resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
     * "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
     * fractional digits, terminated by 's'. Example: "3.5s".
     */
    lifetime?: pulumi.Input<string>;
    /**
     * Location of the Certificate. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     *
     *
     * - - -
     */
    location?: pulumi.Input<string>;
    /**
     * The name for this Certificate.
     */
    name?: pulumi.Input<string>;
    /**
     * Output only. The pem-encoded, signed X.509 certificate.
     */
    pemCertificate?: pulumi.Input<string>;
    /**
     * The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
     */
    pemCertificateChains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Deprecated)
     * Required. Expected to be in leaf-to-root order according to RFC 5246.
     *
     * @deprecated `pem_certificates` is deprecated and will be removed in a future major release. Use `pem_certificate_chain` instead.
     */
    pemCertificates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Immutable. A pem-encoded X.509 certificate signing request (CSR).
     */
    pemCsr?: pulumi.Input<string>;
    /**
     * The name of the CaPool this Certificate belongs to.
     */
    pool?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Output only. Details regarding the revocation of this Certificate. This Certificate is
     * considered revoked if and only if this field is present.
     * Structure is documented below.
     */
    revocationDetails?: pulumi.Input<pulumi.Input<inputs.certificateauthority.CertificateRevocationDetail>[]>;
    /**
     * Output only. The time at which this CertificateAuthority was updated.
     * This is in RFC3339 text format.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * The Certificate Authority ID that should issue the certificate. For example, to issue a Certificate from
     * a Certificate Authority with resource name `projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca`,
     * argument `pool` should be set to `projects/my-project/locations/us-central1/caPools/my-pool`, argument `certificateAuthority`
     * should be set to `my-ca`.
     */
    certificateAuthority?: pulumi.Input<string>;
    /**
     * The resource name for a CertificateTemplate used to issue this certificate,
     * in the format `projects/*&#47;locations/*&#47;certificateTemplates/*`. If this is specified,
     * the caller must have the necessary permission to use this template. If this is
     * omitted, no template will be used. This template must be in the same location
     * as the Certificate.
     */
    certificateTemplate?: pulumi.Input<string>;
    /**
     * The config used to create a self-signed X.509 certificate or CSR.
     * Structure is documented below.
     */
    config?: pulumi.Input<inputs.certificateauthority.CertificateConfig>;
    /**
     * Labels with user-defined metadata to apply to this resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
     * "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
     * fractional digits, terminated by 's'. Example: "3.5s".
     */
    lifetime?: pulumi.Input<string>;
    /**
     * Location of the Certificate. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     *
     *
     * - - -
     */
    location: pulumi.Input<string>;
    /**
     * The name for this Certificate.
     */
    name?: pulumi.Input<string>;
    /**
     * Immutable. A pem-encoded X.509 certificate signing request (CSR).
     */
    pemCsr?: pulumi.Input<string>;
    /**
     * The name of the CaPool this Certificate belongs to.
     */
    pool: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

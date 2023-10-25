// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * An Environment Keystore Alias for Self Signed Certificate Format in Apigee
 *
 * To get more information about KeystoresAliasesSelfSignedCert, see:
 *
 * * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.keystores.aliases/create)
 * * How-to Guides
 *     * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)
 *
 * ## Example Usage
 * ### Apigee Env Keystore Alias Self Signed Cert
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = new gcp.organizations.Project("project", {
 *     projectId: "my-project",
 *     orgId: "123456789",
 *     billingAccount: "000000-0000000-0000000-000000",
 * });
 * const apigee = new gcp.projects.Service("apigee", {
 *     project: project.projectId,
 *     service: "apigee.googleapis.com",
 * });
 * const servicenetworking = new gcp.projects.Service("servicenetworking", {
 *     project: project.projectId,
 *     service: "servicenetworking.googleapis.com",
 * }, {
 *     dependsOn: [apigee],
 * });
 * const compute = new gcp.projects.Service("compute", {
 *     project: project.projectId,
 *     service: "compute.googleapis.com",
 * }, {
 *     dependsOn: [servicenetworking],
 * });
 * const apigeeNetwork = new gcp.compute.Network("apigeeNetwork", {project: project.projectId}, {
 *     dependsOn: [compute],
 * });
 * const apigeeRange = new gcp.compute.GlobalAddress("apigeeRange", {
 *     purpose: "VPC_PEERING",
 *     addressType: "INTERNAL",
 *     prefixLength: 16,
 *     network: apigeeNetwork.id,
 *     project: project.projectId,
 * });
 * const apigeeVpcConnection = new gcp.servicenetworking.Connection("apigeeVpcConnection", {
 *     network: apigeeNetwork.id,
 *     service: "servicenetworking.googleapis.com",
 *     reservedPeeringRanges: [apigeeRange.name],
 * }, {
 *     dependsOn: [servicenetworking],
 * });
 * const apigeeOrg = new gcp.apigee.Organization("apigeeOrg", {
 *     analyticsRegion: "us-central1",
 *     projectId: project.projectId,
 *     authorizedNetwork: apigeeNetwork.id,
 * }, {
 *     dependsOn: [
 *         apigeeVpcConnection,
 *         apigee,
 *     ],
 * });
 * const apigeeEnvironmentKeystoreSsAliasEnvironment = new gcp.apigee.Environment("apigeeEnvironmentKeystoreSsAliasEnvironment", {
 *     orgId: apigeeOrg.id,
 *     description: "Apigee Environment",
 *     displayName: "environment-1",
 * });
 * const apigeeEnvironmentKeystoreAlias = new gcp.apigee.EnvKeystore("apigeeEnvironmentKeystoreAlias", {envId: apigeeEnvironmentKeystoreSsAliasEnvironment.id});
 * const apigeeEnvironmentKeystoreSsAliasKeystoresAliasesSelfSignedCert = new gcp.apigee.KeystoresAliasesSelfSignedCert("apigeeEnvironmentKeystoreSsAliasKeystoresAliasesSelfSignedCert", {
 *     environment: apigeeEnvironmentKeystoreSsAliasEnvironment.name,
 *     orgId: apigeeOrg.name,
 *     keystore: apigeeEnvironmentKeystoreAlias.name,
 *     alias: "alias",
 *     keySize: "1024",
 *     sigAlg: "SHA512withRSA",
 *     certValidityInDays: 4,
 *     subject: {
 *         commonName: "selfsigned_example",
 *         countryCode: "US",
 *         locality: "TX",
 *         org: "CCE",
 *         orgUnit: "PSO",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * KeystoresAliasesSelfSignedCert can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:apigee/keystoresAliasesSelfSignedCert:KeystoresAliasesSelfSignedCert default organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:apigee/keystoresAliasesSelfSignedCert:KeystoresAliasesSelfSignedCert default {{org_id}}/{{environment}}/{{keystore}}/{{alias}}
 * ```
 */
export class KeystoresAliasesSelfSignedCert extends pulumi.CustomResource {
    /**
     * Get an existing KeystoresAliasesSelfSignedCert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeystoresAliasesSelfSignedCertState, opts?: pulumi.CustomResourceOptions): KeystoresAliasesSelfSignedCert {
        return new KeystoresAliasesSelfSignedCert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:apigee/keystoresAliasesSelfSignedCert:KeystoresAliasesSelfSignedCert';

    /**
     * Returns true if the given object is an instance of KeystoresAliasesSelfSignedCert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeystoresAliasesSelfSignedCert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeystoresAliasesSelfSignedCert.__pulumiType;
    }

    /**
     * Alias for the key/certificate pair. Values must match the regular expression [\w\s-.]{1,255}.
     * This must be provided for all formats except selfsignedcert; self-signed certs may specify the alias in either
     * this parameter or the JSON body.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * Validity duration of certificate, in days. Accepts positive non-zero value. Defaults to 365.
     */
    public readonly certValidityInDays!: pulumi.Output<number | undefined>;
    /**
     * Chain of certificates under this alias.
     * Structure is documented below.
     */
    public /*out*/ readonly certsInfos!: pulumi.Output<outputs.apigee.KeystoresAliasesSelfSignedCertCertsInfo[]>;
    /**
     * The Apigee environment name
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * Key size. Default and maximum value is 2048 bits.
     */
    public readonly keySize!: pulumi.Output<string | undefined>;
    /**
     * The Apigee keystore name associated in an Apigee environment
     */
    public readonly keystore!: pulumi.Output<string>;
    /**
     * The Apigee Organization name associated with the Apigee environment
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * Signature algorithm to generate private key. Valid values are SHA512withRSA, SHA384withRSA, and SHA256withRSA
     */
    public readonly sigAlg!: pulumi.Output<string>;
    /**
     * Subject details.
     * Structure is documented below.
     */
    public readonly subject!: pulumi.Output<outputs.apigee.KeystoresAliasesSelfSignedCertSubject>;
    /**
     * List of alternative host names. Maximum length is 255 characters for each value.
     * Structure is documented below.
     */
    public readonly subjectAlternativeDnsNames!: pulumi.Output<outputs.apigee.KeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames | undefined>;
    /**
     * Optional.Type of Alias
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a KeystoresAliasesSelfSignedCert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeystoresAliasesSelfSignedCertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeystoresAliasesSelfSignedCertArgs | KeystoresAliasesSelfSignedCertState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeystoresAliasesSelfSignedCertState | undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["certValidityInDays"] = state ? state.certValidityInDays : undefined;
            resourceInputs["certsInfos"] = state ? state.certsInfos : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["keySize"] = state ? state.keySize : undefined;
            resourceInputs["keystore"] = state ? state.keystore : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["sigAlg"] = state ? state.sigAlg : undefined;
            resourceInputs["subject"] = state ? state.subject : undefined;
            resourceInputs["subjectAlternativeDnsNames"] = state ? state.subjectAlternativeDnsNames : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as KeystoresAliasesSelfSignedCertArgs | undefined;
            if ((!args || args.alias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alias'");
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
            if ((!args || args.sigAlg === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sigAlg'");
            }
            if ((!args || args.subject === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subject'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["certValidityInDays"] = args ? args.certValidityInDays : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["keySize"] = args ? args.keySize : undefined;
            resourceInputs["keystore"] = args ? args.keystore : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["sigAlg"] = args ? args.sigAlg : undefined;
            resourceInputs["subject"] = args ? args.subject : undefined;
            resourceInputs["subjectAlternativeDnsNames"] = args ? args.subjectAlternativeDnsNames : undefined;
            resourceInputs["certsInfos"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KeystoresAliasesSelfSignedCert.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeystoresAliasesSelfSignedCert resources.
 */
export interface KeystoresAliasesSelfSignedCertState {
    /**
     * Alias for the key/certificate pair. Values must match the regular expression [\w\s-.]{1,255}.
     * This must be provided for all formats except selfsignedcert; self-signed certs may specify the alias in either
     * this parameter or the JSON body.
     */
    alias?: pulumi.Input<string>;
    /**
     * Validity duration of certificate, in days. Accepts positive non-zero value. Defaults to 365.
     */
    certValidityInDays?: pulumi.Input<number>;
    /**
     * Chain of certificates under this alias.
     * Structure is documented below.
     */
    certsInfos?: pulumi.Input<pulumi.Input<inputs.apigee.KeystoresAliasesSelfSignedCertCertsInfo>[]>;
    /**
     * The Apigee environment name
     */
    environment?: pulumi.Input<string>;
    /**
     * Key size. Default and maximum value is 2048 bits.
     */
    keySize?: pulumi.Input<string>;
    /**
     * The Apigee keystore name associated in an Apigee environment
     */
    keystore?: pulumi.Input<string>;
    /**
     * The Apigee Organization name associated with the Apigee environment
     */
    orgId?: pulumi.Input<string>;
    /**
     * Signature algorithm to generate private key. Valid values are SHA512withRSA, SHA384withRSA, and SHA256withRSA
     */
    sigAlg?: pulumi.Input<string>;
    /**
     * Subject details.
     * Structure is documented below.
     */
    subject?: pulumi.Input<inputs.apigee.KeystoresAliasesSelfSignedCertSubject>;
    /**
     * List of alternative host names. Maximum length is 255 characters for each value.
     * Structure is documented below.
     */
    subjectAlternativeDnsNames?: pulumi.Input<inputs.apigee.KeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames>;
    /**
     * Optional.Type of Alias
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeystoresAliasesSelfSignedCert resource.
 */
export interface KeystoresAliasesSelfSignedCertArgs {
    /**
     * Alias for the key/certificate pair. Values must match the regular expression [\w\s-.]{1,255}.
     * This must be provided for all formats except selfsignedcert; self-signed certs may specify the alias in either
     * this parameter or the JSON body.
     */
    alias: pulumi.Input<string>;
    /**
     * Validity duration of certificate, in days. Accepts positive non-zero value. Defaults to 365.
     */
    certValidityInDays?: pulumi.Input<number>;
    /**
     * The Apigee environment name
     */
    environment: pulumi.Input<string>;
    /**
     * Key size. Default and maximum value is 2048 bits.
     */
    keySize?: pulumi.Input<string>;
    /**
     * The Apigee keystore name associated in an Apigee environment
     */
    keystore: pulumi.Input<string>;
    /**
     * The Apigee Organization name associated with the Apigee environment
     */
    orgId: pulumi.Input<string>;
    /**
     * Signature algorithm to generate private key. Valid values are SHA512withRSA, SHA384withRSA, and SHA256withRSA
     */
    sigAlg: pulumi.Input<string>;
    /**
     * Subject details.
     * Structure is documented below.
     */
    subject: pulumi.Input<inputs.apigee.KeystoresAliasesSelfSignedCertSubject>;
    /**
     * List of alternative host names. Maximum length is 255 characters for each value.
     * Structure is documented below.
     */
    subjectAlternativeDnsNames?: pulumi.Input<inputs.apigee.KeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames>;
}

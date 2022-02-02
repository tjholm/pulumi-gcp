// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Certificate Authority Service provides reusable and parameterized templates that you can use for common certificate issuance scenarios. A certificate template represents a relatively static and well-defined certificate issuance schema within an organization.  A certificate template can essentially become a full-fledged vertical certificate issuance framework.
 *
 * For more information, see:
 * * [Understanding Certificate Templates](https://cloud.google.com/certificate-authority-service/docs/certificate-template)
 * * [Common configurations and Certificate Profiles](https://cloud.google.com/certificate-authority-service/docs/certificate-profile)
 * ## Example Usage
 * ### Basic_certificate_template
 * An example of a basic privateca certificate template
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const primary = new gcp.certificateauthority.CertificateTemplate("primary", {
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
 *     labels: {
 *         "label-two": "value-two",
 *     },
 *     location: "us-west1",
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
 *             critical: true,
 *             objectId: {
 *                 objectIdPaths: [
 *                     1,
 *                     6,
 *                 ],
 *             },
 *             value: "c3RyaW5nCg==",
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
 *     project: "my-project-name",
 * });
 * ```
 *
 * ## Import
 *
 * CertificateTemplate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default {{location}}/{{name}}
 * ```
 */
export class CertificateTemplate extends pulumi.CustomResource {
    /**
     * Get an existing CertificateTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateTemplateState, opts?: pulumi.CustomResourceOptions): CertificateTemplate {
        return new CertificateTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:certificateauthority/certificateTemplate:CertificateTemplate';

    /**
     * Returns true if the given object is an instance of CertificateTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateTemplate.__pulumiType;
    }

    /**
     * Output only. The time at which this CertificateTemplate was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
     */
    public readonly identityConstraints!: pulumi.Output<outputs.certificateauthority.CertificateTemplateIdentityConstraints | undefined>;
    /**
     * Optional. Labels with user-defined metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name for this CertificateTemplate in the format `projects/*&#47;locations/*&#47;certificateTemplates/*`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baselineValues that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
     */
    public readonly passthroughExtensions!: pulumi.Output<outputs.certificateauthority.CertificateTemplatePassthroughExtensions | undefined>;
    /**
     * Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baselineValues for the same properties, the certificate issuance request will fail.
     */
    public readonly predefinedValues!: pulumi.Output<outputs.certificateauthority.CertificateTemplatePredefinedValues | undefined>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Output only. The time at which this CertificateTemplate was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a CertificateTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateTemplateArgs | CertificateTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateTemplateState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["identityConstraints"] = state ? state.identityConstraints : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["passthroughExtensions"] = state ? state.passthroughExtensions : undefined;
            resourceInputs["predefinedValues"] = state ? state.predefinedValues : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as CertificateTemplateArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["identityConstraints"] = args ? args.identityConstraints : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["passthroughExtensions"] = args ? args.passthroughExtensions : undefined;
            resourceInputs["predefinedValues"] = args ? args.predefinedValues : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CertificateTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertificateTemplate resources.
 */
export interface CertificateTemplateState {
    /**
     * Output only. The time at which this CertificateTemplate was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
     */
    identityConstraints?: pulumi.Input<inputs.certificateauthority.CertificateTemplateIdentityConstraints>;
    /**
     * Optional. Labels with user-defined metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
     */
    location?: pulumi.Input<string>;
    /**
     * The resource name for this CertificateTemplate in the format `projects/*&#47;locations/*&#47;certificateTemplates/*`.
     */
    name?: pulumi.Input<string>;
    /**
     * Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baselineValues that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
     */
    passthroughExtensions?: pulumi.Input<inputs.certificateauthority.CertificateTemplatePassthroughExtensions>;
    /**
     * Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baselineValues for the same properties, the certificate issuance request will fail.
     */
    predefinedValues?: pulumi.Input<inputs.certificateauthority.CertificateTemplatePredefinedValues>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Output only. The time at which this CertificateTemplate was updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CertificateTemplate resource.
 */
export interface CertificateTemplateArgs {
    /**
     * Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
     */
    identityConstraints?: pulumi.Input<inputs.certificateauthority.CertificateTemplateIdentityConstraints>;
    /**
     * Optional. Labels with user-defined metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
     */
    location: pulumi.Input<string>;
    /**
     * The resource name for this CertificateTemplate in the format `projects/*&#47;locations/*&#47;certificateTemplates/*`.
     */
    name?: pulumi.Input<string>;
    /**
     * Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baselineValues that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
     */
    passthroughExtensions?: pulumi.Input<inputs.certificateauthority.CertificateTemplatePassthroughExtensions>;
    /**
     * Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baselineValues for the same properties, the certificate issuance request will fail.
     */
    predefinedValues?: pulumi.Input<inputs.certificateauthority.CertificateTemplatePredefinedValues>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
}

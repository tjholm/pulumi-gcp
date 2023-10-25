// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A RegionSslCertificate resource, used for HTTPS load balancing. This resource
 * provides a mechanism to upload an SSL key and certificate to
 * the load balancer to serve secure connections from the user.
 *
 * To get more information about RegionSslCertificate, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionSslCertificates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
 *
 * > **Warning:** All arguments including `certificate` and `privateKey` will be stored in the raw
 * state as plain-text.
 *
 * ## Example Usage
 * ### Region Ssl Certificate Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.compute.RegionSslCertificate("default", {
 *     region: "us-central1",
 *     namePrefix: "my-certificate-",
 *     description: "a description",
 *     privateKey: fs.readFileSync("path/to/private.key"),
 *     certificate: fs.readFileSync("path/to/certificate.crt"),
 * });
 * ```
 * ### Region Ssl Certificate Random Provider
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as crypto from "crypto";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 * import * as random from "@pulumi/random";
 *
 * function computeFilebase64sha256(path string) string {
 * 	const fileData = Buffer.from(fs.readFileSync(path), 'binary')
 * 	return crypto.createHash('sha256').update(fileData).digest('hex')
 * }
 *
 * // You may also want to control name generation explicitly:
 * const _default = new gcp.compute.RegionSslCertificate("default", {
 *     region: "us-central1",
 *     privateKey: fs.readFileSync("path/to/private.key"),
 *     certificate: fs.readFileSync("path/to/certificate.crt"),
 * });
 * const certificate = new random.RandomId("certificate", {
 *     byteLength: 4,
 *     prefix: "my-certificate-",
 *     keepers: {
 *         private_key: computeFilebase64sha256("path/to/private.key"),
 *         certificate: computeFilebase64sha256("path/to/certificate.crt"),
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * RegionSslCertificate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionSslCertificate:RegionSslCertificate default projects/{{project}}/regions/{{region}}/sslCertificates/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionSslCertificate:RegionSslCertificate default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionSslCertificate:RegionSslCertificate default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionSslCertificate:RegionSslCertificate default {{name}}
 * ```
 */
export class RegionSslCertificate extends pulumi.CustomResource {
    /**
     * Get an existing RegionSslCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionSslCertificateState, opts?: pulumi.CustomResourceOptions): RegionSslCertificate {
        return new RegionSslCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionSslCertificate:RegionSslCertificate';

    /**
     * Returns true if the given object is an instance of RegionSslCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionSslCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionSslCertificate.__pulumiType;
    }

    /**
     * The certificate in PEM format.
     * The certificate chain must be no greater than 5 certs long.
     * The chain must include at least one intermediate cert.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * The unique identifier for the resource.
     */
    public /*out*/ readonly certificateId!: pulumi.Output<number>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Expire time of the certificate in RFC3339 text format.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     *
     * These are in the same namespace as the managed SSL certificates.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * The write-only private key in PEM format.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     *
     *
     * - - -
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The Region in which the created regional ssl certificate should reside.
     * If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a RegionSslCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionSslCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionSslCertificateArgs | RegionSslCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegionSslCertificateState | undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expireTime"] = state ? state.expireTime : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as RegionSslCertificateArgs | undefined;
            if ((!args || args.certificate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificate'");
            }
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            resourceInputs["certificate"] = args?.certificate ? pulumi.secret(args.certificate) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["certificateId"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["certificate", "privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(RegionSslCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionSslCertificate resources.
 */
export interface RegionSslCertificateState {
    /**
     * The certificate in PEM format.
     * The certificate chain must be no greater than 5 certs long.
     * The chain must include at least one intermediate cert.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    certificate?: pulumi.Input<string>;
    /**
     * The unique identifier for the resource.
     */
    certificateId?: pulumi.Input<number>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Expire time of the certificate in RFC3339 text format.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     *
     * These are in the same namespace as the managed SSL certificates.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The write-only private key in PEM format.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     *
     *
     * - - -
     */
    privateKey?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The Region in which the created regional ssl certificate should reside.
     * If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RegionSslCertificate resource.
 */
export interface RegionSslCertificateArgs {
    /**
     * The certificate in PEM format.
     * The certificate chain must be no greater than 5 certs long.
     * The chain must include at least one intermediate cert.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    certificate: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     *
     * These are in the same namespace as the managed SSL certificates.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The write-only private key in PEM format.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     *
     *
     * - - -
     */
    privateKey: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The Region in which the created regional ssl certificate should reside.
     * If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
}

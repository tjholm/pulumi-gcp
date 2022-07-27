// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * An SslCertificate resource, used for HTTPS load balancing.  This resource
 * represents a certificate for which the certificate secrets are created and
 * managed by Google.
 *
 * For a resource where you provide the key, see the
 * SSL Certificate resource.
 *
 * To get more information about ManagedSslCertificate, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
 *
 * > **Warning:** This resource should be used with extreme caution!  Provisioning an SSL
 * certificate is complex.  Ensure that you understand the lifecycle of a
 * certificate before attempting complex tasks like cert rotation automatically.
 * This resource will "return" as soon as the certificate object is created,
 * but post-creation the certificate object will go through a "provisioning"
 * process.  The provisioning process can complete only when the domain name
 * for which the certificate is created points to a target pool which, itself,
 * points at the certificate.  Depending on your DNS provider, this may take
 * some time, and migrating from self-managed certificates to Google-managed
 * certificates may entail some downtime while the certificate provisions.
 *
 * In conclusion: Be extremely cautious.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * ManagedSslCertificate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/mangedSslCertificate:MangedSslCertificate default projects/{{project}}/global/sslCertificates/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/mangedSslCertificate:MangedSslCertificate default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/mangedSslCertificate:MangedSslCertificate default {{name}}
 * ```
 *
 * @deprecated gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate
 */
export class MangedSslCertificate extends pulumi.CustomResource {
    /**
     * Get an existing MangedSslCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MangedSslCertificateState, opts?: pulumi.CustomResourceOptions): MangedSslCertificate {
        pulumi.log.warn("MangedSslCertificate is deprecated: gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate")
        return new MangedSslCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/mangedSslCertificate:MangedSslCertificate';

    /**
     * Returns true if the given object is an instance of MangedSslCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MangedSslCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MangedSslCertificate.__pulumiType;
    }

    /**
     * The unique identifier for the resource.
     */
    public readonly certificateId!: pulumi.Output<number>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Expire time of the certificate.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * Properties relevant to a managed certificate.  These will be used if the
     * certificate is managed (as indicated by a value of `MANAGED` in `type`).
     * Structure is documented below.
     */
    public readonly managed!: pulumi.Output<outputs.compute.MangedSslCertificateManaged | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Domains associated with the certificate via Subject Alternative Name.
     */
    public /*out*/ readonly subjectAlternativeNames!: pulumi.Output<string[]>;
    /**
     * Enum field whose value is always `MANAGED` - used to signal to the API
     * which type this is.
     * Default value is `MANAGED`.
     * Possible values are `MANAGED`.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a MangedSslCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate */
    constructor(name: string, args?: MangedSslCertificateArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate */
    constructor(name: string, argsOrState?: MangedSslCertificateArgs | MangedSslCertificateState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("MangedSslCertificate is deprecated: gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MangedSslCertificateState | undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expireTime"] = state ? state.expireTime : undefined;
            resourceInputs["managed"] = state ? state.managed : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["subjectAlternativeNames"] = state ? state.subjectAlternativeNames : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as MangedSslCertificateArgs | undefined;
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["managed"] = args ? args.managed : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["subjectAlternativeNames"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MangedSslCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MangedSslCertificate resources.
 */
export interface MangedSslCertificateState {
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
     * Expire time of the certificate.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Properties relevant to a managed certificate.  These will be used if the
     * certificate is managed (as indicated by a value of `MANAGED` in `type`).
     * Structure is documented below.
     */
    managed?: pulumi.Input<inputs.compute.MangedSslCertificateManaged>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * Domains associated with the certificate via Subject Alternative Name.
     */
    subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enum field whose value is always `MANAGED` - used to signal to the API
     * which type this is.
     * Default value is `MANAGED`.
     * Possible values are `MANAGED`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MangedSslCertificate resource.
 */
export interface MangedSslCertificateArgs {
    /**
     * The unique identifier for the resource.
     */
    certificateId?: pulumi.Input<number>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Properties relevant to a managed certificate.  These will be used if the
     * certificate is managed (as indicated by a value of `MANAGED` in `type`).
     * Structure is documented below.
     */
    managed?: pulumi.Input<inputs.compute.MangedSslCertificateManaged>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Enum field whose value is always `MANAGED` - used to signal to the API
     * which type this is.
     * Default value is `MANAGED`.
     * Possible values are `MANAGED`.
     */
    type?: pulumi.Input<string>;
}

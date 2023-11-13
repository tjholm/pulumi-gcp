// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a Microsoft AD domain
 *
 * To get more information about Domain, see:
 *
 * * [API documentation](https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains)
 * * How-to Guides
 *     * [Managed Microsoft Active Directory Quickstart](https://cloud.google.com/managed-microsoft-ad/docs/quickstarts)
 *
 * ## Example Usage
 * ### Active Directory Domain Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const ad_domain = new gcp.activedirectory.Domain("ad-domain", {
 *     domainName: "tfgen.org.com",
 *     locations: ["us-central1"],
 *     reservedIpRange: "192.168.255.0/24",
 * });
 * ```
 *
 * ## Import
 *
 * Domain can be imported using any of these accepted formats:
 *
 * ```sh
 *  $ pulumi import gcp:activedirectory/domain:Domain default {{name}}
 * ```
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:activedirectory/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * The name of delegated administrator account used to perform Active Directory operations.
     * If not specified, setupadmin will be used.
     */
    public readonly admin!: pulumi.Output<string | undefined>;
    /**
     * The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
     * If CIDR subnets overlap between networks, domain creation will fail.
     */
    public readonly authorizedNetworks!: pulumi.Output<string[] | undefined>;
    /**
     * The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
     * https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
     *
     *
     * - - -
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The fully-qualified domain name of the exposed domain used by clients to connect to the service.
     * Similar to what would be chosen for an Active Directory set up on an internal network.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * Resource labels that can contain user-provided metadata
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
     * e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
     */
    public readonly locations!: pulumi.Output<string[]>;
    /**
     * The unique name of the domain using the format: `projects/{project}/locations/global/domains/{domainName}`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
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
     * The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
     * Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
     */
    public readonly reservedIpRange!: pulumi.Output<string>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainState | undefined;
            resourceInputs["admin"] = state ? state.admin : undefined;
            resourceInputs["authorizedNetworks"] = state ? state.authorizedNetworks : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["fqdn"] = state ? state.fqdn : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["locations"] = state ? state.locations : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["reservedIpRange"] = state ? state.reservedIpRange : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.locations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locations'");
            }
            if ((!args || args.reservedIpRange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reservedIpRange'");
            }
            resourceInputs["admin"] = args ? args.admin : undefined;
            resourceInputs["authorizedNetworks"] = args ? args.authorizedNetworks : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["locations"] = args ? args.locations : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["reservedIpRange"] = args ? args.reservedIpRange : undefined;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["fqdn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Domain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * The name of delegated administrator account used to perform Active Directory operations.
     * If not specified, setupadmin will be used.
     */
    admin?: pulumi.Input<string>;
    /**
     * The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
     * If CIDR subnets overlap between networks, domain creation will fail.
     */
    authorizedNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
     * https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
     *
     *
     * - - -
     */
    domainName?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The fully-qualified domain name of the exposed domain used by clients to connect to the service.
     * Similar to what would be chosen for an Active Directory set up on an internal network.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Resource labels that can contain user-provided metadata
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
     * e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
     */
    locations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique name of the domain using the format: `projects/{project}/locations/global/domains/{domainName}`.
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
     * The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
     * Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
     */
    reservedIpRange?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * The name of delegated administrator account used to perform Active Directory operations.
     * If not specified, setupadmin will be used.
     */
    admin?: pulumi.Input<string>;
    /**
     * The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
     * If CIDR subnets overlap between networks, domain creation will fail.
     */
    authorizedNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
     * https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
     *
     *
     * - - -
     */
    domainName: pulumi.Input<string>;
    /**
     * Resource labels that can contain user-provided metadata
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
     * e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
     */
    locations: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
     * Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
     */
    reservedIpRange: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Clouddomains Registration Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myRegistration = new gcp.clouddomains.Registration("myRegistration", {
 *     contactSettings: {
 *         adminContact: {
 *             email: "user@example.com",
 *             phoneNumber: "+12345000000",
 *             postalAddress: {
 *                 addressLines: ["1234 Example street"],
 *                 administrativeArea: "CA",
 *                 locality: "Example City",
 *                 postalCode: "95050",
 *                 recipients: ["example recipient"],
 *                 regionCode: "US",
 *             },
 *         },
 *         privacy: "REDACTED_CONTACT_DATA",
 *         registrantContact: {
 *             email: "user@example.com",
 *             phoneNumber: "+12345000000",
 *             postalAddress: {
 *                 addressLines: ["1234 Example street"],
 *                 administrativeArea: "CA",
 *                 locality: "Example City",
 *                 postalCode: "95050",
 *                 recipients: ["example recipient"],
 *                 regionCode: "US",
 *             },
 *         },
 *         technicalContact: {
 *             email: "user@example.com",
 *             phoneNumber: "+12345000000",
 *             postalAddress: {
 *                 addressLines: ["1234 Example street"],
 *                 administrativeArea: "CA",
 *                 locality: "Example City",
 *                 postalCode: "95050",
 *                 recipients: ["example recipient"],
 *                 regionCode: "US",
 *             },
 *         },
 *     },
 *     dnsSettings: {
 *         customDns: {
 *             nameServers: [
 *                 "ns-cloud-a1.googledomains.com.",
 *                 "ns-cloud-a2.googledomains.com.",
 *                 "ns-cloud-a3.googledomains.com.",
 *                 "ns-cloud-a4.googledomains.com.",
 *             ],
 *         },
 *     },
 *     domainName: "example-domain.com",
 *     labels: {
 *         labelkey: "labelvalue",
 *     },
 *     location: "global",
 *     yearlyPrice: {
 *         currencyCode: "USD",
 *         units: "12",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Registration can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/registrations/{{domain_name}}` * `{{project}}/{{location}}/{{domain_name}}` * `{{location}}/{{domain_name}}` In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Registration using one of the formats above. For exampletf import {
 *
 *  id = "projects/{{project}}/locations/{{location}}/registrations/{{domain_name}}"
 *
 *  to = google_clouddomains_registration.default }
 *
 * ```sh
 *  $ pulumi import gcp:clouddomains/registration:Registration When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Registration can be imported using one of the formats above. For example
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:clouddomains/registration:Registration default projects/{{project}}/locations/{{location}}/registrations/{{domain_name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:clouddomains/registration:Registration default {{project}}/{{location}}/{{domain_name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:clouddomains/registration:Registration default {{location}}/{{domain_name}}
 * ```
 */
export class Registration extends pulumi.CustomResource {
    /**
     * Get an existing Registration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistrationState, opts?: pulumi.CustomResourceOptions): Registration {
        return new Registration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:clouddomains/registration:Registration';

    /**
     * Returns true if the given object is an instance of Registration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registration.__pulumiType;
    }

    /**
     * The list of contact notices that the caller acknowledges. Possible value is PUBLIC_CONTACT_DATA_ACKNOWLEDGEMENT
     */
    public readonly contactNotices!: pulumi.Output<string[] | undefined>;
    /**
     * Required. Settings for contact information linked to the Registration.
     * Structure is documented below.
     */
    public readonly contactSettings!: pulumi.Output<outputs.clouddomains.RegistrationContactSettings>;
    /**
     * Output only. Time at which the automation was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Settings controlling the DNS configuration of the Registration.
     * Structure is documented below.
     */
    public readonly dnsSettings!: pulumi.Output<outputs.clouddomains.RegistrationDnsSettings | undefined>;
    /**
     * Required. The domain name. Unicode domain names must be expressed in Punycode format.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The list of domain notices that you acknowledge. Possible value is HSTS_PRELOADED
     */
    public readonly domainNotices!: pulumi.Output<string[] | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Output only. Time at which the automation was updated.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * Output only. The set of issues with the Registration that require attention.
     */
    public /*out*/ readonly issues!: pulumi.Output<string[]>;
    /**
     * Set of labels associated with the Registration.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location for the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Settings for management of the Registration, including renewal, billing, and transfer
     * Structure is documented below.
     */
    public readonly managementSettings!: pulumi.Output<outputs.clouddomains.RegistrationManagementSettings>;
    /**
     * Output only. Name of the Registration resource, in the format projects/*&#47;locations/*&#47;registrations/<domain_name>.
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
     * Output only. The reason the domain registration failed. Only set for domains in REGISTRATION_FAILED state.
     */
    public /*out*/ readonly registerFailureReason!: pulumi.Output<string>;
    /**
     * Output only. The current state of the Registration.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Output only. Set of options for the contactSettings.privacy field that this Registration supports.
     */
    public /*out*/ readonly supportedPrivacies!: pulumi.Output<string[]>;
    /**
     * Required. Yearly price to register or renew the domain. The value that should be put here can be obtained from
     * registrations.retrieveRegisterParameters or registrations.searchDomains calls.
     * Structure is documented below.
     */
    public readonly yearlyPrice!: pulumi.Output<outputs.clouddomains.RegistrationYearlyPrice>;

    /**
     * Create a Registration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistrationArgs | RegistrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegistrationState | undefined;
            resourceInputs["contactNotices"] = state ? state.contactNotices : undefined;
            resourceInputs["contactSettings"] = state ? state.contactSettings : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["dnsSettings"] = state ? state.dnsSettings : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["domainNotices"] = state ? state.domainNotices : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["expireTime"] = state ? state.expireTime : undefined;
            resourceInputs["issues"] = state ? state.issues : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["managementSettings"] = state ? state.managementSettings : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["registerFailureReason"] = state ? state.registerFailureReason : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["supportedPrivacies"] = state ? state.supportedPrivacies : undefined;
            resourceInputs["yearlyPrice"] = state ? state.yearlyPrice : undefined;
        } else {
            const args = argsOrState as RegistrationArgs | undefined;
            if ((!args || args.contactSettings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contactSettings'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.yearlyPrice === undefined) && !opts.urn) {
                throw new Error("Missing required property 'yearlyPrice'");
            }
            resourceInputs["contactNotices"] = args ? args.contactNotices : undefined;
            resourceInputs["contactSettings"] = args ? args.contactSettings : undefined;
            resourceInputs["dnsSettings"] = args ? args.dnsSettings : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["domainNotices"] = args ? args.domainNotices : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["managementSettings"] = args ? args.managementSettings : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["yearlyPrice"] = args ? args.yearlyPrice : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["issues"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["registerFailureReason"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["supportedPrivacies"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Registration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Registration resources.
 */
export interface RegistrationState {
    /**
     * The list of contact notices that the caller acknowledges. Possible value is PUBLIC_CONTACT_DATA_ACKNOWLEDGEMENT
     */
    contactNotices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. Settings for contact information linked to the Registration.
     * Structure is documented below.
     */
    contactSettings?: pulumi.Input<inputs.clouddomains.RegistrationContactSettings>;
    /**
     * Output only. Time at which the automation was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Settings controlling the DNS configuration of the Registration.
     * Structure is documented below.
     */
    dnsSettings?: pulumi.Input<inputs.clouddomains.RegistrationDnsSettings>;
    /**
     * Required. The domain name. Unicode domain names must be expressed in Punycode format.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The list of domain notices that you acknowledge. Possible value is HSTS_PRELOADED
     */
    domainNotices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Output only. Time at which the automation was updated.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Output only. The set of issues with the Registration that require attention.
     */
    issues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of labels associated with the Registration.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location?: pulumi.Input<string>;
    /**
     * Settings for management of the Registration, including renewal, billing, and transfer
     * Structure is documented below.
     */
    managementSettings?: pulumi.Input<inputs.clouddomains.RegistrationManagementSettings>;
    /**
     * Output only. Name of the Registration resource, in the format projects/*&#47;locations/*&#47;registrations/<domain_name>.
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
     * Output only. The reason the domain registration failed. Only set for domains in REGISTRATION_FAILED state.
     */
    registerFailureReason?: pulumi.Input<string>;
    /**
     * Output only. The current state of the Registration.
     */
    state?: pulumi.Input<string>;
    /**
     * Output only. Set of options for the contactSettings.privacy field that this Registration supports.
     */
    supportedPrivacies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. Yearly price to register or renew the domain. The value that should be put here can be obtained from
     * registrations.retrieveRegisterParameters or registrations.searchDomains calls.
     * Structure is documented below.
     */
    yearlyPrice?: pulumi.Input<inputs.clouddomains.RegistrationYearlyPrice>;
}

/**
 * The set of arguments for constructing a Registration resource.
 */
export interface RegistrationArgs {
    /**
     * The list of contact notices that the caller acknowledges. Possible value is PUBLIC_CONTACT_DATA_ACKNOWLEDGEMENT
     */
    contactNotices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. Settings for contact information linked to the Registration.
     * Structure is documented below.
     */
    contactSettings: pulumi.Input<inputs.clouddomains.RegistrationContactSettings>;
    /**
     * Settings controlling the DNS configuration of the Registration.
     * Structure is documented below.
     */
    dnsSettings?: pulumi.Input<inputs.clouddomains.RegistrationDnsSettings>;
    /**
     * Required. The domain name. Unicode domain names must be expressed in Punycode format.
     */
    domainName: pulumi.Input<string>;
    /**
     * The list of domain notices that you acknowledge. Possible value is HSTS_PRELOADED
     */
    domainNotices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of labels associated with the Registration.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location: pulumi.Input<string>;
    /**
     * Settings for management of the Registration, including renewal, billing, and transfer
     * Structure is documented below.
     */
    managementSettings?: pulumi.Input<inputs.clouddomains.RegistrationManagementSettings>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Required. Yearly price to register or renew the domain. The value that should be put here can be obtained from
     * registrations.retrieveRegisterParameters or registrations.searchDomains calls.
     * Structure is documented below.
     */
    yearlyPrice: pulumi.Input<inputs.clouddomains.RegistrationYearlyPrice>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Google Cloud Looker instance.
 *
 * To get more information about Instance, see:
 *
 * * [API documentation](https://cloud.google.com/looker/docs/reference/rest/v1/projects.locations.instances)
 * * How-to Guides
 *     * [Create a Looker (Google Cloud core) instance](https://cloud.google.com/looker/docs/looker-core-instance-create)
 *     * [Configure a Looker (Google Cloud core) instance](https://cloud.google.com/looker/docs/looker-core-instance-setup)
 *
 * ## Example Usage
 * ### Looker Instance Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const looker_instance = new gcp.looker.Instance("looker-instance", {
 *     oauthConfig: {
 *         clientId: "my-client-id",
 *         clientSecret: "my-client-secret",
 *     },
 *     platformEdition: "LOOKER_CORE_STANDARD",
 *     region: "us-central1",
 * });
 * ```
 * ### Looker Instance Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const looker_instance = new gcp.looker.Instance("looker-instance", {
 *     adminSettings: {
 *         allowedEmailDomains: ["google.com"],
 *     },
 *     denyMaintenancePeriod: {
 *         endDate: {
 *             day: 1,
 *             month: 2,
 *             year: 2050,
 *         },
 *         startDate: {
 *             day: 1,
 *             month: 1,
 *             year: 2050,
 *         },
 *         time: {
 *             hours: 10,
 *             minutes: 0,
 *             nanos: 0,
 *             seconds: 0,
 *         },
 *     },
 *     maintenanceWindow: {
 *         dayOfWeek: "THURSDAY",
 *         startTime: {
 *             hours: 22,
 *             minutes: 0,
 *             nanos: 0,
 *             seconds: 0,
 *         },
 *     },
 *     oauthConfig: {
 *         clientId: "my-client-id",
 *         clientSecret: "my-client-secret",
 *     },
 *     platformEdition: "LOOKER_CORE_STANDARD",
 *     publicIpEnabled: true,
 *     region: "us-central1",
 *     userMetadata: {
 *         additionalDeveloperUserCount: 10,
 *         additionalStandardUserCount: 10,
 *         additionalViewerUserCount: 10,
 *     },
 * });
 * ```
 * ### Looker Instance Enterprise Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const lookerNetwork = new gcp.compute.Network("lookerNetwork", {});
 * const lookerRange = new gcp.compute.GlobalAddress("lookerRange", {
 *     purpose: "VPC_PEERING",
 *     addressType: "INTERNAL",
 *     prefixLength: 20,
 *     network: lookerNetwork.id,
 * });
 * const lookerVpcConnection = new gcp.servicenetworking.Connection("lookerVpcConnection", {
 *     network: lookerNetwork.id,
 *     service: "servicenetworking.googleapis.com",
 *     reservedPeeringRanges: [lookerRange.name],
 * });
 * const looker_instance = new gcp.looker.Instance("looker-instance", {
 *     platformEdition: "LOOKER_CORE_ENTERPRISE_ANNUAL",
 *     region: "us-central1",
 *     privateIpEnabled: true,
 *     publicIpEnabled: false,
 *     reservedRange: lookerRange.name,
 *     consumerNetwork: lookerNetwork.id,
 *     adminSettings: {
 *         allowedEmailDomains: ["google.com"],
 *     },
 *     encryptionConfig: {
 *         kmsKeyName: "looker-kms-key",
 *     },
 *     maintenanceWindow: {
 *         dayOfWeek: "THURSDAY",
 *         startTime: {
 *             hours: 22,
 *             minutes: 0,
 *             seconds: 0,
 *             nanos: 0,
 *         },
 *     },
 *     denyMaintenancePeriod: {
 *         startDate: {
 *             year: 2050,
 *             month: 1,
 *             day: 1,
 *         },
 *         endDate: {
 *             year: 2050,
 *             month: 2,
 *             day: 1,
 *         },
 *         time: {
 *             hours: 10,
 *             minutes: 0,
 *             seconds: 0,
 *             nanos: 0,
 *         },
 *     },
 *     oauthConfig: {
 *         clientId: "my-client-id",
 *         clientSecret: "my-client-secret",
 *     },
 * }, {
 *     dependsOn: [lookerVpcConnection],
 * });
 * const project = gcp.organizations.getProject({});
 * const cryptoKey = new gcp.kms.CryptoKeyIAMMember("cryptoKey", {
 *     cryptoKeyId: "looker-kms-key",
 *     role: "roles/cloudkms.cryptoKeyEncrypterDecrypter",
 *     member: project.then(project => `serviceAccount:service-${project.number}@gcp-sa-looker.iam.gserviceaccount.com`),
 * });
 * ```
 *
 * ## Import
 *
 * Instance can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/instances/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Instance can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:looker/instance:Instance default projects/{{project}}/locations/{{region}}/instances/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:looker/instance:Instance default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:looker/instance:Instance default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:looker/instance:Instance default {{name}}
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:looker/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Looker instance Admin settings.
     * Structure is documented below.
     */
    public readonly adminSettings!: pulumi.Output<outputs.looker.InstanceAdminSettings | undefined>;
    /**
     * Network name in the consumer project in the format of: projects/{project}/global/networks/{network}
     * Note that the consumer network may be in a different GCP project than the consumer
     * project that is hosting the Looker Instance.
     */
    public readonly consumerNetwork!: pulumi.Output<string | undefined>;
    /**
     * The time the instance was created in RFC3339 UTC "Zulu" format,
     * accurate to nanoseconds.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Maintenance denial period for this instance.
     * You must allow at least 14 days of maintenance availability
     * between any two deny maintenance periods.
     * Structure is documented below.
     */
    public readonly denyMaintenancePeriod!: pulumi.Output<outputs.looker.InstanceDenyMaintenancePeriod | undefined>;
    /**
     * Public Egress IP (IPv4).
     */
    public /*out*/ readonly egressPublicIp!: pulumi.Output<string>;
    /**
     * Looker instance encryption settings.
     * Structure is documented below.
     */
    public readonly encryptionConfig!: pulumi.Output<outputs.looker.InstanceEncryptionConfig>;
    /**
     * Private Ingress IP (IPv4).
     */
    public /*out*/ readonly ingressPrivateIp!: pulumi.Output<string>;
    /**
     * Public Ingress IP (IPv4).
     */
    public /*out*/ readonly ingressPublicIp!: pulumi.Output<string>;
    /**
     * Looker instance URI which can be used to access the Looker Instance UI.
     */
    public /*out*/ readonly lookerUri!: pulumi.Output<string>;
    /**
     * The Looker version that the instance is using.
     */
    public /*out*/ readonly lookerVersion!: pulumi.Output<string>;
    /**
     * Maintenance window for an instance.
     * Maintenance of your instance takes place once a month, and will require
     * your instance to be restarted during updates, which will temporarily
     * disrupt service.
     * Structure is documented below.
     */
    public readonly maintenanceWindow!: pulumi.Output<outputs.looker.InstanceMaintenanceWindow | undefined>;
    /**
     * The ID of the instance or a fully qualified identifier for the instance.
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Looker Instance OAuth login settings.
     * Structure is documented below.
     */
    public readonly oauthConfig!: pulumi.Output<outputs.looker.InstanceOauthConfig | undefined>;
    /**
     * Platform editions for a Looker instance. Each edition maps to a set of instance features, like its size. Must be one of these values:
     * - LOOKER_CORE_TRIAL: trial instance
     * - LOOKER_CORE_STANDARD: pay as you go standard instance
     * - LOOKER_CORE_STANDARD_ANNUAL: subscription standard instance
     * - LOOKER_CORE_ENTERPRISE_ANNUAL: subscription enterprise instance
     * - LOOKER_CORE_EMBED_ANNUAL: subscription embed instance
     * Default value is `LOOKER_CORE_TRIAL`.
     * Possible values are: `LOOKER_CORE_TRIAL`, `LOOKER_CORE_STANDARD`, `LOOKER_CORE_STANDARD_ANNUAL`, `LOOKER_CORE_ENTERPRISE_ANNUAL`, `LOOKER_CORE_EMBED_ANNUAL`.
     */
    public readonly platformEdition!: pulumi.Output<string | undefined>;
    /**
     * Whether private IP is enabled on the Looker instance.
     */
    public readonly privateIpEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Whether public IP is enabled on the Looker instance.
     */
    public readonly publicIpEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Looker region of the instance.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Name of a reserved IP address range within the consumer network, to be used for
     * private service access connection. User may or may not specify this in a request.
     */
    public readonly reservedRange!: pulumi.Output<string | undefined>;
    /**
     * The time the instance was updated in RFC3339 UTC "Zulu" format,
     * accurate to nanoseconds.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Metadata about users for a Looker instance.
     * These settings are only available when platform edition LOOKER_CORE_STANDARD is set.
     * There are ten Standard and two Developer users included in the cost of the product.
     * You can allocate additional Standard, Viewer, and Developer users for this instance.
     * It is an optional step and can be modified later.
     * With the Standard edition of Looker (Google Cloud core), you can provision up to 50
     * total users, distributed across Viewer, Standard, and Developer.
     * Structure is documented below.
     */
    public readonly userMetadata!: pulumi.Output<outputs.looker.InstanceUserMetadata | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["adminSettings"] = state ? state.adminSettings : undefined;
            resourceInputs["consumerNetwork"] = state ? state.consumerNetwork : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["denyMaintenancePeriod"] = state ? state.denyMaintenancePeriod : undefined;
            resourceInputs["egressPublicIp"] = state ? state.egressPublicIp : undefined;
            resourceInputs["encryptionConfig"] = state ? state.encryptionConfig : undefined;
            resourceInputs["ingressPrivateIp"] = state ? state.ingressPrivateIp : undefined;
            resourceInputs["ingressPublicIp"] = state ? state.ingressPublicIp : undefined;
            resourceInputs["lookerUri"] = state ? state.lookerUri : undefined;
            resourceInputs["lookerVersion"] = state ? state.lookerVersion : undefined;
            resourceInputs["maintenanceWindow"] = state ? state.maintenanceWindow : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["oauthConfig"] = state ? state.oauthConfig : undefined;
            resourceInputs["platformEdition"] = state ? state.platformEdition : undefined;
            resourceInputs["privateIpEnabled"] = state ? state.privateIpEnabled : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["publicIpEnabled"] = state ? state.publicIpEnabled : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["reservedRange"] = state ? state.reservedRange : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["userMetadata"] = state ? state.userMetadata : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            resourceInputs["adminSettings"] = args ? args.adminSettings : undefined;
            resourceInputs["consumerNetwork"] = args ? args.consumerNetwork : undefined;
            resourceInputs["denyMaintenancePeriod"] = args ? args.denyMaintenancePeriod : undefined;
            resourceInputs["encryptionConfig"] = args ? args.encryptionConfig : undefined;
            resourceInputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oauthConfig"] = args ? args.oauthConfig : undefined;
            resourceInputs["platformEdition"] = args ? args.platformEdition : undefined;
            resourceInputs["privateIpEnabled"] = args ? args.privateIpEnabled : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["publicIpEnabled"] = args ? args.publicIpEnabled : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["reservedRange"] = args ? args.reservedRange : undefined;
            resourceInputs["userMetadata"] = args ? args.userMetadata : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["egressPublicIp"] = undefined /*out*/;
            resourceInputs["ingressPrivateIp"] = undefined /*out*/;
            resourceInputs["ingressPublicIp"] = undefined /*out*/;
            resourceInputs["lookerUri"] = undefined /*out*/;
            resourceInputs["lookerVersion"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Looker instance Admin settings.
     * Structure is documented below.
     */
    adminSettings?: pulumi.Input<inputs.looker.InstanceAdminSettings>;
    /**
     * Network name in the consumer project in the format of: projects/{project}/global/networks/{network}
     * Note that the consumer network may be in a different GCP project than the consumer
     * project that is hosting the Looker Instance.
     */
    consumerNetwork?: pulumi.Input<string>;
    /**
     * The time the instance was created in RFC3339 UTC "Zulu" format,
     * accurate to nanoseconds.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Maintenance denial period for this instance.
     * You must allow at least 14 days of maintenance availability
     * between any two deny maintenance periods.
     * Structure is documented below.
     */
    denyMaintenancePeriod?: pulumi.Input<inputs.looker.InstanceDenyMaintenancePeriod>;
    /**
     * Public Egress IP (IPv4).
     */
    egressPublicIp?: pulumi.Input<string>;
    /**
     * Looker instance encryption settings.
     * Structure is documented below.
     */
    encryptionConfig?: pulumi.Input<inputs.looker.InstanceEncryptionConfig>;
    /**
     * Private Ingress IP (IPv4).
     */
    ingressPrivateIp?: pulumi.Input<string>;
    /**
     * Public Ingress IP (IPv4).
     */
    ingressPublicIp?: pulumi.Input<string>;
    /**
     * Looker instance URI which can be used to access the Looker Instance UI.
     */
    lookerUri?: pulumi.Input<string>;
    /**
     * The Looker version that the instance is using.
     */
    lookerVersion?: pulumi.Input<string>;
    /**
     * Maintenance window for an instance.
     * Maintenance of your instance takes place once a month, and will require
     * your instance to be restarted during updates, which will temporarily
     * disrupt service.
     * Structure is documented below.
     */
    maintenanceWindow?: pulumi.Input<inputs.looker.InstanceMaintenanceWindow>;
    /**
     * The ID of the instance or a fully qualified identifier for the instance.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * Looker Instance OAuth login settings.
     * Structure is documented below.
     */
    oauthConfig?: pulumi.Input<inputs.looker.InstanceOauthConfig>;
    /**
     * Platform editions for a Looker instance. Each edition maps to a set of instance features, like its size. Must be one of these values:
     * - LOOKER_CORE_TRIAL: trial instance
     * - LOOKER_CORE_STANDARD: pay as you go standard instance
     * - LOOKER_CORE_STANDARD_ANNUAL: subscription standard instance
     * - LOOKER_CORE_ENTERPRISE_ANNUAL: subscription enterprise instance
     * - LOOKER_CORE_EMBED_ANNUAL: subscription embed instance
     * Default value is `LOOKER_CORE_TRIAL`.
     * Possible values are: `LOOKER_CORE_TRIAL`, `LOOKER_CORE_STANDARD`, `LOOKER_CORE_STANDARD_ANNUAL`, `LOOKER_CORE_ENTERPRISE_ANNUAL`, `LOOKER_CORE_EMBED_ANNUAL`.
     */
    platformEdition?: pulumi.Input<string>;
    /**
     * Whether private IP is enabled on the Looker instance.
     */
    privateIpEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Whether public IP is enabled on the Looker instance.
     */
    publicIpEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Looker region of the instance.
     */
    region?: pulumi.Input<string>;
    /**
     * Name of a reserved IP address range within the consumer network, to be used for
     * private service access connection. User may or may not specify this in a request.
     */
    reservedRange?: pulumi.Input<string>;
    /**
     * The time the instance was updated in RFC3339 UTC "Zulu" format,
     * accurate to nanoseconds.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * Metadata about users for a Looker instance.
     * These settings are only available when platform edition LOOKER_CORE_STANDARD is set.
     * There are ten Standard and two Developer users included in the cost of the product.
     * You can allocate additional Standard, Viewer, and Developer users for this instance.
     * It is an optional step and can be modified later.
     * With the Standard edition of Looker (Google Cloud core), you can provision up to 50
     * total users, distributed across Viewer, Standard, and Developer.
     * Structure is documented below.
     */
    userMetadata?: pulumi.Input<inputs.looker.InstanceUserMetadata>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Looker instance Admin settings.
     * Structure is documented below.
     */
    adminSettings?: pulumi.Input<inputs.looker.InstanceAdminSettings>;
    /**
     * Network name in the consumer project in the format of: projects/{project}/global/networks/{network}
     * Note that the consumer network may be in a different GCP project than the consumer
     * project that is hosting the Looker Instance.
     */
    consumerNetwork?: pulumi.Input<string>;
    /**
     * Maintenance denial period for this instance.
     * You must allow at least 14 days of maintenance availability
     * between any two deny maintenance periods.
     * Structure is documented below.
     */
    denyMaintenancePeriod?: pulumi.Input<inputs.looker.InstanceDenyMaintenancePeriod>;
    /**
     * Looker instance encryption settings.
     * Structure is documented below.
     */
    encryptionConfig?: pulumi.Input<inputs.looker.InstanceEncryptionConfig>;
    /**
     * Maintenance window for an instance.
     * Maintenance of your instance takes place once a month, and will require
     * your instance to be restarted during updates, which will temporarily
     * disrupt service.
     * Structure is documented below.
     */
    maintenanceWindow?: pulumi.Input<inputs.looker.InstanceMaintenanceWindow>;
    /**
     * The ID of the instance or a fully qualified identifier for the instance.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * Looker Instance OAuth login settings.
     * Structure is documented below.
     */
    oauthConfig?: pulumi.Input<inputs.looker.InstanceOauthConfig>;
    /**
     * Platform editions for a Looker instance. Each edition maps to a set of instance features, like its size. Must be one of these values:
     * - LOOKER_CORE_TRIAL: trial instance
     * - LOOKER_CORE_STANDARD: pay as you go standard instance
     * - LOOKER_CORE_STANDARD_ANNUAL: subscription standard instance
     * - LOOKER_CORE_ENTERPRISE_ANNUAL: subscription enterprise instance
     * - LOOKER_CORE_EMBED_ANNUAL: subscription embed instance
     * Default value is `LOOKER_CORE_TRIAL`.
     * Possible values are: `LOOKER_CORE_TRIAL`, `LOOKER_CORE_STANDARD`, `LOOKER_CORE_STANDARD_ANNUAL`, `LOOKER_CORE_ENTERPRISE_ANNUAL`, `LOOKER_CORE_EMBED_ANNUAL`.
     */
    platformEdition?: pulumi.Input<string>;
    /**
     * Whether private IP is enabled on the Looker instance.
     */
    privateIpEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Whether public IP is enabled on the Looker instance.
     */
    publicIpEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Looker region of the instance.
     */
    region?: pulumi.Input<string>;
    /**
     * Name of a reserved IP address range within the consumer network, to be used for
     * private service access connection. User may or may not specify this in a request.
     */
    reservedRange?: pulumi.Input<string>;
    /**
     * Metadata about users for a Looker instance.
     * These settings are only available when platform edition LOOKER_CORE_STANDARD is set.
     * There are ten Standard and two Developer users included in the cost of the product.
     * You can allocate additional Standard, Viewer, and Developer users for this instance.
     * It is an optional step and can be modified later.
     * With the Standard edition of Looker (Google Cloud core), you can provision up to 50
     * total users, distributed across Viewer, Standard, and Developer.
     * Structure is documented below.
     */
    userMetadata?: pulumi.Input<inputs.looker.InstanceUserMetadata>;
}

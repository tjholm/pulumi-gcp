// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Backup Dr Management Server
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {}, {
 *     provider: google_beta,
 * });
 * const privateIpAddress = new gcp.compute.GlobalAddress("privateIpAddress", {
 *     addressType: "INTERNAL",
 *     purpose: "VPC_PEERING",
 *     prefixLength: 20,
 *     network: defaultNetwork.id,
 * }, {
 *     provider: google_beta,
 * });
 * const defaultConnection = new gcp.servicenetworking.Connection("defaultConnection", {
 *     network: defaultNetwork.id,
 *     service: "servicenetworking.googleapis.com",
 *     reservedPeeringRanges: [privateIpAddress.name],
 * }, {
 *     provider: google_beta,
 * });
 * const ms_console = new gcp.backupdisasterrecovery.ManagementServer("ms-console", {
 *     location: "us-central1",
 *     type: "BACKUP_RESTORE",
 *     networks: [{
 *         network: defaultNetwork.id,
 *         peeringMode: "PRIVATE_SERVICE_ACCESS",
 *     }],
 * }, {
 *     provider: google_beta,
 *     dependsOn: [defaultConnection],
 * });
 * ```
 *
 * ## Import
 *
 * ManagementServer can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:backupdisasterrecovery/managementServer:ManagementServer default projects/{{project}}/locations/{{location}}/managementServers/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:backupdisasterrecovery/managementServer:ManagementServer default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:backupdisasterrecovery/managementServer:ManagementServer default {{location}}/{{name}}
 * ```
 */
export class ManagementServer extends pulumi.CustomResource {
    /**
     * Get an existing ManagementServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagementServerState, opts?: pulumi.CustomResourceOptions): ManagementServer {
        return new ManagementServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:backupdisasterrecovery/managementServer:ManagementServer';

    /**
     * Returns true if the given object is an instance of ManagementServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagementServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagementServer.__pulumiType;
    }

    /**
     * The location for the management server (management console)
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The management console URI
     * Structure is documented below.
     */
    public /*out*/ readonly managementUris!: pulumi.Output<outputs.backupdisasterrecovery.ManagementServerManagementUri[]>;
    /**
     * The name of management server (management console)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Network details to create management server (management console).
     * Structure is documented below.
     */
    public readonly networks!: pulumi.Output<outputs.backupdisasterrecovery.ManagementServerNetwork[]>;
    /**
     * The oauth2ClientId of management console.
     */
    public /*out*/ readonly oauth2ClientId!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The type of management server (management console).
     * Default value is `BACKUP_RESTORE`.
     * Possible values are: `BACKUP_RESTORE`.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a ManagementServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagementServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagementServerArgs | ManagementServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ManagementServerState | undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["managementUris"] = state ? state.managementUris : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["oauth2ClientId"] = state ? state.oauth2ClientId : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ManagementServerArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.networks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networks'");
            }
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["managementUris"] = undefined /*out*/;
            resourceInputs["oauth2ClientId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ManagementServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagementServer resources.
 */
export interface ManagementServerState {
    /**
     * The location for the management server (management console)
     */
    location?: pulumi.Input<string>;
    /**
     * The management console URI
     * Structure is documented below.
     */
    managementUris?: pulumi.Input<pulumi.Input<inputs.backupdisasterrecovery.ManagementServerManagementUri>[]>;
    /**
     * The name of management server (management console)
     */
    name?: pulumi.Input<string>;
    /**
     * Network details to create management server (management console).
     * Structure is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.backupdisasterrecovery.ManagementServerNetwork>[]>;
    /**
     * The oauth2ClientId of management console.
     */
    oauth2ClientId?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The type of management server (management console).
     * Default value is `BACKUP_RESTORE`.
     * Possible values are: `BACKUP_RESTORE`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagementServer resource.
 */
export interface ManagementServerArgs {
    /**
     * The location for the management server (management console)
     */
    location: pulumi.Input<string>;
    /**
     * The name of management server (management console)
     */
    name?: pulumi.Input<string>;
    /**
     * Network details to create management server (management console).
     * Structure is documented below.
     */
    networks: pulumi.Input<pulumi.Input<inputs.backupdisasterrecovery.ManagementServerNetwork>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The type of management server (management console).
     * Default value is `BACKUP_RESTORE`.
     * Possible values are: `BACKUP_RESTORE`.
     */
    type?: pulumi.Input<string>;
}

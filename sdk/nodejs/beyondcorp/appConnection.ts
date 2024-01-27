// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A BeyondCorp AppConnection resource represents a BeyondCorp protected AppConnection to a remote application.
 * It creates all the necessary GCP components needed for creating a BeyondCorp protected AppConnection.
 * Multiple connectors can be authorised for a single AppConnection.
 *
 * To get more information about AppConnection, see:
 *
 * * [API documentation](https://cloud.google.com/beyondcorp/docs/reference/rest#rest-resource:-v1.projects.locations.appconnections)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/beyondcorp-enterprise/docs/enable-app-connector)
 *
 * ## Example Usage
 * ### Beyondcorp App Connection Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const serviceAccount = new gcp.serviceaccount.Account("serviceAccount", {
 *     accountId: "my-account",
 *     displayName: "Test Service Account",
 * });
 * const appConnector = new gcp.beyondcorp.AppConnector("appConnector", {principalInfo: {
 *     serviceAccount: {
 *         email: serviceAccount.email,
 *     },
 * }});
 * const appConnection = new gcp.beyondcorp.AppConnection("appConnection", {
 *     type: "TCP_PROXY",
 *     applicationEndpoint: {
 *         host: "foo-host",
 *         port: 8080,
 *     },
 *     connectors: [appConnector.id],
 * });
 * ```
 * ### Beyondcorp App Connection Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const serviceAccount = new gcp.serviceaccount.Account("serviceAccount", {
 *     accountId: "my-account",
 *     displayName: "Test Service Account",
 * });
 * const appGateway = new gcp.beyondcorp.AppGateway("appGateway", {
 *     type: "TCP_PROXY",
 *     hostType: "GCP_REGIONAL_MIG",
 * });
 * const appConnector = new gcp.beyondcorp.AppConnector("appConnector", {principalInfo: {
 *     serviceAccount: {
 *         email: serviceAccount.email,
 *     },
 * }});
 * const appConnection = new gcp.beyondcorp.AppConnection("appConnection", {
 *     type: "TCP_PROXY",
 *     displayName: "some display name",
 *     applicationEndpoint: {
 *         host: "foo-host",
 *         port: 8080,
 *     },
 *     connectors: [appConnector.id],
 *     gateway: {
 *         appGateway: appGateway.id,
 *     },
 *     labels: {
 *         foo: "bar",
 *         bar: "baz",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * AppConnection can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/appConnections/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, AppConnection can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default projects/{{project}}/locations/{{region}}/appConnections/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default {{name}}
 * ```
 */
export class AppConnection extends pulumi.CustomResource {
    /**
     * Get an existing AppConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppConnectionState, opts?: pulumi.CustomResourceOptions): AppConnection {
        return new AppConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:beyondcorp/appConnection:AppConnection';

    /**
     * Returns true if the given object is an instance of AppConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppConnection.__pulumiType;
    }

    /**
     * Address of the remote application endpoint for the BeyondCorp AppConnection.
     * Structure is documented below.
     */
    public readonly applicationEndpoint!: pulumi.Output<outputs.beyondcorp.AppConnectionApplicationEndpoint>;
    /**
     * List of AppConnectors that are authorised to be associated with this AppConnection
     */
    public readonly connectors!: pulumi.Output<string[] | undefined>;
    /**
     * An arbitrary user-provided name for the AppConnection.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Gateway used by the AppConnection.
     * Structure is documented below.
     */
    public readonly gateway!: pulumi.Output<outputs.beyondcorp.AppConnectionGateway>;
    /**
     * Resource labels to represent user provided metadata.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * ID of the AppConnection.
     */
    public readonly name!: pulumi.Output<string>;
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
     * The region of the AppConnection.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * The type of network connectivity used by the AppConnection. Refer to
     * https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
     * for a list of possible values.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a AppConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppConnectionArgs | AppConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppConnectionState | undefined;
            resourceInputs["applicationEndpoint"] = state ? state.applicationEndpoint : undefined;
            resourceInputs["connectors"] = state ? state.connectors : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AppConnectionArgs | undefined;
            if ((!args || args.applicationEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationEndpoint'");
            }
            resourceInputs["applicationEndpoint"] = args ? args.applicationEndpoint : undefined;
            resourceInputs["connectors"] = args ? args.connectors : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["gateway"] = args ? args.gateway : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AppConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppConnection resources.
 */
export interface AppConnectionState {
    /**
     * Address of the remote application endpoint for the BeyondCorp AppConnection.
     * Structure is documented below.
     */
    applicationEndpoint?: pulumi.Input<inputs.beyondcorp.AppConnectionApplicationEndpoint>;
    /**
     * List of AppConnectors that are authorised to be associated with this AppConnection
     */
    connectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An arbitrary user-provided name for the AppConnection.
     */
    displayName?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Gateway used by the AppConnection.
     * Structure is documented below.
     */
    gateway?: pulumi.Input<inputs.beyondcorp.AppConnectionGateway>;
    /**
     * Resource labels to represent user provided metadata.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * ID of the AppConnection.
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
     * The region of the AppConnection.
     */
    region?: pulumi.Input<string>;
    /**
     * The type of network connectivity used by the AppConnection. Refer to
     * https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
     * for a list of possible values.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppConnection resource.
 */
export interface AppConnectionArgs {
    /**
     * Address of the remote application endpoint for the BeyondCorp AppConnection.
     * Structure is documented below.
     */
    applicationEndpoint: pulumi.Input<inputs.beyondcorp.AppConnectionApplicationEndpoint>;
    /**
     * List of AppConnectors that are authorised to be associated with this AppConnection
     */
    connectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An arbitrary user-provided name for the AppConnection.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Gateway used by the AppConnection.
     * Structure is documented below.
     */
    gateway?: pulumi.Input<inputs.beyondcorp.AppConnectionGateway>;
    /**
     * Resource labels to represent user provided metadata.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * ID of the AppConnection.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region of the AppConnection.
     */
    region?: pulumi.Input<string>;
    /**
     * The type of network connectivity used by the AppConnection. Refer to
     * https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
     * for a list of possible values.
     */
    type?: pulumi.Input<string>;
}

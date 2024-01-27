// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Contains the data that describes an Identity Aware Proxy owned client.
 *
 * > **Note:** Only internal org clients can be created via declarative tools. External clients must be
 * manually created via the GCP console. This restriction is due to the existing APIs and not lack of support
 * in this tool.
 *
 * To get more information about Client, see:
 *
 * * [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.brands.identityAwareProxyClients)
 * * How-to Guides
 *     * [Setting up IAP Client](https://cloud.google.com/iap/docs/authentication-howto)
 *
 * ## Example Usage
 * ### Iap Client
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = new gcp.organizations.Project("project", {orgId: "123456789"});
 * const projectService = new gcp.projects.Service("projectService", {
 *     project: project.projectId,
 *     service: "iap.googleapis.com",
 * });
 * const projectBrand = new gcp.iap.Brand("projectBrand", {
 *     supportEmail: "support@example.com",
 *     applicationTitle: "Cloud IAP protected Application",
 *     project: projectService.project,
 * });
 * const projectClient = new gcp.iap.Client("projectClient", {
 *     displayName: "Test Client",
 *     brand: projectBrand.name,
 * });
 * ```
 *
 * ## Import
 *
 * Client can be imported using any of these accepted formats* `{{brand}}/identityAwareProxyClients/{{client_id}}` * `{{brand}}/{{client_id}}` When using the `pulumi import` command, Client can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:iap/client:Client default {{brand}}/identityAwareProxyClients/{{client_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:iap/client:Client default {{brand}}/{{client_id}}
 * ```
 */
export class Client extends pulumi.CustomResource {
    /**
     * Get an existing Client resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientState, opts?: pulumi.CustomResourceOptions): Client {
        return new Client(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/client:Client';

    /**
     * Returns true if the given object is an instance of Client.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Client {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Client.__pulumiType;
    }

    /**
     * Identifier of the brand to which this client
     * is attached to. The format is
     * `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
     *
     *
     * - - -
     */
    public readonly brand!: pulumi.Output<string>;
    /**
     * The OAuth2 ID of the client.
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * Human-friendly name given to the OAuth client.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Output only. Client secret of the OAuth client.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    public /*out*/ readonly secret!: pulumi.Output<string>;

    /**
     * Create a Client resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientArgs | ClientState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientState | undefined;
            resourceInputs["brand"] = state ? state.brand : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
        } else {
            const args = argsOrState as ClientArgs | undefined;
            if ((!args || args.brand === undefined) && !opts.urn) {
                throw new Error("Missing required property 'brand'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["brand"] = args ? args.brand : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["clientId"] = undefined /*out*/;
            resourceInputs["secret"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Client.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Client resources.
 */
export interface ClientState {
    /**
     * Identifier of the brand to which this client
     * is attached to. The format is
     * `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
     *
     *
     * - - -
     */
    brand?: pulumi.Input<string>;
    /**
     * The OAuth2 ID of the client.
     */
    clientId?: pulumi.Input<string>;
    /**
     * Human-friendly name given to the OAuth client.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Output only. Client secret of the OAuth client.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    secret?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Client resource.
 */
export interface ClientArgs {
    /**
     * Identifier of the brand to which this client
     * is attached to. The format is
     * `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
     *
     *
     * - - -
     */
    brand: pulumi.Input<string>;
    /**
     * Human-friendly name given to the OAuth client.
     */
    displayName: pulumi.Input<string>;
}

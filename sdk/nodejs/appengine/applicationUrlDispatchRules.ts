// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Rules to match an HTTP request and dispatch that request to a service.
 *
 * To get more information about ApplicationUrlDispatchRules, see:
 *
 * * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#UrlDispatchRule)
 *
 * ## Example Usage
 * ### App Engine Application Url Dispatch Rules Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const bucket = new gcp.storage.Bucket("bucket", {location: "US"});
 * const object = new gcp.storage.BucketObject("object", {
 *     bucket: bucket.name,
 *     source: new pulumi.asset.FileAsset("./test-fixtures/hello-world.zip"),
 * });
 * const adminV3 = new gcp.appengine.StandardAppVersion("adminV3", {
 *     versionId: "v3",
 *     service: "admin",
 *     runtime: "nodejs10",
 *     entrypoint: {
 *         shell: "node ./app.js",
 *     },
 *     deployment: {
 *         zip: {
 *             sourceUrl: pulumi.interpolate`https://storage.googleapis.com/${bucket.name}/${object.name}`,
 *         },
 *     },
 *     envVariables: {
 *         port: "8080",
 *     },
 *     deleteServiceOnDestroy: true,
 * });
 * const webService = new gcp.appengine.ApplicationUrlDispatchRules("webService", {dispatchRules: [
 *     {
 *         domain: "*",
 *         path: "/*",
 *         service: "default",
 *     },
 *     {
 *         domain: "*",
 *         path: "/admin/*",
 *         service: adminV3.service,
 *     },
 * ]});
 * ```
 *
 * ## Import
 *
 * ApplicationUrlDispatchRules can be imported using any of these accepted formats* `{{project}}` When using the `pulumi import` command, ApplicationUrlDispatchRules can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules default {{project}}
 * ```
 */
export class ApplicationUrlDispatchRules extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationUrlDispatchRules resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationUrlDispatchRulesState, opts?: pulumi.CustomResourceOptions): ApplicationUrlDispatchRules {
        return new ApplicationUrlDispatchRules(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules';

    /**
     * Returns true if the given object is an instance of ApplicationUrlDispatchRules.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationUrlDispatchRules {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationUrlDispatchRules.__pulumiType;
    }

    /**
     * Rules to match an HTTP request and dispatch that request to a service.
     * Structure is documented below.
     */
    public readonly dispatchRules!: pulumi.Output<outputs.appengine.ApplicationUrlDispatchRulesDispatchRule[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a ApplicationUrlDispatchRules resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationUrlDispatchRulesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationUrlDispatchRulesArgs | ApplicationUrlDispatchRulesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationUrlDispatchRulesState | undefined;
            resourceInputs["dispatchRules"] = state ? state.dispatchRules : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ApplicationUrlDispatchRulesArgs | undefined;
            if ((!args || args.dispatchRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dispatchRules'");
            }
            resourceInputs["dispatchRules"] = args ? args.dispatchRules : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationUrlDispatchRules.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationUrlDispatchRules resources.
 */
export interface ApplicationUrlDispatchRulesState {
    /**
     * Rules to match an HTTP request and dispatch that request to a service.
     * Structure is documented below.
     */
    dispatchRules?: pulumi.Input<pulumi.Input<inputs.appengine.ApplicationUrlDispatchRulesDispatchRule>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationUrlDispatchRules resource.
 */
export interface ApplicationUrlDispatchRulesArgs {
    /**
     * Rules to match an HTTP request and dispatch that request to a service.
     * Structure is documented below.
     */
    dispatchRules: pulumi.Input<pulumi.Input<inputs.appengine.ApplicationUrlDispatchRulesDispatchRule>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

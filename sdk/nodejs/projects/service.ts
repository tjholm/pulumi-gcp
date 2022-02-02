// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows management of a single API service for a Google Cloud Platform project.
 *
 * For a list of services available, visit the [API library page](https://console.cloud.google.com/apis/library)
 * or run `gcloud services list --available`.
 *
 * This resource requires the [Service Usage API](https://console.cloud.google.com/apis/library/serviceusage.googleapis.com)
 * to use.
 *
 * To get more information about `gcp.projects.Service`, see:
 *
 * * [API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1/services)
 * * How-to Guides
 *     * [Enabling and Disabling Services](https://cloud.google.com/service-usage/docs/enable-disable)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = new gcp.projects.Service("project", {
 *     disableDependentServices: true,
 *     project: "your-project-id",
 *     service: "iam.googleapis.com",
 * }, { timeouts: {
 *     create: "30m",
 *     update: "40m",
 * } });
 * ```
 *
 * ## Import
 *
 * Project services can be imported using the `project_id` and `service`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:projects/service:Service my_project your-project-id/iam.googleapis.com
 * ```
 *
 *  Note that unlike other resources that fail if they already exist, `terraform apply` can be successfully used to verify already enabled services. This means that when importing existing resources into Terraform, you can either import the `google_project_service` resources or treat them as new infrastructure and run `terraform apply` to add them to state.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:projects/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * If `true`, services that are enabled
     * and which depend on this service should also be disabled when this service is
     * destroyed. If `false` or unset, an error will be generated if any enabled
     * services depend on this service when destroying it.
     */
    public readonly disableDependentServices!: pulumi.Output<boolean | undefined>;
    /**
     * If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
     */
    public readonly disableOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The project ID. If not provided, the provider project
     * is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The service to enable.
     */
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["disableDependentServices"] = state ? state.disableDependentServices : undefined;
            resourceInputs["disableOnDestroy"] = state ? state.disableOnDestroy : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            resourceInputs["disableDependentServices"] = args ? args.disableDependentServices : undefined;
            resourceInputs["disableOnDestroy"] = args ? args.disableOnDestroy : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * If `true`, services that are enabled
     * and which depend on this service should also be disabled when this service is
     * destroyed. If `false` or unset, an error will be generated if any enabled
     * services depend on this service when destroying it.
     */
    disableDependentServices?: pulumi.Input<boolean>;
    /**
     * If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
     */
    disableOnDestroy?: pulumi.Input<boolean>;
    /**
     * The project ID. If not provided, the provider project
     * is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The service to enable.
     */
    service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * If `true`, services that are enabled
     * and which depend on this service should also be disabled when this service is
     * destroyed. If `false` or unset, an error will be generated if any enabled
     * services depend on this service when destroying it.
     */
    disableDependentServices?: pulumi.Input<boolean>;
    /**
     * If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
     */
    disableOnDestroy?: pulumi.Input<boolean>;
    /**
     * The project ID. If not provided, the provider project
     * is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The service to enable.
     */
    service: pulumi.Input<string>;
}

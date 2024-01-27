// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Network Services Service Binding Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNamespace = new gcp.servicedirectory.Namespace("defaultNamespace", {
 *     namespaceId: "my-namespace",
 *     location: "us-central1",
 * }, {
 *     provider: google_beta,
 * });
 * const defaultService = new gcp.servicedirectory.Service("defaultService", {
 *     serviceId: "my-service",
 *     namespace: defaultNamespace.id,
 *     metadata: {
 *         stage: "prod",
 *         region: "us-central1",
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const defaultServiceBinding = new gcp.networkservices.ServiceBinding("defaultServiceBinding", {
 *     labels: {
 *         foo: "bar",
 *     },
 *     description: "my description",
 *     service: defaultService.id,
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * ServiceBinding can be imported using any of these accepted formats* `projects/{{project}}/locations/global/serviceBindings/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, ServiceBinding can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/serviceBinding:ServiceBinding default projects/{{project}}/locations/global/serviceBindings/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/serviceBinding:ServiceBinding default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/serviceBinding:ServiceBinding default {{name}}
 * ```
 */
export class ServiceBinding extends pulumi.CustomResource {
    /**
     * Get an existing ServiceBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceBindingState, opts?: pulumi.CustomResourceOptions): ServiceBinding {
        return new ServiceBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:networkservices/serviceBinding:ServiceBinding';

    /**
     * Returns true if the given object is an instance of ServiceBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceBinding.__pulumiType;
    }

    /**
     * Time the ServiceBinding was created in UTC.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A free-text description of the resource. Max length 1024 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Set of label tags associated with the ServiceBinding resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the ServiceBinding resource.
     *
     *
     * - - -
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
     * The full Service Directory Service name of the format
     * projects/*&#47;locations/*&#47;namespaces/*&#47;services/*
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * Time the ServiceBinding was updated in UTC.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ServiceBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceBindingArgs | ServiceBindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceBindingState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as ServiceBindingArgs | undefined;
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceBinding resources.
 */
export interface ServiceBindingState {
    /**
     * Time the ServiceBinding was created in UTC.
     */
    createTime?: pulumi.Input<string>;
    /**
     * A free-text description of the resource. Max length 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Set of label tags associated with the ServiceBinding resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the ServiceBinding resource.
     *
     *
     * - - -
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
     * The full Service Directory Service name of the format
     * projects/*&#47;locations/*&#47;namespaces/*&#47;services/*
     */
    service?: pulumi.Input<string>;
    /**
     * Time the ServiceBinding was updated in UTC.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceBinding resource.
 */
export interface ServiceBindingArgs {
    /**
     * A free-text description of the resource. Max length 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Set of label tags associated with the ServiceBinding resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the ServiceBinding resource.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The full Service Directory Service name of the format
     * projects/*&#47;locations/*&#47;namespaces/*&#47;services/*
     */
    service: pulumi.Input<string>;
}

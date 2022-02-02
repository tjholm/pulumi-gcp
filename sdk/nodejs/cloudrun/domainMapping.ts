// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Resource to hold the state and status of a user's domain mapping.
 *
 * To get more information about DomainMapping, see:
 *
 * * [API documentation](https://cloud.google.com/run/docs/reference/rest/v1/projects.locations.domainmappings)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/run/docs/mapping-custom-domains)
 *
 * ## Example Usage
 * ### Cloud Run Domain Mapping Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultService = new gcp.cloudrun.Service("defaultService", {
 *     location: "us-central1",
 *     metadata: {
 *         namespace: "my-project-name",
 *     },
 *     template: {
 *         spec: {
 *             containers: [{
 *                 image: "us-docker.pkg.dev/cloudrun/container/hello",
 *             }],
 *         },
 *     },
 * });
 * const defaultDomainMapping = new gcp.cloudrun.DomainMapping("defaultDomainMapping", {
 *     location: "us-central1",
 *     metadata: {
 *         namespace: "my-project-name",
 *     },
 *     spec: {
 *         routeName: defaultService.name,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * DomainMapping can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:cloudrun/domainMapping:DomainMapping default locations/{{location}}/namespaces/{{project}}/domainmappings/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudrun/domainMapping:DomainMapping default {{location}}/{{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudrun/domainMapping:DomainMapping default {{location}}/{{name}}
 * ```
 */
export class DomainMapping extends pulumi.CustomResource {
    /**
     * Get an existing DomainMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainMappingState, opts?: pulumi.CustomResourceOptions): DomainMapping {
        return new DomainMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudrun/domainMapping:DomainMapping';

    /**
     * Returns true if the given object is an instance of DomainMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainMapping.__pulumiType;
    }

    /**
     * The location of the cloud run instance. eg us-central1
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Metadata associated with this DomainMapping.
     * Structure is documented below.
     */
    public readonly metadata!: pulumi.Output<outputs.cloudrun.DomainMappingMetadata>;
    /**
     * Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The spec for this DomainMapping.
     * Structure is documented below.
     */
    public readonly spec!: pulumi.Output<outputs.cloudrun.DomainMappingSpec>;
    /**
     * The current status of the DomainMapping.
     */
    public /*out*/ readonly statuses!: pulumi.Output<outputs.cloudrun.DomainMappingStatus[]>;

    /**
     * Create a DomainMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainMappingArgs | DomainMappingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainMappingState | undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["spec"] = state ? state.spec : undefined;
            resourceInputs["statuses"] = state ? state.statuses : undefined;
        } else {
            const args = argsOrState as DomainMappingArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.metadata === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metadata'");
            }
            if ((!args || args.spec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spec'");
            }
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["statuses"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainMapping resources.
 */
export interface DomainMappingState {
    /**
     * The location of the cloud run instance. eg us-central1
     */
    location?: pulumi.Input<string>;
    /**
     * Metadata associated with this DomainMapping.
     * Structure is documented below.
     */
    metadata?: pulumi.Input<inputs.cloudrun.DomainMappingMetadata>;
    /**
     * Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The spec for this DomainMapping.
     * Structure is documented below.
     */
    spec?: pulumi.Input<inputs.cloudrun.DomainMappingSpec>;
    /**
     * The current status of the DomainMapping.
     */
    statuses?: pulumi.Input<pulumi.Input<inputs.cloudrun.DomainMappingStatus>[]>;
}

/**
 * The set of arguments for constructing a DomainMapping resource.
 */
export interface DomainMappingArgs {
    /**
     * The location of the cloud run instance. eg us-central1
     */
    location: pulumi.Input<string>;
    /**
     * Metadata associated with this DomainMapping.
     * Structure is documented below.
     */
    metadata: pulumi.Input<inputs.cloudrun.DomainMappingMetadata>;
    /**
     * Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The spec for this DomainMapping.
     * Structure is documented below.
     */
    spec: pulumi.Input<inputs.cloudrun.DomainMappingSpec>;
}

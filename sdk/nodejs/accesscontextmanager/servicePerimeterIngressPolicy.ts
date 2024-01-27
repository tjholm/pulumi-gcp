// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * IngressPolicies match requests based on ingressFrom and ingressTo stanzas. For an ingress policy to match,
 * both the ingressFrom and ingressTo stanzas must be matched. If an IngressPolicy matches a request,
 * the request is allowed through the perimeter boundary from outside the perimeter.
 * For example, access from the internet can be allowed either based on an AccessLevel or,
 * for traffic hosted on Google Cloud, the project of the source network.
 * For access from private networks, using the project of the hosting network is required.
 * Individual ingress policies can be limited by restricting which services and/
 * or actions they match using the ingressTo field.
 *
 * To get more information about ServicePerimeterIngressPolicy, see:
 *
 * * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters#ingresspolicy)
 *
 * ## Import
 *
 * ServicePerimeterIngressPolicy can be imported using any of these accepted formats* `{{perimeter}}` When using the `pulumi import` command, ServicePerimeterIngressPolicy can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:accesscontextmanager/servicePerimeterIngressPolicy:ServicePerimeterIngressPolicy default {{perimeter}}
 * ```
 */
export class ServicePerimeterIngressPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ServicePerimeterIngressPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicePerimeterIngressPolicyState, opts?: pulumi.CustomResourceOptions): ServicePerimeterIngressPolicy {
        return new ServicePerimeterIngressPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:accesscontextmanager/servicePerimeterIngressPolicy:ServicePerimeterIngressPolicy';

    /**
     * Returns true if the given object is an instance of ServicePerimeterIngressPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePerimeterIngressPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePerimeterIngressPolicy.__pulumiType;
    }

    /**
     * Defines the conditions on the source of a request causing this `IngressPolicy`
     * to apply.
     * Structure is documented below.
     */
    public readonly ingressFrom!: pulumi.Output<outputs.accesscontextmanager.ServicePerimeterIngressPolicyIngressFrom | undefined>;
    /**
     * Defines the conditions on the `ApiOperation` and request destination that cause
     * this `IngressPolicy` to apply.
     * Structure is documented below.
     */
    public readonly ingressTo!: pulumi.Output<outputs.accesscontextmanager.ServicePerimeterIngressPolicyIngressTo | undefined>;
    /**
     * The name of the Service Perimeter to add this resource to.
     *
     *
     * - - -
     */
    public readonly perimeter!: pulumi.Output<string>;

    /**
     * Create a ServicePerimeterIngressPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePerimeterIngressPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicePerimeterIngressPolicyArgs | ServicePerimeterIngressPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServicePerimeterIngressPolicyState | undefined;
            resourceInputs["ingressFrom"] = state ? state.ingressFrom : undefined;
            resourceInputs["ingressTo"] = state ? state.ingressTo : undefined;
            resourceInputs["perimeter"] = state ? state.perimeter : undefined;
        } else {
            const args = argsOrState as ServicePerimeterIngressPolicyArgs | undefined;
            if ((!args || args.perimeter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'perimeter'");
            }
            resourceInputs["ingressFrom"] = args ? args.ingressFrom : undefined;
            resourceInputs["ingressTo"] = args ? args.ingressTo : undefined;
            resourceInputs["perimeter"] = args ? args.perimeter : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServicePerimeterIngressPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePerimeterIngressPolicy resources.
 */
export interface ServicePerimeterIngressPolicyState {
    /**
     * Defines the conditions on the source of a request causing this `IngressPolicy`
     * to apply.
     * Structure is documented below.
     */
    ingressFrom?: pulumi.Input<inputs.accesscontextmanager.ServicePerimeterIngressPolicyIngressFrom>;
    /**
     * Defines the conditions on the `ApiOperation` and request destination that cause
     * this `IngressPolicy` to apply.
     * Structure is documented below.
     */
    ingressTo?: pulumi.Input<inputs.accesscontextmanager.ServicePerimeterIngressPolicyIngressTo>;
    /**
     * The name of the Service Perimeter to add this resource to.
     *
     *
     * - - -
     */
    perimeter?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServicePerimeterIngressPolicy resource.
 */
export interface ServicePerimeterIngressPolicyArgs {
    /**
     * Defines the conditions on the source of a request causing this `IngressPolicy`
     * to apply.
     * Structure is documented below.
     */
    ingressFrom?: pulumi.Input<inputs.accesscontextmanager.ServicePerimeterIngressPolicyIngressFrom>;
    /**
     * Defines the conditions on the `ApiOperation` and request destination that cause
     * this `IngressPolicy` to apply.
     * Structure is documented below.
     */
    ingressTo?: pulumi.Input<inputs.accesscontextmanager.ServicePerimeterIngressPolicyIngressTo>;
    /**
     * The name of the Service Perimeter to add this resource to.
     *
     *
     * - - -
     */
    perimeter: pulumi.Input<string>;
}

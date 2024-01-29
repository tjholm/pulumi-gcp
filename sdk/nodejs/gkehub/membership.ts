// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Membership contains information about a member cluster.
 *
 * To get more information about Membership, see:
 *
 * * [API documentation](https://cloud.google.com/anthos/multicluster-management/reference/rest/v1/projects.locations.memberships)
 * * How-to Guides
 *     * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
 *
 * ## Example Usage
 * ### Gkehub Membership Regional
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const primary = new gcp.container.Cluster("primary", {
 *     deletionProtection: false,
 *     initialNodeCount: 1,
 *     location: "us-central1-a",
 *     network: "default",
 *     subnetwork: "default",
 * });
 * const membership = new gcp.gkehub.Membership("membership", {
 *     endpoint: {
 *         gkeCluster: {
 *             resourceLink: pulumi.interpolate`//container.googleapis.com/${primary.id}`,
 *         },
 *     },
 *     location: "us-west1",
 *     membershipId: "basic",
 * });
 * ```
 * ### Gkehub Membership Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const primary = new gcp.container.Cluster("primary", {
 *     deletionProtection: true,
 *     initialNodeCount: 1,
 *     location: "us-central1-a",
 *     network: "default",
 *     subnetwork: "default",
 * });
 * const membership = new gcp.gkehub.Membership("membership", {
 *     endpoint: {
 *         gkeCluster: {
 *             resourceLink: pulumi.interpolate`//container.googleapis.com/${primary.id}`,
 *         },
 *     },
 *     labels: {
 *         env: "test",
 *     },
 *     membershipId: "basic",
 * });
 * ```
 * ### Gkehub Membership Issuer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const primary = new gcp.container.Cluster("primary", {
 *     location: "us-central1-a",
 *     initialNodeCount: 1,
 *     workloadIdentityConfig: {
 *         workloadPool: "my-project-name.svc.id.goog",
 *     },
 *     deletionProtection: true,
 *     network: "default",
 *     subnetwork: "default",
 * });
 * const membership = new gcp.gkehub.Membership("membership", {
 *     membershipId: "basic",
 *     endpoint: {
 *         gkeCluster: {
 *             resourceLink: primary.id,
 *         },
 *     },
 *     authority: {
 *         issuer: pulumi.interpolate`https://container.googleapis.com/v1/${primary.id}`,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Membership can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}` * `{{project}}/{{location}}/{{membership_id}}` * `{{location}}/{{membership_id}}` When using the `pulumi import` command, Membership can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:gkehub/membership:Membership default projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:gkehub/membership:Membership default {{project}}/{{location}}/{{membership_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:gkehub/membership:Membership default {{location}}/{{membership_id}}
 * ```
 */
export class Membership extends pulumi.CustomResource {
    /**
     * Get an existing Membership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MembershipState, opts?: pulumi.CustomResourceOptions): Membership {
        return new Membership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:gkehub/membership:Membership';

    /**
     * Returns true if the given object is an instance of Membership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Membership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Membership.__pulumiType;
    }

    /**
     * Authority encodes how Google will recognize identities from this Membership.
     * See the workload identity documentation for more details:
     * https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
     * Structure is documented below.
     */
    public readonly authority!: pulumi.Output<outputs.gkehub.MembershipAuthority | undefined>;
    /**
     * The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
     *
     * > **Warning:** `description` is deprecated and will be removed in a future major release.
     *
     * @deprecated `description` is deprecated and will be removed in a future major release.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
     * Structure is documented below.
     */
    public readonly endpoint!: pulumi.Output<outputs.gkehub.MembershipEndpoint | undefined>;
    /**
     * Labels to apply to this membership.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Location of the membership.
     * The default value is `global`.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The client-provided identifier of the membership.
     *
     *
     * - - -
     */
    public readonly membershipId!: pulumi.Output<string>;
    /**
     * The unique identifier of the membership.
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
     * Create a Membership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MembershipArgs | MembershipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MembershipState | undefined;
            resourceInputs["authority"] = state ? state.authority : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["membershipId"] = state ? state.membershipId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
        } else {
            const args = argsOrState as MembershipArgs | undefined;
            if ((!args || args.membershipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'membershipId'");
            }
            resourceInputs["authority"] = args ? args.authority : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["membershipId"] = args ? args.membershipId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Membership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Membership resources.
 */
export interface MembershipState {
    /**
     * Authority encodes how Google will recognize identities from this Membership.
     * See the workload identity documentation for more details:
     * https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
     * Structure is documented below.
     */
    authority?: pulumi.Input<inputs.gkehub.MembershipAuthority>;
    /**
     * The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
     *
     * > **Warning:** `description` is deprecated and will be removed in a future major release.
     *
     * @deprecated `description` is deprecated and will be removed in a future major release.
     */
    description?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
     * Structure is documented below.
     */
    endpoint?: pulumi.Input<inputs.gkehub.MembershipEndpoint>;
    /**
     * Labels to apply to this membership.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Location of the membership.
     * The default value is `global`.
     */
    location?: pulumi.Input<string>;
    /**
     * The client-provided identifier of the membership.
     *
     *
     * - - -
     */
    membershipId?: pulumi.Input<string>;
    /**
     * The unique identifier of the membership.
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
}

/**
 * The set of arguments for constructing a Membership resource.
 */
export interface MembershipArgs {
    /**
     * Authority encodes how Google will recognize identities from this Membership.
     * See the workload identity documentation for more details:
     * https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
     * Structure is documented below.
     */
    authority?: pulumi.Input<inputs.gkehub.MembershipAuthority>;
    /**
     * The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
     *
     * > **Warning:** `description` is deprecated and will be removed in a future major release.
     *
     * @deprecated `description` is deprecated and will be removed in a future major release.
     */
    description?: pulumi.Input<string>;
    /**
     * If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
     * Structure is documented below.
     */
    endpoint?: pulumi.Input<inputs.gkehub.MembershipEndpoint>;
    /**
     * Labels to apply to this membership.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Location of the membership.
     * The default value is `global`.
     */
    location?: pulumi.Input<string>;
    /**
     * The client-provided identifier of the membership.
     *
     *
     * - - -
     */
    membershipId: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

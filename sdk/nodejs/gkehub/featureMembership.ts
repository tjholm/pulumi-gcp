// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Contains information about a GKEHub Feature Memberships. Feature Memberships configure GKEHub Features that apply to specific memberships rather than the project as a whole. The googleGkeHub is the Fleet API.
 *
 * ## Example Usage
 * ### Config Management
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cluster = new gcp.container.Cluster("cluster", {
 *     location: "us-central1-a",
 *     initialNodeCount: 1,
 * });
 * const membership = new gcp.gkehub.Membership("membership", {
 *     membershipId: "my-membership",
 *     endpoint: {
 *         gkeCluster: {
 *             resourceLink: pulumi.interpolate`//container.googleapis.com/${cluster.id}`,
 *         },
 *     },
 * });
 * const feature = new gcp.gkehub.Feature("feature", {
 *     location: "global",
 *     labels: {
 *         foo: "bar",
 *     },
 * });
 * const featureMember = new gcp.gkehub.FeatureMembership("featureMember", {
 *     location: "global",
 *     feature: feature.name,
 *     membership: membership.membershipId,
 *     configmanagement: {
 *         version: "1.6.2",
 *         configSync: {
 *             git: {
 *                 syncRepo: "https://github.com/hashicorp/terraform",
 *             },
 *         },
 *     },
 * });
 * ```
 * ### Config Management With OCI
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cluster = new gcp.container.Cluster("cluster", {
 *     location: "us-central1-a",
 *     initialNodeCount: 1,
 * });
 * const membership = new gcp.gkehub.Membership("membership", {
 *     membershipId: "my-membership",
 *     endpoint: {
 *         gkeCluster: {
 *             resourceLink: pulumi.interpolate`//container.googleapis.com/${cluster.id}`,
 *         },
 *     },
 * });
 * const feature = new gcp.gkehub.Feature("feature", {
 *     location: "global",
 *     labels: {
 *         foo: "bar",
 *     },
 * });
 * const featureMember = new gcp.gkehub.FeatureMembership("featureMember", {
 *     location: "global",
 *     feature: feature.name,
 *     membership: membership.membershipId,
 *     configmanagement: {
 *         version: "1.15.1",
 *         configSync: {
 *             oci: {
 *                 syncRepo: "us-central1-docker.pkg.dev/sample-project/config-repo/config-sync-gke:latest",
 *                 policyDir: "config-connector",
 *                 syncWaitSecs: "20",
 *                 secretType: "gcpserviceaccount",
 *                 gcpServiceAccountEmail: "sa@project-id.iam.gserviceaccount.com",
 *             },
 *         },
 *     },
 * });
 * ```
 * ### Multi Cluster Service Discovery
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const feature = new gcp.gkehub.Feature("feature", {
 *     labels: {
 *         foo: "bar",
 *     },
 *     location: "global",
 * });
 * ```
 * ### Service Mesh
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cluster = new gcp.container.Cluster("cluster", {
 *     location: "us-central1-a",
 *     initialNodeCount: 1,
 * });
 * const membership = new gcp.gkehub.Membership("membership", {
 *     membershipId: "my-membership",
 *     endpoint: {
 *         gkeCluster: {
 *             resourceLink: pulumi.interpolate`//container.googleapis.com/${cluster.id}`,
 *         },
 *     },
 * });
 * const feature = new gcp.gkehub.Feature("feature", {location: "global"});
 * const featureMember = new gcp.gkehub.FeatureMembership("featureMember", {
 *     location: "global",
 *     feature: feature.name,
 *     membership: membership.membershipId,
 *     mesh: {
 *         management: "MANAGEMENT_AUTOMATIC",
 *     },
 * });
 * ```
 * ### Config Management With Regional Membership
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cluster = new gcp.container.Cluster("cluster", {
 *     location: "us-central1-a",
 *     initialNodeCount: 1,
 * });
 * const membership = new gcp.gkehub.Membership("membership", {
 *     membershipId: "my-membership",
 *     location: "us-central1",
 *     endpoint: {
 *         gkeCluster: {
 *             resourceLink: pulumi.interpolate`//container.googleapis.com/${cluster.id}`,
 *         },
 *     },
 * });
 * const feature = new gcp.gkehub.Feature("feature", {
 *     location: "global",
 *     labels: {
 *         foo: "bar",
 *     },
 * });
 * const featureMember = new gcp.gkehub.FeatureMembership("featureMember", {
 *     location: "global",
 *     feature: feature.name,
 *     membership: membership.membershipId,
 *     membershipLocation: membership.location,
 *     configmanagement: {
 *         version: "1.6.2",
 *         configSync: {
 *             git: {
 *                 syncRepo: "https://github.com/hashicorp/terraform",
 *             },
 *         },
 *     },
 * });
 * ```
 * ### Policy Controller With Minimal Configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cluster = new gcp.container.Cluster("cluster", {
 *     location: "us-central1-a",
 *     initialNodeCount: 1,
 * });
 * const membership = new gcp.gkehub.Membership("membership", {
 *     membershipId: "my-membership",
 *     endpoint: {
 *         gkeCluster: {
 *             resourceLink: pulumi.interpolate`//container.googleapis.com/${cluster.id}`,
 *         },
 *     },
 * });
 * const feature = new gcp.gkehub.Feature("feature", {location: "global"});
 * const featureMember = new gcp.gkehub.FeatureMembership("featureMember", {
 *     location: "global",
 *     feature: feature.name,
 *     membership: membership.membershipId,
 *     policycontroller: {
 *         policyControllerHubConfig: {
 *             installSpec: "INSTALL_SPEC_ENABLED",
 *         },
 *     },
 * });
 * ```
 * ### Policy Controller With Custom Configurations
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cluster = new gcp.container.Cluster("cluster", {
 *     location: "us-central1-a",
 *     initialNodeCount: 1,
 * });
 * const membership = new gcp.gkehub.Membership("membership", {
 *     membershipId: "my-membership",
 *     endpoint: {
 *         gkeCluster: {
 *             resourceLink: pulumi.interpolate`//container.googleapis.com/${cluster.id}`,
 *         },
 *     },
 * });
 * const feature = new gcp.gkehub.Feature("feature", {location: "global"});
 * const featureMember = new gcp.gkehub.FeatureMembership("featureMember", {
 *     location: "global",
 *     feature: feature.name,
 *     membership: membership.membershipId,
 *     policycontroller: {
 *         policyControllerHubConfig: {
 *             installSpec: "INSTALL_SPEC_SUSPENDED",
 *             policyContent: {
 *                 templateLibrary: {
 *                     installation: "NOT_INSTALLED",
 *                 },
 *             },
 *             constraintViolationLimit: 50,
 *             auditIntervalSeconds: 120,
 *             referentialRulesEnabled: true,
 *             logDeniesEnabled: true,
 *             mutationEnabled: true,
 *         },
 *         version: "1.17.0",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * FeatureMembership can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/features/{{feature}}/membershipId/{{membership}}` * `{{project}}/{{location}}/{{feature}}/{{membership}}` * `{{location}}/{{feature}}/{{membership}}` In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import FeatureMembership using one of the formats above. For exampletf import {
 *
 *  id = "projects/{{project}}/locations/{{location}}/features/{{feature}}/membershipId/{{membership}}"
 *
 *  to = google_gke_hub_feature_membership.default }
 *
 * ```sh
 *  $ pulumi import gcp:gkehub/featureMembership:FeatureMembership When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), FeatureMembership can be imported using one of the formats above. For example
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:gkehub/featureMembership:FeatureMembership default projects/{{project}}/locations/{{location}}/features/{{feature}}/membershipId/{{membership}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:gkehub/featureMembership:FeatureMembership default {{project}}/{{location}}/{{feature}}/{{membership}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:gkehub/featureMembership:FeatureMembership default {{location}}/{{feature}}/{{membership}}
 * ```
 */
export class FeatureMembership extends pulumi.CustomResource {
    /**
     * Get an existing FeatureMembership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FeatureMembershipState, opts?: pulumi.CustomResourceOptions): FeatureMembership {
        return new FeatureMembership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:gkehub/featureMembership:FeatureMembership';

    /**
     * Returns true if the given object is an instance of FeatureMembership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FeatureMembership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FeatureMembership.__pulumiType;
    }

    /**
     * Config Management-specific spec. Structure is documented below.
     */
    public readonly configmanagement!: pulumi.Output<outputs.gkehub.FeatureMembershipConfigmanagement | undefined>;
    /**
     * The name of the feature
     */
    public readonly feature!: pulumi.Output<string>;
    /**
     * The location of the feature
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the membership
     */
    public readonly membership!: pulumi.Output<string>;
    /**
     * The location of the membership, for example, "us-central1". Default is "global".
     */
    public readonly membershipLocation!: pulumi.Output<string | undefined>;
    /**
     * Service mesh specific spec. Structure is documented below.
     */
    public readonly mesh!: pulumi.Output<outputs.gkehub.FeatureMembershipMesh | undefined>;
    /**
     * Policy Controller-specific spec. Structure is documented below.
     */
    public readonly policycontroller!: pulumi.Output<outputs.gkehub.FeatureMembershipPolicycontroller | undefined>;
    /**
     * The project of the feature
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a FeatureMembership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FeatureMembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FeatureMembershipArgs | FeatureMembershipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FeatureMembershipState | undefined;
            resourceInputs["configmanagement"] = state ? state.configmanagement : undefined;
            resourceInputs["feature"] = state ? state.feature : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["membership"] = state ? state.membership : undefined;
            resourceInputs["membershipLocation"] = state ? state.membershipLocation : undefined;
            resourceInputs["mesh"] = state ? state.mesh : undefined;
            resourceInputs["policycontroller"] = state ? state.policycontroller : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as FeatureMembershipArgs | undefined;
            if ((!args || args.feature === undefined) && !opts.urn) {
                throw new Error("Missing required property 'feature'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.membership === undefined) && !opts.urn) {
                throw new Error("Missing required property 'membership'");
            }
            resourceInputs["configmanagement"] = args ? args.configmanagement : undefined;
            resourceInputs["feature"] = args ? args.feature : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["membership"] = args ? args.membership : undefined;
            resourceInputs["membershipLocation"] = args ? args.membershipLocation : undefined;
            resourceInputs["mesh"] = args ? args.mesh : undefined;
            resourceInputs["policycontroller"] = args ? args.policycontroller : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FeatureMembership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FeatureMembership resources.
 */
export interface FeatureMembershipState {
    /**
     * Config Management-specific spec. Structure is documented below.
     */
    configmanagement?: pulumi.Input<inputs.gkehub.FeatureMembershipConfigmanagement>;
    /**
     * The name of the feature
     */
    feature?: pulumi.Input<string>;
    /**
     * The location of the feature
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the membership
     */
    membership?: pulumi.Input<string>;
    /**
     * The location of the membership, for example, "us-central1". Default is "global".
     */
    membershipLocation?: pulumi.Input<string>;
    /**
     * Service mesh specific spec. Structure is documented below.
     */
    mesh?: pulumi.Input<inputs.gkehub.FeatureMembershipMesh>;
    /**
     * Policy Controller-specific spec. Structure is documented below.
     */
    policycontroller?: pulumi.Input<inputs.gkehub.FeatureMembershipPolicycontroller>;
    /**
     * The project of the feature
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FeatureMembership resource.
 */
export interface FeatureMembershipArgs {
    /**
     * Config Management-specific spec. Structure is documented below.
     */
    configmanagement?: pulumi.Input<inputs.gkehub.FeatureMembershipConfigmanagement>;
    /**
     * The name of the feature
     */
    feature: pulumi.Input<string>;
    /**
     * The location of the feature
     */
    location: pulumi.Input<string>;
    /**
     * The name of the membership
     */
    membership: pulumi.Input<string>;
    /**
     * The location of the membership, for example, "us-central1". Default is "global".
     */
    membershipLocation?: pulumi.Input<string>;
    /**
     * Service mesh specific spec. Structure is documented below.
     */
    mesh?: pulumi.Input<inputs.gkehub.FeatureMembershipMesh>;
    /**
     * Policy Controller-specific spec. Structure is documented below.
     */
    policycontroller?: pulumi.Input<inputs.gkehub.FeatureMembershipPolicycontroller>;
    /**
     * The project of the feature
     */
    project?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The Google Compute Engine Regional Instance Group Manager API creates and manages pools
 * of homogeneous Compute Engine virtual machine instances from a common instance
 * template.
 *
 * To get more information about regionInstanceGroupManagers, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroupManagers)
 * * How-to Guides
 *     * [Regional Instance Groups Guide](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups)
 *
 * > **Note:** Use [gcp.compute.InstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_instance_group_manager.html) to create a zonal instance group manager.
 *
 * ## Example Usage
 * ### With Top Level Instance Template (`Google` Provider)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const autohealing = new gcp.compute.HealthCheck("autohealing", {
 *     checkIntervalSec: 5,
 *     timeoutSec: 5,
 *     healthyThreshold: 2,
 *     unhealthyThreshold: 10,
 *     httpHealthCheck: {
 *         requestPath: "/healthz",
 *         port: 8080,
 *     },
 * });
 * const appserver = new gcp.compute.RegionInstanceGroupManager("appserver", {
 *     baseInstanceName: "app",
 *     region: "us-central1",
 *     distributionPolicyZones: [
 *         "us-central1-a",
 *         "us-central1-f",
 *     ],
 *     versions: [{
 *         instanceTemplate: google_compute_instance_template.appserver.self_link_unique,
 *     }],
 *     allInstancesConfig: {
 *         metadata: {
 *             metadata_key: "metadata_value",
 *         },
 *         labels: {
 *             label_key: "label_value",
 *         },
 *     },
 *     targetPools: [google_compute_target_pool.appserver.id],
 *     targetSize: 2,
 *     namedPorts: [{
 *         name: "custom",
 *         port: 8888,
 *     }],
 *     autoHealingPolicies: {
 *         healthCheck: autohealing.id,
 *         initialDelaySec: 300,
 *     },
 * });
 * ```
 * ### With Multiple Versions
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const appserver = new gcp.compute.RegionInstanceGroupManager("appserver", {
 *     baseInstanceName: "app",
 *     region: "us-central1",
 *     targetSize: 5,
 *     versions: [
 *         {
 *             instanceTemplate: google_compute_instance_template.appserver.self_link_unique,
 *         },
 *         {
 *             instanceTemplate: google_compute_instance_template["appserver-canary"].self_link_unique,
 *             targetSize: {
 *                 fixed: 1,
 *             },
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Instance group managers can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager appserver appserver-igm
 * ```
 */
export class RegionInstanceGroupManager extends pulumi.CustomResource {
    /**
     * Get an existing RegionInstanceGroupManager resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionInstanceGroupManagerState, opts?: pulumi.CustomResourceOptions): RegionInstanceGroupManager {
        return new RegionInstanceGroupManager(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager';

    /**
     * Returns true if the given object is an instance of RegionInstanceGroupManager.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionInstanceGroupManager {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionInstanceGroupManager.__pulumiType;
    }

    /**
     * )
     * Properties to set on all instances in the group. After setting
     * allInstancesConfig on the group, you must update the group's instances to
     * apply the configuration.
     */
    public readonly allInstancesConfig!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerAllInstancesConfig | undefined>;
    /**
     * The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     */
    public readonly autoHealingPolicies!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerAutoHealingPolicies | undefined>;
    /**
     * The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     */
    public readonly baseInstanceName!: pulumi.Output<string>;
    /**
     * An optional textual description of the instance
     * group manager.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The shape to which the group converges either proactively or on resize events (depending on the value set in update_policy.0.instance_redistribution_type). For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/regional-mig-distribution-shape).
     */
    public readonly distributionPolicyTargetShape!: pulumi.Output<string>;
    /**
     * The distribution policy for this managed instance
     * group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
     */
    public readonly distributionPolicyZones!: pulumi.Output<string[]>;
    /**
     * The fingerprint of the instance group manager.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The full URL of the instance group created by the manager.
     */
    public /*out*/ readonly instanceGroup!: pulumi.Output<string>;
    /**
     * The instance lifecycle policy for this managed instance group.
     */
    public readonly instanceLifecyclePolicy!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerInstanceLifecyclePolicy>;
    /**
     * Pagination behavior of the `listManagedInstances` API
     * method for this managed instance group. Valid values are: `PAGELESS`, `PAGINATED`.
     * If `PAGELESS` (default), Pagination is disabled for the group's `listManagedInstances` API method.
     * `maxResults` and `pageToken` query parameters are ignored and all instances are returned in a single
     * response. If `PAGINATED`, pagination is enabled, `maxResults` and `pageToken` query parameters are
     * respected.
     */
    public readonly listManagedInstancesResults!: pulumi.Output<string | undefined>;
    /**
     * The name of the instance group manager. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    public readonly namedPorts!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerNamedPort[] | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region where the managed instance group resides. If not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URL of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs). Proactive cross zone instance redistribution must be disabled before you can update stateful disks on existing instance group managers. This can be controlled via the `updatePolicy`.
     */
    public readonly statefulDisks!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerStatefulDisk[] | undefined>;
    /**
     * ) External network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
     */
    public readonly statefulExternalIps!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerStatefulExternalIp[] | undefined>;
    /**
     * ) Internal network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
     */
    public readonly statefulInternalIps!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerStatefulInternalIp[] | undefined>;
    /**
     * The status of this managed instance group.
     */
    public /*out*/ readonly statuses!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerStatus[]>;
    /**
     * The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     */
    public readonly targetPools!: pulumi.Output<string[] | undefined>;
    /**
     * The target number of running instances for this managed
     * instance group. This value should always be explicitly set unless this resource is attached to
     * an autoscaler, in which case it should never be set. Defaults to `0`.
     */
    public readonly targetSize!: pulumi.Output<number>;
    /**
     * The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
     */
    public readonly updatePolicy!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerUpdatePolicy>;
    /**
     * Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Structure is documented below.
     */
    public readonly versions!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerVersion[]>;
    /**
     * Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, the provider will
     * continue trying until it times out.
     */
    public readonly waitForInstances!: pulumi.Output<boolean | undefined>;
    /**
     * When used with `waitForInstances` it specifies the status to wait for.
     * When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
     * set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
     * instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
     */
    public readonly waitForInstancesStatus!: pulumi.Output<string | undefined>;

    /**
     * Create a RegionInstanceGroupManager resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionInstanceGroupManagerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionInstanceGroupManagerArgs | RegionInstanceGroupManagerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegionInstanceGroupManagerState | undefined;
            resourceInputs["allInstancesConfig"] = state ? state.allInstancesConfig : undefined;
            resourceInputs["autoHealingPolicies"] = state ? state.autoHealingPolicies : undefined;
            resourceInputs["baseInstanceName"] = state ? state.baseInstanceName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["distributionPolicyTargetShape"] = state ? state.distributionPolicyTargetShape : undefined;
            resourceInputs["distributionPolicyZones"] = state ? state.distributionPolicyZones : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["instanceGroup"] = state ? state.instanceGroup : undefined;
            resourceInputs["instanceLifecyclePolicy"] = state ? state.instanceLifecyclePolicy : undefined;
            resourceInputs["listManagedInstancesResults"] = state ? state.listManagedInstancesResults : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namedPorts"] = state ? state.namedPorts : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["statefulDisks"] = state ? state.statefulDisks : undefined;
            resourceInputs["statefulExternalIps"] = state ? state.statefulExternalIps : undefined;
            resourceInputs["statefulInternalIps"] = state ? state.statefulInternalIps : undefined;
            resourceInputs["statuses"] = state ? state.statuses : undefined;
            resourceInputs["targetPools"] = state ? state.targetPools : undefined;
            resourceInputs["targetSize"] = state ? state.targetSize : undefined;
            resourceInputs["updatePolicy"] = state ? state.updatePolicy : undefined;
            resourceInputs["versions"] = state ? state.versions : undefined;
            resourceInputs["waitForInstances"] = state ? state.waitForInstances : undefined;
            resourceInputs["waitForInstancesStatus"] = state ? state.waitForInstancesStatus : undefined;
        } else {
            const args = argsOrState as RegionInstanceGroupManagerArgs | undefined;
            if ((!args || args.baseInstanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseInstanceName'");
            }
            if ((!args || args.versions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'versions'");
            }
            resourceInputs["allInstancesConfig"] = args ? args.allInstancesConfig : undefined;
            resourceInputs["autoHealingPolicies"] = args ? args.autoHealingPolicies : undefined;
            resourceInputs["baseInstanceName"] = args ? args.baseInstanceName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["distributionPolicyTargetShape"] = args ? args.distributionPolicyTargetShape : undefined;
            resourceInputs["distributionPolicyZones"] = args ? args.distributionPolicyZones : undefined;
            resourceInputs["instanceLifecyclePolicy"] = args ? args.instanceLifecyclePolicy : undefined;
            resourceInputs["listManagedInstancesResults"] = args ? args.listManagedInstancesResults : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namedPorts"] = args ? args.namedPorts : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["statefulDisks"] = args ? args.statefulDisks : undefined;
            resourceInputs["statefulExternalIps"] = args ? args.statefulExternalIps : undefined;
            resourceInputs["statefulInternalIps"] = args ? args.statefulInternalIps : undefined;
            resourceInputs["targetPools"] = args ? args.targetPools : undefined;
            resourceInputs["targetSize"] = args ? args.targetSize : undefined;
            resourceInputs["updatePolicy"] = args ? args.updatePolicy : undefined;
            resourceInputs["versions"] = args ? args.versions : undefined;
            resourceInputs["waitForInstances"] = args ? args.waitForInstances : undefined;
            resourceInputs["waitForInstancesStatus"] = args ? args.waitForInstancesStatus : undefined;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["instanceGroup"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["statuses"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegionInstanceGroupManager.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionInstanceGroupManager resources.
 */
export interface RegionInstanceGroupManagerState {
    /**
     * )
     * Properties to set on all instances in the group. After setting
     * allInstancesConfig on the group, you must update the group's instances to
     * apply the configuration.
     */
    allInstancesConfig?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerAllInstancesConfig>;
    /**
     * The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     */
    autoHealingPolicies?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerAutoHealingPolicies>;
    /**
     * The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     */
    baseInstanceName?: pulumi.Input<string>;
    /**
     * An optional textual description of the instance
     * group manager.
     */
    description?: pulumi.Input<string>;
    /**
     * The shape to which the group converges either proactively or on resize events (depending on the value set in update_policy.0.instance_redistribution_type). For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/regional-mig-distribution-shape).
     */
    distributionPolicyTargetShape?: pulumi.Input<string>;
    /**
     * The distribution policy for this managed instance
     * group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
     */
    distributionPolicyZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The fingerprint of the instance group manager.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The full URL of the instance group created by the manager.
     */
    instanceGroup?: pulumi.Input<string>;
    /**
     * The instance lifecycle policy for this managed instance group.
     */
    instanceLifecyclePolicy?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerInstanceLifecyclePolicy>;
    /**
     * Pagination behavior of the `listManagedInstances` API
     * method for this managed instance group. Valid values are: `PAGELESS`, `PAGINATED`.
     * If `PAGELESS` (default), Pagination is disabled for the group's `listManagedInstances` API method.
     * `maxResults` and `pageToken` query parameters are ignored and all instances are returned in a single
     * response. If `PAGINATED`, pagination is enabled, `maxResults` and `pageToken` query parameters are
     * respected.
     */
    listManagedInstancesResults?: pulumi.Input<string>;
    /**
     * The name of the instance group manager. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     */
    name?: pulumi.Input<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerNamedPort>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region where the managed instance group resides. If not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The URL of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs). Proactive cross zone instance redistribution must be disabled before you can update stateful disks on existing instance group managers. This can be controlled via the `updatePolicy`.
     */
    statefulDisks?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerStatefulDisk>[]>;
    /**
     * ) External network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
     */
    statefulExternalIps?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerStatefulExternalIp>[]>;
    /**
     * ) Internal network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
     */
    statefulInternalIps?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerStatefulInternalIp>[]>;
    /**
     * The status of this managed instance group.
     */
    statuses?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerStatus>[]>;
    /**
     * The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     */
    targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The target number of running instances for this managed
     * instance group. This value should always be explicitly set unless this resource is attached to
     * an autoscaler, in which case it should never be set. Defaults to `0`.
     */
    targetSize?: pulumi.Input<number>;
    /**
     * The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
     */
    updatePolicy?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerUpdatePolicy>;
    /**
     * Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Structure is documented below.
     */
    versions?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerVersion>[]>;
    /**
     * Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, the provider will
     * continue trying until it times out.
     */
    waitForInstances?: pulumi.Input<boolean>;
    /**
     * When used with `waitForInstances` it specifies the status to wait for.
     * When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
     * set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
     * instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
     */
    waitForInstancesStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RegionInstanceGroupManager resource.
 */
export interface RegionInstanceGroupManagerArgs {
    /**
     * )
     * Properties to set on all instances in the group. After setting
     * allInstancesConfig on the group, you must update the group's instances to
     * apply the configuration.
     */
    allInstancesConfig?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerAllInstancesConfig>;
    /**
     * The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     */
    autoHealingPolicies?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerAutoHealingPolicies>;
    /**
     * The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     */
    baseInstanceName: pulumi.Input<string>;
    /**
     * An optional textual description of the instance
     * group manager.
     */
    description?: pulumi.Input<string>;
    /**
     * The shape to which the group converges either proactively or on resize events (depending on the value set in update_policy.0.instance_redistribution_type). For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/regional-mig-distribution-shape).
     */
    distributionPolicyTargetShape?: pulumi.Input<string>;
    /**
     * The distribution policy for this managed instance
     * group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
     */
    distributionPolicyZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instance lifecycle policy for this managed instance group.
     */
    instanceLifecyclePolicy?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerInstanceLifecyclePolicy>;
    /**
     * Pagination behavior of the `listManagedInstances` API
     * method for this managed instance group. Valid values are: `PAGELESS`, `PAGINATED`.
     * If `PAGELESS` (default), Pagination is disabled for the group's `listManagedInstances` API method.
     * `maxResults` and `pageToken` query parameters are ignored and all instances are returned in a single
     * response. If `PAGINATED`, pagination is enabled, `maxResults` and `pageToken` query parameters are
     * respected.
     */
    listManagedInstancesResults?: pulumi.Input<string>;
    /**
     * The name of the instance group manager. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     */
    name?: pulumi.Input<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerNamedPort>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region where the managed instance group resides. If not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs). Proactive cross zone instance redistribution must be disabled before you can update stateful disks on existing instance group managers. This can be controlled via the `updatePolicy`.
     */
    statefulDisks?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerStatefulDisk>[]>;
    /**
     * ) External network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
     */
    statefulExternalIps?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerStatefulExternalIp>[]>;
    /**
     * ) Internal network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
     */
    statefulInternalIps?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerStatefulInternalIp>[]>;
    /**
     * The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     */
    targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The target number of running instances for this managed
     * instance group. This value should always be explicitly set unless this resource is attached to
     * an autoscaler, in which case it should never be set. Defaults to `0`.
     */
    targetSize?: pulumi.Input<number>;
    /**
     * The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
     */
    updatePolicy?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerUpdatePolicy>;
    /**
     * Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Structure is documented below.
     */
    versions: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerVersion>[]>;
    /**
     * Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, the provider will
     * continue trying until it times out.
     */
    waitForInstances?: pulumi.Input<boolean>;
    /**
     * When used with `waitForInstances` it specifies the status to wait for.
     * When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
     * set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
     * instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
     */
    waitForInstancesStatus?: pulumi.Input<string>;
}

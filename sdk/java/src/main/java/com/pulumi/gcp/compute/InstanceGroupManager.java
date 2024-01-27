// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.InstanceGroupManagerArgs;
import com.pulumi.gcp.compute.inputs.InstanceGroupManagerState;
import com.pulumi.gcp.compute.outputs.InstanceGroupManagerAllInstancesConfig;
import com.pulumi.gcp.compute.outputs.InstanceGroupManagerAutoHealingPolicies;
import com.pulumi.gcp.compute.outputs.InstanceGroupManagerInstanceLifecyclePolicy;
import com.pulumi.gcp.compute.outputs.InstanceGroupManagerNamedPort;
import com.pulumi.gcp.compute.outputs.InstanceGroupManagerStatefulDisk;
import com.pulumi.gcp.compute.outputs.InstanceGroupManagerStatefulExternalIp;
import com.pulumi.gcp.compute.outputs.InstanceGroupManagerStatefulInternalIp;
import com.pulumi.gcp.compute.outputs.InstanceGroupManagerStatus;
import com.pulumi.gcp.compute.outputs.InstanceGroupManagerUpdatePolicy;
import com.pulumi.gcp.compute.outputs.InstanceGroupManagerVersion;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The Google Compute Engine Instance Group Manager API creates and manages pools
 * of homogeneous Compute Engine virtual machine instances from a common instance
 * template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/manager)
 * and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroupManagers)
 * 
 * &gt; **Note:** Use [gcp.compute.RegionInstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager.html) to create a regional (multi-zone) instance group manager.
 * 
 * ## Example Usage
 * ### With Top Level Instance Template (`Google` Provider)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.HealthCheck;
 * import com.pulumi.gcp.compute.HealthCheckArgs;
 * import com.pulumi.gcp.compute.inputs.HealthCheckHttpHealthCheckArgs;
 * import com.pulumi.gcp.compute.InstanceGroupManager;
 * import com.pulumi.gcp.compute.InstanceGroupManagerArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceGroupManagerVersionArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceGroupManagerAllInstancesConfigArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceGroupManagerNamedPortArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceGroupManagerAutoHealingPoliciesArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var autohealing = new HealthCheck(&#34;autohealing&#34;, HealthCheckArgs.builder()        
 *             .checkIntervalSec(5)
 *             .timeoutSec(5)
 *             .healthyThreshold(2)
 *             .unhealthyThreshold(10)
 *             .httpHealthCheck(HealthCheckHttpHealthCheckArgs.builder()
 *                 .requestPath(&#34;/healthz&#34;)
 *                 .port(&#34;8080&#34;)
 *                 .build())
 *             .build());
 * 
 *         var appserver = new InstanceGroupManager(&#34;appserver&#34;, InstanceGroupManagerArgs.builder()        
 *             .baseInstanceName(&#34;app&#34;)
 *             .zone(&#34;us-central1-a&#34;)
 *             .versions(InstanceGroupManagerVersionArgs.builder()
 *                 .instanceTemplate(google_compute_instance_template.appserver().self_link_unique())
 *                 .build())
 *             .allInstancesConfig(InstanceGroupManagerAllInstancesConfigArgs.builder()
 *                 .metadata(Map.of(&#34;metadata_key&#34;, &#34;metadata_value&#34;))
 *                 .labels(Map.of(&#34;label_key&#34;, &#34;label_value&#34;))
 *                 .build())
 *             .targetPools(google_compute_target_pool.appserver().id())
 *             .targetSize(2)
 *             .namedPorts(InstanceGroupManagerNamedPortArgs.builder()
 *                 .name(&#34;customhttp&#34;)
 *                 .port(8888)
 *                 .build())
 *             .autoHealingPolicies(InstanceGroupManagerAutoHealingPoliciesArgs.builder()
 *                 .healthCheck(autohealing.id())
 *                 .initialDelaySec(300)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With Multiple Versions (`Google-Beta` Provider)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceGroupManager;
 * import com.pulumi.gcp.compute.InstanceGroupManagerArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceGroupManagerVersionArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceGroupManagerVersionTargetSizeArgs;
 * import com.pulumi.resources.CustomResourceOptions;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var appserver = new InstanceGroupManager(&#34;appserver&#34;, InstanceGroupManagerArgs.builder()        
 *             .baseInstanceName(&#34;app&#34;)
 *             .zone(&#34;us-central1-a&#34;)
 *             .targetSize(5)
 *             .versions(            
 *                 InstanceGroupManagerVersionArgs.builder()
 *                     .name(&#34;appserver&#34;)
 *                     .instanceTemplate(google_compute_instance_template.appserver().self_link_unique())
 *                     .build(),
 *                 InstanceGroupManagerVersionArgs.builder()
 *                     .name(&#34;appserver-canary&#34;)
 *                     .instanceTemplate(google_compute_instance_template.appserver-canary().self_link_unique())
 *                     .targetSize(InstanceGroupManagerVersionTargetSizeArgs.builder()
 *                         .fixed(1)
 *                         .build())
 *                     .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Instance group managers can be imported using any of these accepted formats* `projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}` * `{{project}}/{{zone}}/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, instance group managers can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager default projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager default {{project}}/{{zone}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/instanceGroupManager:InstanceGroupManager")
public class InstanceGroupManager extends com.pulumi.resources.CustomResource {
    /**
     * Properties to set on all instances in the group. After setting
     * allInstancesConfig on the group, you must update the group&#39;s instances to
     * apply the configuration.
     * 
     */
    @Export(name="allInstancesConfig", refs={InstanceGroupManagerAllInstancesConfig.class}, tree="[0]")
    private Output</* @Nullable */ InstanceGroupManagerAllInstancesConfig> allInstancesConfig;

    /**
     * @return Properties to set on all instances in the group. After setting
     * allInstancesConfig on the group, you must update the group&#39;s instances to
     * apply the configuration.
     * 
     */
    public Output<Optional<InstanceGroupManagerAllInstancesConfig>> allInstancesConfig() {
        return Codegen.optional(this.allInstancesConfig);
    }
    /**
     * The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     * 
     */
    @Export(name="autoHealingPolicies", refs={InstanceGroupManagerAutoHealingPolicies.class}, tree="[0]")
    private Output</* @Nullable */ InstanceGroupManagerAutoHealingPolicies> autoHealingPolicies;

    /**
     * @return The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     * 
     */
    public Output<Optional<InstanceGroupManagerAutoHealingPolicies>> autoHealingPolicies() {
        return Codegen.optional(this.autoHealingPolicies);
    }
    /**
     * The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     * 
     */
    @Export(name="baseInstanceName", refs={String.class}, tree="[0]")
    private Output<String> baseInstanceName;

    /**
     * @return The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     * 
     */
    public Output<String> baseInstanceName() {
        return this.baseInstanceName;
    }
    /**
     * An optional textual description of the instance
     * group manager.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional textual description of the instance
     * group manager.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The fingerprint of the instance group manager.
     * 
     */
    @Export(name="fingerprint", refs={String.class}, tree="[0]")
    private Output<String> fingerprint;

    /**
     * @return The fingerprint of the instance group manager.
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * The full URL of the instance group created by the manager.
     * 
     */
    @Export(name="instanceGroup", refs={String.class}, tree="[0]")
    private Output<String> instanceGroup;

    /**
     * @return The full URL of the instance group created by the manager.
     * 
     */
    public Output<String> instanceGroup() {
        return this.instanceGroup;
    }
    /**
     * The instance lifecycle policy for this managed instance group.
     * 
     */
    @Export(name="instanceLifecyclePolicy", refs={InstanceGroupManagerInstanceLifecyclePolicy.class}, tree="[0]")
    private Output<InstanceGroupManagerInstanceLifecyclePolicy> instanceLifecyclePolicy;

    /**
     * @return The instance lifecycle policy for this managed instance group.
     * 
     */
    public Output<InstanceGroupManagerInstanceLifecyclePolicy> instanceLifecyclePolicy() {
        return this.instanceLifecyclePolicy;
    }
    /**
     * Pagination behavior of the `listManagedInstances` API
     * method for this managed instance group. Valid values are: `PAGELESS`, `PAGINATED`.
     * If `PAGELESS` (default), Pagination is disabled for the group&#39;s `listManagedInstances` API method.
     * `maxResults` and `pageToken` query parameters are ignored and all instances are returned in a single
     * response. If `PAGINATED`, pagination is enabled, `maxResults` and `pageToken` query parameters are
     * respected.
     * 
     */
    @Export(name="listManagedInstancesResults", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> listManagedInstancesResults;

    /**
     * @return Pagination behavior of the `listManagedInstances` API
     * method for this managed instance group. Valid values are: `PAGELESS`, `PAGINATED`.
     * If `PAGELESS` (default), Pagination is disabled for the group&#39;s `listManagedInstances` API method.
     * `maxResults` and `pageToken` query parameters are ignored and all instances are returned in a single
     * response. If `PAGINATED`, pagination is enabled, `maxResults` and `pageToken` query parameters are
     * respected.
     * 
     */
    public Output<Optional<String>> listManagedInstancesResults() {
        return Codegen.optional(this.listManagedInstancesResults);
    }
    /**
     * The name of the instance group manager. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the instance group manager. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     * 
     */
    @Export(name="namedPorts", refs={List.class,InstanceGroupManagerNamedPort.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceGroupManagerNamedPort>> namedPorts;

    /**
     * @return The named port configuration. See the section below
     * for details on configuration.
     * 
     */
    public Output<Optional<List<InstanceGroupManagerNamedPort>>> namedPorts() {
        return Codegen.optional(this.namedPorts);
    }
    @Export(name="operation", refs={String.class}, tree="[0]")
    private Output<String> operation;

    public Output<String> operation() {
        return this.operation;
    }
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The URL of the created resource.
     * 
     */
    @Export(name="selfLink", refs={String.class}, tree="[0]")
    private Output<String> selfLink;

    /**
     * @return The URL of the created resource.
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }
    /**
     * Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
     * 
     */
    @Export(name="statefulDisks", refs={List.class,InstanceGroupManagerStatefulDisk.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceGroupManagerStatefulDisk>> statefulDisks;

    /**
     * @return Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
     * 
     */
    public Output<Optional<List<InstanceGroupManagerStatefulDisk>>> statefulDisks() {
        return Codegen.optional(this.statefulDisks);
    }
    /**
     * External network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
     * 
     */
    @Export(name="statefulExternalIps", refs={List.class,InstanceGroupManagerStatefulExternalIp.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceGroupManagerStatefulExternalIp>> statefulExternalIps;

    /**
     * @return External network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
     * 
     */
    public Output<Optional<List<InstanceGroupManagerStatefulExternalIp>>> statefulExternalIps() {
        return Codegen.optional(this.statefulExternalIps);
    }
    /**
     * Internal network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
     * 
     */
    @Export(name="statefulInternalIps", refs={List.class,InstanceGroupManagerStatefulInternalIp.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceGroupManagerStatefulInternalIp>> statefulInternalIps;

    /**
     * @return Internal network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
     * 
     */
    public Output<Optional<List<InstanceGroupManagerStatefulInternalIp>>> statefulInternalIps() {
        return Codegen.optional(this.statefulInternalIps);
    }
    /**
     * The status of this managed instance group.
     * 
     */
    @Export(name="statuses", refs={List.class,InstanceGroupManagerStatus.class}, tree="[0,1]")
    private Output<List<InstanceGroupManagerStatus>> statuses;

    /**
     * @return The status of this managed instance group.
     * 
     */
    public Output<List<InstanceGroupManagerStatus>> statuses() {
        return this.statuses;
    }
    /**
     * The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     * 
     */
    @Export(name="targetPools", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> targetPools;

    /**
     * @return The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     * 
     */
    public Output<Optional<List<String>>> targetPools() {
        return Codegen.optional(this.targetPools);
    }
    /**
     * The target number of running instances for this managed instance group. This value should always be explicitly set
     * unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to 0.
     * 
     */
    @Export(name="targetSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> targetSize;

    /**
     * @return The target number of running instances for this managed instance group. This value should always be explicitly set
     * unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to 0.
     * 
     */
    public Output<Integer> targetSize() {
        return this.targetSize;
    }
    /**
     * The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
     * 
     * ***
     * 
     */
    @Export(name="updatePolicy", refs={InstanceGroupManagerUpdatePolicy.class}, tree="[0]")
    private Output<InstanceGroupManagerUpdatePolicy> updatePolicy;

    /**
     * @return The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
     * 
     * ***
     * 
     */
    public Output<InstanceGroupManagerUpdatePolicy> updatePolicy() {
        return this.updatePolicy;
    }
    /**
     * Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Structure is documented below.
     * 
     */
    @Export(name="versions", refs={List.class,InstanceGroupManagerVersion.class}, tree="[0,1]")
    private Output<List<InstanceGroupManagerVersion>> versions;

    /**
     * @return Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Structure is documented below.
     * 
     */
    public Output<List<InstanceGroupManagerVersion>> versions() {
        return this.versions;
    }
    /**
     * Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, this provider will
     * continue trying until it times out.
     * 
     */
    @Export(name="waitForInstances", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> waitForInstances;

    /**
     * @return Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, this provider will
     * continue trying until it times out.
     * 
     */
    public Output<Optional<Boolean>> waitForInstances() {
        return Codegen.optional(this.waitForInstances);
    }
    /**
     * When used with `wait_for_instances` it specifies the status to wait for.
     * When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
     * set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
     * instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
     * 
     */
    @Export(name="waitForInstancesStatus", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> waitForInstancesStatus;

    /**
     * @return When used with `wait_for_instances` it specifies the status to wait for.
     * When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
     * set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
     * instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
     * 
     */
    public Output<Optional<String>> waitForInstancesStatus() {
        return Codegen.optional(this.waitForInstancesStatus);
    }
    /**
     * The zone that instances in this group should be created
     * in.
     * 
     * ***
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return The zone that instances in this group should be created
     * in.
     * 
     * ***
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceGroupManager(String name) {
        this(name, InstanceGroupManagerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceGroupManager(String name, InstanceGroupManagerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceGroupManager(String name, InstanceGroupManagerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instanceGroupManager:InstanceGroupManager", name, args == null ? InstanceGroupManagerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceGroupManager(String name, Output<String> id, @Nullable InstanceGroupManagerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instanceGroupManager:InstanceGroupManager", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static InstanceGroupManager get(String name, Output<String> id, @Nullable InstanceGroupManagerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceGroupManager(name, id, state, options);
    }
}

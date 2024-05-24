// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.alloydb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.alloydb.InstanceArgs;
import com.pulumi.gcp.alloydb.inputs.InstanceState;
import com.pulumi.gcp.alloydb.outputs.InstanceClientConnectionConfig;
import com.pulumi.gcp.alloydb.outputs.InstanceMachineConfig;
import com.pulumi.gcp.alloydb.outputs.InstanceNetworkConfig;
import com.pulumi.gcp.alloydb.outputs.InstanceQueryInsightsConfig;
import com.pulumi.gcp.alloydb.outputs.InstanceReadPoolConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### Alloydb Instance Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.alloydb.Cluster;
 * import com.pulumi.gcp.alloydb.ClusterArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterInitialUserArgs;
 * import com.pulumi.gcp.alloydb.Instance;
 * import com.pulumi.gcp.alloydb.InstanceArgs;
 * import com.pulumi.gcp.alloydb.inputs.InstanceMachineConfigArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
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
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .name("alloydb-network")
 *             .build());
 * 
 *         var defaultCluster = new Cluster("defaultCluster", ClusterArgs.builder()
 *             .clusterId("alloydb-cluster")
 *             .location("us-central1")
 *             .network(defaultNetwork.id())
 *             .initialUser(ClusterInitialUserArgs.builder()
 *                 .password("alloydb-cluster")
 *                 .build())
 *             .build());
 * 
 *         var default_ = new Instance("default", InstanceArgs.builder()
 *             .cluster(defaultCluster.name())
 *             .instanceId("alloydb-instance")
 *             .instanceType("PRIMARY")
 *             .machineConfig(InstanceMachineConfigArgs.builder()
 *                 .cpuCount(2)
 *                 .build())
 *             .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var privateIpAlloc = new GlobalAddress("privateIpAlloc", GlobalAddressArgs.builder()
 *             .name("alloydb-cluster")
 *             .addressType("INTERNAL")
 *             .purpose("VPC_PEERING")
 *             .prefixLength(16)
 *             .network(defaultNetwork.id())
 *             .build());
 * 
 *         var vpcConnection = new Connection("vpcConnection", ConnectionArgs.builder()
 *             .network(defaultNetwork.id())
 *             .service("servicenetworking.googleapis.com")
 *             .reservedPeeringRanges(privateIpAlloc.name())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Alloydb Secondary Instance Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.alloydb.Cluster;
 * import com.pulumi.gcp.alloydb.ClusterArgs;
 * import com.pulumi.gcp.alloydb.Instance;
 * import com.pulumi.gcp.alloydb.InstanceArgs;
 * import com.pulumi.gcp.alloydb.inputs.InstanceMachineConfigArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterContinuousBackupConfigArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterSecondaryConfigArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
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
 *         var default_ = new Network("default", NetworkArgs.builder()
 *             .name("alloydb-secondary-network")
 *             .build());
 * 
 *         var primary = new Cluster("primary", ClusterArgs.builder()
 *             .clusterId("alloydb-primary-cluster")
 *             .location("us-central1")
 *             .network(default_.id())
 *             .build());
 * 
 *         var primaryInstance = new Instance("primaryInstance", InstanceArgs.builder()
 *             .cluster(primary.name())
 *             .instanceId("alloydb-primary-instance")
 *             .instanceType("PRIMARY")
 *             .machineConfig(InstanceMachineConfigArgs.builder()
 *                 .cpuCount(2)
 *                 .build())
 *             .build());
 * 
 *         var secondary = new Cluster("secondary", ClusterArgs.builder()
 *             .clusterId("alloydb-secondary-cluster")
 *             .location("us-east1")
 *             .network(default_.id())
 *             .clusterType("SECONDARY")
 *             .continuousBackupConfig(ClusterContinuousBackupConfigArgs.builder()
 *                 .enabled(false)
 *                 .build())
 *             .secondaryConfig(ClusterSecondaryConfigArgs.builder()
 *                 .primaryClusterName(primary.name())
 *                 .build())
 *             .deletionPolicy("FORCE")
 *             .build());
 * 
 *         var secondaryInstance = new Instance("secondaryInstance", InstanceArgs.builder()
 *             .cluster(secondary.name())
 *             .instanceId("alloydb-secondary-instance")
 *             .instanceType(secondary.clusterType())
 *             .machineConfig(InstanceMachineConfigArgs.builder()
 *                 .cpuCount(2)
 *                 .build())
 *             .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var privateIpAlloc = new GlobalAddress("privateIpAlloc", GlobalAddressArgs.builder()
 *             .name("alloydb-secondary-instance")
 *             .addressType("INTERNAL")
 *             .purpose("VPC_PEERING")
 *             .prefixLength(16)
 *             .network(default_.id())
 *             .build());
 * 
 *         var vpcConnection = new Connection("vpcConnection", ConnectionArgs.builder()
 *             .network(default_.id())
 *             .service("servicenetworking.googleapis.com")
 *             .reservedPeeringRanges(privateIpAlloc.name())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Instance can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/instances/{{instance_id}}`
 * 
 * * `{{project}}/{{location}}/{{cluster}}/{{instance_id}}`
 * 
 * * `{{location}}/{{cluster}}/{{instance_id}}`
 * 
 * When using the `pulumi import` command, Instance can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:alloydb/instance:Instance default projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/instances/{{instance_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:alloydb/instance:Instance default {{project}}/{{location}}/{{cluster}}/{{instance_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:alloydb/instance:Instance default {{location}}/{{cluster}}/{{instance_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:alloydb/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    @Export(name="annotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> annotations;

    /**
     * @return Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> annotations() {
        return Codegen.optional(this.annotations);
    }
    /**
     * &#39;Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
     * Note that primary and read instances can have different availability types.
     * Only READ_POOL instance supports ZONAL type. Users can&#39;t specify the zone for READ_POOL instance.
     * Zone is automatically chosen from the list of zones in the region specified.
     * Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
     * can have regional availability (nodes are present in 2 or more zones in a region).&#39;
     * Possible values are: `AVAILABILITY_TYPE_UNSPECIFIED`, `ZONAL`, `REGIONAL`.
     * 
     */
    @Export(name="availabilityType", refs={String.class}, tree="[0]")
    private Output<String> availabilityType;

    /**
     * @return &#39;Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
     * Note that primary and read instances can have different availability types.
     * Only READ_POOL instance supports ZONAL type. Users can&#39;t specify the zone for READ_POOL instance.
     * Zone is automatically chosen from the list of zones in the region specified.
     * Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
     * can have regional availability (nodes are present in 2 or more zones in a region).&#39;
     * Possible values are: `AVAILABILITY_TYPE_UNSPECIFIED`, `ZONAL`, `REGIONAL`.
     * 
     */
    public Output<String> availabilityType() {
        return this.availabilityType;
    }
    /**
     * Client connection specific configurations.
     * Structure is documented below.
     * 
     */
    @Export(name="clientConnectionConfig", refs={InstanceClientConnectionConfig.class}, tree="[0]")
    private Output<InstanceClientConnectionConfig> clientConnectionConfig;

    /**
     * @return Client connection specific configurations.
     * Structure is documented below.
     * 
     */
    public Output<InstanceClientConnectionConfig> clientConnectionConfig() {
        return this.clientConnectionConfig;
    }
    /**
     * Identifies the alloydb cluster. Must be in the format
     * &#39;projects/{project}/locations/{location}/clusters/{cluster_id}&#39;
     * 
     */
    @Export(name="cluster", refs={String.class}, tree="[0]")
    private Output<String> cluster;

    /**
     * @return Identifies the alloydb cluster. Must be in the format
     * &#39;projects/{project}/locations/{location}/clusters/{cluster_id}&#39;
     * 
     */
    public Output<String> cluster() {
        return this.cluster;
    }
    /**
     * Time the Instance was created in UTC.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Time the Instance was created in UTC.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
     * 
     */
    @Export(name="databaseFlags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> databaseFlags;

    /**
     * @return Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
     * 
     */
    public Output<Map<String,String>> databaseFlags() {
        return this.databaseFlags;
    }
    /**
     * User-settable and human-readable display name for the Instance.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return User-settable and human-readable display name for the Instance.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    @Export(name="effectiveAnnotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveAnnotations;

    public Output<Map<String,String>> effectiveAnnotations() {
        return this.effectiveAnnotations;
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
     * 
     */
    @Export(name="gceZone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> gceZone;

    /**
     * @return The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
     * 
     */
    public Output<Optional<String>> gceZone() {
        return Codegen.optional(this.gceZone);
    }
    /**
     * The ID of the alloydb instance.
     * 
     * ***
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the alloydb instance.
     * 
     * ***
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output<String> instanceType;

    public Output<String> instanceType() {
        return this.instanceType;
    }
    /**
     * The IP address for the Instance. This is the connection endpoint for an end-user application.
     * 
     */
    @Export(name="ipAddress", refs={String.class}, tree="[0]")
    private Output<String> ipAddress;

    /**
     * @return The IP address for the Instance. This is the connection endpoint for an end-user application.
     * 
     */
    public Output<String> ipAddress() {
        return this.ipAddress;
    }
    /**
     * User-defined labels for the alloydb instance.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return User-defined labels for the alloydb instance.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Configurations for the machines that host the underlying database engine.
     * Structure is documented below.
     * 
     */
    @Export(name="machineConfig", refs={InstanceMachineConfig.class}, tree="[0]")
    private Output<InstanceMachineConfig> machineConfig;

    /**
     * @return Configurations for the machines that host the underlying database engine.
     * Structure is documented below.
     * 
     */
    public Output<InstanceMachineConfig> machineConfig() {
        return this.machineConfig;
    }
    /**
     * The name of the instance resource.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the instance resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Instance level network configuration.
     * Structure is documented below.
     * 
     */
    @Export(name="networkConfig", refs={InstanceNetworkConfig.class}, tree="[0]")
    private Output</* @Nullable */ InstanceNetworkConfig> networkConfig;

    /**
     * @return Instance level network configuration.
     * Structure is documented below.
     * 
     */
    public Output<Optional<InstanceNetworkConfig>> networkConfig() {
        return Codegen.optional(this.networkConfig);
    }
    /**
     * The public IP addresses for the Instance. This is available ONLY when
     * networkConfig.enablePublicIp is set to true. This is the connection
     * endpoint for an end-user application.
     * 
     */
    @Export(name="publicIpAddress", refs={String.class}, tree="[0]")
    private Output<String> publicIpAddress;

    /**
     * @return The public IP addresses for the Instance. This is available ONLY when
     * networkConfig.enablePublicIp is set to true. This is the connection
     * endpoint for an end-user application.
     * 
     */
    public Output<String> publicIpAddress() {
        return this.publicIpAddress;
    }
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * Configuration for query insights.
     * Structure is documented below.
     * 
     */
    @Export(name="queryInsightsConfig", refs={InstanceQueryInsightsConfig.class}, tree="[0]")
    private Output<InstanceQueryInsightsConfig> queryInsightsConfig;

    /**
     * @return Configuration for query insights.
     * Structure is documented below.
     * 
     */
    public Output<InstanceQueryInsightsConfig> queryInsightsConfig() {
        return this.queryInsightsConfig;
    }
    /**
     * Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.
     * Structure is documented below.
     * 
     */
    @Export(name="readPoolConfig", refs={InstanceReadPoolConfig.class}, tree="[0]")
    private Output</* @Nullable */ InstanceReadPoolConfig> readPoolConfig;

    /**
     * @return Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.
     * Structure is documented below.
     * 
     */
    public Output<Optional<InstanceReadPoolConfig>> readPoolConfig() {
        return Codegen.optional(this.readPoolConfig);
    }
    /**
     * Set to true if the current state of Instance does not match the user&#39;s intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.
     * 
     */
    @Export(name="reconciling", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reconciling;

    /**
     * @return Set to true if the current state of Instance does not match the user&#39;s intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.
     * 
     */
    public Output<Boolean> reconciling() {
        return this.reconciling;
    }
    /**
     * The current state of the alloydb instance.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The current state of the alloydb instance.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The system-generated UID of the resource.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return The system-generated UID of the resource.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Time the Instance was updated in UTC.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Time the Instance was updated in UTC.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:alloydb/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:alloydb/instance:Instance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "effectiveLabels",
                "pulumiLabels"
            ))
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}

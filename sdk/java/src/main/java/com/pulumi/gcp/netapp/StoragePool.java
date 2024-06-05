// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.netapp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.netapp.StoragePoolArgs;
import com.pulumi.gcp.netapp.inputs.StoragePoolState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Storage pools act as containers for volumes. All volumes in a storage pool share the following information:
 * * Location
 * * Service level
 * * Virtual Private Cloud (VPC) network
 * * Active Directory policy
 * * LDAP use for NFS volumes, if applicable
 * * Customer-managed encryption key (CMEK) policy
 * 
 * The capacity of the pool can be split up and assigned to volumes within the pool. Storage pools are a billable
 * component of NetApp Volumes. Billing is based on the location, service level, and capacity allocated to a pool
 * independent of consumption at the volume level.
 * 
 * To get more information about storagePool, see:
 * 
 * * [API documentation](https://cloud.google.com/netapp/volumes/docs/reference/rest/v1/projects.locations.storagePools)
 * * How-to Guides
 *     * [Quickstart documentation](https://cloud.google.com/netapp/volumes/docs/get-started/quickstarts/create-storage-pool)
 * 
 * ## Example Usage
 * 
 * ### Storage Pool Create
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
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.compute.NetworkPeeringRoutesConfig;
 * import com.pulumi.gcp.compute.NetworkPeeringRoutesConfigArgs;
 * import com.pulumi.gcp.netapp.StoragePool;
 * import com.pulumi.gcp.netapp.StoragePoolArgs;
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
 *         // Create a network or use datasource to reference existing network
 *         var peeringNetwork = new Network("peeringNetwork", NetworkArgs.builder()
 *             .name("test-network")
 *             .build());
 * 
 *         // Reserve a CIDR for NetApp Volumes to use
 *         // When using shared-VPCs, this resource needs to be created in host project
 *         var privateIpAlloc = new GlobalAddress("privateIpAlloc", GlobalAddressArgs.builder()
 *             .name("test-address")
 *             .purpose("VPC_PEERING")
 *             .addressType("INTERNAL")
 *             .prefixLength(16)
 *             .network(peeringNetwork.id())
 *             .build());
 * 
 *         // Create a Private Service Access connection
 *         // When using shared-VPCs, this resource needs to be created in host project
 *         var default_ = new Connection("default", ConnectionArgs.builder()
 *             .network(peeringNetwork.id())
 *             .service("netapp.servicenetworking.goog")
 *             .reservedPeeringRanges(privateIpAlloc.name())
 *             .build());
 * 
 *         // Modify the PSA Connection to allow import/export of custom routes
 *         // When using shared-VPCs, this resource needs to be created in host project
 *         var routeUpdates = new NetworkPeeringRoutesConfig("routeUpdates", NetworkPeeringRoutesConfigArgs.builder()
 *             .peering(default_.peering())
 *             .network(peeringNetwork.name())
 *             .importCustomRoutes(true)
 *             .exportCustomRoutes(true)
 *             .build());
 * 
 *         // Create a storage pool
 *         // Create this resource in the project which is expected to own the volumes
 *         var testPool = new StoragePool("testPool", StoragePoolArgs.builder()
 *             .name("test-pool")
 *             .location("us-central1")
 *             .serviceLevel("PREMIUM")
 *             .capacityGib("2048")
 *             .network(peeringNetwork.id())
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
 * storagePool can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/storagePools/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{name}}`
 * 
 * * `{{location}}/{{name}}`
 * 
 * When using the `pulumi import` command, storagePool can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:netapp/storagePool:StoragePool default projects/{{project}}/locations/{{location}}/storagePools/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:netapp/storagePool:StoragePool default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:netapp/storagePool:StoragePool default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:netapp/storagePool:StoragePool")
public class StoragePool extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the Active Directory policy to be used. Format: `projects/{{project}}/locations/{{location}}/activeDirectories/{{name}}`.
     * The policy needs to be in the same location as the storage pool.
     * 
     */
    @Export(name="activeDirectory", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> activeDirectory;

    /**
     * @return Specifies the Active Directory policy to be used. Format: `projects/{{project}}/locations/{{location}}/activeDirectories/{{name}}`.
     * The policy needs to be in the same location as the storage pool.
     * 
     */
    public Output<Optional<String>> activeDirectory() {
        return Codegen.optional(this.activeDirectory);
    }
    /**
     * Capacity of the storage pool (in GiB).
     * 
     */
    @Export(name="capacityGib", refs={String.class}, tree="[0]")
    private Output<String> capacityGib;

    /**
     * @return Capacity of the storage pool (in GiB).
     * 
     */
    public Output<String> capacityGib() {
        return this.capacityGib;
    }
    /**
     * An optional description of this resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * Reports if volumes in the pool are encrypted using a Google-managed encryption key or CMEK.
     * 
     */
    @Export(name="encryptionType", refs={String.class}, tree="[0]")
    private Output<String> encryptionType;

    /**
     * @return Reports if volumes in the pool are encrypted using a Google-managed encryption key or CMEK.
     * 
     */
    public Output<String> encryptionType() {
        return this.encryptionType;
    }
    /**
     * Specifies the CMEK policy to be used for volume encryption. Format: `projects/{{project}}/locations/{{location}}/kmsConfigs/{{name}}`.
     * The policy needs to be in the same location as the storage pool.
     * 
     */
    @Export(name="kmsConfig", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsConfig;

    /**
     * @return Specifies the CMEK policy to be used for volume encryption. Format: `projects/{{project}}/locations/{{location}}/kmsConfigs/{{name}}`.
     * The policy needs to be in the same location as the storage pool.
     * 
     */
    public Output<Optional<String>> kmsConfig() {
        return Codegen.optional(this.kmsConfig);
    }
    /**
     * Labels as key value pairs. Example: `{ &#34;owner&#34;: &#34;Bob&#34;, &#34;department&#34;: &#34;finance&#34;, &#34;purpose&#34;: &#34;testing&#34; }`.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels as key value pairs. Example: `{ &#34;owner&#34;: &#34;Bob&#34;, &#34;department&#34;: &#34;finance&#34;, &#34;purpose&#34;: &#34;testing&#34; }`.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * When enabled, the volumes uses Active Directory as LDAP name service for UID/GID lookups. Required to enable extended group support for NFSv3,
     * using security identifiers for NFSv4.1 or principal names for kerberized NFSv4.1.
     * 
     */
    @Export(name="ldapEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ldapEnabled;

    /**
     * @return When enabled, the volumes uses Active Directory as LDAP name service for UID/GID lookups. Required to enable extended group support for NFSv3,
     * using security identifiers for NFSv4.1 or principal names for kerberized NFSv4.1.
     * 
     */
    public Output<Optional<Boolean>> ldapEnabled() {
        return Codegen.optional(this.ldapEnabled);
    }
    /**
     * Name of the location. Usually a region name, expect for some FLEX service level pools which require a zone name.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return Name of the location. Usually a region name, expect for some FLEX service level pools which require a zone name.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The resource name of the storage pool. Needs to be unique per location.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the storage pool. Needs to be unique per location.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * VPC network name with format: `projects/{{project}}/global/networks/{{network}}`
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return VPC network name with format: `projects/{{project}}/global/networks/{{network}}`
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
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
     * Service level of the storage pool.
     * Possible values are: `PREMIUM`, `EXTREME`, `STANDARD`, `FLEX`.
     * 
     */
    @Export(name="serviceLevel", refs={String.class}, tree="[0]")
    private Output<String> serviceLevel;

    /**
     * @return Service level of the storage pool.
     * Possible values are: `PREMIUM`, `EXTREME`, `STANDARD`, `FLEX`.
     * 
     */
    public Output<String> serviceLevel() {
        return this.serviceLevel;
    }
    /**
     * Size allocated to volumes in the storage pool (in GiB).
     * 
     */
    @Export(name="volumeCapacityGib", refs={String.class}, tree="[0]")
    private Output<String> volumeCapacityGib;

    /**
     * @return Size allocated to volumes in the storage pool (in GiB).
     * 
     */
    public Output<String> volumeCapacityGib() {
        return this.volumeCapacityGib;
    }
    /**
     * Number of volume in the storage pool.
     * 
     */
    @Export(name="volumeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> volumeCount;

    /**
     * @return Number of volume in the storage pool.
     * 
     */
    public Output<Integer> volumeCount() {
        return this.volumeCount;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public StoragePool(String name) {
        this(name, StoragePoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public StoragePool(String name, StoragePoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public StoragePool(String name, StoragePoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:netapp/storagePool:StoragePool", name, args == null ? StoragePoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private StoragePool(String name, Output<String> id, @Nullable StoragePoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:netapp/storagePool:StoragePool", name, state, makeResourceOptions(options, id));
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
    public static StoragePool get(String name, Output<String> id, @Nullable StoragePoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new StoragePool(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vmwareengine;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vmwareengine.ExternalAddressArgs;
import com.pulumi.gcp.vmwareengine.inputs.ExternalAddressState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An allocated external IP address and its corresponding internal IP address in a private cloud.
 * 
 * To get more information about ExternalAddress, see:
 * 
 * * [API documentation](https://cloud.google.com/vmware-engine/docs/reference/rest/v1/projects.locations.privateClouds.externalAddresses)
 * 
 * ## Example Usage
 * 
 * ### Vmware Engine External Address Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vmwareengine.Network;
 * import com.pulumi.gcp.vmwareengine.NetworkArgs;
 * import com.pulumi.gcp.vmwareengine.PrivateCloud;
 * import com.pulumi.gcp.vmwareengine.PrivateCloudArgs;
 * import com.pulumi.gcp.vmwareengine.inputs.PrivateCloudNetworkConfigArgs;
 * import com.pulumi.gcp.vmwareengine.inputs.PrivateCloudManagementClusterArgs;
 * import com.pulumi.gcp.vmwareengine.NetworkPolicy;
 * import com.pulumi.gcp.vmwareengine.NetworkPolicyArgs;
 * import com.pulumi.gcp.vmwareengine.ExternalAddress;
 * import com.pulumi.gcp.vmwareengine.ExternalAddressArgs;
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
 *         var external_address_nw = new Network("external-address-nw", NetworkArgs.builder()
 *             .name("pc-nw")
 *             .location("global")
 *             .type("STANDARD")
 *             .description("PC network description.")
 *             .build());
 * 
 *         var external_address_pc = new PrivateCloud("external-address-pc", PrivateCloudArgs.builder()
 *             .location("-a")
 *             .name("sample-pc")
 *             .description("Sample test PC.")
 *             .networkConfig(PrivateCloudNetworkConfigArgs.builder()
 *                 .managementCidr("192.168.50.0/24")
 *                 .vmwareEngineNetwork(external_address_nw.id())
 *                 .build())
 *             .managementCluster(PrivateCloudManagementClusterArgs.builder()
 *                 .clusterId("sample-mgmt-cluster")
 *                 .nodeTypeConfigs(PrivateCloudManagementClusterNodeTypeConfigArgs.builder()
 *                     .nodeTypeId("standard-72")
 *                     .nodeCount(3)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var external_address_np = new NetworkPolicy("external-address-np", NetworkPolicyArgs.builder()
 *             .location("")
 *             .name("sample-np")
 *             .edgeServicesCidr("192.168.30.0/26")
 *             .vmwareEngineNetwork(external_address_nw.id())
 *             .build());
 * 
 *         var vmw_engine_external_address = new ExternalAddress("vmw-engine-external-address", ExternalAddressArgs.builder()
 *             .name("sample-external-address")
 *             .parent(external_address_pc.id())
 *             .internalIp("192.168.0.66")
 *             .description("Sample description.")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(external_address_np)
 *                 .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ExternalAddress can be imported using any of these accepted formats:
 * 
 * * `{{parent}}/externalAddresses/{{name}}`
 * 
 * When using the `pulumi import` command, ExternalAddress can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:vmwareengine/externalAddress:ExternalAddress default {{parent}}/externalAddresses/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:vmwareengine/externalAddress:ExternalAddress")
public class ExternalAddress extends com.pulumi.resources.CustomResource {
    /**
     * Creation time of this resource.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and
     * up to nine fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Creation time of this resource.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and
     * up to nine fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * User-provided description for this resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return User-provided description for this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The external IP address of a workload VM.
     * 
     */
    @Export(name="externalIp", refs={String.class}, tree="[0]")
    private Output<String> externalIp;

    /**
     * @return The external IP address of a workload VM.
     * 
     */
    public Output<String> externalIp() {
        return this.externalIp;
    }
    /**
     * The internal IP address of a workload VM.
     * 
     */
    @Export(name="internalIp", refs={String.class}, tree="[0]")
    private Output<String> internalIp;

    /**
     * @return The internal IP address of a workload VM.
     * 
     */
    public Output<String> internalIp() {
        return this.internalIp;
    }
    /**
     * The ID of the external IP Address.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The ID of the external IP Address.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The resource name of the private cloud to create a new external address in.
     * Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
     * For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return The resource name of the private cloud to create a new external address in.
     * Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
     * For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * State of the resource.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State of the resource.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * System-generated unique identifier for the resource.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return System-generated unique identifier for the resource.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Last updated time of this resource.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine
     * fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Last updated time of this resource.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine
     * fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ExternalAddress(java.lang.String name) {
        this(name, ExternalAddressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ExternalAddress(java.lang.String name, ExternalAddressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ExternalAddress(java.lang.String name, ExternalAddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vmwareengine/externalAddress:ExternalAddress", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ExternalAddress(java.lang.String name, Output<java.lang.String> id, @Nullable ExternalAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vmwareengine/externalAddress:ExternalAddress", name, state, makeResourceOptions(options, id), false);
    }

    private static ExternalAddressArgs makeArgs(ExternalAddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ExternalAddressArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ExternalAddress get(java.lang.String name, Output<java.lang.String> id, @Nullable ExternalAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ExternalAddress(name, id, state, options);
    }
}

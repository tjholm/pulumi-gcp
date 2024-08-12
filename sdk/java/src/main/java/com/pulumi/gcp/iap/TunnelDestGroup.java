// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.iap.TunnelDestGroupArgs;
import com.pulumi.gcp.iap.inputs.TunnelDestGroupState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Tunnel destination groups represent resources that have the same tunnel access restrictions.
 * 
 * To get more information about TunnelDestGroup, see:
 * 
 * * [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.iap_tunnel.locations.destGroups)
 * * How-to Guides
 *     * [Set up IAP TCP forwarding with an IP address or hostname in a Google Cloud or non-Google Cloud environment](https://cloud.google.com/iap/docs/tcp-by-host)
 * 
 * ## Example Usage
 * 
 * ### Iap Destgroup
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.TunnelDestGroup;
 * import com.pulumi.gcp.iap.TunnelDestGroupArgs;
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
 *         var destGroup = new TunnelDestGroup("destGroup", TunnelDestGroupArgs.builder()
 *             .region("us-central1")
 *             .groupName("testgroup_2605")
 *             .cidrs(            
 *                 "10.1.0.0/16",
 *                 "192.168.10.0/24")
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
 * TunnelDestGroup can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}`
 * 
 * * `{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}`
 * 
 * * `{{project}}/{{region}}/{{group_name}}`
 * 
 * * `{{region}}/destGroups/{{group_name}}`
 * 
 * * `{{region}}/{{group_name}}`
 * 
 * * `{{group_name}}`
 * 
 * When using the `pulumi import` command, TunnelDestGroup can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{project}}/{{region}}/{{group_name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{region}}/destGroups/{{group_name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{region}}/{{group_name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{group_name}}
 * ```
 * 
 */
@ResourceType(type="gcp:iap/tunnelDestGroup:TunnelDestGroup")
public class TunnelDestGroup extends com.pulumi.resources.CustomResource {
    /**
     * List of CIDRs that this group applies to.
     * 
     */
    @Export(name="cidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> cidrs;

    /**
     * @return List of CIDRs that this group applies to.
     * 
     */
    public Output<Optional<List<String>>> cidrs() {
        return Codegen.optional(this.cidrs);
    }
    /**
     * List of FQDNs that this group applies to.
     * 
     */
    @Export(name="fqdns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> fqdns;

    /**
     * @return List of FQDNs that this group applies to.
     * 
     */
    public Output<Optional<List<String>>> fqdns() {
        return Codegen.optional(this.fqdns);
    }
    /**
     * Unique tunnel destination group name.
     * 
     * ***
     * 
     */
    @Export(name="groupName", refs={String.class}, tree="[0]")
    private Output<String> groupName;

    /**
     * @return Unique tunnel destination group name.
     * 
     * ***
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
    }
    /**
     * Full resource name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Full resource name.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * The region of the tunnel group. Must be the same as the network resources in the group.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region of the tunnel group. Must be the same as the network resources in the group.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TunnelDestGroup(java.lang.String name) {
        this(name, TunnelDestGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TunnelDestGroup(java.lang.String name, TunnelDestGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TunnelDestGroup(java.lang.String name, TunnelDestGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/tunnelDestGroup:TunnelDestGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TunnelDestGroup(java.lang.String name, Output<java.lang.String> id, @Nullable TunnelDestGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/tunnelDestGroup:TunnelDestGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static TunnelDestGroupArgs makeArgs(TunnelDestGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TunnelDestGroupArgs.Empty : args;
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
    public static TunnelDestGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable TunnelDestGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TunnelDestGroup(name, id, state, options);
    }
}

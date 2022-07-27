// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.GlobalNetworkEndpointGroupArgs;
import com.pulumi.gcp.compute.inputs.GlobalNetworkEndpointGroupState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A global network endpoint group contains endpoints that reside outside of Google Cloud.
 * Currently a global network endpoint group can only support a single endpoint.
 * 
 * Recreating a global network endpoint group that&#39;s in use by another resource will give a
 * `resourceInUseByAnotherResource` error. Use `lifecycle.create_before_destroy`
 * to avoid this type of error.
 * 
 * To get more information about GlobalNetworkEndpointGroup, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/internet-neg-concepts)
 * 
 * ## Example Usage
 * ### Global Network Endpoint Group
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.GlobalNetworkEndpointGroup;
 * import com.pulumi.gcp.compute.GlobalNetworkEndpointGroupArgs;
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
 *         var neg = new GlobalNetworkEndpointGroup(&#34;neg&#34;, GlobalNetworkEndpointGroupArgs.builder()        
 *             .defaultPort(&#34;90&#34;)
 *             .networkEndpointType(&#34;INTERNET_FQDN_PORT&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Global Network Endpoint Group Ip Address
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.GlobalNetworkEndpointGroup;
 * import com.pulumi.gcp.compute.GlobalNetworkEndpointGroupArgs;
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
 *         var neg = new GlobalNetworkEndpointGroup(&#34;neg&#34;, GlobalNetworkEndpointGroupArgs.builder()        
 *             .defaultPort(90)
 *             .networkEndpointType(&#34;INTERNET_IP_PORT&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GlobalNetworkEndpointGroup can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup default projects/{{project}}/global/networkEndpointGroups/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup")
public class GlobalNetworkEndpointGroup extends com.pulumi.resources.CustomResource {
    /**
     * The default port used if the port number is not specified in the
     * network endpoint.
     * 
     */
    @Export(name="defaultPort", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> defaultPort;

    /**
     * @return The default port used if the port number is not specified in the
     * network endpoint.
     * 
     */
    public Output<Optional<Integer>> defaultPort() {
        return Codegen.optional(this.defaultPort);
    }
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource. Provide this property when
     * you create the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Type of network endpoints in this network endpoint group.
     * Possible values are `INTERNET_IP_PORT` and `INTERNET_FQDN_PORT`.
     * 
     */
    @Export(name="networkEndpointType", type=String.class, parameters={})
    private Output<String> networkEndpointType;

    /**
     * @return Type of network endpoints in this network endpoint group.
     * Possible values are `INTERNET_IP_PORT` and `INTERNET_FQDN_PORT`.
     * 
     */
    public Output<String> networkEndpointType() {
        return this.networkEndpointType;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
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
     * The URI of the created resource.
     * 
     */
    @Export(name="selfLink", type=String.class, parameters={})
    private Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GlobalNetworkEndpointGroup(String name) {
        this(name, GlobalNetworkEndpointGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GlobalNetworkEndpointGroup(String name, GlobalNetworkEndpointGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GlobalNetworkEndpointGroup(String name, GlobalNetworkEndpointGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup", name, args == null ? GlobalNetworkEndpointGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GlobalNetworkEndpointGroup(String name, Output<String> id, @Nullable GlobalNetworkEndpointGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup", name, state, makeResourceOptions(options, id));
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
    public static GlobalNetworkEndpointGroup get(String name, Output<String> id, @Nullable GlobalNetworkEndpointGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GlobalNetworkEndpointGroup(name, id, state, options);
    }
}

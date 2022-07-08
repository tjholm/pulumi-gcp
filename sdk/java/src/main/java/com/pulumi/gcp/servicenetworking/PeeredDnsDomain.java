// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.servicenetworking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.servicenetworking.PeeredDnsDomainArgs;
import com.pulumi.gcp.servicenetworking.inputs.PeeredDnsDomainState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows management of a single peered DNS domain for an existing Google Cloud Platform project.
 * 
 * When using Google Cloud DNS to manage internal DNS, create peered DNS domains to make your DNS available to services like Google Cloud Build.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var name = new PeeredDnsDomain(&#34;name&#34;, PeeredDnsDomainArgs.builder()        
 *             .dnsSuffix(&#34;example.com.&#34;)
 *             .network(&#34;default&#34;)
 *             .project(10000000)
 *             .service(&#34;peering-service&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Project peered DNS domains can be imported using the `service`, `project`, `network` and `name`, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:servicenetworking/peeredDnsDomain:PeeredDnsDomain my_domain services/{service}/projects/{project}/global/networks/{network}/peeredDnsDomains/{name}
 * ```
 * 
 *  Where- `service` is the service connection, defaults to `servicenetworking.googleapis.com`. - `project` is the producer project name. - `network` is the consumer network name. - `name` is the name of your peered DNS domain.
 * 
 */
@ResourceType(type="gcp:servicenetworking/peeredDnsDomain:PeeredDnsDomain")
public class PeeredDnsDomain extends com.pulumi.resources.CustomResource {
    /**
     * The DNS domain suffix of the peered DNS domain. Make sure to suffix with a `.` (dot).
     * 
     */
    @Export(name="dnsSuffix", type=String.class, parameters={})
    private Output<String> dnsSuffix;

    /**
     * @return The DNS domain suffix of the peered DNS domain. Make sure to suffix with a `.` (dot).
     * 
     */
    public Output<String> dnsSuffix() {
        return this.dnsSuffix;
    }
    /**
     * Internal name used for the peered DNS domain.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Internal name used for the peered DNS domain.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The network in the consumer project.
     * 
     */
    @Export(name="network", type=String.class, parameters={})
    private Output<String> network;

    /**
     * @return The network in the consumer project.
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * an identifier for the resource with format `services/{{service}}/projects/{{project}}/global/networks/{{network}}`
     * 
     */
    @Export(name="parent", type=String.class, parameters={})
    private Output<String> parent;

    /**
     * @return an identifier for the resource with format `services/{{service}}/projects/{{project}}/global/networks/{{network}}`
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * The producer project number. If not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The producer project number. If not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Private service connection between service and consumer network, defaults to `servicenetworking.googleapis.com`
     * 
     */
    @Export(name="service", type=String.class, parameters={})
    private Output</* @Nullable */ String> service;

    /**
     * @return Private service connection between service and consumer network, defaults to `servicenetworking.googleapis.com`
     * 
     */
    public Output<Optional<String>> service() {
        return Codegen.optional(this.service);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PeeredDnsDomain(String name) {
        this(name, PeeredDnsDomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PeeredDnsDomain(String name, PeeredDnsDomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PeeredDnsDomain(String name, PeeredDnsDomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:servicenetworking/peeredDnsDomain:PeeredDnsDomain", name, args == null ? PeeredDnsDomainArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PeeredDnsDomain(String name, Output<String> id, @Nullable PeeredDnsDomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:servicenetworking/peeredDnsDomain:PeeredDnsDomain", name, state, makeResourceOptions(options, id));
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
    public static PeeredDnsDomain get(String name, Output<String> id, @Nullable PeeredDnsDomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PeeredDnsDomain(name, id, state, options);
    }
}

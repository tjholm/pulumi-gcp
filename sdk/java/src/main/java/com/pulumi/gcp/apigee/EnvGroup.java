// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigee;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigee.EnvGroupArgs;
import com.pulumi.gcp.apigee.inputs.EnvGroupState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An `Environment group` in Apigee.
 * 
 * To get more information about Envgroup, see:
 * 
 * * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.envgroups/create)
 * * How-to Guides
 *     * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)
 * 
 * ## Example Usage
 * ### Apigee Environment Group Basic
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * import com.pulumi.resources.CustomResourceOptions;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var current = Output.of(OrganizationsFunctions.getClientConfig());
 * 
 *         var apigeeNetwork = new Network(&#34;apigeeNetwork&#34;);
 * 
 *         var apigeeRange = new GlobalAddress(&#34;apigeeRange&#34;, GlobalAddressArgs.builder()        
 *             .purpose(&#34;VPC_PEERING&#34;)
 *             .addressType(&#34;INTERNAL&#34;)
 *             .prefixLength(16)
 *             .network(apigeeNetwork.id())
 *             .build());
 * 
 *         var apigeeVpcConnection = new Connection(&#34;apigeeVpcConnection&#34;, ConnectionArgs.builder()        
 *             .network(apigeeNetwork.id())
 *             .service(&#34;servicenetworking.googleapis.com&#34;)
 *             .reservedPeeringRanges(apigeeRange.name())
 *             .build());
 * 
 *         var apigeeOrg = new Organization(&#34;apigeeOrg&#34;, OrganizationArgs.builder()        
 *             .analyticsRegion(&#34;us-central1&#34;)
 *             .projectId(current.apply(getClientConfigResult -&gt; getClientConfigResult.project()))
 *             .authorizedNetwork(apigeeNetwork.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(apigeeVpcConnection)
 *                 .build());
 * 
 *         var envGrp = new EnvGroup(&#34;envGrp&#34;, EnvGroupArgs.builder()        
 *             .hostnames(&#34;abc.foo.com&#34;)
 *             .orgId(apigeeOrg.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Envgroup can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:apigee/envGroup:EnvGroup default {{org_id}}/envgroups/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:apigee/envGroup:EnvGroup default {{org_id}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:apigee/envGroup:EnvGroup")
public class EnvGroup extends com.pulumi.resources.CustomResource {
    /**
     * Hostnames of the environment group.
     * 
     */
    @Export(name="hostnames", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> hostnames;

    /**
     * @return Hostnames of the environment group.
     * 
     */
    public Output<Optional<List<String>>> hostnames() {
        return Codegen.optional(this.hostnames);
    }
    /**
     * The resource ID of the environment group.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The resource ID of the environment group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The Apigee Organization associated with the Apigee environment group,
     * in the format `organizations/{{org_name}}`.
     * 
     */
    @Export(name="orgId", type=String.class, parameters={})
    private Output<String> orgId;

    /**
     * @return The Apigee Organization associated with the Apigee environment group,
     * in the format `organizations/{{org_name}}`.
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EnvGroup(String name) {
        this(name, EnvGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EnvGroup(String name, EnvGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EnvGroup(String name, EnvGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/envGroup:EnvGroup", name, args == null ? EnvGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EnvGroup(String name, Output<String> id, @Nullable EnvGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/envGroup:EnvGroup", name, state, makeResourceOptions(options, id));
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
    public static EnvGroup get(String name, Output<String> id, @Nullable EnvGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EnvGroup(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.TargetPoolArgs;
import com.pulumi.gcp.compute.inputs.TargetPoolState;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Target Pool within GCE. This is a collection of instances used as
 * target of a network load balancer (Forwarding Rule). For more information see
 * [the official
 * documentation](https://cloud.google.com/compute/docs/load-balancing/network/target-pools)
 * and [API](https://cloud.google.com/compute/docs/reference/latest/targetPools).
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.HttpHealthCheck;
 * import com.pulumi.gcp.compute.HttpHealthCheckArgs;
 * import com.pulumi.gcp.compute.TargetPool;
 * import com.pulumi.gcp.compute.TargetPoolArgs;
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
 *         var defaultHttpHealthCheck = new HttpHealthCheck("defaultHttpHealthCheck", HttpHealthCheckArgs.builder()
 *             .name("default")
 *             .requestPath("/")
 *             .checkIntervalSec(1)
 *             .timeoutSec(1)
 *             .build());
 * 
 *         var default_ = new TargetPool("default", TargetPoolArgs.builder()
 *             .name("instance-pool")
 *             .instances(            
 *                 "us-central1-a/myinstance1",
 *                 "us-central1-b/myinstance2")
 *             .healthChecks(defaultHttpHealthCheck.name())
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
 * Target pools can be imported using any of the following formats:
 * 
 * * `projects/{{project}}/regions/{{region}}/targetPools/{{name}}`
 * 
 * * `{{project}}/{{region}}/{{name}}`
 * 
 * * `{{region}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, target pools can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:compute/targetPool:TargetPool default projects/{{project}}/regions/{{region}}/targetPools/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/targetPool:TargetPool default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/targetPool:TargetPool default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/targetPool:TargetPool default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/targetPool:TargetPool")
public class TargetPool extends com.pulumi.resources.CustomResource {
    /**
     * URL to the backup target pool. Must also set
     * failover\_ratio.
     * 
     */
    @Export(name="backupPool", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backupPool;

    /**
     * @return URL to the backup target pool. Must also set
     * failover\_ratio.
     * 
     */
    public Output<Optional<String>> backupPool() {
        return Codegen.optional(this.backupPool);
    }
    /**
     * Textual description field.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Textual description field.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Ratio (0 to 1) of failed nodes before using the
     * backup pool (which must also be set).
     * 
     */
    @Export(name="failoverRatio", refs={Double.class}, tree="[0]")
    private Output</* @Nullable */ Double> failoverRatio;

    /**
     * @return Ratio (0 to 1) of failed nodes before using the
     * backup pool (which must also be set).
     * 
     */
    public Output<Optional<Double>> failoverRatio() {
        return Codegen.optional(this.failoverRatio);
    }
    /**
     * List of zero or one health check name or self_link. Only
     * legacy `gcp.compute.HttpHealthCheck` is supported.
     * 
     */
    @Export(name="healthChecks", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthChecks;

    /**
     * @return List of zero or one health check name or self_link. Only
     * legacy `gcp.compute.HttpHealthCheck` is supported.
     * 
     */
    public Output<Optional<String>> healthChecks() {
        return Codegen.optional(this.healthChecks);
    }
    /**
     * List of instances in the pool. They can be given as
     * URLs, or in the form of &#34;zone/name&#34;. Note that the instances need not exist
     * at the time of target pool creation, so there is no need to use the
     * interpolation to create a dependency on the instances from the
     * target pool.
     * 
     */
    @Export(name="instances", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> instances;

    /**
     * @return List of instances in the pool. They can be given as
     * URLs, or in the form of &#34;zone/name&#34;. Note that the instances need not exist
     * at the time of target pool creation, so there is no need to use the
     * interpolation to create a dependency on the instances from the
     * target pool.
     * 
     */
    public Output<List<String>> instances() {
        return this.instances;
    }
    /**
     * A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * Where the target pool resides. Defaults to project
     * region.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Where the target pool resides. Defaults to project
     * region.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The resource URL for the security policy associated with this target pool.
     * 
     */
    @Export(name="securityPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityPolicy;

    /**
     * @return The resource URL for the security policy associated with this target pool.
     * 
     */
    public Output<Optional<String>> securityPolicy() {
        return Codegen.optional(this.securityPolicy);
    }
    /**
     * The URI of the created resource.
     * 
     */
    @Export(name="selfLink", refs={String.class}, tree="[0]")
    private Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }
    /**
     * How to distribute load. Options are &#34;NONE&#34; (no
     * affinity). &#34;CLIENT\_IP&#34; (hash of the source/dest addresses / ports), and
     * &#34;CLIENT\_IP\_PROTO&#34; also includes the protocol (default &#34;NONE&#34;).
     * 
     */
    @Export(name="sessionAffinity", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sessionAffinity;

    /**
     * @return How to distribute load. Options are &#34;NONE&#34; (no
     * affinity). &#34;CLIENT\_IP&#34; (hash of the source/dest addresses / ports), and
     * &#34;CLIENT\_IP\_PROTO&#34; also includes the protocol (default &#34;NONE&#34;).
     * 
     */
    public Output<Optional<String>> sessionAffinity() {
        return Codegen.optional(this.sessionAffinity);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TargetPool(String name) {
        this(name, TargetPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TargetPool(String name, @Nullable TargetPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TargetPool(String name, @Nullable TargetPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/targetPool:TargetPool", name, args == null ? TargetPoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TargetPool(String name, Output<String> id, @Nullable TargetPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/targetPool:TargetPool", name, state, makeResourceOptions(options, id));
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
    public static TargetPool get(String name, Output<String> id, @Nullable TargetPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TargetPool(name, id, state, options);
    }
}

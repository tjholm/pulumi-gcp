// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudtasks;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.cloudtasks.QueueArgs;
import com.pulumi.gcp.cloudtasks.inputs.QueueState;
import com.pulumi.gcp.cloudtasks.outputs.QueueAppEngineRoutingOverride;
import com.pulumi.gcp.cloudtasks.outputs.QueueRateLimits;
import com.pulumi.gcp.cloudtasks.outputs.QueueRetryConfig;
import com.pulumi.gcp.cloudtasks.outputs.QueueStackdriverLoggingConfig;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A named resource to which messages are sent by publishers.
 * 
 * &gt; **Warning:** This resource requires an App Engine application to be created on the
 * project you&#39;re provisioning it on. If you haven&#39;t already enabled it, you
 * can create a `gcp.appengine.Application` resource to do so. This
 * resource&#39;s location will be the same as the App Engine location specified.
 * 
 * ## Example Usage
 * ### Queue Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.cloudtasks.Queue;
 * import com.pulumi.gcp.cloudtasks.QueueArgs;
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
 *         var default_ = new Queue(&#34;default&#34;, QueueArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Cloud Tasks Queue Advanced
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.cloudtasks.Queue;
 * import com.pulumi.gcp.cloudtasks.QueueArgs;
 * import com.pulumi.gcp.cloudtasks.inputs.QueueAppEngineRoutingOverrideArgs;
 * import com.pulumi.gcp.cloudtasks.inputs.QueueRateLimitsArgs;
 * import com.pulumi.gcp.cloudtasks.inputs.QueueRetryConfigArgs;
 * import com.pulumi.gcp.cloudtasks.inputs.QueueStackdriverLoggingConfigArgs;
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
 *         var advancedConfiguration = new Queue(&#34;advancedConfiguration&#34;, QueueArgs.builder()        
 *             .appEngineRoutingOverride(QueueAppEngineRoutingOverrideArgs.builder()
 *                 .instance(&#34;test&#34;)
 *                 .service(&#34;worker&#34;)
 *                 .version(&#34;1.0&#34;)
 *                 .build())
 *             .location(&#34;us-central1&#34;)
 *             .rateLimits(QueueRateLimitsArgs.builder()
 *                 .maxConcurrentDispatches(3)
 *                 .maxDispatchesPerSecond(2)
 *                 .build())
 *             .retryConfig(QueueRetryConfigArgs.builder()
 *                 .maxAttempts(5)
 *                 .maxBackoff(&#34;3s&#34;)
 *                 .maxDoublings(1)
 *                 .maxRetryDuration(&#34;4s&#34;)
 *                 .minBackoff(&#34;2s&#34;)
 *                 .build())
 *             .stackdriverLoggingConfig(QueueStackdriverLoggingConfigArgs.builder()
 *                 .samplingRatio(0.9)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Queue can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:cloudtasks/queue:Queue default projects/{{project}}/locations/{{location}}/queues/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:cloudtasks/queue:Queue default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:cloudtasks/queue:Queue default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:cloudtasks/queue:Queue")
public class Queue extends com.pulumi.resources.CustomResource {
    /**
     * Overrides for task-level appEngineRouting. These settings apply only
     * to App Engine tasks in this queue
     * Structure is documented below.
     * 
     */
    @Export(name="appEngineRoutingOverride", type=QueueAppEngineRoutingOverride.class, parameters={})
    private Output</* @Nullable */ QueueAppEngineRoutingOverride> appEngineRoutingOverride;

    /**
     * @return Overrides for task-level appEngineRouting. These settings apply only
     * to App Engine tasks in this queue
     * Structure is documented below.
     * 
     */
    public Output<Optional<QueueAppEngineRoutingOverride>> appEngineRoutingOverride() {
        return Codegen.optional(this.appEngineRoutingOverride);
    }
    /**
     * The location of the queue
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return The location of the queue
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The queue name.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The queue name.
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
     * Rate limits for task dispatches.
     * The queue&#39;s actual dispatch rate is the result of:
     * * Number of tasks in the queue
     * * User-specified throttling: rateLimits, retryConfig, and the queue&#39;s state.
     * * System throttling due to 429 (Too Many Requests) or 503 (Service
     *   Unavailable) responses from the worker, high error rates, or to
     *   smooth sudden large traffic spikes.
     *   Structure is documented below.
     * 
     */
    @Export(name="rateLimits", type=QueueRateLimits.class, parameters={})
    private Output<QueueRateLimits> rateLimits;

    /**
     * @return Rate limits for task dispatches.
     * The queue&#39;s actual dispatch rate is the result of:
     * * Number of tasks in the queue
     * * User-specified throttling: rateLimits, retryConfig, and the queue&#39;s state.
     * * System throttling due to 429 (Too Many Requests) or 503 (Service
     *   Unavailable) responses from the worker, high error rates, or to
     *   smooth sudden large traffic spikes.
     *   Structure is documented below.
     * 
     */
    public Output<QueueRateLimits> rateLimits() {
        return this.rateLimits;
    }
    /**
     * Settings that determine the retry behavior.
     * Structure is documented below.
     * 
     */
    @Export(name="retryConfig", type=QueueRetryConfig.class, parameters={})
    private Output<QueueRetryConfig> retryConfig;

    /**
     * @return Settings that determine the retry behavior.
     * Structure is documented below.
     * 
     */
    public Output<QueueRetryConfig> retryConfig() {
        return this.retryConfig;
    }
    /**
     * Configuration options for writing logs to Stackdriver Logging.
     * Structure is documented below.
     * 
     */
    @Export(name="stackdriverLoggingConfig", type=QueueStackdriverLoggingConfig.class, parameters={})
    private Output</* @Nullable */ QueueStackdriverLoggingConfig> stackdriverLoggingConfig;

    /**
     * @return Configuration options for writing logs to Stackdriver Logging.
     * Structure is documented below.
     * 
     */
    public Output<Optional<QueueStackdriverLoggingConfig>> stackdriverLoggingConfig() {
        return Codegen.optional(this.stackdriverLoggingConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Queue(String name) {
        this(name, QueueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Queue(String name, QueueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Queue(String name, QueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudtasks/queue:Queue", name, args == null ? QueueArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Queue(String name, Output<String> id, @Nullable QueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudtasks/queue:Queue", name, state, makeResourceOptions(options, id));
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
    public static Queue get(String name, Output<String> id, @Nullable QueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Queue(name, id, state, options);
    }
}

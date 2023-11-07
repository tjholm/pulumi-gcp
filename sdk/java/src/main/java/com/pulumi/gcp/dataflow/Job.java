// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataflow;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataflow.JobArgs;
import com.pulumi.gcp.dataflow.inputs.JobState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a job on Dataflow, which is an implementation of Apache Beam running on Google Compute Engine. For more information see
 * the official documentation for
 * [Beam](https://beam.apache.org) and [Dataflow](https://cloud.google.com/dataflow/).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataflow.Job;
 * import com.pulumi.gcp.dataflow.JobArgs;
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
 *         var bigDataJob = new Job(&#34;bigDataJob&#34;, JobArgs.builder()        
 *             .parameters(Map.ofEntries(
 *                 Map.entry(&#34;baz&#34;, &#34;qux&#34;),
 *                 Map.entry(&#34;foo&#34;, &#34;bar&#34;)
 *             ))
 *             .tempGcsLocation(&#34;gs://my-bucket/tmp_dir&#34;)
 *             .templateGcsPath(&#34;gs://my-bucket/templates/template_file&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Streaming Job
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.pubsub.Topic;
 * import com.pulumi.gcp.storage.Bucket;
 * import com.pulumi.gcp.storage.BucketArgs;
 * import com.pulumi.gcp.dataflow.Job;
 * import com.pulumi.gcp.dataflow.JobArgs;
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
 *         var topic = new Topic(&#34;topic&#34;);
 * 
 *         var bucket1 = new Bucket(&#34;bucket1&#34;, BucketArgs.builder()        
 *             .location(&#34;US&#34;)
 *             .forceDestroy(true)
 *             .build());
 * 
 *         var bucket2 = new Bucket(&#34;bucket2&#34;, BucketArgs.builder()        
 *             .location(&#34;US&#34;)
 *             .forceDestroy(true)
 *             .build());
 * 
 *         var pubsubStream = new Job(&#34;pubsubStream&#34;, JobArgs.builder()        
 *             .templateGcsPath(&#34;gs://my-bucket/templates/template_file&#34;)
 *             .tempGcsLocation(&#34;gs://my-bucket/tmp_dir&#34;)
 *             .enableStreamingEngine(true)
 *             .parameters(Map.ofEntries(
 *                 Map.entry(&#34;inputFilePattern&#34;, bucket1.url().applyValue(url -&gt; String.format(&#34;%s/*.json&#34;, url))),
 *                 Map.entry(&#34;outputTopic&#34;, topic.id())
 *             ))
 *             .transformNameMapping(Map.ofEntries(
 *                 Map.entry(&#34;name&#34;, &#34;test_job&#34;),
 *                 Map.entry(&#34;env&#34;, &#34;test&#34;)
 *             ))
 *             .onDelete(&#34;cancel&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Note on &#34;destroy&#34; / &#34;apply&#34;
 * 
 * There are many types of Dataflow jobs.  Some Dataflow jobs run constantly, getting new data from (e.g.) a GCS bucket, and outputting data continuously.  Some jobs process a set amount of data then terminate.  All jobs can fail while running due to programming errors or other issues.  In this way, Dataflow jobs are different from most other Google resources.
 * 
 * The Dataflow resource is considered &#39;existing&#39; while it is in a nonterminal state.  If it reaches a terminal state (e.g. &#39;FAILED&#39;, &#39;COMPLETE&#39;, &#39;CANCELLED&#39;), it will be recreated on the next &#39;apply&#39;.  This is as expected for jobs which run continuously, but may surprise users who use this resource for other kinds of Dataflow jobs.
 * 
 * A Dataflow job which is &#39;destroyed&#39; may be &#34;cancelled&#34; or &#34;drained&#34;.  If &#34;cancelled&#34;, the job terminates - any data written remains where it is, but no new data will be processed.  If &#34;drained&#34;, no new data will enter the pipeline, but any data currently in the pipeline will finish being processed.  The default is &#34;drain&#34;. When `on_delete` is set to `&#34;drain&#34;` in the configuration, you may experience a long wait for your `pulumi destroy` to complete.
 * 
 * You can potentially short-circuit the wait by setting `skip_wait_on_job_termination` to `true`, but beware that unless you take active steps to ensure that the job `name` parameter changes between instances, the name will conflict and the launch of the new job will fail. One way to do this is with a random_id resource, for example:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomId;
 * import com.pulumi.random.RandomIdArgs;
 * import com.pulumi.gcp.dataflow.FlexTemplateJob;
 * import com.pulumi.gcp.dataflow.FlexTemplateJobArgs;
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
 *         final var config = ctx.config();
 *         final var bigDataJobSubscriptionId = config.get(&#34;bigDataJobSubscriptionId&#34;).orElse(&#34;projects/myproject/subscriptions/messages&#34;);
 *         var bigDataJobNameSuffix = new RandomId(&#34;bigDataJobNameSuffix&#34;, RandomIdArgs.builder()        
 *             .byteLength(4)
 *             .keepers(Map.ofEntries(
 *                 Map.entry(&#34;region&#34;, var_.region()),
 *                 Map.entry(&#34;subscription_id&#34;, bigDataJobSubscriptionId)
 *             ))
 *             .build());
 * 
 *         var bigDataJob = new FlexTemplateJob(&#34;bigDataJob&#34;, FlexTemplateJobArgs.builder()        
 *             .region(var_.region())
 *             .containerSpecGcsPath(&#34;gs://my-bucket/templates/template.json&#34;)
 *             .skipWaitOnJobTermination(true)
 *             .parameters(Map.of(&#34;inputSubscription&#34;, bigDataJobSubscriptionId))
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
 * Dataflow jobs can be imported using the job `id` e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:dataflow/job:Job example 2022-07-31_06_25_42-11926927532632678660
 * ```
 * 
 */
@ResourceType(type="gcp:dataflow/job:Job")
public class Job extends com.pulumi.resources.CustomResource {
    /**
     * List of experiments that should be used by the job. An example value is `[&#34;enable_stackdriver_agent_metrics&#34;]`.
     * 
     */
    @Export(name="additionalExperiments", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> additionalExperiments;

    /**
     * @return List of experiments that should be used by the job. An example value is `[&#34;enable_stackdriver_agent_metrics&#34;]`.
     * 
     */
    public Output<List<String>> additionalExperiments() {
        return this.additionalExperiments;
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Enable/disable the use of [Streaming Engine](https://cloud.google.com/dataflow/docs/guides/deploying-a-pipeline#streaming-engine) for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
     * 
     */
    @Export(name="enableStreamingEngine", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableStreamingEngine;

    /**
     * @return Enable/disable the use of [Streaming Engine](https://cloud.google.com/dataflow/docs/guides/deploying-a-pipeline#streaming-engine) for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
     * 
     */
    public Output<Optional<Boolean>> enableStreamingEngine() {
        return Codegen.optional(this.enableStreamingEngine);
    }
    /**
     * The configuration for VM IPs.  Options are `&#34;WORKER_IP_PUBLIC&#34;` or `&#34;WORKER_IP_PRIVATE&#34;`.
     * 
     */
    @Export(name="ipConfiguration", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipConfiguration;

    /**
     * @return The configuration for VM IPs.  Options are `&#34;WORKER_IP_PUBLIC&#34;` or `&#34;WORKER_IP_PRIVATE&#34;`.
     * 
     */
    public Output<Optional<String>> ipConfiguration() {
        return Codegen.optional(this.ipConfiguration);
    }
    /**
     * The unique ID of this job.
     * 
     */
    @Export(name="jobId", refs={String.class}, tree="[0]")
    private Output<String> jobId;

    /**
     * @return The unique ID of this job.
     * 
     */
    public Output<String> jobId() {
        return this.jobId;
    }
    /**
     * The name for the Cloud KMS key for the job. Key format is: `projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`
     * 
     */
    @Export(name="kmsKeyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyName;

    /**
     * @return The name for the Cloud KMS key for the job. Key format is: `projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`
     * 
     */
    public Output<Optional<String>> kmsKeyName() {
        return Codegen.optional(this.kmsKeyName);
    }
    /**
     * User labels to be specified for the job. Keys and values should follow the restrictions
     * specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> labels;

    /**
     * @return User labels to be specified for the job. Keys and values should follow the restrictions
     * specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The machine type to use for the job.
     * 
     */
    @Export(name="machineType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> machineType;

    /**
     * @return The machine type to use for the job.
     * 
     */
    public Output<Optional<String>> machineType() {
        return Codegen.optional(this.machineType);
    }
    /**
     * The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
     * 
     */
    @Export(name="maxWorkers", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxWorkers;

    /**
     * @return The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
     * 
     */
    public Output<Optional<Integer>> maxWorkers() {
        return Codegen.optional(this.maxWorkers);
    }
    /**
     * A unique name for the resource, required by Dataflow.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the resource, required by Dataflow.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The network to which VMs will be assigned. If it is not provided, &#34;default&#34; will be used.
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> network;

    /**
     * @return The network to which VMs will be assigned. If it is not provided, &#34;default&#34; will be used.
     * 
     */
    public Output<Optional<String>> network() {
        return Codegen.optional(this.network);
    }
    /**
     * One of &#34;drain&#34; or &#34;cancel&#34;.  Specifies behavior of deletion during `pulumi destroy`.  See above note.
     * 
     */
    @Export(name="onDelete", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> onDelete;

    /**
     * @return One of &#34;drain&#34; or &#34;cancel&#34;.  Specifies behavior of deletion during `pulumi destroy`.  See above note.
     * 
     */
    public Output<Optional<String>> onDelete() {
        return Codegen.optional(this.onDelete);
    }
    /**
     * Key/Value pairs to be passed to the Dataflow job (as used in the template).
     * 
     */
    @Export(name="parameters", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> parameters;

    /**
     * @return Key/Value pairs to be passed to the Dataflow job (as used in the template).
     * 
     */
    public Output<Optional<Map<String,Object>>> parameters() {
        return Codegen.optional(this.parameters);
    }
    /**
     * The project in which the resource belongs. If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The project in which the resource belongs. If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * The region in which the created job should run.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> region;

    /**
     * @return The region in which the created job should run.
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * The Service Account email used to create the job.
     * 
     */
    @Export(name="serviceAccountEmail", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serviceAccountEmail;

    /**
     * @return The Service Account email used to create the job.
     * 
     */
    public Output<Optional<String>> serviceAccountEmail() {
        return Codegen.optional(this.serviceAccountEmail);
    }
    /**
     * If set to `true`, Pulumi will treat `DRAINING` and `CANCELLING` as terminal states when deleting the resource, and will remove the resource from Pulumi state and move on.  See above note.
     * 
     */
    @Export(name="skipWaitOnJobTermination", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> skipWaitOnJobTermination;

    /**
     * @return If set to `true`, Pulumi will treat `DRAINING` and `CANCELLING` as terminal states when deleting the resource, and will remove the resource from Pulumi state and move on.  See above note.
     * 
     */
    public Output<Optional<Boolean>> skipWaitOnJobTermination() {
        return Codegen.optional(this.skipWaitOnJobTermination);
    }
    /**
     * The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The subnetwork to which VMs will be assigned. Should be of the form &#34;regions/REGION/subnetworks/SUBNETWORK&#34;. If the [subnetwork is located in a Shared VPC network](https://cloud.google.com/dataflow/docs/guides/specifying-networks#shared), you must use the complete URL. For example `&#34;googleapis.com/compute/v1/projects/PROJECT_ID/regions/REGION/subnetworks/SUBNET_NAME&#34;`
     * 
     */
    @Export(name="subnetwork", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> subnetwork;

    /**
     * @return The subnetwork to which VMs will be assigned. Should be of the form &#34;regions/REGION/subnetworks/SUBNETWORK&#34;. If the [subnetwork is located in a Shared VPC network](https://cloud.google.com/dataflow/docs/guides/specifying-networks#shared), you must use the complete URL. For example `&#34;googleapis.com/compute/v1/projects/PROJECT_ID/regions/REGION/subnetworks/SUBNET_NAME&#34;`
     * 
     */
    public Output<Optional<String>> subnetwork() {
        return Codegen.optional(this.subnetwork);
    }
    /**
     * A writeable location on GCS for the Dataflow job to dump its temporary data.
     * 
     * ***
     * 
     */
    @Export(name="tempGcsLocation", refs={String.class}, tree="[0]")
    private Output<String> tempGcsLocation;

    /**
     * @return A writeable location on GCS for the Dataflow job to dump its temporary data.
     * 
     * ***
     * 
     */
    public Output<String> tempGcsLocation() {
        return this.tempGcsLocation;
    }
    /**
     * The GCS path to the Dataflow job template.
     * 
     */
    @Export(name="templateGcsPath", refs={String.class}, tree="[0]")
    private Output<String> templateGcsPath;

    /**
     * @return The GCS path to the Dataflow job template.
     * 
     */
    public Output<String> templateGcsPath() {
        return this.templateGcsPath;
    }
    /**
     * Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
     * 
     */
    @Export(name="transformNameMapping", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> transformNameMapping;

    /**
     * @return Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
     * 
     */
    public Output<Optional<Map<String,Object>>> transformNameMapping() {
        return Codegen.optional(this.transformNameMapping);
    }
    /**
     * The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The zone in which the created job should run. If it is not provided, the provider zone is used.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> zone;

    /**
     * @return The zone in which the created job should run. If it is not provided, the provider zone is used.
     * 
     */
    public Output<Optional<String>> zone() {
        return Codegen.optional(this.zone);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Job(String name) {
        this(name, JobArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Job(String name, JobArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Job(String name, JobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataflow/job:Job", name, args == null ? JobArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Job(String name, Output<String> id, @Nullable JobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataflow/job:Job", name, state, makeResourceOptions(options, id));
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
    public static Job get(String name, Output<String> id, @Nullable JobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Job(name, id, state, options);
    }
}

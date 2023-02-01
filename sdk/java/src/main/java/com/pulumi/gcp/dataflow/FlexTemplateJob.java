// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataflow;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataflow.FlexTemplateJobArgs;
import com.pulumi.gcp.dataflow.inputs.FlexTemplateJobState;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * This resource does not support import.
 * 
 */
@ResourceType(type="gcp:dataflow/flexTemplateJob:FlexTemplateJob")
public class FlexTemplateJob extends com.pulumi.resources.CustomResource {
    /**
     * The GCS path to the Dataflow job Flex
     * Template.
     * 
     */
    @Export(name="containerSpecGcsPath", type=String.class, parameters={})
    private Output<String> containerSpecGcsPath;

    /**
     * @return The GCS path to the Dataflow job Flex
     * Template.
     * 
     */
    public Output<String> containerSpecGcsPath() {
        return this.containerSpecGcsPath;
    }
    /**
     * The unique ID of this job.
     * 
     */
    @Export(name="jobId", type=String.class, parameters={})
    private Output<String> jobId;

    /**
     * @return The unique ID of this job.
     * 
     */
    public Output<String> jobId() {
        return this.jobId;
    }
    /**
     * User labels to be specified for the job. Keys and values
     * should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
     * page.
     * **NOTE**: Google-provided Dataflow templates often provide default labels
     * that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
     * labels will be ignored to prevent diffs on re-apply.
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> labels;

    /**
     * @return User labels to be specified for the job. Keys and values
     * should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
     * page.
     * **NOTE**: Google-provided Dataflow templates often provide default labels
     * that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
     * labels will be ignored to prevent diffs on re-apply.
     * 
     */
    public Output<Optional<Map<String,Object>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * A unique name for the resource, required by Dataflow.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return A unique name for the resource, required by Dataflow.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="onDelete", type=String.class, parameters={})
    private Output</* @Nullable */ String> onDelete;

    public Output<Optional<String>> onDelete() {
        return Codegen.optional(this.onDelete);
    }
    /**
     * Key/Value pairs to be passed to the Dataflow job (as
     * used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
     * such as `serviceAccount`, `workerMachineType`, etc can be specified here.
     * 
     */
    @Export(name="parameters", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> parameters;

    /**
     * @return Key/Value pairs to be passed to the Dataflow job (as
     * used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
     * such as `serviceAccount`, `workerMachineType`, etc can be specified here.
     * 
     */
    public Output<Optional<Map<String,Object>>> parameters() {
        return Codegen.optional(this.parameters);
    }
    /**
     * The project in which the resource belongs. If it is not
     * provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The project in which the resource belongs. If it is not
     * provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The region in which the created job should run.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which the created job should run.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * If true, treat DRAINING and CANCELLING as terminal job states and do not wait for further changes before removing from
     * terraform state and moving on. WARNING: this will lead to job name conflicts if you do not ensure that the job names are
     * different, e.g. by embedding a release ID or by using a random_id.
     * 
     */
    @Export(name="skipWaitOnJobTermination", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> skipWaitOnJobTermination;

    /**
     * @return If true, treat DRAINING and CANCELLING as terminal job states and do not wait for further changes before removing from
     * terraform state and moving on. WARNING: this will lead to job name conflicts if you do not ensure that the job names are
     * different, e.g. by embedding a release ID or by using a random_id.
     * 
     */
    public Output<Optional<Boolean>> skipWaitOnJobTermination() {
        return Codegen.optional(this.skipWaitOnJobTermination);
    }
    /**
     * The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
     * 
     */
    public Output<String> state() {
        return this.state;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FlexTemplateJob(String name) {
        this(name, FlexTemplateJobArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FlexTemplateJob(String name, FlexTemplateJobArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FlexTemplateJob(String name, FlexTemplateJobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataflow/flexTemplateJob:FlexTemplateJob", name, args == null ? FlexTemplateJobArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FlexTemplateJob(String name, Output<String> id, @Nullable FlexTemplateJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataflow/flexTemplateJob:FlexTemplateJob", name, state, makeResourceOptions(options, id));
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
    public static FlexTemplateJob get(String name, Output<String> id, @Nullable FlexTemplateJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FlexTemplateJob(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.clouddeploy;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.clouddeploy.TargetArgs;
import com.pulumi.gcp.clouddeploy.inputs.TargetState;
import com.pulumi.gcp.clouddeploy.outputs.TargetAnthosCluster;
import com.pulumi.gcp.clouddeploy.outputs.TargetCustomTarget;
import com.pulumi.gcp.clouddeploy.outputs.TargetExecutionConfig;
import com.pulumi.gcp.clouddeploy.outputs.TargetGke;
import com.pulumi.gcp.clouddeploy.outputs.TargetMultiTarget;
import com.pulumi.gcp.clouddeploy.outputs.TargetRun;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The Cloud Deploy `Target` resource
 * 
 * ## Example Usage
 * 
 * ### Multi_target
 * tests creating and updating a multi-target
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.clouddeploy.Target;
 * import com.pulumi.gcp.clouddeploy.TargetArgs;
 * import com.pulumi.gcp.clouddeploy.inputs.TargetExecutionConfigArgs;
 * import com.pulumi.gcp.clouddeploy.inputs.TargetMultiTargetArgs;
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
 *         var primary = new Target("primary", TargetArgs.builder()
 *             .location("us-west1")
 *             .name("target")
 *             .deployParameters()
 *             .description("multi-target description")
 *             .executionConfigs(TargetExecutionConfigArgs.builder()
 *                 .usages(                
 *                     "RENDER",
 *                     "DEPLOY")
 *                 .executionTimeout("3600s")
 *                 .build())
 *             .multiTarget(TargetMultiTargetArgs.builder()
 *                 .targetIds(                
 *                     "1",
 *                     "2")
 *                 .build())
 *             .project("my-project-name")
 *             .requireApproval(false)
 *             .annotations(Map.ofEntries(
 *                 Map.entry("my_first_annotation", "example-annotation-1"),
 *                 Map.entry("my_second_annotation", "example-annotation-2")
 *             ))
 *             .labels(Map.ofEntries(
 *                 Map.entry("my_first_label", "example-label-1"),
 *                 Map.entry("my_second_label", "example-label-2")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Run_target
 * tests creating and updating a cloud run target
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.clouddeploy.Target;
 * import com.pulumi.gcp.clouddeploy.TargetArgs;
 * import com.pulumi.gcp.clouddeploy.inputs.TargetExecutionConfigArgs;
 * import com.pulumi.gcp.clouddeploy.inputs.TargetRunArgs;
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
 *         var primary = new Target("primary", TargetArgs.builder()
 *             .location("us-west1")
 *             .name("target")
 *             .deployParameters()
 *             .description("basic description")
 *             .executionConfigs(TargetExecutionConfigArgs.builder()
 *                 .usages(                
 *                     "RENDER",
 *                     "DEPLOY")
 *                 .executionTimeout("3600s")
 *                 .build())
 *             .project("my-project-name")
 *             .requireApproval(false)
 *             .run(TargetRunArgs.builder()
 *                 .location("projects/my-project-name/locations/us-west1")
 *                 .build())
 *             .annotations(Map.ofEntries(
 *                 Map.entry("my_first_annotation", "example-annotation-1"),
 *                 Map.entry("my_second_annotation", "example-annotation-2")
 *             ))
 *             .labels(Map.ofEntries(
 *                 Map.entry("my_first_label", "example-label-1"),
 *                 Map.entry("my_second_label", "example-label-2")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Target
 * Creates a basic Cloud Deploy target
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.clouddeploy.Target;
 * import com.pulumi.gcp.clouddeploy.TargetArgs;
 * import com.pulumi.gcp.clouddeploy.inputs.TargetGkeArgs;
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
 *         var primary = new Target("primary", TargetArgs.builder()
 *             .location("us-west1")
 *             .name("target")
 *             .deployParameters(Map.of("deployParameterKey", "deployParameterValue"))
 *             .description("basic description")
 *             .gke(TargetGkeArgs.builder()
 *                 .cluster("projects/my-project-name/locations/us-west1/clusters/example-cluster-name")
 *                 .build())
 *             .project("my-project-name")
 *             .requireApproval(false)
 *             .annotations(Map.ofEntries(
 *                 Map.entry("my_first_annotation", "example-annotation-1"),
 *                 Map.entry("my_second_annotation", "example-annotation-2")
 *             ))
 *             .labels(Map.ofEntries(
 *                 Map.entry("my_first_label", "example-label-1"),
 *                 Map.entry("my_second_label", "example-label-2")
 *             ))
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
 * Target can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/targets/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{name}}`
 * 
 * * `{{location}}/{{name}}`
 * 
 * When using the `pulumi import` command, Target can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:clouddeploy/target:Target default projects/{{project}}/locations/{{location}}/targets/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:clouddeploy/target:Target default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:clouddeploy/target:Target default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:clouddeploy/target:Target")
public class Target extends com.pulumi.resources.CustomResource {
    /**
     * Optional. User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     * 
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    @Export(name="annotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> annotations;

    /**
     * @return Optional. User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     * 
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> annotations() {
        return Codegen.optional(this.annotations);
    }
    /**
     * Information specifying an Anthos Cluster.
     * 
     */
    @Export(name="anthosCluster", refs={TargetAnthosCluster.class}, tree="[0]")
    private Output</* @Nullable */ TargetAnthosCluster> anthosCluster;

    /**
     * @return Information specifying an Anthos Cluster.
     * 
     */
    public Output<Optional<TargetAnthosCluster>> anthosCluster() {
        return Codegen.optional(this.anthosCluster);
    }
    /**
     * Output only. Time at which the `Target` was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. Time at which the `Target` was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Optional. Information specifying a Custom Target.
     * 
     */
    @Export(name="customTarget", refs={TargetCustomTarget.class}, tree="[0]")
    private Output</* @Nullable */ TargetCustomTarget> customTarget;

    /**
     * @return Optional. Information specifying a Custom Target.
     * 
     */
    public Output<Optional<TargetCustomTarget>> customTarget() {
        return Codegen.optional(this.customTarget);
    }
    /**
     * Optional. The deploy parameters to use for this target.
     * 
     */
    @Export(name="deployParameters", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> deployParameters;

    /**
     * @return Optional. The deploy parameters to use for this target.
     * 
     */
    public Output<Optional<Map<String,String>>> deployParameters() {
        return Codegen.optional(this.deployParameters);
    }
    /**
     * Optional. Description of the `Target`. Max length is 255 characters.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Optional. Description of the `Target`. Max length is 255 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="effectiveAnnotations", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> effectiveAnnotations;

    public Output<Map<String,Object>> effectiveAnnotations() {
        return this.effectiveAnnotations;
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,Object>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Optional. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return Optional. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Configurations for all execution that relates to this `Target`. Each `ExecutionEnvironmentUsage` value may only be used in a single configuration; using the same value multiple times is an error. When one or more configurations are specified, they must include the `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values. When no configurations are specified, execution will use the default specified in `DefaultPool`.
     * 
     */
    @Export(name="executionConfigs", refs={List.class,TargetExecutionConfig.class}, tree="[0,1]")
    private Output<List<TargetExecutionConfig>> executionConfigs;

    /**
     * @return Configurations for all execution that relates to this `Target`. Each `ExecutionEnvironmentUsage` value may only be used in a single configuration; using the same value multiple times is an error. When one or more configurations are specified, they must include the `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values. When no configurations are specified, execution will use the default specified in `DefaultPool`.
     * 
     */
    public Output<List<TargetExecutionConfig>> executionConfigs() {
        return this.executionConfigs;
    }
    /**
     * Information specifying a GKE Cluster.
     * 
     */
    @Export(name="gke", refs={TargetGke.class}, tree="[0]")
    private Output</* @Nullable */ TargetGke> gke;

    /**
     * @return Information specifying a GKE Cluster.
     * 
     */
    public Output<Optional<TargetGke>> gke() {
        return Codegen.optional(this.gke);
    }
    /**
     * Optional. Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be &lt;= 128 bytes.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Optional. Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be &lt;= 128 bytes.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location for the resource
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location for the resource
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Information specifying a multiTarget.
     * 
     */
    @Export(name="multiTarget", refs={TargetMultiTarget.class}, tree="[0]")
    private Output</* @Nullable */ TargetMultiTarget> multiTarget;

    /**
     * @return Information specifying a multiTarget.
     * 
     */
    public Output<Optional<TargetMultiTarget>> multiTarget() {
        return Codegen.optional(this.multiTarget);
    }
    /**
     * Name of the `Target`. Format is `a-z?`.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the `Target`. Format is `a-z?`.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The project for the resource
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The project for the resource
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    public Output<Map<String,Object>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * Optional. Whether or not the `Target` requires approval.
     * 
     */
    @Export(name="requireApproval", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> requireApproval;

    /**
     * @return Optional. Whether or not the `Target` requires approval.
     * 
     */
    public Output<Optional<Boolean>> requireApproval() {
        return Codegen.optional(this.requireApproval);
    }
    /**
     * Information specifying a Cloud Run deployment target.
     * 
     */
    @Export(name="run", refs={TargetRun.class}, tree="[0]")
    private Output</* @Nullable */ TargetRun> run;

    /**
     * @return Information specifying a Cloud Run deployment target.
     * 
     */
    public Output<Optional<TargetRun>> run() {
        return Codegen.optional(this.run);
    }
    /**
     * Output only. Resource id of the `Target`.
     * 
     */
    @Export(name="targetId", refs={String.class}, tree="[0]")
    private Output<String> targetId;

    /**
     * @return Output only. Resource id of the `Target`.
     * 
     */
    public Output<String> targetId() {
        return this.targetId;
    }
    /**
     * Output only. Unique identifier of the `Target`.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return Output only. Unique identifier of the `Target`.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Output only. Most recent time at which the `Target` was updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. Most recent time at which the `Target` was updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Target(java.lang.String name) {
        this(name, TargetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Target(java.lang.String name, TargetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Target(java.lang.String name, TargetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:clouddeploy/target:Target", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Target(java.lang.String name, Output<java.lang.String> id, @Nullable TargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:clouddeploy/target:Target", name, state, makeResourceOptions(options, id), false);
    }

    private static TargetArgs makeArgs(TargetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TargetArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Target get(java.lang.String name, Output<java.lang.String> id, @Nullable TargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Target(name, id, state, options);
    }
}

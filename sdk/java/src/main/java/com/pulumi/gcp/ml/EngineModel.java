// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.ml;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.ml.EngineModelArgs;
import com.pulumi.gcp.ml.inputs.EngineModelState;
import com.pulumi.gcp.ml.outputs.EngineModelDefaultVersion;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a machine learning solution.
 * 
 * A model can have multiple versions, each of which is a deployed, trained model
 * ready to receive prediction requests. The model itself is just a container.
 * 
 * To get more information about Model, see:
 * 
 * * [API documentation](https://cloud.google.com/ai-platform/prediction/docs/reference/rest/v1/projects.models)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/ai-platform/prediction/docs/deploying-models)
 * 
 * ## Example Usage
 * 
 * ### Ml Model Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.ml.EngineModel;
 * import com.pulumi.gcp.ml.EngineModelArgs;
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
 *         var default_ = new EngineModel("default", EngineModelArgs.builder()
 *             .name("default")
 *             .description("My model")
 *             .regions("us-central1")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Ml Model Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.ml.EngineModel;
 * import com.pulumi.gcp.ml.EngineModelArgs;
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
 *         var default_ = new EngineModel("default", EngineModelArgs.builder()
 *             .name("default")
 *             .description("My model")
 *             .regions("us-central1")
 *             .labels(Map.of("my_model", "foo"))
 *             .onlinePredictionLogging(true)
 *             .onlinePredictionConsoleLogging(true)
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
 * Model can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/models/{{name}}`
 * 
 * * `{{project}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, Model can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:ml/engineModel:EngineModel default projects/{{project}}/models/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:ml/engineModel:EngineModel default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:ml/engineModel:EngineModel default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:ml/engineModel:EngineModel")
public class EngineModel extends com.pulumi.resources.CustomResource {
    /**
     * The default version of the model. This version will be used to handle
     * prediction requests that do not specify a version.
     * Structure is documented below.
     * 
     */
    @Export(name="defaultVersion", refs={EngineModelDefaultVersion.class}, tree="[0]")
    private Output</* @Nullable */ EngineModelDefaultVersion> defaultVersion;

    /**
     * @return The default version of the model. This version will be used to handle
     * prediction requests that do not specify a version.
     * Structure is documented below.
     * 
     */
    public Output<Optional<EngineModelDefaultVersion>> defaultVersion() {
        return Codegen.optional(this.defaultVersion);
    }
    /**
     * The description specified for the model when it was created.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description specified for the model when it was created.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * One or more labels that you can add, to organize your models.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return One or more labels that you can add, to organize your models.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The name specified for the model.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name specified for the model.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
     * 
     */
    @Export(name="onlinePredictionConsoleLogging", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> onlinePredictionConsoleLogging;

    /**
     * @return If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
     * 
     */
    public Output<Optional<Boolean>> onlinePredictionConsoleLogging() {
        return Codegen.optional(this.onlinePredictionConsoleLogging);
    }
    /**
     * If true, online prediction access logs are sent to StackDriver Logging.
     * 
     */
    @Export(name="onlinePredictionLogging", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> onlinePredictionLogging;

    /**
     * @return If true, online prediction access logs are sent to StackDriver Logging.
     * 
     */
    public Output<Optional<Boolean>> onlinePredictionLogging() {
        return Codegen.optional(this.onlinePredictionLogging);
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
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * The list of regions where the model is going to be deployed.
     * Currently only one region per model is supported
     * 
     */
    @Export(name="regions", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> regions;

    /**
     * @return The list of regions where the model is going to be deployed.
     * Currently only one region per model is supported
     * 
     */
    public Output<Optional<String>> regions() {
        return Codegen.optional(this.regions);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EngineModel(java.lang.String name) {
        this(name, EngineModelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EngineModel(java.lang.String name, @Nullable EngineModelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EngineModel(java.lang.String name, @Nullable EngineModelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:ml/engineModel:EngineModel", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EngineModel(java.lang.String name, Output<java.lang.String> id, @Nullable EngineModelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:ml/engineModel:EngineModel", name, state, makeResourceOptions(options, id), false);
    }

    private static EngineModelArgs makeArgs(@Nullable EngineModelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EngineModelArgs.Empty : args;
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
    public static EngineModel get(java.lang.String name, Output<java.lang.String> id, @Nullable EngineModelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EngineModel(name, id, state, options);
    }
}

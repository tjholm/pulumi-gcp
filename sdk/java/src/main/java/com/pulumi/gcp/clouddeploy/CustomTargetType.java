// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.clouddeploy;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.clouddeploy.CustomTargetTypeArgs;
import com.pulumi.gcp.clouddeploy.inputs.CustomTargetTypeState;
import com.pulumi.gcp.clouddeploy.outputs.CustomTargetTypeCustomActions;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Cloud Deploy `CustomTargetType` defines a type of custom target that can be referenced in a
 * Cloud Deploy `Target` in order to facilitate deploying to other systems besides the supported runtimes.
 * 
 * To get more information about CustomTargetType, see:
 * 
 * * [API documentation](https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.customTargetTypes)
 * * How-to Guides
 *     * [Define and use a custom target type](https://cloud.google.com/deploy/docs/deploy-app-custom-target)
 * 
 * ## Example Usage
 * 
 * ### Clouddeploy Custom Target Type Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.clouddeploy.CustomTargetType;
 * import com.pulumi.gcp.clouddeploy.CustomTargetTypeArgs;
 * import com.pulumi.gcp.clouddeploy.inputs.CustomTargetTypeCustomActionsArgs;
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
 *         var custom_target_type = new CustomTargetType("custom-target-type", CustomTargetTypeArgs.builder()
 *             .location("us-central1")
 *             .name("my-custom-target-type")
 *             .description("My custom target type")
 *             .annotations(Map.ofEntries(
 *                 Map.entry("my_first_annotation", "example-annotation-1"),
 *                 Map.entry("my_second_annotation", "example-annotation-2")
 *             ))
 *             .labels(Map.ofEntries(
 *                 Map.entry("my_first_label", "example-label-1"),
 *                 Map.entry("my_second_label", "example-label-2")
 *             ))
 *             .customActions(CustomTargetTypeCustomActionsArgs.builder()
 *                 .renderAction("renderAction")
 *                 .deployAction("deployAction")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Clouddeploy Custom Target Type Git Skaffold Modules
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.clouddeploy.CustomTargetType;
 * import com.pulumi.gcp.clouddeploy.CustomTargetTypeArgs;
 * import com.pulumi.gcp.clouddeploy.inputs.CustomTargetTypeCustomActionsArgs;
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
 *         var custom_target_type = new CustomTargetType("custom-target-type", CustomTargetTypeArgs.builder()
 *             .location("us-central1")
 *             .name("my-custom-target-type")
 *             .description("My custom target type")
 *             .customActions(CustomTargetTypeCustomActionsArgs.builder()
 *                 .renderAction("renderAction")
 *                 .deployAction("deployAction")
 *                 .includeSkaffoldModules(CustomTargetTypeCustomActionsIncludeSkaffoldModuleArgs.builder()
 *                     .configs("my-config")
 *                     .git(CustomTargetTypeCustomActionsIncludeSkaffoldModuleGitArgs.builder()
 *                         .repo("http://github.com/example/example-repo.git")
 *                         .path("configs/skaffold.yaml")
 *                         .ref("main")
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Clouddeploy Custom Target Type Gcs Skaffold Modules
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.clouddeploy.CustomTargetType;
 * import com.pulumi.gcp.clouddeploy.CustomTargetTypeArgs;
 * import com.pulumi.gcp.clouddeploy.inputs.CustomTargetTypeCustomActionsArgs;
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
 *         var custom_target_type = new CustomTargetType("custom-target-type", CustomTargetTypeArgs.builder()
 *             .location("us-central1")
 *             .name("my-custom-target-type")
 *             .description("My custom target type")
 *             .customActions(CustomTargetTypeCustomActionsArgs.builder()
 *                 .renderAction("renderAction")
 *                 .deployAction("deployAction")
 *                 .includeSkaffoldModules(CustomTargetTypeCustomActionsIncludeSkaffoldModuleArgs.builder()
 *                     .configs("my-config")
 *                     .googleCloudStorage(CustomTargetTypeCustomActionsIncludeSkaffoldModuleGoogleCloudStorageArgs.builder()
 *                         .source("gs://example-bucket/dir/configs/*")
 *                         .path("skaffold.yaml")
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Clouddeploy Custom Target Type Gcb Repo Skaffold Modules
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.clouddeploy.CustomTargetType;
 * import com.pulumi.gcp.clouddeploy.CustomTargetTypeArgs;
 * import com.pulumi.gcp.clouddeploy.inputs.CustomTargetTypeCustomActionsArgs;
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
 *         var custom_target_type = new CustomTargetType("custom-target-type", CustomTargetTypeArgs.builder()
 *             .location("us-central1")
 *             .name("my-custom-target-type")
 *             .description("My custom target type")
 *             .customActions(CustomTargetTypeCustomActionsArgs.builder()
 *                 .renderAction("renderAction")
 *                 .deployAction("deployAction")
 *                 .includeSkaffoldModules(CustomTargetTypeCustomActionsIncludeSkaffoldModuleArgs.builder()
 *                     .configs("my-config")
 *                     .googleCloudBuildRepo(CustomTargetTypeCustomActionsIncludeSkaffoldModuleGoogleCloudBuildRepoArgs.builder()
 *                         .repository("projects/example/locations/us-central1/connections/git/repositories/example-repo")
 *                         .path("configs/skaffold.yaml")
 *                         .ref("main")
 *                         .build())
 *                     .build())
 *                 .build())
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
 * CustomTargetType can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/customTargetTypes/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{name}}`
 * 
 * * `{{location}}/{{name}}`
 * 
 * When using the `pulumi import` command, CustomTargetType can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:clouddeploy/customTargetType:CustomTargetType default projects/{{project}}/locations/{{location}}/customTargetTypes/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:clouddeploy/customTargetType:CustomTargetType default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:clouddeploy/customTargetType:CustomTargetType default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:clouddeploy/customTargetType:CustomTargetType")
public class CustomTargetType extends com.pulumi.resources.CustomResource {
    /**
     * User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    @Export(name="annotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> annotations;

    /**
     * @return User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> annotations() {
        return Codegen.optional(this.annotations);
    }
    /**
     * Time at which the `CustomTargetType` was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Time at which the `CustomTargetType` was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Configures render and deploy for the `CustomTargetType` using Skaffold custom actions.
     * Structure is documented below.
     * 
     */
    @Export(name="customActions", refs={CustomTargetTypeCustomActions.class}, tree="[0]")
    private Output</* @Nullable */ CustomTargetTypeCustomActions> customActions;

    /**
     * @return Configures render and deploy for the `CustomTargetType` using Skaffold custom actions.
     * Structure is documented below.
     * 
     */
    public Output<Optional<CustomTargetTypeCustomActions>> customActions() {
        return Codegen.optional(this.customActions);
    }
    /**
     * Resource id of the `CustomTargetType`.
     * 
     */
    @Export(name="customTargetTypeId", refs={String.class}, tree="[0]")
    private Output<String> customTargetTypeId;

    /**
     * @return Resource id of the `CustomTargetType`.
     * 
     */
    public Output<String> customTargetTypeId() {
        return this.customTargetTypeId;
    }
    /**
     * Description of the `CustomTargetType`. Max length is 255 characters.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the `CustomTargetType`. Max length is 255 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="effectiveAnnotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveAnnotations;

    public Output<Map<String,String>> effectiveAnnotations() {
        return this.effectiveAnnotations;
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
     * The weak etag of the `CustomTargetType` resource. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return The weak etag of the `CustomTargetType` resource. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be &lt;= 128 bytes.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be &lt;= 128 bytes.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location of the source.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location of the source.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Name of the `CustomTargetType`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the `CustomTargetType`.
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
     * Unique identifier of the `CustomTargetType`.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return Unique identifier of the `CustomTargetType`.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Time at which the `CustomTargetType` was updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Time at which the `CustomTargetType` was updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomTargetType(java.lang.String name) {
        this(name, CustomTargetTypeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomTargetType(java.lang.String name, CustomTargetTypeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomTargetType(java.lang.String name, CustomTargetTypeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:clouddeploy/customTargetType:CustomTargetType", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CustomTargetType(java.lang.String name, Output<java.lang.String> id, @Nullable CustomTargetTypeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:clouddeploy/customTargetType:CustomTargetType", name, state, makeResourceOptions(options, id), false);
    }

    private static CustomTargetTypeArgs makeArgs(CustomTargetTypeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CustomTargetTypeArgs.Empty : args;
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
    public static CustomTargetType get(java.lang.String name, Output<java.lang.String> id, @Nullable CustomTargetTypeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomTargetType(name, id, state, options);
    }
}

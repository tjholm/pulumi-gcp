// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataloss.PreventionStoredInfoTypeArgs;
import com.pulumi.gcp.dataloss.inputs.PreventionStoredInfoTypeState;
import com.pulumi.gcp.dataloss.outputs.PreventionStoredInfoTypeDictionary;
import com.pulumi.gcp.dataloss.outputs.PreventionStoredInfoTypeLargeCustomDictionary;
import com.pulumi.gcp.dataloss.outputs.PreventionStoredInfoTypeRegex;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows creation of custom info types.
 * 
 * To get more information about StoredInfoType, see:
 * 
 * * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.storedInfoTypes)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dlp/docs/creating-stored-infotypes)
 * 
 * ## Example Usage
 * 
 * ### Dlp Stored Info Type Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataloss.PreventionStoredInfoType;
 * import com.pulumi.gcp.dataloss.PreventionStoredInfoTypeArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionStoredInfoTypeRegexArgs;
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
 *         var basic = new PreventionStoredInfoType("basic", PreventionStoredInfoTypeArgs.builder()
 *             .parent("projects/my-project-name")
 *             .description("Description")
 *             .displayName("Displayname")
 *             .regex(PreventionStoredInfoTypeRegexArgs.builder()
 *                 .pattern("patient")
 *                 .groupIndexes(2)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Dlp Stored Info Type Dictionary
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataloss.PreventionStoredInfoType;
 * import com.pulumi.gcp.dataloss.PreventionStoredInfoTypeArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionStoredInfoTypeDictionaryArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionStoredInfoTypeDictionaryWordListArgs;
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
 *         var dictionary = new PreventionStoredInfoType("dictionary", PreventionStoredInfoTypeArgs.builder()
 *             .parent("projects/my-project-name")
 *             .description("Description")
 *             .displayName("Displayname")
 *             .dictionary(PreventionStoredInfoTypeDictionaryArgs.builder()
 *                 .wordList(PreventionStoredInfoTypeDictionaryWordListArgs.builder()
 *                     .words(                    
 *                         "word",
 *                         "word2")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Dlp Stored Info Type Large Custom Dictionary
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.storage.Bucket;
 * import com.pulumi.gcp.storage.BucketArgs;
 * import com.pulumi.gcp.storage.BucketObject;
 * import com.pulumi.gcp.storage.BucketObjectArgs;
 * import com.pulumi.gcp.dataloss.PreventionStoredInfoType;
 * import com.pulumi.gcp.dataloss.PreventionStoredInfoTypeArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionStoredInfoTypeLargeCustomDictionaryArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionStoredInfoTypeLargeCustomDictionaryOutputPathArgs;
 * import com.pulumi.asset.FileAsset;
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
 *         var bucket = new Bucket("bucket", BucketArgs.builder()
 *             .name("tf-test-bucket")
 *             .location("US")
 *             .forceDestroy(true)
 *             .build());
 * 
 *         var object = new BucketObject("object", BucketObjectArgs.builder()
 *             .name("tf-test-object")
 *             .bucket(bucket.name())
 *             .source(new FileAsset("./test-fixtures/words.txt"))
 *             .build());
 * 
 *         var large = new PreventionStoredInfoType("large", PreventionStoredInfoTypeArgs.builder()
 *             .parent("projects/my-project-name")
 *             .description("Description")
 *             .displayName("Displayname")
 *             .largeCustomDictionary(PreventionStoredInfoTypeLargeCustomDictionaryArgs.builder()
 *                 .cloudStorageFileSet(PreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetArgs.builder()
 *                     .url(Output.tuple(bucket.name(), object.name()).applyValue(values -> {
 *                         var bucketName = values.t1;
 *                         var objectName = values.t2;
 *                         return String.format("gs://%s/%s", bucketName,objectName);
 *                     }))
 *                     .build())
 *                 .outputPath(PreventionStoredInfoTypeLargeCustomDictionaryOutputPathArgs.builder()
 *                     .path(bucket.name().applyValue(name -> String.format("gs://%s/output/dictionary.txt", name)))
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Dlp Stored Info Type With Id
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataloss.PreventionStoredInfoType;
 * import com.pulumi.gcp.dataloss.PreventionStoredInfoTypeArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionStoredInfoTypeRegexArgs;
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
 *         var withStoredInfoTypeId = new PreventionStoredInfoType("withStoredInfoTypeId", PreventionStoredInfoTypeArgs.builder()
 *             .parent("projects/my-project-name")
 *             .description("Description")
 *             .displayName("Displayname")
 *             .storedInfoTypeId("id-")
 *             .regex(PreventionStoredInfoTypeRegexArgs.builder()
 *                 .pattern("patient")
 *                 .groupIndexes(2)
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
 * StoredInfoType can be imported using any of these accepted formats:
 * 
 * * `{{parent}}/storedInfoTypes/{{name}}`
 * 
 * * `{{parent}}/{{name}}`
 * 
 * When using the `pulumi import` command, StoredInfoType can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType default {{parent}}/storedInfoTypes/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType default {{parent}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType")
public class PreventionStoredInfoType extends com.pulumi.resources.CustomResource {
    /**
     * A description of the info type.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the info type.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Dictionary which defines the rule.
     * Structure is documented below.
     * 
     */
    @Export(name="dictionary", refs={PreventionStoredInfoTypeDictionary.class}, tree="[0]")
    private Output</* @Nullable */ PreventionStoredInfoTypeDictionary> dictionary;

    /**
     * @return Dictionary which defines the rule.
     * Structure is documented below.
     * 
     */
    public Output<Optional<PreventionStoredInfoTypeDictionary>> dictionary() {
        return Codegen.optional(this.dictionary);
    }
    /**
     * User set display name of the info type.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return User set display name of the info type.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Dictionary which defines the rule.
     * Structure is documented below.
     * 
     */
    @Export(name="largeCustomDictionary", refs={PreventionStoredInfoTypeLargeCustomDictionary.class}, tree="[0]")
    private Output</* @Nullable */ PreventionStoredInfoTypeLargeCustomDictionary> largeCustomDictionary;

    /**
     * @return Dictionary which defines the rule.
     * Structure is documented below.
     * 
     */
    public Output<Optional<PreventionStoredInfoTypeLargeCustomDictionary>> largeCustomDictionary() {
        return Codegen.optional(this.largeCustomDictionary);
    }
    /**
     * The resource name of the info type. Set by the server.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the info type. Set by the server.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The parent of the info type in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     * 
     * ***
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return The parent of the info type in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     * 
     * ***
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * Regular expression which defines the rule.
     * Structure is documented below.
     * 
     */
    @Export(name="regex", refs={PreventionStoredInfoTypeRegex.class}, tree="[0]")
    private Output</* @Nullable */ PreventionStoredInfoTypeRegex> regex;

    /**
     * @return Regular expression which defines the rule.
     * Structure is documented below.
     * 
     */
    public Output<Optional<PreventionStoredInfoTypeRegex>> regex() {
        return Codegen.optional(this.regex);
    }
    /**
     * The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens;
     * that is, it must match the regular expression: [a-zA-Z\d-_]+. The maximum length is 100
     * characters. Can be empty to allow the system to generate one.
     * 
     */
    @Export(name="storedInfoTypeId", refs={String.class}, tree="[0]")
    private Output<String> storedInfoTypeId;

    /**
     * @return The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens;
     * that is, it must match the regular expression: [a-zA-Z\d-_]+. The maximum length is 100
     * characters. Can be empty to allow the system to generate one.
     * 
     */
    public Output<String> storedInfoTypeId() {
        return this.storedInfoTypeId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PreventionStoredInfoType(String name) {
        this(name, PreventionStoredInfoTypeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PreventionStoredInfoType(String name, PreventionStoredInfoTypeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PreventionStoredInfoType(String name, PreventionStoredInfoTypeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType", name, args == null ? PreventionStoredInfoTypeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PreventionStoredInfoType(String name, Output<String> id, @Nullable PreventionStoredInfoTypeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType", name, state, makeResourceOptions(options, id));
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
    public static PreventionStoredInfoType get(String name, Output<String> id, @Nullable PreventionStoredInfoTypeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PreventionStoredInfoType(name, id, state, options);
    }
}

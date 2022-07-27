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
 * ### Dlp Stored Info Type Basic
 * ```java
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
 *         var basic = new PreventionStoredInfoType(&#34;basic&#34;, PreventionStoredInfoTypeArgs.builder()        
 *             .description(&#34;Description&#34;)
 *             .displayName(&#34;Displayname&#34;)
 *             .parent(&#34;projects/my-project-name&#34;)
 *             .regex(PreventionStoredInfoTypeRegexArgs.builder()
 *                 .groupIndexes(2)
 *                 .pattern(&#34;patient&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Dlp Stored Info Type Dictionary
 * ```java
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
 *         var dictionary = new PreventionStoredInfoType(&#34;dictionary&#34;, PreventionStoredInfoTypeArgs.builder()        
 *             .description(&#34;Description&#34;)
 *             .dictionary(PreventionStoredInfoTypeDictionaryArgs.builder()
 *                 .wordList(PreventionStoredInfoTypeDictionaryWordListArgs.builder()
 *                     .words(                    
 *                         &#34;word&#34;,
 *                         &#34;word2&#34;)
 *                     .build())
 *                 .build())
 *             .displayName(&#34;Displayname&#34;)
 *             .parent(&#34;projects/my-project-name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Dlp Stored Info Type Large Custom Dictionary
 * ```java
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
 *         var bucket = new Bucket(&#34;bucket&#34;, BucketArgs.builder()        
 *             .location(&#34;US&#34;)
 *             .forceDestroy(true)
 *             .build());
 * 
 *         var object = new BucketObject(&#34;object&#34;, BucketObjectArgs.builder()        
 *             .bucket(bucket.name())
 *             .source(new FileAsset(&#34;./test-fixtures/dlp/words.txt&#34;))
 *             .build());
 * 
 *         var large = new PreventionStoredInfoType(&#34;large&#34;, PreventionStoredInfoTypeArgs.builder()        
 *             .parent(&#34;projects/my-project-name&#34;)
 *             .description(&#34;Description&#34;)
 *             .displayName(&#34;Displayname&#34;)
 *             .largeCustomDictionary(PreventionStoredInfoTypeLargeCustomDictionaryArgs.builder()
 *                 .cloudStorageFileSet(PreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetArgs.builder()
 *                     .url(Output.tuple(bucket.name(), object.name()).applyValue(values -&gt; {
 *                         var bucketName = values.t1;
 *                         var objectName = values.t2;
 *                         return String.format(&#34;gs://%s/%s&#34;, bucketName,objectName);
 *                     }))
 *                     .build())
 *                 .outputPath(PreventionStoredInfoTypeLargeCustomDictionaryOutputPathArgs.builder()
 *                     .path(bucket.name().applyValue(name -&gt; String.format(&#34;gs://%s/output/dictionary.txt&#34;, name)))
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * StoredInfoType can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType default {{parent}}/storedInfoTypes/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType default {{parent}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType")
public class PreventionStoredInfoType extends com.pulumi.resources.CustomResource {
    /**
     * A description of the info type.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
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
    @Export(name="dictionary", type=PreventionStoredInfoTypeDictionary.class, parameters={})
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
    @Export(name="displayName", type=String.class, parameters={})
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
    @Export(name="largeCustomDictionary", type=PreventionStoredInfoTypeLargeCustomDictionary.class, parameters={})
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
     * Name describing the field.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name describing the field.
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
     */
    @Export(name="parent", type=String.class, parameters={})
    private Output<String> parent;

    /**
     * @return The parent of the info type in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
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
    @Export(name="regex", type=PreventionStoredInfoTypeRegex.class, parameters={})
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

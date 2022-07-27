// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firebaserules;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.firebaserules.RulesetArgs;
import com.pulumi.gcp.firebaserules.inputs.RulesetState;
import com.pulumi.gcp.firebaserules.outputs.RulesetMetadata;
import com.pulumi.gcp.firebaserules.outputs.RulesetSource;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * For more information, see:
 * * [Get started with Firebase Security Rules](https://firebase.google.com/docs/rules/get-started)
 * ## Example Usage
 * ### Basic_ruleset
 * Creates a basic Firestore ruleset
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.firebaserules.Ruleset;
 * import com.pulumi.gcp.firebaserules.RulesetArgs;
 * import com.pulumi.gcp.firebaserules.inputs.RulesetSourceArgs;
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
 *         var primary = new Ruleset(&#34;primary&#34;, RulesetArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .source(RulesetSourceArgs.builder()
 *                 .files(RulesetSourceFileArgs.builder()
 *                     .content(&#34;service cloud.firestore {match /databases/{database}/documents { match /{document=**} { allow read, write: if false; } } }&#34;)
 *                     .fingerprint(&#34;&#34;)
 *                     .name(&#34;firestore.rules&#34;)
 *                     .build())
 *                 .language(&#34;&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Minimal_ruleset
 * Creates a minimal Firestore ruleset
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.firebaserules.Ruleset;
 * import com.pulumi.gcp.firebaserules.RulesetArgs;
 * import com.pulumi.gcp.firebaserules.inputs.RulesetSourceArgs;
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
 *         var primary = new Ruleset(&#34;primary&#34;, RulesetArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .source(RulesetSourceArgs.builder()
 *                 .files(RulesetSourceFileArgs.builder()
 *                     .content(&#34;service cloud.firestore {match /databases/{database}/documents { match /{document=**} { allow read, write: if false; } } }&#34;)
 *                     .name(&#34;firestore.rules&#34;)
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
 * Ruleset can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:firebaserules/ruleset:Ruleset default projects/{{project}}/rulesets/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:firebaserules/ruleset:Ruleset default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:firebaserules/ruleset:Ruleset default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:firebaserules/ruleset:Ruleset")
public class Ruleset extends com.pulumi.resources.CustomResource {
    /**
     * Output only. Time the `Ruleset` was created.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return Output only. Time the `Ruleset` was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Output only. The metadata for this ruleset.
     * 
     */
    @Export(name="metadatas", type=List.class, parameters={RulesetMetadata.class})
    private Output<List<RulesetMetadata>> metadatas;

    /**
     * @return Output only. The metadata for this ruleset.
     * 
     */
    public Output<List<RulesetMetadata>> metadatas() {
        return this.metadatas;
    }
    /**
     * File name.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return File name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The project for the resource
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The project for the resource
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * `Source` for the `Ruleset`.
     * 
     */
    @Export(name="source", type=RulesetSource.class, parameters={})
    private Output<RulesetSource> source;

    /**
     * @return `Source` for the `Ruleset`.
     * 
     */
    public Output<RulesetSource> source() {
        return this.source;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ruleset(String name) {
        this(name, RulesetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ruleset(String name, RulesetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ruleset(String name, RulesetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firebaserules/ruleset:Ruleset", name, args == null ? RulesetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Ruleset(String name, Output<String> id, @Nullable RulesetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firebaserules/ruleset:Ruleset", name, state, makeResourceOptions(options, id));
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
    public static Ruleset get(String name, Output<String> id, @Nullable RulesetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ruleset(name, id, state, options);
    }
}

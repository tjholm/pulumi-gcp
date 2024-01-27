// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.essentialcontacts;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.essentialcontacts.DocumentAiProcessorArgs;
import com.pulumi.gcp.essentialcontacts.inputs.DocumentAiProcessorState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The first-class citizen for Document AI. Each processor defines how to extract structural information from a document.
 * 
 * To get more information about Processor, see:
 * 
 * * [API documentation](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations.processors)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/document-ai/docs/overview)
 * 
 * ## Example Usage
 * ### Documentai Processor
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.essentialcontacts.DocumentAiProcessor;
 * import com.pulumi.gcp.essentialcontacts.DocumentAiProcessorArgs;
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
 *         var processor = new DocumentAiProcessor(&#34;processor&#34;, DocumentAiProcessorArgs.builder()        
 *             .displayName(&#34;test-processor&#34;)
 *             .location(&#34;us&#34;)
 *             .type(&#34;OCR_PROCESSOR&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Processor can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/processors/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Processor can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor default projects/{{project}}/locations/{{location}}/processors/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor")
public class DocumentAiProcessor extends com.pulumi.resources.CustomResource {
    /**
     * The display name. Must be unique.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name. Must be unique.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
     * 
     */
    @Export(name="kmsKeyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyName;

    /**
     * @return The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
     * 
     */
    public Output<Optional<String>> kmsKeyName() {
        return Codegen.optional(this.kmsKeyName);
    }
    /**
     * The location of the resource.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location of the resource.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The resource name of the processor.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the processor.
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
     * The type of processor. For possible types see the [official list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of processor. For possible types see the [official list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DocumentAiProcessor(String name) {
        this(name, DocumentAiProcessorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DocumentAiProcessor(String name, DocumentAiProcessorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DocumentAiProcessor(String name, DocumentAiProcessorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor", name, args == null ? DocumentAiProcessorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DocumentAiProcessor(String name, Output<String> id, @Nullable DocumentAiProcessorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor", name, state, makeResourceOptions(options, id));
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
    public static DocumentAiProcessor get(String name, Output<String> id, @Nullable DocumentAiProcessorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DocumentAiProcessor(name, id, state, options);
    }
}

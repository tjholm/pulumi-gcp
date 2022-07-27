// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.datacatalog;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.datacatalog.TagTemplateArgs;
import com.pulumi.gcp.datacatalog.inputs.TagTemplateState;
import com.pulumi.gcp.datacatalog.outputs.TagTemplateField;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A tag template defines a tag, which can have one or more typed fields.
 * The template is used to create and attach the tag to GCP resources.
 * 
 * To get more information about TagTemplate, see:
 * 
 * * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.tagTemplates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
 * 
 * ## Example Usage
 * ### Data Catalog Tag Template Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.datacatalog.TagTemplate;
 * import com.pulumi.gcp.datacatalog.TagTemplateArgs;
 * import com.pulumi.gcp.datacatalog.inputs.TagTemplateFieldArgs;
 * import com.pulumi.gcp.datacatalog.inputs.TagTemplateFieldTypeArgs;
 * import com.pulumi.gcp.datacatalog.inputs.TagTemplateFieldTypeEnumTypeArgs;
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
 *         var basicTagTemplate = new TagTemplate(&#34;basicTagTemplate&#34;, TagTemplateArgs.builder()        
 *             .displayName(&#34;Demo Tag Template&#34;)
 *             .fields(            
 *                 TagTemplateFieldArgs.builder()
 *                     .displayName(&#34;Source of data asset&#34;)
 *                     .fieldId(&#34;source&#34;)
 *                     .isRequired(true)
 *                     .type(TagTemplateFieldTypeArgs.builder()
 *                         .primitiveType(&#34;STRING&#34;)
 *                         .build())
 *                     .build(),
 *                 TagTemplateFieldArgs.builder()
 *                     .displayName(&#34;Number of rows in the data asset&#34;)
 *                     .fieldId(&#34;num_rows&#34;)
 *                     .type(TagTemplateFieldTypeArgs.builder()
 *                         .primitiveType(&#34;DOUBLE&#34;)
 *                         .build())
 *                     .build(),
 *                 TagTemplateFieldArgs.builder()
 *                     .displayName(&#34;PII type&#34;)
 *                     .fieldId(&#34;pii_type&#34;)
 *                     .type(TagTemplateFieldTypeArgs.builder()
 *                         .enumType(TagTemplateFieldTypeEnumTypeArgs.builder()
 *                             .allowedValues(                            
 *                                 TagTemplateFieldTypeEnumTypeAllowedValueArgs.builder()
 *                                     .displayName(&#34;EMAIL&#34;)
 *                                     .build(),
 *                                 TagTemplateFieldTypeEnumTypeAllowedValueArgs.builder()
 *                                     .displayName(&#34;SOCIAL SECURITY NUMBER&#34;)
 *                                     .build(),
 *                                 TagTemplateFieldTypeEnumTypeAllowedValueArgs.builder()
 *                                     .displayName(&#34;NONE&#34;)
 *                                     .build())
 *                             .build())
 *                         .build())
 *                     .build())
 *             .forceDelete(&#34;false&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .tagTemplateId(&#34;my_template&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * TagTemplate can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:datacatalog/tagTemplate:TagTemplate default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:datacatalog/tagTemplate:TagTemplate")
public class TagTemplate extends com.pulumi.resources.CustomResource {
    /**
     * The display name for this template.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output</* @Nullable */ String> displayName;

    /**
     * @return The display name for this template.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
     * Structure is documented below.
     * 
     */
    @Export(name="fields", type=List.class, parameters={TagTemplateField.class})
    private Output<List<TagTemplateField>> fields;

    /**
     * @return Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
     * Structure is documented below.
     * 
     */
    public Output<List<TagTemplateField>> fields() {
        return this.fields;
    }
    /**
     * This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
     * 
     */
    @Export(name="forceDelete", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> forceDelete;

    /**
     * @return This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
     * 
     */
    public Output<Optional<Boolean>> forceDelete() {
        return Codegen.optional(this.forceDelete);
    }
    /**
     * - 
     * The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return -
     * The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
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
     * Template location region.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return Template location region.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The id of the tag template to create.
     * 
     */
    @Export(name="tagTemplateId", type=String.class, parameters={})
    private Output<String> tagTemplateId;

    /**
     * @return The id of the tag template to create.
     * 
     */
    public Output<String> tagTemplateId() {
        return this.tagTemplateId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TagTemplate(String name) {
        this(name, TagTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TagTemplate(String name, TagTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TagTemplate(String name, TagTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:datacatalog/tagTemplate:TagTemplate", name, args == null ? TagTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TagTemplate(String name, Output<String> id, @Nullable TagTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:datacatalog/tagTemplate:TagTemplate", name, state, makeResourceOptions(options, id));
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
    public static TagTemplate get(String name, Output<String> id, @Nullable TagTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TagTemplate(name, id, state, options);
    }
}

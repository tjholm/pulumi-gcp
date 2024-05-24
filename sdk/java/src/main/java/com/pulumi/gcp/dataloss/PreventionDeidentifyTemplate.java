// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataloss.PreventionDeidentifyTemplateArgs;
import com.pulumi.gcp.dataloss.inputs.PreventionDeidentifyTemplateState;
import com.pulumi.gcp.dataloss.outputs.PreventionDeidentifyTemplateDeidentifyConfig;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows creation of templates to de-identify content.
 * 
 * To get more information about DeidentifyTemplate, see:
 * 
 * * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.deidentifyTemplates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dlp/docs/concepts-templates)
 * 
 * ## Example Usage
 * 
 * ### Dlp Deidentify Template Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataloss.PreventionDeidentifyTemplate;
 * import com.pulumi.gcp.dataloss.PreventionDeidentifyTemplateArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionDeidentifyTemplateDeidentifyConfigArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsArgs;
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
 *         var basic = new PreventionDeidentifyTemplate("basic", PreventionDeidentifyTemplateArgs.builder()
 *             .parent("projects/my-project-name")
 *             .description("Description")
 *             .displayName("Displayname")
 *             .deidentifyConfig(PreventionDeidentifyTemplateDeidentifyConfigArgs.builder()
 *                 .infoTypeTransformations(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsArgs.builder()
 *                     .transformations(                    
 *                         PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs.builder()
 *                             .infoTypes(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                 .name("FIRST_NAME")
 *                                 .build())
 *                             .primitiveTransformation(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs.builder()
 *                                 .replaceWithInfoTypeConfig(true)
 *                                 .build())
 *                             .build(),
 *                         PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs.builder()
 *                             .infoTypes(                            
 *                                 PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                     .name("PHONE_NUMBER")
 *                                     .build(),
 *                                 PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                     .name("AGE")
 *                                     .build())
 *                             .primitiveTransformation(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs.builder()
 *                                 .replaceConfig(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigArgs.builder()
 *                                     .newValue(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueArgs.builder()
 *                                         .integerValue(9)
 *                                         .build())
 *                                     .build())
 *                                 .build())
 *                             .build(),
 *                         PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs.builder()
 *                             .infoTypes(                            
 *                                 PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                     .name("EMAIL_ADDRESS")
 *                                     .build(),
 *                                 PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                     .name("LAST_NAME")
 *                                     .build())
 *                             .primitiveTransformation(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs.builder()
 *                                 .characterMaskConfig(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCharacterMaskConfigArgs.builder()
 *                                     .maskingCharacter("X")
 *                                     .numberToMask(4)
 *                                     .reverseOrder(true)
 *                                     .charactersToIgnores(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArgs.builder()
 *                                         .commonCharactersToIgnore("PUNCTUATION")
 *                                         .build())
 *                                     .build())
 *                                 .build())
 *                             .build(),
 *                         PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs.builder()
 *                             .infoTypes(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                 .name("DATE_OF_BIRTH")
 *                                 .build())
 *                             .primitiveTransformation(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs.builder()
 *                                 .replaceConfig(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigArgs.builder()
 *                                     .newValue(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueArgs.builder()
 *                                         .dateValue(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueDateValueArgs.builder()
 *                                             .year(2020)
 *                                             .month(1)
 *                                             .day(1)
 *                                             .build())
 *                                         .build())
 *                                     .build())
 *                                 .build())
 *                             .build(),
 *                         PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs.builder()
 *                             .infoTypes(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                 .name("CREDIT_CARD_NUMBER")
 *                                 .build())
 *                             .primitiveTransformation(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs.builder()
 *                                 .cryptoDeterministicConfig(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigArgs.builder()
 *                                     .context(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigContextArgs.builder()
 *                                         .name("sometweak")
 *                                         .build())
 *                                     .cryptoKey(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigCryptoKeyArgs.builder()
 *                                         .transient_(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransientArgs.builder()
 *                                             .name("beep")
 *                                             .build())
 *                                         .build())
 *                                     .surrogateInfoType(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoTypeArgs.builder()
 *                                         .name("abc")
 *                                         .build())
 *                                     .build())
 *                                 .build())
 *                             .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Dlp Deidentify Template Image Transformations
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataloss.PreventionDeidentifyTemplate;
 * import com.pulumi.gcp.dataloss.PreventionDeidentifyTemplateArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionDeidentifyTemplateDeidentifyConfigArgs;
 * import com.pulumi.gcp.dataloss.inputs.PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsArgs;
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
 *         var basic = new PreventionDeidentifyTemplate("basic", PreventionDeidentifyTemplateArgs.builder()
 *             .parent("projects/my-project-name")
 *             .description("Description")
 *             .displayName("Displayname")
 *             .deidentifyConfig(PreventionDeidentifyTemplateDeidentifyConfigArgs.builder()
 *                 .imageTransformations(PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsArgs.builder()
 *                     .transforms(                    
 *                         PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformArgs.builder()
 *                             .redactionColor(PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColorArgs.builder()
 *                                 .red(0.5)
 *                                 .blue(1)
 *                                 .green(0.2)
 *                                 .build())
 *                             .selectedInfoTypes(PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypesArgs.builder()
 *                                 .infoTypes(PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypesInfoTypeArgs.builder()
 *                                     .name("COLOR_INFO")
 *                                     .version("latest")
 *                                     .build())
 *                                 .build())
 *                             .build(),
 *                         PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformArgs.builder()
 *                             .allInfoTypes()
 *                             .build(),
 *                         PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformArgs.builder()
 *                             .allText()
 *                             .build())
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
 * DeidentifyTemplate can be imported using any of these accepted formats:
 * 
 * * `{{parent}}/deidentifyTemplates/{{name}}`
 * 
 * * `{{parent}}/{{name}}`
 * 
 * When using the `pulumi import` command, DeidentifyTemplate can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/deidentifyTemplates/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate")
public class PreventionDeidentifyTemplate extends com.pulumi.resources.CustomResource {
    /**
     * The creation timestamp of an deidentifyTemplate. Set by the server.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation timestamp of an deidentifyTemplate. Set by the server.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Configuration of the deidentify template
     * Structure is documented below.
     * 
     */
    @Export(name="deidentifyConfig", refs={PreventionDeidentifyTemplateDeidentifyConfig.class}, tree="[0]")
    private Output<PreventionDeidentifyTemplateDeidentifyConfig> deidentifyConfig;

    /**
     * @return Configuration of the deidentify template
     * Structure is documented below.
     * 
     */
    public Output<PreventionDeidentifyTemplateDeidentifyConfig> deidentifyConfig() {
        return this.deidentifyConfig;
    }
    /**
     * A description of the template.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the template.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * User set display name of the template.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return User set display name of the template.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * The resource name of the template. Set by the server.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the template. Set by the server.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The parent of the template in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return The parent of the template in any of the following formats:
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
     * The template id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular
     * expression: [a-zA-Z\d-_]+. The maximum length is 100 characters. Can be empty to allow the system to generate one.
     * 
     */
    @Export(name="templateId", refs={String.class}, tree="[0]")
    private Output<String> templateId;

    /**
     * @return The template id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular
     * expression: [a-zA-Z\d-_]+. The maximum length is 100 characters. Can be empty to allow the system to generate one.
     * 
     */
    public Output<String> templateId() {
        return this.templateId;
    }
    /**
     * The last update timestamp of an deidentifyTemplate. Set by the server.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The last update timestamp of an deidentifyTemplate. Set by the server.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PreventionDeidentifyTemplate(String name) {
        this(name, PreventionDeidentifyTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PreventionDeidentifyTemplate(String name, PreventionDeidentifyTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PreventionDeidentifyTemplate(String name, PreventionDeidentifyTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate", name, args == null ? PreventionDeidentifyTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PreventionDeidentifyTemplate(String name, Output<String> id, @Nullable PreventionDeidentifyTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate", name, state, makeResourceOptions(options, id));
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
    public static PreventionDeidentifyTemplate get(String name, Output<String> id, @Nullable PreventionDeidentifyTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PreventionDeidentifyTemplate(name, id, state, options);
    }
}

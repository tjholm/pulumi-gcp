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
 * ### Dlp Deidentify Template Basic
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var basic = new PreventionDeidentifyTemplate(&#34;basic&#34;, PreventionDeidentifyTemplateArgs.builder()        
 *             .deidentifyConfig(PreventionDeidentifyTemplateDeidentifyConfigArgs.builder()
 *                 .infoTypeTransformations(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsArgs.builder()
 *                     .transformations(                    
 *                         PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs.builder()
 *                             .infoTypes(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                 .name(&#34;FIRST_NAME&#34;)
 *                                 .build())
 *                             .primitiveTransformation(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs.builder()
 *                                 .replaceWithInfoTypeConfig(true)
 *                                 .build())
 *                             .build(),
 *                         PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs.builder()
 *                             .infoTypes(                            
 *                                 PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                     .name(&#34;PHONE_NUMBER&#34;)
 *                                     .build(),
 *                                 PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                     .name(&#34;AGE&#34;)
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
 *                                     .name(&#34;EMAIL_ADDRESS&#34;)
 *                                     .build(),
 *                                 PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                     .name(&#34;LAST_NAME&#34;)
 *                                     .build())
 *                             .primitiveTransformation(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs.builder()
 *                                 .characterMaskConfig(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCharacterMaskConfigArgs.builder()
 *                                     .charactersToIgnore(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *                                     .maskingCharacter(&#34;X&#34;)
 *                                     .numberToMask(4)
 *                                     .reverseOrder(true)
 *                                     .build())
 *                                 .build())
 *                             .build(),
 *                         PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs.builder()
 *                             .infoTypes(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                 .name(&#34;DATE_OF_BIRTH&#34;)
 *                                 .build())
 *                             .primitiveTransformation(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs.builder()
 *                                 .replaceConfig(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigArgs.builder()
 *                                     .newValue(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueArgs.builder()
 *                                         .dateValue(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueDateValueArgs.builder()
 *                                             .day(1)
 *                                             .month(1)
 *                                             .year(2020)
 *                                             .build())
 *                                         .build())
 *                                     .build())
 *                                 .build())
 *                             .build(),
 *                         PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs.builder()
 *                             .infoTypes(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs.builder()
 *                                 .name(&#34;CREDIT_CARD_NUMBER&#34;)
 *                                 .build())
 *                             .primitiveTransformation(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs.builder()
 *                                 .cryptoDeterministicConfig(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigArgs.builder()
 *                                     .context(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigContextArgs.builder()
 *                                         .name(&#34;sometweak&#34;)
 *                                         .build())
 *                                     .cryptoKey(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigCryptoKeyArgs.builder()
 *                                         .transient_(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransientArgs.builder()
 *                                             .name(&#34;beep&#34;)
 *                                             .build())
 *                                         .build())
 *                                     .surrogateInfoType(PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoTypeArgs.builder()
 *                                         .name(&#34;abc&#34;)
 *                                         .build())
 *                                     .build())
 *                                 .build())
 *                             .build())
 *                     .build())
 *                 .build())
 *             .description(&#34;Description&#34;)
 *             .displayName(&#34;Displayname&#34;)
 *             .parent(&#34;projects/my-project-name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * DeidentifyTemplate can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/deidentifyTemplates/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate")
public class PreventionDeidentifyTemplate extends com.pulumi.resources.CustomResource {
    /**
     * Configuration of the deidentify template
     * Structure is documented below.
     * 
     */
    @Export(name="deidentifyConfig", type=PreventionDeidentifyTemplateDeidentifyConfig.class, parameters={})
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
    @Export(name="description", type=String.class, parameters={})
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
    @Export(name="displayName", type=String.class, parameters={})
    private Output</* @Nullable */ String> displayName;

    /**
     * @return User set display name of the template.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at [https://cloud.google.com/dlp/docs/infotypes-reference](https://cloud.google.com/dlp/docs/infotypes-reference) when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at [https://cloud.google.com/dlp/docs/infotypes-reference](https://cloud.google.com/dlp/docs/infotypes-reference) when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.
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
    @Export(name="parent", type=String.class, parameters={})
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

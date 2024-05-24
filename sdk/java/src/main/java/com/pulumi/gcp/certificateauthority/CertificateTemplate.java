// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.certificateauthority.CertificateTemplateArgs;
import com.pulumi.gcp.certificateauthority.inputs.CertificateTemplateState;
import com.pulumi.gcp.certificateauthority.outputs.CertificateTemplateIdentityConstraints;
import com.pulumi.gcp.certificateauthority.outputs.CertificateTemplatePassthroughExtensions;
import com.pulumi.gcp.certificateauthority.outputs.CertificateTemplatePredefinedValues;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Certificate Authority Service provides reusable and parameterized templates that you can use for common certificate issuance scenarios. A certificate template represents a relatively static and well-defined certificate issuance schema within an organization.  A certificate template can essentially become a full-fledged vertical certificate issuance framework.
 * 
 * For more information, see:
 * * [Understanding Certificate Templates](https://cloud.google.com/certificate-authority-service/docs/certificate-template)
 * * [Common configurations and Certificate Profiles](https://cloud.google.com/certificate-authority-service/docs/certificate-profile)
 * ## Example Usage
 * 
 * ### Basic_certificate_template
 * An example of a basic privateca certificate template
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificateauthority.CertificateTemplate;
 * import com.pulumi.gcp.certificateauthority.CertificateTemplateArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CertificateTemplateIdentityConstraintsArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CertificateTemplateIdentityConstraintsCelExpressionArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CertificateTemplatePassthroughExtensionsArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CertificateTemplatePredefinedValuesArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CertificateTemplatePredefinedValuesCaOptionsArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CertificateTemplatePredefinedValuesKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageArgs;
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
 *         var primary = new CertificateTemplate("primary", CertificateTemplateArgs.builder()
 *             .location("us-west1")
 *             .name("template")
 *             .description("An updated sample certificate template")
 *             .identityConstraints(CertificateTemplateIdentityConstraintsArgs.builder()
 *                 .allowSubjectAltNamesPassthrough(true)
 *                 .allowSubjectPassthrough(true)
 *                 .celExpression(CertificateTemplateIdentityConstraintsCelExpressionArgs.builder()
 *                     .description("Always true")
 *                     .expression("true")
 *                     .location("any.file.anywhere")
 *                     .title("Sample expression")
 *                     .build())
 *                 .build())
 *             .maximumLifetime("86400s")
 *             .passthroughExtensions(CertificateTemplatePassthroughExtensionsArgs.builder()
 *                 .additionalExtensions(CertificateTemplatePassthroughExtensionsAdditionalExtensionArgs.builder()
 *                     .objectIdPaths(                    
 *                         1,
 *                         6)
 *                     .build())
 *                 .knownExtensions("EXTENDED_KEY_USAGE")
 *                 .build())
 *             .predefinedValues(CertificateTemplatePredefinedValuesArgs.builder()
 *                 .additionalExtensions(CertificateTemplatePredefinedValuesAdditionalExtensionArgs.builder()
 *                     .objectId(CertificateTemplatePredefinedValuesAdditionalExtensionObjectIdArgs.builder()
 *                         .objectIdPaths(                        
 *                             1,
 *                             6)
 *                         .build())
 *                     .value("c3RyaW5nCg==")
 *                     .critical(true)
 *                     .build())
 *                 .aiaOcspServers("string")
 *                 .caOptions(CertificateTemplatePredefinedValuesCaOptionsArgs.builder()
 *                     .isCa(false)
 *                     .maxIssuerPathLength(6)
 *                     .build())
 *                 .keyUsage(CertificateTemplatePredefinedValuesKeyUsageArgs.builder()
 *                     .baseKeyUsage(CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageArgs.builder()
 *                         .certSign(false)
 *                         .contentCommitment(true)
 *                         .crlSign(false)
 *                         .dataEncipherment(true)
 *                         .decipherOnly(true)
 *                         .digitalSignature(true)
 *                         .encipherOnly(true)
 *                         .keyAgreement(true)
 *                         .keyEncipherment(true)
 *                         .build())
 *                     .extendedKeyUsage(CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageArgs.builder()
 *                         .clientAuth(true)
 *                         .codeSigning(true)
 *                         .emailProtection(true)
 *                         .ocspSigning(true)
 *                         .serverAuth(true)
 *                         .timeStamping(true)
 *                         .build())
 *                     .unknownExtendedKeyUsages(CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs.builder()
 *                         .objectIdPaths(                        
 *                             1,
 *                             6)
 *                         .build())
 *                     .build())
 *                 .policyIds(CertificateTemplatePredefinedValuesPolicyIdArgs.builder()
 *                     .objectIdPaths(                    
 *                         1,
 *                         6)
 *                     .build())
 *                 .build())
 *             .project("my-project-name")
 *             .labels(Map.of("label-two", "value-two"))
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
 * CertificateTemplate can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{name}}`
 * 
 * * `{{location}}/{{name}}`
 * 
 * When using the `pulumi import` command, CertificateTemplate can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:certificateauthority/certificateTemplate:CertificateTemplate")
public class CertificateTemplate extends com.pulumi.resources.CustomResource {
    /**
     * Output only. The time at which this CertificateTemplate was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. The time at which this CertificateTemplate was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Optional. A human-readable description of scenarios this template is intended for.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Optional. A human-readable description of scenarios this template is intended for.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is
     * omitted, then this template will not add restrictions on a certificate&#39;s identity.
     * 
     */
    @Export(name="identityConstraints", refs={CertificateTemplateIdentityConstraints.class}, tree="[0]")
    private Output</* @Nullable */ CertificateTemplateIdentityConstraints> identityConstraints;

    /**
     * @return Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is
     * omitted, then this template will not add restrictions on a certificate&#39;s identity.
     * 
     */
    public Output<Optional<CertificateTemplateIdentityConstraints>> identityConstraints() {
        return Codegen.optional(this.identityConstraints);
    }
    /**
     * Optional. Labels with user-defined metadata. **Note**: This field is non-authoritative, and will only manage the labels
     * present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the
     * resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Optional. Labels with user-defined metadata. **Note**: This field is non-authoritative, and will only manage the labels
     * present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the
     * resource.
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
     * Optional. The maximum lifetime allowed for all issued certificates that use this template. If the issuing CaPool&#39;s
     * IssuancePolicy specifies a maximum lifetime the minimum of the two durations will be the maximum lifetime for issued.
     * Note that if the issuing CertificateAuthority expires before a Certificate&#39;s requested maximum_lifetime, the effective
     * lifetime will be explicitly truncated to match it.
     * 
     */
    @Export(name="maximumLifetime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> maximumLifetime;

    /**
     * @return Optional. The maximum lifetime allowed for all issued certificates that use this template. If the issuing CaPool&#39;s
     * IssuancePolicy specifies a maximum lifetime the minimum of the two durations will be the maximum lifetime for issued.
     * Note that if the issuing CertificateAuthority expires before a Certificate&#39;s requested maximum_lifetime, the effective
     * lifetime will be explicitly truncated to match it.
     * 
     */
    public Output<Optional<String>> maximumLifetime() {
        return Codegen.optional(this.maximumLifetime);
    }
    /**
     * The resource name for this CertificateTemplate in the format `projects/*{@literal /}locations/*{@literal /}certificateTemplates/*`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name for this CertificateTemplate in the format `projects/*{@literal /}locations/*{@literal /}certificateTemplates/*`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate.
     * If a certificate request sets extensions that don&#39;t appear in the passthrough_extensions, those extensions will be
     * dropped. If the issuing CaPool&#39;s IssuancePolicy defines baseline_values that don&#39;t appear here, the certificate issuance
     * request will fail. If this is omitted, then this template will not add restrictions on a certificate&#39;s X.509 extensions.
     * These constraints do not apply to X.509 extensions set in this CertificateTemplate&#39;s predefined_values.
     * 
     */
    @Export(name="passthroughExtensions", refs={CertificateTemplatePassthroughExtensions.class}, tree="[0]")
    private Output</* @Nullable */ CertificateTemplatePassthroughExtensions> passthroughExtensions;

    /**
     * @return Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate.
     * If a certificate request sets extensions that don&#39;t appear in the passthrough_extensions, those extensions will be
     * dropped. If the issuing CaPool&#39;s IssuancePolicy defines baseline_values that don&#39;t appear here, the certificate issuance
     * request will fail. If this is omitted, then this template will not add restrictions on a certificate&#39;s X.509 extensions.
     * These constraints do not apply to X.509 extensions set in this CertificateTemplate&#39;s predefined_values.
     * 
     */
    public Output<Optional<CertificateTemplatePassthroughExtensions>> passthroughExtensions() {
        return Codegen.optional(this.passthroughExtensions);
    }
    /**
     * Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the
     * certificate request includes conflicting values for the same properties, they will be overwritten by the values defined
     * here. If the issuing CaPool&#39;s IssuancePolicy defines conflicting baseline_values for the same properties, the
     * certificate issuance request will fail.
     * 
     */
    @Export(name="predefinedValues", refs={CertificateTemplatePredefinedValues.class}, tree="[0]")
    private Output</* @Nullable */ CertificateTemplatePredefinedValues> predefinedValues;

    /**
     * @return Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the
     * certificate request includes conflicting values for the same properties, they will be overwritten by the values defined
     * here. If the issuing CaPool&#39;s IssuancePolicy defines conflicting baseline_values for the same properties, the
     * certificate issuance request will fail.
     * 
     */
    public Output<Optional<CertificateTemplatePredefinedValues>> predefinedValues() {
        return Codegen.optional(this.predefinedValues);
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
     * Output only. The time at which this CertificateTemplate was updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. The time at which this CertificateTemplate was updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CertificateTemplate(String name) {
        this(name, CertificateTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CertificateTemplate(String name, CertificateTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CertificateTemplate(String name, CertificateTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificateauthority/certificateTemplate:CertificateTemplate", name, args == null ? CertificateTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CertificateTemplate(String name, Output<String> id, @Nullable CertificateTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificateauthority/certificateTemplate:CertificateTemplate", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static CertificateTemplate get(String name, Output<String> id, @Nullable CertificateTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CertificateTemplate(name, id, state, options);
    }
}

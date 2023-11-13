// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.certificateauthority.CaPoolArgs;
import com.pulumi.gcp.certificateauthority.inputs.CaPoolState;
import com.pulumi.gcp.certificateauthority.outputs.CaPoolIssuancePolicy;
import com.pulumi.gcp.certificateauthority.outputs.CaPoolPublishingOptions;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A CaPool represents a group of CertificateAuthorities that form a trust anchor. A CaPool can be used to manage
 * issuance policies for one or more CertificateAuthority resources and to rotate CA certificates in and out of the
 * trust anchor.
 * 
 * ## Example Usage
 * ### Privateca Capool Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificateauthority.CaPool;
 * import com.pulumi.gcp.certificateauthority.CaPoolArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolPublishingOptionsArgs;
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
 *         var default_ = new CaPool(&#34;default&#34;, CaPoolArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .location(&#34;us-central1&#34;)
 *             .publishingOptions(CaPoolPublishingOptionsArgs.builder()
 *                 .publishCaCert(true)
 *                 .publishCrl(true)
 *                 .build())
 *             .tier(&#34;ENTERPRISE&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Privateca Capool All Fields
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificateauthority.CaPool;
 * import com.pulumi.gcp.certificateauthority.CaPoolArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolIssuancePolicyArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolIssuancePolicyAllowedIssuanceModesArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolIssuancePolicyBaselineValuesArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolIssuancePolicyBaselineValuesCaOptionsArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolIssuancePolicyBaselineValuesKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolIssuancePolicyBaselineValuesNameConstraintsArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolIssuancePolicyIdentityConstraintsArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolIssuancePolicyIdentityConstraintsCelExpressionArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.CaPoolPublishingOptionsArgs;
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
 *         var default_ = new CaPool(&#34;default&#34;, CaPoolArgs.builder()        
 *             .issuancePolicy(CaPoolIssuancePolicyArgs.builder()
 *                 .allowedIssuanceModes(CaPoolIssuancePolicyAllowedIssuanceModesArgs.builder()
 *                     .allowConfigBasedIssuance(true)
 *                     .allowCsrBasedIssuance(true)
 *                     .build())
 *                 .allowedKeyTypes(                
 *                     CaPoolIssuancePolicyAllowedKeyTypeArgs.builder()
 *                         .ellipticCurve(CaPoolIssuancePolicyAllowedKeyTypeEllipticCurveArgs.builder()
 *                             .signatureAlgorithm(&#34;ECDSA_P256&#34;)
 *                             .build())
 *                         .build(),
 *                     CaPoolIssuancePolicyAllowedKeyTypeArgs.builder()
 *                         .rsa(CaPoolIssuancePolicyAllowedKeyTypeRsaArgs.builder()
 *                             .maxModulusSize(10)
 *                             .minModulusSize(5)
 *                             .build())
 *                         .build())
 *                 .baselineValues(CaPoolIssuancePolicyBaselineValuesArgs.builder()
 *                     .additionalExtensions(CaPoolIssuancePolicyBaselineValuesAdditionalExtensionArgs.builder()
 *                         .critical(true)
 *                         .objectId(CaPoolIssuancePolicyBaselineValuesAdditionalExtensionObjectIdArgs.builder()
 *                             .objectIdPath(                            
 *                                 1,
 *                                 7)
 *                             .build())
 *                         .value(&#34;asdf&#34;)
 *                         .build())
 *                     .aiaOcspServers(&#34;example.com&#34;)
 *                     .caOptions(CaPoolIssuancePolicyBaselineValuesCaOptionsArgs.builder()
 *                         .isCa(true)
 *                         .maxIssuerPathLength(10)
 *                         .build())
 *                     .keyUsage(CaPoolIssuancePolicyBaselineValuesKeyUsageArgs.builder()
 *                         .baseKeyUsage(CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageArgs.builder()
 *                             .certSign(false)
 *                             .contentCommitment(true)
 *                             .crlSign(true)
 *                             .dataEncipherment(true)
 *                             .decipherOnly(true)
 *                             .digitalSignature(true)
 *                             .keyAgreement(true)
 *                             .keyEncipherment(false)
 *                             .build())
 *                         .extendedKeyUsage(CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageArgs.builder()
 *                             .clientAuth(false)
 *                             .codeSigning(true)
 *                             .emailProtection(true)
 *                             .serverAuth(true)
 *                             .timeStamping(true)
 *                             .build())
 *                         .build())
 *                     .nameConstraints(CaPoolIssuancePolicyBaselineValuesNameConstraintsArgs.builder()
 *                         .critical(true)
 *                         .excludedDnsNames(                        
 *                             &#34;*.deny.example1.com&#34;,
 *                             &#34;*.deny.example2.com&#34;)
 *                         .excludedEmailAddresses(                        
 *                             &#34;.deny.example1.com&#34;,
 *                             &#34;.deny.example2.com&#34;)
 *                         .excludedIpRanges(                        
 *                             &#34;10.1.1.0/24&#34;,
 *                             &#34;11.1.1.0/24&#34;)
 *                         .excludedUris(                        
 *                             &#34;.deny.example1.com&#34;,
 *                             &#34;.deny.example2.com&#34;)
 *                         .permittedDnsNames(                        
 *                             &#34;*.example1.com&#34;,
 *                             &#34;*.example2.com&#34;)
 *                         .permittedEmailAddresses(                        
 *                             &#34;.example1.com&#34;,
 *                             &#34;.example2.com&#34;)
 *                         .permittedIpRanges(                        
 *                             &#34;10.0.0.0/8&#34;,
 *                             &#34;11.0.0.0/8&#34;)
 *                         .permittedUris(                        
 *                             &#34;.example1.com&#34;,
 *                             &#34;.example2.com&#34;)
 *                         .build())
 *                     .policyIds(                    
 *                         CaPoolIssuancePolicyBaselineValuesPolicyIdArgs.builder()
 *                             .objectIdPath(                            
 *                                 1,
 *                                 5)
 *                             .build(),
 *                         CaPoolIssuancePolicyBaselineValuesPolicyIdArgs.builder()
 *                             .objectIdPath(                            
 *                                 1,
 *                                 5,
 *                                 7)
 *                             .build())
 *                     .build())
 *                 .identityConstraints(CaPoolIssuancePolicyIdentityConstraintsArgs.builder()
 *                     .allowSubjectAltNamesPassthrough(true)
 *                     .allowSubjectPassthrough(true)
 *                     .celExpression(CaPoolIssuancePolicyIdentityConstraintsCelExpressionArgs.builder()
 *                         .expression(&#34;subject_alt_names.all(san, san.type == DNS || san.type == EMAIL )&#34;)
 *                         .title(&#34;My title&#34;)
 *                         .build())
 *                     .build())
 *                 .maximumLifetime(&#34;50000s&#34;)
 *                 .build())
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .location(&#34;us-central1&#34;)
 *             .publishingOptions(CaPoolPublishingOptionsArgs.builder()
 *                 .encodingFormat(&#34;PEM&#34;)
 *                 .publishCaCert(false)
 *                 .publishCrl(true)
 *                 .build())
 *             .tier(&#34;ENTERPRISE&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * CaPool can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:certificateauthority/caPool:CaPool default projects/{{project}}/locations/{{location}}/caPools/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:certificateauthority/caPool:CaPool default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:certificateauthority/caPool:CaPool default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:certificateauthority/caPool:CaPool")
public class CaPool extends com.pulumi.resources.CustomResource {
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
     * The IssuancePolicy to control how Certificates will be issued from this CaPool.
     * Structure is documented below.
     * 
     */
    @Export(name="issuancePolicy", refs={CaPoolIssuancePolicy.class}, tree="[0]")
    private Output</* @Nullable */ CaPoolIssuancePolicy> issuancePolicy;

    /**
     * @return The IssuancePolicy to control how Certificates will be issued from this CaPool.
     * Structure is documented below.
     * 
     */
    public Output<Optional<CaPoolIssuancePolicy>> issuancePolicy() {
        return Codegen.optional(this.issuancePolicy);
    }
    /**
     * Labels with user-defined metadata.
     * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;:
     * &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels with user-defined metadata.
     * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;:
     * &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Location of the CaPool. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return Location of the CaPool. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The name for this CaPool.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name for this CaPool.
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
     * The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
     * Structure is documented below.
     * 
     */
    @Export(name="publishingOptions", refs={CaPoolPublishingOptions.class}, tree="[0]")
    private Output</* @Nullable */ CaPoolPublishingOptions> publishingOptions;

    /**
     * @return The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
     * Structure is documented below.
     * 
     */
    public Output<Optional<CaPoolPublishingOptions>> publishingOptions() {
        return Codegen.optional(this.publishingOptions);
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
     * The Tier of this CaPool.
     * Possible values are: `ENTERPRISE`, `DEVOPS`.
     * 
     */
    @Export(name="tier", refs={String.class}, tree="[0]")
    private Output<String> tier;

    /**
     * @return The Tier of this CaPool.
     * Possible values are: `ENTERPRISE`, `DEVOPS`.
     * 
     */
    public Output<String> tier() {
        return this.tier;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CaPool(String name) {
        this(name, CaPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CaPool(String name, CaPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CaPool(String name, CaPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificateauthority/caPool:CaPool", name, args == null ? CaPoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CaPool(String name, Output<String> id, @Nullable CaPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificateauthority/caPool:CaPool", name, state, makeResourceOptions(options, id));
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
    public static CaPool get(String name, Output<String> id, @Nullable CaPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CaPool(name, id, state, options);
    }
}

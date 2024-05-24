// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificatemanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.certificatemanager.CertificateIssuanceConfigArgs;
import com.pulumi.gcp.certificatemanager.inputs.CertificateIssuanceConfigState;
import com.pulumi.gcp.certificatemanager.outputs.CertificateIssuanceConfigCertificateAuthorityConfig;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Certificate represents a HTTP-reachable backend for a Certificate.
 * 
 * To get more information about CertificateIssuanceConfig, see:
 * 
 * * [API documentation](https://cloud.google.com/certificate-manager/docs/reference/certificate-manager/rest/v1/projects.locations.certificateIssuanceConfigs)
 * * How-to Guides
 *     * [Manage certificate issuance configs](https://cloud.google.com/certificate-manager/docs/issuance-configs)
 * 
 * ## Example Usage
 * 
 * ### Certificate Manager Certificate Issuance Config
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificateauthority.CaPool;
 * import com.pulumi.gcp.certificateauthority.CaPoolArgs;
 * import com.pulumi.gcp.certificatemanager.CertificateIssuanceConfig;
 * import com.pulumi.gcp.certificatemanager.CertificateIssuanceConfigArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateIssuanceConfigCertificateAuthorityConfigArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigArgs;
 * import com.pulumi.gcp.certificateauthority.Authority;
 * import com.pulumi.gcp.certificateauthority.AuthorityArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigSubjectConfigArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigSubjectConfigSubjectArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigSubjectConfigSubjectAltNameArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigCaOptionsArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigKeyUsageBaseKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigKeyUsageExtendedKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityKeySpecArgs;
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
 *         var pool = new CaPool("pool", CaPoolArgs.builder()
 *             .name("ca-pool")
 *             .location("us-central1")
 *             .tier("ENTERPRISE")
 *             .build());
 * 
 *         var default_ = new CertificateIssuanceConfig("default", CertificateIssuanceConfigArgs.builder()
 *             .name("issuance-config")
 *             .description("sample description for the certificate issuanceConfigs")
 *             .certificateAuthorityConfig(CertificateIssuanceConfigCertificateAuthorityConfigArgs.builder()
 *                 .certificateAuthorityServiceConfig(CertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigArgs.builder()
 *                     .caPool(pool.id())
 *                     .build())
 *                 .build())
 *             .lifetime("1814400s")
 *             .rotationWindowPercentage(34)
 *             .keyAlgorithm("ECDSA_P256")
 *             .labels(Map.ofEntries(
 *                 Map.entry("name", "wrench"),
 *                 Map.entry("count", "3")
 *             ))
 *             .build());
 * 
 *         var caAuthority = new Authority("caAuthority", AuthorityArgs.builder()
 *             .location("us-central1")
 *             .pool(pool.name())
 *             .certificateAuthorityId("ca-authority")
 *             .config(AuthorityConfigArgs.builder()
 *                 .subjectConfig(AuthorityConfigSubjectConfigArgs.builder()
 *                     .subject(AuthorityConfigSubjectConfigSubjectArgs.builder()
 *                         .organization("HashiCorp")
 *                         .commonName("my-certificate-authority")
 *                         .build())
 *                     .subjectAltName(AuthorityConfigSubjectConfigSubjectAltNameArgs.builder()
 *                         .dnsNames("hashicorp.com")
 *                         .build())
 *                     .build())
 *                 .x509Config(AuthorityConfigX509ConfigArgs.builder()
 *                     .caOptions(AuthorityConfigX509ConfigCaOptionsArgs.builder()
 *                         .isCa(true)
 *                         .build())
 *                     .keyUsage(AuthorityConfigX509ConfigKeyUsageArgs.builder()
 *                         .baseKeyUsage(AuthorityConfigX509ConfigKeyUsageBaseKeyUsageArgs.builder()
 *                             .certSign(true)
 *                             .crlSign(true)
 *                             .build())
 *                         .extendedKeyUsage(AuthorityConfigX509ConfigKeyUsageExtendedKeyUsageArgs.builder()
 *                             .serverAuth(true)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .keySpec(AuthorityKeySpecArgs.builder()
 *                 .algorithm("RSA_PKCS1_4096_SHA256")
 *                 .build())
 *             .deletionProtection(false)
 *             .skipGracePeriod(true)
 *             .ignoreActiveCertificatesOnDeletion(true)
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
 * CertificateIssuanceConfig can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/certificateIssuanceConfigs/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{name}}`
 * 
 * * `{{location}}/{{name}}`
 * 
 * When using the `pulumi import` command, CertificateIssuanceConfig can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:certificatemanager/certificateIssuanceConfig:CertificateIssuanceConfig default projects/{{project}}/locations/{{location}}/certificateIssuanceConfigs/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:certificatemanager/certificateIssuanceConfig:CertificateIssuanceConfig default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:certificatemanager/certificateIssuanceConfig:CertificateIssuanceConfig default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:certificatemanager/certificateIssuanceConfig:CertificateIssuanceConfig")
public class CertificateIssuanceConfig extends com.pulumi.resources.CustomResource {
    /**
     * The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.
     * Structure is documented below.
     * 
     */
    @Export(name="certificateAuthorityConfig", refs={CertificateIssuanceConfigCertificateAuthorityConfig.class}, tree="[0]")
    private Output<CertificateIssuanceConfigCertificateAuthorityConfig> certificateAuthorityConfig;

    /**
     * @return The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.
     * Structure is documented below.
     * 
     */
    public Output<CertificateIssuanceConfigCertificateAuthorityConfig> certificateAuthorityConfig() {
        return this.certificateAuthorityConfig;
    }
    /**
     * The creation timestamp of a CertificateIssuanceConfig. Timestamp is in RFC3339 UTC &#34;Zulu&#34; format,
     * accurate to nanoseconds with up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation timestamp of a CertificateIssuanceConfig. Timestamp is in RFC3339 UTC &#34;Zulu&#34; format,
     * accurate to nanoseconds with up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * One or more paragraphs of text description of a CertificateIssuanceConfig.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return One or more paragraphs of text description of a CertificateIssuanceConfig.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * Key algorithm to use when generating the private key.
     * Possible values are: `RSA_2048`, `ECDSA_P256`.
     * 
     */
    @Export(name="keyAlgorithm", refs={String.class}, tree="[0]")
    private Output<String> keyAlgorithm;

    /**
     * @return Key algorithm to use when generating the private key.
     * Possible values are: `RSA_2048`, `ECDSA_P256`.
     * 
     */
    public Output<String> keyAlgorithm() {
        return this.keyAlgorithm;
    }
    /**
     * &#39;Set of label tags associated with the CertificateIssuanceConfig resource. An object containing a list of &#34;key&#34;: value
     * pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;count&#34;: &#34;3&#34; }. **Note**: This field is non-authoritative, and will only manage the
     * labels present in your configuration. Please refer to the field &#39;effective_labels&#39; for all of the labels present on the
     * resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return &#39;Set of label tags associated with the CertificateIssuanceConfig resource. An object containing a list of &#34;key&#34;: value
     * pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;count&#34;: &#34;3&#34; }. **Note**: This field is non-authoritative, and will only manage the
     * labels present in your configuration. Please refer to the field &#39;effective_labels&#39; for all of the labels present on the
     * resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Lifetime of issued certificates. A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;.
     * Example: &#34;1814400s&#34;. Valid values are from 21 days (1814400s) to 30 days (2592000s)
     * 
     */
    @Export(name="lifetime", refs={String.class}, tree="[0]")
    private Output<String> lifetime;

    /**
     * @return Lifetime of issued certificates. A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;.
     * Example: &#34;1814400s&#34;. Valid values are from 21 days (1814400s) to 30 days (2592000s)
     * 
     */
    public Output<String> lifetime() {
        return this.lifetime;
    }
    /**
     * The Certificate Manager location. If not specified, &#34;global&#34; is used.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> location;

    /**
     * @return The Certificate Manager location. If not specified, &#34;global&#34; is used.
     * 
     */
    public Output<Optional<String>> location() {
        return Codegen.optional(this.location);
    }
    /**
     * A user-defined name of the certificate issuance config.
     * CertificateIssuanceConfig names must be unique globally.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A user-defined name of the certificate issuance config.
     * CertificateIssuanceConfig names must be unique globally.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

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
     * It specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate.
     * Must be a number between 1-99, inclusive.
     * You must set the rotation window percentage in relation to the certificate lifetime so that certificate renewal occurs at least 7 days after
     * the certificate has been issued and at least 7 days before it expires.
     * 
     */
    @Export(name="rotationWindowPercentage", refs={Integer.class}, tree="[0]")
    private Output<Integer> rotationWindowPercentage;

    /**
     * @return It specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate.
     * Must be a number between 1-99, inclusive.
     * You must set the rotation window percentage in relation to the certificate lifetime so that certificate renewal occurs at least 7 days after
     * the certificate has been issued and at least 7 days before it expires.
     * 
     */
    public Output<Integer> rotationWindowPercentage() {
        return this.rotationWindowPercentage;
    }
    /**
     * The last update timestamp of a CertificateIssuanceConfig. Timestamp is in RFC3339 UTC &#34;Zulu&#34; format,
     * accurate to nanoseconds with up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The last update timestamp of a CertificateIssuanceConfig. Timestamp is in RFC3339 UTC &#34;Zulu&#34; format,
     * accurate to nanoseconds with up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CertificateIssuanceConfig(String name) {
        this(name, CertificateIssuanceConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CertificateIssuanceConfig(String name, CertificateIssuanceConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CertificateIssuanceConfig(String name, CertificateIssuanceConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificatemanager/certificateIssuanceConfig:CertificateIssuanceConfig", name, args == null ? CertificateIssuanceConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CertificateIssuanceConfig(String name, Output<String> id, @Nullable CertificateIssuanceConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificatemanager/certificateIssuanceConfig:CertificateIssuanceConfig", name, state, makeResourceOptions(options, id));
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
    public static CertificateIssuanceConfig get(String name, Output<String> id, @Nullable CertificateIssuanceConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CertificateIssuanceConfig(name, id, state, options);
    }
}

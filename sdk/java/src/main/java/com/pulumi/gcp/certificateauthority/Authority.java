// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.certificateauthority.AuthorityArgs;
import com.pulumi.gcp.certificateauthority.inputs.AuthorityState;
import com.pulumi.gcp.certificateauthority.outputs.AuthorityAccessUrl;
import com.pulumi.gcp.certificateauthority.outputs.AuthorityConfig;
import com.pulumi.gcp.certificateauthority.outputs.AuthorityKeySpec;
import com.pulumi.gcp.certificateauthority.outputs.AuthoritySubordinateConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A CertificateAuthority represents an individual Certificate Authority. A
 * CertificateAuthority can be used to create Certificates.
 * 
 * To get more information about CertificateAuthority, see:
 * 
 * * [API documentation](https://cloud.google.com/certificate-authority-service/docs/reference/rest)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/certificate-authority-service)
 * 
 * &gt; **Warning:** On newer versions of the provider, you must explicitly set `deletion_protection=false`
 * (and run `pulumi up` to write the field to state) in order to destroy a CertificateAuthority.
 * It is recommended to not set this field (or set it to true) until you&#39;re ready to destroy.
 * 
 * ## Example Usage
 * ### Privateca Certificate Authority Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *         var default_ = new Authority(&#34;default&#34;, AuthorityArgs.builder()        
 *             .certificateAuthorityId(&#34;my-certificate-authority&#34;)
 *             .config(AuthorityConfigArgs.builder()
 *                 .subjectConfig(AuthorityConfigSubjectConfigArgs.builder()
 *                     .subject(AuthorityConfigSubjectConfigSubjectArgs.builder()
 *                         .commonName(&#34;my-certificate-authority&#34;)
 *                         .organization(&#34;HashiCorp&#34;)
 *                         .build())
 *                     .subjectAltName(AuthorityConfigSubjectConfigSubjectAltNameArgs.builder()
 *                         .dnsNames(&#34;hashicorp.com&#34;)
 *                         .build())
 *                     .build())
 *                 .x509Config(AuthorityConfigX509ConfigArgs.builder()
 *                     .caOptions(AuthorityConfigX509ConfigCaOptionsArgs.builder()
 *                         .isCa(true)
 *                         .maxIssuerPathLength(10)
 *                         .build())
 *                     .keyUsage(AuthorityConfigX509ConfigKeyUsageArgs.builder()
 *                         .baseKeyUsage(AuthorityConfigX509ConfigKeyUsageBaseKeyUsageArgs.builder()
 *                             .certSign(true)
 *                             .contentCommitment(true)
 *                             .crlSign(true)
 *                             .dataEncipherment(true)
 *                             .decipherOnly(true)
 *                             .digitalSignature(true)
 *                             .keyAgreement(true)
 *                             .keyEncipherment(false)
 *                             .build())
 *                         .extendedKeyUsage(AuthorityConfigX509ConfigKeyUsageExtendedKeyUsageArgs.builder()
 *                             .clientAuth(false)
 *                             .codeSigning(true)
 *                             .emailProtection(true)
 *                             .serverAuth(true)
 *                             .timeStamping(true)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .deletionProtection(&#34;true&#34;)
 *             .keySpec(AuthorityKeySpecArgs.builder()
 *                 .algorithm(&#34;RSA_PKCS1_4096_SHA256&#34;)
 *                 .build())
 *             .lifetime(&#34;86400s&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .pool(&#34;ca-pool&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Privateca Certificate Authority Subordinate
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 * import com.pulumi.gcp.certificateauthority.inputs.AuthoritySubordinateConfigArgs;
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
 *         var root_ca = new Authority(&#34;root-ca&#34;, AuthorityArgs.builder()        
 *             .pool(&#34;ca-pool&#34;)
 *             .certificateAuthorityId(&#34;my-certificate-authority-root&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .config(AuthorityConfigArgs.builder()
 *                 .subjectConfig(AuthorityConfigSubjectConfigArgs.builder()
 *                     .subject(AuthorityConfigSubjectConfigSubjectArgs.builder()
 *                         .organization(&#34;HashiCorp&#34;)
 *                         .commonName(&#34;my-certificate-authority&#34;)
 *                         .build())
 *                     .subjectAltName(AuthorityConfigSubjectConfigSubjectAltNameArgs.builder()
 *                         .dnsNames(&#34;hashicorp.com&#34;)
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
 *                             .serverAuth(false)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .keySpec(AuthorityKeySpecArgs.builder()
 *                 .algorithm(&#34;RSA_PKCS1_4096_SHA256&#34;)
 *                 .build())
 *             .deletionProtection(false)
 *             .skipGracePeriod(true)
 *             .ignoreActiveCertificatesOnDeletion(true)
 *             .build());
 * 
 *         var default_ = new Authority(&#34;default&#34;, AuthorityArgs.builder()        
 *             .pool(&#34;ca-pool&#34;)
 *             .certificateAuthorityId(&#34;my-certificate-authority-sub&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .deletionProtection(&#34;true&#34;)
 *             .subordinateConfig(AuthoritySubordinateConfigArgs.builder()
 *                 .certificateAuthority(root_ca.name())
 *                 .build())
 *             .config(AuthorityConfigArgs.builder()
 *                 .subjectConfig(AuthorityConfigSubjectConfigArgs.builder()
 *                     .subject(AuthorityConfigSubjectConfigSubjectArgs.builder()
 *                         .organization(&#34;HashiCorp&#34;)
 *                         .commonName(&#34;my-subordinate-authority&#34;)
 *                         .build())
 *                     .subjectAltName(AuthorityConfigSubjectConfigSubjectAltNameArgs.builder()
 *                         .dnsNames(&#34;hashicorp.com&#34;)
 *                         .build())
 *                     .build())
 *                 .x509Config(AuthorityConfigX509ConfigArgs.builder()
 *                     .caOptions(AuthorityConfigX509ConfigCaOptionsArgs.builder()
 *                         .isCa(true)
 *                         .maxIssuerPathLength(0)
 *                         .build())
 *                     .keyUsage(AuthorityConfigX509ConfigKeyUsageArgs.builder()
 *                         .baseKeyUsage(AuthorityConfigX509ConfigKeyUsageBaseKeyUsageArgs.builder()
 *                             .digitalSignature(true)
 *                             .contentCommitment(true)
 *                             .keyEncipherment(false)
 *                             .dataEncipherment(true)
 *                             .keyAgreement(true)
 *                             .certSign(true)
 *                             .crlSign(true)
 *                             .decipherOnly(true)
 *                             .build())
 *                         .extendedKeyUsage(AuthorityConfigX509ConfigKeyUsageExtendedKeyUsageArgs.builder()
 *                             .serverAuth(true)
 *                             .clientAuth(false)
 *                             .emailProtection(true)
 *                             .codeSigning(true)
 *                             .timeStamping(true)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .lifetime(&#34;86400s&#34;)
 *             .keySpec(AuthorityKeySpecArgs.builder()
 *                 .algorithm(&#34;RSA_PKCS1_4096_SHA256&#34;)
 *                 .build())
 *             .type(&#34;SUBORDINATE&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Privateca Certificate Authority Byo Key
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.projects.ServiceIdentity;
 * import com.pulumi.gcp.projects.ServiceIdentityArgs;
 * import com.pulumi.gcp.kms.CryptoKeyIAMBinding;
 * import com.pulumi.gcp.kms.CryptoKeyIAMBindingArgs;
 * import com.pulumi.gcp.certificateauthority.Authority;
 * import com.pulumi.gcp.certificateauthority.AuthorityArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityKeySpecArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigSubjectConfigArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigSubjectConfigSubjectArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigCaOptionsArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigKeyUsageBaseKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigKeyUsageExtendedKeyUsageArgs;
 * import com.pulumi.gcp.certificateauthority.inputs.AuthorityConfigX509ConfigNameConstraintsArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var privatecaSa = new ServiceIdentity(&#34;privatecaSa&#34;, ServiceIdentityArgs.builder()        
 *             .service(&#34;privateca.googleapis.com&#34;)
 *             .build());
 * 
 *         var privatecaSaKeyuserSignerverifier = new CryptoKeyIAMBinding(&#34;privatecaSaKeyuserSignerverifier&#34;, CryptoKeyIAMBindingArgs.builder()        
 *             .cryptoKeyId(&#34;projects/keys-project/locations/us-central1/keyRings/key-ring/cryptoKeys/crypto-key&#34;)
 *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
 *             .members(privatecaSa.email().applyValue(email -&gt; String.format(&#34;serviceAccount:%s&#34;, email)))
 *             .build());
 * 
 *         var privatecaSaKeyuserViewer = new CryptoKeyIAMBinding(&#34;privatecaSaKeyuserViewer&#34;, CryptoKeyIAMBindingArgs.builder()        
 *             .cryptoKeyId(&#34;projects/keys-project/locations/us-central1/keyRings/key-ring/cryptoKeys/crypto-key&#34;)
 *             .role(&#34;roles/viewer&#34;)
 *             .members(privatecaSa.email().applyValue(email -&gt; String.format(&#34;serviceAccount:%s&#34;, email)))
 *             .build());
 * 
 *         var default_ = new Authority(&#34;default&#34;, AuthorityArgs.builder()        
 *             .pool(&#34;ca-pool&#34;)
 *             .certificateAuthorityId(&#34;my-certificate-authority&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .deletionProtection(&#34;true&#34;)
 *             .keySpec(AuthorityKeySpecArgs.builder()
 *                 .cloudKmsKeyVersion(&#34;projects/keys-project/locations/us-central1/keyRings/key-ring/cryptoKeys/crypto-key/cryptoKeyVersions/1&#34;)
 *                 .build())
 *             .config(AuthorityConfigArgs.builder()
 *                 .subjectConfig(AuthorityConfigSubjectConfigArgs.builder()
 *                     .subject(AuthorityConfigSubjectConfigSubjectArgs.builder()
 *                         .organization(&#34;Example, Org.&#34;)
 *                         .commonName(&#34;Example Authority&#34;)
 *                         .build())
 *                     .build())
 *                 .x509Config(AuthorityConfigX509ConfigArgs.builder()
 *                     .caOptions(AuthorityConfigX509ConfigCaOptionsArgs.builder()
 *                         .isCa(true)
 *                         .maxIssuerPathLength(10)
 *                         .build())
 *                     .keyUsage(AuthorityConfigX509ConfigKeyUsageArgs.builder()
 *                         .baseKeyUsage(AuthorityConfigX509ConfigKeyUsageBaseKeyUsageArgs.builder()
 *                             .certSign(true)
 *                             .crlSign(true)
 *                             .build())
 *                         .extendedKeyUsage(AuthorityConfigX509ConfigKeyUsageExtendedKeyUsageArgs.builder()
 *                             .serverAuth(false)
 *                             .build())
 *                         .build())
 *                     .nameConstraints(AuthorityConfigX509ConfigNameConstraintsArgs.builder()
 *                         .critical(true)
 *                         .permittedDnsNames(&#34;*.example.com&#34;)
 *                         .excludedDnsNames(&#34;*.deny.example.com&#34;)
 *                         .permittedIpRanges(&#34;10.0.0.0/8&#34;)
 *                         .excludedIpRanges(&#34;10.1.1.0/24&#34;)
 *                         .permittedEmailAddresses(&#34;.example.com&#34;)
 *                         .excludedEmailAddresses(&#34;.deny.example.com&#34;)
 *                         .permittedUris(&#34;.example.com&#34;)
 *                         .excludedUris(&#34;.deny.example.com&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     privatecaSaKeyuserSignerverifier,
 *                     privatecaSaKeyuserViewer)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * CertificateAuthority can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:certificateauthority/authority:Authority default projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificateAuthorities/{{certificate_authority_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:certificateauthority/authority:Authority default {{project}}/{{location}}/{{pool}}/{{certificate_authority_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:certificateauthority/authority:Authority default {{location}}/{{pool}}/{{certificate_authority_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:certificateauthority/authority:Authority")
public class Authority extends com.pulumi.resources.CustomResource {
    /**
     * URLs for accessing content published by this CA, such as the CA certificate and CRLs.
     * Structure is documented below.
     * 
     */
    @Export(name="accessUrls", refs={List.class,AuthorityAccessUrl.class}, tree="[0,1]")
    private Output<List<AuthorityAccessUrl>> accessUrls;

    /**
     * @return URLs for accessing content published by this CA, such as the CA certificate and CRLs.
     * Structure is documented below.
     * 
     */
    public Output<List<AuthorityAccessUrl>> accessUrls() {
        return this.accessUrls;
    }
    /**
     * The user provided Resource ID for this Certificate Authority.
     * 
     */
    @Export(name="certificateAuthorityId", refs={String.class}, tree="[0]")
    private Output<String> certificateAuthorityId;

    /**
     * @return The user provided Resource ID for this Certificate Authority.
     * 
     */
    public Output<String> certificateAuthorityId() {
        return this.certificateAuthorityId;
    }
    /**
     * The config used to create a self-signed X.509 certificate or CSR.
     * Structure is documented below.
     * 
     */
    @Export(name="config", refs={AuthorityConfig.class}, tree="[0]")
    private Output<AuthorityConfig> config;

    /**
     * @return The config used to create a self-signed X.509 certificate or CSR.
     * Structure is documented below.
     * 
     */
    public Output<AuthorityConfig> config() {
        return this.config;
    }
    /**
     * The time at which this CertificateAuthority was created.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine
     * fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The time at which this CertificateAuthority was created.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine
     * fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Whether or not to allow Terraform to destroy the CertificateAuthority. Unless this field is set to false in Terraform
     * state, a &#39;terraform destroy&#39; or &#39;terraform apply&#39; that would delete the instance will fail.
     * 
     */
    @Export(name="deletionProtection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deletionProtection;

    /**
     * @return Whether or not to allow Terraform to destroy the CertificateAuthority. Unless this field is set to false in Terraform
     * state, a &#39;terraform destroy&#39; or &#39;terraform apply&#39; that would delete the instance will fail.
     * 
     */
    public Output<Optional<Boolean>> deletionProtection() {
        return Codegen.optional(this.deletionProtection);
    }
    /**
     * Desired state of the CertificateAuthority. Set this field to `STAGED` to create a `STAGED` root CA.
     * 
     */
    @Export(name="desiredState", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> desiredState;

    /**
     * @return Desired state of the CertificateAuthority. Set this field to `STAGED` to create a `STAGED` root CA.
     * 
     */
    public Output<Optional<String>> desiredState() {
        return Codegen.optional(this.desiredState);
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
     * The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
     * such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
     * (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
     * my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
     * created.
     * 
     */
    @Export(name="gcsBucket", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> gcsBucket;

    /**
     * @return The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
     * such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
     * (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
     * my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
     * created.
     * 
     */
    public Output<Optional<String>> gcsBucket() {
        return Codegen.optional(this.gcsBucket);
    }
    /**
     * This field allows the CA to be deleted even if the CA has active certs. Active certs include both unrevoked and unexpired certs.
     * Use with care. Defaults to `false`.
     * 
     */
    @Export(name="ignoreActiveCertificatesOnDeletion", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ignoreActiveCertificatesOnDeletion;

    /**
     * @return This field allows the CA to be deleted even if the CA has active certs. Active certs include both unrevoked and unexpired certs.
     * Use with care. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> ignoreActiveCertificatesOnDeletion() {
        return Codegen.optional(this.ignoreActiveCertificatesOnDeletion);
    }
    /**
     * Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
     * is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
     * certificate. Otherwise, it is used to sign a CSR.
     * Structure is documented below.
     * 
     */
    @Export(name="keySpec", refs={AuthorityKeySpec.class}, tree="[0]")
    private Output<AuthorityKeySpec> keySpec;

    /**
     * @return Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
     * is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
     * certificate. Otherwise, it is used to sign a CSR.
     * Structure is documented below.
     * 
     */
    public Output<AuthorityKeySpec> keySpec() {
        return this.keySpec;
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
     * The desired lifetime of the CA certificate. Used to create the &#34;notBeforeTime&#34; and
     * &#34;notAfterTime&#34; fields inside an X.509 certificate. A duration in seconds with up to nine
     * fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    @Export(name="lifetime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lifetime;

    /**
     * @return The desired lifetime of the CA certificate. Used to create the &#34;notBeforeTime&#34; and
     * &#34;notAfterTime&#34; fields inside an X.509 certificate. A duration in seconds with up to nine
     * fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    public Output<Optional<String>> lifetime() {
        return Codegen.optional(this.lifetime);
    }
    /**
     * Location of the CertificateAuthority. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return Location of the CertificateAuthority. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The resource name for this CertificateAuthority in the format
     * projects/*{@literal /}locations/*{@literal /}certificateAuthorities/*.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name for this CertificateAuthority in the format
     * projects/*{@literal /}locations/*{@literal /}certificateAuthorities/*.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The signed CA certificate issued from the subordinated CA&#39;s CSR. This is needed when activating the subordiante CA with a third party issuer.
     * 
     */
    @Export(name="pemCaCertificate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pemCaCertificate;

    /**
     * @return The signed CA certificate issued from the subordinated CA&#39;s CSR. This is needed when activating the subordiante CA with a third party issuer.
     * 
     */
    public Output<Optional<String>> pemCaCertificate() {
        return Codegen.optional(this.pemCaCertificate);
    }
    /**
     * This CertificateAuthority&#39;s certificate chain, including the current
     * CertificateAuthority&#39;s certificate. Ordered such that the root issuer is the final
     * element (consistent with RFC 5246). For a self-signed CA, this will only list the current
     * CertificateAuthority&#39;s certificate.
     * 
     */
    @Export(name="pemCaCertificates", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> pemCaCertificates;

    /**
     * @return This CertificateAuthority&#39;s certificate chain, including the current
     * CertificateAuthority&#39;s certificate. Ordered such that the root issuer is the final
     * element (consistent with RFC 5246). For a self-signed CA, this will only list the current
     * CertificateAuthority&#39;s certificate.
     * 
     */
    public Output<List<String>> pemCaCertificates() {
        return this.pemCaCertificates;
    }
    /**
     * The name of the CaPool this Certificate Authority belongs to.
     * 
     */
    @Export(name="pool", refs={String.class}, tree="[0]")
    private Output<String> pool;

    /**
     * @return The name of the CaPool this Certificate Authority belongs to.
     * 
     */
    public Output<String> pool() {
        return this.pool;
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
     * If this flag is set, the Certificate Authority will be deleted as soon as
     * possible without a 30-day grace period where undeletion would have been
     * allowed. If you proceed, there will be no way to recover this CA.
     * Use with care. Defaults to `false`.
     * 
     */
    @Export(name="skipGracePeriod", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> skipGracePeriod;

    /**
     * @return If this flag is set, the Certificate Authority will be deleted as soon as
     * possible without a 30-day grace period where undeletion would have been
     * allowed. If you proceed, there will be no way to recover this CA.
     * Use with care. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> skipGracePeriod() {
        return Codegen.optional(this.skipGracePeriod);
    }
    /**
     * The State for this CertificateAuthority.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The State for this CertificateAuthority.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * If this is a subordinate CertificateAuthority, this field will be set
     * with the subordinate configuration, which describes its issuers.
     * Structure is documented below.
     * 
     */
    @Export(name="subordinateConfig", refs={AuthoritySubordinateConfig.class}, tree="[0]")
    private Output</* @Nullable */ AuthoritySubordinateConfig> subordinateConfig;

    /**
     * @return If this is a subordinate CertificateAuthority, this field will be set
     * with the subordinate configuration, which describes its issuers.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AuthoritySubordinateConfig>> subordinateConfig() {
        return Codegen.optional(this.subordinateConfig);
    }
    /**
     * The Type of this CertificateAuthority.
     * &gt; **Note:** For `SUBORDINATE` Certificate Authorities, they need to
     * be activated before they can issue certificates.
     * Default value is `SELF_SIGNED`.
     * Possible values are: `SELF_SIGNED`, `SUBORDINATE`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The Type of this CertificateAuthority.
     * &gt; **Note:** For `SUBORDINATE` Certificate Authorities, they need to
     * be activated before they can issue certificates.
     * Default value is `SELF_SIGNED`.
     * Possible values are: `SELF_SIGNED`, `SUBORDINATE`.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * The time at which this CertificateAuthority was updated.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine
     * fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The time at which this CertificateAuthority was updated.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine
     * fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Authority(String name) {
        this(name, AuthorityArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Authority(String name, AuthorityArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Authority(String name, AuthorityArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificateauthority/authority:Authority", name, args == null ? AuthorityArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Authority(String name, Output<String> id, @Nullable AuthorityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificateauthority/authority:Authority", name, state, makeResourceOptions(options, id));
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
    public static Authority get(String name, Output<String> id, @Nullable AuthorityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Authority(name, id, state, options);
    }
}

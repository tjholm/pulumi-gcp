// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificatemanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.certificatemanager.CertificateArgs;
import com.pulumi.gcp.certificatemanager.inputs.CertificateState;
import com.pulumi.gcp.certificatemanager.outputs.CertificateManaged;
import com.pulumi.gcp.certificatemanager.outputs.CertificateSelfManaged;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Certificate represents a HTTP-reachable backend for a Certificate.
 * 
 * ## Example Usage
 * ### Certificate Manager Google Managed Certificate Dns
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificatemanager.DnsAuthorization;
 * import com.pulumi.gcp.certificatemanager.DnsAuthorizationArgs;
 * import com.pulumi.gcp.certificatemanager.Certificate;
 * import com.pulumi.gcp.certificatemanager.CertificateArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateManagedArgs;
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
 *         var instance = new DnsAuthorization(&#34;instance&#34;, DnsAuthorizationArgs.builder()        
 *             .description(&#34;The default dnss&#34;)
 *             .domain(&#34;subdomain.hashicorptest.com&#34;)
 *             .build());
 * 
 *         var instance2 = new DnsAuthorization(&#34;instance2&#34;, DnsAuthorizationArgs.builder()        
 *             .description(&#34;The default dnss&#34;)
 *             .domain(&#34;subdomain2.hashicorptest.com&#34;)
 *             .build());
 * 
 *         var default_ = new Certificate(&#34;default&#34;, CertificateArgs.builder()        
 *             .description(&#34;The default cert&#34;)
 *             .scope(&#34;EDGE_CACHE&#34;)
 *             .labels(Map.of(&#34;env&#34;, &#34;test&#34;))
 *             .managed(CertificateManagedArgs.builder()
 *                 .domains(                
 *                     instance.domain(),
 *                     instance2.domain())
 *                 .dnsAuthorizations(                
 *                     instance.id(),
 *                     instance2.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Certificate Manager Google Managed Certificate Issuance Config
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificateauthority.CaPool;
 * import com.pulumi.gcp.certificateauthority.CaPoolArgs;
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
 * import com.pulumi.gcp.certificatemanager.CertificateIssuanceConfig;
 * import com.pulumi.gcp.certificatemanager.CertificateIssuanceConfigArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateIssuanceConfigCertificateAuthorityConfigArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigArgs;
 * import com.pulumi.gcp.certificatemanager.Certificate;
 * import com.pulumi.gcp.certificatemanager.CertificateArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateManagedArgs;
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
 *         var pool = new CaPool(&#34;pool&#34;, CaPoolArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .tier(&#34;ENTERPRISE&#34;)
 *             .build());
 * 
 *         var caAuthority = new Authority(&#34;caAuthority&#34;, AuthorityArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .pool(pool.name())
 *             .certificateAuthorityId(&#34;ca-authority&#34;)
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
 *                             .serverAuth(true)
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
 *         var issuanceconfig = new CertificateIssuanceConfig(&#34;issuanceconfig&#34;, CertificateIssuanceConfigArgs.builder()        
 *             .description(&#34;sample description for the certificate issuanceConfigs&#34;)
 *             .certificateAuthorityConfig(CertificateIssuanceConfigCertificateAuthorityConfigArgs.builder()
 *                 .certificateAuthorityServiceConfig(CertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigArgs.builder()
 *                     .caPool(pool.id())
 *                     .build())
 *                 .build())
 *             .lifetime(&#34;1814400s&#34;)
 *             .rotationWindowPercentage(34)
 *             .keyAlgorithm(&#34;ECDSA_P256&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(caAuthority)
 *                 .build());
 * 
 *         var default_ = new Certificate(&#34;default&#34;, CertificateArgs.builder()        
 *             .description(&#34;The default cert&#34;)
 *             .scope(&#34;EDGE_CACHE&#34;)
 *             .managed(CertificateManagedArgs.builder()
 *                 .domains(&#34;terraform.subdomain1.com&#34;)
 *                 .issuanceConfig(issuanceconfig.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Certificate Manager Certificate Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificatemanager.DnsAuthorization;
 * import com.pulumi.gcp.certificatemanager.DnsAuthorizationArgs;
 * import com.pulumi.gcp.certificatemanager.Certificate;
 * import com.pulumi.gcp.certificatemanager.CertificateArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateManagedArgs;
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
 *         var instance = new DnsAuthorization(&#34;instance&#34;, DnsAuthorizationArgs.builder()        
 *             .description(&#34;The default dnss&#34;)
 *             .domain(&#34;subdomain.hashicorptest.com&#34;)
 *             .build());
 * 
 *         var instance2 = new DnsAuthorization(&#34;instance2&#34;, DnsAuthorizationArgs.builder()        
 *             .description(&#34;The default dnss&#34;)
 *             .domain(&#34;subdomain2.hashicorptest.com&#34;)
 *             .build());
 * 
 *         var default_ = new Certificate(&#34;default&#34;, CertificateArgs.builder()        
 *             .description(&#34;Global cert&#34;)
 *             .scope(&#34;EDGE_CACHE&#34;)
 *             .managed(CertificateManagedArgs.builder()
 *                 .domains(                
 *                     instance.domain(),
 *                     instance2.domain())
 *                 .dnsAuthorizations(                
 *                     instance.id(),
 *                     instance2.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Certificate Manager Self Managed Certificate Regional
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificatemanager.Certificate;
 * import com.pulumi.gcp.certificatemanager.CertificateArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateSelfManagedArgs;
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
 *         var default_ = new Certificate(&#34;default&#34;, CertificateArgs.builder()        
 *             .description(&#34;Regional cert&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .selfManaged(CertificateSelfManagedArgs.builder()
 *                 .pemCertificate(Files.readString(Paths.get(&#34;test-fixtures/cert.pem&#34;)))
 *                 .pemPrivateKey(Files.readString(Paths.get(&#34;test-fixtures/private-key.pem&#34;)))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Certificate Manager Google Managed Certificate Issuance Config All Regions
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificateauthority.CaPool;
 * import com.pulumi.gcp.certificateauthority.CaPoolArgs;
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
 * import com.pulumi.gcp.certificatemanager.CertificateIssuanceConfig;
 * import com.pulumi.gcp.certificatemanager.CertificateIssuanceConfigArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateIssuanceConfigCertificateAuthorityConfigArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigArgs;
 * import com.pulumi.gcp.certificatemanager.Certificate;
 * import com.pulumi.gcp.certificatemanager.CertificateArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateManagedArgs;
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
 *         var pool = new CaPool(&#34;pool&#34;, CaPoolArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .tier(&#34;ENTERPRISE&#34;)
 *             .build());
 * 
 *         var caAuthority = new Authority(&#34;caAuthority&#34;, AuthorityArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .pool(pool.name())
 *             .certificateAuthorityId(&#34;ca-authority&#34;)
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
 *                             .serverAuth(true)
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
 *         var issuanceconfig = new CertificateIssuanceConfig(&#34;issuanceconfig&#34;, CertificateIssuanceConfigArgs.builder()        
 *             .description(&#34;sample description for the certificate issuanceConfigs&#34;)
 *             .certificateAuthorityConfig(CertificateIssuanceConfigCertificateAuthorityConfigArgs.builder()
 *                 .certificateAuthorityServiceConfig(CertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigArgs.builder()
 *                     .caPool(pool.id())
 *                     .build())
 *                 .build())
 *             .lifetime(&#34;1814400s&#34;)
 *             .rotationWindowPercentage(34)
 *             .keyAlgorithm(&#34;ECDSA_P256&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(caAuthority)
 *                 .build());
 * 
 *         var default_ = new Certificate(&#34;default&#34;, CertificateArgs.builder()        
 *             .description(&#34;sample google managed all_regions certificate with issuance config for terraform&#34;)
 *             .scope(&#34;ALL_REGIONS&#34;)
 *             .managed(CertificateManagedArgs.builder()
 *                 .domains(&#34;terraform.subdomain1.com&#34;)
 *                 .issuanceConfig(issuanceconfig.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Certificate Manager Google Managed Certificate Dns All Regions
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificatemanager.DnsAuthorization;
 * import com.pulumi.gcp.certificatemanager.DnsAuthorizationArgs;
 * import com.pulumi.gcp.certificatemanager.Certificate;
 * import com.pulumi.gcp.certificatemanager.CertificateArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateManagedArgs;
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
 *         var instance = new DnsAuthorization(&#34;instance&#34;, DnsAuthorizationArgs.builder()        
 *             .description(&#34;The default dnss&#34;)
 *             .domain(&#34;subdomain.hashicorptest.com&#34;)
 *             .build());
 * 
 *         var instance2 = new DnsAuthorization(&#34;instance2&#34;, DnsAuthorizationArgs.builder()        
 *             .description(&#34;The default dnss&#34;)
 *             .domain(&#34;subdomain2.hashicorptest.com&#34;)
 *             .build());
 * 
 *         var default_ = new Certificate(&#34;default&#34;, CertificateArgs.builder()        
 *             .description(&#34;The default cert&#34;)
 *             .scope(&#34;ALL_REGIONS&#34;)
 *             .managed(CertificateManagedArgs.builder()
 *                 .domains(                
 *                     instance.domain(),
 *                     instance2.domain())
 *                 .dnsAuthorizations(                
 *                     instance.id(),
 *                     instance2.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Certificate can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/certificates/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Certificate can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:certificatemanager/certificate:Certificate default projects/{{project}}/locations/{{location}}/certificates/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:certificatemanager/certificate:Certificate default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:certificatemanager/certificate:Certificate default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:certificatemanager/certificate:Certificate")
public class Certificate extends com.pulumi.resources.CustomResource {
    /**
     * A human-readable description of the resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A human-readable description of the resource.
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
     * Set of label tags associated with the Certificate resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Set of label tags associated with the Certificate resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
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
     * Configuration and state of a Managed Certificate.
     * Certificate Manager provisions and renews Managed Certificates
     * automatically, for as long as it&#39;s authorized to do so.
     * Structure is documented below.
     * 
     */
    @Export(name="managed", refs={CertificateManaged.class}, tree="[0]")
    private Output</* @Nullable */ CertificateManaged> managed;

    /**
     * @return Configuration and state of a Managed Certificate.
     * Certificate Manager provisions and renews Managed Certificates
     * automatically, for as long as it&#39;s authorized to do so.
     * Structure is documented below.
     * 
     */
    public Output<Optional<CertificateManaged>> managed() {
        return Codegen.optional(this.managed);
    }
    /**
     * A user-defined name of the certificate. Certificate names must be unique
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A user-defined name of the certificate. Certificate names must be unique
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     * 
     * ***
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
     * The scope of the certificate.
     * DEFAULT: Certificates with default scope are served from core Google data centers.
     * If unsure, choose this option.
     * EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served from Edge Points of Presence.
     * See https://cloud.google.com/vpc/docs/edge-locations.
     * ALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).
     * See https://cloud.google.com/compute/docs/regions-zones
     * 
     */
    @Export(name="scope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scope;

    /**
     * @return The scope of the certificate.
     * DEFAULT: Certificates with default scope are served from core Google data centers.
     * If unsure, choose this option.
     * EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served from Edge Points of Presence.
     * See https://cloud.google.com/vpc/docs/edge-locations.
     * ALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).
     * See https://cloud.google.com/compute/docs/regions-zones
     * 
     */
    public Output<Optional<String>> scope() {
        return Codegen.optional(this.scope);
    }
    /**
     * Certificate data for a SelfManaged Certificate.
     * SelfManaged Certificates are uploaded by the user. Updating such
     * certificates before they expire remains the user&#39;s responsibility.
     * Structure is documented below.
     * 
     */
    @Export(name="selfManaged", refs={CertificateSelfManaged.class}, tree="[0]")
    private Output</* @Nullable */ CertificateSelfManaged> selfManaged;

    /**
     * @return Certificate data for a SelfManaged Certificate.
     * SelfManaged Certificates are uploaded by the user. Updating such
     * certificates before they expire remains the user&#39;s responsibility.
     * Structure is documented below.
     * 
     */
    public Output<Optional<CertificateSelfManaged>> selfManaged() {
        return Codegen.optional(this.selfManaged);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Certificate(String name) {
        this(name, CertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Certificate(String name, @Nullable CertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Certificate(String name, @Nullable CertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificatemanager/certificate:Certificate", name, args == null ? CertificateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Certificate(String name, Output<String> id, @Nullable CertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificatemanager/certificate:Certificate", name, state, makeResourceOptions(options, id));
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
    public static Certificate get(String name, Output<String> id, @Nullable CertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Certificate(name, id, state, options);
    }
}

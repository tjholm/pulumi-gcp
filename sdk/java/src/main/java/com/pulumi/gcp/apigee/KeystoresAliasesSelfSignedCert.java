// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigee;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigee.KeystoresAliasesSelfSignedCertArgs;
import com.pulumi.gcp.apigee.inputs.KeystoresAliasesSelfSignedCertState;
import com.pulumi.gcp.apigee.outputs.KeystoresAliasesSelfSignedCertCertsInfo;
import com.pulumi.gcp.apigee.outputs.KeystoresAliasesSelfSignedCertSubject;
import com.pulumi.gcp.apigee.outputs.KeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An Environment Keystore Alias for Self Signed Certificate Format in Apigee
 * 
 * To get more information about KeystoresAliasesSelfSignedCert, see:
 * 
 * * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.keystores.aliases/create)
 * * How-to Guides
 *     * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)
 * 
 * ## Example Usage
 * 
 * ### Apigee Env Keystore Alias Self Signed Cert
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.projects.Service;
 * import com.pulumi.gcp.projects.ServiceArgs;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.apigee.Organization;
 * import com.pulumi.gcp.apigee.OrganizationArgs;
 * import com.pulumi.gcp.apigee.Environment;
 * import com.pulumi.gcp.apigee.EnvironmentArgs;
 * import com.pulumi.gcp.apigee.EnvKeystore;
 * import com.pulumi.gcp.apigee.EnvKeystoreArgs;
 * import com.pulumi.gcp.apigee.KeystoresAliasesSelfSignedCert;
 * import com.pulumi.gcp.apigee.KeystoresAliasesSelfSignedCertArgs;
 * import com.pulumi.gcp.apigee.inputs.KeystoresAliasesSelfSignedCertSubjectArgs;
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
 *         var project = new Project("project", ProjectArgs.builder()
 *             .projectId("my-project")
 *             .name("my-project")
 *             .orgId("123456789")
 *             .billingAccount("000000-0000000-0000000-000000")
 *             .build());
 * 
 *         var apigee = new Service("apigee", ServiceArgs.builder()
 *             .project(project.projectId())
 *             .service("apigee.googleapis.com")
 *             .build());
 * 
 *         var servicenetworking = new Service("servicenetworking", ServiceArgs.builder()
 *             .project(project.projectId())
 *             .service("servicenetworking.googleapis.com")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(apigee)
 *                 .build());
 * 
 *         var compute = new Service("compute", ServiceArgs.builder()
 *             .project(project.projectId())
 *             .service("compute.googleapis.com")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(servicenetworking)
 *                 .build());
 * 
 *         var apigeeNetwork = new Network("apigeeNetwork", NetworkArgs.builder()
 *             .name("apigee-network")
 *             .project(project.projectId())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(compute)
 *                 .build());
 * 
 *         var apigeeRange = new GlobalAddress("apigeeRange", GlobalAddressArgs.builder()
 *             .name("apigee-range")
 *             .purpose("VPC_PEERING")
 *             .addressType("INTERNAL")
 *             .prefixLength(16)
 *             .network(apigeeNetwork.id())
 *             .project(project.projectId())
 *             .build());
 * 
 *         var apigeeVpcConnection = new Connection("apigeeVpcConnection", ConnectionArgs.builder()
 *             .network(apigeeNetwork.id())
 *             .service("servicenetworking.googleapis.com")
 *             .reservedPeeringRanges(apigeeRange.name())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(servicenetworking)
 *                 .build());
 * 
 *         var apigeeOrg = new Organization("apigeeOrg", OrganizationArgs.builder()
 *             .analyticsRegion("us-central1")
 *             .projectId(project.projectId())
 *             .authorizedNetwork(apigeeNetwork.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     apigeeVpcConnection,
 *                     apigee)
 *                 .build());
 * 
 *         var apigeeEnvironmentKeystoreSsAlias = new Environment("apigeeEnvironmentKeystoreSsAlias", EnvironmentArgs.builder()
 *             .orgId(apigeeOrg.id())
 *             .name("env-name")
 *             .description("Apigee Environment")
 *             .displayName("environment-1")
 *             .build());
 * 
 *         var apigeeEnvironmentKeystoreAlias = new EnvKeystore("apigeeEnvironmentKeystoreAlias", EnvKeystoreArgs.builder()
 *             .name("env-keystore")
 *             .envId(apigeeEnvironmentKeystoreSsAlias.id())
 *             .build());
 * 
 *         var apigeeEnvironmentKeystoreSsAliasKeystoresAliasesSelfSignedCert = new KeystoresAliasesSelfSignedCert("apigeeEnvironmentKeystoreSsAliasKeystoresAliasesSelfSignedCert", KeystoresAliasesSelfSignedCertArgs.builder()
 *             .environment(apigeeEnvironmentKeystoreSsAlias.name())
 *             .orgId(apigeeOrg.name())
 *             .keystore(apigeeEnvironmentKeystoreAlias.name())
 *             .alias("alias")
 *             .keySize(1024)
 *             .sigAlg("SHA512withRSA")
 *             .certValidityInDays(4)
 *             .subject(KeystoresAliasesSelfSignedCertSubjectArgs.builder()
 *                 .commonName("selfsigned_example")
 *                 .countryCode("US")
 *                 .locality("TX")
 *                 .org("CCE")
 *                 .orgUnit("PSO")
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
 * KeystoresAliasesSelfSignedCert can be imported using any of these accepted formats:
 * 
 * * `organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}`
 * 
 * * `{{org_id}}/{{environment}}/{{keystore}}/{{alias}}`
 * 
 * When using the `pulumi import` command, KeystoresAliasesSelfSignedCert can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:apigee/keystoresAliasesSelfSignedCert:KeystoresAliasesSelfSignedCert default organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:apigee/keystoresAliasesSelfSignedCert:KeystoresAliasesSelfSignedCert default {{org_id}}/{{environment}}/{{keystore}}/{{alias}}
 * ```
 * 
 */
@ResourceType(type="gcp:apigee/keystoresAliasesSelfSignedCert:KeystoresAliasesSelfSignedCert")
public class KeystoresAliasesSelfSignedCert extends com.pulumi.resources.CustomResource {
    /**
     * Alias for the key/certificate pair. Values must match the regular expression [\w\s-.]{1,255}.
     * This must be provided for all formats except selfsignedcert; self-signed certs may specify the alias in either
     * this parameter or the JSON body.
     * 
     */
    @Export(name="alias", refs={String.class}, tree="[0]")
    private Output<String> alias;

    /**
     * @return Alias for the key/certificate pair. Values must match the regular expression [\w\s-.]{1,255}.
     * This must be provided for all formats except selfsignedcert; self-signed certs may specify the alias in either
     * this parameter or the JSON body.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }
    /**
     * Validity duration of certificate, in days. Accepts positive non-zero value. Defaults to 365.
     * 
     */
    @Export(name="certValidityInDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> certValidityInDays;

    /**
     * @return Validity duration of certificate, in days. Accepts positive non-zero value. Defaults to 365.
     * 
     */
    public Output<Optional<Integer>> certValidityInDays() {
        return Codegen.optional(this.certValidityInDays);
    }
    /**
     * Chain of certificates under this alias.
     * Structure is documented below.
     * 
     */
    @Export(name="certsInfos", refs={List.class,KeystoresAliasesSelfSignedCertCertsInfo.class}, tree="[0,1]")
    private Output<List<KeystoresAliasesSelfSignedCertCertsInfo>> certsInfos;

    /**
     * @return Chain of certificates under this alias.
     * Structure is documented below.
     * 
     */
    public Output<List<KeystoresAliasesSelfSignedCertCertsInfo>> certsInfos() {
        return this.certsInfos;
    }
    /**
     * The Apigee environment name
     * 
     */
    @Export(name="environment", refs={String.class}, tree="[0]")
    private Output<String> environment;

    /**
     * @return The Apigee environment name
     * 
     */
    public Output<String> environment() {
        return this.environment;
    }
    /**
     * Key size. Default and maximum value is 2048 bits.
     * 
     */
    @Export(name="keySize", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> keySize;

    /**
     * @return Key size. Default and maximum value is 2048 bits.
     * 
     */
    public Output<Optional<String>> keySize() {
        return Codegen.optional(this.keySize);
    }
    /**
     * The Apigee keystore name associated in an Apigee environment
     * 
     */
    @Export(name="keystore", refs={String.class}, tree="[0]")
    private Output<String> keystore;

    /**
     * @return The Apigee keystore name associated in an Apigee environment
     * 
     */
    public Output<String> keystore() {
        return this.keystore;
    }
    /**
     * The Apigee Organization name associated with the Apigee environment
     * 
     */
    @Export(name="orgId", refs={String.class}, tree="[0]")
    private Output<String> orgId;

    /**
     * @return The Apigee Organization name associated with the Apigee environment
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }
    /**
     * Signature algorithm to generate private key. Valid values are SHA512withRSA, SHA384withRSA, and SHA256withRSA
     * 
     */
    @Export(name="sigAlg", refs={String.class}, tree="[0]")
    private Output<String> sigAlg;

    /**
     * @return Signature algorithm to generate private key. Valid values are SHA512withRSA, SHA384withRSA, and SHA256withRSA
     * 
     */
    public Output<String> sigAlg() {
        return this.sigAlg;
    }
    /**
     * Subject details.
     * Structure is documented below.
     * 
     */
    @Export(name="subject", refs={KeystoresAliasesSelfSignedCertSubject.class}, tree="[0]")
    private Output<KeystoresAliasesSelfSignedCertSubject> subject;

    /**
     * @return Subject details.
     * Structure is documented below.
     * 
     */
    public Output<KeystoresAliasesSelfSignedCertSubject> subject() {
        return this.subject;
    }
    /**
     * List of alternative host names. Maximum length is 255 characters for each value.
     * 
     */
    @Export(name="subjectAlternativeDnsNames", refs={KeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames.class}, tree="[0]")
    private Output</* @Nullable */ KeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames> subjectAlternativeDnsNames;

    /**
     * @return List of alternative host names. Maximum length is 255 characters for each value.
     * 
     */
    public Output<Optional<KeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames>> subjectAlternativeDnsNames() {
        return Codegen.optional(this.subjectAlternativeDnsNames);
    }
    /**
     * Optional.Type of Alias
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Optional.Type of Alias
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KeystoresAliasesSelfSignedCert(java.lang.String name) {
        this(name, KeystoresAliasesSelfSignedCertArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KeystoresAliasesSelfSignedCert(java.lang.String name, KeystoresAliasesSelfSignedCertArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KeystoresAliasesSelfSignedCert(java.lang.String name, KeystoresAliasesSelfSignedCertArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/keystoresAliasesSelfSignedCert:KeystoresAliasesSelfSignedCert", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private KeystoresAliasesSelfSignedCert(java.lang.String name, Output<java.lang.String> id, @Nullable KeystoresAliasesSelfSignedCertState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/keystoresAliasesSelfSignedCert:KeystoresAliasesSelfSignedCert", name, state, makeResourceOptions(options, id), false);
    }

    private static KeystoresAliasesSelfSignedCertArgs makeArgs(KeystoresAliasesSelfSignedCertArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? KeystoresAliasesSelfSignedCertArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static KeystoresAliasesSelfSignedCert get(java.lang.String name, Output<java.lang.String> id, @Nullable KeystoresAliasesSelfSignedCertState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KeystoresAliasesSelfSignedCert(name, id, state, options);
    }
}

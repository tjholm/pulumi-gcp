// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigee;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigee.OrganizationArgs;
import com.pulumi.gcp.apigee.inputs.OrganizationState;
import com.pulumi.gcp.apigee.outputs.OrganizationProperties;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An `Organization` is the top-level container in Apigee.
 * 
 * To get more information about Organization, see:
 * 
 * * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations)
 * * How-to Guides
 *     * [Creating an API organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org)
 * 
 * ## Example Usage
 * ### Apigee Organization Cloud Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.apigee.Organization;
 * import com.pulumi.gcp.apigee.OrganizationArgs;
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
 *         final var current = OrganizationsFunctions.getClientConfig();
 * 
 *         var apigeeNetwork = new Network(&#34;apigeeNetwork&#34;);
 * 
 *         var apigeeRange = new GlobalAddress(&#34;apigeeRange&#34;, GlobalAddressArgs.builder()        
 *             .purpose(&#34;VPC_PEERING&#34;)
 *             .addressType(&#34;INTERNAL&#34;)
 *             .prefixLength(16)
 *             .network(apigeeNetwork.id())
 *             .build());
 * 
 *         var apigeeVpcConnection = new Connection(&#34;apigeeVpcConnection&#34;, ConnectionArgs.builder()        
 *             .network(apigeeNetwork.id())
 *             .service(&#34;servicenetworking.googleapis.com&#34;)
 *             .reservedPeeringRanges(apigeeRange.name())
 *             .build());
 * 
 *         var org = new Organization(&#34;org&#34;, OrganizationArgs.builder()        
 *             .analyticsRegion(&#34;us-central1&#34;)
 *             .projectId(current.applyValue(getClientConfigResult -&gt; getClientConfigResult.project()))
 *             .authorizedNetwork(apigeeNetwork.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(apigeeVpcConnection)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Apigee Organization Cloud Full
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.kms.CryptoKey;
 * import com.pulumi.gcp.kms.CryptoKeyArgs;
 * import com.pulumi.gcp.projects.ServiceIdentity;
 * import com.pulumi.gcp.projects.ServiceIdentityArgs;
 * import com.pulumi.gcp.kms.CryptoKeyIAMBinding;
 * import com.pulumi.gcp.kms.CryptoKeyIAMBindingArgs;
 * import com.pulumi.gcp.apigee.Organization;
 * import com.pulumi.gcp.apigee.OrganizationArgs;
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
 *         final var current = OrganizationsFunctions.getClientConfig();
 * 
 *         var apigeeNetwork = new Network(&#34;apigeeNetwork&#34;);
 * 
 *         var apigeeRange = new GlobalAddress(&#34;apigeeRange&#34;, GlobalAddressArgs.builder()        
 *             .purpose(&#34;VPC_PEERING&#34;)
 *             .addressType(&#34;INTERNAL&#34;)
 *             .prefixLength(16)
 *             .network(apigeeNetwork.id())
 *             .build());
 * 
 *         var apigeeVpcConnection = new Connection(&#34;apigeeVpcConnection&#34;, ConnectionArgs.builder()        
 *             .network(apigeeNetwork.id())
 *             .service(&#34;servicenetworking.googleapis.com&#34;)
 *             .reservedPeeringRanges(apigeeRange.name())
 *             .build());
 * 
 *         var apigeeKeyring = new KeyRing(&#34;apigeeKeyring&#34;, KeyRingArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .build());
 * 
 *         var apigeeKey = new CryptoKey(&#34;apigeeKey&#34;, CryptoKeyArgs.builder()        
 *             .keyRing(apigeeKeyring.id())
 *             .build());
 * 
 *         var apigeeSa = new ServiceIdentity(&#34;apigeeSa&#34;, ServiceIdentityArgs.builder()        
 *             .project(google_project.project().project_id())
 *             .service(google_project_service.apigee().service())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var apigeeSaKeyuser = new CryptoKeyIAMBinding(&#34;apigeeSaKeyuser&#34;, CryptoKeyIAMBindingArgs.builder()        
 *             .cryptoKeyId(apigeeKey.id())
 *             .role(&#34;roles/cloudkms.cryptoKeyEncrypterDecrypter&#34;)
 *             .members(apigeeSa.email().applyValue(email -&gt; String.format(&#34;serviceAccount:%s&#34;, email)))
 *             .build());
 * 
 *         var org = new Organization(&#34;org&#34;, OrganizationArgs.builder()        
 *             .analyticsRegion(&#34;us-central1&#34;)
 *             .displayName(&#34;apigee-org&#34;)
 *             .description(&#34;Auto-provisioned Apigee Org.&#34;)
 *             .projectId(current.applyValue(getClientConfigResult -&gt; getClientConfigResult.project()))
 *             .authorizedNetwork(apigeeNetwork.id())
 *             .runtimeDatabaseEncryptionKeyName(apigeeKey.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     apigeeVpcConnection,
 *                     apigeeSaKeyuser)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Organization can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:apigee/organization:Organization default organizations/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:apigee/organization:Organization default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:apigee/organization:Organization")
public class Organization extends com.pulumi.resources.CustomResource {
    /**
     * Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
     * 
     */
    @Export(name="analyticsRegion", type=String.class, parameters={})
    private Output</* @Nullable */ String> analyticsRegion;

    /**
     * @return Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
     * 
     */
    public Output<Optional<String>> analyticsRegion() {
        return Codegen.optional(this.analyticsRegion);
    }
    /**
     * Output only. Project ID of the Apigee Tenant Project.
     * 
     */
    @Export(name="apigeeProjectId", type=String.class, parameters={})
    private Output<String> apigeeProjectId;

    /**
     * @return Output only. Project ID of the Apigee Tenant Project.
     * 
     */
    public Output<String> apigeeProjectId() {
        return this.apigeeProjectId;
    }
    /**
     * Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
     * See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
     * Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: &#34;default&#34;.
     * 
     */
    @Export(name="authorizedNetwork", type=String.class, parameters={})
    private Output</* @Nullable */ String> authorizedNetwork;

    /**
     * @return Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
     * See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
     * Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: &#34;default&#34;.
     * 
     */
    public Output<Optional<String>> authorizedNetwork() {
        return Codegen.optional(this.authorizedNetwork);
    }
    /**
     * Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
     * 
     */
    @Export(name="billingType", type=String.class, parameters={})
    private Output<String> billingType;

    /**
     * @return Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
     * 
     */
    public Output<String> billingType() {
        return this.billingType;
    }
    /**
     * Output only. Base64-encoded public certificate for the root CA of the Apigee organization.
     * Valid only when `RuntimeType` is CLOUD. A base64-encoded string.
     * 
     */
    @Export(name="caCertificate", type=String.class, parameters={})
    private Output<String> caCertificate;

    /**
     * @return Output only. Base64-encoded public certificate for the root CA of the Apigee organization.
     * Valid only when `RuntimeType` is CLOUD. A base64-encoded string.
     * 
     */
    public Output<String> caCertificate() {
        return this.caCertificate;
    }
    /**
     * Description of the Apigee organization.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the Apigee organization.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The display name of the Apigee organization.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output</* @Nullable */ String> displayName;

    /**
     * @return The display name of the Apigee organization.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Name of the property.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the property.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The project ID associated with the Apigee organization.
     * 
     * ***
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The project ID associated with the Apigee organization.
     * 
     * ***
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Properties defined in the Apigee organization profile.
     * Structure is documented below.
     * 
     */
    @Export(name="properties", type=OrganizationProperties.class, parameters={})
    private Output<OrganizationProperties> properties;

    /**
     * @return Properties defined in the Apigee organization profile.
     * Structure is documented below.
     * 
     */
    public Output<OrganizationProperties> properties() {
        return this.properties;
    }
    /**
     * Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
     * is not EVALUATION). It controls how long Organization data will be retained after the initial delete
     * operation completes. During this period, the Organization may be restored to its last known state.
     * After this period, the Organization will no longer be able to be restored.
     * Default value is `DELETION_RETENTION_UNSPECIFIED`.
     * Possible values are: `DELETION_RETENTION_UNSPECIFIED`, `MINIMUM`.
     * 
     */
    @Export(name="retention", type=String.class, parameters={})
    private Output</* @Nullable */ String> retention;

    /**
     * @return Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
     * is not EVALUATION). It controls how long Organization data will be retained after the initial delete
     * operation completes. During this period, the Organization may be restored to its last known state.
     * After this period, the Organization will no longer be able to be restored.
     * Default value is `DELETION_RETENTION_UNSPECIFIED`.
     * Possible values are: `DELETION_RETENTION_UNSPECIFIED`, `MINIMUM`.
     * 
     */
    public Output<Optional<String>> retention() {
        return Codegen.optional(this.retention);
    }
    /**
     * Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
     * Update is not allowed after the organization is created.
     * If not specified, a Google-Managed encryption key will be used.
     * Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
     * 
     */
    @Export(name="runtimeDatabaseEncryptionKeyName", type=String.class, parameters={})
    private Output</* @Nullable */ String> runtimeDatabaseEncryptionKeyName;

    /**
     * @return Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
     * Update is not allowed after the organization is created.
     * If not specified, a Google-Managed encryption key will be used.
     * Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
     * 
     */
    public Output<Optional<String>> runtimeDatabaseEncryptionKeyName() {
        return Codegen.optional(this.runtimeDatabaseEncryptionKeyName);
    }
    /**
     * Runtime type of the Apigee organization based on the Apigee subscription purchased.
     * Default value is `CLOUD`.
     * Possible values are: `CLOUD`, `HYBRID`.
     * 
     */
    @Export(name="runtimeType", type=String.class, parameters={})
    private Output</* @Nullable */ String> runtimeType;

    /**
     * @return Runtime type of the Apigee organization based on the Apigee subscription purchased.
     * Default value is `CLOUD`.
     * Possible values are: `CLOUD`, `HYBRID`.
     * 
     */
    public Output<Optional<String>> runtimeType() {
        return Codegen.optional(this.runtimeType);
    }
    /**
     * Output only. Subscription type of the Apigee organization.
     * Valid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased).
     * 
     */
    @Export(name="subscriptionType", type=String.class, parameters={})
    private Output<String> subscriptionType;

    /**
     * @return Output only. Subscription type of the Apigee organization.
     * Valid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased).
     * 
     */
    public Output<String> subscriptionType() {
        return this.subscriptionType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Organization(String name) {
        this(name, OrganizationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Organization(String name, OrganizationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Organization(String name, OrganizationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/organization:Organization", name, args == null ? OrganizationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Organization(String name, Output<String> id, @Nullable OrganizationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/organization:Organization", name, state, makeResourceOptions(options, id));
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
    public static Organization get(String name, Output<String> id, @Nullable OrganizationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Organization(name, id, state, options);
    }
}

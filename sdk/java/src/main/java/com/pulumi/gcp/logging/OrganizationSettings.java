// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.logging;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.logging.OrganizationSettingsArgs;
import com.pulumi.gcp.logging.inputs.OrganizationSettingsState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Default resource settings control whether CMEK is required for new log buckets. These settings also determine the storage location for the _Default and _Required log buckets, and whether the _Default sink is enabled or disabled.
 * 
 * To get more information about OrganizationSettings, see:
 * 
 * * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/TopLevel/getSettings)
 * * How-to Guides
 *     * [Configure default settings for organizations and folders](https://cloud.google.com/logging/docs/default-settings)
 * 
 * ## Example Usage
 * 
 * ### Logging Organization Settings All
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.logging.LoggingFunctions;
 * import com.pulumi.gcp.logging.inputs.GetOrganizationSettingsArgs;
 * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
 * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
 * import com.pulumi.gcp.logging.OrganizationSettings;
 * import com.pulumi.gcp.logging.OrganizationSettingsArgs;
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
 *         final var settings = LoggingFunctions.getOrganizationSettings(GetOrganizationSettingsArgs.builder()
 *             .organization("123456789")
 *             .build());
 * 
 *         var iam = new CryptoKeyIAMMember("iam", CryptoKeyIAMMemberArgs.builder()
 *             .cryptoKeyId("kms-key")
 *             .role("roles/cloudkms.cryptoKeyEncrypterDecrypter")
 *             .member(String.format("serviceAccount:%s", settings.applyValue(getOrganizationSettingsResult -> getOrganizationSettingsResult.kmsServiceAccountId())))
 *             .build());
 * 
 *         var example = new OrganizationSettings("example", OrganizationSettingsArgs.builder()
 *             .disableDefaultSink(true)
 *             .kmsKeyName("kms-key")
 *             .organization("123456789")
 *             .storageLocation("us-central1")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(iam)
 *                 .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * OrganizationSettings can be imported using any of these accepted formats:
 * 
 * * `organizations/{{organization}}/settings`
 * 
 * * `{{organization}}`
 * 
 * When using the `pulumi import` command, OrganizationSettings can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:logging/organizationSettings:OrganizationSettings default organizations/{{organization}}/settings
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:logging/organizationSettings:OrganizationSettings default {{organization}}
 * ```
 * 
 */
@ResourceType(type="gcp:logging/organizationSettings:OrganizationSettings")
public class OrganizationSettings extends com.pulumi.resources.CustomResource {
    /**
     * If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The _Default sink can be re-enabled manually if needed.
     * 
     */
    @Export(name="disableDefaultSink", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disableDefaultSink;

    /**
     * @return If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The _Default sink can be re-enabled manually if needed.
     * 
     */
    public Output<Boolean> disableDefaultSink() {
        return this.disableDefaultSink;
    }
    /**
     * The resource name for the configured Cloud KMS key.
     * 
     */
    @Export(name="kmsKeyName", refs={String.class}, tree="[0]")
    private Output<String> kmsKeyName;

    /**
     * @return The resource name for the configured Cloud KMS key.
     * 
     */
    public Output<String> kmsKeyName() {
        return this.kmsKeyName;
    }
    /**
     * The service account that will be used by the Log Router to access your Cloud KMS key.
     * 
     */
    @Export(name="kmsServiceAccountId", refs={String.class}, tree="[0]")
    private Output<String> kmsServiceAccountId;

    /**
     * @return The service account that will be used by the Log Router to access your Cloud KMS key.
     * 
     */
    public Output<String> kmsServiceAccountId() {
        return this.kmsServiceAccountId;
    }
    /**
     * The service account for the given container. Sinks use this service account as their writerIdentity if no custom service account is provided.
     * 
     */
    @Export(name="loggingServiceAccountId", refs={String.class}, tree="[0]")
    private Output<String> loggingServiceAccountId;

    /**
     * @return The service account for the given container. Sinks use this service account as their writerIdentity if no custom service account is provided.
     * 
     */
    public Output<String> loggingServiceAccountId() {
        return this.loggingServiceAccountId;
    }
    /**
     * The resource name of the settings.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the settings.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization for which to retrieve or configure settings.
     * 
     * ***
     * 
     */
    @Export(name="organization", refs={String.class}, tree="[0]")
    private Output<String> organization;

    /**
     * @return The organization for which to retrieve or configure settings.
     * 
     * ***
     * 
     */
    public Output<String> organization() {
        return this.organization;
    }
    /**
     * The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly provided.
     * 
     */
    @Export(name="storageLocation", refs={String.class}, tree="[0]")
    private Output<String> storageLocation;

    /**
     * @return The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly provided.
     * 
     */
    public Output<String> storageLocation() {
        return this.storageLocation;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OrganizationSettings(java.lang.String name) {
        this(name, OrganizationSettingsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OrganizationSettings(java.lang.String name, OrganizationSettingsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OrganizationSettings(java.lang.String name, OrganizationSettingsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:logging/organizationSettings:OrganizationSettings", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private OrganizationSettings(java.lang.String name, Output<java.lang.String> id, @Nullable OrganizationSettingsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:logging/organizationSettings:OrganizationSettings", name, state, makeResourceOptions(options, id), false);
    }

    private static OrganizationSettingsArgs makeArgs(OrganizationSettingsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? OrganizationSettingsArgs.Empty : args;
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
    public static OrganizationSettings get(java.lang.String name, Output<java.lang.String> id, @Nullable OrganizationSettingsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OrganizationSettings(name, id, state, options);
    }
}

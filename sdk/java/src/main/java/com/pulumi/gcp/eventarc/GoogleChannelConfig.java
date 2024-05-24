// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.eventarc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.eventarc.GoogleChannelConfigArgs;
import com.pulumi.gcp.eventarc.inputs.GoogleChannelConfigState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The Eventarc GoogleChannelConfig resource
 * 
 * ## Example Usage
 * 
 * ### Basic
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.kms.KmsFunctions;
 * import com.pulumi.gcp.kms.inputs.GetKMSKeyRingArgs;
 * import com.pulumi.gcp.kms.inputs.GetKMSCryptoKeyArgs;
 * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
 * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
 * import com.pulumi.gcp.eventarc.GoogleChannelConfig;
 * import com.pulumi.gcp.eventarc.GoogleChannelConfigArgs;
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
 *         final var testProject = OrganizationsFunctions.getProject(GetProjectArgs.builder()
 *             .projectId("my-project-name")
 *             .build());
 * 
 *         final var testKeyRing = KmsFunctions.getKMSKeyRing(GetKMSKeyRingArgs.builder()
 *             .name("keyring")
 *             .location("us-west1")
 *             .build());
 * 
 *         final var key = KmsFunctions.getKMSCryptoKey(GetKMSCryptoKeyArgs.builder()
 *             .name("key")
 *             .keyRing(testKeyRing.applyValue(getKMSKeyRingResult -> getKMSKeyRingResult.id()))
 *             .build());
 * 
 *         var key1Member = new CryptoKeyIAMMember("key1Member", CryptoKeyIAMMemberArgs.builder()
 *             .cryptoKeyId(key1.id())
 *             .role("roles/cloudkms.cryptoKeyEncrypterDecrypter")
 *             .member(String.format("serviceAccount:service-%s{@literal @}gcp-sa-eventarc.iam.gserviceaccount.com", testProject.applyValue(getProjectResult -> getProjectResult.number())))
 *             .build());
 * 
 *         var primary = new GoogleChannelConfig("primary", GoogleChannelConfigArgs.builder()
 *             .location("us-west1")
 *             .name("channel")
 *             .project(testProject.applyValue(getProjectResult -> getProjectResult.projectId()))
 *             .cryptoKeyName(key1.id())
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
 * GoogleChannelConfig can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/googleChannelConfig`
 * 
 * * `{{project}}/{{location}}`
 * 
 * * `{{location}}`
 * 
 * When using the `pulumi import` command, GoogleChannelConfig can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:eventarc/googleChannelConfig:GoogleChannelConfig default projects/{{project}}/locations/{{location}}/googleChannelConfig
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:eventarc/googleChannelConfig:GoogleChannelConfig default {{project}}/{{location}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:eventarc/googleChannelConfig:GoogleChannelConfig default {{location}}
 * ```
 * 
 */
@ResourceType(type="gcp:eventarc/googleChannelConfig:GoogleChannelConfig")
public class GoogleChannelConfig extends com.pulumi.resources.CustomResource {
    /**
     * Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*{@literal /}locations/*{@literal /}keyRings/*{@literal /}cryptoKeys/*`.
     * 
     */
    @Export(name="cryptoKeyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cryptoKeyName;

    /**
     * @return Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*{@literal /}locations/*{@literal /}keyRings/*{@literal /}cryptoKeys/*`.
     * 
     */
    public Output<Optional<String>> cryptoKeyName() {
        return Codegen.optional(this.cryptoKeyName);
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
     * Required. The resource name of the config. Must be in the format of, `projects/{project}/locations/{location}/googleChannelConfig`.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Required. The resource name of the config. Must be in the format of, `projects/{project}/locations/{location}/googleChannelConfig`.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * Output only. The last-modified time.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. The last-modified time.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GoogleChannelConfig(String name) {
        this(name, GoogleChannelConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GoogleChannelConfig(String name, GoogleChannelConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GoogleChannelConfig(String name, GoogleChannelConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:eventarc/googleChannelConfig:GoogleChannelConfig", name, args == null ? GoogleChannelConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GoogleChannelConfig(String name, Output<String> id, @Nullable GoogleChannelConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:eventarc/googleChannelConfig:GoogleChannelConfig", name, state, makeResourceOptions(options, id));
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
    public static GoogleChannelConfig get(String name, Output<String> id, @Nullable GoogleChannelConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GoogleChannelConfig(name, id, state, options);
    }
}

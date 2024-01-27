// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.eventarc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.eventarc.ChannelArgs;
import com.pulumi.gcp.eventarc.inputs.ChannelState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The Eventarc Channel resource
 * 
 * ## Example Usage
 * ### Basic
 * ```java
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
 * import com.pulumi.gcp.eventarc.Channel;
 * import com.pulumi.gcp.eventarc.ChannelArgs;
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
 *         final var testProject = OrganizationsFunctions.getProject(GetProjectArgs.builder()
 *             .projectId(&#34;my-project-name&#34;)
 *             .build());
 * 
 *         final var testKeyRing = KmsFunctions.getKMSKeyRing(GetKMSKeyRingArgs.builder()
 *             .name(&#34;keyring&#34;)
 *             .location(&#34;us-west1&#34;)
 *             .build());
 * 
 *         final var key = KmsFunctions.getKMSCryptoKey(GetKMSCryptoKeyArgs.builder()
 *             .name(&#34;key&#34;)
 *             .keyRing(testKeyRing.applyValue(getKMSKeyRingResult -&gt; getKMSKeyRingResult.id()))
 *             .build());
 * 
 *         var key1Member = new CryptoKeyIAMMember(&#34;key1Member&#34;, CryptoKeyIAMMemberArgs.builder()        
 *             .cryptoKeyId(data.google_kms_crypto_key().key1().id())
 *             .role(&#34;roles/cloudkms.cryptoKeyEncrypterDecrypter&#34;)
 *             .member(String.format(&#34;serviceAccount:service-%s@gcp-sa-eventarc.iam.gserviceaccount.com&#34;, testProject.applyValue(getProjectResult -&gt; getProjectResult.number())))
 *             .build());
 * 
 *         var primary = new Channel(&#34;primary&#34;, ChannelArgs.builder()        
 *             .location(&#34;us-west1&#34;)
 *             .project(testProject.applyValue(getProjectResult -&gt; getProjectResult.projectId()))
 *             .cryptoKeyName(data.google_kms_crypto_key().key1().id())
 *             .thirdPartyProvider(String.format(&#34;projects/%s/locations/us-west1/providers/datadog&#34;, testProject.applyValue(getProjectResult -&gt; getProjectResult.projectId())))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(key1Member)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Channel can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/channels/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Channel can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:eventarc/channel:Channel default projects/{{project}}/locations/{{location}}/channels/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:eventarc/channel:Channel default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:eventarc/channel:Channel default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:eventarc/channel:Channel")
public class Channel extends com.pulumi.resources.CustomResource {
    /**
     * Output only. The activation token for the channel. The token must be used by the provider to register the channel for publishing.
     * 
     */
    @Export(name="activationToken", refs={String.class}, tree="[0]")
    private Output<String> activationToken;

    /**
     * @return Output only. The activation token for the channel. The token must be used by the provider to register the channel for publishing.
     * 
     */
    public Output<String> activationToken() {
        return this.activationToken;
    }
    /**
     * Output only. The creation time.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. The creation time.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
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
     * Required. The resource name of the channel. Must be unique within the location on the project.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Required. The resource name of the channel. Must be unique within the location on the project.
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
     * Output only. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{project}/topics/{topic_id}`.
     * 
     */
    @Export(name="pubsubTopic", refs={String.class}, tree="[0]")
    private Output<String> pubsubTopic;

    /**
     * @return Output only. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{project}/topics/{topic_id}`.
     * 
     */
    public Output<String> pubsubTopic() {
        return this.pubsubTopic;
    }
    /**
     * Output only. The state of a Channel. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE, INACTIVE
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Output only. The state of a Channel. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE, INACTIVE
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
     * 
     */
    @Export(name="thirdPartyProvider", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> thirdPartyProvider;

    /**
     * @return The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
     * 
     */
    public Output<Optional<String>> thirdPartyProvider() {
        return Codegen.optional(this.thirdPartyProvider);
    }
    /**
     * Output only. Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return Output only. Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
     * 
     */
    public Output<String> uid() {
        return this.uid;
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
    public Channel(String name) {
        this(name, ChannelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Channel(String name, ChannelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Channel(String name, ChannelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:eventarc/channel:Channel", name, args == null ? ChannelArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Channel(String name, Output<String> id, @Nullable ChannelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:eventarc/channel:Channel", name, state, makeResourceOptions(options, id));
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
    public static Channel get(String name, Output<String> id, @Nullable ChannelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Channel(name, id, state, options);
    }
}

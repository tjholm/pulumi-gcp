// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.pubsub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.pubsub.TopicArgs;
import com.pulumi.gcp.pubsub.inputs.TopicState;
import com.pulumi.gcp.pubsub.outputs.TopicMessageStoragePolicy;
import com.pulumi.gcp.pubsub.outputs.TopicSchemaSettings;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A named resource to which messages are sent by publishers.
 * 
 * To get more information about Topic, see:
 * 
 * * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics)
 * * How-to Guides
 *     * [Managing Topics](https://cloud.google.com/pubsub/docs/admin#managing_topics)
 * 
 * &gt; **Note:** You can retrieve the email of the Google Managed Pub/Sub Service Account used for forwarding
 * by using the `gcp.projects.ServiceIdentity` resource.
 * 
 * ## Example Usage
 * ### Pubsub Topic Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.pubsub.Topic;
 * import com.pulumi.gcp.pubsub.TopicArgs;
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
 *         var example = new Topic(&#34;example&#34;, TopicArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .messageRetentionDuration(&#34;86600s&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Pubsub Topic Cmek
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.kms.CryptoKey;
 * import com.pulumi.gcp.kms.CryptoKeyArgs;
 * import com.pulumi.gcp.pubsub.Topic;
 * import com.pulumi.gcp.pubsub.TopicArgs;
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
 *         var keyRing = new KeyRing(&#34;keyRing&#34;, KeyRingArgs.builder()        
 *             .location(&#34;global&#34;)
 *             .build());
 * 
 *         var cryptoKey = new CryptoKey(&#34;cryptoKey&#34;, CryptoKeyArgs.builder()        
 *             .keyRing(keyRing.id())
 *             .build());
 * 
 *         var example = new Topic(&#34;example&#34;, TopicArgs.builder()        
 *             .kmsKeyName(cryptoKey.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Pubsub Topic Geo Restricted
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.pubsub.Topic;
 * import com.pulumi.gcp.pubsub.TopicArgs;
 * import com.pulumi.gcp.pubsub.inputs.TopicMessageStoragePolicyArgs;
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
 *         var example = new Topic(&#34;example&#34;, TopicArgs.builder()        
 *             .messageStoragePolicy(TopicMessageStoragePolicyArgs.builder()
 *                 .allowedPersistenceRegions(&#34;europe-west3&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Pubsub Topic Schema Settings
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.pubsub.Schema;
 * import com.pulumi.gcp.pubsub.SchemaArgs;
 * import com.pulumi.gcp.pubsub.Topic;
 * import com.pulumi.gcp.pubsub.TopicArgs;
 * import com.pulumi.gcp.pubsub.inputs.TopicSchemaSettingsArgs;
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
 *         var exampleSchema = new Schema(&#34;exampleSchema&#34;, SchemaArgs.builder()        
 *             .type(&#34;AVRO&#34;)
 *             .definition(&#34;&#34;&#34;
 * {
 *   &#34;type&#34; : &#34;record&#34;,
 *   &#34;name&#34; : &#34;Avro&#34;,
 *   &#34;fields&#34; : [
 *     {
 *       &#34;name&#34; : &#34;StringField&#34;,
 *       &#34;type&#34; : &#34;string&#34;
 *     },
 *     {
 *       &#34;name&#34; : &#34;IntField&#34;,
 *       &#34;type&#34; : &#34;int&#34;
 *     }
 *   ]
 * }
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *         var exampleTopic = new Topic(&#34;exampleTopic&#34;, TopicArgs.builder()        
 *             .schemaSettings(TopicSchemaSettingsArgs.builder()
 *                 .schema(&#34;projects/my-project-name/schemas/example&#34;)
 *                 .encoding(&#34;JSON&#34;)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(exampleSchema)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Topic can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:pubsub/topic:Topic default projects/{{project}}/topics/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:pubsub/topic:Topic default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:pubsub/topic:Topic default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:pubsub/topic:Topic")
public class Topic extends com.pulumi.resources.CustomResource {
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * The resource name of the Cloud KMS CryptoKey to be used to protect access
     * to messages published on this topic. Your project&#39;s PubSub service account
     * (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
     * `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
     * The expected format is `projects/*{@literal /}locations/*{@literal /}keyRings/*{@literal /}cryptoKeys/*`
     * 
     */
    @Export(name="kmsKeyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyName;

    /**
     * @return The resource name of the Cloud KMS CryptoKey to be used to protect access
     * to messages published on this topic. Your project&#39;s PubSub service account
     * (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
     * `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
     * The expected format is `projects/*{@literal /}locations/*{@literal /}keyRings/*{@literal /}cryptoKeys/*`
     * 
     */
    public Output<Optional<String>> kmsKeyName() {
        return Codegen.optional(this.kmsKeyName);
    }
    /**
     * A set of key/value label pairs to assign to this Topic.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs to assign to this Topic.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Indicates the minimum duration to retain a message after it is published
     * to the topic. If this field is set, messages published to the topic in
     * the last messageRetentionDuration are always available to subscribers.
     * For instance, it allows any attached subscription to seek to a timestamp
     * that is up to messageRetentionDuration in the past. If this field is not
     * set, message retention is controlled by settings on individual subscriptions.
     * Cannot be more than 31 days or less than 10 minutes.
     * 
     */
    @Export(name="messageRetentionDuration", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> messageRetentionDuration;

    /**
     * @return Indicates the minimum duration to retain a message after it is published
     * to the topic. If this field is set, messages published to the topic in
     * the last messageRetentionDuration are always available to subscribers.
     * For instance, it allows any attached subscription to seek to a timestamp
     * that is up to messageRetentionDuration in the past. If this field is not
     * set, message retention is controlled by settings on individual subscriptions.
     * Cannot be more than 31 days or less than 10 minutes.
     * 
     */
    public Output<Optional<String>> messageRetentionDuration() {
        return Codegen.optional(this.messageRetentionDuration);
    }
    /**
     * Policy constraining the set of Google Cloud Platform regions where
     * messages published to the topic may be stored. If not present, then no
     * constraints are in effect.
     * Structure is documented below.
     * 
     */
    @Export(name="messageStoragePolicy", refs={TopicMessageStoragePolicy.class}, tree="[0]")
    private Output<TopicMessageStoragePolicy> messageStoragePolicy;

    /**
     * @return Policy constraining the set of Google Cloud Platform regions where
     * messages published to the topic may be stored. If not present, then no
     * constraints are in effect.
     * Structure is documented below.
     * 
     */
    public Output<TopicMessageStoragePolicy> messageStoragePolicy() {
        return this.messageStoragePolicy;
    }
    /**
     * Name of the topic.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the topic.
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
     * Settings for validating messages published against a schema.
     * Structure is documented below.
     * 
     */
    @Export(name="schemaSettings", refs={TopicSchemaSettings.class}, tree="[0]")
    private Output<TopicSchemaSettings> schemaSettings;

    /**
     * @return Settings for validating messages published against a schema.
     * Structure is documented below.
     * 
     */
    public Output<TopicSchemaSettings> schemaSettings() {
        return this.schemaSettings;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Topic(String name) {
        this(name, TopicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Topic(String name, @Nullable TopicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Topic(String name, @Nullable TopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:pubsub/topic:Topic", name, args == null ? TopicArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Topic(String name, Output<String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:pubsub/topic:Topic", name, state, makeResourceOptions(options, id));
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
    public static Topic get(String name, Output<String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Topic(name, id, state, options);
    }
}

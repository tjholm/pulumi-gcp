// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.kms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.kms.RegistryArgs;
import com.pulumi.gcp.kms.inputs.RegistryState;
import com.pulumi.gcp.kms.outputs.RegistryCredential;
import com.pulumi.gcp.kms.outputs.RegistryEventNotificationConfigItem;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Google Cloud IoT Core device registry.
 * 
 * To get more information about DeviceRegistry, see:
 * 
 * * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/iot/docs/)
 * 
 * ## Example Usage
 * ### Cloudiot Device Registry Basic
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var test_registry = new Registry(&#34;test-registry&#34;);
 * 
 *     }
 * }
 * ```
 * ### Cloudiot Device Registry Single Event Notification Configs
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var default_telemetry = new Topic(&#34;default-telemetry&#34;);
 * 
 *         var test_registry = new Registry(&#34;test-registry&#34;, RegistryArgs.builder()        
 *             .eventNotificationConfigs(RegistryEventNotificationConfigItemArgs.builder()
 *                 .pubsubTopicName(default_telemetry.id())
 *                 .subfolderMatches(&#34;&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Cloudiot Device Registry Full
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var default_devicestatus = new Topic(&#34;default-devicestatus&#34;);
 * 
 *         var default_telemetry = new Topic(&#34;default-telemetry&#34;);
 * 
 *         var additional_telemetry = new Topic(&#34;additional-telemetry&#34;);
 * 
 *         var test_registry = new Registry(&#34;test-registry&#34;, RegistryArgs.builder()        
 *             .eventNotificationConfigs(            
 *                 RegistryEventNotificationConfigItemArgs.builder()
 *                     .pubsubTopicName(additional_telemetry.id())
 *                     .subfolderMatches(&#34;test/path&#34;)
 *                     .build(),
 *                 RegistryEventNotificationConfigItemArgs.builder()
 *                     .pubsubTopicName(default_telemetry.id())
 *                     .subfolderMatches(&#34;&#34;)
 *                     .build())
 *             .stateNotificationConfig(Map.of(&#34;pubsub_topic_name&#34;, default_devicestatus.id()))
 *             .mqttConfig(Map.of(&#34;mqtt_enabled_state&#34;, &#34;MQTT_ENABLED&#34;))
 *             .httpConfig(Map.of(&#34;http_enabled_state&#34;, &#34;HTTP_ENABLED&#34;))
 *             .logLevel(&#34;INFO&#34;)
 *             .credentials(RegistryCredentialArgs.builder()
 *                 .publicKeyCertificate(Map.ofEntries(
 *                     Map.entry(&#34;format&#34;, &#34;X509_CERTIFICATE_PEM&#34;),
 *                     Map.entry(&#34;certificate&#34;, Files.readString(&#34;test-fixtures/rsa_cert.pem&#34;))
 *                 ))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * DeviceRegistry can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:kms/registry:Registry default {{project}}/locations/{{region}}/registries/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:kms/registry:Registry default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:kms/registry:Registry default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:kms/registry:Registry default {{name}}
 * ```
 * 
 * @deprecated
 * gcp.kms.Registry has been deprecated in favor of gcp.iot.Registry
 * 
 */
@Deprecated /* gcp.kms.Registry has been deprecated in favor of gcp.iot.Registry */
@ResourceType(type="gcp:kms/registry:Registry")
public class Registry extends com.pulumi.resources.CustomResource {
    /**
     * List of public key certificates to authenticate devices.
     * The structure is documented below.
     * 
     */
    @Export(name="credentials", type=List.class, parameters={RegistryCredential.class})
    private Output</* @Nullable */ List<RegistryCredential>> credentials;

    /**
     * @return List of public key certificates to authenticate devices.
     * The structure is documented below.
     * 
     */
    public Output<Optional<List<RegistryCredential>>> credentials() {
        return Codegen.optional(this.credentials);
    }
    /**
     * List of configurations for event notifications, such as PubSub topics
     * to publish device events to.
     * Structure is documented below.
     * 
     */
    @Export(name="eventNotificationConfigs", type=List.class, parameters={RegistryEventNotificationConfigItem.class})
    private Output<List<RegistryEventNotificationConfigItem>> eventNotificationConfigs;

    /**
     * @return List of configurations for event notifications, such as PubSub topics
     * to publish device events to.
     * Structure is documented below.
     * 
     */
    public Output<List<RegistryEventNotificationConfigItem>> eventNotificationConfigs() {
        return this.eventNotificationConfigs;
    }
    /**
     * Activate or deactivate HTTP.
     * The structure is documented below.
     * 
     */
    @Export(name="httpConfig", type=Map.class, parameters={String.class, Object.class})
    private Output<Map<String,Object>> httpConfig;

    /**
     * @return Activate or deactivate HTTP.
     * The structure is documented below.
     * 
     */
    public Output<Map<String,Object>> httpConfig() {
        return this.httpConfig;
    }
    /**
     * The default logging verbosity for activity from devices in this
     * registry. Specifies which events should be written to logs. For
     * example, if the LogLevel is ERROR, only events that terminate in
     * errors will be logged. LogLevel is inclusive; enabling INFO logging
     * will also enable ERROR logging.
     * Default value is `NONE`.
     * Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
     * 
     */
    @Export(name="logLevel", type=String.class, parameters={})
    private Output</* @Nullable */ String> logLevel;

    /**
     * @return The default logging verbosity for activity from devices in this
     * registry. Specifies which events should be written to logs. For
     * example, if the LogLevel is ERROR, only events that terminate in
     * errors will be logged. LogLevel is inclusive; enabling INFO logging
     * will also enable ERROR logging.
     * Default value is `NONE`.
     * Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
     * 
     */
    public Output<Optional<String>> logLevel() {
        return Codegen.optional(this.logLevel);
    }
    /**
     * Activate or deactivate MQTT.
     * The structure is documented below.
     * 
     */
    @Export(name="mqttConfig", type=Map.class, parameters={String.class, Object.class})
    private Output<Map<String,Object>> mqttConfig;

    /**
     * @return Activate or deactivate MQTT.
     * The structure is documented below.
     * 
     */
    public Output<Map<String,Object>> mqttConfig() {
        return this.mqttConfig;
    }
    /**
     * A unique name for the resource, required by device registry.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return A unique name for the resource, required by device registry.
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
    @Export(name="project", type=String.class, parameters={})
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
     * The region in which the created registry should reside.
     * If it is not provided, the provider region is used.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which the created registry should reside.
     * If it is not provided, the provider region is used.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * A PubSub topic to publish device state updates.
     * The structure is documented below.
     * 
     */
    @Export(name="stateNotificationConfig", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> stateNotificationConfig;

    /**
     * @return A PubSub topic to publish device state updates.
     * The structure is documented below.
     * 
     */
    public Output<Optional<Map<String,Object>>> stateNotificationConfig() {
        return Codegen.optional(this.stateNotificationConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Registry(String name) {
        this(name, RegistryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Registry(String name, @Nullable RegistryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Registry(String name, @Nullable RegistryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:kms/registry:Registry", name, args == null ? RegistryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Registry(String name, Output<String> id, @Nullable RegistryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:kms/registry:Registry", name, state, makeResourceOptions(options, id));
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
    public static Registry get(String name, Output<String> id, @Nullable RegistryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Registry(name, id, state, options);
    }
}

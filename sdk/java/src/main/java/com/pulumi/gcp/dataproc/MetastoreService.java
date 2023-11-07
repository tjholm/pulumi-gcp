// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataproc.MetastoreServiceArgs;
import com.pulumi.gcp.dataproc.inputs.MetastoreServiceState;
import com.pulumi.gcp.dataproc.outputs.MetastoreServiceEncryptionConfig;
import com.pulumi.gcp.dataproc.outputs.MetastoreServiceHiveMetastoreConfig;
import com.pulumi.gcp.dataproc.outputs.MetastoreServiceMaintenanceWindow;
import com.pulumi.gcp.dataproc.outputs.MetastoreServiceMetadataIntegration;
import com.pulumi.gcp.dataproc.outputs.MetastoreServiceNetworkConfig;
import com.pulumi.gcp.dataproc.outputs.MetastoreServiceScalingConfig;
import com.pulumi.gcp.dataproc.outputs.MetastoreServiceTelemetryConfig;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A managed metastore service that serves metadata queries.
 * 
 * To get more information about Service, see:
 * 
 * * [API documentation](https://cloud.google.com/dataproc-metastore/docs/reference/rest/v1/projects.locations.services)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dataproc-metastore/docs/overview)
 * 
 * ## Example Usage
 * ### Dataproc Metastore Service Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataproc.MetastoreService;
 * import com.pulumi.gcp.dataproc.MetastoreServiceArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceHiveMetastoreConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceMaintenanceWindowArgs;
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
 *         var default_ = new MetastoreService(&#34;default&#34;, MetastoreServiceArgs.builder()        
 *             .hiveMetastoreConfig(MetastoreServiceHiveMetastoreConfigArgs.builder()
 *                 .version(&#34;2.3.6&#34;)
 *                 .build())
 *             .labels(Map.of(&#34;env&#34;, &#34;test&#34;))
 *             .location(&#34;us-central1&#34;)
 *             .maintenanceWindow(MetastoreServiceMaintenanceWindowArgs.builder()
 *                 .dayOfWeek(&#34;SUNDAY&#34;)
 *                 .hourOfDay(2)
 *                 .build())
 *             .port(9080)
 *             .serviceId(&#34;metastore-srv&#34;)
 *             .tier(&#34;DEVELOPER&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Dataproc Metastore Service Cmek Example
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
 * import com.pulumi.gcp.dataproc.MetastoreService;
 * import com.pulumi.gcp.dataproc.MetastoreServiceArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceEncryptionConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceHiveMetastoreConfigArgs;
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
 *         var keyRing = new KeyRing(&#34;keyRing&#34;, KeyRingArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var cryptoKey = new CryptoKey(&#34;cryptoKey&#34;, CryptoKeyArgs.builder()        
 *             .keyRing(keyRing.id())
 *             .purpose(&#34;ENCRYPT_DECRYPT&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var default_ = new MetastoreService(&#34;default&#34;, MetastoreServiceArgs.builder()        
 *             .serviceId(&#34;example-service&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .encryptionConfig(MetastoreServiceEncryptionConfigArgs.builder()
 *                 .kmsKey(cryptoKey.id())
 *                 .build())
 *             .hiveMetastoreConfig(MetastoreServiceHiveMetastoreConfigArgs.builder()
 *                 .version(&#34;3.1.2&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Dataproc Metastore Service Private Service Connect
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
 * import com.pulumi.gcp.dataproc.MetastoreService;
 * import com.pulumi.gcp.dataproc.MetastoreServiceArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceHiveMetastoreConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceNetworkConfigArgs;
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
 *         var net = new Network(&#34;net&#34;, NetworkArgs.builder()        
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var subnet = new Subnetwork(&#34;subnet&#34;, SubnetworkArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .network(net.id())
 *             .ipCidrRange(&#34;10.0.0.0/22&#34;)
 *             .privateIpGoogleAccess(true)
 *             .build());
 * 
 *         var default_ = new MetastoreService(&#34;default&#34;, MetastoreServiceArgs.builder()        
 *             .serviceId(&#34;metastore-srv&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .hiveMetastoreConfig(MetastoreServiceHiveMetastoreConfigArgs.builder()
 *                 .version(&#34;3.1.2&#34;)
 *                 .build())
 *             .networkConfig(MetastoreServiceNetworkConfigArgs.builder()
 *                 .consumers(MetastoreServiceNetworkConfigConsumerArgs.builder()
 *                     .subnetwork(subnet.id())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Dataproc Metastore Service Dpms2
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataproc.MetastoreService;
 * import com.pulumi.gcp.dataproc.MetastoreServiceArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceHiveMetastoreConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceScalingConfigArgs;
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
 *         var dpms2 = new MetastoreService(&#34;dpms2&#34;, MetastoreServiceArgs.builder()        
 *             .databaseType(&#34;SPANNER&#34;)
 *             .hiveMetastoreConfig(MetastoreServiceHiveMetastoreConfigArgs.builder()
 *                 .version(&#34;3.1.2&#34;)
 *                 .build())
 *             .location(&#34;us-central1&#34;)
 *             .scalingConfig(MetastoreServiceScalingConfigArgs.builder()
 *                 .instanceSize(&#34;EXTRA_SMALL&#34;)
 *                 .build())
 *             .serviceId(&#34;dpms2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Dataproc Metastore Service Dpms2 Scaling Factor
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataproc.MetastoreService;
 * import com.pulumi.gcp.dataproc.MetastoreServiceArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceHiveMetastoreConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceScalingConfigArgs;
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
 *         var dpms2ScalingFactor = new MetastoreService(&#34;dpms2ScalingFactor&#34;, MetastoreServiceArgs.builder()        
 *             .databaseType(&#34;SPANNER&#34;)
 *             .hiveMetastoreConfig(MetastoreServiceHiveMetastoreConfigArgs.builder()
 *                 .version(&#34;3.1.2&#34;)
 *                 .build())
 *             .location(&#34;us-central1&#34;)
 *             .scalingConfig(MetastoreServiceScalingConfigArgs.builder()
 *                 .scalingFactor(&#34;2&#34;)
 *                 .build())
 *             .serviceId(&#34;dpms2sf&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Service can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:dataproc/metastoreService:MetastoreService default projects/{{project}}/locations/{{location}}/services/{{service_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:dataproc/metastoreService:MetastoreService default {{project}}/{{location}}/{{service_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:dataproc/metastoreService:MetastoreService default {{location}}/{{service_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:dataproc/metastoreService:MetastoreService")
public class MetastoreService extends com.pulumi.resources.CustomResource {
    /**
     * A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
     * 
     */
    @Export(name="artifactGcsUri", refs={String.class}, tree="[0]")
    private Output<String> artifactGcsUri;

    /**
     * @return A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
     * 
     */
    public Output<String> artifactGcsUri() {
        return this.artifactGcsUri;
    }
    /**
     * The database type that the Metastore service stores its data.
     * Default value is `MYSQL`.
     * Possible values are: `MYSQL`, `SPANNER`.
     * 
     */
    @Export(name="databaseType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> databaseType;

    /**
     * @return The database type that the Metastore service stores its data.
     * Default value is `MYSQL`.
     * Possible values are: `MYSQL`, `SPANNER`.
     * 
     */
    public Output<Optional<String>> databaseType() {
        return Codegen.optional(this.databaseType);
    }
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
     * Information used to configure the Dataproc Metastore service to encrypt
     * customer data at rest.
     * Structure is documented below.
     * 
     */
    @Export(name="encryptionConfig", refs={MetastoreServiceEncryptionConfig.class}, tree="[0]")
    private Output</* @Nullable */ MetastoreServiceEncryptionConfig> encryptionConfig;

    /**
     * @return Information used to configure the Dataproc Metastore service to encrypt
     * customer data at rest.
     * Structure is documented below.
     * 
     */
    public Output<Optional<MetastoreServiceEncryptionConfig>> encryptionConfig() {
        return Codegen.optional(this.encryptionConfig);
    }
    /**
     * (Output)
     * The URI of the endpoint used to access the metastore service.
     * 
     */
    @Export(name="endpointUri", refs={String.class}, tree="[0]")
    private Output<String> endpointUri;

    /**
     * @return (Output)
     * The URI of the endpoint used to access the metastore service.
     * 
     */
    public Output<String> endpointUri() {
        return this.endpointUri;
    }
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     * Structure is documented below.
     * 
     */
    @Export(name="hiveMetastoreConfig", refs={MetastoreServiceHiveMetastoreConfig.class}, tree="[0]")
    private Output</* @Nullable */ MetastoreServiceHiveMetastoreConfig> hiveMetastoreConfig;

    /**
     * @return Configuration information specific to running Hive metastore software as the metastore service.
     * Structure is documented below.
     * 
     */
    public Output<Optional<MetastoreServiceHiveMetastoreConfig>> hiveMetastoreConfig() {
        return Codegen.optional(this.hiveMetastoreConfig);
    }
    /**
     * User-defined labels for the metastore service.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return User-defined labels for the metastore service.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location where the metastore service should reside.
     * The default value is `global`.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> location;

    /**
     * @return The location where the metastore service should reside.
     * The default value is `global`.
     * 
     */
    public Output<Optional<String>> location() {
        return Codegen.optional(this.location);
    }
    /**
     * The one hour maintenance window of the metastore service.
     * This specifies when the service can be restarted for maintenance purposes in UTC time.
     * Maintenance window is not needed for services with the `SPANNER` database type.
     * Structure is documented below.
     * 
     */
    @Export(name="maintenanceWindow", refs={MetastoreServiceMaintenanceWindow.class}, tree="[0]")
    private Output</* @Nullable */ MetastoreServiceMaintenanceWindow> maintenanceWindow;

    /**
     * @return The one hour maintenance window of the metastore service.
     * This specifies when the service can be restarted for maintenance purposes in UTC time.
     * Maintenance window is not needed for services with the `SPANNER` database type.
     * Structure is documented below.
     * 
     */
    public Output<Optional<MetastoreServiceMaintenanceWindow>> maintenanceWindow() {
        return Codegen.optional(this.maintenanceWindow);
    }
    /**
     * The setting that defines how metastore metadata should be integrated with external services and systems.
     * 
     */
    @Export(name="metadataIntegration", refs={MetastoreServiceMetadataIntegration.class}, tree="[0]")
    private Output</* @Nullable */ MetastoreServiceMetadataIntegration> metadataIntegration;

    /**
     * @return The setting that defines how metastore metadata should be integrated with external services and systems.
     * 
     */
    public Output<Optional<MetastoreServiceMetadataIntegration>> metadataIntegration() {
        return Codegen.optional(this.metadataIntegration);
    }
    /**
     * The relative resource name of the metastore service.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The relative resource name of the metastore service.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
     * &#34;projects/{projectNumber}/global/networks/{network_id}&#34;.
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
     * &#34;projects/{projectNumber}/global/networks/{network_id}&#34;.
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * The configuration specifying the network settings for the Dataproc Metastore service.
     * Structure is documented below.
     * 
     */
    @Export(name="networkConfig", refs={MetastoreServiceNetworkConfig.class}, tree="[0]")
    private Output</* @Nullable */ MetastoreServiceNetworkConfig> networkConfig;

    /**
     * @return The configuration specifying the network settings for the Dataproc Metastore service.
     * Structure is documented below.
     * 
     */
    public Output<Optional<MetastoreServiceNetworkConfig>> networkConfig() {
        return Codegen.optional(this.networkConfig);
    }
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output<Integer> port;

    /**
     * @return The TCP port at which the metastore service is reached. Default: 9083.
     * 
     */
    public Output<Integer> port() {
        return this.port;
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
     * The release channel of the service. If unspecified, defaults to `STABLE`.
     * Default value is `STABLE`.
     * Possible values are: `CANARY`, `STABLE`.
     * 
     */
    @Export(name="releaseChannel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> releaseChannel;

    /**
     * @return The release channel of the service. If unspecified, defaults to `STABLE`.
     * Default value is `STABLE`.
     * Possible values are: `CANARY`, `STABLE`.
     * 
     */
    public Output<Optional<String>> releaseChannel() {
        return Codegen.optional(this.releaseChannel);
    }
    /**
     * Represents the scaling configuration of a metastore service.
     * Structure is documented below.
     * 
     */
    @Export(name="scalingConfig", refs={MetastoreServiceScalingConfig.class}, tree="[0]")
    private Output</* @Nullable */ MetastoreServiceScalingConfig> scalingConfig;

    /**
     * @return Represents the scaling configuration of a metastore service.
     * Structure is documented below.
     * 
     */
    public Output<Optional<MetastoreServiceScalingConfig>> scalingConfig() {
        return Codegen.optional(this.scalingConfig);
    }
    /**
     * The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 63 characters.
     * 
     * ***
     * 
     */
    @Export(name="serviceId", refs={String.class}, tree="[0]")
    private Output<String> serviceId;

    /**
     * @return The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 63 characters.
     * 
     * ***
     * 
     */
    public Output<String> serviceId() {
        return this.serviceId;
    }
    /**
     * The current state of the metastore service.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The current state of the metastore service.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Additional information about the current state of the metastore service, if available.
     * 
     */
    @Export(name="stateMessage", refs={String.class}, tree="[0]")
    private Output<String> stateMessage;

    /**
     * @return Additional information about the current state of the metastore service, if available.
     * 
     */
    public Output<String> stateMessage() {
        return this.stateMessage;
    }
    /**
     * The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
     * Structure is documented below.
     * 
     */
    @Export(name="telemetryConfig", refs={MetastoreServiceTelemetryConfig.class}, tree="[0]")
    private Output<MetastoreServiceTelemetryConfig> telemetryConfig;

    /**
     * @return The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
     * Structure is documented below.
     * 
     */
    public Output<MetastoreServiceTelemetryConfig> telemetryConfig() {
        return this.telemetryConfig;
    }
    /**
     * The tier of the service.
     * Possible values are: `DEVELOPER`, `ENTERPRISE`.
     * 
     */
    @Export(name="tier", refs={String.class}, tree="[0]")
    private Output<String> tier;

    /**
     * @return The tier of the service.
     * Possible values are: `DEVELOPER`, `ENTERPRISE`.
     * 
     */
    public Output<String> tier() {
        return this.tier;
    }
    /**
     * The globally unique resource identifier of the metastore service.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return The globally unique resource identifier of the metastore service.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MetastoreService(String name) {
        this(name, MetastoreServiceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MetastoreService(String name, MetastoreServiceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MetastoreService(String name, MetastoreServiceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataproc/metastoreService:MetastoreService", name, args == null ? MetastoreServiceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MetastoreService(String name, Output<String> id, @Nullable MetastoreServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataproc/metastoreService:MetastoreService", name, state, makeResourceOptions(options, id));
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
    public static MetastoreService get(String name, Output<String> id, @Nullable MetastoreServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MetastoreService(name, id, state, options);
    }
}

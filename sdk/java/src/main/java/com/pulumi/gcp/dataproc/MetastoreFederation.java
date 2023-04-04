// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataproc.MetastoreFederationArgs;
import com.pulumi.gcp.dataproc.inputs.MetastoreFederationState;
import com.pulumi.gcp.dataproc.outputs.MetastoreFederationBackendMetastore;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### Dataproc Metastore Federation Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataproc.MetastoreService;
 * import com.pulumi.gcp.dataproc.MetastoreServiceArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceHiveMetastoreConfigArgs;
 * import com.pulumi.gcp.dataproc.MetastoreFederation;
 * import com.pulumi.gcp.dataproc.MetastoreFederationArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreFederationBackendMetastoreArgs;
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
 *         var defaultMetastoreService = new MetastoreService(&#34;defaultMetastoreService&#34;, MetastoreServiceArgs.builder()        
 *             .serviceId(&#34;&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .tier(&#34;DEVELOPER&#34;)
 *             .hiveMetastoreConfig(MetastoreServiceHiveMetastoreConfigArgs.builder()
 *                 .version(&#34;3.1.2&#34;)
 *                 .endpointProtocol(&#34;GRPC&#34;)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var defaultMetastoreFederation = new MetastoreFederation(&#34;defaultMetastoreFederation&#34;, MetastoreFederationArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .federationId(&#34;&#34;)
 *             .version(&#34;3.1.2&#34;)
 *             .backendMetastores(MetastoreFederationBackendMetastoreArgs.builder()
 *                 .rank(&#34;1&#34;)
 *                 .name(defaultMetastoreService.id())
 *                 .metastoreType(&#34;DATAPROC_METASTORE&#34;)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Dataproc Metastore Federation Bigquery
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataproc.MetastoreService;
 * import com.pulumi.gcp.dataproc.MetastoreServiceArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreServiceHiveMetastoreConfigArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.dataproc.MetastoreFederation;
 * import com.pulumi.gcp.dataproc.MetastoreFederationArgs;
 * import com.pulumi.gcp.dataproc.inputs.MetastoreFederationBackendMetastoreArgs;
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
 *         var defaultMetastoreService = new MetastoreService(&#34;defaultMetastoreService&#34;, MetastoreServiceArgs.builder()        
 *             .serviceId(&#34;&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .tier(&#34;DEVELOPER&#34;)
 *             .hiveMetastoreConfig(MetastoreServiceHiveMetastoreConfigArgs.builder()
 *                 .version(&#34;3.1.2&#34;)
 *                 .endpointProtocol(&#34;GRPC&#34;)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var defaultMetastoreFederation = new MetastoreFederation(&#34;defaultMetastoreFederation&#34;, MetastoreFederationArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .federationId(&#34;&#34;)
 *             .version(&#34;3.1.2&#34;)
 *             .backendMetastores(            
 *                 MetastoreFederationBackendMetastoreArgs.builder()
 *                     .rank(&#34;2&#34;)
 *                     .name(project.applyValue(getProjectResult -&gt; getProjectResult.id()))
 *                     .metastoreType(&#34;BIGQUERY&#34;)
 *                     .build(),
 *                 MetastoreFederationBackendMetastoreArgs.builder()
 *                     .rank(&#34;1&#34;)
 *                     .name(defaultMetastoreService.id())
 *                     .metastoreType(&#34;DATAPROC_METASTORE&#34;)
 *                     .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Federation can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:dataproc/metastoreFederation:MetastoreFederation default projects/{{project}}/locations/{{location}}/federations/{{federation_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:dataproc/metastoreFederation:MetastoreFederation default {{project}}/{{location}}/{{federation_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:dataproc/metastoreFederation:MetastoreFederation default {{location}}/{{federation_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:dataproc/metastoreFederation:MetastoreFederation")
public class MetastoreFederation extends com.pulumi.resources.CustomResource {
    /**
     * A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
     * Structure is documented below.
     * 
     */
    @Export(name="backendMetastores", type=List.class, parameters={MetastoreFederationBackendMetastore.class})
    private Output<List<MetastoreFederationBackendMetastore>> backendMetastores;

    /**
     * @return A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
     * Structure is documented below.
     * 
     */
    public Output<List<MetastoreFederationBackendMetastore>> backendMetastores() {
        return this.backendMetastores;
    }
    /**
     * The URI of the endpoint used to access the metastore federation.
     * 
     */
    @Export(name="endpointUri", type=String.class, parameters={})
    private Output<String> endpointUri;

    /**
     * @return The URI of the endpoint used to access the metastore federation.
     * 
     */
    public Output<String> endpointUri() {
        return this.endpointUri;
    }
    /**
     * The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 63 characters.
     * 
     */
    @Export(name="federationId", type=String.class, parameters={})
    private Output<String> federationId;

    /**
     * @return The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 63 characters.
     * 
     */
    public Output<String> federationId() {
        return this.federationId;
    }
    /**
     * User-defined labels for the metastore federation.
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return User-defined labels for the metastore federation.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location where the metastore federation should reside.
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output</* @Nullable */ String> location;

    /**
     * @return The location where the metastore federation should reside.
     * 
     */
    public Output<Optional<String>> location() {
        return Codegen.optional(this.location);
    }
    /**
     * The relative resource name of the metastore that is being federated. The formats of the relative resource names for the currently supported metastores are listed below: Dataplex: projects/{projectId}/locations/{location}/lakes/{lake_id} BigQuery: projects/{projectId} Dataproc Metastore: projects/{projectId}/locations/{location}/services/{serviceId}
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The relative resource name of the metastore that is being federated. The formats of the relative resource names for the currently supported metastores are listed below: Dataplex: projects/{projectId}/locations/{location}/lakes/{lake_id} BigQuery: projects/{projectId} Dataproc Metastore: projects/{projectId}/locations/{location}/services/{serviceId}
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
     * The current state of the metastore federation.
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return The current state of the metastore federation.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Additional information about the current state of the metastore federation, if available.
     * 
     */
    @Export(name="stateMessage", type=String.class, parameters={})
    private Output<String> stateMessage;

    /**
     * @return Additional information about the current state of the metastore federation, if available.
     * 
     */
    public Output<String> stateMessage() {
        return this.stateMessage;
    }
    /**
     * The globally unique resource identifier of the metastore federation.
     * 
     */
    @Export(name="uid", type=String.class, parameters={})
    private Output<String> uid;

    /**
     * @return The globally unique resource identifier of the metastore federation.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
     * 
     */
    @Export(name="version", type=String.class, parameters={})
    private Output<String> version;

    /**
     * @return The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MetastoreFederation(String name) {
        this(name, MetastoreFederationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MetastoreFederation(String name, MetastoreFederationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MetastoreFederation(String name, MetastoreFederationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataproc/metastoreFederation:MetastoreFederation", name, args == null ? MetastoreFederationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MetastoreFederation(String name, Output<String> id, @Nullable MetastoreFederationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataproc/metastoreFederation:MetastoreFederation", name, state, makeResourceOptions(options, id));
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
    public static MetastoreFederation get(String name, Output<String> id, @Nullable MetastoreFederationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MetastoreFederation(name, id, state, options);
    }
}

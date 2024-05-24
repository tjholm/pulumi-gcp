// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.container.AttachedClusterArgs;
import com.pulumi.gcp.container.inputs.AttachedClusterState;
import com.pulumi.gcp.container.outputs.AttachedClusterAuthorization;
import com.pulumi.gcp.container.outputs.AttachedClusterBinaryAuthorization;
import com.pulumi.gcp.container.outputs.AttachedClusterError;
import com.pulumi.gcp.container.outputs.AttachedClusterFleet;
import com.pulumi.gcp.container.outputs.AttachedClusterLoggingConfig;
import com.pulumi.gcp.container.outputs.AttachedClusterMonitoringConfig;
import com.pulumi.gcp.container.outputs.AttachedClusterOidcConfig;
import com.pulumi.gcp.container.outputs.AttachedClusterProxyConfig;
import com.pulumi.gcp.container.outputs.AttachedClusterWorkloadIdentityConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An Anthos cluster running on customer owned infrastructure.
 * 
 * To get more information about Cluster, see:
 * 
 * * [API documentation](https://cloud.google.com/anthos/clusters/docs/multi-cloud/reference/rest)
 * * How-to Guides
 *     * [API reference](https://cloud.google.com/anthos/clusters/docs/multi-cloud/reference/rest/v1/projects.locations.attachedClusters)
 *     * [Multicloud overview](https://cloud.google.com/anthos/clusters/docs/multi-cloud)
 * 
 * ## Example Usage
 * 
 * ### Container Attached Cluster Basic
 * 
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
 * import com.pulumi.gcp.container.ContainerFunctions;
 * import com.pulumi.gcp.container.inputs.GetAttachedVersionsArgs;
 * import com.pulumi.gcp.container.AttachedCluster;
 * import com.pulumi.gcp.container.AttachedClusterArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterOidcConfigArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterFleetArgs;
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
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         final var versions = ContainerFunctions.getAttachedVersions(GetAttachedVersionsArgs.builder()
 *             .location("us-west1")
 *             .project(project.applyValue(getProjectResult -> getProjectResult.projectId()))
 *             .build());
 * 
 *         var primary = new AttachedCluster("primary", AttachedClusterArgs.builder()
 *             .name("basic")
 *             .location("us-west1")
 *             .project(project.applyValue(getProjectResult -> getProjectResult.projectId()))
 *             .description("Test cluster")
 *             .distribution("aks")
 *             .oidcConfig(AttachedClusterOidcConfigArgs.builder()
 *                 .issuerUrl("https://oidc.issuer.url")
 *                 .build())
 *             .platformVersion(versions.applyValue(getAttachedVersionsResult -> getAttachedVersionsResult.validVersions()[0]))
 *             .fleet(AttachedClusterFleetArgs.builder()
 *                 .project(String.format("projects/%s", project.applyValue(getProjectResult -> getProjectResult.number())))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Container Attached Cluster Full
 * 
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
 * import com.pulumi.gcp.container.ContainerFunctions;
 * import com.pulumi.gcp.container.inputs.GetAttachedVersionsArgs;
 * import com.pulumi.gcp.container.AttachedCluster;
 * import com.pulumi.gcp.container.AttachedClusterArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterAuthorizationArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterOidcConfigArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterFleetArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterLoggingConfigArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterLoggingConfigComponentConfigArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterMonitoringConfigArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterMonitoringConfigManagedPrometheusConfigArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterBinaryAuthorizationArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterProxyConfigArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterProxyConfigKubernetesSecretArgs;
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
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         final var versions = ContainerFunctions.getAttachedVersions(GetAttachedVersionsArgs.builder()
 *             .location("us-west1")
 *             .project(project.applyValue(getProjectResult -> getProjectResult.projectId()))
 *             .build());
 * 
 *         var primary = new AttachedCluster("primary", AttachedClusterArgs.builder()
 *             .name("basic")
 *             .project(project.applyValue(getProjectResult -> getProjectResult.projectId()))
 *             .location("us-west1")
 *             .description("Test cluster")
 *             .distribution("aks")
 *             .annotations(Map.of("label-one", "value-one"))
 *             .authorization(AttachedClusterAuthorizationArgs.builder()
 *                 .adminUsers(                
 *                     "user1{@literal @}example.com",
 *                     "user2{@literal @}example.com")
 *                 .adminGroups(                
 *                     "group1{@literal @}example.com",
 *                     "group2{@literal @}example.com")
 *                 .build())
 *             .oidcConfig(AttachedClusterOidcConfigArgs.builder()
 *                 .issuerUrl("https://oidc.issuer.url")
 *                 .jwks(StdFunctions.base64encode(Base64encodeArgs.builder()
 *                     .input("{\"keys\":[{\"use\":\"sig\",\"kty\":\"RSA\",\"kid\":\"testid\",\"alg\":\"RS256\",\"n\":\"somedata\",\"e\":\"AQAB\"}]}")
 *                     .build()).result())
 *                 .build())
 *             .platformVersion(versions.applyValue(getAttachedVersionsResult -> getAttachedVersionsResult.validVersions()[0]))
 *             .fleet(AttachedClusterFleetArgs.builder()
 *                 .project(String.format("projects/%s", project.applyValue(getProjectResult -> getProjectResult.number())))
 *                 .build())
 *             .loggingConfig(AttachedClusterLoggingConfigArgs.builder()
 *                 .componentConfig(AttachedClusterLoggingConfigComponentConfigArgs.builder()
 *                     .enableComponents(                    
 *                         "SYSTEM_COMPONENTS",
 *                         "WORKLOADS")
 *                     .build())
 *                 .build())
 *             .monitoringConfig(AttachedClusterMonitoringConfigArgs.builder()
 *                 .managedPrometheusConfig(AttachedClusterMonitoringConfigManagedPrometheusConfigArgs.builder()
 *                     .enabled(true)
 *                     .build())
 *                 .build())
 *             .binaryAuthorization(AttachedClusterBinaryAuthorizationArgs.builder()
 *                 .evaluationMode("PROJECT_SINGLETON_POLICY_ENFORCE")
 *                 .build())
 *             .proxyConfig(AttachedClusterProxyConfigArgs.builder()
 *                 .kubernetesSecret(AttachedClusterProxyConfigKubernetesSecretArgs.builder()
 *                     .name("proxy-config")
 *                     .namespace("default")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Container Attached Cluster Ignore Errors
 * 
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
 * import com.pulumi.gcp.container.ContainerFunctions;
 * import com.pulumi.gcp.container.inputs.GetAttachedVersionsArgs;
 * import com.pulumi.gcp.container.AttachedCluster;
 * import com.pulumi.gcp.container.AttachedClusterArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterOidcConfigArgs;
 * import com.pulumi.gcp.container.inputs.AttachedClusterFleetArgs;
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
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         final var versions = ContainerFunctions.getAttachedVersions(GetAttachedVersionsArgs.builder()
 *             .location("us-west1")
 *             .project(project.applyValue(getProjectResult -> getProjectResult.projectId()))
 *             .build());
 * 
 *         var primary = new AttachedCluster("primary", AttachedClusterArgs.builder()
 *             .name("basic")
 *             .location("us-west1")
 *             .project(project.applyValue(getProjectResult -> getProjectResult.projectId()))
 *             .description("Test cluster")
 *             .distribution("aks")
 *             .oidcConfig(AttachedClusterOidcConfigArgs.builder()
 *                 .issuerUrl("https://oidc.issuer.url")
 *                 .build())
 *             .platformVersion(versions.applyValue(getAttachedVersionsResult -> getAttachedVersionsResult.validVersions()[0]))
 *             .fleet(AttachedClusterFleetArgs.builder()
 *                 .project(String.format("projects/%s", project.applyValue(getProjectResult -> getProjectResult.number())))
 *                 .build())
 *             .deletionPolicy("DELETE_IGNORE_ERRORS")
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
 * Cluster can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/attachedClusters/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{name}}`
 * 
 * * `{{location}}/{{name}}`
 * 
 * When using the `pulumi import` command, Cluster can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:container/attachedCluster:AttachedCluster default projects/{{project}}/locations/{{location}}/attachedClusters/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:container/attachedCluster:AttachedCluster default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:container/attachedCluster:AttachedCluster default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:container/attachedCluster:AttachedCluster")
public class AttachedCluster extends com.pulumi.resources.CustomResource {
    /**
     * Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
     * all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required),
     * separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
     * alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
     * non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
     * &#39;effective_annotations&#39; for all of the annotations present on the resource.
     * 
     */
    @Export(name="annotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> annotations;

    /**
     * @return Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
     * all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required),
     * separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
     * alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
     * non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
     * &#39;effective_annotations&#39; for all of the annotations present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> annotations() {
        return Codegen.optional(this.annotations);
    }
    /**
     * Configuration related to the cluster RBAC settings.
     * 
     */
    @Export(name="authorization", refs={AttachedClusterAuthorization.class}, tree="[0]")
    private Output</* @Nullable */ AttachedClusterAuthorization> authorization;

    /**
     * @return Configuration related to the cluster RBAC settings.
     * 
     */
    public Output<Optional<AttachedClusterAuthorization>> authorization() {
        return Codegen.optional(this.authorization);
    }
    /**
     * Binary Authorization configuration.
     * 
     */
    @Export(name="binaryAuthorization", refs={AttachedClusterBinaryAuthorization.class}, tree="[0]")
    private Output<AttachedClusterBinaryAuthorization> binaryAuthorization;

    /**
     * @return Binary Authorization configuration.
     * 
     */
    public Output<AttachedClusterBinaryAuthorization> binaryAuthorization() {
        return this.binaryAuthorization;
    }
    /**
     * Output only. The region where this cluster runs.
     * For EKS clusters, this is an AWS region. For AKS clusters,
     * this is an Azure region.
     * 
     */
    @Export(name="clusterRegion", refs={String.class}, tree="[0]")
    private Output<String> clusterRegion;

    /**
     * @return Output only. The region where this cluster runs.
     * For EKS clusters, this is an AWS region. For AKS clusters,
     * this is an Azure region.
     * 
     */
    public Output<String> clusterRegion() {
        return this.clusterRegion;
    }
    /**
     * Output only. The time at which this cluster was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. The time at which this cluster was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Policy to determine what flags to send on delete.
     * 
     */
    @Export(name="deletionPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deletionPolicy;

    /**
     * @return Policy to determine what flags to send on delete.
     * 
     */
    public Output<Optional<String>> deletionPolicy() {
        return Codegen.optional(this.deletionPolicy);
    }
    /**
     * A human readable description of this attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A human readable description of this attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Kubernetes distribution of the underlying attached cluster. Supported values:
     * &#34;eks&#34;, &#34;aks&#34;.
     * 
     */
    @Export(name="distribution", refs={String.class}, tree="[0]")
    private Output<String> distribution;

    /**
     * @return The Kubernetes distribution of the underlying attached cluster. Supported values:
     * &#34;eks&#34;, &#34;aks&#34;.
     * 
     */
    public Output<String> distribution() {
        return this.distribution;
    }
    @Export(name="effectiveAnnotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveAnnotations;

    public Output<Map<String,String>> effectiveAnnotations() {
        return this.effectiveAnnotations;
    }
    /**
     * A set of errors found in the cluster.
     * Structure is documented below.
     * 
     */
    @Export(name="errors", refs={List.class,AttachedClusterError.class}, tree="[0,1]")
    private Output<List<AttachedClusterError>> errors;

    /**
     * @return A set of errors found in the cluster.
     * Structure is documented below.
     * 
     */
    public Output<List<AttachedClusterError>> errors() {
        return this.errors;
    }
    /**
     * Fleet configuration.
     * Structure is documented below.
     * 
     */
    @Export(name="fleet", refs={AttachedClusterFleet.class}, tree="[0]")
    private Output<AttachedClusterFleet> fleet;

    /**
     * @return Fleet configuration.
     * Structure is documented below.
     * 
     */
    public Output<AttachedClusterFleet> fleet() {
        return this.fleet;
    }
    /**
     * The Kubernetes version of the cluster.
     * 
     */
    @Export(name="kubernetesVersion", refs={String.class}, tree="[0]")
    private Output<String> kubernetesVersion;

    /**
     * @return The Kubernetes version of the cluster.
     * 
     */
    public Output<String> kubernetesVersion() {
        return this.kubernetesVersion;
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
     * Logging configuration.
     * 
     */
    @Export(name="loggingConfig", refs={AttachedClusterLoggingConfig.class}, tree="[0]")
    private Output</* @Nullable */ AttachedClusterLoggingConfig> loggingConfig;

    /**
     * @return Logging configuration.
     * 
     */
    public Output<Optional<AttachedClusterLoggingConfig>> loggingConfig() {
        return Codegen.optional(this.loggingConfig);
    }
    /**
     * Monitoring configuration.
     * 
     */
    @Export(name="monitoringConfig", refs={AttachedClusterMonitoringConfig.class}, tree="[0]")
    private Output<AttachedClusterMonitoringConfig> monitoringConfig;

    /**
     * @return Monitoring configuration.
     * 
     */
    public Output<AttachedClusterMonitoringConfig> monitoringConfig() {
        return this.monitoringConfig;
    }
    /**
     * The name of this resource.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of this resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * OIDC discovery information of the target cluster.
     * Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
     * API server. This fields indicates how GCP services
     * validate KSA tokens in order to allow system workloads (such as GKE Connect
     * and telemetry agents) to authenticate back to GCP.
     * Both clusters with public and private issuer URLs are supported.
     * Clusters with public issuers only need to specify the `issuer_url` field
     * while clusters with private issuers need to provide both
     * `issuer_url` and `jwks`.
     * Structure is documented below.
     * 
     */
    @Export(name="oidcConfig", refs={AttachedClusterOidcConfig.class}, tree="[0]")
    private Output<AttachedClusterOidcConfig> oidcConfig;

    /**
     * @return OIDC discovery information of the target cluster.
     * Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
     * API server. This fields indicates how GCP services
     * validate KSA tokens in order to allow system workloads (such as GKE Connect
     * and telemetry agents) to authenticate back to GCP.
     * Both clusters with public and private issuer URLs are supported.
     * Clusters with public issuers only need to specify the `issuer_url` field
     * while clusters with private issuers need to provide both
     * `issuer_url` and `jwks`.
     * Structure is documented below.
     * 
     */
    public Output<AttachedClusterOidcConfig> oidcConfig() {
        return this.oidcConfig;
    }
    /**
     * The platform version for the cluster (e.g. `1.23.0-gke.1`).
     * 
     */
    @Export(name="platformVersion", refs={String.class}, tree="[0]")
    private Output<String> platformVersion;

    /**
     * @return The platform version for the cluster (e.g. `1.23.0-gke.1`).
     * 
     */
    public Output<String> platformVersion() {
        return this.platformVersion;
    }
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }
    /**
     * Support for proxy configuration.
     * 
     */
    @Export(name="proxyConfig", refs={AttachedClusterProxyConfig.class}, tree="[0]")
    private Output</* @Nullable */ AttachedClusterProxyConfig> proxyConfig;

    /**
     * @return Support for proxy configuration.
     * 
     */
    public Output<Optional<AttachedClusterProxyConfig>> proxyConfig() {
        return Codegen.optional(this.proxyConfig);
    }
    /**
     * If set, there are currently changes in flight to the cluster.
     * 
     */
    @Export(name="reconciling", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reconciling;

    /**
     * @return If set, there are currently changes in flight to the cluster.
     * 
     */
    public Output<Boolean> reconciling() {
        return this.reconciling;
    }
    /**
     * The current state of the cluster. Possible values:
     * STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,
     * DEGRADED
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The current state of the cluster. Possible values:
     * STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,
     * DEGRADED
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * A globally unique identifier for the cluster.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return A globally unique identifier for the cluster.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * The time at which this cluster was last updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The time at which this cluster was last updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Workload Identity settings.
     * Structure is documented below.
     * 
     */
    @Export(name="workloadIdentityConfigs", refs={List.class,AttachedClusterWorkloadIdentityConfig.class}, tree="[0,1]")
    private Output<List<AttachedClusterWorkloadIdentityConfig>> workloadIdentityConfigs;

    /**
     * @return Workload Identity settings.
     * Structure is documented below.
     * 
     */
    public Output<List<AttachedClusterWorkloadIdentityConfig>> workloadIdentityConfigs() {
        return this.workloadIdentityConfigs;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AttachedCluster(String name) {
        this(name, AttachedClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AttachedCluster(String name, AttachedClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AttachedCluster(String name, AttachedClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:container/attachedCluster:AttachedCluster", name, args == null ? AttachedClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AttachedCluster(String name, Output<String> id, @Nullable AttachedClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:container/attachedCluster:AttachedCluster", name, state, makeResourceOptions(options, id));
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
    public static AttachedCluster get(String name, Output<String> id, @Nullable AttachedClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AttachedCluster(name, id, state, options);
    }
}

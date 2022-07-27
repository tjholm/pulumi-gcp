// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.container.AzureClusterArgs;
import com.pulumi.gcp.container.inputs.AzureClusterState;
import com.pulumi.gcp.container.outputs.AzureClusterAuthorization;
import com.pulumi.gcp.container.outputs.AzureClusterControlPlane;
import com.pulumi.gcp.container.outputs.AzureClusterFleet;
import com.pulumi.gcp.container.outputs.AzureClusterLoggingConfig;
import com.pulumi.gcp.container.outputs.AzureClusterNetworking;
import com.pulumi.gcp.container.outputs.AzureClusterWorkloadIdentityConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An Anthos cluster running on Azure.
 * 
 * For more information, see:
 * * [Multicloud overview](https://cloud.google.com/anthos/clusters/docs/multi-cloud)
 * ## Example Usage
 * ### Basic_azure_cluster
 * A basic example of a containerazure azure cluster
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.container.ContainerFunctions;
 * import com.pulumi.gcp.container.inputs.GetAzureVersionsArgs;
 * import com.pulumi.gcp.container.AzureClient;
 * import com.pulumi.gcp.container.AzureClientArgs;
 * import com.pulumi.gcp.container.AzureCluster;
 * import com.pulumi.gcp.container.AzureClusterArgs;
 * import com.pulumi.gcp.container.inputs.AzureClusterAuthorizationArgs;
 * import com.pulumi.gcp.container.inputs.AzureClusterControlPlaneArgs;
 * import com.pulumi.gcp.container.inputs.AzureClusterControlPlaneSshConfigArgs;
 * import com.pulumi.gcp.container.inputs.AzureClusterFleetArgs;
 * import com.pulumi.gcp.container.inputs.AzureClusterNetworkingArgs;
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
 *         final var versions = ContainerFunctions.getAzureVersions(GetAzureVersionsArgs.builder()
 *             .location(&#34;us-west1&#34;)
 *             .project(&#34;my-project-name&#34;)
 *             .build());
 * 
 *         var basic = new AzureClient(&#34;basic&#34;, AzureClientArgs.builder()        
 *             .applicationId(&#34;12345678-1234-1234-1234-123456789111&#34;)
 *             .location(&#34;us-west1&#34;)
 *             .project(&#34;my-project-name&#34;)
 *             .tenantId(&#34;12345678-1234-1234-1234-123456789111&#34;)
 *             .build());
 * 
 *         var primary = new AzureCluster(&#34;primary&#34;, AzureClusterArgs.builder()        
 *             .authorization(AzureClusterAuthorizationArgs.builder()
 *                 .adminUsers(AzureClusterAuthorizationAdminUserArgs.builder()
 *                     .username(&#34;mmv2@google.com&#34;)
 *                     .build())
 *                 .build())
 *             .azureRegion(&#34;westus2&#34;)
 *             .client(basic.name().applyValue(name -&gt; String.format(&#34;projects/my-project-number/locations/us-west1/azureClients/%s&#34;, name)))
 *             .controlPlane(AzureClusterControlPlaneArgs.builder()
 *                 .sshConfig(AzureClusterControlPlaneSshConfigArgs.builder()
 *                     .authorizedKey(&#34;ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8yaayO6lnb2v+SedxUMa2c8vtIEzCzBjM3EJJsv8Vm9zUDWR7dXWKoNGARUb2mNGXASvI6mFIDXTIlkQ0poDEPpMaXR0g2cb5xT8jAAJq7fqXL3+0rcJhY/uigQ+MrT6s+ub0BFVbsmGHNrMQttXX9gtmwkeAEvj3mra9e5pkNf90qlKnZz6U0SVArxVsLx07vHPHDIYrl0OPG4zUREF52igbBPiNrHJFDQJT/4YlDMJmo/QT/A1D6n9ocemvZSzhRx15/Arjowhr+VVKSbaxzPtEfY0oIg2SrqJnnr/l3Du5qIefwh5VmCZe4xopPUaDDoOIEFriZ88sB+3zz8ib8sk8zJJQCgeP78tQvXCgS+4e5W3TUg9mxjB6KjXTyHIVhDZqhqde0OI3Fy1UuVzRUwnBaLjBnAwP5EoFQGRmDYk/rEYe7HTmovLeEBUDQocBQKT4Ripm/xJkkWY7B07K/tfo56dGUCkvyIVXKBInCh+dLK7gZapnd4UWkY0xBYcwo1geMLRq58iFTLA2j/JmpmHXp7m0l7jJii7d44uD3tTIFYThn7NlOnvhLim/YcBK07GMGIN7XwrrKZKmxXaspw6KBWVhzuw1UPxctxshYEaMLfFg/bwOw8HvMPr9VtrElpSB7oiOh91PDIPdPBgHCi7N2QgQ5l/ZDBHieSpNrQ== thomasrodgers&#34;)
 *                     .build())
 *                 .subnetId(&#34;/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-byo/providers/Microsoft.Network/virtualNetworks/my--dev-vnet/subnets/default&#34;)
 *                 .version(versions.applyValue(getAzureVersionsResult -&gt; getAzureVersionsResult.validVersions()[0]))
 *                 .build())
 *             .fleet(AzureClusterFleetArgs.builder()
 *                 .project(&#34;my-project-number&#34;)
 *                 .build())
 *             .location(&#34;us-west1&#34;)
 *             .networking(AzureClusterNetworkingArgs.builder()
 *                 .podAddressCidrBlocks(&#34;10.200.0.0/16&#34;)
 *                 .serviceAddressCidrBlocks(&#34;10.32.0.0/24&#34;)
 *                 .virtualNetworkId(&#34;/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-byo/providers/Microsoft.Network/virtualNetworks/my--dev-vnet&#34;)
 *                 .build())
 *             .project(&#34;my-project-name&#34;)
 *             .resourceGroupId(&#34;/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-cluster&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cluster can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:container/azureCluster:AzureCluster default projects/{{project}}/locations/{{location}}/azureClusters/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:container/azureCluster:AzureCluster default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:container/azureCluster:AzureCluster default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:container/azureCluster:AzureCluster")
public class AzureCluster extends com.pulumi.resources.CustomResource {
    /**
     * Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
     * 
     */
    @Export(name="annotations", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> annotations;

    /**
     * @return Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
     * 
     */
    public Output<Optional<Map<String,String>>> annotations() {
        return Codegen.optional(this.annotations);
    }
    /**
     * Configuration related to the cluster RBAC settings.
     * 
     */
    @Export(name="authorization", type=AzureClusterAuthorization.class, parameters={})
    private Output<AzureClusterAuthorization> authorization;

    /**
     * @return Configuration related to the cluster RBAC settings.
     * 
     */
    public Output<AzureClusterAuthorization> authorization() {
        return this.authorization;
    }
    /**
     * The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
     * 
     */
    @Export(name="azureRegion", type=String.class, parameters={})
    private Output<String> azureRegion;

    /**
     * @return The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
     * 
     */
    public Output<String> azureRegion() {
        return this.azureRegion;
    }
    /**
     * Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the `AzureCluster`. `AzureClient` names are formatted as `projects/&lt;project-number&gt;/locations/&lt;region&gt;/azureClients/&lt;client-id&gt;`. See Resource Names (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
     * 
     */
    @Export(name="client", type=String.class, parameters={})
    private Output<String> client;

    /**
     * @return Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the `AzureCluster`. `AzureClient` names are formatted as `projects/&lt;project-number&gt;/locations/&lt;region&gt;/azureClients/&lt;client-id&gt;`. See Resource Names (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
     * 
     */
    public Output<String> client() {
        return this.client;
    }
    /**
     * Configuration related to the cluster control plane.
     * 
     */
    @Export(name="controlPlane", type=AzureClusterControlPlane.class, parameters={})
    private Output<AzureClusterControlPlane> controlPlane;

    /**
     * @return Configuration related to the cluster control plane.
     * 
     */
    public Output<AzureClusterControlPlane> controlPlane() {
        return this.controlPlane;
    }
    /**
     * Output only. The time at which this cluster was created.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return Output only. The time at which this cluster was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Output only. The endpoint of the cluster&#39;s API server.
     * 
     */
    @Export(name="endpoint", type=String.class, parameters={})
    private Output<String> endpoint;

    /**
     * @return Output only. The endpoint of the cluster&#39;s API server.
     * 
     */
    public Output<String> endpoint() {
        return this.endpoint;
    }
    /**
     * Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update
     * and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update
     * and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Fleet configuration.
     * 
     */
    @Export(name="fleet", type=AzureClusterFleet.class, parameters={})
    private Output<AzureClusterFleet> fleet;

    /**
     * @return Fleet configuration.
     * 
     */
    public Output<AzureClusterFleet> fleet() {
        return this.fleet;
    }
    /**
     * The location for the resource
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return The location for the resource
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * (Beta only) Logging configuration.
     * 
     */
    @Export(name="loggingConfig", type=AzureClusterLoggingConfig.class, parameters={})
    private Output<AzureClusterLoggingConfig> loggingConfig;

    /**
     * @return (Beta only) Logging configuration.
     * 
     */
    public Output<AzureClusterLoggingConfig> loggingConfig() {
        return this.loggingConfig;
    }
    /**
     * The name of this resource.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of this resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Cluster-wide networking configuration.
     * 
     */
    @Export(name="networking", type=AzureClusterNetworking.class, parameters={})
    private Output<AzureClusterNetworking> networking;

    /**
     * @return Cluster-wide networking configuration.
     * 
     */
    public Output<AzureClusterNetworking> networking() {
        return this.networking;
    }
    /**
     * The project for the resource
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The project for the resource
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Output only. If set, there are currently changes in flight to the cluster.
     * 
     */
    @Export(name="reconciling", type=Boolean.class, parameters={})
    private Output<Boolean> reconciling;

    /**
     * @return Output only. If set, there are currently changes in flight to the cluster.
     * 
     */
    public Output<Boolean> reconciling() {
        return this.reconciling;
    }
    /**
     * The ARM ID the of the resource group containing proxy keyvault. Resource group ids are formatted as `/subscriptions/&lt;subscription-id&gt;/resourceGroups/&lt;resource-group-name&gt;`
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The ARM ID the of the resource group containing proxy keyvault. Resource group ids are formatted as `/subscriptions/&lt;subscription-id&gt;/resourceGroups/&lt;resource-group-name&gt;`
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING,
     * STOPPING, ERROR, DEGRADED
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING,
     * STOPPING, ERROR, DEGRADED
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Output only. A globally unique identifier for the cluster.
     * 
     */
    @Export(name="uid", type=String.class, parameters={})
    private Output<String> uid;

    /**
     * @return Output only. A globally unique identifier for the cluster.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Output only. The time at which this cluster was last updated.
     * 
     */
    @Export(name="updateTime", type=String.class, parameters={})
    private Output<String> updateTime;

    /**
     * @return Output only. The time at which this cluster was last updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Output only. Workload Identity settings.
     * 
     */
    @Export(name="workloadIdentityConfigs", type=List.class, parameters={AzureClusterWorkloadIdentityConfig.class})
    private Output<List<AzureClusterWorkloadIdentityConfig>> workloadIdentityConfigs;

    /**
     * @return Output only. Workload Identity settings.
     * 
     */
    public Output<List<AzureClusterWorkloadIdentityConfig>> workloadIdentityConfigs() {
        return this.workloadIdentityConfigs;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AzureCluster(String name) {
        this(name, AzureClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AzureCluster(String name, AzureClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AzureCluster(String name, AzureClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:container/azureCluster:AzureCluster", name, args == null ? AzureClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AzureCluster(String name, Output<String> id, @Nullable AzureClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:container/azureCluster:AzureCluster", name, state, makeResourceOptions(options, id));
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
    public static AzureCluster get(String name, Output<String> id, @Nullable AzureClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AzureCluster(name, id, state, options);
    }
}

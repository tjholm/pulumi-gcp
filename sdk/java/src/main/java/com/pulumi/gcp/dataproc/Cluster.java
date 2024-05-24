// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataproc.ClusterArgs;
import com.pulumi.gcp.dataproc.inputs.ClusterState;
import com.pulumi.gcp.dataproc.outputs.ClusterClusterConfig;
import com.pulumi.gcp.dataproc.outputs.ClusterVirtualClusterConfig;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Cloud Dataproc cluster resource within GCP.
 * 
 * * [API documentation](https://cloud.google.com/dataproc/docs/reference/rest/v1/projects.regions.clusters)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dataproc/docs)
 * 
 * !&gt; **Warning:** Due to limitations of the API, all arguments except
 * `labels`,`cluster_config.worker_config.num_instances` and `cluster_config.preemptible_worker_config.num_instances` are non-updatable. Changing `cluster_config.worker_config.min_num_instances` will be ignored. Changing others will cause recreation of the
 * whole cluster!
 * 
 * ## Example Usage
 * 
 * ### Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataproc.Cluster;
 * import com.pulumi.gcp.dataproc.ClusterArgs;
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
 *         var simplecluster = new Cluster("simplecluster", ClusterArgs.builder()
 *             .name("simplecluster")
 *             .region("us-central1")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Advanced
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.serviceaccount.Account;
 * import com.pulumi.gcp.serviceaccount.AccountArgs;
 * import com.pulumi.gcp.dataproc.Cluster;
 * import com.pulumi.gcp.dataproc.ClusterArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigMasterConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigMasterConfigDiskConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigWorkerConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigWorkerConfigDiskConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigPreemptibleWorkerConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigSoftwareConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigGceClusterConfigArgs;
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
 *         var default_ = new Account("default", AccountArgs.builder()
 *             .accountId("service-account-id")
 *             .displayName("Service Account")
 *             .build());
 * 
 *         var mycluster = new Cluster("mycluster", ClusterArgs.builder()
 *             .name("mycluster")
 *             .region("us-central1")
 *             .gracefulDecommissionTimeout("120s")
 *             .labels(Map.of("foo", "bar"))
 *             .clusterConfig(ClusterClusterConfigArgs.builder()
 *                 .stagingBucket("dataproc-staging-bucket")
 *                 .masterConfig(ClusterClusterConfigMasterConfigArgs.builder()
 *                     .numInstances(1)
 *                     .machineType("e2-medium")
 *                     .diskConfig(ClusterClusterConfigMasterConfigDiskConfigArgs.builder()
 *                         .bootDiskType("pd-ssd")
 *                         .bootDiskSizeGb(30)
 *                         .build())
 *                     .build())
 *                 .workerConfig(ClusterClusterConfigWorkerConfigArgs.builder()
 *                     .numInstances(2)
 *                     .machineType("e2-medium")
 *                     .minCpuPlatform("Intel Skylake")
 *                     .diskConfig(ClusterClusterConfigWorkerConfigDiskConfigArgs.builder()
 *                         .bootDiskSizeGb(30)
 *                         .numLocalSsds(1)
 *                         .build())
 *                     .build())
 *                 .preemptibleWorkerConfig(ClusterClusterConfigPreemptibleWorkerConfigArgs.builder()
 *                     .numInstances(0)
 *                     .build())
 *                 .softwareConfig(ClusterClusterConfigSoftwareConfigArgs.builder()
 *                     .imageVersion("2.0.35-debian10")
 *                     .overrideProperties(Map.of("dataproc:dataproc.allow.zero.workers", "true"))
 *                     .build())
 *                 .gceClusterConfig(ClusterClusterConfigGceClusterConfigArgs.builder()
 *                     .tags(                    
 *                         "foo",
 *                         "bar")
 *                     .serviceAccount(default_.email())
 *                     .serviceAccountScopes("cloud-platform")
 *                     .build())
 *                 .initializationActions(ClusterClusterConfigInitializationActionArgs.builder()
 *                     .script("gs://dataproc-initialization-actions/stackdriver/stackdriver.sh")
 *                     .timeoutSec(500)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Using A GPU Accelerator
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataproc.Cluster;
 * import com.pulumi.gcp.dataproc.ClusterArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigGceClusterConfigArgs;
 * import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigMasterConfigArgs;
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
 *         var acceleratedCluster = new Cluster("acceleratedCluster", ClusterArgs.builder()
 *             .name("my-cluster-with-gpu")
 *             .region("us-central1")
 *             .clusterConfig(ClusterClusterConfigArgs.builder()
 *                 .gceClusterConfig(ClusterClusterConfigGceClusterConfigArgs.builder()
 *                     .zone("us-central1-a")
 *                     .build())
 *                 .masterConfig(ClusterClusterConfigMasterConfigArgs.builder()
 *                     .accelerators(ClusterClusterConfigMasterConfigAcceleratorArgs.builder()
 *                         .acceleratorType("nvidia-tesla-k80")
 *                         .acceleratorCount("1")
 *                         .build())
 *                     .build())
 *                 .build())
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
 * This resource does not support import.
 * 
 */
@ResourceType(type="gcp:dataproc/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    /**
     * Allows you to configure various aspects of the cluster.
     * Structure defined below.
     * 
     */
    @Export(name="clusterConfig", refs={ClusterClusterConfig.class}, tree="[0]")
    private Output<ClusterClusterConfig> clusterConfig;

    /**
     * @return Allows you to configure various aspects of the cluster.
     * Structure defined below.
     * 
     */
    public Output<ClusterClusterConfig> clusterConfig() {
        return this.clusterConfig;
    }
    /**
     * The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    @Export(name="gracefulDecommissionTimeout", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> gracefulDecommissionTimeout;

    public Output<Optional<String>> gracefulDecommissionTimeout() {
        return Codegen.optional(this.gracefulDecommissionTimeout);
    }
    /**
     * The list of the labels (key/value pairs) configured on the resource and to be applied to instances in the cluster.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer
     * to the field &#39;effective_labels&#39; for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return The list of the labels (key/value pairs) configured on the resource and to be applied to instances in the cluster.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer
     * to the field &#39;effective_labels&#39; for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The name of the cluster, unique within the project and
     * zone.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the cluster, unique within the project and
     * zone.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> region;

    /**
     * @return The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * Allows you to configure a virtual Dataproc on GKE cluster.
     * Structure defined below.
     * 
     */
    @Export(name="virtualClusterConfig", refs={ClusterVirtualClusterConfig.class}, tree="[0]")
    private Output<ClusterVirtualClusterConfig> virtualClusterConfig;

    /**
     * @return Allows you to configure a virtual Dataproc on GKE cluster.
     * Structure defined below.
     * 
     */
    public Output<ClusterVirtualClusterConfig> virtualClusterConfig() {
        return this.virtualClusterConfig;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Cluster(String name) {
        this(name, ClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cluster(String name, @Nullable ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, @Nullable ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataproc/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataproc/cluster:Cluster", name, state, makeResourceOptions(options, id));
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
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}

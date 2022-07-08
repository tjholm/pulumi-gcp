// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs extends com.pulumi.resources.ResourceArgs {

    public static final WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs Empty = new WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs();

    /**
     * Optional. A namespace within the GKE cluster to deploy into.
     * 
     */
    @Import(name="clusterNamespace")
    private @Nullable Output<String> clusterNamespace;

    /**
     * @return Optional. A namespace within the GKE cluster to deploy into.
     * 
     */
    public Optional<Output<String>> clusterNamespace() {
        return Optional.ofNullable(this.clusterNamespace);
    }

    /**
     * Optional. The target GKE cluster to deploy to. Format: &#39;projects/{project}/locations/{location}/clusters/{cluster_id}&#39;
     * 
     */
    @Import(name="targetGkeCluster")
    private @Nullable Output<String> targetGkeCluster;

    /**
     * @return Optional. The target GKE cluster to deploy to. Format: &#39;projects/{project}/locations/{location}/clusters/{cluster_id}&#39;
     * 
     */
    public Optional<Output<String>> targetGkeCluster() {
        return Optional.ofNullable(this.targetGkeCluster);
    }

    private WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs() {}

    private WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs(WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs $) {
        this.clusterNamespace = $.clusterNamespace;
        this.targetGkeCluster = $.targetGkeCluster;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs $;

        public Builder() {
            $ = new WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs();
        }

        public Builder(WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs defaults) {
            $ = new WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterNamespace Optional. A namespace within the GKE cluster to deploy into.
         * 
         * @return builder
         * 
         */
        public Builder clusterNamespace(@Nullable Output<String> clusterNamespace) {
            $.clusterNamespace = clusterNamespace;
            return this;
        }

        /**
         * @param clusterNamespace Optional. A namespace within the GKE cluster to deploy into.
         * 
         * @return builder
         * 
         */
        public Builder clusterNamespace(String clusterNamespace) {
            return clusterNamespace(Output.of(clusterNamespace));
        }

        /**
         * @param targetGkeCluster Optional. The target GKE cluster to deploy to. Format: &#39;projects/{project}/locations/{location}/clusters/{cluster_id}&#39;
         * 
         * @return builder
         * 
         */
        public Builder targetGkeCluster(@Nullable Output<String> targetGkeCluster) {
            $.targetGkeCluster = targetGkeCluster;
            return this;
        }

        /**
         * @param targetGkeCluster Optional. The target GKE cluster to deploy to. Format: &#39;projects/{project}/locations/{location}/clusters/{cluster_id}&#39;
         * 
         * @return builder
         * 
         */
        public Builder targetGkeCluster(String targetGkeCluster) {
            return targetGkeCluster(Output.of(targetGkeCluster));
        }

        public WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArgs build() {
            return $;
        }
    }

}

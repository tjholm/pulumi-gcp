// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.container.inputs.ClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterNodePoolDefaultsNodeConfigDefaultsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterNodePoolDefaultsNodeConfigDefaultsArgs Empty = new ClusterNodePoolDefaultsNodeConfigDefaultsArgs();

    /**
     * The default Google Container Filesystem (GCFS) configuration at the cluster level. e.g. enable [image streaming](https://cloud.google.com/kubernetes-engine/docs/how-to/image-streaming) across all the node pools within the cluster. Structure is documented below.
     * 
     */
    @Import(name="gcfsConfig")
    private @Nullable Output<ClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigArgs> gcfsConfig;

    /**
     * @return The default Google Container Filesystem (GCFS) configuration at the cluster level. e.g. enable [image streaming](https://cloud.google.com/kubernetes-engine/docs/how-to/image-streaming) across all the node pools within the cluster. Structure is documented below.
     * 
     */
    public Optional<Output<ClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigArgs>> gcfsConfig() {
        return Optional.ofNullable(this.gcfsConfig);
    }

    /**
     * The type of logging agent that is deployed by default for newly created node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT. See [Increasing logging agent throughput](https://cloud.google.com/stackdriver/docs/solutions/gke/managing-logs#throughput) for more information.
     * 
     */
    @Import(name="loggingVariant")
    private @Nullable Output<String> loggingVariant;

    /**
     * @return The type of logging agent that is deployed by default for newly created node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT. See [Increasing logging agent throughput](https://cloud.google.com/stackdriver/docs/solutions/gke/managing-logs#throughput) for more information.
     * 
     */
    public Optional<Output<String>> loggingVariant() {
        return Optional.ofNullable(this.loggingVariant);
    }

    private ClusterNodePoolDefaultsNodeConfigDefaultsArgs() {}

    private ClusterNodePoolDefaultsNodeConfigDefaultsArgs(ClusterNodePoolDefaultsNodeConfigDefaultsArgs $) {
        this.gcfsConfig = $.gcfsConfig;
        this.loggingVariant = $.loggingVariant;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterNodePoolDefaultsNodeConfigDefaultsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterNodePoolDefaultsNodeConfigDefaultsArgs $;

        public Builder() {
            $ = new ClusterNodePoolDefaultsNodeConfigDefaultsArgs();
        }

        public Builder(ClusterNodePoolDefaultsNodeConfigDefaultsArgs defaults) {
            $ = new ClusterNodePoolDefaultsNodeConfigDefaultsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param gcfsConfig The default Google Container Filesystem (GCFS) configuration at the cluster level. e.g. enable [image streaming](https://cloud.google.com/kubernetes-engine/docs/how-to/image-streaming) across all the node pools within the cluster. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder gcfsConfig(@Nullable Output<ClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigArgs> gcfsConfig) {
            $.gcfsConfig = gcfsConfig;
            return this;
        }

        /**
         * @param gcfsConfig The default Google Container Filesystem (GCFS) configuration at the cluster level. e.g. enable [image streaming](https://cloud.google.com/kubernetes-engine/docs/how-to/image-streaming) across all the node pools within the cluster. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder gcfsConfig(ClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigArgs gcfsConfig) {
            return gcfsConfig(Output.of(gcfsConfig));
        }

        /**
         * @param loggingVariant The type of logging agent that is deployed by default for newly created node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT. See [Increasing logging agent throughput](https://cloud.google.com/stackdriver/docs/solutions/gke/managing-logs#throughput) for more information.
         * 
         * @return builder
         * 
         */
        public Builder loggingVariant(@Nullable Output<String> loggingVariant) {
            $.loggingVariant = loggingVariant;
            return this;
        }

        /**
         * @param loggingVariant The type of logging agent that is deployed by default for newly created node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT. See [Increasing logging agent throughput](https://cloud.google.com/stackdriver/docs/solutions/gke/managing-logs#throughput) for more information.
         * 
         * @return builder
         * 
         */
        public Builder loggingVariant(String loggingVariant) {
            return loggingVariant(Output.of(loggingVariant));
        }

        public ClusterNodePoolDefaultsNodeConfigDefaultsArgs build() {
            return $;
        }
    }

}

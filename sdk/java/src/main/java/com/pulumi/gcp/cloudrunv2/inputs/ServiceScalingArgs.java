// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudrunv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceScalingArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceScalingArgs Empty = new ServiceScalingArgs();

    /**
     * Minimum number of instances for the service, to be divided among all revisions receiving traffic.
     * 
     */
    @Import(name="minInstanceCount")
    private @Nullable Output<Integer> minInstanceCount;

    /**
     * @return Minimum number of instances for the service, to be divided among all revisions receiving traffic.
     * 
     */
    public Optional<Output<Integer>> minInstanceCount() {
        return Optional.ofNullable(this.minInstanceCount);
    }

    private ServiceScalingArgs() {}

    private ServiceScalingArgs(ServiceScalingArgs $) {
        this.minInstanceCount = $.minInstanceCount;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceScalingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceScalingArgs $;

        public Builder() {
            $ = new ServiceScalingArgs();
        }

        public Builder(ServiceScalingArgs defaults) {
            $ = new ServiceScalingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param minInstanceCount Minimum number of instances for the service, to be divided among all revisions receiving traffic.
         * 
         * @return builder
         * 
         */
        public Builder minInstanceCount(@Nullable Output<Integer> minInstanceCount) {
            $.minInstanceCount = minInstanceCount;
            return this;
        }

        /**
         * @param minInstanceCount Minimum number of instances for the service, to be divided among all revisions receiving traffic.
         * 
         * @return builder
         * 
         */
        public Builder minInstanceCount(Integer minInstanceCount) {
            return minInstanceCount(Output.of(minInstanceCount));
        }

        public ServiceScalingArgs build() {
            return $;
        }
    }

}

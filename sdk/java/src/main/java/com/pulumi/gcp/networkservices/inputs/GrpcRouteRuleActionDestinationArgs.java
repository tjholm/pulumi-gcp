// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networkservices.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GrpcRouteRuleActionDestinationArgs extends com.pulumi.resources.ResourceArgs {

    public static final GrpcRouteRuleActionDestinationArgs Empty = new GrpcRouteRuleActionDestinationArgs();

    /**
     * The URL of a BackendService to route traffic to.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The URL of a BackendService to route traffic to.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.
     * 
     */
    @Import(name="weight")
    private @Nullable Output<Integer> weight;

    /**
     * @return Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.
     * 
     */
    public Optional<Output<Integer>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private GrpcRouteRuleActionDestinationArgs() {}

    private GrpcRouteRuleActionDestinationArgs(GrpcRouteRuleActionDestinationArgs $) {
        this.serviceName = $.serviceName;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GrpcRouteRuleActionDestinationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GrpcRouteRuleActionDestinationArgs $;

        public Builder() {
            $ = new GrpcRouteRuleActionDestinationArgs();
        }

        public Builder(GrpcRouteRuleActionDestinationArgs defaults) {
            $ = new GrpcRouteRuleActionDestinationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The URL of a BackendService to route traffic to.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The URL of a BackendService to route traffic to.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param weight Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.
         * 
         * @return builder
         * 
         */
        public Builder weight(@Nullable Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param weight Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.
         * 
         * @return builder
         * 
         */
        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public GrpcRouteRuleActionDestinationArgs build() {
            return $;
        }
    }

}

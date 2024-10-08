// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networkservices.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.networkservices.inputs.GrpcRouteRuleActionArgs;
import com.pulumi.gcp.networkservices.inputs.GrpcRouteRuleMatchArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GrpcRouteRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final GrpcRouteRuleArgs Empty = new GrpcRouteRuleArgs();

    /**
     * Required. A detailed rule defining how to route traffic.
     * Structure is documented below.
     * 
     */
    @Import(name="action")
    private @Nullable Output<GrpcRouteRuleActionArgs> action;

    /**
     * @return Required. A detailed rule defining how to route traffic.
     * Structure is documented below.
     * 
     */
    public Optional<Output<GrpcRouteRuleActionArgs>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * Matches define conditions used for matching the rule against incoming gRPC requests.
     * Structure is documented below.
     * 
     */
    @Import(name="matches")
    private @Nullable Output<List<GrpcRouteRuleMatchArgs>> matches;

    /**
     * @return Matches define conditions used for matching the rule against incoming gRPC requests.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<GrpcRouteRuleMatchArgs>>> matches() {
        return Optional.ofNullable(this.matches);
    }

    private GrpcRouteRuleArgs() {}

    private GrpcRouteRuleArgs(GrpcRouteRuleArgs $) {
        this.action = $.action;
        this.matches = $.matches;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GrpcRouteRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GrpcRouteRuleArgs $;

        public Builder() {
            $ = new GrpcRouteRuleArgs();
        }

        public Builder(GrpcRouteRuleArgs defaults) {
            $ = new GrpcRouteRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param action Required. A detailed rule defining how to route traffic.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<GrpcRouteRuleActionArgs> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action Required. A detailed rule defining how to route traffic.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder action(GrpcRouteRuleActionArgs action) {
            return action(Output.of(action));
        }

        /**
         * @param matches Matches define conditions used for matching the rule against incoming gRPC requests.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder matches(@Nullable Output<List<GrpcRouteRuleMatchArgs>> matches) {
            $.matches = matches;
            return this;
        }

        /**
         * @param matches Matches define conditions used for matching the rule against incoming gRPC requests.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder matches(List<GrpcRouteRuleMatchArgs> matches) {
            return matches(Output.of(matches));
        }

        /**
         * @param matches Matches define conditions used for matching the rule against incoming gRPC requests.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder matches(GrpcRouteRuleMatchArgs... matches) {
            return matches(List.of(matches));
        }

        public GrpcRouteRuleArgs build() {
            return $;
        }
    }

}

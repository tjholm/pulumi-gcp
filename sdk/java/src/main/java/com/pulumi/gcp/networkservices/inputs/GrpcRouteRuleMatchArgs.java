// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networkservices.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.networkservices.inputs.GrpcRouteRuleMatchHeaderArgs;
import com.pulumi.gcp.networkservices.inputs.GrpcRouteRuleMatchMethodArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GrpcRouteRuleMatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final GrpcRouteRuleMatchArgs Empty = new GrpcRouteRuleMatchArgs();

    /**
     * Specifies a list of HTTP request headers to match against.
     * Structure is documented below.
     * 
     */
    @Import(name="headers")
    private @Nullable Output<List<GrpcRouteRuleMatchHeaderArgs>> headers;

    /**
     * @return Specifies a list of HTTP request headers to match against.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<GrpcRouteRuleMatchHeaderArgs>>> headers() {
        return Optional.ofNullable(this.headers);
    }

    /**
     * A gRPC method to match against. If this field is empty or omitted, will match all methods.
     * Structure is documented below.
     * 
     */
    @Import(name="method")
    private @Nullable Output<GrpcRouteRuleMatchMethodArgs> method;

    /**
     * @return A gRPC method to match against. If this field is empty or omitted, will match all methods.
     * Structure is documented below.
     * 
     */
    public Optional<Output<GrpcRouteRuleMatchMethodArgs>> method() {
        return Optional.ofNullable(this.method);
    }

    private GrpcRouteRuleMatchArgs() {}

    private GrpcRouteRuleMatchArgs(GrpcRouteRuleMatchArgs $) {
        this.headers = $.headers;
        this.method = $.method;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GrpcRouteRuleMatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GrpcRouteRuleMatchArgs $;

        public Builder() {
            $ = new GrpcRouteRuleMatchArgs();
        }

        public Builder(GrpcRouteRuleMatchArgs defaults) {
            $ = new GrpcRouteRuleMatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param headers Specifies a list of HTTP request headers to match against.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder headers(@Nullable Output<List<GrpcRouteRuleMatchHeaderArgs>> headers) {
            $.headers = headers;
            return this;
        }

        /**
         * @param headers Specifies a list of HTTP request headers to match against.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder headers(List<GrpcRouteRuleMatchHeaderArgs> headers) {
            return headers(Output.of(headers));
        }

        /**
         * @param headers Specifies a list of HTTP request headers to match against.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder headers(GrpcRouteRuleMatchHeaderArgs... headers) {
            return headers(List.of(headers));
        }

        /**
         * @param method A gRPC method to match against. If this field is empty or omitted, will match all methods.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder method(@Nullable Output<GrpcRouteRuleMatchMethodArgs> method) {
            $.method = method;
            return this;
        }

        /**
         * @param method A gRPC method to match against. If this field is empty or omitted, will match all methods.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder method(GrpcRouteRuleMatchMethodArgs method) {
            return method(Output.of(method));
        }

        public GrpcRouteRuleMatchArgs build() {
            return $;
        }
    }

}

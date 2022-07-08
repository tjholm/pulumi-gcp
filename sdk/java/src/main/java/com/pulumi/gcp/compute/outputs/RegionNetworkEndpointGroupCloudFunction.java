// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RegionNetworkEndpointGroupCloudFunction {
    /**
     * @return A user-defined name of the Cloud Function.
     * The function name is case-sensitive and must be 1-63 characters long.
     * Example value: &#34;func1&#34;.
     * 
     */
    private final @Nullable String function;
    /**
     * @return A template to parse platform-specific fields from a request URL. URL mask allows for routing to multiple resources
     * on the same serverless platform without having to create multiple Network Endpoint Groups and backend resources.
     * The fields parsed by this template are platform-specific and are as follows: API Gateway: The gateway ID,
     * App Engine: The service and version, Cloud Functions: The function name, Cloud Run: The service and tag
     * 
     */
    private final @Nullable String urlMask;

    @CustomType.Constructor
    private RegionNetworkEndpointGroupCloudFunction(
        @CustomType.Parameter("function") @Nullable String function,
        @CustomType.Parameter("urlMask") @Nullable String urlMask) {
        this.function = function;
        this.urlMask = urlMask;
    }

    /**
     * @return A user-defined name of the Cloud Function.
     * The function name is case-sensitive and must be 1-63 characters long.
     * Example value: &#34;func1&#34;.
     * 
     */
    public Optional<String> function() {
        return Optional.ofNullable(this.function);
    }
    /**
     * @return A template to parse platform-specific fields from a request URL. URL mask allows for routing to multiple resources
     * on the same serverless platform without having to create multiple Network Endpoint Groups and backend resources.
     * The fields parsed by this template are platform-specific and are as follows: API Gateway: The gateway ID,
     * App Engine: The service and version, Cloud Functions: The function name, Cloud Run: The service and tag
     * 
     */
    public Optional<String> urlMask() {
        return Optional.ofNullable(this.urlMask);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegionNetworkEndpointGroupCloudFunction defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String function;
        private @Nullable String urlMask;

        public Builder() {
    	      // Empty
        }

        public Builder(RegionNetworkEndpointGroupCloudFunction defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.function = defaults.function;
    	      this.urlMask = defaults.urlMask;
        }

        public Builder function(@Nullable String function) {
            this.function = function;
            return this;
        }
        public Builder urlMask(@Nullable String urlMask) {
            this.urlMask = urlMask;
            return this;
        }        public RegionNetworkEndpointGroupCloudFunction build() {
            return new RegionNetworkEndpointGroupCloudFunction(function, urlMask);
        }
    }
}

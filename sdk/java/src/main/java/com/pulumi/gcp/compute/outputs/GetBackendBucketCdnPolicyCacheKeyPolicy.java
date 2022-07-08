// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBackendBucketCdnPolicyCacheKeyPolicy {
    private final List<String> includeHttpHeaders;
    private final List<String> queryStringWhitelists;

    @CustomType.Constructor
    private GetBackendBucketCdnPolicyCacheKeyPolicy(
        @CustomType.Parameter("includeHttpHeaders") List<String> includeHttpHeaders,
        @CustomType.Parameter("queryStringWhitelists") List<String> queryStringWhitelists) {
        this.includeHttpHeaders = includeHttpHeaders;
        this.queryStringWhitelists = queryStringWhitelists;
    }

    public List<String> includeHttpHeaders() {
        return this.includeHttpHeaders;
    }
    public List<String> queryStringWhitelists() {
        return this.queryStringWhitelists;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackendBucketCdnPolicyCacheKeyPolicy defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private List<String> includeHttpHeaders;
        private List<String> queryStringWhitelists;

        public Builder() {
    	      // Empty
        }

        public Builder(GetBackendBucketCdnPolicyCacheKeyPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.includeHttpHeaders = defaults.includeHttpHeaders;
    	      this.queryStringWhitelists = defaults.queryStringWhitelists;
        }

        public Builder includeHttpHeaders(List<String> includeHttpHeaders) {
            this.includeHttpHeaders = Objects.requireNonNull(includeHttpHeaders);
            return this;
        }
        public Builder includeHttpHeaders(String... includeHttpHeaders) {
            return includeHttpHeaders(List.of(includeHttpHeaders));
        }
        public Builder queryStringWhitelists(List<String> queryStringWhitelists) {
            this.queryStringWhitelists = Objects.requireNonNull(queryStringWhitelists);
            return this;
        }
        public Builder queryStringWhitelists(String... queryStringWhitelists) {
            return queryStringWhitelists(List.of(queryStringWhitelists));
        }        public GetBackendBucketCdnPolicyCacheKeyPolicy build() {
            return new GetBackendBucketCdnPolicyCacheKeyPolicy(includeHttpHeaders, queryStringWhitelists);
        }
    }
}

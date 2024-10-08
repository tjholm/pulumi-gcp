// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.storage.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBucketCor {
    /**
     * @return The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses.
     * 
     */
    private Integer maxAgeSeconds;
    /**
     * @return The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: &#34;*&#34; is permitted in the list of methods, and means &#34;any method&#34;.
     * 
     */
    private List<String> methods;
    /**
     * @return The list of Origins eligible to receive CORS response headers. Note: &#34;*&#34; is permitted in the list of origins, and means &#34;any Origin&#34;.
     * 
     */
    private List<String> origins;
    /**
     * @return The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains.
     * 
     */
    private List<String> responseHeaders;

    private GetBucketCor() {}
    /**
     * @return The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses.
     * 
     */
    public Integer maxAgeSeconds() {
        return this.maxAgeSeconds;
    }
    /**
     * @return The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: &#34;*&#34; is permitted in the list of methods, and means &#34;any method&#34;.
     * 
     */
    public List<String> methods() {
        return this.methods;
    }
    /**
     * @return The list of Origins eligible to receive CORS response headers. Note: &#34;*&#34; is permitted in the list of origins, and means &#34;any Origin&#34;.
     * 
     */
    public List<String> origins() {
        return this.origins;
    }
    /**
     * @return The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains.
     * 
     */
    public List<String> responseHeaders() {
        return this.responseHeaders;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBucketCor defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer maxAgeSeconds;
        private List<String> methods;
        private List<String> origins;
        private List<String> responseHeaders;
        public Builder() {}
        public Builder(GetBucketCor defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.maxAgeSeconds = defaults.maxAgeSeconds;
    	      this.methods = defaults.methods;
    	      this.origins = defaults.origins;
    	      this.responseHeaders = defaults.responseHeaders;
        }

        @CustomType.Setter
        public Builder maxAgeSeconds(Integer maxAgeSeconds) {
            if (maxAgeSeconds == null) {
              throw new MissingRequiredPropertyException("GetBucketCor", "maxAgeSeconds");
            }
            this.maxAgeSeconds = maxAgeSeconds;
            return this;
        }
        @CustomType.Setter
        public Builder methods(List<String> methods) {
            if (methods == null) {
              throw new MissingRequiredPropertyException("GetBucketCor", "methods");
            }
            this.methods = methods;
            return this;
        }
        public Builder methods(String... methods) {
            return methods(List.of(methods));
        }
        @CustomType.Setter
        public Builder origins(List<String> origins) {
            if (origins == null) {
              throw new MissingRequiredPropertyException("GetBucketCor", "origins");
            }
            this.origins = origins;
            return this;
        }
        public Builder origins(String... origins) {
            return origins(List.of(origins));
        }
        @CustomType.Setter
        public Builder responseHeaders(List<String> responseHeaders) {
            if (responseHeaders == null) {
              throw new MissingRequiredPropertyException("GetBucketCor", "responseHeaders");
            }
            this.responseHeaders = responseHeaders;
            return this;
        }
        public Builder responseHeaders(String... responseHeaders) {
            return responseHeaders(List.of(responseHeaders));
        }
        public GetBucketCor build() {
            final var _resultValue = new GetBucketCor();
            _resultValue.maxAgeSeconds = maxAgeSeconds;
            _resultValue.methods = methods;
            _resultValue.origins = origins;
            _resultValue.responseHeaders = responseHeaders;
            return _resultValue;
        }
    }
}

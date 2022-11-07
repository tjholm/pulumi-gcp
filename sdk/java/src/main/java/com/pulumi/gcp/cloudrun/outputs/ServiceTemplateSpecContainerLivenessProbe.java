// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudrun.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.cloudrun.outputs.ServiceTemplateSpecContainerLivenessProbeHttpGet;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTemplateSpecContainerLivenessProbe {
    /**
     * @return Minimum consecutive failures for the probe to be considered failed after
     * having succeeded. Defaults to 3. Minimum value is 1.
     * 
     */
    private @Nullable Integer failureThreshold;
    /**
     * @return HttpGet specifies the http request to perform.
     * Structure is documented below.
     * 
     */
    private @Nullable ServiceTemplateSpecContainerLivenessProbeHttpGet httpGet;
    /**
     * @return Number of seconds after the container has started before the probe is
     * initiated.
     * Defaults to 0 seconds. Minimum value is 0. Maximum value is 3600.
     * 
     */
    private @Nullable Integer initialDelaySeconds;
    /**
     * @return How often (in seconds) to perform the probe.
     * Default to 10 seconds. Minimum value is 1. Maximum value is 3600.
     * 
     */
    private @Nullable Integer periodSeconds;
    /**
     * @return Number of seconds after which the probe times out.
     * Defaults to 1 second. Minimum value is 1. Maximum value is 3600.
     * Must be smaller than period_seconds.
     * 
     */
    private @Nullable Integer timeoutSeconds;

    private ServiceTemplateSpecContainerLivenessProbe() {}
    /**
     * @return Minimum consecutive failures for the probe to be considered failed after
     * having succeeded. Defaults to 3. Minimum value is 1.
     * 
     */
    public Optional<Integer> failureThreshold() {
        return Optional.ofNullable(this.failureThreshold);
    }
    /**
     * @return HttpGet specifies the http request to perform.
     * Structure is documented below.
     * 
     */
    public Optional<ServiceTemplateSpecContainerLivenessProbeHttpGet> httpGet() {
        return Optional.ofNullable(this.httpGet);
    }
    /**
     * @return Number of seconds after the container has started before the probe is
     * initiated.
     * Defaults to 0 seconds. Minimum value is 0. Maximum value is 3600.
     * 
     */
    public Optional<Integer> initialDelaySeconds() {
        return Optional.ofNullable(this.initialDelaySeconds);
    }
    /**
     * @return How often (in seconds) to perform the probe.
     * Default to 10 seconds. Minimum value is 1. Maximum value is 3600.
     * 
     */
    public Optional<Integer> periodSeconds() {
        return Optional.ofNullable(this.periodSeconds);
    }
    /**
     * @return Number of seconds after which the probe times out.
     * Defaults to 1 second. Minimum value is 1. Maximum value is 3600.
     * Must be smaller than period_seconds.
     * 
     */
    public Optional<Integer> timeoutSeconds() {
        return Optional.ofNullable(this.timeoutSeconds);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTemplateSpecContainerLivenessProbe defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer failureThreshold;
        private @Nullable ServiceTemplateSpecContainerLivenessProbeHttpGet httpGet;
        private @Nullable Integer initialDelaySeconds;
        private @Nullable Integer periodSeconds;
        private @Nullable Integer timeoutSeconds;
        public Builder() {}
        public Builder(ServiceTemplateSpecContainerLivenessProbe defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.failureThreshold = defaults.failureThreshold;
    	      this.httpGet = defaults.httpGet;
    	      this.initialDelaySeconds = defaults.initialDelaySeconds;
    	      this.periodSeconds = defaults.periodSeconds;
    	      this.timeoutSeconds = defaults.timeoutSeconds;
        }

        @CustomType.Setter
        public Builder failureThreshold(@Nullable Integer failureThreshold) {
            this.failureThreshold = failureThreshold;
            return this;
        }
        @CustomType.Setter
        public Builder httpGet(@Nullable ServiceTemplateSpecContainerLivenessProbeHttpGet httpGet) {
            this.httpGet = httpGet;
            return this;
        }
        @CustomType.Setter
        public Builder initialDelaySeconds(@Nullable Integer initialDelaySeconds) {
            this.initialDelaySeconds = initialDelaySeconds;
            return this;
        }
        @CustomType.Setter
        public Builder periodSeconds(@Nullable Integer periodSeconds) {
            this.periodSeconds = periodSeconds;
            return this;
        }
        @CustomType.Setter
        public Builder timeoutSeconds(@Nullable Integer timeoutSeconds) {
            this.timeoutSeconds = timeoutSeconds;
            return this;
        }
        public ServiceTemplateSpecContainerLivenessProbe build() {
            final var o = new ServiceTemplateSpecContainerLivenessProbe();
            o.failureThreshold = failureThreshold;
            o.httpGet = httpGet;
            o.initialDelaySeconds = initialDelaySeconds;
            o.periodSeconds = periodSeconds;
            o.timeoutSeconds = timeoutSeconds;
            return o;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.compute.outputs.URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbort;
import com.pulumi.gcp.compute.outputs.URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelay;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class URLMapPathMatcherDefaultRouteActionFaultInjectionPolicy {
    /**
     * @return The specification for how client requests are aborted as part of fault injection.
     * Structure is documented below.
     * 
     */
    private final @Nullable URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbort abort;
    /**
     * @return The specification for how client requests are delayed as part of fault injection, before being sent to a backend service.
     * Structure is documented below.
     * 
     */
    private final @Nullable URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelay delay;

    @CustomType.Constructor
    private URLMapPathMatcherDefaultRouteActionFaultInjectionPolicy(
        @CustomType.Parameter("abort") @Nullable URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbort abort,
        @CustomType.Parameter("delay") @Nullable URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelay delay) {
        this.abort = abort;
        this.delay = delay;
    }

    /**
     * @return The specification for how client requests are aborted as part of fault injection.
     * Structure is documented below.
     * 
     */
    public Optional<URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbort> abort() {
        return Optional.ofNullable(this.abort);
    }
    /**
     * @return The specification for how client requests are delayed as part of fault injection, before being sent to a backend service.
     * Structure is documented below.
     * 
     */
    public Optional<URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelay> delay() {
        return Optional.ofNullable(this.delay);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(URLMapPathMatcherDefaultRouteActionFaultInjectionPolicy defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbort abort;
        private @Nullable URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelay delay;

        public Builder() {
    	      // Empty
        }

        public Builder(URLMapPathMatcherDefaultRouteActionFaultInjectionPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.abort = defaults.abort;
    	      this.delay = defaults.delay;
        }

        public Builder abort(@Nullable URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbort abort) {
            this.abort = abort;
            return this;
        }
        public Builder delay(@Nullable URLMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelay delay) {
            this.delay = delay;
            return this;
        }        public URLMapPathMatcherDefaultRouteActionFaultInjectionPolicy build() {
            return new URLMapPathMatcherDefaultRouteActionFaultInjectionPolicy(abort, delay);
        }
    }
}

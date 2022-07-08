// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortArgs;
import com.pulumi.gcp.compute.inputs.RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs Empty = new RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs();

    /**
     * The specification for how client requests are aborted as part of fault
     * injection.
     * Structure is documented below.
     * 
     */
    @Import(name="abort")
    private @Nullable Output<RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortArgs> abort;

    /**
     * @return The specification for how client requests are aborted as part of fault
     * injection.
     * Structure is documented below.
     * 
     */
    public Optional<Output<RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortArgs>> abort() {
        return Optional.ofNullable(this.abort);
    }

    /**
     * The specification for how client requests are delayed as part of fault
     * injection, before being sent to a backend service.
     * Structure is documented below.
     * 
     */
    @Import(name="delay")
    private @Nullable Output<RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayArgs> delay;

    /**
     * @return The specification for how client requests are delayed as part of fault
     * injection, before being sent to a backend service.
     * Structure is documented below.
     * 
     */
    public Optional<Output<RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayArgs>> delay() {
        return Optional.ofNullable(this.delay);
    }

    private RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs() {}

    private RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs(RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs $) {
        this.abort = $.abort;
        this.delay = $.delay;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs $;

        public Builder() {
            $ = new RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs();
        }

        public Builder(RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs defaults) {
            $ = new RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param abort The specification for how client requests are aborted as part of fault
         * injection.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder abort(@Nullable Output<RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortArgs> abort) {
            $.abort = abort;
            return this;
        }

        /**
         * @param abort The specification for how client requests are aborted as part of fault
         * injection.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder abort(RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortArgs abort) {
            return abort(Output.of(abort));
        }

        /**
         * @param delay The specification for how client requests are delayed as part of fault
         * injection, before being sent to a backend service.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder delay(@Nullable Output<RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayArgs> delay) {
            $.delay = delay;
            return this;
        }

        /**
         * @param delay The specification for how client requests are delayed as part of fault
         * injection, before being sent to a backend service.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder delay(RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayArgs delay) {
            return delay(Output.of(delay));
        }

        public RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs build() {
            return $;
        }
    }

}

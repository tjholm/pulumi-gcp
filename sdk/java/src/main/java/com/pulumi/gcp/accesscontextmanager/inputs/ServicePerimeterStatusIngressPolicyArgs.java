// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.accesscontextmanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimeterStatusIngressPolicyIngressFromArgs;
import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimeterStatusIngressPolicyIngressToArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServicePerimeterStatusIngressPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServicePerimeterStatusIngressPolicyArgs Empty = new ServicePerimeterStatusIngressPolicyArgs();

    /**
     * Defines the conditions on the source of a request causing this `IngressPolicy`
     * to apply.
     * Structure is documented below.
     * 
     */
    @Import(name="ingressFrom")
    private @Nullable Output<ServicePerimeterStatusIngressPolicyIngressFromArgs> ingressFrom;

    /**
     * @return Defines the conditions on the source of a request causing this `IngressPolicy`
     * to apply.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ServicePerimeterStatusIngressPolicyIngressFromArgs>> ingressFrom() {
        return Optional.ofNullable(this.ingressFrom);
    }

    /**
     * Defines the conditions on the `ApiOperation` and request destination that cause
     * this `IngressPolicy` to apply.
     * Structure is documented below.
     * 
     */
    @Import(name="ingressTo")
    private @Nullable Output<ServicePerimeterStatusIngressPolicyIngressToArgs> ingressTo;

    /**
     * @return Defines the conditions on the `ApiOperation` and request destination that cause
     * this `IngressPolicy` to apply.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ServicePerimeterStatusIngressPolicyIngressToArgs>> ingressTo() {
        return Optional.ofNullable(this.ingressTo);
    }

    private ServicePerimeterStatusIngressPolicyArgs() {}

    private ServicePerimeterStatusIngressPolicyArgs(ServicePerimeterStatusIngressPolicyArgs $) {
        this.ingressFrom = $.ingressFrom;
        this.ingressTo = $.ingressTo;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServicePerimeterStatusIngressPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServicePerimeterStatusIngressPolicyArgs $;

        public Builder() {
            $ = new ServicePerimeterStatusIngressPolicyArgs();
        }

        public Builder(ServicePerimeterStatusIngressPolicyArgs defaults) {
            $ = new ServicePerimeterStatusIngressPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ingressFrom Defines the conditions on the source of a request causing this `IngressPolicy`
         * to apply.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder ingressFrom(@Nullable Output<ServicePerimeterStatusIngressPolicyIngressFromArgs> ingressFrom) {
            $.ingressFrom = ingressFrom;
            return this;
        }

        /**
         * @param ingressFrom Defines the conditions on the source of a request causing this `IngressPolicy`
         * to apply.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder ingressFrom(ServicePerimeterStatusIngressPolicyIngressFromArgs ingressFrom) {
            return ingressFrom(Output.of(ingressFrom));
        }

        /**
         * @param ingressTo Defines the conditions on the `ApiOperation` and request destination that cause
         * this `IngressPolicy` to apply.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder ingressTo(@Nullable Output<ServicePerimeterStatusIngressPolicyIngressToArgs> ingressTo) {
            $.ingressTo = ingressTo;
            return this;
        }

        /**
         * @param ingressTo Defines the conditions on the `ApiOperation` and request destination that cause
         * this `IngressPolicy` to apply.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder ingressTo(ServicePerimeterStatusIngressPolicyIngressToArgs ingressTo) {
            return ingressTo(Output.of(ingressTo));
        }

        public ServicePerimeterStatusIngressPolicyArgs build() {
            return $;
        }
    }

}

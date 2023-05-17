// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudrun.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceTemplateSpecContainerStartupProbeTcpSocketArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceTemplateSpecContainerStartupProbeTcpSocketArgs Empty = new ServiceTemplateSpecContainerStartupProbeTcpSocketArgs();

    /**
     * Port number to access on the container. Number must be in the range 1 to 65535.
     * If not specified, defaults to the same value as container.ports[0].containerPort.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Port number to access on the container. Number must be in the range 1 to 65535.
     * If not specified, defaults to the same value as container.ports[0].containerPort.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    private ServiceTemplateSpecContainerStartupProbeTcpSocketArgs() {}

    private ServiceTemplateSpecContainerStartupProbeTcpSocketArgs(ServiceTemplateSpecContainerStartupProbeTcpSocketArgs $) {
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceTemplateSpecContainerStartupProbeTcpSocketArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceTemplateSpecContainerStartupProbeTcpSocketArgs $;

        public Builder() {
            $ = new ServiceTemplateSpecContainerStartupProbeTcpSocketArgs();
        }

        public Builder(ServiceTemplateSpecContainerStartupProbeTcpSocketArgs defaults) {
            $ = new ServiceTemplateSpecContainerStartupProbeTcpSocketArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param port Port number to access on the container. Number must be in the range 1 to 65535.
         * If not specified, defaults to the same value as container.ports[0].containerPort.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port number to access on the container. Number must be in the range 1 to 65535.
         * If not specified, defaults to the same value as container.ports[0].containerPort.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public ServiceTemplateSpecContainerStartupProbeTcpSocketArgs build() {
            return $;
        }
    }

}

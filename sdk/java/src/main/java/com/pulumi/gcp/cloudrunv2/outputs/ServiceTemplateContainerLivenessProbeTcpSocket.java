// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudrunv2.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTemplateContainerLivenessProbeTcpSocket {
    /**
     * @return Port number to access on the container. Must be in the range 1 to 65535.
     * If not specified, defaults to the same value as container.ports[0].containerPort.
     * 
     */
    private @Nullable Integer port;

    private ServiceTemplateContainerLivenessProbeTcpSocket() {}
    /**
     * @return Port number to access on the container. Must be in the range 1 to 65535.
     * If not specified, defaults to the same value as container.ports[0].containerPort.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTemplateContainerLivenessProbeTcpSocket defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer port;
        public Builder() {}
        public Builder(ServiceTemplateContainerLivenessProbeTcpSocket defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        public ServiceTemplateContainerLivenessProbeTcpSocket build() {
            final var o = new ServiceTemplateContainerLivenessProbeTcpSocket();
            o.port = port;
            return o;
        }
    }
}

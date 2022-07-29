// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificatemanager.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CertificateMapGclbTargetIpConfig {
    private final @Nullable String ipAddress;
    private final @Nullable List<Integer> ports;

    @CustomType.Constructor
    private CertificateMapGclbTargetIpConfig(
        @CustomType.Parameter("ipAddress") @Nullable String ipAddress,
        @CustomType.Parameter("ports") @Nullable List<Integer> ports) {
        this.ipAddress = ipAddress;
        this.ports = ports;
    }

    public Optional<String> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }
    public List<Integer> ports() {
        return this.ports == null ? List.of() : this.ports;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CertificateMapGclbTargetIpConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String ipAddress;
        private @Nullable List<Integer> ports;

        public Builder() {
    	      // Empty
        }

        public Builder(CertificateMapGclbTargetIpConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipAddress = defaults.ipAddress;
    	      this.ports = defaults.ports;
        }

        public Builder ipAddress(@Nullable String ipAddress) {
            this.ipAddress = ipAddress;
            return this;
        }
        public Builder ports(@Nullable List<Integer> ports) {
            this.ports = ports;
            return this;
        }
        public Builder ports(Integer... ports) {
            return ports(List.of(ports));
        }        public CertificateMapGclbTargetIpConfig build() {
            return new CertificateMapGclbTargetIpConfig(ipAddress, ports);
        }
    }
}

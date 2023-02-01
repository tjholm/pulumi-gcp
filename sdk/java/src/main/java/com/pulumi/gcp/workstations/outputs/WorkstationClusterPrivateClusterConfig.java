// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.workstations.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class WorkstationClusterPrivateClusterConfig {
    /**
     * @return Hostname for the workstation cluster.
     * This field will be populated only when private endpoint is enabled.
     * To access workstations in the cluster, create a new DNS zone mapping this domain name to an internal IP address and a forwarding rule mapping that address to the service attachment.
     * 
     */
    private @Nullable String clusterHostname;
    /**
     * @return Whether Workstations endpoint is private.
     * 
     */
    private Boolean enablePrivateEndpoint;
    /**
     * @return Service attachment URI for the workstation cluster.
     * The service attachemnt is created when private endpoint is enabled.
     * To access workstations in the cluster, configure access to the managed service using (Private Service Connect)[https://cloud.google.com/vpc/docs/configure-private-service-connect-services].
     * 
     */
    private @Nullable String serviceAttachmentUri;

    private WorkstationClusterPrivateClusterConfig() {}
    /**
     * @return Hostname for the workstation cluster.
     * This field will be populated only when private endpoint is enabled.
     * To access workstations in the cluster, create a new DNS zone mapping this domain name to an internal IP address and a forwarding rule mapping that address to the service attachment.
     * 
     */
    public Optional<String> clusterHostname() {
        return Optional.ofNullable(this.clusterHostname);
    }
    /**
     * @return Whether Workstations endpoint is private.
     * 
     */
    public Boolean enablePrivateEndpoint() {
        return this.enablePrivateEndpoint;
    }
    /**
     * @return Service attachment URI for the workstation cluster.
     * The service attachemnt is created when private endpoint is enabled.
     * To access workstations in the cluster, configure access to the managed service using (Private Service Connect)[https://cloud.google.com/vpc/docs/configure-private-service-connect-services].
     * 
     */
    public Optional<String> serviceAttachmentUri() {
        return Optional.ofNullable(this.serviceAttachmentUri);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WorkstationClusterPrivateClusterConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String clusterHostname;
        private Boolean enablePrivateEndpoint;
        private @Nullable String serviceAttachmentUri;
        public Builder() {}
        public Builder(WorkstationClusterPrivateClusterConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterHostname = defaults.clusterHostname;
    	      this.enablePrivateEndpoint = defaults.enablePrivateEndpoint;
    	      this.serviceAttachmentUri = defaults.serviceAttachmentUri;
        }

        @CustomType.Setter
        public Builder clusterHostname(@Nullable String clusterHostname) {
            this.clusterHostname = clusterHostname;
            return this;
        }
        @CustomType.Setter
        public Builder enablePrivateEndpoint(Boolean enablePrivateEndpoint) {
            this.enablePrivateEndpoint = Objects.requireNonNull(enablePrivateEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder serviceAttachmentUri(@Nullable String serviceAttachmentUri) {
            this.serviceAttachmentUri = serviceAttachmentUri;
            return this;
        }
        public WorkstationClusterPrivateClusterConfig build() {
            final var o = new WorkstationClusterPrivateClusterConfig();
            o.clusterHostname = clusterHostname;
            o.enablePrivateEndpoint = enablePrivateEndpoint;
            o.serviceAttachmentUri = serviceAttachmentUri;
            return o;
        }
    }
}

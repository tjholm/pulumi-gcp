// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.kms.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.kms.outputs.EkmConnectionServiceResolverServerCertificate;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EkmConnectionServiceResolver {
    /**
     * @return Optional. The filter applied to the endpoints of the resolved service. If no filter is specified, all endpoints will be considered. An endpoint will be chosen arbitrarily from the filtered list for each request. For endpoint filter syntax and examples, see https://cloud.google.com/service-directory/docs/reference/rpc/google.cloud.servicedirectory.v1#resolveservicerequest.
     * 
     */
    private @Nullable String endpointFilter;
    /**
     * @return Required. The hostname of the EKM replica used at TLS and HTTP layers.
     * 
     */
    private String hostname;
    /**
     * @return Required. A list of leaf server certificates used to authenticate HTTPS connections to the EKM replica. Currently, a maximum of 10 Certificate is supported.
     * Structure is documented below.
     * 
     */
    private List<EkmConnectionServiceResolverServerCertificate> serverCertificates;
    /**
     * @return Required. The resource name of the Service Directory service pointing to an EKM replica, in the format projects/*&#47;locations/*&#47;namespaces/*&#47;services/*
     * 
     */
    private String serviceDirectoryService;

    private EkmConnectionServiceResolver() {}
    /**
     * @return Optional. The filter applied to the endpoints of the resolved service. If no filter is specified, all endpoints will be considered. An endpoint will be chosen arbitrarily from the filtered list for each request. For endpoint filter syntax and examples, see https://cloud.google.com/service-directory/docs/reference/rpc/google.cloud.servicedirectory.v1#resolveservicerequest.
     * 
     */
    public Optional<String> endpointFilter() {
        return Optional.ofNullable(this.endpointFilter);
    }
    /**
     * @return Required. The hostname of the EKM replica used at TLS and HTTP layers.
     * 
     */
    public String hostname() {
        return this.hostname;
    }
    /**
     * @return Required. A list of leaf server certificates used to authenticate HTTPS connections to the EKM replica. Currently, a maximum of 10 Certificate is supported.
     * Structure is documented below.
     * 
     */
    public List<EkmConnectionServiceResolverServerCertificate> serverCertificates() {
        return this.serverCertificates;
    }
    /**
     * @return Required. The resource name of the Service Directory service pointing to an EKM replica, in the format projects/*&#47;locations/*&#47;namespaces/*&#47;services/*
     * 
     */
    public String serviceDirectoryService() {
        return this.serviceDirectoryService;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EkmConnectionServiceResolver defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String endpointFilter;
        private String hostname;
        private List<EkmConnectionServiceResolverServerCertificate> serverCertificates;
        private String serviceDirectoryService;
        public Builder() {}
        public Builder(EkmConnectionServiceResolver defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endpointFilter = defaults.endpointFilter;
    	      this.hostname = defaults.hostname;
    	      this.serverCertificates = defaults.serverCertificates;
    	      this.serviceDirectoryService = defaults.serviceDirectoryService;
        }

        @CustomType.Setter
        public Builder endpointFilter(@Nullable String endpointFilter) {

            this.endpointFilter = endpointFilter;
            return this;
        }
        @CustomType.Setter
        public Builder hostname(String hostname) {
            if (hostname == null) {
              throw new MissingRequiredPropertyException("EkmConnectionServiceResolver", "hostname");
            }
            this.hostname = hostname;
            return this;
        }
        @CustomType.Setter
        public Builder serverCertificates(List<EkmConnectionServiceResolverServerCertificate> serverCertificates) {
            if (serverCertificates == null) {
              throw new MissingRequiredPropertyException("EkmConnectionServiceResolver", "serverCertificates");
            }
            this.serverCertificates = serverCertificates;
            return this;
        }
        public Builder serverCertificates(EkmConnectionServiceResolverServerCertificate... serverCertificates) {
            return serverCertificates(List.of(serverCertificates));
        }
        @CustomType.Setter
        public Builder serviceDirectoryService(String serviceDirectoryService) {
            if (serviceDirectoryService == null) {
              throw new MissingRequiredPropertyException("EkmConnectionServiceResolver", "serviceDirectoryService");
            }
            this.serviceDirectoryService = serviceDirectoryService;
            return this;
        }
        public EkmConnectionServiceResolver build() {
            final var _resultValue = new EkmConnectionServiceResolver();
            _resultValue.endpointFilter = endpointFilter;
            _resultValue.hostname = hostname;
            _resultValue.serverCertificates = serverCertificates;
            _resultValue.serviceDirectoryService = serviceDirectoryService;
            return _resultValue;
        }
    }
}

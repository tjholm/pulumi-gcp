// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetClusterClusterAutoscalingAutoProvisioningDefault {
    private String bootDiskKmsKey;
    private Integer diskSize;
    private String diskType;
    private String imageType;
    private String minCpuPlatform;
    private List<String> oauthScopes;
    private String serviceAccount;

    private GetClusterClusterAutoscalingAutoProvisioningDefault() {}
    public String bootDiskKmsKey() {
        return this.bootDiskKmsKey;
    }
    public Integer diskSize() {
        return this.diskSize;
    }
    public String diskType() {
        return this.diskType;
    }
    public String imageType() {
        return this.imageType;
    }
    public String minCpuPlatform() {
        return this.minCpuPlatform;
    }
    public List<String> oauthScopes() {
        return this.oauthScopes;
    }
    public String serviceAccount() {
        return this.serviceAccount;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClusterClusterAutoscalingAutoProvisioningDefault defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String bootDiskKmsKey;
        private Integer diskSize;
        private String diskType;
        private String imageType;
        private String minCpuPlatform;
        private List<String> oauthScopes;
        private String serviceAccount;
        public Builder() {}
        public Builder(GetClusterClusterAutoscalingAutoProvisioningDefault defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bootDiskKmsKey = defaults.bootDiskKmsKey;
    	      this.diskSize = defaults.diskSize;
    	      this.diskType = defaults.diskType;
    	      this.imageType = defaults.imageType;
    	      this.minCpuPlatform = defaults.minCpuPlatform;
    	      this.oauthScopes = defaults.oauthScopes;
    	      this.serviceAccount = defaults.serviceAccount;
        }

        @CustomType.Setter
        public Builder bootDiskKmsKey(String bootDiskKmsKey) {
            this.bootDiskKmsKey = Objects.requireNonNull(bootDiskKmsKey);
            return this;
        }
        @CustomType.Setter
        public Builder diskSize(Integer diskSize) {
            this.diskSize = Objects.requireNonNull(diskSize);
            return this;
        }
        @CustomType.Setter
        public Builder diskType(String diskType) {
            this.diskType = Objects.requireNonNull(diskType);
            return this;
        }
        @CustomType.Setter
        public Builder imageType(String imageType) {
            this.imageType = Objects.requireNonNull(imageType);
            return this;
        }
        @CustomType.Setter
        public Builder minCpuPlatform(String minCpuPlatform) {
            this.minCpuPlatform = Objects.requireNonNull(minCpuPlatform);
            return this;
        }
        @CustomType.Setter
        public Builder oauthScopes(List<String> oauthScopes) {
            this.oauthScopes = Objects.requireNonNull(oauthScopes);
            return this;
        }
        public Builder oauthScopes(String... oauthScopes) {
            return oauthScopes(List.of(oauthScopes));
        }
        @CustomType.Setter
        public Builder serviceAccount(String serviceAccount) {
            this.serviceAccount = Objects.requireNonNull(serviceAccount);
            return this;
        }
        public GetClusterClusterAutoscalingAutoProvisioningDefault build() {
            final var o = new GetClusterClusterAutoscalingAutoProvisioningDefault();
            o.bootDiskKmsKey = bootDiskKmsKey;
            o.diskSize = diskSize;
            o.diskType = diskType;
            o.imageType = imageType;
            o.minCpuPlatform = minCpuPlatform;
            o.oauthScopes = oauthScopes;
            o.serviceAccount = serviceAccount;
            return o;
        }
    }
}

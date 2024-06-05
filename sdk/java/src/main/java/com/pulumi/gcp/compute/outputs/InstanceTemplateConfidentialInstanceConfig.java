// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceTemplateConfidentialInstanceConfig {
    /**
     * @return Defines the confidential computing technology the instance uses. SEV is an AMD feature. One of the following values: `SEV`, `SEV_SNP`. `on_host_maintenance` can be set to MIGRATE if `confidential_instance_type` is set to `SEV` and `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM. If `SEV_SNP`, currently `min_cpu_platform` has to be set to `&#34;AMD Milan&#34;` or this will fail to create the VM.
     * 
     */
    private @Nullable String confidentialInstanceType;
    /**
     * @return Defines whether the instance should have confidential compute enabled with AMD SEV. If enabled, `on_host_maintenance` can be set to MIGRATE if `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM.
     * 
     */
    private @Nullable Boolean enableConfidentialCompute;

    private InstanceTemplateConfidentialInstanceConfig() {}
    /**
     * @return Defines the confidential computing technology the instance uses. SEV is an AMD feature. One of the following values: `SEV`, `SEV_SNP`. `on_host_maintenance` can be set to MIGRATE if `confidential_instance_type` is set to `SEV` and `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM. If `SEV_SNP`, currently `min_cpu_platform` has to be set to `&#34;AMD Milan&#34;` or this will fail to create the VM.
     * 
     */
    public Optional<String> confidentialInstanceType() {
        return Optional.ofNullable(this.confidentialInstanceType);
    }
    /**
     * @return Defines whether the instance should have confidential compute enabled with AMD SEV. If enabled, `on_host_maintenance` can be set to MIGRATE if `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM.
     * 
     */
    public Optional<Boolean> enableConfidentialCompute() {
        return Optional.ofNullable(this.enableConfidentialCompute);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceTemplateConfidentialInstanceConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String confidentialInstanceType;
        private @Nullable Boolean enableConfidentialCompute;
        public Builder() {}
        public Builder(InstanceTemplateConfidentialInstanceConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.confidentialInstanceType = defaults.confidentialInstanceType;
    	      this.enableConfidentialCompute = defaults.enableConfidentialCompute;
        }

        @CustomType.Setter
        public Builder confidentialInstanceType(@Nullable String confidentialInstanceType) {

            this.confidentialInstanceType = confidentialInstanceType;
            return this;
        }
        @CustomType.Setter
        public Builder enableConfidentialCompute(@Nullable Boolean enableConfidentialCompute) {

            this.enableConfidentialCompute = enableConfidentialCompute;
            return this;
        }
        public InstanceTemplateConfidentialInstanceConfig build() {
            final var _resultValue = new InstanceTemplateConfidentialInstanceConfig();
            _resultValue.confidentialInstanceType = confidentialInstanceType;
            _resultValue.enableConfidentialCompute = enableConfidentialCompute;
            return _resultValue;
        }
    }
}

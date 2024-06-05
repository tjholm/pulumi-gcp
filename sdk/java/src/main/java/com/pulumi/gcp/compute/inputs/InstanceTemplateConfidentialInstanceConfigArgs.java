// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceTemplateConfidentialInstanceConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceTemplateConfidentialInstanceConfigArgs Empty = new InstanceTemplateConfidentialInstanceConfigArgs();

    /**
     * Defines the confidential computing technology the instance uses. SEV is an AMD feature. One of the following values: `SEV`, `SEV_SNP`. `on_host_maintenance` can be set to MIGRATE if `confidential_instance_type` is set to `SEV` and `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM. If `SEV_SNP`, currently `min_cpu_platform` has to be set to `&#34;AMD Milan&#34;` or this will fail to create the VM.
     * 
     */
    @Import(name="confidentialInstanceType")
    private @Nullable Output<String> confidentialInstanceType;

    /**
     * @return Defines the confidential computing technology the instance uses. SEV is an AMD feature. One of the following values: `SEV`, `SEV_SNP`. `on_host_maintenance` can be set to MIGRATE if `confidential_instance_type` is set to `SEV` and `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM. If `SEV_SNP`, currently `min_cpu_platform` has to be set to `&#34;AMD Milan&#34;` or this will fail to create the VM.
     * 
     */
    public Optional<Output<String>> confidentialInstanceType() {
        return Optional.ofNullable(this.confidentialInstanceType);
    }

    /**
     * Defines whether the instance should have confidential compute enabled with AMD SEV. If enabled, `on_host_maintenance` can be set to MIGRATE if `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM.
     * 
     */
    @Import(name="enableConfidentialCompute")
    private @Nullable Output<Boolean> enableConfidentialCompute;

    /**
     * @return Defines whether the instance should have confidential compute enabled with AMD SEV. If enabled, `on_host_maintenance` can be set to MIGRATE if `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM.
     * 
     */
    public Optional<Output<Boolean>> enableConfidentialCompute() {
        return Optional.ofNullable(this.enableConfidentialCompute);
    }

    private InstanceTemplateConfidentialInstanceConfigArgs() {}

    private InstanceTemplateConfidentialInstanceConfigArgs(InstanceTemplateConfidentialInstanceConfigArgs $) {
        this.confidentialInstanceType = $.confidentialInstanceType;
        this.enableConfidentialCompute = $.enableConfidentialCompute;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceTemplateConfidentialInstanceConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceTemplateConfidentialInstanceConfigArgs $;

        public Builder() {
            $ = new InstanceTemplateConfidentialInstanceConfigArgs();
        }

        public Builder(InstanceTemplateConfidentialInstanceConfigArgs defaults) {
            $ = new InstanceTemplateConfidentialInstanceConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param confidentialInstanceType Defines the confidential computing technology the instance uses. SEV is an AMD feature. One of the following values: `SEV`, `SEV_SNP`. `on_host_maintenance` can be set to MIGRATE if `confidential_instance_type` is set to `SEV` and `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM. If `SEV_SNP`, currently `min_cpu_platform` has to be set to `&#34;AMD Milan&#34;` or this will fail to create the VM.
         * 
         * @return builder
         * 
         */
        public Builder confidentialInstanceType(@Nullable Output<String> confidentialInstanceType) {
            $.confidentialInstanceType = confidentialInstanceType;
            return this;
        }

        /**
         * @param confidentialInstanceType Defines the confidential computing technology the instance uses. SEV is an AMD feature. One of the following values: `SEV`, `SEV_SNP`. `on_host_maintenance` can be set to MIGRATE if `confidential_instance_type` is set to `SEV` and `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM. If `SEV_SNP`, currently `min_cpu_platform` has to be set to `&#34;AMD Milan&#34;` or this will fail to create the VM.
         * 
         * @return builder
         * 
         */
        public Builder confidentialInstanceType(String confidentialInstanceType) {
            return confidentialInstanceType(Output.of(confidentialInstanceType));
        }

        /**
         * @param enableConfidentialCompute Defines whether the instance should have confidential compute enabled with AMD SEV. If enabled, `on_host_maintenance` can be set to MIGRATE if `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM.
         * 
         * @return builder
         * 
         */
        public Builder enableConfidentialCompute(@Nullable Output<Boolean> enableConfidentialCompute) {
            $.enableConfidentialCompute = enableConfidentialCompute;
            return this;
        }

        /**
         * @param enableConfidentialCompute Defines whether the instance should have confidential compute enabled with AMD SEV. If enabled, `on_host_maintenance` can be set to MIGRATE if `min_cpu_platform` is set to `&#34;AMD Milan&#34;`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM.
         * 
         * @return builder
         * 
         */
        public Builder enableConfidentialCompute(Boolean enableConfidentialCompute) {
            return enableConfidentialCompute(Output.of(enableConfidentialCompute));
        }

        public InstanceTemplateConfidentialInstanceConfigArgs build() {
            return $;
        }
    }

}

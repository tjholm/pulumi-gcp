// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudrun.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTemplateSpecVolumeEmptyDir {
    /**
     * @return The medium on which the data is stored. The default is &#34;&#34; which means to use the node&#39;s default medium. Must be an empty string (default) or Memory.
     * 
     */
    private @Nullable String medium;
    /**
     * @return Limit on the storage usable by this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. This field&#39;s values are of the &#39;Quantity&#39; k8s type: https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir.
     * 
     * ***
     * 
     */
    private @Nullable String sizeLimit;

    private ServiceTemplateSpecVolumeEmptyDir() {}
    /**
     * @return The medium on which the data is stored. The default is &#34;&#34; which means to use the node&#39;s default medium. Must be an empty string (default) or Memory.
     * 
     */
    public Optional<String> medium() {
        return Optional.ofNullable(this.medium);
    }
    /**
     * @return Limit on the storage usable by this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. This field&#39;s values are of the &#39;Quantity&#39; k8s type: https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir.
     * 
     * ***
     * 
     */
    public Optional<String> sizeLimit() {
        return Optional.ofNullable(this.sizeLimit);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTemplateSpecVolumeEmptyDir defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String medium;
        private @Nullable String sizeLimit;
        public Builder() {}
        public Builder(ServiceTemplateSpecVolumeEmptyDir defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.medium = defaults.medium;
    	      this.sizeLimit = defaults.sizeLimit;
        }

        @CustomType.Setter
        public Builder medium(@Nullable String medium) {
            this.medium = medium;
            return this;
        }
        @CustomType.Setter
        public Builder sizeLimit(@Nullable String sizeLimit) {
            this.sizeLimit = sizeLimit;
            return this;
        }
        public ServiceTemplateSpecVolumeEmptyDir build() {
            final var o = new ServiceTemplateSpecVolumeEmptyDir();
            o.medium = medium;
            o.sizeLimit = sizeLimit;
            return o;
        }
    }
}

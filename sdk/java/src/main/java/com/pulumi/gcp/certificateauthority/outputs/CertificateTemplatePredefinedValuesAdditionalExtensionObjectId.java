// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;

@CustomType
public final class CertificateTemplatePredefinedValuesAdditionalExtensionObjectId {
    /**
     * @return Required. The parts of an OID path. The most significant parts of the path come first.
     * 
     */
    private List<Integer> objectIdPaths;

    private CertificateTemplatePredefinedValuesAdditionalExtensionObjectId() {}
    /**
     * @return Required. The parts of an OID path. The most significant parts of the path come first.
     * 
     */
    public List<Integer> objectIdPaths() {
        return this.objectIdPaths;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CertificateTemplatePredefinedValuesAdditionalExtensionObjectId defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<Integer> objectIdPaths;
        public Builder() {}
        public Builder(CertificateTemplatePredefinedValuesAdditionalExtensionObjectId defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.objectIdPaths = defaults.objectIdPaths;
        }

        @CustomType.Setter
        public Builder objectIdPaths(List<Integer> objectIdPaths) {
            if (objectIdPaths == null) {
              throw new MissingRequiredPropertyException("CertificateTemplatePredefinedValuesAdditionalExtensionObjectId", "objectIdPaths");
            }
            this.objectIdPaths = objectIdPaths;
            return this;
        }
        public Builder objectIdPaths(Integer... objectIdPaths) {
            return objectIdPaths(List.of(objectIdPaths));
        }
        public CertificateTemplatePredefinedValuesAdditionalExtensionObjectId build() {
            final var _resultValue = new CertificateTemplatePredefinedValuesAdditionalExtensionObjectId();
            _resultValue.objectIdPaths = objectIdPaths;
            return _resultValue;
        }
    }
}

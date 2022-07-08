// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;


public final class CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs extends com.pulumi.resources.ResourceArgs {

    public static final CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs Empty = new CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs();

    /**
     * Required. The parts of an OID path. The most significant parts of the path come first.
     * 
     */
    @Import(name="objectIdPaths", required=true)
    private Output<List<Integer>> objectIdPaths;

    /**
     * @return Required. The parts of an OID path. The most significant parts of the path come first.
     * 
     */
    public Output<List<Integer>> objectIdPaths() {
        return this.objectIdPaths;
    }

    private CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs() {}

    private CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs(CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs $) {
        this.objectIdPaths = $.objectIdPaths;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs $;

        public Builder() {
            $ = new CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs();
        }

        public Builder(CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs defaults) {
            $ = new CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param objectIdPaths Required. The parts of an OID path. The most significant parts of the path come first.
         * 
         * @return builder
         * 
         */
        public Builder objectIdPaths(Output<List<Integer>> objectIdPaths) {
            $.objectIdPaths = objectIdPaths;
            return this;
        }

        /**
         * @param objectIdPaths Required. The parts of an OID path. The most significant parts of the path come first.
         * 
         * @return builder
         * 
         */
        public Builder objectIdPaths(List<Integer> objectIdPaths) {
            return objectIdPaths(Output.of(objectIdPaths));
        }

        /**
         * @param objectIdPaths Required. The parts of an OID path. The most significant parts of the path come first.
         * 
         * @return builder
         * 
         */
        public Builder objectIdPaths(Integer... objectIdPaths) {
            return objectIdPaths(List.of(objectIdPaths));
        }

        public CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs build() {
            $.objectIdPaths = Objects.requireNonNull($.objectIdPaths, "expected parameter 'objectIdPaths' to be non-null");
            return $;
        }
    }

}

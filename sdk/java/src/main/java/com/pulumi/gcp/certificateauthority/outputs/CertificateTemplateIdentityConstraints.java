// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.certificateauthority.outputs.CertificateTemplateIdentityConstraintsCelExpression;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CertificateTemplateIdentityConstraints {
    /**
     * @return Required. If this is true, the SubjectAltNames extension may be copied from a certificate request into the signed certificate. Otherwise, the requested SubjectAltNames will be discarded.
     * 
     */
    private Boolean allowSubjectAltNamesPassthrough;
    /**
     * @return Required. If this is true, the Subject field may be copied from a certificate request into the signed certificate. Otherwise, the requested Subject will be discarded.
     * 
     */
    private Boolean allowSubjectPassthrough;
    /**
     * @return Optional. A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a certificate is signed. To see the full allowed syntax and some examples, see https://cloud.google.com/certificate-authority-service/docs/using-cel
     * Structure is documented below.
     * 
     */
    private @Nullable CertificateTemplateIdentityConstraintsCelExpression celExpression;

    private CertificateTemplateIdentityConstraints() {}
    /**
     * @return Required. If this is true, the SubjectAltNames extension may be copied from a certificate request into the signed certificate. Otherwise, the requested SubjectAltNames will be discarded.
     * 
     */
    public Boolean allowSubjectAltNamesPassthrough() {
        return this.allowSubjectAltNamesPassthrough;
    }
    /**
     * @return Required. If this is true, the Subject field may be copied from a certificate request into the signed certificate. Otherwise, the requested Subject will be discarded.
     * 
     */
    public Boolean allowSubjectPassthrough() {
        return this.allowSubjectPassthrough;
    }
    /**
     * @return Optional. A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a certificate is signed. To see the full allowed syntax and some examples, see https://cloud.google.com/certificate-authority-service/docs/using-cel
     * Structure is documented below.
     * 
     */
    public Optional<CertificateTemplateIdentityConstraintsCelExpression> celExpression() {
        return Optional.ofNullable(this.celExpression);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CertificateTemplateIdentityConstraints defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean allowSubjectAltNamesPassthrough;
        private Boolean allowSubjectPassthrough;
        private @Nullable CertificateTemplateIdentityConstraintsCelExpression celExpression;
        public Builder() {}
        public Builder(CertificateTemplateIdentityConstraints defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowSubjectAltNamesPassthrough = defaults.allowSubjectAltNamesPassthrough;
    	      this.allowSubjectPassthrough = defaults.allowSubjectPassthrough;
    	      this.celExpression = defaults.celExpression;
        }

        @CustomType.Setter
        public Builder allowSubjectAltNamesPassthrough(Boolean allowSubjectAltNamesPassthrough) {
            if (allowSubjectAltNamesPassthrough == null) {
              throw new MissingRequiredPropertyException("CertificateTemplateIdentityConstraints", "allowSubjectAltNamesPassthrough");
            }
            this.allowSubjectAltNamesPassthrough = allowSubjectAltNamesPassthrough;
            return this;
        }
        @CustomType.Setter
        public Builder allowSubjectPassthrough(Boolean allowSubjectPassthrough) {
            if (allowSubjectPassthrough == null) {
              throw new MissingRequiredPropertyException("CertificateTemplateIdentityConstraints", "allowSubjectPassthrough");
            }
            this.allowSubjectPassthrough = allowSubjectPassthrough;
            return this;
        }
        @CustomType.Setter
        public Builder celExpression(@Nullable CertificateTemplateIdentityConstraintsCelExpression celExpression) {

            this.celExpression = celExpression;
            return this;
        }
        public CertificateTemplateIdentityConstraints build() {
            final var _resultValue = new CertificateTemplateIdentityConstraints();
            _resultValue.allowSubjectAltNamesPassthrough = allowSubjectAltNamesPassthrough;
            _resultValue.allowSubjectPassthrough = allowSubjectPassthrough;
            _resultValue.celExpression = celExpression;
            return _resultValue;
        }
    }
}

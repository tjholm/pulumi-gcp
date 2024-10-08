// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataplex.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetZoneIamPolicyResult {
    private String dataplexZone;
    /**
     * @return (Computed) The etag of the IAM policy.
     * 
     */
    private String etag;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String lake;
    private String location;
    /**
     * @return (Required only by `gcp.dataplex.ZoneIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    private String policyData;
    private String project;

    private GetZoneIamPolicyResult() {}
    public String dataplexZone() {
        return this.dataplexZone;
    }
    /**
     * @return (Computed) The etag of the IAM policy.
     * 
     */
    public String etag() {
        return this.etag;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String lake() {
        return this.lake;
    }
    public String location() {
        return this.location;
    }
    /**
     * @return (Required only by `gcp.dataplex.ZoneIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    public String policyData() {
        return this.policyData;
    }
    public String project() {
        return this.project;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetZoneIamPolicyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String dataplexZone;
        private String etag;
        private String id;
        private String lake;
        private String location;
        private String policyData;
        private String project;
        public Builder() {}
        public Builder(GetZoneIamPolicyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dataplexZone = defaults.dataplexZone;
    	      this.etag = defaults.etag;
    	      this.id = defaults.id;
    	      this.lake = defaults.lake;
    	      this.location = defaults.location;
    	      this.policyData = defaults.policyData;
    	      this.project = defaults.project;
        }

        @CustomType.Setter
        public Builder dataplexZone(String dataplexZone) {
            if (dataplexZone == null) {
              throw new MissingRequiredPropertyException("GetZoneIamPolicyResult", "dataplexZone");
            }
            this.dataplexZone = dataplexZone;
            return this;
        }
        @CustomType.Setter
        public Builder etag(String etag) {
            if (etag == null) {
              throw new MissingRequiredPropertyException("GetZoneIamPolicyResult", "etag");
            }
            this.etag = etag;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetZoneIamPolicyResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder lake(String lake) {
            if (lake == null) {
              throw new MissingRequiredPropertyException("GetZoneIamPolicyResult", "lake");
            }
            this.lake = lake;
            return this;
        }
        @CustomType.Setter
        public Builder location(String location) {
            if (location == null) {
              throw new MissingRequiredPropertyException("GetZoneIamPolicyResult", "location");
            }
            this.location = location;
            return this;
        }
        @CustomType.Setter
        public Builder policyData(String policyData) {
            if (policyData == null) {
              throw new MissingRequiredPropertyException("GetZoneIamPolicyResult", "policyData");
            }
            this.policyData = policyData;
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            if (project == null) {
              throw new MissingRequiredPropertyException("GetZoneIamPolicyResult", "project");
            }
            this.project = project;
            return this;
        }
        public GetZoneIamPolicyResult build() {
            final var _resultValue = new GetZoneIamPolicyResult();
            _resultValue.dataplexZone = dataplexZone;
            _resultValue.etag = etag;
            _resultValue.id = id;
            _resultValue.lake = lake;
            _resultValue.location = location;
            _resultValue.policyData = policyData;
            _resultValue.project = project;
            return _resultValue;
        }
    }
}

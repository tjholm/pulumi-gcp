// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.filestore.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.filestore.outputs.GetInstanceFileShare;
import com.pulumi.gcp.filestore.outputs.GetInstanceNetwork;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetInstanceResult {
    private String createTime;
    private String description;
    private Map<String,String> effectiveLabels;
    private String etag;
    private List<GetInstanceFileShare> fileShares;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String kmsKeyName;
    private Map<String,String> labels;
    private @Nullable String location;
    private String name;
    private List<GetInstanceNetwork> networks;
    private @Nullable String project;
    private Map<String,String> pulumiLabels;
    private String tier;
    private String zone;

    private GetInstanceResult() {}
    public String createTime() {
        return this.createTime;
    }
    public String description() {
        return this.description;
    }
    public Map<String,String> effectiveLabels() {
        return this.effectiveLabels;
    }
    public String etag() {
        return this.etag;
    }
    public List<GetInstanceFileShare> fileShares() {
        return this.fileShares;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String kmsKeyName() {
        return this.kmsKeyName;
    }
    public Map<String,String> labels() {
        return this.labels;
    }
    public Optional<String> location() {
        return Optional.ofNullable(this.location);
    }
    public String name() {
        return this.name;
    }
    public List<GetInstanceNetwork> networks() {
        return this.networks;
    }
    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }
    public Map<String,String> pulumiLabels() {
        return this.pulumiLabels;
    }
    public String tier() {
        return this.tier;
    }
    public String zone() {
        return this.zone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createTime;
        private String description;
        private Map<String,String> effectiveLabels;
        private String etag;
        private List<GetInstanceFileShare> fileShares;
        private String id;
        private String kmsKeyName;
        private Map<String,String> labels;
        private @Nullable String location;
        private String name;
        private List<GetInstanceNetwork> networks;
        private @Nullable String project;
        private Map<String,String> pulumiLabels;
        private String tier;
        private String zone;
        public Builder() {}
        public Builder(GetInstanceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createTime = defaults.createTime;
    	      this.description = defaults.description;
    	      this.effectiveLabels = defaults.effectiveLabels;
    	      this.etag = defaults.etag;
    	      this.fileShares = defaults.fileShares;
    	      this.id = defaults.id;
    	      this.kmsKeyName = defaults.kmsKeyName;
    	      this.labels = defaults.labels;
    	      this.location = defaults.location;
    	      this.name = defaults.name;
    	      this.networks = defaults.networks;
    	      this.project = defaults.project;
    	      this.pulumiLabels = defaults.pulumiLabels;
    	      this.tier = defaults.tier;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder createTime(String createTime) {
            if (createTime == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "createTime");
            }
            this.createTime = createTime;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder effectiveLabels(Map<String,String> effectiveLabels) {
            if (effectiveLabels == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "effectiveLabels");
            }
            this.effectiveLabels = effectiveLabels;
            return this;
        }
        @CustomType.Setter
        public Builder etag(String etag) {
            if (etag == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "etag");
            }
            this.etag = etag;
            return this;
        }
        @CustomType.Setter
        public Builder fileShares(List<GetInstanceFileShare> fileShares) {
            if (fileShares == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "fileShares");
            }
            this.fileShares = fileShares;
            return this;
        }
        public Builder fileShares(GetInstanceFileShare... fileShares) {
            return fileShares(List.of(fileShares));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder kmsKeyName(String kmsKeyName) {
            if (kmsKeyName == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "kmsKeyName");
            }
            this.kmsKeyName = kmsKeyName;
            return this;
        }
        @CustomType.Setter
        public Builder labels(Map<String,String> labels) {
            if (labels == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "labels");
            }
            this.labels = labels;
            return this;
        }
        @CustomType.Setter
        public Builder location(@Nullable String location) {

            this.location = location;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder networks(List<GetInstanceNetwork> networks) {
            if (networks == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "networks");
            }
            this.networks = networks;
            return this;
        }
        public Builder networks(GetInstanceNetwork... networks) {
            return networks(List.of(networks));
        }
        @CustomType.Setter
        public Builder project(@Nullable String project) {

            this.project = project;
            return this;
        }
        @CustomType.Setter
        public Builder pulumiLabels(Map<String,String> pulumiLabels) {
            if (pulumiLabels == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "pulumiLabels");
            }
            this.pulumiLabels = pulumiLabels;
            return this;
        }
        @CustomType.Setter
        public Builder tier(String tier) {
            if (tier == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "tier");
            }
            this.tier = tier;
            return this;
        }
        @CustomType.Setter
        public Builder zone(String zone) {
            if (zone == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "zone");
            }
            this.zone = zone;
            return this;
        }
        public GetInstanceResult build() {
            final var _resultValue = new GetInstanceResult();
            _resultValue.createTime = createTime;
            _resultValue.description = description;
            _resultValue.effectiveLabels = effectiveLabels;
            _resultValue.etag = etag;
            _resultValue.fileShares = fileShares;
            _resultValue.id = id;
            _resultValue.kmsKeyName = kmsKeyName;
            _resultValue.labels = labels;
            _resultValue.location = location;
            _resultValue.name = name;
            _resultValue.networks = networks;
            _resultValue.project = project;
            _resultValue.pulumiLabels = pulumiLabels;
            _resultValue.tier = tier;
            _resultValue.zone = zone;
            return _resultValue;
        }
    }
}

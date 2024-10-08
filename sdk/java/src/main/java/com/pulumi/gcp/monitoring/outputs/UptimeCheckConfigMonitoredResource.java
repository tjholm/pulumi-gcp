// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.monitoring.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class UptimeCheckConfigMonitoredResource {
    /**
     * @return Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels `project_id`, `instance_id`, and `zone`.
     * 
     */
    private Map<String,String> labels;
    /**
     * @return The monitored resource type. This field must match the type field of a [`MonitoredResourceDescriptor`](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is `gce_instance`. For a list of types, see [Monitoring resource types](https://cloud.google.com/monitoring/api/resources) and [Logging resource types](https://cloud.google.com/logging/docs/api/v2/resource-list).
     * 
     */
    private String type;

    private UptimeCheckConfigMonitoredResource() {}
    /**
     * @return Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels `project_id`, `instance_id`, and `zone`.
     * 
     */
    public Map<String,String> labels() {
        return this.labels;
    }
    /**
     * @return The monitored resource type. This field must match the type field of a [`MonitoredResourceDescriptor`](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is `gce_instance`. For a list of types, see [Monitoring resource types](https://cloud.google.com/monitoring/api/resources) and [Logging resource types](https://cloud.google.com/logging/docs/api/v2/resource-list).
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UptimeCheckConfigMonitoredResource defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> labels;
        private String type;
        public Builder() {}
        public Builder(UptimeCheckConfigMonitoredResource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.labels = defaults.labels;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder labels(Map<String,String> labels) {
            if (labels == null) {
              throw new MissingRequiredPropertyException("UptimeCheckConfigMonitoredResource", "labels");
            }
            this.labels = labels;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("UptimeCheckConfigMonitoredResource", "type");
            }
            this.type = type;
            return this;
        }
        public UptimeCheckConfigMonitoredResource build() {
            final var _resultValue = new UptimeCheckConfigMonitoredResource();
            _resultValue.labels = labels;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}

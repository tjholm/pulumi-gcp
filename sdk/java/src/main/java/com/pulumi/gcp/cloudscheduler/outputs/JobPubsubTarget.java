// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudscheduler.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class JobPubsubTarget {
    /**
     * @return Attributes for PubsubMessage.
     * Pubsub message must contain either non-empty data, or at least one attribute.
     * 
     */
    private final @Nullable Map<String,String> attributes;
    /**
     * @return The message payload for PubsubMessage.
     * Pubsub message must contain either non-empty data, or at least one attribute.
     * A base64-encoded string.
     * 
     */
    private final @Nullable String data;
    /**
     * @return The full resource name for the Cloud Pub/Sub topic to which
     * messages will be published when a job is delivered. ~&gt;**NOTE:**
     * The topic name must be in the same format as required by PubSub&#39;s
     * PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.
     * 
     */
    private final String topicName;

    @CustomType.Constructor
    private JobPubsubTarget(
        @CustomType.Parameter("attributes") @Nullable Map<String,String> attributes,
        @CustomType.Parameter("data") @Nullable String data,
        @CustomType.Parameter("topicName") String topicName) {
        this.attributes = attributes;
        this.data = data;
        this.topicName = topicName;
    }

    /**
     * @return Attributes for PubsubMessage.
     * Pubsub message must contain either non-empty data, or at least one attribute.
     * 
     */
    public Map<String,String> attributes() {
        return this.attributes == null ? Map.of() : this.attributes;
    }
    /**
     * @return The message payload for PubsubMessage.
     * Pubsub message must contain either non-empty data, or at least one attribute.
     * A base64-encoded string.
     * 
     */
    public Optional<String> data() {
        return Optional.ofNullable(this.data);
    }
    /**
     * @return The full resource name for the Cloud Pub/Sub topic to which
     * messages will be published when a job is delivered. ~&gt;**NOTE:**
     * The topic name must be in the same format as required by PubSub&#39;s
     * PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.
     * 
     */
    public String topicName() {
        return this.topicName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(JobPubsubTarget defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Map<String,String> attributes;
        private @Nullable String data;
        private String topicName;

        public Builder() {
    	      // Empty
        }

        public Builder(JobPubsubTarget defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attributes = defaults.attributes;
    	      this.data = defaults.data;
    	      this.topicName = defaults.topicName;
        }

        public Builder attributes(@Nullable Map<String,String> attributes) {
            this.attributes = attributes;
            return this;
        }
        public Builder data(@Nullable String data) {
            this.data = data;
            return this;
        }
        public Builder topicName(String topicName) {
            this.topicName = Objects.requireNonNull(topicName);
            return this;
        }        public JobPubsubTarget build() {
            return new JobPubsubTarget(attributes, data, topicName);
        }
    }
}

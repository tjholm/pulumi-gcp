// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.kms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RegistryEventNotificationConfigItem {
    /**
     * @return PubSub topic name to publish device events.
     * 
     */
    private final String pubsubTopicName;
    /**
     * @return If the subfolder name matches this string exactly, this
     * configuration will be used. The string must not include the
     * leading &#39;/&#39; character. If empty, all strings are matched. Empty
     * value can only be used for the last `event_notification_configs`
     * item.
     * 
     */
    private final @Nullable String subfolderMatches;

    @CustomType.Constructor
    private RegistryEventNotificationConfigItem(
        @CustomType.Parameter("pubsubTopicName") String pubsubTopicName,
        @CustomType.Parameter("subfolderMatches") @Nullable String subfolderMatches) {
        this.pubsubTopicName = pubsubTopicName;
        this.subfolderMatches = subfolderMatches;
    }

    /**
     * @return PubSub topic name to publish device events.
     * 
     */
    public String pubsubTopicName() {
        return this.pubsubTopicName;
    }
    /**
     * @return If the subfolder name matches this string exactly, this
     * configuration will be used. The string must not include the
     * leading &#39;/&#39; character. If empty, all strings are matched. Empty
     * value can only be used for the last `event_notification_configs`
     * item.
     * 
     */
    public Optional<String> subfolderMatches() {
        return Optional.ofNullable(this.subfolderMatches);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegistryEventNotificationConfigItem defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String pubsubTopicName;
        private @Nullable String subfolderMatches;

        public Builder() {
    	      // Empty
        }

        public Builder(RegistryEventNotificationConfigItem defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.pubsubTopicName = defaults.pubsubTopicName;
    	      this.subfolderMatches = defaults.subfolderMatches;
        }

        public Builder pubsubTopicName(String pubsubTopicName) {
            this.pubsubTopicName = Objects.requireNonNull(pubsubTopicName);
            return this;
        }
        public Builder subfolderMatches(@Nullable String subfolderMatches) {
            this.subfolderMatches = subfolderMatches;
            return this;
        }        public RegistryEventNotificationConfigItem build() {
            return new RegistryEventNotificationConfigItem(pubsubTopicName, subfolderMatches);
        }
    }
}

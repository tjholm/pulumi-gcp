// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TableTimePartitioning {
    /**
     * @return Number of milliseconds for which to keep the
     * storage for a partition.
     * 
     */
    private final @Nullable Integer expirationMs;
    /**
     * @return The field used to determine how to create a range-based
     * partition.
     * 
     */
    private final @Nullable String field;
    /**
     * @return If set to true, queries over this table
     * require a partition filter that can be used for partition elimination to be
     * specified.
     * 
     */
    private final @Nullable Boolean requirePartitionFilter;
    /**
     * @return The supported types are DAY, HOUR, MONTH, and YEAR,
     * which will generate one partition per day, hour, month, and year, respectively.
     * 
     */
    private final String type;

    @CustomType.Constructor
    private TableTimePartitioning(
        @CustomType.Parameter("expirationMs") @Nullable Integer expirationMs,
        @CustomType.Parameter("field") @Nullable String field,
        @CustomType.Parameter("requirePartitionFilter") @Nullable Boolean requirePartitionFilter,
        @CustomType.Parameter("type") String type) {
        this.expirationMs = expirationMs;
        this.field = field;
        this.requirePartitionFilter = requirePartitionFilter;
        this.type = type;
    }

    /**
     * @return Number of milliseconds for which to keep the
     * storage for a partition.
     * 
     */
    public Optional<Integer> expirationMs() {
        return Optional.ofNullable(this.expirationMs);
    }
    /**
     * @return The field used to determine how to create a range-based
     * partition.
     * 
     */
    public Optional<String> field() {
        return Optional.ofNullable(this.field);
    }
    /**
     * @return If set to true, queries over this table
     * require a partition filter that can be used for partition elimination to be
     * specified.
     * 
     */
    public Optional<Boolean> requirePartitionFilter() {
        return Optional.ofNullable(this.requirePartitionFilter);
    }
    /**
     * @return The supported types are DAY, HOUR, MONTH, and YEAR,
     * which will generate one partition per day, hour, month, and year, respectively.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TableTimePartitioning defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Integer expirationMs;
        private @Nullable String field;
        private @Nullable Boolean requirePartitionFilter;
        private String type;

        public Builder() {
    	      // Empty
        }

        public Builder(TableTimePartitioning defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expirationMs = defaults.expirationMs;
    	      this.field = defaults.field;
    	      this.requirePartitionFilter = defaults.requirePartitionFilter;
    	      this.type = defaults.type;
        }

        public Builder expirationMs(@Nullable Integer expirationMs) {
            this.expirationMs = expirationMs;
            return this;
        }
        public Builder field(@Nullable String field) {
            this.field = field;
            return this;
        }
        public Builder requirePartitionFilter(@Nullable Boolean requirePartitionFilter) {
            this.requirePartitionFilter = requirePartitionFilter;
            return this;
        }
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }        public TableTimePartitioning build() {
            return new TableTimePartitioning(expirationMs, field, requirePartitionFilter, type);
        }
    }
}

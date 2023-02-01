// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery.Inputs
{

    public sealed class JobLoadGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept rows that are missing trailing optional columns. The missing values are treated as nulls.
        /// If false, records with missing trailing columns are treated as bad records, and if there are too many bad records,
        /// an invalid error is returned in the job result. The default value is false. Only applicable to CSV, ignored for other formats.
        /// </summary>
        [Input("allowJaggedRows")]
        public Input<bool>? AllowJaggedRows { get; set; }

        /// <summary>
        /// Indicates if BigQuery should allow quoted data sections that contain newline characters in a CSV file.
        /// The default value is false.
        /// </summary>
        [Input("allowQuotedNewlines")]
        public Input<bool>? AllowQuotedNewlines { get; set; }

        /// <summary>
        /// Indicates if we should automatically infer the options and schema for CSV and JSON sources.
        /// </summary>
        [Input("autodetect")]
        public Input<bool>? Autodetect { get; set; }

        /// <summary>
        /// Specifies whether the job is allowed to create new tables. The following values are supported:
        /// CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table.
        /// CREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result.
        /// Creation, truncation and append actions occur as one atomic update upon job completion
        /// Default value is `CREATE_IF_NEEDED`.
        /// Possible values are `CREATE_IF_NEEDED` and `CREATE_NEVER`.
        /// </summary>
        [Input("createDisposition")]
        public Input<string>? CreateDisposition { get; set; }

        /// <summary>
        /// Custom encryption configuration (e.g., Cloud KMS keys)
        /// Structure is documented below.
        /// </summary>
        [Input("destinationEncryptionConfiguration")]
        public Input<Inputs.JobLoadDestinationEncryptionConfigurationGetArgs>? DestinationEncryptionConfiguration { get; set; }

        /// <summary>
        /// The destination table to load the data into.
        /// Structure is documented below.
        /// </summary>
        [Input("destinationTable", required: true)]
        public Input<Inputs.JobLoadDestinationTableGetArgs> DestinationTable { get; set; } = null!;

        /// <summary>
        /// The character encoding of the data. The supported values are UTF-8 or ISO-8859-1.
        /// The default value is UTF-8. BigQuery decodes the data after the raw, binary data
        /// has been split using the values of the quote and fieldDelimiter properties.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// The separator for fields in a CSV file. The separator can be any ISO-8859-1 single-byte character.
        /// To use a character in the range 128-255, you must encode the character as UTF8. BigQuery converts
        /// the string to ISO-8859-1 encoding, and then uses the first byte of the encoded string to split the
        /// data in its raw, binary state. BigQuery also supports the escape sequence "\t" to specify a tab separator.
        /// The default value is a comma (',').
        /// </summary>
        [Input("fieldDelimiter")]
        public Input<string>? FieldDelimiter { get; set; }

        /// <summary>
        /// Indicates if BigQuery should allow extra values that are not represented in the table schema.
        /// If true, the extra values are ignored. If false, records with extra columns are treated as bad records,
        /// and if there are too many bad records, an invalid error is returned in the job result.
        /// The default value is false. The sourceFormat property determines what BigQuery treats as an extra value:
        /// CSV: Trailing columns
        /// JSON: Named values that don't match any column names
        /// </summary>
        [Input("ignoreUnknownValues")]
        public Input<bool>? IgnoreUnknownValues { get; set; }

        /// <summary>
        /// If sourceFormat is set to newline-delimited JSON, indicates whether it should be processed as a JSON variant such as GeoJSON.
        /// For a sourceFormat other than JSON, omit this field. If the sourceFormat is newline-delimited JSON: - for newline-delimited
        /// GeoJSON: set to GEOJSON.
        /// </summary>
        [Input("jsonExtension")]
        public Input<string>? JsonExtension { get; set; }

        /// <summary>
        /// The maximum number of bad records that BigQuery can ignore when running the job. If the number of bad records exceeds this value,
        /// an invalid error is returned in the job result. The default value is 0, which requires that all records are valid.
        /// </summary>
        [Input("maxBadRecords")]
        public Input<int>? MaxBadRecords { get; set; }

        /// <summary>
        /// Specifies a string that represents a null value in a CSV file. For example, if you specify "\N", BigQuery interprets "\N" as a null value
        /// when loading a CSV file. The default value is the empty string. If you set this property to a custom value, BigQuery throws an error if an
        /// empty string is present for all data types except for STRING and BYTE. For STRING and BYTE columns, BigQuery interprets the empty string as
        /// an empty value.
        /// </summary>
        [Input("nullMarker")]
        public Input<string>? NullMarker { get; set; }

        [Input("projectionFields")]
        private InputList<string>? _projectionFields;

        /// <summary>
        /// If sourceFormat is set to "DATASTORE_BACKUP", indicates which entity properties to load into BigQuery from a Cloud Datastore backup.
        /// Property names are case sensitive and must be top-level properties. If no properties are specified, BigQuery loads all properties.
        /// If any named property isn't found in the Cloud Datastore backup, an invalid error is returned in the job result.
        /// </summary>
        public InputList<string> ProjectionFields
        {
            get => _projectionFields ?? (_projectionFields = new InputList<string>());
            set => _projectionFields = value;
        }

        /// <summary>
        /// The value that is used to quote data sections in a CSV file. BigQuery converts the string to ISO-8859-1 encoding,
        /// and then uses the first byte of the encoded string to split the data in its raw, binary state.
        /// The default value is a double-quote ('"'). If your data does not contain quoted sections, set the property value to an empty string.
        /// If your data contains quoted newline characters, you must also set the allowQuotedNewlines property to true.
        /// </summary>
        [Input("quote")]
        public Input<string>? Quote { get; set; }

        [Input("schemaUpdateOptions")]
        private InputList<string>? _schemaUpdateOptions;

        /// <summary>
        /// Allows the schema of the destination table to be updated as a side effect of the load job if a schema is autodetected or
        /// supplied in the job configuration. Schema update options are supported in two cases: when writeDisposition is WRITE_APPEND;
        /// when writeDisposition is WRITE_TRUNCATE and the destination table is a partition of a table, specified by partition decorators.
        /// For normal tables, WRITE_TRUNCATE will always overwrite the schema. One or more of the following values are specified:
        /// ALLOW_FIELD_ADDITION: allow adding a nullable field to the schema.
        /// ALLOW_FIELD_RELAXATION: allow relaxing a required field in the original schema to nullable.
        /// </summary>
        public InputList<string> SchemaUpdateOptions
        {
            get => _schemaUpdateOptions ?? (_schemaUpdateOptions = new InputList<string>());
            set => _schemaUpdateOptions = value;
        }

        /// <summary>
        /// The number of rows at the top of a CSV file that BigQuery will skip when loading the data.
        /// The default value is 0. This property is useful if you have header rows in the file that should be skipped.
        /// When autodetect is on, the behavior is the following:
        /// skipLeadingRows unspecified - Autodetect tries to detect headers in the first row. If they are not detected,
        /// the row is read as data. Otherwise data is read starting from the second row.
        /// skipLeadingRows is 0 - Instructs autodetect that there are no headers and data should be read starting from the first row.
        /// skipLeadingRows = N &gt; 0 - Autodetect skips N-1 rows and tries to detect headers in row N. If headers are not detected,
        /// row N is just skipped. Otherwise row N is used to extract column names for the detected schema.
        /// </summary>
        [Input("skipLeadingRows")]
        public Input<int>? SkipLeadingRows { get; set; }

        /// <summary>
        /// The format of the data files. For CSV files, specify "CSV". For datastore backups, specify "DATASTORE_BACKUP".
        /// For newline-delimited JSON, specify "NEWLINE_DELIMITED_JSON". For Avro, specify "AVRO". For parquet, specify "PARQUET".
        /// For orc, specify "ORC". [Beta] For Bigtable, specify "BIGTABLE".
        /// The default value is CSV.
        /// </summary>
        [Input("sourceFormat")]
        public Input<string>? SourceFormat { get; set; }

        [Input("sourceUris", required: true)]
        private InputList<string>? _sourceUris;

        /// <summary>
        /// The fully-qualified URIs that point to your data in Google Cloud.
        /// For Google Cloud Storage URIs: Each URI can contain one '\*' wildcard character
        /// and it must come after the 'bucket' name. Size limits related to load jobs apply
        /// to external data sources. For Google Cloud Bigtable URIs: Exactly one URI can be
        /// specified and it has be a fully specified and valid HTTPS URL for a Google Cloud Bigtable table.
        /// For Google Cloud Datastore backups: Exactly one URI can be specified. Also, the '\*' wildcard character is not allowed.
        /// </summary>
        public InputList<string> SourceUris
        {
            get => _sourceUris ?? (_sourceUris = new InputList<string>());
            set => _sourceUris = value;
        }

        /// <summary>
        /// Time-based partitioning specification for the destination table.
        /// Structure is documented below.
        /// </summary>
        [Input("timePartitioning")]
        public Input<Inputs.JobLoadTimePartitioningGetArgs>? TimePartitioning { get; set; }

        /// <summary>
        /// Specifies the action that occurs if the destination table already exists. The following values are supported:
        /// WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the table data and uses the schema from the query result.
        /// WRITE_APPEND: If the table already exists, BigQuery appends the data to the table.
        /// WRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result.
        /// Each action is atomic and only occurs if BigQuery is able to complete the job successfully.
        /// Creation, truncation and append actions occur as one atomic update upon job completion.
        /// Default value is `WRITE_EMPTY`.
        /// Possible values are `WRITE_TRUNCATE`, `WRITE_APPEND`, and `WRITE_EMPTY`.
        /// </summary>
        [Input("writeDisposition")]
        public Input<string>? WriteDisposition { get; set; }

        public JobLoadGetArgs()
        {
        }
        public static new JobLoadGetArgs Empty => new JobLoadGetArgs();
    }
}
